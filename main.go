// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>

package main

import (
	"github.com/StageAutoControl/controller/cmd"
	_ "github.com/StageAutoControl/controller/cmd/artnet"
	_ "github.com/StageAutoControl/controller/cmd/audio"
	_ "github.com/StageAutoControl/controller/cmd/midi"
)

func main() {
	cmd.Execute()
}
