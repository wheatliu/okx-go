# GetTradeOrderV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccFillSz** | Pointer to **string** | Accumulated fill quantity  The unit is &#x60;base_ccy&#x60; for SPOT and MARGIN, e.g. BTC-USDT, the unit is BTC; For market orders, the unit both is &#x60;base_ccy&#x60; when the tgtCcy is &#x60;base_ccy&#x60; or &#x60;quote_ccy&#x60;;  The unit is contract for &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [optional] [default to ""]
**AlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID. There will be a value when algo order attaching &#x60;algoClOrdId&#x60; is triggered, or it will be \&quot;\&quot;. | [optional] [default to ""]
**AlgoId** | Pointer to **string** | Algo ID. There will be a value when algo order is triggered, or it will be \&quot;\&quot;. | [optional] [default to ""]
**AttachAlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID when placing order attaching TP/SL. | [optional] [default to ""]
**AttachAlgoOrds** | Pointer to [**[]GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner**](GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner.md) | TP/SL information attached when placing order | [optional] 
**AvgPx** | Pointer to **string** | Average filled price. If none is filled, it will return \&quot;\&quot;. | [optional] [default to ""]
**CTime** | Pointer to **string** | Creation time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**CancelSource** | Pointer to **string** | Code of the cancellation source. | [optional] [default to ""]
**CancelSourceReason** | Pointer to **string** | Reason for the cancellation. | [optional] [default to ""]
**Category** | Pointer to **string** | Category  &#x60;normal&#x60;  &#x60;twap&#x60;   &#x60;adl&#x60;  &#x60;full_liquidation&#x60;  &#x60;partial_liquidation&#x60;  &#x60;delivery&#x60;    &#x60;ddh&#x60;: Delta dynamic hedge | [optional] [default to ""]
**Ccy** | Pointer to **string** | Margin currency   Applicable to all &#x60;isolated&#x60; &#x60;MARGIN&#x60; orders and &#x60;cross&#x60; &#x60;MARGIN&#x60; orders in &#x60;Spot and futures mode&#x60;. | [optional] [default to ""]
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client | [optional] [default to ""]
**Fee** | Pointer to **string** | Fee and rebate  For spot and margin, it is accumulated fee charged by the platform. It is always negative, e.g. -0.01.   For Expiry Futures, Perpetual Futures and Options, it is accumulated fee and rebate | [optional] [default to ""]
**FeeCcy** | Pointer to **string** | Fee currency | [optional] [default to ""]
**FillPx** | Pointer to **string** | Last filled price. If none is filled, it will return \&quot;\&quot;. | [optional] [default to ""]
**FillSz** | Pointer to **string** | Last filled quantity  The unit is &#x60;base_ccy&#x60; for SPOT and MARGIN, e.g. BTC-USDT, the unit is BTC; For market orders, the unit both is &#x60;base_ccy&#x60; when the tgtCcy is &#x60;base_ccy&#x60; or &#x60;quote_ccy&#x60;;  The unit is contract for &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [optional] [default to ""]
**FillTime** | Pointer to **string** | Last filled time | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;MARGIN&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [optional] [default to ""]
**IsTpLimit** | Pointer to **string** | Whether it is TP limit order. true or false | [optional] [default to ""]
**Lever** | Pointer to **string** | Leverage, from &#x60;0.01&#x60; to &#x60;125&#x60;.   Only applicable to &#x60;MARGIN/FUTURES/SWAP&#x60; | [optional] [default to ""]
**LinkedAlgoOrd** | Pointer to [**GetTradeOrderV5RespDataInnerLinkedAlgoOrd**](GetTradeOrderV5RespDataInnerLinkedAlgoOrd.md) |  | [optional] 
**OrdId** | Pointer to **string** | Order ID | [optional] [default to ""]
**OrdType** | Pointer to **string** | Order type   &#x60;market&#x60;: Market order   &#x60;limit&#x60;: Limit order   &#x60;post_only&#x60;: Post-only order   &#x60;fok&#x60;: Fill-or-kill order   &#x60;ioc&#x60;: Immediate-or-cancel order    &#x60;optimal_limit_ioc&#x60;: Market order with immediate-or-cancel order  &#x60;mmp&#x60;: Market Maker Protection (only applicable to Option in Portfolio Margin mode)  &#x60;mmp_and_post_only&#x60;: Market Maker Protection and Post-only order(only applicable to Option in Portfolio Margin mode)   &#x60;op_fok&#x60;: Simple options (fok) | [optional] [default to ""]
**Pnl** | Pointer to **string** | Profit and loss, Applicable to orders which have a trade and aim to close position. It always is 0 in other conditions | [optional] [default to ""]
**PosSide** | Pointer to **string** | Position side | [optional] [default to ""]
**Px** | Pointer to **string** | Price  For options, use coin as unit (e.g. BTC, ETH) | [optional] [default to ""]
**PxType** | Pointer to **string** | Price type of options   &#x60;px&#x60;: Place an order based on price, in the unit of coin (the unit for the request parameter px is BTC or ETH)   &#x60;pxVol&#x60;: Place an order based on pxVol   &#x60;pxUsd&#x60;: Place an order based on pxUsd, in the unit of USD (the unit for the request parameter px is USD) | [optional] [default to ""]
**PxUsd** | Pointer to **string** | Options price in USDOnly applicable to options; return \&quot;\&quot; for other instrument types | [optional] [default to ""]
**PxVol** | Pointer to **string** | Implied volatility of the options orderOnly applicable to options; return \&quot;\&quot; for other instrument types | [optional] [default to ""]
**QuickMgnType** | Pointer to **string** | Quick Margin type, Only applicable to Quick Margin Mode of isolated margin  &#x60;manual&#x60;, &#x60;auto_borrow&#x60;, &#x60;auto_repay&#x60; | [optional] [default to ""]
**Rebate** | Pointer to **string** | Rebate amount, only applicable to spot and margin, the reward of placing orders from the platform (rebate) given to user who has reached the specified trading level. If there is no rebate, this field is \&quot;\&quot;. | [optional] [default to ""]
**RebateCcy** | Pointer to **string** | Rebate currency | [optional] [default to ""]
**ReduceOnly** | Pointer to **string** | Whether the order can only reduce the position size. Valid options: true or false. | [optional] [default to ""]
**Side** | Pointer to **string** | Order side | [optional] [default to ""]
**SlOrdPx** | Pointer to **string** | Stop-loss order price. | [optional] [default to ""]
**SlTriggerPx** | Pointer to **string** | Stop-loss trigger price. | [optional] [default to ""]
**SlTriggerPxType** | Pointer to **string** | Stop-loss trigger price type.   &#x60;last&#x60;: last price  &#x60;index&#x60;: index price  &#x60;mark&#x60;: mark price | [optional] [default to ""]
**Source** | Pointer to **string** | Order source  &#x60;6&#x60;: The normal order triggered by the &#x60;trigger order&#x60;  &#x60;7&#x60;:The normal order triggered by the &#x60;TP/SL order&#x60;   &#x60;13&#x60;: The normal order triggered by the algo order  &#x60;25&#x60;:The normal order triggered by the &#x60;trailing stop order&#x60;  &#x60;34&#x60;: The normal order triggered by the chase order | [optional] [default to ""]
**State** | Pointer to **string** | State   &#x60;canceled&#x60;  &#x60;live&#x60;   &#x60;partially_filled&#x60;  &#x60;filled&#x60;  &#x60;mmp_canceled&#x60; | [optional] [default to ""]
**StpId** | Pointer to **string** | Self trade prevention ID  Return \&quot;\&quot; if self trade prevention is not applicable (deprecated) | [optional] [default to ""]
**StpMode** | Pointer to **string** | Self trade prevention mode   Return \&quot;\&quot; if self trade prevention is not applicable | [optional] [default to ""]
**Sz** | Pointer to **string** | Quantity to buy or sell | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**TdMode** | Pointer to **string** | Trade mode | [optional] [default to ""]
**TgtCcy** | Pointer to **string** | Order quantity unit setting for &#x60;sz&#x60;  &#x60;base_ccy&#x60;: Base currency ,&#x60;quote_ccy&#x60;: Quote currency    Only applicable to &#x60;SPOT&#x60; Market Orders  Default is &#x60;quote_ccy&#x60; for buy, &#x60;base_ccy&#x60; for sell | [optional] [default to ""]
**TpOrdPx** | Pointer to **string** | Take-profit order price. | [optional] [default to ""]
**TpTriggerPx** | Pointer to **string** | Take-profit trigger price. | [optional] [default to ""]
**TpTriggerPxType** | Pointer to **string** | Take-profit trigger price type.   &#x60;last&#x60;: last price  &#x60;index&#x60;: index price  &#x60;mark&#x60;: mark price | [optional] [default to ""]
**TradeId** | Pointer to **string** | Last traded ID | [optional] [default to ""]
**UTime** | Pointer to **string** | Update time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetTradeOrderV5RespDataInner

`func NewGetTradeOrderV5RespDataInner() *GetTradeOrderV5RespDataInner`

NewGetTradeOrderV5RespDataInner instantiates a new GetTradeOrderV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradeOrderV5RespDataInnerWithDefaults

`func NewGetTradeOrderV5RespDataInnerWithDefaults() *GetTradeOrderV5RespDataInner`

NewGetTradeOrderV5RespDataInnerWithDefaults instantiates a new GetTradeOrderV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccFillSz

`func (o *GetTradeOrderV5RespDataInner) GetAccFillSz() string`

GetAccFillSz returns the AccFillSz field if non-nil, zero value otherwise.

### GetAccFillSzOk

`func (o *GetTradeOrderV5RespDataInner) GetAccFillSzOk() (*string, bool)`

GetAccFillSzOk returns a tuple with the AccFillSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccFillSz

`func (o *GetTradeOrderV5RespDataInner) SetAccFillSz(v string)`

SetAccFillSz sets AccFillSz field to given value.

### HasAccFillSz

`func (o *GetTradeOrderV5RespDataInner) HasAccFillSz() bool`

HasAccFillSz returns a boolean if a field has been set.

### GetAlgoClOrdId

`func (o *GetTradeOrderV5RespDataInner) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *GetTradeOrderV5RespDataInner) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *GetTradeOrderV5RespDataInner) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *GetTradeOrderV5RespDataInner) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAlgoId

`func (o *GetTradeOrderV5RespDataInner) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetTradeOrderV5RespDataInner) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetTradeOrderV5RespDataInner) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetTradeOrderV5RespDataInner) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetAttachAlgoClOrdId

`func (o *GetTradeOrderV5RespDataInner) GetAttachAlgoClOrdId() string`

GetAttachAlgoClOrdId returns the AttachAlgoClOrdId field if non-nil, zero value otherwise.

### GetAttachAlgoClOrdIdOk

`func (o *GetTradeOrderV5RespDataInner) GetAttachAlgoClOrdIdOk() (*string, bool)`

GetAttachAlgoClOrdIdOk returns a tuple with the AttachAlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachAlgoClOrdId

`func (o *GetTradeOrderV5RespDataInner) SetAttachAlgoClOrdId(v string)`

SetAttachAlgoClOrdId sets AttachAlgoClOrdId field to given value.

### HasAttachAlgoClOrdId

`func (o *GetTradeOrderV5RespDataInner) HasAttachAlgoClOrdId() bool`

HasAttachAlgoClOrdId returns a boolean if a field has been set.

### GetAttachAlgoOrds

`func (o *GetTradeOrderV5RespDataInner) GetAttachAlgoOrds() []GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner`

GetAttachAlgoOrds returns the AttachAlgoOrds field if non-nil, zero value otherwise.

### GetAttachAlgoOrdsOk

`func (o *GetTradeOrderV5RespDataInner) GetAttachAlgoOrdsOk() (*[]GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner, bool)`

GetAttachAlgoOrdsOk returns a tuple with the AttachAlgoOrds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachAlgoOrds

`func (o *GetTradeOrderV5RespDataInner) SetAttachAlgoOrds(v []GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner)`

SetAttachAlgoOrds sets AttachAlgoOrds field to given value.

### HasAttachAlgoOrds

`func (o *GetTradeOrderV5RespDataInner) HasAttachAlgoOrds() bool`

HasAttachAlgoOrds returns a boolean if a field has been set.

### GetAvgPx

`func (o *GetTradeOrderV5RespDataInner) GetAvgPx() string`

GetAvgPx returns the AvgPx field if non-nil, zero value otherwise.

### GetAvgPxOk

`func (o *GetTradeOrderV5RespDataInner) GetAvgPxOk() (*string, bool)`

GetAvgPxOk returns a tuple with the AvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPx

`func (o *GetTradeOrderV5RespDataInner) SetAvgPx(v string)`

SetAvgPx sets AvgPx field to given value.

### HasAvgPx

`func (o *GetTradeOrderV5RespDataInner) HasAvgPx() bool`

HasAvgPx returns a boolean if a field has been set.

### GetCTime

`func (o *GetTradeOrderV5RespDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetTradeOrderV5RespDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetTradeOrderV5RespDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetTradeOrderV5RespDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCancelSource

`func (o *GetTradeOrderV5RespDataInner) GetCancelSource() string`

GetCancelSource returns the CancelSource field if non-nil, zero value otherwise.

### GetCancelSourceOk

`func (o *GetTradeOrderV5RespDataInner) GetCancelSourceOk() (*string, bool)`

GetCancelSourceOk returns a tuple with the CancelSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelSource

`func (o *GetTradeOrderV5RespDataInner) SetCancelSource(v string)`

SetCancelSource sets CancelSource field to given value.

### HasCancelSource

`func (o *GetTradeOrderV5RespDataInner) HasCancelSource() bool`

HasCancelSource returns a boolean if a field has been set.

### GetCancelSourceReason

`func (o *GetTradeOrderV5RespDataInner) GetCancelSourceReason() string`

GetCancelSourceReason returns the CancelSourceReason field if non-nil, zero value otherwise.

### GetCancelSourceReasonOk

`func (o *GetTradeOrderV5RespDataInner) GetCancelSourceReasonOk() (*string, bool)`

GetCancelSourceReasonOk returns a tuple with the CancelSourceReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelSourceReason

`func (o *GetTradeOrderV5RespDataInner) SetCancelSourceReason(v string)`

SetCancelSourceReason sets CancelSourceReason field to given value.

### HasCancelSourceReason

`func (o *GetTradeOrderV5RespDataInner) HasCancelSourceReason() bool`

HasCancelSourceReason returns a boolean if a field has been set.

### GetCategory

`func (o *GetTradeOrderV5RespDataInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetTradeOrderV5RespDataInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetTradeOrderV5RespDataInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetTradeOrderV5RespDataInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCcy

`func (o *GetTradeOrderV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetTradeOrderV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetTradeOrderV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetTradeOrderV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetClOrdId

`func (o *GetTradeOrderV5RespDataInner) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *GetTradeOrderV5RespDataInner) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *GetTradeOrderV5RespDataInner) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *GetTradeOrderV5RespDataInner) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetFee

`func (o *GetTradeOrderV5RespDataInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetTradeOrderV5RespDataInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetTradeOrderV5RespDataInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetTradeOrderV5RespDataInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFeeCcy

`func (o *GetTradeOrderV5RespDataInner) GetFeeCcy() string`

GetFeeCcy returns the FeeCcy field if non-nil, zero value otherwise.

### GetFeeCcyOk

`func (o *GetTradeOrderV5RespDataInner) GetFeeCcyOk() (*string, bool)`

GetFeeCcyOk returns a tuple with the FeeCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCcy

`func (o *GetTradeOrderV5RespDataInner) SetFeeCcy(v string)`

SetFeeCcy sets FeeCcy field to given value.

### HasFeeCcy

`func (o *GetTradeOrderV5RespDataInner) HasFeeCcy() bool`

HasFeeCcy returns a boolean if a field has been set.

### GetFillPx

`func (o *GetTradeOrderV5RespDataInner) GetFillPx() string`

GetFillPx returns the FillPx field if non-nil, zero value otherwise.

### GetFillPxOk

`func (o *GetTradeOrderV5RespDataInner) GetFillPxOk() (*string, bool)`

GetFillPxOk returns a tuple with the FillPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPx

`func (o *GetTradeOrderV5RespDataInner) SetFillPx(v string)`

SetFillPx sets FillPx field to given value.

### HasFillPx

`func (o *GetTradeOrderV5RespDataInner) HasFillPx() bool`

HasFillPx returns a boolean if a field has been set.

### GetFillSz

`func (o *GetTradeOrderV5RespDataInner) GetFillSz() string`

GetFillSz returns the FillSz field if non-nil, zero value otherwise.

### GetFillSzOk

`func (o *GetTradeOrderV5RespDataInner) GetFillSzOk() (*string, bool)`

GetFillSzOk returns a tuple with the FillSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillSz

`func (o *GetTradeOrderV5RespDataInner) SetFillSz(v string)`

SetFillSz sets FillSz field to given value.

### HasFillSz

`func (o *GetTradeOrderV5RespDataInner) HasFillSz() bool`

HasFillSz returns a boolean if a field has been set.

### GetFillTime

`func (o *GetTradeOrderV5RespDataInner) GetFillTime() string`

GetFillTime returns the FillTime field if non-nil, zero value otherwise.

### GetFillTimeOk

`func (o *GetTradeOrderV5RespDataInner) GetFillTimeOk() (*string, bool)`

GetFillTimeOk returns a tuple with the FillTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillTime

`func (o *GetTradeOrderV5RespDataInner) SetFillTime(v string)`

SetFillTime sets FillTime field to given value.

### HasFillTime

`func (o *GetTradeOrderV5RespDataInner) HasFillTime() bool`

HasFillTime returns a boolean if a field has been set.

### GetInstId

`func (o *GetTradeOrderV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetTradeOrderV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetTradeOrderV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetTradeOrderV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetTradeOrderV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetTradeOrderV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetTradeOrderV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetTradeOrderV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetIsTpLimit

`func (o *GetTradeOrderV5RespDataInner) GetIsTpLimit() string`

GetIsTpLimit returns the IsTpLimit field if non-nil, zero value otherwise.

### GetIsTpLimitOk

`func (o *GetTradeOrderV5RespDataInner) GetIsTpLimitOk() (*string, bool)`

GetIsTpLimitOk returns a tuple with the IsTpLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTpLimit

`func (o *GetTradeOrderV5RespDataInner) SetIsTpLimit(v string)`

SetIsTpLimit sets IsTpLimit field to given value.

### HasIsTpLimit

`func (o *GetTradeOrderV5RespDataInner) HasIsTpLimit() bool`

HasIsTpLimit returns a boolean if a field has been set.

### GetLever

`func (o *GetTradeOrderV5RespDataInner) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetTradeOrderV5RespDataInner) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetTradeOrderV5RespDataInner) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetTradeOrderV5RespDataInner) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetLinkedAlgoOrd

`func (o *GetTradeOrderV5RespDataInner) GetLinkedAlgoOrd() GetTradeOrderV5RespDataInnerLinkedAlgoOrd`

GetLinkedAlgoOrd returns the LinkedAlgoOrd field if non-nil, zero value otherwise.

### GetLinkedAlgoOrdOk

`func (o *GetTradeOrderV5RespDataInner) GetLinkedAlgoOrdOk() (*GetTradeOrderV5RespDataInnerLinkedAlgoOrd, bool)`

GetLinkedAlgoOrdOk returns a tuple with the LinkedAlgoOrd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAlgoOrd

`func (o *GetTradeOrderV5RespDataInner) SetLinkedAlgoOrd(v GetTradeOrderV5RespDataInnerLinkedAlgoOrd)`

SetLinkedAlgoOrd sets LinkedAlgoOrd field to given value.

### HasLinkedAlgoOrd

`func (o *GetTradeOrderV5RespDataInner) HasLinkedAlgoOrd() bool`

HasLinkedAlgoOrd returns a boolean if a field has been set.

### GetOrdId

`func (o *GetTradeOrderV5RespDataInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *GetTradeOrderV5RespDataInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *GetTradeOrderV5RespDataInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *GetTradeOrderV5RespDataInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetOrdType

`func (o *GetTradeOrderV5RespDataInner) GetOrdType() string`

GetOrdType returns the OrdType field if non-nil, zero value otherwise.

### GetOrdTypeOk

`func (o *GetTradeOrderV5RespDataInner) GetOrdTypeOk() (*string, bool)`

GetOrdTypeOk returns a tuple with the OrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdType

`func (o *GetTradeOrderV5RespDataInner) SetOrdType(v string)`

SetOrdType sets OrdType field to given value.

### HasOrdType

`func (o *GetTradeOrderV5RespDataInner) HasOrdType() bool`

HasOrdType returns a boolean if a field has been set.

### GetPnl

`func (o *GetTradeOrderV5RespDataInner) GetPnl() string`

GetPnl returns the Pnl field if non-nil, zero value otherwise.

### GetPnlOk

`func (o *GetTradeOrderV5RespDataInner) GetPnlOk() (*string, bool)`

GetPnlOk returns a tuple with the Pnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnl

`func (o *GetTradeOrderV5RespDataInner) SetPnl(v string)`

SetPnl sets Pnl field to given value.

### HasPnl

`func (o *GetTradeOrderV5RespDataInner) HasPnl() bool`

HasPnl returns a boolean if a field has been set.

### GetPosSide

`func (o *GetTradeOrderV5RespDataInner) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *GetTradeOrderV5RespDataInner) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *GetTradeOrderV5RespDataInner) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *GetTradeOrderV5RespDataInner) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetPx

`func (o *GetTradeOrderV5RespDataInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetTradeOrderV5RespDataInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetTradeOrderV5RespDataInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetTradeOrderV5RespDataInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetPxType

`func (o *GetTradeOrderV5RespDataInner) GetPxType() string`

GetPxType returns the PxType field if non-nil, zero value otherwise.

### GetPxTypeOk

`func (o *GetTradeOrderV5RespDataInner) GetPxTypeOk() (*string, bool)`

GetPxTypeOk returns a tuple with the PxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxType

`func (o *GetTradeOrderV5RespDataInner) SetPxType(v string)`

SetPxType sets PxType field to given value.

### HasPxType

`func (o *GetTradeOrderV5RespDataInner) HasPxType() bool`

HasPxType returns a boolean if a field has been set.

### GetPxUsd

`func (o *GetTradeOrderV5RespDataInner) GetPxUsd() string`

GetPxUsd returns the PxUsd field if non-nil, zero value otherwise.

### GetPxUsdOk

`func (o *GetTradeOrderV5RespDataInner) GetPxUsdOk() (*string, bool)`

GetPxUsdOk returns a tuple with the PxUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxUsd

`func (o *GetTradeOrderV5RespDataInner) SetPxUsd(v string)`

SetPxUsd sets PxUsd field to given value.

### HasPxUsd

`func (o *GetTradeOrderV5RespDataInner) HasPxUsd() bool`

HasPxUsd returns a boolean if a field has been set.

### GetPxVol

`func (o *GetTradeOrderV5RespDataInner) GetPxVol() string`

GetPxVol returns the PxVol field if non-nil, zero value otherwise.

### GetPxVolOk

`func (o *GetTradeOrderV5RespDataInner) GetPxVolOk() (*string, bool)`

GetPxVolOk returns a tuple with the PxVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxVol

`func (o *GetTradeOrderV5RespDataInner) SetPxVol(v string)`

SetPxVol sets PxVol field to given value.

### HasPxVol

`func (o *GetTradeOrderV5RespDataInner) HasPxVol() bool`

HasPxVol returns a boolean if a field has been set.

### GetQuickMgnType

`func (o *GetTradeOrderV5RespDataInner) GetQuickMgnType() string`

GetQuickMgnType returns the QuickMgnType field if non-nil, zero value otherwise.

### GetQuickMgnTypeOk

`func (o *GetTradeOrderV5RespDataInner) GetQuickMgnTypeOk() (*string, bool)`

GetQuickMgnTypeOk returns a tuple with the QuickMgnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickMgnType

`func (o *GetTradeOrderV5RespDataInner) SetQuickMgnType(v string)`

SetQuickMgnType sets QuickMgnType field to given value.

### HasQuickMgnType

`func (o *GetTradeOrderV5RespDataInner) HasQuickMgnType() bool`

HasQuickMgnType returns a boolean if a field has been set.

### GetRebate

`func (o *GetTradeOrderV5RespDataInner) GetRebate() string`

GetRebate returns the Rebate field if non-nil, zero value otherwise.

### GetRebateOk

`func (o *GetTradeOrderV5RespDataInner) GetRebateOk() (*string, bool)`

GetRebateOk returns a tuple with the Rebate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebate

`func (o *GetTradeOrderV5RespDataInner) SetRebate(v string)`

SetRebate sets Rebate field to given value.

### HasRebate

`func (o *GetTradeOrderV5RespDataInner) HasRebate() bool`

HasRebate returns a boolean if a field has been set.

### GetRebateCcy

`func (o *GetTradeOrderV5RespDataInner) GetRebateCcy() string`

GetRebateCcy returns the RebateCcy field if non-nil, zero value otherwise.

### GetRebateCcyOk

`func (o *GetTradeOrderV5RespDataInner) GetRebateCcyOk() (*string, bool)`

GetRebateCcyOk returns a tuple with the RebateCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebateCcy

`func (o *GetTradeOrderV5RespDataInner) SetRebateCcy(v string)`

SetRebateCcy sets RebateCcy field to given value.

### HasRebateCcy

`func (o *GetTradeOrderV5RespDataInner) HasRebateCcy() bool`

HasRebateCcy returns a boolean if a field has been set.

### GetReduceOnly

`func (o *GetTradeOrderV5RespDataInner) GetReduceOnly() string`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *GetTradeOrderV5RespDataInner) GetReduceOnlyOk() (*string, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *GetTradeOrderV5RespDataInner) SetReduceOnly(v string)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *GetTradeOrderV5RespDataInner) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSide

`func (o *GetTradeOrderV5RespDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetTradeOrderV5RespDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetTradeOrderV5RespDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetTradeOrderV5RespDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSlOrdPx

`func (o *GetTradeOrderV5RespDataInner) GetSlOrdPx() string`

GetSlOrdPx returns the SlOrdPx field if non-nil, zero value otherwise.

### GetSlOrdPxOk

`func (o *GetTradeOrderV5RespDataInner) GetSlOrdPxOk() (*string, bool)`

GetSlOrdPxOk returns a tuple with the SlOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlOrdPx

`func (o *GetTradeOrderV5RespDataInner) SetSlOrdPx(v string)`

SetSlOrdPx sets SlOrdPx field to given value.

### HasSlOrdPx

`func (o *GetTradeOrderV5RespDataInner) HasSlOrdPx() bool`

HasSlOrdPx returns a boolean if a field has been set.

### GetSlTriggerPx

`func (o *GetTradeOrderV5RespDataInner) GetSlTriggerPx() string`

GetSlTriggerPx returns the SlTriggerPx field if non-nil, zero value otherwise.

### GetSlTriggerPxOk

`func (o *GetTradeOrderV5RespDataInner) GetSlTriggerPxOk() (*string, bool)`

GetSlTriggerPxOk returns a tuple with the SlTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPx

`func (o *GetTradeOrderV5RespDataInner) SetSlTriggerPx(v string)`

SetSlTriggerPx sets SlTriggerPx field to given value.

### HasSlTriggerPx

`func (o *GetTradeOrderV5RespDataInner) HasSlTriggerPx() bool`

HasSlTriggerPx returns a boolean if a field has been set.

### GetSlTriggerPxType

`func (o *GetTradeOrderV5RespDataInner) GetSlTriggerPxType() string`

GetSlTriggerPxType returns the SlTriggerPxType field if non-nil, zero value otherwise.

### GetSlTriggerPxTypeOk

`func (o *GetTradeOrderV5RespDataInner) GetSlTriggerPxTypeOk() (*string, bool)`

GetSlTriggerPxTypeOk returns a tuple with the SlTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPxType

`func (o *GetTradeOrderV5RespDataInner) SetSlTriggerPxType(v string)`

SetSlTriggerPxType sets SlTriggerPxType field to given value.

### HasSlTriggerPxType

`func (o *GetTradeOrderV5RespDataInner) HasSlTriggerPxType() bool`

HasSlTriggerPxType returns a boolean if a field has been set.

### GetSource

`func (o *GetTradeOrderV5RespDataInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetTradeOrderV5RespDataInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetTradeOrderV5RespDataInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetTradeOrderV5RespDataInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetState

`func (o *GetTradeOrderV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetTradeOrderV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetTradeOrderV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetTradeOrderV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStpId

`func (o *GetTradeOrderV5RespDataInner) GetStpId() string`

GetStpId returns the StpId field if non-nil, zero value otherwise.

### GetStpIdOk

`func (o *GetTradeOrderV5RespDataInner) GetStpIdOk() (*string, bool)`

GetStpIdOk returns a tuple with the StpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpId

`func (o *GetTradeOrderV5RespDataInner) SetStpId(v string)`

SetStpId sets StpId field to given value.

### HasStpId

`func (o *GetTradeOrderV5RespDataInner) HasStpId() bool`

HasStpId returns a boolean if a field has been set.

### GetStpMode

`func (o *GetTradeOrderV5RespDataInner) GetStpMode() string`

GetStpMode returns the StpMode field if non-nil, zero value otherwise.

### GetStpModeOk

`func (o *GetTradeOrderV5RespDataInner) GetStpModeOk() (*string, bool)`

GetStpModeOk returns a tuple with the StpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpMode

`func (o *GetTradeOrderV5RespDataInner) SetStpMode(v string)`

SetStpMode sets StpMode field to given value.

### HasStpMode

`func (o *GetTradeOrderV5RespDataInner) HasStpMode() bool`

HasStpMode returns a boolean if a field has been set.

### GetSz

`func (o *GetTradeOrderV5RespDataInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetTradeOrderV5RespDataInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetTradeOrderV5RespDataInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetTradeOrderV5RespDataInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTag

`func (o *GetTradeOrderV5RespDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetTradeOrderV5RespDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetTradeOrderV5RespDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetTradeOrderV5RespDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTdMode

`func (o *GetTradeOrderV5RespDataInner) GetTdMode() string`

GetTdMode returns the TdMode field if non-nil, zero value otherwise.

### GetTdModeOk

`func (o *GetTradeOrderV5RespDataInner) GetTdModeOk() (*string, bool)`

GetTdModeOk returns a tuple with the TdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdMode

`func (o *GetTradeOrderV5RespDataInner) SetTdMode(v string)`

SetTdMode sets TdMode field to given value.

### HasTdMode

`func (o *GetTradeOrderV5RespDataInner) HasTdMode() bool`

HasTdMode returns a boolean if a field has been set.

### GetTgtCcy

`func (o *GetTradeOrderV5RespDataInner) GetTgtCcy() string`

GetTgtCcy returns the TgtCcy field if non-nil, zero value otherwise.

### GetTgtCcyOk

`func (o *GetTradeOrderV5RespDataInner) GetTgtCcyOk() (*string, bool)`

GetTgtCcyOk returns a tuple with the TgtCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCcy

`func (o *GetTradeOrderV5RespDataInner) SetTgtCcy(v string)`

SetTgtCcy sets TgtCcy field to given value.

### HasTgtCcy

`func (o *GetTradeOrderV5RespDataInner) HasTgtCcy() bool`

HasTgtCcy returns a boolean if a field has been set.

### GetTpOrdPx

`func (o *GetTradeOrderV5RespDataInner) GetTpOrdPx() string`

GetTpOrdPx returns the TpOrdPx field if non-nil, zero value otherwise.

### GetTpOrdPxOk

`func (o *GetTradeOrderV5RespDataInner) GetTpOrdPxOk() (*string, bool)`

GetTpOrdPxOk returns a tuple with the TpOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpOrdPx

`func (o *GetTradeOrderV5RespDataInner) SetTpOrdPx(v string)`

SetTpOrdPx sets TpOrdPx field to given value.

### HasTpOrdPx

`func (o *GetTradeOrderV5RespDataInner) HasTpOrdPx() bool`

HasTpOrdPx returns a boolean if a field has been set.

### GetTpTriggerPx

`func (o *GetTradeOrderV5RespDataInner) GetTpTriggerPx() string`

GetTpTriggerPx returns the TpTriggerPx field if non-nil, zero value otherwise.

### GetTpTriggerPxOk

`func (o *GetTradeOrderV5RespDataInner) GetTpTriggerPxOk() (*string, bool)`

GetTpTriggerPxOk returns a tuple with the TpTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPx

`func (o *GetTradeOrderV5RespDataInner) SetTpTriggerPx(v string)`

SetTpTriggerPx sets TpTriggerPx field to given value.

### HasTpTriggerPx

`func (o *GetTradeOrderV5RespDataInner) HasTpTriggerPx() bool`

HasTpTriggerPx returns a boolean if a field has been set.

### GetTpTriggerPxType

`func (o *GetTradeOrderV5RespDataInner) GetTpTriggerPxType() string`

GetTpTriggerPxType returns the TpTriggerPxType field if non-nil, zero value otherwise.

### GetTpTriggerPxTypeOk

`func (o *GetTradeOrderV5RespDataInner) GetTpTriggerPxTypeOk() (*string, bool)`

GetTpTriggerPxTypeOk returns a tuple with the TpTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPxType

`func (o *GetTradeOrderV5RespDataInner) SetTpTriggerPxType(v string)`

SetTpTriggerPxType sets TpTriggerPxType field to given value.

### HasTpTriggerPxType

`func (o *GetTradeOrderV5RespDataInner) HasTpTriggerPxType() bool`

HasTpTriggerPxType returns a boolean if a field has been set.

### GetTradeId

`func (o *GetTradeOrderV5RespDataInner) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *GetTradeOrderV5RespDataInner) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *GetTradeOrderV5RespDataInner) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *GetTradeOrderV5RespDataInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.

### GetUTime

`func (o *GetTradeOrderV5RespDataInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetTradeOrderV5RespDataInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetTradeOrderV5RespDataInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetTradeOrderV5RespDataInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


