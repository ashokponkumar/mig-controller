package securityinsight

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// CaseRelationsClient is the API spec for Microsoft.SecurityInsights (Azure Security Insights) resource provider
type CaseRelationsClient struct {
	BaseClient
}

// NewCaseRelationsClient creates an instance of the CaseRelationsClient client.
func NewCaseRelationsClient(subscriptionID string) CaseRelationsClient {
	return NewCaseRelationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCaseRelationsClientWithBaseURI creates an instance of the CaseRelationsClient client.
func NewCaseRelationsClientWithBaseURI(baseURI string, subscriptionID string) CaseRelationsClient {
	return CaseRelationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdateRelation creates or updates the case relation.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// caseID - case ID
// relationName - relation Name
// relationInputModel - the relation input model
func (client CaseRelationsClient) CreateOrUpdateRelation(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, relationName string, relationInputModel RelationsModelInput) (result CaseRelation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CaseRelationsClient.CreateOrUpdateRelation")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.CaseRelationsClient", "CreateOrUpdateRelation", err.Error())
	}

	req, err := client.CreateOrUpdateRelationPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, caseID, relationName, relationInputModel)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CaseRelationsClient", "CreateOrUpdateRelation", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateRelationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.CaseRelationsClient", "CreateOrUpdateRelation", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateRelationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CaseRelationsClient", "CreateOrUpdateRelation", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateRelationPreparer prepares the CreateOrUpdateRelation request.
func (client CaseRelationsClient) CreateOrUpdateRelationPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, relationName string, relationInputModel RelationsModelInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"caseId":                              autorest.Encode("path", caseID),
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"relationName":                        autorest.Encode("path", relationName),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/cases/{caseId}/relations/{relationName}", pathParameters),
		autorest.WithJSON(relationInputModel),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateRelationSender sends the CreateOrUpdateRelation request. The method will close the
// http.Response Body if it receives an error.
func (client CaseRelationsClient) CreateOrUpdateRelationSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateRelationResponder handles the response to the CreateOrUpdateRelation request. The method always
// closes the http.Response Body.
func (client CaseRelationsClient) CreateOrUpdateRelationResponder(resp *http.Response) (result CaseRelation, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteRelation delete the case relation.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// caseID - case ID
// relationName - relation Name
func (client CaseRelationsClient) DeleteRelation(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, relationName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CaseRelationsClient.DeleteRelation")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.CaseRelationsClient", "DeleteRelation", err.Error())
	}

	req, err := client.DeleteRelationPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, caseID, relationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CaseRelationsClient", "DeleteRelation", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteRelationSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "securityinsight.CaseRelationsClient", "DeleteRelation", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteRelationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CaseRelationsClient", "DeleteRelation", resp, "Failure responding to request")
	}

	return
}

// DeleteRelationPreparer prepares the DeleteRelation request.
func (client CaseRelationsClient) DeleteRelationPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, relationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"caseId":                              autorest.Encode("path", caseID),
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"relationName":                        autorest.Encode("path", relationName),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/cases/{caseId}/relations/{relationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteRelationSender sends the DeleteRelation request. The method will close the
// http.Response Body if it receives an error.
func (client CaseRelationsClient) DeleteRelationSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteRelationResponder handles the response to the DeleteRelation request. The method always
// closes the http.Response Body.
func (client CaseRelationsClient) DeleteRelationResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetRelation gets a case relation.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// caseID - case ID
// relationName - relation Name
func (client CaseRelationsClient) GetRelation(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, relationName string) (result CaseRelation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CaseRelationsClient.GetRelation")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.CaseRelationsClient", "GetRelation", err.Error())
	}

	req, err := client.GetRelationPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, caseID, relationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CaseRelationsClient", "GetRelation", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetRelationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.CaseRelationsClient", "GetRelation", resp, "Failure sending request")
		return
	}

	result, err = client.GetRelationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CaseRelationsClient", "GetRelation", resp, "Failure responding to request")
	}

	return
}

// GetRelationPreparer prepares the GetRelation request.
func (client CaseRelationsClient) GetRelationPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, relationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"caseId":                              autorest.Encode("path", caseID),
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"relationName":                        autorest.Encode("path", relationName),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/cases/{caseId}/relations/{relationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetRelationSender sends the GetRelation request. The method will close the
// http.Response Body if it receives an error.
func (client CaseRelationsClient) GetRelationSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetRelationResponder handles the response to the GetRelation request. The method always
// closes the http.Response Body.
func (client CaseRelationsClient) GetRelationResponder(resp *http.Response) (result CaseRelation, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all case relations.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// caseID - case ID
// filter - filters the results, based on a Boolean condition. Optional.
// orderby - sorts the results. Optional.
// top - returns only the first n results. Optional.
// skipToken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls. Optional.
func (client CaseRelationsClient) List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, filter string, orderby string, top *int32, skipToken string) (result CaseRelationListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CaseRelationsClient.List")
		defer func() {
			sc := -1
			if result.crl.Response.Response != nil {
				sc = result.crl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.CaseRelationsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, caseID, filter, orderby, top, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CaseRelationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.crl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.CaseRelationsClient", "List", resp, "Failure sending request")
		return
	}

	result.crl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CaseRelationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client CaseRelationsClient) ListPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, filter string, orderby string, top *int32, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"caseId":                              autorest.Encode("path", caseID),
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/cases/{caseId}/relations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client CaseRelationsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client CaseRelationsClient) ListResponder(resp *http.Response) (result CaseRelationList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client CaseRelationsClient) listNextResults(ctx context.Context, lastResults CaseRelationList) (result CaseRelationList, err error) {
	req, err := lastResults.caseRelationListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "securityinsight.CaseRelationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "securityinsight.CaseRelationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CaseRelationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client CaseRelationsClient) ListComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, filter string, orderby string, top *int32, skipToken string) (result CaseRelationListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CaseRelationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, caseID, filter, orderby, top, skipToken)
	return
}
