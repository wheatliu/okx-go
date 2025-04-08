# CreateAccountSetAccountLevelV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctLv** | **string** | Account mode  &#x60;1&#x60;: Spot mode  &#x60;2&#x60;: Spot and futures mode   &#x60;3&#x60;: Multi-currency margin code   &#x60;4&#x60;: Portfolio margin mode | [default to ""]

## Methods

### NewCreateAccountSetAccountLevelV5Req

`func NewCreateAccountSetAccountLevelV5Req(acctLv string, ) *CreateAccountSetAccountLevelV5Req`

NewCreateAccountSetAccountLevelV5Req instantiates a new CreateAccountSetAccountLevelV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountSetAccountLevelV5ReqWithDefaults

`func NewCreateAccountSetAccountLevelV5ReqWithDefaults() *CreateAccountSetAccountLevelV5Req`

NewCreateAccountSetAccountLevelV5ReqWithDefaults instantiates a new CreateAccountSetAccountLevelV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctLv

`func (o *CreateAccountSetAccountLevelV5Req) GetAcctLv() string`

GetAcctLv returns the AcctLv field if non-nil, zero value otherwise.

### GetAcctLvOk

`func (o *CreateAccountSetAccountLevelV5Req) GetAcctLvOk() (*string, bool)`

GetAcctLvOk returns a tuple with the AcctLv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctLv

`func (o *CreateAccountSetAccountLevelV5Req) SetAcctLv(v string)`

SetAcctLv sets AcctLv field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


