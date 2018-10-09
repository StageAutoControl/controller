// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>

package cmd

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// The Logger used by the whole application
	Logger   *logrus.Entry
	logLevel string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "controller",
	Short: "Stage automatic controlling, triggering state changes.",
	Long:  `Automatic stage controlling, including midi and DMX, by analyzing audio signals and pre defined light scenes`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger := logrus.New()
		level, err := logrus.ParseLevel(logLevel)
		if err != nil {
			logger.Panicf("Unable to parse log level %q: %v\n", logLevel, err)
			os.Exit(1)
		}

		logger.Infof("Using log level %s", logLevel)

		logger.SetLevel(level)
		Logger = logger.WithFields(logrus.Fields{})
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
}
