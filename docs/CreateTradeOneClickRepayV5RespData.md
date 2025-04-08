# CreateTradeOneClickRepayV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DebtCcy** | Pointer to **string** | Debt currency type | [optional] [default to ""]
**FillDebtSz** | Pointer to **string** | Filled amount of debt currency | [optional] [default to ""]
**FillRepaySz** | Pointer to **string** | Filled amount of repay currency | [optional] [default to ""]
**RepayCcy** | Pointer to **string** | Repay currency type | [optional] [default to ""]
**Status** | Pointer to **string** | Current status of one-click repay   &#x60;running&#x60;: Running   &#x60;filled&#x60;: Filled   &#x60;failed&#x60;: Failed | [optional] [default to ""]
**UTime** | Pointer to **string** | Trade time, Unix timestamp format in milliseconds, e.g. 1597026383085 | [optional] [default to ""]

## Methods

### NewCreateTradeOneClickRepayV5RespData

`func NewCreateTradeOneClickRepayV5RespData() *CreateTradeOneClickRepayV5RespData`

NewCreateTradeOneClickRepayV5RespData instantiates a new CreateTradeOneClickRepayV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeOneClickRepayV5RespDataWithDefaults

`func NewCreateTradeOneClickRepayV5RespDataWithDefaults() *CreateTradeOneClickRepayV5RespData`

NewCreateTradeOneClickRepayV5RespDataWithDefaults instantiates a new CreateTradeOneClickRepayV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebtCcy

`func (o *CreateTradeOneClickRepayV5RespData) GetDebtCcy() string`

GetDebtCcy returns the DebtCcy field if non-nil, zero value otherwise.

### GetDebtCcyOk

`func (o *CreateTradeOneClickRepayV5RespData) GetDebtCcyOk() (*string, bool)`

GetDebtCcyOk returns a tuple with the DebtCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebtCcy

`func (o *CreateTradeOneClickRepayV5RespData) SetDebtCcy(v string)`

SetDebtCcy sets DebtCcy field to given value.

### HasDebtCcy

`func (o *CreateTradeOneClickRepayV5RespData) HasDebtCcy() bool`

HasDebtCcy returns a boolean if a field has been set.

### GetFillDebtSz

`func (o *CreateTradeOneClickRepayV5RespData) GetFillDebtSz() string`

GetFillDebtSz returns the FillDebtSz field if non-nil, zero value otherwise.

### GetFillDebtSzOk

`func (o *CreateTradeOneClickRepayV5RespData) GetFillDebtSzOk() (*string, bool)`

GetFillDebtSzOk returns a tuple with the FillDebtSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillDebtSz

`func (o *CreateTradeOneClickRepayV5RespData) SetFillDebtSz(v string)`

SetFillDebtSz sets FillDebtSz field to given value.

### HasFillDebtSz

`func (o *CreateTradeOneClickRepayV5RespData) HasFillDebtSz() bool`

HasFillDebtSz returns a boolean if a field has been set.

### GetFillRepaySz

`func (o *CreateTradeOneClickRepayV5RespData) GetFillRepaySz() string`

GetFillRepaySz returns the FillRepaySz field if non-nil, zero value otherwise.

### GetFillRepaySzOk

`func (o *CreateTradeOneClickRepayV5RespData) GetFillRepaySzOk() (*string, bool)`

GetFillRepaySzOk returns a tuple with the FillRepaySz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillRepaySz

`func (o *CreateTradeOneClickRepayV5RespData) SetFillRepaySz(v string)`

SetFillRepaySz sets FillRepaySz field to given value.

### HasFillRepaySz

`func (o *CreateTradeOneClickRepayV5RespData) HasFillRepaySz() bool`

HasFillRepaySz returns a boolean if a field has been set.

### GetRepayCcy

`func (o *CreateTradeOneClickRepayV5RespData) GetRepayCcy() string`

GetRepayCcy returns the RepayCcy field if non-nil, zero value otherwise.

### GetRepayCcyOk

`func (o *CreateTradeOneClickRepayV5RespData) GetRepayCcyOk() (*string, bool)`

GetRepayCcyOk returns a tuple with the RepayCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepayCcy

`func (o *CreateTradeOneClickRepayV5RespData) SetRepayCcy(v string)`

SetRepayCcy sets RepayCcy field to given value.

### HasRepayCcy

`func (o *CreateTradeOneClickRepayV5RespData) HasRepayCcy() bool`

HasRepayCcy returns a boolean if a field has been set.

### GetStatus

`func (o *CreateTradeOneClickRepayV5RespData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateTradeOneClickRepayV5RespData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateTradeOneClickRepayV5RespData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateTradeOneClickRepayV5RespData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUTime

`func (o *CreateTradeOneClickRepayV5RespData) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *CreateTradeOneClickRepayV5RespData) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *CreateTradeOneClickRepayV5RespData) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *CreateTradeOneClickRepayV5RespData) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


