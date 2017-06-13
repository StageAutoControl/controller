package transport

import (
	"fmt"
	"log"
	"net"

	"strings"

	"github.com/StageAutoControl/controller/cntl"
	"github.com/golang/protobuf/proto"
)

// Visualizer is a writer to the visualizer tool
type Visualizer struct {
	endpoint string
	socket   net.Conn
}

// NewVisualizer creates a new Visualizer
func NewVisualizer(endpoint string) (*Visualizer, error) {
	socket, err := net.Dial("tcp", endpoint)
	if err != nil {
		return nil, err
	}

	return &Visualizer{
		socket:   socket,
		endpoint: endpoint,
	}, nil
}

// Write writes to the visualizer stream
func (t *Visualizer) Write(cmd cntl.Command) error {
	tc := Command{
		DmxCommands:  make([]*DMXCommand, len(cmd.DMXCommands)),
		MidiCommands: make([]*MIDICommand, len(cmd.MIDICommands)),
		BarChange:    convertBarChange(cmd.BarChange),
	}
	for i, c := range cmd.DMXCommands {
		tc.DmxCommands[i] = &DMXCommand{
			Universe: uint32(c.Universe),
			Channel:  uint32(c.Channel),
			Value:    uint32(c.Value.Value),
		}
	}

	b, err := proto.Marshal(&tc)
	if err != nil {
		return err
	}

	// append delimiter byte
	b = append(b, 0x0)
	if n, err := t.socket.Write(b); err != nil {
		return err
	} else if n == 0 {
		return fmt.Errorf("Did not sent anything, sent %d bytes.", n)
	}

	go t.debug(tc, b)

	return nil
}

func convertBarChange(bc *cntl.BarChange) *BarChange {
	if bc == nil {
		return nil
	}
	return &BarChange{
		At:        bc.At,
		NoteCount: uint32(bc.NoteCount),
		NoteValue: uint32(bc.NoteValue),
		Speed:     uint32(bc.Speed),
	}
}

func (t *Visualizer) debug(cs Command, b []byte) {
	log.Printf("Sent %d commands to visualizer: %v", len(cs.DmxCommands), renderDMXCommands(cs))
}

func renderDMXCommands(cmds Command) string {
	s := make([]string, len(cmds.DmxCommands))

	for i, c := range cmds.DmxCommands {
		s[i] = fmt.Sprintf("%d:%d -> %d", c.Universe, c.Channel, c.Value)
	}

	return fmt.Sprintf("%s --> %s", renderBarChange(cmds.BarChange), strings.Join(s, " | "))
}

func renderBarChange(bc *BarChange) string {
	if bc == nil {
		return strings.Repeat(" ", 20)
	}

	return fmt.Sprintf("%19s", fmt.Sprintf("#%d %d/%d @%d bpm", bc.At, bc.NoteCount, bc.NoteValue, bc.Speed))
}
