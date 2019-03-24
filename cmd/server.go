package cmd

import (
	"fmt"

	"github.com/StageAutoControl/controller/pkg/api/server"
	"github.com/StageAutoControl/controller/pkg/cntl/playback"
	"github.com/StageAutoControl/controller/pkg/disk"
	"github.com/StageAutoControl/controller/pkg/process"
	"github.com/apinnecke/go-exitcontext"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Opens the RPC API to manage the data and control the processes",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := exitcontext.New()
		pm := process.NewManager(ctx, logger)
		server, err := server.New(logger.WithField("module", "api"), storage, controller, pm)
		if err != nil {
			logger.Fatal(err)
		}

		port, err := cmd.Flags().GetUint16("port")
		if err != nil {
			logger.Fatal(err)
		}

		endpoint := fmt.Sprintf("0.0.0.0:%d", port)
		loader := disk.NewLoader(storage)

		if !disableController {
			if err := playback.EnsureDefaultConfig(storage); err != nil {
				logger.Fatal(err)
			}
			if err := pm.AddProcess(playback.ProcessName, playback.NewProcess(loader, storage, controller), true); err != nil {
				logger.Fatal(err)
			}

		}
		if err := server.Run(ctx, endpoint); err != nil {
			logger.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)

	serverCmd.Flags().Uint16P("port", "p", 8080, "TCP port the API should listen on")
}
