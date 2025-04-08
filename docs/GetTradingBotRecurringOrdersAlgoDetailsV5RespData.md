# GetTradingBotRecurringOrdersAlgoDetailsV5RespData

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
**NextInvestTime** | Pointer to **string** | Next invest time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Period** | Pointer to **string** | Period  &#x60;monthly&#x60;  &#x60;weekly&#x60;  &#x60;daily&#x60;  &#x60;hourly&#x60; | [optional] [default to ""]
**PnlRatio** | Pointer to **string** | Rate of yield | [optional] [default to ""]
**RecurringDay** | Pointer to **string** | Recurring buy date  When the period is &#x60;monthly&#x60;, the value range is an integer of [1,28]  When the period is &#x60;weekly&#x60;, the value range is an integer of [1,7] | [optional] [default to ""]
**RecurringHour** | Pointer to **string** | Recurring buy by hourly  &#x60;1&#x60;/&#x60;4&#x60;/&#x60;8&#x60;/&#x60;12&#x60;, e.g. &#x60;4&#x60; represents \&quot;recurring buy every 4 hour\&quot; | [optional] [default to ""]
**RecurringList** | Pointer to [**[]GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner**](GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner.md) | Recurring buy info | [optional] 
**RecurringTime** | Pointer to **string** | Recurring buy time, the value range is an integer of [0,23] | [optional] [default to ""]
**State** | Pointer to **string** | Algo order state  &#x60;running&#x60;  &#x60;stopping&#x60;  &#x60;stopped&#x60;  &#x60;pause&#x60; | [optional] [default to ""]
**StgyName** | Pointer to **string** | Custom name for trading bot, no more than 40 characters | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**TimeZone** | Pointer to **string** | UTC time zone, the value range is an integer of [-12,14]  e.g. \&quot;8\&quot; representing UTC+8 (East 8 District), Beijing Time | [optional] [default to ""]
**TotalAnnRate** | Pointer to **string** | Total annualized rate of yield | [optional] [default to ""]
**TotalPnl** | Pointer to **string** | Total P&amp;L | [optional] [default to ""]
**UTime** | Pointer to **string** | Algo order updated time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetTradingBotRecurringOrdersAlgoDetailsV5RespData

`func NewGetTradingBotRecurringOrdersAlgoDetailsV5RespData() *GetTradingBotRecurringOrdersAlgoDetailsV5RespData`

NewGetTradingBotRecurringOrdersAlgoDetailsV5RespData instantiates a new GetTradingBotRecurringOrdersAlgoDetailsV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotRecurringOrdersAlgoDetailsV5RespDataWithDefaults

`func NewGetTradingBotRecurringOrdersAlgoDetailsV5RespDataWithDefaults() *GetTradingBotRecurringOrdersAlgoDetailsV5RespData`

NewGetTradingBotRecurringOrdersAlgoDetailsV5RespDataWithDefaults instantiates a new GetTradingBotRecurringOrdersAlgoDetailsV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoClOrdId

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAlgoId

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetAlgoOrdType

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetAlgoOrdType() string`

GetAlgoOrdType returns the AlgoOrdType field if non-nil, zero value otherwise.

### GetAlgoOrdTypeOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetAlgoOrdTypeOk() (*string, bool)`

GetAlgoOrdTypeOk returns a tuple with the AlgoOrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoOrdType

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetAlgoOrdType(v string)`

SetAlgoOrdType sets AlgoOrdType field to given value.

### HasAlgoOrdType

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasAlgoOrdType() bool`

HasAlgoOrdType returns a boolean if a field has been set.

### GetAmt

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetCTime

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCycles

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetCycles() string`

GetCycles returns the Cycles field if non-nil, zero value otherwise.

### GetCyclesOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetCyclesOk() (*string, bool)`

GetCyclesOk returns a tuple with the Cycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycles

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetCycles(v string)`

SetCycles sets Cycles field to given value.

### HasCycles

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasCycles() bool`

HasCycles returns a boolean if a field has been set.

### GetInstType

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetInvestmentAmt

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetInvestmentAmt() string`

GetInvestmentAmt returns the InvestmentAmt field if non-nil, zero value otherwise.

### GetInvestmentAmtOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetInvestmentAmtOk() (*string, bool)`

GetInvestmentAmtOk returns a tuple with the InvestmentAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentAmt

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetInvestmentAmt(v string)`

SetInvestmentAmt sets InvestmentAmt field to given value.

### HasInvestmentAmt

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasInvestmentAmt() bool`

HasInvestmentAmt returns a boolean if a field has been set.

### GetInvestmentCcy

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetInvestmentCcy() string`

GetInvestmentCcy returns the InvestmentCcy field if non-nil, zero value otherwise.

### GetInvestmentCcyOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetInvestmentCcyOk() (*string, bool)`

GetInvestmentCcyOk returns a tuple with the InvestmentCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentCcy

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetInvestmentCcy(v string)`

SetInvestmentCcy sets InvestmentCcy field to given value.

### HasInvestmentCcy

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasInvestmentCcy() bool`

HasInvestmentCcy returns a boolean if a field has been set.

### GetMktCap

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetMktCap() string`

GetMktCap returns the MktCap field if non-nil, zero value otherwise.

### GetMktCapOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetMktCapOk() (*string, bool)`

GetMktCapOk returns a tuple with the MktCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMktCap

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetMktCap(v string)`

SetMktCap sets MktCap field to given value.

### HasMktCap

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasMktCap() bool`

HasMktCap returns a boolean if a field has been set.

### GetNextInvestTime

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetNextInvestTime() string`

GetNextInvestTime returns the NextInvestTime field if non-nil, zero value otherwise.

### GetNextInvestTimeOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetNextInvestTimeOk() (*string, bool)`

GetNextInvestTimeOk returns a tuple with the NextInvestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvestTime

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetNextInvestTime(v string)`

SetNextInvestTime sets NextInvestTime field to given value.

### HasNextInvestTime

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasNextInvestTime() bool`

HasNextInvestTime returns a boolean if a field has been set.

### GetPeriod

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPnlRatio

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetPnlRatio() string`

GetPnlRatio returns the PnlRatio field if non-nil, zero value otherwise.

### GetPnlRatioOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetPnlRatioOk() (*string, bool)`

GetPnlRatioOk returns a tuple with the PnlRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnlRatio

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetPnlRatio(v string)`

SetPnlRatio sets PnlRatio field to given value.

### HasPnlRatio

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasPnlRatio() bool`

HasPnlRatio returns a boolean if a field has been set.

### GetRecurringDay

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetRecurringDay() string`

GetRecurringDay returns the RecurringDay field if non-nil, zero value otherwise.

### GetRecurringDayOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetRecurringDayOk() (*string, bool)`

GetRecurringDayOk returns a tuple with the RecurringDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDay

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetRecurringDay(v string)`

SetRecurringDay sets RecurringDay field to given value.

### HasRecurringDay

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasRecurringDay() bool`

HasRecurringDay returns a boolean if a field has been set.

### GetRecurringHour

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetRecurringHour() string`

GetRecurringHour returns the RecurringHour field if non-nil, zero value otherwise.

### GetRecurringHourOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetRecurringHourOk() (*string, bool)`

GetRecurringHourOk returns a tuple with the RecurringHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringHour

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetRecurringHour(v string)`

SetRecurringHour sets RecurringHour field to given value.

### HasRecurringHour

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasRecurringHour() bool`

HasRecurringHour returns a boolean if a field has been set.

### GetRecurringList

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetRecurringList() []GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner`

GetRecurringList returns the RecurringList field if non-nil, zero value otherwise.

### GetRecurringListOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetRecurringListOk() (*[]GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner, bool)`

GetRecurringListOk returns a tuple with the RecurringList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringList

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetRecurringList(v []GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner)`

SetRecurringList sets RecurringList field to given value.

### HasRecurringList

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasRecurringList() bool`

HasRecurringList returns a boolean if a field has been set.

### GetRecurringTime

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetRecurringTime() string`

GetRecurringTime returns the RecurringTime field if non-nil, zero value otherwise.

### GetRecurringTimeOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetRecurringTimeOk() (*string, bool)`

GetRecurringTimeOk returns a tuple with the RecurringTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringTime

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetRecurringTime(v string)`

SetRecurringTime sets RecurringTime field to given value.

### HasRecurringTime

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasRecurringTime() bool`

HasRecurringTime returns a boolean if a field has been set.

### GetState

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStgyName

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetStgyName() string`

GetStgyName returns the StgyName field if non-nil, zero value otherwise.

### GetStgyNameOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetStgyNameOk() (*string, bool)`

GetStgyNameOk returns a tuple with the StgyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStgyName

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetStgyName(v string)`

SetStgyName sets StgyName field to given value.

### HasStgyName

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasStgyName() bool`

HasStgyName returns a boolean if a field has been set.

### GetTag

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTimeZone

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetTotalAnnRate

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetTotalAnnRate() string`

GetTotalAnnRate returns the TotalAnnRate field if non-nil, zero value otherwise.

### GetTotalAnnRateOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetTotalAnnRateOk() (*string, bool)`

GetTotalAnnRateOk returns a tuple with the TotalAnnRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAnnRate

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetTotalAnnRate(v string)`

SetTotalAnnRate sets TotalAnnRate field to given value.

### HasTotalAnnRate

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasTotalAnnRate() bool`

HasTotalAnnRate returns a boolean if a field has been set.

### GetTotalPnl

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetTotalPnl() string`

GetTotalPnl returns the TotalPnl field if non-nil, zero value otherwise.

### GetTotalPnlOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetTotalPnlOk() (*string, bool)`

GetTotalPnlOk returns a tuple with the TotalPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPnl

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetTotalPnl(v string)`

SetTotalPnl sets TotalPnl field to given value.

### HasTotalPnl

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasTotalPnl() bool`

HasTotalPnl returns a boolean if a field has been set.

### GetUTime

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespData) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


