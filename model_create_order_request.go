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

// CreateOrderRequest struct for CreateOrderRequest
type CreateOrderRequest struct {
	Symbol string `json:"symbol"`
	Qty *float64 `json:"qty,omitempty"`
	Notional *float64 `json:"notional,omitempty"`
	Side string `json:"side"`
	Type string `json:"type"`
	TimeInForce string `json:"time_in_force"`
	LimitPrice *float64 `json:"limit_price,omitempty"`
	StopPrice *float64 `json:"stop_price,omitempty"`
	TrailPrice *float64 `json:"trail_price,omitempty"`
	TrailPercent *float64 `json:"trail_percent,omitempty"`
	ExtendedHours *bool `json:"extended_hours,omitempty"`
	ClientOrderId *string `json:"client_order_id,omitempty"`
	OrderClass *string `json:"order_class,omitempty"`
	TakeProfit *CreateOrderRequestTakeProfit `json:"take_profit,omitempty"`
	StopLoss *CreateOrderRequestStopLoss `json:"stop_loss,omitempty"`
	Commission *float64 `json:"commission,omitempty"`
}

// NewCreateOrderRequest instantiates a new CreateOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrderRequest(symbol string, side string, type_ string, timeInForce string) *CreateOrderRequest {
	this := CreateOrderRequest{}
	this.Symbol = symbol
	this.Side = side
	this.Type = type_
	this.TimeInForce = timeInForce
	return &this
}

// NewCreateOrderRequestWithDefaults instantiates a new CreateOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrderRequestWithDefaults() *CreateOrderRequest {
	this := CreateOrderRequest{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *CreateOrderRequest) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetSymbolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *CreateOrderRequest) SetSymbol(v string) {
	o.Symbol = v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *CreateOrderRequest) GetQty() float64 {
	if o == nil || o.Qty == nil {
		var ret float64
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetQtyOk() (*float64, bool) {
	if o == nil || o.Qty == nil {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *CreateOrderRequest) HasQty() bool {
	if o != nil && o.Qty != nil {
		return true
	}

	return false
}

// SetQty gets a reference to the given float64 and assigns it to the Qty field.
func (o *CreateOrderRequest) SetQty(v float64) {
	o.Qty = &v
}

// GetNotional returns the Notional field value if set, zero value otherwise.
func (o *CreateOrderRequest) GetNotional() float64 {
	if o == nil || o.Notional == nil {
		var ret float64
		return ret
	}
	return *o.Notional
}

// GetNotionalOk returns a tuple with the Notional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetNotionalOk() (*float64, bool) {
	if o == nil || o.Notional == nil {
		return nil, false
	}
	return o.Notional, true
}

// HasNotional returns a boolean if a field has been set.
func (o *CreateOrderRequest) HasNotional() bool {
	if o != nil && o.Notional != nil {
		return true
	}

	return false
}

// SetNotional gets a reference to the given float64 and assigns it to the Notional field.
func (o *CreateOrderRequest) SetNotional(v float64) {
	o.Notional = &v
}

// GetSide returns the Side field value
func (o *CreateOrderRequest) GetSide() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Side
}

// GetSideOk returns a tuple with the Side field value
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetSideOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Side, true
}

// SetSide sets field value
func (o *CreateOrderRequest) SetSide(v string) {
	o.Side = v
}

// GetType returns the Type field value
func (o *CreateOrderRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateOrderRequest) SetType(v string) {
	o.Type = v
}

// GetTimeInForce returns the TimeInForce field value
func (o *CreateOrderRequest) GetTimeInForce() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetTimeInForceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TimeInForce, true
}

// SetTimeInForce sets field value
func (o *CreateOrderRequest) SetTimeInForce(v string) {
	o.TimeInForce = v
}

// GetLimitPrice returns the LimitPrice field value if set, zero value otherwise.
func (o *CreateOrderRequest) GetLimitPrice() float64 {
	if o == nil || o.LimitPrice == nil {
		var ret float64
		return ret
	}
	return *o.LimitPrice
}

// GetLimitPriceOk returns a tuple with the LimitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetLimitPriceOk() (*float64, bool) {
	if o == nil || o.LimitPrice == nil {
		return nil, false
	}
	return o.LimitPrice, true
}

// HasLimitPrice returns a boolean if a field has been set.
func (o *CreateOrderRequest) HasLimitPrice() bool {
	if o != nil && o.LimitPrice != nil {
		return true
	}

	return false
}

// SetLimitPrice gets a reference to the given float64 and assigns it to the LimitPrice field.
func (o *CreateOrderRequest) SetLimitPrice(v float64) {
	o.LimitPrice = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *CreateOrderRequest) GetStopPrice() float64 {
	if o == nil || o.StopPrice == nil {
		var ret float64
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetStopPriceOk() (*float64, bool) {
	if o == nil || o.StopPrice == nil {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *CreateOrderRequest) HasStopPrice() bool {
	if o != nil && o.StopPrice != nil {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given float64 and assigns it to the StopPrice field.
func (o *CreateOrderRequest) SetStopPrice(v float64) {
	o.StopPrice = &v
}

// GetTrailPrice returns the TrailPrice field value if set, zero value otherwise.
func (o *CreateOrderRequest) GetTrailPrice() float64 {
	if o == nil || o.TrailPrice == nil {
		var ret float64
		return ret
	}
	return *o.TrailPrice
}

// GetTrailPriceOk returns a tuple with the TrailPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetTrailPriceOk() (*float64, bool) {
	if o == nil || o.TrailPrice == nil {
		return nil, false
	}
	return o.TrailPrice, true
}

// HasTrailPrice returns a boolean if a field has been set.
func (o *CreateOrderRequest) HasTrailPrice() bool {
	if o != nil && o.TrailPrice != nil {
		return true
	}

	return false
}

// SetTrailPrice gets a reference to the given float64 and assigns it to the TrailPrice field.
func (o *CreateOrderRequest) SetTrailPrice(v float64) {
	o.TrailPrice = &v
}

// GetTrailPercent returns the TrailPercent field value if set, zero value otherwise.
func (o *CreateOrderRequest) GetTrailPercent() float64 {
	if o == nil || o.TrailPercent == nil {
		var ret float64
		return ret
	}
	return *o.TrailPercent
}

// GetTrailPercentOk returns a tuple with the TrailPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetTrailPercentOk() (*float64, bool) {
	if o == nil || o.TrailPercent == nil {
		return nil, false
	}
	return o.TrailPercent, true
}

// HasTrailPercent returns a boolean if a field has been set.
func (o *CreateOrderRequest) HasTrailPercent() bool {
	if o != nil && o.TrailPercent != nil {
		return true
	}

	return false
}

// SetTrailPercent gets a reference to the given float64 and assigns it to the TrailPercent field.
func (o *CreateOrderRequest) SetTrailPercent(v float64) {
	o.TrailPercent = &v
}

// GetExtendedHours returns the ExtendedHours field value if set, zero value otherwise.
func (o *CreateOrderRequest) GetExtendedHours() bool {
	if o == nil || o.ExtendedHours == nil {
		var ret bool
		return ret
	}
	return *o.ExtendedHours
}

// GetExtendedHoursOk returns a tuple with the ExtendedHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetExtendedHoursOk() (*bool, bool) {
	if o == nil || o.ExtendedHours == nil {
		return nil, false
	}
	return o.ExtendedHours, true
}

// HasExtendedHours returns a boolean if a field has been set.
func (o *CreateOrderRequest) HasExtendedHours() bool {
	if o != nil && o.ExtendedHours != nil {
		return true
	}

	return false
}

// SetExtendedHours gets a reference to the given bool and assigns it to the ExtendedHours field.
func (o *CreateOrderRequest) SetExtendedHours(v bool) {
	o.ExtendedHours = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *CreateOrderRequest) GetClientOrderId() string {
	if o == nil || o.ClientOrderId == nil {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetClientOrderIdOk() (*string, bool) {
	if o == nil || o.ClientOrderId == nil {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *CreateOrderRequest) HasClientOrderId() bool {
	if o != nil && o.ClientOrderId != nil {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *CreateOrderRequest) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetOrderClass returns the OrderClass field value if set, zero value otherwise.
func (o *CreateOrderRequest) GetOrderClass() string {
	if o == nil || o.OrderClass == nil {
		var ret string
		return ret
	}
	return *o.OrderClass
}

// GetOrderClassOk returns a tuple with the OrderClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetOrderClassOk() (*string, bool) {
	if o == nil || o.OrderClass == nil {
		return nil, false
	}
	return o.OrderClass, true
}

// HasOrderClass returns a boolean if a field has been set.
func (o *CreateOrderRequest) HasOrderClass() bool {
	if o != nil && o.OrderClass != nil {
		return true
	}

	return false
}

// SetOrderClass gets a reference to the given string and assigns it to the OrderClass field.
func (o *CreateOrderRequest) SetOrderClass(v string) {
	o.OrderClass = &v
}

// GetTakeProfit returns the TakeProfit field value if set, zero value otherwise.
func (o *CreateOrderRequest) GetTakeProfit() CreateOrderRequestTakeProfit {
	if o == nil || o.TakeProfit == nil {
		var ret CreateOrderRequestTakeProfit
		return ret
	}
	return *o.TakeProfit
}

// GetTakeProfitOk returns a tuple with the TakeProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetTakeProfitOk() (*CreateOrderRequestTakeProfit, bool) {
	if o == nil || o.TakeProfit == nil {
		return nil, false
	}
	return o.TakeProfit, true
}

// HasTakeProfit returns a boolean if a field has been set.
func (o *CreateOrderRequest) HasTakeProfit() bool {
	if o != nil && o.TakeProfit != nil {
		return true
	}

	return false
}

// SetTakeProfit gets a reference to the given CreateOrderRequestTakeProfit and assigns it to the TakeProfit field.
func (o *CreateOrderRequest) SetTakeProfit(v CreateOrderRequestTakeProfit) {
	o.TakeProfit = &v
}

// GetStopLoss returns the StopLoss field value if set, zero value otherwise.
func (o *CreateOrderRequest) GetStopLoss() CreateOrderRequestStopLoss {
	if o == nil || o.StopLoss == nil {
		var ret CreateOrderRequestStopLoss
		return ret
	}
	return *o.StopLoss
}

// GetStopLossOk returns a tuple with the StopLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetStopLossOk() (*CreateOrderRequestStopLoss, bool) {
	if o == nil || o.StopLoss == nil {
		return nil, false
	}
	return o.StopLoss, true
}

// HasStopLoss returns a boolean if a field has been set.
func (o *CreateOrderRequest) HasStopLoss() bool {
	if o != nil && o.StopLoss != nil {
		return true
	}

	return false
}

// SetStopLoss gets a reference to the given CreateOrderRequestStopLoss and assigns it to the StopLoss field.
func (o *CreateOrderRequest) SetStopLoss(v CreateOrderRequestStopLoss) {
	o.StopLoss = &v
}

// GetCommission returns the Commission field value if set, zero value otherwise.
func (o *CreateOrderRequest) GetCommission() float64 {
	if o == nil || o.Commission == nil {
		var ret float64
		return ret
	}
	return *o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetCommissionOk() (*float64, bool) {
	if o == nil || o.Commission == nil {
		return nil, false
	}
	return o.Commission, true
}

// HasCommission returns a boolean if a field has been set.
func (o *CreateOrderRequest) HasCommission() bool {
	if o != nil && o.Commission != nil {
		return true
	}

	return false
}

// SetCommission gets a reference to the given float64 and assigns it to the Commission field.
func (o *CreateOrderRequest) SetCommission(v float64) {
	o.Commission = &v
}

func (o CreateOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Qty != nil {
		toSerialize["qty"] = o.Qty
	}
	if o.Notional != nil {
		toSerialize["notional"] = o.Notional
	}
	if true {
		toSerialize["side"] = o.Side
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["time_in_force"] = o.TimeInForce
	}
	if o.LimitPrice != nil {
		toSerialize["limit_price"] = o.LimitPrice
	}
	if o.StopPrice != nil {
		toSerialize["stop_price"] = o.StopPrice
	}
	if o.TrailPrice != nil {
		toSerialize["trail_price"] = o.TrailPrice
	}
	if o.TrailPercent != nil {
		toSerialize["trail_percent"] = o.TrailPercent
	}
	if o.ExtendedHours != nil {
		toSerialize["extended_hours"] = o.ExtendedHours
	}
	if o.ClientOrderId != nil {
		toSerialize["client_order_id"] = o.ClientOrderId
	}
	if o.OrderClass != nil {
		toSerialize["order_class"] = o.OrderClass
	}
	if o.TakeProfit != nil {
		toSerialize["take_profit"] = o.TakeProfit
	}
	if o.StopLoss != nil {
		toSerialize["stop_loss"] = o.StopLoss
	}
	if o.Commission != nil {
		toSerialize["commission"] = o.Commission
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrderRequest struct {
	value *CreateOrderRequest
	isSet bool
}

func (v NullableCreateOrderRequest) Get() *CreateOrderRequest {
	return v.value
}

func (v *NullableCreateOrderRequest) Set(val *CreateOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrderRequest(val *CreateOrderRequest) *NullableCreateOrderRequest {
	return &NullableCreateOrderRequest{value: val, isSet: true}
}

func (v NullableCreateOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

