# GetRfqMakerInstrumentSettingsV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The result code, &#x60;0&#x60; means success. | [optional] [default to ""]
**Data** | Pointer to [**[]GetRfqMakerInstrumentSettingsV5RespDataDataInner**](GetRfqMakerInstrumentSettingsV5RespDataDataInner.md) | Return data of the request. | [optional] 
**Msg** | Pointer to **string** | The error message, not empty if the code is not &#x60;0&#x60;. | [optional] [default to ""]

## Methods

### NewGetRfqMakerInstrumentSettingsV5RespData

`func NewGetRfqMakerInstrumentSettingsV5RespData() *GetRfqMakerInstrumentSettingsV5RespData`

NewGetRfqMakerInstrumentSettingsV5RespData instantiates a new GetRfqMakerInstrumentSettingsV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRfqMakerInstrumentSettingsV5RespDataWithDefaults

`func NewGetRfqMakerInstrumentSettingsV5RespDataWithDefaults() *GetRfqMakerInstrumentSettingsV5RespData`

NewGetRfqMakerInstrumentSettingsV5RespDataWithDefaults instantiates a new GetRfqMakerInstrumentSettingsV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetRfqMakerInstrumentSettingsV5RespData) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetRfqMakerInstrumentSettingsV5RespData) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetRfqMakerInstrumentSettingsV5RespData) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetRfqMakerInstrumentSettingsV5RespData) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *GetRfqMakerInstrumentSettingsV5RespData) GetData() []GetRfqMakerInstrumentSettingsV5RespDataDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRfqMakerInstrumentSettingsV5RespData) GetDataOk() (*[]GetRfqMakerInstrumentSettingsV5RespDataDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRfqMakerInstrumentSettingsV5RespData) SetData(v []GetRfqMakerInstrumentSettingsV5RespDataDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetRfqMakerInstrumentSettingsV5RespData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *GetRfqMakerInstrumentSettingsV5RespData) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetRfqMakerInstrumentSettingsV5RespData) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetRfqMakerInstrumentSettingsV5RespData) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetRfqMakerInstrumentSettingsV5RespData) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


