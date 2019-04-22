package playback

import (
	"context"
	"fmt"
	"time"

	"github.com/StageAutoControl/controller/pkg/cntl"
	"github.com/StageAutoControl/controller/pkg/cntl/song"
	"github.com/StageAutoControl/controller/pkg/internal/logging"
)

// Player plays various things from a given data store, for example songs or a whole SetList.
type Player struct {
	logger    logging.Logger
	dataStore *cntl.DataStore
	writers   []TransportWriter
	waiters   []Waiter
}

// NewPlayer returns a new Player instance
func NewPlayer(logger logging.Logger, ds *cntl.DataStore, writers []TransportWriter, waiters []Waiter) *Player {
	return &Player{
		logger:    logger,
		dataStore: ds,
		writers:   writers,
		waiters:   waiters,
	}
}

func (p *Player) checkSetList(setList *cntl.SetList) error {
	for _, songID := range setList.Songs {
		if _, ok := p.dataStore.Songs[songID]; !ok {
			return fmt.Errorf("cannot find Process %q", songID)
		}
	}

	return nil
}

// PlaySetList plays a full SetList
func (p *Player) PlaySetList(ctx context.Context, setListID string) error {
	setList, ok := p.dataStore.SetLists[setListID]
	if !ok {
		return fmt.Errorf("cannot find SetList %q", setListID)
	}

	if err := p.checkSetList(setList); err != nil {
		return err
	}

	for _, songID := range setList.Songs {
		select {
		case <-ctx.Done():
			p.logger.Warn("Aborting")
			return nil
		default:
		}

		s, ok := p.dataStore.Songs[songID]
		if !ok {
			return fmt.Errorf("failed to find song %v", songID)
		}

		p.logger.Infof("Playing song %v", s.Name)

		if err := p.PlaySong(ctx, songID); err != nil {
			return err
		}
	}

	return nil
}

func (p *Player) wait(ctx context.Context) error {
	if len(p.waiters) == 0 {
		return nil
	}

	chanLen := len(p.waiters) + 1
	done := make(chan struct{}, chanLen)
	cancel := make(chan struct{}, chanLen)

	defer func() {
		cancel <- struct{}{}
	}()

	for _, w := range p.waiters {
		go func() {
			if err := w.Wait(done, cancel); err != nil {
				p.logger.Error(err)
			}
		}()
	}

	select {
	case <-ctx.Done():
		return ErrCancelled
	case <-done:
		return nil
	}
}

// PlaySong plays a full song
func (p *Player) PlaySong(ctx context.Context, songID string) error {
	commands, err := song.Render(p.dataStore, songID)
	if err != nil {
		return err
	}

	s, ok := p.dataStore.Songs[songID]
	if !ok {
		return fmt.Errorf("failed to find song %v", songID)
	}

	p.logger.Infof("Playing song %v", s.Name)

	p.logger.Infof("Waiting for waiters before playing song %v", s.Name)
	if err := p.wait(ctx); err != nil {
		return err
	}

	p.logger.Infof("Playing song %v", s.Name)
	return Play(ctx, p.logger, p.writers, commands)
}

// CalcRenderSpeed calculates the render speed of a BarChange to a time.Duration
func CalcRenderSpeed(bc *cntl.BarChange) time.Duration {
	return time.Minute / time.Duration(bc.Speed*uint16(bc.NoteValue)/4) / time.Duration(cntl.RenderFrames/bc.NoteValue)
}

// Play plays a given slice of commands and send it to the given writers
func Play(ctx context.Context, logger logging.Logger, writers []TransportWriter, commands []cntl.Command) error {
	l := len(commands)
	t := time.NewTicker(1 * time.Nanosecond)
	done := ctx.Done()

	var i int
	var cmd cntl.Command
	for {
		select {
		case <-done:
			return ErrCancelled

		case <-t.C:
			if i >= l {
				t.Stop()
				return nil
			}

			cmd = commands[i]
			if cmd.BarChange != nil {
				t.Stop()
				t = time.NewTicker(CalcRenderSpeed(cmd.BarChange))
			}

			for _, w := range writers {
				go func() {
					if err := w.Write(cmd); err != nil {
						logger.Error(err)
					}
				}()
			}

			i++
		}
	}
}

// ToPlayable takes a slice of DMXCommands and combines it with the given BarParams to a playable slice of Commands
func ToPlayable(bp cntl.BarParams, dmxCommands []cntl.DMXCommands) []cntl.Command {
	commands := make([]cntl.Command, len(dmxCommands))
	for i, cmd := range dmxCommands {
		commands[i] = cntl.Command{

			DMXCommands:  cmd,
			MIDICommands: []cntl.MIDICommand{},
		}
	}

	commands[0].BarChange = &cntl.BarChange{
		At:        0,
		BarParams: bp,
	}

	return commands
}
