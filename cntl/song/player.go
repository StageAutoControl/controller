package song

import (
	"fmt"

	"github.com/StageAutoControl/controller/cntl"
)

// OutputWriter is a writer to an output stream, for example a websocket or Stdout.
type OutputWriter interface {
	Write(cntl.Command)
}

// Player plays various things from a given data store, for example songs or whole set lists.
type Player struct {
	ds *cntl.DataStore
	w  OutputWriter
}

// NewPlayer returns a new Player instance
func NewPlayer(ds *cntl.DataStore, w OutputWriter) *Player {
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

	fmt.Println(len(cmds))

	for _, cmd := range cmds {
		p.w.Write(cmd)
	}

	return nil
}
