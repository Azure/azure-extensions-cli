package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/management/storageservice"
	"github.com/Azure/azure-sdk-for-go/storage"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

const (
	containerName = "extension-packages"
)

func createExtension(c *cli.Context) {
	cl := mkClient(checkFlag(c, flSubsID.Name), checkFlag(c, flSubsCert.Name))

	// Upload blob
	blobURL := uploadBlob(cl, checkFlag(c, flStorageAccount.Name), checkFlag(c, flPackage.Name))
	log.Debugf("Extension package uploaded to: %s", blobURL)

	// Read manifest
	b, err := ioutil.ReadFile(checkFlag(c, flManifest.Name))
	if err != nil {
		log.Fatalf("Error reading manifest: %v", err)
	}

	// Replace %BLOB_URL% in the manifest.
	manifest := string(b)
	manifest = strings.Replace(manifest, "%BLOB_URL%", blobURL, 1)

	// Initiate Create Extension operation
	op, err := cl.CreateExtension([]byte(manifest))
	if err != nil {
		log.Fatalf("Error creating extension: %v", err)
	}
	log.Debug("CreateExtension operation started.")
	if err := cl.WaitForOperation(op); err != nil {
		log.Fatalf("CreateExtension failed: %v", err)
	}
	log.Info("CreateExtension operation finished.")
	log.Info("Next steps: Test in your subscription, promote to pilot regions and then PROD.")
}

func updateExtension(c *cli.Context) {
}

func uploadBlob(cl ExtensionsClient, storageAccount, packagePath string) string {
	// Fetch keys for storage account
	svc := storageservice.NewClient(cl.client)
	keys, err := svc.GetStorageServiceKeys(storageAccount)
	if err != nil {
		log.WithField("name", "storageAccount").Fatalf("Could not fetch keys for storage account. Make sure it is in publisher subscription. Error: %v", err)
	}
	log.Debug("Retrieved storage account keys.")

	// Read package
	pkg, err := os.OpenFile(packagePath, os.O_RDONLY, 0777)
	if err != nil {
		log.Fatalf("Could not reach package file: %v", err)
	}
	stat, err := pkg.Stat()
	if err != nil {
		log.Fatalf("Could not stat the package file: %v", err)
	}
	defer pkg.Close()

	// Upload blob
	sc, err := storage.NewBasicClient(storageAccount, keys.PrimaryKey)
	if err != nil {
		log.Fatalf("Could not create storage client: %v", err)
	}
	bs := sc.GetBlobService()
	if _, err := bs.CreateContainerIfNotExists(containerName, storage.ContainerAccessTypeBlob); err != nil {
		log.Fatalf("Error creating blob container: %v", err)
	}
	blobName := fmt.Sprintf("%d.zip", time.Now().Unix())
	if err := bs.CreateBlockBlobFromReader(containerName, blobName, uint64(stat.Size()), pkg, nil); err != nil {
		log.Fatalf("Error uploading blob: %v", err)
	}
	return bs.GetBlobURL(containerName, blobName)
}
