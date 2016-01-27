package main

import (
	"os"
	"text/template"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func newExtensionManifest(c *cli.Context) {
	cl := mkClient(checkFlag(c, flSubsID.Name), checkFlag(c, flSubsCert.Name))
	storageAccount := checkFlag(c, flStorageAccount.Name)
	extensionPkg := checkFlag(c, flPackage.Name)

	var p struct {
		Namespace, Name, Version, BlobURL, Label, Description, Eula, Privacy, Homepage, Company, OS string
	}
	flags := []struct {
		ref *string
		fl  string
	}{
		{&p.Namespace, flNamespace.Name},
		{&p.Name, flName.Name},
		{&p.Version, flVersion.Name},
		{&p.Label, "label"},
		{&p.Description, "description"},
		{&p.Eula, "eula-url"},
		{&p.Privacy, "privacy-url"},
		{&p.Homepage, "homepage-url"},
		{&p.Company, "company"},
		{&p.OS, "supported-os"},
	}
	for _, f := range flags {
		*f.ref = checkFlag(c, f.fl)
	}

	// Upload extension blob
	blobURL, err := uploadBlob(cl, storageAccount, extensionPkg)
	if err != nil {
		log.Fatal(err)
	}
	log.Debugf("Extension package uploaded to: %s", blobURL)
	p.BlobURL = blobURL

	// doing a text template is easier and let us create comments (xml encoder can't)
	// that are used as placeholders later on.
	manifestXml := `<?xml version="1.0" encoding="utf-8" ?>
<ExtensionImage xmlns="http://schemas.microsoft.com/windowsazure"  xmlns:i="http://www.w3.org/2001/XMLSchema-instance">
  <!-- WARNING: Ordering of fields matter in this file. -->
  <ProviderNameSpace>{{.Namespace}}</ProviderNameSpace>
  <Type>{{.Name}}</Type>
  <Version>{{.Version}}</Version>
  <Label>{{.Label}}</Label>
  <HostingResources>VmRole</HostingResources>
  <MediaLink>{{.BlobURL}}</MediaLink>
  <Description>{{.Description}}</Description>
  <IsInternalExtension>true</IsInternalExtension>
  <Eula>{{.Eula}}</Eula>
  <PrivacyUri>{{.Privacy}}</PrivacyUri>
  <HomepageUri>{{.Homepage}}</HomepageUri>
  <IsJsonExtension>true</IsJsonExtension>
  <CompanyName>{{.Company}}</CompanyName>
  <SupportedOS>{{.OS}}</SupportedOS>
  <!--%REGIONS%-->
</ExtensionImage>
`
	tpl, err := template.New("manifest").Parse(manifestXml)
	if err != nil {
		log.Fatalf("template parse error: %v", err)
	}
	if err = tpl.Execute(os.Stdout, p); err != nil {
		log.Fatalf("template execute error: %v", err)
	}
}
