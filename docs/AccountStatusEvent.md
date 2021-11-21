# AccountStatusEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**AccountNumber** | **string** |  | 
**StatusFrom** | **string** |  | 
**StatusTo** | **string** |  | 
**Reason** | **string** |  | 
**At** | **string** |  | 
**KycResult** | Pointer to [**KYCResult**](KYCResult.md) |  | [optional] 

## Methods

### NewAccountStatusEvent

`func NewAccountStatusEvent(accountId string, accountNumber string, statusFrom string, statusTo string, reason string, at string, ) *AccountStatusEvent`

NewAccountStatusEvent instantiates a new AccountStatusEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountStatusEventWithDefaults

`func NewAccountStatusEventWithDefaults() *AccountStatusEvent`

NewAccountStatusEventWithDefaults instantiates a new AccountStatusEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountStatusEvent) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountStatusEvent) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountStatusEvent) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAccountNumber

`func (o *AccountStatusEvent) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AccountStatusEvent) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AccountStatusEvent) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetStatusFrom

`func (o *AccountStatusEvent) GetStatusFrom() string`

GetStatusFrom returns the StatusFrom field if non-nil, zero value otherwise.

### GetStatusFromOk

`func (o *AccountStatusEvent) GetStatusFromOk() (*string, bool)`

GetStatusFromOk returns a tuple with the StatusFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusFrom

`func (o *AccountStatusEvent) SetStatusFrom(v string)`

SetStatusFrom sets StatusFrom field to given value.


### GetStatusTo

`func (o *AccountStatusEvent) GetStatusTo() string`

GetStatusTo returns the StatusTo field if non-nil, zero value otherwise.

### GetStatusToOk

`func (o *AccountStatusEvent) GetStatusToOk() (*string, bool)`

GetStatusToOk returns a tuple with the StatusTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTo

`func (o *AccountStatusEvent) SetStatusTo(v string)`

SetStatusTo sets StatusTo field to given value.


### GetReason

`func (o *AccountStatusEvent) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AccountStatusEvent) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AccountStatusEvent) SetReason(v string)`

SetReason sets Reason field to given value.


### GetAt

`func (o *AccountStatusEvent) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *AccountStatusEvent) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *AccountStatusEvent) SetAt(v string)`

SetAt sets At field to given value.


### GetKycResult

`func (o *AccountStatusEvent) GetKycResult() KYCResult`

GetKycResult returns the KycResult field if non-nil, zero value otherwise.

### GetKycResultOk

`func (o *AccountStatusEvent) GetKycResultOk() (*KYCResult, bool)`

GetKycResultOk returns a tuple with the KycResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycResult

`func (o *AccountStatusEvent) SetKycResult(v KYCResult)`

SetKycResult sets KycResult field to given value.

### HasKycResult

`func (o *AccountStatusEvent) HasKycResult() bool`

HasKycResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


