# \CalendarApi

All URIs are relative to *https://broker-api.sandbox.alpaca.markets/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalendarGet**](CalendarApi.md#CalendarGet) | **Get** /calendar | Query market calendar



## CalendarGet

> []MarketDay CalendarGet(ctx).Start(start).End(end).Execute()

Query market calendar

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
    start := "start_example" // string | The first date to retrieve data for. (Inclusive) (optional)
    end := "end_example" // string | The last date to retrieve data for. (Inclusive) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalendarApi.CalendarGet(context.Background()).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarApi.CalendarGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CalendarGet`: []MarketDay
    fmt.Fprintf(os.Stdout, "Response from `CalendarApi.CalendarGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCalendarGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | The first date to retrieve data for. (Inclusive) | 
 **end** | **string** | The last date to retrieve data for. (Inclusive) | 

### Return type

[**[]MarketDay**](MarketDay.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

