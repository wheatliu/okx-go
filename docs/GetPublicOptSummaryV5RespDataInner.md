# GetPublicOptSummaryV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AskVol** | Pointer to **string** | Ask volatility | [optional] [default to ""]
**BidVol** | Pointer to **string** | Bid volatility | [optional] [default to ""]
**Delta** | Pointer to **string** | Sensitivity of option price to &#x60;uly&#x60; price | [optional] [default to ""]
**DeltaBS** | Pointer to **string** | Sensitivity of option price to &#x60;uly&#x60; price in BS mode | [optional] [default to ""]
**FwdPx** | Pointer to **string** | Forward price | [optional] [default to ""]
**Gamma** | Pointer to **string** | The delta is sensitivity to &#x60;uly&#x60; price | [optional] [default to ""]
**GammaBS** | Pointer to **string** | The delta is sensitivity to &#x60;uly&#x60; price in BS mode | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. &#x60;BTC-USD-200103-5500-C&#x60; | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type  &#x60;OPTION&#x60; | [optional] [default to ""]
**Lever** | Pointer to **string** | Leverage | [optional] [default to ""]
**MarkVol** | Pointer to **string** | Mark volatility | [optional] [default to ""]
**RealVol** | Pointer to **string** | Realized volatility (not currently used) | [optional] [default to ""]
**Theta** | Pointer to **string** | Sensitivity of option price to remaining maturity | [optional] [default to ""]
**ThetaBS** | Pointer to **string** | Sensitivity of option price to remaining maturity in BS mode | [optional] [default to ""]
**Ts** | Pointer to **string** | Data update time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Uly** | Pointer to **string** | Underlying | [optional] [default to ""]
**Vega** | Pointer to **string** | Sensitivity of option price to implied volatility | [optional] [default to ""]
**VegaBS** | Pointer to **string** | Sensitivity of option price to implied volatility in BS mode | [optional] [default to ""]
**VolLv** | Pointer to **string** | Implied volatility of at-the-money options | [optional] [default to ""]

## Methods

### NewGetPublicOptSummaryV5RespDataInner

`func NewGetPublicOptSummaryV5RespDataInner() *GetPublicOptSummaryV5RespDataInner`

NewGetPublicOptSummaryV5RespDataInner instantiates a new GetPublicOptSummaryV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicOptSummaryV5RespDataInnerWithDefaults

`func NewGetPublicOptSummaryV5RespDataInnerWithDefaults() *GetPublicOptSummaryV5RespDataInner`

NewGetPublicOptSummaryV5RespDataInnerWithDefaults instantiates a new GetPublicOptSummaryV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskVol

`func (o *GetPublicOptSummaryV5RespDataInner) GetAskVol() string`

GetAskVol returns the AskVol field if non-nil, zero value otherwise.

### GetAskVolOk

`func (o *GetPublicOptSummaryV5RespDataInner) GetAskVolOk() (*string, bool)`

GetAskVolOk returns a tuple with the AskVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskVol

`func (o *GetPublicOptSummaryV5RespDataInner) SetAskVol(v string)`

SetAskVol sets AskVol field to given value.

### HasAskVol

`func (o *GetPublicOptSummaryV5RespDataInner) HasAskVol() bool`

HasAskVol returns a boolean if a field has been set.

### GetBidVol

`func (o *GetPublicOptSummaryV5RespDataInner) GetBidVol() string`

GetBidVol returns the BidVol field if non-nil, zero value otherwise.

### GetBidVolOk

`func (o *GetPublicOptSummaryV5RespDataInner) GetBidVolOk() (*string, bool)`

GetBidVolOk returns a tuple with the BidVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidVol

`func (o *GetPublicOptSummaryV5RespDataInner) SetBidVol(v string)`

SetBidVol sets BidVol field to given value.

### HasBidVol

`func (o *GetPublicOptSummaryV5RespDataInner) HasBidVol() bool`

HasBidVol returns a boolean if a field has been set.

### GetDelta

`func (o *GetPublicOptSummaryV5RespDataInner) GetDelta() string`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *GetPublicOptSummaryV5RespDataInner) GetDeltaOk() (*string, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *GetPublicOptSummaryV5RespDataInner) SetDelta(v string)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *GetPublicOptSummaryV5RespDataInner) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### GetDeltaBS

`func (o *GetPublicOptSummaryV5RespDataInner) GetDeltaBS() string`

GetDeltaBS returns the DeltaBS field if non-nil, zero value otherwise.

### GetDeltaBSOk

`func (o *GetPublicOptSummaryV5RespDataInner) GetDeltaBSOk() (*string, bool)`

GetDeltaBSOk returns a tuple with the DeltaBS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaBS

`func (o *GetPublicOptSummaryV5RespDataInner) SetDeltaBS(v string)`

SetDeltaBS sets DeltaBS field to given value.

### HasDeltaBS

`func (o *GetPublicOptSummaryV5RespDataInner) HasDeltaBS() bool`

HasDeltaBS returns a boolean if a field has been set.

### GetFwdPx

`func (o *GetPublicOptSummaryV5RespDataInner) GetFwdPx() string`

GetFwdPx returns the FwdPx field if non-nil, zero value otherwise.

### GetFwdPxOk

`func (o *GetPublicOptSummaryV5RespDataInner) GetFwdPxOk() (*string, bool)`

GetFwdPxOk returns a tuple with the FwdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwdPx

`func (o *GetPublicOptSummaryV5RespDataInner) SetFwdPx(v string)`

SetFwdPx sets FwdPx field to given value.

### HasFwdPx

`func (o *GetPublicOptSummaryV5RespDataInner) HasFwdPx() bool`

HasFwdPx returns a boolean if a field has been set.

### GetGamma

`func (o *GetPublicOptSummaryV5RespDataInner) GetGamma() string`

GetGamma returns the Gamma field if non-nil, zero value otherwise.

### GetGammaOk

`func (o *GetPublicOptSummaryV5RespDataInner) GetGammaOk() (*string, bool)`

GetGammaOk returns a tuple with the Gamma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGamma

`func (o *GetPublicOptSummaryV5RespDataInner) SetGamma(v string)`

SetGamma sets Gamma field to given value.

### HasGamma

`func (o *GetPublicOptSummaryV5RespDataInner) HasGamma() bool`

HasGamma returns a boolean if a field has been set.

### GetGammaBS

`func (o *GetPublicOptSummaryV5RespDataInner) GetGammaBS() string`

GetGammaBS returns the GammaBS field if non-nil, zero value otherwise.

### GetGammaBSOk

`func (o *GetPublicOptSummaryV5RespDataInner) GetGammaBSOk() (*string, bool)`

GetGammaBSOk returns a tuple with the GammaBS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGammaBS

`func (o *GetPublicOptSummaryV5RespDataInner) SetGammaBS(v string)`

SetGammaBS sets GammaBS field to given value.

### HasGammaBS

`func (o *GetPublicOptSummaryV5RespDataInner) HasGammaBS() bool`

HasGammaBS returns a boolean if a field has been set.

### GetInstId

`func (o *GetPublicOptSummaryV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetPublicOptSummaryV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetPublicOptSummaryV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetPublicOptSummaryV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetPublicOptSummaryV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetPublicOptSummaryV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetPublicOptSummaryV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetPublicOptSummaryV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetLever

`func (o *GetPublicOptSummaryV5RespDataInner) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetPublicOptSummaryV5RespDataInner) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetPublicOptSummaryV5RespDataInner) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetPublicOptSummaryV5RespDataInner) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetMarkVol

`func (o *GetPublicOptSummaryV5RespDataInner) GetMarkVol() string`

GetMarkVol returns the MarkVol field if non-nil, zero value otherwise.

### GetMarkVolOk

`func (o *GetPublicOptSummaryV5RespDataInner) GetMarkVolOk() (*string, bool)`

GetMarkVolOk returns a tuple with the MarkVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkVol

`func (o *GetPublicOptSummaryV5RespDataInner) SetMarkVol(v string)`

SetMarkVol sets MarkVol field to given value.

### HasMarkVol

`func (o *GetPublicOptSummaryV5RespDataInner) HasMarkVol() bool`

HasMarkVol returns a boolean if a field has been set.

### GetRealVol

`func (o *GetPublicOptSummaryV5RespDataInner) GetRealVol() string`

GetRealVol returns the RealVol field if non-nil, zero value otherwise.

### GetRealVolOk

`func (o *GetPublicOptSummaryV5RespDataInner) GetRealVolOk() (*string, bool)`

GetRealVolOk returns a tuple with the RealVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealVol

`func (o *GetPublicOptSummaryV5RespDataInner) SetRealVol(v string)`

SetRealVol sets RealVol field to given value.

### HasRealVol

`func (o *GetPublicOptSummaryV5RespDataInner) HasRealVol() bool`

HasRealVol returns a boolean if a field has been set.

### GetTheta

`func (o *GetPublicOptSummaryV5RespDataInner) GetTheta() string`

GetTheta returns the Theta field if non-nil, zero value otherwise.

### GetThetaOk

`func (o *GetPublicOptSummaryV5RespDataInner) GetThetaOk() (*string, bool)`

GetThetaOk returns a tuple with the Theta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheta

`func (o *GetPublicOptSummaryV5RespDataInner) SetTheta(v string)`

SetTheta sets Theta field to given value.

### HasTheta

`func (o *GetPublicOptSummaryV5RespDataInner) HasTheta() bool`

HasTheta returns a boolean if a field has been set.

### GetThetaBS

`func (o *GetPublicOptSummaryV5RespDataInner) GetThetaBS() string`

GetThetaBS returns the ThetaBS field if non-nil, zero value otherwise.

### GetThetaBSOk

`func (o *GetPublicOptSummaryV5RespDataInner) GetThetaBSOk() (*string, bool)`

GetThetaBSOk returns a tuple with the ThetaBS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThetaBS

`func (o *GetPublicOptSummaryV5RespDataInner) SetThetaBS(v string)`

SetThetaBS sets ThetaBS field to given value.

### HasThetaBS

`func (o *GetPublicOptSummaryV5RespDataInner) HasThetaBS() bool`

HasThetaBS returns a boolean if a field has been set.

### GetTs

`func (o *GetPublicOptSummaryV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetPublicOptSummaryV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetPublicOptSummaryV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetPublicOptSummaryV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetUly

`func (o *GetPublicOptSummaryV5RespDataInner) GetUly() string`

GetUly returns the Uly field if non-nil, zero value otherwise.

### GetUlyOk

`func (o *GetPublicOptSummaryV5RespDataInner) GetUlyOk() (*string, bool)`

GetUlyOk returns a tuple with the Uly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUly

`func (o *GetPublicOptSummaryV5RespDataInner) SetUly(v string)`

SetUly sets Uly field to given value.

### HasUly

`func (o *GetPublicOptSummaryV5RespDataInner) HasUly() bool`

HasUly returns a boolean if a field has been set.

### GetVega

`func (o *GetPublicOptSummaryV5RespDataInner) GetVega() string`

GetVega returns the Vega field if non-nil, zero value otherwise.

### GetVegaOk

`func (o *GetPublicOptSummaryV5RespDataInner) GetVegaOk() (*string, bool)`

GetVegaOk returns a tuple with the Vega field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVega

`func (o *GetPublicOptSummaryV5RespDataInner) SetVega(v string)`

SetVega sets Vega field to given value.

### HasVega

`func (o *GetPublicOptSummaryV5RespDataInner) HasVega() bool`

HasVega returns a boolean if a field has been set.

### GetVegaBS

`func (o *GetPublicOptSummaryV5RespDataInner) GetVegaBS() string`

GetVegaBS returns the VegaBS field if non-nil, zero value otherwise.

### GetVegaBSOk

`func (o *GetPublicOptSummaryV5RespDataInner) GetVegaBSOk() (*string, bool)`

GetVegaBSOk returns a tuple with the VegaBS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVegaBS

`func (o *GetPublicOptSummaryV5RespDataInner) SetVegaBS(v string)`

SetVegaBS sets VegaBS field to given value.

### HasVegaBS

`func (o *GetPublicOptSummaryV5RespDataInner) HasVegaBS() bool`

HasVegaBS returns a boolean if a field has been set.

### GetVolLv

`func (o *GetPublicOptSummaryV5RespDataInner) GetVolLv() string`

GetVolLv returns the VolLv field if non-nil, zero value otherwise.

### GetVolLvOk

`func (o *GetPublicOptSummaryV5RespDataInner) GetVolLvOk() (*string, bool)`

GetVolLvOk returns a tuple with the VolLv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolLv

`func (o *GetPublicOptSummaryV5RespDataInner) SetVolLv(v string)`

SetVolLv sets VolLv field to given value.

### HasVolLv

`func (o *GetPublicOptSummaryV5RespDataInner) HasVolLv() bool`

HasVolLv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


