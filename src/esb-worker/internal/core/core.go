package core

import (
	"esb-worker/internal/core/driver"
	"esb-worker/internal/core/libs"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

//Global variables
var (
	ESB *libs.ESB
)

var log = logrus.New()

type ESBApp struct {
	Name    string `default:"T7G App"`
	Usage   string `default:"T7G App worker"`
	Author  string `default:"T7G"`
	Email   string `default:"t7g-app@gmail.com"`
	Version string `default:"0.0.1"`
}

type App struct {
	app *cli.App
}

// Start process
func Start(esb libs.ESB) {

	var esbApp ESBApp
	var process App

	// Initialise a CLI app
	process.app = cli.NewApp()
	process.app.Name = esbApp.Name
	process.app.Usage = esbApp.Usage
	process.app.Author = esbApp.Author
	process.app.Email = esbApp.Email
	process.app.Version = esbApp.Version
	process.app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "c",
			Value:       "",
			Destination: &esb.ConfigPath,
			Usage:       "Path to a configuration file",
		},
		cli.StringFlag{
			Name:        "service",
			Value:       "",
			Destination: &esb.Service,
			Usage:       "a service name registed",
		},
	}

	process.app.Commands = []cli.Command{
		{
			Name:  "worker",
			Usage: "launch a worker",
			Action: func(c *cli.Context) error {

				var Setting, err = esb.GetConfig()

				//Connect database
				var adapter driver.Postgres
				driver.Connection, _ = adapter.Engine(esb.Config)

				if err != nil {
					return cli.NewExitError(fmt.Errorf("Config Parse error %v", Setting), 1)
				}

				if err := esb.GetConsumer(); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
	}

	// Run the CLI app
	process.app.Run(os.Args)
}
