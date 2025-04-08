# GetAccountRiskStateV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtRisk** | Pointer to **bool** | Account risk status in auto-borrow mode   true:  the account is currently in a specific risk state   false:  the account is currently not in a specific risk state | [optional] 
**AtRiskIdx** | Pointer to **[]string** | derivatives risk unit list | [optional] 
**AtRiskMgn** | Pointer to **[]string** | margin risk unit list | [optional] 
**Ts** | Pointer to **string** | Unix timestamp format in milliseconds, e.g.&#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetAccountRiskStateV5RespDataInner

`func NewGetAccountRiskStateV5RespDataInner() *GetAccountRiskStateV5RespDataInner`

NewGetAccountRiskStateV5RespDataInner instantiates a new GetAccountRiskStateV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountRiskStateV5RespDataInnerWithDefaults

`func NewGetAccountRiskStateV5RespDataInnerWithDefaults() *GetAccountRiskStateV5RespDataInner`

NewGetAccountRiskStateV5RespDataInnerWithDefaults instantiates a new GetAccountRiskStateV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtRisk

`func (o *GetAccountRiskStateV5RespDataInner) GetAtRisk() bool`

GetAtRisk returns the AtRisk field if non-nil, zero value otherwise.

### GetAtRiskOk

`func (o *GetAccountRiskStateV5RespDataInner) GetAtRiskOk() (*bool, bool)`

GetAtRiskOk returns a tuple with the AtRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtRisk

`func (o *GetAccountRiskStateV5RespDataInner) SetAtRisk(v bool)`

SetAtRisk sets AtRisk field to given value.

### HasAtRisk

`func (o *GetAccountRiskStateV5RespDataInner) HasAtRisk() bool`

HasAtRisk returns a boolean if a field has been set.

### GetAtRiskIdx

`func (o *GetAccountRiskStateV5RespDataInner) GetAtRiskIdx() []string`

GetAtRiskIdx returns the AtRiskIdx field if non-nil, zero value otherwise.

### GetAtRiskIdxOk

`func (o *GetAccountRiskStateV5RespDataInner) GetAtRiskIdxOk() (*[]string, bool)`

GetAtRiskIdxOk returns a tuple with the AtRiskIdx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtRiskIdx

`func (o *GetAccountRiskStateV5RespDataInner) SetAtRiskIdx(v []string)`

SetAtRiskIdx sets AtRiskIdx field to given value.

### HasAtRiskIdx

`func (o *GetAccountRiskStateV5RespDataInner) HasAtRiskIdx() bool`

HasAtRiskIdx returns a boolean if a field has been set.

### GetAtRiskMgn

`func (o *GetAccountRiskStateV5RespDataInner) GetAtRiskMgn() []string`

GetAtRiskMgn returns the AtRiskMgn field if non-nil, zero value otherwise.

### GetAtRiskMgnOk

`func (o *GetAccountRiskStateV5RespDataInner) GetAtRiskMgnOk() (*[]string, bool)`

GetAtRiskMgnOk returns a tuple with the AtRiskMgn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtRiskMgn

`func (o *GetAccountRiskStateV5RespDataInner) SetAtRiskMgn(v []string)`

SetAtRiskMgn sets AtRiskMgn field to given value.

### HasAtRiskMgn

`func (o *GetAccountRiskStateV5RespDataInner) HasAtRiskMgn() bool`

HasAtRiskMgn returns a boolean if a field has been set.

### GetTs

`func (o *GetAccountRiskStateV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetAccountRiskStateV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetAccountRiskStateV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetAccountRiskStateV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


