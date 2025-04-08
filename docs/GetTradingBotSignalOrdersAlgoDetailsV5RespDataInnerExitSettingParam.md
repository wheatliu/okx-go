# GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SlPct** | Pointer to **string** | Stop-loss percentage | [optional] [default to ""]
**TpPct** | Pointer to **string** | Take-profit percentage | [optional] [default to ""]
**TpSlType** | Pointer to **string** | Type of set the take-profit and stop-loss trigger price  &#x60;pnl&#x60;: Based on the estimated profit and loss percentage from the entry point  &#x60;price&#x60;: Based on price increase or decrease from the crypto’s entry price | [optional] [default to ""]

## Methods

### NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam

`func NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam() *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam`

NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam instantiates a new GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParamWithDefaults

`func NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParamWithDefaults() *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam`

NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParamWithDefaults instantiates a new GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlPct

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) GetSlPct() string`

GetSlPct returns the SlPct field if non-nil, zero value otherwise.

### GetSlPctOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) GetSlPctOk() (*string, bool)`

GetSlPctOk returns a tuple with the SlPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlPct

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) SetSlPct(v string)`

SetSlPct sets SlPct field to given value.

### HasSlPct

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) HasSlPct() bool`

HasSlPct returns a boolean if a field has been set.

### GetTpPct

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) GetTpPct() string`

GetTpPct returns the TpPct field if non-nil, zero value otherwise.

### GetTpPctOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) GetTpPctOk() (*string, bool)`

GetTpPctOk returns a tuple with the TpPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpPct

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) SetTpPct(v string)`

SetTpPct sets TpPct field to given value.

### HasTpPct

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) HasTpPct() bool`

HasTpPct returns a boolean if a field has been set.

### GetTpSlType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) GetTpSlType() string`

GetTpSlType returns the TpSlType field if non-nil, zero value otherwise.

### GetTpSlTypeOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) GetTpSlTypeOk() (*string, bool)`

GetTpSlTypeOk returns a tuple with the TpSlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpSlType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) SetTpSlType(v string)`

SetTpSlType sets TpSlType field to given value.

### HasTpSlType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) HasTpSlType() bool`

HasTpSlType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


