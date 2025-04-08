# CreateAccountPositionBuilderV5RespDataPositionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | Pointer to **string** | When &#x60;instType&#x60; is &#x60;SPOT&#x60;, it represents spot in use.  When &#x60;instType&#x60; is &#x60;SWAP&#x60;/&#x60;FUTURES&#x60;/&#x60;OPTION&#x60;, it represents position amount. | [optional] [default to ""]
**AvgPx** | Pointer to **string** | Average open price | [optional] [default to ""]
**FloatPnl** | Pointer to **string** | Float P&amp;L | [optional] [default to ""]
**Imr** | Pointer to **string** | IMR | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60; | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [optional] [default to ""]
**IsRealPos** | Pointer to **bool** | Whether it is a real position  If &#x60;instType&#x60; is &#x60;SWAP&#x60;/&#x60;FUTURES&#x60;/&#x60;OPTION&#x60;, it is a valid parameter, else it will returns &#x60;false&#x60; | [optional] 
**Lever** | Pointer to **string** | Leverage | [optional] [default to ""]
**MarkPx** | Pointer to **string** | Mark price | [optional] [default to ""]
**MgnRatio** | Pointer to **string** | Margin ratio | [optional] [default to ""]
**NotionalUsd** | Pointer to **string** | Notional in &#x60;USD&#x60; | [optional] [default to ""]
**PosSide** | Pointer to **string** | Position side  &#x60;long&#x60;  &#x60;short&#x60;  &#x60;net&#x60; | [optional] [default to ""]

## Methods

### NewCreateAccountPositionBuilderV5RespDataPositionsInner

`func NewCreateAccountPositionBuilderV5RespDataPositionsInner() *CreateAccountPositionBuilderV5RespDataPositionsInner`

NewCreateAccountPositionBuilderV5RespDataPositionsInner instantiates a new CreateAccountPositionBuilderV5RespDataPositionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountPositionBuilderV5RespDataPositionsInnerWithDefaults

`func NewCreateAccountPositionBuilderV5RespDataPositionsInnerWithDefaults() *CreateAccountPositionBuilderV5RespDataPositionsInner`

NewCreateAccountPositionBuilderV5RespDataPositionsInnerWithDefaults instantiates a new CreateAccountPositionBuilderV5RespDataPositionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetAvgPx

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetAvgPx() string`

GetAvgPx returns the AvgPx field if non-nil, zero value otherwise.

### GetAvgPxOk

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetAvgPxOk() (*string, bool)`

GetAvgPxOk returns a tuple with the AvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPx

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) SetAvgPx(v string)`

SetAvgPx sets AvgPx field to given value.

### HasAvgPx

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) HasAvgPx() bool`

HasAvgPx returns a boolean if a field has been set.

### GetFloatPnl

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetFloatPnl() string`

GetFloatPnl returns the FloatPnl field if non-nil, zero value otherwise.

### GetFloatPnlOk

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetFloatPnlOk() (*string, bool)`

GetFloatPnlOk returns a tuple with the FloatPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatPnl

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) SetFloatPnl(v string)`

SetFloatPnl sets FloatPnl field to given value.

### HasFloatPnl

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) HasFloatPnl() bool`

HasFloatPnl returns a boolean if a field has been set.

### GetImr

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetImr() string`

GetImr returns the Imr field if non-nil, zero value otherwise.

### GetImrOk

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetImrOk() (*string, bool)`

GetImrOk returns a tuple with the Imr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImr

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) SetImr(v string)`

SetImr sets Imr field to given value.

### HasImr

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) HasImr() bool`

HasImr returns a boolean if a field has been set.

### GetInstId

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetIsRealPos

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetIsRealPos() bool`

GetIsRealPos returns the IsRealPos field if non-nil, zero value otherwise.

### GetIsRealPosOk

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetIsRealPosOk() (*bool, bool)`

GetIsRealPosOk returns a tuple with the IsRealPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRealPos

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) SetIsRealPos(v bool)`

SetIsRealPos sets IsRealPos field to given value.

### HasIsRealPos

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) HasIsRealPos() bool`

HasIsRealPos returns a boolean if a field has been set.

### GetLever

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetMarkPx

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetMarkPx() string`

GetMarkPx returns the MarkPx field if non-nil, zero value otherwise.

### GetMarkPxOk

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetMarkPxOk() (*string, bool)`

GetMarkPxOk returns a tuple with the MarkPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPx

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) SetMarkPx(v string)`

SetMarkPx sets MarkPx field to given value.

### HasMarkPx

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) HasMarkPx() bool`

HasMarkPx returns a boolean if a field has been set.

### GetMgnRatio

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetMgnRatio() string`

GetMgnRatio returns the MgnRatio field if non-nil, zero value otherwise.

### GetMgnRatioOk

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetMgnRatioOk() (*string, bool)`

GetMgnRatioOk returns a tuple with the MgnRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnRatio

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) SetMgnRatio(v string)`

SetMgnRatio sets MgnRatio field to given value.

### HasMgnRatio

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) HasMgnRatio() bool`

HasMgnRatio returns a boolean if a field has been set.

### GetNotionalUsd

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetNotionalUsd() string`

GetNotionalUsd returns the NotionalUsd field if non-nil, zero value otherwise.

### GetNotionalUsdOk

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetNotionalUsdOk() (*string, bool)`

GetNotionalUsdOk returns a tuple with the NotionalUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalUsd

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) SetNotionalUsd(v string)`

SetNotionalUsd sets NotionalUsd field to given value.

### HasNotionalUsd

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) HasNotionalUsd() bool`

HasNotionalUsd returns a boolean if a field has been set.

### GetPosSide

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *CreateAccountPositionBuilderV5RespDataPositionsInner) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


