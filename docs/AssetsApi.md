# \AssetsApi

All URIs are relative to *https://broker-api.sandbox.alpaca.markets/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsAssetIdGet**](AssetsApi.md#AssetsAssetIdGet) | **Get** /assets/{asset_id} | Retrieve an asset by UUID
[**AssetsSymbolGet**](AssetsApi.md#AssetsSymbolGet) | **Get** /assets/{symbol} | Retrieve an asset by symbol
[**GetAssets**](AssetsApi.md#GetAssets) | **Get** /assets | Retrieve all assets



## AssetsAssetIdGet

> AssetResource AssetsAssetIdGet(ctx, assetId).Execute()

Retrieve an asset by UUID



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
    assetId := "assetId_example" // string | The UUID of the required asset

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetsApi.AssetsAssetIdGet(context.Background(), assetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetsAssetIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetsAssetIdGet`: AssetResource
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetsAssetIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | The UUID of the required asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsAssetIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssetResource**](AssetResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsSymbolGet

> AssetResource AssetsSymbolGet(ctx, symbol).Execute()

Retrieve an asset by symbol



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
    symbol := "symbol_example" // string | The symbol of the required asset

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetsApi.AssetsSymbolGet(context.Background(), symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.AssetsSymbolGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetsSymbolGet`: AssetResource
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.AssetsSymbolGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string** | The symbol of the required asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsSymbolGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssetResource**](AssetResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssets

> []AssetResource GetAssets(ctx).Execute()

Retrieve all assets



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
    resp, r, err := api_client.AssetsApi.GetAssets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.GetAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssets`: []AssetResource
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.GetAssets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetsRequest struct via the builder pattern


### Return type

[**[]AssetResource**](AssetResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

