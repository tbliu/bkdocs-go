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

// AccountDocument If an account has documents on the application submission, it has the ApplicationDocument model in exchange with DocumentUploadRequest. 
type AccountDocument struct {
	Id string `json:"id"`
	DocumentType DocumentType `json:"document_type"`
	DocumentSubType *string `json:"document_sub_type,omitempty"`
	MimeType *string `json:"mime_type,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// NewAccountDocument instantiates a new AccountDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountDocument(id string, documentType DocumentType, createdAt time.Time) *AccountDocument {
	this := AccountDocument{}
	this.Id = id
	this.DocumentType = documentType
	this.CreatedAt = createdAt
	return &this
}

// NewAccountDocumentWithDefaults instantiates a new AccountDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountDocumentWithDefaults() *AccountDocument {
	this := AccountDocument{}
	return &this
}

// GetId returns the Id field value
func (o *AccountDocument) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountDocument) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccountDocument) SetId(v string) {
	o.Id = v
}

// GetDocumentType returns the DocumentType field value
func (o *AccountDocument) GetDocumentType() DocumentType {
	if o == nil {
		var ret DocumentType
		return ret
	}

	return o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value
// and a boolean to check if the value has been set.
func (o *AccountDocument) GetDocumentTypeOk() (*DocumentType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentType, true
}

// SetDocumentType sets field value
func (o *AccountDocument) SetDocumentType(v DocumentType) {
	o.DocumentType = v
}

// GetDocumentSubType returns the DocumentSubType field value if set, zero value otherwise.
func (o *AccountDocument) GetDocumentSubType() string {
	if o == nil || o.DocumentSubType == nil {
		var ret string
		return ret
	}
	return *o.DocumentSubType
}

// GetDocumentSubTypeOk returns a tuple with the DocumentSubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocument) GetDocumentSubTypeOk() (*string, bool) {
	if o == nil || o.DocumentSubType == nil {
		return nil, false
	}
	return o.DocumentSubType, true
}

// HasDocumentSubType returns a boolean if a field has been set.
func (o *AccountDocument) HasDocumentSubType() bool {
	if o != nil && o.DocumentSubType != nil {
		return true
	}

	return false
}

// SetDocumentSubType gets a reference to the given string and assigns it to the DocumentSubType field.
func (o *AccountDocument) SetDocumentSubType(v string) {
	o.DocumentSubType = &v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *AccountDocument) GetMimeType() string {
	if o == nil || o.MimeType == nil {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocument) GetMimeTypeOk() (*string, bool) {
	if o == nil || o.MimeType == nil {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *AccountDocument) HasMimeType() bool {
	if o != nil && o.MimeType != nil {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *AccountDocument) SetMimeType(v string) {
	o.MimeType = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AccountDocument) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AccountDocument) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AccountDocument) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o AccountDocument) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["document_type"] = o.DocumentType
	}
	if o.DocumentSubType != nil {
		toSerialize["document_sub_type"] = o.DocumentSubType
	}
	if o.MimeType != nil {
		toSerialize["mime_type"] = o.MimeType
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableAccountDocument struct {
	value *AccountDocument
	isSet bool
}

func (v NullableAccountDocument) Get() *AccountDocument {
	return v.value
}

func (v *NullableAccountDocument) Set(val *AccountDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountDocument(val *AccountDocument) *NullableAccountDocument {
	return &NullableAccountDocument{value: val, isSet: true}
}

func (v NullableAccountDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

