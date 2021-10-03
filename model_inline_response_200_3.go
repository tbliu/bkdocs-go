/*
Alpaca Broker API

Open brokerage accounts, enable commission-free trading, and manage the ongoing user experience with Alpaca Broker API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// InlineResponse2003 struct for InlineResponse2003
type InlineResponse2003 struct {
	AccountId *string `json:"account_id,omitempty"`
	AccountNumber *string `json:"account_number,omitempty"`
	StatusFrom *string `json:"status_from,omitempty"`
	StatusTo *string `json:"status_to,omitempty"`
	Reason *string `json:"reason,omitempty"`
	At *time.Time `json:"at,omitempty"`
}

// NewInlineResponse2003 instantiates a new InlineResponse2003 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2003() *InlineResponse2003 {
	this := InlineResponse2003{}
	return &this
}

// NewInlineResponse2003WithDefaults instantiates a new InlineResponse2003 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2003WithDefaults() *InlineResponse2003 {
	this := InlineResponse2003{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *InlineResponse2003) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *InlineResponse2003) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *InlineResponse2003) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *InlineResponse2003) GetAccountNumber() string {
	if o == nil || o.AccountNumber == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetAccountNumberOk() (*string, bool) {
	if o == nil || o.AccountNumber == nil {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *InlineResponse2003) HasAccountNumber() bool {
	if o != nil && o.AccountNumber != nil {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *InlineResponse2003) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetStatusFrom returns the StatusFrom field value if set, zero value otherwise.
func (o *InlineResponse2003) GetStatusFrom() string {
	if o == nil || o.StatusFrom == nil {
		var ret string
		return ret
	}
	return *o.StatusFrom
}

// GetStatusFromOk returns a tuple with the StatusFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetStatusFromOk() (*string, bool) {
	if o == nil || o.StatusFrom == nil {
		return nil, false
	}
	return o.StatusFrom, true
}

// HasStatusFrom returns a boolean if a field has been set.
func (o *InlineResponse2003) HasStatusFrom() bool {
	if o != nil && o.StatusFrom != nil {
		return true
	}

	return false
}

// SetStatusFrom gets a reference to the given string and assigns it to the StatusFrom field.
func (o *InlineResponse2003) SetStatusFrom(v string) {
	o.StatusFrom = &v
}

// GetStatusTo returns the StatusTo field value if set, zero value otherwise.
func (o *InlineResponse2003) GetStatusTo() string {
	if o == nil || o.StatusTo == nil {
		var ret string
		return ret
	}
	return *o.StatusTo
}

// GetStatusToOk returns a tuple with the StatusTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetStatusToOk() (*string, bool) {
	if o == nil || o.StatusTo == nil {
		return nil, false
	}
	return o.StatusTo, true
}

// HasStatusTo returns a boolean if a field has been set.
func (o *InlineResponse2003) HasStatusTo() bool {
	if o != nil && o.StatusTo != nil {
		return true
	}

	return false
}

// SetStatusTo gets a reference to the given string and assigns it to the StatusTo field.
func (o *InlineResponse2003) SetStatusTo(v string) {
	o.StatusTo = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *InlineResponse2003) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *InlineResponse2003) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *InlineResponse2003) SetReason(v string) {
	o.Reason = &v
}

// GetAt returns the At field value if set, zero value otherwise.
func (o *InlineResponse2003) GetAt() time.Time {
	if o == nil || o.At == nil {
		var ret time.Time
		return ret
	}
	return *o.At
}

// GetAtOk returns a tuple with the At field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetAtOk() (*time.Time, bool) {
	if o == nil || o.At == nil {
		return nil, false
	}
	return o.At, true
}

// HasAt returns a boolean if a field has been set.
func (o *InlineResponse2003) HasAt() bool {
	if o != nil && o.At != nil {
		return true
	}

	return false
}

// SetAt gets a reference to the given time.Time and assigns it to the At field.
func (o *InlineResponse2003) SetAt(v time.Time) {
	o.At = &v
}

func (o InlineResponse2003) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.AccountNumber != nil {
		toSerialize["account_number"] = o.AccountNumber
	}
	if o.StatusFrom != nil {
		toSerialize["status_from"] = o.StatusFrom
	}
	if o.StatusTo != nil {
		toSerialize["status_to"] = o.StatusTo
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.At != nil {
		toSerialize["at"] = o.At
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2003 struct {
	value *InlineResponse2003
	isSet bool
}

func (v NullableInlineResponse2003) Get() *InlineResponse2003 {
	return v.value
}

func (v *NullableInlineResponse2003) Set(val *InlineResponse2003) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2003) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2003) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2003(val *InlineResponse2003) *NullableInlineResponse2003 {
	return &NullableInlineResponse2003{value: val, isSet: true}
}

func (v NullableInlineResponse2003) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2003) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


