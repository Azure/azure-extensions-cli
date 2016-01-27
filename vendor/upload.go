package vendor

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
