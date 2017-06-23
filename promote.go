package main

import (
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func promoteRegions(c *cli.Context) {
	regions := c.StringSlice(flRegion.Name)

	log.Info(fmt.Sprintf("regions=%s", strings.Join(regions, ";")))

	if err := promoteExtension(c, func() (extensionManifest, error) {
		return newExtensionImageManifest(checkFlag(c, flManifest.Name), regions)
	}); err != nil {
		log.Fatal(err)
	}

	log.Info("Extension is promoted to PROD in one region. See replication-status.")
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
