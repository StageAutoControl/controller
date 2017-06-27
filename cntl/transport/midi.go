package transport

import (
	"fmt"
	"log"
	"strconv"

	"github.com/StageAutoControl/controller/cntl"
	"github.com/rakyll/portmidi"
)

// MIDI is a transport that sends MIDI signals using portmidi.
type MIDI struct {
	deviceInfo *portmidi.DeviceInfo
	deviceID   portmidi.DeviceID
	out        *portmidi.Stream
}

// NewMIDI creates a new MIDI transport
func NewMIDI(deviceID string) (*MIDI, error) {
	if err := portmidi.Initialize(); err != nil {
		return nil, err
	}

	var d portmidi.DeviceID
	if deviceID == "" {
		d = portmidi.DefaultOutputDeviceID()
	} else {
		i, err := strconv.Atoi(deviceID)
		if err != nil {
			return nil, fmt.Errorf("Failed to transform deviceID %q to int: %v", deviceID, err)
		}
		d = portmidi.DeviceID(i)
	}

	info := portmidi.Info(d)
	if info == nil {
		log.Fatal("Unable to read default output device")
	}

	out, err := portmidi.NewOutputStream(d, 10, 0)
	if err != nil {
		return nil, err
	}

	log.Printf("Using midi device %d \n", d)

	return &MIDI{info, d, out}, nil
}

// Write writes MIDI signals to portmidi
func (m *MIDI) Write(cmd cntl.Command) error {
	if len(cmd.MIDICommands) == 0 {
		return nil
	}

	events := m.convertEvents(cmd)
	if err := m.out.Write(events); err != nil {
		return err
	}

	go m.log(events)

	return nil
}

func (m *MIDI) convertEvents(cmd cntl.Command) (events []portmidi.Event) {
	for _, c := range cmd.MIDICommands {
		events = append(events, portmidi.Event{
			Status: int64(c.Status),
			Data1:  int64(c.Data1),
			Data2:  int64(c.Data2),
		})
	}
	return
}

func (m *MIDI) log(events []portmidi.Event) {
	log.Printf("%#v\n", events)
}
