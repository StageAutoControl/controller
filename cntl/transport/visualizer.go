package transport

import (
	"fmt"
	"log"
	"net"

	"github.com/StageAutoControl/controller/cntl"
	"github.com/golang/protobuf/proto"
)

type VisualizerTransport struct {
	endpoint string
	socket   net.Conn
}

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
	log.Printf("Sent %d commands to visualizer: %v", len(cs.Commands), b)
}
