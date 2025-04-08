# CreateCopytradingAmendProfitSharingRatioV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstType** | Pointer to **string** | Instrument type  &#x60;SWAP&#x60; | [optional] [default to ""]
**ProfitSharingRatio** | **string** | Profit sharing ratio.   0.1 represents 10% | [default to ""]

## Methods

### NewCreateCopytradingAmendProfitSharingRatioV5Req

`func NewCreateCopytradingAmendProfitSharingRatioV5Req(profitSharingRatio string, ) *CreateCopytradingAmendProfitSharingRatioV5Req`

NewCreateCopytradingAmendProfitSharingRatioV5Req instantiates a new CreateCopytradingAmendProfitSharingRatioV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCopytradingAmendProfitSharingRatioV5ReqWithDefaults

`func NewCreateCopytradingAmendProfitSharingRatioV5ReqWithDefaults() *CreateCopytradingAmendProfitSharingRatioV5Req`

NewCreateCopytradingAmendProfitSharingRatioV5ReqWithDefaults instantiates a new CreateCopytradingAmendProfitSharingRatioV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstType

`func (o *CreateCopytradingAmendProfitSharingRatioV5Req) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *CreateCopytradingAmendProfitSharingRatioV5Req) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *CreateCopytradingAmendProfitSharingRatioV5Req) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *CreateCopytradingAmendProfitSharingRatioV5Req) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetProfitSharingRatio

`func (o *CreateCopytradingAmendProfitSharingRatioV5Req) GetProfitSharingRatio() string`

GetProfitSharingRatio returns the ProfitSharingRatio field if non-nil, zero value otherwise.

### GetProfitSharingRatioOk

`func (o *CreateCopytradingAmendProfitSharingRatioV5Req) GetProfitSharingRatioOk() (*string, bool)`

GetProfitSharingRatioOk returns a tuple with the ProfitSharingRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitSharingRatio

`func (o *CreateCopytradingAmendProfitSharingRatioV5Req) SetProfitSharingRatio(v string)`

SetProfitSharingRatio sets ProfitSharingRatio field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


