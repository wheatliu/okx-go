# CreateUsersSubaccountSetTransferOutV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanTransOut** | Pointer to **bool** | Whether the sub-account has the right to transfer out. The default is &#x60;true&#x60;.  &#x60;false&#x60;: cannot transfer out   &#x60;true&#x60;: can transfer out | [optional] 
**SubAcct** | **string** | Name of the sub-account. Single sub-account or multiple sub-account (no more than 20) separated with comma. | [default to ""]

## Methods

### NewCreateUsersSubaccountSetTransferOutV5Req

`func NewCreateUsersSubaccountSetTransferOutV5Req(subAcct string, ) *CreateUsersSubaccountSetTransferOutV5Req`

NewCreateUsersSubaccountSetTransferOutV5Req instantiates a new CreateUsersSubaccountSetTransferOutV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUsersSubaccountSetTransferOutV5ReqWithDefaults

`func NewCreateUsersSubaccountSetTransferOutV5ReqWithDefaults() *CreateUsersSubaccountSetTransferOutV5Req`

NewCreateUsersSubaccountSetTransferOutV5ReqWithDefaults instantiates a new CreateUsersSubaccountSetTransferOutV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanTransOut

`func (o *CreateUsersSubaccountSetTransferOutV5Req) GetCanTransOut() bool`

GetCanTransOut returns the CanTransOut field if non-nil, zero value otherwise.

### GetCanTransOutOk

`func (o *CreateUsersSubaccountSetTransferOutV5Req) GetCanTransOutOk() (*bool, bool)`

GetCanTransOutOk returns a tuple with the CanTransOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTransOut

`func (o *CreateUsersSubaccountSetTransferOutV5Req) SetCanTransOut(v bool)`

SetCanTransOut sets CanTransOut field to given value.

### HasCanTransOut

`func (o *CreateUsersSubaccountSetTransferOutV5Req) HasCanTransOut() bool`

HasCanTransOut returns a boolean if a field has been set.

### GetSubAcct

`func (o *CreateUsersSubaccountSetTransferOutV5Req) GetSubAcct() string`

GetSubAcct returns the SubAcct field if non-nil, zero value otherwise.

### GetSubAcctOk

`func (o *CreateUsersSubaccountSetTransferOutV5Req) GetSubAcctOk() (*string, bool)`

GetSubAcctOk returns a tuple with the SubAcct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAcct

`func (o *CreateUsersSubaccountSetTransferOutV5Req) SetSubAcct(v string)`

SetSubAcct sets SubAcct field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


