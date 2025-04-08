# GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstId** | Pointer to **string** | The Instrument ID | [optional] [default to ""]
**Px** | Pointer to **string** | Trade price | [optional] [default to ""]
**Side** | Pointer to **string** | Trade side  &#x60;buy&#x60;  &#x60;sell&#x60; | [optional] [default to ""]
**Sz** | Pointer to **string** | Trade quantity. The unit is contract. | [optional] [default to ""]
**TradeId** | Pointer to **string** | Trade ID | [optional] [default to ""]
**Ts** | Pointer to **string** | Trade time, Unix timestamp format in milliseconds, e.g. 1597026383085. | [optional] [default to ""]

## Methods

### NewGetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner

`func NewGetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner() *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner`

NewGetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner instantiates a new GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInnerWithDefaults

`func NewGetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInnerWithDefaults() *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner`

NewGetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInnerWithDefaults instantiates a new GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstId

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetPx

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSide

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSz

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTradeId

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.

### GetTs

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


