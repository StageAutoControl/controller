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

func datesAreEqual(t1 *JSONTime, t2 *JSONTime) bool {
	if (t1 == nil && t2 != nil) || (t1 != nil && t2 == nil) {
		return false
	}

	if t1 == nil && t2 == nil {
		return true
	}

	return (*t1).Equal(t2.Time)
}
