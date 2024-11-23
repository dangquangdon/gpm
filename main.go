package main

import (
	"log"
	"os"

	"github.com/dangquangdon/gpm/cli"
	"github.com/dangquangdon/gpm/config"
)

var cfg config.Config
var logger *log.Logger

func init() {
	var err error
	logger = log.Default()
	cfg, err = config.LoadConfig(".")
	if err != nil {
		logger.Printf("No config file found. Create one automatically")
		err = config.CreateEmptyConfigFile(".")
		if err != nil {
			logger.Fatal(err)
		}
	}
}

func main() {
	app := cli.InitCli(cfg)
	app.Run(os.Args)
}
