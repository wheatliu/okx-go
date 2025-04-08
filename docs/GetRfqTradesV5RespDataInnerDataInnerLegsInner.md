# GetRfqTradesV5RespDataInnerDataInnerLegsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fee** | Pointer to **string** | Fee. Negative number represents the user transaction fee charged by the platform. Positive number represents rebate. | [optional] [default to ""]
**FeeCcy** | Pointer to **string** | Fee currency | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60; | [optional] [default to ""]
**Px** | Pointer to **string** | The price the leg executed | [optional] [default to ""]
**Side** | Pointer to **string** | The direction of the leg. Valid value can be buy or sell. | [optional] [default to ""]
**Sz** | Pointer to **string** | Size of the leg in contracts or spot. | [optional] [default to ""]
**TradeId** | Pointer to **string** | Last traded ID. | [optional] [default to ""]

## Methods

### NewGetRfqTradesV5RespDataInnerDataInnerLegsInner

`func NewGetRfqTradesV5RespDataInnerDataInnerLegsInner() *GetRfqTradesV5RespDataInnerDataInnerLegsInner`

NewGetRfqTradesV5RespDataInnerDataInnerLegsInner instantiates a new GetRfqTradesV5RespDataInnerDataInnerLegsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRfqTradesV5RespDataInnerDataInnerLegsInnerWithDefaults

`func NewGetRfqTradesV5RespDataInnerDataInnerLegsInnerWithDefaults() *GetRfqTradesV5RespDataInnerDataInnerLegsInner`

NewGetRfqTradesV5RespDataInnerDataInnerLegsInnerWithDefaults instantiates a new GetRfqTradesV5RespDataInnerDataInnerLegsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFee

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFeeCcy

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) GetFeeCcy() string`

GetFeeCcy returns the FeeCcy field if non-nil, zero value otherwise.

### GetFeeCcyOk

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) GetFeeCcyOk() (*string, bool)`

GetFeeCcyOk returns a tuple with the FeeCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCcy

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) SetFeeCcy(v string)`

SetFeeCcy sets FeeCcy field to given value.

### HasFeeCcy

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) HasFeeCcy() bool`

HasFeeCcy returns a boolean if a field has been set.

### GetInstId

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetPx

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSide

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSz

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTradeId

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *GetRfqTradesV5RespDataInnerDataInnerLegsInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


