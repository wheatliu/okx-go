# CreateTradeOneClickRepayV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DebtCcy** | **[]string** | Debt currency type   Maximum 5 currencies can be selected in one order. If there are multiple currencies, separate them with commas. | 
**RepayCcy** | **string** | Repay currency type   Only one receiving currency type can be selected in one order and cannot be the same as the small payment currencies. | [default to ""]

## Methods

### NewCreateTradeOneClickRepayV5Req

`func NewCreateTradeOneClickRepayV5Req(debtCcy []string, repayCcy string, ) *CreateTradeOneClickRepayV5Req`

NewCreateTradeOneClickRepayV5Req instantiates a new CreateTradeOneClickRepayV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeOneClickRepayV5ReqWithDefaults

`func NewCreateTradeOneClickRepayV5ReqWithDefaults() *CreateTradeOneClickRepayV5Req`

NewCreateTradeOneClickRepayV5ReqWithDefaults instantiates a new CreateTradeOneClickRepayV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebtCcy

`func (o *CreateTradeOneClickRepayV5Req) GetDebtCcy() []string`

GetDebtCcy returns the DebtCcy field if non-nil, zero value otherwise.

### GetDebtCcyOk

`func (o *CreateTradeOneClickRepayV5Req) GetDebtCcyOk() (*[]string, bool)`

GetDebtCcyOk returns a tuple with the DebtCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebtCcy

`func (o *CreateTradeOneClickRepayV5Req) SetDebtCcy(v []string)`

SetDebtCcy sets DebtCcy field to given value.


### GetRepayCcy

`func (o *CreateTradeOneClickRepayV5Req) GetRepayCcy() string`

GetRepayCcy returns the RepayCcy field if non-nil, zero value otherwise.

### GetRepayCcyOk

`func (o *CreateTradeOneClickRepayV5Req) GetRepayCcyOk() (*string, bool)`

GetRepayCcyOk returns a tuple with the RepayCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepayCcy

`func (o *CreateTradeOneClickRepayV5Req) SetRepayCcy(v string)`

SetRepayCcy sets RepayCcy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


