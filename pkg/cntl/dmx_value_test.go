package cntl

import "testing"

func TestDMXValue_Byte(t *testing.T) {
	tests := []struct {
		v uint8
		b byte
	}{
		{255, 0xff},
		{0, 0x00},
	}

	for _, test := range tests {
		v := DMXValue{test.v}
		b := v.Byte()

		if test.b != b {
			t.Errorf("Expected to get %v, got %v", test.b, b)
		}
	}
}
