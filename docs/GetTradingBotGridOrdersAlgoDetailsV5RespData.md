# GetTradingBotGridOrdersAlgoDetailsV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveOrdNum** | Pointer to **string** | Total count of pending sub orders | [optional] [default to ""]
**ActualLever** | Pointer to **string** | Actual Leverage  Only applicable to &#x60;contract grid&#x60; | [optional] [default to ""]
**AlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID | [optional] [default to ""]
**AlgoId** | Pointer to **string** | Algo ID | [optional] [default to ""]
**AlgoOrdType** | Pointer to **string** | Algo order type  &#x60;grid&#x60;: Spot grid  &#x60;contract_grid&#x60;: Contract grid | [optional] [default to ""]
**AnnualizedRate** | Pointer to **string** | Grid annualized rate | [optional] [default to ""]
**ArbitrageNum** | Pointer to **string** | The number of arbitrages executed | [optional] [default to ""]
**AvailEq** | Pointer to **string** | Available margin  Only applicable to &#x60;contract grid&#x60; | [optional] [default to ""]
**BasePos** | Pointer to **bool** | Whether or not to open a position when the strategy is activated  Only applicable to &#x60;contract grid&#x60; | [optional] 
**BaseSz** | Pointer to **string** | Base currency investment amount  Only applicable to &#x60;Spot grid&#x60; | [optional] [default to ""]
**CTime** | Pointer to **string** | Algo order created time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**CancelType** | Pointer to **string** | Algo order stop reason  &#x60;0&#x60;: None  &#x60;1&#x60;: Manual stop  &#x60;2&#x60;: Take profit  &#x60;3&#x60;: Stop loss  &#x60;4&#x60;: Risk control  &#x60;5&#x60;: Delivery  &#x60;6&#x60;: Signal | [optional] [default to ""]
**CopyType** | Pointer to **string** | Profit sharing order type  &#x60;0&#x60;: Normal order  &#x60;1&#x60;: Copy order without profit sharing  &#x60;2&#x60;: Copy order with profit sharing  &#x60;3&#x60;: Lead order | [optional] [default to ""]
**CurBaseSz** | Pointer to **string** | Assets of base currency currently held  Only applicable to &#x60;Spot grid&#x60; | [optional] [default to ""]
**CurQuoteSz** | Pointer to **string** | Assets of quote currency currently held  Only applicable to &#x60;Spot grid&#x60; | [optional] [default to ""]
**Direction** | Pointer to **string** | Contract grid type  &#x60;long&#x60;,&#x60;short&#x60;,&#x60;neutral&#x60;  Only applicable to &#x60;contract grid&#x60; | [optional] [default to ""]
**Eq** | Pointer to **string** | Total equity of strategy account  Only applicable to &#x60;contract grid&#x60; | [optional] [default to ""]
**Fee** | Pointer to **string** | Accumulated fee. Only applicable to contract grid, or it will be \&quot;\&quot; | [optional] [default to ""]
**FloatProfit** | Pointer to **string** | Variable P&amp;L | [optional] [default to ""]
**FundingFee** | Pointer to **string** | Accumulated funding fee. Only applicable to contract grid, or it will be \&quot;\&quot; | [optional] [default to ""]
**GridNum** | Pointer to **string** | Grid quantity | [optional] [default to ""]
**GridProfit** | Pointer to **string** | Grid profit | [optional] [default to ""]
**InstFamily** | Pointer to **string** | Instrument family  Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;  Only applicable to &#x60;contract grid&#x60; | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**Investment** | Pointer to **string** | Accumulated investment amount  Spot grid investment amount calculated on quote currency | [optional] [default to ""]
**Lever** | Pointer to **string** | Leverage  Only applicable to &#x60;contract grid&#x60; | [optional] [default to ""]
**LiqPx** | Pointer to **string** | Estimated liquidation price  Only applicable to &#x60;contract grid&#x60; | [optional] [default to ""]
**MaxPx** | Pointer to **string** | Upper price of price range | [optional] [default to ""]
**MinPx** | Pointer to **string** | Lower price of price range | [optional] [default to ""]
**OrdFrozen** | Pointer to **string** | Margin used by pending orders  Only applicable to &#x60;contract grid&#x60; | [optional] [default to ""]
**PerMaxProfitRate** | Pointer to **string** | Estimated maximum Profit margin per grid | [optional] [default to ""]
**PerMinProfitRate** | Pointer to **string** | Estimated minimum Profit margin per grid | [optional] [default to ""]
**PnlRatio** | Pointer to **string** | P&amp;L ratio | [optional] [default to ""]
**Profit** | Pointer to **string** | Current available profit based on quote currency  Only applicable to &#x60;Spot grid&#x60; | [optional] [default to ""]
**ProfitSharingRatio** | Pointer to **string** | Profit sharing ratio  Value range [0, 0.3]  If it is a normal order (neither copy order nor lead order), this field returns \&quot;\&quot; | [optional] [default to ""]
**QuoteSz** | Pointer to **string** | Quote currency investment amount  Only applicable to &#x60;Spot grid&#x60; | [optional] [default to ""]
**RebateTrans** | Pointer to [**[]GetTradingBotGridOrdersAlgoDetailsV5RespDataRebateTransInner**](GetTradingBotGridOrdersAlgoDetailsV5RespDataRebateTransInner.md) | Rebate transfer info | [optional] 
**RunPx** | Pointer to **string** | Price at launch | [optional] [default to ""]
**RunType** | Pointer to **string** | Grid type  &#x60;1&#x60;: Arithmetic, &#x60;2&#x60;: Geometric | [optional] [default to ""]
**SingleAmt** | Pointer to **string** | Amount per grid | [optional] [default to ""]
**SlRatio** | Pointer to **string** | Stop loss ratio, 0.1 represents 10% | [optional] [default to ""]
**SlTriggerPx** | Pointer to **string** | Stop-loss trigger price | [optional] [default to ""]
**State** | Pointer to **string** | Algo order state  &#x60;starting&#x60;  &#x60;running&#x60;  &#x60;stopping&#x60;  &#x60;no_close_position&#x60;: stopped algo order but have not closed position yet  &#x60;stopped&#x60; | [optional] [default to ""]
**StopResult** | Pointer to **string** | Stop result  &#x60;0&#x60;: default, &#x60;1&#x60;: Successful selling of currency at market price, &#x60;-1&#x60;: Failed to sell currency at market price  Only applicable to &#x60;Spot grid&#x60; | [optional] [default to ""]
**StopType** | Pointer to **string** | Stop type  Spot grid &#x60;1&#x60;: Sell base currency &#x60;2&#x60;: Keep base currency  Contract grid &#x60;1&#x60;: Market Close All positions &#x60;2&#x60;: Keep positions | [optional] [default to ""]
**Sz** | Pointer to **string** | Used margin based on &#x60;USDT&#x60;  Only applicable to &#x60;contract grid&#x60; | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**TotalAnnualizedRate** | Pointer to **string** | Total annualized rate | [optional] [default to ""]
**TotalPnl** | Pointer to **string** | Total P&amp;L | [optional] [default to ""]
**TpRatio** | Pointer to **string** | Take profit ratio, 0.1 represents 10% | [optional] [default to ""]
**TpTriggerPx** | Pointer to **string** | Take-profit trigger price | [optional] [default to ""]
**TradeNum** | Pointer to **string** | The number of trades executed | [optional] [default to ""]
**TriggerParams** | Pointer to [**[]GetTradingBotGridOrdersAlgoDetailsV5RespDataTriggerParamsInner**](GetTradingBotGridOrdersAlgoDetailsV5RespDataTriggerParamsInner.md) | Trigger Parameters | [optional] 
**UTime** | Pointer to **string** | Algo order updated time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Uly** | Pointer to **string** | Underlying  Only applicable to &#x60;contract grid&#x60; | [optional] [default to ""]

## Methods

### NewGetTradingBotGridOrdersAlgoDetailsV5RespData

`func NewGetTradingBotGridOrdersAlgoDetailsV5RespData() *GetTradingBotGridOrdersAlgoDetailsV5RespData`

NewGetTradingBotGridOrdersAlgoDetailsV5RespData instantiates a new GetTradingBotGridOrdersAlgoDetailsV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotGridOrdersAlgoDetailsV5RespDataWithDefaults

`func NewGetTradingBotGridOrdersAlgoDetailsV5RespDataWithDefaults() *GetTradingBotGridOrdersAlgoDetailsV5RespData`

NewGetTradingBotGridOrdersAlgoDetailsV5RespDataWithDefaults instantiates a new GetTradingBotGridOrdersAlgoDetailsV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveOrdNum

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetActiveOrdNum() string`

GetActiveOrdNum returns the ActiveOrdNum field if non-nil, zero value otherwise.

### GetActiveOrdNumOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetActiveOrdNumOk() (*string, bool)`

GetActiveOrdNumOk returns a tuple with the ActiveOrdNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveOrdNum

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetActiveOrdNum(v string)`

SetActiveOrdNum sets ActiveOrdNum field to given value.

### HasActiveOrdNum

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasActiveOrdNum() bool`

HasActiveOrdNum returns a boolean if a field has been set.

### GetActualLever

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetActualLever() string`

GetActualLever returns the ActualLever field if non-nil, zero value otherwise.

### GetActualLeverOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetActualLeverOk() (*string, bool)`

GetActualLeverOk returns a tuple with the ActualLever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualLever

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetActualLever(v string)`

SetActualLever sets ActualLever field to given value.

### HasActualLever

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasActualLever() bool`

HasActualLever returns a boolean if a field has been set.

### GetAlgoClOrdId

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAlgoId

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetAlgoOrdType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetAlgoOrdType() string`

GetAlgoOrdType returns the AlgoOrdType field if non-nil, zero value otherwise.

### GetAlgoOrdTypeOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetAlgoOrdTypeOk() (*string, bool)`

GetAlgoOrdTypeOk returns a tuple with the AlgoOrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoOrdType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetAlgoOrdType(v string)`

SetAlgoOrdType sets AlgoOrdType field to given value.

### HasAlgoOrdType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasAlgoOrdType() bool`

HasAlgoOrdType returns a boolean if a field has been set.

### GetAnnualizedRate

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetAnnualizedRate() string`

GetAnnualizedRate returns the AnnualizedRate field if non-nil, zero value otherwise.

### GetAnnualizedRateOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetAnnualizedRateOk() (*string, bool)`

GetAnnualizedRateOk returns a tuple with the AnnualizedRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualizedRate

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetAnnualizedRate(v string)`

SetAnnualizedRate sets AnnualizedRate field to given value.

### HasAnnualizedRate

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasAnnualizedRate() bool`

HasAnnualizedRate returns a boolean if a field has been set.

### GetArbitrageNum

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetArbitrageNum() string`

GetArbitrageNum returns the ArbitrageNum field if non-nil, zero value otherwise.

### GetArbitrageNumOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetArbitrageNumOk() (*string, bool)`

GetArbitrageNumOk returns a tuple with the ArbitrageNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArbitrageNum

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetArbitrageNum(v string)`

SetArbitrageNum sets ArbitrageNum field to given value.

### HasArbitrageNum

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasArbitrageNum() bool`

HasArbitrageNum returns a boolean if a field has been set.

### GetAvailEq

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetAvailEq() string`

GetAvailEq returns the AvailEq field if non-nil, zero value otherwise.

### GetAvailEqOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetAvailEqOk() (*string, bool)`

GetAvailEqOk returns a tuple with the AvailEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailEq

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetAvailEq(v string)`

SetAvailEq sets AvailEq field to given value.

### HasAvailEq

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasAvailEq() bool`

HasAvailEq returns a boolean if a field has been set.

### GetBasePos

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetBasePos() bool`

GetBasePos returns the BasePos field if non-nil, zero value otherwise.

### GetBasePosOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetBasePosOk() (*bool, bool)`

GetBasePosOk returns a tuple with the BasePos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePos

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetBasePos(v bool)`

SetBasePos sets BasePos field to given value.

### HasBasePos

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasBasePos() bool`

HasBasePos returns a boolean if a field has been set.

### GetBaseSz

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetBaseSz() string`

GetBaseSz returns the BaseSz field if non-nil, zero value otherwise.

### GetBaseSzOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetBaseSzOk() (*string, bool)`

GetBaseSzOk returns a tuple with the BaseSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSz

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetBaseSz(v string)`

SetBaseSz sets BaseSz field to given value.

### HasBaseSz

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasBaseSz() bool`

HasBaseSz returns a boolean if a field has been set.

### GetCTime

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCancelType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetCancelType() string`

GetCancelType returns the CancelType field if non-nil, zero value otherwise.

### GetCancelTypeOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetCancelTypeOk() (*string, bool)`

GetCancelTypeOk returns a tuple with the CancelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetCancelType(v string)`

SetCancelType sets CancelType field to given value.

### HasCancelType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasCancelType() bool`

HasCancelType returns a boolean if a field has been set.

### GetCopyType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetCopyType() string`

GetCopyType returns the CopyType field if non-nil, zero value otherwise.

### GetCopyTypeOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetCopyTypeOk() (*string, bool)`

GetCopyTypeOk returns a tuple with the CopyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetCopyType(v string)`

SetCopyType sets CopyType field to given value.

### HasCopyType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasCopyType() bool`

HasCopyType returns a boolean if a field has been set.

### GetCurBaseSz

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetCurBaseSz() string`

GetCurBaseSz returns the CurBaseSz field if non-nil, zero value otherwise.

### GetCurBaseSzOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetCurBaseSzOk() (*string, bool)`

GetCurBaseSzOk returns a tuple with the CurBaseSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurBaseSz

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetCurBaseSz(v string)`

SetCurBaseSz sets CurBaseSz field to given value.

### HasCurBaseSz

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasCurBaseSz() bool`

HasCurBaseSz returns a boolean if a field has been set.

### GetCurQuoteSz

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetCurQuoteSz() string`

GetCurQuoteSz returns the CurQuoteSz field if non-nil, zero value otherwise.

### GetCurQuoteSzOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetCurQuoteSzOk() (*string, bool)`

GetCurQuoteSzOk returns a tuple with the CurQuoteSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurQuoteSz

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetCurQuoteSz(v string)`

SetCurQuoteSz sets CurQuoteSz field to given value.

### HasCurQuoteSz

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasCurQuoteSz() bool`

HasCurQuoteSz returns a boolean if a field has been set.

### GetDirection

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetEq

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetEq() string`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetEqOk() (*string, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetEq(v string)`

SetEq sets Eq field to given value.

### HasEq

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasEq() bool`

HasEq returns a boolean if a field has been set.

### GetFee

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFloatProfit

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetFloatProfit() string`

GetFloatProfit returns the FloatProfit field if non-nil, zero value otherwise.

### GetFloatProfitOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetFloatProfitOk() (*string, bool)`

GetFloatProfitOk returns a tuple with the FloatProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatProfit

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetFloatProfit(v string)`

SetFloatProfit sets FloatProfit field to given value.

### HasFloatProfit

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasFloatProfit() bool`

HasFloatProfit returns a boolean if a field has been set.

### GetFundingFee

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetFundingFee() string`

GetFundingFee returns the FundingFee field if non-nil, zero value otherwise.

### GetFundingFeeOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetFundingFeeOk() (*string, bool)`

GetFundingFeeOk returns a tuple with the FundingFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingFee

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetFundingFee(v string)`

SetFundingFee sets FundingFee field to given value.

### HasFundingFee

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasFundingFee() bool`

HasFundingFee returns a boolean if a field has been set.

### GetGridNum

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetGridNum() string`

GetGridNum returns the GridNum field if non-nil, zero value otherwise.

### GetGridNumOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetGridNumOk() (*string, bool)`

GetGridNumOk returns a tuple with the GridNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridNum

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetGridNum(v string)`

SetGridNum sets GridNum field to given value.

### HasGridNum

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasGridNum() bool`

HasGridNum returns a boolean if a field has been set.

### GetGridProfit

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetGridProfit() string`

GetGridProfit returns the GridProfit field if non-nil, zero value otherwise.

### GetGridProfitOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetGridProfitOk() (*string, bool)`

GetGridProfitOk returns a tuple with the GridProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridProfit

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetGridProfit(v string)`

SetGridProfit sets GridProfit field to given value.

### HasGridProfit

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasGridProfit() bool`

HasGridProfit returns a boolean if a field has been set.

### GetInstFamily

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetInstFamily() string`

GetInstFamily returns the InstFamily field if non-nil, zero value otherwise.

### GetInstFamilyOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetInstFamilyOk() (*string, bool)`

GetInstFamilyOk returns a tuple with the InstFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstFamily

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetInstFamily(v string)`

SetInstFamily sets InstFamily field to given value.

### HasInstFamily

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasInstFamily() bool`

HasInstFamily returns a boolean if a field has been set.

### GetInstId

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetInvestment

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetInvestment() string`

GetInvestment returns the Investment field if non-nil, zero value otherwise.

### GetInvestmentOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetInvestmentOk() (*string, bool)`

GetInvestmentOk returns a tuple with the Investment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestment

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetInvestment(v string)`

SetInvestment sets Investment field to given value.

### HasInvestment

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasInvestment() bool`

HasInvestment returns a boolean if a field has been set.

### GetLever

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetLiqPx

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetLiqPx() string`

GetLiqPx returns the LiqPx field if non-nil, zero value otherwise.

### GetLiqPxOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetLiqPxOk() (*string, bool)`

GetLiqPxOk returns a tuple with the LiqPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiqPx

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetLiqPx(v string)`

SetLiqPx sets LiqPx field to given value.

### HasLiqPx

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasLiqPx() bool`

HasLiqPx returns a boolean if a field has been set.

### GetMaxPx

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetMaxPx() string`

GetMaxPx returns the MaxPx field if non-nil, zero value otherwise.

### GetMaxPxOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetMaxPxOk() (*string, bool)`

GetMaxPxOk returns a tuple with the MaxPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPx

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetMaxPx(v string)`

SetMaxPx sets MaxPx field to given value.

### HasMaxPx

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasMaxPx() bool`

HasMaxPx returns a boolean if a field has been set.

### GetMinPx

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetMinPx() string`

GetMinPx returns the MinPx field if non-nil, zero value otherwise.

### GetMinPxOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetMinPxOk() (*string, bool)`

GetMinPxOk returns a tuple with the MinPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPx

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetMinPx(v string)`

SetMinPx sets MinPx field to given value.

### HasMinPx

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasMinPx() bool`

HasMinPx returns a boolean if a field has been set.

### GetOrdFrozen

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetOrdFrozen() string`

GetOrdFrozen returns the OrdFrozen field if non-nil, zero value otherwise.

### GetOrdFrozenOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetOrdFrozenOk() (*string, bool)`

GetOrdFrozenOk returns a tuple with the OrdFrozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdFrozen

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetOrdFrozen(v string)`

SetOrdFrozen sets OrdFrozen field to given value.

### HasOrdFrozen

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasOrdFrozen() bool`

HasOrdFrozen returns a boolean if a field has been set.

### GetPerMaxProfitRate

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetPerMaxProfitRate() string`

GetPerMaxProfitRate returns the PerMaxProfitRate field if non-nil, zero value otherwise.

### GetPerMaxProfitRateOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetPerMaxProfitRateOk() (*string, bool)`

GetPerMaxProfitRateOk returns a tuple with the PerMaxProfitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerMaxProfitRate

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetPerMaxProfitRate(v string)`

SetPerMaxProfitRate sets PerMaxProfitRate field to given value.

### HasPerMaxProfitRate

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasPerMaxProfitRate() bool`

HasPerMaxProfitRate returns a boolean if a field has been set.

### GetPerMinProfitRate

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetPerMinProfitRate() string`

GetPerMinProfitRate returns the PerMinProfitRate field if non-nil, zero value otherwise.

### GetPerMinProfitRateOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetPerMinProfitRateOk() (*string, bool)`

GetPerMinProfitRateOk returns a tuple with the PerMinProfitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerMinProfitRate

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetPerMinProfitRate(v string)`

SetPerMinProfitRate sets PerMinProfitRate field to given value.

### HasPerMinProfitRate

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasPerMinProfitRate() bool`

HasPerMinProfitRate returns a boolean if a field has been set.

### GetPnlRatio

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetPnlRatio() string`

GetPnlRatio returns the PnlRatio field if non-nil, zero value otherwise.

### GetPnlRatioOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetPnlRatioOk() (*string, bool)`

GetPnlRatioOk returns a tuple with the PnlRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnlRatio

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetPnlRatio(v string)`

SetPnlRatio sets PnlRatio field to given value.

### HasPnlRatio

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasPnlRatio() bool`

HasPnlRatio returns a boolean if a field has been set.

### GetProfit

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetProfit() string`

GetProfit returns the Profit field if non-nil, zero value otherwise.

### GetProfitOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetProfitOk() (*string, bool)`

GetProfitOk returns a tuple with the Profit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfit

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetProfit(v string)`

SetProfit sets Profit field to given value.

### HasProfit

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasProfit() bool`

HasProfit returns a boolean if a field has been set.

### GetProfitSharingRatio

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetProfitSharingRatio() string`

GetProfitSharingRatio returns the ProfitSharingRatio field if non-nil, zero value otherwise.

### GetProfitSharingRatioOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetProfitSharingRatioOk() (*string, bool)`

GetProfitSharingRatioOk returns a tuple with the ProfitSharingRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitSharingRatio

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetProfitSharingRatio(v string)`

SetProfitSharingRatio sets ProfitSharingRatio field to given value.

### HasProfitSharingRatio

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasProfitSharingRatio() bool`

HasProfitSharingRatio returns a boolean if a field has been set.

### GetQuoteSz

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetQuoteSz() string`

GetQuoteSz returns the QuoteSz field if non-nil, zero value otherwise.

### GetQuoteSzOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetQuoteSzOk() (*string, bool)`

GetQuoteSzOk returns a tuple with the QuoteSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteSz

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetQuoteSz(v string)`

SetQuoteSz sets QuoteSz field to given value.

### HasQuoteSz

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasQuoteSz() bool`

HasQuoteSz returns a boolean if a field has been set.

### GetRebateTrans

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetRebateTrans() []GetTradingBotGridOrdersAlgoDetailsV5RespDataRebateTransInner`

GetRebateTrans returns the RebateTrans field if non-nil, zero value otherwise.

### GetRebateTransOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetRebateTransOk() (*[]GetTradingBotGridOrdersAlgoDetailsV5RespDataRebateTransInner, bool)`

GetRebateTransOk returns a tuple with the RebateTrans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebateTrans

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetRebateTrans(v []GetTradingBotGridOrdersAlgoDetailsV5RespDataRebateTransInner)`

SetRebateTrans sets RebateTrans field to given value.

### HasRebateTrans

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasRebateTrans() bool`

HasRebateTrans returns a boolean if a field has been set.

### GetRunPx

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetRunPx() string`

GetRunPx returns the RunPx field if non-nil, zero value otherwise.

### GetRunPxOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetRunPxOk() (*string, bool)`

GetRunPxOk returns a tuple with the RunPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunPx

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetRunPx(v string)`

SetRunPx sets RunPx field to given value.

### HasRunPx

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasRunPx() bool`

HasRunPx returns a boolean if a field has been set.

### GetRunType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### GetSingleAmt

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetSingleAmt() string`

GetSingleAmt returns the SingleAmt field if non-nil, zero value otherwise.

### GetSingleAmtOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetSingleAmtOk() (*string, bool)`

GetSingleAmtOk returns a tuple with the SingleAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleAmt

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetSingleAmt(v string)`

SetSingleAmt sets SingleAmt field to given value.

### HasSingleAmt

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasSingleAmt() bool`

HasSingleAmt returns a boolean if a field has been set.

### GetSlRatio

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetSlRatio() string`

GetSlRatio returns the SlRatio field if non-nil, zero value otherwise.

### GetSlRatioOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetSlRatioOk() (*string, bool)`

GetSlRatioOk returns a tuple with the SlRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlRatio

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetSlRatio(v string)`

SetSlRatio sets SlRatio field to given value.

### HasSlRatio

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasSlRatio() bool`

HasSlRatio returns a boolean if a field has been set.

### GetSlTriggerPx

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetSlTriggerPx() string`

GetSlTriggerPx returns the SlTriggerPx field if non-nil, zero value otherwise.

### GetSlTriggerPxOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetSlTriggerPxOk() (*string, bool)`

GetSlTriggerPxOk returns a tuple with the SlTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPx

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetSlTriggerPx(v string)`

SetSlTriggerPx sets SlTriggerPx field to given value.

### HasSlTriggerPx

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasSlTriggerPx() bool`

HasSlTriggerPx returns a boolean if a field has been set.

### GetState

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStopResult

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetStopResult() string`

GetStopResult returns the StopResult field if non-nil, zero value otherwise.

### GetStopResultOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetStopResultOk() (*string, bool)`

GetStopResultOk returns a tuple with the StopResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopResult

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetStopResult(v string)`

SetStopResult sets StopResult field to given value.

### HasStopResult

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasStopResult() bool`

HasStopResult returns a boolean if a field has been set.

### GetStopType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetStopType() string`

GetStopType returns the StopType field if non-nil, zero value otherwise.

### GetStopTypeOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetStopTypeOk() (*string, bool)`

GetStopTypeOk returns a tuple with the StopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetStopType(v string)`

SetStopType sets StopType field to given value.

### HasStopType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasStopType() bool`

HasStopType returns a boolean if a field has been set.

### GetSz

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTag

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTotalAnnualizedRate

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetTotalAnnualizedRate() string`

GetTotalAnnualizedRate returns the TotalAnnualizedRate field if non-nil, zero value otherwise.

### GetTotalAnnualizedRateOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetTotalAnnualizedRateOk() (*string, bool)`

GetTotalAnnualizedRateOk returns a tuple with the TotalAnnualizedRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAnnualizedRate

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetTotalAnnualizedRate(v string)`

SetTotalAnnualizedRate sets TotalAnnualizedRate field to given value.

### HasTotalAnnualizedRate

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasTotalAnnualizedRate() bool`

HasTotalAnnualizedRate returns a boolean if a field has been set.

### GetTotalPnl

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetTotalPnl() string`

GetTotalPnl returns the TotalPnl field if non-nil, zero value otherwise.

### GetTotalPnlOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetTotalPnlOk() (*string, bool)`

GetTotalPnlOk returns a tuple with the TotalPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPnl

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetTotalPnl(v string)`

SetTotalPnl sets TotalPnl field to given value.

### HasTotalPnl

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasTotalPnl() bool`

HasTotalPnl returns a boolean if a field has been set.

### GetTpRatio

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetTpRatio() string`

GetTpRatio returns the TpRatio field if non-nil, zero value otherwise.

### GetTpRatioOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetTpRatioOk() (*string, bool)`

GetTpRatioOk returns a tuple with the TpRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpRatio

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetTpRatio(v string)`

SetTpRatio sets TpRatio field to given value.

### HasTpRatio

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasTpRatio() bool`

HasTpRatio returns a boolean if a field has been set.

### GetTpTriggerPx

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetTpTriggerPx() string`

GetTpTriggerPx returns the TpTriggerPx field if non-nil, zero value otherwise.

### GetTpTriggerPxOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetTpTriggerPxOk() (*string, bool)`

GetTpTriggerPxOk returns a tuple with the TpTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPx

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetTpTriggerPx(v string)`

SetTpTriggerPx sets TpTriggerPx field to given value.

### HasTpTriggerPx

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasTpTriggerPx() bool`

HasTpTriggerPx returns a boolean if a field has been set.

### GetTradeNum

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetTradeNum() string`

GetTradeNum returns the TradeNum field if non-nil, zero value otherwise.

### GetTradeNumOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetTradeNumOk() (*string, bool)`

GetTradeNumOk returns a tuple with the TradeNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeNum

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetTradeNum(v string)`

SetTradeNum sets TradeNum field to given value.

### HasTradeNum

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasTradeNum() bool`

HasTradeNum returns a boolean if a field has been set.

### GetTriggerParams

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetTriggerParams() []GetTradingBotGridOrdersAlgoDetailsV5RespDataTriggerParamsInner`

GetTriggerParams returns the TriggerParams field if non-nil, zero value otherwise.

### GetTriggerParamsOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetTriggerParamsOk() (*[]GetTradingBotGridOrdersAlgoDetailsV5RespDataTriggerParamsInner, bool)`

GetTriggerParamsOk returns a tuple with the TriggerParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerParams

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetTriggerParams(v []GetTradingBotGridOrdersAlgoDetailsV5RespDataTriggerParamsInner)`

SetTriggerParams sets TriggerParams field to given value.

### HasTriggerParams

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasTriggerParams() bool`

HasTriggerParams returns a boolean if a field has been set.

### GetUTime

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasUTime() bool`

HasUTime returns a boolean if a field has been set.

### GetUly

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetUly() string`

GetUly returns the Uly field if non-nil, zero value otherwise.

### GetUlyOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) GetUlyOk() (*string, bool)`

GetUlyOk returns a tuple with the Uly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUly

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) SetUly(v string)`

SetUly sets Uly field to given value.

### HasUly

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespData) HasUly() bool`

HasUly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


