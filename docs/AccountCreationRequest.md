# AccountCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**Identity** | Pointer to [**Identity**](Identity.md) |  | [optional] 
**Disclosures** | Pointer to [**Disclosures**](Disclosures.md) |  | [optional] 
**Agreements** | Pointer to [**[]Agreement**](Agreement.md) | The client has to present Alpaca Account Agreement and Margin Agreement to the end user, and have them read full sentences.  | [optional] 
**Documents** | Pointer to [**[]DocumentUploadRequest**](DocumentUploadRequest.md) |  | [optional] 
**TrustedContact** | Pointer to [**TrustedContact**](TrustedContact.md) |  | [optional] 

## Methods

### NewAccountCreationRequest

`func NewAccountCreationRequest() *AccountCreationRequest`

NewAccountCreationRequest instantiates a new AccountCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreationRequestWithDefaults

`func NewAccountCreationRequestWithDefaults() *AccountCreationRequest`

NewAccountCreationRequestWithDefaults instantiates a new AccountCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *AccountCreationRequest) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *AccountCreationRequest) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *AccountCreationRequest) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *AccountCreationRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetIdentity

`func (o *AccountCreationRequest) GetIdentity() Identity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *AccountCreationRequest) GetIdentityOk() (*Identity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *AccountCreationRequest) SetIdentity(v Identity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *AccountCreationRequest) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetDisclosures

`func (o *AccountCreationRequest) GetDisclosures() Disclosures`

GetDisclosures returns the Disclosures field if non-nil, zero value otherwise.

### GetDisclosuresOk

`func (o *AccountCreationRequest) GetDisclosuresOk() (*Disclosures, bool)`

GetDisclosuresOk returns a tuple with the Disclosures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosures

`func (o *AccountCreationRequest) SetDisclosures(v Disclosures)`

SetDisclosures sets Disclosures field to given value.

### HasDisclosures

`func (o *AccountCreationRequest) HasDisclosures() bool`

HasDisclosures returns a boolean if a field has been set.

### GetAgreements

`func (o *AccountCreationRequest) GetAgreements() []Agreement`

GetAgreements returns the Agreements field if non-nil, zero value otherwise.

### GetAgreementsOk

`func (o *AccountCreationRequest) GetAgreementsOk() (*[]Agreement, bool)`

GetAgreementsOk returns a tuple with the Agreements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreements

`func (o *AccountCreationRequest) SetAgreements(v []Agreement)`

SetAgreements sets Agreements field to given value.

### HasAgreements

`func (o *AccountCreationRequest) HasAgreements() bool`

HasAgreements returns a boolean if a field has been set.

### GetDocuments

`func (o *AccountCreationRequest) GetDocuments() []DocumentUploadRequest`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *AccountCreationRequest) GetDocumentsOk() (*[]DocumentUploadRequest, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *AccountCreationRequest) SetDocuments(v []DocumentUploadRequest)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *AccountCreationRequest) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetTrustedContact

`func (o *AccountCreationRequest) GetTrustedContact() TrustedContact`

GetTrustedContact returns the TrustedContact field if non-nil, zero value otherwise.

### GetTrustedContactOk

`func (o *AccountCreationRequest) GetTrustedContactOk() (*TrustedContact, bool)`

GetTrustedContactOk returns a tuple with the TrustedContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedContact

`func (o *AccountCreationRequest) SetTrustedContact(v TrustedContact)`

SetTrustedContact sets TrustedContact field to given value.

### HasTrustedContact

`func (o *AccountCreationRequest) HasTrustedContact() bool`

HasTrustedContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


