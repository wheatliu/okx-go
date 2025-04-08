# CreateRfqMakerInstrumentSettingsV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The result code, &#x60;0&#x60; means success. | [optional] [default to ""]
**Data** | Pointer to [**[]CreateRfqMakerInstrumentSettingsV5RespDataInnerDataInner**](CreateRfqMakerInstrumentSettingsV5RespDataInnerDataInner.md) | Array of objects containing the results. | [optional] 
**Msg** | Pointer to **string** | The error message, not empty if the code is not &#x60;0&#x60;. | [optional] [default to ""]

## Methods

### NewCreateRfqMakerInstrumentSettingsV5RespDataInner

`func NewCreateRfqMakerInstrumentSettingsV5RespDataInner() *CreateRfqMakerInstrumentSettingsV5RespDataInner`

NewCreateRfqMakerInstrumentSettingsV5RespDataInner instantiates a new CreateRfqMakerInstrumentSettingsV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqMakerInstrumentSettingsV5RespDataInnerWithDefaults

`func NewCreateRfqMakerInstrumentSettingsV5RespDataInnerWithDefaults() *CreateRfqMakerInstrumentSettingsV5RespDataInner`

NewCreateRfqMakerInstrumentSettingsV5RespDataInnerWithDefaults instantiates a new CreateRfqMakerInstrumentSettingsV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateRfqMakerInstrumentSettingsV5RespDataInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateRfqMakerInstrumentSettingsV5RespDataInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateRfqMakerInstrumentSettingsV5RespDataInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateRfqMakerInstrumentSettingsV5RespDataInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *CreateRfqMakerInstrumentSettingsV5RespDataInner) GetData() []CreateRfqMakerInstrumentSettingsV5RespDataInnerDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateRfqMakerInstrumentSettingsV5RespDataInner) GetDataOk() (*[]CreateRfqMakerInstrumentSettingsV5RespDataInnerDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateRfqMakerInstrumentSettingsV5RespDataInner) SetData(v []CreateRfqMakerInstrumentSettingsV5RespDataInnerDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *CreateRfqMakerInstrumentSettingsV5RespDataInner) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CreateRfqMakerInstrumentSettingsV5RespDataInner) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateRfqMakerInstrumentSettingsV5RespDataInner) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateRfqMakerInstrumentSettingsV5RespDataInner) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CreateRfqMakerInstrumentSettingsV5RespDataInner) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


