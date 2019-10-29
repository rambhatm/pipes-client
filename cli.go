package main

import (
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "CopyMyFile"
	app.Usage = "P2P copy of files between nodes"
	app.Author = "rambhatm"
	app.Version = "1.0.0"
}
