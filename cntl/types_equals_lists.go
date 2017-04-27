package cntl

type songSelectorList []SongSelector

func (v1 songSelectorList) Equals(v2 songSelectorList) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i := range v1 {
		if !v1[i].Equals(v2[i]) {
			return false
		}
	}

	return true
}

type barChangeList []BarChange

func (v1 barChangeList) Equals(v2 barChangeList) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i := range v1 {
		if !v1[i].Equals(v2[i]) {
			return false
		}
	}

	return true
}

type scenePositionList []ScenePosition

func (v1 scenePositionList) Equals(v2 scenePositionList) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i := range v1 {
		if !v1[i].Equals(v2[i]) {
			return false
		}
	}

	return true
}

type dmxDeviceParamsList []DMXDeviceParams

func (v1 dmxDeviceParamsList) Equals(v2 dmxDeviceParamsList) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i := range v1 {
		if !v1[i].Equals(v2[i]) {
			return false
		}
	}

	return true
}

type tagList []Tag

func (v1 tagList) Equals(v2 tagList) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i := range v1 {
		if !v1[i].Equals(v2[i]) {
			return false
		}
	}

	return true
}

type ledList []LED

func (v1 ledList) Equals(v2 ledList) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i := range v1 {
		if !v1[i].Equals(v2[i]) {
			return false
		}
	}

	return true
}

type dmxDeviceSelectorList []DMXDeviceSelector

func (v1 dmxDeviceSelectorList) Equals(v2 dmxDeviceSelectorList) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i := range v1 {
		if !v1[i].Equals(v2[i]) {
			return false
		}
	}

	return true
}

type dmxSubSceneList []DMXSubScene

func (v1 dmxSubSceneList) Equals(v2 dmxSubSceneList) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i := range v1 {
		if !v1[i].Equals(v2[i]) {
			return false
		}
	}

	return true
}

type atList []uint8

func (v1 atList) Equals(v2 atList) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i := range v1 {
		if v1[i] != v2[i] {
			return false
		}
	}

	return true
}

type dmxAnimationFrameList []DMXAnimationFrame

func (v1 dmxAnimationFrameList) Equals(v2 dmxAnimationFrameList) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i := range v1 {
		if !v1[i].Equals(v2[i]) {
			return false
		}
	}

	return true
}
