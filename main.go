package main

import (
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

const (
	version = "1.0.0-beta1"
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stderr)
}

var (
	flManifest = cli.StringFlag{
		Name:  "manifest",
		Usage: "Path of XML manifest file (output of 'new-extension-manifest')"}
	flSubsID = cli.StringFlag{
		Name:  "subscription-id",
		Usage: "Subscription ID for the publisher subscription"}
	flSubsCert = cli.StringFlag{
		Name:  "subscription-cert",
		Usage: "Path of subscription management certificate (.pem) file"}
	flVersion = cli.StringFlag{
		Name:  "version",
		Usage: "Version of the extension package e.g. 1.0.0"}
	flNamespace = cli.StringFlag{
		Name:  "namespace",
		Usage: "Publisher namespace e.g. Microsoft.Azure.Extensions"}
	flName = cli.StringFlag{
		Name:  "name",
		Usage: "Name of the extension e.g. FooExtension"}
	flStorageAccount = cli.StringFlag{
		Name:  "storage-account",
		Usage: "Name of an existing storage account to be used in uploading the extension package temporarily."}
)

func main() {
	app := cli.NewApp()
	app.Name = "azure-extensions-cli"
	app.Version = version
	app.Usage = "This tool is designed for Microsoft internal extension publishers to release, update and manage Virtual Machine extensions."
	app.Authors = []cli.Author{{Name: "Ahmet Alp Balkan", Email: "ahmetb at microsoft döt com"}}
	app.Commands = []cli.Command{
		{Name: "new-extension-manifest",
			Usage:  "Creates an XML file used to publish or update extension.",
			Action: newExtensionManifest,
			Flags: []cli.Flag{
				flNamespace,
				flName,
				flVersion,
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
			}},
		{Name: "new-extension",
			Usage: "Creates a new type of extension, not for releasing new versions.",
			Flags: []cli.Flag{
				flSubsID,
				flSubsCert,
				flManifest,
				flStorageAccount},
			Action: createExtension},
		{Name: "list-versions",
			Usage:  "Lists all published extension versions for subscription",
			Flags:  []cli.Flag{flSubsID, flSubsCert},
			Action: listVersions},
		{Name: "replication-status",
			Usage:  "Retrieves replication status for an uploaded extension package",
			Flags:  []cli.Flag{flSubsID, flSubsCert, flNamespace, flName, flVersion},
			Action: replicationStatus},
		{Name: "unpublish-version",
			Usage:  "Marks the specified version of the extension internal. Does not delete.",
			Flags:  []cli.Flag{flSubsID, flSubsCert, flNamespace, flName, flVersion},
			Action: unpublishVersion},
		{Name: "delete-version",
			Usage:  "Deletes the extension version. It should be unpublished first.",
			Flags:  []cli.Flag{flSubsID, flSubsCert, flNamespace, flName, flVersion},
			Action: deleteVersion},
	}
	app.RunAndExitOnError()
}

func mkClient(subscriptionID, certFile string) ExtensionsClient {
	b, err := ioutil.ReadFile(certFile)
	if err != nil {
		log.Fatal("Cannot read certificate %s: %v", certFile, err)
	}
	cl, err := NewClient(subscriptionID, b)
	if err != nil {
		log.Fatal("Cannot create client: %v", err)
	}
	return cl
}

func checkFlag(c *cli.Context, fl string) string {
	v := c.String(fl)
	if v == "" {
		log.Fatalf("argument %s must be provided", fl)
	}
	return v
}
