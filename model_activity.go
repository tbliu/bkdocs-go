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

// Activity Base for activity types
type Activity struct {
	Id *string `json:"id,omitempty"`
	AccountId *string `json:"account_id,omitempty"`
	ActivityType *ActivityType `json:"activity_type,omitempty"`
}

// NewActivity instantiates a new Activity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivity() *Activity {
	this := Activity{}
	return &this
}

// NewActivityWithDefaults instantiates a new Activity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityWithDefaults() *Activity {
	this := Activity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Activity) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Activity) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Activity) SetId(v string) {
	o.Id = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Activity) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *Activity) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *Activity) SetAccountId(v string) {
	o.AccountId = &v
}

// GetActivityType returns the ActivityType field value if set, zero value otherwise.
func (o *Activity) GetActivityType() ActivityType {
	if o == nil || o.ActivityType == nil {
		var ret ActivityType
		return ret
	}
	return *o.ActivityType
}

// GetActivityTypeOk returns a tuple with the ActivityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetActivityTypeOk() (*ActivityType, bool) {
	if o == nil || o.ActivityType == nil {
		return nil, false
	}
	return o.ActivityType, true
}

// HasActivityType returns a boolean if a field has been set.
func (o *Activity) HasActivityType() bool {
	if o != nil && o.ActivityType != nil {
		return true
	}

	return false
}

// SetActivityType gets a reference to the given ActivityType and assigns it to the ActivityType field.
func (o *Activity) SetActivityType(v ActivityType) {
	o.ActivityType = &v
}

func (o Activity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.ActivityType != nil {
		toSerialize["activity_type"] = o.ActivityType
	}
	return json.Marshal(toSerialize)
}

type NullableActivity struct {
	value *Activity
	isSet bool
}

func (v NullableActivity) Get() *Activity {
	return v.value
}

func (v *NullableActivity) Set(val *Activity) {
	v.value = val
	v.isSet = true
}

func (v NullableActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivity(val *Activity) *NullableActivity {
	return &NullableActivity{value: val, isSet: true}
}

func (v NullableActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

