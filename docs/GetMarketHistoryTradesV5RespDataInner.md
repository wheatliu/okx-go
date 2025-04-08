# GetMarketHistoryTradesV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**Px** | Pointer to **string** | Trade price | [optional] [default to ""]
**Side** | Pointer to **string** | Trade side   &#x60;buy&#x60;    &#x60;sell&#x60; | [optional] [default to ""]
**Sz** | Pointer to **string** | Trade quantity   For spot trading, the unit is base currency  For &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;, the unit is contract. | [optional] [default to ""]
**TradeId** | Pointer to **string** | Trade ID | [optional] [default to ""]
**Ts** | Pointer to **string** | Trade time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60;. | [optional] [default to ""]

## Methods

### NewGetMarketHistoryTradesV5RespDataInner

`func NewGetMarketHistoryTradesV5RespDataInner() *GetMarketHistoryTradesV5RespDataInner`

NewGetMarketHistoryTradesV5RespDataInner instantiates a new GetMarketHistoryTradesV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketHistoryTradesV5RespDataInnerWithDefaults

`func NewGetMarketHistoryTradesV5RespDataInnerWithDefaults() *GetMarketHistoryTradesV5RespDataInner`

NewGetMarketHistoryTradesV5RespDataInnerWithDefaults instantiates a new GetMarketHistoryTradesV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstId

`func (o *GetMarketHistoryTradesV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetMarketHistoryTradesV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetMarketHistoryTradesV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetMarketHistoryTradesV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetPx

`func (o *GetMarketHistoryTradesV5RespDataInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetMarketHistoryTradesV5RespDataInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetMarketHistoryTradesV5RespDataInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetMarketHistoryTradesV5RespDataInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSide

`func (o *GetMarketHistoryTradesV5RespDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetMarketHistoryTradesV5RespDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetMarketHistoryTradesV5RespDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetMarketHistoryTradesV5RespDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSz

`func (o *GetMarketHistoryTradesV5RespDataInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetMarketHistoryTradesV5RespDataInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetMarketHistoryTradesV5RespDataInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetMarketHistoryTradesV5RespDataInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTradeId

`func (o *GetMarketHistoryTradesV5RespDataInner) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *GetMarketHistoryTradesV5RespDataInner) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *GetMarketHistoryTradesV5RespDataInner) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *GetMarketHistoryTradesV5RespDataInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.

### GetTs

`func (o *GetMarketHistoryTradesV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetMarketHistoryTradesV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetMarketHistoryTradesV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetMarketHistoryTradesV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


