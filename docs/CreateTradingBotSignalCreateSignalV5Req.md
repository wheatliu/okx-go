# CreateTradingBotSignalCreateSignalV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignalChanDesc** | Pointer to **string** | Signal channel description | [optional] [default to ""]
**SignalChanName** | **string** | Signal channel name | [default to ""]

## Methods

### NewCreateTradingBotSignalCreateSignalV5Req

`func NewCreateTradingBotSignalCreateSignalV5Req(signalChanName string, ) *CreateTradingBotSignalCreateSignalV5Req`

NewCreateTradingBotSignalCreateSignalV5Req instantiates a new CreateTradingBotSignalCreateSignalV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradingBotSignalCreateSignalV5ReqWithDefaults

`func NewCreateTradingBotSignalCreateSignalV5ReqWithDefaults() *CreateTradingBotSignalCreateSignalV5Req`

NewCreateTradingBotSignalCreateSignalV5ReqWithDefaults instantiates a new CreateTradingBotSignalCreateSignalV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignalChanDesc

`func (o *CreateTradingBotSignalCreateSignalV5Req) GetSignalChanDesc() string`

GetSignalChanDesc returns the SignalChanDesc field if non-nil, zero value otherwise.

### GetSignalChanDescOk

`func (o *CreateTradingBotSignalCreateSignalV5Req) GetSignalChanDescOk() (*string, bool)`

GetSignalChanDescOk returns a tuple with the SignalChanDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalChanDesc

`func (o *CreateTradingBotSignalCreateSignalV5Req) SetSignalChanDesc(v string)`

SetSignalChanDesc sets SignalChanDesc field to given value.

### HasSignalChanDesc

`func (o *CreateTradingBotSignalCreateSignalV5Req) HasSignalChanDesc() bool`

HasSignalChanDesc returns a boolean if a field has been set.

### GetSignalChanName

`func (o *CreateTradingBotSignalCreateSignalV5Req) GetSignalChanName() string`

GetSignalChanName returns the SignalChanName field if non-nil, zero value otherwise.

### GetSignalChanNameOk

`func (o *CreateTradingBotSignalCreateSignalV5Req) GetSignalChanNameOk() (*string, bool)`

GetSignalChanNameOk returns a tuple with the SignalChanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalChanName

`func (o *CreateTradingBotSignalCreateSignalV5Req) SetSignalChanName(v string)`

SetSignalChanName sets SignalChanName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


