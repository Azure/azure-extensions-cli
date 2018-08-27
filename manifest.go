package main

import (
	"encoding/xml"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"strings"
)

type certificate struct {
	StoreLocation       string `xml:"StoreLocation,omitempty"`
	StoreName           string `xml:"StoreName,omitempty"`
	ThumbprintRequired  bool   `xml:"ThumbprintRequired,omitempty"`
	ThumbprintAlgorithm string `xml:"ThumbprintAlgorithm,omitempty"`
}

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
type extensionImage struct {
	XMLName                     string       `xml:"ExtensionImage"`
	NS                          string       `xml:"xmlns,attr"`
	ProviderNameSpace           string       `xml:"ProviderNameSpace"`
	Type                        string       `xml:"Type"`
	Version                     string       `xml:"Version"`
	Label                       string       `xml:"Label"`
	HostingResources            string       `xml:"HostingResources"`
	MediaLink                   string       `xml:"MediaLink"`
	Endpoints                   string       `xml:"Endpoints"`
	Certificate                 *certificate `xml:"Certificate,omitempty"`
	PublicConfigurationSchema   string       `xml:"PublicConfigurationSchema,omitempty"`
	PrivateConfigurationSchema  string       `xml:"PrivateConfigurationSchema,omitempty"`
	Description                 string       `xml:"Description"`
	LocalResources              string       `xml:"LocalResources"`
	BlockRoleUponFailure        string       `xml:"BlockRoleUponFailure,omitempty"`
	IsInternalExtension         bool         `xml:"IsInternalExtension"`
	Eula                        string       `xml:"Eula,omitempty"`
	PrivacyURI                  string       `xml:"PrivacyUri,omitempty"`
	HomepageURI                 string       `xml:"HomepageUri,omitempty"`
	IsJSONExtension             bool         `xml:"IsJsonExtension,omitempty"`
	DisallowMajorVersionUpgrade bool         `xml:"DisallowMajorVersionUpgrade,omitempty"`
	CompanyName                 string       `xml:"CompanyName,omitempty"`
	SupportedOS                 string       `xml:"SupportedOS,omitempty"`
	Regions                     string       `xml:"Regions,omitempty"`
	Boumenot                    string       `xml:"Regions,omitempty"`
}

type extensionImageGlobal struct {
	XMLName                     string       `xml:"ExtensionImage"`
	NS                          string       `xml:"xmlns,attr"`
	ProviderNameSpace           string       `xml:"ProviderNameSpace"`
	Type                        string       `xml:"Type"`
	Version                     string       `xml:"Version"`
	Label                       string       `xml:"Label"`
	HostingResources            string       `xml:"HostingResources"`
	Endpoints                   string       `xml:"Endpoints"`
	MediaLink                   string       `xml:"MediaLink"`
	Certificate                 *certificate `xml:"Certificate,omitempty"`
	PublicConfigurationSchema   string       `xml:"PublicConfigurationSchema,omitempty"`
	PrivateConfigurationSchema  string       `xml:"PrivateConfigurationSchema,omitempty"`
	Description                 string       `xml:"Description"`
	LocalResources              string       `xml:"LocalResources"`
	BlockRoleUponFailure        string       `xml:"BlockRoleUponFailure,omitempty"`
	IsInternalExtension         bool         `xml:"IsInternalExtension"`
	Eula                        string       `xml:"Eula,omitempty"`
	PrivacyURI                  string       `xml:"PrivacyUri,omitempty"`
	HomepageURI                 string       `xml:"HomepageUri,omitempty"`
	IsJSONExtension             bool         `xml:"IsJsonExtension,omitempty"`
	DisallowMajorVersionUpgrade bool         `xml:"DisallowMajorVersionUpgrade,omitempty"`
	CompanyName                 string       `xml:"CompanyName,omitempty"`
	SupportedOS                 string       `xml:"SupportedOS,omitempty"`
	Regions                     string       `xml:"Regions"`
}

type extensionManifest interface {
	Marshal() ([]byte, error)
}

func isGuestAgent(providerNameSpace string) bool {
	return "Microsoft.OSTCLinuxAgent" == providerNameSpace
}

func newExtensionImageManifest(filename string, regions []string) (extensionManifest, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var manifest extensionImage
	err = xml.Unmarshal(b, &manifest)
	if err != nil {
		return nil, err
	}

	manifest.Regions = strings.Join(regions, ";")
	manifest.IsInternalExtension = isGuestAgent(manifest.ProviderNameSpace)

	return &manifest, nil
}

func newExtensionImageGlobalManifest(filename string) (extensionManifest, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var manifest extensionImageGlobal
	err = xml.Unmarshal(b, &manifest)
	if err != nil {
		return nil, err
	}

	manifest.IsInternalExtension = isGuestAgent(manifest.ProviderNameSpace)
	return &manifest, nil
}

func (ext *extensionImage) Marshal() ([]byte, error) {
	return xml.Marshal(*ext)
}

func (ext *extensionImageGlobal) Marshal() ([]byte, error) {
	return xml.Marshal(*ext)
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

	manifest := extensionImage{
		ProviderNameSpace:   checkFlag(c, flNamespace.Name),
		Type:                checkFlag(c, flName.Name),
		Version:             checkFlag(c, flVersion.Name),
		Label:               "label",
		Description:         "description",
		IsInternalExtension: true,
		MediaLink:           blobURL,
		Eula:                "eula-url",
		PrivacyURI:          "privacy-url",
		HomepageURI:         "homepage-url",
		IsJSONExtension:     true,
		CompanyName:         "company",
		SupportedOS:         "supported-os",
	}

	bs, err := xml.MarshalIndent(manifest, "", "  ")
	if err != nil {
		log.Fatalf("xml marshall error: %v", err)
	}

	fmt.Println(string(bs))
}
