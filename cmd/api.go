// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>
//

package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/StageAutoControl/controller/api"
	"github.com/StageAutoControl/controller/database/files"
	"github.com/gorilla/handlers"
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
		repo := files.New(dataDir)

		handler := api.NewRepoLockingHandler(repo, api.NewHandler(repo))
		mux.Handle("/rpc", handler)

		h := handlers.CORS(
			handlers.AllowCredentials(),
			handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Language", "Origin", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"HEAD", "POST"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(mux)

		// Add Logging middleware
		h = handlers.LoggingHandler(os.Stdout, h)
		// Add recovery handler to fetch panics
		h = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(h)

		log.Printf("Starting api server at %s", apiListen)
		if err := http.ListenAndServe(apiListen, h); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(apiCmd)

	apiCmd.Flags().StringVarP(&apiListen, "listen", "l", "0.0.0.0:8080", "The listen string to bind the api server to")
}
