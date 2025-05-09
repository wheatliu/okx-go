# CreateRfqExecuteQuoteV5RespDataInner

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
**Legs** | Pointer to [**[]CreateRfqExecuteQuoteV5RespDataInnerLegsInner**](CreateRfqExecuteQuoteV5RespDataInnerLegsInner.md) | Legs of trade | [optional] 
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

### NewCreateRfqExecuteQuoteV5RespDataInner

`func NewCreateRfqExecuteQuoteV5RespDataInner() *CreateRfqExecuteQuoteV5RespDataInner`

NewCreateRfqExecuteQuoteV5RespDataInner instantiates a new CreateRfqExecuteQuoteV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqExecuteQuoteV5RespDataInnerWithDefaults

`func NewCreateRfqExecuteQuoteV5RespDataInnerWithDefaults() *CreateRfqExecuteQuoteV5RespDataInner`

NewCreateRfqExecuteQuoteV5RespDataInnerWithDefaults instantiates a new CreateRfqExecuteQuoteV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockTdId

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetBlockTdId() string`

GetBlockTdId returns the BlockTdId field if non-nil, zero value otherwise.

### GetBlockTdIdOk

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetBlockTdIdOk() (*string, bool)`

GetBlockTdIdOk returns a tuple with the BlockTdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTdId

`func (o *CreateRfqExecuteQuoteV5RespDataInner) SetBlockTdId(v string)`

SetBlockTdId sets BlockTdId field to given value.

### HasBlockTdId

`func (o *CreateRfqExecuteQuoteV5RespDataInner) HasBlockTdId() bool`

HasBlockTdId returns a boolean if a field has been set.

### GetCTime

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *CreateRfqExecuteQuoteV5RespDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *CreateRfqExecuteQuoteV5RespDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetClQuoteId

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetClQuoteId() string`

GetClQuoteId returns the ClQuoteId field if non-nil, zero value otherwise.

### GetClQuoteIdOk

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetClQuoteIdOk() (*string, bool)`

GetClQuoteIdOk returns a tuple with the ClQuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClQuoteId

`func (o *CreateRfqExecuteQuoteV5RespDataInner) SetClQuoteId(v string)`

SetClQuoteId sets ClQuoteId field to given value.

### HasClQuoteId

`func (o *CreateRfqExecuteQuoteV5RespDataInner) HasClQuoteId() bool`

HasClQuoteId returns a boolean if a field has been set.

### GetClRfqId

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetClRfqId() string`

GetClRfqId returns the ClRfqId field if non-nil, zero value otherwise.

### GetClRfqIdOk

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetClRfqIdOk() (*string, bool)`

GetClRfqIdOk returns a tuple with the ClRfqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClRfqId

`func (o *CreateRfqExecuteQuoteV5RespDataInner) SetClRfqId(v string)`

SetClRfqId sets ClRfqId field to given value.

### HasClRfqId

`func (o *CreateRfqExecuteQuoteV5RespDataInner) HasClRfqId() bool`

HasClRfqId returns a boolean if a field has been set.

### GetFee

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CreateRfqExecuteQuoteV5RespDataInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *CreateRfqExecuteQuoteV5RespDataInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFeeCcy

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetFeeCcy() string`

GetFeeCcy returns the FeeCcy field if non-nil, zero value otherwise.

### GetFeeCcyOk

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetFeeCcyOk() (*string, bool)`

GetFeeCcyOk returns a tuple with the FeeCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCcy

`func (o *CreateRfqExecuteQuoteV5RespDataInner) SetFeeCcy(v string)`

SetFeeCcy sets FeeCcy field to given value.

### HasFeeCcy

`func (o *CreateRfqExecuteQuoteV5RespDataInner) HasFeeCcy() bool`

HasFeeCcy returns a boolean if a field has been set.

### GetInstId

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateRfqExecuteQuoteV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *CreateRfqExecuteQuoteV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetLegs

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetLegs() []CreateRfqExecuteQuoteV5RespDataInnerLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetLegsOk() (*[]CreateRfqExecuteQuoteV5RespDataInnerLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *CreateRfqExecuteQuoteV5RespDataInner) SetLegs(v []CreateRfqExecuteQuoteV5RespDataInnerLegsInner)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *CreateRfqExecuteQuoteV5RespDataInner) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetMTraderCode

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetMTraderCode() string`

GetMTraderCode returns the MTraderCode field if non-nil, zero value otherwise.

### GetMTraderCodeOk

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetMTraderCodeOk() (*string, bool)`

GetMTraderCodeOk returns a tuple with the MTraderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMTraderCode

`func (o *CreateRfqExecuteQuoteV5RespDataInner) SetMTraderCode(v string)`

SetMTraderCode sets MTraderCode field to given value.

### HasMTraderCode

`func (o *CreateRfqExecuteQuoteV5RespDataInner) HasMTraderCode() bool`

HasMTraderCode returns a boolean if a field has been set.

### GetPx

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *CreateRfqExecuteQuoteV5RespDataInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *CreateRfqExecuteQuoteV5RespDataInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetQuoteId

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *CreateRfqExecuteQuoteV5RespDataInner) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.

### HasQuoteId

`func (o *CreateRfqExecuteQuoteV5RespDataInner) HasQuoteId() bool`

HasQuoteId returns a boolean if a field has been set.

### GetRfqId

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetRfqId() string`

GetRfqId returns the RfqId field if non-nil, zero value otherwise.

### GetRfqIdOk

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetRfqIdOk() (*string, bool)`

GetRfqIdOk returns a tuple with the RfqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfqId

`func (o *CreateRfqExecuteQuoteV5RespDataInner) SetRfqId(v string)`

SetRfqId sets RfqId field to given value.

### HasRfqId

`func (o *CreateRfqExecuteQuoteV5RespDataInner) HasRfqId() bool`

HasRfqId returns a boolean if a field has been set.

### GetSide

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateRfqExecuteQuoteV5RespDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *CreateRfqExecuteQuoteV5RespDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSz

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *CreateRfqExecuteQuoteV5RespDataInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *CreateRfqExecuteQuoteV5RespDataInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTTraderCode

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetTTraderCode() string`

GetTTraderCode returns the TTraderCode field if non-nil, zero value otherwise.

### GetTTraderCodeOk

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetTTraderCodeOk() (*string, bool)`

GetTTraderCodeOk returns a tuple with the TTraderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTTraderCode

`func (o *CreateRfqExecuteQuoteV5RespDataInner) SetTTraderCode(v string)`

SetTTraderCode sets TTraderCode field to given value.

### HasTTraderCode

`func (o *CreateRfqExecuteQuoteV5RespDataInner) HasTTraderCode() bool`

HasTTraderCode returns a boolean if a field has been set.

### GetTag

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateRfqExecuteQuoteV5RespDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateRfqExecuteQuoteV5RespDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTradeId

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *CreateRfqExecuteQuoteV5RespDataInner) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *CreateRfqExecuteQuoteV5RespDataInner) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *CreateRfqExecuteQuoteV5RespDataInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


