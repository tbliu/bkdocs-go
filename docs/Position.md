# Position

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Exchange** | Pointer to **string** |  | [optional] 
**AssetClass** | Pointer to **string** |  | [optional] 
**AvgEntryPrice** | Pointer to **string** |  | [optional] 
**Qty** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**MarketValue** | Pointer to **string** |  | [optional] 
**CostBasis** | Pointer to **string** |  | [optional] 
**UnrealizedPl** | Pointer to **string** |  | [optional] 
**UnrealizedPlpc** | Pointer to **string** |  | [optional] 
**UnrealizedIntradayPl** | Pointer to **string** |  | [optional] 
**UnrealizedIntradayPlpc** | Pointer to **string** |  | [optional] 
**CurrentPrice** | Pointer to **string** |  | [optional] 
**LastdayPrice** | Pointer to **string** |  | [optional] 
**ChangeToday** | Pointer to **string** |  | [optional] 

## Methods

### NewPosition

`func NewPosition() *Position`

NewPosition instantiates a new Position object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositionWithDefaults

`func NewPositionWithDefaults() *Position`

NewPositionWithDefaults instantiates a new Position object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *Position) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *Position) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *Position) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *Position) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetSymbol

`func (o *Position) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Position) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Position) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Position) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetExchange

`func (o *Position) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *Position) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *Position) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *Position) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetAssetClass

`func (o *Position) GetAssetClass() string`

GetAssetClass returns the AssetClass field if non-nil, zero value otherwise.

### GetAssetClassOk

`func (o *Position) GetAssetClassOk() (*string, bool)`

GetAssetClassOk returns a tuple with the AssetClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetClass

`func (o *Position) SetAssetClass(v string)`

SetAssetClass sets AssetClass field to given value.

### HasAssetClass

`func (o *Position) HasAssetClass() bool`

HasAssetClass returns a boolean if a field has been set.

### GetAvgEntryPrice

`func (o *Position) GetAvgEntryPrice() string`

GetAvgEntryPrice returns the AvgEntryPrice field if non-nil, zero value otherwise.

### GetAvgEntryPriceOk

`func (o *Position) GetAvgEntryPriceOk() (*string, bool)`

GetAvgEntryPriceOk returns a tuple with the AvgEntryPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgEntryPrice

`func (o *Position) SetAvgEntryPrice(v string)`

SetAvgEntryPrice sets AvgEntryPrice field to given value.

### HasAvgEntryPrice

`func (o *Position) HasAvgEntryPrice() bool`

HasAvgEntryPrice returns a boolean if a field has been set.

### GetQty

`func (o *Position) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *Position) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *Position) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *Position) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetSide

`func (o *Position) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *Position) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *Position) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *Position) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetMarketValue

`func (o *Position) GetMarketValue() string`

GetMarketValue returns the MarketValue field if non-nil, zero value otherwise.

### GetMarketValueOk

`func (o *Position) GetMarketValueOk() (*string, bool)`

GetMarketValueOk returns a tuple with the MarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketValue

`func (o *Position) SetMarketValue(v string)`

SetMarketValue sets MarketValue field to given value.

### HasMarketValue

`func (o *Position) HasMarketValue() bool`

HasMarketValue returns a boolean if a field has been set.

### GetCostBasis

`func (o *Position) GetCostBasis() string`

GetCostBasis returns the CostBasis field if non-nil, zero value otherwise.

### GetCostBasisOk

`func (o *Position) GetCostBasisOk() (*string, bool)`

GetCostBasisOk returns a tuple with the CostBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostBasis

`func (o *Position) SetCostBasis(v string)`

SetCostBasis sets CostBasis field to given value.

### HasCostBasis

`func (o *Position) HasCostBasis() bool`

HasCostBasis returns a boolean if a field has been set.

### GetUnrealizedPl

`func (o *Position) GetUnrealizedPl() string`

GetUnrealizedPl returns the UnrealizedPl field if non-nil, zero value otherwise.

### GetUnrealizedPlOk

`func (o *Position) GetUnrealizedPlOk() (*string, bool)`

GetUnrealizedPlOk returns a tuple with the UnrealizedPl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedPl

`func (o *Position) SetUnrealizedPl(v string)`

SetUnrealizedPl sets UnrealizedPl field to given value.

### HasUnrealizedPl

`func (o *Position) HasUnrealizedPl() bool`

HasUnrealizedPl returns a boolean if a field has been set.

### GetUnrealizedPlpc

`func (o *Position) GetUnrealizedPlpc() string`

GetUnrealizedPlpc returns the UnrealizedPlpc field if non-nil, zero value otherwise.

### GetUnrealizedPlpcOk

`func (o *Position) GetUnrealizedPlpcOk() (*string, bool)`

GetUnrealizedPlpcOk returns a tuple with the UnrealizedPlpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedPlpc

`func (o *Position) SetUnrealizedPlpc(v string)`

SetUnrealizedPlpc sets UnrealizedPlpc field to given value.

### HasUnrealizedPlpc

`func (o *Position) HasUnrealizedPlpc() bool`

HasUnrealizedPlpc returns a boolean if a field has been set.

### GetUnrealizedIntradayPl

`func (o *Position) GetUnrealizedIntradayPl() string`

GetUnrealizedIntradayPl returns the UnrealizedIntradayPl field if non-nil, zero value otherwise.

### GetUnrealizedIntradayPlOk

`func (o *Position) GetUnrealizedIntradayPlOk() (*string, bool)`

GetUnrealizedIntradayPlOk returns a tuple with the UnrealizedIntradayPl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedIntradayPl

`func (o *Position) SetUnrealizedIntradayPl(v string)`

SetUnrealizedIntradayPl sets UnrealizedIntradayPl field to given value.

### HasUnrealizedIntradayPl

`func (o *Position) HasUnrealizedIntradayPl() bool`

HasUnrealizedIntradayPl returns a boolean if a field has been set.

### GetUnrealizedIntradayPlpc

`func (o *Position) GetUnrealizedIntradayPlpc() string`

GetUnrealizedIntradayPlpc returns the UnrealizedIntradayPlpc field if non-nil, zero value otherwise.

### GetUnrealizedIntradayPlpcOk

`func (o *Position) GetUnrealizedIntradayPlpcOk() (*string, bool)`

GetUnrealizedIntradayPlpcOk returns a tuple with the UnrealizedIntradayPlpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedIntradayPlpc

`func (o *Position) SetUnrealizedIntradayPlpc(v string)`

SetUnrealizedIntradayPlpc sets UnrealizedIntradayPlpc field to given value.

### HasUnrealizedIntradayPlpc

`func (o *Position) HasUnrealizedIntradayPlpc() bool`

HasUnrealizedIntradayPlpc returns a boolean if a field has been set.

### GetCurrentPrice

`func (o *Position) GetCurrentPrice() string`

GetCurrentPrice returns the CurrentPrice field if non-nil, zero value otherwise.

### GetCurrentPriceOk

`func (o *Position) GetCurrentPriceOk() (*string, bool)`

GetCurrentPriceOk returns a tuple with the CurrentPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPrice

`func (o *Position) SetCurrentPrice(v string)`

SetCurrentPrice sets CurrentPrice field to given value.

### HasCurrentPrice

`func (o *Position) HasCurrentPrice() bool`

HasCurrentPrice returns a boolean if a field has been set.

### GetLastdayPrice

`func (o *Position) GetLastdayPrice() string`

GetLastdayPrice returns the LastdayPrice field if non-nil, zero value otherwise.

### GetLastdayPriceOk

`func (o *Position) GetLastdayPriceOk() (*string, bool)`

GetLastdayPriceOk returns a tuple with the LastdayPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastdayPrice

`func (o *Position) SetLastdayPrice(v string)`

SetLastdayPrice sets LastdayPrice field to given value.

### HasLastdayPrice

`func (o *Position) HasLastdayPrice() bool`

HasLastdayPrice returns a boolean if a field has been set.

### GetChangeToday

`func (o *Position) GetChangeToday() string`

GetChangeToday returns the ChangeToday field if non-nil, zero value otherwise.

### GetChangeTodayOk

`func (o *Position) GetChangeTodayOk() (*string, bool)`

GetChangeTodayOk returns a tuple with the ChangeToday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeToday

`func (o *Position) SetChangeToday(v string)`

SetChangeToday sets ChangeToday field to given value.

### HasChangeToday

`func (o *Position) HasChangeToday() bool`

HasChangeToday returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


