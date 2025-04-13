# GetSprdOrdersPendingV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccFillSz** | Pointer to **string** | Accumulated fill quantity | [optional] [default to ""]
**AvgPx** | Pointer to **string** | Average filled price. If none is filled, it will return \&quot;0\&quot;. | [optional] [default to ""]
**CTime** | Pointer to **string** | Creation time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**CancelSource** | Pointer to **string** | Source of the order cancellation.Valid values and the corresponding meanings are:   &#x60;0&#x60;: Order canceled by system   &#x60;1&#x60;: Order canceled by user   &#x60;14&#x60;: Order canceled: IOC order was partially canceled due to incompletely filled  &#x60;15&#x60;: Order canceled: The order price is beyond the limit   &#x60;20&#x60;: Cancel all after triggered   &#x60;31&#x60;: The post-only order will take liquidity in maker orders   &#x60;32&#x60;: Self trade prevention   &#x60;34&#x60;: Order failed to settle due to insufficient margin   &#x60;35&#x60;: Order cancellation due to insufficient margin from another order | [optional] [default to ""]
**CanceledSz** | Pointer to **string** | Quantity canceled due order cancellations or trade rejections | [optional] [default to ""]
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client | [optional] [default to ""]
**FillPx** | Pointer to **string** | Last fill price | [optional] [default to ""]
**FillSz** | Pointer to **string** | Last fill quantity | [optional] [default to ""]
**OrdId** | Pointer to **string** | Order ID | [optional] [default to ""]
**OrdType** | Pointer to **string** | Order type  &#x60;market&#x60;: Market order   &#x60;limit&#x60;: Limit order   &#x60;post_only&#x60;: Post-only order   &#x60;ioc&#x60;: Immediate-or-cancel order | [optional] [default to ""]
**PendingFillSz** | Pointer to **string** | Quantity still remaining to be filled | [optional] [default to ""]
**PendingSettleSz** | Pointer to **string** | Quantity that&#39;s pending settlement | [optional] [default to ""]
**Px** | Pointer to **string** | Price | [optional] [default to ""]
**Side** | Pointer to **string** | Order side | [optional] [default to ""]
**SprdId** | Pointer to **string** | spread ID | [optional] [default to ""]
**State** | Pointer to **string** | State   &#x60;live&#x60;   &#x60;partially_filled&#x60; | [optional] [default to ""]
**Sz** | Pointer to **string** | Quantity to buy or sell | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**TradeId** | Pointer to **string** | Last trade ID | [optional] [default to ""]
**UTime** | Pointer to **string** | Update time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetSprdOrdersPendingV5RespDataInner

`func NewGetSprdOrdersPendingV5RespDataInner() *GetSprdOrdersPendingV5RespDataInner`

NewGetSprdOrdersPendingV5RespDataInner instantiates a new GetSprdOrdersPendingV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSprdOrdersPendingV5RespDataInnerWithDefaults

`func NewGetSprdOrdersPendingV5RespDataInnerWithDefaults() *GetSprdOrdersPendingV5RespDataInner`

NewGetSprdOrdersPendingV5RespDataInnerWithDefaults instantiates a new GetSprdOrdersPendingV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccFillSz

`func (o *GetSprdOrdersPendingV5RespDataInner) GetAccFillSz() string`

GetAccFillSz returns the AccFillSz field if non-nil, zero value otherwise.

### GetAccFillSzOk

`func (o *GetSprdOrdersPendingV5RespDataInner) GetAccFillSzOk() (*string, bool)`

GetAccFillSzOk returns a tuple with the AccFillSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccFillSz

`func (o *GetSprdOrdersPendingV5RespDataInner) SetAccFillSz(v string)`

SetAccFillSz sets AccFillSz field to given value.

### HasAccFillSz

`func (o *GetSprdOrdersPendingV5RespDataInner) HasAccFillSz() bool`

HasAccFillSz returns a boolean if a field has been set.

### GetAvgPx

`func (o *GetSprdOrdersPendingV5RespDataInner) GetAvgPx() string`

GetAvgPx returns the AvgPx field if non-nil, zero value otherwise.

### GetAvgPxOk

`func (o *GetSprdOrdersPendingV5RespDataInner) GetAvgPxOk() (*string, bool)`

GetAvgPxOk returns a tuple with the AvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPx

`func (o *GetSprdOrdersPendingV5RespDataInner) SetAvgPx(v string)`

SetAvgPx sets AvgPx field to given value.

### HasAvgPx

`func (o *GetSprdOrdersPendingV5RespDataInner) HasAvgPx() bool`

HasAvgPx returns a boolean if a field has been set.

### GetCTime

`func (o *GetSprdOrdersPendingV5RespDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetSprdOrdersPendingV5RespDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetSprdOrdersPendingV5RespDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetSprdOrdersPendingV5RespDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCancelSource

`func (o *GetSprdOrdersPendingV5RespDataInner) GetCancelSource() string`

GetCancelSource returns the CancelSource field if non-nil, zero value otherwise.

### GetCancelSourceOk

`func (o *GetSprdOrdersPendingV5RespDataInner) GetCancelSourceOk() (*string, bool)`

GetCancelSourceOk returns a tuple with the CancelSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelSource

`func (o *GetSprdOrdersPendingV5RespDataInner) SetCancelSource(v string)`

SetCancelSource sets CancelSource field to given value.

### HasCancelSource

`func (o *GetSprdOrdersPendingV5RespDataInner) HasCancelSource() bool`

HasCancelSource returns a boolean if a field has been set.

### GetCanceledSz

`func (o *GetSprdOrdersPendingV5RespDataInner) GetCanceledSz() string`

GetCanceledSz returns the CanceledSz field if non-nil, zero value otherwise.

### GetCanceledSzOk

`func (o *GetSprdOrdersPendingV5RespDataInner) GetCanceledSzOk() (*string, bool)`

GetCanceledSzOk returns a tuple with the CanceledSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledSz

`func (o *GetSprdOrdersPendingV5RespDataInner) SetCanceledSz(v string)`

SetCanceledSz sets CanceledSz field to given value.

### HasCanceledSz

`func (o *GetSprdOrdersPendingV5RespDataInner) HasCanceledSz() bool`

HasCanceledSz returns a boolean if a field has been set.

### GetClOrdId

`func (o *GetSprdOrdersPendingV5RespDataInner) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *GetSprdOrdersPendingV5RespDataInner) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *GetSprdOrdersPendingV5RespDataInner) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *GetSprdOrdersPendingV5RespDataInner) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetFillPx

`func (o *GetSprdOrdersPendingV5RespDataInner) GetFillPx() string`

GetFillPx returns the FillPx field if non-nil, zero value otherwise.

### GetFillPxOk

`func (o *GetSprdOrdersPendingV5RespDataInner) GetFillPxOk() (*string, bool)`

GetFillPxOk returns a tuple with the FillPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPx

`func (o *GetSprdOrdersPendingV5RespDataInner) SetFillPx(v string)`

SetFillPx sets FillPx field to given value.

### HasFillPx

`func (o *GetSprdOrdersPendingV5RespDataInner) HasFillPx() bool`

HasFillPx returns a boolean if a field has been set.

### GetFillSz

`func (o *GetSprdOrdersPendingV5RespDataInner) GetFillSz() string`

GetFillSz returns the FillSz field if non-nil, zero value otherwise.

### GetFillSzOk

`func (o *GetSprdOrdersPendingV5RespDataInner) GetFillSzOk() (*string, bool)`

GetFillSzOk returns a tuple with the FillSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillSz

`func (o *GetSprdOrdersPendingV5RespDataInner) SetFillSz(v string)`

SetFillSz sets FillSz field to given value.

### HasFillSz

`func (o *GetSprdOrdersPendingV5RespDataInner) HasFillSz() bool`

HasFillSz returns a boolean if a field has been set.

### GetOrdId

`func (o *GetSprdOrdersPendingV5RespDataInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *GetSprdOrdersPendingV5RespDataInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *GetSprdOrdersPendingV5RespDataInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *GetSprdOrdersPendingV5RespDataInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetOrdType

`func (o *GetSprdOrdersPendingV5RespDataInner) GetOrdType() string`

GetOrdType returns the OrdType field if non-nil, zero value otherwise.

### GetOrdTypeOk

`func (o *GetSprdOrdersPendingV5RespDataInner) GetOrdTypeOk() (*string, bool)`

GetOrdTypeOk returns a tuple with the OrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdType

`func (o *GetSprdOrdersPendingV5RespDataInner) SetOrdType(v string)`

SetOrdType sets OrdType field to given value.

### HasOrdType

`func (o *GetSprdOrdersPendingV5RespDataInner) HasOrdType() bool`

HasOrdType returns a boolean if a field has been set.

### GetPendingFillSz

`func (o *GetSprdOrdersPendingV5RespDataInner) GetPendingFillSz() string`

GetPendingFillSz returns the PendingFillSz field if non-nil, zero value otherwise.

### GetPendingFillSzOk

`func (o *GetSprdOrdersPendingV5RespDataInner) GetPendingFillSzOk() (*string, bool)`

GetPendingFillSzOk returns a tuple with the PendingFillSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingFillSz

`func (o *GetSprdOrdersPendingV5RespDataInner) SetPendingFillSz(v string)`

SetPendingFillSz sets PendingFillSz field to given value.

### HasPendingFillSz

`func (o *GetSprdOrdersPendingV5RespDataInner) HasPendingFillSz() bool`

HasPendingFillSz returns a boolean if a field has been set.

### GetPendingSettleSz

`func (o *GetSprdOrdersPendingV5RespDataInner) GetPendingSettleSz() string`

GetPendingSettleSz returns the PendingSettleSz field if non-nil, zero value otherwise.

### GetPendingSettleSzOk

`func (o *GetSprdOrdersPendingV5RespDataInner) GetPendingSettleSzOk() (*string, bool)`

GetPendingSettleSzOk returns a tuple with the PendingSettleSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingSettleSz

`func (o *GetSprdOrdersPendingV5RespDataInner) SetPendingSettleSz(v string)`

SetPendingSettleSz sets PendingSettleSz field to given value.

### HasPendingSettleSz

`func (o *GetSprdOrdersPendingV5RespDataInner) HasPendingSettleSz() bool`

HasPendingSettleSz returns a boolean if a field has been set.

### GetPx

`func (o *GetSprdOrdersPendingV5RespDataInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetSprdOrdersPendingV5RespDataInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetSprdOrdersPendingV5RespDataInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetSprdOrdersPendingV5RespDataInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSide

`func (o *GetSprdOrdersPendingV5RespDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetSprdOrdersPendingV5RespDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetSprdOrdersPendingV5RespDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetSprdOrdersPendingV5RespDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSprdId

`func (o *GetSprdOrdersPendingV5RespDataInner) GetSprdId() string`

GetSprdId returns the SprdId field if non-nil, zero value otherwise.

### GetSprdIdOk

`func (o *GetSprdOrdersPendingV5RespDataInner) GetSprdIdOk() (*string, bool)`

GetSprdIdOk returns a tuple with the SprdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSprdId

`func (o *GetSprdOrdersPendingV5RespDataInner) SetSprdId(v string)`

SetSprdId sets SprdId field to given value.

### HasSprdId

`func (o *GetSprdOrdersPendingV5RespDataInner) HasSprdId() bool`

HasSprdId returns a boolean if a field has been set.

### GetState

`func (o *GetSprdOrdersPendingV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetSprdOrdersPendingV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetSprdOrdersPendingV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetSprdOrdersPendingV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSz

`func (o *GetSprdOrdersPendingV5RespDataInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetSprdOrdersPendingV5RespDataInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetSprdOrdersPendingV5RespDataInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetSprdOrdersPendingV5RespDataInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTag

`func (o *GetSprdOrdersPendingV5RespDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetSprdOrdersPendingV5RespDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetSprdOrdersPendingV5RespDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetSprdOrdersPendingV5RespDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTradeId

`func (o *GetSprdOrdersPendingV5RespDataInner) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *GetSprdOrdersPendingV5RespDataInner) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *GetSprdOrdersPendingV5RespDataInner) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *GetSprdOrdersPendingV5RespDataInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.

### GetUTime

`func (o *GetSprdOrdersPendingV5RespDataInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetSprdOrdersPendingV5RespDataInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetSprdOrdersPendingV5RespDataInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetSprdOrdersPendingV5RespDataInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


