# AssetResource

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

### NewAssetResource

`func NewAssetResource() *AssetResource`

NewAssetResource instantiates a new AssetResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetResourceWithDefaults

`func NewAssetResourceWithDefaults() *AssetResource`

NewAssetResourceWithDefaults instantiates a new AssetResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssetResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssetResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClass

`func (o *AssetResource) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *AssetResource) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *AssetResource) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *AssetResource) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetExchange

`func (o *AssetResource) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *AssetResource) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *AssetResource) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *AssetResource) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetSymbol

`func (o *AssetResource) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *AssetResource) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *AssetResource) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *AssetResource) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetName

`func (o *AssetResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssetResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *AssetResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssetResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssetResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssetResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTradable

`func (o *AssetResource) GetTradable() bool`

GetTradable returns the Tradable field if non-nil, zero value otherwise.

### GetTradableOk

`func (o *AssetResource) GetTradableOk() (*bool, bool)`

GetTradableOk returns a tuple with the Tradable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradable

`func (o *AssetResource) SetTradable(v bool)`

SetTradable sets Tradable field to given value.

### HasTradable

`func (o *AssetResource) HasTradable() bool`

HasTradable returns a boolean if a field has been set.

### GetMarginable

`func (o *AssetResource) GetMarginable() bool`

GetMarginable returns the Marginable field if non-nil, zero value otherwise.

### GetMarginableOk

`func (o *AssetResource) GetMarginableOk() (*bool, bool)`

GetMarginableOk returns a tuple with the Marginable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginable

`func (o *AssetResource) SetMarginable(v bool)`

SetMarginable sets Marginable field to given value.

### HasMarginable

`func (o *AssetResource) HasMarginable() bool`

HasMarginable returns a boolean if a field has been set.

### GetShortable

`func (o *AssetResource) GetShortable() bool`

GetShortable returns the Shortable field if non-nil, zero value otherwise.

### GetShortableOk

`func (o *AssetResource) GetShortableOk() (*bool, bool)`

GetShortableOk returns a tuple with the Shortable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortable

`func (o *AssetResource) SetShortable(v bool)`

SetShortable sets Shortable field to given value.

### HasShortable

`func (o *AssetResource) HasShortable() bool`

HasShortable returns a boolean if a field has been set.

### GetEasyToBorrow

`func (o *AssetResource) GetEasyToBorrow() bool`

GetEasyToBorrow returns the EasyToBorrow field if non-nil, zero value otherwise.

### GetEasyToBorrowOk

`func (o *AssetResource) GetEasyToBorrowOk() (*bool, bool)`

GetEasyToBorrowOk returns a tuple with the EasyToBorrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasyToBorrow

`func (o *AssetResource) SetEasyToBorrow(v bool)`

SetEasyToBorrow sets EasyToBorrow field to given value.

### HasEasyToBorrow

`func (o *AssetResource) HasEasyToBorrow() bool`

HasEasyToBorrow returns a boolean if a field has been set.

### GetFractionable

`func (o *AssetResource) GetFractionable() bool`

GetFractionable returns the Fractionable field if non-nil, zero value otherwise.

### GetFractionableOk

`func (o *AssetResource) GetFractionableOk() (*bool, bool)`

GetFractionableOk returns a tuple with the Fractionable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFractionable

`func (o *AssetResource) SetFractionable(v bool)`

SetFractionable sets Fractionable field to given value.

### HasFractionable

`func (o *AssetResource) HasFractionable() bool`

HasFractionable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


