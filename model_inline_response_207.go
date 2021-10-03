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

// InlineResponse207 struct for InlineResponse207
type InlineResponse207 struct {
	Id *string `json:"id,omitempty"`
	Status *int32 `json:"status,omitempty"`
	Body *OrderObject `json:"body,omitempty"`
}

// NewInlineResponse207 instantiates a new InlineResponse207 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse207() *InlineResponse207 {
	this := InlineResponse207{}
	return &this
}

// NewInlineResponse207WithDefaults instantiates a new InlineResponse207 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse207WithDefaults() *InlineResponse207 {
	this := InlineResponse207{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse207) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse207) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse207) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse207) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse207) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse207) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse207) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *InlineResponse207) SetStatus(v int32) {
	o.Status = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *InlineResponse207) GetBody() OrderObject {
	if o == nil || o.Body == nil {
		var ret OrderObject
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse207) GetBodyOk() (*OrderObject, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *InlineResponse207) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given OrderObject and assigns it to the Body field.
func (o *InlineResponse207) SetBody(v OrderObject) {
	o.Body = &v
}

func (o InlineResponse207) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse207 struct {
	value *InlineResponse207
	isSet bool
}

func (v NullableInlineResponse207) Get() *InlineResponse207 {
	return v.value
}

func (v *NullableInlineResponse207) Set(val *InlineResponse207) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse207) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse207) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse207(val *InlineResponse207) *NullableInlineResponse207 {
	return &NullableInlineResponse207{value: val, isSet: true}
}

func (v NullableInlineResponse207) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse207) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

