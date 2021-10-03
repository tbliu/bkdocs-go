# TransferData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransferType** | **string** |  | 
**Timing** | Pointer to **string** |  | [optional] 
**Amount** | **float64** |  | 
**Direction** | **string** |  | 
**RelationshipId** | **string** |  | 
**AdditionalInformation** | Pointer to **string** |  | [optional] 
**BankId** | **string** |  | 

## Methods

### NewTransferData

`func NewTransferData(transferType string, amount float64, direction string, relationshipId string, bankId string, ) *TransferData`

NewTransferData instantiates a new TransferData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferDataWithDefaults

`func NewTransferDataWithDefaults() *TransferData`

NewTransferDataWithDefaults instantiates a new TransferData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransferType

`func (o *TransferData) GetTransferType() string`

GetTransferType returns the TransferType field if non-nil, zero value otherwise.

### GetTransferTypeOk

`func (o *TransferData) GetTransferTypeOk() (*string, bool)`

GetTransferTypeOk returns a tuple with the TransferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferType

`func (o *TransferData) SetTransferType(v string)`

SetTransferType sets TransferType field to given value.


### GetTiming

`func (o *TransferData) GetTiming() string`

GetTiming returns the Timing field if non-nil, zero value otherwise.

### GetTimingOk

`func (o *TransferData) GetTimingOk() (*string, bool)`

GetTimingOk returns a tuple with the Timing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiming

`func (o *TransferData) SetTiming(v string)`

SetTiming sets Timing field to given value.

### HasTiming

`func (o *TransferData) HasTiming() bool`

HasTiming returns a boolean if a field has been set.

### GetAmount

`func (o *TransferData) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferData) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferData) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetDirection

`func (o *TransferData) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *TransferData) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *TransferData) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetRelationshipId

`func (o *TransferData) GetRelationshipId() string`

GetRelationshipId returns the RelationshipId field if non-nil, zero value otherwise.

### GetRelationshipIdOk

`func (o *TransferData) GetRelationshipIdOk() (*string, bool)`

GetRelationshipIdOk returns a tuple with the RelationshipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipId

`func (o *TransferData) SetRelationshipId(v string)`

SetRelationshipId sets RelationshipId field to given value.


### GetAdditionalInformation

`func (o *TransferData) GetAdditionalInformation() string`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *TransferData) GetAdditionalInformationOk() (*string, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *TransferData) SetAdditionalInformation(v string)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *TransferData) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetBankId

`func (o *TransferData) GetBankId() string`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *TransferData) GetBankIdOk() (*string, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *TransferData) SetBankId(v string)`

SetBankId sets BankId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


