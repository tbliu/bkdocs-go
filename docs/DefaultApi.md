# \DefaultApi

All URIs are relative to *https://broker-api.sandbox.alpaca.markets/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAccountsAccountIdWatchlistsWatchlistId**](DefaultApi.md#DeleteAccountsAccountIdWatchlistsWatchlistId) | **Delete** /accounts/{account_id}/watchlists/{watchlist_id} | Remove a watchlist
[**GetAccountsAccountIdWatchlistsWatchlistId**](DefaultApi.md#GetAccountsAccountIdWatchlistsWatchlistId) | **Get** /accounts/{account_id}/watchlists/{watchlist_id} | Manage watchlists
[**GetTradingAccountsAccountIdWatchlists**](DefaultApi.md#GetTradingAccountsAccountIdWatchlists) | **Get** /trading/accounts/{account_id}/watchlists | Retrieve all watchlists
[**PostTradingAccountsAccountIdWatchlists**](DefaultApi.md#PostTradingAccountsAccountIdWatchlists) | **Post** /trading/accounts/{account_id}/watchlists | Create a new watchlist
[**PutAccountsAccountIdWatchlistsWatchlistId**](DefaultApi.md#PutAccountsAccountIdWatchlistsWatchlistId) | **Put** /accounts/{account_id}/watchlists/{watchlist_id} | Update an existing watchlist



## DeleteAccountsAccountIdWatchlistsWatchlistId

> DeleteAccountsAccountIdWatchlistsWatchlistId(ctx, accountId, watchlistId).Execute()

Remove a watchlist



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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of an account
    watchlistId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of a watchlist

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteAccountsAccountIdWatchlistsWatchlistId(context.Background(), accountId, watchlistId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAccountsAccountIdWatchlistsWatchlistId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Unique identifier of an account | 
**watchlistId** | **string** | Unique identifier of a watchlist | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountsAccountIdWatchlistsWatchlistIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountsAccountIdWatchlistsWatchlistId

> Watchlist GetAccountsAccountIdWatchlistsWatchlistId(ctx, accountId, watchlistId).Execute()

Manage watchlists



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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of an account
    watchlistId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of a watchlist

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetAccountsAccountIdWatchlistsWatchlistId(context.Background(), accountId, watchlistId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAccountsAccountIdWatchlistsWatchlistId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountsAccountIdWatchlistsWatchlistId`: Watchlist
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAccountsAccountIdWatchlistsWatchlistId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Unique identifier of an account | 
**watchlistId** | **string** | Unique identifier of a watchlist | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsAccountIdWatchlistsWatchlistIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Watchlist**](Watchlist.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingAccountsAccountIdWatchlists

> []Watchlist GetTradingAccountsAccountIdWatchlists(ctx, accountId).Execute()

Retrieve all watchlists



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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of an account.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetTradingAccountsAccountIdWatchlists(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTradingAccountsAccountIdWatchlists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTradingAccountsAccountIdWatchlists`: []Watchlist
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTradingAccountsAccountIdWatchlists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Unique identifier of an account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingAccountsAccountIdWatchlistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Watchlist**](Watchlist.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTradingAccountsAccountIdWatchlists

> Watchlist PostTradingAccountsAccountIdWatchlists(ctx, accountId).Watchlist(watchlist).Execute()

Create a new watchlist

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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of an account.
    watchlist := *openapiclient.NewWatchlist() // Watchlist |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.PostTradingAccountsAccountIdWatchlists(context.Background(), accountId).Watchlist(watchlist).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostTradingAccountsAccountIdWatchlists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTradingAccountsAccountIdWatchlists`: Watchlist
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostTradingAccountsAccountIdWatchlists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Unique identifier of an account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTradingAccountsAccountIdWatchlistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **watchlist** | [**Watchlist**](Watchlist.md) |  | 

### Return type

[**Watchlist**](Watchlist.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAccountsAccountIdWatchlistsWatchlistId

> Watchlist PutAccountsAccountIdWatchlistsWatchlistId(ctx, accountId, watchlistId).InlineObject2(inlineObject2).Execute()

Update an existing watchlist



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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of an account
    watchlistId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of a watchlist
    inlineObject2 := *openapiclient.NewInlineObject2() // InlineObject2 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.PutAccountsAccountIdWatchlistsWatchlistId(context.Background(), accountId, watchlistId).InlineObject2(inlineObject2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PutAccountsAccountIdWatchlistsWatchlistId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutAccountsAccountIdWatchlistsWatchlistId`: Watchlist
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PutAccountsAccountIdWatchlistsWatchlistId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Unique identifier of an account | 
**watchlistId** | **string** | Unique identifier of a watchlist | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAccountsAccountIdWatchlistsWatchlistIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject2** | [**InlineObject2**](InlineObject2.md) |  | 

### Return type

[**Watchlist**](Watchlist.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

