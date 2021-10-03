# ACHRelationshipResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**AccountOwnerName** | **string** |  | 
**BankAccountType** | **string** |  | 
**BankAccountNumber** | **string** |  | 
**BankRoutingNumber** | **string** |  | 
**Nickname** | Pointer to **string** |  | [optional] 
**AccountId** | **string** |  | 
**Status** | **string** |  | 

## Methods

### NewACHRelationshipResource

`func NewACHRelationshipResource(id string, createdAt time.Time, updatedAt time.Time, accountOwnerName string, bankAccountType string, bankAccountNumber string, bankRoutingNumber string, accountId string, status string, ) *ACHRelationshipResource`

NewACHRelationshipResource instantiates a new ACHRelationshipResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACHRelationshipResourceWithDefaults

`func NewACHRelationshipResourceWithDefaults() *ACHRelationshipResource`

NewACHRelationshipResourceWithDefaults instantiates a new ACHRelationshipResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ACHRelationshipResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ACHRelationshipResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ACHRelationshipResource) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ACHRelationshipResource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ACHRelationshipResource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ACHRelationshipResource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ACHRelationshipResource) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ACHRelationshipResource) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ACHRelationshipResource) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAccountOwnerName

`func (o *ACHRelationshipResource) GetAccountOwnerName() string`

GetAccountOwnerName returns the AccountOwnerName field if non-nil, zero value otherwise.

### GetAccountOwnerNameOk

`func (o *ACHRelationshipResource) GetAccountOwnerNameOk() (*string, bool)`

GetAccountOwnerNameOk returns a tuple with the AccountOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOwnerName

`func (o *ACHRelationshipResource) SetAccountOwnerName(v string)`

SetAccountOwnerName sets AccountOwnerName field to given value.


### GetBankAccountType

`func (o *ACHRelationshipResource) GetBankAccountType() string`

GetBankAccountType returns the BankAccountType field if non-nil, zero value otherwise.

### GetBankAccountTypeOk

`func (o *ACHRelationshipResource) GetBankAccountTypeOk() (*string, bool)`

GetBankAccountTypeOk returns a tuple with the BankAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountType

`func (o *ACHRelationshipResource) SetBankAccountType(v string)`

SetBankAccountType sets BankAccountType field to given value.


### GetBankAccountNumber

`func (o *ACHRelationshipResource) GetBankAccountNumber() string`

GetBankAccountNumber returns the BankAccountNumber field if non-nil, zero value otherwise.

### GetBankAccountNumberOk

`func (o *ACHRelationshipResource) GetBankAccountNumberOk() (*string, bool)`

GetBankAccountNumberOk returns a tuple with the BankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountNumber

`func (o *ACHRelationshipResource) SetBankAccountNumber(v string)`

SetBankAccountNumber sets BankAccountNumber field to given value.


### GetBankRoutingNumber

`func (o *ACHRelationshipResource) GetBankRoutingNumber() string`

GetBankRoutingNumber returns the BankRoutingNumber field if non-nil, zero value otherwise.

### GetBankRoutingNumberOk

`func (o *ACHRelationshipResource) GetBankRoutingNumberOk() (*string, bool)`

GetBankRoutingNumberOk returns a tuple with the BankRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankRoutingNumber

`func (o *ACHRelationshipResource) SetBankRoutingNumber(v string)`

SetBankRoutingNumber sets BankRoutingNumber field to given value.


### GetNickname

`func (o *ACHRelationshipResource) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *ACHRelationshipResource) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *ACHRelationshipResource) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *ACHRelationshipResource) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetAccountId

`func (o *ACHRelationshipResource) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ACHRelationshipResource) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ACHRelationshipResource) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetStatus

`func (o *ACHRelationshipResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ACHRelationshipResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ACHRelationshipResource) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


