# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**AccountStatus**](AccountStatus.md) |  | [optional] 
**Currency** | Pointer to **string** | Always \&quot;USD\&quot; | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastEquity** | Pointer to **float64** |  | [optional] 
**KycResults** | Pointer to [**KYCResult**](KYCResult.md) |  | [optional] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Account) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Account) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountNumber

`func (o *Account) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Account) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Account) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *Account) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetStatus

`func (o *Account) GetStatus() AccountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Account) GetStatusOk() (*AccountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Account) SetStatus(v AccountStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Account) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCurrency

`func (o *Account) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Account) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Account) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Account) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Account) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Account) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Account) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Account) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastEquity

`func (o *Account) GetLastEquity() float64`

GetLastEquity returns the LastEquity field if non-nil, zero value otherwise.

### GetLastEquityOk

`func (o *Account) GetLastEquityOk() (*float64, bool)`

GetLastEquityOk returns a tuple with the LastEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEquity

`func (o *Account) SetLastEquity(v float64)`

SetLastEquity sets LastEquity field to given value.

### HasLastEquity

`func (o *Account) HasLastEquity() bool`

HasLastEquity returns a boolean if a field has been set.

### GetKycResults

`func (o *Account) GetKycResults() KYCResult`

GetKycResults returns the KycResults field if non-nil, zero value otherwise.

### GetKycResultsOk

`func (o *Account) GetKycResultsOk() (*KYCResult, bool)`

GetKycResultsOk returns a tuple with the KycResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycResults

`func (o *Account) SetKycResults(v KYCResult)`

SetKycResults sets KycResults field to given value.

### HasKycResults

`func (o *Account) HasKycResults() bool`

HasKycResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


