package output

import "math"

func GenerateSin(sampleRate, samples uint, freq float64) []float64 {
	b := make([]float64, samples)
	step := freq / float64(sampleRate)
	var phase float64

	for i := range b {
		b[i] = float64(math.Sin(2 * math.Pi * phase))
		_, phase = math.Modf(phase + step)
	}

	return b
}

