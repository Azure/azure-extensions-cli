package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

func publishExtension(c *cli.Context, operationName string, manifest []byte, op func([]byte) (management.OperationID, error)) error {
	log.Infof("%s operation starting...", operationName)

	mPath, err := saveManifestForDebugging(manifest)
	if err != nil {
		return fmt.Errorf("Error saving manifest for debugging: %v", err)
	}
	log.Debugf("Saving used manifest for debugging: %s", mPath)

	cl := mkClient(checkFlag(c, flMgtURL.Name), checkFlag(c, flSubsID.Name), checkFlag(c, flSubsCert.Name))
	opID, err := op(manifest)
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

func saveManifestForDebugging(contents []byte) (string, error) {
	dir, err := ioutil.TempDir("", "extension-manifests")
	if err != nil {
		return "", err
	}
	f, err := ioutil.TempFile(dir, "manifest")
	if err != nil {
		return "", err
	}
	if _, err := f.Write(contents); err != nil {
		return "", err
	}
	fi, err := f.Stat()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, fi.Name()), nil
}

func publishExtensionFromManifestFile(c *cli.Context, operationName, manifestPath string, op func([]byte) (management.OperationID, error)) error {
	b, err := ioutil.ReadFile(manifestPath)
	if err != nil {
		return fmt.Errorf("Error reading manifest: %v", err)
	}
	return publishExtension(c, operationName, b, op)
}

func createExtension(c *cli.Context) {
	cl := mkClient(checkFlag(c, flMgtURL.Name), checkFlag(c, flSubsID.Name), checkFlag(c, flSubsCert.Name))
	if err := publishExtensionFromManifestFile(c, "CreateExtension",
		checkFlag(c, flManifest.Name), cl.CreateExtension); err != nil {
		log.Fatal(err)
	}
}

func updateExtension(c *cli.Context) {
	cl := mkClient(checkFlag(c, flMgtURL.Name), checkFlag(c, flSubsID.Name), checkFlag(c, flSubsCert.Name))
	if err := publishExtensionFromManifestFile(c, "UpdateExtension", checkFlag(c, flManifest.Name),
		cl.UpdateExtension); err != nil {
		log.Fatal(err)
	}
}

func uploadBlob(cl ExtensionsClient, storageRealm, storageAccount, packagePath string) (string, error) {
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
	sc, err := storage.NewClient(storageAccount, keys.PrimaryKey, storageRealm, storage.DefaultAPIVersion, true)
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
