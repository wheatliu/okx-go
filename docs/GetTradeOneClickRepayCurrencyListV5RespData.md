# GetTradeOneClickRepayCurrencyListV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DebtData** | Pointer to [**[]GetTradeOneClickRepayCurrencyListV5RespDataDebtDataInner**](GetTradeOneClickRepayCurrencyListV5RespDataDebtDataInner.md) | Debt currency data list | [optional] 
**DebtType** | Pointer to **string** | Debt type   &#x60;cross&#x60;: cross   &#x60;isolated&#x60;: isolated | [optional] [default to ""]
**RepayData** | Pointer to [**[]GetTradeOneClickRepayCurrencyListV2V5RespDataRepayDataInner**](GetTradeOneClickRepayCurrencyListV2V5RespDataRepayDataInner.md) | Repay currency data list | [optional] 

## Methods

### NewGetTradeOneClickRepayCurrencyListV5RespData

`func NewGetTradeOneClickRepayCurrencyListV5RespData() *GetTradeOneClickRepayCurrencyListV5RespData`

NewGetTradeOneClickRepayCurrencyListV5RespData instantiates a new GetTradeOneClickRepayCurrencyListV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradeOneClickRepayCurrencyListV5RespDataWithDefaults

`func NewGetTradeOneClickRepayCurrencyListV5RespDataWithDefaults() *GetTradeOneClickRepayCurrencyListV5RespData`

NewGetTradeOneClickRepayCurrencyListV5RespDataWithDefaults instantiates a new GetTradeOneClickRepayCurrencyListV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebtData

`func (o *GetTradeOneClickRepayCurrencyListV5RespData) GetDebtData() []GetTradeOneClickRepayCurrencyListV5RespDataDebtDataInner`

GetDebtData returns the DebtData field if non-nil, zero value otherwise.

### GetDebtDataOk

`func (o *GetTradeOneClickRepayCurrencyListV5RespData) GetDebtDataOk() (*[]GetTradeOneClickRepayCurrencyListV5RespDataDebtDataInner, bool)`

GetDebtDataOk returns a tuple with the DebtData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebtData

`func (o *GetTradeOneClickRepayCurrencyListV5RespData) SetDebtData(v []GetTradeOneClickRepayCurrencyListV5RespDataDebtDataInner)`

SetDebtData sets DebtData field to given value.

### HasDebtData

`func (o *GetTradeOneClickRepayCurrencyListV5RespData) HasDebtData() bool`

HasDebtData returns a boolean if a field has been set.

### GetDebtType

`func (o *GetTradeOneClickRepayCurrencyListV5RespData) GetDebtType() string`

GetDebtType returns the DebtType field if non-nil, zero value otherwise.

### GetDebtTypeOk

`func (o *GetTradeOneClickRepayCurrencyListV5RespData) GetDebtTypeOk() (*string, bool)`

GetDebtTypeOk returns a tuple with the DebtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebtType

`func (o *GetTradeOneClickRepayCurrencyListV5RespData) SetDebtType(v string)`

SetDebtType sets DebtType field to given value.

### HasDebtType

`func (o *GetTradeOneClickRepayCurrencyListV5RespData) HasDebtType() bool`

HasDebtType returns a boolean if a field has been set.

### GetRepayData

`func (o *GetTradeOneClickRepayCurrencyListV5RespData) GetRepayData() []GetTradeOneClickRepayCurrencyListV2V5RespDataRepayDataInner`

GetRepayData returns the RepayData field if non-nil, zero value otherwise.

### GetRepayDataOk

`func (o *GetTradeOneClickRepayCurrencyListV5RespData) GetRepayDataOk() (*[]GetTradeOneClickRepayCurrencyListV2V5RespDataRepayDataInner, bool)`

GetRepayDataOk returns a tuple with the RepayData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepayData

`func (o *GetTradeOneClickRepayCurrencyListV5RespData) SetRepayData(v []GetTradeOneClickRepayCurrencyListV2V5RespDataRepayDataInner)`

SetRepayData sets RepayData field to given value.

### HasRepayData

`func (o *GetTradeOneClickRepayCurrencyListV5RespData) HasRepayData() bool`

HasRepayData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


