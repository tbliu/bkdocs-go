/*
Alpaca Broker API

Open brokerage accounts, enable commission-free trading, and manage the ongoing user experience with Alpaca Broker API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ActivityItem struct for ActivityItem
type ActivityItem struct {
	NonTradeActivity *NonTradeActivity
	TradeActivity *TradeActivity
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ActivityItem) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NonTradeActivity
	err = json.Unmarshal(data, &dst.NonTradeActivity);
	if err == nil {
		jsonNonTradeActivity, _ := json.Marshal(dst.NonTradeActivity)
		if string(jsonNonTradeActivity) == "{}" { // empty struct
			dst.NonTradeActivity = nil
		} else {
			return nil // data stored in dst.NonTradeActivity, return on the first match
		}
	} else {
		dst.NonTradeActivity = nil
	}

	// try to unmarshal JSON data into TradeActivity
	err = json.Unmarshal(data, &dst.TradeActivity);
	if err == nil {
		jsonTradeActivity, _ := json.Marshal(dst.TradeActivity)
		if string(jsonTradeActivity) == "{}" { // empty struct
			dst.TradeActivity = nil
		} else {
			return nil // data stored in dst.TradeActivity, return on the first match
		}
	} else {
		dst.TradeActivity = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(ActivityItem)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ActivityItem) MarshalJSON() ([]byte, error) {
	if src.NonTradeActivity != nil {
		return json.Marshal(&src.NonTradeActivity)
	}

	if src.TradeActivity != nil {
		return json.Marshal(&src.TradeActivity)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableActivityItem struct {
	value *ActivityItem
	isSet bool
}

func (v NullableActivityItem) Get() *ActivityItem {
	return v.value
}

func (v *NullableActivityItem) Set(val *ActivityItem) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityItem) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityItem(val *ActivityItem) *NullableActivityItem {
	return &NullableActivityItem{value: val, isSet: true}
}

func (v NullableActivityItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

