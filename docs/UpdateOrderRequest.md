# UpdateOrderRequest

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

### NewUpdateOrderRequest

`func NewUpdateOrderRequest(createdAt time.Time, updatedAt time.Time, ) *UpdateOrderRequest`

NewUpdateOrderRequest instantiates a new UpdateOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrderRequestWithDefaults

`func NewUpdateOrderRequestWithDefaults() *UpdateOrderRequest`

NewUpdateOrderRequestWithDefaults instantiates a new UpdateOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQty

`func (o *UpdateOrderRequest) GetQty() float64`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *UpdateOrderRequest) GetQtyOk() (*float64, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *UpdateOrderRequest) SetQty(v float64)`

SetQty sets Qty field to given value.

### HasQty

`func (o *UpdateOrderRequest) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetTimeInForce

`func (o *UpdateOrderRequest) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *UpdateOrderRequest) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *UpdateOrderRequest) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *UpdateOrderRequest) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetLimitPrice

`func (o *UpdateOrderRequest) GetLimitPrice() float64`

GetLimitPrice returns the LimitPrice field if non-nil, zero value otherwise.

### GetLimitPriceOk

`func (o *UpdateOrderRequest) GetLimitPriceOk() (*float64, bool)`

GetLimitPriceOk returns a tuple with the LimitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPrice

`func (o *UpdateOrderRequest) SetLimitPrice(v float64)`

SetLimitPrice sets LimitPrice field to given value.

### HasLimitPrice

`func (o *UpdateOrderRequest) HasLimitPrice() bool`

HasLimitPrice returns a boolean if a field has been set.

### GetStopPrice

`func (o *UpdateOrderRequest) GetStopPrice() float64`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *UpdateOrderRequest) GetStopPriceOk() (*float64, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *UpdateOrderRequest) SetStopPrice(v float64)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *UpdateOrderRequest) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetTrail

`func (o *UpdateOrderRequest) GetTrail() float64`

GetTrail returns the Trail field if non-nil, zero value otherwise.

### GetTrailOk

`func (o *UpdateOrderRequest) GetTrailOk() (*float64, bool)`

GetTrailOk returns a tuple with the Trail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrail

`func (o *UpdateOrderRequest) SetTrail(v float64)`

SetTrail sets Trail field to given value.

### HasTrail

`func (o *UpdateOrderRequest) HasTrail() bool`

HasTrail returns a boolean if a field has been set.

### GetClientOrderId

`func (o *UpdateOrderRequest) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *UpdateOrderRequest) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *UpdateOrderRequest) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *UpdateOrderRequest) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UpdateOrderRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdateOrderRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdateOrderRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UpdateOrderRequest) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UpdateOrderRequest) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UpdateOrderRequest) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


