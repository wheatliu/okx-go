# CreateTradingBotGridMarginBalanceV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoId** | **string** | Algo ID | [default to ""]
**Amt** | Pointer to **string** | Adjust margin balance amount  Either &#x60;amt&#x60; or &#x60;percent&#x60; is required. | [optional] [default to ""]
**Percent** | Pointer to **string** | Adjust margin balance percentage | [optional] [default to ""]
**Type** | **string** | Adjust margin balance type  &#x60;add&#x60; &#x60;reduce&#x60; | [default to ""]

## Methods

### NewCreateTradingBotGridMarginBalanceV5Req

`func NewCreateTradingBotGridMarginBalanceV5Req(algoId string, type_ string, ) *CreateTradingBotGridMarginBalanceV5Req`

NewCreateTradingBotGridMarginBalanceV5Req instantiates a new CreateTradingBotGridMarginBalanceV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradingBotGridMarginBalanceV5ReqWithDefaults

`func NewCreateTradingBotGridMarginBalanceV5ReqWithDefaults() *CreateTradingBotGridMarginBalanceV5Req`

NewCreateTradingBotGridMarginBalanceV5ReqWithDefaults instantiates a new CreateTradingBotGridMarginBalanceV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoId

`func (o *CreateTradingBotGridMarginBalanceV5Req) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *CreateTradingBotGridMarginBalanceV5Req) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *CreateTradingBotGridMarginBalanceV5Req) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.


### GetAmt

`func (o *CreateTradingBotGridMarginBalanceV5Req) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *CreateTradingBotGridMarginBalanceV5Req) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *CreateTradingBotGridMarginBalanceV5Req) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *CreateTradingBotGridMarginBalanceV5Req) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetPercent

`func (o *CreateTradingBotGridMarginBalanceV5Req) GetPercent() string`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *CreateTradingBotGridMarginBalanceV5Req) GetPercentOk() (*string, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *CreateTradingBotGridMarginBalanceV5Req) SetPercent(v string)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *CreateTradingBotGridMarginBalanceV5Req) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetType

`func (o *CreateTradingBotGridMarginBalanceV5Req) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTradingBotGridMarginBalanceV5Req) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTradingBotGridMarginBalanceV5Req) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


