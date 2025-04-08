# GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID | [optional] [default to ""]
**AlgoId** | Pointer to **string** | Algo ID | [optional] [default to ""]
**AlgoOrdType** | Pointer to **string** | Algo order type  &#x60;contract&#x60;: Contract signal | [optional] [default to ""]
**AvailBal** | Pointer to **string** | Avail balance | [optional] [default to ""]
**CTime** | Pointer to **string** | Algo order created time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**CancelType** | Pointer to **string** | Algo order stop reason  &#x60;1&#x60;: Manual stop | [optional] [default to ""]
**EntrySettingParam** | Pointer to [**GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam**](GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam.md) |  | [optional] 
**ExitSettingParam** | Pointer to [**GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam**](GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam.md) |  | [optional] 
**FloatPnl** | Pointer to **string** | Float P&amp;L | [optional] [default to ""]
**FrozenBal** | Pointer to **string** | Frozen balance | [optional] [default to ""]
**InstIds** | Pointer to **[]string** | Instrument IDs | [optional] 
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**InvestAmt** | Pointer to **string** | Investment amount | [optional] [default to ""]
**Lever** | Pointer to **string** | Leverage  Only applicable to &#x60;contract signal&#x60; | [optional] [default to ""]
**Ratio** | Pointer to **string** | Price offset ratio, calculate the limit price as a percentage offset from the best bid/ask price  Only applicable to &#x60;subOrdType&#x60; is &#x60;limit order&#x60; | [optional] [default to ""]
**RealizedPnl** | Pointer to **string** | Realized P&amp;L | [optional] [default to ""]
**SignalChanId** | Pointer to **string** | Signal channel Id | [optional] [default to ""]
**SignalChanName** | Pointer to **string** | Signal channel name | [optional] [default to ""]
**SignalSourceType** | Pointer to **string** | Signal source type  &#x60;1&#x60;: Created by yourself  &#x60;2&#x60;: Subscribe  &#x60;3&#x60;: Free signal | [optional] [default to ""]
**State** | Pointer to **string** | Algo order state  &#x60;stopped&#x60; | [optional] [default to ""]
**SubOrdType** | Pointer to **string** | Sub order type  &#x60;1&#x60;：limit order  &#x60;2&#x60;：market order  &#x60;9&#x60;：tradingView signal | [optional] [default to ""]
**TotalEq** | Pointer to **string** | Total equity of strategy account | [optional] [default to ""]
**TotalPnl** | Pointer to **string** | Total P&amp;L | [optional] [default to ""]
**TotalPnlRatio** | Pointer to **string** | Total P&amp;L ratio | [optional] [default to ""]
**UTime** | Pointer to **string** | Algo order updated time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetTradingBotSignalOrdersAlgoHistoryV5RespDataInner

`func NewGetTradingBotSignalOrdersAlgoHistoryV5RespDataInner() *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner`

NewGetTradingBotSignalOrdersAlgoHistoryV5RespDataInner instantiates a new GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotSignalOrdersAlgoHistoryV5RespDataInnerWithDefaults

`func NewGetTradingBotSignalOrdersAlgoHistoryV5RespDataInnerWithDefaults() *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner`

NewGetTradingBotSignalOrdersAlgoHistoryV5RespDataInnerWithDefaults instantiates a new GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoClOrdId

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAlgoId

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetAlgoOrdType

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetAlgoOrdType() string`

GetAlgoOrdType returns the AlgoOrdType field if non-nil, zero value otherwise.

### GetAlgoOrdTypeOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetAlgoOrdTypeOk() (*string, bool)`

GetAlgoOrdTypeOk returns a tuple with the AlgoOrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoOrdType

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetAlgoOrdType(v string)`

SetAlgoOrdType sets AlgoOrdType field to given value.

### HasAlgoOrdType

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasAlgoOrdType() bool`

HasAlgoOrdType returns a boolean if a field has been set.

### GetAvailBal

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetAvailBal() string`

GetAvailBal returns the AvailBal field if non-nil, zero value otherwise.

### GetAvailBalOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetAvailBalOk() (*string, bool)`

GetAvailBalOk returns a tuple with the AvailBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailBal

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetAvailBal(v string)`

SetAvailBal sets AvailBal field to given value.

### HasAvailBal

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasAvailBal() bool`

HasAvailBal returns a boolean if a field has been set.

### GetCTime

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCancelType

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetCancelType() string`

GetCancelType returns the CancelType field if non-nil, zero value otherwise.

### GetCancelTypeOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetCancelTypeOk() (*string, bool)`

GetCancelTypeOk returns a tuple with the CancelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelType

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetCancelType(v string)`

SetCancelType sets CancelType field to given value.

### HasCancelType

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasCancelType() bool`

HasCancelType returns a boolean if a field has been set.

### GetEntrySettingParam

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetEntrySettingParam() GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam`

GetEntrySettingParam returns the EntrySettingParam field if non-nil, zero value otherwise.

### GetEntrySettingParamOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetEntrySettingParamOk() (*GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam, bool)`

GetEntrySettingParamOk returns a tuple with the EntrySettingParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrySettingParam

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetEntrySettingParam(v GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam)`

SetEntrySettingParam sets EntrySettingParam field to given value.

### HasEntrySettingParam

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasEntrySettingParam() bool`

HasEntrySettingParam returns a boolean if a field has been set.

### GetExitSettingParam

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetExitSettingParam() GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam`

GetExitSettingParam returns the ExitSettingParam field if non-nil, zero value otherwise.

### GetExitSettingParamOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetExitSettingParamOk() (*GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam, bool)`

GetExitSettingParamOk returns a tuple with the ExitSettingParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitSettingParam

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetExitSettingParam(v GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam)`

SetExitSettingParam sets ExitSettingParam field to given value.

### HasExitSettingParam

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasExitSettingParam() bool`

HasExitSettingParam returns a boolean if a field has been set.

### GetFloatPnl

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetFloatPnl() string`

GetFloatPnl returns the FloatPnl field if non-nil, zero value otherwise.

### GetFloatPnlOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetFloatPnlOk() (*string, bool)`

GetFloatPnlOk returns a tuple with the FloatPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatPnl

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetFloatPnl(v string)`

SetFloatPnl sets FloatPnl field to given value.

### HasFloatPnl

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasFloatPnl() bool`

HasFloatPnl returns a boolean if a field has been set.

### GetFrozenBal

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetFrozenBal() string`

GetFrozenBal returns the FrozenBal field if non-nil, zero value otherwise.

### GetFrozenBalOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetFrozenBalOk() (*string, bool)`

GetFrozenBalOk returns a tuple with the FrozenBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenBal

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetFrozenBal(v string)`

SetFrozenBal sets FrozenBal field to given value.

### HasFrozenBal

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasFrozenBal() bool`

HasFrozenBal returns a boolean if a field has been set.

### GetInstIds

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetInstIds() []string`

GetInstIds returns the InstIds field if non-nil, zero value otherwise.

### GetInstIdsOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetInstIdsOk() (*[]string, bool)`

GetInstIdsOk returns a tuple with the InstIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstIds

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetInstIds(v []string)`

SetInstIds sets InstIds field to given value.

### HasInstIds

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasInstIds() bool`

HasInstIds returns a boolean if a field has been set.

### GetInstType

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetInvestAmt

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetInvestAmt() string`

GetInvestAmt returns the InvestAmt field if non-nil, zero value otherwise.

### GetInvestAmtOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetInvestAmtOk() (*string, bool)`

GetInvestAmtOk returns a tuple with the InvestAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestAmt

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetInvestAmt(v string)`

SetInvestAmt sets InvestAmt field to given value.

### HasInvestAmt

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasInvestAmt() bool`

HasInvestAmt returns a boolean if a field has been set.

### GetLever

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetRatio

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetRatio() string`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetRatioOk() (*string, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetRatio(v string)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetRealizedPnl

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetRealizedPnl() string`

GetRealizedPnl returns the RealizedPnl field if non-nil, zero value otherwise.

### GetRealizedPnlOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetRealizedPnlOk() (*string, bool)`

GetRealizedPnlOk returns a tuple with the RealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedPnl

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetRealizedPnl(v string)`

SetRealizedPnl sets RealizedPnl field to given value.

### HasRealizedPnl

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasRealizedPnl() bool`

HasRealizedPnl returns a boolean if a field has been set.

### GetSignalChanId

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetSignalChanId() string`

GetSignalChanId returns the SignalChanId field if non-nil, zero value otherwise.

### GetSignalChanIdOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetSignalChanIdOk() (*string, bool)`

GetSignalChanIdOk returns a tuple with the SignalChanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalChanId

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetSignalChanId(v string)`

SetSignalChanId sets SignalChanId field to given value.

### HasSignalChanId

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasSignalChanId() bool`

HasSignalChanId returns a boolean if a field has been set.

### GetSignalChanName

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetSignalChanName() string`

GetSignalChanName returns the SignalChanName field if non-nil, zero value otherwise.

### GetSignalChanNameOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetSignalChanNameOk() (*string, bool)`

GetSignalChanNameOk returns a tuple with the SignalChanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalChanName

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetSignalChanName(v string)`

SetSignalChanName sets SignalChanName field to given value.

### HasSignalChanName

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasSignalChanName() bool`

HasSignalChanName returns a boolean if a field has been set.

### GetSignalSourceType

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetSignalSourceType() string`

GetSignalSourceType returns the SignalSourceType field if non-nil, zero value otherwise.

### GetSignalSourceTypeOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetSignalSourceTypeOk() (*string, bool)`

GetSignalSourceTypeOk returns a tuple with the SignalSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalSourceType

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetSignalSourceType(v string)`

SetSignalSourceType sets SignalSourceType field to given value.

### HasSignalSourceType

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasSignalSourceType() bool`

HasSignalSourceType returns a boolean if a field has been set.

### GetState

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubOrdType

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetSubOrdType() string`

GetSubOrdType returns the SubOrdType field if non-nil, zero value otherwise.

### GetSubOrdTypeOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetSubOrdTypeOk() (*string, bool)`

GetSubOrdTypeOk returns a tuple with the SubOrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrdType

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetSubOrdType(v string)`

SetSubOrdType sets SubOrdType field to given value.

### HasSubOrdType

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasSubOrdType() bool`

HasSubOrdType returns a boolean if a field has been set.

### GetTotalEq

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetTotalEq() string`

GetTotalEq returns the TotalEq field if non-nil, zero value otherwise.

### GetTotalEqOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetTotalEqOk() (*string, bool)`

GetTotalEqOk returns a tuple with the TotalEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEq

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetTotalEq(v string)`

SetTotalEq sets TotalEq field to given value.

### HasTotalEq

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasTotalEq() bool`

HasTotalEq returns a boolean if a field has been set.

### GetTotalPnl

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetTotalPnl() string`

GetTotalPnl returns the TotalPnl field if non-nil, zero value otherwise.

### GetTotalPnlOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetTotalPnlOk() (*string, bool)`

GetTotalPnlOk returns a tuple with the TotalPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPnl

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetTotalPnl(v string)`

SetTotalPnl sets TotalPnl field to given value.

### HasTotalPnl

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasTotalPnl() bool`

HasTotalPnl returns a boolean if a field has been set.

### GetTotalPnlRatio

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetTotalPnlRatio() string`

GetTotalPnlRatio returns the TotalPnlRatio field if non-nil, zero value otherwise.

### GetTotalPnlRatioOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetTotalPnlRatioOk() (*string, bool)`

GetTotalPnlRatioOk returns a tuple with the TotalPnlRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPnlRatio

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetTotalPnlRatio(v string)`

SetTotalPnlRatio sets TotalPnlRatio field to given value.

### HasTotalPnlRatio

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasTotalPnlRatio() bool`

HasTotalPnlRatio returns a boolean if a field has been set.

### GetUTime

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetTradingBotSignalOrdersAlgoHistoryV5RespDataInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


