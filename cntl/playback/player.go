package playback

import (
	"time"

	"fmt"

	"log"

	"github.com/StageAutoControl/controller/cntl"
	"github.com/StageAutoControl/controller/cntl/song"
)

// TransportWriter is a writer to an output stream, for example a websocket or Stdout.
type TransportWriter interface {
	Write(cntl.Command) error
}

// Waiter waits for a trigger to happen
type Waiter interface {
	Wait(done chan struct{}, cancel chan struct{}, err chan error) error
}

// Player plays various things from a given data store, for example songs or a whole SetList.
type Player struct {
	dataStore *cntl.DataStore
	writers   []TransportWriter
	waiters   []Waiter
}

// NewPlayer returns a new Player instance
func NewPlayer(ds *cntl.DataStore, writers []TransportWriter, waiters []Waiter) *Player {
	return &Player{ds, writers, waiters}
}

func (p *Player) checkSetList(setList *cntl.SetList) error {
	for _, songSel := range setList.Songs {
		if _, ok := p.dataStore.Songs[songSel.ID]; !ok {
			return fmt.Errorf("Cannot find Song %q", songSel.ID)
		}
	}

	return nil
}

// PlaySetList plays a full SetList
func (p *Player) PlaySetList(setListId string) error {
	setList, ok := p.dataStore.SetLists[setListId]
	if !ok {
		return fmt.Errorf("Cannot find SetList %q", setListId)
	}

	if err := p.checkSetList(setList); err != nil {
		return err
	}

	for _, songSel := range setList.Songs {
		log.Printf("Playing song %s \n", songSel.ID)

		if err := p.PlaySong(songSel.ID); err != nil {
			return err
		}
	}

	return nil
}

func (p *Player) wait() error {
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
	case <-done:
		return nil
	case err := <-err:
		return err
	}
}

// PlaySong plays a full song
func (p *Player) PlaySong(songID string) error {
	cmds, err := song.Render(p.dataStore, songID)
	if err != nil {
		return err
	}

	log.Printf("Waiting for waiters before playing song %s \n", songID)
	if err := p.wait(); err != nil {
		return err
	}

	l := len(cmds)
	t := time.NewTicker(1 * time.Nanosecond)

	var i int
	var cmd cntl.Command
	for {
		select {
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

	return nil
}

// CalcRenderSpeed calculates the render speed of a BarChange to a time.Duration
func CalcRenderSpeed(bc *cntl.BarChange) time.Duration {
	return time.Minute / time.Duration(bc.Speed*uint16(bc.NoteValue)/4) / time.Duration(cntl.RenderFrames/bc.NoteValue)
}
