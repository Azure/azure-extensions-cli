package main

import (
	"encoding/xml"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

// NOTE(@boumenot): there is probably a better way to express this.  If
// you know please share...
//
// The only difference between ExtensionImage and ExtensionImageGlobal is the
// Regions element.  This element can be in three different states to my
// knowledge.
//
//   1. not defined
//   2. <Regions>Region1;Region2</Regions>
//   3. <Regions></Regions>
//
// Case (1) occurs when an extension is first published.  Case(2) occurs when
// an extension is promoted to one or two regions.  Case (3) occurs when an
// extension is published to all regions.
//
// I do not know how to express all three cases using Go's XML serializer.
//
type ExtensionImage struct {
	XMLName             string `xml:"ExtensionImage"`
	NS                  string `xml:"xmlns,attr"`
	ProviderNameSpace   string `xml:"ProviderNameSpace"`
	Type                string `xml:"Type"`
	Version             string `xml:"Version"`
	Label               string `xml:"Label"`
	HostingResources    string `xml:"HostingResources"`
	MediaLink           string `xml:"MediaLink"`
	Description         string `xml:"Description"`
	IsInternalExtension bool   `xml:"IsInternalExtension"`
	Eula                string `xml:"Eula"`
	PrivacyUri          string `xml:"PrivacyUri"`
	HomepageUri         string `xml:"HomepageUri"`
	IsJsonExtension     bool   `xml:"IsJsonExtension"`
	CompanyName         string `xml:"CompanyName"`
	SupportedOS         string `xml:"SupportedOS"`
	Regions             string `xml:"Regions,omitempty"`
}

type ExtensionImageGlobal struct {
	XMLName             string `xml:"ExtensionImage"`
	NS                  string `xml:"xmlns,attr"`
	ProviderNameSpace   string `xml:"ProviderNameSpace"`
	Type                string `xml:"Type"`
	Version             string `xml:"Version"`
	Label               string `xml:"Label"`
	HostingResources    string `xml:"HostingResources"`
	MediaLink           string `xml:"MediaLink"`
	Description         string `xml:"Description"`
	IsInternalExtension bool   `xml:"IsInternalExtension"`
	Eula                string `xml:"Eula"`
	PrivacyUri          string `xml:"PrivacyUri"`
	HomepageUri         string `xml:"HomepageUri"`
	IsJsonExtension     bool   `xml:"IsJsonExtension"`
	CompanyName         string `xml:"CompanyName"`
	SupportedOS         string `xml:"SupportedOS"`
	Regions             string `xml:"Regions"`
}

func newExtensionManifest(c *cli.Context) {
	cl := mkClient(checkFlag(c, flMgtURL.Name), checkFlag(c, flSubsID.Name), checkFlag(c, flSubsCert.Name))
	storageRealm := checkFlag(c, flStorageRealm.Name)
	storageAccount := checkFlag(c, flStorageAccount.Name)
	extensionPkg := checkFlag(c, flPackage.Name)

	// Upload extension blob
	blobURL, err := uploadBlob(cl, storageRealm, storageAccount, extensionPkg)
	if err != nil {
		log.Fatal(err)
	}
	log.Debugf("Extension package uploaded to: %s", blobURL)

	manifest := ExtensionImage{
		ProviderNameSpace: checkFlag(c, flNamespace.Name),
		Type: checkFlag(c, flName.Name),
		Version: checkFlag(c, flVersion.Name),
		Label: "label",
		Description: "description",
		MediaLink: blobURL,
		Eula: "eula-url",
		PrivacyUri: "privacy-url",
		HomepageUri: "homepage-url",
		CompanyName: "company",
		SupportedOS: "supported-os",
	}

	bs, err := xml.MarshalIndent(manifest, "", "  ")
	if err != nil {
		log.Fatalf("xml marshall error: %v", err)
	}

	fmt.Println(string(bs))
}
