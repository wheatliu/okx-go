# GetAccountCollateralAssetsV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | Pointer to **string** | Currency, e.g. &#x60;BTC&#x60; | [optional] [default to ""]
**CollateralEnabled** | Pointer to **bool** | Whether or not to be a collateral asset | [optional] 

## Methods

### NewGetAccountCollateralAssetsV5RespDataInner

`func NewGetAccountCollateralAssetsV5RespDataInner() *GetAccountCollateralAssetsV5RespDataInner`

NewGetAccountCollateralAssetsV5RespDataInner instantiates a new GetAccountCollateralAssetsV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountCollateralAssetsV5RespDataInnerWithDefaults

`func NewGetAccountCollateralAssetsV5RespDataInnerWithDefaults() *GetAccountCollateralAssetsV5RespDataInner`

NewGetAccountCollateralAssetsV5RespDataInnerWithDefaults instantiates a new GetAccountCollateralAssetsV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *GetAccountCollateralAssetsV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAccountCollateralAssetsV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAccountCollateralAssetsV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAccountCollateralAssetsV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetCollateralEnabled

`func (o *GetAccountCollateralAssetsV5RespDataInner) GetCollateralEnabled() bool`

GetCollateralEnabled returns the CollateralEnabled field if non-nil, zero value otherwise.

### GetCollateralEnabledOk

`func (o *GetAccountCollateralAssetsV5RespDataInner) GetCollateralEnabledOk() (*bool, bool)`

GetCollateralEnabledOk returns a tuple with the CollateralEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralEnabled

`func (o *GetAccountCollateralAssetsV5RespDataInner) SetCollateralEnabled(v bool)`

SetCollateralEnabled sets CollateralEnabled field to given value.

### HasCollateralEnabled

`func (o *GetAccountCollateralAssetsV5RespDataInner) HasCollateralEnabled() bool`

HasCollateralEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


