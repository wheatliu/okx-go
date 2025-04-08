# CreateTradeOneClickRepayV2V5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DebtCcy** | Pointer to **string** | Debt currency | [optional] [default to ""]
**RepayCcyList** | Pointer to **[]string** | Repay currency list, e.g. [\&quot;USDC\&quot;,\&quot;BTC\&quot;] | [optional] 
**Ts** | Pointer to **string** | Request time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewCreateTradeOneClickRepayV2V5RespDataInner

`func NewCreateTradeOneClickRepayV2V5RespDataInner() *CreateTradeOneClickRepayV2V5RespDataInner`

NewCreateTradeOneClickRepayV2V5RespDataInner instantiates a new CreateTradeOneClickRepayV2V5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeOneClickRepayV2V5RespDataInnerWithDefaults

`func NewCreateTradeOneClickRepayV2V5RespDataInnerWithDefaults() *CreateTradeOneClickRepayV2V5RespDataInner`

NewCreateTradeOneClickRepayV2V5RespDataInnerWithDefaults instantiates a new CreateTradeOneClickRepayV2V5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebtCcy

`func (o *CreateTradeOneClickRepayV2V5RespDataInner) GetDebtCcy() string`

GetDebtCcy returns the DebtCcy field if non-nil, zero value otherwise.

### GetDebtCcyOk

`func (o *CreateTradeOneClickRepayV2V5RespDataInner) GetDebtCcyOk() (*string, bool)`

GetDebtCcyOk returns a tuple with the DebtCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebtCcy

`func (o *CreateTradeOneClickRepayV2V5RespDataInner) SetDebtCcy(v string)`

SetDebtCcy sets DebtCcy field to given value.

### HasDebtCcy

`func (o *CreateTradeOneClickRepayV2V5RespDataInner) HasDebtCcy() bool`

HasDebtCcy returns a boolean if a field has been set.

### GetRepayCcyList

`func (o *CreateTradeOneClickRepayV2V5RespDataInner) GetRepayCcyList() []string`

GetRepayCcyList returns the RepayCcyList field if non-nil, zero value otherwise.

### GetRepayCcyListOk

`func (o *CreateTradeOneClickRepayV2V5RespDataInner) GetRepayCcyListOk() (*[]string, bool)`

GetRepayCcyListOk returns a tuple with the RepayCcyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepayCcyList

`func (o *CreateTradeOneClickRepayV2V5RespDataInner) SetRepayCcyList(v []string)`

SetRepayCcyList sets RepayCcyList field to given value.

### HasRepayCcyList

`func (o *CreateTradeOneClickRepayV2V5RespDataInner) HasRepayCcyList() bool`

HasRepayCcyList returns a boolean if a field has been set.

### GetTs

`func (o *CreateTradeOneClickRepayV2V5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *CreateTradeOneClickRepayV2V5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *CreateTradeOneClickRepayV2V5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *CreateTradeOneClickRepayV2V5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


