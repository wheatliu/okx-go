# CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | Pointer to **string** | When &#x60;instType&#x60; is &#x60;SPOT&#x60;, it represents spot in use.  When &#x60;instType&#x60; is &#x60;SWAP&#x60;/&#x60;FUTURES&#x60;/&#x60;OPTION&#x60;, it represents position amount. | [optional] [default to ""]
**AvgPx** | Pointer to **string** | Average open price | [optional] [default to ""]
**Delta** | Pointer to **string** | When &#x60;instType&#x60; is &#x60;SPOT&#x60;, it represents asset amount.  When &#x60;instType&#x60; is &#x60;SWAP&#x60;/&#x60;FUTURES&#x60;/&#x60;OPTION&#x60;, it represents the rate of change in the contract’s price with respect to changes in the underlying asset’s price (by Instrument ID). | [optional] [default to ""]
**FloatPnl** | Pointer to **string** | Float P&amp;L | [optional] [default to ""]
**Gamma** | Pointer to **string** | The rate of change in the delta with respect to changes in the underlying price (by Instrument ID).   When &#x60;instType&#x60; is &#x60;SPOT&#x60;, it will returns \&quot;\&quot;. | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60; | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [optional] [default to ""]
**IsRealPos** | Pointer to **bool** | Whether it is a real position  If &#x60;instType&#x60; is &#x60;SWAP&#x60;/&#x60;FUTURES&#x60;/&#x60;OPTION&#x60;, it is a valid parameter, else it will returns &#x60;false&#x60; | [optional] 
**MarkPx** | Pointer to **string** | Mark price | [optional] [default to ""]
**NotionalUsd** | Pointer to **string** | Notional in &#x60;USD&#x60; | [optional] [default to ""]
**PosSide** | Pointer to **string** | Position side  &#x60;long&#x60;  &#x60;short&#x60;  &#x60;net&#x60; | [optional] [default to ""]
**Theta** | Pointer to **string** | The change in contract price each day closer to expiry (by Instrument ID).  When &#x60;instType&#x60; is &#x60;SPOT&#x60;, it will returns \&quot;\&quot;. | [optional] [default to ""]
**Vega** | Pointer to **string** | The change of the option price when underlying volatility increases by 1% (by Instrument ID).  When &#x60;instType&#x60; is &#x60;SPOT&#x60;, it will returns \&quot;\&quot;. | [optional] [default to ""]

## Methods

### NewCreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner

`func NewCreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner() *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner`

NewCreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner instantiates a new CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInnerWithDefaults

`func NewCreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInnerWithDefaults() *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner`

NewCreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInnerWithDefaults instantiates a new CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetAvgPx

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetAvgPx() string`

GetAvgPx returns the AvgPx field if non-nil, zero value otherwise.

### GetAvgPxOk

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetAvgPxOk() (*string, bool)`

GetAvgPxOk returns a tuple with the AvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPx

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) SetAvgPx(v string)`

SetAvgPx sets AvgPx field to given value.

### HasAvgPx

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) HasAvgPx() bool`

HasAvgPx returns a boolean if a field has been set.

### GetDelta

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetDelta() string`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetDeltaOk() (*string, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) SetDelta(v string)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### GetFloatPnl

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetFloatPnl() string`

GetFloatPnl returns the FloatPnl field if non-nil, zero value otherwise.

### GetFloatPnlOk

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetFloatPnlOk() (*string, bool)`

GetFloatPnlOk returns a tuple with the FloatPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatPnl

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) SetFloatPnl(v string)`

SetFloatPnl sets FloatPnl field to given value.

### HasFloatPnl

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) HasFloatPnl() bool`

HasFloatPnl returns a boolean if a field has been set.

### GetGamma

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetGamma() string`

GetGamma returns the Gamma field if non-nil, zero value otherwise.

### GetGammaOk

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetGammaOk() (*string, bool)`

GetGammaOk returns a tuple with the Gamma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGamma

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) SetGamma(v string)`

SetGamma sets Gamma field to given value.

### HasGamma

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) HasGamma() bool`

HasGamma returns a boolean if a field has been set.

### GetInstId

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetIsRealPos

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetIsRealPos() bool`

GetIsRealPos returns the IsRealPos field if non-nil, zero value otherwise.

### GetIsRealPosOk

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetIsRealPosOk() (*bool, bool)`

GetIsRealPosOk returns a tuple with the IsRealPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRealPos

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) SetIsRealPos(v bool)`

SetIsRealPos sets IsRealPos field to given value.

### HasIsRealPos

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) HasIsRealPos() bool`

HasIsRealPos returns a boolean if a field has been set.

### GetMarkPx

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetMarkPx() string`

GetMarkPx returns the MarkPx field if non-nil, zero value otherwise.

### GetMarkPxOk

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetMarkPxOk() (*string, bool)`

GetMarkPxOk returns a tuple with the MarkPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPx

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) SetMarkPx(v string)`

SetMarkPx sets MarkPx field to given value.

### HasMarkPx

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) HasMarkPx() bool`

HasMarkPx returns a boolean if a field has been set.

### GetNotionalUsd

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetNotionalUsd() string`

GetNotionalUsd returns the NotionalUsd field if non-nil, zero value otherwise.

### GetNotionalUsdOk

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetNotionalUsdOk() (*string, bool)`

GetNotionalUsdOk returns a tuple with the NotionalUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalUsd

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) SetNotionalUsd(v string)`

SetNotionalUsd sets NotionalUsd field to given value.

### HasNotionalUsd

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) HasNotionalUsd() bool`

HasNotionalUsd returns a boolean if a field has been set.

### GetPosSide

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetTheta

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetTheta() string`

GetTheta returns the Theta field if non-nil, zero value otherwise.

### GetThetaOk

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetThetaOk() (*string, bool)`

GetThetaOk returns a tuple with the Theta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheta

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) SetTheta(v string)`

SetTheta sets Theta field to given value.

### HasTheta

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) HasTheta() bool`

HasTheta returns a boolean if a field has been set.

### GetVega

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetVega() string`

GetVega returns the Vega field if non-nil, zero value otherwise.

### GetVegaOk

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) GetVegaOk() (*string, bool)`

GetVegaOk returns a tuple with the Vega field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVega

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) SetVega(v string)`

SetVega sets Vega field to given value.

### HasVega

`func (o *CreateAccountPositionBuilderV5RespDataRiskUnitDataInnerPortfoliosInner) HasVega() bool`

HasVega returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


