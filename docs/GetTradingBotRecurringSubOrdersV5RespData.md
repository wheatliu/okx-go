# GetTradingBotRecurringSubOrdersV5RespData

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

### NewGetTradingBotRecurringSubOrdersV5RespData

`func NewGetTradingBotRecurringSubOrdersV5RespData() *GetTradingBotRecurringSubOrdersV5RespData`

NewGetTradingBotRecurringSubOrdersV5RespData instantiates a new GetTradingBotRecurringSubOrdersV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotRecurringSubOrdersV5RespDataWithDefaults

`func NewGetTradingBotRecurringSubOrdersV5RespDataWithDefaults() *GetTradingBotRecurringSubOrdersV5RespData`

NewGetTradingBotRecurringSubOrdersV5RespDataWithDefaults instantiates a new GetTradingBotRecurringSubOrdersV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccFillSz

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetAccFillSz() string`

GetAccFillSz returns the AccFillSz field if non-nil, zero value otherwise.

### GetAccFillSzOk

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetAccFillSzOk() (*string, bool)`

GetAccFillSzOk returns a tuple with the AccFillSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccFillSz

`func (o *GetTradingBotRecurringSubOrdersV5RespData) SetAccFillSz(v string)`

SetAccFillSz sets AccFillSz field to given value.

### HasAccFillSz

`func (o *GetTradingBotRecurringSubOrdersV5RespData) HasAccFillSz() bool`

HasAccFillSz returns a boolean if a field has been set.

### GetAlgoClOrdId

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *GetTradingBotRecurringSubOrdersV5RespData) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *GetTradingBotRecurringSubOrdersV5RespData) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAlgoId

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetTradingBotRecurringSubOrdersV5RespData) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetTradingBotRecurringSubOrdersV5RespData) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetAlgoOrdType

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetAlgoOrdType() string`

GetAlgoOrdType returns the AlgoOrdType field if non-nil, zero value otherwise.

### GetAlgoOrdTypeOk

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetAlgoOrdTypeOk() (*string, bool)`

GetAlgoOrdTypeOk returns a tuple with the AlgoOrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoOrdType

`func (o *GetTradingBotRecurringSubOrdersV5RespData) SetAlgoOrdType(v string)`

SetAlgoOrdType sets AlgoOrdType field to given value.

### HasAlgoOrdType

`func (o *GetTradingBotRecurringSubOrdersV5RespData) HasAlgoOrdType() bool`

HasAlgoOrdType returns a boolean if a field has been set.

### GetAvgPx

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetAvgPx() string`

GetAvgPx returns the AvgPx field if non-nil, zero value otherwise.

### GetAvgPxOk

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetAvgPxOk() (*string, bool)`

GetAvgPxOk returns a tuple with the AvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPx

`func (o *GetTradingBotRecurringSubOrdersV5RespData) SetAvgPx(v string)`

SetAvgPx sets AvgPx field to given value.

### HasAvgPx

`func (o *GetTradingBotRecurringSubOrdersV5RespData) HasAvgPx() bool`

HasAvgPx returns a boolean if a field has been set.

### GetCTime

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetTradingBotRecurringSubOrdersV5RespData) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetTradingBotRecurringSubOrdersV5RespData) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetFee

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetTradingBotRecurringSubOrdersV5RespData) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetTradingBotRecurringSubOrdersV5RespData) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFeeCcy

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetFeeCcy() string`

GetFeeCcy returns the FeeCcy field if non-nil, zero value otherwise.

### GetFeeCcyOk

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetFeeCcyOk() (*string, bool)`

GetFeeCcyOk returns a tuple with the FeeCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCcy

`func (o *GetTradingBotRecurringSubOrdersV5RespData) SetFeeCcy(v string)`

SetFeeCcy sets FeeCcy field to given value.

### HasFeeCcy

`func (o *GetTradingBotRecurringSubOrdersV5RespData) HasFeeCcy() bool`

HasFeeCcy returns a boolean if a field has been set.

### GetInstId

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetTradingBotRecurringSubOrdersV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetTradingBotRecurringSubOrdersV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetTradingBotRecurringSubOrdersV5RespData) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetTradingBotRecurringSubOrdersV5RespData) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetOrdId

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *GetTradingBotRecurringSubOrdersV5RespData) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *GetTradingBotRecurringSubOrdersV5RespData) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetOrdType

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetOrdType() string`

GetOrdType returns the OrdType field if non-nil, zero value otherwise.

### GetOrdTypeOk

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetOrdTypeOk() (*string, bool)`

GetOrdTypeOk returns a tuple with the OrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdType

`func (o *GetTradingBotRecurringSubOrdersV5RespData) SetOrdType(v string)`

SetOrdType sets OrdType field to given value.

### HasOrdType

`func (o *GetTradingBotRecurringSubOrdersV5RespData) HasOrdType() bool`

HasOrdType returns a boolean if a field has been set.

### GetPx

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetTradingBotRecurringSubOrdersV5RespData) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetTradingBotRecurringSubOrdersV5RespData) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSide

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetTradingBotRecurringSubOrdersV5RespData) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetTradingBotRecurringSubOrdersV5RespData) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetState

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetTradingBotRecurringSubOrdersV5RespData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetTradingBotRecurringSubOrdersV5RespData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSz

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetTradingBotRecurringSubOrdersV5RespData) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetTradingBotRecurringSubOrdersV5RespData) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTag

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetTradingBotRecurringSubOrdersV5RespData) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetTradingBotRecurringSubOrdersV5RespData) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTdMode

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetTdMode() string`

GetTdMode returns the TdMode field if non-nil, zero value otherwise.

### GetTdModeOk

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetTdModeOk() (*string, bool)`

GetTdModeOk returns a tuple with the TdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdMode

`func (o *GetTradingBotRecurringSubOrdersV5RespData) SetTdMode(v string)`

SetTdMode sets TdMode field to given value.

### HasTdMode

`func (o *GetTradingBotRecurringSubOrdersV5RespData) HasTdMode() bool`

HasTdMode returns a boolean if a field has been set.

### GetUTime

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetTradingBotRecurringSubOrdersV5RespData) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetTradingBotRecurringSubOrdersV5RespData) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetTradingBotRecurringSubOrdersV5RespData) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


