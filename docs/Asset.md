# Asset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Class** | Pointer to **string** |  | [optional] 
**Exchange** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tradable** | Pointer to **bool** |  | [optional] 
**Marginable** | Pointer to **bool** |  | [optional] 
**Shortable** | Pointer to **bool** |  | [optional] 
**EasyToBorrow** | Pointer to **bool** |  | [optional] 
**Fractionable** | Pointer to **bool** |  | [optional] 

## Methods

### NewAsset

`func NewAsset() *Asset`

NewAsset instantiates a new Asset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWithDefaults

`func NewAssetWithDefaults() *Asset`

NewAssetWithDefaults instantiates a new Asset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Asset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Asset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Asset) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Asset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClass

`func (o *Asset) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *Asset) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *Asset) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *Asset) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetExchange

`func (o *Asset) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *Asset) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *Asset) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *Asset) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetSymbol

`func (o *Asset) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Asset) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Asset) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Asset) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetName

`func (o *Asset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Asset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Asset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Asset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *Asset) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Asset) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Asset) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Asset) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTradable

`func (o *Asset) GetTradable() bool`

GetTradable returns the Tradable field if non-nil, zero value otherwise.

### GetTradableOk

`func (o *Asset) GetTradableOk() (*bool, bool)`

GetTradableOk returns a tuple with the Tradable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradable

`func (o *Asset) SetTradable(v bool)`

SetTradable sets Tradable field to given value.

### HasTradable

`func (o *Asset) HasTradable() bool`

HasTradable returns a boolean if a field has been set.

### GetMarginable

`func (o *Asset) GetMarginable() bool`

GetMarginable returns the Marginable field if non-nil, zero value otherwise.

### GetMarginableOk

`func (o *Asset) GetMarginableOk() (*bool, bool)`

GetMarginableOk returns a tuple with the Marginable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginable

`func (o *Asset) SetMarginable(v bool)`

SetMarginable sets Marginable field to given value.

### HasMarginable

`func (o *Asset) HasMarginable() bool`

HasMarginable returns a boolean if a field has been set.

### GetShortable

`func (o *Asset) GetShortable() bool`

GetShortable returns the Shortable field if non-nil, zero value otherwise.

### GetShortableOk

`func (o *Asset) GetShortableOk() (*bool, bool)`

GetShortableOk returns a tuple with the Shortable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortable

`func (o *Asset) SetShortable(v bool)`

SetShortable sets Shortable field to given value.

### HasShortable

`func (o *Asset) HasShortable() bool`

HasShortable returns a boolean if a field has been set.

### GetEasyToBorrow

`func (o *Asset) GetEasyToBorrow() bool`

GetEasyToBorrow returns the EasyToBorrow field if non-nil, zero value otherwise.

### GetEasyToBorrowOk

`func (o *Asset) GetEasyToBorrowOk() (*bool, bool)`

GetEasyToBorrowOk returns a tuple with the EasyToBorrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasyToBorrow

`func (o *Asset) SetEasyToBorrow(v bool)`

SetEasyToBorrow sets EasyToBorrow field to given value.

### HasEasyToBorrow

`func (o *Asset) HasEasyToBorrow() bool`

HasEasyToBorrow returns a boolean if a field has been set.

### GetFractionable

`func (o *Asset) GetFractionable() bool`

GetFractionable returns the Fractionable field if non-nil, zero value otherwise.

### GetFractionableOk

`func (o *Asset) GetFractionableOk() (*bool, bool)`

GetFractionableOk returns a tuple with the Fractionable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFractionable

`func (o *Asset) SetFractionable(v bool)`

SetFractionable sets Fractionable field to given value.

### HasFractionable

`func (o *Asset) HasFractionable() bool`

HasFractionable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


