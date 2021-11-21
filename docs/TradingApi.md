# \TradingApi

All URIs are relative to *https://broker-api.sandbox.alpaca.markets/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrder**](TradingApi.md#DeleteOrder) | **Delete** /trading/accounts/{account_id}/orders/{order_id} | Attempts to cancel an open order.
[**DeleteOrders**](TradingApi.md#DeleteOrders) | **Delete** /trading/accounts/{account_id}/orders | Attempts to cancel all open orders. A response will be provided for each order that is attempted to be cancelled.
[**GetOrder**](TradingApi.md#GetOrder) | **Get** /trading/accounts/{account_id}/orders/{order_id} | Retrieves a single order for the given order_id.
[**GetOrders**](TradingApi.md#GetOrders) | **Get** /trading/accounts/{account_id}/orders | Retrieves a list of orders for the account, filtered by the supplied query parameters.
[**GetPositions**](TradingApi.md#GetPositions) | **Get** /trading/accounts/{account_id}/positions | List open positions for an account
[**PatchOrder**](TradingApi.md#PatchOrder) | **Patch** /trading/accounts/{account_id}/orders/{order_id} | Replaces a single order with updated parameters. Each parameter overrides the corresponding attribute of the existing order.
[**PostOrders**](TradingApi.md#PostOrders) | **Post** /trading/accounts/{account_id}/orders | Create an order for an account.



## DeleteOrder

> DeleteOrder(ctx, accountId, orderId).Execute()

Attempts to cancel an open order.



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
    orderId := "orderId_example" // string | Order identifier.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradingApi.DeleteOrder(context.Background(), accountId, orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradingApi.DeleteOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account identifier. | 
**orderId** | **string** | Order identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrderRequest struct via the builder pattern


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


## DeleteOrders

> []InlineResponse207 DeleteOrders(ctx, accountId).Execute()

Attempts to cancel all open orders. A response will be provided for each order that is attempted to be cancelled.



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
    resp, r, err := api_client.TradingApi.DeleteOrders(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradingApi.DeleteOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrders`: []InlineResponse207
    fmt.Fprintf(os.Stdout, "Response from `TradingApi.DeleteOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse207**](InlineResponse207.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrder

> Order GetOrder(ctx, accountId, orderId).Execute()

Retrieves a single order for the given order_id.



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
    orderId := "orderId_example" // string | Order identifier.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradingApi.GetOrder(context.Background(), accountId, orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradingApi.GetOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrder`: Order
    fmt.Fprintf(os.Stdout, "Response from `TradingApi.GetOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account identifier. | 
**orderId** | **string** | Order identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Order**](Order.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrders

> []Order GetOrders(ctx, accountId).Status(status).Limit(limit).After(after).Until(until).Direction(direction).Nested(nested).Symbols(symbols).Execute()

Retrieves a list of orders for the account, filtered by the supplied query parameters.



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
    status := "status_example" // string | Status of the orders to list. (optional)
    limit := int32(500) // int32 | The maximum number of orders in response. (optional)
    after := time.Now() // time.Time | The response will include only ones submitted after this timestamp (exclusive.) (optional)
    until := time.Now() // time.Time | The response will include only ones submitted until this timestamp (exclusive.) (optional)
    direction := "desc" // string | The chronological order of response based on the submission time. asc or desc. Defaults to desc. (optional)
    nested := false // bool | If true, the result will roll up multi-leg orders under the legs field of primary order. (optional)
    symbols := "AAPL,TSLA,MSFT" // string | A comma-separated list of symbols to filter by. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradingApi.GetOrders(context.Background(), accountId).Status(status).Limit(limit).After(after).Until(until).Direction(direction).Nested(nested).Symbols(symbols).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradingApi.GetOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrders`: []Order
    fmt.Fprintf(os.Stdout, "Response from `TradingApi.GetOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Status of the orders to list. | 
 **limit** | **int32** | The maximum number of orders in response. | 
 **after** | **time.Time** | The response will include only ones submitted after this timestamp (exclusive.) | 
 **until** | **time.Time** | The response will include only ones submitted until this timestamp (exclusive.) | 
 **direction** | **string** | The chronological order of response based on the submission time. asc or desc. Defaults to desc. | 
 **nested** | **bool** | If true, the result will roll up multi-leg orders under the legs field of primary order. | 
 **symbols** | **string** | A comma-separated list of symbols to filter by. | 

### Return type

[**[]Order**](Order.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPositions

> []Position GetPositions(ctx, accountId).Execute()

List open positions for an account

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
    resp, r, err := api_client.TradingApi.GetPositions(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradingApi.GetPositions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPositions`: []Position
    fmt.Fprintf(os.Stdout, "Response from `TradingApi.GetPositions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPositionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Position**](Position.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchOrder

> Order PatchOrder(ctx, accountId, orderId).UpdateOrderRequest(updateOrderRequest).Execute()

Replaces a single order with updated parameters. Each parameter overrides the corresponding attribute of the existing order.



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
    orderId := "orderId_example" // string | Order identifier.
    updateOrderRequest := *openapiclient.NewUpdateOrderRequest(time.Now(), time.Now()) // UpdateOrderRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradingApi.PatchOrder(context.Background(), accountId, orderId).UpdateOrderRequest(updateOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradingApi.PatchOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchOrder`: Order
    fmt.Fprintf(os.Stdout, "Response from `TradingApi.PatchOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account identifier. | 
**orderId** | **string** | Order identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrderRequest** | [**UpdateOrderRequest**](UpdateOrderRequest.md) |  | 

### Return type

[**Order**](Order.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrders

> Order PostOrders(ctx, accountId).CreateOrderRequest(createOrderRequest).Execute()

Create an order for an account.



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
    createOrderRequest := *openapiclient.NewCreateOrderRequest("AAPL", "buy", "limit", "gtc") // CreateOrderRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradingApi.PostOrders(context.Background(), accountId).CreateOrderRequest(createOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradingApi.PostOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostOrders`: Order
    fmt.Fprintf(os.Stdout, "Response from `TradingApi.PostOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrderRequest** | [**CreateOrderRequest**](CreateOrderRequest.md) |  | 

### Return type

[**Order**](Order.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

