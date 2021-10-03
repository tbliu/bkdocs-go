# MarketDay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** |  | 
**Open** | **string** |  | 
**Close** | **string** |  | 
**SessionOpen** | Pointer to **string** |  | [optional] 
**SessionClose** | Pointer to **string** |  | [optional] 

## Methods

### NewMarketDay

`func NewMarketDay(date string, open string, close string, ) *MarketDay`

NewMarketDay instantiates a new MarketDay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketDayWithDefaults

`func NewMarketDayWithDefaults() *MarketDay`

NewMarketDayWithDefaults instantiates a new MarketDay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *MarketDay) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *MarketDay) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *MarketDay) SetDate(v string)`

SetDate sets Date field to given value.


### GetOpen

`func (o *MarketDay) GetOpen() string`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *MarketDay) GetOpenOk() (*string, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *MarketDay) SetOpen(v string)`

SetOpen sets Open field to given value.


### GetClose

`func (o *MarketDay) GetClose() string`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *MarketDay) GetCloseOk() (*string, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *MarketDay) SetClose(v string)`

SetClose sets Close field to given value.


### GetSessionOpen

`func (o *MarketDay) GetSessionOpen() string`

GetSessionOpen returns the SessionOpen field if non-nil, zero value otherwise.

### GetSessionOpenOk

`func (o *MarketDay) GetSessionOpenOk() (*string, bool)`

GetSessionOpenOk returns a tuple with the SessionOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionOpen

`func (o *MarketDay) SetSessionOpen(v string)`

SetSessionOpen sets SessionOpen field to given value.

### HasSessionOpen

`func (o *MarketDay) HasSessionOpen() bool`

HasSessionOpen returns a boolean if a field has been set.

### GetSessionClose

`func (o *MarketDay) GetSessionClose() string`

GetSessionClose returns the SessionClose field if non-nil, zero value otherwise.

### GetSessionCloseOk

`func (o *MarketDay) GetSessionCloseOk() (*string, bool)`

GetSessionCloseOk returns a tuple with the SessionClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionClose

`func (o *MarketDay) SetSessionClose(v string)`

SetSessionClose sets SessionClose field to given value.

### HasSessionClose

`func (o *MarketDay) HasSessionClose() bool`

HasSessionClose returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


