package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func promoteToFirstSlice(c *cli.Context) {
	regions := []string{
		checkFlag(c, flRegion1.Name),
	}

	if err := promoteExtension(c, func() (extensionManifest, error) {
		return newExtensionImageManifest(checkFlag(c, flManifest.Name), regions)
	}); err != nil {
		log.Fatal(err)
	}

	log.Info("Extension is promoted to PROD in one region. See replication-status.")
}
func promoteToSecondSlice(c *cli.Context) {
	regions := []string{
		checkFlag(c, flRegion1.Name),
		checkFlag(c, flRegion2.Name),
	}

	if err := promoteExtension(c, func() (extensionManifest, error) {
		return newExtensionImageManifest(checkFlag(c, flManifest.Name), regions)
	}); err != nil {
		log.Fatal(err)
	}

	log.Info("Extension is promoted to PROD in two regions. See replication-status.")
}

func promoteToAllRegions(c *cli.Context) {
	if err := promoteExtension(c, func() (extensionManifest, error) {
		return newExtensionImageGlobalManifest(checkFlag(c, flManifest.Name))
	}); err != nil {
		log.Fatal(err)
	}

	log.Info("Extension is promoted to all regions. See replication-status.")
}

func promoteExtension(c *cli.Context, factory func() (extensionManifest, error)) error {
	manifest, err := factory()
	if err != nil {
		return err
	}

	b, err := manifest.Marshal()
	if err != nil {
		return err
	}

	if err := publishExtension(c, "UpdateExtension", b,
		mkClient(checkFlag(c, flMgtURL.Name), checkFlag(c, flSubsID.Name), checkFlag(c, flSubsCert.Name)).UpdateExtension); err != nil {
		return err
	}
	return nil
}
