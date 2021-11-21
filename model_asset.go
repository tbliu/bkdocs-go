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

// Asset struct for Asset
type Asset struct {
	Id *string `json:"id,omitempty"`
	Class *string `json:"class,omitempty"`
	Exchange *string `json:"exchange,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
	Tradable *bool `json:"tradable,omitempty"`
	Marginable *bool `json:"marginable,omitempty"`
	Shortable *bool `json:"shortable,omitempty"`
	EasyToBorrow *bool `json:"easy_to_borrow,omitempty"`
	Fractionable *bool `json:"fractionable,omitempty"`
}

// NewAsset instantiates a new Asset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsset() *Asset {
	this := Asset{}
	return &this
}

// NewAssetWithDefaults instantiates a new Asset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWithDefaults() *Asset {
	this := Asset{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Asset) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Asset) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Asset) SetId(v string) {
	o.Id = &v
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *Asset) GetClass() string {
	if o == nil || o.Class == nil {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetClassOk() (*string, bool) {
	if o == nil || o.Class == nil {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *Asset) HasClass() bool {
	if o != nil && o.Class != nil {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *Asset) SetClass(v string) {
	o.Class = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *Asset) GetExchange() string {
	if o == nil || o.Exchange == nil {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetExchangeOk() (*string, bool) {
	if o == nil || o.Exchange == nil {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *Asset) HasExchange() bool {
	if o != nil && o.Exchange != nil {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *Asset) SetExchange(v string) {
	o.Exchange = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *Asset) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *Asset) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *Asset) SetSymbol(v string) {
	o.Symbol = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Asset) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Asset) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Asset) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Asset) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Asset) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Asset) SetStatus(v string) {
	o.Status = &v
}

// GetTradable returns the Tradable field value if set, zero value otherwise.
func (o *Asset) GetTradable() bool {
	if o == nil || o.Tradable == nil {
		var ret bool
		return ret
	}
	return *o.Tradable
}

// GetTradableOk returns a tuple with the Tradable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetTradableOk() (*bool, bool) {
	if o == nil || o.Tradable == nil {
		return nil, false
	}
	return o.Tradable, true
}

// HasTradable returns a boolean if a field has been set.
func (o *Asset) HasTradable() bool {
	if o != nil && o.Tradable != nil {
		return true
	}

	return false
}

// SetTradable gets a reference to the given bool and assigns it to the Tradable field.
func (o *Asset) SetTradable(v bool) {
	o.Tradable = &v
}

// GetMarginable returns the Marginable field value if set, zero value otherwise.
func (o *Asset) GetMarginable() bool {
	if o == nil || o.Marginable == nil {
		var ret bool
		return ret
	}
	return *o.Marginable
}

// GetMarginableOk returns a tuple with the Marginable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetMarginableOk() (*bool, bool) {
	if o == nil || o.Marginable == nil {
		return nil, false
	}
	return o.Marginable, true
}

// HasMarginable returns a boolean if a field has been set.
func (o *Asset) HasMarginable() bool {
	if o != nil && o.Marginable != nil {
		return true
	}

	return false
}

// SetMarginable gets a reference to the given bool and assigns it to the Marginable field.
func (o *Asset) SetMarginable(v bool) {
	o.Marginable = &v
}

// GetShortable returns the Shortable field value if set, zero value otherwise.
func (o *Asset) GetShortable() bool {
	if o == nil || o.Shortable == nil {
		var ret bool
		return ret
	}
	return *o.Shortable
}

// GetShortableOk returns a tuple with the Shortable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetShortableOk() (*bool, bool) {
	if o == nil || o.Shortable == nil {
		return nil, false
	}
	return o.Shortable, true
}

// HasShortable returns a boolean if a field has been set.
func (o *Asset) HasShortable() bool {
	if o != nil && o.Shortable != nil {
		return true
	}

	return false
}

// SetShortable gets a reference to the given bool and assigns it to the Shortable field.
func (o *Asset) SetShortable(v bool) {
	o.Shortable = &v
}

// GetEasyToBorrow returns the EasyToBorrow field value if set, zero value otherwise.
func (o *Asset) GetEasyToBorrow() bool {
	if o == nil || o.EasyToBorrow == nil {
		var ret bool
		return ret
	}
	return *o.EasyToBorrow
}

// GetEasyToBorrowOk returns a tuple with the EasyToBorrow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetEasyToBorrowOk() (*bool, bool) {
	if o == nil || o.EasyToBorrow == nil {
		return nil, false
	}
	return o.EasyToBorrow, true
}

// HasEasyToBorrow returns a boolean if a field has been set.
func (o *Asset) HasEasyToBorrow() bool {
	if o != nil && o.EasyToBorrow != nil {
		return true
	}

	return false
}

// SetEasyToBorrow gets a reference to the given bool and assigns it to the EasyToBorrow field.
func (o *Asset) SetEasyToBorrow(v bool) {
	o.EasyToBorrow = &v
}

// GetFractionable returns the Fractionable field value if set, zero value otherwise.
func (o *Asset) GetFractionable() bool {
	if o == nil || o.Fractionable == nil {
		var ret bool
		return ret
	}
	return *o.Fractionable
}

// GetFractionableOk returns a tuple with the Fractionable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetFractionableOk() (*bool, bool) {
	if o == nil || o.Fractionable == nil {
		return nil, false
	}
	return o.Fractionable, true
}

// HasFractionable returns a boolean if a field has been set.
func (o *Asset) HasFractionable() bool {
	if o != nil && o.Fractionable != nil {
		return true
	}

	return false
}

// SetFractionable gets a reference to the given bool and assigns it to the Fractionable field.
func (o *Asset) SetFractionable(v bool) {
	o.Fractionable = &v
}

func (o Asset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Class != nil {
		toSerialize["class"] = o.Class
	}
	if o.Exchange != nil {
		toSerialize["exchange"] = o.Exchange
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Tradable != nil {
		toSerialize["tradable"] = o.Tradable
	}
	if o.Marginable != nil {
		toSerialize["marginable"] = o.Marginable
	}
	if o.Shortable != nil {
		toSerialize["shortable"] = o.Shortable
	}
	if o.EasyToBorrow != nil {
		toSerialize["easy_to_borrow"] = o.EasyToBorrow
	}
	if o.Fractionable != nil {
		toSerialize["fractionable"] = o.Fractionable
	}
	return json.Marshal(toSerialize)
}

type NullableAsset struct {
	value *Asset
	isSet bool
}

func (v NullableAsset) Get() *Asset {
	return v.value
}

func (v *NullableAsset) Set(val *Asset) {
	v.value = val
	v.isSet = true
}

func (v NullableAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsset(val *Asset) *NullableAsset {
	return &NullableAsset{value: val, isSet: true}
}

func (v NullableAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

