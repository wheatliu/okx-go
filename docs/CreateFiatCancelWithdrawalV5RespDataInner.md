# CreateFiatCancelWithdrawalV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrdId** | Pointer to **string** | Payment Order ID | [optional] [default to ""]
**State** | Pointer to **string** | The state of the transaction, e.g.&#x60;canceled&#x60; | [optional] [default to ""]

## Methods

### NewCreateFiatCancelWithdrawalV5RespDataInner

`func NewCreateFiatCancelWithdrawalV5RespDataInner() *CreateFiatCancelWithdrawalV5RespDataInner`

NewCreateFiatCancelWithdrawalV5RespDataInner instantiates a new CreateFiatCancelWithdrawalV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFiatCancelWithdrawalV5RespDataInnerWithDefaults

`func NewCreateFiatCancelWithdrawalV5RespDataInnerWithDefaults() *CreateFiatCancelWithdrawalV5RespDataInner`

NewCreateFiatCancelWithdrawalV5RespDataInnerWithDefaults instantiates a new CreateFiatCancelWithdrawalV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrdId

`func (o *CreateFiatCancelWithdrawalV5RespDataInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *CreateFiatCancelWithdrawalV5RespDataInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *CreateFiatCancelWithdrawalV5RespDataInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *CreateFiatCancelWithdrawalV5RespDataInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetState

`func (o *CreateFiatCancelWithdrawalV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateFiatCancelWithdrawalV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateFiatCancelWithdrawalV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CreateFiatCancelWithdrawalV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


