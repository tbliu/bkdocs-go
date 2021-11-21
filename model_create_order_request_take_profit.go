/*
Alpaca Broker API

Open brokerage accounts, enable commission-free trading, and manage the ongoing user experience with Alpaca Broker API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CreateOrderRequestTakeProfit struct for CreateOrderRequestTakeProfit
type CreateOrderRequestTakeProfit struct {
	LimitPrice *float64 `json:"limit_price,omitempty"`
}

// NewCreateOrderRequestTakeProfit instantiates a new CreateOrderRequestTakeProfit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrderRequestTakeProfit() *CreateOrderRequestTakeProfit {
	this := CreateOrderRequestTakeProfit{}
	return &this
}

// NewCreateOrderRequestTakeProfitWithDefaults instantiates a new CreateOrderRequestTakeProfit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrderRequestTakeProfitWithDefaults() *CreateOrderRequestTakeProfit {
	this := CreateOrderRequestTakeProfit{}
	return &this
}

// GetLimitPrice returns the LimitPrice field value if set, zero value otherwise.
func (o *CreateOrderRequestTakeProfit) GetLimitPrice() float64 {
	if o == nil || o.LimitPrice == nil {
		var ret float64
		return ret
	}
	return *o.LimitPrice
}

// GetLimitPriceOk returns a tuple with the LimitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequestTakeProfit) GetLimitPriceOk() (*float64, bool) {
	if o == nil || o.LimitPrice == nil {
		return nil, false
	}
	return o.LimitPrice, true
}

// HasLimitPrice returns a boolean if a field has been set.
func (o *CreateOrderRequestTakeProfit) HasLimitPrice() bool {
	if o != nil && o.LimitPrice != nil {
		return true
	}

	return false
}

// SetLimitPrice gets a reference to the given float64 and assigns it to the LimitPrice field.
func (o *CreateOrderRequestTakeProfit) SetLimitPrice(v float64) {
	o.LimitPrice = &v
}

func (o CreateOrderRequestTakeProfit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LimitPrice != nil {
		toSerialize["limit_price"] = o.LimitPrice
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrderRequestTakeProfit struct {
	value *CreateOrderRequestTakeProfit
	isSet bool
}

func (v NullableCreateOrderRequestTakeProfit) Get() *CreateOrderRequestTakeProfit {
	return v.value
}

func (v *NullableCreateOrderRequestTakeProfit) Set(val *CreateOrderRequestTakeProfit) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrderRequestTakeProfit) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrderRequestTakeProfit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrderRequestTakeProfit(val *CreateOrderRequestTakeProfit) *NullableCreateOrderRequestTakeProfit {
	return &NullableCreateOrderRequestTakeProfit{value: val, isSet: true}
}

func (v NullableCreateOrderRequestTakeProfit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrderRequestTakeProfit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

