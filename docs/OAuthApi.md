# \OAuthApi

All URIs are relative to *https://broker-api.sandbox.alpaca.markets/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OauthAuthorizePost**](OAuthApi.md#OauthAuthorizePost) | **Post** /oauth/authorize | Issue a code.
[**OauthClientsClientIdGet**](OAuthApi.md#OauthClientsClientIdGet) | **Get** /oauth/clients/{client_id} | Returns an OAuth client.
[**OauthTokenPost**](OAuthApi.md#OauthTokenPost) | **Post** /oauth/token | Issue a token.



## OauthAuthorizePost

> InlineResponse2007 OauthAuthorizePost(ctx).InlineObject1(inlineObject1).Execute()

Issue a code.



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
    inlineObject1 := *openapiclient.NewInlineObject1() // InlineObject1 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OAuthApi.OauthAuthorizePost(context.Background()).InlineObject1(inlineObject1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthApi.OauthAuthorizePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OauthAuthorizePost`: InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `OAuthApi.OauthAuthorizePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauthAuthorizePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject1** | [**InlineObject1**](InlineObject1.md) |  | 

### Return type

[**InlineResponse2007**](InlineResponse2007.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OauthClientsClientIdGet

> InlineResponse2005 OauthClientsClientIdGet(ctx, clientId).ResponseType(responseType).RedirectUri(redirectUri).Scope(scope).Execute()

Returns an OAuth client.



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
    clientId := TODO // string | 
    responseType := "token" // string |  (optional)
    redirectUri := "https://example.com/authorize" // string |  (optional)
    scope := "general" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OAuthApi.OauthClientsClientIdGet(context.Background(), clientId).ResponseType(responseType).RedirectUri(redirectUri).Scope(scope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthApi.OauthClientsClientIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OauthClientsClientIdGet`: InlineResponse2005
    fmt.Fprintf(os.Stdout, "Response from `OAuthApi.OauthClientsClientIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauthClientsClientIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **responseType** | **string** |  | 
 **redirectUri** | **string** |  | 
 **scope** | **string** |  | 

### Return type

[**InlineResponse2005**](InlineResponse2005.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OauthTokenPost

> InlineResponse2006 OauthTokenPost(ctx).InlineObject(inlineObject).Execute()

Issue a token.



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
    inlineObject := *openapiclient.NewInlineObject() // InlineObject | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OAuthApi.OauthTokenPost(context.Background()).InlineObject(inlineObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthApi.OauthTokenPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OauthTokenPost`: InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `OAuthApi.OauthTokenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauthTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**InlineObject**](InlineObject.md) |  | 

### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

