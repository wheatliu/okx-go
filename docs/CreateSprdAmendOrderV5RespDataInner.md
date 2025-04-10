# CreateSprdAmendOrderV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client. | [optional] [default to ""]
**OrdId** | Pointer to **string** | Order ID | [optional] [default to ""]
**ReqId** | Pointer to **string** | Client Request ID as assigned by the client for order amendment. | [optional] [default to ""]
**SCode** | Pointer to **string** | The code of the event execution result, 0 means success. | [optional] [default to ""]
**SMsg** | Pointer to **string** | Rejection message if the request is unsuccessful. | [optional] [default to ""]

## Methods

### NewCreateSprdAmendOrderV5RespDataInner

`func NewCreateSprdAmendOrderV5RespDataInner() *CreateSprdAmendOrderV5RespDataInner`

NewCreateSprdAmendOrderV5RespDataInner instantiates a new CreateSprdAmendOrderV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSprdAmendOrderV5RespDataInnerWithDefaults

`func NewCreateSprdAmendOrderV5RespDataInnerWithDefaults() *CreateSprdAmendOrderV5RespDataInner`

NewCreateSprdAmendOrderV5RespDataInnerWithDefaults instantiates a new CreateSprdAmendOrderV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClOrdId

`func (o *CreateSprdAmendOrderV5RespDataInner) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *CreateSprdAmendOrderV5RespDataInner) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *CreateSprdAmendOrderV5RespDataInner) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *CreateSprdAmendOrderV5RespDataInner) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetOrdId

`func (o *CreateSprdAmendOrderV5RespDataInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *CreateSprdAmendOrderV5RespDataInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *CreateSprdAmendOrderV5RespDataInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *CreateSprdAmendOrderV5RespDataInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetReqId

`func (o *CreateSprdAmendOrderV5RespDataInner) GetReqId() string`

GetReqId returns the ReqId field if non-nil, zero value otherwise.

### GetReqIdOk

`func (o *CreateSprdAmendOrderV5RespDataInner) GetReqIdOk() (*string, bool)`

GetReqIdOk returns a tuple with the ReqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqId

`func (o *CreateSprdAmendOrderV5RespDataInner) SetReqId(v string)`

SetReqId sets ReqId field to given value.

### HasReqId

`func (o *CreateSprdAmendOrderV5RespDataInner) HasReqId() bool`

HasReqId returns a boolean if a field has been set.

### GetSCode

`func (o *CreateSprdAmendOrderV5RespDataInner) GetSCode() string`

GetSCode returns the SCode field if non-nil, zero value otherwise.

### GetSCodeOk

`func (o *CreateSprdAmendOrderV5RespDataInner) GetSCodeOk() (*string, bool)`

GetSCodeOk returns a tuple with the SCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCode

`func (o *CreateSprdAmendOrderV5RespDataInner) SetSCode(v string)`

SetSCode sets SCode field to given value.

### HasSCode

`func (o *CreateSprdAmendOrderV5RespDataInner) HasSCode() bool`

HasSCode returns a boolean if a field has been set.

### GetSMsg

`func (o *CreateSprdAmendOrderV5RespDataInner) GetSMsg() string`

GetSMsg returns the SMsg field if non-nil, zero value otherwise.

### GetSMsgOk

`func (o *CreateSprdAmendOrderV5RespDataInner) GetSMsgOk() (*string, bool)`

GetSMsgOk returns a tuple with the SMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMsg

`func (o *CreateSprdAmendOrderV5RespDataInner) SetSMsg(v string)`

SetSMsg sets SMsg field to given value.

### HasSMsg

`func (o *CreateSprdAmendOrderV5RespDataInner) HasSMsg() bool`

HasSMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


