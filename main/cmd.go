package main

import (
	"avr/logger"
	"flag"
	"os"
	"path/filepath"
)

var (
	queryPath = flag.String("config", "D://", "")
	build = ""
)

func main() {
	flag.Parse()
	logger.Info("Build: " + build)

	err := filepath.Walk(*queryPath, func (path string, info os.FileInfo, err error) error {
		logger.Info(path, info, err)
		return nil
	})
	if err != nil {
		logger.Info("Filepath.Walk Error:", err.Error())
	}
}
