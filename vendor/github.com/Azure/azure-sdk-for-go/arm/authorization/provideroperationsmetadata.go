package authorization

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
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ProviderOperationsMetadataClient is the role based access control provides
// you a way to apply granular level policy administration down to individual
// resources or resource groups. These operations enable you to manage role
// definitions and role assignments. A role definition describes the set of
// actions that can be performed on resources. A role assignment grants access
// to Azure Active Directory users.
type ProviderOperationsMetadataClient struct {
	ManagementClient
}

// NewProviderOperationsMetadataClient creates an instance of the
// ProviderOperationsMetadataClient client.
func NewProviderOperationsMetadataClient(subscriptionID string) ProviderOperationsMetadataClient {
	return NewProviderOperationsMetadataClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProviderOperationsMetadataClientWithBaseURI creates an instance of the
// ProviderOperationsMetadataClient client.
func NewProviderOperationsMetadataClientWithBaseURI(baseURI string, subscriptionID string) ProviderOperationsMetadataClient {
	return ProviderOperationsMetadataClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets provider operations metadata for the specified resource provider.
//
// resourceProviderNamespace is the namespace of the resource provider. expand
// is specifies whether to expand the values.
func (client ProviderOperationsMetadataClient) Get(resourceProviderNamespace string, expand string) (result ProviderOperationsMetadata, err error) {
	req, err := client.GetPreparer(resourceProviderNamespace, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.ProviderOperationsMetadataClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.ProviderOperationsMetadataClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.ProviderOperationsMetadataClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProviderOperationsMetadataClient) GetPreparer(resourceProviderNamespace string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
	}

	const APIVersion = "2015-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Authorization/providerOperations/{resourceProviderNamespace}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderOperationsMetadataClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProviderOperationsMetadataClient) GetResponder(resp *http.Response) (result ProviderOperationsMetadata, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets provider operations metadata for all resource providers.
//
// expand is specifies whether to expand the values.
func (client ProviderOperationsMetadataClient) List(expand string) (result ProviderOperationsMetadataListResult, err error) {
	req, err := client.ListPreparer(expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.ProviderOperationsMetadataClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.ProviderOperationsMetadataClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.ProviderOperationsMetadataClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ProviderOperationsMetadataClient) ListPreparer(expand string) (*http.Request, error) {
	const APIVersion = "2015-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Authorization/providerOperations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderOperationsMetadataClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ProviderOperationsMetadataClient) ListResponder(resp *http.Response) (result ProviderOperationsMetadataListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client ProviderOperationsMetadataClient) ListNextResults(lastResults ProviderOperationsMetadataListResult) (result ProviderOperationsMetadataListResult, err error) {
	req, err := lastResults.ProviderOperationsMetadataListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ProviderOperationsMetadataClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization.ProviderOperationsMetadataClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.ProviderOperationsMetadataClient", "List", resp, "Failure responding to next results request")
	}

	return
}
