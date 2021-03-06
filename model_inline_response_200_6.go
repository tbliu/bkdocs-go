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

// InlineResponse2006 struct for InlineResponse2006
type InlineResponse2006 struct {
	AccessToken *string `json:"access_token,omitempty"`
	// constant `Bearer`
	TokenType *string `json:"token_type,omitempty"`
	// token's scope
	Scope *string `json:"scope,omitempty"`
}

// NewInlineResponse2006 instantiates a new InlineResponse2006 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2006() *InlineResponse2006 {
	this := InlineResponse2006{}
	return &this
}

// NewInlineResponse2006WithDefaults instantiates a new InlineResponse2006 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2006WithDefaults() *InlineResponse2006 {
	this := InlineResponse2006{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *InlineResponse2006) GetAccessToken() string {
	if o == nil || o.AccessToken == nil {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006) GetAccessTokenOk() (*string, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *InlineResponse2006) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *InlineResponse2006) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *InlineResponse2006) GetTokenType() string {
	if o == nil || o.TokenType == nil {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006) GetTokenTypeOk() (*string, bool) {
	if o == nil || o.TokenType == nil {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *InlineResponse2006) HasTokenType() bool {
	if o != nil && o.TokenType != nil {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *InlineResponse2006) SetTokenType(v string) {
	o.TokenType = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *InlineResponse2006) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *InlineResponse2006) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *InlineResponse2006) SetScope(v string) {
	o.Scope = &v
}

func (o InlineResponse2006) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessToken != nil {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.TokenType != nil {
		toSerialize["token_type"] = o.TokenType
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2006 struct {
	value *InlineResponse2006
	isSet bool
}

func (v NullableInlineResponse2006) Get() *InlineResponse2006 {
	return v.value
}

func (v *NullableInlineResponse2006) Set(val *InlineResponse2006) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2006) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2006) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2006(val *InlineResponse2006) *NullableInlineResponse2006 {
	return &NullableInlineResponse2006{value: val, isSet: true}
}

func (v NullableInlineResponse2006) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2006) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


