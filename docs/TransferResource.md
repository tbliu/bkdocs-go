# TransferResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Type** | **string** |  | 
**Status** | **string** |  | 
**AccountId** | **string** |  | 
**Reason** | Pointer to **NullableString** |  | [optional] 
**ExpiresAt** | **time.Time** |  | 
**Amount** | **float64** |  | 
**Direction** | **string** |  | 
**RelationshipId** | **string** |  | 
**AdditionalInformation** | Pointer to **string** |  | [optional] 
**BankId** | **string** |  | 

## Methods

### NewTransferResource

`func NewTransferResource(id string, createdAt time.Time, updatedAt time.Time, type_ string, status string, accountId string, expiresAt time.Time, amount float64, direction string, relationshipId string, bankId string, ) *TransferResource`

NewTransferResource instantiates a new TransferResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferResourceWithDefaults

`func NewTransferResourceWithDefaults() *TransferResource`

NewTransferResourceWithDefaults instantiates a new TransferResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransferResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransferResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransferResource) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *TransferResource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransferResource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransferResource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TransferResource) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TransferResource) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TransferResource) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetType

`func (o *TransferResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransferResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransferResource) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *TransferResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransferResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransferResource) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAccountId

`func (o *TransferResource) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransferResource) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransferResource) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetReason

`func (o *TransferResource) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TransferResource) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TransferResource) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TransferResource) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *TransferResource) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *TransferResource) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetExpiresAt

`func (o *TransferResource) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TransferResource) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TransferResource) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetAmount

`func (o *TransferResource) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferResource) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferResource) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetDirection

`func (o *TransferResource) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *TransferResource) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *TransferResource) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetRelationshipId

`func (o *TransferResource) GetRelationshipId() string`

GetRelationshipId returns the RelationshipId field if non-nil, zero value otherwise.

### GetRelationshipIdOk

`func (o *TransferResource) GetRelationshipIdOk() (*string, bool)`

GetRelationshipIdOk returns a tuple with the RelationshipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipId

`func (o *TransferResource) SetRelationshipId(v string)`

SetRelationshipId sets RelationshipId field to given value.


### GetAdditionalInformation

`func (o *TransferResource) GetAdditionalInformation() string`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *TransferResource) GetAdditionalInformationOk() (*string, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *TransferResource) SetAdditionalInformation(v string)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *TransferResource) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetBankId

`func (o *TransferResource) GetBankId() string`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *TransferResource) GetBankIdOk() (*string, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *TransferResource) SetBankId(v string)`

SetBankId sets BankId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


