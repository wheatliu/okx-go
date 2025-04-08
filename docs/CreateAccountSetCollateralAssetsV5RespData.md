# CreateAccountSetCollateralAssetsV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CcyList** | Pointer to **[]string** | Currency list, e.g. [\&quot;BTC\&quot;,\&quot;ETH\&quot;] | [optional] 
**CollateralEnabled** | Pointer to **bool** | Whether or not set the assets to be collateral  &#x60;true&#x60;: Set to be collateral  &#x60;false&#x60;: Set to be non-collateral | [optional] 
**Type** | Pointer to **string** | Type  &#x60;all&#x60;  &#x60;custom&#x60; | [optional] [default to ""]

## Methods

### NewCreateAccountSetCollateralAssetsV5RespData

`func NewCreateAccountSetCollateralAssetsV5RespData() *CreateAccountSetCollateralAssetsV5RespData`

NewCreateAccountSetCollateralAssetsV5RespData instantiates a new CreateAccountSetCollateralAssetsV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountSetCollateralAssetsV5RespDataWithDefaults

`func NewCreateAccountSetCollateralAssetsV5RespDataWithDefaults() *CreateAccountSetCollateralAssetsV5RespData`

NewCreateAccountSetCollateralAssetsV5RespDataWithDefaults instantiates a new CreateAccountSetCollateralAssetsV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcyList

`func (o *CreateAccountSetCollateralAssetsV5RespData) GetCcyList() []string`

GetCcyList returns the CcyList field if non-nil, zero value otherwise.

### GetCcyListOk

`func (o *CreateAccountSetCollateralAssetsV5RespData) GetCcyListOk() (*[]string, bool)`

GetCcyListOk returns a tuple with the CcyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcyList

`func (o *CreateAccountSetCollateralAssetsV5RespData) SetCcyList(v []string)`

SetCcyList sets CcyList field to given value.

### HasCcyList

`func (o *CreateAccountSetCollateralAssetsV5RespData) HasCcyList() bool`

HasCcyList returns a boolean if a field has been set.

### GetCollateralEnabled

`func (o *CreateAccountSetCollateralAssetsV5RespData) GetCollateralEnabled() bool`

GetCollateralEnabled returns the CollateralEnabled field if non-nil, zero value otherwise.

### GetCollateralEnabledOk

`func (o *CreateAccountSetCollateralAssetsV5RespData) GetCollateralEnabledOk() (*bool, bool)`

GetCollateralEnabledOk returns a tuple with the CollateralEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralEnabled

`func (o *CreateAccountSetCollateralAssetsV5RespData) SetCollateralEnabled(v bool)`

SetCollateralEnabled sets CollateralEnabled field to given value.

### HasCollateralEnabled

`func (o *CreateAccountSetCollateralAssetsV5RespData) HasCollateralEnabled() bool`

HasCollateralEnabled returns a boolean if a field has been set.

### GetType

`func (o *CreateAccountSetCollateralAssetsV5RespData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateAccountSetCollateralAssetsV5RespData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateAccountSetCollateralAssetsV5RespData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateAccountSetCollateralAssetsV5RespData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


