# ActivityItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**ActivityType** | Pointer to [**ActivityType**](ActivityType.md) |  | [optional] 
**TransactionTime** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Qty** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**LeavesQty** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 
**CumQty** | Pointer to **string** |  | [optional] 
**OrderStatus** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**NetAmount** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**PerShareAmount** | Pointer to **string** |  | [optional] 

## Methods

### NewActivityItem

`func NewActivityItem() *ActivityItem`

NewActivityItem instantiates a new ActivityItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityItemWithDefaults

`func NewActivityItemWithDefaults() *ActivityItem`

NewActivityItemWithDefaults instantiates a new ActivityItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivityItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *ActivityItem) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ActivityItem) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ActivityItem) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ActivityItem) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetActivityType

`func (o *ActivityItem) GetActivityType() ActivityType`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *ActivityItem) GetActivityTypeOk() (*ActivityType, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *ActivityItem) SetActivityType(v ActivityType)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *ActivityItem) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetTransactionTime

`func (o *ActivityItem) GetTransactionTime() string`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *ActivityItem) GetTransactionTimeOk() (*string, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *ActivityItem) SetTransactionTime(v string)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *ActivityItem) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.

### GetType

`func (o *ActivityItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActivityItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrice

`func (o *ActivityItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ActivityItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ActivityItem) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ActivityItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQty

`func (o *ActivityItem) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *ActivityItem) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *ActivityItem) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *ActivityItem) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetSide

`func (o *ActivityItem) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *ActivityItem) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *ActivityItem) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *ActivityItem) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSymbol

`func (o *ActivityItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ActivityItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ActivityItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ActivityItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetLeavesQty

`func (o *ActivityItem) GetLeavesQty() string`

GetLeavesQty returns the LeavesQty field if non-nil, zero value otherwise.

### GetLeavesQtyOk

`func (o *ActivityItem) GetLeavesQtyOk() (*string, bool)`

GetLeavesQtyOk returns a tuple with the LeavesQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeavesQty

`func (o *ActivityItem) SetLeavesQty(v string)`

SetLeavesQty sets LeavesQty field to given value.

### HasLeavesQty

`func (o *ActivityItem) HasLeavesQty() bool`

HasLeavesQty returns a boolean if a field has been set.

### GetOrderId

`func (o *ActivityItem) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ActivityItem) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ActivityItem) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ActivityItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetCumQty

`func (o *ActivityItem) GetCumQty() string`

GetCumQty returns the CumQty field if non-nil, zero value otherwise.

### GetCumQtyOk

`func (o *ActivityItem) GetCumQtyOk() (*string, bool)`

GetCumQtyOk returns a tuple with the CumQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumQty

`func (o *ActivityItem) SetCumQty(v string)`

SetCumQty sets CumQty field to given value.

### HasCumQty

`func (o *ActivityItem) HasCumQty() bool`

HasCumQty returns a boolean if a field has been set.

### GetOrderStatus

`func (o *ActivityItem) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *ActivityItem) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *ActivityItem) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *ActivityItem) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetDate

`func (o *ActivityItem) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ActivityItem) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ActivityItem) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *ActivityItem) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetNetAmount

`func (o *ActivityItem) GetNetAmount() string`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *ActivityItem) GetNetAmountOk() (*string, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *ActivityItem) SetNetAmount(v string)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmount

`func (o *ActivityItem) HasNetAmount() bool`

HasNetAmount returns a boolean if a field has been set.

### GetDescription

`func (o *ActivityItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ActivityItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ActivityItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ActivityItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *ActivityItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActivityItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActivityItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActivityItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPerShareAmount

`func (o *ActivityItem) GetPerShareAmount() string`

GetPerShareAmount returns the PerShareAmount field if non-nil, zero value otherwise.

### GetPerShareAmountOk

`func (o *ActivityItem) GetPerShareAmountOk() (*string, bool)`

GetPerShareAmountOk returns a tuple with the PerShareAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerShareAmount

`func (o *ActivityItem) SetPerShareAmount(v string)`

SetPerShareAmount sets PerShareAmount field to given value.

### HasPerShareAmount

`func (o *ActivityItem) HasPerShareAmount() bool`

HasPerShareAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


