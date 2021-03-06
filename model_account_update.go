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

// AccountUpdate struct for AccountUpdate
type AccountUpdate struct {
	Contact *Contact `json:"contact,omitempty"`
	Identity *Identity `json:"identity,omitempty"`
	Disclosures *Disclosures `json:"disclosures,omitempty"`
	TrustedContact *TrustedContact `json:"trustedContact,omitempty"`
}

// NewAccountUpdate instantiates a new AccountUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUpdate() *AccountUpdate {
	this := AccountUpdate{}
	return &this
}

// NewAccountUpdateWithDefaults instantiates a new AccountUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUpdateWithDefaults() *AccountUpdate {
	this := AccountUpdate{}
	return &this
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *AccountUpdate) GetContact() Contact {
	if o == nil || o.Contact == nil {
		var ret Contact
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetContactOk() (*Contact, bool) {
	if o == nil || o.Contact == nil {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *AccountUpdate) HasContact() bool {
	if o != nil && o.Contact != nil {
		return true
	}

	return false
}

// SetContact gets a reference to the given Contact and assigns it to the Contact field.
func (o *AccountUpdate) SetContact(v Contact) {
	o.Contact = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *AccountUpdate) GetIdentity() Identity {
	if o == nil || o.Identity == nil {
		var ret Identity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetIdentityOk() (*Identity, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *AccountUpdate) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given Identity and assigns it to the Identity field.
func (o *AccountUpdate) SetIdentity(v Identity) {
	o.Identity = &v
}

// GetDisclosures returns the Disclosures field value if set, zero value otherwise.
func (o *AccountUpdate) GetDisclosures() Disclosures {
	if o == nil || o.Disclosures == nil {
		var ret Disclosures
		return ret
	}
	return *o.Disclosures
}

// GetDisclosuresOk returns a tuple with the Disclosures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetDisclosuresOk() (*Disclosures, bool) {
	if o == nil || o.Disclosures == nil {
		return nil, false
	}
	return o.Disclosures, true
}

// HasDisclosures returns a boolean if a field has been set.
func (o *AccountUpdate) HasDisclosures() bool {
	if o != nil && o.Disclosures != nil {
		return true
	}

	return false
}

// SetDisclosures gets a reference to the given Disclosures and assigns it to the Disclosures field.
func (o *AccountUpdate) SetDisclosures(v Disclosures) {
	o.Disclosures = &v
}

// GetTrustedContact returns the TrustedContact field value if set, zero value otherwise.
func (o *AccountUpdate) GetTrustedContact() TrustedContact {
	if o == nil || o.TrustedContact == nil {
		var ret TrustedContact
		return ret
	}
	return *o.TrustedContact
}

// GetTrustedContactOk returns a tuple with the TrustedContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetTrustedContactOk() (*TrustedContact, bool) {
	if o == nil || o.TrustedContact == nil {
		return nil, false
	}
	return o.TrustedContact, true
}

// HasTrustedContact returns a boolean if a field has been set.
func (o *AccountUpdate) HasTrustedContact() bool {
	if o != nil && o.TrustedContact != nil {
		return true
	}

	return false
}

// SetTrustedContact gets a reference to the given TrustedContact and assigns it to the TrustedContact field.
func (o *AccountUpdate) SetTrustedContact(v TrustedContact) {
	o.TrustedContact = &v
}

func (o AccountUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Contact != nil {
		toSerialize["contact"] = o.Contact
	}
	if o.Identity != nil {
		toSerialize["identity"] = o.Identity
	}
	if o.Disclosures != nil {
		toSerialize["disclosures"] = o.Disclosures
	}
	if o.TrustedContact != nil {
		toSerialize["trustedContact"] = o.TrustedContact
	}
	return json.Marshal(toSerialize)
}

type NullableAccountUpdate struct {
	value *AccountUpdate
	isSet bool
}

func (v NullableAccountUpdate) Get() *AccountUpdate {
	return v.value
}

func (v *NullableAccountUpdate) Set(val *AccountUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUpdate(val *AccountUpdate) *NullableAccountUpdate {
	return &NullableAccountUpdate{value: val, isSet: true}
}

func (v NullableAccountUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


