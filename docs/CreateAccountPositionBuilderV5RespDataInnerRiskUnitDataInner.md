# CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | Pointer to **string** | When &#x60;instType&#x60; is &#x60;SPOT&#x60;, it represents spot in use.  When &#x60;instType&#x60; is &#x60;SWAP&#x60;/&#x60;FUTURES&#x60;/&#x60;OPTION&#x60;, it represents position amount. | [optional] [default to ""]
**AvgPx** | Pointer to **string** | Average open price | [optional] [default to ""]
**Delta** | Pointer to **string** | When &#x60;instType&#x60; is &#x60;SPOT&#x60;, it represents asset amount.  When &#x60;instType&#x60; is &#x60;SWAP&#x60;/&#x60;FUTURES&#x60;/&#x60;OPTION&#x60;, it represents the rate of change in the contract’s price with respect to changes in the underlying asset’s price (by Instrument ID). | [optional] [default to ""]
**FloatPnl** | Pointer to **string** | Float P&amp;L | [optional] [default to ""]
**Gamma** | Pointer to **string** | The rate of change in the delta with respect to changes in the underlying price (by Instrument ID).   When &#x60;instType&#x60; is &#x60;SPOT&#x60;, it will returns \&quot;\&quot;. | [optional] [default to ""]
**Imr** | Pointer to **string** | Risk unit IMR (&#x60;USD&#x60;) | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60; | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [optional] [default to ""]
**IsRealPos** | Pointer to **bool** | Whether it is a real position  If &#x60;instType&#x60; is &#x60;SWAP&#x60;/&#x60;FUTURES&#x60;/&#x60;OPTION&#x60;, it is a valid parameter, else it will returns &#x60;false&#x60; | [optional] 
**MarkPx** | Pointer to **string** | Mark price | [optional] [default to ""]
**Mmr** | Pointer to **string** | Risk unit MMR (&#x60;USD&#x60;) | [optional] [default to ""]
**Mr1** | Pointer to **string** | Stress testing value of spot and volatility (all derivatives, and spot trading in spot-derivatives risk offset mode) | [optional] [default to ""]
**Mr1FinalResult** | Pointer to [**CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult**](CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult.md) |  | [optional] 
**Mr1Scenarios** | Pointer to [**CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios**](CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios.md) |  | [optional] 
**Mr2** | Pointer to **string** | Stress testing value of time value of money (TVM) (for options) | [optional] [default to ""]
**Mr3** | Pointer to **string** | Stress testing value of volatility span (for options) | [optional] [default to ""]
**Mr4** | Pointer to **string** | Stress testing value of basis (for all derivatives) | [optional] [default to ""]
**Mr5** | Pointer to **string** | Stress testing value of interest rate risk (for options) | [optional] [default to ""]
**Mr6** | Pointer to **string** | Stress testing value of extremely volatile markets (for all derivatives, and spot trading in spot-derivatives risk offset mode) | [optional] [default to ""]
**Mr6FinalResult** | Pointer to [**CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr6FinalResult**](CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr6FinalResult.md) |  | [optional] 
**Mr7** | Pointer to **string** | Stress testing value of position reduction cost (for all derivatives) | [optional] [default to ""]
**Mr8** | Pointer to **string** | Borrowing MMR/IMR | [optional] [default to ""]
**Mr9** | Pointer to **string** | USDT-USDC-USD hedge risk | [optional] [default to ""]
**NotionalUsd** | Pointer to **string** | Notional in &#x60;USD&#x60; | [optional] [default to ""]
**Pnl** | Pointer to **string** | MR6 stress P&amp;L (&#x60;USD&#x60;) | [optional] [default to ""]
**Portfolios** | Pointer to [**[]CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerPortfoliosInner**](CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerPortfoliosInner.md) | Portfolios info  Only applicable to &#x60;Portfolio margin&#x60; | [optional] 
**PosSide** | Pointer to **string** | Position side  &#x60;long&#x60;  &#x60;short&#x60;  &#x60;net&#x60; | [optional] [default to ""]
**RiskUnit** | Pointer to **string** | Risk unit, e.g. &#x60;BTC&#x60; | [optional] [default to ""]
**SpotShock** | Pointer to **string** | MR6 worst-case scenario spot shock (in percentage), e.g. &#x60;0.01&#x60; representing &#x60;1%&#x60; | [optional] [default to ""]
**Theta** | Pointer to **string** | The change in contract price each day closer to expiry (by Instrument ID).  When &#x60;instType&#x60; is &#x60;SPOT&#x60;, it will returns \&quot;\&quot;. | [optional] [default to ""]
**Upl** | Pointer to **string** | Risk unit UPL (&#x60;USD&#x60;) | [optional] [default to ""]
**Vega** | Pointer to **string** | The change of the option price when underlying volatility increases by 1% (by Instrument ID).  When &#x60;instType&#x60; is &#x60;SPOT&#x60;, it will returns \&quot;\&quot;. | [optional] [default to ""]
**VolSame** | Pointer to **map[string]interface{}** | When the volatility keep the same, the P&amp;L of stress tests under different price volatility ratios, format in {&#x60;change&#x60;: &#x60;value&#x60;,...}  &#x60;change&#x60;: price volatility ratio (in percentage), e.g. &#x60;0.01&#x60; representing &#x60;1%&#x60;  &#x60;value&#x60;: P&amp;L under stress tests, measured in &#x60;USD&#x60;  e.g. {\&quot;-0.15\&quot;:\&quot;-2333.23\&quot;, ...} | [optional] 
**VolShock** | Pointer to **string** | MR1 worst-case scenario volatility shock  &#x60;down&#x60;: volatility shock down  &#x60;unchange&#x60;: volatility unchanged  &#x60;up&#x60;: volatility shock up | [optional] [default to ""]
**VolShockDown** | Pointer to **map[string]interface{}** | When the volatility shock down, the P&amp;L of stress tests under different price volatility ratios, format in {&#x60;change&#x60;: &#x60;value&#x60;,...}  &#x60;change&#x60;: price volatility ratio (in percentage), e.g. &#x60;0.01&#x60; representing &#x60;1%&#x60;  &#x60;value&#x60;: P&amp;L under stress tests, measured in &#x60;USD&#x60;  e.g. {\&quot;-0.15\&quot;:\&quot;-2333.23\&quot;, ...} | [optional] 
**VolShockUp** | Pointer to **map[string]interface{}** | When the volatility shock up, the P&amp;L of stress tests under different price volatility ratios, format in {&#x60;change&#x60;: &#x60;value&#x60;,...}  &#x60;change&#x60;: price volatility ratio (in percentage), e.g. &#x60;0.01&#x60; representing &#x60;1%&#x60;  &#x60;value&#x60;: P&amp;L under stress tests, measured in &#x60;USD&#x60;  e.g. {\&quot;-0.15\&quot;:\&quot;-2333.23\&quot;, ...} | [optional] 

## Methods

### NewCreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner

`func NewCreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner() *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner`

NewCreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner instantiates a new CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerWithDefaults

`func NewCreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerWithDefaults() *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner`

NewCreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerWithDefaults instantiates a new CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetAvgPx

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetAvgPx() string`

GetAvgPx returns the AvgPx field if non-nil, zero value otherwise.

### GetAvgPxOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetAvgPxOk() (*string, bool)`

GetAvgPxOk returns a tuple with the AvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPx

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetAvgPx(v string)`

SetAvgPx sets AvgPx field to given value.

### HasAvgPx

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasAvgPx() bool`

HasAvgPx returns a boolean if a field has been set.

### GetDelta

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetDelta() string`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetDeltaOk() (*string, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetDelta(v string)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### GetFloatPnl

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetFloatPnl() string`

GetFloatPnl returns the FloatPnl field if non-nil, zero value otherwise.

### GetFloatPnlOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetFloatPnlOk() (*string, bool)`

GetFloatPnlOk returns a tuple with the FloatPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatPnl

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetFloatPnl(v string)`

SetFloatPnl sets FloatPnl field to given value.

### HasFloatPnl

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasFloatPnl() bool`

HasFloatPnl returns a boolean if a field has been set.

### GetGamma

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetGamma() string`

GetGamma returns the Gamma field if non-nil, zero value otherwise.

### GetGammaOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetGammaOk() (*string, bool)`

GetGammaOk returns a tuple with the Gamma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGamma

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetGamma(v string)`

SetGamma sets Gamma field to given value.

### HasGamma

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasGamma() bool`

HasGamma returns a boolean if a field has been set.

### GetImr

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetImr() string`

GetImr returns the Imr field if non-nil, zero value otherwise.

### GetImrOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetImrOk() (*string, bool)`

GetImrOk returns a tuple with the Imr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImr

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetImr(v string)`

SetImr sets Imr field to given value.

### HasImr

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasImr() bool`

HasImr returns a boolean if a field has been set.

### GetInstId

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetIsRealPos

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetIsRealPos() bool`

GetIsRealPos returns the IsRealPos field if non-nil, zero value otherwise.

### GetIsRealPosOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetIsRealPosOk() (*bool, bool)`

GetIsRealPosOk returns a tuple with the IsRealPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRealPos

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetIsRealPos(v bool)`

SetIsRealPos sets IsRealPos field to given value.

### HasIsRealPos

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasIsRealPos() bool`

HasIsRealPos returns a boolean if a field has been set.

### GetMarkPx

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMarkPx() string`

GetMarkPx returns the MarkPx field if non-nil, zero value otherwise.

### GetMarkPxOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMarkPxOk() (*string, bool)`

GetMarkPxOk returns a tuple with the MarkPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPx

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetMarkPx(v string)`

SetMarkPx sets MarkPx field to given value.

### HasMarkPx

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasMarkPx() bool`

HasMarkPx returns a boolean if a field has been set.

### GetMmr

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMmr() string`

GetMmr returns the Mmr field if non-nil, zero value otherwise.

### GetMmrOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMmrOk() (*string, bool)`

GetMmrOk returns a tuple with the Mmr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmr

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetMmr(v string)`

SetMmr sets Mmr field to given value.

### HasMmr

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasMmr() bool`

HasMmr returns a boolean if a field has been set.

### GetMr1

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr1() string`

GetMr1 returns the Mr1 field if non-nil, zero value otherwise.

### GetMr1Ok

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr1Ok() (*string, bool)`

GetMr1Ok returns a tuple with the Mr1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMr1

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetMr1(v string)`

SetMr1 sets Mr1 field to given value.

### HasMr1

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasMr1() bool`

HasMr1 returns a boolean if a field has been set.

### GetMr1FinalResult

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr1FinalResult() CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult`

GetMr1FinalResult returns the Mr1FinalResult field if non-nil, zero value otherwise.

### GetMr1FinalResultOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr1FinalResultOk() (*CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult, bool)`

GetMr1FinalResultOk returns a tuple with the Mr1FinalResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMr1FinalResult

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetMr1FinalResult(v CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult)`

SetMr1FinalResult sets Mr1FinalResult field to given value.

### HasMr1FinalResult

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasMr1FinalResult() bool`

HasMr1FinalResult returns a boolean if a field has been set.

### GetMr1Scenarios

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr1Scenarios() CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios`

GetMr1Scenarios returns the Mr1Scenarios field if non-nil, zero value otherwise.

### GetMr1ScenariosOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr1ScenariosOk() (*CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios, bool)`

GetMr1ScenariosOk returns a tuple with the Mr1Scenarios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMr1Scenarios

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetMr1Scenarios(v CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios)`

SetMr1Scenarios sets Mr1Scenarios field to given value.

### HasMr1Scenarios

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasMr1Scenarios() bool`

HasMr1Scenarios returns a boolean if a field has been set.

### GetMr2

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr2() string`

GetMr2 returns the Mr2 field if non-nil, zero value otherwise.

### GetMr2Ok

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr2Ok() (*string, bool)`

GetMr2Ok returns a tuple with the Mr2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMr2

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetMr2(v string)`

SetMr2 sets Mr2 field to given value.

### HasMr2

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasMr2() bool`

HasMr2 returns a boolean if a field has been set.

### GetMr3

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr3() string`

GetMr3 returns the Mr3 field if non-nil, zero value otherwise.

### GetMr3Ok

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr3Ok() (*string, bool)`

GetMr3Ok returns a tuple with the Mr3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMr3

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetMr3(v string)`

SetMr3 sets Mr3 field to given value.

### HasMr3

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasMr3() bool`

HasMr3 returns a boolean if a field has been set.

### GetMr4

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr4() string`

GetMr4 returns the Mr4 field if non-nil, zero value otherwise.

### GetMr4Ok

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr4Ok() (*string, bool)`

GetMr4Ok returns a tuple with the Mr4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMr4

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetMr4(v string)`

SetMr4 sets Mr4 field to given value.

### HasMr4

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasMr4() bool`

HasMr4 returns a boolean if a field has been set.

### GetMr5

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr5() string`

GetMr5 returns the Mr5 field if non-nil, zero value otherwise.

### GetMr5Ok

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr5Ok() (*string, bool)`

GetMr5Ok returns a tuple with the Mr5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMr5

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetMr5(v string)`

SetMr5 sets Mr5 field to given value.

### HasMr5

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasMr5() bool`

HasMr5 returns a boolean if a field has been set.

### GetMr6

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr6() string`

GetMr6 returns the Mr6 field if non-nil, zero value otherwise.

### GetMr6Ok

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr6Ok() (*string, bool)`

GetMr6Ok returns a tuple with the Mr6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMr6

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetMr6(v string)`

SetMr6 sets Mr6 field to given value.

### HasMr6

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasMr6() bool`

HasMr6 returns a boolean if a field has been set.

### GetMr6FinalResult

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr6FinalResult() CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr6FinalResult`

GetMr6FinalResult returns the Mr6FinalResult field if non-nil, zero value otherwise.

### GetMr6FinalResultOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr6FinalResultOk() (*CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr6FinalResult, bool)`

GetMr6FinalResultOk returns a tuple with the Mr6FinalResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMr6FinalResult

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetMr6FinalResult(v CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr6FinalResult)`

SetMr6FinalResult sets Mr6FinalResult field to given value.

### HasMr6FinalResult

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasMr6FinalResult() bool`

HasMr6FinalResult returns a boolean if a field has been set.

### GetMr7

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr7() string`

GetMr7 returns the Mr7 field if non-nil, zero value otherwise.

### GetMr7Ok

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr7Ok() (*string, bool)`

GetMr7Ok returns a tuple with the Mr7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMr7

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetMr7(v string)`

SetMr7 sets Mr7 field to given value.

### HasMr7

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasMr7() bool`

HasMr7 returns a boolean if a field has been set.

### GetMr8

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr8() string`

GetMr8 returns the Mr8 field if non-nil, zero value otherwise.

### GetMr8Ok

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr8Ok() (*string, bool)`

GetMr8Ok returns a tuple with the Mr8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMr8

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetMr8(v string)`

SetMr8 sets Mr8 field to given value.

### HasMr8

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasMr8() bool`

HasMr8 returns a boolean if a field has been set.

### GetMr9

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr9() string`

GetMr9 returns the Mr9 field if non-nil, zero value otherwise.

### GetMr9Ok

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetMr9Ok() (*string, bool)`

GetMr9Ok returns a tuple with the Mr9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMr9

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetMr9(v string)`

SetMr9 sets Mr9 field to given value.

### HasMr9

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasMr9() bool`

HasMr9 returns a boolean if a field has been set.

### GetNotionalUsd

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetNotionalUsd() string`

GetNotionalUsd returns the NotionalUsd field if non-nil, zero value otherwise.

### GetNotionalUsdOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetNotionalUsdOk() (*string, bool)`

GetNotionalUsdOk returns a tuple with the NotionalUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalUsd

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetNotionalUsd(v string)`

SetNotionalUsd sets NotionalUsd field to given value.

### HasNotionalUsd

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasNotionalUsd() bool`

HasNotionalUsd returns a boolean if a field has been set.

### GetPnl

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetPnl() string`

GetPnl returns the Pnl field if non-nil, zero value otherwise.

### GetPnlOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetPnlOk() (*string, bool)`

GetPnlOk returns a tuple with the Pnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnl

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetPnl(v string)`

SetPnl sets Pnl field to given value.

### HasPnl

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasPnl() bool`

HasPnl returns a boolean if a field has been set.

### GetPortfolios

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetPortfolios() []CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerPortfoliosInner`

GetPortfolios returns the Portfolios field if non-nil, zero value otherwise.

### GetPortfoliosOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetPortfoliosOk() (*[]CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerPortfoliosInner, bool)`

GetPortfoliosOk returns a tuple with the Portfolios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolios

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetPortfolios(v []CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerPortfoliosInner)`

SetPortfolios sets Portfolios field to given value.

### HasPortfolios

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasPortfolios() bool`

HasPortfolios returns a boolean if a field has been set.

### GetPosSide

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetRiskUnit

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetRiskUnit() string`

GetRiskUnit returns the RiskUnit field if non-nil, zero value otherwise.

### GetRiskUnitOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetRiskUnitOk() (*string, bool)`

GetRiskUnitOk returns a tuple with the RiskUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskUnit

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetRiskUnit(v string)`

SetRiskUnit sets RiskUnit field to given value.

### HasRiskUnit

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasRiskUnit() bool`

HasRiskUnit returns a boolean if a field has been set.

### GetSpotShock

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetSpotShock() string`

GetSpotShock returns the SpotShock field if non-nil, zero value otherwise.

### GetSpotShockOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetSpotShockOk() (*string, bool)`

GetSpotShockOk returns a tuple with the SpotShock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotShock

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetSpotShock(v string)`

SetSpotShock sets SpotShock field to given value.

### HasSpotShock

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasSpotShock() bool`

HasSpotShock returns a boolean if a field has been set.

### GetTheta

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetTheta() string`

GetTheta returns the Theta field if non-nil, zero value otherwise.

### GetThetaOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetThetaOk() (*string, bool)`

GetThetaOk returns a tuple with the Theta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheta

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetTheta(v string)`

SetTheta sets Theta field to given value.

### HasTheta

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasTheta() bool`

HasTheta returns a boolean if a field has been set.

### GetUpl

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetUpl() string`

GetUpl returns the Upl field if non-nil, zero value otherwise.

### GetUplOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetUplOk() (*string, bool)`

GetUplOk returns a tuple with the Upl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpl

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetUpl(v string)`

SetUpl sets Upl field to given value.

### HasUpl

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasUpl() bool`

HasUpl returns a boolean if a field has been set.

### GetVega

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetVega() string`

GetVega returns the Vega field if non-nil, zero value otherwise.

### GetVegaOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetVegaOk() (*string, bool)`

GetVegaOk returns a tuple with the Vega field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVega

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetVega(v string)`

SetVega sets Vega field to given value.

### HasVega

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasVega() bool`

HasVega returns a boolean if a field has been set.

### GetVolSame

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetVolSame() map[string]interface{}`

GetVolSame returns the VolSame field if non-nil, zero value otherwise.

### GetVolSameOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetVolSameOk() (*map[string]interface{}, bool)`

GetVolSameOk returns a tuple with the VolSame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolSame

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetVolSame(v map[string]interface{})`

SetVolSame sets VolSame field to given value.

### HasVolSame

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasVolSame() bool`

HasVolSame returns a boolean if a field has been set.

### GetVolShock

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetVolShock() string`

GetVolShock returns the VolShock field if non-nil, zero value otherwise.

### GetVolShockOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetVolShockOk() (*string, bool)`

GetVolShockOk returns a tuple with the VolShock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolShock

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetVolShock(v string)`

SetVolShock sets VolShock field to given value.

### HasVolShock

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasVolShock() bool`

HasVolShock returns a boolean if a field has been set.

### GetVolShockDown

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetVolShockDown() map[string]interface{}`

GetVolShockDown returns the VolShockDown field if non-nil, zero value otherwise.

### GetVolShockDownOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetVolShockDownOk() (*map[string]interface{}, bool)`

GetVolShockDownOk returns a tuple with the VolShockDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolShockDown

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetVolShockDown(v map[string]interface{})`

SetVolShockDown sets VolShockDown field to given value.

### HasVolShockDown

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasVolShockDown() bool`

HasVolShockDown returns a boolean if a field has been set.

### GetVolShockUp

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetVolShockUp() map[string]interface{}`

GetVolShockUp returns the VolShockUp field if non-nil, zero value otherwise.

### GetVolShockUpOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) GetVolShockUpOk() (*map[string]interface{}, bool)`

GetVolShockUpOk returns a tuple with the VolShockUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolShockUp

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) SetVolShockUp(v map[string]interface{})`

SetVolShockUp sets VolShockUp field to given value.

### HasVolShockUp

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner) HasVolShockUp() bool`

HasVolShockUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


