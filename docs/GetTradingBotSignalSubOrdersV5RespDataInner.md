# GetTradingBotSignalSubOrdersV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccFillSz** | Pointer to **string** | Sub order accumulated fill quantity | [optional] [default to ""]
**AlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID. Used to be extended in the future | [optional] [default to ""]
**AlgoId** | Pointer to **string** | Algo ID | [optional] [default to ""]
**AlgoOrdType** | Pointer to **string** | Algo order type  &#x60;contract&#x60;: Contract signal | [optional] [default to ""]
**AvgPx** | Pointer to **string** | Sub order average filled price | [optional] [default to ""]
**CTime** | Pointer to **string** | Sub order created time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Ccy** | Pointer to **string** | Margin currency  Only applicable to cross MARGIN orders in &#x60;Spot and futures mode&#x60;. | [optional] [default to ""]
**ClOrdId** | Pointer to **string** | Sub order client-supplied ID.    It is equal to &#x60;signalOrdId&#x60; | [optional] [default to ""]
**CtVal** | Pointer to **string** | Contract value  Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60; | [optional] [default to ""]
**Fee** | Pointer to **string** | Sub order fee amount | [optional] [default to ""]
**FeeCcy** | Pointer to **string** | Sub order fee currency | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**Lever** | Pointer to **string** | Leverage | [optional] [default to ""]
**OrdId** | Pointer to **string** | Sub order ID | [optional] [default to ""]
**OrdType** | Pointer to **string** | Sub order type  &#x60;market&#x60;: Market order  &#x60;limit&#x60;: Limit order  &#x60;ioc&#x60;: Immediate-or-cancel order | [optional] [default to ""]
**Pnl** | Pointer to **string** | Sub order profit and loss | [optional] [default to ""]
**PosSide** | Pointer to **string** | Sub order position side  &#x60;net&#x60; | [optional] [default to ""]
**Px** | Pointer to **string** | Sub order price | [optional] [default to ""]
**Side** | Pointer to **string** | Sub order side  &#x60;buy&#x60;,&#x60;sell&#x60; | [optional] [default to ""]
**State** | Pointer to **string** | Sub order state  &#x60;canceled&#x60;  &#x60;live&#x60;  &#x60;partially_filled&#x60;  &#x60;filled&#x60;  &#x60;cancelling&#x60; | [optional] [default to ""]
**Sz** | Pointer to **string** | Sub order quantity to buy or sell | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**TdMode** | Pointer to **string** | Sub order trade mode  Margin mode: &#x60;cross&#x60;/&#x60;isolated&#x60;  Non-Margin mode: &#x60;cash&#x60; | [optional] [default to ""]
**UTime** | Pointer to **string** | Sub order updated time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetTradingBotSignalSubOrdersV5RespDataInner

`func NewGetTradingBotSignalSubOrdersV5RespDataInner() *GetTradingBotSignalSubOrdersV5RespDataInner`

NewGetTradingBotSignalSubOrdersV5RespDataInner instantiates a new GetTradingBotSignalSubOrdersV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotSignalSubOrdersV5RespDataInnerWithDefaults

`func NewGetTradingBotSignalSubOrdersV5RespDataInnerWithDefaults() *GetTradingBotSignalSubOrdersV5RespDataInner`

NewGetTradingBotSignalSubOrdersV5RespDataInnerWithDefaults instantiates a new GetTradingBotSignalSubOrdersV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccFillSz

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetAccFillSz() string`

GetAccFillSz returns the AccFillSz field if non-nil, zero value otherwise.

### GetAccFillSzOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetAccFillSzOk() (*string, bool)`

GetAccFillSzOk returns a tuple with the AccFillSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccFillSz

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetAccFillSz(v string)`

SetAccFillSz sets AccFillSz field to given value.

### HasAccFillSz

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasAccFillSz() bool`

HasAccFillSz returns a boolean if a field has been set.

### GetAlgoClOrdId

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAlgoId

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetAlgoOrdType

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetAlgoOrdType() string`

GetAlgoOrdType returns the AlgoOrdType field if non-nil, zero value otherwise.

### GetAlgoOrdTypeOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetAlgoOrdTypeOk() (*string, bool)`

GetAlgoOrdTypeOk returns a tuple with the AlgoOrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoOrdType

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetAlgoOrdType(v string)`

SetAlgoOrdType sets AlgoOrdType field to given value.

### HasAlgoOrdType

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasAlgoOrdType() bool`

HasAlgoOrdType returns a boolean if a field has been set.

### GetAvgPx

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetAvgPx() string`

GetAvgPx returns the AvgPx field if non-nil, zero value otherwise.

### GetAvgPxOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetAvgPxOk() (*string, bool)`

GetAvgPxOk returns a tuple with the AvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPx

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetAvgPx(v string)`

SetAvgPx sets AvgPx field to given value.

### HasAvgPx

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasAvgPx() bool`

HasAvgPx returns a boolean if a field has been set.

### GetCTime

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCcy

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetClOrdId

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetCtVal

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetCtVal() string`

GetCtVal returns the CtVal field if non-nil, zero value otherwise.

### GetCtValOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetCtValOk() (*string, bool)`

GetCtValOk returns a tuple with the CtVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtVal

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetCtVal(v string)`

SetCtVal sets CtVal field to given value.

### HasCtVal

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasCtVal() bool`

HasCtVal returns a boolean if a field has been set.

### GetFee

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFeeCcy

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetFeeCcy() string`

GetFeeCcy returns the FeeCcy field if non-nil, zero value otherwise.

### GetFeeCcyOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetFeeCcyOk() (*string, bool)`

GetFeeCcyOk returns a tuple with the FeeCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCcy

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetFeeCcy(v string)`

SetFeeCcy sets FeeCcy field to given value.

### HasFeeCcy

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasFeeCcy() bool`

HasFeeCcy returns a boolean if a field has been set.

### GetInstId

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetLever

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetOrdId

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetOrdType

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetOrdType() string`

GetOrdType returns the OrdType field if non-nil, zero value otherwise.

### GetOrdTypeOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetOrdTypeOk() (*string, bool)`

GetOrdTypeOk returns a tuple with the OrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdType

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetOrdType(v string)`

SetOrdType sets OrdType field to given value.

### HasOrdType

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasOrdType() bool`

HasOrdType returns a boolean if a field has been set.

### GetPnl

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetPnl() string`

GetPnl returns the Pnl field if non-nil, zero value otherwise.

### GetPnlOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetPnlOk() (*string, bool)`

GetPnlOk returns a tuple with the Pnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnl

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetPnl(v string)`

SetPnl sets Pnl field to given value.

### HasPnl

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasPnl() bool`

HasPnl returns a boolean if a field has been set.

### GetPosSide

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetPx

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSide

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetState

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSz

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTag

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTdMode

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetTdMode() string`

GetTdMode returns the TdMode field if non-nil, zero value otherwise.

### GetTdModeOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetTdModeOk() (*string, bool)`

GetTdModeOk returns a tuple with the TdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdMode

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetTdMode(v string)`

SetTdMode sets TdMode field to given value.

### HasTdMode

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasTdMode() bool`

HasTdMode returns a boolean if a field has been set.

### GetUTime

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetTradingBotSignalSubOrdersV5RespDataInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


