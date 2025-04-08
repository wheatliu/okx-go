# GetTradeOneClickRepayCurrencyListV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DebtData** | Pointer to [**[]GetTradeOneClickRepayCurrencyListV5RespDataInnerDebtDataInner**](GetTradeOneClickRepayCurrencyListV5RespDataInnerDebtDataInner.md) | Debt currency data list | [optional] 
**DebtType** | Pointer to **string** | Debt type   &#x60;cross&#x60;: cross   &#x60;isolated&#x60;: isolated | [optional] [default to ""]
**RepayData** | Pointer to [**[]GetTradeOneClickRepayCurrencyListV2V5RespDataInnerRepayDataInner**](GetTradeOneClickRepayCurrencyListV2V5RespDataInnerRepayDataInner.md) | Repay currency data list | [optional] 

## Methods

### NewGetTradeOneClickRepayCurrencyListV5RespDataInner

`func NewGetTradeOneClickRepayCurrencyListV5RespDataInner() *GetTradeOneClickRepayCurrencyListV5RespDataInner`

NewGetTradeOneClickRepayCurrencyListV5RespDataInner instantiates a new GetTradeOneClickRepayCurrencyListV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradeOneClickRepayCurrencyListV5RespDataInnerWithDefaults

`func NewGetTradeOneClickRepayCurrencyListV5RespDataInnerWithDefaults() *GetTradeOneClickRepayCurrencyListV5RespDataInner`

NewGetTradeOneClickRepayCurrencyListV5RespDataInnerWithDefaults instantiates a new GetTradeOneClickRepayCurrencyListV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebtData

`func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) GetDebtData() []GetTradeOneClickRepayCurrencyListV5RespDataInnerDebtDataInner`

GetDebtData returns the DebtData field if non-nil, zero value otherwise.

### GetDebtDataOk

`func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) GetDebtDataOk() (*[]GetTradeOneClickRepayCurrencyListV5RespDataInnerDebtDataInner, bool)`

GetDebtDataOk returns a tuple with the DebtData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebtData

`func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) SetDebtData(v []GetTradeOneClickRepayCurrencyListV5RespDataInnerDebtDataInner)`

SetDebtData sets DebtData field to given value.

### HasDebtData

`func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) HasDebtData() bool`

HasDebtData returns a boolean if a field has been set.

### GetDebtType

`func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) GetDebtType() string`

GetDebtType returns the DebtType field if non-nil, zero value otherwise.

### GetDebtTypeOk

`func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) GetDebtTypeOk() (*string, bool)`

GetDebtTypeOk returns a tuple with the DebtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebtType

`func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) SetDebtType(v string)`

SetDebtType sets DebtType field to given value.

### HasDebtType

`func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) HasDebtType() bool`

HasDebtType returns a boolean if a field has been set.

### GetRepayData

`func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) GetRepayData() []GetTradeOneClickRepayCurrencyListV2V5RespDataInnerRepayDataInner`

GetRepayData returns the RepayData field if non-nil, zero value otherwise.

### GetRepayDataOk

`func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) GetRepayDataOk() (*[]GetTradeOneClickRepayCurrencyListV2V5RespDataInnerRepayDataInner, bool)`

GetRepayDataOk returns a tuple with the RepayData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepayData

`func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) SetRepayData(v []GetTradeOneClickRepayCurrencyListV2V5RespDataInnerRepayDataInner)`

SetRepayData sets RepayData field to given value.

### HasRepayData

`func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) HasRepayData() bool`

HasRepayData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


