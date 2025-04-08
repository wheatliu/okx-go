# CreateTradingBotSignalOrderAlgoV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntrySettingParam** | Pointer to **string** | Entry setting | [optional] [default to ""]
**ExitSettingParam** | Pointer to **string** | Exit setting | [optional] [default to ""]
**IncludeAll** | Pointer to **bool** | Whether to include all USDT-margined contract.The default value is &#x60;false&#x60;. &#x60;true&#x60;: include &#x60;false&#x60; : exclude | [optional] 
**InstIds** | Pointer to **string** | Instrument IDs. Single currency or multiple currencies separated with comma. When &#x60;includeAll&#x60; is &#x60;true&#x60;, it is ignored | [optional] [default to ""]
**InvestAmt** | **string** | Investment amount | [default to ""]
**Lever** | **string** | Leverage  Only applicable to &#x60;contract signal&#x60; | [default to ""]
**Ratio** | Pointer to **string** | Price offset ratio, calculate the limit price as a percentage offset from the best bid/ask price.  Only applicable to &#x60;subOrdType&#x60; is &#x60;limit&#x60; order | [optional] [default to ""]
**SignalChanId** | **string** | Signal channel Id | [default to ""]
**SubOrdType** | **string** | Sub order type &#x60;1&#x60;：limit order &#x60;2&#x60;：market order &#x60;9&#x60;：tradingView signal | [default to ""]

## Methods

### NewCreateTradingBotSignalOrderAlgoV5Req

`func NewCreateTradingBotSignalOrderAlgoV5Req(investAmt string, lever string, signalChanId string, subOrdType string, ) *CreateTradingBotSignalOrderAlgoV5Req`

NewCreateTradingBotSignalOrderAlgoV5Req instantiates a new CreateTradingBotSignalOrderAlgoV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradingBotSignalOrderAlgoV5ReqWithDefaults

`func NewCreateTradingBotSignalOrderAlgoV5ReqWithDefaults() *CreateTradingBotSignalOrderAlgoV5Req`

NewCreateTradingBotSignalOrderAlgoV5ReqWithDefaults instantiates a new CreateTradingBotSignalOrderAlgoV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntrySettingParam

`func (o *CreateTradingBotSignalOrderAlgoV5Req) GetEntrySettingParam() string`

GetEntrySettingParam returns the EntrySettingParam field if non-nil, zero value otherwise.

### GetEntrySettingParamOk

`func (o *CreateTradingBotSignalOrderAlgoV5Req) GetEntrySettingParamOk() (*string, bool)`

GetEntrySettingParamOk returns a tuple with the EntrySettingParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrySettingParam

`func (o *CreateTradingBotSignalOrderAlgoV5Req) SetEntrySettingParam(v string)`

SetEntrySettingParam sets EntrySettingParam field to given value.

### HasEntrySettingParam

`func (o *CreateTradingBotSignalOrderAlgoV5Req) HasEntrySettingParam() bool`

HasEntrySettingParam returns a boolean if a field has been set.

### GetExitSettingParam

`func (o *CreateTradingBotSignalOrderAlgoV5Req) GetExitSettingParam() string`

GetExitSettingParam returns the ExitSettingParam field if non-nil, zero value otherwise.

### GetExitSettingParamOk

`func (o *CreateTradingBotSignalOrderAlgoV5Req) GetExitSettingParamOk() (*string, bool)`

GetExitSettingParamOk returns a tuple with the ExitSettingParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitSettingParam

`func (o *CreateTradingBotSignalOrderAlgoV5Req) SetExitSettingParam(v string)`

SetExitSettingParam sets ExitSettingParam field to given value.

### HasExitSettingParam

`func (o *CreateTradingBotSignalOrderAlgoV5Req) HasExitSettingParam() bool`

HasExitSettingParam returns a boolean if a field has been set.

### GetIncludeAll

`func (o *CreateTradingBotSignalOrderAlgoV5Req) GetIncludeAll() bool`

GetIncludeAll returns the IncludeAll field if non-nil, zero value otherwise.

### GetIncludeAllOk

`func (o *CreateTradingBotSignalOrderAlgoV5Req) GetIncludeAllOk() (*bool, bool)`

GetIncludeAllOk returns a tuple with the IncludeAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAll

`func (o *CreateTradingBotSignalOrderAlgoV5Req) SetIncludeAll(v bool)`

SetIncludeAll sets IncludeAll field to given value.

### HasIncludeAll

`func (o *CreateTradingBotSignalOrderAlgoV5Req) HasIncludeAll() bool`

HasIncludeAll returns a boolean if a field has been set.

### GetInstIds

`func (o *CreateTradingBotSignalOrderAlgoV5Req) GetInstIds() string`

GetInstIds returns the InstIds field if non-nil, zero value otherwise.

### GetInstIdsOk

`func (o *CreateTradingBotSignalOrderAlgoV5Req) GetInstIdsOk() (*string, bool)`

GetInstIdsOk returns a tuple with the InstIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstIds

`func (o *CreateTradingBotSignalOrderAlgoV5Req) SetInstIds(v string)`

SetInstIds sets InstIds field to given value.

### HasInstIds

`func (o *CreateTradingBotSignalOrderAlgoV5Req) HasInstIds() bool`

HasInstIds returns a boolean if a field has been set.

### GetInvestAmt

`func (o *CreateTradingBotSignalOrderAlgoV5Req) GetInvestAmt() string`

GetInvestAmt returns the InvestAmt field if non-nil, zero value otherwise.

### GetInvestAmtOk

`func (o *CreateTradingBotSignalOrderAlgoV5Req) GetInvestAmtOk() (*string, bool)`

GetInvestAmtOk returns a tuple with the InvestAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestAmt

`func (o *CreateTradingBotSignalOrderAlgoV5Req) SetInvestAmt(v string)`

SetInvestAmt sets InvestAmt field to given value.


### GetLever

`func (o *CreateTradingBotSignalOrderAlgoV5Req) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *CreateTradingBotSignalOrderAlgoV5Req) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *CreateTradingBotSignalOrderAlgoV5Req) SetLever(v string)`

SetLever sets Lever field to given value.


### GetRatio

`func (o *CreateTradingBotSignalOrderAlgoV5Req) GetRatio() string`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *CreateTradingBotSignalOrderAlgoV5Req) GetRatioOk() (*string, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *CreateTradingBotSignalOrderAlgoV5Req) SetRatio(v string)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *CreateTradingBotSignalOrderAlgoV5Req) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetSignalChanId

`func (o *CreateTradingBotSignalOrderAlgoV5Req) GetSignalChanId() string`

GetSignalChanId returns the SignalChanId field if non-nil, zero value otherwise.

### GetSignalChanIdOk

`func (o *CreateTradingBotSignalOrderAlgoV5Req) GetSignalChanIdOk() (*string, bool)`

GetSignalChanIdOk returns a tuple with the SignalChanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalChanId

`func (o *CreateTradingBotSignalOrderAlgoV5Req) SetSignalChanId(v string)`

SetSignalChanId sets SignalChanId field to given value.


### GetSubOrdType

`func (o *CreateTradingBotSignalOrderAlgoV5Req) GetSubOrdType() string`

GetSubOrdType returns the SubOrdType field if non-nil, zero value otherwise.

### GetSubOrdTypeOk

`func (o *CreateTradingBotSignalOrderAlgoV5Req) GetSubOrdTypeOk() (*string, bool)`

GetSubOrdTypeOk returns a tuple with the SubOrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrdType

`func (o *CreateTradingBotSignalOrderAlgoV5Req) SetSubOrdType(v string)`

SetSubOrdType sets SubOrdType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


