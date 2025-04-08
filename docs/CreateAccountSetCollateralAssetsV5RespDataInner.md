# CreateAccountSetCollateralAssetsV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CcyList** | Pointer to **[]string** | Currency list, e.g. [\&quot;BTC\&quot;,\&quot;ETH\&quot;] | [optional] 
**CollateralEnabled** | Pointer to **bool** | Whether or not set the assets to be collateral  &#x60;true&#x60;: Set to be collateral  &#x60;false&#x60;: Set to be non-collateral | [optional] 
**Type** | Pointer to **string** | Type  &#x60;all&#x60;  &#x60;custom&#x60; | [optional] [default to ""]

## Methods

### NewCreateAccountSetCollateralAssetsV5RespDataInner

`func NewCreateAccountSetCollateralAssetsV5RespDataInner() *CreateAccountSetCollateralAssetsV5RespDataInner`

NewCreateAccountSetCollateralAssetsV5RespDataInner instantiates a new CreateAccountSetCollateralAssetsV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountSetCollateralAssetsV5RespDataInnerWithDefaults

`func NewCreateAccountSetCollateralAssetsV5RespDataInnerWithDefaults() *CreateAccountSetCollateralAssetsV5RespDataInner`

NewCreateAccountSetCollateralAssetsV5RespDataInnerWithDefaults instantiates a new CreateAccountSetCollateralAssetsV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcyList

`func (o *CreateAccountSetCollateralAssetsV5RespDataInner) GetCcyList() []string`

GetCcyList returns the CcyList field if non-nil, zero value otherwise.

### GetCcyListOk

`func (o *CreateAccountSetCollateralAssetsV5RespDataInner) GetCcyListOk() (*[]string, bool)`

GetCcyListOk returns a tuple with the CcyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcyList

`func (o *CreateAccountSetCollateralAssetsV5RespDataInner) SetCcyList(v []string)`

SetCcyList sets CcyList field to given value.

### HasCcyList

`func (o *CreateAccountSetCollateralAssetsV5RespDataInner) HasCcyList() bool`

HasCcyList returns a boolean if a field has been set.

### GetCollateralEnabled

`func (o *CreateAccountSetCollateralAssetsV5RespDataInner) GetCollateralEnabled() bool`

GetCollateralEnabled returns the CollateralEnabled field if non-nil, zero value otherwise.

### GetCollateralEnabledOk

`func (o *CreateAccountSetCollateralAssetsV5RespDataInner) GetCollateralEnabledOk() (*bool, bool)`

GetCollateralEnabledOk returns a tuple with the CollateralEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralEnabled

`func (o *CreateAccountSetCollateralAssetsV5RespDataInner) SetCollateralEnabled(v bool)`

SetCollateralEnabled sets CollateralEnabled field to given value.

### HasCollateralEnabled

`func (o *CreateAccountSetCollateralAssetsV5RespDataInner) HasCollateralEnabled() bool`

HasCollateralEnabled returns a boolean if a field has been set.

### GetType

`func (o *CreateAccountSetCollateralAssetsV5RespDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateAccountSetCollateralAssetsV5RespDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateAccountSetCollateralAssetsV5RespDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateAccountSetCollateralAssetsV5RespDataInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


