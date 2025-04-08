# CreateTradingBotGridClosePositionV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoId** | **string** | Algo ID | [default to ""]
**MktClose** | **bool** | Market close all the positions or not  &#x60;true&#x60;: Market close all position, &#x60;false&#x60;: Close part of position | 
**Px** | Pointer to **string** | Close position price  If &#x60;mktClose&#x60; is &#x60;false&#x60;, the parameter is required. | [optional] [default to ""]
**Sz** | Pointer to **string** | Close position amount, with unit of &#x60;contract&#x60;  If &#x60;mktClose&#x60; is &#x60;false&#x60;, the parameter is required. | [optional] [default to ""]

## Methods

### NewCreateTradingBotGridClosePositionV5Req

`func NewCreateTradingBotGridClosePositionV5Req(algoId string, mktClose bool, ) *CreateTradingBotGridClosePositionV5Req`

NewCreateTradingBotGridClosePositionV5Req instantiates a new CreateTradingBotGridClosePositionV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradingBotGridClosePositionV5ReqWithDefaults

`func NewCreateTradingBotGridClosePositionV5ReqWithDefaults() *CreateTradingBotGridClosePositionV5Req`

NewCreateTradingBotGridClosePositionV5ReqWithDefaults instantiates a new CreateTradingBotGridClosePositionV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoId

`func (o *CreateTradingBotGridClosePositionV5Req) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *CreateTradingBotGridClosePositionV5Req) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *CreateTradingBotGridClosePositionV5Req) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.


### GetMktClose

`func (o *CreateTradingBotGridClosePositionV5Req) GetMktClose() bool`

GetMktClose returns the MktClose field if non-nil, zero value otherwise.

### GetMktCloseOk

`func (o *CreateTradingBotGridClosePositionV5Req) GetMktCloseOk() (*bool, bool)`

GetMktCloseOk returns a tuple with the MktClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMktClose

`func (o *CreateTradingBotGridClosePositionV5Req) SetMktClose(v bool)`

SetMktClose sets MktClose field to given value.


### GetPx

`func (o *CreateTradingBotGridClosePositionV5Req) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *CreateTradingBotGridClosePositionV5Req) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *CreateTradingBotGridClosePositionV5Req) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *CreateTradingBotGridClosePositionV5Req) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSz

`func (o *CreateTradingBotGridClosePositionV5Req) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *CreateTradingBotGridClosePositionV5Req) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *CreateTradingBotGridClosePositionV5Req) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *CreateTradingBotGridClosePositionV5Req) HasSz() bool`

HasSz returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


