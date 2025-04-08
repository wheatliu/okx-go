# CreateAccountSetAutoRepayV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRepay** | Pointer to **bool** | Whether auto repay is allowed or not under &#x60;Spot mode&#x60;  &#x60;true&#x60;: Enable auto repay  &#x60;false&#x60;: Disable auto repay | [optional] 

## Methods

### NewCreateAccountSetAutoRepayV5RespDataInner

`func NewCreateAccountSetAutoRepayV5RespDataInner() *CreateAccountSetAutoRepayV5RespDataInner`

NewCreateAccountSetAutoRepayV5RespDataInner instantiates a new CreateAccountSetAutoRepayV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountSetAutoRepayV5RespDataInnerWithDefaults

`func NewCreateAccountSetAutoRepayV5RespDataInnerWithDefaults() *CreateAccountSetAutoRepayV5RespDataInner`

NewCreateAccountSetAutoRepayV5RespDataInnerWithDefaults instantiates a new CreateAccountSetAutoRepayV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoRepay

`func (o *CreateAccountSetAutoRepayV5RespDataInner) GetAutoRepay() bool`

GetAutoRepay returns the AutoRepay field if non-nil, zero value otherwise.

### GetAutoRepayOk

`func (o *CreateAccountSetAutoRepayV5RespDataInner) GetAutoRepayOk() (*bool, bool)`

GetAutoRepayOk returns a tuple with the AutoRepay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRepay

`func (o *CreateAccountSetAutoRepayV5RespDataInner) SetAutoRepay(v bool)`

SetAutoRepay sets AutoRepay field to given value.

### HasAutoRepay

`func (o *CreateAccountSetAutoRepayV5RespDataInner) HasAutoRepay() bool`

HasAutoRepay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


