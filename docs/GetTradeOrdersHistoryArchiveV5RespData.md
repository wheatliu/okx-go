# GetTradeOrdersHistoryArchiveV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccFillSz** | Pointer to **string** | Accumulated fill quantity | [optional] [default to ""]
**AlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID. There will be a value when algo order attaching &#x60;algoClOrdId&#x60; is triggered, or it will be \&quot;\&quot;. | [optional] [default to ""]
**AlgoId** | Pointer to **string** | Algo ID. There will be a value when algo order is triggered, or it will be \&quot;\&quot;. | [optional] [default to ""]
**AttachAlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID when placing order attaching TP/SL. | [optional] [default to ""]
**AttachAlgoOrds** | Pointer to [**[]GetTradeOrderV5RespDataAttachAlgoOrdsInner**](GetTradeOrderV5RespDataAttachAlgoOrdsInner.md) | TP/SL information attached when placing order | [optional] 
**AvgPx** | Pointer to **string** | Average filled price. If none is filled, it will return \&quot;\&quot;. | [optional] [default to ""]
**CTime** | Pointer to **string** | Creation time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**CancelSource** | Pointer to **string** | Code of the cancellation source. | [optional] [default to ""]
**CancelSourceReason** | Pointer to **string** | Reason for the cancellation. | [optional] [default to ""]
**Category** | Pointer to **string** | Category   &#x60;normal&#x60;  &#x60;twap&#x60;   &#x60;adl&#x60;  &#x60;full_liquidation&#x60;  &#x60;partial_liquidation&#x60;   &#x60;delivery&#x60;    &#x60;ddh&#x60;: Delta dynamic hedge | [optional] [default to ""]
**Ccy** | Pointer to **string** | Margin currency   Applicable to all &#x60;isolated&#x60; &#x60;MARGIN&#x60; orders and &#x60;cross&#x60; &#x60;MARGIN&#x60; orders in &#x60;Spot and futures mode&#x60;. | [optional] [default to ""]
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client | [optional] [default to ""]
**Fee** | Pointer to **string** | Fee and rebate  For spot and margin, it is accumulated fee charged by the platform. It is always negative, e.g. -0.01.   For Expiry Futures, Perpetual Futures and Options, it is accumulated fee and rebate | [optional] [default to ""]
**FeeCcy** | Pointer to **string** | Fee currency | [optional] [default to ""]
**FillPx** | Pointer to **string** | Last filled price. If none is filled, it will return \&quot;\&quot;. | [optional] [default to ""]
**FillSz** | Pointer to **string** | Last filled quantity | [optional] [default to ""]
**FillTime** | Pointer to **string** | Last filled time | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**IsTpLimit** | Pointer to **string** | Whether it is TP limit order. true or false | [optional] [default to ""]
**Lever** | Pointer to **string** | Leverage, from &#x60;0.01&#x60; to &#x60;125&#x60;.   Only applicable to &#x60;MARGIN/FUTURES/SWAP&#x60; | [optional] [default to ""]
**LinkedAlgoOrd** | Pointer to [**GetTradeOrderV5RespDataLinkedAlgoOrd**](GetTradeOrderV5RespDataLinkedAlgoOrd.md) |  | [optional] 
**OrdId** | Pointer to **string** | Order ID | [optional] [default to ""]
**OrdType** | Pointer to **string** | Order type   &#x60;market&#x60;: Market order   &#x60;limit&#x60;: Limit order   &#x60;post_only&#x60;: Post-only order   &#x60;fok&#x60;: Fill-or-kill order   &#x60;ioc&#x60;: Immediate-or-cancel order    &#x60;optimal_limit_ioc&#x60;: Market order with immediate-or-cancel order  &#x60;mmp&#x60;: Market Maker Protection (only applicable to Option in Portfolio Margin mode)  &#x60;mmp_and_post_only&#x60;: Market Maker Protection and Post-only order(only applicable to Option in Portfolio Margin mode)   &#x60;op_fok&#x60;: Simple options (fok) | [optional] [default to ""]
**Pnl** | Pointer to **string** | Profit and loss, Applicable to orders which have a trade and aim to close position. It always is 0 in other conditions | [optional] [default to ""]
**PosSide** | Pointer to **string** | Position side | [optional] [default to ""]
**Px** | Pointer to **string** | Price   For options, use coin as unit (e.g. BTC, ETH) | [optional] [default to ""]
**PxType** | Pointer to **string** | Price type of options   &#x60;px&#x60;: Place an order based on price, in the unit of coin (the unit for the request parameter px is BTC or ETH)   &#x60;pxVol&#x60;: Place an order based on pxVol   &#x60;pxUsd&#x60;: Place an order based on pxUsd, in the unit of USD (the unit for the request parameter px is USD) | [optional] [default to ""]
**PxUsd** | Pointer to **string** | Options price in USDOnly applicable to options; return \&quot;\&quot; for other instrument types | [optional] [default to ""]
**PxVol** | Pointer to **string** | Implied volatility of the options orderOnly applicable to options; return \&quot;\&quot; for other instrument types | [optional] [default to ""]
**QuickMgnType** | Pointer to **string** | Quick Margin type, Only applicable to Quick Margin Mode of isolated margin  &#x60;manual&#x60;, &#x60;auto_borrow&#x60;, &#x60;auto_repay&#x60;  (Deprecated) | [optional] [default to ""]
**Rebate** | Pointer to **string** | Rebate amount, only applicable to spot and margin, the reward of placing orders from the platform (rebate) given to user who has reached the specified trading level. If there is no rebate, this field is \&quot;\&quot;. | [optional] [default to ""]
**RebateCcy** | Pointer to **string** | Rebate currency | [optional] [default to ""]
**ReduceOnly** | Pointer to **string** | Whether the order can only reduce the position size. Valid options: true or false. | [optional] [default to ""]
**Side** | Pointer to **string** | Order side | [optional] [default to ""]
**SlOrdPx** | Pointer to **string** | Stop-loss order price. | [optional] [default to ""]
**SlTriggerPx** | Pointer to **string** | Stop-loss trigger price. | [optional] [default to ""]
**SlTriggerPxType** | Pointer to **string** | Stop-loss trigger price type.   &#x60;last&#x60;: last price  &#x60;index&#x60;: index price  &#x60;mark&#x60;: mark price | [optional] [default to ""]
**Source** | Pointer to **string** | Order source  &#x60;6&#x60;: The normal order triggered by the &#x60;trigger order&#x60;  &#x60;7&#x60;:The normal order triggered by the &#x60;TP/SL order&#x60;   &#x60;13&#x60;: The normal order triggered by the algo order  &#x60;25&#x60;:The normal order triggered by the &#x60;trailing stop order&#x60;  &#x60;34&#x60;: The normal order triggered by the &#x60;chase order&#x60; | [optional] [default to ""]
**State** | Pointer to **string** | State   &#x60;canceled&#x60;   &#x60;filled&#x60;   &#x60;mmp_canceled&#x60; | [optional] [default to ""]
**StpId** | Pointer to **string** | Self trade prevention ID  Return \&quot;\&quot; if self trade prevention is not applicable (deprecated) | [optional] [default to ""]
**StpMode** | Pointer to **string** | Self trade prevention mode   Return \&quot;\&quot; if self trade prevention is not applicable | [optional] [default to ""]
**Sz** | Pointer to **string** | Quantity to buy or sell | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**TdMode** | Pointer to **string** | Trade mode | [optional] [default to ""]
**TgtCcy** | Pointer to **string** | Order quantity unit setting for &#x60;sz&#x60;  &#x60;base_ccy&#x60;: Base currency ,&#x60;quote_ccy&#x60;: Quote currency    Only applicable to &#x60;SPOT&#x60; Market Orders  Default is &#x60;quote_ccy&#x60; for buy, &#x60;base_ccy&#x60; for sell | [optional] [default to ""]
**TpOrdPx** | Pointer to **string** | Take-profit order price. | [optional] [default to ""]
**TpTriggerPx** | Pointer to **string** | Take-profit trigger price. | [optional] [default to ""]
**TpTriggerPxType** | Pointer to **string** | Take-profit trigger price type.   &#x60;last&#x60;: last price  &#x60;index&#x60;: index price  &#x60;mark&#x60;: mark price | [optional] [default to ""]
**TradeId** | Pointer to **string** | Last trade ID | [optional] [default to ""]
**UTime** | Pointer to **string** | Update time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetTradeOrdersHistoryArchiveV5RespData

`func NewGetTradeOrdersHistoryArchiveV5RespData() *GetTradeOrdersHistoryArchiveV5RespData`

NewGetTradeOrdersHistoryArchiveV5RespData instantiates a new GetTradeOrdersHistoryArchiveV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradeOrdersHistoryArchiveV5RespDataWithDefaults

`func NewGetTradeOrdersHistoryArchiveV5RespDataWithDefaults() *GetTradeOrdersHistoryArchiveV5RespData`

NewGetTradeOrdersHistoryArchiveV5RespDataWithDefaults instantiates a new GetTradeOrdersHistoryArchiveV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccFillSz

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetAccFillSz() string`

GetAccFillSz returns the AccFillSz field if non-nil, zero value otherwise.

### GetAccFillSzOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetAccFillSzOk() (*string, bool)`

GetAccFillSzOk returns a tuple with the AccFillSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccFillSz

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetAccFillSz(v string)`

SetAccFillSz sets AccFillSz field to given value.

### HasAccFillSz

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasAccFillSz() bool`

HasAccFillSz returns a boolean if a field has been set.

### GetAlgoClOrdId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAlgoId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetAttachAlgoClOrdId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetAttachAlgoClOrdId() string`

GetAttachAlgoClOrdId returns the AttachAlgoClOrdId field if non-nil, zero value otherwise.

### GetAttachAlgoClOrdIdOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetAttachAlgoClOrdIdOk() (*string, bool)`

GetAttachAlgoClOrdIdOk returns a tuple with the AttachAlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachAlgoClOrdId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetAttachAlgoClOrdId(v string)`

SetAttachAlgoClOrdId sets AttachAlgoClOrdId field to given value.

### HasAttachAlgoClOrdId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasAttachAlgoClOrdId() bool`

HasAttachAlgoClOrdId returns a boolean if a field has been set.

### GetAttachAlgoOrds

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetAttachAlgoOrds() []GetTradeOrderV5RespDataAttachAlgoOrdsInner`

GetAttachAlgoOrds returns the AttachAlgoOrds field if non-nil, zero value otherwise.

### GetAttachAlgoOrdsOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetAttachAlgoOrdsOk() (*[]GetTradeOrderV5RespDataAttachAlgoOrdsInner, bool)`

GetAttachAlgoOrdsOk returns a tuple with the AttachAlgoOrds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachAlgoOrds

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetAttachAlgoOrds(v []GetTradeOrderV5RespDataAttachAlgoOrdsInner)`

SetAttachAlgoOrds sets AttachAlgoOrds field to given value.

### HasAttachAlgoOrds

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasAttachAlgoOrds() bool`

HasAttachAlgoOrds returns a boolean if a field has been set.

### GetAvgPx

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetAvgPx() string`

GetAvgPx returns the AvgPx field if non-nil, zero value otherwise.

### GetAvgPxOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetAvgPxOk() (*string, bool)`

GetAvgPxOk returns a tuple with the AvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPx

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetAvgPx(v string)`

SetAvgPx sets AvgPx field to given value.

### HasAvgPx

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasAvgPx() bool`

HasAvgPx returns a boolean if a field has been set.

### GetCTime

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCancelSource

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetCancelSource() string`

GetCancelSource returns the CancelSource field if non-nil, zero value otherwise.

### GetCancelSourceOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetCancelSourceOk() (*string, bool)`

GetCancelSourceOk returns a tuple with the CancelSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelSource

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetCancelSource(v string)`

SetCancelSource sets CancelSource field to given value.

### HasCancelSource

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasCancelSource() bool`

HasCancelSource returns a boolean if a field has been set.

### GetCancelSourceReason

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetCancelSourceReason() string`

GetCancelSourceReason returns the CancelSourceReason field if non-nil, zero value otherwise.

### GetCancelSourceReasonOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetCancelSourceReasonOk() (*string, bool)`

GetCancelSourceReasonOk returns a tuple with the CancelSourceReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelSourceReason

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetCancelSourceReason(v string)`

SetCancelSourceReason sets CancelSourceReason field to given value.

### HasCancelSourceReason

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasCancelSourceReason() bool`

HasCancelSourceReason returns a boolean if a field has been set.

### GetCategory

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCcy

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetClOrdId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetFee

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFeeCcy

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetFeeCcy() string`

GetFeeCcy returns the FeeCcy field if non-nil, zero value otherwise.

### GetFeeCcyOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetFeeCcyOk() (*string, bool)`

GetFeeCcyOk returns a tuple with the FeeCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCcy

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetFeeCcy(v string)`

SetFeeCcy sets FeeCcy field to given value.

### HasFeeCcy

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasFeeCcy() bool`

HasFeeCcy returns a boolean if a field has been set.

### GetFillPx

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetFillPx() string`

GetFillPx returns the FillPx field if non-nil, zero value otherwise.

### GetFillPxOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetFillPxOk() (*string, bool)`

GetFillPxOk returns a tuple with the FillPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPx

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetFillPx(v string)`

SetFillPx sets FillPx field to given value.

### HasFillPx

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasFillPx() bool`

HasFillPx returns a boolean if a field has been set.

### GetFillSz

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetFillSz() string`

GetFillSz returns the FillSz field if non-nil, zero value otherwise.

### GetFillSzOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetFillSzOk() (*string, bool)`

GetFillSzOk returns a tuple with the FillSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillSz

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetFillSz(v string)`

SetFillSz sets FillSz field to given value.

### HasFillSz

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasFillSz() bool`

HasFillSz returns a boolean if a field has been set.

### GetFillTime

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetFillTime() string`

GetFillTime returns the FillTime field if non-nil, zero value otherwise.

### GetFillTimeOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetFillTimeOk() (*string, bool)`

GetFillTimeOk returns a tuple with the FillTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillTime

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetFillTime(v string)`

SetFillTime sets FillTime field to given value.

### HasFillTime

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasFillTime() bool`

HasFillTime returns a boolean if a field has been set.

### GetInstId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetIsTpLimit

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetIsTpLimit() string`

GetIsTpLimit returns the IsTpLimit field if non-nil, zero value otherwise.

### GetIsTpLimitOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetIsTpLimitOk() (*string, bool)`

GetIsTpLimitOk returns a tuple with the IsTpLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTpLimit

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetIsTpLimit(v string)`

SetIsTpLimit sets IsTpLimit field to given value.

### HasIsTpLimit

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasIsTpLimit() bool`

HasIsTpLimit returns a boolean if a field has been set.

### GetLever

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetLinkedAlgoOrd

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetLinkedAlgoOrd() GetTradeOrderV5RespDataLinkedAlgoOrd`

GetLinkedAlgoOrd returns the LinkedAlgoOrd field if non-nil, zero value otherwise.

### GetLinkedAlgoOrdOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetLinkedAlgoOrdOk() (*GetTradeOrderV5RespDataLinkedAlgoOrd, bool)`

GetLinkedAlgoOrdOk returns a tuple with the LinkedAlgoOrd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAlgoOrd

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetLinkedAlgoOrd(v GetTradeOrderV5RespDataLinkedAlgoOrd)`

SetLinkedAlgoOrd sets LinkedAlgoOrd field to given value.

### HasLinkedAlgoOrd

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasLinkedAlgoOrd() bool`

HasLinkedAlgoOrd returns a boolean if a field has been set.

### GetOrdId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetOrdType

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetOrdType() string`

GetOrdType returns the OrdType field if non-nil, zero value otherwise.

### GetOrdTypeOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetOrdTypeOk() (*string, bool)`

GetOrdTypeOk returns a tuple with the OrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdType

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetOrdType(v string)`

SetOrdType sets OrdType field to given value.

### HasOrdType

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasOrdType() bool`

HasOrdType returns a boolean if a field has been set.

### GetPnl

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetPnl() string`

GetPnl returns the Pnl field if non-nil, zero value otherwise.

### GetPnlOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetPnlOk() (*string, bool)`

GetPnlOk returns a tuple with the Pnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnl

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetPnl(v string)`

SetPnl sets Pnl field to given value.

### HasPnl

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasPnl() bool`

HasPnl returns a boolean if a field has been set.

### GetPosSide

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetPx

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetPxType

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetPxType() string`

GetPxType returns the PxType field if non-nil, zero value otherwise.

### GetPxTypeOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetPxTypeOk() (*string, bool)`

GetPxTypeOk returns a tuple with the PxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxType

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetPxType(v string)`

SetPxType sets PxType field to given value.

### HasPxType

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasPxType() bool`

HasPxType returns a boolean if a field has been set.

### GetPxUsd

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetPxUsd() string`

GetPxUsd returns the PxUsd field if non-nil, zero value otherwise.

### GetPxUsdOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetPxUsdOk() (*string, bool)`

GetPxUsdOk returns a tuple with the PxUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxUsd

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetPxUsd(v string)`

SetPxUsd sets PxUsd field to given value.

### HasPxUsd

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasPxUsd() bool`

HasPxUsd returns a boolean if a field has been set.

### GetPxVol

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetPxVol() string`

GetPxVol returns the PxVol field if non-nil, zero value otherwise.

### GetPxVolOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetPxVolOk() (*string, bool)`

GetPxVolOk returns a tuple with the PxVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxVol

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetPxVol(v string)`

SetPxVol sets PxVol field to given value.

### HasPxVol

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasPxVol() bool`

HasPxVol returns a boolean if a field has been set.

### GetQuickMgnType

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetQuickMgnType() string`

GetQuickMgnType returns the QuickMgnType field if non-nil, zero value otherwise.

### GetQuickMgnTypeOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetQuickMgnTypeOk() (*string, bool)`

GetQuickMgnTypeOk returns a tuple with the QuickMgnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickMgnType

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetQuickMgnType(v string)`

SetQuickMgnType sets QuickMgnType field to given value.

### HasQuickMgnType

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasQuickMgnType() bool`

HasQuickMgnType returns a boolean if a field has been set.

### GetRebate

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetRebate() string`

GetRebate returns the Rebate field if non-nil, zero value otherwise.

### GetRebateOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetRebateOk() (*string, bool)`

GetRebateOk returns a tuple with the Rebate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebate

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetRebate(v string)`

SetRebate sets Rebate field to given value.

### HasRebate

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasRebate() bool`

HasRebate returns a boolean if a field has been set.

### GetRebateCcy

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetRebateCcy() string`

GetRebateCcy returns the RebateCcy field if non-nil, zero value otherwise.

### GetRebateCcyOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetRebateCcyOk() (*string, bool)`

GetRebateCcyOk returns a tuple with the RebateCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebateCcy

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetRebateCcy(v string)`

SetRebateCcy sets RebateCcy field to given value.

### HasRebateCcy

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasRebateCcy() bool`

HasRebateCcy returns a boolean if a field has been set.

### GetReduceOnly

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetReduceOnly() string`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetReduceOnlyOk() (*string, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetReduceOnly(v string)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSide

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSlOrdPx

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetSlOrdPx() string`

GetSlOrdPx returns the SlOrdPx field if non-nil, zero value otherwise.

### GetSlOrdPxOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetSlOrdPxOk() (*string, bool)`

GetSlOrdPxOk returns a tuple with the SlOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlOrdPx

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetSlOrdPx(v string)`

SetSlOrdPx sets SlOrdPx field to given value.

### HasSlOrdPx

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasSlOrdPx() bool`

HasSlOrdPx returns a boolean if a field has been set.

### GetSlTriggerPx

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetSlTriggerPx() string`

GetSlTriggerPx returns the SlTriggerPx field if non-nil, zero value otherwise.

### GetSlTriggerPxOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetSlTriggerPxOk() (*string, bool)`

GetSlTriggerPxOk returns a tuple with the SlTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPx

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetSlTriggerPx(v string)`

SetSlTriggerPx sets SlTriggerPx field to given value.

### HasSlTriggerPx

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasSlTriggerPx() bool`

HasSlTriggerPx returns a boolean if a field has been set.

### GetSlTriggerPxType

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetSlTriggerPxType() string`

GetSlTriggerPxType returns the SlTriggerPxType field if non-nil, zero value otherwise.

### GetSlTriggerPxTypeOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetSlTriggerPxTypeOk() (*string, bool)`

GetSlTriggerPxTypeOk returns a tuple with the SlTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPxType

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetSlTriggerPxType(v string)`

SetSlTriggerPxType sets SlTriggerPxType field to given value.

### HasSlTriggerPxType

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasSlTriggerPxType() bool`

HasSlTriggerPxType returns a boolean if a field has been set.

### GetSource

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetState

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStpId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetStpId() string`

GetStpId returns the StpId field if non-nil, zero value otherwise.

### GetStpIdOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetStpIdOk() (*string, bool)`

GetStpIdOk returns a tuple with the StpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetStpId(v string)`

SetStpId sets StpId field to given value.

### HasStpId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasStpId() bool`

HasStpId returns a boolean if a field has been set.

### GetStpMode

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetStpMode() string`

GetStpMode returns the StpMode field if non-nil, zero value otherwise.

### GetStpModeOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetStpModeOk() (*string, bool)`

GetStpModeOk returns a tuple with the StpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpMode

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetStpMode(v string)`

SetStpMode sets StpMode field to given value.

### HasStpMode

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasStpMode() bool`

HasStpMode returns a boolean if a field has been set.

### GetSz

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTag

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTdMode

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetTdMode() string`

GetTdMode returns the TdMode field if non-nil, zero value otherwise.

### GetTdModeOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetTdModeOk() (*string, bool)`

GetTdModeOk returns a tuple with the TdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdMode

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetTdMode(v string)`

SetTdMode sets TdMode field to given value.

### HasTdMode

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasTdMode() bool`

HasTdMode returns a boolean if a field has been set.

### GetTgtCcy

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetTgtCcy() string`

GetTgtCcy returns the TgtCcy field if non-nil, zero value otherwise.

### GetTgtCcyOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetTgtCcyOk() (*string, bool)`

GetTgtCcyOk returns a tuple with the TgtCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCcy

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetTgtCcy(v string)`

SetTgtCcy sets TgtCcy field to given value.

### HasTgtCcy

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasTgtCcy() bool`

HasTgtCcy returns a boolean if a field has been set.

### GetTpOrdPx

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetTpOrdPx() string`

GetTpOrdPx returns the TpOrdPx field if non-nil, zero value otherwise.

### GetTpOrdPxOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetTpOrdPxOk() (*string, bool)`

GetTpOrdPxOk returns a tuple with the TpOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpOrdPx

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetTpOrdPx(v string)`

SetTpOrdPx sets TpOrdPx field to given value.

### HasTpOrdPx

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasTpOrdPx() bool`

HasTpOrdPx returns a boolean if a field has been set.

### GetTpTriggerPx

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetTpTriggerPx() string`

GetTpTriggerPx returns the TpTriggerPx field if non-nil, zero value otherwise.

### GetTpTriggerPxOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetTpTriggerPxOk() (*string, bool)`

GetTpTriggerPxOk returns a tuple with the TpTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPx

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetTpTriggerPx(v string)`

SetTpTriggerPx sets TpTriggerPx field to given value.

### HasTpTriggerPx

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasTpTriggerPx() bool`

HasTpTriggerPx returns a boolean if a field has been set.

### GetTpTriggerPxType

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetTpTriggerPxType() string`

GetTpTriggerPxType returns the TpTriggerPxType field if non-nil, zero value otherwise.

### GetTpTriggerPxTypeOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetTpTriggerPxTypeOk() (*string, bool)`

GetTpTriggerPxTypeOk returns a tuple with the TpTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPxType

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetTpTriggerPxType(v string)`

SetTpTriggerPxType sets TpTriggerPxType field to given value.

### HasTpTriggerPxType

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasTpTriggerPxType() bool`

HasTpTriggerPxType returns a boolean if a field has been set.

### GetTradeId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.

### GetUTime

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetTradeOrdersHistoryArchiveV5RespData) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetTradeOrdersHistoryArchiveV5RespData) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetTradeOrdersHistoryArchiveV5RespData) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


