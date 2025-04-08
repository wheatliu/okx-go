# GetAssetCurrenciesV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BurningFeeRate** | Pointer to **string** | Burning fee rate, e.g \&quot;0.05\&quot; represents \&quot;5%\&quot;.  Some currencies may charge combustion fees. The burning fee is deducted based on the withdrawal quantity (excluding gas fee) multiplied by the burning fee rate.  Apply to &#x60;on-chain withdrawal&#x60; | [optional] [default to ""]
**CanDep** | Pointer to **bool** | The availability to deposit from chain   &#x60;false&#x60;: not available   &#x60;true&#x60;: available | [optional] 
**CanInternal** | Pointer to **bool** | The availability to internal transfer   &#x60;false&#x60;: not available   &#x60;true&#x60;: available | [optional] 
**CanWd** | Pointer to **bool** | The availability to withdraw to chain   &#x60;false&#x60;: not available   &#x60;true&#x60;: available | [optional] 
**Ccy** | Pointer to **string** | Currency, e.g. &#x60;BTC&#x60; | [optional] [default to ""]
**Chain** | Pointer to **string** | Chain name, e.g. &#x60;USDT-ERC20&#x60;, &#x60;USDT-TRC20&#x60; | [optional] [default to ""]
**CtAddr** | Pointer to **string** | Contract address | [optional] [default to ""]
**DepEstOpenTime** | Pointer to **string** | Estimated opening time for deposit, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**DepQuotaFixed** | Pointer to **string** | The fixed deposit limit, unit in &#x60;USD&#x60;  Return empty string if there is no deposit limit | [optional] [default to ""]
**DepQuoteDailyLayer2** | Pointer to **string** | The layer2 network daily deposit limit | [optional] [default to ""]
**Fee** | Pointer to **string** | The fixed withdrawal fee  Apply to &#x60;on-chain withdrawal&#x60; | [optional] [default to ""]
**LogoLink** | Pointer to **string** | The logo link of currency | [optional] [default to ""]
**MainNet** | Pointer to **bool** | If current chain is main net, then it will return &#x60;true&#x60;, otherwise it will return &#x60;false&#x60; | [optional] 
**MaxFee** | Pointer to **string** | The maximum withdrawal fee for normal address  Apply to &#x60;on-chain withdrawal&#x60;  (Deprecated) | [optional] [default to ""]
**MaxFeeForCtAddr** | Pointer to **string** | The maximum withdrawal fee for contract address  Apply to &#x60;on-chain withdrawal&#x60;  (Deprecated) | [optional] [default to ""]
**MaxWd** | Pointer to **string** | The maximum amount of currency &#x60;on-chain withdrawal&#x60; in a single transaction | [optional] [default to ""]
**MinDep** | Pointer to **string** | The minimum deposit amount of currency in a single transaction | [optional] [default to ""]
**MinDepArrivalConfirm** | Pointer to **string** | The minimum number of blockchain confirmations to acknowledge fund deposit. The account is credited after that, but the deposit can not be withdrawn | [optional] [default to ""]
**MinFee** | Pointer to **string** | The minimum withdrawal fee for normal address  Apply to &#x60;on-chain withdrawal&#x60;  (Deprecated) | [optional] [default to ""]
**MinFeeForCtAddr** | Pointer to **string** | The minimum withdrawal fee for contract address  Apply to &#x60;on-chain withdrawal&#x60;  (Deprecated) | [optional] [default to ""]
**MinInternal** | Pointer to **string** | The minimum &#x60;internal transfer&#x60; amount of currency in a single transaction  No maximum &#x60;internal transfer&#x60; limit in a single transaction, subject to the withdrawal limit in the past 24 hours(&#x60;wdQuota&#x60;). | [optional] [default to ""]
**MinWd** | Pointer to **string** | The minimum &#x60;on-chain withdrawal&#x60; amount of currency in a single transaction | [optional] [default to ""]
**MinWdUnlockConfirm** | Pointer to **string** | The minimum number of blockchain confirmations required for withdrawal of a deposit | [optional] [default to ""]
**Name** | Pointer to **string** | Name of currency. There is no related name when it is not shown. | [optional] [default to ""]
**NeedTag** | Pointer to **bool** | Whether tag/memo information is required for withdrawal, e.g. &#x60;EOS&#x60; will return &#x60;true&#x60; | [optional] 
**UsedDepQuotaFixed** | Pointer to **string** | The used amount of fixed deposit quota, unit in &#x60;USD&#x60;  Return empty string if there is no deposit limit | [optional] [default to ""]
**UsedWdQuota** | Pointer to **string** | The amount of currency withdrawal used in the past 24 hours, unit in &#x60;USD&#x60; | [optional] [default to ""]
**WdEstOpenTime** | Pointer to **string** | Estimated opening time for withdraw, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**WdQuota** | Pointer to **string** | The withdrawal limit in the past 24 hours (including &#x60;on-chain withdrawal&#x60; and &#x60;internal transfer&#x60;), unit in &#x60;USD&#x60; | [optional] [default to ""]
**WdTickSz** | Pointer to **string** | The withdrawal precision, indicating the number of digits after the decimal point.  The withdrawal fee precision kept the same as withdrawal precision.  The accuracy of internal transfer withdrawal is 8 decimal places. | [optional] [default to ""]

## Methods

### NewGetAssetCurrenciesV5RespData

`func NewGetAssetCurrenciesV5RespData() *GetAssetCurrenciesV5RespData`

NewGetAssetCurrenciesV5RespData instantiates a new GetAssetCurrenciesV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetCurrenciesV5RespDataWithDefaults

`func NewGetAssetCurrenciesV5RespDataWithDefaults() *GetAssetCurrenciesV5RespData`

NewGetAssetCurrenciesV5RespDataWithDefaults instantiates a new GetAssetCurrenciesV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBurningFeeRate

`func (o *GetAssetCurrenciesV5RespData) GetBurningFeeRate() string`

GetBurningFeeRate returns the BurningFeeRate field if non-nil, zero value otherwise.

### GetBurningFeeRateOk

`func (o *GetAssetCurrenciesV5RespData) GetBurningFeeRateOk() (*string, bool)`

GetBurningFeeRateOk returns a tuple with the BurningFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurningFeeRate

`func (o *GetAssetCurrenciesV5RespData) SetBurningFeeRate(v string)`

SetBurningFeeRate sets BurningFeeRate field to given value.

### HasBurningFeeRate

`func (o *GetAssetCurrenciesV5RespData) HasBurningFeeRate() bool`

HasBurningFeeRate returns a boolean if a field has been set.

### GetCanDep

`func (o *GetAssetCurrenciesV5RespData) GetCanDep() bool`

GetCanDep returns the CanDep field if non-nil, zero value otherwise.

### GetCanDepOk

`func (o *GetAssetCurrenciesV5RespData) GetCanDepOk() (*bool, bool)`

GetCanDepOk returns a tuple with the CanDep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDep

`func (o *GetAssetCurrenciesV5RespData) SetCanDep(v bool)`

SetCanDep sets CanDep field to given value.

### HasCanDep

`func (o *GetAssetCurrenciesV5RespData) HasCanDep() bool`

HasCanDep returns a boolean if a field has been set.

### GetCanInternal

`func (o *GetAssetCurrenciesV5RespData) GetCanInternal() bool`

GetCanInternal returns the CanInternal field if non-nil, zero value otherwise.

### GetCanInternalOk

`func (o *GetAssetCurrenciesV5RespData) GetCanInternalOk() (*bool, bool)`

GetCanInternalOk returns a tuple with the CanInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanInternal

`func (o *GetAssetCurrenciesV5RespData) SetCanInternal(v bool)`

SetCanInternal sets CanInternal field to given value.

### HasCanInternal

`func (o *GetAssetCurrenciesV5RespData) HasCanInternal() bool`

HasCanInternal returns a boolean if a field has been set.

### GetCanWd

`func (o *GetAssetCurrenciesV5RespData) GetCanWd() bool`

GetCanWd returns the CanWd field if non-nil, zero value otherwise.

### GetCanWdOk

`func (o *GetAssetCurrenciesV5RespData) GetCanWdOk() (*bool, bool)`

GetCanWdOk returns a tuple with the CanWd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWd

`func (o *GetAssetCurrenciesV5RespData) SetCanWd(v bool)`

SetCanWd sets CanWd field to given value.

### HasCanWd

`func (o *GetAssetCurrenciesV5RespData) HasCanWd() bool`

HasCanWd returns a boolean if a field has been set.

### GetCcy

`func (o *GetAssetCurrenciesV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAssetCurrenciesV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAssetCurrenciesV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAssetCurrenciesV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetChain

`func (o *GetAssetCurrenciesV5RespData) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *GetAssetCurrenciesV5RespData) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *GetAssetCurrenciesV5RespData) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *GetAssetCurrenciesV5RespData) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetCtAddr

`func (o *GetAssetCurrenciesV5RespData) GetCtAddr() string`

GetCtAddr returns the CtAddr field if non-nil, zero value otherwise.

### GetCtAddrOk

`func (o *GetAssetCurrenciesV5RespData) GetCtAddrOk() (*string, bool)`

GetCtAddrOk returns a tuple with the CtAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtAddr

`func (o *GetAssetCurrenciesV5RespData) SetCtAddr(v string)`

SetCtAddr sets CtAddr field to given value.

### HasCtAddr

`func (o *GetAssetCurrenciesV5RespData) HasCtAddr() bool`

HasCtAddr returns a boolean if a field has been set.

### GetDepEstOpenTime

`func (o *GetAssetCurrenciesV5RespData) GetDepEstOpenTime() string`

GetDepEstOpenTime returns the DepEstOpenTime field if non-nil, zero value otherwise.

### GetDepEstOpenTimeOk

`func (o *GetAssetCurrenciesV5RespData) GetDepEstOpenTimeOk() (*string, bool)`

GetDepEstOpenTimeOk returns a tuple with the DepEstOpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepEstOpenTime

`func (o *GetAssetCurrenciesV5RespData) SetDepEstOpenTime(v string)`

SetDepEstOpenTime sets DepEstOpenTime field to given value.

### HasDepEstOpenTime

`func (o *GetAssetCurrenciesV5RespData) HasDepEstOpenTime() bool`

HasDepEstOpenTime returns a boolean if a field has been set.

### GetDepQuotaFixed

`func (o *GetAssetCurrenciesV5RespData) GetDepQuotaFixed() string`

GetDepQuotaFixed returns the DepQuotaFixed field if non-nil, zero value otherwise.

### GetDepQuotaFixedOk

`func (o *GetAssetCurrenciesV5RespData) GetDepQuotaFixedOk() (*string, bool)`

GetDepQuotaFixedOk returns a tuple with the DepQuotaFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepQuotaFixed

`func (o *GetAssetCurrenciesV5RespData) SetDepQuotaFixed(v string)`

SetDepQuotaFixed sets DepQuotaFixed field to given value.

### HasDepQuotaFixed

`func (o *GetAssetCurrenciesV5RespData) HasDepQuotaFixed() bool`

HasDepQuotaFixed returns a boolean if a field has been set.

### GetDepQuoteDailyLayer2

`func (o *GetAssetCurrenciesV5RespData) GetDepQuoteDailyLayer2() string`

GetDepQuoteDailyLayer2 returns the DepQuoteDailyLayer2 field if non-nil, zero value otherwise.

### GetDepQuoteDailyLayer2Ok

`func (o *GetAssetCurrenciesV5RespData) GetDepQuoteDailyLayer2Ok() (*string, bool)`

GetDepQuoteDailyLayer2Ok returns a tuple with the DepQuoteDailyLayer2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepQuoteDailyLayer2

`func (o *GetAssetCurrenciesV5RespData) SetDepQuoteDailyLayer2(v string)`

SetDepQuoteDailyLayer2 sets DepQuoteDailyLayer2 field to given value.

### HasDepQuoteDailyLayer2

`func (o *GetAssetCurrenciesV5RespData) HasDepQuoteDailyLayer2() bool`

HasDepQuoteDailyLayer2 returns a boolean if a field has been set.

### GetFee

`func (o *GetAssetCurrenciesV5RespData) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetAssetCurrenciesV5RespData) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetAssetCurrenciesV5RespData) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetAssetCurrenciesV5RespData) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetLogoLink

`func (o *GetAssetCurrenciesV5RespData) GetLogoLink() string`

GetLogoLink returns the LogoLink field if non-nil, zero value otherwise.

### GetLogoLinkOk

`func (o *GetAssetCurrenciesV5RespData) GetLogoLinkOk() (*string, bool)`

GetLogoLinkOk returns a tuple with the LogoLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoLink

`func (o *GetAssetCurrenciesV5RespData) SetLogoLink(v string)`

SetLogoLink sets LogoLink field to given value.

### HasLogoLink

`func (o *GetAssetCurrenciesV5RespData) HasLogoLink() bool`

HasLogoLink returns a boolean if a field has been set.

### GetMainNet

`func (o *GetAssetCurrenciesV5RespData) GetMainNet() bool`

GetMainNet returns the MainNet field if non-nil, zero value otherwise.

### GetMainNetOk

`func (o *GetAssetCurrenciesV5RespData) GetMainNetOk() (*bool, bool)`

GetMainNetOk returns a tuple with the MainNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainNet

`func (o *GetAssetCurrenciesV5RespData) SetMainNet(v bool)`

SetMainNet sets MainNet field to given value.

### HasMainNet

`func (o *GetAssetCurrenciesV5RespData) HasMainNet() bool`

HasMainNet returns a boolean if a field has been set.

### GetMaxFee

`func (o *GetAssetCurrenciesV5RespData) GetMaxFee() string`

GetMaxFee returns the MaxFee field if non-nil, zero value otherwise.

### GetMaxFeeOk

`func (o *GetAssetCurrenciesV5RespData) GetMaxFeeOk() (*string, bool)`

GetMaxFeeOk returns a tuple with the MaxFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFee

`func (o *GetAssetCurrenciesV5RespData) SetMaxFee(v string)`

SetMaxFee sets MaxFee field to given value.

### HasMaxFee

`func (o *GetAssetCurrenciesV5RespData) HasMaxFee() bool`

HasMaxFee returns a boolean if a field has been set.

### GetMaxFeeForCtAddr

`func (o *GetAssetCurrenciesV5RespData) GetMaxFeeForCtAddr() string`

GetMaxFeeForCtAddr returns the MaxFeeForCtAddr field if non-nil, zero value otherwise.

### GetMaxFeeForCtAddrOk

`func (o *GetAssetCurrenciesV5RespData) GetMaxFeeForCtAddrOk() (*string, bool)`

GetMaxFeeForCtAddrOk returns a tuple with the MaxFeeForCtAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeeForCtAddr

`func (o *GetAssetCurrenciesV5RespData) SetMaxFeeForCtAddr(v string)`

SetMaxFeeForCtAddr sets MaxFeeForCtAddr field to given value.

### HasMaxFeeForCtAddr

`func (o *GetAssetCurrenciesV5RespData) HasMaxFeeForCtAddr() bool`

HasMaxFeeForCtAddr returns a boolean if a field has been set.

### GetMaxWd

`func (o *GetAssetCurrenciesV5RespData) GetMaxWd() string`

GetMaxWd returns the MaxWd field if non-nil, zero value otherwise.

### GetMaxWdOk

`func (o *GetAssetCurrenciesV5RespData) GetMaxWdOk() (*string, bool)`

GetMaxWdOk returns a tuple with the MaxWd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWd

`func (o *GetAssetCurrenciesV5RespData) SetMaxWd(v string)`

SetMaxWd sets MaxWd field to given value.

### HasMaxWd

`func (o *GetAssetCurrenciesV5RespData) HasMaxWd() bool`

HasMaxWd returns a boolean if a field has been set.

### GetMinDep

`func (o *GetAssetCurrenciesV5RespData) GetMinDep() string`

GetMinDep returns the MinDep field if non-nil, zero value otherwise.

### GetMinDepOk

`func (o *GetAssetCurrenciesV5RespData) GetMinDepOk() (*string, bool)`

GetMinDepOk returns a tuple with the MinDep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDep

`func (o *GetAssetCurrenciesV5RespData) SetMinDep(v string)`

SetMinDep sets MinDep field to given value.

### HasMinDep

`func (o *GetAssetCurrenciesV5RespData) HasMinDep() bool`

HasMinDep returns a boolean if a field has been set.

### GetMinDepArrivalConfirm

`func (o *GetAssetCurrenciesV5RespData) GetMinDepArrivalConfirm() string`

GetMinDepArrivalConfirm returns the MinDepArrivalConfirm field if non-nil, zero value otherwise.

### GetMinDepArrivalConfirmOk

`func (o *GetAssetCurrenciesV5RespData) GetMinDepArrivalConfirmOk() (*string, bool)`

GetMinDepArrivalConfirmOk returns a tuple with the MinDepArrivalConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDepArrivalConfirm

`func (o *GetAssetCurrenciesV5RespData) SetMinDepArrivalConfirm(v string)`

SetMinDepArrivalConfirm sets MinDepArrivalConfirm field to given value.

### HasMinDepArrivalConfirm

`func (o *GetAssetCurrenciesV5RespData) HasMinDepArrivalConfirm() bool`

HasMinDepArrivalConfirm returns a boolean if a field has been set.

### GetMinFee

`func (o *GetAssetCurrenciesV5RespData) GetMinFee() string`

GetMinFee returns the MinFee field if non-nil, zero value otherwise.

### GetMinFeeOk

`func (o *GetAssetCurrenciesV5RespData) GetMinFeeOk() (*string, bool)`

GetMinFeeOk returns a tuple with the MinFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFee

`func (o *GetAssetCurrenciesV5RespData) SetMinFee(v string)`

SetMinFee sets MinFee field to given value.

### HasMinFee

`func (o *GetAssetCurrenciesV5RespData) HasMinFee() bool`

HasMinFee returns a boolean if a field has been set.

### GetMinFeeForCtAddr

`func (o *GetAssetCurrenciesV5RespData) GetMinFeeForCtAddr() string`

GetMinFeeForCtAddr returns the MinFeeForCtAddr field if non-nil, zero value otherwise.

### GetMinFeeForCtAddrOk

`func (o *GetAssetCurrenciesV5RespData) GetMinFeeForCtAddrOk() (*string, bool)`

GetMinFeeForCtAddrOk returns a tuple with the MinFeeForCtAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFeeForCtAddr

`func (o *GetAssetCurrenciesV5RespData) SetMinFeeForCtAddr(v string)`

SetMinFeeForCtAddr sets MinFeeForCtAddr field to given value.

### HasMinFeeForCtAddr

`func (o *GetAssetCurrenciesV5RespData) HasMinFeeForCtAddr() bool`

HasMinFeeForCtAddr returns a boolean if a field has been set.

### GetMinInternal

`func (o *GetAssetCurrenciesV5RespData) GetMinInternal() string`

GetMinInternal returns the MinInternal field if non-nil, zero value otherwise.

### GetMinInternalOk

`func (o *GetAssetCurrenciesV5RespData) GetMinInternalOk() (*string, bool)`

GetMinInternalOk returns a tuple with the MinInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInternal

`func (o *GetAssetCurrenciesV5RespData) SetMinInternal(v string)`

SetMinInternal sets MinInternal field to given value.

### HasMinInternal

`func (o *GetAssetCurrenciesV5RespData) HasMinInternal() bool`

HasMinInternal returns a boolean if a field has been set.

### GetMinWd

`func (o *GetAssetCurrenciesV5RespData) GetMinWd() string`

GetMinWd returns the MinWd field if non-nil, zero value otherwise.

### GetMinWdOk

`func (o *GetAssetCurrenciesV5RespData) GetMinWdOk() (*string, bool)`

GetMinWdOk returns a tuple with the MinWd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWd

`func (o *GetAssetCurrenciesV5RespData) SetMinWd(v string)`

SetMinWd sets MinWd field to given value.

### HasMinWd

`func (o *GetAssetCurrenciesV5RespData) HasMinWd() bool`

HasMinWd returns a boolean if a field has been set.

### GetMinWdUnlockConfirm

`func (o *GetAssetCurrenciesV5RespData) GetMinWdUnlockConfirm() string`

GetMinWdUnlockConfirm returns the MinWdUnlockConfirm field if non-nil, zero value otherwise.

### GetMinWdUnlockConfirmOk

`func (o *GetAssetCurrenciesV5RespData) GetMinWdUnlockConfirmOk() (*string, bool)`

GetMinWdUnlockConfirmOk returns a tuple with the MinWdUnlockConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWdUnlockConfirm

`func (o *GetAssetCurrenciesV5RespData) SetMinWdUnlockConfirm(v string)`

SetMinWdUnlockConfirm sets MinWdUnlockConfirm field to given value.

### HasMinWdUnlockConfirm

`func (o *GetAssetCurrenciesV5RespData) HasMinWdUnlockConfirm() bool`

HasMinWdUnlockConfirm returns a boolean if a field has been set.

### GetName

`func (o *GetAssetCurrenciesV5RespData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAssetCurrenciesV5RespData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAssetCurrenciesV5RespData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAssetCurrenciesV5RespData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedTag

`func (o *GetAssetCurrenciesV5RespData) GetNeedTag() bool`

GetNeedTag returns the NeedTag field if non-nil, zero value otherwise.

### GetNeedTagOk

`func (o *GetAssetCurrenciesV5RespData) GetNeedTagOk() (*bool, bool)`

GetNeedTagOk returns a tuple with the NeedTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedTag

`func (o *GetAssetCurrenciesV5RespData) SetNeedTag(v bool)`

SetNeedTag sets NeedTag field to given value.

### HasNeedTag

`func (o *GetAssetCurrenciesV5RespData) HasNeedTag() bool`

HasNeedTag returns a boolean if a field has been set.

### GetUsedDepQuotaFixed

`func (o *GetAssetCurrenciesV5RespData) GetUsedDepQuotaFixed() string`

GetUsedDepQuotaFixed returns the UsedDepQuotaFixed field if non-nil, zero value otherwise.

### GetUsedDepQuotaFixedOk

`func (o *GetAssetCurrenciesV5RespData) GetUsedDepQuotaFixedOk() (*string, bool)`

GetUsedDepQuotaFixedOk returns a tuple with the UsedDepQuotaFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedDepQuotaFixed

`func (o *GetAssetCurrenciesV5RespData) SetUsedDepQuotaFixed(v string)`

SetUsedDepQuotaFixed sets UsedDepQuotaFixed field to given value.

### HasUsedDepQuotaFixed

`func (o *GetAssetCurrenciesV5RespData) HasUsedDepQuotaFixed() bool`

HasUsedDepQuotaFixed returns a boolean if a field has been set.

### GetUsedWdQuota

`func (o *GetAssetCurrenciesV5RespData) GetUsedWdQuota() string`

GetUsedWdQuota returns the UsedWdQuota field if non-nil, zero value otherwise.

### GetUsedWdQuotaOk

`func (o *GetAssetCurrenciesV5RespData) GetUsedWdQuotaOk() (*string, bool)`

GetUsedWdQuotaOk returns a tuple with the UsedWdQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedWdQuota

`func (o *GetAssetCurrenciesV5RespData) SetUsedWdQuota(v string)`

SetUsedWdQuota sets UsedWdQuota field to given value.

### HasUsedWdQuota

`func (o *GetAssetCurrenciesV5RespData) HasUsedWdQuota() bool`

HasUsedWdQuota returns a boolean if a field has been set.

### GetWdEstOpenTime

`func (o *GetAssetCurrenciesV5RespData) GetWdEstOpenTime() string`

GetWdEstOpenTime returns the WdEstOpenTime field if non-nil, zero value otherwise.

### GetWdEstOpenTimeOk

`func (o *GetAssetCurrenciesV5RespData) GetWdEstOpenTimeOk() (*string, bool)`

GetWdEstOpenTimeOk returns a tuple with the WdEstOpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWdEstOpenTime

`func (o *GetAssetCurrenciesV5RespData) SetWdEstOpenTime(v string)`

SetWdEstOpenTime sets WdEstOpenTime field to given value.

### HasWdEstOpenTime

`func (o *GetAssetCurrenciesV5RespData) HasWdEstOpenTime() bool`

HasWdEstOpenTime returns a boolean if a field has been set.

### GetWdQuota

`func (o *GetAssetCurrenciesV5RespData) GetWdQuota() string`

GetWdQuota returns the WdQuota field if non-nil, zero value otherwise.

### GetWdQuotaOk

`func (o *GetAssetCurrenciesV5RespData) GetWdQuotaOk() (*string, bool)`

GetWdQuotaOk returns a tuple with the WdQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWdQuota

`func (o *GetAssetCurrenciesV5RespData) SetWdQuota(v string)`

SetWdQuota sets WdQuota field to given value.

### HasWdQuota

`func (o *GetAssetCurrenciesV5RespData) HasWdQuota() bool`

HasWdQuota returns a boolean if a field has been set.

### GetWdTickSz

`func (o *GetAssetCurrenciesV5RespData) GetWdTickSz() string`

GetWdTickSz returns the WdTickSz field if non-nil, zero value otherwise.

### GetWdTickSzOk

`func (o *GetAssetCurrenciesV5RespData) GetWdTickSzOk() (*string, bool)`

GetWdTickSzOk returns a tuple with the WdTickSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWdTickSz

`func (o *GetAssetCurrenciesV5RespData) SetWdTickSz(v string)`

SetWdTickSz sets WdTickSz field to given value.

### HasWdTickSz

`func (o *GetAssetCurrenciesV5RespData) HasWdTickSz() bool`

HasWdTickSz returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


