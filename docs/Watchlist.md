# Watchlist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the watchlist itself.  | [optional] 
**AccountId** | Pointer to **string** | Unique identifier of the account that owns this watchlist.  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**** | Pointer to **string** |  | [optional] 
**Assets** | Pointer to [**[]Asset**](Asset.md) |  | [optional] 

## Methods

### NewWatchlist

`func NewWatchlist() *Watchlist`

NewWatchlist instantiates a new Watchlist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatchlistWithDefaults

`func NewWatchlistWithDefaults() *Watchlist`

NewWatchlistWithDefaults instantiates a new Watchlist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Watchlist) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Watchlist) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Watchlist) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Watchlist) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *Watchlist) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Watchlist) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Watchlist) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Watchlist) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Watchlist) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Watchlist) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Watchlist) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Watchlist) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Watchlist) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Watchlist) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Watchlist) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Watchlist) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### Get

`func (o *Watchlist) Get() string`

Get returns the  field if non-nil, zero value otherwise.

### GetOk

`func (o *Watchlist) GetOk() (*string, bool)`

GetOk returns a tuple with the  field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### Set

`func (o *Watchlist) Set(v string)`

Set sets  field to given value.

### Has

`func (o *Watchlist) Has() bool`

Has returns a boolean if a field has been set.

### GetAssets

`func (o *Watchlist) GetAssets() []Asset`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *Watchlist) GetAssetsOk() (*[]Asset, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *Watchlist) SetAssets(v []Asset)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *Watchlist) HasAssets() bool`

HasAssets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


