# CreateOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** |  | 
**Qty** | Pointer to **float64** |  | [optional] 
**Notional** | Pointer to **float64** |  | [optional] 
**Side** | **string** |  | 
**Type** | **string** |  | 
**TimeInForce** | **string** |  | 
**LimitPrice** | Pointer to **float64** |  | [optional] 
**StopPrice** | Pointer to **float64** |  | [optional] 
**TrailPrice** | Pointer to **float64** |  | [optional] 
**TrailPercent** | Pointer to **float64** |  | [optional] 
**ExtendedHours** | Pointer to **bool** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**OrderClass** | Pointer to **string** |  | [optional] 
**TakeProfit** | Pointer to [**CreateOrderRequestTakeProfit**](CreateOrderRequestTakeProfit.md) |  | [optional] 
**StopLoss** | Pointer to [**CreateOrderRequestStopLoss**](CreateOrderRequestStopLoss.md) |  | [optional] 
**Commission** | Pointer to **float64** |  | [optional] 

## Methods

### NewCreateOrderRequest

`func NewCreateOrderRequest(symbol string, side string, type_ string, timeInForce string, ) *CreateOrderRequest`

NewCreateOrderRequest instantiates a new CreateOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderRequestWithDefaults

`func NewCreateOrderRequestWithDefaults() *CreateOrderRequest`

NewCreateOrderRequestWithDefaults instantiates a new CreateOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *CreateOrderRequest) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CreateOrderRequest) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CreateOrderRequest) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetQty

`func (o *CreateOrderRequest) GetQty() float64`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *CreateOrderRequest) GetQtyOk() (*float64, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *CreateOrderRequest) SetQty(v float64)`

SetQty sets Qty field to given value.

### HasQty

`func (o *CreateOrderRequest) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetNotional

`func (o *CreateOrderRequest) GetNotional() float64`

GetNotional returns the Notional field if non-nil, zero value otherwise.

### GetNotionalOk

`func (o *CreateOrderRequest) GetNotionalOk() (*float64, bool)`

GetNotionalOk returns a tuple with the Notional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotional

`func (o *CreateOrderRequest) SetNotional(v float64)`

SetNotional sets Notional field to given value.

### HasNotional

`func (o *CreateOrderRequest) HasNotional() bool`

HasNotional returns a boolean if a field has been set.

### GetSide

`func (o *CreateOrderRequest) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateOrderRequest) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateOrderRequest) SetSide(v string)`

SetSide sets Side field to given value.


### GetType

`func (o *CreateOrderRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOrderRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOrderRequest) SetType(v string)`

SetType sets Type field to given value.


### GetTimeInForce

`func (o *CreateOrderRequest) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *CreateOrderRequest) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *CreateOrderRequest) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.


### GetLimitPrice

`func (o *CreateOrderRequest) GetLimitPrice() float64`

GetLimitPrice returns the LimitPrice field if non-nil, zero value otherwise.

### GetLimitPriceOk

`func (o *CreateOrderRequest) GetLimitPriceOk() (*float64, bool)`

GetLimitPriceOk returns a tuple with the LimitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPrice

`func (o *CreateOrderRequest) SetLimitPrice(v float64)`

SetLimitPrice sets LimitPrice field to given value.

### HasLimitPrice

`func (o *CreateOrderRequest) HasLimitPrice() bool`

HasLimitPrice returns a boolean if a field has been set.

### GetStopPrice

`func (o *CreateOrderRequest) GetStopPrice() float64`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *CreateOrderRequest) GetStopPriceOk() (*float64, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *CreateOrderRequest) SetStopPrice(v float64)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *CreateOrderRequest) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetTrailPrice

`func (o *CreateOrderRequest) GetTrailPrice() float64`

GetTrailPrice returns the TrailPrice field if non-nil, zero value otherwise.

### GetTrailPriceOk

`func (o *CreateOrderRequest) GetTrailPriceOk() (*float64, bool)`

GetTrailPriceOk returns a tuple with the TrailPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailPrice

`func (o *CreateOrderRequest) SetTrailPrice(v float64)`

SetTrailPrice sets TrailPrice field to given value.

### HasTrailPrice

`func (o *CreateOrderRequest) HasTrailPrice() bool`

HasTrailPrice returns a boolean if a field has been set.

### GetTrailPercent

`func (o *CreateOrderRequest) GetTrailPercent() float64`

GetTrailPercent returns the TrailPercent field if non-nil, zero value otherwise.

### GetTrailPercentOk

`func (o *CreateOrderRequest) GetTrailPercentOk() (*float64, bool)`

GetTrailPercentOk returns a tuple with the TrailPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailPercent

`func (o *CreateOrderRequest) SetTrailPercent(v float64)`

SetTrailPercent sets TrailPercent field to given value.

### HasTrailPercent

`func (o *CreateOrderRequest) HasTrailPercent() bool`

HasTrailPercent returns a boolean if a field has been set.

### GetExtendedHours

`func (o *CreateOrderRequest) GetExtendedHours() bool`

GetExtendedHours returns the ExtendedHours field if non-nil, zero value otherwise.

### GetExtendedHoursOk

`func (o *CreateOrderRequest) GetExtendedHoursOk() (*bool, bool)`

GetExtendedHoursOk returns a tuple with the ExtendedHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedHours

`func (o *CreateOrderRequest) SetExtendedHours(v bool)`

SetExtendedHours sets ExtendedHours field to given value.

### HasExtendedHours

`func (o *CreateOrderRequest) HasExtendedHours() bool`

HasExtendedHours returns a boolean if a field has been set.

### GetClientOrderId

`func (o *CreateOrderRequest) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *CreateOrderRequest) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *CreateOrderRequest) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *CreateOrderRequest) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetOrderClass

`func (o *CreateOrderRequest) GetOrderClass() string`

GetOrderClass returns the OrderClass field if non-nil, zero value otherwise.

### GetOrderClassOk

`func (o *CreateOrderRequest) GetOrderClassOk() (*string, bool)`

GetOrderClassOk returns a tuple with the OrderClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderClass

`func (o *CreateOrderRequest) SetOrderClass(v string)`

SetOrderClass sets OrderClass field to given value.

### HasOrderClass

`func (o *CreateOrderRequest) HasOrderClass() bool`

HasOrderClass returns a boolean if a field has been set.

### GetTakeProfit

`func (o *CreateOrderRequest) GetTakeProfit() CreateOrderRequestTakeProfit`

GetTakeProfit returns the TakeProfit field if non-nil, zero value otherwise.

### GetTakeProfitOk

`func (o *CreateOrderRequest) GetTakeProfitOk() (*CreateOrderRequestTakeProfit, bool)`

GetTakeProfitOk returns a tuple with the TakeProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakeProfit

`func (o *CreateOrderRequest) SetTakeProfit(v CreateOrderRequestTakeProfit)`

SetTakeProfit sets TakeProfit field to given value.

### HasTakeProfit

`func (o *CreateOrderRequest) HasTakeProfit() bool`

HasTakeProfit returns a boolean if a field has been set.

### GetStopLoss

`func (o *CreateOrderRequest) GetStopLoss() CreateOrderRequestStopLoss`

GetStopLoss returns the StopLoss field if non-nil, zero value otherwise.

### GetStopLossOk

`func (o *CreateOrderRequest) GetStopLossOk() (*CreateOrderRequestStopLoss, bool)`

GetStopLossOk returns a tuple with the StopLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopLoss

`func (o *CreateOrderRequest) SetStopLoss(v CreateOrderRequestStopLoss)`

SetStopLoss sets StopLoss field to given value.

### HasStopLoss

`func (o *CreateOrderRequest) HasStopLoss() bool`

HasStopLoss returns a boolean if a field has been set.

### GetCommission

`func (o *CreateOrderRequest) GetCommission() float64`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *CreateOrderRequest) GetCommissionOk() (*float64, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *CreateOrderRequest) SetCommission(v float64)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *CreateOrderRequest) HasCommission() bool`

HasCommission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


