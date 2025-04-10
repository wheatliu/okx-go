# GetAssetConvertCurrencyPairV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCcy** | Pointer to **string** | Base currency, e.g. &#x60;BTC&#x60; in &#x60;BTC-USDT&#x60; | [optional] [default to ""]
**BaseCcyMax** | Pointer to **string** | Maximum amount of base currency | [optional] [default to ""]
**BaseCcyMin** | Pointer to **string** | Minimum amount of base currency | [optional] [default to ""]
**InstId** | Pointer to **string** | Currency pair, e.g. &#x60;BTC-USDT&#x60; | [optional] [default to ""]
**QuoteCcy** | Pointer to **string** | Quote currency, e.g. &#x60;USDT&#x60; in &#x60;BTC-USDT&#x60; | [optional] [default to ""]
**QuoteCcyMax** | Pointer to **string** | Maximum amount of quote currency | [optional] [default to ""]
**QuoteCcyMin** | Pointer to **string** | Minimum amount of quote currency | [optional] [default to ""]

## Methods

### NewGetAssetConvertCurrencyPairV5RespDataInner

`func NewGetAssetConvertCurrencyPairV5RespDataInner() *GetAssetConvertCurrencyPairV5RespDataInner`

NewGetAssetConvertCurrencyPairV5RespDataInner instantiates a new GetAssetConvertCurrencyPairV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetConvertCurrencyPairV5RespDataInnerWithDefaults

`func NewGetAssetConvertCurrencyPairV5RespDataInnerWithDefaults() *GetAssetConvertCurrencyPairV5RespDataInner`

NewGetAssetConvertCurrencyPairV5RespDataInnerWithDefaults instantiates a new GetAssetConvertCurrencyPairV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseCcy

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) GetBaseCcy() string`

GetBaseCcy returns the BaseCcy field if non-nil, zero value otherwise.

### GetBaseCcyOk

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) GetBaseCcyOk() (*string, bool)`

GetBaseCcyOk returns a tuple with the BaseCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCcy

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) SetBaseCcy(v string)`

SetBaseCcy sets BaseCcy field to given value.

### HasBaseCcy

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) HasBaseCcy() bool`

HasBaseCcy returns a boolean if a field has been set.

### GetBaseCcyMax

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) GetBaseCcyMax() string`

GetBaseCcyMax returns the BaseCcyMax field if non-nil, zero value otherwise.

### GetBaseCcyMaxOk

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) GetBaseCcyMaxOk() (*string, bool)`

GetBaseCcyMaxOk returns a tuple with the BaseCcyMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCcyMax

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) SetBaseCcyMax(v string)`

SetBaseCcyMax sets BaseCcyMax field to given value.

### HasBaseCcyMax

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) HasBaseCcyMax() bool`

HasBaseCcyMax returns a boolean if a field has been set.

### GetBaseCcyMin

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) GetBaseCcyMin() string`

GetBaseCcyMin returns the BaseCcyMin field if non-nil, zero value otherwise.

### GetBaseCcyMinOk

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) GetBaseCcyMinOk() (*string, bool)`

GetBaseCcyMinOk returns a tuple with the BaseCcyMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCcyMin

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) SetBaseCcyMin(v string)`

SetBaseCcyMin sets BaseCcyMin field to given value.

### HasBaseCcyMin

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) HasBaseCcyMin() bool`

HasBaseCcyMin returns a boolean if a field has been set.

### GetInstId

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetQuoteCcy

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) GetQuoteCcy() string`

GetQuoteCcy returns the QuoteCcy field if non-nil, zero value otherwise.

### GetQuoteCcyOk

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) GetQuoteCcyOk() (*string, bool)`

GetQuoteCcyOk returns a tuple with the QuoteCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCcy

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) SetQuoteCcy(v string)`

SetQuoteCcy sets QuoteCcy field to given value.

### HasQuoteCcy

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) HasQuoteCcy() bool`

HasQuoteCcy returns a boolean if a field has been set.

### GetQuoteCcyMax

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) GetQuoteCcyMax() string`

GetQuoteCcyMax returns the QuoteCcyMax field if non-nil, zero value otherwise.

### GetQuoteCcyMaxOk

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) GetQuoteCcyMaxOk() (*string, bool)`

GetQuoteCcyMaxOk returns a tuple with the QuoteCcyMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCcyMax

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) SetQuoteCcyMax(v string)`

SetQuoteCcyMax sets QuoteCcyMax field to given value.

### HasQuoteCcyMax

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) HasQuoteCcyMax() bool`

HasQuoteCcyMax returns a boolean if a field has been set.

### GetQuoteCcyMin

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) GetQuoteCcyMin() string`

GetQuoteCcyMin returns the QuoteCcyMin field if non-nil, zero value otherwise.

### GetQuoteCcyMinOk

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) GetQuoteCcyMinOk() (*string, bool)`

GetQuoteCcyMinOk returns a tuple with the QuoteCcyMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCcyMin

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) SetQuoteCcyMin(v string)`

SetQuoteCcyMin sets QuoteCcyMin field to given value.

### HasQuoteCcyMin

`func (o *GetAssetConvertCurrencyPairV5RespDataInner) HasQuoteCcyMin() bool`

HasQuoteCcyMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


