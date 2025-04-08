# GetUsersSubaccountListV5Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] [default to ""]
**Data** | Pointer to [**[]GetUsersSubaccountListV5RespDataInner**](GetUsersSubaccountListV5RespDataInner.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewGetUsersSubaccountListV5Resp

`func NewGetUsersSubaccountListV5Resp() *GetUsersSubaccountListV5Resp`

NewGetUsersSubaccountListV5Resp instantiates a new GetUsersSubaccountListV5Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUsersSubaccountListV5RespWithDefaults

`func NewGetUsersSubaccountListV5RespWithDefaults() *GetUsersSubaccountListV5Resp`

NewGetUsersSubaccountListV5RespWithDefaults instantiates a new GetUsersSubaccountListV5Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetUsersSubaccountListV5Resp) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetUsersSubaccountListV5Resp) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetUsersSubaccountListV5Resp) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetUsersSubaccountListV5Resp) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *GetUsersSubaccountListV5Resp) GetData() []GetUsersSubaccountListV5RespDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUsersSubaccountListV5Resp) GetDataOk() (*[]GetUsersSubaccountListV5RespDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUsersSubaccountListV5Resp) SetData(v []GetUsersSubaccountListV5RespDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetUsersSubaccountListV5Resp) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *GetUsersSubaccountListV5Resp) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetUsersSubaccountListV5Resp) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetUsersSubaccountListV5Resp) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetUsersSubaccountListV5Resp) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


