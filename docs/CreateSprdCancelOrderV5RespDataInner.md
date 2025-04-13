# CreateSprdCancelOrderV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client | [optional] [default to ""]
**OrdId** | Pointer to **string** | Order ID | [optional] [default to ""]
**SCode** | Pointer to **string** | The code of the event execution result, 0 means success. | [optional] [default to ""]
**SMsg** | Pointer to **string** | Rejection message if the request is unsuccessful. | [optional] [default to ""]

## Methods

### NewCreateSprdCancelOrderV5RespDataInner

`func NewCreateSprdCancelOrderV5RespDataInner() *CreateSprdCancelOrderV5RespDataInner`

NewCreateSprdCancelOrderV5RespDataInner instantiates a new CreateSprdCancelOrderV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSprdCancelOrderV5RespDataInnerWithDefaults

`func NewCreateSprdCancelOrderV5RespDataInnerWithDefaults() *CreateSprdCancelOrderV5RespDataInner`

NewCreateSprdCancelOrderV5RespDataInnerWithDefaults instantiates a new CreateSprdCancelOrderV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClOrdId

`func (o *CreateSprdCancelOrderV5RespDataInner) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *CreateSprdCancelOrderV5RespDataInner) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *CreateSprdCancelOrderV5RespDataInner) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *CreateSprdCancelOrderV5RespDataInner) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetOrdId

`func (o *CreateSprdCancelOrderV5RespDataInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *CreateSprdCancelOrderV5RespDataInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *CreateSprdCancelOrderV5RespDataInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *CreateSprdCancelOrderV5RespDataInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetSCode

`func (o *CreateSprdCancelOrderV5RespDataInner) GetSCode() string`

GetSCode returns the SCode field if non-nil, zero value otherwise.

### GetSCodeOk

`func (o *CreateSprdCancelOrderV5RespDataInner) GetSCodeOk() (*string, bool)`

GetSCodeOk returns a tuple with the SCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCode

`func (o *CreateSprdCancelOrderV5RespDataInner) SetSCode(v string)`

SetSCode sets SCode field to given value.

### HasSCode

`func (o *CreateSprdCancelOrderV5RespDataInner) HasSCode() bool`

HasSCode returns a boolean if a field has been set.

### GetSMsg

`func (o *CreateSprdCancelOrderV5RespDataInner) GetSMsg() string`

GetSMsg returns the SMsg field if non-nil, zero value otherwise.

### GetSMsgOk

`func (o *CreateSprdCancelOrderV5RespDataInner) GetSMsgOk() (*string, bool)`

GetSMsgOk returns a tuple with the SMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMsg

`func (o *CreateSprdCancelOrderV5RespDataInner) SetSMsg(v string)`

SetSMsg sets SMsg field to given value.

### HasSMsg

`func (o *CreateSprdCancelOrderV5RespDataInner) HasSMsg() bool`

HasSMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


