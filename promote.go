package main

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func promoteToRegions(c *cli.Context) {
	regions := c.StringSlice(flRegion.Name)

	if len(regions) == 0 {
		log.Fatalf("At least one region must be specified!")
		return
	}

	normalizedRegions := normalizeRegionList(regions)

	if err := promoteExtension(c, func() (extensionManifest, error) {
		return newExtensionImageManifest(checkFlag(c, flManifest.Name), normalizedRegions)
	}); err != nil {
		log.Fatal(err)
	}

	log.Infof("Extension is promoted to PROD in %s. See replication-status.", strings.Join(regions, ","))
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

	return publishExtension(c, "UpdateExtension", b,
		mkClient(checkFlag(c, flMgtURL.Name), checkFlag(c, flSubsID.Name), checkFlag(c, flSubsCert.Name)).UpdateExtension)
}
