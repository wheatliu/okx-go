# GetRfqPublicTradesV5RespDataInnerLegsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstId** | Pointer to **string** | Instrument ID, e.g. BTC-USDT-SWAP | [optional] [default to ""]
**Px** | Pointer to **string** | The price the leg executed | [optional] [default to ""]
**Side** | Pointer to **string** | The direction of the leg from the Takers perspective. Valid value can be buy or sell. | [optional] [default to ""]
**Sz** | Pointer to **string** | Trade quantity   For spot trading, the unit is base currency  For &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;, the unit is contract. | [optional] [default to ""]
**TradeId** | Pointer to **string** | Last traded ID. | [optional] [default to ""]

## Methods

### NewGetRfqPublicTradesV5RespDataInnerLegsInner

`func NewGetRfqPublicTradesV5RespDataInnerLegsInner() *GetRfqPublicTradesV5RespDataInnerLegsInner`

NewGetRfqPublicTradesV5RespDataInnerLegsInner instantiates a new GetRfqPublicTradesV5RespDataInnerLegsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRfqPublicTradesV5RespDataInnerLegsInnerWithDefaults

`func NewGetRfqPublicTradesV5RespDataInnerLegsInnerWithDefaults() *GetRfqPublicTradesV5RespDataInnerLegsInner`

NewGetRfqPublicTradesV5RespDataInnerLegsInnerWithDefaults instantiates a new GetRfqPublicTradesV5RespDataInnerLegsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstId

`func (o *GetRfqPublicTradesV5RespDataInnerLegsInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetRfqPublicTradesV5RespDataInnerLegsInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetRfqPublicTradesV5RespDataInnerLegsInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetRfqPublicTradesV5RespDataInnerLegsInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetPx

`func (o *GetRfqPublicTradesV5RespDataInnerLegsInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetRfqPublicTradesV5RespDataInnerLegsInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetRfqPublicTradesV5RespDataInnerLegsInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetRfqPublicTradesV5RespDataInnerLegsInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSide

`func (o *GetRfqPublicTradesV5RespDataInnerLegsInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetRfqPublicTradesV5RespDataInnerLegsInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetRfqPublicTradesV5RespDataInnerLegsInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetRfqPublicTradesV5RespDataInnerLegsInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSz

`func (o *GetRfqPublicTradesV5RespDataInnerLegsInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetRfqPublicTradesV5RespDataInnerLegsInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetRfqPublicTradesV5RespDataInnerLegsInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetRfqPublicTradesV5RespDataInnerLegsInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTradeId

`func (o *GetRfqPublicTradesV5RespDataInnerLegsInner) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *GetRfqPublicTradesV5RespDataInnerLegsInner) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *GetRfqPublicTradesV5RespDataInnerLegsInner) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *GetRfqPublicTradesV5RespDataInnerLegsInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


