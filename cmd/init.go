package cmd

import (
	"os"
	"path/filepath"

	"github.com/StageAutoControl/controller/pkg/artnet"
	"github.com/StageAutoControl/controller/pkg/disk"
	"github.com/sirupsen/logrus"
)

func createLogger(logLevel string) *logrus.Entry {
	logger := logrus.New()
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logger.Panicf("Unable to parse log level %q: %v\n", logLevel, err)
		os.Exit(1)
	}

	logger.Infof("Using log level %s", logLevel)

	logger.SetLevel(level)
	return logger.WithFields(logrus.Fields{})
}

func createStorage(logger *logrus.Entry, storagePath string) *disk.Storage {
	if filepath.IsAbs(storagePath) {
		return disk.New(storagePath)
	}

	cwd, err := os.Getwd()
	if err != nil {
		logger.Fatal(err)
	}
	storagePath = filepath.Clean(filepath.Join(cwd, storagePath))
	if err != nil {
		logger.Fatal(err)
	}

	return disk.New(storagePath)
}

func createController(logger *logrus.Entry, disable bool) artnet.Controller {
	if disable {
		logger.Warn("ArtNet controller is disabled, so no playback or playground will be possible!")
		return nil
	}

	c, err := artnet.NewController(logger.WithField("module", "controller"))
	if err != nil {
		logger.Fatal(err)
	}

	return c
}
