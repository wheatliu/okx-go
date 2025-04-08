# GetTradeFillsHistoryV5RespData

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
**FillPx** | Pointer to **string** | Last filled price | [optional] [default to ""]
**FillPxUsd** | Pointer to **string** | Options price when filled, in the unit of USD   Only applicable to options; return \&quot;\&quot; for other instrument types | [optional] [default to ""]
**FillPxVol** | Pointer to **string** | Implied volatility when filled   Only applicable to options; return \&quot;\&quot; for other instrument types | [optional] [default to ""]
**FillSz** | Pointer to **string** | Last filled quantity | [optional] [default to ""]
**FillTime** | Pointer to **string** | Trade time which is the same as &#x60;fillTime&#x60; for the order channel. | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**OrdId** | Pointer to **string** | Order ID | [optional] [default to ""]
**PosSide** | Pointer to **string** | Position side  &#x60;long&#x60;  &#x60;short&#x60;  it returns &#x60;net&#x60; in&#x60;net&#x60; mode. | [optional] [default to ""]
**Side** | Pointer to **string** | Order side  &#x60;buy&#x60;  &#x60;sell&#x60; | [optional] [default to ""]
**SubType** | Pointer to **string** | Transaction type | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**TradeId** | Pointer to **string** | Last trade ID | [optional] [default to ""]
**Ts** | Pointer to **string** | Data generation time,  Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60;. | [optional] [default to ""]

## Methods

### NewGetTradeFillsHistoryV5RespData

`func NewGetTradeFillsHistoryV5RespData() *GetTradeFillsHistoryV5RespData`

NewGetTradeFillsHistoryV5RespData instantiates a new GetTradeFillsHistoryV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradeFillsHistoryV5RespDataWithDefaults

`func NewGetTradeFillsHistoryV5RespDataWithDefaults() *GetTradeFillsHistoryV5RespData`

NewGetTradeFillsHistoryV5RespDataWithDefaults instantiates a new GetTradeFillsHistoryV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillId

`func (o *GetTradeFillsHistoryV5RespData) GetBillId() string`

GetBillId returns the BillId field if non-nil, zero value otherwise.

### GetBillIdOk

`func (o *GetTradeFillsHistoryV5RespData) GetBillIdOk() (*string, bool)`

GetBillIdOk returns a tuple with the BillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillId

`func (o *GetTradeFillsHistoryV5RespData) SetBillId(v string)`

SetBillId sets BillId field to given value.

### HasBillId

`func (o *GetTradeFillsHistoryV5RespData) HasBillId() bool`

HasBillId returns a boolean if a field has been set.

### GetClOrdId

`func (o *GetTradeFillsHistoryV5RespData) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *GetTradeFillsHistoryV5RespData) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *GetTradeFillsHistoryV5RespData) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *GetTradeFillsHistoryV5RespData) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetExecType

`func (o *GetTradeFillsHistoryV5RespData) GetExecType() string`

GetExecType returns the ExecType field if non-nil, zero value otherwise.

### GetExecTypeOk

`func (o *GetTradeFillsHistoryV5RespData) GetExecTypeOk() (*string, bool)`

GetExecTypeOk returns a tuple with the ExecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecType

`func (o *GetTradeFillsHistoryV5RespData) SetExecType(v string)`

SetExecType sets ExecType field to given value.

### HasExecType

`func (o *GetTradeFillsHistoryV5RespData) HasExecType() bool`

HasExecType returns a boolean if a field has been set.

### GetFee

`func (o *GetTradeFillsHistoryV5RespData) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetTradeFillsHistoryV5RespData) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetTradeFillsHistoryV5RespData) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetTradeFillsHistoryV5RespData) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFeeCcy

`func (o *GetTradeFillsHistoryV5RespData) GetFeeCcy() string`

GetFeeCcy returns the FeeCcy field if non-nil, zero value otherwise.

### GetFeeCcyOk

`func (o *GetTradeFillsHistoryV5RespData) GetFeeCcyOk() (*string, bool)`

GetFeeCcyOk returns a tuple with the FeeCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCcy

`func (o *GetTradeFillsHistoryV5RespData) SetFeeCcy(v string)`

SetFeeCcy sets FeeCcy field to given value.

### HasFeeCcy

`func (o *GetTradeFillsHistoryV5RespData) HasFeeCcy() bool`

HasFeeCcy returns a boolean if a field has been set.

### GetFeeRate

`func (o *GetTradeFillsHistoryV5RespData) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *GetTradeFillsHistoryV5RespData) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *GetTradeFillsHistoryV5RespData) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.

### HasFeeRate

`func (o *GetTradeFillsHistoryV5RespData) HasFeeRate() bool`

HasFeeRate returns a boolean if a field has been set.

### GetFillFwdPx

`func (o *GetTradeFillsHistoryV5RespData) GetFillFwdPx() string`

GetFillFwdPx returns the FillFwdPx field if non-nil, zero value otherwise.

### GetFillFwdPxOk

`func (o *GetTradeFillsHistoryV5RespData) GetFillFwdPxOk() (*string, bool)`

GetFillFwdPxOk returns a tuple with the FillFwdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillFwdPx

`func (o *GetTradeFillsHistoryV5RespData) SetFillFwdPx(v string)`

SetFillFwdPx sets FillFwdPx field to given value.

### HasFillFwdPx

`func (o *GetTradeFillsHistoryV5RespData) HasFillFwdPx() bool`

HasFillFwdPx returns a boolean if a field has been set.

### GetFillIdxPx

`func (o *GetTradeFillsHistoryV5RespData) GetFillIdxPx() string`

GetFillIdxPx returns the FillIdxPx field if non-nil, zero value otherwise.

### GetFillIdxPxOk

`func (o *GetTradeFillsHistoryV5RespData) GetFillIdxPxOk() (*string, bool)`

GetFillIdxPxOk returns a tuple with the FillIdxPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillIdxPx

`func (o *GetTradeFillsHistoryV5RespData) SetFillIdxPx(v string)`

SetFillIdxPx sets FillIdxPx field to given value.

### HasFillIdxPx

`func (o *GetTradeFillsHistoryV5RespData) HasFillIdxPx() bool`

HasFillIdxPx returns a boolean if a field has been set.

### GetFillMarkPx

`func (o *GetTradeFillsHistoryV5RespData) GetFillMarkPx() string`

GetFillMarkPx returns the FillMarkPx field if non-nil, zero value otherwise.

### GetFillMarkPxOk

`func (o *GetTradeFillsHistoryV5RespData) GetFillMarkPxOk() (*string, bool)`

GetFillMarkPxOk returns a tuple with the FillMarkPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillMarkPx

`func (o *GetTradeFillsHistoryV5RespData) SetFillMarkPx(v string)`

SetFillMarkPx sets FillMarkPx field to given value.

### HasFillMarkPx

`func (o *GetTradeFillsHistoryV5RespData) HasFillMarkPx() bool`

HasFillMarkPx returns a boolean if a field has been set.

### GetFillMarkVol

`func (o *GetTradeFillsHistoryV5RespData) GetFillMarkVol() string`

GetFillMarkVol returns the FillMarkVol field if non-nil, zero value otherwise.

### GetFillMarkVolOk

`func (o *GetTradeFillsHistoryV5RespData) GetFillMarkVolOk() (*string, bool)`

GetFillMarkVolOk returns a tuple with the FillMarkVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillMarkVol

`func (o *GetTradeFillsHistoryV5RespData) SetFillMarkVol(v string)`

SetFillMarkVol sets FillMarkVol field to given value.

### HasFillMarkVol

`func (o *GetTradeFillsHistoryV5RespData) HasFillMarkVol() bool`

HasFillMarkVol returns a boolean if a field has been set.

### GetFillPnl

`func (o *GetTradeFillsHistoryV5RespData) GetFillPnl() string`

GetFillPnl returns the FillPnl field if non-nil, zero value otherwise.

### GetFillPnlOk

`func (o *GetTradeFillsHistoryV5RespData) GetFillPnlOk() (*string, bool)`

GetFillPnlOk returns a tuple with the FillPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPnl

`func (o *GetTradeFillsHistoryV5RespData) SetFillPnl(v string)`

SetFillPnl sets FillPnl field to given value.

### HasFillPnl

`func (o *GetTradeFillsHistoryV5RespData) HasFillPnl() bool`

HasFillPnl returns a boolean if a field has been set.

### GetFillPx

`func (o *GetTradeFillsHistoryV5RespData) GetFillPx() string`

GetFillPx returns the FillPx field if non-nil, zero value otherwise.

### GetFillPxOk

`func (o *GetTradeFillsHistoryV5RespData) GetFillPxOk() (*string, bool)`

GetFillPxOk returns a tuple with the FillPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPx

`func (o *GetTradeFillsHistoryV5RespData) SetFillPx(v string)`

SetFillPx sets FillPx field to given value.

### HasFillPx

`func (o *GetTradeFillsHistoryV5RespData) HasFillPx() bool`

HasFillPx returns a boolean if a field has been set.

### GetFillPxUsd

`func (o *GetTradeFillsHistoryV5RespData) GetFillPxUsd() string`

GetFillPxUsd returns the FillPxUsd field if non-nil, zero value otherwise.

### GetFillPxUsdOk

`func (o *GetTradeFillsHistoryV5RespData) GetFillPxUsdOk() (*string, bool)`

GetFillPxUsdOk returns a tuple with the FillPxUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPxUsd

`func (o *GetTradeFillsHistoryV5RespData) SetFillPxUsd(v string)`

SetFillPxUsd sets FillPxUsd field to given value.

### HasFillPxUsd

`func (o *GetTradeFillsHistoryV5RespData) HasFillPxUsd() bool`

HasFillPxUsd returns a boolean if a field has been set.

### GetFillPxVol

`func (o *GetTradeFillsHistoryV5RespData) GetFillPxVol() string`

GetFillPxVol returns the FillPxVol field if non-nil, zero value otherwise.

### GetFillPxVolOk

`func (o *GetTradeFillsHistoryV5RespData) GetFillPxVolOk() (*string, bool)`

GetFillPxVolOk returns a tuple with the FillPxVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPxVol

`func (o *GetTradeFillsHistoryV5RespData) SetFillPxVol(v string)`

SetFillPxVol sets FillPxVol field to given value.

### HasFillPxVol

`func (o *GetTradeFillsHistoryV5RespData) HasFillPxVol() bool`

HasFillPxVol returns a boolean if a field has been set.

### GetFillSz

`func (o *GetTradeFillsHistoryV5RespData) GetFillSz() string`

GetFillSz returns the FillSz field if non-nil, zero value otherwise.

### GetFillSzOk

`func (o *GetTradeFillsHistoryV5RespData) GetFillSzOk() (*string, bool)`

GetFillSzOk returns a tuple with the FillSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillSz

`func (o *GetTradeFillsHistoryV5RespData) SetFillSz(v string)`

SetFillSz sets FillSz field to given value.

### HasFillSz

`func (o *GetTradeFillsHistoryV5RespData) HasFillSz() bool`

HasFillSz returns a boolean if a field has been set.

### GetFillTime

`func (o *GetTradeFillsHistoryV5RespData) GetFillTime() string`

GetFillTime returns the FillTime field if non-nil, zero value otherwise.

### GetFillTimeOk

`func (o *GetTradeFillsHistoryV5RespData) GetFillTimeOk() (*string, bool)`

GetFillTimeOk returns a tuple with the FillTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillTime

`func (o *GetTradeFillsHistoryV5RespData) SetFillTime(v string)`

SetFillTime sets FillTime field to given value.

### HasFillTime

`func (o *GetTradeFillsHistoryV5RespData) HasFillTime() bool`

HasFillTime returns a boolean if a field has been set.

### GetInstId

`func (o *GetTradeFillsHistoryV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetTradeFillsHistoryV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetTradeFillsHistoryV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetTradeFillsHistoryV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetTradeFillsHistoryV5RespData) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetTradeFillsHistoryV5RespData) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetTradeFillsHistoryV5RespData) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetTradeFillsHistoryV5RespData) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetOrdId

`func (o *GetTradeFillsHistoryV5RespData) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *GetTradeFillsHistoryV5RespData) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *GetTradeFillsHistoryV5RespData) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *GetTradeFillsHistoryV5RespData) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetPosSide

`func (o *GetTradeFillsHistoryV5RespData) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *GetTradeFillsHistoryV5RespData) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *GetTradeFillsHistoryV5RespData) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *GetTradeFillsHistoryV5RespData) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetSide

`func (o *GetTradeFillsHistoryV5RespData) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetTradeFillsHistoryV5RespData) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetTradeFillsHistoryV5RespData) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetTradeFillsHistoryV5RespData) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSubType

`func (o *GetTradeFillsHistoryV5RespData) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *GetTradeFillsHistoryV5RespData) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *GetTradeFillsHistoryV5RespData) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *GetTradeFillsHistoryV5RespData) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetTag

`func (o *GetTradeFillsHistoryV5RespData) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetTradeFillsHistoryV5RespData) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetTradeFillsHistoryV5RespData) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetTradeFillsHistoryV5RespData) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTradeId

`func (o *GetTradeFillsHistoryV5RespData) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *GetTradeFillsHistoryV5RespData) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *GetTradeFillsHistoryV5RespData) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *GetTradeFillsHistoryV5RespData) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.

### GetTs

`func (o *GetTradeFillsHistoryV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetTradeFillsHistoryV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetTradeFillsHistoryV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetTradeFillsHistoryV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


