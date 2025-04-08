# CreateTradingBotSignalSubOrderV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoId** | **string** | Algo ID | [default to ""]
**InstId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60; | [default to ""]
**OrdType** | **string** | Order type   &#x60;market&#x60;: Market order   &#x60;limit&#x60;: Limit order | [default to ""]
**Px** | Pointer to **string** | Order price. Only applicable to &#x60;limit&#x60; order. | [optional] [default to ""]
**ReduceOnly** | Pointer to **bool** | Whether orders can only reduce in position size.   Valid options: &#x60;true&#x60; or &#x60;false&#x60;. The default value is &#x60;false&#x60;.   Only applicable to &#x60;Spot and futures mode&#x60;/&#x60;Multi-currency margin&#x60; | [optional] 
**Side** | **string** | Order side, &#x60;buy&#x60; &#x60;sell&#x60; | [default to ""]
**Sz** | **string** | Quantity to buy or sell | [default to ""]

## Methods

### NewCreateTradingBotSignalSubOrderV5Req

`func NewCreateTradingBotSignalSubOrderV5Req(algoId string, instId string, ordType string, side string, sz string, ) *CreateTradingBotSignalSubOrderV5Req`

NewCreateTradingBotSignalSubOrderV5Req instantiates a new CreateTradingBotSignalSubOrderV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradingBotSignalSubOrderV5ReqWithDefaults

`func NewCreateTradingBotSignalSubOrderV5ReqWithDefaults() *CreateTradingBotSignalSubOrderV5Req`

NewCreateTradingBotSignalSubOrderV5ReqWithDefaults instantiates a new CreateTradingBotSignalSubOrderV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoId

`func (o *CreateTradingBotSignalSubOrderV5Req) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *CreateTradingBotSignalSubOrderV5Req) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *CreateTradingBotSignalSubOrderV5Req) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.


### GetInstId

`func (o *CreateTradingBotSignalSubOrderV5Req) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateTradingBotSignalSubOrderV5Req) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateTradingBotSignalSubOrderV5Req) SetInstId(v string)`

SetInstId sets InstId field to given value.


### GetOrdType

`func (o *CreateTradingBotSignalSubOrderV5Req) GetOrdType() string`

GetOrdType returns the OrdType field if non-nil, zero value otherwise.

### GetOrdTypeOk

`func (o *CreateTradingBotSignalSubOrderV5Req) GetOrdTypeOk() (*string, bool)`

GetOrdTypeOk returns a tuple with the OrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdType

`func (o *CreateTradingBotSignalSubOrderV5Req) SetOrdType(v string)`

SetOrdType sets OrdType field to given value.


### GetPx

`func (o *CreateTradingBotSignalSubOrderV5Req) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *CreateTradingBotSignalSubOrderV5Req) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *CreateTradingBotSignalSubOrderV5Req) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *CreateTradingBotSignalSubOrderV5Req) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetReduceOnly

`func (o *CreateTradingBotSignalSubOrderV5Req) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *CreateTradingBotSignalSubOrderV5Req) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *CreateTradingBotSignalSubOrderV5Req) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *CreateTradingBotSignalSubOrderV5Req) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSide

`func (o *CreateTradingBotSignalSubOrderV5Req) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateTradingBotSignalSubOrderV5Req) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateTradingBotSignalSubOrderV5Req) SetSide(v string)`

SetSide sets Side field to given value.


### GetSz

`func (o *CreateTradingBotSignalSubOrderV5Req) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *CreateTradingBotSignalSubOrderV5Req) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *CreateTradingBotSignalSubOrderV5Req) SetSz(v string)`

SetSz sets Sz field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


