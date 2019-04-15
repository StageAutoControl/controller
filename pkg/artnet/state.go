package artnet

// State stores the state of universes
type State struct {
	data UniverseStateMap
}

// NewState returns a new state instance
func NewState() *State {
	return &State{
		data: UniverseStateMap{},
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
func (s *State) SetChannel(u uint16, c, v uint8) {
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
	s.data[u] = dmx
}

// GetUniverse gets a complete DMX universe data
func (s *State) GetUniverse(u uint16) Universe {
	dmx, ok := s.data[u]
	if !ok {
		dmx = Universe{}
	}

	return dmx
}

// Get returns all of the current state
func (s *State) Get() UniverseStateMap {
	return s.data
}

// GetUniverses returns a slice of all available universe indexes
func (s *State) GetUniverses() []uint16 {
	universes := make([]uint16, 0)

	for u := range s.data {
		universes = append(universes, u)
	}

	return universes
}
