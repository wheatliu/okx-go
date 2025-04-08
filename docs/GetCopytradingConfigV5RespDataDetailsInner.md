# GetCopytradingConfigV5RespDataDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopyTraderNum** | Pointer to **string** | Current number of copy traders | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;SWAP&#x60; | [optional] [default to ""]
**MaxCopyTraderNum** | Pointer to **string** | Maximum number of copy traders | [optional] [default to ""]
**ProfitSharingRatio** | Pointer to **string** | Profit sharing ratio.   Only applicable to lead trader, or it will be \&quot;\&quot;. 0.1 represents 10% | [optional] [default to ""]
**RoleType** | Pointer to **string** | Role type  &#x60;0&#x60;: General user  &#x60;1&#x60;: Leading trader  &#x60;2&#x60;: Copy trader | [optional] [default to ""]

## Methods

### NewGetCopytradingConfigV5RespDataDetailsInner

`func NewGetCopytradingConfigV5RespDataDetailsInner() *GetCopytradingConfigV5RespDataDetailsInner`

NewGetCopytradingConfigV5RespDataDetailsInner instantiates a new GetCopytradingConfigV5RespDataDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCopytradingConfigV5RespDataDetailsInnerWithDefaults

`func NewGetCopytradingConfigV5RespDataDetailsInnerWithDefaults() *GetCopytradingConfigV5RespDataDetailsInner`

NewGetCopytradingConfigV5RespDataDetailsInnerWithDefaults instantiates a new GetCopytradingConfigV5RespDataDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyTraderNum

`func (o *GetCopytradingConfigV5RespDataDetailsInner) GetCopyTraderNum() string`

GetCopyTraderNum returns the CopyTraderNum field if non-nil, zero value otherwise.

### GetCopyTraderNumOk

`func (o *GetCopytradingConfigV5RespDataDetailsInner) GetCopyTraderNumOk() (*string, bool)`

GetCopyTraderNumOk returns a tuple with the CopyTraderNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTraderNum

`func (o *GetCopytradingConfigV5RespDataDetailsInner) SetCopyTraderNum(v string)`

SetCopyTraderNum sets CopyTraderNum field to given value.

### HasCopyTraderNum

`func (o *GetCopytradingConfigV5RespDataDetailsInner) HasCopyTraderNum() bool`

HasCopyTraderNum returns a boolean if a field has been set.

### GetInstType

`func (o *GetCopytradingConfigV5RespDataDetailsInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetCopytradingConfigV5RespDataDetailsInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetCopytradingConfigV5RespDataDetailsInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetCopytradingConfigV5RespDataDetailsInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetMaxCopyTraderNum

`func (o *GetCopytradingConfigV5RespDataDetailsInner) GetMaxCopyTraderNum() string`

GetMaxCopyTraderNum returns the MaxCopyTraderNum field if non-nil, zero value otherwise.

### GetMaxCopyTraderNumOk

`func (o *GetCopytradingConfigV5RespDataDetailsInner) GetMaxCopyTraderNumOk() (*string, bool)`

GetMaxCopyTraderNumOk returns a tuple with the MaxCopyTraderNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCopyTraderNum

`func (o *GetCopytradingConfigV5RespDataDetailsInner) SetMaxCopyTraderNum(v string)`

SetMaxCopyTraderNum sets MaxCopyTraderNum field to given value.

### HasMaxCopyTraderNum

`func (o *GetCopytradingConfigV5RespDataDetailsInner) HasMaxCopyTraderNum() bool`

HasMaxCopyTraderNum returns a boolean if a field has been set.

### GetProfitSharingRatio

`func (o *GetCopytradingConfigV5RespDataDetailsInner) GetProfitSharingRatio() string`

GetProfitSharingRatio returns the ProfitSharingRatio field if non-nil, zero value otherwise.

### GetProfitSharingRatioOk

`func (o *GetCopytradingConfigV5RespDataDetailsInner) GetProfitSharingRatioOk() (*string, bool)`

GetProfitSharingRatioOk returns a tuple with the ProfitSharingRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitSharingRatio

`func (o *GetCopytradingConfigV5RespDataDetailsInner) SetProfitSharingRatio(v string)`

SetProfitSharingRatio sets ProfitSharingRatio field to given value.

### HasProfitSharingRatio

`func (o *GetCopytradingConfigV5RespDataDetailsInner) HasProfitSharingRatio() bool`

HasProfitSharingRatio returns a boolean if a field has been set.

### GetRoleType

`func (o *GetCopytradingConfigV5RespDataDetailsInner) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *GetCopytradingConfigV5RespDataDetailsInner) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *GetCopytradingConfigV5RespDataDetailsInner) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *GetCopytradingConfigV5RespDataDetailsInner) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


