# AccountExtended

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**AccountStatus**](AccountStatus.md) |  | [optional] 
**Currency** | Pointer to **string** | Always \&quot;USD\&quot; | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastEquity** | Pointer to **float64** |  | [optional] 
**KycResults** | Pointer to [**KYCResult**](KYCResult.md) |  | [optional] 
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**Identity** | Pointer to [**Identity**](Identity.md) |  | [optional] 
**Disclosures** | Pointer to [**Disclosures**](Disclosures.md) |  | [optional] 
**Agreements** | Pointer to [**[]Agreement**](Agreement.md) | The client has to present Alpaca Account Agreement and Margin Agreement to the end user, and have them read full sentences.  | [optional] 
**Documents** | Pointer to [**[]ApplicationDocument**](ApplicationDocument.md) |  | [optional] 
**TrustedContact** | Pointer to [**TrustedContact**](TrustedContact.md) |  | [optional] 

## Methods

### NewAccountExtended

`func NewAccountExtended() *AccountExtended`

NewAccountExtended instantiates a new AccountExtended object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountExtendedWithDefaults

`func NewAccountExtendedWithDefaults() *AccountExtended`

NewAccountExtendedWithDefaults instantiates a new AccountExtended object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountExtended) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountExtended) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountExtended) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountExtended) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountNumber

`func (o *AccountExtended) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AccountExtended) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AccountExtended) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *AccountExtended) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetStatus

`func (o *AccountExtended) GetStatus() AccountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountExtended) GetStatusOk() (*AccountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountExtended) SetStatus(v AccountStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountExtended) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCurrency

`func (o *AccountExtended) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AccountExtended) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AccountExtended) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AccountExtended) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccountExtended) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountExtended) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountExtended) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccountExtended) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastEquity

`func (o *AccountExtended) GetLastEquity() float64`

GetLastEquity returns the LastEquity field if non-nil, zero value otherwise.

### GetLastEquityOk

`func (o *AccountExtended) GetLastEquityOk() (*float64, bool)`

GetLastEquityOk returns a tuple with the LastEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEquity

`func (o *AccountExtended) SetLastEquity(v float64)`

SetLastEquity sets LastEquity field to given value.

### HasLastEquity

`func (o *AccountExtended) HasLastEquity() bool`

HasLastEquity returns a boolean if a field has been set.

### GetKycResults

`func (o *AccountExtended) GetKycResults() KYCResult`

GetKycResults returns the KycResults field if non-nil, zero value otherwise.

### GetKycResultsOk

`func (o *AccountExtended) GetKycResultsOk() (*KYCResult, bool)`

GetKycResultsOk returns a tuple with the KycResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycResults

`func (o *AccountExtended) SetKycResults(v KYCResult)`

SetKycResults sets KycResults field to given value.

### HasKycResults

`func (o *AccountExtended) HasKycResults() bool`

HasKycResults returns a boolean if a field has been set.

### GetContact

`func (o *AccountExtended) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *AccountExtended) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *AccountExtended) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *AccountExtended) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetIdentity

`func (o *AccountExtended) GetIdentity() Identity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *AccountExtended) GetIdentityOk() (*Identity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *AccountExtended) SetIdentity(v Identity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *AccountExtended) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetDisclosures

`func (o *AccountExtended) GetDisclosures() Disclosures`

GetDisclosures returns the Disclosures field if non-nil, zero value otherwise.

### GetDisclosuresOk

`func (o *AccountExtended) GetDisclosuresOk() (*Disclosures, bool)`

GetDisclosuresOk returns a tuple with the Disclosures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosures

`func (o *AccountExtended) SetDisclosures(v Disclosures)`

SetDisclosures sets Disclosures field to given value.

### HasDisclosures

`func (o *AccountExtended) HasDisclosures() bool`

HasDisclosures returns a boolean if a field has been set.

### GetAgreements

`func (o *AccountExtended) GetAgreements() []Agreement`

GetAgreements returns the Agreements field if non-nil, zero value otherwise.

### GetAgreementsOk

`func (o *AccountExtended) GetAgreementsOk() (*[]Agreement, bool)`

GetAgreementsOk returns a tuple with the Agreements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreements

`func (o *AccountExtended) SetAgreements(v []Agreement)`

SetAgreements sets Agreements field to given value.

### HasAgreements

`func (o *AccountExtended) HasAgreements() bool`

HasAgreements returns a boolean if a field has been set.

### GetDocuments

`func (o *AccountExtended) GetDocuments() []ApplicationDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *AccountExtended) GetDocumentsOk() (*[]ApplicationDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *AccountExtended) SetDocuments(v []ApplicationDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *AccountExtended) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetTrustedContact

`func (o *AccountExtended) GetTrustedContact() TrustedContact`

GetTrustedContact returns the TrustedContact field if non-nil, zero value otherwise.

### GetTrustedContactOk

`func (o *AccountExtended) GetTrustedContactOk() (*TrustedContact, bool)`

GetTrustedContactOk returns a tuple with the TrustedContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedContact

`func (o *AccountExtended) SetTrustedContact(v TrustedContact)`

SetTrustedContact sets TrustedContact field to given value.

### HasTrustedContact

`func (o *AccountExtended) HasTrustedContact() bool`

HasTrustedContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


