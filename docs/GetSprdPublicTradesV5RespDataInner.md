# GetSprdPublicTradesV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Px** | Pointer to **string** | Trade price | [optional] [default to ""]
**Side** | Pointer to **string** | Trade side of the taker.   &#x60;buy&#x60;   &#x60;sell&#x60; | [optional] [default to ""]
**SprdId** | Pointer to **string** | spread ID | [optional] [default to ""]
**Sz** | Pointer to **string** | Trade quantity | [optional] [default to ""]
**TradeId** | Pointer to **string** | Trade ID | [optional] [default to ""]
**Ts** | Pointer to **string** | Trade time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60;. | [optional] [default to ""]

## Methods

### NewGetSprdPublicTradesV5RespDataInner

`func NewGetSprdPublicTradesV5RespDataInner() *GetSprdPublicTradesV5RespDataInner`

NewGetSprdPublicTradesV5RespDataInner instantiates a new GetSprdPublicTradesV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSprdPublicTradesV5RespDataInnerWithDefaults

`func NewGetSprdPublicTradesV5RespDataInnerWithDefaults() *GetSprdPublicTradesV5RespDataInner`

NewGetSprdPublicTradesV5RespDataInnerWithDefaults instantiates a new GetSprdPublicTradesV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPx

`func (o *GetSprdPublicTradesV5RespDataInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetSprdPublicTradesV5RespDataInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetSprdPublicTradesV5RespDataInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetSprdPublicTradesV5RespDataInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSide

`func (o *GetSprdPublicTradesV5RespDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetSprdPublicTradesV5RespDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetSprdPublicTradesV5RespDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetSprdPublicTradesV5RespDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSprdId

`func (o *GetSprdPublicTradesV5RespDataInner) GetSprdId() string`

GetSprdId returns the SprdId field if non-nil, zero value otherwise.

### GetSprdIdOk

`func (o *GetSprdPublicTradesV5RespDataInner) GetSprdIdOk() (*string, bool)`

GetSprdIdOk returns a tuple with the SprdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSprdId

`func (o *GetSprdPublicTradesV5RespDataInner) SetSprdId(v string)`

SetSprdId sets SprdId field to given value.

### HasSprdId

`func (o *GetSprdPublicTradesV5RespDataInner) HasSprdId() bool`

HasSprdId returns a boolean if a field has been set.

### GetSz

`func (o *GetSprdPublicTradesV5RespDataInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetSprdPublicTradesV5RespDataInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetSprdPublicTradesV5RespDataInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetSprdPublicTradesV5RespDataInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTradeId

`func (o *GetSprdPublicTradesV5RespDataInner) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *GetSprdPublicTradesV5RespDataInner) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *GetSprdPublicTradesV5RespDataInner) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *GetSprdPublicTradesV5RespDataInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.

### GetTs

`func (o *GetSprdPublicTradesV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetSprdPublicTradesV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetSprdPublicTradesV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetSprdPublicTradesV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


