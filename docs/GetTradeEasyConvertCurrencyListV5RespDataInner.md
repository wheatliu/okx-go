# GetTradeEasyConvertCurrencyListV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromData** | Pointer to [**[]GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner**](GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner.md) | Currently owned and convertible small currency list | [optional] 
**ToCcy** | Pointer to **[]string** | Type of mainstream currency convert to, e.g. &#x60;USDT&#x60; | [optional] 

## Methods

### NewGetTradeEasyConvertCurrencyListV5RespDataInner

`func NewGetTradeEasyConvertCurrencyListV5RespDataInner() *GetTradeEasyConvertCurrencyListV5RespDataInner`

NewGetTradeEasyConvertCurrencyListV5RespDataInner instantiates a new GetTradeEasyConvertCurrencyListV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradeEasyConvertCurrencyListV5RespDataInnerWithDefaults

`func NewGetTradeEasyConvertCurrencyListV5RespDataInnerWithDefaults() *GetTradeEasyConvertCurrencyListV5RespDataInner`

NewGetTradeEasyConvertCurrencyListV5RespDataInnerWithDefaults instantiates a new GetTradeEasyConvertCurrencyListV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromData

`func (o *GetTradeEasyConvertCurrencyListV5RespDataInner) GetFromData() []GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner`

GetFromData returns the FromData field if non-nil, zero value otherwise.

### GetFromDataOk

`func (o *GetTradeEasyConvertCurrencyListV5RespDataInner) GetFromDataOk() (*[]GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner, bool)`

GetFromDataOk returns a tuple with the FromData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromData

`func (o *GetTradeEasyConvertCurrencyListV5RespDataInner) SetFromData(v []GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner)`

SetFromData sets FromData field to given value.

### HasFromData

`func (o *GetTradeEasyConvertCurrencyListV5RespDataInner) HasFromData() bool`

HasFromData returns a boolean if a field has been set.

### GetToCcy

`func (o *GetTradeEasyConvertCurrencyListV5RespDataInner) GetToCcy() []string`

GetToCcy returns the ToCcy field if non-nil, zero value otherwise.

### GetToCcyOk

`func (o *GetTradeEasyConvertCurrencyListV5RespDataInner) GetToCcyOk() (*[]string, bool)`

GetToCcyOk returns a tuple with the ToCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCcy

`func (o *GetTradeEasyConvertCurrencyListV5RespDataInner) SetToCcy(v []string)`

SetToCcy sets ToCcy field to given value.

### HasToCcy

`func (o *GetTradeEasyConvertCurrencyListV5RespDataInner) HasToCcy() bool`

HasToCcy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


