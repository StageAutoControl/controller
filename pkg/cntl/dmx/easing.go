package dmx

import (
	"fmt"

	"github.com/StageAutoControl/controller/pkg/cntl"

	"github.com/creasty/go-easing"
)

type easingFunc easing.EaseFunc

// getEasingFunc returns the easingFunc for a given cntl.EaseFunc function name
func getEasingFunc(easingFunc cntl.EaseFunc) (easingFunc, error) {
	switch easingFunc {
	case cntl.EaseLinear:
		return easing.Linear, nil
	case cntl.EaseQuadIn:
		return easing.QuadEaseIn, nil
	case cntl.EaseQuadOut:
		return easing.QuadEaseOut, nil
	case cntl.EaseQuadInOut:
		return easing.QuadEaseInOut, nil
	case cntl.EaseCubicIn:
		return easing.CubicEaseIn, nil
	case cntl.EaseCubicOut:
		return easing.CubicEaseOut, nil
	case cntl.EaseCubicInOut:
		return easing.CubicEaseInOut, nil
	case cntl.EaseQuartIn:
		return easing.QuartEaseIn, nil
	case cntl.EaseQuartOut:
		return easing.QuartEaseOut, nil
	case cntl.EaseQuartInOut:
		return easing.QuartEaseInOut, nil
	case cntl.EaseQuintIn:
		return easing.QuintEaseIn, nil
	case cntl.EaseQuintOut:
		return easing.QuintEaseOut, nil
	case cntl.EaseQuintInOut:
		return easing.QuintEaseInOut, nil
	case cntl.EaseSineIn:
		return easing.SineEaseIn, nil
	case cntl.EaseSineOut:
		return easing.SineEaseOut, nil
	case cntl.EaseSineInOut:
		return easing.SineEaseInOut, nil
	case cntl.EaseCircularIn:
		return easing.CircularEaseIn, nil
	case cntl.EaseCircularOut:
		return easing.CircularEaseOut, nil
	case cntl.EaseCircularInOut:
		return easing.CircularEaseInOut, nil
	case cntl.EaseExpoIn:
		return easing.ExpoEaseIn, nil
	case cntl.EaseExpoOut:
		return easing.ExpoEaseOut, nil
	case cntl.EaseExpoInOut:
		return easing.ExpoEaseInOut, nil
	case cntl.EaseElasticIn:
		return easing.ElasticEaseIn, nil
	case cntl.EaseElasticOut:
		return easing.ElasticEaseOut, nil
	case cntl.EaseElasticInOut:
		return easing.ElasticEaseInOut, nil
	case cntl.EaseBackIn:
		return easing.BackEaseIn, nil
	case cntl.EaseBackOut:
		return easing.BackEaseOut, nil
	case cntl.EaseBackInOut:
		return easing.BackEaseInOut, nil
	case cntl.EaseBounceIn:
		return easing.BounceEaseIn, nil
	case cntl.EaseBounceOut:
		return easing.BounceEaseOut, nil
	case cntl.EaseBounceInOut:
		return easing.BounceEaseInOut, nil
	default:
		return nil, fmt.Errorf("unable to find easing function %v", easingFunc)
	}
}
