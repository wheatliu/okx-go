# CreateTradeEasyConvertV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FillFromSz** | Pointer to **string** | Filled amount of small payment currency convert from | [optional] [default to ""]
**FillToSz** | Pointer to **string** | Filled amount of mainstream currency convert to | [optional] [default to ""]
**FromCcy** | Pointer to **string** | Type of small payment currency convert from | [optional] [default to ""]
**Status** | Pointer to **string** | Current status of easy convert   &#x60;running&#x60;: Running   &#x60;filled&#x60;: Filled   &#x60;failed&#x60;: Failed | [optional] [default to ""]
**ToCcy** | Pointer to **string** | Type of mainstream currency convert to | [optional] [default to ""]
**UTime** | Pointer to **string** | Trade time, Unix timestamp format in milliseconds, e.g. 1597026383085 | [optional] [default to ""]

## Methods

### NewCreateTradeEasyConvertV5RespDataInner

`func NewCreateTradeEasyConvertV5RespDataInner() *CreateTradeEasyConvertV5RespDataInner`

NewCreateTradeEasyConvertV5RespDataInner instantiates a new CreateTradeEasyConvertV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeEasyConvertV5RespDataInnerWithDefaults

`func NewCreateTradeEasyConvertV5RespDataInnerWithDefaults() *CreateTradeEasyConvertV5RespDataInner`

NewCreateTradeEasyConvertV5RespDataInnerWithDefaults instantiates a new CreateTradeEasyConvertV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFillFromSz

`func (o *CreateTradeEasyConvertV5RespDataInner) GetFillFromSz() string`

GetFillFromSz returns the FillFromSz field if non-nil, zero value otherwise.

### GetFillFromSzOk

`func (o *CreateTradeEasyConvertV5RespDataInner) GetFillFromSzOk() (*string, bool)`

GetFillFromSzOk returns a tuple with the FillFromSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillFromSz

`func (o *CreateTradeEasyConvertV5RespDataInner) SetFillFromSz(v string)`

SetFillFromSz sets FillFromSz field to given value.

### HasFillFromSz

`func (o *CreateTradeEasyConvertV5RespDataInner) HasFillFromSz() bool`

HasFillFromSz returns a boolean if a field has been set.

### GetFillToSz

`func (o *CreateTradeEasyConvertV5RespDataInner) GetFillToSz() string`

GetFillToSz returns the FillToSz field if non-nil, zero value otherwise.

### GetFillToSzOk

`func (o *CreateTradeEasyConvertV5RespDataInner) GetFillToSzOk() (*string, bool)`

GetFillToSzOk returns a tuple with the FillToSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillToSz

`func (o *CreateTradeEasyConvertV5RespDataInner) SetFillToSz(v string)`

SetFillToSz sets FillToSz field to given value.

### HasFillToSz

`func (o *CreateTradeEasyConvertV5RespDataInner) HasFillToSz() bool`

HasFillToSz returns a boolean if a field has been set.

### GetFromCcy

`func (o *CreateTradeEasyConvertV5RespDataInner) GetFromCcy() string`

GetFromCcy returns the FromCcy field if non-nil, zero value otherwise.

### GetFromCcyOk

`func (o *CreateTradeEasyConvertV5RespDataInner) GetFromCcyOk() (*string, bool)`

GetFromCcyOk returns a tuple with the FromCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCcy

`func (o *CreateTradeEasyConvertV5RespDataInner) SetFromCcy(v string)`

SetFromCcy sets FromCcy field to given value.

### HasFromCcy

`func (o *CreateTradeEasyConvertV5RespDataInner) HasFromCcy() bool`

HasFromCcy returns a boolean if a field has been set.

### GetStatus

`func (o *CreateTradeEasyConvertV5RespDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateTradeEasyConvertV5RespDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateTradeEasyConvertV5RespDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateTradeEasyConvertV5RespDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToCcy

`func (o *CreateTradeEasyConvertV5RespDataInner) GetToCcy() string`

GetToCcy returns the ToCcy field if non-nil, zero value otherwise.

### GetToCcyOk

`func (o *CreateTradeEasyConvertV5RespDataInner) GetToCcyOk() (*string, bool)`

GetToCcyOk returns a tuple with the ToCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCcy

`func (o *CreateTradeEasyConvertV5RespDataInner) SetToCcy(v string)`

SetToCcy sets ToCcy field to given value.

### HasToCcy

`func (o *CreateTradeEasyConvertV5RespDataInner) HasToCcy() bool`

HasToCcy returns a boolean if a field has been set.

### GetUTime

`func (o *CreateTradeEasyConvertV5RespDataInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *CreateTradeEasyConvertV5RespDataInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *CreateTradeEasyConvertV5RespDataInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *CreateTradeEasyConvertV5RespDataInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


