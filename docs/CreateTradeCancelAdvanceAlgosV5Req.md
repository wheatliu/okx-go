# CreateTradeCancelAdvanceAlgosV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoId** | **string** | Algo ID | [default to ""]
**InstId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to ""]

## Methods

### NewCreateTradeCancelAdvanceAlgosV5Req

`func NewCreateTradeCancelAdvanceAlgosV5Req(algoId string, instId string, ) *CreateTradeCancelAdvanceAlgosV5Req`

NewCreateTradeCancelAdvanceAlgosV5Req instantiates a new CreateTradeCancelAdvanceAlgosV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeCancelAdvanceAlgosV5ReqWithDefaults

`func NewCreateTradeCancelAdvanceAlgosV5ReqWithDefaults() *CreateTradeCancelAdvanceAlgosV5Req`

NewCreateTradeCancelAdvanceAlgosV5ReqWithDefaults instantiates a new CreateTradeCancelAdvanceAlgosV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoId

`func (o *CreateTradeCancelAdvanceAlgosV5Req) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *CreateTradeCancelAdvanceAlgosV5Req) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *CreateTradeCancelAdvanceAlgosV5Req) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.


### GetInstId

`func (o *CreateTradeCancelAdvanceAlgosV5Req) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateTradeCancelAdvanceAlgosV5Req) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateTradeCancelAdvanceAlgosV5Req) SetInstId(v string)`

SetInstId sets InstId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


