# JournalData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryType** | **string** |  | 
**FromAccount** | **string** |  | 
**ToAccount** | **string** |  | 
**Amount** | Pointer to **float64** | Required for JNLC. The dollar amount to move. It has to be a positive value.  | [optional] 
**Symbol** | Pointer to **string** | Required for JNLS.  | [optional] 
**Qty** | Pointer to **float64** | Required for JNLS. The number of shares to move. It has to be a positive value.  | [optional] 

## Methods

### NewJournalData

`func NewJournalData(entryType string, fromAccount string, toAccount string, ) *JournalData`

NewJournalData instantiates a new JournalData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalDataWithDefaults

`func NewJournalDataWithDefaults() *JournalData`

NewJournalDataWithDefaults instantiates a new JournalData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryType

`func (o *JournalData) GetEntryType() string`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *JournalData) GetEntryTypeOk() (*string, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *JournalData) SetEntryType(v string)`

SetEntryType sets EntryType field to given value.


### GetFromAccount

`func (o *JournalData) GetFromAccount() string`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *JournalData) GetFromAccountOk() (*string, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *JournalData) SetFromAccount(v string)`

SetFromAccount sets FromAccount field to given value.


### GetToAccount

`func (o *JournalData) GetToAccount() string`

GetToAccount returns the ToAccount field if non-nil, zero value otherwise.

### GetToAccountOk

`func (o *JournalData) GetToAccountOk() (*string, bool)`

GetToAccountOk returns a tuple with the ToAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccount

`func (o *JournalData) SetToAccount(v string)`

SetToAccount sets ToAccount field to given value.


### GetAmount

`func (o *JournalData) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *JournalData) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *JournalData) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *JournalData) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetSymbol

`func (o *JournalData) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *JournalData) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *JournalData) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *JournalData) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetQty

`func (o *JournalData) GetQty() float64`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *JournalData) GetQtyOk() (*float64, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *JournalData) SetQty(v float64)`

SetQty sets Qty field to given value.

### HasQty

`func (o *JournalData) HasQty() bool`

HasQty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


