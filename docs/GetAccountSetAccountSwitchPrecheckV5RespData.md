# GetAccountSetAccountSwitchPrecheckV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctLv** | Pointer to **string** | Account mode   &#x60;1&#x60;: Spot mode   &#x60;2&#x60;: Spot and futures mode   &#x60;3&#x60;: Multi-currency margin code   &#x60;4&#x60;: Portfolio margin mode   Applicable to all scenarios | [optional] [default to ""]
**CurAcctLv** | Pointer to **string** | Account mode   &#x60;1&#x60;: Spot mode   &#x60;2&#x60;: Spot and futures mode   &#x60;3&#x60;: Multi-currency margin code   &#x60;4&#x60;: Portfolio margin mode   Applicable to all scenarios | [optional] [default to ""]
**MgnAft** | Pointer to [**GetAccountSetAccountSwitchPrecheckV5RespDataMgnAft**](GetAccountSetAccountSwitchPrecheckV5RespDataMgnAft.md) |  | [optional] 
**MgnBf** | Pointer to [**GetAccountSetAccountSwitchPrecheckV5RespDataMgnBf**](GetAccountSetAccountSwitchPrecheckV5RespDataMgnBf.md) |  | [optional] 
**PosList** | Pointer to [**[]GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner**](GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner.md) | Cross margin contract position list   Applicable when curAcctLv is &#x60;4&#x60;, acctLv is &#x60;2/3&#x60; and user has cross margin contract positions   Applicable when sCode is &#x60;0/3/4&#x60; | [optional] 
**PosTierCheck** | Pointer to [**[]GetAccountSetAccountSwitchPrecheckV5RespDataPosTierCheckInner**](GetAccountSetAccountSwitchPrecheckV5RespDataPosTierCheckInner.md) | Cross margin contract positions that don&#39;t pass the position tier check   Only applicable when sCode is &#x60;4&#x60; | [optional] 
**RiskOffsetType** | Pointer to **string** | Risk offset type   &#x60;1&#x60;: Spot-derivatives (USDT) risk offset   &#x60;2&#x60;: Spot-derivatives (Crypto) risk offset   &#x60;3&#x60;: Derivatives only mode   &#x60;4&#x60;: Spot-derivatives (USDC) risk offset   Applicable when acctLv is &#x60;4&#x60;, return \&quot;\&quot; for other scenarios   If the user preset before, it will use the user&#39;s specified value; if not, the default value &#x60;3&#x60; will be applied(Deprecated) | [optional] [default to ""]
**SCode** | Pointer to **string** | Check code   &#x60;0&#x60;: pass all checks   &#x60;1&#x60;: unmatched information   &#x60;3&#x60;: leverage setting is not finished   &#x60;4&#x60;: position tier or margin check is not passed | [optional] [default to ""]
**UnmatchedInfoCheck** | Pointer to [**[]GetAccountSetAccountSwitchPrecheckV5RespDataUnmatchedInfoCheckInner**](GetAccountSetAccountSwitchPrecheckV5RespDataUnmatchedInfoCheckInner.md) | Unmatched information list   Applicable when sCode is &#x60;1&#x60;, indicating there is unmatched information; return [] for other scenarios | [optional] 

## Methods

### NewGetAccountSetAccountSwitchPrecheckV5RespData

`func NewGetAccountSetAccountSwitchPrecheckV5RespData() *GetAccountSetAccountSwitchPrecheckV5RespData`

NewGetAccountSetAccountSwitchPrecheckV5RespData instantiates a new GetAccountSetAccountSwitchPrecheckV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountSetAccountSwitchPrecheckV5RespDataWithDefaults

`func NewGetAccountSetAccountSwitchPrecheckV5RespDataWithDefaults() *GetAccountSetAccountSwitchPrecheckV5RespData`

NewGetAccountSetAccountSwitchPrecheckV5RespDataWithDefaults instantiates a new GetAccountSetAccountSwitchPrecheckV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctLv

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) GetAcctLv() string`

GetAcctLv returns the AcctLv field if non-nil, zero value otherwise.

### GetAcctLvOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) GetAcctLvOk() (*string, bool)`

GetAcctLvOk returns a tuple with the AcctLv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctLv

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) SetAcctLv(v string)`

SetAcctLv sets AcctLv field to given value.

### HasAcctLv

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) HasAcctLv() bool`

HasAcctLv returns a boolean if a field has been set.

### GetCurAcctLv

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) GetCurAcctLv() string`

GetCurAcctLv returns the CurAcctLv field if non-nil, zero value otherwise.

### GetCurAcctLvOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) GetCurAcctLvOk() (*string, bool)`

GetCurAcctLvOk returns a tuple with the CurAcctLv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurAcctLv

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) SetCurAcctLv(v string)`

SetCurAcctLv sets CurAcctLv field to given value.

### HasCurAcctLv

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) HasCurAcctLv() bool`

HasCurAcctLv returns a boolean if a field has been set.

### GetMgnAft

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) GetMgnAft() GetAccountSetAccountSwitchPrecheckV5RespDataMgnAft`

GetMgnAft returns the MgnAft field if non-nil, zero value otherwise.

### GetMgnAftOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) GetMgnAftOk() (*GetAccountSetAccountSwitchPrecheckV5RespDataMgnAft, bool)`

GetMgnAftOk returns a tuple with the MgnAft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnAft

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) SetMgnAft(v GetAccountSetAccountSwitchPrecheckV5RespDataMgnAft)`

SetMgnAft sets MgnAft field to given value.

### HasMgnAft

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) HasMgnAft() bool`

HasMgnAft returns a boolean if a field has been set.

### GetMgnBf

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) GetMgnBf() GetAccountSetAccountSwitchPrecheckV5RespDataMgnBf`

GetMgnBf returns the MgnBf field if non-nil, zero value otherwise.

### GetMgnBfOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) GetMgnBfOk() (*GetAccountSetAccountSwitchPrecheckV5RespDataMgnBf, bool)`

GetMgnBfOk returns a tuple with the MgnBf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnBf

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) SetMgnBf(v GetAccountSetAccountSwitchPrecheckV5RespDataMgnBf)`

SetMgnBf sets MgnBf field to given value.

### HasMgnBf

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) HasMgnBf() bool`

HasMgnBf returns a boolean if a field has been set.

### GetPosList

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) GetPosList() []GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner`

GetPosList returns the PosList field if non-nil, zero value otherwise.

### GetPosListOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) GetPosListOk() (*[]GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner, bool)`

GetPosListOk returns a tuple with the PosList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosList

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) SetPosList(v []GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner)`

SetPosList sets PosList field to given value.

### HasPosList

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) HasPosList() bool`

HasPosList returns a boolean if a field has been set.

### GetPosTierCheck

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) GetPosTierCheck() []GetAccountSetAccountSwitchPrecheckV5RespDataPosTierCheckInner`

GetPosTierCheck returns the PosTierCheck field if non-nil, zero value otherwise.

### GetPosTierCheckOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) GetPosTierCheckOk() (*[]GetAccountSetAccountSwitchPrecheckV5RespDataPosTierCheckInner, bool)`

GetPosTierCheckOk returns a tuple with the PosTierCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosTierCheck

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) SetPosTierCheck(v []GetAccountSetAccountSwitchPrecheckV5RespDataPosTierCheckInner)`

SetPosTierCheck sets PosTierCheck field to given value.

### HasPosTierCheck

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) HasPosTierCheck() bool`

HasPosTierCheck returns a boolean if a field has been set.

### GetRiskOffsetType

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) GetRiskOffsetType() string`

GetRiskOffsetType returns the RiskOffsetType field if non-nil, zero value otherwise.

### GetRiskOffsetTypeOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) GetRiskOffsetTypeOk() (*string, bool)`

GetRiskOffsetTypeOk returns a tuple with the RiskOffsetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskOffsetType

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) SetRiskOffsetType(v string)`

SetRiskOffsetType sets RiskOffsetType field to given value.

### HasRiskOffsetType

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) HasRiskOffsetType() bool`

HasRiskOffsetType returns a boolean if a field has been set.

### GetSCode

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) GetSCode() string`

GetSCode returns the SCode field if non-nil, zero value otherwise.

### GetSCodeOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) GetSCodeOk() (*string, bool)`

GetSCodeOk returns a tuple with the SCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCode

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) SetSCode(v string)`

SetSCode sets SCode field to given value.

### HasSCode

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) HasSCode() bool`

HasSCode returns a boolean if a field has been set.

### GetUnmatchedInfoCheck

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) GetUnmatchedInfoCheck() []GetAccountSetAccountSwitchPrecheckV5RespDataUnmatchedInfoCheckInner`

GetUnmatchedInfoCheck returns the UnmatchedInfoCheck field if non-nil, zero value otherwise.

### GetUnmatchedInfoCheckOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) GetUnmatchedInfoCheckOk() (*[]GetAccountSetAccountSwitchPrecheckV5RespDataUnmatchedInfoCheckInner, bool)`

GetUnmatchedInfoCheckOk returns a tuple with the UnmatchedInfoCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedInfoCheck

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) SetUnmatchedInfoCheck(v []GetAccountSetAccountSwitchPrecheckV5RespDataUnmatchedInfoCheckInner)`

SetUnmatchedInfoCheck sets UnmatchedInfoCheck field to given value.

### HasUnmatchedInfoCheck

`func (o *GetAccountSetAccountSwitchPrecheckV5RespData) HasUnmatchedInfoCheck() bool`

HasUnmatchedInfoCheck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


