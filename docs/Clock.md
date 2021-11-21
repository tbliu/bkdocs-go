# Clock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **string** |  | 
**IsOpen** | **bool** |  | 
**NextOpen** | **time.Time** |  | 
**NextClose** | **time.Time** |  | 

## Methods

### NewClock

`func NewClock(timestamp string, isOpen bool, nextOpen time.Time, nextClose time.Time, ) *Clock`

NewClock instantiates a new Clock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClockWithDefaults

`func NewClockWithDefaults() *Clock`

NewClockWithDefaults instantiates a new Clock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *Clock) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Clock) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Clock) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetIsOpen

`func (o *Clock) GetIsOpen() bool`

GetIsOpen returns the IsOpen field if non-nil, zero value otherwise.

### GetIsOpenOk

`func (o *Clock) GetIsOpenOk() (*bool, bool)`

GetIsOpenOk returns a tuple with the IsOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpen

`func (o *Clock) SetIsOpen(v bool)`

SetIsOpen sets IsOpen field to given value.


### GetNextOpen

`func (o *Clock) GetNextOpen() time.Time`

GetNextOpen returns the NextOpen field if non-nil, zero value otherwise.

### GetNextOpenOk

`func (o *Clock) GetNextOpenOk() (*time.Time, bool)`

GetNextOpenOk returns a tuple with the NextOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOpen

`func (o *Clock) SetNextOpen(v time.Time)`

SetNextOpen sets NextOpen field to given value.


### GetNextClose

`func (o *Clock) GetNextClose() time.Time`

GetNextClose returns the NextClose field if non-nil, zero value otherwise.

### GetNextCloseOk

`func (o *Clock) GetNextCloseOk() (*time.Time, bool)`

GetNextCloseOk returns a tuple with the NextClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextClose

`func (o *Clock) SetNextClose(v time.Time)`

SetNextClose sets NextClose field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


