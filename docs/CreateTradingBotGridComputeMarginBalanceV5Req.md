# CreateTradingBotGridComputeMarginBalanceV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoId** | **string** | Algo ID | [default to ""]
**Amt** | Pointer to **string** | Adjust margin balance amount  Default is zero. | [optional] [default to ""]
**Type** | **string** | Adjust margin balance type  &#x60;add&#x60; &#x60;reduce&#x60; | [default to ""]

## Methods

### NewCreateTradingBotGridComputeMarginBalanceV5Req

`func NewCreateTradingBotGridComputeMarginBalanceV5Req(algoId string, type_ string, ) *CreateTradingBotGridComputeMarginBalanceV5Req`

NewCreateTradingBotGridComputeMarginBalanceV5Req instantiates a new CreateTradingBotGridComputeMarginBalanceV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradingBotGridComputeMarginBalanceV5ReqWithDefaults

`func NewCreateTradingBotGridComputeMarginBalanceV5ReqWithDefaults() *CreateTradingBotGridComputeMarginBalanceV5Req`

NewCreateTradingBotGridComputeMarginBalanceV5ReqWithDefaults instantiates a new CreateTradingBotGridComputeMarginBalanceV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoId

`func (o *CreateTradingBotGridComputeMarginBalanceV5Req) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *CreateTradingBotGridComputeMarginBalanceV5Req) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *CreateTradingBotGridComputeMarginBalanceV5Req) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.


### GetAmt

`func (o *CreateTradingBotGridComputeMarginBalanceV5Req) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *CreateTradingBotGridComputeMarginBalanceV5Req) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *CreateTradingBotGridComputeMarginBalanceV5Req) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *CreateTradingBotGridComputeMarginBalanceV5Req) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetType

`func (o *CreateTradingBotGridComputeMarginBalanceV5Req) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTradingBotGridComputeMarginBalanceV5Req) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTradingBotGridComputeMarginBalanceV5Req) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


