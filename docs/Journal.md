# Journal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | journal ID | 
**EntryType** | **string** | JNLS, for transfers of securities, or JNLC, for transfers of cash. | 
**FromAccount** | **string** | account ID the shares go from | 
**ToAccount** | **string** | account ID the shares go to | 
**SettleDate** | **string** |  | 
**Status** | Pointer to **string** |  | [optional] 
**Symbol** | **string** | Only valid for JNLS journals. Null for JNLC. | 
**Qty** | **float64** | Only valid for JNLS journals. Null for JNLC. | 
**Price** | **float64** | Only valid for JNLS journals. Null for JNLC. | 
**Description** | Pointer to **string** | ID the amount goes to. Only valid for JNLC journals. Null for JNLS. | [optional] 
**NetAmount** | **float64** | Only valid for JNLC journals. Null for JNLS. | 
**TransmitterName** | Pointer to **string** | Only valid for JNLC journals. Null for JNLS. Max 255 characters. | [optional] 
**TransmitterAccountNumber** | Pointer to **string** | Only valid for JNLC journals. Null for JNLS.max 255 characters | [optional] 
**TransmitterAddress** | Pointer to **string** | Only valid for JNLC journals. Null for JNLS.max 255 characters | [optional] 
**TransmitterFinancialInstitution** | Pointer to **string** | Only valid for JNLC journals. Null for JNLS.max 255 characters | [optional] 
**TransmitterTimestamp** | Pointer to **time.Time** | Only valid for JNLC journals. Null for JNLS. | [optional] 

## Methods

### NewJournal

`func NewJournal(id string, entryType string, fromAccount string, toAccount string, settleDate string, symbol string, qty float64, price float64, netAmount float64, ) *Journal`

NewJournal instantiates a new Journal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalWithDefaults

`func NewJournalWithDefaults() *Journal`

NewJournalWithDefaults instantiates a new Journal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Journal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Journal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Journal) SetId(v string)`

SetId sets Id field to given value.


### GetEntryType

`func (o *Journal) GetEntryType() string`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *Journal) GetEntryTypeOk() (*string, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *Journal) SetEntryType(v string)`

SetEntryType sets EntryType field to given value.


### GetFromAccount

`func (o *Journal) GetFromAccount() string`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *Journal) GetFromAccountOk() (*string, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *Journal) SetFromAccount(v string)`

SetFromAccount sets FromAccount field to given value.


### GetToAccount

`func (o *Journal) GetToAccount() string`

GetToAccount returns the ToAccount field if non-nil, zero value otherwise.

### GetToAccountOk

`func (o *Journal) GetToAccountOk() (*string, bool)`

GetToAccountOk returns a tuple with the ToAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccount

`func (o *Journal) SetToAccount(v string)`

SetToAccount sets ToAccount field to given value.


### GetSettleDate

`func (o *Journal) GetSettleDate() string`

GetSettleDate returns the SettleDate field if non-nil, zero value otherwise.

### GetSettleDateOk

`func (o *Journal) GetSettleDateOk() (*string, bool)`

GetSettleDateOk returns a tuple with the SettleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettleDate

`func (o *Journal) SetSettleDate(v string)`

SetSettleDate sets SettleDate field to given value.


### GetStatus

`func (o *Journal) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Journal) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Journal) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Journal) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSymbol

`func (o *Journal) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Journal) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Journal) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetQty

`func (o *Journal) GetQty() float64`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *Journal) GetQtyOk() (*float64, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *Journal) SetQty(v float64)`

SetQty sets Qty field to given value.


### GetPrice

`func (o *Journal) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Journal) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Journal) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetDescription

`func (o *Journal) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Journal) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Journal) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Journal) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNetAmount

`func (o *Journal) GetNetAmount() float64`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *Journal) GetNetAmountOk() (*float64, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *Journal) SetNetAmount(v float64)`

SetNetAmount sets NetAmount field to given value.


### GetTransmitterName

`func (o *Journal) GetTransmitterName() string`

GetTransmitterName returns the TransmitterName field if non-nil, zero value otherwise.

### GetTransmitterNameOk

`func (o *Journal) GetTransmitterNameOk() (*string, bool)`

GetTransmitterNameOk returns a tuple with the TransmitterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterName

`func (o *Journal) SetTransmitterName(v string)`

SetTransmitterName sets TransmitterName field to given value.

### HasTransmitterName

`func (o *Journal) HasTransmitterName() bool`

HasTransmitterName returns a boolean if a field has been set.

### GetTransmitterAccountNumber

`func (o *Journal) GetTransmitterAccountNumber() string`

GetTransmitterAccountNumber returns the TransmitterAccountNumber field if non-nil, zero value otherwise.

### GetTransmitterAccountNumberOk

`func (o *Journal) GetTransmitterAccountNumberOk() (*string, bool)`

GetTransmitterAccountNumberOk returns a tuple with the TransmitterAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterAccountNumber

`func (o *Journal) SetTransmitterAccountNumber(v string)`

SetTransmitterAccountNumber sets TransmitterAccountNumber field to given value.

### HasTransmitterAccountNumber

`func (o *Journal) HasTransmitterAccountNumber() bool`

HasTransmitterAccountNumber returns a boolean if a field has been set.

### GetTransmitterAddress

`func (o *Journal) GetTransmitterAddress() string`

GetTransmitterAddress returns the TransmitterAddress field if non-nil, zero value otherwise.

### GetTransmitterAddressOk

`func (o *Journal) GetTransmitterAddressOk() (*string, bool)`

GetTransmitterAddressOk returns a tuple with the TransmitterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterAddress

`func (o *Journal) SetTransmitterAddress(v string)`

SetTransmitterAddress sets TransmitterAddress field to given value.

### HasTransmitterAddress

`func (o *Journal) HasTransmitterAddress() bool`

HasTransmitterAddress returns a boolean if a field has been set.

### GetTransmitterFinancialInstitution

`func (o *Journal) GetTransmitterFinancialInstitution() string`

GetTransmitterFinancialInstitution returns the TransmitterFinancialInstitution field if non-nil, zero value otherwise.

### GetTransmitterFinancialInstitutionOk

`func (o *Journal) GetTransmitterFinancialInstitutionOk() (*string, bool)`

GetTransmitterFinancialInstitutionOk returns a tuple with the TransmitterFinancialInstitution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterFinancialInstitution

`func (o *Journal) SetTransmitterFinancialInstitution(v string)`

SetTransmitterFinancialInstitution sets TransmitterFinancialInstitution field to given value.

### HasTransmitterFinancialInstitution

`func (o *Journal) HasTransmitterFinancialInstitution() bool`

HasTransmitterFinancialInstitution returns a boolean if a field has been set.

### GetTransmitterTimestamp

`func (o *Journal) GetTransmitterTimestamp() time.Time`

GetTransmitterTimestamp returns the TransmitterTimestamp field if non-nil, zero value otherwise.

### GetTransmitterTimestampOk

`func (o *Journal) GetTransmitterTimestampOk() (*time.Time, bool)`

GetTransmitterTimestampOk returns a tuple with the TransmitterTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterTimestamp

`func (o *Journal) SetTransmitterTimestamp(v time.Time)`

SetTransmitterTimestamp sets TransmitterTimestamp field to given value.

### HasTransmitterTimestamp

`func (o *Journal) HasTransmitterTimestamp() bool`

HasTransmitterTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


