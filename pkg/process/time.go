package process

import (
	"fmt"
	"time"
)

// JSONTime handles parsing and formatting timestamps according the ISO8601 standard
type JSONTime struct {
	time.Time
}

// String returns a string representation of the time.
func (t JSONTime) String() string {
	return t.Format(time.RFC3339)
}

// MarshalJSON formats the timestamp as JSON
func (t JSONTime) MarshalJSON() ([]byte, error) {
	date := fmt.Sprintf("%q", t.String())
	return []byte(date), nil
}
