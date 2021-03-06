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

// BatchJournalRequestEntries struct for BatchJournalRequestEntries
type BatchJournalRequestEntries struct {
	ToAccount string `json:"to_account"`
	Amount string `json:"amount"`
}

// NewBatchJournalRequestEntries instantiates a new BatchJournalRequestEntries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchJournalRequestEntries(toAccount string, amount string) *BatchJournalRequestEntries {
	this := BatchJournalRequestEntries{}
	this.ToAccount = toAccount
	this.Amount = amount
	return &this
}

// NewBatchJournalRequestEntriesWithDefaults instantiates a new BatchJournalRequestEntries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchJournalRequestEntriesWithDefaults() *BatchJournalRequestEntries {
	this := BatchJournalRequestEntries{}
	return &this
}

// GetToAccount returns the ToAccount field value
func (o *BatchJournalRequestEntries) GetToAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToAccount
}

// GetToAccountOk returns a tuple with the ToAccount field value
// and a boolean to check if the value has been set.
func (o *BatchJournalRequestEntries) GetToAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ToAccount, true
}

// SetToAccount sets field value
func (o *BatchJournalRequestEntries) SetToAccount(v string) {
	o.ToAccount = v
}

// GetAmount returns the Amount field value
func (o *BatchJournalRequestEntries) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *BatchJournalRequestEntries) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *BatchJournalRequestEntries) SetAmount(v string) {
	o.Amount = v
}

func (o BatchJournalRequestEntries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["to_account"] = o.ToAccount
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableBatchJournalRequestEntries struct {
	value *BatchJournalRequestEntries
	isSet bool
}

func (v NullableBatchJournalRequestEntries) Get() *BatchJournalRequestEntries {
	return v.value
}

func (v *NullableBatchJournalRequestEntries) Set(val *BatchJournalRequestEntries) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchJournalRequestEntries) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchJournalRequestEntries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchJournalRequestEntries(val *BatchJournalRequestEntries) *NullableBatchJournalRequestEntries {
	return &NullableBatchJournalRequestEntries{value: val, isSet: true}
}

func (v NullableBatchJournalRequestEntries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchJournalRequestEntries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


