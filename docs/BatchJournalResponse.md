# BatchJournalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ErrorMessage** | **string** |  | 
**EntryType** | **string** |  | 
**FromAccount** | **string** |  | 
**ToAccount** | **string** |  | 
**Symbol** | **string** |  | 
**Qty** | **NullableString** |  | 
**Price** | **string** |  | 
**Status** | **string** |  | 
**SettleDate** | **NullableString** |  | 
**SystemDate** | **NullableString** |  | 
**NetAmount** | **string** |  | 
**Description** | **string** |  | 

## Methods

### NewBatchJournalResponse

`func NewBatchJournalResponse(id string, errorMessage string, entryType string, fromAccount string, toAccount string, symbol string, qty NullableString, price string, status string, settleDate NullableString, systemDate NullableString, netAmount string, description string, ) *BatchJournalResponse`

NewBatchJournalResponse instantiates a new BatchJournalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchJournalResponseWithDefaults

`func NewBatchJournalResponseWithDefaults() *BatchJournalResponse`

NewBatchJournalResponseWithDefaults instantiates a new BatchJournalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BatchJournalResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchJournalResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchJournalResponse) SetId(v string)`

SetId sets Id field to given value.


### GetErrorMessage

`func (o *BatchJournalResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BatchJournalResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BatchJournalResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### GetEntryType

`func (o *BatchJournalResponse) GetEntryType() string`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *BatchJournalResponse) GetEntryTypeOk() (*string, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *BatchJournalResponse) SetEntryType(v string)`

SetEntryType sets EntryType field to given value.


### GetFromAccount

`func (o *BatchJournalResponse) GetFromAccount() string`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *BatchJournalResponse) GetFromAccountOk() (*string, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *BatchJournalResponse) SetFromAccount(v string)`

SetFromAccount sets FromAccount field to given value.


### GetToAccount

`func (o *BatchJournalResponse) GetToAccount() string`

GetToAccount returns the ToAccount field if non-nil, zero value otherwise.

### GetToAccountOk

`func (o *BatchJournalResponse) GetToAccountOk() (*string, bool)`

GetToAccountOk returns a tuple with the ToAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccount

`func (o *BatchJournalResponse) SetToAccount(v string)`

SetToAccount sets ToAccount field to given value.


### GetSymbol

`func (o *BatchJournalResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *BatchJournalResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *BatchJournalResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetQty

`func (o *BatchJournalResponse) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *BatchJournalResponse) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *BatchJournalResponse) SetQty(v string)`

SetQty sets Qty field to given value.


### SetQtyNil

`func (o *BatchJournalResponse) SetQtyNil(b bool)`

 SetQtyNil sets the value for Qty to be an explicit nil

### UnsetQty
`func (o *BatchJournalResponse) UnsetQty()`

UnsetQty ensures that no value is present for Qty, not even an explicit nil
### GetPrice

`func (o *BatchJournalResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BatchJournalResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BatchJournalResponse) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetStatus

`func (o *BatchJournalResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchJournalResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchJournalResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSettleDate

`func (o *BatchJournalResponse) GetSettleDate() string`

GetSettleDate returns the SettleDate field if non-nil, zero value otherwise.

### GetSettleDateOk

`func (o *BatchJournalResponse) GetSettleDateOk() (*string, bool)`

GetSettleDateOk returns a tuple with the SettleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettleDate

`func (o *BatchJournalResponse) SetSettleDate(v string)`

SetSettleDate sets SettleDate field to given value.


### SetSettleDateNil

`func (o *BatchJournalResponse) SetSettleDateNil(b bool)`

 SetSettleDateNil sets the value for SettleDate to be an explicit nil

### UnsetSettleDate
`func (o *BatchJournalResponse) UnsetSettleDate()`

UnsetSettleDate ensures that no value is present for SettleDate, not even an explicit nil
### GetSystemDate

`func (o *BatchJournalResponse) GetSystemDate() string`

GetSystemDate returns the SystemDate field if non-nil, zero value otherwise.

### GetSystemDateOk

`func (o *BatchJournalResponse) GetSystemDateOk() (*string, bool)`

GetSystemDateOk returns a tuple with the SystemDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDate

`func (o *BatchJournalResponse) SetSystemDate(v string)`

SetSystemDate sets SystemDate field to given value.


### SetSystemDateNil

`func (o *BatchJournalResponse) SetSystemDateNil(b bool)`

 SetSystemDateNil sets the value for SystemDate to be an explicit nil

### UnsetSystemDate
`func (o *BatchJournalResponse) UnsetSystemDate()`

UnsetSystemDate ensures that no value is present for SystemDate, not even an explicit nil
### GetNetAmount

`func (o *BatchJournalResponse) GetNetAmount() string`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *BatchJournalResponse) GetNetAmountOk() (*string, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *BatchJournalResponse) SetNetAmount(v string)`

SetNetAmount sets NetAmount field to given value.


### GetDescription

`func (o *BatchJournalResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BatchJournalResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BatchJournalResponse) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


