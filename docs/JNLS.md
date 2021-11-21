# JNLS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Only valid for JNLS journals. Null for JNLC. | 
**Qty** | **float64** | Only valid for JNLS journals. Null for JNLC. | 
**Price** | **float64** | Only valid for JNLS journals. Null for JNLC. | 

## Methods

### NewJNLS

`func NewJNLS(symbol string, qty float64, price float64, ) *JNLS`

NewJNLS instantiates a new JNLS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJNLSWithDefaults

`func NewJNLSWithDefaults() *JNLS`

NewJNLSWithDefaults instantiates a new JNLS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *JNLS) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *JNLS) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *JNLS) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetQty

`func (o *JNLS) GetQty() float64`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *JNLS) GetQtyOk() (*float64, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *JNLS) SetQty(v float64)`

SetQty sets Qty field to given value.


### GetPrice

`func (o *JNLS) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *JNLS) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *JNLS) SetPrice(v float64)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


