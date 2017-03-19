package cntl

import "testing"

func TestDMXCommand_Equals(t *testing.T) {
	exp := []struct {
		cmd1  DMXCommand
		cmd2  DMXCommand
		equal bool
	}{
		{
			cmd1:  DMXCommand{Universe: 1, Channel: 1, Value: 12},
			cmd2:  DMXCommand{Universe: 1, Channel: 1, Value: 12},
			equal: true,
		},
		{
			cmd1:  DMXCommand{Universe: 2, Channel: 1, Value: 12},
			cmd2:  DMXCommand{Universe: 1, Channel: 1, Value: 12},
			equal: false,
		},
		{
			cmd1:  DMXCommand{Universe: 1, Channel: 2, Value: 12},
			cmd2:  DMXCommand{Universe: 1, Channel: 1, Value: 12},
			equal: false,
		},
		{
			cmd1:  DMXCommand{Universe: 1, Channel: 1, Value: 12},
			cmd2:  DMXCommand{Universe: 1, Channel: 1, Value: 14},
			equal: false,
		},
		{
			cmd1:  DMXCommand{},
			cmd2:  DMXCommand{},
			equal: true,
		},
	}

	for _, e := range exp {
		eq := e.cmd1.Equals(e.cmd2)
		if eq != e.equal {
			t.Errorf("Expected %+v to be equal to %+v, but Equals sais no.", e.cmd1, e.cmd2)
		}
	}
}

func TestDMXCommands_Equals(t *testing.T) {
	exp := []struct {
		cmd1  DMXCommands
		cmd2  DMXCommands
		equal bool
	}{
		{
			cmd1:  DMXCommands{{Universe: 1, Channel: 1, Value: 12}},
			cmd2:  DMXCommands{{Universe: 1, Channel: 1, Value: 12}},
			equal: true,
		},
		{
			cmd1:  DMXCommands{{Universe: 2, Channel: 1, Value: 12}},
			cmd2:  DMXCommands{{Universe: 1, Channel: 1, Value: 12}},
			equal: false,
		},
		{
			cmd1:  DMXCommands{{Universe: 1, Channel: 2, Value: 12}},
			cmd2:  DMXCommands{{Universe: 1, Channel: 1, Value: 12}},
			equal: false,
		},
		{
			cmd1:  DMXCommands{{Universe: 1, Channel: 1, Value: 12}},
			cmd2:  DMXCommands{{Universe: 1, Channel: 1, Value: 14}},
			equal: false,
		},
		{
			cmd1:  DMXCommands{{Universe: 1, Channel: 1, Value: 12}},
			cmd2:  DMXCommands{{Universe: 1, Channel: 1, Value: 12}, {Universe: 1, Channel: 1, Value: 12}},
			equal: false,
		},
		{
			cmd1:  DMXCommands{{Universe: 1, Channel: 1, Value: 12}},
			cmd2:  DMXCommands{},
			equal: false,
		},
		{
			cmd1:  DMXCommands{},
			cmd2:  DMXCommands{},
			equal: true,
		},
	}

	for _, e := range exp {
		eq := e.cmd1.Equals(e.cmd2)
		if eq != e.equal {
			t.Errorf("Expected %+v to be equal to %+v, but Equals sais no.", e.cmd1, e.cmd2)
		}
	}

}

func TestDMXCommands_Contains(t *testing.T) {
	exp := []struct {
		cmd1     DMXCommands
		cmd2     DMXCommand
		contains bool
	}{
		{
			cmd1:     DMXCommands{{Universe: 1, Channel: 1, Value: 12}},
			cmd2:     DMXCommand{Universe: 1, Channel: 1, Value: 12},
			contains: true,
		},
		{
			cmd1:     DMXCommands{{Universe: 2, Channel: 1, Value: 12}},
			cmd2:     DMXCommand{Universe: 1, Channel: 1, Value: 12},
			contains: false,
		},
		{
			cmd1:     DMXCommands{{Universe: 1, Channel: 2, Value: 12}},
			cmd2:     DMXCommand{Universe: 1, Channel: 1, Value: 12},
			contains: false,
		},
		{
			cmd1:     DMXCommands{{Universe: 1, Channel: 1, Value: 12}},
			cmd2:     DMXCommand{Universe: 1, Channel: 1, Value: 14},
			contains: false,
		},
		{
			cmd1:     DMXCommands{{Universe: 1, Channel: 1, Value: 12}, {Universe: 1, Channel: 1, Value: 13}, {Universe: 1, Channel: 1, Value: 2}},
			cmd2:     DMXCommand{Universe: 1, Channel: 1, Value: 12},
			contains: true,
		},
		{
			cmd1:     DMXCommands{{Universe: 1, Channel: 1, Value: 12}, {}},
			cmd2:     DMXCommand{},
			contains: true,
		},
		{
			cmd1:     DMXCommands{{Universe: 1, Channel: 1, Value: 12}},
			cmd2:     DMXCommand{},
			contains: false,
		},
		{
			cmd1:     DMXCommands{},
			cmd2:     DMXCommand{},
			contains: false,
		},
	}

	for _, e := range exp {
		eq := e.cmd1.Contains(e.cmd2)
		if eq != e.contains {
			t.Errorf("Expected %+v to be contain to %+v, but Equals sais no.", e.cmd1, e.cmd2)
		}
	}
}
