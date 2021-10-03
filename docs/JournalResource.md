# JournalResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | journal ID | 
**EntryType** | **string** | JNLS (constant) | 
**FromAccount** | **string** | account ID the shares go from | 
**ToAccount** | **string** | account ID the shares go to | 
**Description** | Pointer to **string** | ID the amount goes to | [optional] 
**SettleDate** | **NullableString** |  | 
**Status** | Pointer to **string** |  | [optional] 
**NetAmount** | **float64** |  | 
**TransmitterName** | Pointer to **string** | max 255 characters | [optional] 
**TransmitterAccountNumber** | Pointer to **string** | max 255 characters | [optional] 
**TransmitterAddress** | Pointer to **string** | max 255 characters | [optional] 
**TransmitterFinancialInstitution** | Pointer to **string** | max 255 characters | [optional] 
**TransmitterTimestamp** | Pointer to **time.Time** |  | [optional] 
**Symbol** | **string** |  | 
**Qty** | **float64** |  | 
**Price** | **float64** |  | 

## Methods

### NewJournalResource

`func NewJournalResource(id string, entryType string, fromAccount string, toAccount string, settleDate NullableString, netAmount float64, symbol string, qty float64, price float64, ) *JournalResource`

NewJournalResource instantiates a new JournalResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalResourceWithDefaults

`func NewJournalResourceWithDefaults() *JournalResource`

NewJournalResourceWithDefaults instantiates a new JournalResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JournalResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JournalResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JournalResource) SetId(v string)`

SetId sets Id field to given value.


### GetEntryType

`func (o *JournalResource) GetEntryType() string`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *JournalResource) GetEntryTypeOk() (*string, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *JournalResource) SetEntryType(v string)`

SetEntryType sets EntryType field to given value.


### GetFromAccount

`func (o *JournalResource) GetFromAccount() string`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *JournalResource) GetFromAccountOk() (*string, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *JournalResource) SetFromAccount(v string)`

SetFromAccount sets FromAccount field to given value.


### GetToAccount

`func (o *JournalResource) GetToAccount() string`

GetToAccount returns the ToAccount field if non-nil, zero value otherwise.

### GetToAccountOk

`func (o *JournalResource) GetToAccountOk() (*string, bool)`

GetToAccountOk returns a tuple with the ToAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccount

`func (o *JournalResource) SetToAccount(v string)`

SetToAccount sets ToAccount field to given value.


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

### GetSettleDate

`func (o *JournalResource) GetSettleDate() string`

GetSettleDate returns the SettleDate field if non-nil, zero value otherwise.

### GetSettleDateOk

`func (o *JournalResource) GetSettleDateOk() (*string, bool)`

GetSettleDateOk returns a tuple with the SettleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettleDate

`func (o *JournalResource) SetSettleDate(v string)`

SetSettleDate sets SettleDate field to given value.


### SetSettleDateNil

`func (o *JournalResource) SetSettleDateNil(b bool)`

 SetSettleDateNil sets the value for SettleDate to be an explicit nil

### UnsetSettleDate
`func (o *JournalResource) UnsetSettleDate()`

UnsetSettleDate ensures that no value is present for SettleDate, not even an explicit nil
### GetStatus

`func (o *JournalResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JournalResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JournalResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JournalResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

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


