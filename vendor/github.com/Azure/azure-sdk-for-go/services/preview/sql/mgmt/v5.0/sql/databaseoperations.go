package sql

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"github.com/gofrs/uuid"
	"net/http"
)

// DatabaseOperationsClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type DatabaseOperationsClient struct {
	BaseClient
}

// NewDatabaseOperationsClient creates an instance of the DatabaseOperationsClient client.
func NewDatabaseOperationsClient(subscriptionID string) DatabaseOperationsClient {
	return NewDatabaseOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDatabaseOperationsClientWithBaseURI creates an instance of the DatabaseOperationsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewDatabaseOperationsClientWithBaseURI(baseURI string, subscriptionID string) DatabaseOperationsClient {
	return DatabaseOperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Cancel cancels the asynchronous operation on the database.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
// operationID - the operation identifier.
func (client DatabaseOperationsClient) Cancel(ctx context.Context, resourceGroupName string, serverName string, databaseName string, operationID uuid.UUID) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabaseOperationsClient.Cancel")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CancelPreparer(ctx, resourceGroupName, serverName, databaseName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseOperationsClient", "Cancel", nil, "Failure preparing request")
		return
	}

	resp, err := client.CancelSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sql.DatabaseOperationsClient", "Cancel", resp, "Failure sending request")
		return
	}

	result, err = client.CancelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseOperationsClient", "Cancel", resp, "Failure responding to request")
		return
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client DatabaseOperationsClient) CancelPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, operationID uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"operationId":       autorest.Encode("path", operationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/operations/{operationId}/cancel", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseOperationsClient) CancelSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client DatabaseOperationsClient) CancelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByDatabase gets a list of operations performed on the database.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
func (client DatabaseOperationsClient) ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result DatabaseOperationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabaseOperationsClient.ListByDatabase")
		defer func() {
			sc := -1
			if result.dolr.Response.Response != nil {
				sc = result.dolr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByDatabaseNextResults
	req, err := client.ListByDatabasePreparer(ctx, resourceGroupName, serverName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseOperationsClient", "ListByDatabase", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDatabaseSender(req)
	if err != nil {
		result.dolr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.DatabaseOperationsClient", "ListByDatabase", resp, "Failure sending request")
		return
	}

	result.dolr, err = client.ListByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseOperationsClient", "ListByDatabase", resp, "Failure responding to request")
		return
	}
	if result.dolr.hasNextLink() && result.dolr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByDatabasePreparer prepares the ListByDatabase request.
func (client DatabaseOperationsClient) ListByDatabasePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/operations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDatabaseSender sends the ListByDatabase request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseOperationsClient) ListByDatabaseSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDatabaseResponder handles the response to the ListByDatabase request. The method always
// closes the http.Response Body.
func (client DatabaseOperationsClient) ListByDatabaseResponder(resp *http.Response) (result DatabaseOperationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByDatabaseNextResults retrieves the next set of results, if any.
func (client DatabaseOperationsClient) listByDatabaseNextResults(ctx context.Context, lastResults DatabaseOperationListResult) (result DatabaseOperationListResult, err error) {
	req, err := lastResults.databaseOperationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.DatabaseOperationsClient", "listByDatabaseNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByDatabaseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.DatabaseOperationsClient", "listByDatabaseNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseOperationsClient", "listByDatabaseNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByDatabaseComplete enumerates all values, automatically crossing page boundaries as required.
func (client DatabaseOperationsClient) ListByDatabaseComplete(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result DatabaseOperationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabaseOperationsClient.ListByDatabase")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByDatabase(ctx, resourceGroupName, serverName, databaseName)
	return
}
