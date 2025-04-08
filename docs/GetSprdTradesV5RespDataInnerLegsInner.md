# GetSprdTradesV5RespDataInnerLegsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fee** | Pointer to **string** | Fee. Negative number represents the user transaction fee charged by the platform. Positive number represents rebate. | [optional] [default to ""]
**FeeCcy** | Pointer to **string** | Fee currency | [optional] [default to ""]
**FillPnl** | Pointer to **string** | Last filled profit and loss, applicable to orders which have a trade and aim to close position. It always is 0 in other conditions | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. BTC-USDT-SWAP | [optional] [default to ""]
**Px** | Pointer to **string** | The price the leg executed | [optional] [default to ""]
**Side** | Pointer to **string** | The direction of the leg. Valid value can be &#x60;buy&#x60; or &#x60;sell&#x60;. | [optional] [default to ""]
**Sz** | Pointer to **string** | The size of each leg | [optional] [default to ""]
**SzCont** | Pointer to **string** | Filled amount of the contract    Only applicable to contracts, return \&quot;\&quot; for spot | [optional] [default to ""]
**TradeId** | Pointer to **string** | Traded ID in the OKX orderbook. | [optional] [default to ""]

## Methods

### NewGetSprdTradesV5RespDataInnerLegsInner

`func NewGetSprdTradesV5RespDataInnerLegsInner() *GetSprdTradesV5RespDataInnerLegsInner`

NewGetSprdTradesV5RespDataInnerLegsInner instantiates a new GetSprdTradesV5RespDataInnerLegsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSprdTradesV5RespDataInnerLegsInnerWithDefaults

`func NewGetSprdTradesV5RespDataInnerLegsInnerWithDefaults() *GetSprdTradesV5RespDataInnerLegsInner`

NewGetSprdTradesV5RespDataInnerLegsInnerWithDefaults instantiates a new GetSprdTradesV5RespDataInnerLegsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFee

`func (o *GetSprdTradesV5RespDataInnerLegsInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetSprdTradesV5RespDataInnerLegsInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetSprdTradesV5RespDataInnerLegsInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetSprdTradesV5RespDataInnerLegsInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFeeCcy

`func (o *GetSprdTradesV5RespDataInnerLegsInner) GetFeeCcy() string`

GetFeeCcy returns the FeeCcy field if non-nil, zero value otherwise.

### GetFeeCcyOk

`func (o *GetSprdTradesV5RespDataInnerLegsInner) GetFeeCcyOk() (*string, bool)`

GetFeeCcyOk returns a tuple with the FeeCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCcy

`func (o *GetSprdTradesV5RespDataInnerLegsInner) SetFeeCcy(v string)`

SetFeeCcy sets FeeCcy field to given value.

### HasFeeCcy

`func (o *GetSprdTradesV5RespDataInnerLegsInner) HasFeeCcy() bool`

HasFeeCcy returns a boolean if a field has been set.

### GetFillPnl

`func (o *GetSprdTradesV5RespDataInnerLegsInner) GetFillPnl() string`

GetFillPnl returns the FillPnl field if non-nil, zero value otherwise.

### GetFillPnlOk

`func (o *GetSprdTradesV5RespDataInnerLegsInner) GetFillPnlOk() (*string, bool)`

GetFillPnlOk returns a tuple with the FillPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPnl

`func (o *GetSprdTradesV5RespDataInnerLegsInner) SetFillPnl(v string)`

SetFillPnl sets FillPnl field to given value.

### HasFillPnl

`func (o *GetSprdTradesV5RespDataInnerLegsInner) HasFillPnl() bool`

HasFillPnl returns a boolean if a field has been set.

### GetInstId

`func (o *GetSprdTradesV5RespDataInnerLegsInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetSprdTradesV5RespDataInnerLegsInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetSprdTradesV5RespDataInnerLegsInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetSprdTradesV5RespDataInnerLegsInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetPx

`func (o *GetSprdTradesV5RespDataInnerLegsInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetSprdTradesV5RespDataInnerLegsInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetSprdTradesV5RespDataInnerLegsInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetSprdTradesV5RespDataInnerLegsInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSide

`func (o *GetSprdTradesV5RespDataInnerLegsInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetSprdTradesV5RespDataInnerLegsInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetSprdTradesV5RespDataInnerLegsInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetSprdTradesV5RespDataInnerLegsInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSz

`func (o *GetSprdTradesV5RespDataInnerLegsInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetSprdTradesV5RespDataInnerLegsInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetSprdTradesV5RespDataInnerLegsInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetSprdTradesV5RespDataInnerLegsInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetSzCont

`func (o *GetSprdTradesV5RespDataInnerLegsInner) GetSzCont() string`

GetSzCont returns the SzCont field if non-nil, zero value otherwise.

### GetSzContOk

`func (o *GetSprdTradesV5RespDataInnerLegsInner) GetSzContOk() (*string, bool)`

GetSzContOk returns a tuple with the SzCont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSzCont

`func (o *GetSprdTradesV5RespDataInnerLegsInner) SetSzCont(v string)`

SetSzCont sets SzCont field to given value.

### HasSzCont

`func (o *GetSprdTradesV5RespDataInnerLegsInner) HasSzCont() bool`

HasSzCont returns a boolean if a field has been set.

### GetTradeId

`func (o *GetSprdTradesV5RespDataInnerLegsInner) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *GetSprdTradesV5RespDataInnerLegsInner) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *GetSprdTradesV5RespDataInnerLegsInner) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *GetSprdTradesV5RespDataInnerLegsInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


