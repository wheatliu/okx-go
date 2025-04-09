/*
OKX v5 API

OpenAPI specification for Okx exchange - Rest API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rest

import (
	"encoding/json"
)

// checks if the GetTradingBotGridPositionsV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTradingBotGridPositionsV5RespDataInner{}

// GetTradingBotGridPositionsV5RespDataInner struct for GetTradingBotGridPositionsV5RespDataInner
type GetTradingBotGridPositionsV5RespDataInner struct {
	// Auto decrease line, signal area  Divided into 5 levels, from 1 to 5, the smaller the number, the weaker the adl intensity.
	Adl *string `json:"adl,omitempty"`
	// Client-supplied Algo ID
	AlgoClOrdId *string `json:"algoClOrdId,omitempty"`
	// Algo ID
	AlgoId *string `json:"algoId,omitempty"`
	// Average open price
	AvgPx *string `json:"avgPx,omitempty"`
	// Algo order created time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	CTime *string `json:"cTime,omitempty"`
	// Margin currency
	Ccy *string `json:"ccy,omitempty"`
	// Initial margin requirement
	Imr *string `json:"imr,omitempty"`
	// Instrument ID, e.g. `BTC-USDT-SWAP`
	InstId *string `json:"instId,omitempty"`
	// Instrument type
	InstType *string `json:"instType,omitempty"`
	// Latest traded price
	Last *string `json:"last,omitempty"`
	// Leverage
	Lever *string `json:"lever,omitempty"`
	// Estimated liquidation price
	LiqPx *string `json:"liqPx,omitempty"`
	// Mark price
	MarkPx *string `json:"markPx,omitempty"`
	// Margin mode  `cross`  `isolated`
	MgnMode *string `json:"mgnMode,omitempty"`
	// Margin ratio
	MgnRatio *string `json:"mgnRatio,omitempty"`
	// Maintenance margin requirement
	Mmr *string `json:"mmr,omitempty"`
	// Notional value of positions in `USD`
	NotionalUsd *string `json:"notionalUsd,omitempty"`
	// Quantity of positions
	Pos *string `json:"pos,omitempty"`
	// Position side  `net`
	PosSide *string `json:"posSide,omitempty"`
	// Algo order updated time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	UTime *string `json:"uTime,omitempty"`
	// Unrealized profit and loss
	Upl *string `json:"upl,omitempty"`
	// Unrealized profit and loss ratio
	UplRatio *string `json:"uplRatio,omitempty"`
}

// NewGetTradingBotGridPositionsV5RespDataInner instantiates a new GetTradingBotGridPositionsV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTradingBotGridPositionsV5RespDataInner() *GetTradingBotGridPositionsV5RespDataInner {
	this := GetTradingBotGridPositionsV5RespDataInner{}
	var adl string = ""
	this.Adl = &adl
	var algoClOrdId string = ""
	this.AlgoClOrdId = &algoClOrdId
	var algoId string = ""
	this.AlgoId = &algoId
	var avgPx string = ""
	this.AvgPx = &avgPx
	var cTime string = ""
	this.CTime = &cTime
	var ccy string = ""
	this.Ccy = &ccy
	var imr string = ""
	this.Imr = &imr
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var last string = ""
	this.Last = &last
	var lever string = ""
	this.Lever = &lever
	var liqPx string = ""
	this.LiqPx = &liqPx
	var markPx string = ""
	this.MarkPx = &markPx
	var mgnMode string = ""
	this.MgnMode = &mgnMode
	var mgnRatio string = ""
	this.MgnRatio = &mgnRatio
	var mmr string = ""
	this.Mmr = &mmr
	var notionalUsd string = ""
	this.NotionalUsd = &notionalUsd
	var pos string = ""
	this.Pos = &pos
	var posSide string = ""
	this.PosSide = &posSide
	var uTime string = ""
	this.UTime = &uTime
	var upl string = ""
	this.Upl = &upl
	var uplRatio string = ""
	this.UplRatio = &uplRatio
	return &this
}

// NewGetTradingBotGridPositionsV5RespDataInnerWithDefaults instantiates a new GetTradingBotGridPositionsV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTradingBotGridPositionsV5RespDataInnerWithDefaults() *GetTradingBotGridPositionsV5RespDataInner {
	this := GetTradingBotGridPositionsV5RespDataInner{}
	var adl string = ""
	this.Adl = &adl
	var algoClOrdId string = ""
	this.AlgoClOrdId = &algoClOrdId
	var algoId string = ""
	this.AlgoId = &algoId
	var avgPx string = ""
	this.AvgPx = &avgPx
	var cTime string = ""
	this.CTime = &cTime
	var ccy string = ""
	this.Ccy = &ccy
	var imr string = ""
	this.Imr = &imr
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var last string = ""
	this.Last = &last
	var lever string = ""
	this.Lever = &lever
	var liqPx string = ""
	this.LiqPx = &liqPx
	var markPx string = ""
	this.MarkPx = &markPx
	var mgnMode string = ""
	this.MgnMode = &mgnMode
	var mgnRatio string = ""
	this.MgnRatio = &mgnRatio
	var mmr string = ""
	this.Mmr = &mmr
	var notionalUsd string = ""
	this.NotionalUsd = &notionalUsd
	var pos string = ""
	this.Pos = &pos
	var posSide string = ""
	this.PosSide = &posSide
	var uTime string = ""
	this.UTime = &uTime
	var upl string = ""
	this.Upl = &upl
	var uplRatio string = ""
	this.UplRatio = &uplRatio
	return &this
}

// GetAdl returns the Adl field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetAdl() string {
	if o == nil || IsNil(o.Adl) {
		var ret string
		return ret
	}
	return *o.Adl
}

// GetAdlOk returns a tuple with the Adl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetAdlOk() (*string, bool) {
	if o == nil || IsNil(o.Adl) {
		return nil, false
	}
	return o.Adl, true
}

// HasAdl returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasAdl() bool {
	if o != nil && !IsNil(o.Adl) {
		return true
	}

	return false
}

// SetAdl gets a reference to the given string and assigns it to the Adl field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetAdl(v string) {
	o.Adl = &v
}

// GetAlgoClOrdId returns the AlgoClOrdId field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetAlgoClOrdId() string {
	if o == nil || IsNil(o.AlgoClOrdId) {
		var ret string
		return ret
	}
	return *o.AlgoClOrdId
}

// GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetAlgoClOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoClOrdId) {
		return nil, false
	}
	return o.AlgoClOrdId, true
}

// HasAlgoClOrdId returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasAlgoClOrdId() bool {
	if o != nil && !IsNil(o.AlgoClOrdId) {
		return true
	}

	return false
}

// SetAlgoClOrdId gets a reference to the given string and assigns it to the AlgoClOrdId field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetAlgoClOrdId(v string) {
	o.AlgoClOrdId = &v
}

// GetAlgoId returns the AlgoId field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetAlgoId() string {
	if o == nil || IsNil(o.AlgoId) {
		var ret string
		return ret
	}
	return *o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetAlgoIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoId) {
		return nil, false
	}
	return o.AlgoId, true
}

// HasAlgoId returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasAlgoId() bool {
	if o != nil && !IsNil(o.AlgoId) {
		return true
	}

	return false
}

// SetAlgoId gets a reference to the given string and assigns it to the AlgoId field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetAlgoId(v string) {
	o.AlgoId = &v
}

// GetAvgPx returns the AvgPx field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetAvgPx() string {
	if o == nil || IsNil(o.AvgPx) {
		var ret string
		return ret
	}
	return *o.AvgPx
}

// GetAvgPxOk returns a tuple with the AvgPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetAvgPxOk() (*string, bool) {
	if o == nil || IsNil(o.AvgPx) {
		return nil, false
	}
	return o.AvgPx, true
}

// HasAvgPx returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasAvgPx() bool {
	if o != nil && !IsNil(o.AvgPx) {
		return true
	}

	return false
}

// SetAvgPx gets a reference to the given string and assigns it to the AvgPx field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetAvgPx(v string) {
	o.AvgPx = &v
}

// GetCTime returns the CTime field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetCTime() string {
	if o == nil || IsNil(o.CTime) {
		var ret string
		return ret
	}
	return *o.CTime
}

// GetCTimeOk returns a tuple with the CTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetCTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CTime) {
		return nil, false
	}
	return o.CTime, true
}

// HasCTime returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasCTime() bool {
	if o != nil && !IsNil(o.CTime) {
		return true
	}

	return false
}

// SetCTime gets a reference to the given string and assigns it to the CTime field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetCTime(v string) {
	o.CTime = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetImr returns the Imr field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetImr() string {
	if o == nil || IsNil(o.Imr) {
		var ret string
		return ret
	}
	return *o.Imr
}

// GetImrOk returns a tuple with the Imr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetImrOk() (*string, bool) {
	if o == nil || IsNil(o.Imr) {
		return nil, false
	}
	return o.Imr, true
}

// HasImr returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasImr() bool {
	if o != nil && !IsNil(o.Imr) {
		return true
	}

	return false
}

// SetImr gets a reference to the given string and assigns it to the Imr field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetImr(v string) {
	o.Imr = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetInstId(v string) {
	o.InstId = &v
}

// GetInstType returns the InstType field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetInstType() string {
	if o == nil || IsNil(o.InstType) {
		var ret string
		return ret
	}
	return *o.InstType
}

// GetInstTypeOk returns a tuple with the InstType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetInstTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstType) {
		return nil, false
	}
	return o.InstType, true
}

// HasInstType returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasInstType() bool {
	if o != nil && !IsNil(o.InstType) {
		return true
	}

	return false
}

// SetInstType gets a reference to the given string and assigns it to the InstType field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetInstType(v string) {
	o.InstType = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetLast() string {
	if o == nil || IsNil(o.Last) {
		var ret string
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetLastOk() (*string, bool) {
	if o == nil || IsNil(o.Last) {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasLast() bool {
	if o != nil && !IsNil(o.Last) {
		return true
	}

	return false
}

// SetLast gets a reference to the given string and assigns it to the Last field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetLast(v string) {
	o.Last = &v
}

// GetLever returns the Lever field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetLever() string {
	if o == nil || IsNil(o.Lever) {
		var ret string
		return ret
	}
	return *o.Lever
}

// GetLeverOk returns a tuple with the Lever field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetLeverOk() (*string, bool) {
	if o == nil || IsNil(o.Lever) {
		return nil, false
	}
	return o.Lever, true
}

// HasLever returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasLever() bool {
	if o != nil && !IsNil(o.Lever) {
		return true
	}

	return false
}

// SetLever gets a reference to the given string and assigns it to the Lever field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetLever(v string) {
	o.Lever = &v
}

// GetLiqPx returns the LiqPx field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetLiqPx() string {
	if o == nil || IsNil(o.LiqPx) {
		var ret string
		return ret
	}
	return *o.LiqPx
}

// GetLiqPxOk returns a tuple with the LiqPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetLiqPxOk() (*string, bool) {
	if o == nil || IsNil(o.LiqPx) {
		return nil, false
	}
	return o.LiqPx, true
}

// HasLiqPx returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasLiqPx() bool {
	if o != nil && !IsNil(o.LiqPx) {
		return true
	}

	return false
}

// SetLiqPx gets a reference to the given string and assigns it to the LiqPx field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetLiqPx(v string) {
	o.LiqPx = &v
}

// GetMarkPx returns the MarkPx field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetMarkPx() string {
	if o == nil || IsNil(o.MarkPx) {
		var ret string
		return ret
	}
	return *o.MarkPx
}

// GetMarkPxOk returns a tuple with the MarkPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetMarkPxOk() (*string, bool) {
	if o == nil || IsNil(o.MarkPx) {
		return nil, false
	}
	return o.MarkPx, true
}

// HasMarkPx returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasMarkPx() bool {
	if o != nil && !IsNil(o.MarkPx) {
		return true
	}

	return false
}

// SetMarkPx gets a reference to the given string and assigns it to the MarkPx field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetMarkPx(v string) {
	o.MarkPx = &v
}

// GetMgnMode returns the MgnMode field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetMgnMode() string {
	if o == nil || IsNil(o.MgnMode) {
		var ret string
		return ret
	}
	return *o.MgnMode
}

// GetMgnModeOk returns a tuple with the MgnMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetMgnModeOk() (*string, bool) {
	if o == nil || IsNil(o.MgnMode) {
		return nil, false
	}
	return o.MgnMode, true
}

// HasMgnMode returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasMgnMode() bool {
	if o != nil && !IsNil(o.MgnMode) {
		return true
	}

	return false
}

// SetMgnMode gets a reference to the given string and assigns it to the MgnMode field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetMgnMode(v string) {
	o.MgnMode = &v
}

// GetMgnRatio returns the MgnRatio field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetMgnRatio() string {
	if o == nil || IsNil(o.MgnRatio) {
		var ret string
		return ret
	}
	return *o.MgnRatio
}

// GetMgnRatioOk returns a tuple with the MgnRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetMgnRatioOk() (*string, bool) {
	if o == nil || IsNil(o.MgnRatio) {
		return nil, false
	}
	return o.MgnRatio, true
}

// HasMgnRatio returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasMgnRatio() bool {
	if o != nil && !IsNil(o.MgnRatio) {
		return true
	}

	return false
}

// SetMgnRatio gets a reference to the given string and assigns it to the MgnRatio field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetMgnRatio(v string) {
	o.MgnRatio = &v
}

// GetMmr returns the Mmr field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetMmr() string {
	if o == nil || IsNil(o.Mmr) {
		var ret string
		return ret
	}
	return *o.Mmr
}

// GetMmrOk returns a tuple with the Mmr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetMmrOk() (*string, bool) {
	if o == nil || IsNil(o.Mmr) {
		return nil, false
	}
	return o.Mmr, true
}

// HasMmr returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasMmr() bool {
	if o != nil && !IsNil(o.Mmr) {
		return true
	}

	return false
}

// SetMmr gets a reference to the given string and assigns it to the Mmr field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetMmr(v string) {
	o.Mmr = &v
}

// GetNotionalUsd returns the NotionalUsd field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetNotionalUsd() string {
	if o == nil || IsNil(o.NotionalUsd) {
		var ret string
		return ret
	}
	return *o.NotionalUsd
}

// GetNotionalUsdOk returns a tuple with the NotionalUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetNotionalUsdOk() (*string, bool) {
	if o == nil || IsNil(o.NotionalUsd) {
		return nil, false
	}
	return o.NotionalUsd, true
}

// HasNotionalUsd returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasNotionalUsd() bool {
	if o != nil && !IsNil(o.NotionalUsd) {
		return true
	}

	return false
}

// SetNotionalUsd gets a reference to the given string and assigns it to the NotionalUsd field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetNotionalUsd(v string) {
	o.NotionalUsd = &v
}

// GetPos returns the Pos field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetPos() string {
	if o == nil || IsNil(o.Pos) {
		var ret string
		return ret
	}
	return *o.Pos
}

// GetPosOk returns a tuple with the Pos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetPosOk() (*string, bool) {
	if o == nil || IsNil(o.Pos) {
		return nil, false
	}
	return o.Pos, true
}

// HasPos returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasPos() bool {
	if o != nil && !IsNil(o.Pos) {
		return true
	}

	return false
}

// SetPos gets a reference to the given string and assigns it to the Pos field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetPos(v string) {
	o.Pos = &v
}

// GetPosSide returns the PosSide field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetPosSide() string {
	if o == nil || IsNil(o.PosSide) {
		var ret string
		return ret
	}
	return *o.PosSide
}

// GetPosSideOk returns a tuple with the PosSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetPosSideOk() (*string, bool) {
	if o == nil || IsNil(o.PosSide) {
		return nil, false
	}
	return o.PosSide, true
}

// HasPosSide returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasPosSide() bool {
	if o != nil && !IsNil(o.PosSide) {
		return true
	}

	return false
}

// SetPosSide gets a reference to the given string and assigns it to the PosSide field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetPosSide(v string) {
	o.PosSide = &v
}

// GetUTime returns the UTime field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetUTime() string {
	if o == nil || IsNil(o.UTime) {
		var ret string
		return ret
	}
	return *o.UTime
}

// GetUTimeOk returns a tuple with the UTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetUTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UTime) {
		return nil, false
	}
	return o.UTime, true
}

// HasUTime returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasUTime() bool {
	if o != nil && !IsNil(o.UTime) {
		return true
	}

	return false
}

// SetUTime gets a reference to the given string and assigns it to the UTime field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetUTime(v string) {
	o.UTime = &v
}

// GetUpl returns the Upl field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetUpl() string {
	if o == nil || IsNil(o.Upl) {
		var ret string
		return ret
	}
	return *o.Upl
}

// GetUplOk returns a tuple with the Upl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetUplOk() (*string, bool) {
	if o == nil || IsNil(o.Upl) {
		return nil, false
	}
	return o.Upl, true
}

// HasUpl returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasUpl() bool {
	if o != nil && !IsNil(o.Upl) {
		return true
	}

	return false
}

// SetUpl gets a reference to the given string and assigns it to the Upl field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetUpl(v string) {
	o.Upl = &v
}

// GetUplRatio returns the UplRatio field value if set, zero value otherwise.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetUplRatio() string {
	if o == nil || IsNil(o.UplRatio) {
		var ret string
		return ret
	}
	return *o.UplRatio
}

// GetUplRatioOk returns a tuple with the UplRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) GetUplRatioOk() (*string, bool) {
	if o == nil || IsNil(o.UplRatio) {
		return nil, false
	}
	return o.UplRatio, true
}

// HasUplRatio returns a boolean if a field has been set.
func (o *GetTradingBotGridPositionsV5RespDataInner) HasUplRatio() bool {
	if o != nil && !IsNil(o.UplRatio) {
		return true
	}

	return false
}

// SetUplRatio gets a reference to the given string and assigns it to the UplRatio field.
func (o *GetTradingBotGridPositionsV5RespDataInner) SetUplRatio(v string) {
	o.UplRatio = &v
}

func (o GetTradingBotGridPositionsV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTradingBotGridPositionsV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Adl) {
		toSerialize["adl"] = o.Adl
	}
	if !IsNil(o.AlgoClOrdId) {
		toSerialize["algoClOrdId"] = o.AlgoClOrdId
	}
	if !IsNil(o.AlgoId) {
		toSerialize["algoId"] = o.AlgoId
	}
	if !IsNil(o.AvgPx) {
		toSerialize["avgPx"] = o.AvgPx
	}
	if !IsNil(o.CTime) {
		toSerialize["cTime"] = o.CTime
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.Imr) {
		toSerialize["imr"] = o.Imr
	}
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	if !IsNil(o.InstType) {
		toSerialize["instType"] = o.InstType
	}
	if !IsNil(o.Last) {
		toSerialize["last"] = o.Last
	}
	if !IsNil(o.Lever) {
		toSerialize["lever"] = o.Lever
	}
	if !IsNil(o.LiqPx) {
		toSerialize["liqPx"] = o.LiqPx
	}
	if !IsNil(o.MarkPx) {
		toSerialize["markPx"] = o.MarkPx
	}
	if !IsNil(o.MgnMode) {
		toSerialize["mgnMode"] = o.MgnMode
	}
	if !IsNil(o.MgnRatio) {
		toSerialize["mgnRatio"] = o.MgnRatio
	}
	if !IsNil(o.Mmr) {
		toSerialize["mmr"] = o.Mmr
	}
	if !IsNil(o.NotionalUsd) {
		toSerialize["notionalUsd"] = o.NotionalUsd
	}
	if !IsNil(o.Pos) {
		toSerialize["pos"] = o.Pos
	}
	if !IsNil(o.PosSide) {
		toSerialize["posSide"] = o.PosSide
	}
	if !IsNil(o.UTime) {
		toSerialize["uTime"] = o.UTime
	}
	if !IsNil(o.Upl) {
		toSerialize["upl"] = o.Upl
	}
	if !IsNil(o.UplRatio) {
		toSerialize["uplRatio"] = o.UplRatio
	}
	return toSerialize, nil
}

type NullableGetTradingBotGridPositionsV5RespDataInner struct {
	value *GetTradingBotGridPositionsV5RespDataInner
	isSet bool
}

func (v NullableGetTradingBotGridPositionsV5RespDataInner) Get() *GetTradingBotGridPositionsV5RespDataInner {
	return v.value
}

func (v *NullableGetTradingBotGridPositionsV5RespDataInner) Set(val *GetTradingBotGridPositionsV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTradingBotGridPositionsV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTradingBotGridPositionsV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTradingBotGridPositionsV5RespDataInner(val *GetTradingBotGridPositionsV5RespDataInner) *NullableGetTradingBotGridPositionsV5RespDataInner {
	return &NullableGetTradingBotGridPositionsV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetTradingBotGridPositionsV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTradingBotGridPositionsV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


