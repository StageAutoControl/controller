package dmx

import (
	"fmt"
	"math"

	"github.com/StageAutoControl/controller/cntl"
)

// RenderTransition renders the given DMXTransition to an array of DMXCommands to be sent to a DMX device
func RenderTransition(ds *cntl.DataStore, dd []*cntl.DMXDevice, t *cntl.DMXTransition) ([]cntl.DMXCommands, error) {
	cmds := make([]cntl.DMXCommands, t.Length)

	for i, p := range t.Params {
		if p.From.LED != p.To.LED {
			return []cntl.DMXCommands{}, ErrTransitionDeviceParamsMustMatchLED
		}

		paramCMDs, err := RenderTransitionParams(ds, dd, t, p)
		if err != nil {
			return []cntl.DMXCommands{}, fmt.Errorf("failed to render animation transition %q param %d: %v", t.ID, i, err)
		}

		cmds = Merge(cmds, paramCMDs)
	}

	return cmds, nil
}

func RenderTransitionParams(ds *cntl.DataStore, dd []*cntl.DMXDevice, t *cntl.DMXTransition, p cntl.DMXTransitionParams) ([]cntl.DMXCommands, error) {
	result := make([]cntl.DMXCommands, t.Length)
	ease, err := getEasingFunc(t.Ease)
	if err != nil {
		return []cntl.DMXCommands{}, err
	}

	if p.From.Blue != nil && p.To.Blue != nil && p.From.Blue.Value != p.To.Blue.Value {
		steps, err := calcTransitionSteps(p.From.Blue.Value, p.To.Blue.Value, t.Length, ease)
		if err != nil {
			return []cntl.DMXCommands{}, err
		}

		for i, step := range steps {
			//stepParam := copyParams(p.From)
			stepParam := p.From
			stepParam.Blue = &cntl.DMXValue{Value: step}

			cmd, err := RenderParams(ds, dd, stepParam)
			if err != nil {
				return []cntl.DMXCommands{}, err
			}

			result[i] = append(result[i], cmd...)
		}
	}

	return result, nil
}


func calcTransitionSteps(from, to, steps uint8, easingFunc easingFunc) ([]uint8, error) {
	result := make([]uint8, steps)
	diff := float64(to - from)
	floatFrom := float64(from)

	// we assume that the transition is done using n steps, but don't want 8 steps but 7 to have the transition
	// completed at the 8th step. e.g. having 8 steps we need 7 steps to have the 8th be 100%
	step := 1 / float64(steps-1)

	for i := float64(0); i < float64(steps); i++ {
		value := uint8(math.Floor(floatFrom + diff*easingFunc(i*step)))
		fmt.Printf("i=%v steps=%v floatFrom=%v diff=%v value=%v step=%v easeValue=%v\n", i, steps, floatFrom, diff, value, i*step, easingFunc(i*step))
		result[int(i)] = value
	}

	return result, nil
}
