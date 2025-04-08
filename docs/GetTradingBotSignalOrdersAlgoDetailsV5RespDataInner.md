# GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID | [optional] [default to ""]
**AlgoId** | Pointer to **string** | Algo ID | [optional] [default to ""]
**AlgoOrdType** | Pointer to **string** | Algo order type  &#x60;contract&#x60;: Contract signal | [optional] [default to ""]
**AvailBal** | Pointer to **string** | Avail balance | [optional] [default to ""]
**CTime** | Pointer to **string** | Algo order created time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**CancelType** | Pointer to **string** | Algo order stop reason  &#x60;0&#x60;: None  &#x60;1&#x60;: Manual stop | [optional] [default to ""]
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
**State** | Pointer to **string** | Algo order state  &#x60;starting&#x60;  &#x60;running&#x60;  &#x60;stopping&#x60;  &#x60;stopped&#x60; | [optional] [default to ""]
**SubOrdType** | Pointer to **string** | Sub order type  &#x60;1&#x60;：limit order  &#x60;2&#x60;：market order  &#x60;9&#x60;：tradingView signal | [optional] [default to ""]
**TotalEq** | Pointer to **string** | Total equity of strategy account | [optional] [default to ""]
**TotalPnl** | Pointer to **string** | Total P&amp;L | [optional] [default to ""]
**TotalPnlRatio** | Pointer to **string** | Total P&amp;L ratio | [optional] [default to ""]
**UTime** | Pointer to **string** | Algo order updated time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataInner

`func NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataInner() *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner`

NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataInner instantiates a new GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerWithDefaults

`func NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerWithDefaults() *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner`

NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerWithDefaults instantiates a new GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoClOrdId

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAlgoId

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetAlgoOrdType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetAlgoOrdType() string`

GetAlgoOrdType returns the AlgoOrdType field if non-nil, zero value otherwise.

### GetAlgoOrdTypeOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetAlgoOrdTypeOk() (*string, bool)`

GetAlgoOrdTypeOk returns a tuple with the AlgoOrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoOrdType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetAlgoOrdType(v string)`

SetAlgoOrdType sets AlgoOrdType field to given value.

### HasAlgoOrdType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasAlgoOrdType() bool`

HasAlgoOrdType returns a boolean if a field has been set.

### GetAvailBal

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetAvailBal() string`

GetAvailBal returns the AvailBal field if non-nil, zero value otherwise.

### GetAvailBalOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetAvailBalOk() (*string, bool)`

GetAvailBalOk returns a tuple with the AvailBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailBal

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetAvailBal(v string)`

SetAvailBal sets AvailBal field to given value.

### HasAvailBal

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasAvailBal() bool`

HasAvailBal returns a boolean if a field has been set.

### GetCTime

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCancelType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetCancelType() string`

GetCancelType returns the CancelType field if non-nil, zero value otherwise.

### GetCancelTypeOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetCancelTypeOk() (*string, bool)`

GetCancelTypeOk returns a tuple with the CancelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetCancelType(v string)`

SetCancelType sets CancelType field to given value.

### HasCancelType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasCancelType() bool`

HasCancelType returns a boolean if a field has been set.

### GetEntrySettingParam

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetEntrySettingParam() GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam`

GetEntrySettingParam returns the EntrySettingParam field if non-nil, zero value otherwise.

### GetEntrySettingParamOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetEntrySettingParamOk() (*GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam, bool)`

GetEntrySettingParamOk returns a tuple with the EntrySettingParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrySettingParam

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetEntrySettingParam(v GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam)`

SetEntrySettingParam sets EntrySettingParam field to given value.

### HasEntrySettingParam

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasEntrySettingParam() bool`

HasEntrySettingParam returns a boolean if a field has been set.

### GetExitSettingParam

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetExitSettingParam() GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam`

GetExitSettingParam returns the ExitSettingParam field if non-nil, zero value otherwise.

### GetExitSettingParamOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetExitSettingParamOk() (*GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam, bool)`

GetExitSettingParamOk returns a tuple with the ExitSettingParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitSettingParam

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetExitSettingParam(v GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam)`

SetExitSettingParam sets ExitSettingParam field to given value.

### HasExitSettingParam

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasExitSettingParam() bool`

HasExitSettingParam returns a boolean if a field has been set.

### GetFloatPnl

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetFloatPnl() string`

GetFloatPnl returns the FloatPnl field if non-nil, zero value otherwise.

### GetFloatPnlOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetFloatPnlOk() (*string, bool)`

GetFloatPnlOk returns a tuple with the FloatPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatPnl

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetFloatPnl(v string)`

SetFloatPnl sets FloatPnl field to given value.

### HasFloatPnl

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasFloatPnl() bool`

HasFloatPnl returns a boolean if a field has been set.

### GetFrozenBal

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetFrozenBal() string`

GetFrozenBal returns the FrozenBal field if non-nil, zero value otherwise.

### GetFrozenBalOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetFrozenBalOk() (*string, bool)`

GetFrozenBalOk returns a tuple with the FrozenBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenBal

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetFrozenBal(v string)`

SetFrozenBal sets FrozenBal field to given value.

### HasFrozenBal

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasFrozenBal() bool`

HasFrozenBal returns a boolean if a field has been set.

### GetInstIds

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetInstIds() []string`

GetInstIds returns the InstIds field if non-nil, zero value otherwise.

### GetInstIdsOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetInstIdsOk() (*[]string, bool)`

GetInstIdsOk returns a tuple with the InstIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstIds

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetInstIds(v []string)`

SetInstIds sets InstIds field to given value.

### HasInstIds

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasInstIds() bool`

HasInstIds returns a boolean if a field has been set.

### GetInstType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetInvestAmt

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetInvestAmt() string`

GetInvestAmt returns the InvestAmt field if non-nil, zero value otherwise.

### GetInvestAmtOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetInvestAmtOk() (*string, bool)`

GetInvestAmtOk returns a tuple with the InvestAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestAmt

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetInvestAmt(v string)`

SetInvestAmt sets InvestAmt field to given value.

### HasInvestAmt

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasInvestAmt() bool`

HasInvestAmt returns a boolean if a field has been set.

### GetLever

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetRatio

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetRatio() string`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetRatioOk() (*string, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetRatio(v string)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetRealizedPnl

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetRealizedPnl() string`

GetRealizedPnl returns the RealizedPnl field if non-nil, zero value otherwise.

### GetRealizedPnlOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetRealizedPnlOk() (*string, bool)`

GetRealizedPnlOk returns a tuple with the RealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedPnl

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetRealizedPnl(v string)`

SetRealizedPnl sets RealizedPnl field to given value.

### HasRealizedPnl

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasRealizedPnl() bool`

HasRealizedPnl returns a boolean if a field has been set.

### GetSignalChanId

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetSignalChanId() string`

GetSignalChanId returns the SignalChanId field if non-nil, zero value otherwise.

### GetSignalChanIdOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetSignalChanIdOk() (*string, bool)`

GetSignalChanIdOk returns a tuple with the SignalChanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalChanId

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetSignalChanId(v string)`

SetSignalChanId sets SignalChanId field to given value.

### HasSignalChanId

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasSignalChanId() bool`

HasSignalChanId returns a boolean if a field has been set.

### GetSignalChanName

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetSignalChanName() string`

GetSignalChanName returns the SignalChanName field if non-nil, zero value otherwise.

### GetSignalChanNameOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetSignalChanNameOk() (*string, bool)`

GetSignalChanNameOk returns a tuple with the SignalChanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalChanName

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetSignalChanName(v string)`

SetSignalChanName sets SignalChanName field to given value.

### HasSignalChanName

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasSignalChanName() bool`

HasSignalChanName returns a boolean if a field has been set.

### GetSignalSourceType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetSignalSourceType() string`

GetSignalSourceType returns the SignalSourceType field if non-nil, zero value otherwise.

### GetSignalSourceTypeOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetSignalSourceTypeOk() (*string, bool)`

GetSignalSourceTypeOk returns a tuple with the SignalSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalSourceType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetSignalSourceType(v string)`

SetSignalSourceType sets SignalSourceType field to given value.

### HasSignalSourceType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasSignalSourceType() bool`

HasSignalSourceType returns a boolean if a field has been set.

### GetState

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubOrdType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetSubOrdType() string`

GetSubOrdType returns the SubOrdType field if non-nil, zero value otherwise.

### GetSubOrdTypeOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetSubOrdTypeOk() (*string, bool)`

GetSubOrdTypeOk returns a tuple with the SubOrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrdType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetSubOrdType(v string)`

SetSubOrdType sets SubOrdType field to given value.

### HasSubOrdType

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasSubOrdType() bool`

HasSubOrdType returns a boolean if a field has been set.

### GetTotalEq

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetTotalEq() string`

GetTotalEq returns the TotalEq field if non-nil, zero value otherwise.

### GetTotalEqOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetTotalEqOk() (*string, bool)`

GetTotalEqOk returns a tuple with the TotalEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEq

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetTotalEq(v string)`

SetTotalEq sets TotalEq field to given value.

### HasTotalEq

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasTotalEq() bool`

HasTotalEq returns a boolean if a field has been set.

### GetTotalPnl

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetTotalPnl() string`

GetTotalPnl returns the TotalPnl field if non-nil, zero value otherwise.

### GetTotalPnlOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetTotalPnlOk() (*string, bool)`

GetTotalPnlOk returns a tuple with the TotalPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPnl

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetTotalPnl(v string)`

SetTotalPnl sets TotalPnl field to given value.

### HasTotalPnl

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasTotalPnl() bool`

HasTotalPnl returns a boolean if a field has been set.

### GetTotalPnlRatio

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetTotalPnlRatio() string`

GetTotalPnlRatio returns the TotalPnlRatio field if non-nil, zero value otherwise.

### GetTotalPnlRatioOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetTotalPnlRatioOk() (*string, bool)`

GetTotalPnlRatioOk returns a tuple with the TotalPnlRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPnlRatio

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetTotalPnlRatio(v string)`

SetTotalPnlRatio sets TotalPnlRatio field to given value.

### HasTotalPnlRatio

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasTotalPnlRatio() bool`

HasTotalPnlRatio returns a boolean if a field has been set.

### GetUTime

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


