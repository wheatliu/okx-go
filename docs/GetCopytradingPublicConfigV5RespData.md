# GetCopytradingPublicConfigV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxCopyAmt** | Pointer to **string** | Maximum copy amount per order in USDT when you are using copy mode &#x60;fixed_amount&#x60; | [optional] [default to ""]
**MaxCopyRatio** | Pointer to **string** | Maximum ratio per order when you are using copy mode &#x60;ratio_copy&#x60; | [optional] [default to ""]
**MaxCopyTotalAmt** | Pointer to **string** | Maximum copy total amount under the certain lead trader, the minimum is the same with &#x60;minCopyAmt&#x60; | [optional] [default to ""]
**MaxSlRatio** | Pointer to **string** | Maximum ratio of stopping loss per order, the minimum is 0 | [optional] [default to ""]
**MaxTpRatio** | Pointer to **string** | Maximum ratio of taking profit per order, the minimum is 0 | [optional] [default to ""]
**MinCopyAmt** | Pointer to **string** | Minimum copy amount per order in USDT when you are using copy mode &#x60;fixed_amount&#x60; | [optional] [default to ""]
**MinCopyRatio** | Pointer to **string** | Minimum ratio per order when you are using copy mode &#x60;ratio_copy&#x60; | [optional] [default to ""]

## Methods

### NewGetCopytradingPublicConfigV5RespData

`func NewGetCopytradingPublicConfigV5RespData() *GetCopytradingPublicConfigV5RespData`

NewGetCopytradingPublicConfigV5RespData instantiates a new GetCopytradingPublicConfigV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCopytradingPublicConfigV5RespDataWithDefaults

`func NewGetCopytradingPublicConfigV5RespDataWithDefaults() *GetCopytradingPublicConfigV5RespData`

NewGetCopytradingPublicConfigV5RespDataWithDefaults instantiates a new GetCopytradingPublicConfigV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxCopyAmt

`func (o *GetCopytradingPublicConfigV5RespData) GetMaxCopyAmt() string`

GetMaxCopyAmt returns the MaxCopyAmt field if non-nil, zero value otherwise.

### GetMaxCopyAmtOk

`func (o *GetCopytradingPublicConfigV5RespData) GetMaxCopyAmtOk() (*string, bool)`

GetMaxCopyAmtOk returns a tuple with the MaxCopyAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCopyAmt

`func (o *GetCopytradingPublicConfigV5RespData) SetMaxCopyAmt(v string)`

SetMaxCopyAmt sets MaxCopyAmt field to given value.

### HasMaxCopyAmt

`func (o *GetCopytradingPublicConfigV5RespData) HasMaxCopyAmt() bool`

HasMaxCopyAmt returns a boolean if a field has been set.

### GetMaxCopyRatio

`func (o *GetCopytradingPublicConfigV5RespData) GetMaxCopyRatio() string`

GetMaxCopyRatio returns the MaxCopyRatio field if non-nil, zero value otherwise.

### GetMaxCopyRatioOk

`func (o *GetCopytradingPublicConfigV5RespData) GetMaxCopyRatioOk() (*string, bool)`

GetMaxCopyRatioOk returns a tuple with the MaxCopyRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCopyRatio

`func (o *GetCopytradingPublicConfigV5RespData) SetMaxCopyRatio(v string)`

SetMaxCopyRatio sets MaxCopyRatio field to given value.

### HasMaxCopyRatio

`func (o *GetCopytradingPublicConfigV5RespData) HasMaxCopyRatio() bool`

HasMaxCopyRatio returns a boolean if a field has been set.

### GetMaxCopyTotalAmt

`func (o *GetCopytradingPublicConfigV5RespData) GetMaxCopyTotalAmt() string`

GetMaxCopyTotalAmt returns the MaxCopyTotalAmt field if non-nil, zero value otherwise.

### GetMaxCopyTotalAmtOk

`func (o *GetCopytradingPublicConfigV5RespData) GetMaxCopyTotalAmtOk() (*string, bool)`

GetMaxCopyTotalAmtOk returns a tuple with the MaxCopyTotalAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCopyTotalAmt

`func (o *GetCopytradingPublicConfigV5RespData) SetMaxCopyTotalAmt(v string)`

SetMaxCopyTotalAmt sets MaxCopyTotalAmt field to given value.

### HasMaxCopyTotalAmt

`func (o *GetCopytradingPublicConfigV5RespData) HasMaxCopyTotalAmt() bool`

HasMaxCopyTotalAmt returns a boolean if a field has been set.

### GetMaxSlRatio

`func (o *GetCopytradingPublicConfigV5RespData) GetMaxSlRatio() string`

GetMaxSlRatio returns the MaxSlRatio field if non-nil, zero value otherwise.

### GetMaxSlRatioOk

`func (o *GetCopytradingPublicConfigV5RespData) GetMaxSlRatioOk() (*string, bool)`

GetMaxSlRatioOk returns a tuple with the MaxSlRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSlRatio

`func (o *GetCopytradingPublicConfigV5RespData) SetMaxSlRatio(v string)`

SetMaxSlRatio sets MaxSlRatio field to given value.

### HasMaxSlRatio

`func (o *GetCopytradingPublicConfigV5RespData) HasMaxSlRatio() bool`

HasMaxSlRatio returns a boolean if a field has been set.

### GetMaxTpRatio

`func (o *GetCopytradingPublicConfigV5RespData) GetMaxTpRatio() string`

GetMaxTpRatio returns the MaxTpRatio field if non-nil, zero value otherwise.

### GetMaxTpRatioOk

`func (o *GetCopytradingPublicConfigV5RespData) GetMaxTpRatioOk() (*string, bool)`

GetMaxTpRatioOk returns a tuple with the MaxTpRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTpRatio

`func (o *GetCopytradingPublicConfigV5RespData) SetMaxTpRatio(v string)`

SetMaxTpRatio sets MaxTpRatio field to given value.

### HasMaxTpRatio

`func (o *GetCopytradingPublicConfigV5RespData) HasMaxTpRatio() bool`

HasMaxTpRatio returns a boolean if a field has been set.

### GetMinCopyAmt

`func (o *GetCopytradingPublicConfigV5RespData) GetMinCopyAmt() string`

GetMinCopyAmt returns the MinCopyAmt field if non-nil, zero value otherwise.

### GetMinCopyAmtOk

`func (o *GetCopytradingPublicConfigV5RespData) GetMinCopyAmtOk() (*string, bool)`

GetMinCopyAmtOk returns a tuple with the MinCopyAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCopyAmt

`func (o *GetCopytradingPublicConfigV5RespData) SetMinCopyAmt(v string)`

SetMinCopyAmt sets MinCopyAmt field to given value.

### HasMinCopyAmt

`func (o *GetCopytradingPublicConfigV5RespData) HasMinCopyAmt() bool`

HasMinCopyAmt returns a boolean if a field has been set.

### GetMinCopyRatio

`func (o *GetCopytradingPublicConfigV5RespData) GetMinCopyRatio() string`

GetMinCopyRatio returns the MinCopyRatio field if non-nil, zero value otherwise.

### GetMinCopyRatioOk

`func (o *GetCopytradingPublicConfigV5RespData) GetMinCopyRatioOk() (*string, bool)`

GetMinCopyRatioOk returns a tuple with the MinCopyRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCopyRatio

`func (o *GetCopytradingPublicConfigV5RespData) SetMinCopyRatio(v string)`

SetMinCopyRatio sets MinCopyRatio field to given value.

### HasMinCopyRatio

`func (o *GetCopytradingPublicConfigV5RespData) HasMinCopyRatio() bool`

HasMinCopyRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


