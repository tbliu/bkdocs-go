# JournalJNLC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | journal ID | 
**EntryType** | **string** | JNLC (constant) | 
**FromAccount** | **string** | account ID the amount goes from | 
**ToAccount** | **string** |  | 
**Description** | Pointer to **string** | ID the amount goes to | [optional] 
**SettleDate** | **NullableString** |  | 
**Status** | Pointer to **string** |  | [optional] 
**NetAmount** | **float64** |  | 
**TransmitterName** | Pointer to **string** | max 255 characters | [optional] 
**TransmitterAccountNumber** | Pointer to **string** | max 255 characters | [optional] 
**TransmitterAddress** | Pointer to **string** | max 255 characters | [optional] 
**TransmitterFinancialInstitution** | Pointer to **string** | max 255 characters | [optional] 
**TransmitterTimestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewJournalJNLC

`func NewJournalJNLC(id string, entryType string, fromAccount string, toAccount string, settleDate NullableString, netAmount float64, ) *JournalJNLC`

NewJournalJNLC instantiates a new JournalJNLC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalJNLCWithDefaults

`func NewJournalJNLCWithDefaults() *JournalJNLC`

NewJournalJNLCWithDefaults instantiates a new JournalJNLC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JournalJNLC) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JournalJNLC) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JournalJNLC) SetId(v string)`

SetId sets Id field to given value.


### GetEntryType

`func (o *JournalJNLC) GetEntryType() string`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *JournalJNLC) GetEntryTypeOk() (*string, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *JournalJNLC) SetEntryType(v string)`

SetEntryType sets EntryType field to given value.


### GetFromAccount

`func (o *JournalJNLC) GetFromAccount() string`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *JournalJNLC) GetFromAccountOk() (*string, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *JournalJNLC) SetFromAccount(v string)`

SetFromAccount sets FromAccount field to given value.


### GetToAccount

`func (o *JournalJNLC) GetToAccount() string`

GetToAccount returns the ToAccount field if non-nil, zero value otherwise.

### GetToAccountOk

`func (o *JournalJNLC) GetToAccountOk() (*string, bool)`

GetToAccountOk returns a tuple with the ToAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccount

`func (o *JournalJNLC) SetToAccount(v string)`

SetToAccount sets ToAccount field to given value.


### GetDescription

`func (o *JournalJNLC) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JournalJNLC) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JournalJNLC) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JournalJNLC) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSettleDate

`func (o *JournalJNLC) GetSettleDate() string`

GetSettleDate returns the SettleDate field if non-nil, zero value otherwise.

### GetSettleDateOk

`func (o *JournalJNLC) GetSettleDateOk() (*string, bool)`

GetSettleDateOk returns a tuple with the SettleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettleDate

`func (o *JournalJNLC) SetSettleDate(v string)`

SetSettleDate sets SettleDate field to given value.


### SetSettleDateNil

`func (o *JournalJNLC) SetSettleDateNil(b bool)`

 SetSettleDateNil sets the value for SettleDate to be an explicit nil

### UnsetSettleDate
`func (o *JournalJNLC) UnsetSettleDate()`

UnsetSettleDate ensures that no value is present for SettleDate, not even an explicit nil
### GetStatus

`func (o *JournalJNLC) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JournalJNLC) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JournalJNLC) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JournalJNLC) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNetAmount

`func (o *JournalJNLC) GetNetAmount() float64`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *JournalJNLC) GetNetAmountOk() (*float64, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *JournalJNLC) SetNetAmount(v float64)`

SetNetAmount sets NetAmount field to given value.


### GetTransmitterName

`func (o *JournalJNLC) GetTransmitterName() string`

GetTransmitterName returns the TransmitterName field if non-nil, zero value otherwise.

### GetTransmitterNameOk

`func (o *JournalJNLC) GetTransmitterNameOk() (*string, bool)`

GetTransmitterNameOk returns a tuple with the TransmitterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterName

`func (o *JournalJNLC) SetTransmitterName(v string)`

SetTransmitterName sets TransmitterName field to given value.

### HasTransmitterName

`func (o *JournalJNLC) HasTransmitterName() bool`

HasTransmitterName returns a boolean if a field has been set.

### GetTransmitterAccountNumber

`func (o *JournalJNLC) GetTransmitterAccountNumber() string`

GetTransmitterAccountNumber returns the TransmitterAccountNumber field if non-nil, zero value otherwise.

### GetTransmitterAccountNumberOk

`func (o *JournalJNLC) GetTransmitterAccountNumberOk() (*string, bool)`

GetTransmitterAccountNumberOk returns a tuple with the TransmitterAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterAccountNumber

`func (o *JournalJNLC) SetTransmitterAccountNumber(v string)`

SetTransmitterAccountNumber sets TransmitterAccountNumber field to given value.

### HasTransmitterAccountNumber

`func (o *JournalJNLC) HasTransmitterAccountNumber() bool`

HasTransmitterAccountNumber returns a boolean if a field has been set.

### GetTransmitterAddress

`func (o *JournalJNLC) GetTransmitterAddress() string`

GetTransmitterAddress returns the TransmitterAddress field if non-nil, zero value otherwise.

### GetTransmitterAddressOk

`func (o *JournalJNLC) GetTransmitterAddressOk() (*string, bool)`

GetTransmitterAddressOk returns a tuple with the TransmitterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterAddress

`func (o *JournalJNLC) SetTransmitterAddress(v string)`

SetTransmitterAddress sets TransmitterAddress field to given value.

### HasTransmitterAddress

`func (o *JournalJNLC) HasTransmitterAddress() bool`

HasTransmitterAddress returns a boolean if a field has been set.

### GetTransmitterFinancialInstitution

`func (o *JournalJNLC) GetTransmitterFinancialInstitution() string`

GetTransmitterFinancialInstitution returns the TransmitterFinancialInstitution field if non-nil, zero value otherwise.

### GetTransmitterFinancialInstitutionOk

`func (o *JournalJNLC) GetTransmitterFinancialInstitutionOk() (*string, bool)`

GetTransmitterFinancialInstitutionOk returns a tuple with the TransmitterFinancialInstitution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterFinancialInstitution

`func (o *JournalJNLC) SetTransmitterFinancialInstitution(v string)`

SetTransmitterFinancialInstitution sets TransmitterFinancialInstitution field to given value.

### HasTransmitterFinancialInstitution

`func (o *JournalJNLC) HasTransmitterFinancialInstitution() bool`

HasTransmitterFinancialInstitution returns a boolean if a field has been set.

### GetTransmitterTimestamp

`func (o *JournalJNLC) GetTransmitterTimestamp() time.Time`

GetTransmitterTimestamp returns the TransmitterTimestamp field if non-nil, zero value otherwise.

### GetTransmitterTimestampOk

`func (o *JournalJNLC) GetTransmitterTimestampOk() (*time.Time, bool)`

GetTransmitterTimestampOk returns a tuple with the TransmitterTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterTimestamp

`func (o *JournalJNLC) SetTransmitterTimestamp(v time.Time)`

SetTransmitterTimestamp sets TransmitterTimestamp field to given value.

### HasTransmitterTimestamp

`func (o *JournalJNLC) HasTransmitterTimestamp() bool`

HasTransmitterTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


