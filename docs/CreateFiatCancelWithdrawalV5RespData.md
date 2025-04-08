# CreateFiatCancelWithdrawalV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrdId** | Pointer to **string** | Payment Order ID | [optional] [default to ""]
**State** | Pointer to **string** | The state of the transaction, e.g.&#x60;canceled&#x60; | [optional] [default to ""]

## Methods

### NewCreateFiatCancelWithdrawalV5RespData

`func NewCreateFiatCancelWithdrawalV5RespData() *CreateFiatCancelWithdrawalV5RespData`

NewCreateFiatCancelWithdrawalV5RespData instantiates a new CreateFiatCancelWithdrawalV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFiatCancelWithdrawalV5RespDataWithDefaults

`func NewCreateFiatCancelWithdrawalV5RespDataWithDefaults() *CreateFiatCancelWithdrawalV5RespData`

NewCreateFiatCancelWithdrawalV5RespDataWithDefaults instantiates a new CreateFiatCancelWithdrawalV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrdId

`func (o *CreateFiatCancelWithdrawalV5RespData) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *CreateFiatCancelWithdrawalV5RespData) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *CreateFiatCancelWithdrawalV5RespData) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *CreateFiatCancelWithdrawalV5RespData) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetState

`func (o *CreateFiatCancelWithdrawalV5RespData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateFiatCancelWithdrawalV5RespData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateFiatCancelWithdrawalV5RespData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CreateFiatCancelWithdrawalV5RespData) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


