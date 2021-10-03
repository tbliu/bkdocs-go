# TradeActivity

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

## Methods

### NewTradeActivity

`func NewTradeActivity() *TradeActivity`

NewTradeActivity instantiates a new TradeActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradeActivityWithDefaults

`func NewTradeActivityWithDefaults() *TradeActivity`

NewTradeActivityWithDefaults instantiates a new TradeActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TradeActivity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TradeActivity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TradeActivity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TradeActivity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *TradeActivity) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TradeActivity) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TradeActivity) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TradeActivity) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetActivityType

`func (o *TradeActivity) GetActivityType() ActivityType`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *TradeActivity) GetActivityTypeOk() (*ActivityType, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *TradeActivity) SetActivityType(v ActivityType)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *TradeActivity) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetTransactionTime

`func (o *TradeActivity) GetTransactionTime() string`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *TradeActivity) GetTransactionTimeOk() (*string, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *TradeActivity) SetTransactionTime(v string)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *TradeActivity) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.

### GetType

`func (o *TradeActivity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TradeActivity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TradeActivity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TradeActivity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrice

`func (o *TradeActivity) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TradeActivity) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TradeActivity) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *TradeActivity) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQty

`func (o *TradeActivity) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *TradeActivity) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *TradeActivity) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *TradeActivity) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetSide

`func (o *TradeActivity) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *TradeActivity) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *TradeActivity) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *TradeActivity) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSymbol

`func (o *TradeActivity) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TradeActivity) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TradeActivity) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TradeActivity) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetLeavesQty

`func (o *TradeActivity) GetLeavesQty() string`

GetLeavesQty returns the LeavesQty field if non-nil, zero value otherwise.

### GetLeavesQtyOk

`func (o *TradeActivity) GetLeavesQtyOk() (*string, bool)`

GetLeavesQtyOk returns a tuple with the LeavesQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeavesQty

`func (o *TradeActivity) SetLeavesQty(v string)`

SetLeavesQty sets LeavesQty field to given value.

### HasLeavesQty

`func (o *TradeActivity) HasLeavesQty() bool`

HasLeavesQty returns a boolean if a field has been set.

### GetOrderId

`func (o *TradeActivity) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *TradeActivity) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *TradeActivity) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *TradeActivity) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetCumQty

`func (o *TradeActivity) GetCumQty() string`

GetCumQty returns the CumQty field if non-nil, zero value otherwise.

### GetCumQtyOk

`func (o *TradeActivity) GetCumQtyOk() (*string, bool)`

GetCumQtyOk returns a tuple with the CumQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumQty

`func (o *TradeActivity) SetCumQty(v string)`

SetCumQty sets CumQty field to given value.

### HasCumQty

`func (o *TradeActivity) HasCumQty() bool`

HasCumQty returns a boolean if a field has been set.

### GetOrderStatus

`func (o *TradeActivity) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *TradeActivity) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *TradeActivity) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *TradeActivity) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


