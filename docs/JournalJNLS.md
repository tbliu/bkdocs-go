# JournalJNLS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | journal ID | 
**EntryType** | **string** | JNLS (constant) | 
**FromAccount** | **string** | account ID the shares go from | 
**ToAccount** | **string** | account ID the shares go to | 
**SettleDate** | **NullableString** |  | 
**Status** | Pointer to **string** |  | [optional] 
**Symbol** | **string** |  | 
**Qty** | **float64** |  | 
**Price** | **float64** |  | 

## Methods

### NewJournalJNLS

`func NewJournalJNLS(id string, entryType string, fromAccount string, toAccount string, settleDate NullableString, symbol string, qty float64, price float64, ) *JournalJNLS`

NewJournalJNLS instantiates a new JournalJNLS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalJNLSWithDefaults

`func NewJournalJNLSWithDefaults() *JournalJNLS`

NewJournalJNLSWithDefaults instantiates a new JournalJNLS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JournalJNLS) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JournalJNLS) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JournalJNLS) SetId(v string)`

SetId sets Id field to given value.


### GetEntryType

`func (o *JournalJNLS) GetEntryType() string`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *JournalJNLS) GetEntryTypeOk() (*string, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *JournalJNLS) SetEntryType(v string)`

SetEntryType sets EntryType field to given value.


### GetFromAccount

`func (o *JournalJNLS) GetFromAccount() string`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *JournalJNLS) GetFromAccountOk() (*string, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *JournalJNLS) SetFromAccount(v string)`

SetFromAccount sets FromAccount field to given value.


### GetToAccount

`func (o *JournalJNLS) GetToAccount() string`

GetToAccount returns the ToAccount field if non-nil, zero value otherwise.

### GetToAccountOk

`func (o *JournalJNLS) GetToAccountOk() (*string, bool)`

GetToAccountOk returns a tuple with the ToAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccount

`func (o *JournalJNLS) SetToAccount(v string)`

SetToAccount sets ToAccount field to given value.


### GetSettleDate

`func (o *JournalJNLS) GetSettleDate() string`

GetSettleDate returns the SettleDate field if non-nil, zero value otherwise.

### GetSettleDateOk

`func (o *JournalJNLS) GetSettleDateOk() (*string, bool)`

GetSettleDateOk returns a tuple with the SettleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettleDate

`func (o *JournalJNLS) SetSettleDate(v string)`

SetSettleDate sets SettleDate field to given value.


### SetSettleDateNil

`func (o *JournalJNLS) SetSettleDateNil(b bool)`

 SetSettleDateNil sets the value for SettleDate to be an explicit nil

### UnsetSettleDate
`func (o *JournalJNLS) UnsetSettleDate()`

UnsetSettleDate ensures that no value is present for SettleDate, not even an explicit nil
### GetStatus

`func (o *JournalJNLS) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JournalJNLS) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JournalJNLS) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JournalJNLS) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSymbol

`func (o *JournalJNLS) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *JournalJNLS) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *JournalJNLS) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetQty

`func (o *JournalJNLS) GetQty() float64`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *JournalJNLS) GetQtyOk() (*float64, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *JournalJNLS) SetQty(v float64)`

SetQty sets Qty field to given value.


### GetPrice

`func (o *JournalJNLS) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *JournalJNLS) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *JournalJNLS) SetPrice(v float64)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


