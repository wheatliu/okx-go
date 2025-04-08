# CreateTradingBotSignalCancelSubOrderV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoId** | **string** | Algo ID | [default to ""]
**InstId** | **string** | Instrument ID, e.g. BTC-USDT-SWAP | [default to ""]
**SignalOrdId** | **string** | Order ID | [default to ""]

## Methods

### NewCreateTradingBotSignalCancelSubOrderV5Req

`func NewCreateTradingBotSignalCancelSubOrderV5Req(algoId string, instId string, signalOrdId string, ) *CreateTradingBotSignalCancelSubOrderV5Req`

NewCreateTradingBotSignalCancelSubOrderV5Req instantiates a new CreateTradingBotSignalCancelSubOrderV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradingBotSignalCancelSubOrderV5ReqWithDefaults

`func NewCreateTradingBotSignalCancelSubOrderV5ReqWithDefaults() *CreateTradingBotSignalCancelSubOrderV5Req`

NewCreateTradingBotSignalCancelSubOrderV5ReqWithDefaults instantiates a new CreateTradingBotSignalCancelSubOrderV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoId

`func (o *CreateTradingBotSignalCancelSubOrderV5Req) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *CreateTradingBotSignalCancelSubOrderV5Req) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *CreateTradingBotSignalCancelSubOrderV5Req) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.


### GetInstId

`func (o *CreateTradingBotSignalCancelSubOrderV5Req) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateTradingBotSignalCancelSubOrderV5Req) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateTradingBotSignalCancelSubOrderV5Req) SetInstId(v string)`

SetInstId sets InstId field to given value.


### GetSignalOrdId

`func (o *CreateTradingBotSignalCancelSubOrderV5Req) GetSignalOrdId() string`

GetSignalOrdId returns the SignalOrdId field if non-nil, zero value otherwise.

### GetSignalOrdIdOk

`func (o *CreateTradingBotSignalCancelSubOrderV5Req) GetSignalOrdIdOk() (*string, bool)`

GetSignalOrdIdOk returns a tuple with the SignalOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalOrdId

`func (o *CreateTradingBotSignalCancelSubOrderV5Req) SetSignalOrdId(v string)`

SetSignalOrdId sets SignalOrdId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


