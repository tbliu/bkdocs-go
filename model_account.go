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

// Account struct for Account
type Account struct {
	Id *string `json:"id,omitempty"`
	AccountNumber *string `json:"account_number,omitempty"`
	Status *AccountStatus `json:"status,omitempty"`
	// Always \"USD\"
	Currency *string `json:"currency,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	LastEquity *float64 `json:"last_equity,omitempty"`
	KycResults *KycResult `json:"kyc_results,omitempty"`
}

// NewAccount instantiates a new Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount() *Account {
	this := Account{}
	return &this
}

// NewAccountWithDefaults instantiates a new Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithDefaults() *Account {
	this := Account{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Account) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Account) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Account) SetId(v string) {
	o.Id = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *Account) GetAccountNumber() string {
	if o == nil || o.AccountNumber == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetAccountNumberOk() (*string, bool) {
	if o == nil || o.AccountNumber == nil {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *Account) HasAccountNumber() bool {
	if o != nil && o.AccountNumber != nil {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *Account) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Account) GetStatus() AccountStatus {
	if o == nil || o.Status == nil {
		var ret AccountStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetStatusOk() (*AccountStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Account) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AccountStatus and assigns it to the Status field.
func (o *Account) SetStatus(v AccountStatus) {
	o.Status = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Account) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Account) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Account) SetCurrency(v string) {
	o.Currency = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Account) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Account) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Account) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastEquity returns the LastEquity field value if set, zero value otherwise.
func (o *Account) GetLastEquity() float64 {
	if o == nil || o.LastEquity == nil {
		var ret float64
		return ret
	}
	return *o.LastEquity
}

// GetLastEquityOk returns a tuple with the LastEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetLastEquityOk() (*float64, bool) {
	if o == nil || o.LastEquity == nil {
		return nil, false
	}
	return o.LastEquity, true
}

// HasLastEquity returns a boolean if a field has been set.
func (o *Account) HasLastEquity() bool {
	if o != nil && o.LastEquity != nil {
		return true
	}

	return false
}

// SetLastEquity gets a reference to the given float64 and assigns it to the LastEquity field.
func (o *Account) SetLastEquity(v float64) {
	o.LastEquity = &v
}

// GetKycResults returns the KycResults field value if set, zero value otherwise.
func (o *Account) GetKycResults() KycResult {
	if o == nil || o.KycResults == nil {
		var ret KycResult
		return ret
	}
	return *o.KycResults
}

// GetKycResultsOk returns a tuple with the KycResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetKycResultsOk() (*KycResult, bool) {
	if o == nil || o.KycResults == nil {
		return nil, false
	}
	return o.KycResults, true
}

// HasKycResults returns a boolean if a field has been set.
func (o *Account) HasKycResults() bool {
	if o != nil && o.KycResults != nil {
		return true
	}

	return false
}

// SetKycResults gets a reference to the given KycResult and assigns it to the KycResults field.
func (o *Account) SetKycResults(v KycResult) {
	o.KycResults = &v
}

func (o Account) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AccountNumber != nil {
		toSerialize["account_number"] = o.AccountNumber
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastEquity != nil {
		toSerialize["last_equity"] = o.LastEquity
	}
	if o.KycResults != nil {
		toSerialize["kyc_results"] = o.KycResults
	}
	return json.Marshal(toSerialize)
}

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


