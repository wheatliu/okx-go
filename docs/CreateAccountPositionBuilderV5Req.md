# CreateAccountPositionBuilderV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctLv** | Pointer to **string** | Switch to account mode  &#x60;3&#x60;: Multi-currency margin  &#x60;4&#x60;: Portfolio margin  The default is &#x60;4&#x60; | [optional] [default to ""]
**GreeksType** | Pointer to **string** | Greeks type  &#x60;BS&#x60;: Black-Scholes Model Greeks  &#x60;PA&#x60;: Crypto Greeks  &#x60;CASH&#x60;: Empirical Greeks  The default is &#x60;BS&#x60; | [optional] [default to ""]
**InclRealPosAndEq** | Pointer to **bool** | Whether import existing positions and assets  The default is &#x60;true&#x60; | [optional] 
**Lever** | Pointer to **string** | Cross margin leverage in Multi-currency margin mode, the default is &#x60;1&#x60;.  If the allowed leverage is exceeded, set according to the maximum leverage.  Only applicable to &#x60;Multi-currency margin&#x60; | [optional] [default to ""]
**SimAsset** | Pointer to [**[]CreateAccountPositionBuilderV5ReqSimAssetInner**](CreateAccountPositionBuilderV5ReqSimAssetInner.md) | List of simulated assets  When &#x60;inclRealPosAndEq&#x60; is &#x60;true&#x60;, only real assets are considered and virtual assets are ignored | [optional] 
**SimPos** | Pointer to [**[]CreateAccountPositionBuilderV5ReqSimPosInner**](CreateAccountPositionBuilderV5ReqSimPosInner.md) | List of simulated positions | [optional] 

## Methods

### NewCreateAccountPositionBuilderV5Req

`func NewCreateAccountPositionBuilderV5Req() *CreateAccountPositionBuilderV5Req`

NewCreateAccountPositionBuilderV5Req instantiates a new CreateAccountPositionBuilderV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountPositionBuilderV5ReqWithDefaults

`func NewCreateAccountPositionBuilderV5ReqWithDefaults() *CreateAccountPositionBuilderV5Req`

NewCreateAccountPositionBuilderV5ReqWithDefaults instantiates a new CreateAccountPositionBuilderV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctLv

`func (o *CreateAccountPositionBuilderV5Req) GetAcctLv() string`

GetAcctLv returns the AcctLv field if non-nil, zero value otherwise.

### GetAcctLvOk

`func (o *CreateAccountPositionBuilderV5Req) GetAcctLvOk() (*string, bool)`

GetAcctLvOk returns a tuple with the AcctLv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctLv

`func (o *CreateAccountPositionBuilderV5Req) SetAcctLv(v string)`

SetAcctLv sets AcctLv field to given value.

### HasAcctLv

`func (o *CreateAccountPositionBuilderV5Req) HasAcctLv() bool`

HasAcctLv returns a boolean if a field has been set.

### GetGreeksType

`func (o *CreateAccountPositionBuilderV5Req) GetGreeksType() string`

GetGreeksType returns the GreeksType field if non-nil, zero value otherwise.

### GetGreeksTypeOk

`func (o *CreateAccountPositionBuilderV5Req) GetGreeksTypeOk() (*string, bool)`

GetGreeksTypeOk returns a tuple with the GreeksType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreeksType

`func (o *CreateAccountPositionBuilderV5Req) SetGreeksType(v string)`

SetGreeksType sets GreeksType field to given value.

### HasGreeksType

`func (o *CreateAccountPositionBuilderV5Req) HasGreeksType() bool`

HasGreeksType returns a boolean if a field has been set.

### GetInclRealPosAndEq

`func (o *CreateAccountPositionBuilderV5Req) GetInclRealPosAndEq() bool`

GetInclRealPosAndEq returns the InclRealPosAndEq field if non-nil, zero value otherwise.

### GetInclRealPosAndEqOk

`func (o *CreateAccountPositionBuilderV5Req) GetInclRealPosAndEqOk() (*bool, bool)`

GetInclRealPosAndEqOk returns a tuple with the InclRealPosAndEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclRealPosAndEq

`func (o *CreateAccountPositionBuilderV5Req) SetInclRealPosAndEq(v bool)`

SetInclRealPosAndEq sets InclRealPosAndEq field to given value.

### HasInclRealPosAndEq

`func (o *CreateAccountPositionBuilderV5Req) HasInclRealPosAndEq() bool`

HasInclRealPosAndEq returns a boolean if a field has been set.

### GetLever

`func (o *CreateAccountPositionBuilderV5Req) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *CreateAccountPositionBuilderV5Req) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *CreateAccountPositionBuilderV5Req) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *CreateAccountPositionBuilderV5Req) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetSimAsset

`func (o *CreateAccountPositionBuilderV5Req) GetSimAsset() []CreateAccountPositionBuilderV5ReqSimAssetInner`

GetSimAsset returns the SimAsset field if non-nil, zero value otherwise.

### GetSimAssetOk

`func (o *CreateAccountPositionBuilderV5Req) GetSimAssetOk() (*[]CreateAccountPositionBuilderV5ReqSimAssetInner, bool)`

GetSimAssetOk returns a tuple with the SimAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimAsset

`func (o *CreateAccountPositionBuilderV5Req) SetSimAsset(v []CreateAccountPositionBuilderV5ReqSimAssetInner)`

SetSimAsset sets SimAsset field to given value.

### HasSimAsset

`func (o *CreateAccountPositionBuilderV5Req) HasSimAsset() bool`

HasSimAsset returns a boolean if a field has been set.

### GetSimPos

`func (o *CreateAccountPositionBuilderV5Req) GetSimPos() []CreateAccountPositionBuilderV5ReqSimPosInner`

GetSimPos returns the SimPos field if non-nil, zero value otherwise.

### GetSimPosOk

`func (o *CreateAccountPositionBuilderV5Req) GetSimPosOk() (*[]CreateAccountPositionBuilderV5ReqSimPosInner, bool)`

GetSimPosOk returns a tuple with the SimPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimPos

`func (o *CreateAccountPositionBuilderV5Req) SetSimPos(v []CreateAccountPositionBuilderV5ReqSimPosInner)`

SetSimPos sets SimPos field to given value.

### HasSimPos

`func (o *CreateAccountPositionBuilderV5Req) HasSimPos() bool`

HasSimPos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


