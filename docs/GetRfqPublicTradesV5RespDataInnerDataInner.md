# GetRfqPublicTradesV5RespDataInnerDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockTdId** | Pointer to **string** | Block trade ID. | [optional] [default to ""]
**CTime** | Pointer to **string** | The time the trade was executed. Unix timestamp in milliseconds. | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. BTC-USDT-SWAP | [optional] [default to ""]
**Legs** | Pointer to [**[]GetRfqPublicTradesV5RespDataInnerDataInnerLegsInner**](GetRfqPublicTradesV5RespDataInnerDataInnerLegsInner.md) | Legs of trade | [optional] 
**Px** | Pointer to **string** | The price the leg executed | [optional] [default to ""]
**Side** | Pointer to **string** | The direction of the leg from the Takers perspective. Valid value can be buy or sell. | [optional] [default to ""]
**Strategy** | Pointer to **string** | Option strategy, e.g. CALL_CALENDAR_SPREAD | [optional] [default to ""]
**Sz** | Pointer to **string** | Trade quantity   For spot trading, the unit is base currency  For &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;, the unit is contract. | [optional] [default to ""]
**TradeId** | Pointer to **string** | Last traded ID. | [optional] [default to ""]

## Methods

### NewGetRfqPublicTradesV5RespDataInnerDataInner

`func NewGetRfqPublicTradesV5RespDataInnerDataInner() *GetRfqPublicTradesV5RespDataInnerDataInner`

NewGetRfqPublicTradesV5RespDataInnerDataInner instantiates a new GetRfqPublicTradesV5RespDataInnerDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRfqPublicTradesV5RespDataInnerDataInnerWithDefaults

`func NewGetRfqPublicTradesV5RespDataInnerDataInnerWithDefaults() *GetRfqPublicTradesV5RespDataInnerDataInner`

NewGetRfqPublicTradesV5RespDataInnerDataInnerWithDefaults instantiates a new GetRfqPublicTradesV5RespDataInnerDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockTdId

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetBlockTdId() string`

GetBlockTdId returns the BlockTdId field if non-nil, zero value otherwise.

### GetBlockTdIdOk

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetBlockTdIdOk() (*string, bool)`

GetBlockTdIdOk returns a tuple with the BlockTdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTdId

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) SetBlockTdId(v string)`

SetBlockTdId sets BlockTdId field to given value.

### HasBlockTdId

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) HasBlockTdId() bool`

HasBlockTdId returns a boolean if a field has been set.

### GetCTime

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetInstId

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetLegs

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetLegs() []GetRfqPublicTradesV5RespDataInnerDataInnerLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetLegsOk() (*[]GetRfqPublicTradesV5RespDataInnerDataInnerLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) SetLegs(v []GetRfqPublicTradesV5RespDataInnerDataInnerLegsInner)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetPx

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSide

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStrategy

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetSz

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTradeId

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *GetRfqPublicTradesV5RespDataInnerDataInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


