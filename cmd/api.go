package cmd

import (
	"fmt"

	"github.com/StageAutoControl/controller/cmd/internal"
	"github.com/StageAutoControl/controller/pkg/api"
	"github.com/StageAutoControl/controller/pkg/storage"
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Opens the RPC API to manage the data and control the processes",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logger := Logger.WithField("module", "api")

		storagePath, err := cmd.Flags().GetString("storage-path")
		if err != nil {
			logger.Fatal(err)
		}

		store := storage.New(storagePath)
		server, err := api.NewServer(logger, store)
		if err != nil {
			logger.Fatal(err)
		}

		port, err := cmd.Flags().GetUint16("port")
		if err != nil {
			logger.Fatal(err)
		}

		endpoint := fmt.Sprintf("0.0.0.0:%d", port)
		ctx := internal.NewExitHandlerContext(logger.Logger)

		logger.Infof("listening on %s", endpoint)

		if err := server.Run(ctx, endpoint); err != nil {
			logger.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(apiCmd)

	apiCmd.Flags().Uint16P("port", "p", 8080, "TCP port the API should listen on")
	apiCmd.Flags().StringP("storage-path", "s", "/var/controller/data", "path where the storage should store the data")
}
