# JNLC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | ID the amount goes to. Only valid for JNLC journals. Null for JNLS. | [optional] 
**NetAmount** | **float64** | Only valid for JNLC journals. Null for JNLS. | 
**TransmitterName** | Pointer to **string** | Only valid for JNLC journals. Null for JNLS. Max 255 characters. | [optional] 
**TransmitterAccountNumber** | Pointer to **string** | Only valid for JNLC journals. Null for JNLS.max 255 characters | [optional] 
**TransmitterAddress** | Pointer to **string** | Only valid for JNLC journals. Null for JNLS.max 255 characters | [optional] 
**TransmitterFinancialInstitution** | Pointer to **string** | Only valid for JNLC journals. Null for JNLS.max 255 characters | [optional] 
**TransmitterTimestamp** | Pointer to **time.Time** | Only valid for JNLC journals. Null for JNLS. | [optional] 

## Methods

### NewJNLC

`func NewJNLC(netAmount float64, ) *JNLC`

NewJNLC instantiates a new JNLC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJNLCWithDefaults

`func NewJNLCWithDefaults() *JNLC`

NewJNLCWithDefaults instantiates a new JNLC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *JNLC) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JNLC) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JNLC) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JNLC) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNetAmount

`func (o *JNLC) GetNetAmount() float64`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *JNLC) GetNetAmountOk() (*float64, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *JNLC) SetNetAmount(v float64)`

SetNetAmount sets NetAmount field to given value.


### GetTransmitterName

`func (o *JNLC) GetTransmitterName() string`

GetTransmitterName returns the TransmitterName field if non-nil, zero value otherwise.

### GetTransmitterNameOk

`func (o *JNLC) GetTransmitterNameOk() (*string, bool)`

GetTransmitterNameOk returns a tuple with the TransmitterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterName

`func (o *JNLC) SetTransmitterName(v string)`

SetTransmitterName sets TransmitterName field to given value.

### HasTransmitterName

`func (o *JNLC) HasTransmitterName() bool`

HasTransmitterName returns a boolean if a field has been set.

### GetTransmitterAccountNumber

`func (o *JNLC) GetTransmitterAccountNumber() string`

GetTransmitterAccountNumber returns the TransmitterAccountNumber field if non-nil, zero value otherwise.

### GetTransmitterAccountNumberOk

`func (o *JNLC) GetTransmitterAccountNumberOk() (*string, bool)`

GetTransmitterAccountNumberOk returns a tuple with the TransmitterAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterAccountNumber

`func (o *JNLC) SetTransmitterAccountNumber(v string)`

SetTransmitterAccountNumber sets TransmitterAccountNumber field to given value.

### HasTransmitterAccountNumber

`func (o *JNLC) HasTransmitterAccountNumber() bool`

HasTransmitterAccountNumber returns a boolean if a field has been set.

### GetTransmitterAddress

`func (o *JNLC) GetTransmitterAddress() string`

GetTransmitterAddress returns the TransmitterAddress field if non-nil, zero value otherwise.

### GetTransmitterAddressOk

`func (o *JNLC) GetTransmitterAddressOk() (*string, bool)`

GetTransmitterAddressOk returns a tuple with the TransmitterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterAddress

`func (o *JNLC) SetTransmitterAddress(v string)`

SetTransmitterAddress sets TransmitterAddress field to given value.

### HasTransmitterAddress

`func (o *JNLC) HasTransmitterAddress() bool`

HasTransmitterAddress returns a boolean if a field has been set.

### GetTransmitterFinancialInstitution

`func (o *JNLC) GetTransmitterFinancialInstitution() string`

GetTransmitterFinancialInstitution returns the TransmitterFinancialInstitution field if non-nil, zero value otherwise.

### GetTransmitterFinancialInstitutionOk

`func (o *JNLC) GetTransmitterFinancialInstitutionOk() (*string, bool)`

GetTransmitterFinancialInstitutionOk returns a tuple with the TransmitterFinancialInstitution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterFinancialInstitution

`func (o *JNLC) SetTransmitterFinancialInstitution(v string)`

SetTransmitterFinancialInstitution sets TransmitterFinancialInstitution field to given value.

### HasTransmitterFinancialInstitution

`func (o *JNLC) HasTransmitterFinancialInstitution() bool`

HasTransmitterFinancialInstitution returns a boolean if a field has been set.

### GetTransmitterTimestamp

`func (o *JNLC) GetTransmitterTimestamp() time.Time`

GetTransmitterTimestamp returns the TransmitterTimestamp field if non-nil, zero value otherwise.

### GetTransmitterTimestampOk

`func (o *JNLC) GetTransmitterTimestampOk() (*time.Time, bool)`

GetTransmitterTimestampOk returns a tuple with the TransmitterTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterTimestamp

`func (o *JNLC) SetTransmitterTimestamp(v time.Time)`

SetTransmitterTimestamp sets TransmitterTimestamp field to given value.

### HasTransmitterTimestamp

`func (o *JNLC) HasTransmitterTimestamp() bool`

HasTransmitterTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


