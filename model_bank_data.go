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

// BankData struct for BankData
type BankData struct {
	Name string `json:"name"`
	BankCode string `json:"bank_code"`
	BankCodeType string `json:"bank_code_type"`
	Country *string `json:"country,omitempty"`
	StateProvince *string `json:"state_province,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
	City *string `json:"city,omitempty"`
	StreetAddress *string `json:"street_address,omitempty"`
	AccountNumber string `json:"account_number"`
}

// NewBankData instantiates a new BankData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankData(name string, bankCode string, bankCodeType string, accountNumber string) *BankData {
	this := BankData{}
	this.Name = name
	this.BankCode = bankCode
	this.BankCodeType = bankCodeType
	this.AccountNumber = accountNumber
	return &this
}

// NewBankDataWithDefaults instantiates a new BankData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankDataWithDefaults() *BankData {
	this := BankData{}
	return &this
}

// GetName returns the Name field value
func (o *BankData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BankData) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BankData) SetName(v string) {
	o.Name = v
}

// GetBankCode returns the BankCode field value
func (o *BankData) GetBankCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankCode
}

// GetBankCodeOk returns a tuple with the BankCode field value
// and a boolean to check if the value has been set.
func (o *BankData) GetBankCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankCode, true
}

// SetBankCode sets field value
func (o *BankData) SetBankCode(v string) {
	o.BankCode = v
}

// GetBankCodeType returns the BankCodeType field value
func (o *BankData) GetBankCodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankCodeType
}

// GetBankCodeTypeOk returns a tuple with the BankCodeType field value
// and a boolean to check if the value has been set.
func (o *BankData) GetBankCodeTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankCodeType, true
}

// SetBankCodeType sets field value
func (o *BankData) SetBankCodeType(v string) {
	o.BankCodeType = v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *BankData) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankData) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *BankData) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *BankData) SetCountry(v string) {
	o.Country = &v
}

// GetStateProvince returns the StateProvince field value if set, zero value otherwise.
func (o *BankData) GetStateProvince() string {
	if o == nil || o.StateProvince == nil {
		var ret string
		return ret
	}
	return *o.StateProvince
}

// GetStateProvinceOk returns a tuple with the StateProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankData) GetStateProvinceOk() (*string, bool) {
	if o == nil || o.StateProvince == nil {
		return nil, false
	}
	return o.StateProvince, true
}

// HasStateProvince returns a boolean if a field has been set.
func (o *BankData) HasStateProvince() bool {
	if o != nil && o.StateProvince != nil {
		return true
	}

	return false
}

// SetStateProvince gets a reference to the given string and assigns it to the StateProvince field.
func (o *BankData) SetStateProvince(v string) {
	o.StateProvince = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *BankData) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankData) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *BankData) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *BankData) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *BankData) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankData) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *BankData) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *BankData) SetCity(v string) {
	o.City = &v
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *BankData) GetStreetAddress() string {
	if o == nil || o.StreetAddress == nil {
		var ret string
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankData) GetStreetAddressOk() (*string, bool) {
	if o == nil || o.StreetAddress == nil {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *BankData) HasStreetAddress() bool {
	if o != nil && o.StreetAddress != nil {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given string and assigns it to the StreetAddress field.
func (o *BankData) SetStreetAddress(v string) {
	o.StreetAddress = &v
}

// GetAccountNumber returns the AccountNumber field value
func (o *BankData) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *BankData) GetAccountNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *BankData) SetAccountNumber(v string) {
	o.AccountNumber = v
}

func (o BankData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["bank_code"] = o.BankCode
	}
	if true {
		toSerialize["bank_code_type"] = o.BankCodeType
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.StateProvince != nil {
		toSerialize["state_province"] = o.StateProvince
	}
	if o.PostalCode != nil {
		toSerialize["postal_code"] = o.PostalCode
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.StreetAddress != nil {
		toSerialize["street_address"] = o.StreetAddress
	}
	if true {
		toSerialize["account_number"] = o.AccountNumber
	}
	return json.Marshal(toSerialize)
}

type NullableBankData struct {
	value *BankData
	isSet bool
}

func (v NullableBankData) Get() *BankData {
	return v.value
}

func (v *NullableBankData) Set(val *BankData) {
	v.value = val
	v.isSet = true
}

func (v NullableBankData) IsSet() bool {
	return v.isSet
}

func (v *NullableBankData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankData(val *BankData) *NullableBankData {
	return &NullableBankData{value: val, isSet: true}
}

func (v NullableBankData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


