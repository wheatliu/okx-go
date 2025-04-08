# CreateTradeMassCancelV5Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] [default to ""]
**Data** | Pointer to [**[]CreateAccountMmpResetV5RespDataInner**](CreateAccountMmpResetV5RespDataInner.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewCreateTradeMassCancelV5Resp

`func NewCreateTradeMassCancelV5Resp() *CreateTradeMassCancelV5Resp`

NewCreateTradeMassCancelV5Resp instantiates a new CreateTradeMassCancelV5Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeMassCancelV5RespWithDefaults

`func NewCreateTradeMassCancelV5RespWithDefaults() *CreateTradeMassCancelV5Resp`

NewCreateTradeMassCancelV5RespWithDefaults instantiates a new CreateTradeMassCancelV5Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateTradeMassCancelV5Resp) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateTradeMassCancelV5Resp) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateTradeMassCancelV5Resp) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateTradeMassCancelV5Resp) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *CreateTradeMassCancelV5Resp) GetData() []CreateAccountMmpResetV5RespDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateTradeMassCancelV5Resp) GetDataOk() (*[]CreateAccountMmpResetV5RespDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateTradeMassCancelV5Resp) SetData(v []CreateAccountMmpResetV5RespDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *CreateTradeMassCancelV5Resp) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CreateTradeMassCancelV5Resp) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateTradeMassCancelV5Resp) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateTradeMassCancelV5Resp) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CreateTradeMassCancelV5Resp) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


