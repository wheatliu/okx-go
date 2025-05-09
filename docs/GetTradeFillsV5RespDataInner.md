# GetTradeFillsV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillId** | Pointer to **string** | Bill ID | [optional] [default to ""]
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client | [optional] [default to ""]
**ExecType** | Pointer to **string** | Liquidity taker or maker  &#x60;T&#x60;: taker  &#x60;M&#x60;: maker  Not applicable to system orders such as ADL and liquidation | [optional] [default to ""]
**Fee** | Pointer to **string** | The amount of trading fee or rebate. The trading fee deduction is negative, such as &#39;-0.01&#39;; the rebate is positive, such as &#39;0.01&#39;. | [optional] [default to ""]
**FeeCcy** | Pointer to **string** | Trading fee or rebate currency | [optional] [default to ""]
**FeeRate** | Pointer to **string** | Fee rate. This field is returned for &#x60;SPOT&#x60; and &#x60;MARGIN&#x60; only | [optional] [default to ""]
**FillFwdPx** | Pointer to **string** | Forward price when filled   Only applicable to options; return \&quot;\&quot; for other instrument types | [optional] [default to ""]
**FillIdxPx** | Pointer to **string** | Index price at the moment of trade execution   For cross currency spot pairs, it returns baseCcy-USDT index price. For example, for LTC-ETH, this field returns the index price of LTC-USDT. | [optional] [default to ""]
**FillMarkPx** | Pointer to **string** | Mark price when filled   Applicable to &#x60;FUTURES&#x60;, &#x60;SWAP&#x60;, &#x60;OPTION&#x60; | [optional] [default to ""]
**FillMarkVol** | Pointer to **string** | Mark volatility when filled   Only applicable to options; return \&quot;\&quot; for other instrument types | [optional] [default to ""]
**FillPnl** | Pointer to **string** | Last filled profit and loss, applicable to orders which have a trade and aim to close position. It always is 0 in other conditions | [optional] [default to ""]
**FillPx** | Pointer to **string** | Last filled price. It is the same as the px from \&quot;Get bills details\&quot;. | [optional] [default to ""]
**FillPxUsd** | Pointer to **string** | Options price when filled, in the unit of USD   Only applicable to options; return \&quot;\&quot; for other instrument types | [optional] [default to ""]
**FillPxVol** | Pointer to **string** | Implied volatility when filled   Only applicable to options; return \&quot;\&quot; for other instrument types | [optional] [default to ""]
**FillSz** | Pointer to **string** | Last filled quantity | [optional] [default to ""]
**FillTime** | Pointer to **string** | Trade time which is the same as &#x60;fillTime&#x60; for the order channel. | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**OrdId** | Pointer to **string** | Order ID | [optional] [default to ""]
**PosSide** | Pointer to **string** | Position side   &#x60;long&#x60;  &#x60;short&#x60;     it returns &#x60;net&#x60; in&#x60;net&#x60; mode. | [optional] [default to ""]
**Side** | Pointer to **string** | Order side,  &#x60;buy&#x60;  &#x60;sell&#x60; | [optional] [default to ""]
**SubType** | Pointer to **string** | Transaction type | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**TradeId** | Pointer to **string** | Last trade ID | [optional] [default to ""]
**Ts** | Pointer to **string** | Data generation time,  Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60;. | [optional] [default to ""]

## Methods

### NewGetTradeFillsV5RespDataInner

`func NewGetTradeFillsV5RespDataInner() *GetTradeFillsV5RespDataInner`

NewGetTradeFillsV5RespDataInner instantiates a new GetTradeFillsV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradeFillsV5RespDataInnerWithDefaults

`func NewGetTradeFillsV5RespDataInnerWithDefaults() *GetTradeFillsV5RespDataInner`

NewGetTradeFillsV5RespDataInnerWithDefaults instantiates a new GetTradeFillsV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillId

`func (o *GetTradeFillsV5RespDataInner) GetBillId() string`

GetBillId returns the BillId field if non-nil, zero value otherwise.

### GetBillIdOk

`func (o *GetTradeFillsV5RespDataInner) GetBillIdOk() (*string, bool)`

GetBillIdOk returns a tuple with the BillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillId

`func (o *GetTradeFillsV5RespDataInner) SetBillId(v string)`

SetBillId sets BillId field to given value.

### HasBillId

`func (o *GetTradeFillsV5RespDataInner) HasBillId() bool`

HasBillId returns a boolean if a field has been set.

### GetClOrdId

`func (o *GetTradeFillsV5RespDataInner) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *GetTradeFillsV5RespDataInner) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *GetTradeFillsV5RespDataInner) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *GetTradeFillsV5RespDataInner) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetExecType

`func (o *GetTradeFillsV5RespDataInner) GetExecType() string`

GetExecType returns the ExecType field if non-nil, zero value otherwise.

### GetExecTypeOk

`func (o *GetTradeFillsV5RespDataInner) GetExecTypeOk() (*string, bool)`

GetExecTypeOk returns a tuple with the ExecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecType

`func (o *GetTradeFillsV5RespDataInner) SetExecType(v string)`

SetExecType sets ExecType field to given value.

### HasExecType

`func (o *GetTradeFillsV5RespDataInner) HasExecType() bool`

HasExecType returns a boolean if a field has been set.

### GetFee

`func (o *GetTradeFillsV5RespDataInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetTradeFillsV5RespDataInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetTradeFillsV5RespDataInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetTradeFillsV5RespDataInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFeeCcy

`func (o *GetTradeFillsV5RespDataInner) GetFeeCcy() string`

GetFeeCcy returns the FeeCcy field if non-nil, zero value otherwise.

### GetFeeCcyOk

`func (o *GetTradeFillsV5RespDataInner) GetFeeCcyOk() (*string, bool)`

GetFeeCcyOk returns a tuple with the FeeCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCcy

`func (o *GetTradeFillsV5RespDataInner) SetFeeCcy(v string)`

SetFeeCcy sets FeeCcy field to given value.

### HasFeeCcy

`func (o *GetTradeFillsV5RespDataInner) HasFeeCcy() bool`

HasFeeCcy returns a boolean if a field has been set.

### GetFeeRate

`func (o *GetTradeFillsV5RespDataInner) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *GetTradeFillsV5RespDataInner) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *GetTradeFillsV5RespDataInner) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.

### HasFeeRate

`func (o *GetTradeFillsV5RespDataInner) HasFeeRate() bool`

HasFeeRate returns a boolean if a field has been set.

### GetFillFwdPx

`func (o *GetTradeFillsV5RespDataInner) GetFillFwdPx() string`

GetFillFwdPx returns the FillFwdPx field if non-nil, zero value otherwise.

### GetFillFwdPxOk

`func (o *GetTradeFillsV5RespDataInner) GetFillFwdPxOk() (*string, bool)`

GetFillFwdPxOk returns a tuple with the FillFwdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillFwdPx

`func (o *GetTradeFillsV5RespDataInner) SetFillFwdPx(v string)`

SetFillFwdPx sets FillFwdPx field to given value.

### HasFillFwdPx

`func (o *GetTradeFillsV5RespDataInner) HasFillFwdPx() bool`

HasFillFwdPx returns a boolean if a field has been set.

### GetFillIdxPx

`func (o *GetTradeFillsV5RespDataInner) GetFillIdxPx() string`

GetFillIdxPx returns the FillIdxPx field if non-nil, zero value otherwise.

### GetFillIdxPxOk

`func (o *GetTradeFillsV5RespDataInner) GetFillIdxPxOk() (*string, bool)`

GetFillIdxPxOk returns a tuple with the FillIdxPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillIdxPx

`func (o *GetTradeFillsV5RespDataInner) SetFillIdxPx(v string)`

SetFillIdxPx sets FillIdxPx field to given value.

### HasFillIdxPx

`func (o *GetTradeFillsV5RespDataInner) HasFillIdxPx() bool`

HasFillIdxPx returns a boolean if a field has been set.

### GetFillMarkPx

`func (o *GetTradeFillsV5RespDataInner) GetFillMarkPx() string`

GetFillMarkPx returns the FillMarkPx field if non-nil, zero value otherwise.

### GetFillMarkPxOk

`func (o *GetTradeFillsV5RespDataInner) GetFillMarkPxOk() (*string, bool)`

GetFillMarkPxOk returns a tuple with the FillMarkPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillMarkPx

`func (o *GetTradeFillsV5RespDataInner) SetFillMarkPx(v string)`

SetFillMarkPx sets FillMarkPx field to given value.

### HasFillMarkPx

`func (o *GetTradeFillsV5RespDataInner) HasFillMarkPx() bool`

HasFillMarkPx returns a boolean if a field has been set.

### GetFillMarkVol

`func (o *GetTradeFillsV5RespDataInner) GetFillMarkVol() string`

GetFillMarkVol returns the FillMarkVol field if non-nil, zero value otherwise.

### GetFillMarkVolOk

`func (o *GetTradeFillsV5RespDataInner) GetFillMarkVolOk() (*string, bool)`

GetFillMarkVolOk returns a tuple with the FillMarkVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillMarkVol

`func (o *GetTradeFillsV5RespDataInner) SetFillMarkVol(v string)`

SetFillMarkVol sets FillMarkVol field to given value.

### HasFillMarkVol

`func (o *GetTradeFillsV5RespDataInner) HasFillMarkVol() bool`

HasFillMarkVol returns a boolean if a field has been set.

### GetFillPnl

`func (o *GetTradeFillsV5RespDataInner) GetFillPnl() string`

GetFillPnl returns the FillPnl field if non-nil, zero value otherwise.

### GetFillPnlOk

`func (o *GetTradeFillsV5RespDataInner) GetFillPnlOk() (*string, bool)`

GetFillPnlOk returns a tuple with the FillPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPnl

`func (o *GetTradeFillsV5RespDataInner) SetFillPnl(v string)`

SetFillPnl sets FillPnl field to given value.

### HasFillPnl

`func (o *GetTradeFillsV5RespDataInner) HasFillPnl() bool`

HasFillPnl returns a boolean if a field has been set.

### GetFillPx

`func (o *GetTradeFillsV5RespDataInner) GetFillPx() string`

GetFillPx returns the FillPx field if non-nil, zero value otherwise.

### GetFillPxOk

`func (o *GetTradeFillsV5RespDataInner) GetFillPxOk() (*string, bool)`

GetFillPxOk returns a tuple with the FillPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPx

`func (o *GetTradeFillsV5RespDataInner) SetFillPx(v string)`

SetFillPx sets FillPx field to given value.

### HasFillPx

`func (o *GetTradeFillsV5RespDataInner) HasFillPx() bool`

HasFillPx returns a boolean if a field has been set.

### GetFillPxUsd

`func (o *GetTradeFillsV5RespDataInner) GetFillPxUsd() string`

GetFillPxUsd returns the FillPxUsd field if non-nil, zero value otherwise.

### GetFillPxUsdOk

`func (o *GetTradeFillsV5RespDataInner) GetFillPxUsdOk() (*string, bool)`

GetFillPxUsdOk returns a tuple with the FillPxUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPxUsd

`func (o *GetTradeFillsV5RespDataInner) SetFillPxUsd(v string)`

SetFillPxUsd sets FillPxUsd field to given value.

### HasFillPxUsd

`func (o *GetTradeFillsV5RespDataInner) HasFillPxUsd() bool`

HasFillPxUsd returns a boolean if a field has been set.

### GetFillPxVol

`func (o *GetTradeFillsV5RespDataInner) GetFillPxVol() string`

GetFillPxVol returns the FillPxVol field if non-nil, zero value otherwise.

### GetFillPxVolOk

`func (o *GetTradeFillsV5RespDataInner) GetFillPxVolOk() (*string, bool)`

GetFillPxVolOk returns a tuple with the FillPxVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPxVol

`func (o *GetTradeFillsV5RespDataInner) SetFillPxVol(v string)`

SetFillPxVol sets FillPxVol field to given value.

### HasFillPxVol

`func (o *GetTradeFillsV5RespDataInner) HasFillPxVol() bool`

HasFillPxVol returns a boolean if a field has been set.

### GetFillSz

`func (o *GetTradeFillsV5RespDataInner) GetFillSz() string`

GetFillSz returns the FillSz field if non-nil, zero value otherwise.

### GetFillSzOk

`func (o *GetTradeFillsV5RespDataInner) GetFillSzOk() (*string, bool)`

GetFillSzOk returns a tuple with the FillSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillSz

`func (o *GetTradeFillsV5RespDataInner) SetFillSz(v string)`

SetFillSz sets FillSz field to given value.

### HasFillSz

`func (o *GetTradeFillsV5RespDataInner) HasFillSz() bool`

HasFillSz returns a boolean if a field has been set.

### GetFillTime

`func (o *GetTradeFillsV5RespDataInner) GetFillTime() string`

GetFillTime returns the FillTime field if non-nil, zero value otherwise.

### GetFillTimeOk

`func (o *GetTradeFillsV5RespDataInner) GetFillTimeOk() (*string, bool)`

GetFillTimeOk returns a tuple with the FillTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillTime

`func (o *GetTradeFillsV5RespDataInner) SetFillTime(v string)`

SetFillTime sets FillTime field to given value.

### HasFillTime

`func (o *GetTradeFillsV5RespDataInner) HasFillTime() bool`

HasFillTime returns a boolean if a field has been set.

### GetInstId

`func (o *GetTradeFillsV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetTradeFillsV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetTradeFillsV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetTradeFillsV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetTradeFillsV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetTradeFillsV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetTradeFillsV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetTradeFillsV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetOrdId

`func (o *GetTradeFillsV5RespDataInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *GetTradeFillsV5RespDataInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *GetTradeFillsV5RespDataInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *GetTradeFillsV5RespDataInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetPosSide

`func (o *GetTradeFillsV5RespDataInner) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *GetTradeFillsV5RespDataInner) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *GetTradeFillsV5RespDataInner) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *GetTradeFillsV5RespDataInner) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetSide

`func (o *GetTradeFillsV5RespDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetTradeFillsV5RespDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetTradeFillsV5RespDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetTradeFillsV5RespDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSubType

`func (o *GetTradeFillsV5RespDataInner) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *GetTradeFillsV5RespDataInner) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *GetTradeFillsV5RespDataInner) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *GetTradeFillsV5RespDataInner) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetTag

`func (o *GetTradeFillsV5RespDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetTradeFillsV5RespDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetTradeFillsV5RespDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetTradeFillsV5RespDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTradeId

`func (o *GetTradeFillsV5RespDataInner) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *GetTradeFillsV5RespDataInner) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *GetTradeFillsV5RespDataInner) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *GetTradeFillsV5RespDataInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.

### GetTs

`func (o *GetTradeFillsV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetTradeFillsV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetTradeFillsV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetTradeFillsV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


