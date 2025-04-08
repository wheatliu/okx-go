# CreateAccountSetIsolatedModeV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsoMode** | **string** | Isolated margin trading settings   &#x60;auto_transfers_ccy&#x60;: New auto transfers, enabling both base and quote currency as the margin for isolated margin trading. Only applicable to &#x60;MARGIN&#x60;.  &#x60;automatic&#x60;: Auto transfers | [default to ""]
**Type** | **string** | Instrument type  &#x60;MARGIN&#x60;  &#x60;CONTRACTS&#x60; | [default to ""]

## Methods

### NewCreateAccountSetIsolatedModeV5Req

`func NewCreateAccountSetIsolatedModeV5Req(isoMode string, type_ string, ) *CreateAccountSetIsolatedModeV5Req`

NewCreateAccountSetIsolatedModeV5Req instantiates a new CreateAccountSetIsolatedModeV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountSetIsolatedModeV5ReqWithDefaults

`func NewCreateAccountSetIsolatedModeV5ReqWithDefaults() *CreateAccountSetIsolatedModeV5Req`

NewCreateAccountSetIsolatedModeV5ReqWithDefaults instantiates a new CreateAccountSetIsolatedModeV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsoMode

`func (o *CreateAccountSetIsolatedModeV5Req) GetIsoMode() string`

GetIsoMode returns the IsoMode field if non-nil, zero value otherwise.

### GetIsoModeOk

`func (o *CreateAccountSetIsolatedModeV5Req) GetIsoModeOk() (*string, bool)`

GetIsoModeOk returns a tuple with the IsoMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoMode

`func (o *CreateAccountSetIsolatedModeV5Req) SetIsoMode(v string)`

SetIsoMode sets IsoMode field to given value.


### GetType

`func (o *CreateAccountSetIsolatedModeV5Req) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateAccountSetIsolatedModeV5Req) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateAccountSetIsolatedModeV5Req) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


