# CreateTradingBotSignalMarginBalanceV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoId** | **string** | Algo ID | [default to ""]
**AllowReinvest** | Pointer to **bool** | Whether to reinvest with newly added margin. The default value is &#x60;false&#x60;.   &#x60;false&#x60;:it will be used as passive margin to prevent liquidation and will not be used as active investment  &#x60;true&#x60;:the margin added here will furthermore be accounted for in calculations of your total investment amount, and furthermore your order size。  Only applicable to your signal comes in with an “investmentType” of “percentage_investment” | [optional] 
**Amt** | **string** | Adjust margin balance amount  Either &#x60;amt&#x60; or &#x60;percent&#x60; is required. | [default to ""]
**Type** | **string** | Adjust margin balance type  &#x60;add&#x60; &#x60;reduce&#x60; | [default to ""]

## Methods

### NewCreateTradingBotSignalMarginBalanceV5Req

`func NewCreateTradingBotSignalMarginBalanceV5Req(algoId string, amt string, type_ string, ) *CreateTradingBotSignalMarginBalanceV5Req`

NewCreateTradingBotSignalMarginBalanceV5Req instantiates a new CreateTradingBotSignalMarginBalanceV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradingBotSignalMarginBalanceV5ReqWithDefaults

`func NewCreateTradingBotSignalMarginBalanceV5ReqWithDefaults() *CreateTradingBotSignalMarginBalanceV5Req`

NewCreateTradingBotSignalMarginBalanceV5ReqWithDefaults instantiates a new CreateTradingBotSignalMarginBalanceV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoId

`func (o *CreateTradingBotSignalMarginBalanceV5Req) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *CreateTradingBotSignalMarginBalanceV5Req) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *CreateTradingBotSignalMarginBalanceV5Req) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.


### GetAllowReinvest

`func (o *CreateTradingBotSignalMarginBalanceV5Req) GetAllowReinvest() bool`

GetAllowReinvest returns the AllowReinvest field if non-nil, zero value otherwise.

### GetAllowReinvestOk

`func (o *CreateTradingBotSignalMarginBalanceV5Req) GetAllowReinvestOk() (*bool, bool)`

GetAllowReinvestOk returns a tuple with the AllowReinvest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowReinvest

`func (o *CreateTradingBotSignalMarginBalanceV5Req) SetAllowReinvest(v bool)`

SetAllowReinvest sets AllowReinvest field to given value.

### HasAllowReinvest

`func (o *CreateTradingBotSignalMarginBalanceV5Req) HasAllowReinvest() bool`

HasAllowReinvest returns a boolean if a field has been set.

### GetAmt

`func (o *CreateTradingBotSignalMarginBalanceV5Req) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *CreateTradingBotSignalMarginBalanceV5Req) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *CreateTradingBotSignalMarginBalanceV5Req) SetAmt(v string)`

SetAmt sets Amt field to given value.


### GetType

`func (o *CreateTradingBotSignalMarginBalanceV5Req) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTradingBotSignalMarginBalanceV5Req) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTradingBotSignalMarginBalanceV5Req) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


