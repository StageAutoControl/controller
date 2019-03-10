package cmd

import (
	"fmt"

	"github.com/StageAutoControl/controller/pkg/api"
	"github.com/apinnecke/go-exitcontext"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Opens the RPC API to manage the data and control the processes",
	Run: func(cmd *cobra.Command, args []string) {
		logger := logger.WithField("module", "server")
		server, err := api.NewServer(logger, storage, controller)
		if err != nil {
			logger.Fatal(err)
		}

		port, err := cmd.Flags().GetUint16("port")
		if err != nil {
			logger.Fatal(err)
		}

		endpoint := fmt.Sprintf("0.0.0.0:%d", port)
		ctx := exitcontext.New()

		logger.Infof("listening on %s", endpoint)

		if err := server.Run(ctx, endpoint); err != nil {
			logger.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)

	serverCmd.Flags().Uint16P("port", "p", 8080, "TCP port the API should listen on")
}