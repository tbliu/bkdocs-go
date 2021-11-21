# \AccountsApi

All URIs are relative to *https://broker-api.sandbox.alpaca.markets/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountsAccountIdDocumentsUploadPost**](AccountsApi.md#AccountsAccountIdDocumentsUploadPost) | **Post** /accounts/{account_id}/documents/upload | Upload a document to an already existing account
[**AccountsActivitiesActivityTypeGet**](AccountsApi.md#AccountsActivitiesActivityTypeGet) | **Get** /accounts/activities/{activity_type} | Retrieve specific account activities
[**AccountsActivitiesGet**](AccountsApi.md#AccountsActivitiesGet) | **Get** /accounts/activities | Retrieve account activities
[**AccountsGet**](AccountsApi.md#AccountsGet) | **Get** /accounts | Retrieve all accounts
[**AccountsPost**](AccountsApi.md#AccountsPost) | **Post** /accounts | Create an account
[**DeleteAccount**](AccountsApi.md#DeleteAccount) | **Delete** /accounts/{account_id} | Request to close an account
[**DeleteAchRelationship**](AccountsApi.md#DeleteAchRelationship) | **Delete** /accounts/{account_id}/ach_relationships/{ach_relationship_id} | Delete an existing ACH relationship
[**DeleteRecipientBank**](AccountsApi.md#DeleteRecipientBank) | **Delete** /accounts/{account_id}/recipient_banks/{bank_id} | Delete a Bank Relationship for an account
[**DeleteTransfer**](AccountsApi.md#DeleteTransfer) | **Delete** /accounts/{account_id}/transfers/{transfer_id} | Request to close a transfer
[**EventsAccountsStatusGet**](AccountsApi.md#EventsAccountsStatusGet) | **Get** /events/accounts/status | Subscribe to account status events (SSE).
[**GetAccount**](AccountsApi.md#GetAccount) | **Get** /accounts/{account_id} | Retrieve an account.
[**GetAchRelationships**](AccountsApi.md#GetAchRelationships) | **Get** /accounts/{account_id}/ach_relationships | Retrieve ACH Relationships for an account
[**GetRecipientBanks**](AccountsApi.md#GetRecipientBanks) | **Get** /accounts/{account_id}/recipient_banks | Retrieve bank relationships for an account
[**GetTradingAccount**](AccountsApi.md#GetTradingAccount) | **Get** /trading/accounts/{account_id}/account | Retrieve trading details for an account.
[**GetTransfers**](AccountsApi.md#GetTransfers) | **Get** /accounts/{account_id}/transfers | Return a list of transfers for an account.
[**PatchAccount**](AccountsApi.md#PatchAccount) | **Patch** /accounts/{account_id} | Update an account
[**PostAchRelationships**](AccountsApi.md#PostAchRelationships) | **Post** /accounts/{account_id}/ach_relationships | Create an ACH Relationship
[**PostRecipientBanks**](AccountsApi.md#PostRecipientBanks) | **Post** /accounts/{account_id}/recipient_banks | Create a Bank Relationship for an account
[**PostTransfers**](AccountsApi.md#PostTransfers) | **Post** /accounts/{account_id}/transfers | Request a new transfer



## AccountsAccountIdDocumentsUploadPost

> AccountsAccountIdDocumentsUploadPost(ctx, accountId).DocumentUploadRequest(documentUploadRequest).Execute()

Upload a document to an already existing account

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
    documentUploadRequest := *openapiclient.NewDocumentUploadRequest(openapiclient.DocumentType("identity_verification"), "QWxwYWNhcyBjYW5ub3QgbGl2ZSBhbG9uZS4=", "image/jpeg") // DocumentUploadRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsAccountIdDocumentsUploadPost(context.Background(), accountId).DocumentUploadRequest(documentUploadRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsAccountIdDocumentsUploadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsAccountIdDocumentsUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **documentUploadRequest** | [**DocumentUploadRequest**](DocumentUploadRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsActivitiesActivityTypeGet

> []ActivityItem AccountsActivitiesActivityTypeGet(ctx, activityType).Date(date).Until(until).After(after).Direction(direction).AccountId(accountId).PageSize(pageSize).PageToken(pageToken).Execute()

Retrieve specific account activities

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
    activityType := "activityType_example" // string | 
    date := "date_example" // string |  (optional)
    until := "until_example" // string |  (optional)
    after := "after_example" // string |  (optional)
    direction := "direction_example" // string |  (optional)
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    pageToken := "pageToken_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsActivitiesActivityTypeGet(context.Background(), activityType).Date(date).Until(until).After(after).Direction(direction).AccountId(accountId).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsActivitiesActivityTypeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsActivitiesActivityTypeGet`: []ActivityItem
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsActivitiesActivityTypeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsActivitiesActivityTypeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **string** |  | 
 **until** | **string** |  | 
 **after** | **string** |  | 
 **direction** | **string** |  | 
 **accountId** | **string** |  | 
 **pageSize** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**[]ActivityItem**](ActivityItem.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsActivitiesGet

> []ActivityItem AccountsActivitiesGet(ctx).Date(date).Until(until).After(after).Direction(direction).AccountId(accountId).PageSize(pageSize).PageToken(pageToken).Execute()

Retrieve account activities

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
    date := "date_example" // string |  (optional)
    until := "until_example" // string |  (optional)
    after := "after_example" // string |  (optional)
    direction := "direction_example" // string |  (optional)
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    pageToken := "pageToken_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsActivitiesGet(context.Background()).Date(date).Until(until).After(after).Direction(direction).AccountId(accountId).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsActivitiesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsActivitiesGet`: []ActivityItem
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsActivitiesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsActivitiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **string** |  | 
 **until** | **string** |  | 
 **after** | **string** |  | 
 **direction** | **string** |  | 
 **accountId** | **string** |  | 
 **pageSize** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**[]ActivityItem**](ActivityItem.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsGet

> []Account AccountsGet(ctx).Query(query).Execute()

Retrieve all accounts

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
    query := "query_example" // string | The query supports partial match of account number, names, emails, etc.. Items can be space delimited.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsGet(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsGet`: []Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The query supports partial match of account number, names, emails, etc.. Items can be space delimited.  | 

### Return type

[**[]Account**](Account.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsPost

> Account AccountsPost(ctx).AccountCreationRequest(accountCreationRequest).Execute()

Create an account

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
    accountCreationRequest := *openapiclient.NewAccountCreationRequest() // AccountCreationRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsPost(context.Background()).AccountCreationRequest(accountCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsPost`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountCreationRequest** | [**AccountCreationRequest**](AccountCreationRequest.md) |  | 

### Return type

[**Account**](Account.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccount

> DeleteAccount(ctx, accountId).Execute()

Request to close an account

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.DeleteAccount(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.DeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


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


## DeleteAchRelationship

> DeleteAchRelationship(ctx, accountId, achRelationshipId).Execute()

Delete an existing ACH relationship

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
    achRelationshipId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ACH relationship identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.DeleteAchRelationship(context.Background(), accountId, achRelationshipId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.DeleteAchRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account identifier. | 
**achRelationshipId** | **string** | ACH relationship identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAchRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecipientBank

> DeleteRecipientBank(ctx, accountId, bankId).Execute()

Delete a Bank Relationship for an account

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
    bankId := "bankId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.DeleteRecipientBank(context.Background(), accountId, bankId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.DeleteRecipientBank``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account identifier. | 
**bankId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecipientBankRequest struct via the builder pattern


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


## DeleteTransfer

> DeleteTransfer(ctx, accountId, transferId).Execute()

Request to close a transfer

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
    accountId := "accountId_example" // string | 
    transferId := "transferId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.DeleteTransfer(context.Background(), accountId, transferId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.DeleteTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**transferId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsAccountsStatusGet

> AccountStatusEvent EventsAccountsStatusGet(ctx).Since(since).Until(until).SinceId(sinceId).UntilId(untilId).Execute()

Subscribe to account status events (SSE).



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
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    sinceId := int32(56) // int32 |  (optional)
    untilId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.EventsAccountsStatusGet(context.Background()).Since(since).Until(until).SinceId(sinceId).UntilId(untilId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.EventsAccountsStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsAccountsStatusGet`: AccountStatusEvent
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.EventsAccountsStatusGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsAccountsStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **sinceId** | **int32** |  | 
 **untilId** | **int32** |  | 

### Return type

[**AccountStatusEvent**](AccountStatusEvent.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> AccountExtended GetAccount(ctx, accountId).Execute()

Retrieve an account.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetAccount(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccount`: AccountExtended
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountExtended**](AccountExtended.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAchRelationships

> []ACHRelationshipResource GetAchRelationships(ctx, accountId).Statuses(statuses).Execute()

Retrieve ACH Relationships for an account

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
    statuses := "statuses_example" // string | Comma-separated status values (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetAchRelationships(context.Background(), accountId).Statuses(statuses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAchRelationships``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAchRelationships`: []ACHRelationshipResource
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAchRelationships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAchRelationshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **statuses** | **string** | Comma-separated status values | 

### Return type

[**[]ACHRelationshipResource**](ACHRelationshipResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecipientBanks

> []BankResource GetRecipientBanks(ctx, accountId).Status(status).BankName(bankName).Execute()

Retrieve bank relationships for an account

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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    status := "status_example" // string |  (optional)
    bankName := "bankName_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetRecipientBanks(context.Background(), accountId).Status(status).BankName(bankName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetRecipientBanks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecipientBanks`: []BankResource
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetRecipientBanks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecipientBanksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** |  | 
 **bankName** | **string** |  | 

### Return type

[**[]BankResource**](BankResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingAccount

> InlineResponse200 GetTradingAccount(ctx, accountId).Execute()

Retrieve trading details for an account.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetTradingAccount(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetTradingAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTradingAccount`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetTradingAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransfers

> []TransferResource GetTransfers(ctx, accountId).Direction(direction).Limit(limit).Offset(offset).Execute()

Return a list of transfers for an account.



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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    direction := "direction_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetTransfers(context.Background(), accountId).Direction(direction).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetTransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransfers`: []TransferResource
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetTransfers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **direction** | **string** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 

### Return type

[**[]TransferResource**](TransferResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAccount

> Account PatchAccount(ctx, accountId).AccountUpdateRequest(accountUpdateRequest).Execute()

Update an account

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
    accountUpdateRequest := *openapiclient.NewAccountUpdateRequest() // AccountUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.PatchAccount(context.Background(), accountId).AccountUpdateRequest(accountUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.PatchAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.PatchAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountUpdateRequest** | [**AccountUpdateRequest**](AccountUpdateRequest.md) |  | 

### Return type

[**Account**](Account.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAchRelationships

> ACHRelationshipResource PostAchRelationships(ctx, accountId).ACHRelationshipData(aCHRelationshipData).Execute()

Create an ACH Relationship

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
    aCHRelationshipData := *openapiclient.NewACHRelationshipData("AccountOwnerName_example", "BankAccountType_example", "BankAccountNumber_example", "BankRoutingNumber_example") // ACHRelationshipData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.PostAchRelationships(context.Background(), accountId).ACHRelationshipData(aCHRelationshipData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.PostAchRelationships``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAchRelationships`: ACHRelationshipResource
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.PostAchRelationships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAchRelationshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aCHRelationshipData** | [**ACHRelationshipData**](ACHRelationshipData.md) |  | 

### Return type

[**ACHRelationshipResource**](ACHRelationshipResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRecipientBanks

> BankResource PostRecipientBanks(ctx, accountId).BankData(bankData).Execute()

Create a Bank Relationship for an account

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
    bankData := *openapiclient.NewBankData("Name_example", "BankCode_example", "BankCodeType_example", "AccountNumber_example") // BankData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.PostRecipientBanks(context.Background(), accountId).BankData(bankData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.PostRecipientBanks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostRecipientBanks`: BankResource
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.PostRecipientBanks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRecipientBanksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bankData** | [**BankData**](BankData.md) |  | 

### Return type

[**BankResource**](BankResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransfers

> TransferResource PostTransfers(ctx, accountId).TransferData(transferData).Execute()

Request a new transfer



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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    transferData := *openapiclient.NewTransferData("TransferType_example", "Amount_example", "Direction_example", "RelationshipId_example", "BankId_example") // TransferData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.PostTransfers(context.Background(), accountId).TransferData(transferData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.PostTransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTransfers`: TransferResource
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.PostTransfers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferData** | [**TransferData**](TransferData.md) |  | 

### Return type

[**TransferResource**](TransferResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

