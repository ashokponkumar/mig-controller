// Package formrecognizer implements the Azure ARM Formrecognizer service API version 1.0-preview.
//
//
package formrecognizer

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
	"github.com/satori/go.uuid"
	"io"
	"net/http"
)

// BaseClient is the base client for Formrecognizer.
type BaseClient struct {
	autorest.Client
	Endpoint string
}

// New creates an instance of the BaseClient client.
func New(endpoint string) BaseClient {
	return NewWithoutDefaults(endpoint)
}

// NewWithoutDefaults creates an instance of the BaseClient client.
func NewWithoutDefaults(endpoint string) BaseClient {
	return BaseClient{
		Client:   autorest.NewClientWithUserAgent(UserAgent()),
		Endpoint: endpoint,
	}
}

// AnalyzeWithCustomModel extract key-value pairs from a given document. The input document must be of one of the
// supported content types - 'application/pdf', 'image/jpeg' or 'image/png'. A success response is returned in JSON.
// Parameters:
// ID - model Identifier to analyze the document with.
// formStream - a pdf document or image (jpg,png) file to analyze.
// keys - an optional list of known keys to extract the values for.
func (client BaseClient) AnalyzeWithCustomModel(ctx context.Context, ID uuid.UUID, formStream io.ReadCloser, keys []string) (result AnalyzeResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.AnalyzeWithCustomModel")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AnalyzeWithCustomModelPreparer(ctx, ID, formStream, keys)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "AnalyzeWithCustomModel", nil, "Failure preparing request")
		return
	}

	resp, err := client.AnalyzeWithCustomModelSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "AnalyzeWithCustomModel", resp, "Failure sending request")
		return
	}

	result, err = client.AnalyzeWithCustomModelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "AnalyzeWithCustomModel", resp, "Failure responding to request")
	}

	return
}

// AnalyzeWithCustomModelPreparer prepares the AnalyzeWithCustomModel request.
func (client BaseClient) AnalyzeWithCustomModelPreparer(ctx context.Context, ID uuid.UUID, formStream io.ReadCloser, keys []string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"id": autorest.Encode("path", ID),
	}

	queryParameters := map[string]interface{}{}
	if keys != nil && len(keys) > 0 {
		queryParameters["keys"] = autorest.Encode("query", keys, ",")
	}

	formDataParameters := map[string]interface{}{
		"form_stream": formStream,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/formrecognizer/v1.0-preview", urlParameters),
		autorest.WithPathParameters("/custom/models/{id}/analyze", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithMultiPartFormData(formDataParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AnalyzeWithCustomModelSender sends the AnalyzeWithCustomModel request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) AnalyzeWithCustomModelSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// AnalyzeWithCustomModelResponder handles the response to the AnalyzeWithCustomModel request. The method always
// closes the http.Response Body.
func (client BaseClient) AnalyzeWithCustomModelResponder(resp *http.Response) (result AnalyzeResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// BatchReadReceipt batch Read Receipt operation. The response contains a field called 'Operation-Location', which
// contains the URL that you must use for your 'Get Read Receipt Result' operation.
// Parameters:
// imageURL - a JSON document with a URL pointing to the image that is to be analyzed.
func (client BaseClient) BatchReadReceipt(ctx context.Context, imageURL ImageURL) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.BatchReadReceipt")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: imageURL,
			Constraints: []validation.Constraint{{Target: "imageURL.URL", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("formrecognizer.BaseClient", "BatchReadReceipt", err.Error())
	}

	req, err := client.BatchReadReceiptPreparer(ctx, imageURL)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "BatchReadReceipt", nil, "Failure preparing request")
		return
	}

	resp, err := client.BatchReadReceiptSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "BatchReadReceipt", resp, "Failure sending request")
		return
	}

	result, err = client.BatchReadReceiptResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "BatchReadReceipt", resp, "Failure responding to request")
	}

	return
}

// BatchReadReceiptPreparer prepares the BatchReadReceipt request.
func (client BaseClient) BatchReadReceiptPreparer(ctx context.Context, imageURL ImageURL) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/formrecognizer/v1.0-preview", urlParameters),
		autorest.WithPath("/prebuilt/receipt/asyncBatchAnalyze"),
		autorest.WithJSON(imageURL))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// BatchReadReceiptSender sends the BatchReadReceipt request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) BatchReadReceiptSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// BatchReadReceiptResponder handles the response to the BatchReadReceipt request. The method always
// closes the http.Response Body.
func (client BaseClient) BatchReadReceiptResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// BatchReadReceiptInStream read Receipt operation. When you use the 'Batch Read Receipt' interface, the response
// contains a field called 'Operation-Location'. The 'Operation-Location' field contains the URL that you must use for
// your 'Get Read Receipt Result' operation.
// Parameters:
// imageParameter - an image stream.
func (client BaseClient) BatchReadReceiptInStream(ctx context.Context, imageParameter io.ReadCloser) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.BatchReadReceiptInStream")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.BatchReadReceiptInStreamPreparer(ctx, imageParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "BatchReadReceiptInStream", nil, "Failure preparing request")
		return
	}

	resp, err := client.BatchReadReceiptInStreamSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "BatchReadReceiptInStream", resp, "Failure sending request")
		return
	}

	result, err = client.BatchReadReceiptInStreamResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "BatchReadReceiptInStream", resp, "Failure responding to request")
	}

	return
}

// BatchReadReceiptInStreamPreparer prepares the BatchReadReceiptInStream request.
func (client BaseClient) BatchReadReceiptInStreamPreparer(ctx context.Context, imageParameter io.ReadCloser) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/octet-stream"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/formrecognizer/v1.0-preview", urlParameters),
		autorest.WithPath("/prebuilt/receipt/asyncBatchAnalyze"),
		autorest.WithFile(imageParameter))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// BatchReadReceiptInStreamSender sends the BatchReadReceiptInStream request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) BatchReadReceiptInStreamSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// BatchReadReceiptInStreamResponder handles the response to the BatchReadReceiptInStream request. The method always
// closes the http.Response Body.
func (client BaseClient) BatchReadReceiptInStreamResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteCustomModel delete model artifacts.
// Parameters:
// ID - the identifier of the model to delete.
func (client BaseClient) DeleteCustomModel(ctx context.Context, ID uuid.UUID) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.DeleteCustomModel")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteCustomModelPreparer(ctx, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "DeleteCustomModel", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteCustomModelSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "DeleteCustomModel", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteCustomModelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "DeleteCustomModel", resp, "Failure responding to request")
	}

	return
}

// DeleteCustomModelPreparer prepares the DeleteCustomModel request.
func (client BaseClient) DeleteCustomModelPreparer(ctx context.Context, ID uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"id": autorest.Encode("path", ID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{Endpoint}/formrecognizer/v1.0-preview", urlParameters),
		autorest.WithPathParameters("/custom/models/{id}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteCustomModelSender sends the DeleteCustomModel request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) DeleteCustomModelSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteCustomModelResponder handles the response to the DeleteCustomModel request. The method always
// closes the http.Response Body.
func (client BaseClient) DeleteCustomModelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetCustomModel get information about a model.
// Parameters:
// ID - model identifier.
func (client BaseClient) GetCustomModel(ctx context.Context, ID uuid.UUID) (result ModelResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetCustomModel")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetCustomModelPreparer(ctx, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "GetCustomModel", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetCustomModelSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "GetCustomModel", resp, "Failure sending request")
		return
	}

	result, err = client.GetCustomModelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "GetCustomModel", resp, "Failure responding to request")
	}

	return
}

// GetCustomModelPreparer prepares the GetCustomModel request.
func (client BaseClient) GetCustomModelPreparer(ctx context.Context, ID uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"id": autorest.Encode("path", ID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/formrecognizer/v1.0-preview", urlParameters),
		autorest.WithPathParameters("/custom/models/{id}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetCustomModelSender sends the GetCustomModel request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetCustomModelSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetCustomModelResponder handles the response to the GetCustomModel request. The method always
// closes the http.Response Body.
func (client BaseClient) GetCustomModelResponder(resp *http.Response) (result ModelResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetCustomModels get information about all trained custom models
func (client BaseClient) GetCustomModels(ctx context.Context) (result ModelsResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetCustomModels")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetCustomModelsPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "GetCustomModels", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetCustomModelsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "GetCustomModels", resp, "Failure sending request")
		return
	}

	result, err = client.GetCustomModelsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "GetCustomModels", resp, "Failure responding to request")
	}

	return
}

// GetCustomModelsPreparer prepares the GetCustomModels request.
func (client BaseClient) GetCustomModelsPreparer(ctx context.Context) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/formrecognizer/v1.0-preview", urlParameters),
		autorest.WithPath("/custom/models"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetCustomModelsSender sends the GetCustomModels request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetCustomModelsSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetCustomModelsResponder handles the response to the GetCustomModels request. The method always
// closes the http.Response Body.
func (client BaseClient) GetCustomModelsResponder(resp *http.Response) (result ModelsResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetExtractedKeys retrieve the keys that were
// extracted during the training of the specified model.
// Parameters:
// ID - model identifier.
func (client BaseClient) GetExtractedKeys(ctx context.Context, ID uuid.UUID) (result KeysResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetExtractedKeys")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetExtractedKeysPreparer(ctx, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "GetExtractedKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetExtractedKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "GetExtractedKeys", resp, "Failure sending request")
		return
	}

	result, err = client.GetExtractedKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "GetExtractedKeys", resp, "Failure responding to request")
	}

	return
}

// GetExtractedKeysPreparer prepares the GetExtractedKeys request.
func (client BaseClient) GetExtractedKeysPreparer(ctx context.Context, ID uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"id": autorest.Encode("path", ID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/formrecognizer/v1.0-preview", urlParameters),
		autorest.WithPathParameters("/custom/models/{id}/keys", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetExtractedKeysSender sends the GetExtractedKeys request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetExtractedKeysSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetExtractedKeysResponder handles the response to the GetExtractedKeys request. The method always
// closes the http.Response Body.
func (client BaseClient) GetExtractedKeysResponder(resp *http.Response) (result KeysResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetReadReceiptResult this interface is used for getting the analysis results of a 'Batch Read Receipt' operation.
// The URL to this interface should be retrieved from the 'Operation-Location' field returned from the 'Batch Read
// Receipt' operation.
// Parameters:
// operationID - id of read operation returned in the response of a 'Batch Read Receipt' operation.
func (client BaseClient) GetReadReceiptResult(ctx context.Context, operationID string) (result ReadReceiptResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetReadReceiptResult")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetReadReceiptResultPreparer(ctx, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "GetReadReceiptResult", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetReadReceiptResultSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "GetReadReceiptResult", resp, "Failure sending request")
		return
	}

	result, err = client.GetReadReceiptResultResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "GetReadReceiptResult", resp, "Failure responding to request")
	}

	return
}

// GetReadReceiptResultPreparer prepares the GetReadReceiptResult request.
func (client BaseClient) GetReadReceiptResultPreparer(ctx context.Context, operationID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"operationId": autorest.Encode("path", operationID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/formrecognizer/v1.0-preview", urlParameters),
		autorest.WithPathParameters("/prebuilt/receipt/operations/{operationId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetReadReceiptResultSender sends the GetReadReceiptResult request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetReadReceiptResultSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetReadReceiptResultResponder handles the response to the GetReadReceiptResult request. The method always
// closes the http.Response Body.
func (client BaseClient) GetReadReceiptResultResponder(resp *http.Response) (result ReadReceiptResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// TrainCustomModel create and train a custom model. The train request must include a source parameter that is either
// an externally accessible Azure Storage blob container Uri (preferably a Shared Access Signature Uri) or valid path
// to a data folder in a locally mounted drive. When local paths are specified, they must follow the Linux/Unix path
// format and be an absolute path rooted to the input mount configuration
// setting value e.g., if '{Mounts:Input}' configuration setting value is '/input' then a valid source path would be
// '/input/contosodataset'. All data to be trained is expected to be directly under the source folder. Subfolders are
// not supported. Models are trained using documents that are of the following content type - 'application/pdf',
// 'image/jpeg' and 'image/png'."
// Other type of content is ignored.
// Parameters:
// trainRequest - request object for training.
func (client BaseClient) TrainCustomModel(ctx context.Context, trainRequest TrainRequest) (result TrainResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.TrainCustomModel")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: trainRequest,
			Constraints: []validation.Constraint{{Target: "trainRequest.Source", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "trainRequest.Source", Name: validation.MaxLength, Rule: 2048, Chain: nil},
					{Target: "trainRequest.Source", Name: validation.MinLength, Rule: 0, Chain: nil},
				}},
				{Target: "trainRequest.SourceFilter", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "trainRequest.SourceFilter.Prefix", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "trainRequest.SourceFilter.Prefix", Name: validation.MaxLength, Rule: 128, Chain: nil},
							{Target: "trainRequest.SourceFilter.Prefix", Name: validation.MinLength, Rule: 0, Chain: nil},
						}},
					}}}}}); err != nil {
		return result, validation.NewError("formrecognizer.BaseClient", "TrainCustomModel", err.Error())
	}

	req, err := client.TrainCustomModelPreparer(ctx, trainRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "TrainCustomModel", nil, "Failure preparing request")
		return
	}

	resp, err := client.TrainCustomModelSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "TrainCustomModel", resp, "Failure sending request")
		return
	}

	result, err = client.TrainCustomModelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "TrainCustomModel", resp, "Failure responding to request")
	}

	return
}

// TrainCustomModelPreparer prepares the TrainCustomModel request.
func (client BaseClient) TrainCustomModelPreparer(ctx context.Context, trainRequest TrainRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/formrecognizer/v1.0-preview", urlParameters),
		autorest.WithPath("/custom/train"),
		autorest.WithJSON(trainRequest))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TrainCustomModelSender sends the TrainCustomModel request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) TrainCustomModelSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// TrainCustomModelResponder handles the response to the TrainCustomModel request. The method always
// closes the http.Response Body.
func (client BaseClient) TrainCustomModelResponder(resp *http.Response) (result TrainResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
