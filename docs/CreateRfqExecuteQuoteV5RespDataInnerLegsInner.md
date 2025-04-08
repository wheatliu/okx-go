# CreateRfqExecuteQuoteV5RespDataInnerLegsInner

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

### NewCreateRfqExecuteQuoteV5RespDataInnerLegsInner

`func NewCreateRfqExecuteQuoteV5RespDataInnerLegsInner() *CreateRfqExecuteQuoteV5RespDataInnerLegsInner`

NewCreateRfqExecuteQuoteV5RespDataInnerLegsInner instantiates a new CreateRfqExecuteQuoteV5RespDataInnerLegsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqExecuteQuoteV5RespDataInnerLegsInnerWithDefaults

`func NewCreateRfqExecuteQuoteV5RespDataInnerLegsInnerWithDefaults() *CreateRfqExecuteQuoteV5RespDataInnerLegsInner`

NewCreateRfqExecuteQuoteV5RespDataInnerLegsInnerWithDefaults instantiates a new CreateRfqExecuteQuoteV5RespDataInnerLegsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFee

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFeeCcy

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetFeeCcy() string`

GetFeeCcy returns the FeeCcy field if non-nil, zero value otherwise.

### GetFeeCcyOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetFeeCcyOk() (*string, bool)`

GetFeeCcyOk returns a tuple with the FeeCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCcy

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) SetFeeCcy(v string)`

SetFeeCcy sets FeeCcy field to given value.

### HasFeeCcy

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) HasFeeCcy() bool`

HasFeeCcy returns a boolean if a field has been set.

### GetInstId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetPx

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSide

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSz

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTradeId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


