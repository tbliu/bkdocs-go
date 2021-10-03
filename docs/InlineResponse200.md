# InlineResponse200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**AccountStatus**](AccountStatus.md) |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**BuyingPower** | Pointer to **string** |  | [optional] 
**RegtBuyingPower** | Pointer to **string** |  | [optional] 
**DaytradingBuyingPower** | Pointer to **string** |  | [optional] 
**Cash** | Pointer to **string** |  | [optional] 
**CashWithdrawable** | Pointer to **string** |  | [optional] 
**CashTransferable** | Pointer to **string** |  | [optional] 
**PendingTransferOut** | Pointer to **string** |  | [optional] 
**PortfolioValue** | Pointer to **string** |  | [optional] 
**PatternDayTrader** | Pointer to **bool** |  | [optional] 
**TradingBlocked** | Pointer to **bool** |  | [optional] 
**TransfersBlocked** | Pointer to **bool** |  | [optional] 
**AccountBlocked** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**TradeSuspendedByUser** | Pointer to **bool** |  | [optional] 
**Multiplier** | Pointer to **string** |  | [optional] 
**ShortingEnabled** | Pointer to **bool** |  | [optional] 
**Equity** | Pointer to **string** |  | [optional] 
**LastEquity** | Pointer to **string** |  | [optional] 
**LongMarketValue** | Pointer to **string** |  | [optional] 
**ShortMarketValue** | Pointer to **string** |  | [optional] 
**InitialMargin** | Pointer to **string** |  | [optional] 
**MaintenanceMargin** | Pointer to **string** |  | [optional] 
**LastMaintenanceMargin** | Pointer to **string** |  | [optional] 
**Sma** | Pointer to **string** |  | [optional] 
**DaytradeCount** | Pointer to **int32** |  | [optional] 
**PreviousClose** | Pointer to **string** |  | [optional] 
**LastLongMarketValue** | Pointer to **string** |  | [optional] 
**LastShortMarketValue** | Pointer to **string** |  | [optional] 
**LastCash** | Pointer to **string** |  | [optional] 
**LastInitialMargin** | Pointer to **string** |  | [optional] 
**LastRegtBuyingPower** | Pointer to **string** |  | [optional] 
**LastDaytradingBuyingPower** | Pointer to **string** |  | [optional] 
**LastBuyingPower** | Pointer to **string** |  | [optional] 
**LastDaytradeCount** | Pointer to **int32** |  | [optional] 
**ClearingBroker** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineResponse200

`func NewInlineResponse200() *InlineResponse200`

NewInlineResponse200 instantiates a new InlineResponse200 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200WithDefaults

`func NewInlineResponse200WithDefaults() *InlineResponse200`

NewInlineResponse200WithDefaults instantiates a new InlineResponse200 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountNumber

`func (o *InlineResponse200) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *InlineResponse200) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *InlineResponse200) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *InlineResponse200) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200) GetStatus() AccountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200) GetStatusOk() (*AccountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200) SetStatus(v AccountStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCurrency

`func (o *InlineResponse200) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InlineResponse200) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InlineResponse200) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InlineResponse200) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetBuyingPower

`func (o *InlineResponse200) GetBuyingPower() string`

GetBuyingPower returns the BuyingPower field if non-nil, zero value otherwise.

### GetBuyingPowerOk

`func (o *InlineResponse200) GetBuyingPowerOk() (*string, bool)`

GetBuyingPowerOk returns a tuple with the BuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyingPower

`func (o *InlineResponse200) SetBuyingPower(v string)`

SetBuyingPower sets BuyingPower field to given value.

### HasBuyingPower

`func (o *InlineResponse200) HasBuyingPower() bool`

HasBuyingPower returns a boolean if a field has been set.

### GetRegtBuyingPower

`func (o *InlineResponse200) GetRegtBuyingPower() string`

GetRegtBuyingPower returns the RegtBuyingPower field if non-nil, zero value otherwise.

### GetRegtBuyingPowerOk

`func (o *InlineResponse200) GetRegtBuyingPowerOk() (*string, bool)`

GetRegtBuyingPowerOk returns a tuple with the RegtBuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegtBuyingPower

`func (o *InlineResponse200) SetRegtBuyingPower(v string)`

SetRegtBuyingPower sets RegtBuyingPower field to given value.

### HasRegtBuyingPower

`func (o *InlineResponse200) HasRegtBuyingPower() bool`

HasRegtBuyingPower returns a boolean if a field has been set.

### GetDaytradingBuyingPower

`func (o *InlineResponse200) GetDaytradingBuyingPower() string`

GetDaytradingBuyingPower returns the DaytradingBuyingPower field if non-nil, zero value otherwise.

### GetDaytradingBuyingPowerOk

`func (o *InlineResponse200) GetDaytradingBuyingPowerOk() (*string, bool)`

GetDaytradingBuyingPowerOk returns a tuple with the DaytradingBuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaytradingBuyingPower

`func (o *InlineResponse200) SetDaytradingBuyingPower(v string)`

SetDaytradingBuyingPower sets DaytradingBuyingPower field to given value.

### HasDaytradingBuyingPower

`func (o *InlineResponse200) HasDaytradingBuyingPower() bool`

HasDaytradingBuyingPower returns a boolean if a field has been set.

### GetCash

`func (o *InlineResponse200) GetCash() string`

GetCash returns the Cash field if non-nil, zero value otherwise.

### GetCashOk

`func (o *InlineResponse200) GetCashOk() (*string, bool)`

GetCashOk returns a tuple with the Cash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCash

`func (o *InlineResponse200) SetCash(v string)`

SetCash sets Cash field to given value.

### HasCash

`func (o *InlineResponse200) HasCash() bool`

HasCash returns a boolean if a field has been set.

### GetCashWithdrawable

`func (o *InlineResponse200) GetCashWithdrawable() string`

GetCashWithdrawable returns the CashWithdrawable field if non-nil, zero value otherwise.

### GetCashWithdrawableOk

`func (o *InlineResponse200) GetCashWithdrawableOk() (*string, bool)`

GetCashWithdrawableOk returns a tuple with the CashWithdrawable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashWithdrawable

`func (o *InlineResponse200) SetCashWithdrawable(v string)`

SetCashWithdrawable sets CashWithdrawable field to given value.

### HasCashWithdrawable

`func (o *InlineResponse200) HasCashWithdrawable() bool`

HasCashWithdrawable returns a boolean if a field has been set.

### GetCashTransferable

`func (o *InlineResponse200) GetCashTransferable() string`

GetCashTransferable returns the CashTransferable field if non-nil, zero value otherwise.

### GetCashTransferableOk

`func (o *InlineResponse200) GetCashTransferableOk() (*string, bool)`

GetCashTransferableOk returns a tuple with the CashTransferable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashTransferable

`func (o *InlineResponse200) SetCashTransferable(v string)`

SetCashTransferable sets CashTransferable field to given value.

### HasCashTransferable

`func (o *InlineResponse200) HasCashTransferable() bool`

HasCashTransferable returns a boolean if a field has been set.

### GetPendingTransferOut

`func (o *InlineResponse200) GetPendingTransferOut() string`

GetPendingTransferOut returns the PendingTransferOut field if non-nil, zero value otherwise.

### GetPendingTransferOutOk

`func (o *InlineResponse200) GetPendingTransferOutOk() (*string, bool)`

GetPendingTransferOutOk returns a tuple with the PendingTransferOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingTransferOut

`func (o *InlineResponse200) SetPendingTransferOut(v string)`

SetPendingTransferOut sets PendingTransferOut field to given value.

### HasPendingTransferOut

`func (o *InlineResponse200) HasPendingTransferOut() bool`

HasPendingTransferOut returns a boolean if a field has been set.

### GetPortfolioValue

`func (o *InlineResponse200) GetPortfolioValue() string`

GetPortfolioValue returns the PortfolioValue field if non-nil, zero value otherwise.

### GetPortfolioValueOk

`func (o *InlineResponse200) GetPortfolioValueOk() (*string, bool)`

GetPortfolioValueOk returns a tuple with the PortfolioValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolioValue

`func (o *InlineResponse200) SetPortfolioValue(v string)`

SetPortfolioValue sets PortfolioValue field to given value.

### HasPortfolioValue

`func (o *InlineResponse200) HasPortfolioValue() bool`

HasPortfolioValue returns a boolean if a field has been set.

### GetPatternDayTrader

`func (o *InlineResponse200) GetPatternDayTrader() bool`

GetPatternDayTrader returns the PatternDayTrader field if non-nil, zero value otherwise.

### GetPatternDayTraderOk

`func (o *InlineResponse200) GetPatternDayTraderOk() (*bool, bool)`

GetPatternDayTraderOk returns a tuple with the PatternDayTrader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternDayTrader

`func (o *InlineResponse200) SetPatternDayTrader(v bool)`

SetPatternDayTrader sets PatternDayTrader field to given value.

### HasPatternDayTrader

`func (o *InlineResponse200) HasPatternDayTrader() bool`

HasPatternDayTrader returns a boolean if a field has been set.

### GetTradingBlocked

`func (o *InlineResponse200) GetTradingBlocked() bool`

GetTradingBlocked returns the TradingBlocked field if non-nil, zero value otherwise.

### GetTradingBlockedOk

`func (o *InlineResponse200) GetTradingBlockedOk() (*bool, bool)`

GetTradingBlockedOk returns a tuple with the TradingBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingBlocked

`func (o *InlineResponse200) SetTradingBlocked(v bool)`

SetTradingBlocked sets TradingBlocked field to given value.

### HasTradingBlocked

`func (o *InlineResponse200) HasTradingBlocked() bool`

HasTradingBlocked returns a boolean if a field has been set.

### GetTransfersBlocked

`func (o *InlineResponse200) GetTransfersBlocked() bool`

GetTransfersBlocked returns the TransfersBlocked field if non-nil, zero value otherwise.

### GetTransfersBlockedOk

`func (o *InlineResponse200) GetTransfersBlockedOk() (*bool, bool)`

GetTransfersBlockedOk returns a tuple with the TransfersBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfersBlocked

`func (o *InlineResponse200) SetTransfersBlocked(v bool)`

SetTransfersBlocked sets TransfersBlocked field to given value.

### HasTransfersBlocked

`func (o *InlineResponse200) HasTransfersBlocked() bool`

HasTransfersBlocked returns a boolean if a field has been set.

### GetAccountBlocked

`func (o *InlineResponse200) GetAccountBlocked() bool`

GetAccountBlocked returns the AccountBlocked field if non-nil, zero value otherwise.

### GetAccountBlockedOk

`func (o *InlineResponse200) GetAccountBlockedOk() (*bool, bool)`

GetAccountBlockedOk returns a tuple with the AccountBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountBlocked

`func (o *InlineResponse200) SetAccountBlocked(v bool)`

SetAccountBlocked sets AccountBlocked field to given value.

### HasAccountBlocked

`func (o *InlineResponse200) HasAccountBlocked() bool`

HasAccountBlocked returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetTradeSuspendedByUser

`func (o *InlineResponse200) GetTradeSuspendedByUser() bool`

GetTradeSuspendedByUser returns the TradeSuspendedByUser field if non-nil, zero value otherwise.

### GetTradeSuspendedByUserOk

`func (o *InlineResponse200) GetTradeSuspendedByUserOk() (*bool, bool)`

GetTradeSuspendedByUserOk returns a tuple with the TradeSuspendedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeSuspendedByUser

`func (o *InlineResponse200) SetTradeSuspendedByUser(v bool)`

SetTradeSuspendedByUser sets TradeSuspendedByUser field to given value.

### HasTradeSuspendedByUser

`func (o *InlineResponse200) HasTradeSuspendedByUser() bool`

HasTradeSuspendedByUser returns a boolean if a field has been set.

### GetMultiplier

`func (o *InlineResponse200) GetMultiplier() string`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *InlineResponse200) GetMultiplierOk() (*string, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *InlineResponse200) SetMultiplier(v string)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *InlineResponse200) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### GetShortingEnabled

`func (o *InlineResponse200) GetShortingEnabled() bool`

GetShortingEnabled returns the ShortingEnabled field if non-nil, zero value otherwise.

### GetShortingEnabledOk

`func (o *InlineResponse200) GetShortingEnabledOk() (*bool, bool)`

GetShortingEnabledOk returns a tuple with the ShortingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortingEnabled

`func (o *InlineResponse200) SetShortingEnabled(v bool)`

SetShortingEnabled sets ShortingEnabled field to given value.

### HasShortingEnabled

`func (o *InlineResponse200) HasShortingEnabled() bool`

HasShortingEnabled returns a boolean if a field has been set.

### GetEquity

`func (o *InlineResponse200) GetEquity() string`

GetEquity returns the Equity field if non-nil, zero value otherwise.

### GetEquityOk

`func (o *InlineResponse200) GetEquityOk() (*string, bool)`

GetEquityOk returns a tuple with the Equity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquity

`func (o *InlineResponse200) SetEquity(v string)`

SetEquity sets Equity field to given value.

### HasEquity

`func (o *InlineResponse200) HasEquity() bool`

HasEquity returns a boolean if a field has been set.

### GetLastEquity

`func (o *InlineResponse200) GetLastEquity() string`

GetLastEquity returns the LastEquity field if non-nil, zero value otherwise.

### GetLastEquityOk

`func (o *InlineResponse200) GetLastEquityOk() (*string, bool)`

GetLastEquityOk returns a tuple with the LastEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEquity

`func (o *InlineResponse200) SetLastEquity(v string)`

SetLastEquity sets LastEquity field to given value.

### HasLastEquity

`func (o *InlineResponse200) HasLastEquity() bool`

HasLastEquity returns a boolean if a field has been set.

### GetLongMarketValue

`func (o *InlineResponse200) GetLongMarketValue() string`

GetLongMarketValue returns the LongMarketValue field if non-nil, zero value otherwise.

### GetLongMarketValueOk

`func (o *InlineResponse200) GetLongMarketValueOk() (*string, bool)`

GetLongMarketValueOk returns a tuple with the LongMarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongMarketValue

`func (o *InlineResponse200) SetLongMarketValue(v string)`

SetLongMarketValue sets LongMarketValue field to given value.

### HasLongMarketValue

`func (o *InlineResponse200) HasLongMarketValue() bool`

HasLongMarketValue returns a boolean if a field has been set.

### GetShortMarketValue

`func (o *InlineResponse200) GetShortMarketValue() string`

GetShortMarketValue returns the ShortMarketValue field if non-nil, zero value otherwise.

### GetShortMarketValueOk

`func (o *InlineResponse200) GetShortMarketValueOk() (*string, bool)`

GetShortMarketValueOk returns a tuple with the ShortMarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortMarketValue

`func (o *InlineResponse200) SetShortMarketValue(v string)`

SetShortMarketValue sets ShortMarketValue field to given value.

### HasShortMarketValue

`func (o *InlineResponse200) HasShortMarketValue() bool`

HasShortMarketValue returns a boolean if a field has been set.

### GetInitialMargin

`func (o *InlineResponse200) GetInitialMargin() string`

GetInitialMargin returns the InitialMargin field if non-nil, zero value otherwise.

### GetInitialMarginOk

`func (o *InlineResponse200) GetInitialMarginOk() (*string, bool)`

GetInitialMarginOk returns a tuple with the InitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialMargin

`func (o *InlineResponse200) SetInitialMargin(v string)`

SetInitialMargin sets InitialMargin field to given value.

### HasInitialMargin

`func (o *InlineResponse200) HasInitialMargin() bool`

HasInitialMargin returns a boolean if a field has been set.

### GetMaintenanceMargin

`func (o *InlineResponse200) GetMaintenanceMargin() string`

GetMaintenanceMargin returns the MaintenanceMargin field if non-nil, zero value otherwise.

### GetMaintenanceMarginOk

`func (o *InlineResponse200) GetMaintenanceMarginOk() (*string, bool)`

GetMaintenanceMarginOk returns a tuple with the MaintenanceMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMargin

`func (o *InlineResponse200) SetMaintenanceMargin(v string)`

SetMaintenanceMargin sets MaintenanceMargin field to given value.

### HasMaintenanceMargin

`func (o *InlineResponse200) HasMaintenanceMargin() bool`

HasMaintenanceMargin returns a boolean if a field has been set.

### GetLastMaintenanceMargin

`func (o *InlineResponse200) GetLastMaintenanceMargin() string`

GetLastMaintenanceMargin returns the LastMaintenanceMargin field if non-nil, zero value otherwise.

### GetLastMaintenanceMarginOk

`func (o *InlineResponse200) GetLastMaintenanceMarginOk() (*string, bool)`

GetLastMaintenanceMarginOk returns a tuple with the LastMaintenanceMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMaintenanceMargin

`func (o *InlineResponse200) SetLastMaintenanceMargin(v string)`

SetLastMaintenanceMargin sets LastMaintenanceMargin field to given value.

### HasLastMaintenanceMargin

`func (o *InlineResponse200) HasLastMaintenanceMargin() bool`

HasLastMaintenanceMargin returns a boolean if a field has been set.

### GetSma

`func (o *InlineResponse200) GetSma() string`

GetSma returns the Sma field if non-nil, zero value otherwise.

### GetSmaOk

`func (o *InlineResponse200) GetSmaOk() (*string, bool)`

GetSmaOk returns a tuple with the Sma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSma

`func (o *InlineResponse200) SetSma(v string)`

SetSma sets Sma field to given value.

### HasSma

`func (o *InlineResponse200) HasSma() bool`

HasSma returns a boolean if a field has been set.

### GetDaytradeCount

`func (o *InlineResponse200) GetDaytradeCount() int32`

GetDaytradeCount returns the DaytradeCount field if non-nil, zero value otherwise.

### GetDaytradeCountOk

`func (o *InlineResponse200) GetDaytradeCountOk() (*int32, bool)`

GetDaytradeCountOk returns a tuple with the DaytradeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaytradeCount

`func (o *InlineResponse200) SetDaytradeCount(v int32)`

SetDaytradeCount sets DaytradeCount field to given value.

### HasDaytradeCount

`func (o *InlineResponse200) HasDaytradeCount() bool`

HasDaytradeCount returns a boolean if a field has been set.

### GetPreviousClose

`func (o *InlineResponse200) GetPreviousClose() string`

GetPreviousClose returns the PreviousClose field if non-nil, zero value otherwise.

### GetPreviousCloseOk

`func (o *InlineResponse200) GetPreviousCloseOk() (*string, bool)`

GetPreviousCloseOk returns a tuple with the PreviousClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousClose

`func (o *InlineResponse200) SetPreviousClose(v string)`

SetPreviousClose sets PreviousClose field to given value.

### HasPreviousClose

`func (o *InlineResponse200) HasPreviousClose() bool`

HasPreviousClose returns a boolean if a field has been set.

### GetLastLongMarketValue

`func (o *InlineResponse200) GetLastLongMarketValue() string`

GetLastLongMarketValue returns the LastLongMarketValue field if non-nil, zero value otherwise.

### GetLastLongMarketValueOk

`func (o *InlineResponse200) GetLastLongMarketValueOk() (*string, bool)`

GetLastLongMarketValueOk returns a tuple with the LastLongMarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLongMarketValue

`func (o *InlineResponse200) SetLastLongMarketValue(v string)`

SetLastLongMarketValue sets LastLongMarketValue field to given value.

### HasLastLongMarketValue

`func (o *InlineResponse200) HasLastLongMarketValue() bool`

HasLastLongMarketValue returns a boolean if a field has been set.

### GetLastShortMarketValue

`func (o *InlineResponse200) GetLastShortMarketValue() string`

GetLastShortMarketValue returns the LastShortMarketValue field if non-nil, zero value otherwise.

### GetLastShortMarketValueOk

`func (o *InlineResponse200) GetLastShortMarketValueOk() (*string, bool)`

GetLastShortMarketValueOk returns a tuple with the LastShortMarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastShortMarketValue

`func (o *InlineResponse200) SetLastShortMarketValue(v string)`

SetLastShortMarketValue sets LastShortMarketValue field to given value.

### HasLastShortMarketValue

`func (o *InlineResponse200) HasLastShortMarketValue() bool`

HasLastShortMarketValue returns a boolean if a field has been set.

### GetLastCash

`func (o *InlineResponse200) GetLastCash() string`

GetLastCash returns the LastCash field if non-nil, zero value otherwise.

### GetLastCashOk

`func (o *InlineResponse200) GetLastCashOk() (*string, bool)`

GetLastCashOk returns a tuple with the LastCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCash

`func (o *InlineResponse200) SetLastCash(v string)`

SetLastCash sets LastCash field to given value.

### HasLastCash

`func (o *InlineResponse200) HasLastCash() bool`

HasLastCash returns a boolean if a field has been set.

### GetLastInitialMargin

`func (o *InlineResponse200) GetLastInitialMargin() string`

GetLastInitialMargin returns the LastInitialMargin field if non-nil, zero value otherwise.

### GetLastInitialMarginOk

`func (o *InlineResponse200) GetLastInitialMarginOk() (*string, bool)`

GetLastInitialMarginOk returns a tuple with the LastInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInitialMargin

`func (o *InlineResponse200) SetLastInitialMargin(v string)`

SetLastInitialMargin sets LastInitialMargin field to given value.

### HasLastInitialMargin

`func (o *InlineResponse200) HasLastInitialMargin() bool`

HasLastInitialMargin returns a boolean if a field has been set.

### GetLastRegtBuyingPower

`func (o *InlineResponse200) GetLastRegtBuyingPower() string`

GetLastRegtBuyingPower returns the LastRegtBuyingPower field if non-nil, zero value otherwise.

### GetLastRegtBuyingPowerOk

`func (o *InlineResponse200) GetLastRegtBuyingPowerOk() (*string, bool)`

GetLastRegtBuyingPowerOk returns a tuple with the LastRegtBuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRegtBuyingPower

`func (o *InlineResponse200) SetLastRegtBuyingPower(v string)`

SetLastRegtBuyingPower sets LastRegtBuyingPower field to given value.

### HasLastRegtBuyingPower

`func (o *InlineResponse200) HasLastRegtBuyingPower() bool`

HasLastRegtBuyingPower returns a boolean if a field has been set.

### GetLastDaytradingBuyingPower

`func (o *InlineResponse200) GetLastDaytradingBuyingPower() string`

GetLastDaytradingBuyingPower returns the LastDaytradingBuyingPower field if non-nil, zero value otherwise.

### GetLastDaytradingBuyingPowerOk

`func (o *InlineResponse200) GetLastDaytradingBuyingPowerOk() (*string, bool)`

GetLastDaytradingBuyingPowerOk returns a tuple with the LastDaytradingBuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDaytradingBuyingPower

`func (o *InlineResponse200) SetLastDaytradingBuyingPower(v string)`

SetLastDaytradingBuyingPower sets LastDaytradingBuyingPower field to given value.

### HasLastDaytradingBuyingPower

`func (o *InlineResponse200) HasLastDaytradingBuyingPower() bool`

HasLastDaytradingBuyingPower returns a boolean if a field has been set.

### GetLastBuyingPower

`func (o *InlineResponse200) GetLastBuyingPower() string`

GetLastBuyingPower returns the LastBuyingPower field if non-nil, zero value otherwise.

### GetLastBuyingPowerOk

`func (o *InlineResponse200) GetLastBuyingPowerOk() (*string, bool)`

GetLastBuyingPowerOk returns a tuple with the LastBuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBuyingPower

`func (o *InlineResponse200) SetLastBuyingPower(v string)`

SetLastBuyingPower sets LastBuyingPower field to given value.

### HasLastBuyingPower

`func (o *InlineResponse200) HasLastBuyingPower() bool`

HasLastBuyingPower returns a boolean if a field has been set.

### GetLastDaytradeCount

`func (o *InlineResponse200) GetLastDaytradeCount() int32`

GetLastDaytradeCount returns the LastDaytradeCount field if non-nil, zero value otherwise.

### GetLastDaytradeCountOk

`func (o *InlineResponse200) GetLastDaytradeCountOk() (*int32, bool)`

GetLastDaytradeCountOk returns a tuple with the LastDaytradeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDaytradeCount

`func (o *InlineResponse200) SetLastDaytradeCount(v int32)`

SetLastDaytradeCount sets LastDaytradeCount field to given value.

### HasLastDaytradeCount

`func (o *InlineResponse200) HasLastDaytradeCount() bool`

HasLastDaytradeCount returns a boolean if a field has been set.

### GetClearingBroker

`func (o *InlineResponse200) GetClearingBroker() string`

GetClearingBroker returns the ClearingBroker field if non-nil, zero value otherwise.

### GetClearingBrokerOk

`func (o *InlineResponse200) GetClearingBrokerOk() (*string, bool)`

GetClearingBrokerOk returns a tuple with the ClearingBroker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearingBroker

`func (o *InlineResponse200) SetClearingBroker(v string)`

SetClearingBroker sets ClearingBroker field to given value.

### HasClearingBroker

`func (o *InlineResponse200) HasClearingBroker() bool`

HasClearingBroker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


