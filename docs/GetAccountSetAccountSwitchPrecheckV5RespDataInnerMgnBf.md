# GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctAvailEq** | Pointer to **string** | Account available equity in USD   Applicable when curAcctLv is &#x60;3/4&#x60;, return \&quot;\&quot; for other scenarios | [optional] [default to ""]
**AvailEq** | Pointer to **string** | Available equity of currency | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency | [optional] [default to ""]
**Details** | Pointer to [**[]GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnAftDetailsInner**](GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnAftDetailsInner.md) | Detailed information   Only applicable when curAcctLv is &#x60;2&#x60;, return \&quot;\&quot; for other scenarios | [optional] 
**MgnRatio** | Pointer to **string** | Margin ratio of currency | [optional] [default to ""]

## Methods

### NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf

`func NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf() *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf`

NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf instantiates a new GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBfWithDefaults

`func NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBfWithDefaults() *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf`

NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBfWithDefaults instantiates a new GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctAvailEq

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf) GetAcctAvailEq() string`

GetAcctAvailEq returns the AcctAvailEq field if non-nil, zero value otherwise.

### GetAcctAvailEqOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf) GetAcctAvailEqOk() (*string, bool)`

GetAcctAvailEqOk returns a tuple with the AcctAvailEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctAvailEq

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf) SetAcctAvailEq(v string)`

SetAcctAvailEq sets AcctAvailEq field to given value.

### HasAcctAvailEq

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf) HasAcctAvailEq() bool`

HasAcctAvailEq returns a boolean if a field has been set.

### GetAvailEq

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf) GetAvailEq() string`

GetAvailEq returns the AvailEq field if non-nil, zero value otherwise.

### GetAvailEqOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf) GetAvailEqOk() (*string, bool)`

GetAvailEqOk returns a tuple with the AvailEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailEq

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf) SetAvailEq(v string)`

SetAvailEq sets AvailEq field to given value.

### HasAvailEq

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf) HasAvailEq() bool`

HasAvailEq returns a boolean if a field has been set.

### GetCcy

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetDetails

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf) GetDetails() []GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnAftDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf) GetDetailsOk() (*[]GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnAftDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf) SetDetails(v []GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnAftDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetMgnRatio

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf) GetMgnRatio() string`

GetMgnRatio returns the MgnRatio field if non-nil, zero value otherwise.

### GetMgnRatioOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf) GetMgnRatioOk() (*string, bool)`

GetMgnRatioOk returns a tuple with the MgnRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnRatio

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf) SetMgnRatio(v string)`

SetMgnRatio sets MgnRatio field to given value.

### HasMgnRatio

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerMgnBf) HasMgnRatio() bool`

HasMgnRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


