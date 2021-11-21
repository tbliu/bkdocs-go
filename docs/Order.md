# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**SubmittedAt** | Pointer to **time.Time** |  | [optional] 
**FilledAt** | Pointer to **NullableTime** |  | [optional] 
**ExpiredAt** | Pointer to **NullableTime** |  | [optional] 
**CanceledAt** | Pointer to **NullableTime** |  | [optional] 
**FailedAt** | Pointer to **NullableTime** |  | [optional] 
**ReplacedAt** | Pointer to **NullableTime** |  | [optional] 
**ReplacedBy** | Pointer to **NullableString** |  | [optional] 
**Replaces** | Pointer to **NullableString** |  | [optional] 
**AssetId** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**AssetClass** | Pointer to **string** |  | [optional] 
**Notional** | Pointer to **NullableFloat64** |  | [optional] 
**Qty** | Pointer to **NullableFloat64** |  | [optional] 
**FilledQty** | Pointer to **float64** |  | [optional] 
**FilledAvgPrice** | Pointer to **NullableFloat64** |  | [optional] 
**OrderClass** | Pointer to **string** |  | [optional] 
**OrderType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**LimitPrice** | Pointer to **NullableFloat64** |  | [optional] 
**StopPrice** | Pointer to **NullableFloat64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ExtendedHours** | Pointer to **bool** |  | [optional] 
**Legs** | Pointer to [**[]Order**](Order.md) |  | [optional] 
**TrailPrice** | Pointer to **NullableFloat64** |  | [optional] 
**TrailPercent** | Pointer to **NullableFloat64** |  | [optional] 
**Hwm** | Pointer to **NullableFloat64** |  | [optional] 
**Commission** | Pointer to **float64** |  | [optional] 

## Methods

### NewOrder

`func NewOrder() *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Order) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Order) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Order) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Order) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientOrderId

`func (o *Order) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *Order) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *Order) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *Order) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Order) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Order) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Order) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Order) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Order) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Order) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Order) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Order) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *Order) GetSubmittedAt() time.Time`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *Order) GetSubmittedAtOk() (*time.Time, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *Order) SetSubmittedAt(v time.Time)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *Order) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetFilledAt

`func (o *Order) GetFilledAt() time.Time`

GetFilledAt returns the FilledAt field if non-nil, zero value otherwise.

### GetFilledAtOk

`func (o *Order) GetFilledAtOk() (*time.Time, bool)`

GetFilledAtOk returns a tuple with the FilledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledAt

`func (o *Order) SetFilledAt(v time.Time)`

SetFilledAt sets FilledAt field to given value.

### HasFilledAt

`func (o *Order) HasFilledAt() bool`

HasFilledAt returns a boolean if a field has been set.

### SetFilledAtNil

`func (o *Order) SetFilledAtNil(b bool)`

 SetFilledAtNil sets the value for FilledAt to be an explicit nil

### UnsetFilledAt
`func (o *Order) UnsetFilledAt()`

UnsetFilledAt ensures that no value is present for FilledAt, not even an explicit nil
### GetExpiredAt

`func (o *Order) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *Order) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *Order) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *Order) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### SetExpiredAtNil

`func (o *Order) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *Order) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
### GetCanceledAt

`func (o *Order) GetCanceledAt() time.Time`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *Order) GetCanceledAtOk() (*time.Time, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *Order) SetCanceledAt(v time.Time)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *Order) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.

### SetCanceledAtNil

`func (o *Order) SetCanceledAtNil(b bool)`

 SetCanceledAtNil sets the value for CanceledAt to be an explicit nil

### UnsetCanceledAt
`func (o *Order) UnsetCanceledAt()`

UnsetCanceledAt ensures that no value is present for CanceledAt, not even an explicit nil
### GetFailedAt

`func (o *Order) GetFailedAt() time.Time`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *Order) GetFailedAtOk() (*time.Time, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *Order) SetFailedAt(v time.Time)`

SetFailedAt sets FailedAt field to given value.

### HasFailedAt

`func (o *Order) HasFailedAt() bool`

HasFailedAt returns a boolean if a field has been set.

### SetFailedAtNil

`func (o *Order) SetFailedAtNil(b bool)`

 SetFailedAtNil sets the value for FailedAt to be an explicit nil

### UnsetFailedAt
`func (o *Order) UnsetFailedAt()`

UnsetFailedAt ensures that no value is present for FailedAt, not even an explicit nil
### GetReplacedAt

`func (o *Order) GetReplacedAt() time.Time`

GetReplacedAt returns the ReplacedAt field if non-nil, zero value otherwise.

### GetReplacedAtOk

`func (o *Order) GetReplacedAtOk() (*time.Time, bool)`

GetReplacedAtOk returns a tuple with the ReplacedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedAt

`func (o *Order) SetReplacedAt(v time.Time)`

SetReplacedAt sets ReplacedAt field to given value.

### HasReplacedAt

`func (o *Order) HasReplacedAt() bool`

HasReplacedAt returns a boolean if a field has been set.

### SetReplacedAtNil

`func (o *Order) SetReplacedAtNil(b bool)`

 SetReplacedAtNil sets the value for ReplacedAt to be an explicit nil

### UnsetReplacedAt
`func (o *Order) UnsetReplacedAt()`

UnsetReplacedAt ensures that no value is present for ReplacedAt, not even an explicit nil
### GetReplacedBy

`func (o *Order) GetReplacedBy() string`

GetReplacedBy returns the ReplacedBy field if non-nil, zero value otherwise.

### GetReplacedByOk

`func (o *Order) GetReplacedByOk() (*string, bool)`

GetReplacedByOk returns a tuple with the ReplacedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedBy

`func (o *Order) SetReplacedBy(v string)`

SetReplacedBy sets ReplacedBy field to given value.

### HasReplacedBy

`func (o *Order) HasReplacedBy() bool`

HasReplacedBy returns a boolean if a field has been set.

### SetReplacedByNil

`func (o *Order) SetReplacedByNil(b bool)`

 SetReplacedByNil sets the value for ReplacedBy to be an explicit nil

### UnsetReplacedBy
`func (o *Order) UnsetReplacedBy()`

UnsetReplacedBy ensures that no value is present for ReplacedBy, not even an explicit nil
### GetReplaces

`func (o *Order) GetReplaces() string`

GetReplaces returns the Replaces field if non-nil, zero value otherwise.

### GetReplacesOk

`func (o *Order) GetReplacesOk() (*string, bool)`

GetReplacesOk returns a tuple with the Replaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaces

`func (o *Order) SetReplaces(v string)`

SetReplaces sets Replaces field to given value.

### HasReplaces

`func (o *Order) HasReplaces() bool`

HasReplaces returns a boolean if a field has been set.

### SetReplacesNil

`func (o *Order) SetReplacesNil(b bool)`

 SetReplacesNil sets the value for Replaces to be an explicit nil

### UnsetReplaces
`func (o *Order) UnsetReplaces()`

UnsetReplaces ensures that no value is present for Replaces, not even an explicit nil
### GetAssetId

`func (o *Order) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *Order) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *Order) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *Order) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetSymbol

`func (o *Order) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Order) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Order) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Order) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetAssetClass

`func (o *Order) GetAssetClass() string`

GetAssetClass returns the AssetClass field if non-nil, zero value otherwise.

### GetAssetClassOk

`func (o *Order) GetAssetClassOk() (*string, bool)`

GetAssetClassOk returns a tuple with the AssetClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetClass

`func (o *Order) SetAssetClass(v string)`

SetAssetClass sets AssetClass field to given value.

### HasAssetClass

`func (o *Order) HasAssetClass() bool`

HasAssetClass returns a boolean if a field has been set.

### GetNotional

`func (o *Order) GetNotional() float64`

GetNotional returns the Notional field if non-nil, zero value otherwise.

### GetNotionalOk

`func (o *Order) GetNotionalOk() (*float64, bool)`

GetNotionalOk returns a tuple with the Notional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotional

`func (o *Order) SetNotional(v float64)`

SetNotional sets Notional field to given value.

### HasNotional

`func (o *Order) HasNotional() bool`

HasNotional returns a boolean if a field has been set.

### SetNotionalNil

`func (o *Order) SetNotionalNil(b bool)`

 SetNotionalNil sets the value for Notional to be an explicit nil

### UnsetNotional
`func (o *Order) UnsetNotional()`

UnsetNotional ensures that no value is present for Notional, not even an explicit nil
### GetQty

`func (o *Order) GetQty() float64`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *Order) GetQtyOk() (*float64, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *Order) SetQty(v float64)`

SetQty sets Qty field to given value.

### HasQty

`func (o *Order) HasQty() bool`

HasQty returns a boolean if a field has been set.

### SetQtyNil

`func (o *Order) SetQtyNil(b bool)`

 SetQtyNil sets the value for Qty to be an explicit nil

### UnsetQty
`func (o *Order) UnsetQty()`

UnsetQty ensures that no value is present for Qty, not even an explicit nil
### GetFilledQty

`func (o *Order) GetFilledQty() float64`

GetFilledQty returns the FilledQty field if non-nil, zero value otherwise.

### GetFilledQtyOk

`func (o *Order) GetFilledQtyOk() (*float64, bool)`

GetFilledQtyOk returns a tuple with the FilledQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledQty

`func (o *Order) SetFilledQty(v float64)`

SetFilledQty sets FilledQty field to given value.

### HasFilledQty

`func (o *Order) HasFilledQty() bool`

HasFilledQty returns a boolean if a field has been set.

### GetFilledAvgPrice

`func (o *Order) GetFilledAvgPrice() float64`

GetFilledAvgPrice returns the FilledAvgPrice field if non-nil, zero value otherwise.

### GetFilledAvgPriceOk

`func (o *Order) GetFilledAvgPriceOk() (*float64, bool)`

GetFilledAvgPriceOk returns a tuple with the FilledAvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledAvgPrice

`func (o *Order) SetFilledAvgPrice(v float64)`

SetFilledAvgPrice sets FilledAvgPrice field to given value.

### HasFilledAvgPrice

`func (o *Order) HasFilledAvgPrice() bool`

HasFilledAvgPrice returns a boolean if a field has been set.

### SetFilledAvgPriceNil

`func (o *Order) SetFilledAvgPriceNil(b bool)`

 SetFilledAvgPriceNil sets the value for FilledAvgPrice to be an explicit nil

### UnsetFilledAvgPrice
`func (o *Order) UnsetFilledAvgPrice()`

UnsetFilledAvgPrice ensures that no value is present for FilledAvgPrice, not even an explicit nil
### GetOrderClass

`func (o *Order) GetOrderClass() string`

GetOrderClass returns the OrderClass field if non-nil, zero value otherwise.

### GetOrderClassOk

`func (o *Order) GetOrderClassOk() (*string, bool)`

GetOrderClassOk returns a tuple with the OrderClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderClass

`func (o *Order) SetOrderClass(v string)`

SetOrderClass sets OrderClass field to given value.

### HasOrderClass

`func (o *Order) HasOrderClass() bool`

HasOrderClass returns a boolean if a field has been set.

### GetOrderType

`func (o *Order) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *Order) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *Order) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *Order) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetType

`func (o *Order) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Order) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Order) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Order) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSide

`func (o *Order) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *Order) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *Order) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *Order) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetTimeInForce

`func (o *Order) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *Order) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *Order) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *Order) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetLimitPrice

`func (o *Order) GetLimitPrice() float64`

GetLimitPrice returns the LimitPrice field if non-nil, zero value otherwise.

### GetLimitPriceOk

`func (o *Order) GetLimitPriceOk() (*float64, bool)`

GetLimitPriceOk returns a tuple with the LimitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPrice

`func (o *Order) SetLimitPrice(v float64)`

SetLimitPrice sets LimitPrice field to given value.

### HasLimitPrice

`func (o *Order) HasLimitPrice() bool`

HasLimitPrice returns a boolean if a field has been set.

### SetLimitPriceNil

`func (o *Order) SetLimitPriceNil(b bool)`

 SetLimitPriceNil sets the value for LimitPrice to be an explicit nil

### UnsetLimitPrice
`func (o *Order) UnsetLimitPrice()`

UnsetLimitPrice ensures that no value is present for LimitPrice, not even an explicit nil
### GetStopPrice

`func (o *Order) GetStopPrice() float64`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *Order) GetStopPriceOk() (*float64, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *Order) SetStopPrice(v float64)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *Order) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### SetStopPriceNil

`func (o *Order) SetStopPriceNil(b bool)`

 SetStopPriceNil sets the value for StopPrice to be an explicit nil

### UnsetStopPrice
`func (o *Order) UnsetStopPrice()`

UnsetStopPrice ensures that no value is present for StopPrice, not even an explicit nil
### GetStatus

`func (o *Order) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Order) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Order) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Order) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExtendedHours

`func (o *Order) GetExtendedHours() bool`

GetExtendedHours returns the ExtendedHours field if non-nil, zero value otherwise.

### GetExtendedHoursOk

`func (o *Order) GetExtendedHoursOk() (*bool, bool)`

GetExtendedHoursOk returns a tuple with the ExtendedHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedHours

`func (o *Order) SetExtendedHours(v bool)`

SetExtendedHours sets ExtendedHours field to given value.

### HasExtendedHours

`func (o *Order) HasExtendedHours() bool`

HasExtendedHours returns a boolean if a field has been set.

### GetLegs

`func (o *Order) GetLegs() []Order`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *Order) GetLegsOk() (*[]Order, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *Order) SetLegs(v []Order)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *Order) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### SetLegsNil

`func (o *Order) SetLegsNil(b bool)`

 SetLegsNil sets the value for Legs to be an explicit nil

### UnsetLegs
`func (o *Order) UnsetLegs()`

UnsetLegs ensures that no value is present for Legs, not even an explicit nil
### GetTrailPrice

`func (o *Order) GetTrailPrice() float64`

GetTrailPrice returns the TrailPrice field if non-nil, zero value otherwise.

### GetTrailPriceOk

`func (o *Order) GetTrailPriceOk() (*float64, bool)`

GetTrailPriceOk returns a tuple with the TrailPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailPrice

`func (o *Order) SetTrailPrice(v float64)`

SetTrailPrice sets TrailPrice field to given value.

### HasTrailPrice

`func (o *Order) HasTrailPrice() bool`

HasTrailPrice returns a boolean if a field has been set.

### SetTrailPriceNil

`func (o *Order) SetTrailPriceNil(b bool)`

 SetTrailPriceNil sets the value for TrailPrice to be an explicit nil

### UnsetTrailPrice
`func (o *Order) UnsetTrailPrice()`

UnsetTrailPrice ensures that no value is present for TrailPrice, not even an explicit nil
### GetTrailPercent

`func (o *Order) GetTrailPercent() float64`

GetTrailPercent returns the TrailPercent field if non-nil, zero value otherwise.

### GetTrailPercentOk

`func (o *Order) GetTrailPercentOk() (*float64, bool)`

GetTrailPercentOk returns a tuple with the TrailPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailPercent

`func (o *Order) SetTrailPercent(v float64)`

SetTrailPercent sets TrailPercent field to given value.

### HasTrailPercent

`func (o *Order) HasTrailPercent() bool`

HasTrailPercent returns a boolean if a field has been set.

### SetTrailPercentNil

`func (o *Order) SetTrailPercentNil(b bool)`

 SetTrailPercentNil sets the value for TrailPercent to be an explicit nil

### UnsetTrailPercent
`func (o *Order) UnsetTrailPercent()`

UnsetTrailPercent ensures that no value is present for TrailPercent, not even an explicit nil
### GetHwm

`func (o *Order) GetHwm() float64`

GetHwm returns the Hwm field if non-nil, zero value otherwise.

### GetHwmOk

`func (o *Order) GetHwmOk() (*float64, bool)`

GetHwmOk returns a tuple with the Hwm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwm

`func (o *Order) SetHwm(v float64)`

SetHwm sets Hwm field to given value.

### HasHwm

`func (o *Order) HasHwm() bool`

HasHwm returns a boolean if a field has been set.

### SetHwmNil

`func (o *Order) SetHwmNil(b bool)`

 SetHwmNil sets the value for Hwm to be an explicit nil

### UnsetHwm
`func (o *Order) UnsetHwm()`

UnsetHwm ensures that no value is present for Hwm, not even an explicit nil
### GetCommission

`func (o *Order) GetCommission() float64`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *Order) GetCommissionOk() (*float64, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *Order) SetCommission(v float64)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *Order) HasCommission() bool`

HasCommission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


