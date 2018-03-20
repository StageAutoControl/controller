package artnet

// State stores the state of universes
type State map[uint16][512]byte

// NewState returns a new state instance
func NewState() State {
	return make(State)
}

// Set sets a given channel on a given universe on a given value.
func (state State) Set(u uint16, c, v uint8) {
	state.addUniverse(u)

	dmx := state[u]
	dmx[c] = byte(v)
	state[u] = dmx
}

func (state State) addUniverse(u uint16) {
	if _, ok := state[u]; ok {
		return
	}

	state[u] = [512]byte{}
}

// Equals compares two states for equality
func (state State) Equals(state2 State) bool {
	if len(state) != len(state2) {
		return false
	}

	for u := range state {
		if _, ok := state2[u]; !ok {
			return false
		}

		if state2[u] != state[u] {
			return false
		}
	}

	return true
}
