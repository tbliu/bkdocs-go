# AccountUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**Identity** | Pointer to [**Identity**](Identity.md) |  | [optional] 
**Disclosures** | Pointer to [**Disclosures**](Disclosures.md) |  | [optional] 
**TrustedContact** | Pointer to [**TrustedContact**](TrustedContact.md) |  | [optional] 

## Methods

### NewAccountUpdate

`func NewAccountUpdate() *AccountUpdate`

NewAccountUpdate instantiates a new AccountUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUpdateWithDefaults

`func NewAccountUpdateWithDefaults() *AccountUpdate`

NewAccountUpdateWithDefaults instantiates a new AccountUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *AccountUpdate) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *AccountUpdate) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *AccountUpdate) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *AccountUpdate) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetIdentity

`func (o *AccountUpdate) GetIdentity() Identity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *AccountUpdate) GetIdentityOk() (*Identity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *AccountUpdate) SetIdentity(v Identity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *AccountUpdate) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetDisclosures

`func (o *AccountUpdate) GetDisclosures() Disclosures`

GetDisclosures returns the Disclosures field if non-nil, zero value otherwise.

### GetDisclosuresOk

`func (o *AccountUpdate) GetDisclosuresOk() (*Disclosures, bool)`

GetDisclosuresOk returns a tuple with the Disclosures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosures

`func (o *AccountUpdate) SetDisclosures(v Disclosures)`

SetDisclosures sets Disclosures field to given value.

### HasDisclosures

`func (o *AccountUpdate) HasDisclosures() bool`

HasDisclosures returns a boolean if a field has been set.

### GetTrustedContact

`func (o *AccountUpdate) GetTrustedContact() TrustedContact`

GetTrustedContact returns the TrustedContact field if non-nil, zero value otherwise.

### GetTrustedContactOk

`func (o *AccountUpdate) GetTrustedContactOk() (*TrustedContact, bool)`

GetTrustedContactOk returns a tuple with the TrustedContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedContact

`func (o *AccountUpdate) SetTrustedContact(v TrustedContact)`

SetTrustedContact sets TrustedContact field to given value.

### HasTrustedContact

`func (o *AccountUpdate) HasTrustedContact() bool`

HasTrustedContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


