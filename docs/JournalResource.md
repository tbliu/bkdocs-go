# JournalResource

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
**Symbol** | **string** | Only valid for JNLS journals. Null for JNLC. | 
**Qty** | **float64** | Only valid for JNLS journals. Null for JNLC. | 
**Price** | **float64** | Only valid for JNLS journals. Null for JNLC. | 

## Methods

### NewJournalResource

`func NewJournalResource(netAmount float64, symbol string, qty float64, price float64, ) *JournalResource`

NewJournalResource instantiates a new JournalResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalResourceWithDefaults

`func NewJournalResourceWithDefaults() *JournalResource`

NewJournalResourceWithDefaults instantiates a new JournalResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *JournalResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JournalResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JournalResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JournalResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNetAmount

`func (o *JournalResource) GetNetAmount() float64`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *JournalResource) GetNetAmountOk() (*float64, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *JournalResource) SetNetAmount(v float64)`

SetNetAmount sets NetAmount field to given value.


### GetTransmitterName

`func (o *JournalResource) GetTransmitterName() string`

GetTransmitterName returns the TransmitterName field if non-nil, zero value otherwise.

### GetTransmitterNameOk

`func (o *JournalResource) GetTransmitterNameOk() (*string, bool)`

GetTransmitterNameOk returns a tuple with the TransmitterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterName

`func (o *JournalResource) SetTransmitterName(v string)`

SetTransmitterName sets TransmitterName field to given value.

### HasTransmitterName

`func (o *JournalResource) HasTransmitterName() bool`

HasTransmitterName returns a boolean if a field has been set.

### GetTransmitterAccountNumber

`func (o *JournalResource) GetTransmitterAccountNumber() string`

GetTransmitterAccountNumber returns the TransmitterAccountNumber field if non-nil, zero value otherwise.

### GetTransmitterAccountNumberOk

`func (o *JournalResource) GetTransmitterAccountNumberOk() (*string, bool)`

GetTransmitterAccountNumberOk returns a tuple with the TransmitterAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterAccountNumber

`func (o *JournalResource) SetTransmitterAccountNumber(v string)`

SetTransmitterAccountNumber sets TransmitterAccountNumber field to given value.

### HasTransmitterAccountNumber

`func (o *JournalResource) HasTransmitterAccountNumber() bool`

HasTransmitterAccountNumber returns a boolean if a field has been set.

### GetTransmitterAddress

`func (o *JournalResource) GetTransmitterAddress() string`

GetTransmitterAddress returns the TransmitterAddress field if non-nil, zero value otherwise.

### GetTransmitterAddressOk

`func (o *JournalResource) GetTransmitterAddressOk() (*string, bool)`

GetTransmitterAddressOk returns a tuple with the TransmitterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterAddress

`func (o *JournalResource) SetTransmitterAddress(v string)`

SetTransmitterAddress sets TransmitterAddress field to given value.

### HasTransmitterAddress

`func (o *JournalResource) HasTransmitterAddress() bool`

HasTransmitterAddress returns a boolean if a field has been set.

### GetTransmitterFinancialInstitution

`func (o *JournalResource) GetTransmitterFinancialInstitution() string`

GetTransmitterFinancialInstitution returns the TransmitterFinancialInstitution field if non-nil, zero value otherwise.

### GetTransmitterFinancialInstitutionOk

`func (o *JournalResource) GetTransmitterFinancialInstitutionOk() (*string, bool)`

GetTransmitterFinancialInstitutionOk returns a tuple with the TransmitterFinancialInstitution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterFinancialInstitution

`func (o *JournalResource) SetTransmitterFinancialInstitution(v string)`

SetTransmitterFinancialInstitution sets TransmitterFinancialInstitution field to given value.

### HasTransmitterFinancialInstitution

`func (o *JournalResource) HasTransmitterFinancialInstitution() bool`

HasTransmitterFinancialInstitution returns a boolean if a field has been set.

### GetTransmitterTimestamp

`func (o *JournalResource) GetTransmitterTimestamp() time.Time`

GetTransmitterTimestamp returns the TransmitterTimestamp field if non-nil, zero value otherwise.

### GetTransmitterTimestampOk

`func (o *JournalResource) GetTransmitterTimestampOk() (*time.Time, bool)`

GetTransmitterTimestampOk returns a tuple with the TransmitterTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterTimestamp

`func (o *JournalResource) SetTransmitterTimestamp(v time.Time)`

SetTransmitterTimestamp sets TransmitterTimestamp field to given value.

### HasTransmitterTimestamp

`func (o *JournalResource) HasTransmitterTimestamp() bool`

HasTransmitterTimestamp returns a boolean if a field has been set.

### GetSymbol

`func (o *JournalResource) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *JournalResource) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *JournalResource) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetQty

`func (o *JournalResource) GetQty() float64`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *JournalResource) GetQtyOk() (*float64, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *JournalResource) SetQty(v float64)`

SetQty sets Qty field to given value.


### GetPrice

`func (o *JournalResource) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *JournalResource) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *JournalResource) SetPrice(v float64)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


