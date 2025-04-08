# GetTradingBotSignalSubOrdersV5RespData

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

### NewGetTradingBotSignalSubOrdersV5RespData

`func NewGetTradingBotSignalSubOrdersV5RespData() *GetTradingBotSignalSubOrdersV5RespData`

NewGetTradingBotSignalSubOrdersV5RespData instantiates a new GetTradingBotSignalSubOrdersV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotSignalSubOrdersV5RespDataWithDefaults

`func NewGetTradingBotSignalSubOrdersV5RespDataWithDefaults() *GetTradingBotSignalSubOrdersV5RespData`

NewGetTradingBotSignalSubOrdersV5RespDataWithDefaults instantiates a new GetTradingBotSignalSubOrdersV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccFillSz

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetAccFillSz() string`

GetAccFillSz returns the AccFillSz field if non-nil, zero value otherwise.

### GetAccFillSzOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetAccFillSzOk() (*string, bool)`

GetAccFillSzOk returns a tuple with the AccFillSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccFillSz

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetAccFillSz(v string)`

SetAccFillSz sets AccFillSz field to given value.

### HasAccFillSz

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasAccFillSz() bool`

HasAccFillSz returns a boolean if a field has been set.

### GetAlgoClOrdId

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAlgoId

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetAlgoOrdType

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetAlgoOrdType() string`

GetAlgoOrdType returns the AlgoOrdType field if non-nil, zero value otherwise.

### GetAlgoOrdTypeOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetAlgoOrdTypeOk() (*string, bool)`

GetAlgoOrdTypeOk returns a tuple with the AlgoOrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoOrdType

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetAlgoOrdType(v string)`

SetAlgoOrdType sets AlgoOrdType field to given value.

### HasAlgoOrdType

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasAlgoOrdType() bool`

HasAlgoOrdType returns a boolean if a field has been set.

### GetAvgPx

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetAvgPx() string`

GetAvgPx returns the AvgPx field if non-nil, zero value otherwise.

### GetAvgPxOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetAvgPxOk() (*string, bool)`

GetAvgPxOk returns a tuple with the AvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPx

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetAvgPx(v string)`

SetAvgPx sets AvgPx field to given value.

### HasAvgPx

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasAvgPx() bool`

HasAvgPx returns a boolean if a field has been set.

### GetCTime

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCcy

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetClOrdId

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetCtVal

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetCtVal() string`

GetCtVal returns the CtVal field if non-nil, zero value otherwise.

### GetCtValOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetCtValOk() (*string, bool)`

GetCtValOk returns a tuple with the CtVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtVal

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetCtVal(v string)`

SetCtVal sets CtVal field to given value.

### HasCtVal

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasCtVal() bool`

HasCtVal returns a boolean if a field has been set.

### GetFee

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFeeCcy

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetFeeCcy() string`

GetFeeCcy returns the FeeCcy field if non-nil, zero value otherwise.

### GetFeeCcyOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetFeeCcyOk() (*string, bool)`

GetFeeCcyOk returns a tuple with the FeeCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCcy

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetFeeCcy(v string)`

SetFeeCcy sets FeeCcy field to given value.

### HasFeeCcy

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasFeeCcy() bool`

HasFeeCcy returns a boolean if a field has been set.

### GetInstId

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetLever

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetOrdId

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetOrdType

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetOrdType() string`

GetOrdType returns the OrdType field if non-nil, zero value otherwise.

### GetOrdTypeOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetOrdTypeOk() (*string, bool)`

GetOrdTypeOk returns a tuple with the OrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdType

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetOrdType(v string)`

SetOrdType sets OrdType field to given value.

### HasOrdType

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasOrdType() bool`

HasOrdType returns a boolean if a field has been set.

### GetPnl

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetPnl() string`

GetPnl returns the Pnl field if non-nil, zero value otherwise.

### GetPnlOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetPnlOk() (*string, bool)`

GetPnlOk returns a tuple with the Pnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnl

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetPnl(v string)`

SetPnl sets Pnl field to given value.

### HasPnl

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasPnl() bool`

HasPnl returns a boolean if a field has been set.

### GetPosSide

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetPx

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSide

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetState

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSz

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTag

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTdMode

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetTdMode() string`

GetTdMode returns the TdMode field if non-nil, zero value otherwise.

### GetTdModeOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetTdModeOk() (*string, bool)`

GetTdModeOk returns a tuple with the TdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdMode

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetTdMode(v string)`

SetTdMode sets TdMode field to given value.

### HasTdMode

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasTdMode() bool`

HasTdMode returns a boolean if a field has been set.

### GetUTime

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetTradingBotSignalSubOrdersV5RespData) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetTradingBotSignalSubOrdersV5RespData) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetTradingBotSignalSubOrdersV5RespData) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


