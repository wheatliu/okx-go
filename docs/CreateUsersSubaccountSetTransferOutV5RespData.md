# CreateUsersSubaccountSetTransferOutV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanTransOut** | Pointer to **bool** | Whether the sub-account has the right to transfer out.   &#x60;false&#x60;: cannot transfer out   &#x60;true&#x60;: can transfer out | [optional] 
**SubAcct** | Pointer to **string** | Name of the sub-account | [optional] [default to ""]

## Methods

### NewCreateUsersSubaccountSetTransferOutV5RespData

`func NewCreateUsersSubaccountSetTransferOutV5RespData() *CreateUsersSubaccountSetTransferOutV5RespData`

NewCreateUsersSubaccountSetTransferOutV5RespData instantiates a new CreateUsersSubaccountSetTransferOutV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUsersSubaccountSetTransferOutV5RespDataWithDefaults

`func NewCreateUsersSubaccountSetTransferOutV5RespDataWithDefaults() *CreateUsersSubaccountSetTransferOutV5RespData`

NewCreateUsersSubaccountSetTransferOutV5RespDataWithDefaults instantiates a new CreateUsersSubaccountSetTransferOutV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanTransOut

`func (o *CreateUsersSubaccountSetTransferOutV5RespData) GetCanTransOut() bool`

GetCanTransOut returns the CanTransOut field if non-nil, zero value otherwise.

### GetCanTransOutOk

`func (o *CreateUsersSubaccountSetTransferOutV5RespData) GetCanTransOutOk() (*bool, bool)`

GetCanTransOutOk returns a tuple with the CanTransOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTransOut

`func (o *CreateUsersSubaccountSetTransferOutV5RespData) SetCanTransOut(v bool)`

SetCanTransOut sets CanTransOut field to given value.

### HasCanTransOut

`func (o *CreateUsersSubaccountSetTransferOutV5RespData) HasCanTransOut() bool`

HasCanTransOut returns a boolean if a field has been set.

### GetSubAcct

`func (o *CreateUsersSubaccountSetTransferOutV5RespData) GetSubAcct() string`

GetSubAcct returns the SubAcct field if non-nil, zero value otherwise.

### GetSubAcctOk

`func (o *CreateUsersSubaccountSetTransferOutV5RespData) GetSubAcctOk() (*string, bool)`

GetSubAcctOk returns a tuple with the SubAcct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAcct

`func (o *CreateUsersSubaccountSetTransferOutV5RespData) SetSubAcct(v string)`

SetSubAcct sets SubAcct field to given value.

### HasSubAcct

`func (o *CreateUsersSubaccountSetTransferOutV5RespData) HasSubAcct() bool`

HasSubAcct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


