package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func promoteToFirstSlice(c *cli.Context) {
	if err := promoteExtension(c, mkRegionElement(
		checkFlag(c, flRegion1.Name))); err != nil {
		log.Fatal(err)
	}
	log.Info("Extension is promoted to PROD in one region. See replication-status.")
}
func promoteToSecondSlice(c *cli.Context) {
	if err := promoteExtension(c, mkRegionElement(
		checkFlag(c, flRegion1.Name),
		checkFlag(c, flRegion2.Name))); err != nil {
		log.Fatal(err)
	}
	log.Info("Extension is promoted to PROD in two regions. See replication-status.")
}

func promoteToAllRegions(c *cli.Context) {
	regions := `` // replace placeholder with empty string to omit the element.
	if err := promoteExtension(c, regions); err != nil {
		log.Fatal(err)
	}
	log.Info("Extension is promoted to all regions. See replication-status.")
}

func promoteExtension(c *cli.Context, regionsXML string) error {
	b, err := updateManifestRegions(checkFlag(c, flManifest.Name), regionsXML)
	if err != nil {
		return err
	}
	if err := publishExtension(c, "UpdateExtension", b,
		mkClient(checkFlag(c, flMgtURL.Name), checkFlag(c, flSubsID.Name), checkFlag(c, flSubsCert.Name)).UpdateExtension); err != nil {
		return err
	}
	return nil
}

func mkRegionElement(regions ...string) string {
	return fmt.Sprintf(`<Regions>%s</Regions>`, strings.Join(regions, ";"))
}

// updateManifestRegions makes an in-memory update to the <!--%REGIONS%-->
// placeholder string in the manifest XML for further usage and replaces
// <IsInternalExtension>true... with ...false.
func updateManifestRegions(manifestPath string, regionsXMLElement string) ([]byte, error) {
	b, err := ioutil.ReadFile(manifestPath)
	if err != nil {
		return nil, fmt.Errorf("Error reading manifest: %v", err)
	}

	// I know I can do better than this, but will I?
	b = bytes.Replace(b, []byte(`<!--%REGIONS%-->`), []byte(regionsXMLElement), 1)
	return bytes.Replace(b, []byte(`<IsInternalExtension>true`), []byte(`<IsInternalExtension>false`), 1), nil
}
