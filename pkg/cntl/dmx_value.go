package cntl

import "encoding/json"

// MarshalYAML encodes the value to YAML
func (v *DMXValue) MarshalYAML() (interface{}, error) {
	return v.Value, nil
}

// UnmarshalYAML takes the value from YAML
func (v *DMXValue) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return unmarshal(&v.Value)
}

// MarshalJSON converts the value to a json byte array
func (v *DMXValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Value)
}

// UnmarshalJSON sets the value from a json byte array
func (v *DMXValue) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &v.Value)
}

// Equals returns whether the two given objects are equal
func (v *DMXValue) Equals(v2 *DMXValue) bool {
	return v != nil && v2 != nil && v.Value == v2.Value
}

// Byte returns the byte representation of the value
func (v *DMXValue) Byte() byte {
	return byte(v.Value)
}

// Uint8 returns the original uint8 representation of the value
func (v *DMXValue) Uint8() uint8 {
	return v.Value
}
