# GetAssetNonTradableAssetsV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bal** | Pointer to **string** | Withdrawable balance | [optional] [default to ""]
**BurningFeeRate** | Pointer to **string** | Burning fee rate, e.g \&quot;0.05\&quot; represents \&quot;5%\&quot;.  Some currencies may charge combustion fees. The burning fee is deducted based on the withdrawal quantity (excluding gas fee) multiplied by the burning fee rate. | [optional] [default to ""]
**CanWd** | Pointer to **bool** | Availability to withdraw to chain.   &#x60;false&#x60;: not available  &#x60;true&#x60;: available | [optional] 
**Ccy** | Pointer to **string** | Currency, e.g. &#x60;CELT&#x60; | [optional] [default to ""]
**Chain** | Pointer to **string** | Chain for withdrawal | [optional] [default to ""]
**CtAddr** | Pointer to **string** | Last 6 digits of contract address | [optional] [default to ""]
**Fee** | Pointer to **string** | Fixed withdrawal fee | [optional] [default to ""]
**FeeCcy** | Pointer to **string** | Fixed withdrawal fee unit, e.g. &#x60;USDT&#x60; | [optional] [default to ""]
**LogoLink** | Pointer to **string** | Logo link of currency | [optional] [default to ""]
**MinWd** | Pointer to **string** | Minimum withdrawal amount of currency in a single transaction | [optional] [default to ""]
**Name** | Pointer to **string** | Chinese name of currency. There is no related name when it is not shown. | [optional] [default to ""]
**NeedTag** | Pointer to **bool** | Whether tag/memo information is required for withdrawal | [optional] 
**WdAll** | Pointer to **bool** | Whether all assets in this currency must be withdrawn at one time | [optional] 
**WdTickSz** | Pointer to **string** | Withdrawal precision, indicating the number of digits after the decimal point | [optional] [default to ""]

## Methods

### NewGetAssetNonTradableAssetsV5RespDataInner

`func NewGetAssetNonTradableAssetsV5RespDataInner() *GetAssetNonTradableAssetsV5RespDataInner`

NewGetAssetNonTradableAssetsV5RespDataInner instantiates a new GetAssetNonTradableAssetsV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetNonTradableAssetsV5RespDataInnerWithDefaults

`func NewGetAssetNonTradableAssetsV5RespDataInnerWithDefaults() *GetAssetNonTradableAssetsV5RespDataInner`

NewGetAssetNonTradableAssetsV5RespDataInnerWithDefaults instantiates a new GetAssetNonTradableAssetsV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBal

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetBal() string`

GetBal returns the Bal field if non-nil, zero value otherwise.

### GetBalOk

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetBalOk() (*string, bool)`

GetBalOk returns a tuple with the Bal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBal

`func (o *GetAssetNonTradableAssetsV5RespDataInner) SetBal(v string)`

SetBal sets Bal field to given value.

### HasBal

`func (o *GetAssetNonTradableAssetsV5RespDataInner) HasBal() bool`

HasBal returns a boolean if a field has been set.

### GetBurningFeeRate

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetBurningFeeRate() string`

GetBurningFeeRate returns the BurningFeeRate field if non-nil, zero value otherwise.

### GetBurningFeeRateOk

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetBurningFeeRateOk() (*string, bool)`

GetBurningFeeRateOk returns a tuple with the BurningFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurningFeeRate

`func (o *GetAssetNonTradableAssetsV5RespDataInner) SetBurningFeeRate(v string)`

SetBurningFeeRate sets BurningFeeRate field to given value.

### HasBurningFeeRate

`func (o *GetAssetNonTradableAssetsV5RespDataInner) HasBurningFeeRate() bool`

HasBurningFeeRate returns a boolean if a field has been set.

### GetCanWd

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetCanWd() bool`

GetCanWd returns the CanWd field if non-nil, zero value otherwise.

### GetCanWdOk

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetCanWdOk() (*bool, bool)`

GetCanWdOk returns a tuple with the CanWd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWd

`func (o *GetAssetNonTradableAssetsV5RespDataInner) SetCanWd(v bool)`

SetCanWd sets CanWd field to given value.

### HasCanWd

`func (o *GetAssetNonTradableAssetsV5RespDataInner) HasCanWd() bool`

HasCanWd returns a boolean if a field has been set.

### GetCcy

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAssetNonTradableAssetsV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAssetNonTradableAssetsV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetChain

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *GetAssetNonTradableAssetsV5RespDataInner) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *GetAssetNonTradableAssetsV5RespDataInner) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetCtAddr

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetCtAddr() string`

GetCtAddr returns the CtAddr field if non-nil, zero value otherwise.

### GetCtAddrOk

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetCtAddrOk() (*string, bool)`

GetCtAddrOk returns a tuple with the CtAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtAddr

`func (o *GetAssetNonTradableAssetsV5RespDataInner) SetCtAddr(v string)`

SetCtAddr sets CtAddr field to given value.

### HasCtAddr

`func (o *GetAssetNonTradableAssetsV5RespDataInner) HasCtAddr() bool`

HasCtAddr returns a boolean if a field has been set.

### GetFee

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetAssetNonTradableAssetsV5RespDataInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetAssetNonTradableAssetsV5RespDataInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFeeCcy

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetFeeCcy() string`

GetFeeCcy returns the FeeCcy field if non-nil, zero value otherwise.

### GetFeeCcyOk

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetFeeCcyOk() (*string, bool)`

GetFeeCcyOk returns a tuple with the FeeCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCcy

`func (o *GetAssetNonTradableAssetsV5RespDataInner) SetFeeCcy(v string)`

SetFeeCcy sets FeeCcy field to given value.

### HasFeeCcy

`func (o *GetAssetNonTradableAssetsV5RespDataInner) HasFeeCcy() bool`

HasFeeCcy returns a boolean if a field has been set.

### GetLogoLink

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetLogoLink() string`

GetLogoLink returns the LogoLink field if non-nil, zero value otherwise.

### GetLogoLinkOk

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetLogoLinkOk() (*string, bool)`

GetLogoLinkOk returns a tuple with the LogoLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoLink

`func (o *GetAssetNonTradableAssetsV5RespDataInner) SetLogoLink(v string)`

SetLogoLink sets LogoLink field to given value.

### HasLogoLink

`func (o *GetAssetNonTradableAssetsV5RespDataInner) HasLogoLink() bool`

HasLogoLink returns a boolean if a field has been set.

### GetMinWd

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetMinWd() string`

GetMinWd returns the MinWd field if non-nil, zero value otherwise.

### GetMinWdOk

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetMinWdOk() (*string, bool)`

GetMinWdOk returns a tuple with the MinWd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWd

`func (o *GetAssetNonTradableAssetsV5RespDataInner) SetMinWd(v string)`

SetMinWd sets MinWd field to given value.

### HasMinWd

`func (o *GetAssetNonTradableAssetsV5RespDataInner) HasMinWd() bool`

HasMinWd returns a boolean if a field has been set.

### GetName

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAssetNonTradableAssetsV5RespDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAssetNonTradableAssetsV5RespDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedTag

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetNeedTag() bool`

GetNeedTag returns the NeedTag field if non-nil, zero value otherwise.

### GetNeedTagOk

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetNeedTagOk() (*bool, bool)`

GetNeedTagOk returns a tuple with the NeedTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedTag

`func (o *GetAssetNonTradableAssetsV5RespDataInner) SetNeedTag(v bool)`

SetNeedTag sets NeedTag field to given value.

### HasNeedTag

`func (o *GetAssetNonTradableAssetsV5RespDataInner) HasNeedTag() bool`

HasNeedTag returns a boolean if a field has been set.

### GetWdAll

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetWdAll() bool`

GetWdAll returns the WdAll field if non-nil, zero value otherwise.

### GetWdAllOk

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetWdAllOk() (*bool, bool)`

GetWdAllOk returns a tuple with the WdAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWdAll

`func (o *GetAssetNonTradableAssetsV5RespDataInner) SetWdAll(v bool)`

SetWdAll sets WdAll field to given value.

### HasWdAll

`func (o *GetAssetNonTradableAssetsV5RespDataInner) HasWdAll() bool`

HasWdAll returns a boolean if a field has been set.

### GetWdTickSz

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetWdTickSz() string`

GetWdTickSz returns the WdTickSz field if non-nil, zero value otherwise.

### GetWdTickSzOk

`func (o *GetAssetNonTradableAssetsV5RespDataInner) GetWdTickSzOk() (*string, bool)`

GetWdTickSzOk returns a tuple with the WdTickSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWdTickSz

`func (o *GetAssetNonTradableAssetsV5RespDataInner) SetWdTickSz(v string)`

SetWdTickSz sets WdTickSz field to given value.

### HasWdTickSz

`func (o *GetAssetNonTradableAssetsV5RespDataInner) HasWdTickSz() bool`

HasWdTickSz returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


