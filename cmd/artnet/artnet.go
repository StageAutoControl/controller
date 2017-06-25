// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>
//

package artnet

import (
	"github.com/StageAutoControl/controller/cmd"
	"github.com/spf13/cobra"
)

// ArtNetCmd represents the ArtNetTest command namespaces
var ArtNetCmd = &cobra.Command{
	Use:   "artnet",
	Short: "ArtNet commands to work with ArtNet devices",
	Long:  ``,
}

func init() {
	cmd.RootCmd.AddCommand(ArtNetCmd)
}
