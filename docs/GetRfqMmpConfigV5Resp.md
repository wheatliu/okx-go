# GetRfqMmpConfigV5Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] [default to ""]
**Data** | Pointer to [**GetRfqMmpConfigV5RespData**](GetRfqMmpConfigV5RespData.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewGetRfqMmpConfigV5Resp

`func NewGetRfqMmpConfigV5Resp() *GetRfqMmpConfigV5Resp`

NewGetRfqMmpConfigV5Resp instantiates a new GetRfqMmpConfigV5Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRfqMmpConfigV5RespWithDefaults

`func NewGetRfqMmpConfigV5RespWithDefaults() *GetRfqMmpConfigV5Resp`

NewGetRfqMmpConfigV5RespWithDefaults instantiates a new GetRfqMmpConfigV5Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetRfqMmpConfigV5Resp) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetRfqMmpConfigV5Resp) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetRfqMmpConfigV5Resp) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetRfqMmpConfigV5Resp) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *GetRfqMmpConfigV5Resp) GetData() GetRfqMmpConfigV5RespData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRfqMmpConfigV5Resp) GetDataOk() (*GetRfqMmpConfigV5RespData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRfqMmpConfigV5Resp) SetData(v GetRfqMmpConfigV5RespData)`

SetData sets Data field to given value.

### HasData

`func (o *GetRfqMmpConfigV5Resp) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *GetRfqMmpConfigV5Resp) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetRfqMmpConfigV5Resp) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetRfqMmpConfigV5Resp) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetRfqMmpConfigV5Resp) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


