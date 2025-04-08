# CreateAccountSetGreeksV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GreeksType** | **string** | Display  type of Greeks.  &#x60;PA&#x60;: Greeks in coins   &#x60;BS&#x60;: Black-Scholes Greeks in dollars | [default to ""]

## Methods

### NewCreateAccountSetGreeksV5Req

`func NewCreateAccountSetGreeksV5Req(greeksType string, ) *CreateAccountSetGreeksV5Req`

NewCreateAccountSetGreeksV5Req instantiates a new CreateAccountSetGreeksV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountSetGreeksV5ReqWithDefaults

`func NewCreateAccountSetGreeksV5ReqWithDefaults() *CreateAccountSetGreeksV5Req`

NewCreateAccountSetGreeksV5ReqWithDefaults instantiates a new CreateAccountSetGreeksV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGreeksType

`func (o *CreateAccountSetGreeksV5Req) GetGreeksType() string`

GetGreeksType returns the GreeksType field if non-nil, zero value otherwise.

### GetGreeksTypeOk

`func (o *CreateAccountSetGreeksV5Req) GetGreeksTypeOk() (*string, bool)`

GetGreeksTypeOk returns a tuple with the GreeksType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreeksType

`func (o *CreateAccountSetGreeksV5Req) SetGreeksType(v string)`

SetGreeksType sets GreeksType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


