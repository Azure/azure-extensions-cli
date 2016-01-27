package main

import (
	"io/ioutil"
	"time"

	"github.com/Azure/azure-sdk-for-go/management/storageservice"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

const (
	containerName = "extension-packages"
)

func createExtension(c *cli.Context) {
	cl := mkClient(checkFlag(c, flSubsID.Name), checkFlag(c, flSubsCert.Name))
	b, err := ioutil.ReadFile(checkFlag(c, flManifest.Name))
	if err != nil {
		log.Fatalf("Error reading manifest: %v", err)
	}

	// Fetch keys for storage account
	stg := checkFlag(c, flStorageAccount.Name)
	svc := storageservice.NewClient(cl.client)
	keys, err := svc.GetStorageServiceKeys(stg)
	if err != nil {
		log.WithField("name", "stg").Fatalf("Could not fetch keys for storage account. Make sure it is in publisher subscription. Error: %v", err)
	}
	key := keys.PrimaryKey

	// Upload blob

	op, err := cl.CreateExtension(ns, name, version)
	if err != nil {
		log.Fatalf("Error creating extension: %v", err)
	}
	log.Debug("CreateExtension operation started.")
	if err := cl.WaitForOperation(op); err != nil {
		log.Fatalf("CreateExtension failed: %v", err)
	}
	log.Info("CreateExtension operation finished.")

}
