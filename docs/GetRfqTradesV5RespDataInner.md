# GetRfqTradesV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockTdId** | Pointer to **string** | Block trade ID. | [optional] [default to ""]
**CTime** | Pointer to **string** | The time the trade was executed. Unix timestamp in milliseconds. | [optional] [default to ""]
**ClQuoteId** | Pointer to **string** | Client-supplied Quote ID. This attribute is treated as client sensitive information. It will not be exposed to the Taker, only return empty string. | [optional] [default to ""]
**ClRfqId** | Pointer to **string** | Client-supplied RFQ ID. This attribute is treated as client sensitive information. It will not be exposed to the Maker, only return empty string. | [optional] [default to ""]
**ErrorCode** | Pointer to **string** | Error code for unsuccessful trades.   It is \&quot;\&quot; for successful trade. | [optional] [default to ""]
**Fee** | Pointer to **string** | Fee. Negative number represents the user transaction fee charged by the platform. Positive number represents rebate. | [optional] [default to ""]
**FeeCcy** | Pointer to **string** | Fee currency | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60; | [optional] [default to ""]
**IsSuccessful** | Pointer to **bool** | Whether the trade is filled successfully | [optional] 
**Legs** | Pointer to [**[]GetRfqTradesV5RespDataInnerLegsInner**](GetRfqTradesV5RespDataInnerLegsInner.md) | Legs of trade | [optional] 
**MTraderCode** | Pointer to **string** | A unique identifier of the Maker. Empty if the anonymous parameter of the Quote is set to be &#x60;true&#x60;. | [optional] [default to ""]
**Px** | Pointer to **string** | The price the leg executed | [optional] [default to ""]
**QuoteId** | Pointer to **string** | Quote ID. | [optional] [default to ""]
**RfqId** | Pointer to **string** | RFQ ID. | [optional] [default to ""]
**Side** | Pointer to **string** | The direction of the leg. Valid value can be buy or sell. | [optional] [default to ""]
**Sz** | Pointer to **string** | Size of the leg in contracts or spot. | [optional] [default to ""]
**TTraderCode** | Pointer to **string** | A unique identifier of the Taker. Empty if the anonymous parameter of the RFQ is set to be &#x60;true&#x60;. | [optional] [default to ""]
**Tag** | Pointer to **string** | Trade tag. The block trade will have the tag of the RFQ or Quote it corresponds to. | [optional] [default to ""]
**TradeId** | Pointer to **string** | Last traded ID. | [optional] [default to ""]

## Methods

### NewGetRfqTradesV5RespDataInner

`func NewGetRfqTradesV5RespDataInner() *GetRfqTradesV5RespDataInner`

NewGetRfqTradesV5RespDataInner instantiates a new GetRfqTradesV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRfqTradesV5RespDataInnerWithDefaults

`func NewGetRfqTradesV5RespDataInnerWithDefaults() *GetRfqTradesV5RespDataInner`

NewGetRfqTradesV5RespDataInnerWithDefaults instantiates a new GetRfqTradesV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockTdId

`func (o *GetRfqTradesV5RespDataInner) GetBlockTdId() string`

GetBlockTdId returns the BlockTdId field if non-nil, zero value otherwise.

### GetBlockTdIdOk

`func (o *GetRfqTradesV5RespDataInner) GetBlockTdIdOk() (*string, bool)`

GetBlockTdIdOk returns a tuple with the BlockTdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTdId

`func (o *GetRfqTradesV5RespDataInner) SetBlockTdId(v string)`

SetBlockTdId sets BlockTdId field to given value.

### HasBlockTdId

`func (o *GetRfqTradesV5RespDataInner) HasBlockTdId() bool`

HasBlockTdId returns a boolean if a field has been set.

### GetCTime

`func (o *GetRfqTradesV5RespDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetRfqTradesV5RespDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetRfqTradesV5RespDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetRfqTradesV5RespDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetClQuoteId

`func (o *GetRfqTradesV5RespDataInner) GetClQuoteId() string`

GetClQuoteId returns the ClQuoteId field if non-nil, zero value otherwise.

### GetClQuoteIdOk

`func (o *GetRfqTradesV5RespDataInner) GetClQuoteIdOk() (*string, bool)`

GetClQuoteIdOk returns a tuple with the ClQuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClQuoteId

`func (o *GetRfqTradesV5RespDataInner) SetClQuoteId(v string)`

SetClQuoteId sets ClQuoteId field to given value.

### HasClQuoteId

`func (o *GetRfqTradesV5RespDataInner) HasClQuoteId() bool`

HasClQuoteId returns a boolean if a field has been set.

### GetClRfqId

`func (o *GetRfqTradesV5RespDataInner) GetClRfqId() string`

GetClRfqId returns the ClRfqId field if non-nil, zero value otherwise.

### GetClRfqIdOk

`func (o *GetRfqTradesV5RespDataInner) GetClRfqIdOk() (*string, bool)`

GetClRfqIdOk returns a tuple with the ClRfqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClRfqId

`func (o *GetRfqTradesV5RespDataInner) SetClRfqId(v string)`

SetClRfqId sets ClRfqId field to given value.

### HasClRfqId

`func (o *GetRfqTradesV5RespDataInner) HasClRfqId() bool`

HasClRfqId returns a boolean if a field has been set.

### GetErrorCode

`func (o *GetRfqTradesV5RespDataInner) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GetRfqTradesV5RespDataInner) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GetRfqTradesV5RespDataInner) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GetRfqTradesV5RespDataInner) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetFee

`func (o *GetRfqTradesV5RespDataInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetRfqTradesV5RespDataInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetRfqTradesV5RespDataInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetRfqTradesV5RespDataInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFeeCcy

`func (o *GetRfqTradesV5RespDataInner) GetFeeCcy() string`

GetFeeCcy returns the FeeCcy field if non-nil, zero value otherwise.

### GetFeeCcyOk

`func (o *GetRfqTradesV5RespDataInner) GetFeeCcyOk() (*string, bool)`

GetFeeCcyOk returns a tuple with the FeeCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCcy

`func (o *GetRfqTradesV5RespDataInner) SetFeeCcy(v string)`

SetFeeCcy sets FeeCcy field to given value.

### HasFeeCcy

`func (o *GetRfqTradesV5RespDataInner) HasFeeCcy() bool`

HasFeeCcy returns a boolean if a field has been set.

### GetInstId

`func (o *GetRfqTradesV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetRfqTradesV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetRfqTradesV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetRfqTradesV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetIsSuccessful

`func (o *GetRfqTradesV5RespDataInner) GetIsSuccessful() bool`

GetIsSuccessful returns the IsSuccessful field if non-nil, zero value otherwise.

### GetIsSuccessfulOk

`func (o *GetRfqTradesV5RespDataInner) GetIsSuccessfulOk() (*bool, bool)`

GetIsSuccessfulOk returns a tuple with the IsSuccessful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccessful

`func (o *GetRfqTradesV5RespDataInner) SetIsSuccessful(v bool)`

SetIsSuccessful sets IsSuccessful field to given value.

### HasIsSuccessful

`func (o *GetRfqTradesV5RespDataInner) HasIsSuccessful() bool`

HasIsSuccessful returns a boolean if a field has been set.

### GetLegs

`func (o *GetRfqTradesV5RespDataInner) GetLegs() []GetRfqTradesV5RespDataInnerLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *GetRfqTradesV5RespDataInner) GetLegsOk() (*[]GetRfqTradesV5RespDataInnerLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *GetRfqTradesV5RespDataInner) SetLegs(v []GetRfqTradesV5RespDataInnerLegsInner)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *GetRfqTradesV5RespDataInner) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetMTraderCode

`func (o *GetRfqTradesV5RespDataInner) GetMTraderCode() string`

GetMTraderCode returns the MTraderCode field if non-nil, zero value otherwise.

### GetMTraderCodeOk

`func (o *GetRfqTradesV5RespDataInner) GetMTraderCodeOk() (*string, bool)`

GetMTraderCodeOk returns a tuple with the MTraderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMTraderCode

`func (o *GetRfqTradesV5RespDataInner) SetMTraderCode(v string)`

SetMTraderCode sets MTraderCode field to given value.

### HasMTraderCode

`func (o *GetRfqTradesV5RespDataInner) HasMTraderCode() bool`

HasMTraderCode returns a boolean if a field has been set.

### GetPx

`func (o *GetRfqTradesV5RespDataInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetRfqTradesV5RespDataInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetRfqTradesV5RespDataInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetRfqTradesV5RespDataInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetQuoteId

`func (o *GetRfqTradesV5RespDataInner) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *GetRfqTradesV5RespDataInner) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *GetRfqTradesV5RespDataInner) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.

### HasQuoteId

`func (o *GetRfqTradesV5RespDataInner) HasQuoteId() bool`

HasQuoteId returns a boolean if a field has been set.

### GetRfqId

`func (o *GetRfqTradesV5RespDataInner) GetRfqId() string`

GetRfqId returns the RfqId field if non-nil, zero value otherwise.

### GetRfqIdOk

`func (o *GetRfqTradesV5RespDataInner) GetRfqIdOk() (*string, bool)`

GetRfqIdOk returns a tuple with the RfqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfqId

`func (o *GetRfqTradesV5RespDataInner) SetRfqId(v string)`

SetRfqId sets RfqId field to given value.

### HasRfqId

`func (o *GetRfqTradesV5RespDataInner) HasRfqId() bool`

HasRfqId returns a boolean if a field has been set.

### GetSide

`func (o *GetRfqTradesV5RespDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetRfqTradesV5RespDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetRfqTradesV5RespDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetRfqTradesV5RespDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSz

`func (o *GetRfqTradesV5RespDataInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetRfqTradesV5RespDataInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetRfqTradesV5RespDataInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetRfqTradesV5RespDataInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTTraderCode

`func (o *GetRfqTradesV5RespDataInner) GetTTraderCode() string`

GetTTraderCode returns the TTraderCode field if non-nil, zero value otherwise.

### GetTTraderCodeOk

`func (o *GetRfqTradesV5RespDataInner) GetTTraderCodeOk() (*string, bool)`

GetTTraderCodeOk returns a tuple with the TTraderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTTraderCode

`func (o *GetRfqTradesV5RespDataInner) SetTTraderCode(v string)`

SetTTraderCode sets TTraderCode field to given value.

### HasTTraderCode

`func (o *GetRfqTradesV5RespDataInner) HasTTraderCode() bool`

HasTTraderCode returns a boolean if a field has been set.

### GetTag

`func (o *GetRfqTradesV5RespDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetRfqTradesV5RespDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetRfqTradesV5RespDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetRfqTradesV5RespDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTradeId

`func (o *GetRfqTradesV5RespDataInner) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *GetRfqTradesV5RespDataInner) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *GetRfqTradesV5RespDataInner) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *GetRfqTradesV5RespDataInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


