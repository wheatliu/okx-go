# CreateTradeCancelBatchOrdersV5RespDataInnerDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client | [optional] [default to ""]
**OrdId** | Pointer to **string** | Order ID | [optional] [default to ""]
**SCode** | Pointer to **string** | The code of the event execution result, &#x60;0&#x60; means success. | [optional] [default to ""]
**SMsg** | Pointer to **string** | Rejection message if the request is unsuccessful. | [optional] [default to ""]
**Ts** | Pointer to **string** | Timestamp when the order request processing is finished by our system, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewCreateTradeCancelBatchOrdersV5RespDataInnerDataInner

`func NewCreateTradeCancelBatchOrdersV5RespDataInnerDataInner() *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner`

NewCreateTradeCancelBatchOrdersV5RespDataInnerDataInner instantiates a new CreateTradeCancelBatchOrdersV5RespDataInnerDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeCancelBatchOrdersV5RespDataInnerDataInnerWithDefaults

`func NewCreateTradeCancelBatchOrdersV5RespDataInnerDataInnerWithDefaults() *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner`

NewCreateTradeCancelBatchOrdersV5RespDataInnerDataInnerWithDefaults instantiates a new CreateTradeCancelBatchOrdersV5RespDataInnerDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClOrdId

`func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetOrdId

`func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetSCode

`func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) GetSCode() string`

GetSCode returns the SCode field if non-nil, zero value otherwise.

### GetSCodeOk

`func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) GetSCodeOk() (*string, bool)`

GetSCodeOk returns a tuple with the SCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCode

`func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) SetSCode(v string)`

SetSCode sets SCode field to given value.

### HasSCode

`func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) HasSCode() bool`

HasSCode returns a boolean if a field has been set.

### GetSMsg

`func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) GetSMsg() string`

GetSMsg returns the SMsg field if non-nil, zero value otherwise.

### GetSMsgOk

`func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) GetSMsgOk() (*string, bool)`

GetSMsgOk returns a tuple with the SMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMsg

`func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) SetSMsg(v string)`

SetSMsg sets SMsg field to given value.

### HasSMsg

`func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) HasSMsg() bool`

HasSMsg returns a boolean if a field has been set.

### GetTs

`func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


