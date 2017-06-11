package transport

import (
	"fmt"
	"log"
	"net"

	"strings"

	"github.com/StageAutoControl/controller/cntl"
	"github.com/golang/protobuf/proto"
)

// VisualizerTransport is a writer to the visualizer tool
type VisualizerTransport struct {
	endpoint string
	socket   net.Conn
}

// NewVisualizerTransport creates a new VisualizerTransport
func NewVisualizerTransport(endpoint string) (*VisualizerTransport, error) {
	socket, err := net.Dial("tcp", endpoint)
	if err != nil {
		return nil, err
	}

	return &VisualizerTransport{
		socket:   socket,
		endpoint: endpoint,
	}, nil
}

// Write writes to the visualizer stream
func (t *VisualizerTransport) Write(cmd cntl.Command) error {
	cs := DMXCommands{
		Commands: make([]*DMXCommand, len(cmd.DMXCommands)),
	}
	for i, c := range cmd.DMXCommands {
		cs.Commands[i] = &DMXCommand{
			Universe: uint32(c.Universe),
			Channel:  uint32(c.Channel),
			Value:    uint32(c.Value.Value),
		}
	}

	b, err := proto.Marshal(&cs)
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

	go t.debug(cs, b)

	return nil
}

func (t *VisualizerTransport) debug(cs DMXCommands, b []byte) {
	log.Printf("Sent %d commands to visualizer: %v", len(cs.Commands), renderDMXCommands(cs))
}

func renderDMXCommands(cmds DMXCommands) string {
	s := make([]string, len(cmds.Commands))

	for i, c := range cmds.Commands {
		s[i] = fmt.Sprintf("%d:%d -> %d", c.Universe, c.Channel, c.Value)
	}

	return strings.Join(s, " | ")
}
