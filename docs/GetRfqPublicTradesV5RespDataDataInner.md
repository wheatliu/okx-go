# GetRfqPublicTradesV5RespDataDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockTdId** | Pointer to **string** | Block trade ID. | [optional] [default to ""]
**CTime** | Pointer to **string** | The time the trade was executed. Unix timestamp in milliseconds. | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. BTC-USDT-SWAP | [optional] [default to ""]
**Legs** | Pointer to [**[]GetRfqPublicTradesV5RespDataDataInnerLegsInner**](GetRfqPublicTradesV5RespDataDataInnerLegsInner.md) | Legs of trade | [optional] 
**Px** | Pointer to **string** | The price the leg executed | [optional] [default to ""]
**Side** | Pointer to **string** | The direction of the leg from the Takers perspective. Valid value can be buy or sell. | [optional] [default to ""]
**Strategy** | Pointer to **string** | Option strategy, e.g. CALL_CALENDAR_SPREAD | [optional] [default to ""]
**Sz** | Pointer to **string** | Trade quantity   For spot trading, the unit is base currency  For &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;, the unit is contract. | [optional] [default to ""]
**TradeId** | Pointer to **string** | Last traded ID. | [optional] [default to ""]

## Methods

### NewGetRfqPublicTradesV5RespDataDataInner

`func NewGetRfqPublicTradesV5RespDataDataInner() *GetRfqPublicTradesV5RespDataDataInner`

NewGetRfqPublicTradesV5RespDataDataInner instantiates a new GetRfqPublicTradesV5RespDataDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRfqPublicTradesV5RespDataDataInnerWithDefaults

`func NewGetRfqPublicTradesV5RespDataDataInnerWithDefaults() *GetRfqPublicTradesV5RespDataDataInner`

NewGetRfqPublicTradesV5RespDataDataInnerWithDefaults instantiates a new GetRfqPublicTradesV5RespDataDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockTdId

`func (o *GetRfqPublicTradesV5RespDataDataInner) GetBlockTdId() string`

GetBlockTdId returns the BlockTdId field if non-nil, zero value otherwise.

### GetBlockTdIdOk

`func (o *GetRfqPublicTradesV5RespDataDataInner) GetBlockTdIdOk() (*string, bool)`

GetBlockTdIdOk returns a tuple with the BlockTdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTdId

`func (o *GetRfqPublicTradesV5RespDataDataInner) SetBlockTdId(v string)`

SetBlockTdId sets BlockTdId field to given value.

### HasBlockTdId

`func (o *GetRfqPublicTradesV5RespDataDataInner) HasBlockTdId() bool`

HasBlockTdId returns a boolean if a field has been set.

### GetCTime

`func (o *GetRfqPublicTradesV5RespDataDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetRfqPublicTradesV5RespDataDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetRfqPublicTradesV5RespDataDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetRfqPublicTradesV5RespDataDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetInstId

`func (o *GetRfqPublicTradesV5RespDataDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetRfqPublicTradesV5RespDataDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetRfqPublicTradesV5RespDataDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetRfqPublicTradesV5RespDataDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetLegs

`func (o *GetRfqPublicTradesV5RespDataDataInner) GetLegs() []GetRfqPublicTradesV5RespDataDataInnerLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *GetRfqPublicTradesV5RespDataDataInner) GetLegsOk() (*[]GetRfqPublicTradesV5RespDataDataInnerLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *GetRfqPublicTradesV5RespDataDataInner) SetLegs(v []GetRfqPublicTradesV5RespDataDataInnerLegsInner)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *GetRfqPublicTradesV5RespDataDataInner) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetPx

`func (o *GetRfqPublicTradesV5RespDataDataInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetRfqPublicTradesV5RespDataDataInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetRfqPublicTradesV5RespDataDataInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetRfqPublicTradesV5RespDataDataInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSide

`func (o *GetRfqPublicTradesV5RespDataDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetRfqPublicTradesV5RespDataDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetRfqPublicTradesV5RespDataDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetRfqPublicTradesV5RespDataDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStrategy

`func (o *GetRfqPublicTradesV5RespDataDataInner) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *GetRfqPublicTradesV5RespDataDataInner) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *GetRfqPublicTradesV5RespDataDataInner) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *GetRfqPublicTradesV5RespDataDataInner) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetSz

`func (o *GetRfqPublicTradesV5RespDataDataInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetRfqPublicTradesV5RespDataDataInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetRfqPublicTradesV5RespDataDataInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetRfqPublicTradesV5RespDataDataInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTradeId

`func (o *GetRfqPublicTradesV5RespDataDataInner) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *GetRfqPublicTradesV5RespDataDataInner) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *GetRfqPublicTradesV5RespDataDataInner) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *GetRfqPublicTradesV5RespDataDataInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


