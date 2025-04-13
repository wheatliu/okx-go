# CreateSprdOrderV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client | [optional] [default to ""]
**OrdId** | Pointer to **string** | Order ID | [optional] [default to ""]
**SCode** | Pointer to **string** | The code of the event execution result, 0 means success. | [optional] [default to ""]
**SMsg** | Pointer to **string** | Rejection or success message of event execution. | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]

## Methods

### NewCreateSprdOrderV5RespDataInner

`func NewCreateSprdOrderV5RespDataInner() *CreateSprdOrderV5RespDataInner`

NewCreateSprdOrderV5RespDataInner instantiates a new CreateSprdOrderV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSprdOrderV5RespDataInnerWithDefaults

`func NewCreateSprdOrderV5RespDataInnerWithDefaults() *CreateSprdOrderV5RespDataInner`

NewCreateSprdOrderV5RespDataInnerWithDefaults instantiates a new CreateSprdOrderV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClOrdId

`func (o *CreateSprdOrderV5RespDataInner) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *CreateSprdOrderV5RespDataInner) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *CreateSprdOrderV5RespDataInner) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *CreateSprdOrderV5RespDataInner) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetOrdId

`func (o *CreateSprdOrderV5RespDataInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *CreateSprdOrderV5RespDataInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *CreateSprdOrderV5RespDataInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *CreateSprdOrderV5RespDataInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetSCode

`func (o *CreateSprdOrderV5RespDataInner) GetSCode() string`

GetSCode returns the SCode field if non-nil, zero value otherwise.

### GetSCodeOk

`func (o *CreateSprdOrderV5RespDataInner) GetSCodeOk() (*string, bool)`

GetSCodeOk returns a tuple with the SCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCode

`func (o *CreateSprdOrderV5RespDataInner) SetSCode(v string)`

SetSCode sets SCode field to given value.

### HasSCode

`func (o *CreateSprdOrderV5RespDataInner) HasSCode() bool`

HasSCode returns a boolean if a field has been set.

### GetSMsg

`func (o *CreateSprdOrderV5RespDataInner) GetSMsg() string`

GetSMsg returns the SMsg field if non-nil, zero value otherwise.

### GetSMsgOk

`func (o *CreateSprdOrderV5RespDataInner) GetSMsgOk() (*string, bool)`

GetSMsgOk returns a tuple with the SMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMsg

`func (o *CreateSprdOrderV5RespDataInner) SetSMsg(v string)`

SetSMsg sets SMsg field to given value.

### HasSMsg

`func (o *CreateSprdOrderV5RespDataInner) HasSMsg() bool`

HasSMsg returns a boolean if a field has been set.

### GetTag

`func (o *CreateSprdOrderV5RespDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateSprdOrderV5RespDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateSprdOrderV5RespDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateSprdOrderV5RespDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


