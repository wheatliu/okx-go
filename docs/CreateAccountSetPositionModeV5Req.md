# CreateAccountSetPositionModeV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PosMode** | **string** | Position mode  &#x60;long_short_mode&#x60;: long/short, only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;  &#x60;net_mode&#x60;: net | [default to ""]

## Methods

### NewCreateAccountSetPositionModeV5Req

`func NewCreateAccountSetPositionModeV5Req(posMode string, ) *CreateAccountSetPositionModeV5Req`

NewCreateAccountSetPositionModeV5Req instantiates a new CreateAccountSetPositionModeV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountSetPositionModeV5ReqWithDefaults

`func NewCreateAccountSetPositionModeV5ReqWithDefaults() *CreateAccountSetPositionModeV5Req`

NewCreateAccountSetPositionModeV5ReqWithDefaults instantiates a new CreateAccountSetPositionModeV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosMode

`func (o *CreateAccountSetPositionModeV5Req) GetPosMode() string`

GetPosMode returns the PosMode field if non-nil, zero value otherwise.

### GetPosModeOk

`func (o *CreateAccountSetPositionModeV5Req) GetPosModeOk() (*string, bool)`

GetPosModeOk returns a tuple with the PosMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosMode

`func (o *CreateAccountSetPositionModeV5Req) SetPosMode(v string)`

SetPosMode sets PosMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


