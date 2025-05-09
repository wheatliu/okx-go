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

// checks if the CreateRfqCreateRfqV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRfqCreateRfqV5RespDataInner{}

// CreateRfqCreateRfqV5RespDataInner struct for CreateRfqCreateRfqV5RespDataInner
type CreateRfqCreateRfqV5RespDataInner struct {
	// Whether the RFQ can be partially filled provided that the shape of legs stays the same.
	AllowPartialExecution *bool `json:"allowPartialExecution,omitempty"`
	// The timestamp the RFQ was created. Unix timestamp format in milliseconds.
	CTime *string `json:"cTime,omitempty"`
	// Margin currency.   Only applicable to `cross` `MARGIN` orders in `Spot and futures mode`. The parameter will be ignored in other scenarios.
	Ccy *string `json:"ccy,omitempty"`
	// Client-supplied RFQ ID. This attribute is treated as client sensitive information. It will not be exposed to the Maker, only return empty string.
	ClRfqId *string `json:"clRfqId,omitempty"`
	// The list of counterparties traderCode the RFQ was broadcast to.
	Counterparties []string `json:"counterparties,omitempty"`
	// Instrument ID, e.g. BTC-USDT-SWAP
	InstId *string `json:"instId,omitempty"`
	// An Array of objects containing each leg of the RFQ.
	Legs []CreateRfqCreateRfqV5RespDataInnerLegsInner `json:"legs,omitempty"`
	// Position side.   The default is `net` in the net mode. If not specified, return \"\", which is equivalent to net.   It can only be `long` or `short` in the long/short mode. If not specified, return \"\", which corresponds to the direction that opens new positions for the trade (buy => long, sell => short).   Only applicable to FUTURES/SWAP.
	PosSide *string `json:"posSide,omitempty"`
	// The unique identifier of the RFQ generated by system.
	RfqId *string `json:"rfqId,omitempty"`
	// The direction of the leg. Valid values can be buy or sell.
	Side *string `json:"side,omitempty"`
	// The status of the RFQ.   Valid values can be `active` `canceled` `pending_fill` `filled` `expired` `traded_away` `failed`.   `traded_away` only applies to Maker
	State *string `json:"state,omitempty"`
	// Size of the leg in contracts or spot.
	Sz *string `json:"sz,omitempty"`
	// RFQ tag. The block trade associated with the RFQ will have the same tag.
	Tag *string `json:"tag,omitempty"`
	// Trade mode   Margin mode: `cross` `isolated`   Non-Margin mode: `cash`.   If not provided, tdMode will inherit default values set by the system shown below:   Spot and futures mode & SPOT: `cash`   Buy options in Spot and futures mode and Multi-currency Margin: `isolated`   Other cases: `cross`
	TdMode *string `json:"tdMode,omitempty"`
	// Defines the unit of the “sz” attribute.   Only applicable to instType = SPOT.   The valid enumerations are `base_ccy` and `quote_ccy`. When not specified this is equal to `base_ccy` by default.
	TgtCcy *string `json:"tgtCcy,omitempty"`
	// A unique identifier of taker.
	TraderCode *string `json:"traderCode,omitempty"`
	// The timestamp the RFQ was last updated. Unix timestamp format in milliseconds.
	UTime *string `json:"uTime,omitempty"`
	// The timestamp the RFQ expires. Unix timestamp format in milliseconds.   If all legs are options, the RFQ will expire after 10 minutes; otherwise, the RFQ will expire after 2 minutes.
	ValidUntil *string `json:"validUntil,omitempty"`
}

// NewCreateRfqCreateRfqV5RespDataInner instantiates a new CreateRfqCreateRfqV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRfqCreateRfqV5RespDataInner() *CreateRfqCreateRfqV5RespDataInner {
	this := CreateRfqCreateRfqV5RespDataInner{}
	var cTime string = ""
	this.CTime = &cTime
	var ccy string = ""
	this.Ccy = &ccy
	var clRfqId string = ""
	this.ClRfqId = &clRfqId
	var instId string = ""
	this.InstId = &instId
	var posSide string = ""
	this.PosSide = &posSide
	var rfqId string = ""
	this.RfqId = &rfqId
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
	var tgtCcy string = ""
	this.TgtCcy = &tgtCcy
	var traderCode string = ""
	this.TraderCode = &traderCode
	var uTime string = ""
	this.UTime = &uTime
	var validUntil string = ""
	this.ValidUntil = &validUntil
	return &this
}

// NewCreateRfqCreateRfqV5RespDataInnerWithDefaults instantiates a new CreateRfqCreateRfqV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRfqCreateRfqV5RespDataInnerWithDefaults() *CreateRfqCreateRfqV5RespDataInner {
	this := CreateRfqCreateRfqV5RespDataInner{}
	var cTime string = ""
	this.CTime = &cTime
	var ccy string = ""
	this.Ccy = &ccy
	var clRfqId string = ""
	this.ClRfqId = &clRfqId
	var instId string = ""
	this.InstId = &instId
	var posSide string = ""
	this.PosSide = &posSide
	var rfqId string = ""
	this.RfqId = &rfqId
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
	var tgtCcy string = ""
	this.TgtCcy = &tgtCcy
	var traderCode string = ""
	this.TraderCode = &traderCode
	var uTime string = ""
	this.UTime = &uTime
	var validUntil string = ""
	this.ValidUntil = &validUntil
	return &this
}

// GetAllowPartialExecution returns the AllowPartialExecution field value if set, zero value otherwise.
func (o *CreateRfqCreateRfqV5RespDataInner) GetAllowPartialExecution() bool {
	if o == nil || IsNil(o.AllowPartialExecution) {
		var ret bool
		return ret
	}
	return *o.AllowPartialExecution
}

// GetAllowPartialExecutionOk returns a tuple with the AllowPartialExecution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) GetAllowPartialExecutionOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowPartialExecution) {
		return nil, false
	}
	return o.AllowPartialExecution, true
}

// HasAllowPartialExecution returns a boolean if a field has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) HasAllowPartialExecution() bool {
	if o != nil && !IsNil(o.AllowPartialExecution) {
		return true
	}

	return false
}

// SetAllowPartialExecution gets a reference to the given bool and assigns it to the AllowPartialExecution field.
func (o *CreateRfqCreateRfqV5RespDataInner) SetAllowPartialExecution(v bool) {
	o.AllowPartialExecution = &v
}

// GetCTime returns the CTime field value if set, zero value otherwise.
func (o *CreateRfqCreateRfqV5RespDataInner) GetCTime() string {
	if o == nil || IsNil(o.CTime) {
		var ret string
		return ret
	}
	return *o.CTime
}

// GetCTimeOk returns a tuple with the CTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) GetCTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CTime) {
		return nil, false
	}
	return o.CTime, true
}

// HasCTime returns a boolean if a field has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) HasCTime() bool {
	if o != nil && !IsNil(o.CTime) {
		return true
	}

	return false
}

// SetCTime gets a reference to the given string and assigns it to the CTime field.
func (o *CreateRfqCreateRfqV5RespDataInner) SetCTime(v string) {
	o.CTime = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *CreateRfqCreateRfqV5RespDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *CreateRfqCreateRfqV5RespDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetClRfqId returns the ClRfqId field value if set, zero value otherwise.
func (o *CreateRfqCreateRfqV5RespDataInner) GetClRfqId() string {
	if o == nil || IsNil(o.ClRfqId) {
		var ret string
		return ret
	}
	return *o.ClRfqId
}

// GetClRfqIdOk returns a tuple with the ClRfqId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) GetClRfqIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClRfqId) {
		return nil, false
	}
	return o.ClRfqId, true
}

// HasClRfqId returns a boolean if a field has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) HasClRfqId() bool {
	if o != nil && !IsNil(o.ClRfqId) {
		return true
	}

	return false
}

// SetClRfqId gets a reference to the given string and assigns it to the ClRfqId field.
func (o *CreateRfqCreateRfqV5RespDataInner) SetClRfqId(v string) {
	o.ClRfqId = &v
}

// GetCounterparties returns the Counterparties field value if set, zero value otherwise.
func (o *CreateRfqCreateRfqV5RespDataInner) GetCounterparties() []string {
	if o == nil || IsNil(o.Counterparties) {
		var ret []string
		return ret
	}
	return o.Counterparties
}

// GetCounterpartiesOk returns a tuple with the Counterparties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) GetCounterpartiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Counterparties) {
		return nil, false
	}
	return o.Counterparties, true
}

// HasCounterparties returns a boolean if a field has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) HasCounterparties() bool {
	if o != nil && !IsNil(o.Counterparties) {
		return true
	}

	return false
}

// SetCounterparties gets a reference to the given []string and assigns it to the Counterparties field.
func (o *CreateRfqCreateRfqV5RespDataInner) SetCounterparties(v []string) {
	o.Counterparties = v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *CreateRfqCreateRfqV5RespDataInner) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *CreateRfqCreateRfqV5RespDataInner) SetInstId(v string) {
	o.InstId = &v
}

// GetLegs returns the Legs field value if set, zero value otherwise.
func (o *CreateRfqCreateRfqV5RespDataInner) GetLegs() []CreateRfqCreateRfqV5RespDataInnerLegsInner {
	if o == nil || IsNil(o.Legs) {
		var ret []CreateRfqCreateRfqV5RespDataInnerLegsInner
		return ret
	}
	return o.Legs
}

// GetLegsOk returns a tuple with the Legs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) GetLegsOk() ([]CreateRfqCreateRfqV5RespDataInnerLegsInner, bool) {
	if o == nil || IsNil(o.Legs) {
		return nil, false
	}
	return o.Legs, true
}

// HasLegs returns a boolean if a field has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) HasLegs() bool {
	if o != nil && !IsNil(o.Legs) {
		return true
	}

	return false
}

// SetLegs gets a reference to the given []CreateRfqCreateRfqV5RespDataInnerLegsInner and assigns it to the Legs field.
func (o *CreateRfqCreateRfqV5RespDataInner) SetLegs(v []CreateRfqCreateRfqV5RespDataInnerLegsInner) {
	o.Legs = v
}

// GetPosSide returns the PosSide field value if set, zero value otherwise.
func (o *CreateRfqCreateRfqV5RespDataInner) GetPosSide() string {
	if o == nil || IsNil(o.PosSide) {
		var ret string
		return ret
	}
	return *o.PosSide
}

// GetPosSideOk returns a tuple with the PosSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) GetPosSideOk() (*string, bool) {
	if o == nil || IsNil(o.PosSide) {
		return nil, false
	}
	return o.PosSide, true
}

// HasPosSide returns a boolean if a field has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) HasPosSide() bool {
	if o != nil && !IsNil(o.PosSide) {
		return true
	}

	return false
}

// SetPosSide gets a reference to the given string and assigns it to the PosSide field.
func (o *CreateRfqCreateRfqV5RespDataInner) SetPosSide(v string) {
	o.PosSide = &v
}

// GetRfqId returns the RfqId field value if set, zero value otherwise.
func (o *CreateRfqCreateRfqV5RespDataInner) GetRfqId() string {
	if o == nil || IsNil(o.RfqId) {
		var ret string
		return ret
	}
	return *o.RfqId
}

// GetRfqIdOk returns a tuple with the RfqId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) GetRfqIdOk() (*string, bool) {
	if o == nil || IsNil(o.RfqId) {
		return nil, false
	}
	return o.RfqId, true
}

// HasRfqId returns a boolean if a field has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) HasRfqId() bool {
	if o != nil && !IsNil(o.RfqId) {
		return true
	}

	return false
}

// SetRfqId gets a reference to the given string and assigns it to the RfqId field.
func (o *CreateRfqCreateRfqV5RespDataInner) SetRfqId(v string) {
	o.RfqId = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *CreateRfqCreateRfqV5RespDataInner) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *CreateRfqCreateRfqV5RespDataInner) SetSide(v string) {
	o.Side = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CreateRfqCreateRfqV5RespDataInner) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CreateRfqCreateRfqV5RespDataInner) SetState(v string) {
	o.State = &v
}

// GetSz returns the Sz field value if set, zero value otherwise.
func (o *CreateRfqCreateRfqV5RespDataInner) GetSz() string {
	if o == nil || IsNil(o.Sz) {
		var ret string
		return ret
	}
	return *o.Sz
}

// GetSzOk returns a tuple with the Sz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) GetSzOk() (*string, bool) {
	if o == nil || IsNil(o.Sz) {
		return nil, false
	}
	return o.Sz, true
}

// HasSz returns a boolean if a field has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) HasSz() bool {
	if o != nil && !IsNil(o.Sz) {
		return true
	}

	return false
}

// SetSz gets a reference to the given string and assigns it to the Sz field.
func (o *CreateRfqCreateRfqV5RespDataInner) SetSz(v string) {
	o.Sz = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *CreateRfqCreateRfqV5RespDataInner) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *CreateRfqCreateRfqV5RespDataInner) SetTag(v string) {
	o.Tag = &v
}

// GetTdMode returns the TdMode field value if set, zero value otherwise.
func (o *CreateRfqCreateRfqV5RespDataInner) GetTdMode() string {
	if o == nil || IsNil(o.TdMode) {
		var ret string
		return ret
	}
	return *o.TdMode
}

// GetTdModeOk returns a tuple with the TdMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) GetTdModeOk() (*string, bool) {
	if o == nil || IsNil(o.TdMode) {
		return nil, false
	}
	return o.TdMode, true
}

// HasTdMode returns a boolean if a field has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) HasTdMode() bool {
	if o != nil && !IsNil(o.TdMode) {
		return true
	}

	return false
}

// SetTdMode gets a reference to the given string and assigns it to the TdMode field.
func (o *CreateRfqCreateRfqV5RespDataInner) SetTdMode(v string) {
	o.TdMode = &v
}

// GetTgtCcy returns the TgtCcy field value if set, zero value otherwise.
func (o *CreateRfqCreateRfqV5RespDataInner) GetTgtCcy() string {
	if o == nil || IsNil(o.TgtCcy) {
		var ret string
		return ret
	}
	return *o.TgtCcy
}

// GetTgtCcyOk returns a tuple with the TgtCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) GetTgtCcyOk() (*string, bool) {
	if o == nil || IsNil(o.TgtCcy) {
		return nil, false
	}
	return o.TgtCcy, true
}

// HasTgtCcy returns a boolean if a field has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) HasTgtCcy() bool {
	if o != nil && !IsNil(o.TgtCcy) {
		return true
	}

	return false
}

// SetTgtCcy gets a reference to the given string and assigns it to the TgtCcy field.
func (o *CreateRfqCreateRfqV5RespDataInner) SetTgtCcy(v string) {
	o.TgtCcy = &v
}

// GetTraderCode returns the TraderCode field value if set, zero value otherwise.
func (o *CreateRfqCreateRfqV5RespDataInner) GetTraderCode() string {
	if o == nil || IsNil(o.TraderCode) {
		var ret string
		return ret
	}
	return *o.TraderCode
}

// GetTraderCodeOk returns a tuple with the TraderCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) GetTraderCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TraderCode) {
		return nil, false
	}
	return o.TraderCode, true
}

// HasTraderCode returns a boolean if a field has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) HasTraderCode() bool {
	if o != nil && !IsNil(o.TraderCode) {
		return true
	}

	return false
}

// SetTraderCode gets a reference to the given string and assigns it to the TraderCode field.
func (o *CreateRfqCreateRfqV5RespDataInner) SetTraderCode(v string) {
	o.TraderCode = &v
}

// GetUTime returns the UTime field value if set, zero value otherwise.
func (o *CreateRfqCreateRfqV5RespDataInner) GetUTime() string {
	if o == nil || IsNil(o.UTime) {
		var ret string
		return ret
	}
	return *o.UTime
}

// GetUTimeOk returns a tuple with the UTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) GetUTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UTime) {
		return nil, false
	}
	return o.UTime, true
}

// HasUTime returns a boolean if a field has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) HasUTime() bool {
	if o != nil && !IsNil(o.UTime) {
		return true
	}

	return false
}

// SetUTime gets a reference to the given string and assigns it to the UTime field.
func (o *CreateRfqCreateRfqV5RespDataInner) SetUTime(v string) {
	o.UTime = &v
}

// GetValidUntil returns the ValidUntil field value if set, zero value otherwise.
func (o *CreateRfqCreateRfqV5RespDataInner) GetValidUntil() string {
	if o == nil || IsNil(o.ValidUntil) {
		var ret string
		return ret
	}
	return *o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) GetValidUntilOk() (*string, bool) {
	if o == nil || IsNil(o.ValidUntil) {
		return nil, false
	}
	return o.ValidUntil, true
}

// HasValidUntil returns a boolean if a field has been set.
func (o *CreateRfqCreateRfqV5RespDataInner) HasValidUntil() bool {
	if o != nil && !IsNil(o.ValidUntil) {
		return true
	}

	return false
}

// SetValidUntil gets a reference to the given string and assigns it to the ValidUntil field.
func (o *CreateRfqCreateRfqV5RespDataInner) SetValidUntil(v string) {
	o.ValidUntil = &v
}

func (o CreateRfqCreateRfqV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRfqCreateRfqV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowPartialExecution) {
		toSerialize["allowPartialExecution"] = o.AllowPartialExecution
	}
	if !IsNil(o.CTime) {
		toSerialize["cTime"] = o.CTime
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.ClRfqId) {
		toSerialize["clRfqId"] = o.ClRfqId
	}
	if !IsNil(o.Counterparties) {
		toSerialize["counterparties"] = o.Counterparties
	}
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	if !IsNil(o.Legs) {
		toSerialize["legs"] = o.Legs
	}
	if !IsNil(o.PosSide) {
		toSerialize["posSide"] = o.PosSide
	}
	if !IsNil(o.RfqId) {
		toSerialize["rfqId"] = o.RfqId
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
	if !IsNil(o.TgtCcy) {
		toSerialize["tgtCcy"] = o.TgtCcy
	}
	if !IsNil(o.TraderCode) {
		toSerialize["traderCode"] = o.TraderCode
	}
	if !IsNil(o.UTime) {
		toSerialize["uTime"] = o.UTime
	}
	if !IsNil(o.ValidUntil) {
		toSerialize["validUntil"] = o.ValidUntil
	}
	return toSerialize, nil
}

type NullableCreateRfqCreateRfqV5RespDataInner struct {
	value *CreateRfqCreateRfqV5RespDataInner
	isSet bool
}

func (v NullableCreateRfqCreateRfqV5RespDataInner) Get() *CreateRfqCreateRfqV5RespDataInner {
	return v.value
}

func (v *NullableCreateRfqCreateRfqV5RespDataInner) Set(val *CreateRfqCreateRfqV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRfqCreateRfqV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRfqCreateRfqV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRfqCreateRfqV5RespDataInner(val *CreateRfqCreateRfqV5RespDataInner) *NullableCreateRfqCreateRfqV5RespDataInner {
	return &NullableCreateRfqCreateRfqV5RespDataInner{value: val, isSet: true}
}

func (v NullableCreateRfqCreateRfqV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRfqCreateRfqV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


