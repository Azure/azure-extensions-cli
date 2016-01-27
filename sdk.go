package main

import (
	"encoding/xml"

	"github.com/Azure/azure-sdk-for-go/management"
)

type ExtensionsClient struct {
	client management.Client
}

func NewClient(subscriptionID string, cert []byte) (ExtensionsClient, error) {
	cl, err := management.NewClient(subscriptionID, cert)
	return ExtensionsClient{cl}, err
}

type ListVersionsResponse struct {
	XMLName    xml.Name `xml:"ExtensionImages"`
	Extensions []struct {
		Ns                   string `xml:"ProviderNameSpace"`
		Name                 string `xml:"Type"`
		Version              string `xml:Version"`
		ReplicationCompleted bool   `xml:"ReplicationCompleted"`
		Regions              string `xml:"Regions"`
	} `xml:"ExtensionImage"`
}

func (c ExtensionsClient) ListVersions() (ListVersionsResponse, error) {
	var l ListVersionsResponse

	response, err := c.client.SendAzureGetRequest("services/publisherextensions")
	if err != nil {
		return l, err
	}

	err = xml.Unmarshal(response, &l)
	return l, err
}
