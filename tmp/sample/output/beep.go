package output

import (
	"fmt"
	"github.com/cryptix/wav"
)

func Beep(meta wav.File, c chan int32, highPass int32) {
	lastSample := int32(0)
	beeped := false
	lastBeepSample := 0
	lastZeroBeepSample := 0
	i := 0

	for s := range c {
		if s > lastSample {
			lastSample = s
		} else {
			if !beeped && s > highPass {
				fmt.Printf("BEEP %v %v \n", i, s)
				lastBeepSample = i
				beeped = true
			}

			if s < 100 && s > -100 && (lastBeepSample+4000 < i) {
				if lastZeroBeepSample+10 < i {
					beeped = false
				}

				lastZeroBeepSample = i
			}
		}

		i++
	}
}
