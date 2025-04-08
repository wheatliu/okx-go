# CreateTradeOrderPrecheckV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdjEq** | Pointer to **string** | Current adjusted / Effective equity in &#x60;USD&#x60; | [optional] [default to ""]
**AdjEqChg** | Pointer to **string** | After placing order, changed quantity of adjusted / Effective equity in &#x60;USD&#x60; | [optional] [default to ""]
**AvailBal** | Pointer to **string** | Current available balance in margin coin currency, only applicable to turn auto borrow off | [optional] [default to ""]
**AvailBalChg** | Pointer to **string** | After placing order, changed quantity of available balance after placing order, only applicable to turn auto borrow off | [optional] [default to ""]
**Imr** | Pointer to **string** | Current initial margin requirement in &#x60;USD&#x60; | [optional] [default to ""]
**ImrChg** | Pointer to **string** | After placing order, changed quantity of initial margin requirement in &#x60;USD&#x60; | [optional] [default to ""]
**Liab** | Pointer to **string** | Current liabilities of currency   For cross, it is cross liabilities  For isolated position, it is isolated liabilities | [optional] [default to ""]
**LiabChg** | Pointer to **string** | After placing order, changed quantity of liabilities   For cross, it is cross liabilities  For isolated position, it is isolated liabilities | [optional] [default to ""]
**LiabChgCcy** | Pointer to **string** | After placing order, the unit of changed liabilities quantity   only applicable cross and in auto borrow | [optional] [default to ""]
**LiqPx** | Pointer to **string** | Current estimated liquidation price | [optional] [default to ""]
**LiqPxDiff** | Pointer to **string** | After placing order, the distance between estimated liquidation price and mark price | [optional] [default to ""]
**LiqPxDiffRatio** | Pointer to **string** | After placing order, the distance rate between estimated liquidation price and mark price | [optional] [default to ""]
**MgnRatio** | Pointer to **string** | Current margin ratio in &#x60;USD&#x60; | [optional] [default to ""]
**MgnRatioChg** | Pointer to **string** | After placing order, changed quantity of margin ratio in &#x60;USD&#x60; | [optional] [default to ""]
**Mmr** | Pointer to **string** | Current Maintenance margin requirement in &#x60;USD&#x60; | [optional] [default to ""]
**MmrChg** | Pointer to **string** | After placing order, changed quantity of maintenance margin requirement in &#x60;USD&#x60; | [optional] [default to ""]
**PosBal** | Pointer to **string** | Current positive asset, only applicable to margin isolated position | [optional] [default to ""]
**PosBalChg** | Pointer to **string** | After placing order, positive asset of margin isolated, only applicable to margin isolated position | [optional] [default to ""]
**Type** | Pointer to **string** | Unit type of positive asset, only applicable to margin isolated position  &#x60;1&#x60;: it is both base currency before and after placing order    &#x60;2&#x60;: before plaing order, it is base currency. after placing order, it is quota currency.  &#x60;3&#x60;: before plaing order, it is quota currency. after placing order, it is base currency   &#x60;4&#x60;: it is both quota currency before and after placing order | [optional] [default to ""]

## Methods

### NewCreateTradeOrderPrecheckV5RespData

`func NewCreateTradeOrderPrecheckV5RespData() *CreateTradeOrderPrecheckV5RespData`

NewCreateTradeOrderPrecheckV5RespData instantiates a new CreateTradeOrderPrecheckV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeOrderPrecheckV5RespDataWithDefaults

`func NewCreateTradeOrderPrecheckV5RespDataWithDefaults() *CreateTradeOrderPrecheckV5RespData`

NewCreateTradeOrderPrecheckV5RespDataWithDefaults instantiates a new CreateTradeOrderPrecheckV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdjEq

`func (o *CreateTradeOrderPrecheckV5RespData) GetAdjEq() string`

GetAdjEq returns the AdjEq field if non-nil, zero value otherwise.

### GetAdjEqOk

`func (o *CreateTradeOrderPrecheckV5RespData) GetAdjEqOk() (*string, bool)`

GetAdjEqOk returns a tuple with the AdjEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjEq

`func (o *CreateTradeOrderPrecheckV5RespData) SetAdjEq(v string)`

SetAdjEq sets AdjEq field to given value.

### HasAdjEq

`func (o *CreateTradeOrderPrecheckV5RespData) HasAdjEq() bool`

HasAdjEq returns a boolean if a field has been set.

### GetAdjEqChg

`func (o *CreateTradeOrderPrecheckV5RespData) GetAdjEqChg() string`

GetAdjEqChg returns the AdjEqChg field if non-nil, zero value otherwise.

### GetAdjEqChgOk

`func (o *CreateTradeOrderPrecheckV5RespData) GetAdjEqChgOk() (*string, bool)`

GetAdjEqChgOk returns a tuple with the AdjEqChg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjEqChg

`func (o *CreateTradeOrderPrecheckV5RespData) SetAdjEqChg(v string)`

SetAdjEqChg sets AdjEqChg field to given value.

### HasAdjEqChg

`func (o *CreateTradeOrderPrecheckV5RespData) HasAdjEqChg() bool`

HasAdjEqChg returns a boolean if a field has been set.

### GetAvailBal

`func (o *CreateTradeOrderPrecheckV5RespData) GetAvailBal() string`

GetAvailBal returns the AvailBal field if non-nil, zero value otherwise.

### GetAvailBalOk

`func (o *CreateTradeOrderPrecheckV5RespData) GetAvailBalOk() (*string, bool)`

GetAvailBalOk returns a tuple with the AvailBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailBal

`func (o *CreateTradeOrderPrecheckV5RespData) SetAvailBal(v string)`

SetAvailBal sets AvailBal field to given value.

### HasAvailBal

`func (o *CreateTradeOrderPrecheckV5RespData) HasAvailBal() bool`

HasAvailBal returns a boolean if a field has been set.

### GetAvailBalChg

`func (o *CreateTradeOrderPrecheckV5RespData) GetAvailBalChg() string`

GetAvailBalChg returns the AvailBalChg field if non-nil, zero value otherwise.

### GetAvailBalChgOk

`func (o *CreateTradeOrderPrecheckV5RespData) GetAvailBalChgOk() (*string, bool)`

GetAvailBalChgOk returns a tuple with the AvailBalChg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailBalChg

`func (o *CreateTradeOrderPrecheckV5RespData) SetAvailBalChg(v string)`

SetAvailBalChg sets AvailBalChg field to given value.

### HasAvailBalChg

`func (o *CreateTradeOrderPrecheckV5RespData) HasAvailBalChg() bool`

HasAvailBalChg returns a boolean if a field has been set.

### GetImr

`func (o *CreateTradeOrderPrecheckV5RespData) GetImr() string`

GetImr returns the Imr field if non-nil, zero value otherwise.

### GetImrOk

`func (o *CreateTradeOrderPrecheckV5RespData) GetImrOk() (*string, bool)`

GetImrOk returns a tuple with the Imr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImr

`func (o *CreateTradeOrderPrecheckV5RespData) SetImr(v string)`

SetImr sets Imr field to given value.

### HasImr

`func (o *CreateTradeOrderPrecheckV5RespData) HasImr() bool`

HasImr returns a boolean if a field has been set.

### GetImrChg

`func (o *CreateTradeOrderPrecheckV5RespData) GetImrChg() string`

GetImrChg returns the ImrChg field if non-nil, zero value otherwise.

### GetImrChgOk

`func (o *CreateTradeOrderPrecheckV5RespData) GetImrChgOk() (*string, bool)`

GetImrChgOk returns a tuple with the ImrChg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImrChg

`func (o *CreateTradeOrderPrecheckV5RespData) SetImrChg(v string)`

SetImrChg sets ImrChg field to given value.

### HasImrChg

`func (o *CreateTradeOrderPrecheckV5RespData) HasImrChg() bool`

HasImrChg returns a boolean if a field has been set.

### GetLiab

`func (o *CreateTradeOrderPrecheckV5RespData) GetLiab() string`

GetLiab returns the Liab field if non-nil, zero value otherwise.

### GetLiabOk

`func (o *CreateTradeOrderPrecheckV5RespData) GetLiabOk() (*string, bool)`

GetLiabOk returns a tuple with the Liab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiab

`func (o *CreateTradeOrderPrecheckV5RespData) SetLiab(v string)`

SetLiab sets Liab field to given value.

### HasLiab

`func (o *CreateTradeOrderPrecheckV5RespData) HasLiab() bool`

HasLiab returns a boolean if a field has been set.

### GetLiabChg

`func (o *CreateTradeOrderPrecheckV5RespData) GetLiabChg() string`

GetLiabChg returns the LiabChg field if non-nil, zero value otherwise.

### GetLiabChgOk

`func (o *CreateTradeOrderPrecheckV5RespData) GetLiabChgOk() (*string, bool)`

GetLiabChgOk returns a tuple with the LiabChg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiabChg

`func (o *CreateTradeOrderPrecheckV5RespData) SetLiabChg(v string)`

SetLiabChg sets LiabChg field to given value.

### HasLiabChg

`func (o *CreateTradeOrderPrecheckV5RespData) HasLiabChg() bool`

HasLiabChg returns a boolean if a field has been set.

### GetLiabChgCcy

`func (o *CreateTradeOrderPrecheckV5RespData) GetLiabChgCcy() string`

GetLiabChgCcy returns the LiabChgCcy field if non-nil, zero value otherwise.

### GetLiabChgCcyOk

`func (o *CreateTradeOrderPrecheckV5RespData) GetLiabChgCcyOk() (*string, bool)`

GetLiabChgCcyOk returns a tuple with the LiabChgCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiabChgCcy

`func (o *CreateTradeOrderPrecheckV5RespData) SetLiabChgCcy(v string)`

SetLiabChgCcy sets LiabChgCcy field to given value.

### HasLiabChgCcy

`func (o *CreateTradeOrderPrecheckV5RespData) HasLiabChgCcy() bool`

HasLiabChgCcy returns a boolean if a field has been set.

### GetLiqPx

`func (o *CreateTradeOrderPrecheckV5RespData) GetLiqPx() string`

GetLiqPx returns the LiqPx field if non-nil, zero value otherwise.

### GetLiqPxOk

`func (o *CreateTradeOrderPrecheckV5RespData) GetLiqPxOk() (*string, bool)`

GetLiqPxOk returns a tuple with the LiqPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiqPx

`func (o *CreateTradeOrderPrecheckV5RespData) SetLiqPx(v string)`

SetLiqPx sets LiqPx field to given value.

### HasLiqPx

`func (o *CreateTradeOrderPrecheckV5RespData) HasLiqPx() bool`

HasLiqPx returns a boolean if a field has been set.

### GetLiqPxDiff

`func (o *CreateTradeOrderPrecheckV5RespData) GetLiqPxDiff() string`

GetLiqPxDiff returns the LiqPxDiff field if non-nil, zero value otherwise.

### GetLiqPxDiffOk

`func (o *CreateTradeOrderPrecheckV5RespData) GetLiqPxDiffOk() (*string, bool)`

GetLiqPxDiffOk returns a tuple with the LiqPxDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiqPxDiff

`func (o *CreateTradeOrderPrecheckV5RespData) SetLiqPxDiff(v string)`

SetLiqPxDiff sets LiqPxDiff field to given value.

### HasLiqPxDiff

`func (o *CreateTradeOrderPrecheckV5RespData) HasLiqPxDiff() bool`

HasLiqPxDiff returns a boolean if a field has been set.

### GetLiqPxDiffRatio

`func (o *CreateTradeOrderPrecheckV5RespData) GetLiqPxDiffRatio() string`

GetLiqPxDiffRatio returns the LiqPxDiffRatio field if non-nil, zero value otherwise.

### GetLiqPxDiffRatioOk

`func (o *CreateTradeOrderPrecheckV5RespData) GetLiqPxDiffRatioOk() (*string, bool)`

GetLiqPxDiffRatioOk returns a tuple with the LiqPxDiffRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiqPxDiffRatio

`func (o *CreateTradeOrderPrecheckV5RespData) SetLiqPxDiffRatio(v string)`

SetLiqPxDiffRatio sets LiqPxDiffRatio field to given value.

### HasLiqPxDiffRatio

`func (o *CreateTradeOrderPrecheckV5RespData) HasLiqPxDiffRatio() bool`

HasLiqPxDiffRatio returns a boolean if a field has been set.

### GetMgnRatio

`func (o *CreateTradeOrderPrecheckV5RespData) GetMgnRatio() string`

GetMgnRatio returns the MgnRatio field if non-nil, zero value otherwise.

### GetMgnRatioOk

`func (o *CreateTradeOrderPrecheckV5RespData) GetMgnRatioOk() (*string, bool)`

GetMgnRatioOk returns a tuple with the MgnRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnRatio

`func (o *CreateTradeOrderPrecheckV5RespData) SetMgnRatio(v string)`

SetMgnRatio sets MgnRatio field to given value.

### HasMgnRatio

`func (o *CreateTradeOrderPrecheckV5RespData) HasMgnRatio() bool`

HasMgnRatio returns a boolean if a field has been set.

### GetMgnRatioChg

`func (o *CreateTradeOrderPrecheckV5RespData) GetMgnRatioChg() string`

GetMgnRatioChg returns the MgnRatioChg field if non-nil, zero value otherwise.

### GetMgnRatioChgOk

`func (o *CreateTradeOrderPrecheckV5RespData) GetMgnRatioChgOk() (*string, bool)`

GetMgnRatioChgOk returns a tuple with the MgnRatioChg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnRatioChg

`func (o *CreateTradeOrderPrecheckV5RespData) SetMgnRatioChg(v string)`

SetMgnRatioChg sets MgnRatioChg field to given value.

### HasMgnRatioChg

`func (o *CreateTradeOrderPrecheckV5RespData) HasMgnRatioChg() bool`

HasMgnRatioChg returns a boolean if a field has been set.

### GetMmr

`func (o *CreateTradeOrderPrecheckV5RespData) GetMmr() string`

GetMmr returns the Mmr field if non-nil, zero value otherwise.

### GetMmrOk

`func (o *CreateTradeOrderPrecheckV5RespData) GetMmrOk() (*string, bool)`

GetMmrOk returns a tuple with the Mmr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmr

`func (o *CreateTradeOrderPrecheckV5RespData) SetMmr(v string)`

SetMmr sets Mmr field to given value.

### HasMmr

`func (o *CreateTradeOrderPrecheckV5RespData) HasMmr() bool`

HasMmr returns a boolean if a field has been set.

### GetMmrChg

`func (o *CreateTradeOrderPrecheckV5RespData) GetMmrChg() string`

GetMmrChg returns the MmrChg field if non-nil, zero value otherwise.

### GetMmrChgOk

`func (o *CreateTradeOrderPrecheckV5RespData) GetMmrChgOk() (*string, bool)`

GetMmrChgOk returns a tuple with the MmrChg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmrChg

`func (o *CreateTradeOrderPrecheckV5RespData) SetMmrChg(v string)`

SetMmrChg sets MmrChg field to given value.

### HasMmrChg

`func (o *CreateTradeOrderPrecheckV5RespData) HasMmrChg() bool`

HasMmrChg returns a boolean if a field has been set.

### GetPosBal

`func (o *CreateTradeOrderPrecheckV5RespData) GetPosBal() string`

GetPosBal returns the PosBal field if non-nil, zero value otherwise.

### GetPosBalOk

`func (o *CreateTradeOrderPrecheckV5RespData) GetPosBalOk() (*string, bool)`

GetPosBalOk returns a tuple with the PosBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosBal

`func (o *CreateTradeOrderPrecheckV5RespData) SetPosBal(v string)`

SetPosBal sets PosBal field to given value.

### HasPosBal

`func (o *CreateTradeOrderPrecheckV5RespData) HasPosBal() bool`

HasPosBal returns a boolean if a field has been set.

### GetPosBalChg

`func (o *CreateTradeOrderPrecheckV5RespData) GetPosBalChg() string`

GetPosBalChg returns the PosBalChg field if non-nil, zero value otherwise.

### GetPosBalChgOk

`func (o *CreateTradeOrderPrecheckV5RespData) GetPosBalChgOk() (*string, bool)`

GetPosBalChgOk returns a tuple with the PosBalChg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosBalChg

`func (o *CreateTradeOrderPrecheckV5RespData) SetPosBalChg(v string)`

SetPosBalChg sets PosBalChg field to given value.

### HasPosBalChg

`func (o *CreateTradeOrderPrecheckV5RespData) HasPosBalChg() bool`

HasPosBalChg returns a boolean if a field has been set.

### GetType

`func (o *CreateTradeOrderPrecheckV5RespData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTradeOrderPrecheckV5RespData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTradeOrderPrecheckV5RespData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateTradeOrderPrecheckV5RespData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


