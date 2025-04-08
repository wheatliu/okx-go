# CreateSprdCancelOrderV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client | [optional] [default to ""]
**OrdId** | Pointer to **string** | Order ID   Either &#x60;ordId&#x60; or &#x60;clOrdId&#x60; is required. If both are passed, &#x60;ordId&#x60; will be used. | [optional] [default to ""]

## Methods

### NewCreateSprdCancelOrderV5Req

`func NewCreateSprdCancelOrderV5Req() *CreateSprdCancelOrderV5Req`

NewCreateSprdCancelOrderV5Req instantiates a new CreateSprdCancelOrderV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSprdCancelOrderV5ReqWithDefaults

`func NewCreateSprdCancelOrderV5ReqWithDefaults() *CreateSprdCancelOrderV5Req`

NewCreateSprdCancelOrderV5ReqWithDefaults instantiates a new CreateSprdCancelOrderV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClOrdId

`func (o *CreateSprdCancelOrderV5Req) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *CreateSprdCancelOrderV5Req) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *CreateSprdCancelOrderV5Req) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *CreateSprdCancelOrderV5Req) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetOrdId

`func (o *CreateSprdCancelOrderV5Req) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *CreateSprdCancelOrderV5Req) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *CreateSprdCancelOrderV5Req) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *CreateSprdCancelOrderV5Req) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


