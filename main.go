package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	app := cli.NewApp()
	app.Name = "azure-extensions-cli"
	app.Usage = "This tool is designed for Microsoft internal extension publishers to release, update and manage Virtual Machine extensions."
	app.Authors = []cli.Author{{Name: "Ahmet Alp Balkan", Email: "ahmetb at microsoft döt com"}}
	app.RunAndExitOnError()
}
