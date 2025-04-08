# CreateRfqCancelAllRfqsV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The result code, &#x60;0&#x60; means success. | [optional] [default to ""]
**Data** | Pointer to [**[]CreateRfqCancelAllRfqsV5RespDataDataInner**](CreateRfqCancelAllRfqsV5RespDataDataInner.md) | Array of objects containing the results | [optional] 
**Msg** | Pointer to **string** | The error message, not empty if the code is not 0. | [optional] [default to ""]

## Methods

### NewCreateRfqCancelAllRfqsV5RespData

`func NewCreateRfqCancelAllRfqsV5RespData() *CreateRfqCancelAllRfqsV5RespData`

NewCreateRfqCancelAllRfqsV5RespData instantiates a new CreateRfqCancelAllRfqsV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqCancelAllRfqsV5RespDataWithDefaults

`func NewCreateRfqCancelAllRfqsV5RespDataWithDefaults() *CreateRfqCancelAllRfqsV5RespData`

NewCreateRfqCancelAllRfqsV5RespDataWithDefaults instantiates a new CreateRfqCancelAllRfqsV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateRfqCancelAllRfqsV5RespData) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateRfqCancelAllRfqsV5RespData) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateRfqCancelAllRfqsV5RespData) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateRfqCancelAllRfqsV5RespData) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *CreateRfqCancelAllRfqsV5RespData) GetData() []CreateRfqCancelAllRfqsV5RespDataDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateRfqCancelAllRfqsV5RespData) GetDataOk() (*[]CreateRfqCancelAllRfqsV5RespDataDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateRfqCancelAllRfqsV5RespData) SetData(v []CreateRfqCancelAllRfqsV5RespDataDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *CreateRfqCancelAllRfqsV5RespData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CreateRfqCancelAllRfqsV5RespData) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateRfqCancelAllRfqsV5RespData) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateRfqCancelAllRfqsV5RespData) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CreateRfqCancelAllRfqsV5RespData) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


