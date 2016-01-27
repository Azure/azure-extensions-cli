package main

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/management"
	log "github.com/Sirupsen/logrus"
)

const (
	operationStatusPollingInterval = time.Second * 10
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

type ReplicationStatusResponse struct {
	XMLName  xml.Name `xml:"ReplicationStatusList"`
	Statuses []struct {
		Location string `xml:"Location"`
		Status   string `xml:"Status"`
	} `xml:"ReplicationStatus"`
}

func (c ExtensionsClient) GetReplicationStatus(publisherNamespace, extension, version string) (ReplicationStatusResponse, error) {
	var l ReplicationStatusResponse

	response, err := c.client.SendAzureGetRequest(fmt.Sprintf("services/extensions/%s/%s/%s/replicationstatus",
		publisherNamespace, extension, version))
	if err != nil {
		return l, err
	}

	err = xml.Unmarshal(response, &l)
	return l, err
}

// UpdateExtension sends the given extension handler definition XML to issue and update
// request. Returned operation ID should be polled.
func (c ExtensionsClient) UpdateExtension(data []byte) (management.OperationID, error) {
	return c.client.SendAzurePutRequest("services/extensions?action=update", "text/xml", data)
}

func (c ExtensionsClient) WaitForOperation(opID management.OperationID) error {
	lg := log.WithField("x-ms-operation-id", opID)
	lg.Debug("Waiting for operation to complete.")
	for {
		op, err := c.client.GetOperationStatus(opID)
		if err != nil {
			log.Errorf("Error fetching operation status: %v", err)
			continue // don't return because of GetOperationStatus flakiness.
		}

		switch op.Status {
		case management.OperationStatusSucceeded:
			lg.Debug("Operation successful.")
			return nil
		case management.OperationStatusFailed:
			lg.Debug("Operation failed.")
			if op.Error != nil {
				return op.Error
			}
			return fmt.Errorf("Azure Operation (x-ms-request-id=%s) has failed", opID)
		case management.OperationStatusInProgress:
			lg.Debug("Operation in progress...")
			time.Sleep(operationStatusPollingInterval)
			continue
		default:
			lg.Errorf("Encoutered unhandled operation status: %v", op.Status)
			return fmt.Errorf("Unhandled operation status returned from API: %s", op.Status)
		}
	}
}
