# GetTradeOrdersAlgoPendingV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePx** | Pointer to **string** | Active price  Only applicable to &#x60;move_order_stop&#x60; order | [optional] [default to ""]
**ActualPx** | Pointer to **string** | Actual order price | [optional] [default to ""]
**ActualSide** | Pointer to **string** | Actual trigger side  &#x60;tp&#x60;: take profit &#x60;sl&#x60;: stop loss  Only applicable to oco order and conditional order | [optional] [default to ""]
**ActualSz** | Pointer to **string** | Actual order quantity | [optional] [default to ""]
**AlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID | [optional] [default to ""]
**AlgoId** | Pointer to **string** | Algo ID | [optional] [default to ""]
**AmendPxOnTriggerType** | Pointer to **string** | Whether to enable Cost-price SL. Only applicable to SL order of split TPs.   &#x60;0&#x60;: disable, the default value   &#x60;1&#x60;: Enable | [optional] [default to ""]
**AttachAlgoOrds** | Pointer to [**[]GetTradeOrderAlgoV5RespDataAttachAlgoOrdsInner**](GetTradeOrderAlgoV5RespDataAttachAlgoOrdsInner.md) | Attached SL/TP orders info  Applicable to &#x60;Spot and futures mode/Multi-currency margin/Portfolio margin&#x60; | [optional] 
**CTime** | Pointer to **string** | Creation time Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**CallbackRatio** | Pointer to **string** | Callback price ratio  Only applicable to &#x60;move_order_stop&#x60; order | [optional] [default to ""]
**CallbackSpread** | Pointer to **string** | Callback price variance  Only applicable to &#x60;move_order_stop&#x60; order | [optional] [default to ""]
**Ccy** | Pointer to **string** | Margin currency   Applicable to all &#x60;isolated&#x60; &#x60;MARGIN&#x60; orders and &#x60;cross&#x60; &#x60;MARGIN&#x60; orders in &#x60;Spot and futures mode&#x60;. | [optional] [default to ""]
**ChaseType** | Pointer to **string** | Chase type. Only applicable to &#x60;chase&#x60; order. | [optional] [default to ""]
**ChaseVal** | Pointer to **string** | Chase value. Only applicable to &#x60;chase&#x60; order. | [optional] [default to ""]
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client | [optional] [default to ""]
**CloseFraction** | Pointer to **string** | Fraction of position to be closed when the algo order is triggered | [optional] [default to ""]
**FailCode** | Pointer to **string** | It represents that the reason that algo order fails to trigger. There will be value when the state is &#x60;order_failed&#x60;, e.g. 51008;  For this endpoint, it always is \&quot;\&quot;. | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**IsTradeBorrowMode** | Pointer to **string** | Whether borrowing currency automatically   true   false  Only applicable to &#x60;trigger order&#x60;, &#x60;trailing order&#x60; and &#x60;twap order&#x60; | [optional] [default to ""]
**Last** | Pointer to **string** | Last filled price while placing | [optional] [default to ""]
**Lever** | Pointer to **string** | Leverage, from &#x60;0.01&#x60; to &#x60;125&#x60;.   Only applicable to &#x60;MARGIN/FUTURES/SWAP&#x60; | [optional] [default to ""]
**LinkedOrd** | Pointer to [**GetTradeOrderAlgoV5RespDataLinkedOrd**](GetTradeOrderAlgoV5RespDataLinkedOrd.md) |  | [optional] 
**MaxChaseType** | Pointer to **string** | Maximum chase type. Only applicable to &#x60;chase&#x60; order. | [optional] [default to ""]
**MaxChaseVal** | Pointer to **string** | Maximum chase value. Only applicable to &#x60;chase&#x60; order. | [optional] [default to ""]
**MoveTriggerPx** | Pointer to **string** | Trigger price  Only applicable to &#x60;move_order_stop&#x60; order | [optional] [default to ""]
**OrdId** | Pointer to **string** | Latest order ID. It will be deprecated soon | [optional] [default to ""]
**OrdIdList** | Pointer to **[]string** | Order ID list. There will be multiple order IDs when there is TP/SL splitting order. | [optional] 
**OrdPx** | Pointer to **string** | Order price for the trigger order | [optional] [default to ""]
**OrdType** | Pointer to **string** | Order type | [optional] [default to ""]
**PosSide** | Pointer to **string** | Position side | [optional] [default to ""]
**PxLimit** | Pointer to **string** | Price Limit   Only applicable to &#x60;iceberg&#x60; order or &#x60;twap&#x60; order | [optional] [default to ""]
**PxSpread** | Pointer to **string** | Price variance   Only applicable to &#x60;iceberg&#x60; order or &#x60;twap&#x60; order | [optional] [default to ""]
**PxVar** | Pointer to **string** | Price ratio   Only applicable to &#x60;iceberg&#x60; order or &#x60;twap&#x60; order | [optional] [default to ""]
**QuickMgnType** | Pointer to **string** | Quick Margin type, Only applicable to Quick Margin Mode of isolated margin  &#x60;manual&#x60;, &#x60;auto_borrow&#x60;, &#x60;auto_repay&#x60; | [optional] [default to ""]
**ReduceOnly** | Pointer to **string** | Whether the order can only reduce the position size. Valid options: true or false. | [optional] [default to ""]
**Side** | Pointer to **string** | Order side | [optional] [default to ""]
**SlOrdPx** | Pointer to **string** | Stop-loss order price | [optional] [default to ""]
**SlTriggerPx** | Pointer to **string** | Stop-loss trigger price | [optional] [default to ""]
**SlTriggerPxType** | Pointer to **string** | Stop-loss trigger price type.   &#x60;last&#x60;: last price  &#x60;index&#x60;: index price  &#x60;mark&#x60;: mark price | [optional] [default to ""]
**State** | Pointer to **string** | State  &#x60;live&#x60;  &#x60;pause&#x60; | [optional] [default to ""]
**Sz** | Pointer to **string** | Quantity to buy or sell | [optional] [default to ""]
**SzLimit** | Pointer to **string** | Average amount   Only applicable to &#x60;iceberg&#x60; order or &#x60;twap&#x60; order | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**TdMode** | Pointer to **string** | Trade mode | [optional] [default to ""]
**TgtCcy** | Pointer to **string** | Order quantity unit setting for &#x60;sz&#x60;  &#x60;base_ccy&#x60;: Base currency ,&#x60;quote_ccy&#x60;: Quote currency    Only applicable to &#x60;SPOT&#x60; traded with Market order | [optional] [default to ""]
**TimeInterval** | Pointer to **string** | Time interval   Only applicable to &#x60;twap&#x60; order | [optional] [default to ""]
**TpOrdPx** | Pointer to **string** | Take-profit order price | [optional] [default to ""]
**TpTriggerPx** | Pointer to **string** | Take-profit trigger price | [optional] [default to ""]
**TpTriggerPxType** | Pointer to **string** | Take-profit trigger price type.   &#x60;last&#x60;: last price  &#x60;index&#x60;: index price  &#x60;mark&#x60;: mark price | [optional] [default to ""]
**TriggerPx** | Pointer to **string** | Trigger price | [optional] [default to ""]
**TriggerPxType** | Pointer to **string** | Trigger price type.   &#x60;last&#x60;: last price  &#x60;index&#x60;: index price  &#x60;mark&#x60;: mark price | [optional] [default to ""]
**TriggerTime** | Pointer to **string** | Trigger time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**UTime** | Pointer to **string** | Order updated time, Unix timestamp format in milliseconds, e.g. 1597026383085 | [optional] [default to ""]

## Methods

### NewGetTradeOrdersAlgoPendingV5RespData

`func NewGetTradeOrdersAlgoPendingV5RespData() *GetTradeOrdersAlgoPendingV5RespData`

NewGetTradeOrdersAlgoPendingV5RespData instantiates a new GetTradeOrdersAlgoPendingV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradeOrdersAlgoPendingV5RespDataWithDefaults

`func NewGetTradeOrdersAlgoPendingV5RespDataWithDefaults() *GetTradeOrdersAlgoPendingV5RespData`

NewGetTradeOrdersAlgoPendingV5RespDataWithDefaults instantiates a new GetTradeOrdersAlgoPendingV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePx

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetActivePx() string`

GetActivePx returns the ActivePx field if non-nil, zero value otherwise.

### GetActivePxOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetActivePxOk() (*string, bool)`

GetActivePxOk returns a tuple with the ActivePx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePx

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetActivePx(v string)`

SetActivePx sets ActivePx field to given value.

### HasActivePx

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasActivePx() bool`

HasActivePx returns a boolean if a field has been set.

### GetActualPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetActualPx() string`

GetActualPx returns the ActualPx field if non-nil, zero value otherwise.

### GetActualPxOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetActualPxOk() (*string, bool)`

GetActualPxOk returns a tuple with the ActualPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetActualPx(v string)`

SetActualPx sets ActualPx field to given value.

### HasActualPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasActualPx() bool`

HasActualPx returns a boolean if a field has been set.

### GetActualSide

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetActualSide() string`

GetActualSide returns the ActualSide field if non-nil, zero value otherwise.

### GetActualSideOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetActualSideOk() (*string, bool)`

GetActualSideOk returns a tuple with the ActualSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualSide

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetActualSide(v string)`

SetActualSide sets ActualSide field to given value.

### HasActualSide

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasActualSide() bool`

HasActualSide returns a boolean if a field has been set.

### GetActualSz

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetActualSz() string`

GetActualSz returns the ActualSz field if non-nil, zero value otherwise.

### GetActualSzOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetActualSzOk() (*string, bool)`

GetActualSzOk returns a tuple with the ActualSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualSz

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetActualSz(v string)`

SetActualSz sets ActualSz field to given value.

### HasActualSz

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasActualSz() bool`

HasActualSz returns a boolean if a field has been set.

### GetAlgoClOrdId

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAlgoId

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetAmendPxOnTriggerType

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetAmendPxOnTriggerType() string`

GetAmendPxOnTriggerType returns the AmendPxOnTriggerType field if non-nil, zero value otherwise.

### GetAmendPxOnTriggerTypeOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetAmendPxOnTriggerTypeOk() (*string, bool)`

GetAmendPxOnTriggerTypeOk returns a tuple with the AmendPxOnTriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendPxOnTriggerType

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetAmendPxOnTriggerType(v string)`

SetAmendPxOnTriggerType sets AmendPxOnTriggerType field to given value.

### HasAmendPxOnTriggerType

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasAmendPxOnTriggerType() bool`

HasAmendPxOnTriggerType returns a boolean if a field has been set.

### GetAttachAlgoOrds

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetAttachAlgoOrds() []GetTradeOrderAlgoV5RespDataAttachAlgoOrdsInner`

GetAttachAlgoOrds returns the AttachAlgoOrds field if non-nil, zero value otherwise.

### GetAttachAlgoOrdsOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetAttachAlgoOrdsOk() (*[]GetTradeOrderAlgoV5RespDataAttachAlgoOrdsInner, bool)`

GetAttachAlgoOrdsOk returns a tuple with the AttachAlgoOrds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachAlgoOrds

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetAttachAlgoOrds(v []GetTradeOrderAlgoV5RespDataAttachAlgoOrdsInner)`

SetAttachAlgoOrds sets AttachAlgoOrds field to given value.

### HasAttachAlgoOrds

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasAttachAlgoOrds() bool`

HasAttachAlgoOrds returns a boolean if a field has been set.

### GetCTime

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCallbackRatio

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetCallbackRatio() string`

GetCallbackRatio returns the CallbackRatio field if non-nil, zero value otherwise.

### GetCallbackRatioOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetCallbackRatioOk() (*string, bool)`

GetCallbackRatioOk returns a tuple with the CallbackRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackRatio

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetCallbackRatio(v string)`

SetCallbackRatio sets CallbackRatio field to given value.

### HasCallbackRatio

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasCallbackRatio() bool`

HasCallbackRatio returns a boolean if a field has been set.

### GetCallbackSpread

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetCallbackSpread() string`

GetCallbackSpread returns the CallbackSpread field if non-nil, zero value otherwise.

### GetCallbackSpreadOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetCallbackSpreadOk() (*string, bool)`

GetCallbackSpreadOk returns a tuple with the CallbackSpread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSpread

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetCallbackSpread(v string)`

SetCallbackSpread sets CallbackSpread field to given value.

### HasCallbackSpread

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasCallbackSpread() bool`

HasCallbackSpread returns a boolean if a field has been set.

### GetCcy

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetChaseType

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetChaseType() string`

GetChaseType returns the ChaseType field if non-nil, zero value otherwise.

### GetChaseTypeOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetChaseTypeOk() (*string, bool)`

GetChaseTypeOk returns a tuple with the ChaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChaseType

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetChaseType(v string)`

SetChaseType sets ChaseType field to given value.

### HasChaseType

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasChaseType() bool`

HasChaseType returns a boolean if a field has been set.

### GetChaseVal

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetChaseVal() string`

GetChaseVal returns the ChaseVal field if non-nil, zero value otherwise.

### GetChaseValOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetChaseValOk() (*string, bool)`

GetChaseValOk returns a tuple with the ChaseVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChaseVal

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetChaseVal(v string)`

SetChaseVal sets ChaseVal field to given value.

### HasChaseVal

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasChaseVal() bool`

HasChaseVal returns a boolean if a field has been set.

### GetClOrdId

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetCloseFraction

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetCloseFraction() string`

GetCloseFraction returns the CloseFraction field if non-nil, zero value otherwise.

### GetCloseFractionOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetCloseFractionOk() (*string, bool)`

GetCloseFractionOk returns a tuple with the CloseFraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseFraction

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetCloseFraction(v string)`

SetCloseFraction sets CloseFraction field to given value.

### HasCloseFraction

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasCloseFraction() bool`

HasCloseFraction returns a boolean if a field has been set.

### GetFailCode

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetFailCode() string`

GetFailCode returns the FailCode field if non-nil, zero value otherwise.

### GetFailCodeOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetFailCodeOk() (*string, bool)`

GetFailCodeOk returns a tuple with the FailCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailCode

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetFailCode(v string)`

SetFailCode sets FailCode field to given value.

### HasFailCode

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasFailCode() bool`

HasFailCode returns a boolean if a field has been set.

### GetInstId

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetIsTradeBorrowMode

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetIsTradeBorrowMode() string`

GetIsTradeBorrowMode returns the IsTradeBorrowMode field if non-nil, zero value otherwise.

### GetIsTradeBorrowModeOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetIsTradeBorrowModeOk() (*string, bool)`

GetIsTradeBorrowModeOk returns a tuple with the IsTradeBorrowMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTradeBorrowMode

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetIsTradeBorrowMode(v string)`

SetIsTradeBorrowMode sets IsTradeBorrowMode field to given value.

### HasIsTradeBorrowMode

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasIsTradeBorrowMode() bool`

HasIsTradeBorrowMode returns a boolean if a field has been set.

### GetLast

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetLever

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetLinkedOrd

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetLinkedOrd() GetTradeOrderAlgoV5RespDataLinkedOrd`

GetLinkedOrd returns the LinkedOrd field if non-nil, zero value otherwise.

### GetLinkedOrdOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetLinkedOrdOk() (*GetTradeOrderAlgoV5RespDataLinkedOrd, bool)`

GetLinkedOrdOk returns a tuple with the LinkedOrd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedOrd

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetLinkedOrd(v GetTradeOrderAlgoV5RespDataLinkedOrd)`

SetLinkedOrd sets LinkedOrd field to given value.

### HasLinkedOrd

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasLinkedOrd() bool`

HasLinkedOrd returns a boolean if a field has been set.

### GetMaxChaseType

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetMaxChaseType() string`

GetMaxChaseType returns the MaxChaseType field if non-nil, zero value otherwise.

### GetMaxChaseTypeOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetMaxChaseTypeOk() (*string, bool)`

GetMaxChaseTypeOk returns a tuple with the MaxChaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxChaseType

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetMaxChaseType(v string)`

SetMaxChaseType sets MaxChaseType field to given value.

### HasMaxChaseType

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasMaxChaseType() bool`

HasMaxChaseType returns a boolean if a field has been set.

### GetMaxChaseVal

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetMaxChaseVal() string`

GetMaxChaseVal returns the MaxChaseVal field if non-nil, zero value otherwise.

### GetMaxChaseValOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetMaxChaseValOk() (*string, bool)`

GetMaxChaseValOk returns a tuple with the MaxChaseVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxChaseVal

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetMaxChaseVal(v string)`

SetMaxChaseVal sets MaxChaseVal field to given value.

### HasMaxChaseVal

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasMaxChaseVal() bool`

HasMaxChaseVal returns a boolean if a field has been set.

### GetMoveTriggerPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetMoveTriggerPx() string`

GetMoveTriggerPx returns the MoveTriggerPx field if non-nil, zero value otherwise.

### GetMoveTriggerPxOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetMoveTriggerPxOk() (*string, bool)`

GetMoveTriggerPxOk returns a tuple with the MoveTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveTriggerPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetMoveTriggerPx(v string)`

SetMoveTriggerPx sets MoveTriggerPx field to given value.

### HasMoveTriggerPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasMoveTriggerPx() bool`

HasMoveTriggerPx returns a boolean if a field has been set.

### GetOrdId

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetOrdIdList

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetOrdIdList() []string`

GetOrdIdList returns the OrdIdList field if non-nil, zero value otherwise.

### GetOrdIdListOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetOrdIdListOk() (*[]string, bool)`

GetOrdIdListOk returns a tuple with the OrdIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdIdList

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetOrdIdList(v []string)`

SetOrdIdList sets OrdIdList field to given value.

### HasOrdIdList

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasOrdIdList() bool`

HasOrdIdList returns a boolean if a field has been set.

### GetOrdPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetOrdPx() string`

GetOrdPx returns the OrdPx field if non-nil, zero value otherwise.

### GetOrdPxOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetOrdPxOk() (*string, bool)`

GetOrdPxOk returns a tuple with the OrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetOrdPx(v string)`

SetOrdPx sets OrdPx field to given value.

### HasOrdPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasOrdPx() bool`

HasOrdPx returns a boolean if a field has been set.

### GetOrdType

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetOrdType() string`

GetOrdType returns the OrdType field if non-nil, zero value otherwise.

### GetOrdTypeOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetOrdTypeOk() (*string, bool)`

GetOrdTypeOk returns a tuple with the OrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdType

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetOrdType(v string)`

SetOrdType sets OrdType field to given value.

### HasOrdType

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasOrdType() bool`

HasOrdType returns a boolean if a field has been set.

### GetPosSide

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetPxLimit

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetPxLimit() string`

GetPxLimit returns the PxLimit field if non-nil, zero value otherwise.

### GetPxLimitOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetPxLimitOk() (*string, bool)`

GetPxLimitOk returns a tuple with the PxLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxLimit

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetPxLimit(v string)`

SetPxLimit sets PxLimit field to given value.

### HasPxLimit

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasPxLimit() bool`

HasPxLimit returns a boolean if a field has been set.

### GetPxSpread

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetPxSpread() string`

GetPxSpread returns the PxSpread field if non-nil, zero value otherwise.

### GetPxSpreadOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetPxSpreadOk() (*string, bool)`

GetPxSpreadOk returns a tuple with the PxSpread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxSpread

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetPxSpread(v string)`

SetPxSpread sets PxSpread field to given value.

### HasPxSpread

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasPxSpread() bool`

HasPxSpread returns a boolean if a field has been set.

### GetPxVar

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetPxVar() string`

GetPxVar returns the PxVar field if non-nil, zero value otherwise.

### GetPxVarOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetPxVarOk() (*string, bool)`

GetPxVarOk returns a tuple with the PxVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxVar

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetPxVar(v string)`

SetPxVar sets PxVar field to given value.

### HasPxVar

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasPxVar() bool`

HasPxVar returns a boolean if a field has been set.

### GetQuickMgnType

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetQuickMgnType() string`

GetQuickMgnType returns the QuickMgnType field if non-nil, zero value otherwise.

### GetQuickMgnTypeOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetQuickMgnTypeOk() (*string, bool)`

GetQuickMgnTypeOk returns a tuple with the QuickMgnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickMgnType

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetQuickMgnType(v string)`

SetQuickMgnType sets QuickMgnType field to given value.

### HasQuickMgnType

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasQuickMgnType() bool`

HasQuickMgnType returns a boolean if a field has been set.

### GetReduceOnly

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetReduceOnly() string`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetReduceOnlyOk() (*string, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetReduceOnly(v string)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSide

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSlOrdPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetSlOrdPx() string`

GetSlOrdPx returns the SlOrdPx field if non-nil, zero value otherwise.

### GetSlOrdPxOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetSlOrdPxOk() (*string, bool)`

GetSlOrdPxOk returns a tuple with the SlOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlOrdPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetSlOrdPx(v string)`

SetSlOrdPx sets SlOrdPx field to given value.

### HasSlOrdPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasSlOrdPx() bool`

HasSlOrdPx returns a boolean if a field has been set.

### GetSlTriggerPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetSlTriggerPx() string`

GetSlTriggerPx returns the SlTriggerPx field if non-nil, zero value otherwise.

### GetSlTriggerPxOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetSlTriggerPxOk() (*string, bool)`

GetSlTriggerPxOk returns a tuple with the SlTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetSlTriggerPx(v string)`

SetSlTriggerPx sets SlTriggerPx field to given value.

### HasSlTriggerPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasSlTriggerPx() bool`

HasSlTriggerPx returns a boolean if a field has been set.

### GetSlTriggerPxType

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetSlTriggerPxType() string`

GetSlTriggerPxType returns the SlTriggerPxType field if non-nil, zero value otherwise.

### GetSlTriggerPxTypeOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetSlTriggerPxTypeOk() (*string, bool)`

GetSlTriggerPxTypeOk returns a tuple with the SlTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPxType

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetSlTriggerPxType(v string)`

SetSlTriggerPxType sets SlTriggerPxType field to given value.

### HasSlTriggerPxType

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasSlTriggerPxType() bool`

HasSlTriggerPxType returns a boolean if a field has been set.

### GetState

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSz

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetSzLimit

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetSzLimit() string`

GetSzLimit returns the SzLimit field if non-nil, zero value otherwise.

### GetSzLimitOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetSzLimitOk() (*string, bool)`

GetSzLimitOk returns a tuple with the SzLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSzLimit

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetSzLimit(v string)`

SetSzLimit sets SzLimit field to given value.

### HasSzLimit

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasSzLimit() bool`

HasSzLimit returns a boolean if a field has been set.

### GetTag

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTdMode

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetTdMode() string`

GetTdMode returns the TdMode field if non-nil, zero value otherwise.

### GetTdModeOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetTdModeOk() (*string, bool)`

GetTdModeOk returns a tuple with the TdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdMode

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetTdMode(v string)`

SetTdMode sets TdMode field to given value.

### HasTdMode

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasTdMode() bool`

HasTdMode returns a boolean if a field has been set.

### GetTgtCcy

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetTgtCcy() string`

GetTgtCcy returns the TgtCcy field if non-nil, zero value otherwise.

### GetTgtCcyOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetTgtCcyOk() (*string, bool)`

GetTgtCcyOk returns a tuple with the TgtCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCcy

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetTgtCcy(v string)`

SetTgtCcy sets TgtCcy field to given value.

### HasTgtCcy

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasTgtCcy() bool`

HasTgtCcy returns a boolean if a field has been set.

### GetTimeInterval

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetTimeInterval() string`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetTimeIntervalOk() (*string, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetTimeInterval(v string)`

SetTimeInterval sets TimeInterval field to given value.

### HasTimeInterval

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasTimeInterval() bool`

HasTimeInterval returns a boolean if a field has been set.

### GetTpOrdPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetTpOrdPx() string`

GetTpOrdPx returns the TpOrdPx field if non-nil, zero value otherwise.

### GetTpOrdPxOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetTpOrdPxOk() (*string, bool)`

GetTpOrdPxOk returns a tuple with the TpOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpOrdPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetTpOrdPx(v string)`

SetTpOrdPx sets TpOrdPx field to given value.

### HasTpOrdPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasTpOrdPx() bool`

HasTpOrdPx returns a boolean if a field has been set.

### GetTpTriggerPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetTpTriggerPx() string`

GetTpTriggerPx returns the TpTriggerPx field if non-nil, zero value otherwise.

### GetTpTriggerPxOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetTpTriggerPxOk() (*string, bool)`

GetTpTriggerPxOk returns a tuple with the TpTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetTpTriggerPx(v string)`

SetTpTriggerPx sets TpTriggerPx field to given value.

### HasTpTriggerPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasTpTriggerPx() bool`

HasTpTriggerPx returns a boolean if a field has been set.

### GetTpTriggerPxType

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetTpTriggerPxType() string`

GetTpTriggerPxType returns the TpTriggerPxType field if non-nil, zero value otherwise.

### GetTpTriggerPxTypeOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetTpTriggerPxTypeOk() (*string, bool)`

GetTpTriggerPxTypeOk returns a tuple with the TpTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPxType

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetTpTriggerPxType(v string)`

SetTpTriggerPxType sets TpTriggerPxType field to given value.

### HasTpTriggerPxType

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasTpTriggerPxType() bool`

HasTpTriggerPxType returns a boolean if a field has been set.

### GetTriggerPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetTriggerPx() string`

GetTriggerPx returns the TriggerPx field if non-nil, zero value otherwise.

### GetTriggerPxOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetTriggerPxOk() (*string, bool)`

GetTriggerPxOk returns a tuple with the TriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetTriggerPx(v string)`

SetTriggerPx sets TriggerPx field to given value.

### HasTriggerPx

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasTriggerPx() bool`

HasTriggerPx returns a boolean if a field has been set.

### GetTriggerPxType

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetTriggerPxType() string`

GetTriggerPxType returns the TriggerPxType field if non-nil, zero value otherwise.

### GetTriggerPxTypeOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetTriggerPxTypeOk() (*string, bool)`

GetTriggerPxTypeOk returns a tuple with the TriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerPxType

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetTriggerPxType(v string)`

SetTriggerPxType sets TriggerPxType field to given value.

### HasTriggerPxType

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasTriggerPxType() bool`

HasTriggerPxType returns a boolean if a field has been set.

### GetTriggerTime

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetTriggerTime() string`

GetTriggerTime returns the TriggerTime field if non-nil, zero value otherwise.

### GetTriggerTimeOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetTriggerTimeOk() (*string, bool)`

GetTriggerTimeOk returns a tuple with the TriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTime

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetTriggerTime(v string)`

SetTriggerTime sets TriggerTime field to given value.

### HasTriggerTime

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasTriggerTime() bool`

HasTriggerTime returns a boolean if a field has been set.

### GetUTime

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetTradeOrdersAlgoPendingV5RespData) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetTradeOrdersAlgoPendingV5RespData) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetTradeOrdersAlgoPendingV5RespData) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


