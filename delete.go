package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func deleteVersion(c *cli.Context) {
	cl := mkClient(getFlag(c, flMgtURL.Name), checkFlag(c, flSubsID.Name), checkFlag(c, flSubsCert.Name))
	ns, name, version := checkFlag(c, flNamespace.Name), checkFlag(c, flName.Name), checkFlag(c, flVersion.Name)
	log.Info("Deleting extension version. Make sure you unpublished before deleting.")

	op, err := cl.DeleteExtension(ns, name, version)
	if err != nil {
		log.Fatalf("Error deleting version: %v", err)
	}
	log.Debug("DeleteExtension operation started.")
	if err := cl.WaitForOperation(op); err != nil {
		log.Fatalf("DeleteExtension failed: %v", err)
	}
	log.Info("DeleteExtension operation finished.")
}
