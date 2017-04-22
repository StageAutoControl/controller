// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>
//

package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/StageAutoControl/controller/api"
	"github.com/spf13/cobra"
)

var (
	apiListen string
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "API to manage the data",
	Long:  `A JSON RPC server to manage the data handled by this controller.`,
	Run: func(cmd *cobra.Command, args []string) {
		mux := http.NewServeMux()
		handler := api.NewHandler()
		mux.Handle("/rpc", handler)

		if err := http.ListenAndServe(apiListen, mux); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(apiCmd)

	apiCmd.Flags().StringVarP(&apiListen, "listen", "l", "0.0.0.0:8080", "The listen string to bind the api server to")
}
