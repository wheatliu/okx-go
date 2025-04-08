# GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID | [optional] [default to ""]
**AlgoId** | Pointer to **string** | Algo ID | [optional] [default to ""]
**AlgoOrdType** | Pointer to **string** | Algo order type  &#x60;recurring&#x60;: recurring buy | [optional] [default to ""]
**Amt** | Pointer to **string** | Quantity invested per cycle | [optional] [default to ""]
**CTime** | Pointer to **string** | Algo order created time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Cycles** | Pointer to **string** | Accumulate recurring buy cycles | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**InvestmentAmt** | Pointer to **string** | Accumulate quantity invested | [optional] [default to ""]
**InvestmentCcy** | Pointer to **string** | The invested quantity unit, can only be &#x60;USDT&#x60;/&#x60;USDC&#x60; | [optional] [default to ""]
**MktCap** | Pointer to **string** | Market value in unit of &#x60;USDT&#x60; | [optional] [default to ""]
**Period** | Pointer to **string** | Period  &#x60;monthly&#x60;  &#x60;weekly&#x60;  &#x60;daily&#x60;  &#x60;hourly&#x60; | [optional] [default to ""]
**PnlRatio** | Pointer to **string** | Rate of yield | [optional] [default to ""]
**RecurringDay** | Pointer to **string** | Recurring buy date  When the period is &#x60;monthly&#x60;, the value range is an integer of [1,28]  When the period is &#x60;weekly&#x60;, the value range is an integer of [1,7] | [optional] [default to ""]
**RecurringHour** | Pointer to **string** | Recurring buy by hourly  &#x60;1&#x60;/&#x60;4&#x60;/&#x60;8&#x60;/&#x60;12&#x60;, e.g. &#x60;4&#x60; represents \&quot;recurring buy every 4 hour\&quot; | [optional] [default to ""]
**RecurringList** | Pointer to [**[]CreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner**](CreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner.md) | Recurring buy info | [optional] 
**RecurringTime** | Pointer to **string** | Recurring buy time, the value range is an integer of [0,23] | [optional] [default to ""]
**State** | Pointer to **string** | Algo order state  &#x60;stopped&#x60; | [optional] [default to ""]
**StgyName** | Pointer to **string** | Custom name for trading bot, no more than 40 characters | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**TimeZone** | Pointer to **string** | UTC time zone, the value range is an integer of [-12,14]  e.g. \&quot;8\&quot; representing UTC+8 (East 8 District), Beijing Time | [optional] [default to ""]
**TotalAnnRate** | Pointer to **string** | Total annualized rate of yield | [optional] [default to ""]
**TotalPnl** | Pointer to **string** | Total P&amp;L | [optional] [default to ""]
**UTime** | Pointer to **string** | Algo order updated time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner

`func NewGetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner() *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner`

NewGetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner instantiates a new GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotRecurringOrdersAlgoHistoryV5RespDataInnerWithDefaults

`func NewGetTradingBotRecurringOrdersAlgoHistoryV5RespDataInnerWithDefaults() *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner`

NewGetTradingBotRecurringOrdersAlgoHistoryV5RespDataInnerWithDefaults instantiates a new GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoClOrdId

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAlgoId

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetAlgoOrdType

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetAlgoOrdType() string`

GetAlgoOrdType returns the AlgoOrdType field if non-nil, zero value otherwise.

### GetAlgoOrdTypeOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetAlgoOrdTypeOk() (*string, bool)`

GetAlgoOrdTypeOk returns a tuple with the AlgoOrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoOrdType

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetAlgoOrdType(v string)`

SetAlgoOrdType sets AlgoOrdType field to given value.

### HasAlgoOrdType

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasAlgoOrdType() bool`

HasAlgoOrdType returns a boolean if a field has been set.

### GetAmt

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetCTime

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCycles

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetCycles() string`

GetCycles returns the Cycles field if non-nil, zero value otherwise.

### GetCyclesOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetCyclesOk() (*string, bool)`

GetCyclesOk returns a tuple with the Cycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycles

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetCycles(v string)`

SetCycles sets Cycles field to given value.

### HasCycles

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasCycles() bool`

HasCycles returns a boolean if a field has been set.

### GetInstType

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetInvestmentAmt

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetInvestmentAmt() string`

GetInvestmentAmt returns the InvestmentAmt field if non-nil, zero value otherwise.

### GetInvestmentAmtOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetInvestmentAmtOk() (*string, bool)`

GetInvestmentAmtOk returns a tuple with the InvestmentAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentAmt

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetInvestmentAmt(v string)`

SetInvestmentAmt sets InvestmentAmt field to given value.

### HasInvestmentAmt

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasInvestmentAmt() bool`

HasInvestmentAmt returns a boolean if a field has been set.

### GetInvestmentCcy

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetInvestmentCcy() string`

GetInvestmentCcy returns the InvestmentCcy field if non-nil, zero value otherwise.

### GetInvestmentCcyOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetInvestmentCcyOk() (*string, bool)`

GetInvestmentCcyOk returns a tuple with the InvestmentCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentCcy

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetInvestmentCcy(v string)`

SetInvestmentCcy sets InvestmentCcy field to given value.

### HasInvestmentCcy

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasInvestmentCcy() bool`

HasInvestmentCcy returns a boolean if a field has been set.

### GetMktCap

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetMktCap() string`

GetMktCap returns the MktCap field if non-nil, zero value otherwise.

### GetMktCapOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetMktCapOk() (*string, bool)`

GetMktCapOk returns a tuple with the MktCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMktCap

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetMktCap(v string)`

SetMktCap sets MktCap field to given value.

### HasMktCap

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasMktCap() bool`

HasMktCap returns a boolean if a field has been set.

### GetPeriod

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPnlRatio

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetPnlRatio() string`

GetPnlRatio returns the PnlRatio field if non-nil, zero value otherwise.

### GetPnlRatioOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetPnlRatioOk() (*string, bool)`

GetPnlRatioOk returns a tuple with the PnlRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnlRatio

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetPnlRatio(v string)`

SetPnlRatio sets PnlRatio field to given value.

### HasPnlRatio

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasPnlRatio() bool`

HasPnlRatio returns a boolean if a field has been set.

### GetRecurringDay

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetRecurringDay() string`

GetRecurringDay returns the RecurringDay field if non-nil, zero value otherwise.

### GetRecurringDayOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetRecurringDayOk() (*string, bool)`

GetRecurringDayOk returns a tuple with the RecurringDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDay

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetRecurringDay(v string)`

SetRecurringDay sets RecurringDay field to given value.

### HasRecurringDay

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasRecurringDay() bool`

HasRecurringDay returns a boolean if a field has been set.

### GetRecurringHour

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetRecurringHour() string`

GetRecurringHour returns the RecurringHour field if non-nil, zero value otherwise.

### GetRecurringHourOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetRecurringHourOk() (*string, bool)`

GetRecurringHourOk returns a tuple with the RecurringHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringHour

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetRecurringHour(v string)`

SetRecurringHour sets RecurringHour field to given value.

### HasRecurringHour

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasRecurringHour() bool`

HasRecurringHour returns a boolean if a field has been set.

### GetRecurringList

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetRecurringList() []CreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner`

GetRecurringList returns the RecurringList field if non-nil, zero value otherwise.

### GetRecurringListOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetRecurringListOk() (*[]CreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner, bool)`

GetRecurringListOk returns a tuple with the RecurringList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringList

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetRecurringList(v []CreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner)`

SetRecurringList sets RecurringList field to given value.

### HasRecurringList

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasRecurringList() bool`

HasRecurringList returns a boolean if a field has been set.

### GetRecurringTime

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetRecurringTime() string`

GetRecurringTime returns the RecurringTime field if non-nil, zero value otherwise.

### GetRecurringTimeOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetRecurringTimeOk() (*string, bool)`

GetRecurringTimeOk returns a tuple with the RecurringTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringTime

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetRecurringTime(v string)`

SetRecurringTime sets RecurringTime field to given value.

### HasRecurringTime

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasRecurringTime() bool`

HasRecurringTime returns a boolean if a field has been set.

### GetState

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStgyName

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetStgyName() string`

GetStgyName returns the StgyName field if non-nil, zero value otherwise.

### GetStgyNameOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetStgyNameOk() (*string, bool)`

GetStgyNameOk returns a tuple with the StgyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStgyName

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetStgyName(v string)`

SetStgyName sets StgyName field to given value.

### HasStgyName

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasStgyName() bool`

HasStgyName returns a boolean if a field has been set.

### GetTag

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTimeZone

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetTotalAnnRate

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetTotalAnnRate() string`

GetTotalAnnRate returns the TotalAnnRate field if non-nil, zero value otherwise.

### GetTotalAnnRateOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetTotalAnnRateOk() (*string, bool)`

GetTotalAnnRateOk returns a tuple with the TotalAnnRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAnnRate

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetTotalAnnRate(v string)`

SetTotalAnnRate sets TotalAnnRate field to given value.

### HasTotalAnnRate

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasTotalAnnRate() bool`

HasTotalAnnRate returns a boolean if a field has been set.

### GetTotalPnl

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetTotalPnl() string`

GetTotalPnl returns the TotalPnl field if non-nil, zero value otherwise.

### GetTotalPnlOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetTotalPnlOk() (*string, bool)`

GetTotalPnlOk returns a tuple with the TotalPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPnl

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetTotalPnl(v string)`

SetTotalPnl sets TotalPnl field to given value.

### HasTotalPnl

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasTotalPnl() bool`

HasTotalPnl returns a boolean if a field has been set.

### GetUTime

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetTradingBotRecurringOrdersAlgoHistoryV5RespDataInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


