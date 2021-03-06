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

// Disclosures Disclosures fields denote if the account owner falls under each category defined by FINRA rule. The client has to ask questions for the end user and the values should reflect their answers.  If one of the answers is true (yes), the account goes into ACTION_REQUIRED status. 
type Disclosures struct {
	EmploymentStatus *string `json:"employment_status,omitempty"`
	EmployerName *string `json:"employer_name,omitempty"`
	EmployerAddress *string `json:"employer_address,omitempty"`
	EmploymentPosition *string `json:"employment_position,omitempty"`
	IsControlPerson bool `json:"is_control_person"`
	IsAffiliatedExchangeOrFinra bool `json:"is_affiliated_exchange_or_finra"`
	IsPoliticallyExposed bool `json:"is_politically_exposed"`
	ImmediateFamilyExposed bool `json:"immediate_family_exposed"`
}

// NewDisclosures instantiates a new Disclosures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisclosures(isControlPerson bool, isAffiliatedExchangeOrFinra bool, isPoliticallyExposed bool, immediateFamilyExposed bool) *Disclosures {
	this := Disclosures{}
	this.IsControlPerson = isControlPerson
	this.IsAffiliatedExchangeOrFinra = isAffiliatedExchangeOrFinra
	this.IsPoliticallyExposed = isPoliticallyExposed
	this.ImmediateFamilyExposed = immediateFamilyExposed
	return &this
}

// NewDisclosuresWithDefaults instantiates a new Disclosures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisclosuresWithDefaults() *Disclosures {
	this := Disclosures{}
	return &this
}

// GetEmploymentStatus returns the EmploymentStatus field value if set, zero value otherwise.
func (o *Disclosures) GetEmploymentStatus() string {
	if o == nil || o.EmploymentStatus == nil {
		var ret string
		return ret
	}
	return *o.EmploymentStatus
}

// GetEmploymentStatusOk returns a tuple with the EmploymentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disclosures) GetEmploymentStatusOk() (*string, bool) {
	if o == nil || o.EmploymentStatus == nil {
		return nil, false
	}
	return o.EmploymentStatus, true
}

// HasEmploymentStatus returns a boolean if a field has been set.
func (o *Disclosures) HasEmploymentStatus() bool {
	if o != nil && o.EmploymentStatus != nil {
		return true
	}

	return false
}

// SetEmploymentStatus gets a reference to the given string and assigns it to the EmploymentStatus field.
func (o *Disclosures) SetEmploymentStatus(v string) {
	o.EmploymentStatus = &v
}

// GetEmployerName returns the EmployerName field value if set, zero value otherwise.
func (o *Disclosures) GetEmployerName() string {
	if o == nil || o.EmployerName == nil {
		var ret string
		return ret
	}
	return *o.EmployerName
}

// GetEmployerNameOk returns a tuple with the EmployerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disclosures) GetEmployerNameOk() (*string, bool) {
	if o == nil || o.EmployerName == nil {
		return nil, false
	}
	return o.EmployerName, true
}

// HasEmployerName returns a boolean if a field has been set.
func (o *Disclosures) HasEmployerName() bool {
	if o != nil && o.EmployerName != nil {
		return true
	}

	return false
}

// SetEmployerName gets a reference to the given string and assigns it to the EmployerName field.
func (o *Disclosures) SetEmployerName(v string) {
	o.EmployerName = &v
}

// GetEmployerAddress returns the EmployerAddress field value if set, zero value otherwise.
func (o *Disclosures) GetEmployerAddress() string {
	if o == nil || o.EmployerAddress == nil {
		var ret string
		return ret
	}
	return *o.EmployerAddress
}

// GetEmployerAddressOk returns a tuple with the EmployerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disclosures) GetEmployerAddressOk() (*string, bool) {
	if o == nil || o.EmployerAddress == nil {
		return nil, false
	}
	return o.EmployerAddress, true
}

// HasEmployerAddress returns a boolean if a field has been set.
func (o *Disclosures) HasEmployerAddress() bool {
	if o != nil && o.EmployerAddress != nil {
		return true
	}

	return false
}

// SetEmployerAddress gets a reference to the given string and assigns it to the EmployerAddress field.
func (o *Disclosures) SetEmployerAddress(v string) {
	o.EmployerAddress = &v
}

// GetEmploymentPosition returns the EmploymentPosition field value if set, zero value otherwise.
func (o *Disclosures) GetEmploymentPosition() string {
	if o == nil || o.EmploymentPosition == nil {
		var ret string
		return ret
	}
	return *o.EmploymentPosition
}

// GetEmploymentPositionOk returns a tuple with the EmploymentPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disclosures) GetEmploymentPositionOk() (*string, bool) {
	if o == nil || o.EmploymentPosition == nil {
		return nil, false
	}
	return o.EmploymentPosition, true
}

// HasEmploymentPosition returns a boolean if a field has been set.
func (o *Disclosures) HasEmploymentPosition() bool {
	if o != nil && o.EmploymentPosition != nil {
		return true
	}

	return false
}

// SetEmploymentPosition gets a reference to the given string and assigns it to the EmploymentPosition field.
func (o *Disclosures) SetEmploymentPosition(v string) {
	o.EmploymentPosition = &v
}

// GetIsControlPerson returns the IsControlPerson field value
func (o *Disclosures) GetIsControlPerson() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsControlPerson
}

// GetIsControlPersonOk returns a tuple with the IsControlPerson field value
// and a boolean to check if the value has been set.
func (o *Disclosures) GetIsControlPersonOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsControlPerson, true
}

// SetIsControlPerson sets field value
func (o *Disclosures) SetIsControlPerson(v bool) {
	o.IsControlPerson = v
}

// GetIsAffiliatedExchangeOrFinra returns the IsAffiliatedExchangeOrFinra field value
func (o *Disclosures) GetIsAffiliatedExchangeOrFinra() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAffiliatedExchangeOrFinra
}

// GetIsAffiliatedExchangeOrFinraOk returns a tuple with the IsAffiliatedExchangeOrFinra field value
// and a boolean to check if the value has been set.
func (o *Disclosures) GetIsAffiliatedExchangeOrFinraOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsAffiliatedExchangeOrFinra, true
}

// SetIsAffiliatedExchangeOrFinra sets field value
func (o *Disclosures) SetIsAffiliatedExchangeOrFinra(v bool) {
	o.IsAffiliatedExchangeOrFinra = v
}

// GetIsPoliticallyExposed returns the IsPoliticallyExposed field value
func (o *Disclosures) GetIsPoliticallyExposed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPoliticallyExposed
}

// GetIsPoliticallyExposedOk returns a tuple with the IsPoliticallyExposed field value
// and a boolean to check if the value has been set.
func (o *Disclosures) GetIsPoliticallyExposedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsPoliticallyExposed, true
}

// SetIsPoliticallyExposed sets field value
func (o *Disclosures) SetIsPoliticallyExposed(v bool) {
	o.IsPoliticallyExposed = v
}

// GetImmediateFamilyExposed returns the ImmediateFamilyExposed field value
func (o *Disclosures) GetImmediateFamilyExposed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ImmediateFamilyExposed
}

// GetImmediateFamilyExposedOk returns a tuple with the ImmediateFamilyExposed field value
// and a boolean to check if the value has been set.
func (o *Disclosures) GetImmediateFamilyExposedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImmediateFamilyExposed, true
}

// SetImmediateFamilyExposed sets field value
func (o *Disclosures) SetImmediateFamilyExposed(v bool) {
	o.ImmediateFamilyExposed = v
}

func (o Disclosures) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EmploymentStatus != nil {
		toSerialize["employment_status"] = o.EmploymentStatus
	}
	if o.EmployerName != nil {
		toSerialize["employer_name"] = o.EmployerName
	}
	if o.EmployerAddress != nil {
		toSerialize["employer_address"] = o.EmployerAddress
	}
	if o.EmploymentPosition != nil {
		toSerialize["employment_position"] = o.EmploymentPosition
	}
	if true {
		toSerialize["is_control_person"] = o.IsControlPerson
	}
	if true {
		toSerialize["is_affiliated_exchange_or_finra"] = o.IsAffiliatedExchangeOrFinra
	}
	if true {
		toSerialize["is_politically_exposed"] = o.IsPoliticallyExposed
	}
	if true {
		toSerialize["immediate_family_exposed"] = o.ImmediateFamilyExposed
	}
	return json.Marshal(toSerialize)
}

type NullableDisclosures struct {
	value *Disclosures
	isSet bool
}

func (v NullableDisclosures) Get() *Disclosures {
	return v.value
}

func (v *NullableDisclosures) Set(val *Disclosures) {
	v.value = val
	v.isSet = true
}

func (v NullableDisclosures) IsSet() bool {
	return v.isSet
}

func (v *NullableDisclosures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisclosures(val *Disclosures) *NullableDisclosures {
	return &NullableDisclosures{value: val, isSet: true}
}

func (v NullableDisclosures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisclosures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


