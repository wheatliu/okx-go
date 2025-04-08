# GetTradingBotRecurringOrdersAlgoPendingV5RespData

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
**State** | Pointer to **string** | Algo order state  &#x60;running&#x60;  &#x60;stopping&#x60;  &#x60;pause&#x60; | [optional] [default to ""]
**StgyName** | Pointer to **string** | Custom name for trading bot, no more than 40 characters | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**TimeZone** | Pointer to **string** | UTC time zone, the value range is an integer of [-12,14]  e.g. \&quot;8\&quot; representing UTC+8 (East 8 District), Beijing Time | [optional] [default to ""]
**TotalAnnRate** | Pointer to **string** | Total annualized rate of yield | [optional] [default to ""]
**TotalPnl** | Pointer to **string** | Total P&amp;L | [optional] [default to ""]
**UTime** | Pointer to **string** | Algo order updated time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetTradingBotRecurringOrdersAlgoPendingV5RespData

`func NewGetTradingBotRecurringOrdersAlgoPendingV5RespData() *GetTradingBotRecurringOrdersAlgoPendingV5RespData`

NewGetTradingBotRecurringOrdersAlgoPendingV5RespData instantiates a new GetTradingBotRecurringOrdersAlgoPendingV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotRecurringOrdersAlgoPendingV5RespDataWithDefaults

`func NewGetTradingBotRecurringOrdersAlgoPendingV5RespDataWithDefaults() *GetTradingBotRecurringOrdersAlgoPendingV5RespData`

NewGetTradingBotRecurringOrdersAlgoPendingV5RespDataWithDefaults instantiates a new GetTradingBotRecurringOrdersAlgoPendingV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoClOrdId

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAlgoId

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetAlgoOrdType

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetAlgoOrdType() string`

GetAlgoOrdType returns the AlgoOrdType field if non-nil, zero value otherwise.

### GetAlgoOrdTypeOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetAlgoOrdTypeOk() (*string, bool)`

GetAlgoOrdTypeOk returns a tuple with the AlgoOrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoOrdType

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetAlgoOrdType(v string)`

SetAlgoOrdType sets AlgoOrdType field to given value.

### HasAlgoOrdType

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasAlgoOrdType() bool`

HasAlgoOrdType returns a boolean if a field has been set.

### GetAmt

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetCTime

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCycles

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetCycles() string`

GetCycles returns the Cycles field if non-nil, zero value otherwise.

### GetCyclesOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetCyclesOk() (*string, bool)`

GetCyclesOk returns a tuple with the Cycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycles

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetCycles(v string)`

SetCycles sets Cycles field to given value.

### HasCycles

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasCycles() bool`

HasCycles returns a boolean if a field has been set.

### GetInstType

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetInvestmentAmt

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetInvestmentAmt() string`

GetInvestmentAmt returns the InvestmentAmt field if non-nil, zero value otherwise.

### GetInvestmentAmtOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetInvestmentAmtOk() (*string, bool)`

GetInvestmentAmtOk returns a tuple with the InvestmentAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentAmt

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetInvestmentAmt(v string)`

SetInvestmentAmt sets InvestmentAmt field to given value.

### HasInvestmentAmt

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasInvestmentAmt() bool`

HasInvestmentAmt returns a boolean if a field has been set.

### GetInvestmentCcy

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetInvestmentCcy() string`

GetInvestmentCcy returns the InvestmentCcy field if non-nil, zero value otherwise.

### GetInvestmentCcyOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetInvestmentCcyOk() (*string, bool)`

GetInvestmentCcyOk returns a tuple with the InvestmentCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentCcy

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetInvestmentCcy(v string)`

SetInvestmentCcy sets InvestmentCcy field to given value.

### HasInvestmentCcy

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasInvestmentCcy() bool`

HasInvestmentCcy returns a boolean if a field has been set.

### GetMktCap

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetMktCap() string`

GetMktCap returns the MktCap field if non-nil, zero value otherwise.

### GetMktCapOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetMktCapOk() (*string, bool)`

GetMktCapOk returns a tuple with the MktCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMktCap

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetMktCap(v string)`

SetMktCap sets MktCap field to given value.

### HasMktCap

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasMktCap() bool`

HasMktCap returns a boolean if a field has been set.

### GetPeriod

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPnlRatio

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetPnlRatio() string`

GetPnlRatio returns the PnlRatio field if non-nil, zero value otherwise.

### GetPnlRatioOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetPnlRatioOk() (*string, bool)`

GetPnlRatioOk returns a tuple with the PnlRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnlRatio

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetPnlRatio(v string)`

SetPnlRatio sets PnlRatio field to given value.

### HasPnlRatio

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasPnlRatio() bool`

HasPnlRatio returns a boolean if a field has been set.

### GetRecurringDay

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetRecurringDay() string`

GetRecurringDay returns the RecurringDay field if non-nil, zero value otherwise.

### GetRecurringDayOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetRecurringDayOk() (*string, bool)`

GetRecurringDayOk returns a tuple with the RecurringDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDay

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetRecurringDay(v string)`

SetRecurringDay sets RecurringDay field to given value.

### HasRecurringDay

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasRecurringDay() bool`

HasRecurringDay returns a boolean if a field has been set.

### GetRecurringHour

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetRecurringHour() string`

GetRecurringHour returns the RecurringHour field if non-nil, zero value otherwise.

### GetRecurringHourOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetRecurringHourOk() (*string, bool)`

GetRecurringHourOk returns a tuple with the RecurringHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringHour

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetRecurringHour(v string)`

SetRecurringHour sets RecurringHour field to given value.

### HasRecurringHour

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasRecurringHour() bool`

HasRecurringHour returns a boolean if a field has been set.

### GetRecurringList

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetRecurringList() []CreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner`

GetRecurringList returns the RecurringList field if non-nil, zero value otherwise.

### GetRecurringListOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetRecurringListOk() (*[]CreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner, bool)`

GetRecurringListOk returns a tuple with the RecurringList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringList

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetRecurringList(v []CreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner)`

SetRecurringList sets RecurringList field to given value.

### HasRecurringList

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasRecurringList() bool`

HasRecurringList returns a boolean if a field has been set.

### GetRecurringTime

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetRecurringTime() string`

GetRecurringTime returns the RecurringTime field if non-nil, zero value otherwise.

### GetRecurringTimeOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetRecurringTimeOk() (*string, bool)`

GetRecurringTimeOk returns a tuple with the RecurringTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringTime

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetRecurringTime(v string)`

SetRecurringTime sets RecurringTime field to given value.

### HasRecurringTime

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasRecurringTime() bool`

HasRecurringTime returns a boolean if a field has been set.

### GetState

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStgyName

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetStgyName() string`

GetStgyName returns the StgyName field if non-nil, zero value otherwise.

### GetStgyNameOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetStgyNameOk() (*string, bool)`

GetStgyNameOk returns a tuple with the StgyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStgyName

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetStgyName(v string)`

SetStgyName sets StgyName field to given value.

### HasStgyName

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasStgyName() bool`

HasStgyName returns a boolean if a field has been set.

### GetTag

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTimeZone

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetTotalAnnRate

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetTotalAnnRate() string`

GetTotalAnnRate returns the TotalAnnRate field if non-nil, zero value otherwise.

### GetTotalAnnRateOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetTotalAnnRateOk() (*string, bool)`

GetTotalAnnRateOk returns a tuple with the TotalAnnRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAnnRate

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetTotalAnnRate(v string)`

SetTotalAnnRate sets TotalAnnRate field to given value.

### HasTotalAnnRate

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasTotalAnnRate() bool`

HasTotalAnnRate returns a boolean if a field has been set.

### GetTotalPnl

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetTotalPnl() string`

GetTotalPnl returns the TotalPnl field if non-nil, zero value otherwise.

### GetTotalPnlOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetTotalPnlOk() (*string, bool)`

GetTotalPnlOk returns a tuple with the TotalPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPnl

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetTotalPnl(v string)`

SetTotalPnl sets TotalPnl field to given value.

### HasTotalPnl

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasTotalPnl() bool`

HasTotalPnl returns a boolean if a field has been set.

### GetUTime

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespData) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


