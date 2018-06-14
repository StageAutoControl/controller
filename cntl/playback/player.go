package playback

import (
	"context"
	"time"

	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/StageAutoControl/controller/cntl"
	"github.com/StageAutoControl/controller/cntl/song"
)

// Player plays various things from a given data store, for example songs or a whole SetList.
type Player struct {
	logger    *logrus.Entry
	dataStore *cntl.DataStore
	writers   []TransportWriter
	waiters   []Waiter
}

// NewPlayer returns a new Player instance
func NewPlayer(logger *logrus.Entry, ds *cntl.DataStore, writers []TransportWriter, waiters []Waiter) *Player {
	return &Player{logger, ds, writers, waiters}
}

func (p *Player) checkSetList(setList *cntl.SetList) error {
	for _, songSel := range setList.Songs {
		if _, ok := p.dataStore.Songs[songSel.ID]; !ok {
			return fmt.Errorf("cannot find Song %q", songSel.ID)
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

	for _, songSel := range setList.Songs {
		select {
		case <-ctx.Done():
			p.logger.Warn("Aborting")
			return nil
		default:
		}

		p.logger.Infof("Playing song %s", songSel.ID)

		if err := p.PlaySong(ctx, songSel.ID); err != nil {
			return err
		}
	}

	return nil
}

func (p *Player) wait(ctx context.Context) error {
	done := make(chan struct{}, len(p.waiters))
	cancel := make(chan struct{}, len(p.waiters))
	err := make(chan error, len(p.waiters))

	defer func() {
		cancel <- struct{}{}
	}()

	for _, w := range p.waiters {
		go w.Wait(done, cancel, err)
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

	var i int
	var cmd cntl.Command
	for {
		select {
		case <-ctx.Done():
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
				go w.Write(cmd)
			}

			i++
		}
	}
}

// CalcRenderSpeed calculates the render speed of a BarChange to a time.Duration
func CalcRenderSpeed(bc *cntl.BarChange) time.Duration {
	return time.Minute / time.Duration(bc.Speed*uint16(bc.NoteValue)/4) / time.Duration(cntl.RenderFrames/bc.NoteValue)
}
