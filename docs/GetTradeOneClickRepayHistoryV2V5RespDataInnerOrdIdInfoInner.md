# GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CTime** | Pointer to **string** | Creation time for order, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**FillPx** | Pointer to **string** | Last filled price.  If none is filled, it will return \&quot;\&quot;. | [optional] [default to ""]
**FillSz** | Pointer to **string** | Last filled quantity | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [optional] [default to ""]
**OrdId** | Pointer to **string** | Order ID | [optional] [default to ""]
**OrdType** | Pointer to **string** | Order type  &#x60;ioc&#x60;: Immediate-or-cancel order | [optional] [default to ""]
**Px** | Pointer to **string** | Price | [optional] [default to ""]
**Side** | Pointer to **string** | Side  &#x60;buy&#x60;  &#x60;sell&#x60; | [optional] [default to ""]
**State** | Pointer to **string** | State  &#x60;filled&#x60;  &#x60;canceled&#x60; | [optional] [default to ""]
**Sz** | Pointer to **string** | Quantity to buy or sell | [optional] [default to ""]

## Methods

### NewGetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner

`func NewGetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner() *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner`

NewGetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner instantiates a new GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInnerWithDefaults

`func NewGetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInnerWithDefaults() *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner`

NewGetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInnerWithDefaults instantiates a new GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCTime

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetFillPx

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) GetFillPx() string`

GetFillPx returns the FillPx field if non-nil, zero value otherwise.

### GetFillPxOk

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) GetFillPxOk() (*string, bool)`

GetFillPxOk returns a tuple with the FillPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPx

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) SetFillPx(v string)`

SetFillPx sets FillPx field to given value.

### HasFillPx

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) HasFillPx() bool`

HasFillPx returns a boolean if a field has been set.

### GetFillSz

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) GetFillSz() string`

GetFillSz returns the FillSz field if non-nil, zero value otherwise.

### GetFillSzOk

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) GetFillSzOk() (*string, bool)`

GetFillSzOk returns a tuple with the FillSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillSz

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) SetFillSz(v string)`

SetFillSz sets FillSz field to given value.

### HasFillSz

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) HasFillSz() bool`

HasFillSz returns a boolean if a field has been set.

### GetInstId

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetOrdId

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetOrdType

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) GetOrdType() string`

GetOrdType returns the OrdType field if non-nil, zero value otherwise.

### GetOrdTypeOk

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) GetOrdTypeOk() (*string, bool)`

GetOrdTypeOk returns a tuple with the OrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdType

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) SetOrdType(v string)`

SetOrdType sets OrdType field to given value.

### HasOrdType

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) HasOrdType() bool`

HasOrdType returns a boolean if a field has been set.

### GetPx

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSide

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetState

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSz

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetTradeOneClickRepayHistoryV2V5RespDataInnerOrdIdInfoInner) HasSz() bool`

HasSz returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


