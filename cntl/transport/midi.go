package transport

import "github.com/StageAutoControl/controller/cntl"

type MIDI struct {
	device string
}

func NewMIDI(device string) (*MIDI, error) {
	return &MIDI{device}, nil
}

func (m *MIDI) Write(cmd cntl.Command) error {
	if len(cmd.MIDICommands) == 0 {
		return nil
	}

	for _, c := range cmd.MIDICommands {
		if err := m.send(c); err != nil {
			return err
		}
	}

	return nil
}

func (m *MIDI) send(c cntl.MIDICommand) error {
	return nil
}
