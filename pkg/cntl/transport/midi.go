package transport

import (
	"errors"

	"github.com/StageAutoControl/controller/pkg/cntl"
	"github.com/StageAutoControl/controller/pkg/internal/logging"

	"github.com/rakyll/portmidi"
)

// MIDI is a transport that sends MIDI signals using portmidi.
type MIDI struct {
	logger     logging.Logger
	deviceInfo *portmidi.DeviceInfo
	deviceID   portmidi.DeviceID
	out        *portmidi.Stream
}

// NewMIDI creates a new MIDI transport
func NewMIDI(logger logging.Logger, deviceID int8) (*MIDI, error) {
	if err := portmidi.Initialize(); err != nil {
		return nil, err
	}

	var d portmidi.DeviceID
	if deviceID < 0 {
		d = portmidi.DefaultOutputDeviceID()
	} else {
		d = portmidi.DeviceID(deviceID)
	}

	info := portmidi.Info(d)
	if info == nil {
		return nil, errors.New("unable to read default output device")
	}

	out, err := portmidi.NewOutputStream(d, 10, 0)
	if err != nil {
		return nil, err
	}

	logger.Infof("Using midi device %d", d)

	return &MIDI{logger, info, d, out}, nil
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
	m.logger.Debugf("%#v", events)
}
