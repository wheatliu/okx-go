# CreateTradingBotSignalSetInstrumentsV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoId** | **string** | Algo ID | [default to ""]
**IncludeAll** | **bool** | Whether to include all USDT-margined contract.The default value is &#x60;false&#x60;. &#x60;true&#x60;: include &#x60;false&#x60; : exclude | 
**InstIds** | **[]string** | Instrument IDs. When &#x60;includeAll&#x60; is &#x60;true&#x60;, it is ignored | 

## Methods

### NewCreateTradingBotSignalSetInstrumentsV5Req

`func NewCreateTradingBotSignalSetInstrumentsV5Req(algoId string, includeAll bool, instIds []string, ) *CreateTradingBotSignalSetInstrumentsV5Req`

NewCreateTradingBotSignalSetInstrumentsV5Req instantiates a new CreateTradingBotSignalSetInstrumentsV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradingBotSignalSetInstrumentsV5ReqWithDefaults

`func NewCreateTradingBotSignalSetInstrumentsV5ReqWithDefaults() *CreateTradingBotSignalSetInstrumentsV5Req`

NewCreateTradingBotSignalSetInstrumentsV5ReqWithDefaults instantiates a new CreateTradingBotSignalSetInstrumentsV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoId

`func (o *CreateTradingBotSignalSetInstrumentsV5Req) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *CreateTradingBotSignalSetInstrumentsV5Req) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *CreateTradingBotSignalSetInstrumentsV5Req) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.


### GetIncludeAll

`func (o *CreateTradingBotSignalSetInstrumentsV5Req) GetIncludeAll() bool`

GetIncludeAll returns the IncludeAll field if non-nil, zero value otherwise.

### GetIncludeAllOk

`func (o *CreateTradingBotSignalSetInstrumentsV5Req) GetIncludeAllOk() (*bool, bool)`

GetIncludeAllOk returns a tuple with the IncludeAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAll

`func (o *CreateTradingBotSignalSetInstrumentsV5Req) SetIncludeAll(v bool)`

SetIncludeAll sets IncludeAll field to given value.


### GetInstIds

`func (o *CreateTradingBotSignalSetInstrumentsV5Req) GetInstIds() []string`

GetInstIds returns the InstIds field if non-nil, zero value otherwise.

### GetInstIdsOk

`func (o *CreateTradingBotSignalSetInstrumentsV5Req) GetInstIdsOk() (*[]string, bool)`

GetInstIdsOk returns a tuple with the InstIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstIds

`func (o *CreateTradingBotSignalSetInstrumentsV5Req) SetInstIds(v []string)`

SetInstIds sets InstIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


