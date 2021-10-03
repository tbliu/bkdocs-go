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

// JournalJNLS struct for JournalJNLS
type JournalJNLS struct {
	// journal ID
	Id string `json:"id"`
	// JNLS (constant)
	EntryType string `json:"entry_type"`
	// account ID the shares go from
	FromAccount string `json:"from_account"`
	// account ID the shares go to
	ToAccount string `json:"to_account"`
	SettleDate NullableString `json:"settle_date"`
	Status *string `json:"status,omitempty"`
	Symbol string `json:"symbol"`
	Qty float64 `json:"qty"`
	Price float64 `json:"price"`
}

// NewJournalJNLS instantiates a new JournalJNLS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJournalJNLS(id string, entryType string, fromAccount string, toAccount string, settleDate NullableString, symbol string, qty float64, price float64) *JournalJNLS {
	this := JournalJNLS{}
	this.Id = id
	this.EntryType = entryType
	this.FromAccount = fromAccount
	this.ToAccount = toAccount
	this.SettleDate = settleDate
	this.Symbol = symbol
	this.Qty = qty
	this.Price = price
	return &this
}

// NewJournalJNLSWithDefaults instantiates a new JournalJNLS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJournalJNLSWithDefaults() *JournalJNLS {
	this := JournalJNLS{}
	return &this
}

// GetId returns the Id field value
func (o *JournalJNLS) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JournalJNLS) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JournalJNLS) SetId(v string) {
	o.Id = v
}

// GetEntryType returns the EntryType field value
func (o *JournalJNLS) GetEntryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryType
}

// GetEntryTypeOk returns a tuple with the EntryType field value
// and a boolean to check if the value has been set.
func (o *JournalJNLS) GetEntryTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntryType, true
}

// SetEntryType sets field value
func (o *JournalJNLS) SetEntryType(v string) {
	o.EntryType = v
}

// GetFromAccount returns the FromAccount field value
func (o *JournalJNLS) GetFromAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromAccount
}

// GetFromAccountOk returns a tuple with the FromAccount field value
// and a boolean to check if the value has been set.
func (o *JournalJNLS) GetFromAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FromAccount, true
}

// SetFromAccount sets field value
func (o *JournalJNLS) SetFromAccount(v string) {
	o.FromAccount = v
}

// GetToAccount returns the ToAccount field value
func (o *JournalJNLS) GetToAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToAccount
}

// GetToAccountOk returns a tuple with the ToAccount field value
// and a boolean to check if the value has been set.
func (o *JournalJNLS) GetToAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ToAccount, true
}

// SetToAccount sets field value
func (o *JournalJNLS) SetToAccount(v string) {
	o.ToAccount = v
}

// GetSettleDate returns the SettleDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *JournalJNLS) GetSettleDate() string {
	if o == nil || o.SettleDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.SettleDate.Get()
}

// GetSettleDateOk returns a tuple with the SettleDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JournalJNLS) GetSettleDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SettleDate.Get(), o.SettleDate.IsSet()
}

// SetSettleDate sets field value
func (o *JournalJNLS) SetSettleDate(v string) {
	o.SettleDate.Set(&v)
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *JournalJNLS) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JournalJNLS) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *JournalJNLS) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *JournalJNLS) SetStatus(v string) {
	o.Status = &v
}

// GetSymbol returns the Symbol field value
func (o *JournalJNLS) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *JournalJNLS) GetSymbolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *JournalJNLS) SetSymbol(v string) {
	o.Symbol = v
}

// GetQty returns the Qty field value
func (o *JournalJNLS) GetQty() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Qty
}

// GetQtyOk returns a tuple with the Qty field value
// and a boolean to check if the value has been set.
func (o *JournalJNLS) GetQtyOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Qty, true
}

// SetQty sets field value
func (o *JournalJNLS) SetQty(v float64) {
	o.Qty = v
}

// GetPrice returns the Price field value
func (o *JournalJNLS) GetPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *JournalJNLS) GetPriceOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *JournalJNLS) SetPrice(v float64) {
	o.Price = v
}

func (o JournalJNLS) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["entry_type"] = o.EntryType
	}
	if true {
		toSerialize["from_account"] = o.FromAccount
	}
	if true {
		toSerialize["to_account"] = o.ToAccount
	}
	if true {
		toSerialize["settle_date"] = o.SettleDate.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["qty"] = o.Qty
	}
	if true {
		toSerialize["price"] = o.Price
	}
	return json.Marshal(toSerialize)
}

type NullableJournalJNLS struct {
	value *JournalJNLS
	isSet bool
}

func (v NullableJournalJNLS) Get() *JournalJNLS {
	return v.value
}

func (v *NullableJournalJNLS) Set(val *JournalJNLS) {
	v.value = val
	v.isSet = true
}

func (v NullableJournalJNLS) IsSet() bool {
	return v.isSet
}

func (v *NullableJournalJNLS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJournalJNLS(val *JournalJNLS) *NullableJournalJNLS {
	return &NullableJournalJNLS{value: val, isSet: true}
}

func (v NullableJournalJNLS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJournalJNLS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

