# GetSprdTradesV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client | [optional] [default to ""]
**Code** | Pointer to **string** | Error Code, the default is 0 | [optional] [default to ""]
**ExecType** | Pointer to **string** | Liquidity taker or maker, &#x60;T&#x60;: taker  &#x60;M&#x60;: maker | [optional] [default to ""]
**FillPx** | Pointer to **string** | Filled price | [optional] [default to ""]
**FillSz** | Pointer to **string** | Filled quantity | [optional] [default to ""]
**Legs** | Pointer to [**[]GetSprdTradesV5RespDataInnerLegsInner**](GetSprdTradesV5RespDataInnerLegsInner.md) | Legs of trade | [optional] 
**Msg** | Pointer to **string** | Error Message, the default is \&quot;\&quot; | [optional] [default to ""]
**OrdId** | Pointer to **string** | Order ID | [optional] [default to ""]
**Side** | Pointer to **string** | Order side, &#x60;buy&#x60; &#x60;sell&#x60; | [optional] [default to ""]
**SprdId** | Pointer to **string** | spread ID | [optional] [default to ""]
**State** | Pointer to **string** | Trade state.   Valid values are &#x60;filled&#x60; and &#x60;rejected&#x60; | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**TradeId** | Pointer to **string** | Trade ID | [optional] [default to ""]
**Ts** | Pointer to **string** | Data generation time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60;. | [optional] [default to ""]

## Methods

### NewGetSprdTradesV5RespDataInner

`func NewGetSprdTradesV5RespDataInner() *GetSprdTradesV5RespDataInner`

NewGetSprdTradesV5RespDataInner instantiates a new GetSprdTradesV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSprdTradesV5RespDataInnerWithDefaults

`func NewGetSprdTradesV5RespDataInnerWithDefaults() *GetSprdTradesV5RespDataInner`

NewGetSprdTradesV5RespDataInnerWithDefaults instantiates a new GetSprdTradesV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClOrdId

`func (o *GetSprdTradesV5RespDataInner) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *GetSprdTradesV5RespDataInner) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *GetSprdTradesV5RespDataInner) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *GetSprdTradesV5RespDataInner) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetCode

`func (o *GetSprdTradesV5RespDataInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetSprdTradesV5RespDataInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetSprdTradesV5RespDataInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetSprdTradesV5RespDataInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetExecType

`func (o *GetSprdTradesV5RespDataInner) GetExecType() string`

GetExecType returns the ExecType field if non-nil, zero value otherwise.

### GetExecTypeOk

`func (o *GetSprdTradesV5RespDataInner) GetExecTypeOk() (*string, bool)`

GetExecTypeOk returns a tuple with the ExecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecType

`func (o *GetSprdTradesV5RespDataInner) SetExecType(v string)`

SetExecType sets ExecType field to given value.

### HasExecType

`func (o *GetSprdTradesV5RespDataInner) HasExecType() bool`

HasExecType returns a boolean if a field has been set.

### GetFillPx

`func (o *GetSprdTradesV5RespDataInner) GetFillPx() string`

GetFillPx returns the FillPx field if non-nil, zero value otherwise.

### GetFillPxOk

`func (o *GetSprdTradesV5RespDataInner) GetFillPxOk() (*string, bool)`

GetFillPxOk returns a tuple with the FillPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPx

`func (o *GetSprdTradesV5RespDataInner) SetFillPx(v string)`

SetFillPx sets FillPx field to given value.

### HasFillPx

`func (o *GetSprdTradesV5RespDataInner) HasFillPx() bool`

HasFillPx returns a boolean if a field has been set.

### GetFillSz

`func (o *GetSprdTradesV5RespDataInner) GetFillSz() string`

GetFillSz returns the FillSz field if non-nil, zero value otherwise.

### GetFillSzOk

`func (o *GetSprdTradesV5RespDataInner) GetFillSzOk() (*string, bool)`

GetFillSzOk returns a tuple with the FillSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillSz

`func (o *GetSprdTradesV5RespDataInner) SetFillSz(v string)`

SetFillSz sets FillSz field to given value.

### HasFillSz

`func (o *GetSprdTradesV5RespDataInner) HasFillSz() bool`

HasFillSz returns a boolean if a field has been set.

### GetLegs

`func (o *GetSprdTradesV5RespDataInner) GetLegs() []GetSprdTradesV5RespDataInnerLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *GetSprdTradesV5RespDataInner) GetLegsOk() (*[]GetSprdTradesV5RespDataInnerLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *GetSprdTradesV5RespDataInner) SetLegs(v []GetSprdTradesV5RespDataInnerLegsInner)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *GetSprdTradesV5RespDataInner) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetMsg

`func (o *GetSprdTradesV5RespDataInner) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetSprdTradesV5RespDataInner) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetSprdTradesV5RespDataInner) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetSprdTradesV5RespDataInner) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetOrdId

`func (o *GetSprdTradesV5RespDataInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *GetSprdTradesV5RespDataInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *GetSprdTradesV5RespDataInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *GetSprdTradesV5RespDataInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetSide

`func (o *GetSprdTradesV5RespDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetSprdTradesV5RespDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetSprdTradesV5RespDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetSprdTradesV5RespDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSprdId

`func (o *GetSprdTradesV5RespDataInner) GetSprdId() string`

GetSprdId returns the SprdId field if non-nil, zero value otherwise.

### GetSprdIdOk

`func (o *GetSprdTradesV5RespDataInner) GetSprdIdOk() (*string, bool)`

GetSprdIdOk returns a tuple with the SprdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSprdId

`func (o *GetSprdTradesV5RespDataInner) SetSprdId(v string)`

SetSprdId sets SprdId field to given value.

### HasSprdId

`func (o *GetSprdTradesV5RespDataInner) HasSprdId() bool`

HasSprdId returns a boolean if a field has been set.

### GetState

`func (o *GetSprdTradesV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetSprdTradesV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetSprdTradesV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetSprdTradesV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTag

`func (o *GetSprdTradesV5RespDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetSprdTradesV5RespDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetSprdTradesV5RespDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetSprdTradesV5RespDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTradeId

`func (o *GetSprdTradesV5RespDataInner) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *GetSprdTradesV5RespDataInner) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *GetSprdTradesV5RespDataInner) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *GetSprdTradesV5RespDataInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.

### GetTs

`func (o *GetSprdTradesV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetSprdTradesV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetSprdTradesV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetSprdTradesV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


