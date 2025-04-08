# CreateAccountPositionBuilderV5RespDataInnerPositionsInner

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

### NewCreateAccountPositionBuilderV5RespDataInnerPositionsInner

`func NewCreateAccountPositionBuilderV5RespDataInnerPositionsInner() *CreateAccountPositionBuilderV5RespDataInnerPositionsInner`

NewCreateAccountPositionBuilderV5RespDataInnerPositionsInner instantiates a new CreateAccountPositionBuilderV5RespDataInnerPositionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountPositionBuilderV5RespDataInnerPositionsInnerWithDefaults

`func NewCreateAccountPositionBuilderV5RespDataInnerPositionsInnerWithDefaults() *CreateAccountPositionBuilderV5RespDataInnerPositionsInner`

NewCreateAccountPositionBuilderV5RespDataInnerPositionsInnerWithDefaults instantiates a new CreateAccountPositionBuilderV5RespDataInnerPositionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetAvgPx

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetAvgPx() string`

GetAvgPx returns the AvgPx field if non-nil, zero value otherwise.

### GetAvgPxOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetAvgPxOk() (*string, bool)`

GetAvgPxOk returns a tuple with the AvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPx

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetAvgPx(v string)`

SetAvgPx sets AvgPx field to given value.

### HasAvgPx

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasAvgPx() bool`

HasAvgPx returns a boolean if a field has been set.

### GetFloatPnl

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetFloatPnl() string`

GetFloatPnl returns the FloatPnl field if non-nil, zero value otherwise.

### GetFloatPnlOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetFloatPnlOk() (*string, bool)`

GetFloatPnlOk returns a tuple with the FloatPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatPnl

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetFloatPnl(v string)`

SetFloatPnl sets FloatPnl field to given value.

### HasFloatPnl

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasFloatPnl() bool`

HasFloatPnl returns a boolean if a field has been set.

### GetImr

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetImr() string`

GetImr returns the Imr field if non-nil, zero value otherwise.

### GetImrOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetImrOk() (*string, bool)`

GetImrOk returns a tuple with the Imr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImr

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetImr(v string)`

SetImr sets Imr field to given value.

### HasImr

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasImr() bool`

HasImr returns a boolean if a field has been set.

### GetInstId

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetIsRealPos

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetIsRealPos() bool`

GetIsRealPos returns the IsRealPos field if non-nil, zero value otherwise.

### GetIsRealPosOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetIsRealPosOk() (*bool, bool)`

GetIsRealPosOk returns a tuple with the IsRealPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRealPos

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetIsRealPos(v bool)`

SetIsRealPos sets IsRealPos field to given value.

### HasIsRealPos

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasIsRealPos() bool`

HasIsRealPos returns a boolean if a field has been set.

### GetLever

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetMarkPx

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetMarkPx() string`

GetMarkPx returns the MarkPx field if non-nil, zero value otherwise.

### GetMarkPxOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetMarkPxOk() (*string, bool)`

GetMarkPxOk returns a tuple with the MarkPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPx

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetMarkPx(v string)`

SetMarkPx sets MarkPx field to given value.

### HasMarkPx

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasMarkPx() bool`

HasMarkPx returns a boolean if a field has been set.

### GetMgnRatio

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetMgnRatio() string`

GetMgnRatio returns the MgnRatio field if non-nil, zero value otherwise.

### GetMgnRatioOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetMgnRatioOk() (*string, bool)`

GetMgnRatioOk returns a tuple with the MgnRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnRatio

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetMgnRatio(v string)`

SetMgnRatio sets MgnRatio field to given value.

### HasMgnRatio

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasMgnRatio() bool`

HasMgnRatio returns a boolean if a field has been set.

### GetNotionalUsd

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetNotionalUsd() string`

GetNotionalUsd returns the NotionalUsd field if non-nil, zero value otherwise.

### GetNotionalUsdOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetNotionalUsdOk() (*string, bool)`

GetNotionalUsdOk returns a tuple with the NotionalUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalUsd

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetNotionalUsd(v string)`

SetNotionalUsd sets NotionalUsd field to given value.

### HasNotionalUsd

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasNotionalUsd() bool`

HasNotionalUsd returns a boolean if a field has been set.

### GetPosSide

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


