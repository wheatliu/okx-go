# CreateAccountBillsHistoryArchiveV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **string** | Whether there is already a download link for this section   &#x60;true&#x60;: Existed, can check from \&quot;Get bills details (since 2021)\&quot;.   &#x60;false&#x60;: Does not exist and is generating, can check the download link after 2 hours   The data of file is in reverse chronological order using &#x60;billId&#x60;. | [optional] [default to ""]
**Ts** | Pointer to **string** | The first request time when the server receives. Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewCreateAccountBillsHistoryArchiveV5RespDataInner

`func NewCreateAccountBillsHistoryArchiveV5RespDataInner() *CreateAccountBillsHistoryArchiveV5RespDataInner`

NewCreateAccountBillsHistoryArchiveV5RespDataInner instantiates a new CreateAccountBillsHistoryArchiveV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountBillsHistoryArchiveV5RespDataInnerWithDefaults

`func NewCreateAccountBillsHistoryArchiveV5RespDataInnerWithDefaults() *CreateAccountBillsHistoryArchiveV5RespDataInner`

NewCreateAccountBillsHistoryArchiveV5RespDataInnerWithDefaults instantiates a new CreateAccountBillsHistoryArchiveV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *CreateAccountBillsHistoryArchiveV5RespDataInner) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CreateAccountBillsHistoryArchiveV5RespDataInner) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CreateAccountBillsHistoryArchiveV5RespDataInner) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *CreateAccountBillsHistoryArchiveV5RespDataInner) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetTs

`func (o *CreateAccountBillsHistoryArchiveV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *CreateAccountBillsHistoryArchiveV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *CreateAccountBillsHistoryArchiveV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *CreateAccountBillsHistoryArchiveV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


