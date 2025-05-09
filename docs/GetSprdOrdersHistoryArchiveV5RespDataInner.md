# GetSprdOrdersHistoryArchiveV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccFillSz** | Pointer to **string** | Accumulated fill quantity | [optional] [default to ""]
**AvgPx** | Pointer to **string** | Average filled price. If none is filled, it will return \&quot;0\&quot;. | [optional] [default to ""]
**CTime** | Pointer to **string** | Creation time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**CancelSource** | Pointer to **string** | Source of the order cancellation. Valid values and the corresponding meanings are:   &#x60;0&#x60;: Order canceled by system   &#x60;1&#x60;: Order canceled by user   &#x60;14&#x60;: Order canceled: IOC order was partially canceled due to incompletely filled  &#x60;15&#x60;: Order canceled: The order price is beyond the limit   &#x60;20&#x60;: Cancel all after triggered   &#x60;31&#x60;: The post-only order will take liquidity in maker orders   &#x60;32&#x60;: Self trade prevention  &#x60;34&#x60;: Order failed to settle due to insufficient margin   &#x60;35&#x60;: Order cancellation due to insufficient margin from another order | [optional] [default to ""]
**CanceledSz** | Pointer to **string** | Quantity canceled due order cancellations or trade rejections | [optional] [default to ""]
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client | [optional] [default to ""]
**FillPx** | Pointer to **string** | Last fill price | [optional] [default to ""]
**FillSz** | Pointer to **string** | Last fill quantity | [optional] [default to ""]
**OrdId** | Pointer to **string** | Order ID | [optional] [default to ""]
**OrdType** | Pointer to **string** | Order type  &#x60;market&#x60;: Market order   &#x60;limit&#x60;: limit order   &#x60;post_only&#x60;: Post-only order   &#x60;ioc&#x60;: Immediate-or-cancel order | [optional] [default to ""]
**PendingFillSz** | Pointer to **string** | Quantity still remaining to be filled, inluding pendingSettleSz | [optional] [default to ""]
**PendingSettleSz** | Pointer to **string** | Quantity that&#39;s pending settlement | [optional] [default to ""]
**Px** | Pointer to **string** | Price | [optional] [default to ""]
**Side** | Pointer to **string** | Order side | [optional] [default to ""]
**SprdId** | Pointer to **string** | spread ID | [optional] [default to ""]
**State** | Pointer to **string** | State   &#x60;canceled&#x60;   &#x60;filled&#x60; | [optional] [default to ""]
**Sz** | Pointer to **string** | Quantity to buy or sell | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**TradeId** | Pointer to **string** | Last trade ID | [optional] [default to ""]
**UTime** | Pointer to **string** | Update time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetSprdOrdersHistoryArchiveV5RespDataInner

`func NewGetSprdOrdersHistoryArchiveV5RespDataInner() *GetSprdOrdersHistoryArchiveV5RespDataInner`

NewGetSprdOrdersHistoryArchiveV5RespDataInner instantiates a new GetSprdOrdersHistoryArchiveV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSprdOrdersHistoryArchiveV5RespDataInnerWithDefaults

`func NewGetSprdOrdersHistoryArchiveV5RespDataInnerWithDefaults() *GetSprdOrdersHistoryArchiveV5RespDataInner`

NewGetSprdOrdersHistoryArchiveV5RespDataInnerWithDefaults instantiates a new GetSprdOrdersHistoryArchiveV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccFillSz

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetAccFillSz() string`

GetAccFillSz returns the AccFillSz field if non-nil, zero value otherwise.

### GetAccFillSzOk

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetAccFillSzOk() (*string, bool)`

GetAccFillSzOk returns a tuple with the AccFillSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccFillSz

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) SetAccFillSz(v string)`

SetAccFillSz sets AccFillSz field to given value.

### HasAccFillSz

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) HasAccFillSz() bool`

HasAccFillSz returns a boolean if a field has been set.

### GetAvgPx

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetAvgPx() string`

GetAvgPx returns the AvgPx field if non-nil, zero value otherwise.

### GetAvgPxOk

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetAvgPxOk() (*string, bool)`

GetAvgPxOk returns a tuple with the AvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPx

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) SetAvgPx(v string)`

SetAvgPx sets AvgPx field to given value.

### HasAvgPx

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) HasAvgPx() bool`

HasAvgPx returns a boolean if a field has been set.

### GetCTime

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCancelSource

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetCancelSource() string`

GetCancelSource returns the CancelSource field if non-nil, zero value otherwise.

### GetCancelSourceOk

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetCancelSourceOk() (*string, bool)`

GetCancelSourceOk returns a tuple with the CancelSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelSource

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) SetCancelSource(v string)`

SetCancelSource sets CancelSource field to given value.

### HasCancelSource

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) HasCancelSource() bool`

HasCancelSource returns a boolean if a field has been set.

### GetCanceledSz

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetCanceledSz() string`

GetCanceledSz returns the CanceledSz field if non-nil, zero value otherwise.

### GetCanceledSzOk

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetCanceledSzOk() (*string, bool)`

GetCanceledSzOk returns a tuple with the CanceledSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledSz

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) SetCanceledSz(v string)`

SetCanceledSz sets CanceledSz field to given value.

### HasCanceledSz

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) HasCanceledSz() bool`

HasCanceledSz returns a boolean if a field has been set.

### GetClOrdId

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetFillPx

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetFillPx() string`

GetFillPx returns the FillPx field if non-nil, zero value otherwise.

### GetFillPxOk

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetFillPxOk() (*string, bool)`

GetFillPxOk returns a tuple with the FillPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPx

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) SetFillPx(v string)`

SetFillPx sets FillPx field to given value.

### HasFillPx

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) HasFillPx() bool`

HasFillPx returns a boolean if a field has been set.

### GetFillSz

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetFillSz() string`

GetFillSz returns the FillSz field if non-nil, zero value otherwise.

### GetFillSzOk

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetFillSzOk() (*string, bool)`

GetFillSzOk returns a tuple with the FillSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillSz

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) SetFillSz(v string)`

SetFillSz sets FillSz field to given value.

### HasFillSz

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) HasFillSz() bool`

HasFillSz returns a boolean if a field has been set.

### GetOrdId

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetOrdType

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetOrdType() string`

GetOrdType returns the OrdType field if non-nil, zero value otherwise.

### GetOrdTypeOk

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetOrdTypeOk() (*string, bool)`

GetOrdTypeOk returns a tuple with the OrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdType

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) SetOrdType(v string)`

SetOrdType sets OrdType field to given value.

### HasOrdType

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) HasOrdType() bool`

HasOrdType returns a boolean if a field has been set.

### GetPendingFillSz

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetPendingFillSz() string`

GetPendingFillSz returns the PendingFillSz field if non-nil, zero value otherwise.

### GetPendingFillSzOk

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetPendingFillSzOk() (*string, bool)`

GetPendingFillSzOk returns a tuple with the PendingFillSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingFillSz

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) SetPendingFillSz(v string)`

SetPendingFillSz sets PendingFillSz field to given value.

### HasPendingFillSz

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) HasPendingFillSz() bool`

HasPendingFillSz returns a boolean if a field has been set.

### GetPendingSettleSz

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetPendingSettleSz() string`

GetPendingSettleSz returns the PendingSettleSz field if non-nil, zero value otherwise.

### GetPendingSettleSzOk

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetPendingSettleSzOk() (*string, bool)`

GetPendingSettleSzOk returns a tuple with the PendingSettleSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingSettleSz

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) SetPendingSettleSz(v string)`

SetPendingSettleSz sets PendingSettleSz field to given value.

### HasPendingSettleSz

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) HasPendingSettleSz() bool`

HasPendingSettleSz returns a boolean if a field has been set.

### GetPx

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSide

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSprdId

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetSprdId() string`

GetSprdId returns the SprdId field if non-nil, zero value otherwise.

### GetSprdIdOk

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetSprdIdOk() (*string, bool)`

GetSprdIdOk returns a tuple with the SprdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSprdId

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) SetSprdId(v string)`

SetSprdId sets SprdId field to given value.

### HasSprdId

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) HasSprdId() bool`

HasSprdId returns a boolean if a field has been set.

### GetState

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSz

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTag

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTradeId

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.

### GetUTime

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetSprdOrdersHistoryArchiveV5RespDataInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


