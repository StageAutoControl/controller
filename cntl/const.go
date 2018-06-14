package cntl

// easing function names
const (
	EaseLinear        EaseFunc = "Linear"
	EaseQuadIn        EaseFunc = "InQuad"
	EaseQuadOut       EaseFunc = "OutQuad"
	EaseQuadInOut     EaseFunc = "InOutQuad"
	EaseCubicIn       EaseFunc = "InCubic"
	EaseCubicOut      EaseFunc = "OutCubic"
	EaseCubicInOut    EaseFunc = "InOutCubic"
	EaseQuartIn       EaseFunc = "InQuart"
	EaseQuartOut      EaseFunc = "OutQuart"
	EaseQuartInOut    EaseFunc = "InOutQuart"
	EaseQuintIn       EaseFunc = "InQuint"
	EaseQuintOut      EaseFunc = "OutQuint"
	EaseQuintInOut    EaseFunc = "InOutQuint"
	EaseSineIn        EaseFunc = "InSine"
	EaseSineOut       EaseFunc = "OutSine"
	EaseSineInOut     EaseFunc = "InOutSine"
	EaseCircularIn    EaseFunc = "InCircular"
	EaseCircularOut   EaseFunc = "OutCircular"
	EaseCircularInOut EaseFunc = "InOutCircular"
	EaseExpoIn        EaseFunc = "InExpo"
	EaseExpoOut       EaseFunc = "OutExpo"
	EaseExpoInOut     EaseFunc = "InOutExpo"
	EaseElasticIn     EaseFunc = "InElastic"
	EaseElasticOut    EaseFunc = "OutElastic"
	EaseElasticInOut  EaseFunc = "InOutElastic"
	EaseBackIn        EaseFunc = "InBack"
	EaseBackOut       EaseFunc = "OutBack"
	EaseBackInOut     EaseFunc = "InOutBack"
	EaseBounceIn      EaseFunc = "InBounce"
	EaseBounceOut     EaseFunc = "OutBounce"
	EaseBounceInOut   EaseFunc = "InOutBounce"
)

// RenderFrames defines the smallest render unit of a bar. Has to be multiplier of 4.
const RenderFrames uint8 = 64

// Logger fields
const (
	LoggerFieldTransport = "transport"
	LoggerFieldWaiter = "waiter"
)
