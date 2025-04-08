# CreateRfqCancelAllRfqsV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The result code, &#x60;0&#x60; means success. | [optional] [default to ""]
**Data** | Pointer to [**[]CreateRfqCancelAllRfqsV5RespDataInnerDataInner**](CreateRfqCancelAllRfqsV5RespDataInnerDataInner.md) | Array of objects containing the results | [optional] 
**Msg** | Pointer to **string** | The error message, not empty if the code is not 0. | [optional] [default to ""]

## Methods

### NewCreateRfqCancelAllRfqsV5RespDataInner

`func NewCreateRfqCancelAllRfqsV5RespDataInner() *CreateRfqCancelAllRfqsV5RespDataInner`

NewCreateRfqCancelAllRfqsV5RespDataInner instantiates a new CreateRfqCancelAllRfqsV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqCancelAllRfqsV5RespDataInnerWithDefaults

`func NewCreateRfqCancelAllRfqsV5RespDataInnerWithDefaults() *CreateRfqCancelAllRfqsV5RespDataInner`

NewCreateRfqCancelAllRfqsV5RespDataInnerWithDefaults instantiates a new CreateRfqCancelAllRfqsV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateRfqCancelAllRfqsV5RespDataInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateRfqCancelAllRfqsV5RespDataInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateRfqCancelAllRfqsV5RespDataInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateRfqCancelAllRfqsV5RespDataInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *CreateRfqCancelAllRfqsV5RespDataInner) GetData() []CreateRfqCancelAllRfqsV5RespDataInnerDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateRfqCancelAllRfqsV5RespDataInner) GetDataOk() (*[]CreateRfqCancelAllRfqsV5RespDataInnerDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateRfqCancelAllRfqsV5RespDataInner) SetData(v []CreateRfqCancelAllRfqsV5RespDataInnerDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *CreateRfqCancelAllRfqsV5RespDataInner) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CreateRfqCancelAllRfqsV5RespDataInner) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateRfqCancelAllRfqsV5RespDataInner) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateRfqCancelAllRfqsV5RespDataInner) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CreateRfqCancelAllRfqsV5RespDataInner) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


