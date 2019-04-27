package artnet

import (
	"sort"
	"sync"
)

// State stores the state of universes
type State struct {
	data UniverseStateMap
	m    sync.RWMutex
}

// NewState returns a new state instance
func NewState() *State {
	return &State{
		data: UniverseStateMap{},
		m:    sync.RWMutex{},
	}
}

// NewStateFromData takes the given data and stores it into a freshly created store
func NewStateFromData(data map[uint16]Universe) *State {
	return NewStateFromUniverseStateMap(UniverseStateMap(data))
}

// NewStateFromUniverseStateMap takes the given data and stores it into a freshly created store
func NewStateFromUniverseStateMap(data map[uint16]Universe) *State {
	s := NewState()
	for k, value := range data {
		s.SetUniverse(k, value)
	}
	return s
}

// SetChannel sets a given channel on a given universe on a given value.
func (s *State) SetChannel(u, c uint16, v uint8) {
	dmx := s.GetUniverse(u)
	dmx[c] = byte(v)
	s.SetUniverse(u, dmx)
}

func (s *State) SetChannelValue(value ChannelValue) {
	s.SetChannel(value.Universe, value.Channel, value.Value)
}

// SetChannelValues sets a range of ChannelValues for convenience
func (s *State) SetChannelValues(values []ChannelValue) {
	for _, value := range values {
		s.SetChannelValue(value)
	}
}

// SetUniverse sets a complete DMX universe data
func (s *State) SetUniverse(u uint16, dmx Universe) {
	s.m.Lock()
	s.data[u] = dmx
	s.m.Unlock()
}

// GetUniverse gets a complete DMX universe data
func (s *State) GetUniverse(u uint16) Universe {
	s.m.RLock()

	dmx, ok := s.data[u]
	if !ok {
		dmx = Universe{}
	}

	s.m.RUnlock()
	return dmx
}

// Get returns all of the current state
func (s *State) Get() UniverseStateMap {
	s.m.RLock()
	c := make(UniverseStateMap)
	for k, v := range s.data {
		c[k] = v
	}
	s.m.RUnlock()
	return c
}

// GetUniverses returns a slice of all available universe indexes
func (s *State) GetUniverses() []uint16 {
	universes := make([]uint16, 0)

	s.m.RLock()

	for u := range s.data {
		universes = append(universes, u)
	}

	sort.Slice(universes, func(i, j int) bool { return universes[i] < universes[j] })
	s.m.RUnlock()

	return universes
}
