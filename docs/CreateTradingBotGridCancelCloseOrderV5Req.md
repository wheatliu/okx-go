# CreateTradingBotGridCancelCloseOrderV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoId** | **string** | Algo ID | [default to ""]
**OrdId** | **string** | Close position order ID | [default to ""]

## Methods

### NewCreateTradingBotGridCancelCloseOrderV5Req

`func NewCreateTradingBotGridCancelCloseOrderV5Req(algoId string, ordId string, ) *CreateTradingBotGridCancelCloseOrderV5Req`

NewCreateTradingBotGridCancelCloseOrderV5Req instantiates a new CreateTradingBotGridCancelCloseOrderV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradingBotGridCancelCloseOrderV5ReqWithDefaults

`func NewCreateTradingBotGridCancelCloseOrderV5ReqWithDefaults() *CreateTradingBotGridCancelCloseOrderV5Req`

NewCreateTradingBotGridCancelCloseOrderV5ReqWithDefaults instantiates a new CreateTradingBotGridCancelCloseOrderV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoId

`func (o *CreateTradingBotGridCancelCloseOrderV5Req) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *CreateTradingBotGridCancelCloseOrderV5Req) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *CreateTradingBotGridCancelCloseOrderV5Req) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.


### GetOrdId

`func (o *CreateTradingBotGridCancelCloseOrderV5Req) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *CreateTradingBotGridCancelCloseOrderV5Req) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *CreateTradingBotGridCancelCloseOrderV5Req) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


