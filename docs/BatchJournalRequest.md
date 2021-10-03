# BatchJournalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryType** | **string** |  | 
**FromAccount** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Entries** | Pointer to [**[]BatchJournalRequestEntries**](BatchJournalRequestEntries.md) |  | [optional] 

## Methods

### NewBatchJournalRequest

`func NewBatchJournalRequest(entryType string, fromAccount string, ) *BatchJournalRequest`

NewBatchJournalRequest instantiates a new BatchJournalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchJournalRequestWithDefaults

`func NewBatchJournalRequestWithDefaults() *BatchJournalRequest`

NewBatchJournalRequestWithDefaults instantiates a new BatchJournalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryType

`func (o *BatchJournalRequest) GetEntryType() string`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *BatchJournalRequest) GetEntryTypeOk() (*string, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *BatchJournalRequest) SetEntryType(v string)`

SetEntryType sets EntryType field to given value.


### GetFromAccount

`func (o *BatchJournalRequest) GetFromAccount() string`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *BatchJournalRequest) GetFromAccountOk() (*string, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *BatchJournalRequest) SetFromAccount(v string)`

SetFromAccount sets FromAccount field to given value.


### GetDescription

`func (o *BatchJournalRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BatchJournalRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BatchJournalRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BatchJournalRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntries

`func (o *BatchJournalRequest) GetEntries() []BatchJournalRequestEntries`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *BatchJournalRequest) GetEntriesOk() (*[]BatchJournalRequestEntries, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *BatchJournalRequest) SetEntries(v []BatchJournalRequestEntries)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *BatchJournalRequest) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


