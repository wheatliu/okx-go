# GetAssetConvertCurrencyPairV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromCcy** | Pointer to **string** | Currency to convert from, e.g. &#x60;USDT&#x60; | [optional] [default to ""]
**ToCcy** | Pointer to **string** | Currency to convert to, e.g. &#x60;BTC&#x60; | [optional] [default to ""]

## Methods

### NewGetAssetConvertCurrencyPairV5RespData

`func NewGetAssetConvertCurrencyPairV5RespData() *GetAssetConvertCurrencyPairV5RespData`

NewGetAssetConvertCurrencyPairV5RespData instantiates a new GetAssetConvertCurrencyPairV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetConvertCurrencyPairV5RespDataWithDefaults

`func NewGetAssetConvertCurrencyPairV5RespDataWithDefaults() *GetAssetConvertCurrencyPairV5RespData`

NewGetAssetConvertCurrencyPairV5RespDataWithDefaults instantiates a new GetAssetConvertCurrencyPairV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromCcy

`func (o *GetAssetConvertCurrencyPairV5RespData) GetFromCcy() string`

GetFromCcy returns the FromCcy field if non-nil, zero value otherwise.

### GetFromCcyOk

`func (o *GetAssetConvertCurrencyPairV5RespData) GetFromCcyOk() (*string, bool)`

GetFromCcyOk returns a tuple with the FromCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCcy

`func (o *GetAssetConvertCurrencyPairV5RespData) SetFromCcy(v string)`

SetFromCcy sets FromCcy field to given value.

### HasFromCcy

`func (o *GetAssetConvertCurrencyPairV5RespData) HasFromCcy() bool`

HasFromCcy returns a boolean if a field has been set.

### GetToCcy

`func (o *GetAssetConvertCurrencyPairV5RespData) GetToCcy() string`

GetToCcy returns the ToCcy field if non-nil, zero value otherwise.

### GetToCcyOk

`func (o *GetAssetConvertCurrencyPairV5RespData) GetToCcyOk() (*string, bool)`

GetToCcyOk returns a tuple with the ToCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCcy

`func (o *GetAssetConvertCurrencyPairV5RespData) SetToCcy(v string)`

SetToCcy sets ToCcy field to given value.

### HasToCcy

`func (o *GetAssetConvertCurrencyPairV5RespData) HasToCcy() bool`

HasToCcy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


