# \DocumentsApi

All URIs are relative to *https://broker-api.sandbox.alpaca.markets/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountsAccountIdDocumentsDocumentIdDownloadGet**](DocumentsApi.md#AccountsAccountIdDocumentsDocumentIdDownloadGet) | **Get** /accounts/{account_id}/documents/{document_id}/download | Download a document file that belongs to an account.
[**AccountsAccountIdDocumentsGet**](DocumentsApi.md#AccountsAccountIdDocumentsGet) | **Get** /accounts/{account_id}/documents | Return a list of account documents.
[**DocumentsDocumentIdGet**](DocumentsApi.md#DocumentsDocumentIdGet) | **Get** /documents/{document_id} | Download a document file directly



## AccountsAccountIdDocumentsDocumentIdDownloadGet

> AccountsAccountIdDocumentsDocumentIdDownloadGet(ctx, accountId, documentId).Execute()

Download a document file that belongs to an account.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Account identifier.
    documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.AccountsAccountIdDocumentsDocumentIdDownloadGet(context.Background(), accountId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.AccountsAccountIdDocumentsDocumentIdDownloadGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account identifier. | 
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsAccountIdDocumentsDocumentIdDownloadGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsAccountIdDocumentsGet

> [][]map[string]interface{} AccountsAccountIdDocumentsGet(ctx, accountId).StartDate(startDate).EndDate(endDate).Execute()

Return a list of account documents.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Account identifier.
    startDate := time.Now() // string | optional date value to filter the list (inclusive). (optional)
    endDate := time.Now() // string | optional date value to filter the list (inclusive). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.AccountsAccountIdDocumentsGet(context.Background(), accountId).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.AccountsAccountIdDocumentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsAccountIdDocumentsGet`: [][]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.AccountsAccountIdDocumentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsAccountIdDocumentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | optional date value to filter the list (inclusive). | 
 **endDate** | **string** | optional date value to filter the list (inclusive). | 

### Return type

[**[][]map[string]interface{}**](set.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentsDocumentIdGet

> DocumentsDocumentIdGet(ctx, documentId).Execute()

Download a document file directly



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.DocumentsDocumentIdGet(context.Background(), documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.DocumentsDocumentIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDocumentsDocumentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

