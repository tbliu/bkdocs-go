# PatchOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Qty** | Pointer to **float64** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**LimitPrice** | Pointer to **float64** |  | [optional] 
**StopPrice** | Pointer to **float64** |  | [optional] 
**Trail** | Pointer to **float64** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewPatchOrder

`func NewPatchOrder(createdAt time.Time, updatedAt time.Time, ) *PatchOrder`

NewPatchOrder instantiates a new PatchOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchOrderWithDefaults

`func NewPatchOrderWithDefaults() *PatchOrder`

NewPatchOrderWithDefaults instantiates a new PatchOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQty

`func (o *PatchOrder) GetQty() float64`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *PatchOrder) GetQtyOk() (*float64, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *PatchOrder) SetQty(v float64)`

SetQty sets Qty field to given value.

### HasQty

`func (o *PatchOrder) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetTimeInForce

`func (o *PatchOrder) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *PatchOrder) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *PatchOrder) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *PatchOrder) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetLimitPrice

`func (o *PatchOrder) GetLimitPrice() float64`

GetLimitPrice returns the LimitPrice field if non-nil, zero value otherwise.

### GetLimitPriceOk

`func (o *PatchOrder) GetLimitPriceOk() (*float64, bool)`

GetLimitPriceOk returns a tuple with the LimitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPrice

`func (o *PatchOrder) SetLimitPrice(v float64)`

SetLimitPrice sets LimitPrice field to given value.

### HasLimitPrice

`func (o *PatchOrder) HasLimitPrice() bool`

HasLimitPrice returns a boolean if a field has been set.

### GetStopPrice

`func (o *PatchOrder) GetStopPrice() float64`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *PatchOrder) GetStopPriceOk() (*float64, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *PatchOrder) SetStopPrice(v float64)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *PatchOrder) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetTrail

`func (o *PatchOrder) GetTrail() float64`

GetTrail returns the Trail field if non-nil, zero value otherwise.

### GetTrailOk

`func (o *PatchOrder) GetTrailOk() (*float64, bool)`

GetTrailOk returns a tuple with the Trail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrail

`func (o *PatchOrder) SetTrail(v float64)`

SetTrail sets Trail field to given value.

### HasTrail

`func (o *PatchOrder) HasTrail() bool`

HasTrail returns a boolean if a field has been set.

### GetClientOrderId

`func (o *PatchOrder) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *PatchOrder) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *PatchOrder) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *PatchOrder) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchOrder) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchOrder) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchOrder) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PatchOrder) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PatchOrder) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PatchOrder) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


