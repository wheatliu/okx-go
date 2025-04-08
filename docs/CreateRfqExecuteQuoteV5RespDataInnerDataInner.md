# CreateRfqExecuteQuoteV5RespDataInnerDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockTdId** | Pointer to **string** | Block trade ID. | [optional] [default to ""]
**CTime** | Pointer to **string** | The execution time for the trade. Unix timestamp in milliseconds. | [optional] [default to ""]
**ClQuoteId** | Pointer to **string** | Client-supplied Quote ID. This attribute is treated as client sensitive information. It will not be exposed to the Taker, only return empty string. | [optional] [default to ""]
**ClRfqId** | Pointer to **string** | Client-supplied RFQ ID. This attribute is treated as client sensitive information. It will not be exposed to the Maker, only return empty string. | [optional] [default to ""]
**Fee** | Pointer to **string** | Fee for the individual leg.   Negative fee represents the user transaction fee charged by the platform. Positive fee represents rebate. | [optional] [default to ""]
**FeeCcy** | Pointer to **string** | Fee currency. To be read in conjunction with fee | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. BTC-USDT-SWAP | [optional] [default to ""]
**Legs** | Pointer to [**[]CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner**](CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner.md) | Legs of trade | [optional] 
**MTraderCode** | Pointer to **string** | A unique identifier of the maker. Empty if the anonymous parameter of the Quote is set to be &#x60;true&#x60;. | [optional] [default to ""]
**Px** | Pointer to **string** | The price the leg executed | [optional] [default to ""]
**QuoteId** | Pointer to **string** | Quote ID. | [optional] [default to ""]
**RfqId** | Pointer to **string** | RFQ ID. | [optional] [default to ""]
**Side** | Pointer to **string** | The direction of the leg from the Takers perspective. Valid value can be buy or sell. | [optional] [default to ""]
**Sz** | Pointer to **string** | Size of the leg in contracts or spot. | [optional] [default to ""]
**TTraderCode** | Pointer to **string** | A unique identifier of the taker. Empty if the anonymous parameter of the RFQ is set to be &#x60;true&#x60;. | [optional] [default to ""]
**Tag** | Pointer to **string** | RFQ tag. | [optional] [default to ""]
**TradeId** | Pointer to **string** | Last traded ID. | [optional] [default to ""]

## Methods

### NewCreateRfqExecuteQuoteV5RespDataInnerDataInner

`func NewCreateRfqExecuteQuoteV5RespDataInnerDataInner() *CreateRfqExecuteQuoteV5RespDataInnerDataInner`

NewCreateRfqExecuteQuoteV5RespDataInnerDataInner instantiates a new CreateRfqExecuteQuoteV5RespDataInnerDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqExecuteQuoteV5RespDataInnerDataInnerWithDefaults

`func NewCreateRfqExecuteQuoteV5RespDataInnerDataInnerWithDefaults() *CreateRfqExecuteQuoteV5RespDataInnerDataInner`

NewCreateRfqExecuteQuoteV5RespDataInnerDataInnerWithDefaults instantiates a new CreateRfqExecuteQuoteV5RespDataInnerDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockTdId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetBlockTdId() string`

GetBlockTdId returns the BlockTdId field if non-nil, zero value otherwise.

### GetBlockTdIdOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetBlockTdIdOk() (*string, bool)`

GetBlockTdIdOk returns a tuple with the BlockTdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTdId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) SetBlockTdId(v string)`

SetBlockTdId sets BlockTdId field to given value.

### HasBlockTdId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) HasBlockTdId() bool`

HasBlockTdId returns a boolean if a field has been set.

### GetCTime

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetClQuoteId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetClQuoteId() string`

GetClQuoteId returns the ClQuoteId field if non-nil, zero value otherwise.

### GetClQuoteIdOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetClQuoteIdOk() (*string, bool)`

GetClQuoteIdOk returns a tuple with the ClQuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClQuoteId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) SetClQuoteId(v string)`

SetClQuoteId sets ClQuoteId field to given value.

### HasClQuoteId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) HasClQuoteId() bool`

HasClQuoteId returns a boolean if a field has been set.

### GetClRfqId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetClRfqId() string`

GetClRfqId returns the ClRfqId field if non-nil, zero value otherwise.

### GetClRfqIdOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetClRfqIdOk() (*string, bool)`

GetClRfqIdOk returns a tuple with the ClRfqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClRfqId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) SetClRfqId(v string)`

SetClRfqId sets ClRfqId field to given value.

### HasClRfqId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) HasClRfqId() bool`

HasClRfqId returns a boolean if a field has been set.

### GetFee

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFeeCcy

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetFeeCcy() string`

GetFeeCcy returns the FeeCcy field if non-nil, zero value otherwise.

### GetFeeCcyOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetFeeCcyOk() (*string, bool)`

GetFeeCcyOk returns a tuple with the FeeCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCcy

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) SetFeeCcy(v string)`

SetFeeCcy sets FeeCcy field to given value.

### HasFeeCcy

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) HasFeeCcy() bool`

HasFeeCcy returns a boolean if a field has been set.

### GetInstId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetLegs

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetLegs() []CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetLegsOk() (*[]CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) SetLegs(v []CreateRfqExecuteQuoteV5RespDataInnerDataInnerLegsInner)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetMTraderCode

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetMTraderCode() string`

GetMTraderCode returns the MTraderCode field if non-nil, zero value otherwise.

### GetMTraderCodeOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetMTraderCodeOk() (*string, bool)`

GetMTraderCodeOk returns a tuple with the MTraderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMTraderCode

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) SetMTraderCode(v string)`

SetMTraderCode sets MTraderCode field to given value.

### HasMTraderCode

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) HasMTraderCode() bool`

HasMTraderCode returns a boolean if a field has been set.

### GetPx

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetQuoteId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.

### HasQuoteId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) HasQuoteId() bool`

HasQuoteId returns a boolean if a field has been set.

### GetRfqId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetRfqId() string`

GetRfqId returns the RfqId field if non-nil, zero value otherwise.

### GetRfqIdOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetRfqIdOk() (*string, bool)`

GetRfqIdOk returns a tuple with the RfqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfqId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) SetRfqId(v string)`

SetRfqId sets RfqId field to given value.

### HasRfqId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) HasRfqId() bool`

HasRfqId returns a boolean if a field has been set.

### GetSide

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSz

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTTraderCode

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetTTraderCode() string`

GetTTraderCode returns the TTraderCode field if non-nil, zero value otherwise.

### GetTTraderCodeOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetTTraderCodeOk() (*string, bool)`

GetTTraderCodeOk returns a tuple with the TTraderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTTraderCode

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) SetTTraderCode(v string)`

SetTTraderCode sets TTraderCode field to given value.

### HasTTraderCode

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) HasTTraderCode() bool`

HasTTraderCode returns a boolean if a field has been set.

### GetTag

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTradeId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *CreateRfqExecuteQuoteV5RespDataInnerDataInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


