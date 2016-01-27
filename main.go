package main

import (
	"os"
	"text/template"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stderr)
}

func main() {
	app := cli.NewApp()
	app.Name = "azure-extensions-cli"
	app.Usage = "This tool is designed for Microsoft internal extension publishers to release, update and manage Virtual Machine extensions."
	app.Authors = []cli.Author{{Name: "Ahmet Alp Balkan", Email: "ahmetb at microsoft döt com"}}
	app.Commands = []cli.Command{
		{Name: "new-extension-manifest",
			Usage:  "Creates an XML file used to publish or update extension.",
			Action: newExtensionManifest,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "namespace",
					Usage: "Publisher namespace e.g. Microsoft.Azure.Extensions"},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the extension e.g. FooExtension"},
				cli.StringFlag{
					Name:  "version",
					Usage: "Version of the extension package e.g. 1.0.0"},
				cli.StringFlag{
					Name:  "label",
					Usage: "Human readable name of the extension"},
				cli.StringFlag{
					Name:  "description",
					Usage: "Description of the extension"},
				cli.StringFlag{
					Name:  "eula-url",
					Usage: "URL to the End-User License Agreement page"},
				cli.StringFlag{
					Name:  "privacy-url",
					Usage: "URL to the Privacy Policy page"},
				cli.StringFlag{
					Name:  "homepage-url",
					Usage: "URL to the homepage of the extension"},
				cli.StringFlag{
					Name:  "company",
					Usage: "Human-readable Company Name of the publisher"},
				cli.StringFlag{
					Name:  "supported-os",
					Usage: "Extension platform e.g. 'Linux'"},
			},
		},
	}
	app.RunAndExitOnError()
}

func newExtensionManifest(c *cli.Context) {
	var p struct {
		Namespace, Name, Version, Label, Description, Eula, Privacy, Homepage, Company, OS string
	}
	flags := []struct {
		ref *string
		fl  string
	}{
		{&p.Namespace, "namespace"},
		{&p.Name, "name"},
		{&p.Version, "version"},
		{&p.Label, "label"},
		{&p.Description, "description"},
		{&p.Eula, "eula-url"},
		{&p.Privacy, "privacy-url"},
		{&p.Homepage, "homepage-url"},
		{&p.Company, "company"},
		{&p.OS, "supported-os"},
	}
	for _, f := range flags {
		if val := c.String(f.fl); val == "" {
			log.Fatalf("argument %s must be provided", f.fl)
		} else {
			*f.ref = val
		}
	}
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
  <MediaLink>%BLOB_URL%</MediaLink>
  <Description>{{.Description}}</Description>
  <IsInternalExtension>true</IsInternalExtension>
  <Eula>{{.Eula}}</Eula>
  <PrivacyUri>{{.Privacy}}</PrivacyUri>
  <HomepageUri>{{.Homepage}}</HomepageUri>
  <IsJsonExtension>true</IsJsonExtension>
  <CompanyName>{{.Company}}</CompanyName>
  <SupportedOS>{{.OS}}</SupportedOS>
  <!--%REGIONS%-->
</ExtensionImage>`
	tpl, err := template.New("manifest").Parse(manifestXml)
	if err != nil {
		log.Fatalf("template parse error: %v", err)
	}
	if err = tpl.Execute(os.Stdout, p); err != nil {
		log.Fatalf("template execute error: %v", err)
	}
}
