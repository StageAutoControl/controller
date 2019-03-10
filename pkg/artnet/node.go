package artnet

import (
	"fmt"
	"strings"

	"github.com/jsimonetti/go-artnet"
)

// NodeToString returns a string representation of the given Node
func NodeToString(n *artnet.ControlledNode) string {
	var inputs, outputs []string

	for _, p := range n.Node.InputPorts {
		inputs = append(inputs, fmt.Sprintf("%s: %s", p.Address.String(), p.Type.String()))
	}

	for _, p := range n.Node.OutputPorts {
		outputs = append(outputs, fmt.Sprintf("%s: %s", p.Address.String(), p.Type.String()))
	}

	return fmt.Sprintf(
		" | IP=%s name=%q type=%q manufacturer=%q desc=%q inputs=%q outputs=%q",
		n.UDPAddress.String(), n.Node.Name, n.Node.Type,
		n.Node.Manufacturer, n.Node.Description,
		strings.Join(inputs, "; "), strings.Join(outputs, "; "),
	)
}
