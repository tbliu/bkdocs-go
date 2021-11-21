# JournalAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | journal ID | 
**EntryType** | **string** | JNLS, for transfers of securities, or JNLC, for transfers of cash. | 
**FromAccount** | **string** | account ID the shares go from | 
**ToAccount** | **string** | account ID the shares go to | 
**SettleDate** | **string** |  | 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewJournalAllOf

`func NewJournalAllOf(id string, entryType string, fromAccount string, toAccount string, settleDate string, ) *JournalAllOf`

NewJournalAllOf instantiates a new JournalAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalAllOfWithDefaults

`func NewJournalAllOfWithDefaults() *JournalAllOf`

NewJournalAllOfWithDefaults instantiates a new JournalAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JournalAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JournalAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JournalAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetEntryType

`func (o *JournalAllOf) GetEntryType() string`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *JournalAllOf) GetEntryTypeOk() (*string, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *JournalAllOf) SetEntryType(v string)`

SetEntryType sets EntryType field to given value.


### GetFromAccount

`func (o *JournalAllOf) GetFromAccount() string`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *JournalAllOf) GetFromAccountOk() (*string, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *JournalAllOf) SetFromAccount(v string)`

SetFromAccount sets FromAccount field to given value.


### GetToAccount

`func (o *JournalAllOf) GetToAccount() string`

GetToAccount returns the ToAccount field if non-nil, zero value otherwise.

### GetToAccountOk

`func (o *JournalAllOf) GetToAccountOk() (*string, bool)`

GetToAccountOk returns a tuple with the ToAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccount

`func (o *JournalAllOf) SetToAccount(v string)`

SetToAccount sets ToAccount field to given value.


### GetSettleDate

`func (o *JournalAllOf) GetSettleDate() string`

GetSettleDate returns the SettleDate field if non-nil, zero value otherwise.

### GetSettleDateOk

`func (o *JournalAllOf) GetSettleDateOk() (*string, bool)`

GetSettleDateOk returns a tuple with the SettleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettleDate

`func (o *JournalAllOf) SetSettleDate(v string)`

SetSettleDate sets SettleDate field to given value.


### GetStatus

`func (o *JournalAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JournalAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JournalAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JournalAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


