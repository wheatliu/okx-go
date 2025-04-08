# CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fee** | Pointer to **string** | Fee for the individual leg.   Negative fee represents the user transaction fee charged by the platform. Positive fee represents rebate. | [optional] [default to ""]
**FeeCcy** | Pointer to **string** | Fee currency. To be read in conjunction with fee | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. BTC-USDT-SWAP | [optional] [default to ""]
**Px** | Pointer to **string** | The price the leg executed | [optional] [default to ""]
**Side** | Pointer to **string** | The direction of the leg from the Takers perspective. Valid value can be buy or sell. | [optional] [default to ""]
**Sz** | Pointer to **string** | Size of the leg in contracts or spot. | [optional] [default to ""]
**TradeId** | Pointer to **string** | Last traded ID. | [optional] [default to ""]

## Methods

### NewCreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner

`func NewCreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner() *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner`

NewCreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner instantiates a new CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInnerWithDefaults

`func NewCreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInnerWithDefaults() *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner`

NewCreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInnerWithDefaults instantiates a new CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFee

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFeeCcy

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) GetFeeCcy() string`

GetFeeCcy returns the FeeCcy field if non-nil, zero value otherwise.

### GetFeeCcyOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) GetFeeCcyOk() (*string, bool)`

GetFeeCcyOk returns a tuple with the FeeCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCcy

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) SetFeeCcy(v string)`

SetFeeCcy sets FeeCcy field to given value.

### HasFeeCcy

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) HasFeeCcy() bool`

HasFeeCcy returns a boolean if a field has been set.

### GetInstId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetPx

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSide

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSz

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTradeId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


