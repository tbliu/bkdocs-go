# AccountCreationObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**Identity** | Pointer to [**Identity**](Identity.md) |  | [optional] 
**Disclosures** | Pointer to [**Disclosures**](Disclosures.md) |  | [optional] 
**Agreements** | Pointer to [**[]Agreement**](Agreement.md) | The client has to present Alpaca Account Agreement and Margin Agreement to the end user, and have them read full sentences.  | [optional] 
**Documents** | Pointer to [**[]DocumentUpload**](DocumentUpload.md) |  | [optional] 
**TrustedContact** | Pointer to [**TrustedContact**](TrustedContact.md) |  | [optional] 

## Methods

### NewAccountCreationObject

`func NewAccountCreationObject() *AccountCreationObject`

NewAccountCreationObject instantiates a new AccountCreationObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreationObjectWithDefaults

`func NewAccountCreationObjectWithDefaults() *AccountCreationObject`

NewAccountCreationObjectWithDefaults instantiates a new AccountCreationObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *AccountCreationObject) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *AccountCreationObject) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *AccountCreationObject) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *AccountCreationObject) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetIdentity

`func (o *AccountCreationObject) GetIdentity() Identity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *AccountCreationObject) GetIdentityOk() (*Identity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *AccountCreationObject) SetIdentity(v Identity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *AccountCreationObject) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetDisclosures

`func (o *AccountCreationObject) GetDisclosures() Disclosures`

GetDisclosures returns the Disclosures field if non-nil, zero value otherwise.

### GetDisclosuresOk

`func (o *AccountCreationObject) GetDisclosuresOk() (*Disclosures, bool)`

GetDisclosuresOk returns a tuple with the Disclosures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosures

`func (o *AccountCreationObject) SetDisclosures(v Disclosures)`

SetDisclosures sets Disclosures field to given value.

### HasDisclosures

`func (o *AccountCreationObject) HasDisclosures() bool`

HasDisclosures returns a boolean if a field has been set.

### GetAgreements

`func (o *AccountCreationObject) GetAgreements() []Agreement`

GetAgreements returns the Agreements field if non-nil, zero value otherwise.

### GetAgreementsOk

`func (o *AccountCreationObject) GetAgreementsOk() (*[]Agreement, bool)`

GetAgreementsOk returns a tuple with the Agreements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreements

`func (o *AccountCreationObject) SetAgreements(v []Agreement)`

SetAgreements sets Agreements field to given value.

### HasAgreements

`func (o *AccountCreationObject) HasAgreements() bool`

HasAgreements returns a boolean if a field has been set.

### GetDocuments

`func (o *AccountCreationObject) GetDocuments() []DocumentUpload`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *AccountCreationObject) GetDocumentsOk() (*[]DocumentUpload, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *AccountCreationObject) SetDocuments(v []DocumentUpload)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *AccountCreationObject) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetTrustedContact

`func (o *AccountCreationObject) GetTrustedContact() TrustedContact`

GetTrustedContact returns the TrustedContact field if non-nil, zero value otherwise.

### GetTrustedContactOk

`func (o *AccountCreationObject) GetTrustedContactOk() (*TrustedContact, bool)`

GetTrustedContactOk returns a tuple with the TrustedContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedContact

`func (o *AccountCreationObject) SetTrustedContact(v TrustedContact)`

SetTrustedContact sets TrustedContact field to given value.

### HasTrustedContact

`func (o *AccountCreationObject) HasTrustedContact() bool`

HasTrustedContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


