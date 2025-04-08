# CreateAccountMmpResetV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstFamily** | **string** | Instrument family | [default to ""]
**InstType** | Pointer to **string** | Instrument type  &#x60;OPTION&#x60;  The default is &#x60;OPTION | [optional] [default to ""]

## Methods

### NewCreateAccountMmpResetV5Req

`func NewCreateAccountMmpResetV5Req(instFamily string, ) *CreateAccountMmpResetV5Req`

NewCreateAccountMmpResetV5Req instantiates a new CreateAccountMmpResetV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountMmpResetV5ReqWithDefaults

`func NewCreateAccountMmpResetV5ReqWithDefaults() *CreateAccountMmpResetV5Req`

NewCreateAccountMmpResetV5ReqWithDefaults instantiates a new CreateAccountMmpResetV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstFamily

`func (o *CreateAccountMmpResetV5Req) GetInstFamily() string`

GetInstFamily returns the InstFamily field if non-nil, zero value otherwise.

### GetInstFamilyOk

`func (o *CreateAccountMmpResetV5Req) GetInstFamilyOk() (*string, bool)`

GetInstFamilyOk returns a tuple with the InstFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstFamily

`func (o *CreateAccountMmpResetV5Req) SetInstFamily(v string)`

SetInstFamily sets InstFamily field to given value.


### GetInstType

`func (o *CreateAccountMmpResetV5Req) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *CreateAccountMmpResetV5Req) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *CreateAccountMmpResetV5Req) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *CreateAccountMmpResetV5Req) HasInstType() bool`

HasInstType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


