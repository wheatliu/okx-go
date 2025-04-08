/*
Okx Rest API

OpenAPI specification for Okx cryptocurrency exchange - Rest API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rest

import (
	"encoding/json"
)

// checks if the GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner{}

// GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner struct for GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner
type GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner struct {
	// Client-supplied Algo ID
	AlgoClOrdId *string `json:"algoClOrdId,omitempty"`
	// Algo ID
	AlgoId *string `json:"algoId,omitempty"`
	// Algo order type  `recurring`: recurring buy
	AlgoOrdType *string `json:"algoOrdType,omitempty"`
	// Quantity invested per cycle
	Amt *string `json:"amt,omitempty"`
	// Algo order created time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	CTime *string `json:"cTime,omitempty"`
	// Accumulate recurring buy cycles
	Cycles *string `json:"cycles,omitempty"`
	// Instrument type
	InstType *string `json:"instType,omitempty"`
	// Accumulate quantity invested
	InvestmentAmt *string `json:"investmentAmt,omitempty"`
	// The invested quantity unit, can only be `USDT`/`USDC`
	InvestmentCcy *string `json:"investmentCcy,omitempty"`
	// Market value in unit of `USDT`
	MktCap *string `json:"mktCap,omitempty"`
	// Period  `monthly`  `weekly`  `daily`  `hourly`
	Period *string `json:"period,omitempty"`
	// Rate of yield
	PnlRatio *string `json:"pnlRatio,omitempty"`
	// Recurring buy date  When the period is `monthly`, the value range is an integer of [1,28]  When the period is `weekly`, the value range is an integer of [1,7]
	RecurringDay *string `json:"recurringDay,omitempty"`
	// Recurring buy by hourly  `1`/`4`/`8`/`12`, e.g. `4` represents \"recurring buy every 4 hour\"
	RecurringHour *string `json:"recurringHour,omitempty"`
	// Recurring buy info
	RecurringList []CreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner `json:"recurringList,omitempty"`
	// Recurring buy time, the value range is an integer of [0,23]
	RecurringTime *string `json:"recurringTime,omitempty"`
	// Algo order state  `running`  `stopping`  `pause`
	State *string `json:"state,omitempty"`
	// Custom name for trading bot, no more than 40 characters
	StgyName *string `json:"stgyName,omitempty"`
	// Order tag
	Tag *string `json:"tag,omitempty"`
	// UTC time zone, the value range is an integer of [-12,14]  e.g. \"8\" representing UTC+8 (East 8 District), Beijing Time
	TimeZone *string `json:"timeZone,omitempty"`
	// Total annualized rate of yield
	TotalAnnRate *string `json:"totalAnnRate,omitempty"`
	// Total P&L
	TotalPnl *string `json:"totalPnl,omitempty"`
	// Algo order updated time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	UTime *string `json:"uTime,omitempty"`
}

// NewGetTradingBotRecurringOrdersAlgoPendingV5RespDataInner instantiates a new GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTradingBotRecurringOrdersAlgoPendingV5RespDataInner() *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner {
	this := GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner{}
	var algoClOrdId string = ""
	this.AlgoClOrdId = &algoClOrdId
	var algoId string = ""
	this.AlgoId = &algoId
	var algoOrdType string = ""
	this.AlgoOrdType = &algoOrdType
	var amt string = ""
	this.Amt = &amt
	var cTime string = ""
	this.CTime = &cTime
	var cycles string = ""
	this.Cycles = &cycles
	var instType string = ""
	this.InstType = &instType
	var investmentAmt string = ""
	this.InvestmentAmt = &investmentAmt
	var investmentCcy string = ""
	this.InvestmentCcy = &investmentCcy
	var mktCap string = ""
	this.MktCap = &mktCap
	var period string = ""
	this.Period = &period
	var pnlRatio string = ""
	this.PnlRatio = &pnlRatio
	var recurringDay string = ""
	this.RecurringDay = &recurringDay
	var recurringHour string = ""
	this.RecurringHour = &recurringHour
	var recurringTime string = ""
	this.RecurringTime = &recurringTime
	var state string = ""
	this.State = &state
	var stgyName string = ""
	this.StgyName = &stgyName
	var tag string = ""
	this.Tag = &tag
	var timeZone string = ""
	this.TimeZone = &timeZone
	var totalAnnRate string = ""
	this.TotalAnnRate = &totalAnnRate
	var totalPnl string = ""
	this.TotalPnl = &totalPnl
	var uTime string = ""
	this.UTime = &uTime
	return &this
}

// NewGetTradingBotRecurringOrdersAlgoPendingV5RespDataInnerWithDefaults instantiates a new GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTradingBotRecurringOrdersAlgoPendingV5RespDataInnerWithDefaults() *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner {
	this := GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner{}
	var algoClOrdId string = ""
	this.AlgoClOrdId = &algoClOrdId
	var algoId string = ""
	this.AlgoId = &algoId
	var algoOrdType string = ""
	this.AlgoOrdType = &algoOrdType
	var amt string = ""
	this.Amt = &amt
	var cTime string = ""
	this.CTime = &cTime
	var cycles string = ""
	this.Cycles = &cycles
	var instType string = ""
	this.InstType = &instType
	var investmentAmt string = ""
	this.InvestmentAmt = &investmentAmt
	var investmentCcy string = ""
	this.InvestmentCcy = &investmentCcy
	var mktCap string = ""
	this.MktCap = &mktCap
	var period string = ""
	this.Period = &period
	var pnlRatio string = ""
	this.PnlRatio = &pnlRatio
	var recurringDay string = ""
	this.RecurringDay = &recurringDay
	var recurringHour string = ""
	this.RecurringHour = &recurringHour
	var recurringTime string = ""
	this.RecurringTime = &recurringTime
	var state string = ""
	this.State = &state
	var stgyName string = ""
	this.StgyName = &stgyName
	var tag string = ""
	this.Tag = &tag
	var timeZone string = ""
	this.TimeZone = &timeZone
	var totalAnnRate string = ""
	this.TotalAnnRate = &totalAnnRate
	var totalPnl string = ""
	this.TotalPnl = &totalPnl
	var uTime string = ""
	this.UTime = &uTime
	return &this
}

// GetAlgoClOrdId returns the AlgoClOrdId field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetAlgoClOrdId() string {
	if o == nil || IsNil(o.AlgoClOrdId) {
		var ret string
		return ret
	}
	return *o.AlgoClOrdId
}

// GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetAlgoClOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoClOrdId) {
		return nil, false
	}
	return o.AlgoClOrdId, true
}

// HasAlgoClOrdId returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasAlgoClOrdId() bool {
	if o != nil && !IsNil(o.AlgoClOrdId) {
		return true
	}

	return false
}

// SetAlgoClOrdId gets a reference to the given string and assigns it to the AlgoClOrdId field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetAlgoClOrdId(v string) {
	o.AlgoClOrdId = &v
}

// GetAlgoId returns the AlgoId field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetAlgoId() string {
	if o == nil || IsNil(o.AlgoId) {
		var ret string
		return ret
	}
	return *o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetAlgoIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoId) {
		return nil, false
	}
	return o.AlgoId, true
}

// HasAlgoId returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasAlgoId() bool {
	if o != nil && !IsNil(o.AlgoId) {
		return true
	}

	return false
}

// SetAlgoId gets a reference to the given string and assigns it to the AlgoId field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetAlgoId(v string) {
	o.AlgoId = &v
}

// GetAlgoOrdType returns the AlgoOrdType field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetAlgoOrdType() string {
	if o == nil || IsNil(o.AlgoOrdType) {
		var ret string
		return ret
	}
	return *o.AlgoOrdType
}

// GetAlgoOrdTypeOk returns a tuple with the AlgoOrdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetAlgoOrdTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoOrdType) {
		return nil, false
	}
	return o.AlgoOrdType, true
}

// HasAlgoOrdType returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasAlgoOrdType() bool {
	if o != nil && !IsNil(o.AlgoOrdType) {
		return true
	}

	return false
}

// SetAlgoOrdType gets a reference to the given string and assigns it to the AlgoOrdType field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetAlgoOrdType(v string) {
	o.AlgoOrdType = &v
}

// GetAmt returns the Amt field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetAmt() string {
	if o == nil || IsNil(o.Amt) {
		var ret string
		return ret
	}
	return *o.Amt
}

// GetAmtOk returns a tuple with the Amt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetAmtOk() (*string, bool) {
	if o == nil || IsNil(o.Amt) {
		return nil, false
	}
	return o.Amt, true
}

// HasAmt returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasAmt() bool {
	if o != nil && !IsNil(o.Amt) {
		return true
	}

	return false
}

// SetAmt gets a reference to the given string and assigns it to the Amt field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetAmt(v string) {
	o.Amt = &v
}

// GetCTime returns the CTime field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetCTime() string {
	if o == nil || IsNil(o.CTime) {
		var ret string
		return ret
	}
	return *o.CTime
}

// GetCTimeOk returns a tuple with the CTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetCTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CTime) {
		return nil, false
	}
	return o.CTime, true
}

// HasCTime returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasCTime() bool {
	if o != nil && !IsNil(o.CTime) {
		return true
	}

	return false
}

// SetCTime gets a reference to the given string and assigns it to the CTime field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetCTime(v string) {
	o.CTime = &v
}

// GetCycles returns the Cycles field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetCycles() string {
	if o == nil || IsNil(o.Cycles) {
		var ret string
		return ret
	}
	return *o.Cycles
}

// GetCyclesOk returns a tuple with the Cycles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetCyclesOk() (*string, bool) {
	if o == nil || IsNil(o.Cycles) {
		return nil, false
	}
	return o.Cycles, true
}

// HasCycles returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasCycles() bool {
	if o != nil && !IsNil(o.Cycles) {
		return true
	}

	return false
}

// SetCycles gets a reference to the given string and assigns it to the Cycles field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetCycles(v string) {
	o.Cycles = &v
}

// GetInstType returns the InstType field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetInstType() string {
	if o == nil || IsNil(o.InstType) {
		var ret string
		return ret
	}
	return *o.InstType
}

// GetInstTypeOk returns a tuple with the InstType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetInstTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstType) {
		return nil, false
	}
	return o.InstType, true
}

// HasInstType returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasInstType() bool {
	if o != nil && !IsNil(o.InstType) {
		return true
	}

	return false
}

// SetInstType gets a reference to the given string and assigns it to the InstType field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetInstType(v string) {
	o.InstType = &v
}

// GetInvestmentAmt returns the InvestmentAmt field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetInvestmentAmt() string {
	if o == nil || IsNil(o.InvestmentAmt) {
		var ret string
		return ret
	}
	return *o.InvestmentAmt
}

// GetInvestmentAmtOk returns a tuple with the InvestmentAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetInvestmentAmtOk() (*string, bool) {
	if o == nil || IsNil(o.InvestmentAmt) {
		return nil, false
	}
	return o.InvestmentAmt, true
}

// HasInvestmentAmt returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasInvestmentAmt() bool {
	if o != nil && !IsNil(o.InvestmentAmt) {
		return true
	}

	return false
}

// SetInvestmentAmt gets a reference to the given string and assigns it to the InvestmentAmt field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetInvestmentAmt(v string) {
	o.InvestmentAmt = &v
}

// GetInvestmentCcy returns the InvestmentCcy field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetInvestmentCcy() string {
	if o == nil || IsNil(o.InvestmentCcy) {
		var ret string
		return ret
	}
	return *o.InvestmentCcy
}

// GetInvestmentCcyOk returns a tuple with the InvestmentCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetInvestmentCcyOk() (*string, bool) {
	if o == nil || IsNil(o.InvestmentCcy) {
		return nil, false
	}
	return o.InvestmentCcy, true
}

// HasInvestmentCcy returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasInvestmentCcy() bool {
	if o != nil && !IsNil(o.InvestmentCcy) {
		return true
	}

	return false
}

// SetInvestmentCcy gets a reference to the given string and assigns it to the InvestmentCcy field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetInvestmentCcy(v string) {
	o.InvestmentCcy = &v
}

// GetMktCap returns the MktCap field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetMktCap() string {
	if o == nil || IsNil(o.MktCap) {
		var ret string
		return ret
	}
	return *o.MktCap
}

// GetMktCapOk returns a tuple with the MktCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetMktCapOk() (*string, bool) {
	if o == nil || IsNil(o.MktCap) {
		return nil, false
	}
	return o.MktCap, true
}

// HasMktCap returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasMktCap() bool {
	if o != nil && !IsNil(o.MktCap) {
		return true
	}

	return false
}

// SetMktCap gets a reference to the given string and assigns it to the MktCap field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetMktCap(v string) {
	o.MktCap = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetPeriod() string {
	if o == nil || IsNil(o.Period) {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetPeriod(v string) {
	o.Period = &v
}

// GetPnlRatio returns the PnlRatio field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetPnlRatio() string {
	if o == nil || IsNil(o.PnlRatio) {
		var ret string
		return ret
	}
	return *o.PnlRatio
}

// GetPnlRatioOk returns a tuple with the PnlRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetPnlRatioOk() (*string, bool) {
	if o == nil || IsNil(o.PnlRatio) {
		return nil, false
	}
	return o.PnlRatio, true
}

// HasPnlRatio returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasPnlRatio() bool {
	if o != nil && !IsNil(o.PnlRatio) {
		return true
	}

	return false
}

// SetPnlRatio gets a reference to the given string and assigns it to the PnlRatio field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetPnlRatio(v string) {
	o.PnlRatio = &v
}

// GetRecurringDay returns the RecurringDay field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetRecurringDay() string {
	if o == nil || IsNil(o.RecurringDay) {
		var ret string
		return ret
	}
	return *o.RecurringDay
}

// GetRecurringDayOk returns a tuple with the RecurringDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetRecurringDayOk() (*string, bool) {
	if o == nil || IsNil(o.RecurringDay) {
		return nil, false
	}
	return o.RecurringDay, true
}

// HasRecurringDay returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasRecurringDay() bool {
	if o != nil && !IsNil(o.RecurringDay) {
		return true
	}

	return false
}

// SetRecurringDay gets a reference to the given string and assigns it to the RecurringDay field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetRecurringDay(v string) {
	o.RecurringDay = &v
}

// GetRecurringHour returns the RecurringHour field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetRecurringHour() string {
	if o == nil || IsNil(o.RecurringHour) {
		var ret string
		return ret
	}
	return *o.RecurringHour
}

// GetRecurringHourOk returns a tuple with the RecurringHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetRecurringHourOk() (*string, bool) {
	if o == nil || IsNil(o.RecurringHour) {
		return nil, false
	}
	return o.RecurringHour, true
}

// HasRecurringHour returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasRecurringHour() bool {
	if o != nil && !IsNil(o.RecurringHour) {
		return true
	}

	return false
}

// SetRecurringHour gets a reference to the given string and assigns it to the RecurringHour field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetRecurringHour(v string) {
	o.RecurringHour = &v
}

// GetRecurringList returns the RecurringList field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetRecurringList() []CreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner {
	if o == nil || IsNil(o.RecurringList) {
		var ret []CreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner
		return ret
	}
	return o.RecurringList
}

// GetRecurringListOk returns a tuple with the RecurringList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetRecurringListOk() ([]CreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner, bool) {
	if o == nil || IsNil(o.RecurringList) {
		return nil, false
	}
	return o.RecurringList, true
}

// HasRecurringList returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasRecurringList() bool {
	if o != nil && !IsNil(o.RecurringList) {
		return true
	}

	return false
}

// SetRecurringList gets a reference to the given []CreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner and assigns it to the RecurringList field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetRecurringList(v []CreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner) {
	o.RecurringList = v
}

// GetRecurringTime returns the RecurringTime field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetRecurringTime() string {
	if o == nil || IsNil(o.RecurringTime) {
		var ret string
		return ret
	}
	return *o.RecurringTime
}

// GetRecurringTimeOk returns a tuple with the RecurringTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetRecurringTimeOk() (*string, bool) {
	if o == nil || IsNil(o.RecurringTime) {
		return nil, false
	}
	return o.RecurringTime, true
}

// HasRecurringTime returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasRecurringTime() bool {
	if o != nil && !IsNil(o.RecurringTime) {
		return true
	}

	return false
}

// SetRecurringTime gets a reference to the given string and assigns it to the RecurringTime field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetRecurringTime(v string) {
	o.RecurringTime = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetState(v string) {
	o.State = &v
}

// GetStgyName returns the StgyName field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetStgyName() string {
	if o == nil || IsNil(o.StgyName) {
		var ret string
		return ret
	}
	return *o.StgyName
}

// GetStgyNameOk returns a tuple with the StgyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetStgyNameOk() (*string, bool) {
	if o == nil || IsNil(o.StgyName) {
		return nil, false
	}
	return o.StgyName, true
}

// HasStgyName returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasStgyName() bool {
	if o != nil && !IsNil(o.StgyName) {
		return true
	}

	return false
}

// SetStgyName gets a reference to the given string and assigns it to the StgyName field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetStgyName(v string) {
	o.StgyName = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetTag(v string) {
	o.Tag = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetTotalAnnRate returns the TotalAnnRate field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetTotalAnnRate() string {
	if o == nil || IsNil(o.TotalAnnRate) {
		var ret string
		return ret
	}
	return *o.TotalAnnRate
}

// GetTotalAnnRateOk returns a tuple with the TotalAnnRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetTotalAnnRateOk() (*string, bool) {
	if o == nil || IsNil(o.TotalAnnRate) {
		return nil, false
	}
	return o.TotalAnnRate, true
}

// HasTotalAnnRate returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasTotalAnnRate() bool {
	if o != nil && !IsNil(o.TotalAnnRate) {
		return true
	}

	return false
}

// SetTotalAnnRate gets a reference to the given string and assigns it to the TotalAnnRate field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetTotalAnnRate(v string) {
	o.TotalAnnRate = &v
}

// GetTotalPnl returns the TotalPnl field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetTotalPnl() string {
	if o == nil || IsNil(o.TotalPnl) {
		var ret string
		return ret
	}
	return *o.TotalPnl
}

// GetTotalPnlOk returns a tuple with the TotalPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetTotalPnlOk() (*string, bool) {
	if o == nil || IsNil(o.TotalPnl) {
		return nil, false
	}
	return o.TotalPnl, true
}

// HasTotalPnl returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasTotalPnl() bool {
	if o != nil && !IsNil(o.TotalPnl) {
		return true
	}

	return false
}

// SetTotalPnl gets a reference to the given string and assigns it to the TotalPnl field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetTotalPnl(v string) {
	o.TotalPnl = &v
}

// GetUTime returns the UTime field value if set, zero value otherwise.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetUTime() string {
	if o == nil || IsNil(o.UTime) {
		var ret string
		return ret
	}
	return *o.UTime
}

// GetUTimeOk returns a tuple with the UTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) GetUTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UTime) {
		return nil, false
	}
	return o.UTime, true
}

// HasUTime returns a boolean if a field has been set.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) HasUTime() bool {
	if o != nil && !IsNil(o.UTime) {
		return true
	}

	return false
}

// SetUTime gets a reference to the given string and assigns it to the UTime field.
func (o *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) SetUTime(v string) {
	o.UTime = &v
}

func (o GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlgoClOrdId) {
		toSerialize["algoClOrdId"] = o.AlgoClOrdId
	}
	if !IsNil(o.AlgoId) {
		toSerialize["algoId"] = o.AlgoId
	}
	if !IsNil(o.AlgoOrdType) {
		toSerialize["algoOrdType"] = o.AlgoOrdType
	}
	if !IsNil(o.Amt) {
		toSerialize["amt"] = o.Amt
	}
	if !IsNil(o.CTime) {
		toSerialize["cTime"] = o.CTime
	}
	if !IsNil(o.Cycles) {
		toSerialize["cycles"] = o.Cycles
	}
	if !IsNil(o.InstType) {
		toSerialize["instType"] = o.InstType
	}
	if !IsNil(o.InvestmentAmt) {
		toSerialize["investmentAmt"] = o.InvestmentAmt
	}
	if !IsNil(o.InvestmentCcy) {
		toSerialize["investmentCcy"] = o.InvestmentCcy
	}
	if !IsNil(o.MktCap) {
		toSerialize["mktCap"] = o.MktCap
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !IsNil(o.PnlRatio) {
		toSerialize["pnlRatio"] = o.PnlRatio
	}
	if !IsNil(o.RecurringDay) {
		toSerialize["recurringDay"] = o.RecurringDay
	}
	if !IsNil(o.RecurringHour) {
		toSerialize["recurringHour"] = o.RecurringHour
	}
	if !IsNil(o.RecurringList) {
		toSerialize["recurringList"] = o.RecurringList
	}
	if !IsNil(o.RecurringTime) {
		toSerialize["recurringTime"] = o.RecurringTime
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.StgyName) {
		toSerialize["stgyName"] = o.StgyName
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !IsNil(o.TotalAnnRate) {
		toSerialize["totalAnnRate"] = o.TotalAnnRate
	}
	if !IsNil(o.TotalPnl) {
		toSerialize["totalPnl"] = o.TotalPnl
	}
	if !IsNil(o.UTime) {
		toSerialize["uTime"] = o.UTime
	}
	return toSerialize, nil
}

type NullableGetTradingBotRecurringOrdersAlgoPendingV5RespDataInner struct {
	value *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner
	isSet bool
}

func (v NullableGetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) Get() *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner {
	return v.value
}

func (v *NullableGetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) Set(val *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTradingBotRecurringOrdersAlgoPendingV5RespDataInner(val *GetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) *NullableGetTradingBotRecurringOrdersAlgoPendingV5RespDataInner {
	return &NullableGetTradingBotRecurringOrdersAlgoPendingV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTradingBotRecurringOrdersAlgoPendingV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


