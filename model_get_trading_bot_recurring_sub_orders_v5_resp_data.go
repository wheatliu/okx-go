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

// checks if the GetTradingBotRecurringSubOrdersV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTradingBotRecurringSubOrdersV5RespData{}

// GetTradingBotRecurringSubOrdersV5RespData struct for GetTradingBotRecurringSubOrdersV5RespData
type GetTradingBotRecurringSubOrdersV5RespData struct {
	// Sub order accumulated fill quantity
	AccFillSz *string `json:"accFillSz,omitempty"`
	// Client-supplied Algo ID
	AlgoClOrdId *string `json:"algoClOrdId,omitempty"`
	// Algo ID
	AlgoId *string `json:"algoId,omitempty"`
	// Algo order type  `recurring`: recurring buy
	AlgoOrdType *string `json:"algoOrdType,omitempty"`
	// Sub order average filled price
	AvgPx *string `json:"avgPx,omitempty"`
	// Sub order created time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	CTime *string `json:"cTime,omitempty"`
	// Sub order fee
	Fee *string `json:"fee,omitempty"`
	// Sub order fee currency
	FeeCcy *string `json:"feeCcy,omitempty"`
	// Instrument ID
	InstId *string `json:"instId,omitempty"`
	// Instrument type
	InstType *string `json:"instType,omitempty"`
	// Sub order ID
	OrdId *string `json:"ordId,omitempty"`
	// Sub order type  `market`: Market order
	OrdType *string `json:"ordType,omitempty"`
	// Sub order limit price  If it's a market order, \"-1\" will be return
	Px *string `json:"px,omitempty"`
	// Sub order side  `buy` `sell`
	Side *string `json:"side,omitempty"`
	// Sub order state  `canceled`  `live`  `partially_filled`  `filled`  `cancelling`
	State *string `json:"state,omitempty"`
	// Sub order quantity to buy or sell
	Sz *string `json:"sz,omitempty"`
	// Order tag
	Tag *string `json:"tag,omitempty"`
	// Sub order trade mode  Margin mode : `cross`  Non-Margin mode : `cash`
	TdMode *string `json:"tdMode,omitempty"`
	// Sub order updated time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	UTime *string `json:"uTime,omitempty"`
}

// NewGetTradingBotRecurringSubOrdersV5RespData instantiates a new GetTradingBotRecurringSubOrdersV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTradingBotRecurringSubOrdersV5RespData() *GetTradingBotRecurringSubOrdersV5RespData {
	this := GetTradingBotRecurringSubOrdersV5RespData{}
	var accFillSz string = ""
	this.AccFillSz = &accFillSz
	var algoClOrdId string = ""
	this.AlgoClOrdId = &algoClOrdId
	var algoId string = ""
	this.AlgoId = &algoId
	var algoOrdType string = ""
	this.AlgoOrdType = &algoOrdType
	var avgPx string = ""
	this.AvgPx = &avgPx
	var cTime string = ""
	this.CTime = &cTime
	var fee string = ""
	this.Fee = &fee
	var feeCcy string = ""
	this.FeeCcy = &feeCcy
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var ordId string = ""
	this.OrdId = &ordId
	var ordType string = ""
	this.OrdType = &ordType
	var px string = ""
	this.Px = &px
	var side string = ""
	this.Side = &side
	var state string = ""
	this.State = &state
	var sz string = ""
	this.Sz = &sz
	var tag string = ""
	this.Tag = &tag
	var tdMode string = ""
	this.TdMode = &tdMode
	var uTime string = ""
	this.UTime = &uTime
	return &this
}

// NewGetTradingBotRecurringSubOrdersV5RespDataWithDefaults instantiates a new GetTradingBotRecurringSubOrdersV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTradingBotRecurringSubOrdersV5RespDataWithDefaults() *GetTradingBotRecurringSubOrdersV5RespData {
	this := GetTradingBotRecurringSubOrdersV5RespData{}
	var accFillSz string = ""
	this.AccFillSz = &accFillSz
	var algoClOrdId string = ""
	this.AlgoClOrdId = &algoClOrdId
	var algoId string = ""
	this.AlgoId = &algoId
	var algoOrdType string = ""
	this.AlgoOrdType = &algoOrdType
	var avgPx string = ""
	this.AvgPx = &avgPx
	var cTime string = ""
	this.CTime = &cTime
	var fee string = ""
	this.Fee = &fee
	var feeCcy string = ""
	this.FeeCcy = &feeCcy
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var ordId string = ""
	this.OrdId = &ordId
	var ordType string = ""
	this.OrdType = &ordType
	var px string = ""
	this.Px = &px
	var side string = ""
	this.Side = &side
	var state string = ""
	this.State = &state
	var sz string = ""
	this.Sz = &sz
	var tag string = ""
	this.Tag = &tag
	var tdMode string = ""
	this.TdMode = &tdMode
	var uTime string = ""
	this.UTime = &uTime
	return &this
}

// GetAccFillSz returns the AccFillSz field value if set, zero value otherwise.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetAccFillSz() string {
	if o == nil || IsNil(o.AccFillSz) {
		var ret string
		return ret
	}
	return *o.AccFillSz
}

// GetAccFillSzOk returns a tuple with the AccFillSz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetAccFillSzOk() (*string, bool) {
	if o == nil || IsNil(o.AccFillSz) {
		return nil, false
	}
	return o.AccFillSz, true
}

// HasAccFillSz returns a boolean if a field has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) HasAccFillSz() bool {
	if o != nil && !IsNil(o.AccFillSz) {
		return true
	}

	return false
}

// SetAccFillSz gets a reference to the given string and assigns it to the AccFillSz field.
func (o *GetTradingBotRecurringSubOrdersV5RespData) SetAccFillSz(v string) {
	o.AccFillSz = &v
}

// GetAlgoClOrdId returns the AlgoClOrdId field value if set, zero value otherwise.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetAlgoClOrdId() string {
	if o == nil || IsNil(o.AlgoClOrdId) {
		var ret string
		return ret
	}
	return *o.AlgoClOrdId
}

// GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetAlgoClOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoClOrdId) {
		return nil, false
	}
	return o.AlgoClOrdId, true
}

// HasAlgoClOrdId returns a boolean if a field has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) HasAlgoClOrdId() bool {
	if o != nil && !IsNil(o.AlgoClOrdId) {
		return true
	}

	return false
}

// SetAlgoClOrdId gets a reference to the given string and assigns it to the AlgoClOrdId field.
func (o *GetTradingBotRecurringSubOrdersV5RespData) SetAlgoClOrdId(v string) {
	o.AlgoClOrdId = &v
}

// GetAlgoId returns the AlgoId field value if set, zero value otherwise.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetAlgoId() string {
	if o == nil || IsNil(o.AlgoId) {
		var ret string
		return ret
	}
	return *o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetAlgoIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoId) {
		return nil, false
	}
	return o.AlgoId, true
}

// HasAlgoId returns a boolean if a field has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) HasAlgoId() bool {
	if o != nil && !IsNil(o.AlgoId) {
		return true
	}

	return false
}

// SetAlgoId gets a reference to the given string and assigns it to the AlgoId field.
func (o *GetTradingBotRecurringSubOrdersV5RespData) SetAlgoId(v string) {
	o.AlgoId = &v
}

// GetAlgoOrdType returns the AlgoOrdType field value if set, zero value otherwise.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetAlgoOrdType() string {
	if o == nil || IsNil(o.AlgoOrdType) {
		var ret string
		return ret
	}
	return *o.AlgoOrdType
}

// GetAlgoOrdTypeOk returns a tuple with the AlgoOrdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetAlgoOrdTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoOrdType) {
		return nil, false
	}
	return o.AlgoOrdType, true
}

// HasAlgoOrdType returns a boolean if a field has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) HasAlgoOrdType() bool {
	if o != nil && !IsNil(o.AlgoOrdType) {
		return true
	}

	return false
}

// SetAlgoOrdType gets a reference to the given string and assigns it to the AlgoOrdType field.
func (o *GetTradingBotRecurringSubOrdersV5RespData) SetAlgoOrdType(v string) {
	o.AlgoOrdType = &v
}

// GetAvgPx returns the AvgPx field value if set, zero value otherwise.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetAvgPx() string {
	if o == nil || IsNil(o.AvgPx) {
		var ret string
		return ret
	}
	return *o.AvgPx
}

// GetAvgPxOk returns a tuple with the AvgPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetAvgPxOk() (*string, bool) {
	if o == nil || IsNil(o.AvgPx) {
		return nil, false
	}
	return o.AvgPx, true
}

// HasAvgPx returns a boolean if a field has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) HasAvgPx() bool {
	if o != nil && !IsNil(o.AvgPx) {
		return true
	}

	return false
}

// SetAvgPx gets a reference to the given string and assigns it to the AvgPx field.
func (o *GetTradingBotRecurringSubOrdersV5RespData) SetAvgPx(v string) {
	o.AvgPx = &v
}

// GetCTime returns the CTime field value if set, zero value otherwise.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetCTime() string {
	if o == nil || IsNil(o.CTime) {
		var ret string
		return ret
	}
	return *o.CTime
}

// GetCTimeOk returns a tuple with the CTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetCTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CTime) {
		return nil, false
	}
	return o.CTime, true
}

// HasCTime returns a boolean if a field has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) HasCTime() bool {
	if o != nil && !IsNil(o.CTime) {
		return true
	}

	return false
}

// SetCTime gets a reference to the given string and assigns it to the CTime field.
func (o *GetTradingBotRecurringSubOrdersV5RespData) SetCTime(v string) {
	o.CTime = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetFee() string {
	if o == nil || IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetFeeOk() (*string, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *GetTradingBotRecurringSubOrdersV5RespData) SetFee(v string) {
	o.Fee = &v
}

// GetFeeCcy returns the FeeCcy field value if set, zero value otherwise.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetFeeCcy() string {
	if o == nil || IsNil(o.FeeCcy) {
		var ret string
		return ret
	}
	return *o.FeeCcy
}

// GetFeeCcyOk returns a tuple with the FeeCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetFeeCcyOk() (*string, bool) {
	if o == nil || IsNil(o.FeeCcy) {
		return nil, false
	}
	return o.FeeCcy, true
}

// HasFeeCcy returns a boolean if a field has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) HasFeeCcy() bool {
	if o != nil && !IsNil(o.FeeCcy) {
		return true
	}

	return false
}

// SetFeeCcy gets a reference to the given string and assigns it to the FeeCcy field.
func (o *GetTradingBotRecurringSubOrdersV5RespData) SetFeeCcy(v string) {
	o.FeeCcy = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetTradingBotRecurringSubOrdersV5RespData) SetInstId(v string) {
	o.InstId = &v
}

// GetInstType returns the InstType field value if set, zero value otherwise.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetInstType() string {
	if o == nil || IsNil(o.InstType) {
		var ret string
		return ret
	}
	return *o.InstType
}

// GetInstTypeOk returns a tuple with the InstType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetInstTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstType) {
		return nil, false
	}
	return o.InstType, true
}

// HasInstType returns a boolean if a field has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) HasInstType() bool {
	if o != nil && !IsNil(o.InstType) {
		return true
	}

	return false
}

// SetInstType gets a reference to the given string and assigns it to the InstType field.
func (o *GetTradingBotRecurringSubOrdersV5RespData) SetInstType(v string) {
	o.InstType = &v
}

// GetOrdId returns the OrdId field value if set, zero value otherwise.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetOrdId() string {
	if o == nil || IsNil(o.OrdId) {
		var ret string
		return ret
	}
	return *o.OrdId
}

// GetOrdIdOk returns a tuple with the OrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrdId) {
		return nil, false
	}
	return o.OrdId, true
}

// HasOrdId returns a boolean if a field has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) HasOrdId() bool {
	if o != nil && !IsNil(o.OrdId) {
		return true
	}

	return false
}

// SetOrdId gets a reference to the given string and assigns it to the OrdId field.
func (o *GetTradingBotRecurringSubOrdersV5RespData) SetOrdId(v string) {
	o.OrdId = &v
}

// GetOrdType returns the OrdType field value if set, zero value otherwise.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetOrdType() string {
	if o == nil || IsNil(o.OrdType) {
		var ret string
		return ret
	}
	return *o.OrdType
}

// GetOrdTypeOk returns a tuple with the OrdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetOrdTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OrdType) {
		return nil, false
	}
	return o.OrdType, true
}

// HasOrdType returns a boolean if a field has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) HasOrdType() bool {
	if o != nil && !IsNil(o.OrdType) {
		return true
	}

	return false
}

// SetOrdType gets a reference to the given string and assigns it to the OrdType field.
func (o *GetTradingBotRecurringSubOrdersV5RespData) SetOrdType(v string) {
	o.OrdType = &v
}

// GetPx returns the Px field value if set, zero value otherwise.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetPx() string {
	if o == nil || IsNil(o.Px) {
		var ret string
		return ret
	}
	return *o.Px
}

// GetPxOk returns a tuple with the Px field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetPxOk() (*string, bool) {
	if o == nil || IsNil(o.Px) {
		return nil, false
	}
	return o.Px, true
}

// HasPx returns a boolean if a field has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) HasPx() bool {
	if o != nil && !IsNil(o.Px) {
		return true
	}

	return false
}

// SetPx gets a reference to the given string and assigns it to the Px field.
func (o *GetTradingBotRecurringSubOrdersV5RespData) SetPx(v string) {
	o.Px = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *GetTradingBotRecurringSubOrdersV5RespData) SetSide(v string) {
	o.Side = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GetTradingBotRecurringSubOrdersV5RespData) SetState(v string) {
	o.State = &v
}

// GetSz returns the Sz field value if set, zero value otherwise.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetSz() string {
	if o == nil || IsNil(o.Sz) {
		var ret string
		return ret
	}
	return *o.Sz
}

// GetSzOk returns a tuple with the Sz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetSzOk() (*string, bool) {
	if o == nil || IsNil(o.Sz) {
		return nil, false
	}
	return o.Sz, true
}

// HasSz returns a boolean if a field has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) HasSz() bool {
	if o != nil && !IsNil(o.Sz) {
		return true
	}

	return false
}

// SetSz gets a reference to the given string and assigns it to the Sz field.
func (o *GetTradingBotRecurringSubOrdersV5RespData) SetSz(v string) {
	o.Sz = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *GetTradingBotRecurringSubOrdersV5RespData) SetTag(v string) {
	o.Tag = &v
}

// GetTdMode returns the TdMode field value if set, zero value otherwise.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetTdMode() string {
	if o == nil || IsNil(o.TdMode) {
		var ret string
		return ret
	}
	return *o.TdMode
}

// GetTdModeOk returns a tuple with the TdMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetTdModeOk() (*string, bool) {
	if o == nil || IsNil(o.TdMode) {
		return nil, false
	}
	return o.TdMode, true
}

// HasTdMode returns a boolean if a field has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) HasTdMode() bool {
	if o != nil && !IsNil(o.TdMode) {
		return true
	}

	return false
}

// SetTdMode gets a reference to the given string and assigns it to the TdMode field.
func (o *GetTradingBotRecurringSubOrdersV5RespData) SetTdMode(v string) {
	o.TdMode = &v
}

// GetUTime returns the UTime field value if set, zero value otherwise.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetUTime() string {
	if o == nil || IsNil(o.UTime) {
		var ret string
		return ret
	}
	return *o.UTime
}

// GetUTimeOk returns a tuple with the UTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) GetUTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UTime) {
		return nil, false
	}
	return o.UTime, true
}

// HasUTime returns a boolean if a field has been set.
func (o *GetTradingBotRecurringSubOrdersV5RespData) HasUTime() bool {
	if o != nil && !IsNil(o.UTime) {
		return true
	}

	return false
}

// SetUTime gets a reference to the given string and assigns it to the UTime field.
func (o *GetTradingBotRecurringSubOrdersV5RespData) SetUTime(v string) {
	o.UTime = &v
}

func (o GetTradingBotRecurringSubOrdersV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTradingBotRecurringSubOrdersV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccFillSz) {
		toSerialize["accFillSz"] = o.AccFillSz
	}
	if !IsNil(o.AlgoClOrdId) {
		toSerialize["algoClOrdId"] = o.AlgoClOrdId
	}
	if !IsNil(o.AlgoId) {
		toSerialize["algoId"] = o.AlgoId
	}
	if !IsNil(o.AlgoOrdType) {
		toSerialize["algoOrdType"] = o.AlgoOrdType
	}
	if !IsNil(o.AvgPx) {
		toSerialize["avgPx"] = o.AvgPx
	}
	if !IsNil(o.CTime) {
		toSerialize["cTime"] = o.CTime
	}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.FeeCcy) {
		toSerialize["feeCcy"] = o.FeeCcy
	}
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	if !IsNil(o.InstType) {
		toSerialize["instType"] = o.InstType
	}
	if !IsNil(o.OrdId) {
		toSerialize["ordId"] = o.OrdId
	}
	if !IsNil(o.OrdType) {
		toSerialize["ordType"] = o.OrdType
	}
	if !IsNil(o.Px) {
		toSerialize["px"] = o.Px
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Sz) {
		toSerialize["sz"] = o.Sz
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.TdMode) {
		toSerialize["tdMode"] = o.TdMode
	}
	if !IsNil(o.UTime) {
		toSerialize["uTime"] = o.UTime
	}
	return toSerialize, nil
}

type NullableGetTradingBotRecurringSubOrdersV5RespData struct {
	value *GetTradingBotRecurringSubOrdersV5RespData
	isSet bool
}

func (v NullableGetTradingBotRecurringSubOrdersV5RespData) Get() *GetTradingBotRecurringSubOrdersV5RespData {
	return v.value
}

func (v *NullableGetTradingBotRecurringSubOrdersV5RespData) Set(val *GetTradingBotRecurringSubOrdersV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTradingBotRecurringSubOrdersV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTradingBotRecurringSubOrdersV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTradingBotRecurringSubOrdersV5RespData(val *GetTradingBotRecurringSubOrdersV5RespData) *NullableGetTradingBotRecurringSubOrdersV5RespData {
	return &NullableGetTradingBotRecurringSubOrdersV5RespData{value: val, isSet: true}
}

func (v NullableGetTradingBotRecurringSubOrdersV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTradingBotRecurringSubOrdersV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


