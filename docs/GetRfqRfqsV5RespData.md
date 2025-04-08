# GetRfqRfqsV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The result code, &#x60;0&#x60; means success. | [optional] [default to ""]
**Data** | Pointer to [**[]GetRfqRfqsV5RespDataDataInner**](GetRfqRfqsV5RespDataDataInner.md) | Array of objects containing the results of the RFQ creation. | [optional] 
**Msg** | Pointer to **string** | The error message, not empty if the code is not 0. | [optional] [default to ""]

## Methods

### NewGetRfqRfqsV5RespData

`func NewGetRfqRfqsV5RespData() *GetRfqRfqsV5RespData`

NewGetRfqRfqsV5RespData instantiates a new GetRfqRfqsV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRfqRfqsV5RespDataWithDefaults

`func NewGetRfqRfqsV5RespDataWithDefaults() *GetRfqRfqsV5RespData`

NewGetRfqRfqsV5RespDataWithDefaults instantiates a new GetRfqRfqsV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetRfqRfqsV5RespData) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetRfqRfqsV5RespData) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetRfqRfqsV5RespData) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetRfqRfqsV5RespData) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *GetRfqRfqsV5RespData) GetData() []GetRfqRfqsV5RespDataDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRfqRfqsV5RespData) GetDataOk() (*[]GetRfqRfqsV5RespDataDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRfqRfqsV5RespData) SetData(v []GetRfqRfqsV5RespDataDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetRfqRfqsV5RespData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *GetRfqRfqsV5RespData) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetRfqRfqsV5RespData) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetRfqRfqsV5RespData) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetRfqRfqsV5RespData) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


