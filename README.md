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
    export MANAGEMENT_URL=https://management.core.windows.net

If you are always operating on the same extension, you can also set:

    export EXTENSION_NAMESPACE=Microsoft.Azure.Extensions
    export EXTENSION_NAME=FooExtension

Then use help to explore the commands and arguments.

    ./azure-extensions-cli --help

> **NOTE:** If you are not familiar with extension publishing
process (i.e. slices, behaviors of extension pipeline) you should read
the relevant documentation first.

## CLI

This might not be up-to-date, but to give an idea, here are the subcommands

```
$./azure-extensions-cli
NAME:
   azure-extensions-cli - This tool is designed for Microsoft internal extension publishers
    to release, update and manage Virtual Machine extensions.

USAGE:
   azure-extensions-cli [global options] command [command options] [arguments...]

COMMANDS:
   new-extension-manifest   Creates an XML file used to publish or update extension.
   new-extension		    Creates a new type of extension, not for releasing new versions.
   new-extension-version    Publishes a new type of extension internally.
   promote                  Promote published internal extension to one or more PROD Locations.
   promote-all-regions      Promote published extension to all PROD Locations.
   list-versions		    Lists all published extension versions for subscription
   replication-status		Retrieves replication status for an uploaded extension package
   unpublish-version		Marks the specified version of the extension internal. Does not delete.
   delete-version		    Deletes the extension version. It should be unpublished first.
   help, h	                Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version 
```

## Installing (or building from source)

You can head over to the **Releases** section to download a binary built for various platforms.

If you need to compile from the source code, make sure you have Go compiler 1.6+ installed.
Check out the project, set the GOPATH environment variable correctly (if necessary) and
run `go build`. This should compile a binary.

## Overview

The CLI makes it easy (easier) to publish an Azure extension.  An example workflow is provided below. This workflow 
assumes an extension type already exists, which is why the command **new-extension-version** is used.  (If the type does 
not exist use substitute for the command new-extension instead.)

Not all command line parameters are shown for each command, only the salient options are shown.

Step 1 - create an extension manifest.

 1. ./azure-extensions-cli new-extension-manifest

Step 2 - publish an extension internally.

 1. ./azure-extensions-cli new-extension-version
 
Step 3 - rollout the extension to Azure, by slowly including more and more regions.  It is recommended that you pause
24 hours between regions.  

> Every time a new region is added, the previous regions must be included with the promote command.
 
 1. ./azure-extensions-cli promote --region "West Central US"
 1. ./azure-extensions-cli promote --region "West Central US" --region "North Central US"
 1. ./azure-extensions-cli promote --region "West Central US" --region "North Central US" --region "West US"
 1. ./azure-extensions-cli promote ...
 
Step 4 - promote the extension to all Azure regions.

 1. ./azure-extensions-cli promote-all-regions

### Regions

As of 13-Sept-2017 this is the list of regions available for the
public cloud.  azure-extensions-cli uses the Service Management
regions.  (The equivalent Resource Manager (RM or ARM) regions
equivalents are also provided, but are not **yet** supported.)

| Service Management  | Resource Manager   |
|---------------------|--------------------|
| Australia East      | australiaeast      |
| Australia Southeast | australiasoutheast |
| Brazil South        | brazilsouth        |
| Canada Central      | canadacentral      |
| Canada East         | canadaeast         |
| Central India       | centralindia       |
| Central US          | centralus          |
| Central US EUAP     | centraluseuap      |
| East Asia           | eastasia           |
| East US             | eastus             |
| East US 2           | eastus2            |
| East US 2 EUAP      | eastus2euap        |
| Japan East          | japaneast          |
| Japan West          | japanwest          |
| Korea Central       | koreacentral       |
| Korea South         | koreasouth         |
| North Central US    | northcentralus     |
| North Europe        | northeurope        |
| South Central US    | southcentralus     |
| South India         | southindia         |
| Southeast Asia      | southeastasia      |
| UK South            | uksouth            |
| UK West             | ukwest             |
| West Central US     | westcentra         |
| West Europe         | westeurope         |
| West India          | westindia          |
| West US             | westus             |
| West US 2           | westus2            |

  
## TODO 

- [ ] make `replication-status` exit with appropriate code if replication is not completed.
- [ ] make `replication-status` `--wait` arg to poll until replication completes.
- [x] add `replication-status --json` flag to output for a programmable output.

## License

See [LICENSE](LICENSE).


-----
This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/). For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.
