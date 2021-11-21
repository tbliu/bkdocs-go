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

// JNLC Journal information specific to cash transfers. This field is required for `Journal`s with an `entry_type` of `jnlc` (cash transfers), but will be null for those with `jnls` (securities transfers).
type JNLC struct {
	// ID the amount goes to. Only valid for JNLC journals. Null for JNLS.
	Description *string `json:"description,omitempty"`
	// Only valid for JNLC journals. Null for JNLS.
	NetAmount float64 `json:"net_amount"`
	// Only valid for JNLC journals. Null for JNLS. Max 255 characters.
	TransmitterName *string `json:"transmitter_name,omitempty"`
	// Only valid for JNLC journals. Null for JNLS.max 255 characters
	TransmitterAccountNumber *string `json:"transmitter_account_number,omitempty"`
	// Only valid for JNLC journals. Null for JNLS.max 255 characters
	TransmitterAddress *string `json:"transmitter_address,omitempty"`
	// Only valid for JNLC journals. Null for JNLS.max 255 characters
	TransmitterFinancialInstitution *string `json:"transmitter_financial_institution,omitempty"`
	// Only valid for JNLC journals. Null for JNLS.
	TransmitterTimestamp *time.Time `json:"transmitter_timestamp,omitempty"`
}

// NewJNLC instantiates a new JNLC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJNLC(netAmount float64) *JNLC {
	this := JNLC{}
	this.NetAmount = netAmount
	return &this
}

// NewJNLCWithDefaults instantiates a new JNLC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJNLCWithDefaults() *JNLC {
	this := JNLC{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JNLC) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JNLC) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JNLC) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JNLC) SetDescription(v string) {
	o.Description = &v
}

// GetNetAmount returns the NetAmount field value
func (o *JNLC) GetNetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.NetAmount
}

// GetNetAmountOk returns a tuple with the NetAmount field value
// and a boolean to check if the value has been set.
func (o *JNLC) GetNetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NetAmount, true
}

// SetNetAmount sets field value
func (o *JNLC) SetNetAmount(v float64) {
	o.NetAmount = v
}

// GetTransmitterName returns the TransmitterName field value if set, zero value otherwise.
func (o *JNLC) GetTransmitterName() string {
	if o == nil || o.TransmitterName == nil {
		var ret string
		return ret
	}
	return *o.TransmitterName
}

// GetTransmitterNameOk returns a tuple with the TransmitterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JNLC) GetTransmitterNameOk() (*string, bool) {
	if o == nil || o.TransmitterName == nil {
		return nil, false
	}
	return o.TransmitterName, true
}

// HasTransmitterName returns a boolean if a field has been set.
func (o *JNLC) HasTransmitterName() bool {
	if o != nil && o.TransmitterName != nil {
		return true
	}

	return false
}

// SetTransmitterName gets a reference to the given string and assigns it to the TransmitterName field.
func (o *JNLC) SetTransmitterName(v string) {
	o.TransmitterName = &v
}

// GetTransmitterAccountNumber returns the TransmitterAccountNumber field value if set, zero value otherwise.
func (o *JNLC) GetTransmitterAccountNumber() string {
	if o == nil || o.TransmitterAccountNumber == nil {
		var ret string
		return ret
	}
	return *o.TransmitterAccountNumber
}

// GetTransmitterAccountNumberOk returns a tuple with the TransmitterAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JNLC) GetTransmitterAccountNumberOk() (*string, bool) {
	if o == nil || o.TransmitterAccountNumber == nil {
		return nil, false
	}
	return o.TransmitterAccountNumber, true
}

// HasTransmitterAccountNumber returns a boolean if a field has been set.
func (o *JNLC) HasTransmitterAccountNumber() bool {
	if o != nil && o.TransmitterAccountNumber != nil {
		return true
	}

	return false
}

// SetTransmitterAccountNumber gets a reference to the given string and assigns it to the TransmitterAccountNumber field.
func (o *JNLC) SetTransmitterAccountNumber(v string) {
	o.TransmitterAccountNumber = &v
}

// GetTransmitterAddress returns the TransmitterAddress field value if set, zero value otherwise.
func (o *JNLC) GetTransmitterAddress() string {
	if o == nil || o.TransmitterAddress == nil {
		var ret string
		return ret
	}
	return *o.TransmitterAddress
}

// GetTransmitterAddressOk returns a tuple with the TransmitterAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JNLC) GetTransmitterAddressOk() (*string, bool) {
	if o == nil || o.TransmitterAddress == nil {
		return nil, false
	}
	return o.TransmitterAddress, true
}

// HasTransmitterAddress returns a boolean if a field has been set.
func (o *JNLC) HasTransmitterAddress() bool {
	if o != nil && o.TransmitterAddress != nil {
		return true
	}

	return false
}

// SetTransmitterAddress gets a reference to the given string and assigns it to the TransmitterAddress field.
func (o *JNLC) SetTransmitterAddress(v string) {
	o.TransmitterAddress = &v
}

// GetTransmitterFinancialInstitution returns the TransmitterFinancialInstitution field value if set, zero value otherwise.
func (o *JNLC) GetTransmitterFinancialInstitution() string {
	if o == nil || o.TransmitterFinancialInstitution == nil {
		var ret string
		return ret
	}
	return *o.TransmitterFinancialInstitution
}

// GetTransmitterFinancialInstitutionOk returns a tuple with the TransmitterFinancialInstitution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JNLC) GetTransmitterFinancialInstitutionOk() (*string, bool) {
	if o == nil || o.TransmitterFinancialInstitution == nil {
		return nil, false
	}
	return o.TransmitterFinancialInstitution, true
}

// HasTransmitterFinancialInstitution returns a boolean if a field has been set.
func (o *JNLC) HasTransmitterFinancialInstitution() bool {
	if o != nil && o.TransmitterFinancialInstitution != nil {
		return true
	}

	return false
}

// SetTransmitterFinancialInstitution gets a reference to the given string and assigns it to the TransmitterFinancialInstitution field.
func (o *JNLC) SetTransmitterFinancialInstitution(v string) {
	o.TransmitterFinancialInstitution = &v
}

// GetTransmitterTimestamp returns the TransmitterTimestamp field value if set, zero value otherwise.
func (o *JNLC) GetTransmitterTimestamp() time.Time {
	if o == nil || o.TransmitterTimestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.TransmitterTimestamp
}

// GetTransmitterTimestampOk returns a tuple with the TransmitterTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JNLC) GetTransmitterTimestampOk() (*time.Time, bool) {
	if o == nil || o.TransmitterTimestamp == nil {
		return nil, false
	}
	return o.TransmitterTimestamp, true
}

// HasTransmitterTimestamp returns a boolean if a field has been set.
func (o *JNLC) HasTransmitterTimestamp() bool {
	if o != nil && o.TransmitterTimestamp != nil {
		return true
	}

	return false
}

// SetTransmitterTimestamp gets a reference to the given time.Time and assigns it to the TransmitterTimestamp field.
func (o *JNLC) SetTransmitterTimestamp(v time.Time) {
	o.TransmitterTimestamp = &v
}

func (o JNLC) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["net_amount"] = o.NetAmount
	}
	if o.TransmitterName != nil {
		toSerialize["transmitter_name"] = o.TransmitterName
	}
	if o.TransmitterAccountNumber != nil {
		toSerialize["transmitter_account_number"] = o.TransmitterAccountNumber
	}
	if o.TransmitterAddress != nil {
		toSerialize["transmitter_address"] = o.TransmitterAddress
	}
	if o.TransmitterFinancialInstitution != nil {
		toSerialize["transmitter_financial_institution"] = o.TransmitterFinancialInstitution
	}
	if o.TransmitterTimestamp != nil {
		toSerialize["transmitter_timestamp"] = o.TransmitterTimestamp
	}
	return json.Marshal(toSerialize)
}

type NullableJNLC struct {
	value *JNLC
	isSet bool
}

func (v NullableJNLC) Get() *JNLC {
	return v.value
}

func (v *NullableJNLC) Set(val *JNLC) {
	v.value = val
	v.isSet = true
}

func (v NullableJNLC) IsSet() bool {
	return v.isSet
}

func (v *NullableJNLC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJNLC(val *JNLC) *NullableJNLC {
	return &NullableJNLC{value: val, isSet: true}
}

func (v NullableJNLC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJNLC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

