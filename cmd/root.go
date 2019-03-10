// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>

package cmd

import (
	"fmt"
	"os"

	"github.com/StageAutoControl/controller/pkg/artnet"
	"github.com/StageAutoControl/controller/pkg/disk"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	logLevel    string
	logger      *logrus.Entry
	storagePath string
	storage     *disk.Storage
	controller  artnet.Controller
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "controller",
	Short: "Stage automatic controlling, triggering state changes.",
	Long:  `Automatic stage controlling, including midi and DMX, by analyzing audio signals and pre defined light scenes`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger = createLogger(logLevel)
		storage = createStorage(logger, storagePath)
		controller = createController(logger)
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "Which log level to use")
	RootCmd.PersistentFlags().StringVarP(&storagePath, "storage-path", "s", "/var/controller/data", "path where the storage should store the data")
}
