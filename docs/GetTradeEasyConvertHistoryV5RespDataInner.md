# GetTradeEasyConvertHistoryV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acct** | Pointer to **string** | The account where the mainstream currency is located  &#x60;6&#x60;: Funding account   &#x60;18&#x60;: Trading account | [optional] [default to ""]
**FillFromSz** | Pointer to **string** | Amount of small payment currency convert from | [optional] [default to ""]
**FillToSz** | Pointer to **string** | Amount of mainstream currency convert to | [optional] [default to ""]
**FromCcy** | Pointer to **string** | Type of small payment currency convert from | [optional] [default to ""]
**Status** | Pointer to **string** | Current status of easy convert   &#x60;running&#x60;: Running   &#x60;filled&#x60;: Filled   &#x60;failed&#x60;: Failed | [optional] [default to ""]
**ToCcy** | Pointer to **string** | Type of mainstream currency convert to | [optional] [default to ""]
**UTime** | Pointer to **string** | Trade time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetTradeEasyConvertHistoryV5RespDataInner

`func NewGetTradeEasyConvertHistoryV5RespDataInner() *GetTradeEasyConvertHistoryV5RespDataInner`

NewGetTradeEasyConvertHistoryV5RespDataInner instantiates a new GetTradeEasyConvertHistoryV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradeEasyConvertHistoryV5RespDataInnerWithDefaults

`func NewGetTradeEasyConvertHistoryV5RespDataInnerWithDefaults() *GetTradeEasyConvertHistoryV5RespDataInner`

NewGetTradeEasyConvertHistoryV5RespDataInnerWithDefaults instantiates a new GetTradeEasyConvertHistoryV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcct

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) GetAcct() string`

GetAcct returns the Acct field if non-nil, zero value otherwise.

### GetAcctOk

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) GetAcctOk() (*string, bool)`

GetAcctOk returns a tuple with the Acct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcct

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) SetAcct(v string)`

SetAcct sets Acct field to given value.

### HasAcct

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) HasAcct() bool`

HasAcct returns a boolean if a field has been set.

### GetFillFromSz

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) GetFillFromSz() string`

GetFillFromSz returns the FillFromSz field if non-nil, zero value otherwise.

### GetFillFromSzOk

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) GetFillFromSzOk() (*string, bool)`

GetFillFromSzOk returns a tuple with the FillFromSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillFromSz

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) SetFillFromSz(v string)`

SetFillFromSz sets FillFromSz field to given value.

### HasFillFromSz

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) HasFillFromSz() bool`

HasFillFromSz returns a boolean if a field has been set.

### GetFillToSz

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) GetFillToSz() string`

GetFillToSz returns the FillToSz field if non-nil, zero value otherwise.

### GetFillToSzOk

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) GetFillToSzOk() (*string, bool)`

GetFillToSzOk returns a tuple with the FillToSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillToSz

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) SetFillToSz(v string)`

SetFillToSz sets FillToSz field to given value.

### HasFillToSz

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) HasFillToSz() bool`

HasFillToSz returns a boolean if a field has been set.

### GetFromCcy

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) GetFromCcy() string`

GetFromCcy returns the FromCcy field if non-nil, zero value otherwise.

### GetFromCcyOk

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) GetFromCcyOk() (*string, bool)`

GetFromCcyOk returns a tuple with the FromCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCcy

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) SetFromCcy(v string)`

SetFromCcy sets FromCcy field to given value.

### HasFromCcy

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) HasFromCcy() bool`

HasFromCcy returns a boolean if a field has been set.

### GetStatus

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToCcy

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) GetToCcy() string`

GetToCcy returns the ToCcy field if non-nil, zero value otherwise.

### GetToCcyOk

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) GetToCcyOk() (*string, bool)`

GetToCcyOk returns a tuple with the ToCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCcy

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) SetToCcy(v string)`

SetToCcy sets ToCcy field to given value.

### HasToCcy

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) HasToCcy() bool`

HasToCcy returns a boolean if a field has been set.

### GetUTime

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetTradeEasyConvertHistoryV5RespDataInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


