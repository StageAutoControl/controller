package dmx

import "errors"

// Render Errors
var (
	ErrDeviceHasDisabledModeChannel   = errors.New("device has disabled Mode channel")
	ErrDeviceHasDisabledStrobeChannel = errors.New("device has disabled Strobe channel")
	ErrDeviceHasDisabledDimmerChannel = errors.New("device has disabled Dimmer channel")
	ErrDeviceIsNotMoving              = errors.New("device is not moving, cannot use tilt and pan")

	ErrDeviceParamsDevicesInvalid          = errors.New("DMXDeviceParams must have either a group or a device")
	ErrDeviceParamsValuesInvalid           = errors.New("DMXDeviceParams must not have more the one of [Animation, Transition, Params]")
	ErrDeviceParamsNoDevices               = errors.New("DMXDeviceParams matches no device")
	ErrDeviceParamsColorVarMustBeExclusive = errors.New("DMXDeviceParams cannot have a $color var and one of [red, green, blue, white]")
	ErrTransitionDeviceParamsMustMatchLED  = errors.New("DMXTransition contains a param set where the LED is not the same")
	ErrDeviceSelectorMustHaveTagsOrID      = errors.New("DMXDeviceSelector must have either tags or an ID")
	ErrDeviceSelectorCannotHaveTagsAndID   = errors.New("DMXDeviceSelector cannot have tags and an ID")
)
