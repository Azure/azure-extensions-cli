package storsimple8000series

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.1.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// BackupsClient is the client for the Backups methods of the
// Storsimple8000series service.
type BackupsClient struct {
	ManagementClient
}

// NewBackupsClient creates an instance of the BackupsClient client.
func NewBackupsClient(subscriptionID string) BackupsClient {
	return NewBackupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBackupsClientWithBaseURI creates an instance of the BackupsClient client.
func NewBackupsClientWithBaseURI(baseURI string, subscriptionID string) BackupsClient {
	return BackupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Clone clones the backup element as a new volume. This method may poll for
// completion. Polling can be canceled by passing the cancel channel argument.
// The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// deviceName is the device name backupName is the backup name.
// backupElementName is the backup element name. parameters is the clone
// request object. resourceGroupName is the resource group name managerName is
// the manager name
func (client BackupsClient) Clone(deviceName string, backupName string, backupElementName string, parameters CloneRequest, resourceGroupName string, managerName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.TargetDeviceID", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.TargetVolumeName", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.TargetAccessControlRecordIds", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.BackupElement", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.BackupElement.ElementID", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.BackupElement.ElementName", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.BackupElement.ElementType", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.BackupElement.SizeInBytes", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.BackupElement.VolumeName", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.BackupElement.VolumeContainerID", Name: validation.Null, Rule: true, Chain: nil},
					}}}},
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "storsimple8000series.BackupsClient", "Clone")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.ClonePreparer(deviceName, backupName, backupElementName, parameters, resourceGroupName, managerName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "storsimple8000series.BackupsClient", "Clone", nil, "Failure preparing request")
			return
		}

		resp, err := client.CloneSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "storsimple8000series.BackupsClient", "Clone", resp, "Failure sending request")
			return
		}

		result, err = client.CloneResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "storsimple8000series.BackupsClient", "Clone", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// ClonePreparer prepares the Clone request.
func (client BackupsClient) ClonePreparer(deviceName string, backupName string, backupElementName string, parameters CloneRequest, resourceGroupName string, managerName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupElementName": backupElementName,
		"backupName":        backupName,
		"deviceName":        deviceName,
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backups/{backupName}/elements/{backupElementName}/clone", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CloneSender sends the Clone request. The method will close the
// http.Response Body if it receives an error.
func (client BackupsClient) CloneSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CloneResponder handles the response to the Clone request. The method always
// closes the http.Response Body.
func (client BackupsClient) CloneResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete deletes the backup. This method may poll for completion. Polling can
// be canceled by passing the cancel channel argument. The channel will be used
// to cancel polling and any outstanding HTTP requests.
//
// deviceName is the device name backupName is the backup name.
// resourceGroupName is the resource group name managerName is the manager name
func (client BackupsClient) Delete(deviceName string, backupName string, resourceGroupName string, managerName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "storsimple8000series.BackupsClient", "Delete")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeletePreparer(deviceName, backupName, resourceGroupName, managerName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "storsimple8000series.BackupsClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "storsimple8000series.BackupsClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "storsimple8000series.BackupsClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client BackupsClient) DeletePreparer(deviceName string, backupName string, resourceGroupName string, managerName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupName":        backupName,
		"deviceName":        deviceName,
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backups/{backupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client BackupsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client BackupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByDevice retrieves all the backups in a device.
//
// deviceName is the device name resourceGroupName is the resource group name
// managerName is the manager name filter is oData Filter options
func (client BackupsClient) ListByDevice(deviceName string, resourceGroupName string, managerName string, filter string) (result BackupList, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "storsimple8000series.BackupsClient", "ListByDevice")
	}

	req, err := client.ListByDevicePreparer(deviceName, resourceGroupName, managerName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple8000series.BackupsClient", "ListByDevice", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDeviceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple8000series.BackupsClient", "ListByDevice", resp, "Failure sending request")
		return
	}

	result, err = client.ListByDeviceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple8000series.BackupsClient", "ListByDevice", resp, "Failure responding to request")
	}

	return
}

// ListByDevicePreparer prepares the ListByDevice request.
func (client BackupsClient) ListByDevicePreparer(deviceName string, resourceGroupName string, managerName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        deviceName,
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByDeviceSender sends the ListByDevice request. The method will close the
// http.Response Body if it receives an error.
func (client BackupsClient) ListByDeviceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByDeviceResponder handles the response to the ListByDevice request. The method always
// closes the http.Response Body.
func (client BackupsClient) ListByDeviceResponder(resp *http.Response) (result BackupList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDeviceNextResults retrieves the next set of results, if any.
func (client BackupsClient) ListByDeviceNextResults(lastResults BackupList) (result BackupList, err error) {
	req, err := lastResults.BackupListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storsimple8000series.BackupsClient", "ListByDevice", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByDeviceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "storsimple8000series.BackupsClient", "ListByDevice", resp, "Failure sending next results request")
	}

	result, err = client.ListByDeviceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple8000series.BackupsClient", "ListByDevice", resp, "Failure responding to next results request")
	}

	return
}

// ListByDeviceComplete gets all elements from the list without paging.
func (client BackupsClient) ListByDeviceComplete(deviceName string, resourceGroupName string, managerName string, filter string, cancel <-chan struct{}) (<-chan Backup, <-chan error) {
	resultChan := make(chan Backup)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByDevice(deviceName, resourceGroupName, managerName, filter)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListByDeviceNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// Restore restores the backup on the device. This method may poll for
// completion. Polling can be canceled by passing the cancel channel argument.
// The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// deviceName is the device name backupName is the backupSet name
// resourceGroupName is the resource group name managerName is the manager name
func (client BackupsClient) Restore(deviceName string, backupName string, resourceGroupName string, managerName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "storsimple8000series.BackupsClient", "Restore")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.RestorePreparer(deviceName, backupName, resourceGroupName, managerName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "storsimple8000series.BackupsClient", "Restore", nil, "Failure preparing request")
			return
		}

		resp, err := client.RestoreSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "storsimple8000series.BackupsClient", "Restore", resp, "Failure sending request")
			return
		}

		result, err = client.RestoreResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "storsimple8000series.BackupsClient", "Restore", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// RestorePreparer prepares the Restore request.
func (client BackupsClient) RestorePreparer(deviceName string, backupName string, resourceGroupName string, managerName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupName":        backupName,
		"deviceName":        deviceName,
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backups/{backupName}/restore", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// RestoreSender sends the Restore request. The method will close the
// http.Response Body if it receives an error.
func (client BackupsClient) RestoreSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// RestoreResponder handles the response to the Restore request. The method always
// closes the http.Response Body.
func (client BackupsClient) RestoreResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
