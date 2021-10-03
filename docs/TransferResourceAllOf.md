# TransferResourceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Status** | **string** |  | 
**AccountId** | **string** |  | 
**Reason** | Pointer to **NullableString** |  | [optional] 
**ExpiresAt** | **time.Time** |  | 

## Methods

### NewTransferResourceAllOf

`func NewTransferResourceAllOf(type_ string, status string, accountId string, expiresAt time.Time, ) *TransferResourceAllOf`

NewTransferResourceAllOf instantiates a new TransferResourceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferResourceAllOfWithDefaults

`func NewTransferResourceAllOfWithDefaults() *TransferResourceAllOf`

NewTransferResourceAllOfWithDefaults instantiates a new TransferResourceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransferResourceAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransferResourceAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransferResourceAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *TransferResourceAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransferResourceAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransferResourceAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAccountId

`func (o *TransferResourceAllOf) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransferResourceAllOf) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransferResourceAllOf) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetReason

`func (o *TransferResourceAllOf) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TransferResourceAllOf) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TransferResourceAllOf) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TransferResourceAllOf) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *TransferResourceAllOf) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *TransferResourceAllOf) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetExpiresAt

`func (o *TransferResourceAllOf) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TransferResourceAllOf) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TransferResourceAllOf) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


