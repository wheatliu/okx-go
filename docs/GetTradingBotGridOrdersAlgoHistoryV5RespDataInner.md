# GetTradingBotGridOrdersAlgoHistoryV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualLever** | Pointer to **string** | Actual Leverage  Only applicable to &#x60;contract grid&#x60; | [optional] [default to ""]
**AlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID | [optional] [default to ""]
**AlgoId** | Pointer to **string** | Algo ID | [optional] [default to ""]
**AlgoOrdType** | Pointer to **string** | Algo order type  &#x60;grid&#x60;: Spot grid  &#x60;contract_grid&#x60;: Contract grid | [optional] [default to ""]
**ArbitrageNum** | Pointer to **string** | The number of arbitrages executed | [optional] [default to ""]
**AvailEq** | Pointer to **string** | Available margin  Only applicable to &#x60;contract grid&#x60; | [optional] [default to ""]
**BasePos** | Pointer to **bool** | Whether or not to open a position when the strategy is activated  Only applicable to &#x60;contract grid&#x60; | [optional] 
**BaseSz** | Pointer to **string** | Base currency investment amount  Only applicable to &#x60;Spot grid&#x60; | [optional] [default to ""]
**CTime** | Pointer to **string** | Algo order created time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**CancelType** | Pointer to **string** | Algo order stop reason  &#x60;0&#x60;: None  &#x60;1&#x60;: Manual stop  &#x60;2&#x60;: Take profit  &#x60;3&#x60;: Stop loss  &#x60;4&#x60;: Risk control  &#x60;5&#x60;: Delivery  &#x60;6&#x60;: Signal | [optional] [default to ""]
**CopyType** | Pointer to **string** | Profit sharing order type  &#x60;0&#x60;: Normal order  &#x60;1&#x60;: Copy order without profit sharing  &#x60;2&#x60;: Copy order with profit sharing  &#x60;3&#x60;: Lead order | [optional] [default to ""]
**Direction** | Pointer to **string** | Contract grid type  &#x60;long&#x60;,&#x60;short&#x60;,&#x60;neutral&#x60;  Only applicable to &#x60;contract grid&#x60; | [optional] [default to ""]
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
**PnlRatio** | Pointer to **string** | P&amp;L ratio | [optional] [default to ""]
**ProfitSharingRatio** | Pointer to **string** | Profit sharing ratio  Value range [0, 0.3]  If it is a normal order (neither copy order nor lead order), this field returns \&quot;\&quot; | [optional] [default to ""]
**QuoteSz** | Pointer to **string** | Quote currency investment amount  Only applicable to &#x60;Spot grid&#x60; | [optional] [default to ""]
**RebateTrans** | Pointer to [**[]GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerRebateTransInner**](GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerRebateTransInner.md) | Rebate transfer info | [optional] 
**RunType** | Pointer to **string** | Grid type  &#x60;1&#x60;: Arithmetic, &#x60;2&#x60;: Geometric | [optional] [default to ""]
**SlTriggerPx** | Pointer to **string** | Stop-loss trigger price | [optional] [default to ""]
**State** | Pointer to **string** | Algo order state  &#x60;stopped&#x60; | [optional] [default to ""]
**StopResult** | Pointer to **string** | Stop result  &#x60;0&#x60;: default, &#x60;1&#x60;: Successful selling of currency at market price, &#x60;-1&#x60;: Failed to sell currency at market price  Only applicable to &#x60;Spot grid&#x60; | [optional] [default to ""]
**StopType** | Pointer to **string** | Actual Stop type  Spot grid &#x60;1&#x60;: Sell base currency &#x60;2&#x60;: Keep base currency  Contract grid &#x60;1&#x60;: Market Close All positions &#x60;2&#x60;: Keep positions | [optional] [default to ""]
**Sz** | Pointer to **string** | Used margin based on &#x60;USDT&#x60;  Only applicable to &#x60;contract grid&#x60; | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**TotalPnl** | Pointer to **string** | Total P&amp;L | [optional] [default to ""]
**TpTriggerPx** | Pointer to **string** | Take-profit trigger price | [optional] [default to ""]
**TriggerParams** | Pointer to [**[]GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner**](GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner.md) | Trigger Parameters | [optional] 
**UTime** | Pointer to **string** | Algo order updated time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Uly** | Pointer to **string** | Underlying  Only applicable to &#x60;contract grid&#x60; | [optional] [default to ""]

## Methods

### NewGetTradingBotGridOrdersAlgoHistoryV5RespDataInner

`func NewGetTradingBotGridOrdersAlgoHistoryV5RespDataInner() *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner`

NewGetTradingBotGridOrdersAlgoHistoryV5RespDataInner instantiates a new GetTradingBotGridOrdersAlgoHistoryV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotGridOrdersAlgoHistoryV5RespDataInnerWithDefaults

`func NewGetTradingBotGridOrdersAlgoHistoryV5RespDataInnerWithDefaults() *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner`

NewGetTradingBotGridOrdersAlgoHistoryV5RespDataInnerWithDefaults instantiates a new GetTradingBotGridOrdersAlgoHistoryV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualLever

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetActualLever() string`

GetActualLever returns the ActualLever field if non-nil, zero value otherwise.

### GetActualLeverOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetActualLeverOk() (*string, bool)`

GetActualLeverOk returns a tuple with the ActualLever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualLever

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetActualLever(v string)`

SetActualLever sets ActualLever field to given value.

### HasActualLever

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasActualLever() bool`

HasActualLever returns a boolean if a field has been set.

### GetAlgoClOrdId

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAlgoId

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetAlgoOrdType

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetAlgoOrdType() string`

GetAlgoOrdType returns the AlgoOrdType field if non-nil, zero value otherwise.

### GetAlgoOrdTypeOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetAlgoOrdTypeOk() (*string, bool)`

GetAlgoOrdTypeOk returns a tuple with the AlgoOrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoOrdType

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetAlgoOrdType(v string)`

SetAlgoOrdType sets AlgoOrdType field to given value.

### HasAlgoOrdType

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasAlgoOrdType() bool`

HasAlgoOrdType returns a boolean if a field has been set.

### GetArbitrageNum

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetArbitrageNum() string`

GetArbitrageNum returns the ArbitrageNum field if non-nil, zero value otherwise.

### GetArbitrageNumOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetArbitrageNumOk() (*string, bool)`

GetArbitrageNumOk returns a tuple with the ArbitrageNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArbitrageNum

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetArbitrageNum(v string)`

SetArbitrageNum sets ArbitrageNum field to given value.

### HasArbitrageNum

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasArbitrageNum() bool`

HasArbitrageNum returns a boolean if a field has been set.

### GetAvailEq

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetAvailEq() string`

GetAvailEq returns the AvailEq field if non-nil, zero value otherwise.

### GetAvailEqOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetAvailEqOk() (*string, bool)`

GetAvailEqOk returns a tuple with the AvailEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailEq

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetAvailEq(v string)`

SetAvailEq sets AvailEq field to given value.

### HasAvailEq

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasAvailEq() bool`

HasAvailEq returns a boolean if a field has been set.

### GetBasePos

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetBasePos() bool`

GetBasePos returns the BasePos field if non-nil, zero value otherwise.

### GetBasePosOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetBasePosOk() (*bool, bool)`

GetBasePosOk returns a tuple with the BasePos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePos

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetBasePos(v bool)`

SetBasePos sets BasePos field to given value.

### HasBasePos

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasBasePos() bool`

HasBasePos returns a boolean if a field has been set.

### GetBaseSz

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetBaseSz() string`

GetBaseSz returns the BaseSz field if non-nil, zero value otherwise.

### GetBaseSzOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetBaseSzOk() (*string, bool)`

GetBaseSzOk returns a tuple with the BaseSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSz

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetBaseSz(v string)`

SetBaseSz sets BaseSz field to given value.

### HasBaseSz

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasBaseSz() bool`

HasBaseSz returns a boolean if a field has been set.

### GetCTime

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCancelType

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetCancelType() string`

GetCancelType returns the CancelType field if non-nil, zero value otherwise.

### GetCancelTypeOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetCancelTypeOk() (*string, bool)`

GetCancelTypeOk returns a tuple with the CancelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelType

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetCancelType(v string)`

SetCancelType sets CancelType field to given value.

### HasCancelType

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasCancelType() bool`

HasCancelType returns a boolean if a field has been set.

### GetCopyType

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetCopyType() string`

GetCopyType returns the CopyType field if non-nil, zero value otherwise.

### GetCopyTypeOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetCopyTypeOk() (*string, bool)`

GetCopyTypeOk returns a tuple with the CopyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyType

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetCopyType(v string)`

SetCopyType sets CopyType field to given value.

### HasCopyType

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasCopyType() bool`

HasCopyType returns a boolean if a field has been set.

### GetDirection

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetFee

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFloatProfit

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetFloatProfit() string`

GetFloatProfit returns the FloatProfit field if non-nil, zero value otherwise.

### GetFloatProfitOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetFloatProfitOk() (*string, bool)`

GetFloatProfitOk returns a tuple with the FloatProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatProfit

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetFloatProfit(v string)`

SetFloatProfit sets FloatProfit field to given value.

### HasFloatProfit

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasFloatProfit() bool`

HasFloatProfit returns a boolean if a field has been set.

### GetFundingFee

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetFundingFee() string`

GetFundingFee returns the FundingFee field if non-nil, zero value otherwise.

### GetFundingFeeOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetFundingFeeOk() (*string, bool)`

GetFundingFeeOk returns a tuple with the FundingFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingFee

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetFundingFee(v string)`

SetFundingFee sets FundingFee field to given value.

### HasFundingFee

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasFundingFee() bool`

HasFundingFee returns a boolean if a field has been set.

### GetGridNum

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetGridNum() string`

GetGridNum returns the GridNum field if non-nil, zero value otherwise.

### GetGridNumOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetGridNumOk() (*string, bool)`

GetGridNumOk returns a tuple with the GridNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridNum

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetGridNum(v string)`

SetGridNum sets GridNum field to given value.

### HasGridNum

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasGridNum() bool`

HasGridNum returns a boolean if a field has been set.

### GetGridProfit

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetGridProfit() string`

GetGridProfit returns the GridProfit field if non-nil, zero value otherwise.

### GetGridProfitOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetGridProfitOk() (*string, bool)`

GetGridProfitOk returns a tuple with the GridProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridProfit

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetGridProfit(v string)`

SetGridProfit sets GridProfit field to given value.

### HasGridProfit

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasGridProfit() bool`

HasGridProfit returns a boolean if a field has been set.

### GetInstFamily

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetInstFamily() string`

GetInstFamily returns the InstFamily field if non-nil, zero value otherwise.

### GetInstFamilyOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetInstFamilyOk() (*string, bool)`

GetInstFamilyOk returns a tuple with the InstFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstFamily

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetInstFamily(v string)`

SetInstFamily sets InstFamily field to given value.

### HasInstFamily

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasInstFamily() bool`

HasInstFamily returns a boolean if a field has been set.

### GetInstId

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetInvestment

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetInvestment() string`

GetInvestment returns the Investment field if non-nil, zero value otherwise.

### GetInvestmentOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetInvestmentOk() (*string, bool)`

GetInvestmentOk returns a tuple with the Investment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestment

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetInvestment(v string)`

SetInvestment sets Investment field to given value.

### HasInvestment

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasInvestment() bool`

HasInvestment returns a boolean if a field has been set.

### GetLever

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetLiqPx

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetLiqPx() string`

GetLiqPx returns the LiqPx field if non-nil, zero value otherwise.

### GetLiqPxOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetLiqPxOk() (*string, bool)`

GetLiqPxOk returns a tuple with the LiqPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiqPx

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetLiqPx(v string)`

SetLiqPx sets LiqPx field to given value.

### HasLiqPx

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasLiqPx() bool`

HasLiqPx returns a boolean if a field has been set.

### GetMaxPx

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetMaxPx() string`

GetMaxPx returns the MaxPx field if non-nil, zero value otherwise.

### GetMaxPxOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetMaxPxOk() (*string, bool)`

GetMaxPxOk returns a tuple with the MaxPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPx

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetMaxPx(v string)`

SetMaxPx sets MaxPx field to given value.

### HasMaxPx

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasMaxPx() bool`

HasMaxPx returns a boolean if a field has been set.

### GetMinPx

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetMinPx() string`

GetMinPx returns the MinPx field if non-nil, zero value otherwise.

### GetMinPxOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetMinPxOk() (*string, bool)`

GetMinPxOk returns a tuple with the MinPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPx

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetMinPx(v string)`

SetMinPx sets MinPx field to given value.

### HasMinPx

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasMinPx() bool`

HasMinPx returns a boolean if a field has been set.

### GetOrdFrozen

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetOrdFrozen() string`

GetOrdFrozen returns the OrdFrozen field if non-nil, zero value otherwise.

### GetOrdFrozenOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetOrdFrozenOk() (*string, bool)`

GetOrdFrozenOk returns a tuple with the OrdFrozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdFrozen

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetOrdFrozen(v string)`

SetOrdFrozen sets OrdFrozen field to given value.

### HasOrdFrozen

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasOrdFrozen() bool`

HasOrdFrozen returns a boolean if a field has been set.

### GetPnlRatio

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetPnlRatio() string`

GetPnlRatio returns the PnlRatio field if non-nil, zero value otherwise.

### GetPnlRatioOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetPnlRatioOk() (*string, bool)`

GetPnlRatioOk returns a tuple with the PnlRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnlRatio

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetPnlRatio(v string)`

SetPnlRatio sets PnlRatio field to given value.

### HasPnlRatio

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasPnlRatio() bool`

HasPnlRatio returns a boolean if a field has been set.

### GetProfitSharingRatio

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetProfitSharingRatio() string`

GetProfitSharingRatio returns the ProfitSharingRatio field if non-nil, zero value otherwise.

### GetProfitSharingRatioOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetProfitSharingRatioOk() (*string, bool)`

GetProfitSharingRatioOk returns a tuple with the ProfitSharingRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitSharingRatio

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetProfitSharingRatio(v string)`

SetProfitSharingRatio sets ProfitSharingRatio field to given value.

### HasProfitSharingRatio

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasProfitSharingRatio() bool`

HasProfitSharingRatio returns a boolean if a field has been set.

### GetQuoteSz

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetQuoteSz() string`

GetQuoteSz returns the QuoteSz field if non-nil, zero value otherwise.

### GetQuoteSzOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetQuoteSzOk() (*string, bool)`

GetQuoteSzOk returns a tuple with the QuoteSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteSz

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetQuoteSz(v string)`

SetQuoteSz sets QuoteSz field to given value.

### HasQuoteSz

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasQuoteSz() bool`

HasQuoteSz returns a boolean if a field has been set.

### GetRebateTrans

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetRebateTrans() []GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerRebateTransInner`

GetRebateTrans returns the RebateTrans field if non-nil, zero value otherwise.

### GetRebateTransOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetRebateTransOk() (*[]GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerRebateTransInner, bool)`

GetRebateTransOk returns a tuple with the RebateTrans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebateTrans

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetRebateTrans(v []GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerRebateTransInner)`

SetRebateTrans sets RebateTrans field to given value.

### HasRebateTrans

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasRebateTrans() bool`

HasRebateTrans returns a boolean if a field has been set.

### GetRunType

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### GetSlTriggerPx

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetSlTriggerPx() string`

GetSlTriggerPx returns the SlTriggerPx field if non-nil, zero value otherwise.

### GetSlTriggerPxOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetSlTriggerPxOk() (*string, bool)`

GetSlTriggerPxOk returns a tuple with the SlTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPx

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetSlTriggerPx(v string)`

SetSlTriggerPx sets SlTriggerPx field to given value.

### HasSlTriggerPx

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasSlTriggerPx() bool`

HasSlTriggerPx returns a boolean if a field has been set.

### GetState

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStopResult

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetStopResult() string`

GetStopResult returns the StopResult field if non-nil, zero value otherwise.

### GetStopResultOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetStopResultOk() (*string, bool)`

GetStopResultOk returns a tuple with the StopResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopResult

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetStopResult(v string)`

SetStopResult sets StopResult field to given value.

### HasStopResult

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasStopResult() bool`

HasStopResult returns a boolean if a field has been set.

### GetStopType

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetStopType() string`

GetStopType returns the StopType field if non-nil, zero value otherwise.

### GetStopTypeOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetStopTypeOk() (*string, bool)`

GetStopTypeOk returns a tuple with the StopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopType

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetStopType(v string)`

SetStopType sets StopType field to given value.

### HasStopType

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasStopType() bool`

HasStopType returns a boolean if a field has been set.

### GetSz

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTag

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTotalPnl

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetTotalPnl() string`

GetTotalPnl returns the TotalPnl field if non-nil, zero value otherwise.

### GetTotalPnlOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetTotalPnlOk() (*string, bool)`

GetTotalPnlOk returns a tuple with the TotalPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPnl

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetTotalPnl(v string)`

SetTotalPnl sets TotalPnl field to given value.

### HasTotalPnl

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasTotalPnl() bool`

HasTotalPnl returns a boolean if a field has been set.

### GetTpTriggerPx

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetTpTriggerPx() string`

GetTpTriggerPx returns the TpTriggerPx field if non-nil, zero value otherwise.

### GetTpTriggerPxOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetTpTriggerPxOk() (*string, bool)`

GetTpTriggerPxOk returns a tuple with the TpTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPx

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetTpTriggerPx(v string)`

SetTpTriggerPx sets TpTriggerPx field to given value.

### HasTpTriggerPx

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasTpTriggerPx() bool`

HasTpTriggerPx returns a boolean if a field has been set.

### GetTriggerParams

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetTriggerParams() []GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner`

GetTriggerParams returns the TriggerParams field if non-nil, zero value otherwise.

### GetTriggerParamsOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetTriggerParamsOk() (*[]GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner, bool)`

GetTriggerParamsOk returns a tuple with the TriggerParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerParams

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetTriggerParams(v []GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner)`

SetTriggerParams sets TriggerParams field to given value.

### HasTriggerParams

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasTriggerParams() bool`

HasTriggerParams returns a boolean if a field has been set.

### GetUTime

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.

### GetUly

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetUly() string`

GetUly returns the Uly field if non-nil, zero value otherwise.

### GetUlyOk

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) GetUlyOk() (*string, bool)`

GetUlyOk returns a tuple with the Uly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUly

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) SetUly(v string)`

SetUly sets Uly field to given value.

### HasUly

`func (o *GetTradingBotGridOrdersAlgoHistoryV5RespDataInner) HasUly() bool`

HasUly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


