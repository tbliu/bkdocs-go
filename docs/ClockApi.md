# \ClockApi

All URIs are relative to *https://broker-api.sandbox.alpaca.markets/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClockGet**](ClockApi.md#ClockGet) | **Get** /clock | Query market clock



## ClockGet

> InlineResponse2001 ClockGet(ctx).Execute()

Query market clock

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClockApi.ClockGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClockApi.ClockGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClockGet`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ClockApi.ClockGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiClockGetRequest struct via the builder pattern


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

