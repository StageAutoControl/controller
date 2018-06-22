package transport

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/StageAutoControl/controller/cntl"
)

// Visualizer is a writer to the visualizer tool
type Visualizer struct {
	logger   *logrus.Entry
	endpoint string
	socket   net.Conn
}

// NewVisualizer creates a new Visualizer
func NewVisualizer(logger *logrus.Entry, endpoint string) (*Visualizer, error) {
	socket, err := net.Dial("tcp", endpoint)
	if err != nil {
		return nil, err
	}

	return &Visualizer{
		logger:   logger,
		socket:   socket,
		endpoint: endpoint,
	}, nil
}

// Write writes to the visualizer stream
func (t *Visualizer) Write(cmd cntl.Command) error {
	b, err := json.Marshal(&cmd)
	if err != nil {
		return err
	}

	// append delimiter byte
	b = append(b, 0x0)
	if n, err := t.socket.Write(b); err != nil {
		return err
	} else if n == 0 {
		return fmt.Errorf("did not sent anything, sent %d bytes", n)
	}

	go t.debug(cmd, b)

	return nil
}

func (t *Visualizer) debug(cs cntl.Command, b []byte) {
	t.logger.Infof("Sent %d commands to visualizer: %v", len(cs.DMXCommands), renderDMXCommands(cs))
}

func renderDMXCommands(cmds cntl.Command) string {
	s := make([]string, len(cmds.DMXCommands))

	for i, c := range cmds.DMXCommands {
		s[i] = fmt.Sprintf("%d:%d -> %d", c.Universe, c.Channel, c.Value)
	}

	return fmt.Sprintf("%s --> %s", renderBarChange(cmds.BarChange), strings.Join(s, " | "))
}

func renderBarChange(bc *cntl.BarChange) string {
	if bc == nil {
		return strings.Repeat(" ", 20)
	}

	return fmt.Sprintf("%19s", fmt.Sprintf("#%d %d/%d @%d bpm", bc.At, bc.NoteCount, bc.NoteValue, bc.Speed))
}
