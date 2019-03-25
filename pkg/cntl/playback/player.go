package playback

import (
	"context"
	"time"

	"github.com/StageAutoControl/controller/pkg/cntl"
	"github.com/StageAutoControl/controller/pkg/internal/logging"

	"fmt"

	"github.com/StageAutoControl/controller/pkg/cntl/song"
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

		p.logger.Infof("Playing song %s", songID)

		if err := p.PlaySong(ctx, songID); err != nil {
			return err
		}
	}

	return nil
}

func (p *Player) wait(ctx context.Context) error {
	chanLen := len(p.waiters) + 1
	done := make(chan struct{}, chanLen)
	cancel := make(chan struct{}, chanLen)
	err := make(chan error, chanLen)

	defer func() {
		cancel <- struct{}{}
	}()

	for _, w := range p.waiters {
		go func() {
			if err := w.Wait(done, cancel, err); err != nil {
				p.logger.Error(err)
			}
		}()
	}

	select {
	case <-ctx.Done():
		return ErrCancelled
	case <-done:
		return nil
	case err := <-err:
		return err
	}
}

// PlaySong plays a full song
func (p *Player) PlaySong(ctx context.Context, songID string) error {
	cmds, err := song.Render(p.dataStore, songID)
	if err != nil {
		return err
	}

	p.logger.Infof("Waiting for waiters before playing song %s", songID)
	if err := p.wait(ctx); err != nil {
		return err
	}

	l := len(cmds)
	t := time.NewTicker(1 * time.Nanosecond)

	p.logger.Infof("Playing song %s", songID)
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

			cmd = cmds[i]
			if cmd.BarChange != nil {
				t.Stop()
				t = time.NewTicker(CalcRenderSpeed(cmd.BarChange))
			}

			for _, w := range p.writers {
				go func() {
					if err := w.Write(cmd); err != nil {
						p.logger.Error(err)
					}
				}()
			}

			i++
		}
	}
}

// CalcRenderSpeed calculates the render speed of a BarChange to a time.Duration
func CalcRenderSpeed(bc *cntl.BarChange) time.Duration {
	return time.Minute / time.Duration(bc.Speed*uint16(bc.NoteValue)/4) / time.Duration(cntl.RenderFrames/bc.NoteValue)
}
