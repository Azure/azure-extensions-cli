package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/management"
	"github.com/Azure/azure-sdk-for-go/management/storageservice"
	"github.com/Azure/azure-sdk-for-go/storage"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

const (
	containerName = "extension-packages"
)

func publishExtension(c *cli.Context, operationName string, op func([]byte) (management.OperationID, error)) error {
	// Read manifest
	b, err := ioutil.ReadFile(checkFlag(c, flManifest.Name))
	if err != nil {
		return fmt.Errorf("Error reading manifest: %v", err)
	}

	// Initiate operation and poll
	log.Debugf("%s operation starting...", operationName)
	cl := mkClient(checkFlag(c, flSubsID.Name), checkFlag(c, flSubsCert.Name))
	opID, err := op(b)
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}
	log.Debugf("%s operation started.", operationName)
	if err := cl.WaitForOperation(opID); err != nil {
		return fmt.Errorf("%s failed: %v", operationName, err)
	}
	log.Infof("%s operation finished.", operationName)
	return nil
}

func createExtension(c *cli.Context) {
	cl := mkClient(checkFlag(c, flSubsID.Name), checkFlag(c, flSubsCert.Name))
	if err := publishExtension(c, "CreateExtension", cl.CreateExtension); err != nil {
		log.Fatal(err)
	}
}

func updateExtension(c *cli.Context) {
	cl := mkClient(checkFlag(c, flSubsID.Name), checkFlag(c, flSubsCert.Name))
	if err := publishExtension(c, "UpdateExtension", cl.UpdateExtension); err != nil {
		log.Fatal(err)
	}
}

func uploadBlob(cl ExtensionsClient, storageAccount, packagePath string) (string, error) {
	// Fetch keys for storage account
	svc := storageservice.NewClient(cl.client)
	keys, err := svc.GetStorageServiceKeys(storageAccount)
	if err != nil {
		return "", fmt.Errorf("Could not fetch keys for storage account. Make sure it is in publisher subscription. Error: %v", err)
	}
	log.Debug("Retrieved storage account keys.")

	// Read package
	pkg, err := os.OpenFile(packagePath, os.O_RDONLY, 0777)
	if err != nil {
		return "", fmt.Errorf("Could not reach package file: %v", err)
	}
	stat, err := pkg.Stat()
	if err != nil {
		return "", fmt.Errorf("Could not stat the package file: %v", err)
	}
	defer pkg.Close()

	// Upload blob
	sc, err := storage.NewBasicClient(storageAccount, keys.PrimaryKey)
	if err != nil {
		return "", fmt.Errorf("Could not create storage client: %v", err)
	}
	bs := sc.GetBlobService()
	if _, err := bs.CreateContainerIfNotExists(containerName, storage.ContainerAccessTypeBlob); err != nil {
		return "", fmt.Errorf("Error creating blob container: %v", err)
	}
	blobName := fmt.Sprintf("%d.zip", time.Now().Unix())
	if err := bs.CreateBlockBlobFromReader(containerName, blobName, uint64(stat.Size()), pkg, nil); err != nil {
		return "", fmt.Errorf("Error uploading blob: %v", err)
	}
	return bs.GetBlobURL(containerName, blobName), nil
}
