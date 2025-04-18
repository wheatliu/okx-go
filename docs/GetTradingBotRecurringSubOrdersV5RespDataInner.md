# GetTradingBotRecurringSubOrdersV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccFillSz** | Pointer to **string** | Sub order accumulated fill quantity | [optional] [default to ""]
**AlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID | [optional] [default to ""]
**AlgoId** | Pointer to **string** | Algo ID | [optional] [default to ""]
**AlgoOrdType** | Pointer to **string** | Algo order type  &#x60;recurring&#x60;: recurring buy | [optional] [default to ""]
**AvgPx** | Pointer to **string** | Sub order average filled price | [optional] [default to ""]
**CTime** | Pointer to **string** | Sub order created time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Fee** | Pointer to **string** | Sub order fee | [optional] [default to ""]
**FeeCcy** | Pointer to **string** | Sub order fee currency | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**OrdId** | Pointer to **string** | Sub order ID | [optional] [default to ""]
**OrdType** | Pointer to **string** | Sub order type  &#x60;market&#x60;: Market order | [optional] [default to ""]
**Px** | Pointer to **string** | Sub order limit price  If it&#39;s a market order, \&quot;-1\&quot; will be return | [optional] [default to ""]
**Side** | Pointer to **string** | Sub order side  &#x60;buy&#x60; &#x60;sell&#x60; | [optional] [default to ""]
**State** | Pointer to **string** | Sub order state  &#x60;canceled&#x60;  &#x60;live&#x60;  &#x60;partially_filled&#x60;  &#x60;filled&#x60;  &#x60;cancelling&#x60; | [optional] [default to ""]
**Sz** | Pointer to **string** | Sub order quantity to buy or sell | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**TdMode** | Pointer to **string** | Sub order trade mode  Margin mode : &#x60;cross&#x60;  Non-Margin mode : &#x60;cash&#x60; | [optional] [default to ""]
**UTime** | Pointer to **string** | Sub order updated time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetTradingBotRecurringSubOrdersV5RespDataInner

`func NewGetTradingBotRecurringSubOrdersV5RespDataInner() *GetTradingBotRecurringSubOrdersV5RespDataInner`

NewGetTradingBotRecurringSubOrdersV5RespDataInner instantiates a new GetTradingBotRecurringSubOrdersV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotRecurringSubOrdersV5RespDataInnerWithDefaults

`func NewGetTradingBotRecurringSubOrdersV5RespDataInnerWithDefaults() *GetTradingBotRecurringSubOrdersV5RespDataInner`

NewGetTradingBotRecurringSubOrdersV5RespDataInnerWithDefaults instantiates a new GetTradingBotRecurringSubOrdersV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccFillSz

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetAccFillSz() string`

GetAccFillSz returns the AccFillSz field if non-nil, zero value otherwise.

### GetAccFillSzOk

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetAccFillSzOk() (*string, bool)`

GetAccFillSzOk returns a tuple with the AccFillSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccFillSz

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) SetAccFillSz(v string)`

SetAccFillSz sets AccFillSz field to given value.

### HasAccFillSz

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) HasAccFillSz() bool`

HasAccFillSz returns a boolean if a field has been set.

### GetAlgoClOrdId

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAlgoId

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetAlgoOrdType

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetAlgoOrdType() string`

GetAlgoOrdType returns the AlgoOrdType field if non-nil, zero value otherwise.

### GetAlgoOrdTypeOk

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetAlgoOrdTypeOk() (*string, bool)`

GetAlgoOrdTypeOk returns a tuple with the AlgoOrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoOrdType

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) SetAlgoOrdType(v string)`

SetAlgoOrdType sets AlgoOrdType field to given value.

### HasAlgoOrdType

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) HasAlgoOrdType() bool`

HasAlgoOrdType returns a boolean if a field has been set.

### GetAvgPx

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetAvgPx() string`

GetAvgPx returns the AvgPx field if non-nil, zero value otherwise.

### GetAvgPxOk

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetAvgPxOk() (*string, bool)`

GetAvgPxOk returns a tuple with the AvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPx

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) SetAvgPx(v string)`

SetAvgPx sets AvgPx field to given value.

### HasAvgPx

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) HasAvgPx() bool`

HasAvgPx returns a boolean if a field has been set.

### GetCTime

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetFee

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFeeCcy

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetFeeCcy() string`

GetFeeCcy returns the FeeCcy field if non-nil, zero value otherwise.

### GetFeeCcyOk

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetFeeCcyOk() (*string, bool)`

GetFeeCcyOk returns a tuple with the FeeCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCcy

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) SetFeeCcy(v string)`

SetFeeCcy sets FeeCcy field to given value.

### HasFeeCcy

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) HasFeeCcy() bool`

HasFeeCcy returns a boolean if a field has been set.

### GetInstId

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetOrdId

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetOrdType

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetOrdType() string`

GetOrdType returns the OrdType field if non-nil, zero value otherwise.

### GetOrdTypeOk

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetOrdTypeOk() (*string, bool)`

GetOrdTypeOk returns a tuple with the OrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdType

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) SetOrdType(v string)`

SetOrdType sets OrdType field to given value.

### HasOrdType

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) HasOrdType() bool`

HasOrdType returns a boolean if a field has been set.

### GetPx

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSide

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetState

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSz

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTag

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTdMode

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetTdMode() string`

GetTdMode returns the TdMode field if non-nil, zero value otherwise.

### GetTdModeOk

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetTdModeOk() (*string, bool)`

GetTdModeOk returns a tuple with the TdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdMode

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) SetTdMode(v string)`

SetTdMode sets TdMode field to given value.

### HasTdMode

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) HasTdMode() bool`

HasTdMode returns a boolean if a field has been set.

### GetUTime

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetTradingBotRecurringSubOrdersV5RespDataInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


