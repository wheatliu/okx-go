# CreateTradeOneClickRepayV2V5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DebtCcy** | **string** | Debt currency | [default to ""]
**RepayCcyList** | **[]string** | Repay currency list, e.g. [\&quot;USDC\&quot;,\&quot;BTC\&quot;]  The priority of currency to repay is consistent with the order in the array. (The first item has the highest priority) | 

## Methods

### NewCreateTradeOneClickRepayV2V5Req

`func NewCreateTradeOneClickRepayV2V5Req(debtCcy string, repayCcyList []string, ) *CreateTradeOneClickRepayV2V5Req`

NewCreateTradeOneClickRepayV2V5Req instantiates a new CreateTradeOneClickRepayV2V5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeOneClickRepayV2V5ReqWithDefaults

`func NewCreateTradeOneClickRepayV2V5ReqWithDefaults() *CreateTradeOneClickRepayV2V5Req`

NewCreateTradeOneClickRepayV2V5ReqWithDefaults instantiates a new CreateTradeOneClickRepayV2V5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebtCcy

`func (o *CreateTradeOneClickRepayV2V5Req) GetDebtCcy() string`

GetDebtCcy returns the DebtCcy field if non-nil, zero value otherwise.

### GetDebtCcyOk

`func (o *CreateTradeOneClickRepayV2V5Req) GetDebtCcyOk() (*string, bool)`

GetDebtCcyOk returns a tuple with the DebtCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebtCcy

`func (o *CreateTradeOneClickRepayV2V5Req) SetDebtCcy(v string)`

SetDebtCcy sets DebtCcy field to given value.


### GetRepayCcyList

`func (o *CreateTradeOneClickRepayV2V5Req) GetRepayCcyList() []string`

GetRepayCcyList returns the RepayCcyList field if non-nil, zero value otherwise.

### GetRepayCcyListOk

`func (o *CreateTradeOneClickRepayV2V5Req) GetRepayCcyListOk() (*[]string, bool)`

GetRepayCcyListOk returns a tuple with the RepayCcyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepayCcyList

`func (o *CreateTradeOneClickRepayV2V5Req) SetRepayCcyList(v []string)`

SetRepayCcyList sets RepayCcyList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


