# \EventsApi

All URIs are relative to *https://broker-api.sandbox.alpaca.markets/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsAccountsStatusGet**](EventsApi.md#EventsAccountsStatusGet) | **Get** /events/accounts/status | Subscribe to account status events (SSE).
[**EventsJournalsStatusGet**](EventsApi.md#EventsJournalsStatusGet) | **Get** /events/journals/status | Subscribe to journal events (SSE).



## EventsAccountsStatusGet

> InlineResponse2003 EventsAccountsStatusGet(ctx).Since(since).Until(until).SinceId(sinceId).UntilId(untilId).Execute()

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
    resp, r, err := api_client.EventsApi.EventsAccountsStatusGet(context.Background()).Since(since).Until(until).SinceId(sinceId).UntilId(untilId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsAccountsStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsAccountsStatusGet`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsAccountsStatusGet`: %v\n", resp)
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

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsJournalsStatusGet

> InlineResponse2004 EventsJournalsStatusGet(ctx).Since(since).Until(until).SinceId(sinceId).UntilId(untilId).Execute()

Subscribe to journal events (SSE).



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
    resp, r, err := api_client.EventsApi.EventsJournalsStatusGet(context.Background()).Since(since).Until(until).SinceId(sinceId).UntilId(untilId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsJournalsStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsJournalsStatusGet`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsJournalsStatusGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsJournalsStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **sinceId** | **int32** |  | 
 **untilId** | **int32** |  | 

### Return type

[**InlineResponse2004**](InlineResponse2004.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

