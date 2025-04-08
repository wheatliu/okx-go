# CreateAccountSetAutoRepayV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRepay** | **bool** | Whether auto repay is allowed or not under &#x60;Spot mode&#x60;  &#x60;true&#x60;: Enable auto repay  &#x60;false&#x60;: Disable auto repay | 

## Methods

### NewCreateAccountSetAutoRepayV5Req

`func NewCreateAccountSetAutoRepayV5Req(autoRepay bool, ) *CreateAccountSetAutoRepayV5Req`

NewCreateAccountSetAutoRepayV5Req instantiates a new CreateAccountSetAutoRepayV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountSetAutoRepayV5ReqWithDefaults

`func NewCreateAccountSetAutoRepayV5ReqWithDefaults() *CreateAccountSetAutoRepayV5Req`

NewCreateAccountSetAutoRepayV5ReqWithDefaults instantiates a new CreateAccountSetAutoRepayV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoRepay

`func (o *CreateAccountSetAutoRepayV5Req) GetAutoRepay() bool`

GetAutoRepay returns the AutoRepay field if non-nil, zero value otherwise.

### GetAutoRepayOk

`func (o *CreateAccountSetAutoRepayV5Req) GetAutoRepayOk() (*bool, bool)`

GetAutoRepayOk returns a tuple with the AutoRepay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRepay

`func (o *CreateAccountSetAutoRepayV5Req) SetAutoRepay(v bool)`

SetAutoRepay sets AutoRepay field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


