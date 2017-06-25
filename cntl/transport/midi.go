package transport

import (
	"log"

	"github.com/StageAutoControl/controller/cntl"
	"github.com/rakyll/portmidi"
)

// MIDI is a transport that sends MIDI signals using portmidi.
type MIDI struct {
	deviceID portmidi.DeviceID
}

// NewMIDI creates a new MIDI transport
func NewMIDI(deviceID portmidi.DeviceID) (*MIDI, error) {
	portmidi.Initialize()

	if deviceID == 0 {
		deviceID = portmidi.DefaultOutputDeviceID()
	}

	info := portmidi.Info(deviceID)
	if info == nil {
		log.Fatal("Unable to read default output device")
	}

	return &MIDI{deviceID}, nil
}

// Write writes MIDI signals to portmidi
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
