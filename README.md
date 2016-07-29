# azure-extensions-cli

:warning::warning:  This tool is intended only for publishers of Azure VM 
extensions. If you are not an Azure employee or a whitelisted extension
publisher, there is nothing here for you.

This command line interface is a simple wrapper on top of RDFE Extension
Publishing APIs.

With this command-line interface you can:

- Upload an extension .zip package as a new extension or a new version.
- Promote versions to required rollout slices.
- List all extension versions published.
- Get replication status of an extension version.
- Mark a version as internal and delete a version.

## Usage

Make sure you have:

- a subscription ID that has extension publishing enabled
- and a .pem file used as the subscription management certificate

Instead of passing these arguments over and over to the CLI,
you can simply set environment variables:

    export SUBSCRIPTION_ID=xxxx-xxxxx-xxxxxx...
    export SUBSCRIPTION_CERT=/path/to/cert.pem

If you are always operating on the same extension, you can also set:

    export EXTENSION_NAMESPACE=Microsoft.Azure.Extensions
    export EXTENSION_NAME=FooExtension

Then use help to explore the commands and arguments.

    ./azure-extensions-cli --help

> **NOTE:** If you are not familiar with extension publishing
process (i.e. slices, behaviors of extension pipeline) you should read
the relevant documentation first.

## CLI

May be up-to-date, but to give an idea:

```
$./azure-extensions-cli
NAME:
   azure-extensions-cli - This tool is designed for Microsoft internal extension publishers to release,
   update and manage Virtual Machine extensions.

USAGE:
   azure-extensions-cli [global options] command [command options] [arguments...]

COMMANDS:
   new-extension-manifest	Creates an XML file used to publish or update extension.
   new-extension		Creates a new type of extension, not for releasing new versions.
   new-extension-version	Publishes a new type of extension internally.
   promote-single-region	Promote published internal extension to a PROD Location.
   promote-two-regions		Promote published extension to two PROD Locations.
   promote-to-prod		Promote published extension to all PROD Locations.
   list-versions		Lists all published extension versions for subscription
   replication-status		Retrieves replication status for an uploaded extension package
   unpublish-version		Marks the specified version of the extension internal. Does not delete.
   delete-version		Deletes the extension version. It should be unpublished first.
   help, h			Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version 
```


## Building

This project is written in Go, make sure you have Go **1.5+** installed
and run:

     export GO15VENDOREXPERIMENT=1
     go build

To cross compile:
    
     go get github.com/mitchellh/gox
     export PATH=$PATH:$GOPATH/bin
     gox -arch="amd64" -os="windows linux darwin" -output "bin/{{.OS}}-{{.Arch}}_{{.Dir}}"

## Author

Ahmet Alp Balkan

## TODO 

- [ ] make `replication-status` exit with appropriate code if replication is not completed.
- [ ] make `replication-status` `--wait` arg to poll until replication completes.
- [ ] add `replication-status --json` flag to output for a programmable output.

## License

See [LICENSE](LICENSE).


-----
This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/). For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.
