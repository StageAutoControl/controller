// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>

package main

import "github.com/StageAutoControl/controller/cmd"
import _ "github.com/StageAutoControl/controller/cmd/artnet"
import _ "github.com/StageAutoControl/controller/cmd/midi"

func main() {
	cmd.Execute()
}
