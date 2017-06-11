package cntl

import "encoding/json"

// MarshalYAML encodes the value to YAML
func (v *DMXValue) MarshalYAML() (interface{}, error) {
	return v.Value, nil
}

// UnmarshalYAML takes the value from YAML
func (v *DMXValue) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := unmarshal(&v.Value); err != nil {
		return err
	}

	return nil
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
