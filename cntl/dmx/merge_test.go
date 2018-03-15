package dmx

import (
	"testing"

	"github.com/StageAutoControl/controller/cntl"
	"github.com/StageAutoControl/controller/fixtures"
)

func TestMergeAtOffset(t *testing.T) {
	cmds := []cntl.DMXCommands{
		{{Universe: 0, Channel: 255, Value: cntl.DMXValue{Value: 12}}},
		{{Universe: 45, Channel: 200, Value: cntl.DMXValue{Value: 15}}},
		{{Universe: 12, Channel: 0, Value: *fixtures.Value255}},
		{{Universe: 44, Channel: 55, Value: cntl.DMXValue{Value: 66}}},
		{{Universe: 41, Channel: 210, Value: cntl.DMXValue{Value: 115}}},
	}
	cs := []cntl.DMXCommands{
		{{Universe: 10, Channel: 15, Value: cntl.DMXValue{Value: 1}}},
		{{Universe: 11, Channel: 16, Value: cntl.DMXValue{Value: 15}}},
	}
	e := []cntl.DMXCommands{
		{{Universe: 0, Channel: 255, Value: cntl.DMXValue{Value: 12}}},
		{{Universe: 45, Channel: 200, Value: cntl.DMXValue{Value: 15}}},
		{{Universe: 12, Channel: 0, Value: *fixtures.Value255}, {Universe: 10, Channel: 15, Value: cntl.DMXValue{Value: 1}}},
		{{Universe: 44, Channel: 55, Value: cntl.DMXValue{Value: 66}}, {Universe: 11, Channel: 16, Value: cntl.DMXValue{Value: 15}}},
		{{Universe: 41, Channel: 210, Value: cntl.DMXValue{Value: 115}}},
	}

	res := MergeAtOffset(cmds, cs, 2)
	for i, c := range e {
		if !c.Equals(res[i]) {
			t.Errorf("Expected %+v to equal %+v at index %d but doesn't.", c, res[i], i)
		}
	}
}
