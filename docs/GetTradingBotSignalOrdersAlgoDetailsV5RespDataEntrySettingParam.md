# GetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowMultipleEntry** | Pointer to **bool** | Whether or not allow multiple entries in the same direction for the same trading pairs | [optional] 
**Amt** | Pointer to **string** | Amount per order  Only applicable to &#x60;entryType&#x60; in &#x60;2&#x60;/&#x60;3&#x60; | [optional] [default to ""]
**EntryType** | Pointer to **string** | Entry type  &#x60;1&#x60;: TradingView signal  &#x60;2&#x60;: Fixed margin  &#x60;3&#x60;: Contracts  &#x60;4&#x60;: Percentage of free margin  &#x60;5&#x60;: Percentage of the initial invested margin | [optional] [default to ""]
**Ratio** | Pointer to **string** | Amount ratio per order  Only applicable to &#x60;entryType&#x60; in &#x60;4&#x60;/&#x60;5&#x60; | [optional] [default to ""]

## Methods

### NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam

`func NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam() *GetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam`

NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam instantiates a new GetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParamWithDefaults

`func NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParamWithDefaults() *GetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam`

NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParamWithDefaults instantiates a new GetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowMultipleEntry

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam) GetAllowMultipleEntry() bool`

GetAllowMultipleEntry returns the AllowMultipleEntry field if non-nil, zero value otherwise.

### GetAllowMultipleEntryOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam) GetAllowMultipleEntryOk() (*bool, bool)`

GetAllowMultipleEntryOk returns a tuple with the AllowMultipleEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleEntry

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam) SetAllowMultipleEntry(v bool)`

SetAllowMultipleEntry sets AllowMultipleEntry field to given value.

### HasAllowMultipleEntry

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam) HasAllowMultipleEntry() bool`

HasAllowMultipleEntry returns a boolean if a field has been set.

### GetAmt

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetEntryType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam) GetEntryType() string`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam) GetEntryTypeOk() (*string, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam) SetEntryType(v string)`

SetEntryType sets EntryType field to given value.

### HasEntryType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam) HasEntryType() bool`

HasEntryType returns a boolean if a field has been set.

### GetRatio

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam) GetRatio() string`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam) GetRatioOk() (*string, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam) SetRatio(v string)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataEntrySettingParam) HasRatio() bool`

HasRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


