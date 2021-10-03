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

// AccountExtended struct for AccountExtended
type AccountExtended struct {
	Id *string `json:"id,omitempty"`
	AccountNumber *string `json:"account_number,omitempty"`
	Status *AccountStatus `json:"status,omitempty"`
	// Always \"USD\"
	Currency *string `json:"currency,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	LastEquity *float64 `json:"last_equity,omitempty"`
	KycResults *KycResult `json:"kyc_results,omitempty"`
	Contact *Contact `json:"contact,omitempty"`
	Identity *Identity `json:"identity,omitempty"`
	Disclosures *Disclosures `json:"disclosures,omitempty"`
	// The client has to present Alpaca Account Agreement and Margin Agreement to the end user, and have them read full sentences. 
	Agreements *[]Agreement `json:"agreements,omitempty"`
	Documents *[]ApplicationDocument `json:"documents,omitempty"`
	TrustedContact *TrustedContact `json:"trusted_contact,omitempty"`
}

// NewAccountExtended instantiates a new AccountExtended object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountExtended() *AccountExtended {
	this := AccountExtended{}
	return &this
}

// NewAccountExtendedWithDefaults instantiates a new AccountExtended object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountExtendedWithDefaults() *AccountExtended {
	this := AccountExtended{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountExtended) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExtended) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountExtended) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountExtended) SetId(v string) {
	o.Id = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *AccountExtended) GetAccountNumber() string {
	if o == nil || o.AccountNumber == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExtended) GetAccountNumberOk() (*string, bool) {
	if o == nil || o.AccountNumber == nil {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *AccountExtended) HasAccountNumber() bool {
	if o != nil && o.AccountNumber != nil {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *AccountExtended) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AccountExtended) GetStatus() AccountStatus {
	if o == nil || o.Status == nil {
		var ret AccountStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExtended) GetStatusOk() (*AccountStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AccountExtended) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AccountStatus and assigns it to the Status field.
func (o *AccountExtended) SetStatus(v AccountStatus) {
	o.Status = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AccountExtended) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExtended) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AccountExtended) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AccountExtended) SetCurrency(v string) {
	o.Currency = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AccountExtended) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExtended) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AccountExtended) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AccountExtended) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastEquity returns the LastEquity field value if set, zero value otherwise.
func (o *AccountExtended) GetLastEquity() float64 {
	if o == nil || o.LastEquity == nil {
		var ret float64
		return ret
	}
	return *o.LastEquity
}

// GetLastEquityOk returns a tuple with the LastEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExtended) GetLastEquityOk() (*float64, bool) {
	if o == nil || o.LastEquity == nil {
		return nil, false
	}
	return o.LastEquity, true
}

// HasLastEquity returns a boolean if a field has been set.
func (o *AccountExtended) HasLastEquity() bool {
	if o != nil && o.LastEquity != nil {
		return true
	}

	return false
}

// SetLastEquity gets a reference to the given float64 and assigns it to the LastEquity field.
func (o *AccountExtended) SetLastEquity(v float64) {
	o.LastEquity = &v
}

// GetKycResults returns the KycResults field value if set, zero value otherwise.
func (o *AccountExtended) GetKycResults() KycResult {
	if o == nil || o.KycResults == nil {
		var ret KycResult
		return ret
	}
	return *o.KycResults
}

// GetKycResultsOk returns a tuple with the KycResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExtended) GetKycResultsOk() (*KycResult, bool) {
	if o == nil || o.KycResults == nil {
		return nil, false
	}
	return o.KycResults, true
}

// HasKycResults returns a boolean if a field has been set.
func (o *AccountExtended) HasKycResults() bool {
	if o != nil && o.KycResults != nil {
		return true
	}

	return false
}

// SetKycResults gets a reference to the given KycResult and assigns it to the KycResults field.
func (o *AccountExtended) SetKycResults(v KycResult) {
	o.KycResults = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *AccountExtended) GetContact() Contact {
	if o == nil || o.Contact == nil {
		var ret Contact
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExtended) GetContactOk() (*Contact, bool) {
	if o == nil || o.Contact == nil {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *AccountExtended) HasContact() bool {
	if o != nil && o.Contact != nil {
		return true
	}

	return false
}

// SetContact gets a reference to the given Contact and assigns it to the Contact field.
func (o *AccountExtended) SetContact(v Contact) {
	o.Contact = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *AccountExtended) GetIdentity() Identity {
	if o == nil || o.Identity == nil {
		var ret Identity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExtended) GetIdentityOk() (*Identity, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *AccountExtended) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given Identity and assigns it to the Identity field.
func (o *AccountExtended) SetIdentity(v Identity) {
	o.Identity = &v
}

// GetDisclosures returns the Disclosures field value if set, zero value otherwise.
func (o *AccountExtended) GetDisclosures() Disclosures {
	if o == nil || o.Disclosures == nil {
		var ret Disclosures
		return ret
	}
	return *o.Disclosures
}

// GetDisclosuresOk returns a tuple with the Disclosures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExtended) GetDisclosuresOk() (*Disclosures, bool) {
	if o == nil || o.Disclosures == nil {
		return nil, false
	}
	return o.Disclosures, true
}

// HasDisclosures returns a boolean if a field has been set.
func (o *AccountExtended) HasDisclosures() bool {
	if o != nil && o.Disclosures != nil {
		return true
	}

	return false
}

// SetDisclosures gets a reference to the given Disclosures and assigns it to the Disclosures field.
func (o *AccountExtended) SetDisclosures(v Disclosures) {
	o.Disclosures = &v
}

// GetAgreements returns the Agreements field value if set, zero value otherwise.
func (o *AccountExtended) GetAgreements() []Agreement {
	if o == nil || o.Agreements == nil {
		var ret []Agreement
		return ret
	}
	return *o.Agreements
}

// GetAgreementsOk returns a tuple with the Agreements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExtended) GetAgreementsOk() (*[]Agreement, bool) {
	if o == nil || o.Agreements == nil {
		return nil, false
	}
	return o.Agreements, true
}

// HasAgreements returns a boolean if a field has been set.
func (o *AccountExtended) HasAgreements() bool {
	if o != nil && o.Agreements != nil {
		return true
	}

	return false
}

// SetAgreements gets a reference to the given []Agreement and assigns it to the Agreements field.
func (o *AccountExtended) SetAgreements(v []Agreement) {
	o.Agreements = &v
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *AccountExtended) GetDocuments() []ApplicationDocument {
	if o == nil || o.Documents == nil {
		var ret []ApplicationDocument
		return ret
	}
	return *o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExtended) GetDocumentsOk() (*[]ApplicationDocument, bool) {
	if o == nil || o.Documents == nil {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *AccountExtended) HasDocuments() bool {
	if o != nil && o.Documents != nil {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []ApplicationDocument and assigns it to the Documents field.
func (o *AccountExtended) SetDocuments(v []ApplicationDocument) {
	o.Documents = &v
}

// GetTrustedContact returns the TrustedContact field value if set, zero value otherwise.
func (o *AccountExtended) GetTrustedContact() TrustedContact {
	if o == nil || o.TrustedContact == nil {
		var ret TrustedContact
		return ret
	}
	return *o.TrustedContact
}

// GetTrustedContactOk returns a tuple with the TrustedContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExtended) GetTrustedContactOk() (*TrustedContact, bool) {
	if o == nil || o.TrustedContact == nil {
		return nil, false
	}
	return o.TrustedContact, true
}

// HasTrustedContact returns a boolean if a field has been set.
func (o *AccountExtended) HasTrustedContact() bool {
	if o != nil && o.TrustedContact != nil {
		return true
	}

	return false
}

// SetTrustedContact gets a reference to the given TrustedContact and assigns it to the TrustedContact field.
func (o *AccountExtended) SetTrustedContact(v TrustedContact) {
	o.TrustedContact = &v
}

func (o AccountExtended) MarshalJSON() ([]byte, error) {
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
	if o.Contact != nil {
		toSerialize["contact"] = o.Contact
	}
	if o.Identity != nil {
		toSerialize["identity"] = o.Identity
	}
	if o.Disclosures != nil {
		toSerialize["disclosures"] = o.Disclosures
	}
	if o.Agreements != nil {
		toSerialize["agreements"] = o.Agreements
	}
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}
	if o.TrustedContact != nil {
		toSerialize["trusted_contact"] = o.TrustedContact
	}
	return json.Marshal(toSerialize)
}

type NullableAccountExtended struct {
	value *AccountExtended
	isSet bool
}

func (v NullableAccountExtended) Get() *AccountExtended {
	return v.value
}

func (v *NullableAccountExtended) Set(val *AccountExtended) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountExtended) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountExtended) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountExtended(val *AccountExtended) *NullableAccountExtended {
	return &NullableAccountExtended{value: val, isSet: true}
}

func (v NullableAccountExtended) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountExtended) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


