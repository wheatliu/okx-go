# CreateTradeEasyConvertV5RespData

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

### NewCreateTradeEasyConvertV5RespData

`func NewCreateTradeEasyConvertV5RespData() *CreateTradeEasyConvertV5RespData`

NewCreateTradeEasyConvertV5RespData instantiates a new CreateTradeEasyConvertV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeEasyConvertV5RespDataWithDefaults

`func NewCreateTradeEasyConvertV5RespDataWithDefaults() *CreateTradeEasyConvertV5RespData`

NewCreateTradeEasyConvertV5RespDataWithDefaults instantiates a new CreateTradeEasyConvertV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFillFromSz

`func (o *CreateTradeEasyConvertV5RespData) GetFillFromSz() string`

GetFillFromSz returns the FillFromSz field if non-nil, zero value otherwise.

### GetFillFromSzOk

`func (o *CreateTradeEasyConvertV5RespData) GetFillFromSzOk() (*string, bool)`

GetFillFromSzOk returns a tuple with the FillFromSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillFromSz

`func (o *CreateTradeEasyConvertV5RespData) SetFillFromSz(v string)`

SetFillFromSz sets FillFromSz field to given value.

### HasFillFromSz

`func (o *CreateTradeEasyConvertV5RespData) HasFillFromSz() bool`

HasFillFromSz returns a boolean if a field has been set.

### GetFillToSz

`func (o *CreateTradeEasyConvertV5RespData) GetFillToSz() string`

GetFillToSz returns the FillToSz field if non-nil, zero value otherwise.

### GetFillToSzOk

`func (o *CreateTradeEasyConvertV5RespData) GetFillToSzOk() (*string, bool)`

GetFillToSzOk returns a tuple with the FillToSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillToSz

`func (o *CreateTradeEasyConvertV5RespData) SetFillToSz(v string)`

SetFillToSz sets FillToSz field to given value.

### HasFillToSz

`func (o *CreateTradeEasyConvertV5RespData) HasFillToSz() bool`

HasFillToSz returns a boolean if a field has been set.

### GetFromCcy

`func (o *CreateTradeEasyConvertV5RespData) GetFromCcy() string`

GetFromCcy returns the FromCcy field if non-nil, zero value otherwise.

### GetFromCcyOk

`func (o *CreateTradeEasyConvertV5RespData) GetFromCcyOk() (*string, bool)`

GetFromCcyOk returns a tuple with the FromCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCcy

`func (o *CreateTradeEasyConvertV5RespData) SetFromCcy(v string)`

SetFromCcy sets FromCcy field to given value.

### HasFromCcy

`func (o *CreateTradeEasyConvertV5RespData) HasFromCcy() bool`

HasFromCcy returns a boolean if a field has been set.

### GetStatus

`func (o *CreateTradeEasyConvertV5RespData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateTradeEasyConvertV5RespData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateTradeEasyConvertV5RespData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateTradeEasyConvertV5RespData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToCcy

`func (o *CreateTradeEasyConvertV5RespData) GetToCcy() string`

GetToCcy returns the ToCcy field if non-nil, zero value otherwise.

### GetToCcyOk

`func (o *CreateTradeEasyConvertV5RespData) GetToCcyOk() (*string, bool)`

GetToCcyOk returns a tuple with the ToCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCcy

`func (o *CreateTradeEasyConvertV5RespData) SetToCcy(v string)`

SetToCcy sets ToCcy field to given value.

### HasToCcy

`func (o *CreateTradeEasyConvertV5RespData) HasToCcy() bool`

HasToCcy returns a boolean if a field has been set.

### GetUTime

`func (o *CreateTradeEasyConvertV5RespData) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *CreateTradeEasyConvertV5RespData) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *CreateTradeEasyConvertV5RespData) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *CreateTradeEasyConvertV5RespData) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


