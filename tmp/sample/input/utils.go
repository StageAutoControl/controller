package input

import (
	"fmt"
	"math"
)

func toInt(s []byte, bytesPerSample uint32) (n int32) {
	switch bytesPerSample {
	case 1:
		n = int32(s[0])
	case 2:
		n = int32(s[0]) + int32(s[1])<<8
	case 3:
		n = int32(s[0]) + int32(s[1])<<8 + int32(s[2])<<16
	case 4:
		n = int32(s[0]) + int32(s[1])<<8 + int32(s[2])<<16 + int32(s[3])<<24
	default:
		n = 0
		panic(fmt.Errorf("Unhandled bytesPerSample! b:%d", bytesPerSample))
	}

	return
}

func checkNegative(i int32, bytesPerSample uint32) int32 {
	maxVal := int32(math.Pow(2, float64(int(bytesPerSample)*8)))

	if i > maxVal/2-1 {
		return i - maxVal
	}

	return i
}
