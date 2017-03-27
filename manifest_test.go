package main

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/approvals/go-approval-tests"
)

func TestRoundTripExtensionImage(t *testing.T) {
	xmlString := []byte(`<ExtensionImage xmlns="http://schemas.microsoft.com/windowsazure">
  <ProviderNameSpace>Microsoft.OSCTExtensions</ProviderNameSpace>
  <Type>CustomScriptForLinux</Type>
  <Version>4.3.2.1</Version>
  <Label>Microsoft Azure Custom Script Extension for Linux Virtual Machines</Label>
  <HostingResources>VmRole</HostingResources>
  <MediaLink>http://localhost/extension.zip</MediaLink>
  <Description>Please consider using Microsoft.Azure.Extensions.CustomScript instead.</Description>
  <IsInternalExtension>true</IsInternalExtension>
  <Eula>https://github.com/Azure/azure-linux-extensions/blob/master/LICENSE-2_0.txt</Eula>
  <PrivacyUri>http://www.microsoft.com/privacystatement/en-us/OnlineServices/Default.aspx</PrivacyUri>
  <HomepageUri>https://github.com/Azure/azure-linux-extensions</HomepageUri>
  <IsJsonExtension>true</IsJsonExtension>
  <CompanyName>Microsoft</CompanyName>
  <SupportedOS>Linux</SupportedOS>
  <Regions>South Central US</Regions>
</ExtensionImage>`)

	var obj extensionImage
	err := xml.Unmarshal(xmlString, &obj)
	if err != nil {
		t.Fatal(err)
	}

	bs, err := xml.MarshalIndent(obj, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	err = approvaltests.Verify(t, bytes.NewReader(bs))
	if err != nil {
		t.Fatal(err)
	}
}

func TestSuppressRegionWhenEmpty(t *testing.T) {
	xmlString := []byte(`<ExtensionImage xmlns="http://schemas.microsoft.com/windowsazure">
  <ProviderNameSpace>Microsoft.OSCTExtensions</ProviderNameSpace>
  <Type>CustomScriptForLinux</Type>
  <Version>4.3.2.1</Version>
  <Label>Microsoft Azure Custom Script Extension for Linux Virtual Machines</Label>
  <HostingResources>VmRole</HostingResources>
  <MediaLink>http://localhost/extension.zip</MediaLink>
  <Description>Please consider using Microsoft.Azure.Extensions.CustomScript instead.</Description>
  <IsInternalExtension>true</IsInternalExtension>
  <Eula>https://github.com/Azure/azure-linux-extensions/blob/master/LICENSE-2_0.txt</Eula>
  <PrivacyUri>http://www.microsoft.com/privacystatement/en-us/OnlineServices/Default.aspx</PrivacyUri>
  <HomepageUri>https://github.com/Azure/azure-linux-extensions</HomepageUri>
  <IsJsonExtension>true</IsJsonExtension>
  <CompanyName>Microsoft</CompanyName>
  <SupportedOS>Linux</SupportedOS>
</ExtensionImage>`)

	var obj extensionImage
	err := xml.Unmarshal(xmlString, &obj)
	if err != nil {
		t.Fatal(err)
	}

	bs, err := xml.MarshalIndent(obj, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	err = approvaltests.Verify(t, bytes.NewReader(bs))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRoundTripExtensionImageGlobal(t *testing.T) {
	xmlString := []byte(`<ExtensionImage xmlns="http://schemas.microsoft.com/windowsazure">
  <ProviderNameSpace>Microsoft.OSCTExtensions</ProviderNameSpace>
  <Type>CustomScriptForLinux</Type>
  <Version>4.3.2.1</Version>
  <Label>Microsoft Azure Custom Script Extension for Linux Virtual Machines</Label>
  <HostingResources>VmRole</HostingResources>
  <MediaLink>http://localhost/extension.zip</MediaLink>
  <Description>Please consider using Microsoft.Azure.Extensions.CustomScript instead.</Description>
  <IsInternalExtension>true</IsInternalExtension>
  <Eula>https://github.com/Azure/azure-linux-extensions/blob/master/LICENSE-2_0.txt</Eula>
  <PrivacyUri>http://www.microsoft.com/privacystatement/en-us/OnlineServices/Default.aspx</PrivacyUri>
  <HomepageUri>https://github.com/Azure/azure-linux-extensions</HomepageUri>
  <IsJsonExtension>true</IsJsonExtension>
  <CompanyName>Microsoft</CompanyName>
  <SupportedOS>Linux</SupportedOS>
  <Regions></Regions>
</ExtensionImage>`)

	var obj extensionImageGlobal
	err := xml.Unmarshal(xmlString, &obj)
	if err != nil {
		t.Fatal(err)
	}

	bs, err := xml.MarshalIndent(obj, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	err = approvaltests.Verify(t, bytes.NewReader(bs))
	if err != nil {
		t.Fatal(err)
	}
}
