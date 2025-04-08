# GetPublicOpenInterestV5Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] [default to ""]
**Data** | Pointer to [**[]GetPublicOpenInterestV5RespDataInner**](GetPublicOpenInterestV5RespDataInner.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewGetPublicOpenInterestV5Resp

`func NewGetPublicOpenInterestV5Resp() *GetPublicOpenInterestV5Resp`

NewGetPublicOpenInterestV5Resp instantiates a new GetPublicOpenInterestV5Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicOpenInterestV5RespWithDefaults

`func NewGetPublicOpenInterestV5RespWithDefaults() *GetPublicOpenInterestV5Resp`

NewGetPublicOpenInterestV5RespWithDefaults instantiates a new GetPublicOpenInterestV5Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetPublicOpenInterestV5Resp) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetPublicOpenInterestV5Resp) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetPublicOpenInterestV5Resp) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetPublicOpenInterestV5Resp) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *GetPublicOpenInterestV5Resp) GetData() []GetPublicOpenInterestV5RespDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetPublicOpenInterestV5Resp) GetDataOk() (*[]GetPublicOpenInterestV5RespDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetPublicOpenInterestV5Resp) SetData(v []GetPublicOpenInterestV5RespDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetPublicOpenInterestV5Resp) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *GetPublicOpenInterestV5Resp) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetPublicOpenInterestV5Resp) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetPublicOpenInterestV5Resp) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetPublicOpenInterestV5Resp) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


