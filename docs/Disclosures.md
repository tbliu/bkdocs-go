# Disclosures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmploymentStatus** | Pointer to **string** |  | [optional] 
**EmployerName** | Pointer to **string** |  | [optional] 
**EmployerAddress** | Pointer to **string** |  | [optional] 
**EmploymentPosition** | Pointer to **string** |  | [optional] 
**IsControlPerson** | **bool** |  | 
**IsAffiliatedExchangeOrFinra** | **bool** |  | 
**IsPoliticallyExposed** | **bool** |  | 
**ImmediateFamilyExposed** | **bool** |  | 

## Methods

### NewDisclosures

`func NewDisclosures(isControlPerson bool, isAffiliatedExchangeOrFinra bool, isPoliticallyExposed bool, immediateFamilyExposed bool, ) *Disclosures`

NewDisclosures instantiates a new Disclosures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisclosuresWithDefaults

`func NewDisclosuresWithDefaults() *Disclosures`

NewDisclosuresWithDefaults instantiates a new Disclosures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmploymentStatus

`func (o *Disclosures) GetEmploymentStatus() string`

GetEmploymentStatus returns the EmploymentStatus field if non-nil, zero value otherwise.

### GetEmploymentStatusOk

`func (o *Disclosures) GetEmploymentStatusOk() (*string, bool)`

GetEmploymentStatusOk returns a tuple with the EmploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentStatus

`func (o *Disclosures) SetEmploymentStatus(v string)`

SetEmploymentStatus sets EmploymentStatus field to given value.

### HasEmploymentStatus

`func (o *Disclosures) HasEmploymentStatus() bool`

HasEmploymentStatus returns a boolean if a field has been set.

### GetEmployerName

`func (o *Disclosures) GetEmployerName() string`

GetEmployerName returns the EmployerName field if non-nil, zero value otherwise.

### GetEmployerNameOk

`func (o *Disclosures) GetEmployerNameOk() (*string, bool)`

GetEmployerNameOk returns a tuple with the EmployerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerName

`func (o *Disclosures) SetEmployerName(v string)`

SetEmployerName sets EmployerName field to given value.

### HasEmployerName

`func (o *Disclosures) HasEmployerName() bool`

HasEmployerName returns a boolean if a field has been set.

### GetEmployerAddress

`func (o *Disclosures) GetEmployerAddress() string`

GetEmployerAddress returns the EmployerAddress field if non-nil, zero value otherwise.

### GetEmployerAddressOk

`func (o *Disclosures) GetEmployerAddressOk() (*string, bool)`

GetEmployerAddressOk returns a tuple with the EmployerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerAddress

`func (o *Disclosures) SetEmployerAddress(v string)`

SetEmployerAddress sets EmployerAddress field to given value.

### HasEmployerAddress

`func (o *Disclosures) HasEmployerAddress() bool`

HasEmployerAddress returns a boolean if a field has been set.

### GetEmploymentPosition

`func (o *Disclosures) GetEmploymentPosition() string`

GetEmploymentPosition returns the EmploymentPosition field if non-nil, zero value otherwise.

### GetEmploymentPositionOk

`func (o *Disclosures) GetEmploymentPositionOk() (*string, bool)`

GetEmploymentPositionOk returns a tuple with the EmploymentPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentPosition

`func (o *Disclosures) SetEmploymentPosition(v string)`

SetEmploymentPosition sets EmploymentPosition field to given value.

### HasEmploymentPosition

`func (o *Disclosures) HasEmploymentPosition() bool`

HasEmploymentPosition returns a boolean if a field has been set.

### GetIsControlPerson

`func (o *Disclosures) GetIsControlPerson() bool`

GetIsControlPerson returns the IsControlPerson field if non-nil, zero value otherwise.

### GetIsControlPersonOk

`func (o *Disclosures) GetIsControlPersonOk() (*bool, bool)`

GetIsControlPersonOk returns a tuple with the IsControlPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsControlPerson

`func (o *Disclosures) SetIsControlPerson(v bool)`

SetIsControlPerson sets IsControlPerson field to given value.


### GetIsAffiliatedExchangeOrFinra

`func (o *Disclosures) GetIsAffiliatedExchangeOrFinra() bool`

GetIsAffiliatedExchangeOrFinra returns the IsAffiliatedExchangeOrFinra field if non-nil, zero value otherwise.

### GetIsAffiliatedExchangeOrFinraOk

`func (o *Disclosures) GetIsAffiliatedExchangeOrFinraOk() (*bool, bool)`

GetIsAffiliatedExchangeOrFinraOk returns a tuple with the IsAffiliatedExchangeOrFinra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAffiliatedExchangeOrFinra

`func (o *Disclosures) SetIsAffiliatedExchangeOrFinra(v bool)`

SetIsAffiliatedExchangeOrFinra sets IsAffiliatedExchangeOrFinra field to given value.


### GetIsPoliticallyExposed

`func (o *Disclosures) GetIsPoliticallyExposed() bool`

GetIsPoliticallyExposed returns the IsPoliticallyExposed field if non-nil, zero value otherwise.

### GetIsPoliticallyExposedOk

`func (o *Disclosures) GetIsPoliticallyExposedOk() (*bool, bool)`

GetIsPoliticallyExposedOk returns a tuple with the IsPoliticallyExposed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPoliticallyExposed

`func (o *Disclosures) SetIsPoliticallyExposed(v bool)`

SetIsPoliticallyExposed sets IsPoliticallyExposed field to given value.


### GetImmediateFamilyExposed

`func (o *Disclosures) GetImmediateFamilyExposed() bool`

GetImmediateFamilyExposed returns the ImmediateFamilyExposed field if non-nil, zero value otherwise.

### GetImmediateFamilyExposedOk

`func (o *Disclosures) GetImmediateFamilyExposedOk() (*bool, bool)`

GetImmediateFamilyExposedOk returns a tuple with the ImmediateFamilyExposed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateFamilyExposed

`func (o *Disclosures) SetImmediateFamilyExposed(v bool)`

SetImmediateFamilyExposed sets ImmediateFamilyExposed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


