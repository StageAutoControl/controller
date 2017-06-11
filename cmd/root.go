// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	directoryLoader = "directory"
	databaseLoader  = "database"
)

var (
	loaders    = []string{directoryLoader, databaseLoader}
	loaderType string
	dataDir    string
	cfgFile    string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "controller",
	Short: "Stage automatic controlling, triggering state changes.",
	Long:  `Automatic stage controlling, including midi and DMX, by analyzing audio signals and pre defined light scenes`,
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
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVarP(&dataDir, "data-dir", "d", "", "Data directory to load (when loader is set to directory)")
	RootCmd.PersistentFlags().StringVar(&loaderType, "loader", directoryLoader, fmt.Sprintf("Which loader to use %s.", loaders))

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sac-controller.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".sac-controller")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
