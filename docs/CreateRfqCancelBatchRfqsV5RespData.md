# CreateRfqCancelBatchRfqsV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The result code, &#x60;0&#x60; means success. | [optional] [default to ""]
**Data** | Pointer to [**[]CreateRfqCancelBatchRfqsV5RespDataDataInner**](CreateRfqCancelBatchRfqsV5RespDataDataInner.md) | Array of objects containing the results | [optional] 
**Msg** | Pointer to **string** | The error message, not empty if the code is not 0. | [optional] [default to ""]

## Methods

### NewCreateRfqCancelBatchRfqsV5RespData

`func NewCreateRfqCancelBatchRfqsV5RespData() *CreateRfqCancelBatchRfqsV5RespData`

NewCreateRfqCancelBatchRfqsV5RespData instantiates a new CreateRfqCancelBatchRfqsV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqCancelBatchRfqsV5RespDataWithDefaults

`func NewCreateRfqCancelBatchRfqsV5RespDataWithDefaults() *CreateRfqCancelBatchRfqsV5RespData`

NewCreateRfqCancelBatchRfqsV5RespDataWithDefaults instantiates a new CreateRfqCancelBatchRfqsV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateRfqCancelBatchRfqsV5RespData) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateRfqCancelBatchRfqsV5RespData) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateRfqCancelBatchRfqsV5RespData) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateRfqCancelBatchRfqsV5RespData) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *CreateRfqCancelBatchRfqsV5RespData) GetData() []CreateRfqCancelBatchRfqsV5RespDataDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateRfqCancelBatchRfqsV5RespData) GetDataOk() (*[]CreateRfqCancelBatchRfqsV5RespDataDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateRfqCancelBatchRfqsV5RespData) SetData(v []CreateRfqCancelBatchRfqsV5RespDataDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *CreateRfqCancelBatchRfqsV5RespData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CreateRfqCancelBatchRfqsV5RespData) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateRfqCancelBatchRfqsV5RespData) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateRfqCancelBatchRfqsV5RespData) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CreateRfqCancelBatchRfqsV5RespData) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


