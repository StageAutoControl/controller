package song

import (
	"time"

	"github.com/StageAutoControl/controller/cntl"
)

// OutputWriter is a writer to an output stream, for example a websocket or Stdout.
type TransportWriter interface {
	Write(cntl.Command) error
}

// Player plays various things from a given data store, for example songs or whole set lists.
type Player struct {
	ds *cntl.DataStore
	w  TransportWriter
}

// NewPlayer returns a new Player instance
func NewPlayer(ds *cntl.DataStore, w TransportWriter) *Player {
	return &Player{ds, w}
}

// Play plays a given length of a song
func (p *Player) Play(songID string, bars uint16) {}

// PlayAll plays a whole song
func (p *Player) PlayAll(songID string) error {
	cmds, err := Render(p.ds, songID)
	if err != nil {
		return err
	}

	t := time.NewTicker(10 * time.Nanosecond)
	l := len(cmds)

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

			go p.w.Write(cmd)

			i++
		}
	}

	return nil
}

func CalcRenderSpeed(bc *cntl.BarChange) time.Duration {
	return time.Minute / time.Duration(bc.Speed*uint16(bc.NoteValue)/4) / time.Duration(cntl.RenderFrames/bc.NoteValue)
}
