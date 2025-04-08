# GetAssetConvertCurrenciesV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | Pointer to **string** | Currency, e.g. BTC | [optional] [default to ""]
**Max** | Pointer to **string** | Maximum amount to convert ( Deprecated ) | [optional] [default to ""]
**Min** | Pointer to **string** | Minimum amount to convert ( Deprecated ) | [optional] [default to ""]

## Methods

### NewGetAssetConvertCurrenciesV5RespData

`func NewGetAssetConvertCurrenciesV5RespData() *GetAssetConvertCurrenciesV5RespData`

NewGetAssetConvertCurrenciesV5RespData instantiates a new GetAssetConvertCurrenciesV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetConvertCurrenciesV5RespDataWithDefaults

`func NewGetAssetConvertCurrenciesV5RespDataWithDefaults() *GetAssetConvertCurrenciesV5RespData`

NewGetAssetConvertCurrenciesV5RespDataWithDefaults instantiates a new GetAssetConvertCurrenciesV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *GetAssetConvertCurrenciesV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAssetConvertCurrenciesV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAssetConvertCurrenciesV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAssetConvertCurrenciesV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetMax

`func (o *GetAssetConvertCurrenciesV5RespData) GetMax() string`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GetAssetConvertCurrenciesV5RespData) GetMaxOk() (*string, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GetAssetConvertCurrenciesV5RespData) SetMax(v string)`

SetMax sets Max field to given value.

### HasMax

`func (o *GetAssetConvertCurrenciesV5RespData) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *GetAssetConvertCurrenciesV5RespData) GetMin() string`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *GetAssetConvertCurrenciesV5RespData) GetMinOk() (*string, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *GetAssetConvertCurrenciesV5RespData) SetMin(v string)`

SetMin sets Min field to given value.

### HasMin

`func (o *GetAssetConvertCurrenciesV5RespData) HasMin() bool`

HasMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


