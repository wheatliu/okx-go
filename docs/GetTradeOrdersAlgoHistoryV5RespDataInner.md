# GetTradeOrdersAlgoHistoryV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePx** | Pointer to **string** | Active price  Only applicable to &#x60;move_order_stop&#x60; order | [optional] [default to ""]
**ActualPx** | Pointer to **string** | Actual order price | [optional] [default to ""]
**ActualSide** | Pointer to **string** | Actual trigger side, &#x60;tp&#x60;: take profit &#x60;sl&#x60;: stop loss  Only applicable to oco order and conditional order | [optional] [default to ""]
**ActualSz** | Pointer to **string** | Actual order quantity | [optional] [default to ""]
**AlgoClOrdId** | Pointer to **string** | Client Algo Order ID as assigned by the client. | [optional] [default to ""]
**AlgoId** | Pointer to **string** | Algo ID | [optional] [default to ""]
**AmendPxOnTriggerType** | Pointer to **string** | Whether to enable Cost-price SL. Only applicable to SL order of split TPs.   &#x60;0&#x60;: disable, the default value   &#x60;1&#x60;: Enable | [optional] [default to ""]
**AttachAlgoOrds** | Pointer to [**[]GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner**](GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner.md) | Attached SL/TP orders info  Applicable to &#x60;Spot and futures mode/Multi-currency margin/Portfolio margin&#x60; | [optional] 
**CTime** | Pointer to **string** | Creation time Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**CallbackRatio** | Pointer to **string** | Callback price ratio  Only applicable to &#x60;move_order_stop&#x60; order | [optional] [default to ""]
**CallbackSpread** | Pointer to **string** | Callback price variance  Only applicable to &#x60;move_order_stop&#x60; order | [optional] [default to ""]
**Ccy** | Pointer to **string** | Margin currency   Applicable to all &#x60;isolated&#x60; &#x60;MARGIN&#x60; orders and &#x60;cross&#x60; &#x60;MARGIN&#x60; orders in &#x60;Spot and futures mode&#x60;. | [optional] [default to ""]
**ChaseType** | Pointer to **string** | Chase type. Only applicable to &#x60;chase&#x60; order. | [optional] [default to ""]
**ChaseVal** | Pointer to **string** | Chase value. Only applicable to &#x60;chase&#x60; order. | [optional] [default to ""]
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client | [optional] [default to ""]
**CloseFraction** | Pointer to **string** | Fraction of position to be closed when the algo order is triggered | [optional] [default to ""]
**FailCode** | Pointer to **string** | It represents that the reason that algo order fails to trigger. It is \&quot;\&quot; when the state is &#x60;effective&#x60;/&#x60;canceled&#x60;. There will be value when the state is &#x60;order_failed&#x60;, e.g. 51008;  Only applicable to Stop Order, Trailing Stop Order, Trigger order. | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**IsTradeBorrowMode** | Pointer to **string** | Whether borrowing currency automatically   true   false  Only applicable to &#x60;trigger order&#x60;, &#x60;trailing order&#x60; and &#x60;twap order&#x60; | [optional] [default to ""]
**Last** | Pointer to **string** | Last filled price while placing | [optional] [default to ""]
**Lever** | Pointer to **string** | Leverage, from &#x60;0.01&#x60; to &#x60;125&#x60;.   Only applicable to &#x60;MARGIN/FUTURES/SWAP&#x60; | [optional] [default to ""]
**LinkedOrd** | Pointer to [**GetTradeOrderAlgoV5RespDataInnerLinkedOrd**](GetTradeOrderAlgoV5RespDataInnerLinkedOrd.md) |  | [optional] 
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
**SlOrdPx** | Pointer to **string** | Stop-loss order price. | [optional] [default to ""]
**SlTriggerPx** | Pointer to **string** | Stop-loss trigger price. | [optional] [default to ""]
**SlTriggerPxType** | Pointer to **string** | Stop-loss trigger price type.   &#x60;last&#x60;: last price  &#x60;index&#x60;: index price  &#x60;mark&#x60;: mark price | [optional] [default to ""]
**State** | Pointer to **string** | State   &#x60;effective&#x60;   &#x60;canceled&#x60;   &#x60;order_failed&#x60;   &#x60;partially_failed&#x60; | [optional] [default to ""]
**Sz** | Pointer to **string** | Quantity to buy or sell | [optional] [default to ""]
**SzLimit** | Pointer to **string** | Average amount   Only applicable to &#x60;iceberg&#x60; order or &#x60;twap&#x60; order | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**TdMode** | Pointer to **string** | Trade mode | [optional] [default to ""]
**TgtCcy** | Pointer to **string** | Order quantity unit setting for &#x60;sz&#x60;  &#x60;base_ccy&#x60;: Base currency ,&#x60;quote_ccy&#x60;: Quote currency    Only applicable to &#x60;SPOT&#x60; Market Orders  Default is &#x60;quote_ccy&#x60; for buy, &#x60;base_ccy&#x60; for sell | [optional] [default to ""]
**TimeInterval** | Pointer to **string** | Time interval   Only applicable to &#x60;twap&#x60; order | [optional] [default to ""]
**TpOrdPx** | Pointer to **string** | Take-profit order price. | [optional] [default to ""]
**TpTriggerPx** | Pointer to **string** | Take-profit trigger price. | [optional] [default to ""]
**TpTriggerPxType** | Pointer to **string** | Take-profit trigger price type.   &#x60;last&#x60;: last price  &#x60;index&#x60;: index price  &#x60;mark&#x60;: mark price | [optional] [default to ""]
**TriggerPx** | Pointer to **string** | trigger price. | [optional] [default to ""]
**TriggerPxType** | Pointer to **string** | trigger price type.   &#x60;last&#x60;: last price  &#x60;index&#x60;: index price  &#x60;mark&#x60;: mark price | [optional] [default to ""]
**TriggerTime** | Pointer to **string** | Trigger time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**UTime** | Pointer to **string** | Order updated time, Unix timestamp format in milliseconds, e.g. 1597026383085 | [optional] [default to ""]

## Methods

### NewGetTradeOrdersAlgoHistoryV5RespDataInner

`func NewGetTradeOrdersAlgoHistoryV5RespDataInner() *GetTradeOrdersAlgoHistoryV5RespDataInner`

NewGetTradeOrdersAlgoHistoryV5RespDataInner instantiates a new GetTradeOrdersAlgoHistoryV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradeOrdersAlgoHistoryV5RespDataInnerWithDefaults

`func NewGetTradeOrdersAlgoHistoryV5RespDataInnerWithDefaults() *GetTradeOrdersAlgoHistoryV5RespDataInner`

NewGetTradeOrdersAlgoHistoryV5RespDataInnerWithDefaults instantiates a new GetTradeOrdersAlgoHistoryV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetActivePx() string`

GetActivePx returns the ActivePx field if non-nil, zero value otherwise.

### GetActivePxOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetActivePxOk() (*string, bool)`

GetActivePxOk returns a tuple with the ActivePx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetActivePx(v string)`

SetActivePx sets ActivePx field to given value.

### HasActivePx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasActivePx() bool`

HasActivePx returns a boolean if a field has been set.

### GetActualPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetActualPx() string`

GetActualPx returns the ActualPx field if non-nil, zero value otherwise.

### GetActualPxOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetActualPxOk() (*string, bool)`

GetActualPxOk returns a tuple with the ActualPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetActualPx(v string)`

SetActualPx sets ActualPx field to given value.

### HasActualPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasActualPx() bool`

HasActualPx returns a boolean if a field has been set.

### GetActualSide

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetActualSide() string`

GetActualSide returns the ActualSide field if non-nil, zero value otherwise.

### GetActualSideOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetActualSideOk() (*string, bool)`

GetActualSideOk returns a tuple with the ActualSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualSide

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetActualSide(v string)`

SetActualSide sets ActualSide field to given value.

### HasActualSide

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasActualSide() bool`

HasActualSide returns a boolean if a field has been set.

### GetActualSz

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetActualSz() string`

GetActualSz returns the ActualSz field if non-nil, zero value otherwise.

### GetActualSzOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetActualSzOk() (*string, bool)`

GetActualSzOk returns a tuple with the ActualSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualSz

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetActualSz(v string)`

SetActualSz sets ActualSz field to given value.

### HasActualSz

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasActualSz() bool`

HasActualSz returns a boolean if a field has been set.

### GetAlgoClOrdId

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAlgoId

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetAmendPxOnTriggerType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetAmendPxOnTriggerType() string`

GetAmendPxOnTriggerType returns the AmendPxOnTriggerType field if non-nil, zero value otherwise.

### GetAmendPxOnTriggerTypeOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetAmendPxOnTriggerTypeOk() (*string, bool)`

GetAmendPxOnTriggerTypeOk returns a tuple with the AmendPxOnTriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendPxOnTriggerType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetAmendPxOnTriggerType(v string)`

SetAmendPxOnTriggerType sets AmendPxOnTriggerType field to given value.

### HasAmendPxOnTriggerType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasAmendPxOnTriggerType() bool`

HasAmendPxOnTriggerType returns a boolean if a field has been set.

### GetAttachAlgoOrds

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetAttachAlgoOrds() []GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner`

GetAttachAlgoOrds returns the AttachAlgoOrds field if non-nil, zero value otherwise.

### GetAttachAlgoOrdsOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetAttachAlgoOrdsOk() (*[]GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner, bool)`

GetAttachAlgoOrdsOk returns a tuple with the AttachAlgoOrds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachAlgoOrds

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetAttachAlgoOrds(v []GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner)`

SetAttachAlgoOrds sets AttachAlgoOrds field to given value.

### HasAttachAlgoOrds

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasAttachAlgoOrds() bool`

HasAttachAlgoOrds returns a boolean if a field has been set.

### GetCTime

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCallbackRatio

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetCallbackRatio() string`

GetCallbackRatio returns the CallbackRatio field if non-nil, zero value otherwise.

### GetCallbackRatioOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetCallbackRatioOk() (*string, bool)`

GetCallbackRatioOk returns a tuple with the CallbackRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackRatio

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetCallbackRatio(v string)`

SetCallbackRatio sets CallbackRatio field to given value.

### HasCallbackRatio

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasCallbackRatio() bool`

HasCallbackRatio returns a boolean if a field has been set.

### GetCallbackSpread

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetCallbackSpread() string`

GetCallbackSpread returns the CallbackSpread field if non-nil, zero value otherwise.

### GetCallbackSpreadOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetCallbackSpreadOk() (*string, bool)`

GetCallbackSpreadOk returns a tuple with the CallbackSpread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSpread

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetCallbackSpread(v string)`

SetCallbackSpread sets CallbackSpread field to given value.

### HasCallbackSpread

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasCallbackSpread() bool`

HasCallbackSpread returns a boolean if a field has been set.

### GetCcy

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetChaseType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetChaseType() string`

GetChaseType returns the ChaseType field if non-nil, zero value otherwise.

### GetChaseTypeOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetChaseTypeOk() (*string, bool)`

GetChaseTypeOk returns a tuple with the ChaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChaseType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetChaseType(v string)`

SetChaseType sets ChaseType field to given value.

### HasChaseType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasChaseType() bool`

HasChaseType returns a boolean if a field has been set.

### GetChaseVal

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetChaseVal() string`

GetChaseVal returns the ChaseVal field if non-nil, zero value otherwise.

### GetChaseValOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetChaseValOk() (*string, bool)`

GetChaseValOk returns a tuple with the ChaseVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChaseVal

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetChaseVal(v string)`

SetChaseVal sets ChaseVal field to given value.

### HasChaseVal

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasChaseVal() bool`

HasChaseVal returns a boolean if a field has been set.

### GetClOrdId

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetCloseFraction

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetCloseFraction() string`

GetCloseFraction returns the CloseFraction field if non-nil, zero value otherwise.

### GetCloseFractionOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetCloseFractionOk() (*string, bool)`

GetCloseFractionOk returns a tuple with the CloseFraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseFraction

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetCloseFraction(v string)`

SetCloseFraction sets CloseFraction field to given value.

### HasCloseFraction

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasCloseFraction() bool`

HasCloseFraction returns a boolean if a field has been set.

### GetFailCode

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetFailCode() string`

GetFailCode returns the FailCode field if non-nil, zero value otherwise.

### GetFailCodeOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetFailCodeOk() (*string, bool)`

GetFailCodeOk returns a tuple with the FailCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailCode

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetFailCode(v string)`

SetFailCode sets FailCode field to given value.

### HasFailCode

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasFailCode() bool`

HasFailCode returns a boolean if a field has been set.

### GetInstId

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetIsTradeBorrowMode

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetIsTradeBorrowMode() string`

GetIsTradeBorrowMode returns the IsTradeBorrowMode field if non-nil, zero value otherwise.

### GetIsTradeBorrowModeOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetIsTradeBorrowModeOk() (*string, bool)`

GetIsTradeBorrowModeOk returns a tuple with the IsTradeBorrowMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTradeBorrowMode

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetIsTradeBorrowMode(v string)`

SetIsTradeBorrowMode sets IsTradeBorrowMode field to given value.

### HasIsTradeBorrowMode

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasIsTradeBorrowMode() bool`

HasIsTradeBorrowMode returns a boolean if a field has been set.

### GetLast

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetLever

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetLinkedOrd

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetLinkedOrd() GetTradeOrderAlgoV5RespDataInnerLinkedOrd`

GetLinkedOrd returns the LinkedOrd field if non-nil, zero value otherwise.

### GetLinkedOrdOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetLinkedOrdOk() (*GetTradeOrderAlgoV5RespDataInnerLinkedOrd, bool)`

GetLinkedOrdOk returns a tuple with the LinkedOrd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedOrd

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetLinkedOrd(v GetTradeOrderAlgoV5RespDataInnerLinkedOrd)`

SetLinkedOrd sets LinkedOrd field to given value.

### HasLinkedOrd

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasLinkedOrd() bool`

HasLinkedOrd returns a boolean if a field has been set.

### GetMaxChaseType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetMaxChaseType() string`

GetMaxChaseType returns the MaxChaseType field if non-nil, zero value otherwise.

### GetMaxChaseTypeOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetMaxChaseTypeOk() (*string, bool)`

GetMaxChaseTypeOk returns a tuple with the MaxChaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxChaseType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetMaxChaseType(v string)`

SetMaxChaseType sets MaxChaseType field to given value.

### HasMaxChaseType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasMaxChaseType() bool`

HasMaxChaseType returns a boolean if a field has been set.

### GetMaxChaseVal

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetMaxChaseVal() string`

GetMaxChaseVal returns the MaxChaseVal field if non-nil, zero value otherwise.

### GetMaxChaseValOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetMaxChaseValOk() (*string, bool)`

GetMaxChaseValOk returns a tuple with the MaxChaseVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxChaseVal

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetMaxChaseVal(v string)`

SetMaxChaseVal sets MaxChaseVal field to given value.

### HasMaxChaseVal

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasMaxChaseVal() bool`

HasMaxChaseVal returns a boolean if a field has been set.

### GetMoveTriggerPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetMoveTriggerPx() string`

GetMoveTriggerPx returns the MoveTriggerPx field if non-nil, zero value otherwise.

### GetMoveTriggerPxOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetMoveTriggerPxOk() (*string, bool)`

GetMoveTriggerPxOk returns a tuple with the MoveTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveTriggerPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetMoveTriggerPx(v string)`

SetMoveTriggerPx sets MoveTriggerPx field to given value.

### HasMoveTriggerPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasMoveTriggerPx() bool`

HasMoveTriggerPx returns a boolean if a field has been set.

### GetOrdId

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetOrdIdList

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetOrdIdList() []string`

GetOrdIdList returns the OrdIdList field if non-nil, zero value otherwise.

### GetOrdIdListOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetOrdIdListOk() (*[]string, bool)`

GetOrdIdListOk returns a tuple with the OrdIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdIdList

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetOrdIdList(v []string)`

SetOrdIdList sets OrdIdList field to given value.

### HasOrdIdList

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasOrdIdList() bool`

HasOrdIdList returns a boolean if a field has been set.

### GetOrdPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetOrdPx() string`

GetOrdPx returns the OrdPx field if non-nil, zero value otherwise.

### GetOrdPxOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetOrdPxOk() (*string, bool)`

GetOrdPxOk returns a tuple with the OrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetOrdPx(v string)`

SetOrdPx sets OrdPx field to given value.

### HasOrdPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasOrdPx() bool`

HasOrdPx returns a boolean if a field has been set.

### GetOrdType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetOrdType() string`

GetOrdType returns the OrdType field if non-nil, zero value otherwise.

### GetOrdTypeOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetOrdTypeOk() (*string, bool)`

GetOrdTypeOk returns a tuple with the OrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetOrdType(v string)`

SetOrdType sets OrdType field to given value.

### HasOrdType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasOrdType() bool`

HasOrdType returns a boolean if a field has been set.

### GetPosSide

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetPxLimit

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetPxLimit() string`

GetPxLimit returns the PxLimit field if non-nil, zero value otherwise.

### GetPxLimitOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetPxLimitOk() (*string, bool)`

GetPxLimitOk returns a tuple with the PxLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxLimit

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetPxLimit(v string)`

SetPxLimit sets PxLimit field to given value.

### HasPxLimit

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasPxLimit() bool`

HasPxLimit returns a boolean if a field has been set.

### GetPxSpread

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetPxSpread() string`

GetPxSpread returns the PxSpread field if non-nil, zero value otherwise.

### GetPxSpreadOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetPxSpreadOk() (*string, bool)`

GetPxSpreadOk returns a tuple with the PxSpread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxSpread

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetPxSpread(v string)`

SetPxSpread sets PxSpread field to given value.

### HasPxSpread

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasPxSpread() bool`

HasPxSpread returns a boolean if a field has been set.

### GetPxVar

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetPxVar() string`

GetPxVar returns the PxVar field if non-nil, zero value otherwise.

### GetPxVarOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetPxVarOk() (*string, bool)`

GetPxVarOk returns a tuple with the PxVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxVar

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetPxVar(v string)`

SetPxVar sets PxVar field to given value.

### HasPxVar

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasPxVar() bool`

HasPxVar returns a boolean if a field has been set.

### GetQuickMgnType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetQuickMgnType() string`

GetQuickMgnType returns the QuickMgnType field if non-nil, zero value otherwise.

### GetQuickMgnTypeOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetQuickMgnTypeOk() (*string, bool)`

GetQuickMgnTypeOk returns a tuple with the QuickMgnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickMgnType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetQuickMgnType(v string)`

SetQuickMgnType sets QuickMgnType field to given value.

### HasQuickMgnType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasQuickMgnType() bool`

HasQuickMgnType returns a boolean if a field has been set.

### GetReduceOnly

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetReduceOnly() string`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetReduceOnlyOk() (*string, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetReduceOnly(v string)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSide

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSlOrdPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetSlOrdPx() string`

GetSlOrdPx returns the SlOrdPx field if non-nil, zero value otherwise.

### GetSlOrdPxOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetSlOrdPxOk() (*string, bool)`

GetSlOrdPxOk returns a tuple with the SlOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlOrdPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetSlOrdPx(v string)`

SetSlOrdPx sets SlOrdPx field to given value.

### HasSlOrdPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasSlOrdPx() bool`

HasSlOrdPx returns a boolean if a field has been set.

### GetSlTriggerPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetSlTriggerPx() string`

GetSlTriggerPx returns the SlTriggerPx field if non-nil, zero value otherwise.

### GetSlTriggerPxOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetSlTriggerPxOk() (*string, bool)`

GetSlTriggerPxOk returns a tuple with the SlTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetSlTriggerPx(v string)`

SetSlTriggerPx sets SlTriggerPx field to given value.

### HasSlTriggerPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasSlTriggerPx() bool`

HasSlTriggerPx returns a boolean if a field has been set.

### GetSlTriggerPxType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetSlTriggerPxType() string`

GetSlTriggerPxType returns the SlTriggerPxType field if non-nil, zero value otherwise.

### GetSlTriggerPxTypeOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetSlTriggerPxTypeOk() (*string, bool)`

GetSlTriggerPxTypeOk returns a tuple with the SlTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPxType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetSlTriggerPxType(v string)`

SetSlTriggerPxType sets SlTriggerPxType field to given value.

### HasSlTriggerPxType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasSlTriggerPxType() bool`

HasSlTriggerPxType returns a boolean if a field has been set.

### GetState

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSz

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetSzLimit

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetSzLimit() string`

GetSzLimit returns the SzLimit field if non-nil, zero value otherwise.

### GetSzLimitOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetSzLimitOk() (*string, bool)`

GetSzLimitOk returns a tuple with the SzLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSzLimit

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetSzLimit(v string)`

SetSzLimit sets SzLimit field to given value.

### HasSzLimit

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasSzLimit() bool`

HasSzLimit returns a boolean if a field has been set.

### GetTag

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTdMode

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetTdMode() string`

GetTdMode returns the TdMode field if non-nil, zero value otherwise.

### GetTdModeOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetTdModeOk() (*string, bool)`

GetTdModeOk returns a tuple with the TdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdMode

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetTdMode(v string)`

SetTdMode sets TdMode field to given value.

### HasTdMode

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasTdMode() bool`

HasTdMode returns a boolean if a field has been set.

### GetTgtCcy

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetTgtCcy() string`

GetTgtCcy returns the TgtCcy field if non-nil, zero value otherwise.

### GetTgtCcyOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetTgtCcyOk() (*string, bool)`

GetTgtCcyOk returns a tuple with the TgtCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCcy

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetTgtCcy(v string)`

SetTgtCcy sets TgtCcy field to given value.

### HasTgtCcy

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasTgtCcy() bool`

HasTgtCcy returns a boolean if a field has been set.

### GetTimeInterval

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetTimeInterval() string`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetTimeIntervalOk() (*string, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetTimeInterval(v string)`

SetTimeInterval sets TimeInterval field to given value.

### HasTimeInterval

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasTimeInterval() bool`

HasTimeInterval returns a boolean if a field has been set.

### GetTpOrdPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetTpOrdPx() string`

GetTpOrdPx returns the TpOrdPx field if non-nil, zero value otherwise.

### GetTpOrdPxOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetTpOrdPxOk() (*string, bool)`

GetTpOrdPxOk returns a tuple with the TpOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpOrdPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetTpOrdPx(v string)`

SetTpOrdPx sets TpOrdPx field to given value.

### HasTpOrdPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasTpOrdPx() bool`

HasTpOrdPx returns a boolean if a field has been set.

### GetTpTriggerPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetTpTriggerPx() string`

GetTpTriggerPx returns the TpTriggerPx field if non-nil, zero value otherwise.

### GetTpTriggerPxOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetTpTriggerPxOk() (*string, bool)`

GetTpTriggerPxOk returns a tuple with the TpTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetTpTriggerPx(v string)`

SetTpTriggerPx sets TpTriggerPx field to given value.

### HasTpTriggerPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasTpTriggerPx() bool`

HasTpTriggerPx returns a boolean if a field has been set.

### GetTpTriggerPxType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetTpTriggerPxType() string`

GetTpTriggerPxType returns the TpTriggerPxType field if non-nil, zero value otherwise.

### GetTpTriggerPxTypeOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetTpTriggerPxTypeOk() (*string, bool)`

GetTpTriggerPxTypeOk returns a tuple with the TpTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPxType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetTpTriggerPxType(v string)`

SetTpTriggerPxType sets TpTriggerPxType field to given value.

### HasTpTriggerPxType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasTpTriggerPxType() bool`

HasTpTriggerPxType returns a boolean if a field has been set.

### GetTriggerPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetTriggerPx() string`

GetTriggerPx returns the TriggerPx field if non-nil, zero value otherwise.

### GetTriggerPxOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetTriggerPxOk() (*string, bool)`

GetTriggerPxOk returns a tuple with the TriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetTriggerPx(v string)`

SetTriggerPx sets TriggerPx field to given value.

### HasTriggerPx

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasTriggerPx() bool`

HasTriggerPx returns a boolean if a field has been set.

### GetTriggerPxType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetTriggerPxType() string`

GetTriggerPxType returns the TriggerPxType field if non-nil, zero value otherwise.

### GetTriggerPxTypeOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetTriggerPxTypeOk() (*string, bool)`

GetTriggerPxTypeOk returns a tuple with the TriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerPxType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetTriggerPxType(v string)`

SetTriggerPxType sets TriggerPxType field to given value.

### HasTriggerPxType

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasTriggerPxType() bool`

HasTriggerPxType returns a boolean if a field has been set.

### GetTriggerTime

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetTriggerTime() string`

GetTriggerTime returns the TriggerTime field if non-nil, zero value otherwise.

### GetTriggerTimeOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetTriggerTimeOk() (*string, bool)`

GetTriggerTimeOk returns a tuple with the TriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTime

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetTriggerTime(v string)`

SetTriggerTime sets TriggerTime field to given value.

### HasTriggerTime

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasTriggerTime() bool`

HasTriggerTime returns a boolean if a field has been set.

### GetUTime

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetTradeOrdersAlgoHistoryV5RespDataInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


