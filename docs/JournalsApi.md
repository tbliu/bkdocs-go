# \JournalsApi

All URIs are relative to *https://broker-api.sandbox.alpaca.markets/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteJournal**](JournalsApi.md#DeleteJournal) | **Delete** /journals/{journal_id} | Cancel a pending journal.
[**EventsJournalsStatusGet**](JournalsApi.md#EventsJournalsStatusGet) | **Get** /events/journals/status | Subscribe to journal events (SSE).
[**GetJournals**](JournalsApi.md#GetJournals) | **Get** /journals | Return a list of requested journals.
[**PostJournals**](JournalsApi.md#PostJournals) | **Post** /journals | Request a journal.
[**PostJournalsBatch**](JournalsApi.md#PostJournalsBatch) | **Post** /journals/batch | Create a batch journal



## DeleteJournal

> DeleteJournal(ctx, journalId).Execute()

Cancel a pending journal.



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
    journalId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.JournalsApi.DeleteJournal(context.Background(), journalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JournalsApi.DeleteJournal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJournalRequest struct via the builder pattern


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
    resp, r, err := api_client.JournalsApi.EventsJournalsStatusGet(context.Background()).Since(since).Until(until).SinceId(sinceId).UntilId(untilId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JournalsApi.EventsJournalsStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsJournalsStatusGet`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `JournalsApi.EventsJournalsStatusGet`: %v\n", resp)
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


## GetJournals

> []JournalResource GetJournals(ctx).After(after).Before(before).Status(status).EntryType(entryType).ToAccount(toAccount).FromAccount(fromAccount).Execute()

Return a list of requested journals.

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
    after := time.Now() // string | by settle_date (optional)
    before := time.Now() // string | by settle_date (optional)
    status := "status_example" // string |  (optional)
    entryType := "entryType_example" // string |  (optional)
    toAccount := TODO // string |  (optional)
    fromAccount := TODO // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.JournalsApi.GetJournals(context.Background()).After(after).Before(before).Status(status).EntryType(entryType).ToAccount(toAccount).FromAccount(fromAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JournalsApi.GetJournals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJournals`: []JournalResource
    fmt.Fprintf(os.Stdout, "Response from `JournalsApi.GetJournals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJournalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | by settle_date | 
 **before** | **string** | by settle_date | 
 **status** | **string** |  | 
 **entryType** | **string** |  | 
 **toAccount** | [**string**](string.md) |  | 
 **fromAccount** | [**string**](string.md) |  | 

### Return type

[**[]JournalResource**](JournalResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostJournals

> JournalResource PostJournals(ctx).JournalData(journalData).Execute()

Request a journal.



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
    journalData := *openapiclient.NewJournalData("EntryType_example", "FromAccount_example", "ToAccount_example") // JournalData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.JournalsApi.PostJournals(context.Background()).JournalData(journalData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JournalsApi.PostJournals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostJournals`: JournalResource
    fmt.Fprintf(os.Stdout, "Response from `JournalsApi.PostJournals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostJournalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **journalData** | [**JournalData**](JournalData.md) |  | 

### Return type

[**JournalResource**](JournalResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostJournalsBatch

> []BatchJournalResponse PostJournalsBatch(ctx).BatchJournalRequest(batchJournalRequest).Execute()

Create a batch journal

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
    batchJournalRequest := *openapiclient.NewBatchJournalRequest("EntryType_example", "FromAccount_example") // BatchJournalRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.JournalsApi.PostJournalsBatch(context.Background()).BatchJournalRequest(batchJournalRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JournalsApi.PostJournalsBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostJournalsBatch`: []BatchJournalResponse
    fmt.Fprintf(os.Stdout, "Response from `JournalsApi.PostJournalsBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostJournalsBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchJournalRequest** | [**BatchJournalRequest**](BatchJournalRequest.md) |  | 

### Return type

[**[]BatchJournalResponse**](BatchJournalResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

