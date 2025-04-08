# GetAccountInterestLimitsV5Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] [default to ""]
**Data** | Pointer to [**[]GetAccountInterestLimitsV5RespDataInner**](GetAccountInterestLimitsV5RespDataInner.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewGetAccountInterestLimitsV5Resp

`func NewGetAccountInterestLimitsV5Resp() *GetAccountInterestLimitsV5Resp`

NewGetAccountInterestLimitsV5Resp instantiates a new GetAccountInterestLimitsV5Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountInterestLimitsV5RespWithDefaults

`func NewGetAccountInterestLimitsV5RespWithDefaults() *GetAccountInterestLimitsV5Resp`

NewGetAccountInterestLimitsV5RespWithDefaults instantiates a new GetAccountInterestLimitsV5Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetAccountInterestLimitsV5Resp) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetAccountInterestLimitsV5Resp) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetAccountInterestLimitsV5Resp) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetAccountInterestLimitsV5Resp) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *GetAccountInterestLimitsV5Resp) GetData() []GetAccountInterestLimitsV5RespDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAccountInterestLimitsV5Resp) GetDataOk() (*[]GetAccountInterestLimitsV5RespDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAccountInterestLimitsV5Resp) SetData(v []GetAccountInterestLimitsV5RespDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetAccountInterestLimitsV5Resp) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *GetAccountInterestLimitsV5Resp) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetAccountInterestLimitsV5Resp) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetAccountInterestLimitsV5Resp) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetAccountInterestLimitsV5Resp) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


