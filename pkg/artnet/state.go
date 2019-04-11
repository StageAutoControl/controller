package artnet

import (
	"sync"
)

// State stores the state of universes
type State struct {
	data sync.Map
}

// NewState returns a new state instance
func NewState() *State {
	return &State{
		data: sync.Map{},
	}
}

// NewStateFromData takes the given data and stores it into a freshly created store
func NewStateFromData(data map[uint16]Universe) *State {
	s := NewState()
	for k, value := range data {
		s.SetUniverse(k, value)
	}
	return s
}

// SetChannel sets a given channel on a given universe on a given value.
func (s *State) SetChannel(u uint16, c, v uint8) {
	dmx := s.GetUniverse(u)
	dmx[c] = byte(v)
	s.SetUniverse(u, dmx)
}

// SetUniverse sets a complete DMX universe data
func (s *State) SetUniverse(u uint16, dmx Universe) {
	s.data.Store(u, dmx)
}

// GetUniverse gets a complete DMX universe data
func (s *State) GetUniverse(u uint16) Universe {
	dmx, ok := s.data.Load(u)
	if !ok {
		return Universe{}
	}

	return dmx.(Universe)
}

// GetUniverses returns a slice of all available universe indexes
func (s *State) GetUniverses() []uint16 {
	universes := make([]uint16, 0)

	s.data.Range(func(key interface{}, value interface{}) bool {
		universes = append(universes, key.(uint16))
		return true
	})

	return universes
}
