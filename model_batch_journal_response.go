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

// BatchJournalResponse struct for BatchJournalResponse
type BatchJournalResponse struct {
	Id string `json:"id"`
	ErrorMessage string `json:"error_message"`
	EntryType string `json:"entry_type"`
	FromAccount string `json:"from_account"`
	ToAccount string `json:"to_account"`
	Symbol string `json:"symbol"`
	Qty NullableString `json:"qty"`
	Price string `json:"price"`
	Status string `json:"status"`
	SettleDate NullableString `json:"settle_date"`
	SystemDate NullableString `json:"system_date"`
	NetAmount string `json:"net_amount"`
	Description string `json:"description"`
}

// NewBatchJournalResponse instantiates a new BatchJournalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchJournalResponse(id string, errorMessage string, entryType string, fromAccount string, toAccount string, symbol string, qty NullableString, price string, status string, settleDate NullableString, systemDate NullableString, netAmount string, description string) *BatchJournalResponse {
	this := BatchJournalResponse{}
	this.Id = id
	this.ErrorMessage = errorMessage
	this.EntryType = entryType
	this.FromAccount = fromAccount
	this.ToAccount = toAccount
	this.Symbol = symbol
	this.Qty = qty
	this.Price = price
	this.Status = status
	this.SettleDate = settleDate
	this.SystemDate = systemDate
	this.NetAmount = netAmount
	this.Description = description
	return &this
}

// NewBatchJournalResponseWithDefaults instantiates a new BatchJournalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchJournalResponseWithDefaults() *BatchJournalResponse {
	this := BatchJournalResponse{}
	return &this
}

// GetId returns the Id field value
func (o *BatchJournalResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BatchJournalResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BatchJournalResponse) SetId(v string) {
	o.Id = v
}

// GetErrorMessage returns the ErrorMessage field value
func (o *BatchJournalResponse) GetErrorMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value
// and a boolean to check if the value has been set.
func (o *BatchJournalResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorMessage, true
}

// SetErrorMessage sets field value
func (o *BatchJournalResponse) SetErrorMessage(v string) {
	o.ErrorMessage = v
}

// GetEntryType returns the EntryType field value
func (o *BatchJournalResponse) GetEntryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryType
}

// GetEntryTypeOk returns a tuple with the EntryType field value
// and a boolean to check if the value has been set.
func (o *BatchJournalResponse) GetEntryTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntryType, true
}

// SetEntryType sets field value
func (o *BatchJournalResponse) SetEntryType(v string) {
	o.EntryType = v
}

// GetFromAccount returns the FromAccount field value
func (o *BatchJournalResponse) GetFromAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromAccount
}

// GetFromAccountOk returns a tuple with the FromAccount field value
// and a boolean to check if the value has been set.
func (o *BatchJournalResponse) GetFromAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FromAccount, true
}

// SetFromAccount sets field value
func (o *BatchJournalResponse) SetFromAccount(v string) {
	o.FromAccount = v
}

// GetToAccount returns the ToAccount field value
func (o *BatchJournalResponse) GetToAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToAccount
}

// GetToAccountOk returns a tuple with the ToAccount field value
// and a boolean to check if the value has been set.
func (o *BatchJournalResponse) GetToAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ToAccount, true
}

// SetToAccount sets field value
func (o *BatchJournalResponse) SetToAccount(v string) {
	o.ToAccount = v
}

// GetSymbol returns the Symbol field value
func (o *BatchJournalResponse) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *BatchJournalResponse) GetSymbolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *BatchJournalResponse) SetSymbol(v string) {
	o.Symbol = v
}

// GetQty returns the Qty field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BatchJournalResponse) GetQty() string {
	if o == nil || o.Qty.Get() == nil {
		var ret string
		return ret
	}

	return *o.Qty.Get()
}

// GetQtyOk returns a tuple with the Qty field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BatchJournalResponse) GetQtyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Qty.Get(), o.Qty.IsSet()
}

// SetQty sets field value
func (o *BatchJournalResponse) SetQty(v string) {
	o.Qty.Set(&v)
}

// GetPrice returns the Price field value
func (o *BatchJournalResponse) GetPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *BatchJournalResponse) GetPriceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *BatchJournalResponse) SetPrice(v string) {
	o.Price = v
}

// GetStatus returns the Status field value
func (o *BatchJournalResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BatchJournalResponse) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BatchJournalResponse) SetStatus(v string) {
	o.Status = v
}

// GetSettleDate returns the SettleDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BatchJournalResponse) GetSettleDate() string {
	if o == nil || o.SettleDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.SettleDate.Get()
}

// GetSettleDateOk returns a tuple with the SettleDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BatchJournalResponse) GetSettleDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SettleDate.Get(), o.SettleDate.IsSet()
}

// SetSettleDate sets field value
func (o *BatchJournalResponse) SetSettleDate(v string) {
	o.SettleDate.Set(&v)
}

// GetSystemDate returns the SystemDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BatchJournalResponse) GetSystemDate() string {
	if o == nil || o.SystemDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.SystemDate.Get()
}

// GetSystemDateOk returns a tuple with the SystemDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BatchJournalResponse) GetSystemDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SystemDate.Get(), o.SystemDate.IsSet()
}

// SetSystemDate sets field value
func (o *BatchJournalResponse) SetSystemDate(v string) {
	o.SystemDate.Set(&v)
}

// GetNetAmount returns the NetAmount field value
func (o *BatchJournalResponse) GetNetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetAmount
}

// GetNetAmountOk returns a tuple with the NetAmount field value
// and a boolean to check if the value has been set.
func (o *BatchJournalResponse) GetNetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NetAmount, true
}

// SetNetAmount sets field value
func (o *BatchJournalResponse) SetNetAmount(v string) {
	o.NetAmount = v
}

// GetDescription returns the Description field value
func (o *BatchJournalResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *BatchJournalResponse) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *BatchJournalResponse) SetDescription(v string) {
	o.Description = v
}

func (o BatchJournalResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["error_message"] = o.ErrorMessage
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
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["qty"] = o.Qty.Get()
	}
	if true {
		toSerialize["price"] = o.Price
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["settle_date"] = o.SettleDate.Get()
	}
	if true {
		toSerialize["system_date"] = o.SystemDate.Get()
	}
	if true {
		toSerialize["net_amount"] = o.NetAmount
	}
	if true {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableBatchJournalResponse struct {
	value *BatchJournalResponse
	isSet bool
}

func (v NullableBatchJournalResponse) Get() *BatchJournalResponse {
	return v.value
}

func (v *NullableBatchJournalResponse) Set(val *BatchJournalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchJournalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchJournalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchJournalResponse(val *BatchJournalResponse) *NullableBatchJournalResponse {
	return &NullableBatchJournalResponse{value: val, isSet: true}
}

func (v NullableBatchJournalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchJournalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


