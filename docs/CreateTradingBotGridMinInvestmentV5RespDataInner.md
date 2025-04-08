# CreateTradingBotGridMinInvestmentV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinInvestmentData** | Pointer to [**[]CreateTradingBotGridMinInvestmentV5RespDataInnerMinInvestmentDataInner**](CreateTradingBotGridMinInvestmentV5RespDataInnerMinInvestmentDataInner.md) | Minimum invest Data | [optional] 
**SingleAmt** | Pointer to **string** | Single grid trading amount  In terms of &#x60;spot grid&#x60;, the unit is &#x60;quote currency&#x60;  In terms of &#x60;contract grid&#x60;, the unit is &#x60;contract&#x60; | [optional] [default to ""]

## Methods

### NewCreateTradingBotGridMinInvestmentV5RespDataInner

`func NewCreateTradingBotGridMinInvestmentV5RespDataInner() *CreateTradingBotGridMinInvestmentV5RespDataInner`

NewCreateTradingBotGridMinInvestmentV5RespDataInner instantiates a new CreateTradingBotGridMinInvestmentV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradingBotGridMinInvestmentV5RespDataInnerWithDefaults

`func NewCreateTradingBotGridMinInvestmentV5RespDataInnerWithDefaults() *CreateTradingBotGridMinInvestmentV5RespDataInner`

NewCreateTradingBotGridMinInvestmentV5RespDataInnerWithDefaults instantiates a new CreateTradingBotGridMinInvestmentV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinInvestmentData

`func (o *CreateTradingBotGridMinInvestmentV5RespDataInner) GetMinInvestmentData() []CreateTradingBotGridMinInvestmentV5RespDataInnerMinInvestmentDataInner`

GetMinInvestmentData returns the MinInvestmentData field if non-nil, zero value otherwise.

### GetMinInvestmentDataOk

`func (o *CreateTradingBotGridMinInvestmentV5RespDataInner) GetMinInvestmentDataOk() (*[]CreateTradingBotGridMinInvestmentV5RespDataInnerMinInvestmentDataInner, bool)`

GetMinInvestmentDataOk returns a tuple with the MinInvestmentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInvestmentData

`func (o *CreateTradingBotGridMinInvestmentV5RespDataInner) SetMinInvestmentData(v []CreateTradingBotGridMinInvestmentV5RespDataInnerMinInvestmentDataInner)`

SetMinInvestmentData sets MinInvestmentData field to given value.

### HasMinInvestmentData

`func (o *CreateTradingBotGridMinInvestmentV5RespDataInner) HasMinInvestmentData() bool`

HasMinInvestmentData returns a boolean if a field has been set.

### GetSingleAmt

`func (o *CreateTradingBotGridMinInvestmentV5RespDataInner) GetSingleAmt() string`

GetSingleAmt returns the SingleAmt field if non-nil, zero value otherwise.

### GetSingleAmtOk

`func (o *CreateTradingBotGridMinInvestmentV5RespDataInner) GetSingleAmtOk() (*string, bool)`

GetSingleAmtOk returns a tuple with the SingleAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleAmt

`func (o *CreateTradingBotGridMinInvestmentV5RespDataInner) SetSingleAmt(v string)`

SetSingleAmt sets SingleAmt field to given value.

### HasSingleAmt

`func (o *CreateTradingBotGridMinInvestmentV5RespDataInner) HasSingleAmt() bool`

HasSingleAmt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


