# \FundingApi

All URIs are relative to *https://broker-api.sandbox.alpaca.markets/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAchRelationship**](FundingApi.md#DeleteAchRelationship) | **Delete** /accounts/{account_id}/ach_relationships/{ach_relationship_id} | Delete an existing ACH relationship
[**DeleteRecipientBank**](FundingApi.md#DeleteRecipientBank) | **Delete** /accounts/{account_id}/recipient_banks/{bank_id} | Delete a Bank Relationship for an account
[**DeleteTransfer**](FundingApi.md#DeleteTransfer) | **Delete** /accounts/{account_id}/transfers/{transfer_id} | Request to close a transfer
[**GetAchRelationships**](FundingApi.md#GetAchRelationships) | **Get** /accounts/{account_id}/ach_relationships | Retrieve ACH Relationships for an account
[**GetRecipientBanks**](FundingApi.md#GetRecipientBanks) | **Get** /accounts/{account_id}/recipient_banks | Retrieve bank relationships for an account
[**GetTransfers**](FundingApi.md#GetTransfers) | **Get** /accounts/{account_id}/transfers | Return a list of transfers for an account.
[**PostAchRelationships**](FundingApi.md#PostAchRelationships) | **Post** /accounts/{account_id}/ach_relationships | Create an ACH Relationship
[**PostRecipientBanks**](FundingApi.md#PostRecipientBanks) | **Post** /accounts/{account_id}/recipient_banks | Create a Bank Relationship for an account
[**PostTransfers**](FundingApi.md#PostTransfers) | **Post** /accounts/{account_id}/transfers | Request a new transfer



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
    resp, r, err := api_client.FundingApi.DeleteAchRelationship(context.Background(), accountId, achRelationshipId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.DeleteAchRelationship``: %v\n", err)
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
    resp, r, err := api_client.FundingApi.DeleteRecipientBank(context.Background(), accountId, bankId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.DeleteRecipientBank``: %v\n", err)
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
    resp, r, err := api_client.FundingApi.DeleteTransfer(context.Background(), accountId, transferId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.DeleteTransfer``: %v\n", err)
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
    resp, r, err := api_client.FundingApi.GetAchRelationships(context.Background(), accountId).Statuses(statuses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.GetAchRelationships``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAchRelationships`: []ACHRelationshipResource
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.GetAchRelationships`: %v\n", resp)
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
    resp, r, err := api_client.FundingApi.GetRecipientBanks(context.Background(), accountId).Status(status).BankName(bankName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.GetRecipientBanks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecipientBanks`: []BankResource
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.GetRecipientBanks`: %v\n", resp)
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
    resp, r, err := api_client.FundingApi.GetTransfers(context.Background(), accountId).Direction(direction).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.GetTransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransfers`: []TransferResource
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.GetTransfers`: %v\n", resp)
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
    resp, r, err := api_client.FundingApi.PostAchRelationships(context.Background(), accountId).ACHRelationshipData(aCHRelationshipData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.PostAchRelationships``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAchRelationships`: ACHRelationshipResource
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.PostAchRelationships`: %v\n", resp)
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
    resp, r, err := api_client.FundingApi.PostRecipientBanks(context.Background(), accountId).BankData(bankData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.PostRecipientBanks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostRecipientBanks`: BankResource
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.PostRecipientBanks`: %v\n", resp)
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
    resp, r, err := api_client.FundingApi.PostTransfers(context.Background(), accountId).TransferData(transferData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.PostTransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTransfers`: TransferResource
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.PostTransfers`: %v\n", resp)
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

