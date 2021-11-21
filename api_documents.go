/*
Alpaca Broker API

Open brokerage accounts, enable commission-free trading, and manage the ongoing user experience with Alpaca Broker API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// DocumentsApiService DocumentsApi service
type DocumentsApiService service

type DocumentsApiApiAccountsAccountIdDocumentsDocumentIdDownloadGetRequest struct {
	ctx _context.Context
	ApiService *DocumentsApiService
	accountId string
	documentId string
}


func (r DocumentsApiApiAccountsAccountIdDocumentsDocumentIdDownloadGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.AccountsAccountIdDocumentsDocumentIdDownloadGetExecute(r)
}

/*
AccountsAccountIdDocumentsDocumentIdDownloadGet Download a document file that belongs to an account.

The operation returns a pre-signed downloadable link as
a redirect with HTTP status code 301 if one is found.


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId Account identifier.
 @param documentId
 @return DocumentsApiApiAccountsAccountIdDocumentsDocumentIdDownloadGetRequest
*/
func (a *DocumentsApiService) AccountsAccountIdDocumentsDocumentIdDownloadGet(ctx _context.Context, accountId string, documentId string) DocumentsApiApiAccountsAccountIdDocumentsDocumentIdDownloadGetRequest {
	return DocumentsApiApiAccountsAccountIdDocumentsDocumentIdDownloadGetRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		documentId: documentId,
	}
}

// Execute executes the request
func (a *DocumentsApiService) AccountsAccountIdDocumentsDocumentIdDownloadGetExecute(r DocumentsApiApiAccountsAccountIdDocumentsDocumentIdDownloadGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DocumentsApiService.AccountsAccountIdDocumentsDocumentIdDownloadGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/{account_id}/documents/{document_id}/download"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", _neturl.PathEscape(parameterToString(r.accountId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"document_id"+"}", _neturl.PathEscape(parameterToString(r.documentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type DocumentsApiApiAccountsAccountIdDocumentsGetRequest struct {
	ctx _context.Context
	ApiService *DocumentsApiService
	accountId string
	startDate *string
	endDate *string
}

// optional date value to filter the list (inclusive).
func (r DocumentsApiApiAccountsAccountIdDocumentsGetRequest) StartDate(startDate string) DocumentsApiApiAccountsAccountIdDocumentsGetRequest {
	r.startDate = &startDate
	return r
}
// optional date value to filter the list (inclusive).
func (r DocumentsApiApiAccountsAccountIdDocumentsGetRequest) EndDate(endDate string) DocumentsApiApiAccountsAccountIdDocumentsGetRequest {
	r.endDate = &endDate
	return r
}

func (r DocumentsApiApiAccountsAccountIdDocumentsGetRequest) Execute() ([][]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.AccountsAccountIdDocumentsGetExecute(r)
}

/*
AccountsAccountIdDocumentsGet Return a list of account documents.

You can query account documents such as monthly 
statements and trade confirms under an account.


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId Account identifier.
 @return DocumentsApiApiAccountsAccountIdDocumentsGetRequest
*/
func (a *DocumentsApiService) AccountsAccountIdDocumentsGet(ctx _context.Context, accountId string) DocumentsApiApiAccountsAccountIdDocumentsGetRequest {
	return DocumentsApiApiAccountsAccountIdDocumentsGetRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return [][]map[string]interface{}
func (a *DocumentsApiService) AccountsAccountIdDocumentsGetExecute(r DocumentsApiApiAccountsAccountIdDocumentsGetRequest) ([][]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  [][]map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DocumentsApiService.AccountsAccountIdDocumentsGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/{account_id}/documents"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", _neturl.PathEscape(parameterToString(r.accountId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.startDate != nil {
		localVarQueryParams.Add("start_date", parameterToString(*r.startDate, ""))
	}
	if r.endDate != nil {
		localVarQueryParams.Add("end_date", parameterToString(*r.endDate, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DocumentsApiApiDocumentsDocumentIdGetRequest struct {
	ctx _context.Context
	ApiService *DocumentsApiService
	documentId string
}


func (r DocumentsApiApiDocumentsDocumentIdGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DocumentsDocumentIdGetExecute(r)
}

/*
DocumentsDocumentIdGet Download a document file directly

The operation returns a pre-signed downloadable link as
a redirect with HTTP status code 301 if one is found.


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param documentId
 @return DocumentsApiApiDocumentsDocumentIdGetRequest
*/
func (a *DocumentsApiService) DocumentsDocumentIdGet(ctx _context.Context, documentId string) DocumentsApiApiDocumentsDocumentIdGetRequest {
	return DocumentsApiApiDocumentsDocumentIdGetRequest{
		ApiService: a,
		ctx: ctx,
		documentId: documentId,
	}
}

// Execute executes the request
func (a *DocumentsApiService) DocumentsDocumentIdGetExecute(r DocumentsApiApiDocumentsDocumentIdGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DocumentsApiService.DocumentsDocumentIdGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/documents/{document_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"document_id"+"}", _neturl.PathEscape(parameterToString(r.documentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
