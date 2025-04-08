# GetAccountRiskStateV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtRisk** | Pointer to **bool** | Account risk status in auto-borrow mode   true:  the account is currently in a specific risk state   false:  the account is currently not in a specific risk state | [optional] 
**AtRiskIdx** | Pointer to **[]string** | derivatives risk unit list | [optional] 
**AtRiskMgn** | Pointer to **[]string** | margin risk unit list | [optional] 
**Ts** | Pointer to **string** | Unix timestamp format in milliseconds, e.g.&#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetAccountRiskStateV5RespData

`func NewGetAccountRiskStateV5RespData() *GetAccountRiskStateV5RespData`

NewGetAccountRiskStateV5RespData instantiates a new GetAccountRiskStateV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountRiskStateV5RespDataWithDefaults

`func NewGetAccountRiskStateV5RespDataWithDefaults() *GetAccountRiskStateV5RespData`

NewGetAccountRiskStateV5RespDataWithDefaults instantiates a new GetAccountRiskStateV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtRisk

`func (o *GetAccountRiskStateV5RespData) GetAtRisk() bool`

GetAtRisk returns the AtRisk field if non-nil, zero value otherwise.

### GetAtRiskOk

`func (o *GetAccountRiskStateV5RespData) GetAtRiskOk() (*bool, bool)`

GetAtRiskOk returns a tuple with the AtRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtRisk

`func (o *GetAccountRiskStateV5RespData) SetAtRisk(v bool)`

SetAtRisk sets AtRisk field to given value.

### HasAtRisk

`func (o *GetAccountRiskStateV5RespData) HasAtRisk() bool`

HasAtRisk returns a boolean if a field has been set.

### GetAtRiskIdx

`func (o *GetAccountRiskStateV5RespData) GetAtRiskIdx() []string`

GetAtRiskIdx returns the AtRiskIdx field if non-nil, zero value otherwise.

### GetAtRiskIdxOk

`func (o *GetAccountRiskStateV5RespData) GetAtRiskIdxOk() (*[]string, bool)`

GetAtRiskIdxOk returns a tuple with the AtRiskIdx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtRiskIdx

`func (o *GetAccountRiskStateV5RespData) SetAtRiskIdx(v []string)`

SetAtRiskIdx sets AtRiskIdx field to given value.

### HasAtRiskIdx

`func (o *GetAccountRiskStateV5RespData) HasAtRiskIdx() bool`

HasAtRiskIdx returns a boolean if a field has been set.

### GetAtRiskMgn

`func (o *GetAccountRiskStateV5RespData) GetAtRiskMgn() []string`

GetAtRiskMgn returns the AtRiskMgn field if non-nil, zero value otherwise.

### GetAtRiskMgnOk

`func (o *GetAccountRiskStateV5RespData) GetAtRiskMgnOk() (*[]string, bool)`

GetAtRiskMgnOk returns a tuple with the AtRiskMgn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtRiskMgn

`func (o *GetAccountRiskStateV5RespData) SetAtRiskMgn(v []string)`

SetAtRiskMgn sets AtRiskMgn field to given value.

### HasAtRiskMgn

`func (o *GetAccountRiskStateV5RespData) HasAtRiskMgn() bool`

HasAtRiskMgn returns a boolean if a field has been set.

### GetTs

`func (o *GetAccountRiskStateV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetAccountRiskStateV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetAccountRiskStateV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetAccountRiskStateV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


