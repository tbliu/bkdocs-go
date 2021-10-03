# UntypedWireTransferData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float64** |  | 
**Direction** | **string** |  | 
**AdditionalInformation** | Pointer to **string** |  | [optional] 
**BankId** | **string** |  | 

## Methods

### NewUntypedWireTransferData

`func NewUntypedWireTransferData(amount float64, direction string, bankId string, ) *UntypedWireTransferData`

NewUntypedWireTransferData instantiates a new UntypedWireTransferData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUntypedWireTransferDataWithDefaults

`func NewUntypedWireTransferDataWithDefaults() *UntypedWireTransferData`

NewUntypedWireTransferDataWithDefaults instantiates a new UntypedWireTransferData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UntypedWireTransferData) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UntypedWireTransferData) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UntypedWireTransferData) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetDirection

`func (o *UntypedWireTransferData) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *UntypedWireTransferData) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *UntypedWireTransferData) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetAdditionalInformation

`func (o *UntypedWireTransferData) GetAdditionalInformation() string`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *UntypedWireTransferData) GetAdditionalInformationOk() (*string, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *UntypedWireTransferData) SetAdditionalInformation(v string)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *UntypedWireTransferData) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetBankId

`func (o *UntypedWireTransferData) GetBankId() string`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *UntypedWireTransferData) GetBankIdOk() (*string, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *UntypedWireTransferData) SetBankId(v string)`

SetBankId sets BankId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


