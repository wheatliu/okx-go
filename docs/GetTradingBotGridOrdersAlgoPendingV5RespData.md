# GetTradingBotGridOrdersAlgoPendingV5RespData

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
**RebateTrans** | Pointer to [**[]GetTradingBotGridOrdersAlgoDetailsV5RespDataRebateTransInner**](GetTradingBotGridOrdersAlgoDetailsV5RespDataRebateTransInner.md) | Rebate transfer info | [optional] 
**RunType** | Pointer to **string** | Grid type  &#x60;1&#x60;: Arithmetic, &#x60;2&#x60;: Geometric | [optional] [default to ""]
**SlTriggerPx** | Pointer to **string** | Stop-loss trigger price | [optional] [default to ""]
**State** | Pointer to **string** | Algo order state  &#x60;starting&#x60;  &#x60;running&#x60;  &#x60;stopping&#x60;  &#x60;pending_signal&#x60;  &#x60;no_close_position&#x60;: stopped algo order but have not closed position yet | [optional] [default to ""]
**StopType** | Pointer to **string** | Actual Stop type  Spot &#x60;1&#x60;: Sell base currency &#x60;2&#x60;: Keep base currency  Contract grid &#x60;1&#x60;: Market Close All positions &#x60;2&#x60;: Keep positions | [optional] [default to ""]
**Sz** | Pointer to **string** | Used margin based on &#x60;USDT&#x60;  Only applicable to &#x60;contract grid&#x60; | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**TotalPnl** | Pointer to **string** | Total P&amp;L | [optional] [default to ""]
**TpTriggerPx** | Pointer to **string** | Take-profit trigger price | [optional] [default to ""]
**TriggerParams** | Pointer to [**[]GetTradingBotGridOrdersAlgoDetailsV5RespDataTriggerParamsInner**](GetTradingBotGridOrdersAlgoDetailsV5RespDataTriggerParamsInner.md) | Trigger Parameters | [optional] 
**UTime** | Pointer to **string** | Algo order updated time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Uly** | Pointer to **string** | Underlying  Only applicable to &#x60;contract grid&#x60; | [optional] [default to ""]

## Methods

### NewGetTradingBotGridOrdersAlgoPendingV5RespData

`func NewGetTradingBotGridOrdersAlgoPendingV5RespData() *GetTradingBotGridOrdersAlgoPendingV5RespData`

NewGetTradingBotGridOrdersAlgoPendingV5RespData instantiates a new GetTradingBotGridOrdersAlgoPendingV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotGridOrdersAlgoPendingV5RespDataWithDefaults

`func NewGetTradingBotGridOrdersAlgoPendingV5RespDataWithDefaults() *GetTradingBotGridOrdersAlgoPendingV5RespData`

NewGetTradingBotGridOrdersAlgoPendingV5RespDataWithDefaults instantiates a new GetTradingBotGridOrdersAlgoPendingV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualLever

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetActualLever() string`

GetActualLever returns the ActualLever field if non-nil, zero value otherwise.

### GetActualLeverOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetActualLeverOk() (*string, bool)`

GetActualLeverOk returns a tuple with the ActualLever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualLever

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetActualLever(v string)`

SetActualLever sets ActualLever field to given value.

### HasActualLever

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasActualLever() bool`

HasActualLever returns a boolean if a field has been set.

### GetAlgoClOrdId

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAlgoId

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetAlgoOrdType

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetAlgoOrdType() string`

GetAlgoOrdType returns the AlgoOrdType field if non-nil, zero value otherwise.

### GetAlgoOrdTypeOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetAlgoOrdTypeOk() (*string, bool)`

GetAlgoOrdTypeOk returns a tuple with the AlgoOrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoOrdType

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetAlgoOrdType(v string)`

SetAlgoOrdType sets AlgoOrdType field to given value.

### HasAlgoOrdType

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasAlgoOrdType() bool`

HasAlgoOrdType returns a boolean if a field has been set.

### GetArbitrageNum

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetArbitrageNum() string`

GetArbitrageNum returns the ArbitrageNum field if non-nil, zero value otherwise.

### GetArbitrageNumOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetArbitrageNumOk() (*string, bool)`

GetArbitrageNumOk returns a tuple with the ArbitrageNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArbitrageNum

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetArbitrageNum(v string)`

SetArbitrageNum sets ArbitrageNum field to given value.

### HasArbitrageNum

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasArbitrageNum() bool`

HasArbitrageNum returns a boolean if a field has been set.

### GetAvailEq

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetAvailEq() string`

GetAvailEq returns the AvailEq field if non-nil, zero value otherwise.

### GetAvailEqOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetAvailEqOk() (*string, bool)`

GetAvailEqOk returns a tuple with the AvailEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailEq

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetAvailEq(v string)`

SetAvailEq sets AvailEq field to given value.

### HasAvailEq

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasAvailEq() bool`

HasAvailEq returns a boolean if a field has been set.

### GetBasePos

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetBasePos() bool`

GetBasePos returns the BasePos field if non-nil, zero value otherwise.

### GetBasePosOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetBasePosOk() (*bool, bool)`

GetBasePosOk returns a tuple with the BasePos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePos

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetBasePos(v bool)`

SetBasePos sets BasePos field to given value.

### HasBasePos

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasBasePos() bool`

HasBasePos returns a boolean if a field has been set.

### GetBaseSz

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetBaseSz() string`

GetBaseSz returns the BaseSz field if non-nil, zero value otherwise.

### GetBaseSzOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetBaseSzOk() (*string, bool)`

GetBaseSzOk returns a tuple with the BaseSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSz

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetBaseSz(v string)`

SetBaseSz sets BaseSz field to given value.

### HasBaseSz

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasBaseSz() bool`

HasBaseSz returns a boolean if a field has been set.

### GetCTime

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCancelType

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetCancelType() string`

GetCancelType returns the CancelType field if non-nil, zero value otherwise.

### GetCancelTypeOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetCancelTypeOk() (*string, bool)`

GetCancelTypeOk returns a tuple with the CancelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelType

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetCancelType(v string)`

SetCancelType sets CancelType field to given value.

### HasCancelType

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasCancelType() bool`

HasCancelType returns a boolean if a field has been set.

### GetCopyType

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetCopyType() string`

GetCopyType returns the CopyType field if non-nil, zero value otherwise.

### GetCopyTypeOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetCopyTypeOk() (*string, bool)`

GetCopyTypeOk returns a tuple with the CopyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyType

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetCopyType(v string)`

SetCopyType sets CopyType field to given value.

### HasCopyType

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasCopyType() bool`

HasCopyType returns a boolean if a field has been set.

### GetDirection

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetFee

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFloatProfit

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetFloatProfit() string`

GetFloatProfit returns the FloatProfit field if non-nil, zero value otherwise.

### GetFloatProfitOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetFloatProfitOk() (*string, bool)`

GetFloatProfitOk returns a tuple with the FloatProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatProfit

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetFloatProfit(v string)`

SetFloatProfit sets FloatProfit field to given value.

### HasFloatProfit

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasFloatProfit() bool`

HasFloatProfit returns a boolean if a field has been set.

### GetFundingFee

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetFundingFee() string`

GetFundingFee returns the FundingFee field if non-nil, zero value otherwise.

### GetFundingFeeOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetFundingFeeOk() (*string, bool)`

GetFundingFeeOk returns a tuple with the FundingFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingFee

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetFundingFee(v string)`

SetFundingFee sets FundingFee field to given value.

### HasFundingFee

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasFundingFee() bool`

HasFundingFee returns a boolean if a field has been set.

### GetGridNum

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetGridNum() string`

GetGridNum returns the GridNum field if non-nil, zero value otherwise.

### GetGridNumOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetGridNumOk() (*string, bool)`

GetGridNumOk returns a tuple with the GridNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridNum

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetGridNum(v string)`

SetGridNum sets GridNum field to given value.

### HasGridNum

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasGridNum() bool`

HasGridNum returns a boolean if a field has been set.

### GetGridProfit

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetGridProfit() string`

GetGridProfit returns the GridProfit field if non-nil, zero value otherwise.

### GetGridProfitOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetGridProfitOk() (*string, bool)`

GetGridProfitOk returns a tuple with the GridProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridProfit

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetGridProfit(v string)`

SetGridProfit sets GridProfit field to given value.

### HasGridProfit

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasGridProfit() bool`

HasGridProfit returns a boolean if a field has been set.

### GetInstFamily

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetInstFamily() string`

GetInstFamily returns the InstFamily field if non-nil, zero value otherwise.

### GetInstFamilyOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetInstFamilyOk() (*string, bool)`

GetInstFamilyOk returns a tuple with the InstFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstFamily

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetInstFamily(v string)`

SetInstFamily sets InstFamily field to given value.

### HasInstFamily

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasInstFamily() bool`

HasInstFamily returns a boolean if a field has been set.

### GetInstId

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetInvestment

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetInvestment() string`

GetInvestment returns the Investment field if non-nil, zero value otherwise.

### GetInvestmentOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetInvestmentOk() (*string, bool)`

GetInvestmentOk returns a tuple with the Investment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestment

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetInvestment(v string)`

SetInvestment sets Investment field to given value.

### HasInvestment

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasInvestment() bool`

HasInvestment returns a boolean if a field has been set.

### GetLever

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetLiqPx

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetLiqPx() string`

GetLiqPx returns the LiqPx field if non-nil, zero value otherwise.

### GetLiqPxOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetLiqPxOk() (*string, bool)`

GetLiqPxOk returns a tuple with the LiqPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiqPx

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetLiqPx(v string)`

SetLiqPx sets LiqPx field to given value.

### HasLiqPx

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasLiqPx() bool`

HasLiqPx returns a boolean if a field has been set.

### GetMaxPx

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetMaxPx() string`

GetMaxPx returns the MaxPx field if non-nil, zero value otherwise.

### GetMaxPxOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetMaxPxOk() (*string, bool)`

GetMaxPxOk returns a tuple with the MaxPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPx

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetMaxPx(v string)`

SetMaxPx sets MaxPx field to given value.

### HasMaxPx

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasMaxPx() bool`

HasMaxPx returns a boolean if a field has been set.

### GetMinPx

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetMinPx() string`

GetMinPx returns the MinPx field if non-nil, zero value otherwise.

### GetMinPxOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetMinPxOk() (*string, bool)`

GetMinPxOk returns a tuple with the MinPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPx

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetMinPx(v string)`

SetMinPx sets MinPx field to given value.

### HasMinPx

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasMinPx() bool`

HasMinPx returns a boolean if a field has been set.

### GetOrdFrozen

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetOrdFrozen() string`

GetOrdFrozen returns the OrdFrozen field if non-nil, zero value otherwise.

### GetOrdFrozenOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetOrdFrozenOk() (*string, bool)`

GetOrdFrozenOk returns a tuple with the OrdFrozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdFrozen

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetOrdFrozen(v string)`

SetOrdFrozen sets OrdFrozen field to given value.

### HasOrdFrozen

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasOrdFrozen() bool`

HasOrdFrozen returns a boolean if a field has been set.

### GetPnlRatio

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetPnlRatio() string`

GetPnlRatio returns the PnlRatio field if non-nil, zero value otherwise.

### GetPnlRatioOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetPnlRatioOk() (*string, bool)`

GetPnlRatioOk returns a tuple with the PnlRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnlRatio

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetPnlRatio(v string)`

SetPnlRatio sets PnlRatio field to given value.

### HasPnlRatio

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasPnlRatio() bool`

HasPnlRatio returns a boolean if a field has been set.

### GetProfitSharingRatio

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetProfitSharingRatio() string`

GetProfitSharingRatio returns the ProfitSharingRatio field if non-nil, zero value otherwise.

### GetProfitSharingRatioOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetProfitSharingRatioOk() (*string, bool)`

GetProfitSharingRatioOk returns a tuple with the ProfitSharingRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitSharingRatio

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetProfitSharingRatio(v string)`

SetProfitSharingRatio sets ProfitSharingRatio field to given value.

### HasProfitSharingRatio

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasProfitSharingRatio() bool`

HasProfitSharingRatio returns a boolean if a field has been set.

### GetQuoteSz

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetQuoteSz() string`

GetQuoteSz returns the QuoteSz field if non-nil, zero value otherwise.

### GetQuoteSzOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetQuoteSzOk() (*string, bool)`

GetQuoteSzOk returns a tuple with the QuoteSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteSz

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetQuoteSz(v string)`

SetQuoteSz sets QuoteSz field to given value.

### HasQuoteSz

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasQuoteSz() bool`

HasQuoteSz returns a boolean if a field has been set.

### GetRebateTrans

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetRebateTrans() []GetTradingBotGridOrdersAlgoDetailsV5RespDataRebateTransInner`

GetRebateTrans returns the RebateTrans field if non-nil, zero value otherwise.

### GetRebateTransOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetRebateTransOk() (*[]GetTradingBotGridOrdersAlgoDetailsV5RespDataRebateTransInner, bool)`

GetRebateTransOk returns a tuple with the RebateTrans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebateTrans

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetRebateTrans(v []GetTradingBotGridOrdersAlgoDetailsV5RespDataRebateTransInner)`

SetRebateTrans sets RebateTrans field to given value.

### HasRebateTrans

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasRebateTrans() bool`

HasRebateTrans returns a boolean if a field has been set.

### GetRunType

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### GetSlTriggerPx

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetSlTriggerPx() string`

GetSlTriggerPx returns the SlTriggerPx field if non-nil, zero value otherwise.

### GetSlTriggerPxOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetSlTriggerPxOk() (*string, bool)`

GetSlTriggerPxOk returns a tuple with the SlTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPx

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetSlTriggerPx(v string)`

SetSlTriggerPx sets SlTriggerPx field to given value.

### HasSlTriggerPx

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasSlTriggerPx() bool`

HasSlTriggerPx returns a boolean if a field has been set.

### GetState

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStopType

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetStopType() string`

GetStopType returns the StopType field if non-nil, zero value otherwise.

### GetStopTypeOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetStopTypeOk() (*string, bool)`

GetStopTypeOk returns a tuple with the StopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopType

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetStopType(v string)`

SetStopType sets StopType field to given value.

### HasStopType

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasStopType() bool`

HasStopType returns a boolean if a field has been set.

### GetSz

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTag

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTotalPnl

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetTotalPnl() string`

GetTotalPnl returns the TotalPnl field if non-nil, zero value otherwise.

### GetTotalPnlOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetTotalPnlOk() (*string, bool)`

GetTotalPnlOk returns a tuple with the TotalPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPnl

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetTotalPnl(v string)`

SetTotalPnl sets TotalPnl field to given value.

### HasTotalPnl

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasTotalPnl() bool`

HasTotalPnl returns a boolean if a field has been set.

### GetTpTriggerPx

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetTpTriggerPx() string`

GetTpTriggerPx returns the TpTriggerPx field if non-nil, zero value otherwise.

### GetTpTriggerPxOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetTpTriggerPxOk() (*string, bool)`

GetTpTriggerPxOk returns a tuple with the TpTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPx

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetTpTriggerPx(v string)`

SetTpTriggerPx sets TpTriggerPx field to given value.

### HasTpTriggerPx

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasTpTriggerPx() bool`

HasTpTriggerPx returns a boolean if a field has been set.

### GetTriggerParams

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetTriggerParams() []GetTradingBotGridOrdersAlgoDetailsV5RespDataTriggerParamsInner`

GetTriggerParams returns the TriggerParams field if non-nil, zero value otherwise.

### GetTriggerParamsOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetTriggerParamsOk() (*[]GetTradingBotGridOrdersAlgoDetailsV5RespDataTriggerParamsInner, bool)`

GetTriggerParamsOk returns a tuple with the TriggerParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerParams

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetTriggerParams(v []GetTradingBotGridOrdersAlgoDetailsV5RespDataTriggerParamsInner)`

SetTriggerParams sets TriggerParams field to given value.

### HasTriggerParams

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasTriggerParams() bool`

HasTriggerParams returns a boolean if a field has been set.

### GetUTime

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasUTime() bool`

HasUTime returns a boolean if a field has been set.

### GetUly

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetUly() string`

GetUly returns the Uly field if non-nil, zero value otherwise.

### GetUlyOk

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) GetUlyOk() (*string, bool)`

GetUlyOk returns a tuple with the Uly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUly

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) SetUly(v string)`

SetUly sets Uly field to given value.

### HasUly

`func (o *GetTradingBotGridOrdersAlgoPendingV5RespData) HasUly() bool`

HasUly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


