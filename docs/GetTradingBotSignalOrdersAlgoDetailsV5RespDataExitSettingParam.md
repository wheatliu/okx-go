# GetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SlPct** | Pointer to **string** | Stop-loss percentage | [optional] [default to ""]
**TpPct** | Pointer to **string** | Take-profit percentage | [optional] [default to ""]
**TpSlType** | Pointer to **string** | Type of set the take-profit and stop-loss trigger price  &#x60;pnl&#x60;: Based on the estimated profit and loss percentage from the entry point  &#x60;price&#x60;: Based on price increase or decrease from the crypto’s entry price | [optional] [default to ""]

## Methods

### NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParam

`func NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParam() *GetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParam`

NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParam instantiates a new GetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParamWithDefaults

`func NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParamWithDefaults() *GetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParam`

NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParamWithDefaults instantiates a new GetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlPct

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParam) GetSlPct() string`

GetSlPct returns the SlPct field if non-nil, zero value otherwise.

### GetSlPctOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParam) GetSlPctOk() (*string, bool)`

GetSlPctOk returns a tuple with the SlPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlPct

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParam) SetSlPct(v string)`

SetSlPct sets SlPct field to given value.

### HasSlPct

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParam) HasSlPct() bool`

HasSlPct returns a boolean if a field has been set.

### GetTpPct

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParam) GetTpPct() string`

GetTpPct returns the TpPct field if non-nil, zero value otherwise.

### GetTpPctOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParam) GetTpPctOk() (*string, bool)`

GetTpPctOk returns a tuple with the TpPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpPct

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParam) SetTpPct(v string)`

SetTpPct sets TpPct field to given value.

### HasTpPct

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParam) HasTpPct() bool`

HasTpPct returns a boolean if a field has been set.

### GetTpSlType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParam) GetTpSlType() string`

GetTpSlType returns the TpSlType field if non-nil, zero value otherwise.

### GetTpSlTypeOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParam) GetTpSlTypeOk() (*string, bool)`

GetTpSlTypeOk returns a tuple with the TpSlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpSlType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParam) SetTpSlType(v string)`

SetTpSlType sets TpSlType field to given value.

### HasTpSlType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataExitSettingParam) HasTpSlType() bool`

HasTpSlType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


