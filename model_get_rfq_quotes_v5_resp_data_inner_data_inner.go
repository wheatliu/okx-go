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

// checks if the GetRfqQuotesV5RespDataInnerDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRfqQuotesV5RespDataInnerDataInner{}

// GetRfqQuotesV5RespDataInnerDataInner struct for GetRfqQuotesV5RespDataInnerDataInner
type GetRfqQuotesV5RespDataInnerDataInner struct {
	// The timestamp the Quote was created, Unix timestamp format in milliseconds.
	CTime *string `json:"cTime,omitempty"`
	// Margin currency.   Only applicable to `cross` `MARGIN` orders in `Spot and futures mode`. The parameter will be ignored in other scenarios.
	Ccy *string `json:"ccy,omitempty"`
	// Client-supplied Quote ID. This attribute is treated as client sensitive information. It will not be exposed to the Taker, only return empty string.
	ClQuoteId *string `json:"clQuoteId,omitempty"`
	// Client-supplied RFQ ID. This attribute is treated as client sensitive information. It will not be exposed to the Maker, only return empty string.
	ClRfqId *string `json:"clRfqId,omitempty"`
	// The instrument ID of the quoted leg.
	InstId *string `json:"instId,omitempty"`
	// The legs of the Quote.
	Legs []GetRfqQuotesV5RespDataInnerDataInnerLegsInner `json:"legs,omitempty"`
	// Position side.   The default is `net` in the net mode. If not specified, return \"\", which is equivalent to net.   It can only be `long` or `short` in the long/short mode. If not specified, return \"\", which corresponds to the direction that opens new positions for the trade (buy => long, sell => short).   Only applicable to `FUTURES`/`SWAP`.
	PosSide *string `json:"posSide,omitempty"`
	// The price of the leg.
	Px *string `json:"px,omitempty"`
	// Quote ID.
	QuoteId *string `json:"quoteId,omitempty"`
	// Top level direction of Quote. Its value can be buy or sell.
	QuoteSide *string `json:"quoteSide,omitempty"`
	// Reasons of state. Valid values can be `mmp_canceled`.
	Reason *string `json:"reason,omitempty"`
	// RFQ ID.
	RfqId *string `json:"rfqId,omitempty"`
	// The direction of the leg. Valid values can be buy or sell.
	Side *string `json:"side,omitempty"`
	// The status of the quote. Valid values can be `active` `canceled` `pending_fill` `filled` `expired` or `failed`.
	State *string `json:"state,omitempty"`
	// Size of the leg in contracts or spot.
	Sz *string `json:"sz,omitempty"`
	// Quote tag. The block trade associated with the Quote will have the same tag.
	Tag *string `json:"tag,omitempty"`
	// Trade mode   Margin mode: `cross` `isolated`   Non-Margin mode: `cash`.   If not provided, tdMode will inherit default values set by the system shown below:   Spot and futures mode & SPOT: `cash`   Buy options in Spot and futures mode and Multi-currency Margin: `isolated`   Other cases: `cross`
	TdMode *string `json:"tdMode,omitempty"`
	// Defines the unit of the “sz” attribute.   Only applicable to instType = SPOT.   The valid enumerations are base_ccy and quote_ccy. When not specified this is equal to base_ccy by default.
	TgtCcy *string `json:"tgtCcy,omitempty"`
	// A unique identifier of maker. Empty If the anonymous parameter of the Quote is set to be `true`.
	TraderCode *string `json:"traderCode,omitempty"`
	// The timestamp the Quote was last updated, Unix timestamp format in milliseconds.
	UTime *string `json:"uTime,omitempty"`
	// The timestamp the Quote expires. Unix timestamp format in milliseconds.
	ValidUntil *string `json:"validUntil,omitempty"`
}

// NewGetRfqQuotesV5RespDataInnerDataInner instantiates a new GetRfqQuotesV5RespDataInnerDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRfqQuotesV5RespDataInnerDataInner() *GetRfqQuotesV5RespDataInnerDataInner {
	this := GetRfqQuotesV5RespDataInnerDataInner{}
	var cTime string = ""
	this.CTime = &cTime
	var ccy string = ""
	this.Ccy = &ccy
	var clQuoteId string = ""
	this.ClQuoteId = &clQuoteId
	var clRfqId string = ""
	this.ClRfqId = &clRfqId
	var instId string = ""
	this.InstId = &instId
	var posSide string = ""
	this.PosSide = &posSide
	var px string = ""
	this.Px = &px
	var quoteId string = ""
	this.QuoteId = &quoteId
	var quoteSide string = ""
	this.QuoteSide = &quoteSide
	var reason string = ""
	this.Reason = &reason
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

// NewGetRfqQuotesV5RespDataInnerDataInnerWithDefaults instantiates a new GetRfqQuotesV5RespDataInnerDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRfqQuotesV5RespDataInnerDataInnerWithDefaults() *GetRfqQuotesV5RespDataInnerDataInner {
	this := GetRfqQuotesV5RespDataInnerDataInner{}
	var cTime string = ""
	this.CTime = &cTime
	var ccy string = ""
	this.Ccy = &ccy
	var clQuoteId string = ""
	this.ClQuoteId = &clQuoteId
	var clRfqId string = ""
	this.ClRfqId = &clRfqId
	var instId string = ""
	this.InstId = &instId
	var posSide string = ""
	this.PosSide = &posSide
	var px string = ""
	this.Px = &px
	var quoteId string = ""
	this.QuoteId = &quoteId
	var quoteSide string = ""
	this.QuoteSide = &quoteSide
	var reason string = ""
	this.Reason = &reason
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

// GetCTime returns the CTime field value if set, zero value otherwise.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetCTime() string {
	if o == nil || IsNil(o.CTime) {
		var ret string
		return ret
	}
	return *o.CTime
}

// GetCTimeOk returns a tuple with the CTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetCTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CTime) {
		return nil, false
	}
	return o.CTime, true
}

// HasCTime returns a boolean if a field has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) HasCTime() bool {
	if o != nil && !IsNil(o.CTime) {
		return true
	}

	return false
}

// SetCTime gets a reference to the given string and assigns it to the CTime field.
func (o *GetRfqQuotesV5RespDataInnerDataInner) SetCTime(v string) {
	o.CTime = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetRfqQuotesV5RespDataInnerDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetClQuoteId returns the ClQuoteId field value if set, zero value otherwise.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetClQuoteId() string {
	if o == nil || IsNil(o.ClQuoteId) {
		var ret string
		return ret
	}
	return *o.ClQuoteId
}

// GetClQuoteIdOk returns a tuple with the ClQuoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetClQuoteIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClQuoteId) {
		return nil, false
	}
	return o.ClQuoteId, true
}

// HasClQuoteId returns a boolean if a field has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) HasClQuoteId() bool {
	if o != nil && !IsNil(o.ClQuoteId) {
		return true
	}

	return false
}

// SetClQuoteId gets a reference to the given string and assigns it to the ClQuoteId field.
func (o *GetRfqQuotesV5RespDataInnerDataInner) SetClQuoteId(v string) {
	o.ClQuoteId = &v
}

// GetClRfqId returns the ClRfqId field value if set, zero value otherwise.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetClRfqId() string {
	if o == nil || IsNil(o.ClRfqId) {
		var ret string
		return ret
	}
	return *o.ClRfqId
}

// GetClRfqIdOk returns a tuple with the ClRfqId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetClRfqIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClRfqId) {
		return nil, false
	}
	return o.ClRfqId, true
}

// HasClRfqId returns a boolean if a field has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) HasClRfqId() bool {
	if o != nil && !IsNil(o.ClRfqId) {
		return true
	}

	return false
}

// SetClRfqId gets a reference to the given string and assigns it to the ClRfqId field.
func (o *GetRfqQuotesV5RespDataInnerDataInner) SetClRfqId(v string) {
	o.ClRfqId = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetRfqQuotesV5RespDataInnerDataInner) SetInstId(v string) {
	o.InstId = &v
}

// GetLegs returns the Legs field value if set, zero value otherwise.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetLegs() []GetRfqQuotesV5RespDataInnerDataInnerLegsInner {
	if o == nil || IsNil(o.Legs) {
		var ret []GetRfqQuotesV5RespDataInnerDataInnerLegsInner
		return ret
	}
	return o.Legs
}

// GetLegsOk returns a tuple with the Legs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetLegsOk() ([]GetRfqQuotesV5RespDataInnerDataInnerLegsInner, bool) {
	if o == nil || IsNil(o.Legs) {
		return nil, false
	}
	return o.Legs, true
}

// HasLegs returns a boolean if a field has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) HasLegs() bool {
	if o != nil && !IsNil(o.Legs) {
		return true
	}

	return false
}

// SetLegs gets a reference to the given []GetRfqQuotesV5RespDataInnerDataInnerLegsInner and assigns it to the Legs field.
func (o *GetRfqQuotesV5RespDataInnerDataInner) SetLegs(v []GetRfqQuotesV5RespDataInnerDataInnerLegsInner) {
	o.Legs = v
}

// GetPosSide returns the PosSide field value if set, zero value otherwise.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetPosSide() string {
	if o == nil || IsNil(o.PosSide) {
		var ret string
		return ret
	}
	return *o.PosSide
}

// GetPosSideOk returns a tuple with the PosSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetPosSideOk() (*string, bool) {
	if o == nil || IsNil(o.PosSide) {
		return nil, false
	}
	return o.PosSide, true
}

// HasPosSide returns a boolean if a field has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) HasPosSide() bool {
	if o != nil && !IsNil(o.PosSide) {
		return true
	}

	return false
}

// SetPosSide gets a reference to the given string and assigns it to the PosSide field.
func (o *GetRfqQuotesV5RespDataInnerDataInner) SetPosSide(v string) {
	o.PosSide = &v
}

// GetPx returns the Px field value if set, zero value otherwise.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetPx() string {
	if o == nil || IsNil(o.Px) {
		var ret string
		return ret
	}
	return *o.Px
}

// GetPxOk returns a tuple with the Px field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetPxOk() (*string, bool) {
	if o == nil || IsNil(o.Px) {
		return nil, false
	}
	return o.Px, true
}

// HasPx returns a boolean if a field has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) HasPx() bool {
	if o != nil && !IsNil(o.Px) {
		return true
	}

	return false
}

// SetPx gets a reference to the given string and assigns it to the Px field.
func (o *GetRfqQuotesV5RespDataInnerDataInner) SetPx(v string) {
	o.Px = &v
}

// GetQuoteId returns the QuoteId field value if set, zero value otherwise.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetQuoteId() string {
	if o == nil || IsNil(o.QuoteId) {
		var ret string
		return ret
	}
	return *o.QuoteId
}

// GetQuoteIdOk returns a tuple with the QuoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetQuoteIdOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteId) {
		return nil, false
	}
	return o.QuoteId, true
}

// HasQuoteId returns a boolean if a field has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) HasQuoteId() bool {
	if o != nil && !IsNil(o.QuoteId) {
		return true
	}

	return false
}

// SetQuoteId gets a reference to the given string and assigns it to the QuoteId field.
func (o *GetRfqQuotesV5RespDataInnerDataInner) SetQuoteId(v string) {
	o.QuoteId = &v
}

// GetQuoteSide returns the QuoteSide field value if set, zero value otherwise.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetQuoteSide() string {
	if o == nil || IsNil(o.QuoteSide) {
		var ret string
		return ret
	}
	return *o.QuoteSide
}

// GetQuoteSideOk returns a tuple with the QuoteSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetQuoteSideOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteSide) {
		return nil, false
	}
	return o.QuoteSide, true
}

// HasQuoteSide returns a boolean if a field has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) HasQuoteSide() bool {
	if o != nil && !IsNil(o.QuoteSide) {
		return true
	}

	return false
}

// SetQuoteSide gets a reference to the given string and assigns it to the QuoteSide field.
func (o *GetRfqQuotesV5RespDataInnerDataInner) SetQuoteSide(v string) {
	o.QuoteSide = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *GetRfqQuotesV5RespDataInnerDataInner) SetReason(v string) {
	o.Reason = &v
}

// GetRfqId returns the RfqId field value if set, zero value otherwise.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetRfqId() string {
	if o == nil || IsNil(o.RfqId) {
		var ret string
		return ret
	}
	return *o.RfqId
}

// GetRfqIdOk returns a tuple with the RfqId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetRfqIdOk() (*string, bool) {
	if o == nil || IsNil(o.RfqId) {
		return nil, false
	}
	return o.RfqId, true
}

// HasRfqId returns a boolean if a field has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) HasRfqId() bool {
	if o != nil && !IsNil(o.RfqId) {
		return true
	}

	return false
}

// SetRfqId gets a reference to the given string and assigns it to the RfqId field.
func (o *GetRfqQuotesV5RespDataInnerDataInner) SetRfqId(v string) {
	o.RfqId = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *GetRfqQuotesV5RespDataInnerDataInner) SetSide(v string) {
	o.Side = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GetRfqQuotesV5RespDataInnerDataInner) SetState(v string) {
	o.State = &v
}

// GetSz returns the Sz field value if set, zero value otherwise.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetSz() string {
	if o == nil || IsNil(o.Sz) {
		var ret string
		return ret
	}
	return *o.Sz
}

// GetSzOk returns a tuple with the Sz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetSzOk() (*string, bool) {
	if o == nil || IsNil(o.Sz) {
		return nil, false
	}
	return o.Sz, true
}

// HasSz returns a boolean if a field has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) HasSz() bool {
	if o != nil && !IsNil(o.Sz) {
		return true
	}

	return false
}

// SetSz gets a reference to the given string and assigns it to the Sz field.
func (o *GetRfqQuotesV5RespDataInnerDataInner) SetSz(v string) {
	o.Sz = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *GetRfqQuotesV5RespDataInnerDataInner) SetTag(v string) {
	o.Tag = &v
}

// GetTdMode returns the TdMode field value if set, zero value otherwise.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetTdMode() string {
	if o == nil || IsNil(o.TdMode) {
		var ret string
		return ret
	}
	return *o.TdMode
}

// GetTdModeOk returns a tuple with the TdMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetTdModeOk() (*string, bool) {
	if o == nil || IsNil(o.TdMode) {
		return nil, false
	}
	return o.TdMode, true
}

// HasTdMode returns a boolean if a field has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) HasTdMode() bool {
	if o != nil && !IsNil(o.TdMode) {
		return true
	}

	return false
}

// SetTdMode gets a reference to the given string and assigns it to the TdMode field.
func (o *GetRfqQuotesV5RespDataInnerDataInner) SetTdMode(v string) {
	o.TdMode = &v
}

// GetTgtCcy returns the TgtCcy field value if set, zero value otherwise.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetTgtCcy() string {
	if o == nil || IsNil(o.TgtCcy) {
		var ret string
		return ret
	}
	return *o.TgtCcy
}

// GetTgtCcyOk returns a tuple with the TgtCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetTgtCcyOk() (*string, bool) {
	if o == nil || IsNil(o.TgtCcy) {
		return nil, false
	}
	return o.TgtCcy, true
}

// HasTgtCcy returns a boolean if a field has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) HasTgtCcy() bool {
	if o != nil && !IsNil(o.TgtCcy) {
		return true
	}

	return false
}

// SetTgtCcy gets a reference to the given string and assigns it to the TgtCcy field.
func (o *GetRfqQuotesV5RespDataInnerDataInner) SetTgtCcy(v string) {
	o.TgtCcy = &v
}

// GetTraderCode returns the TraderCode field value if set, zero value otherwise.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetTraderCode() string {
	if o == nil || IsNil(o.TraderCode) {
		var ret string
		return ret
	}
	return *o.TraderCode
}

// GetTraderCodeOk returns a tuple with the TraderCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetTraderCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TraderCode) {
		return nil, false
	}
	return o.TraderCode, true
}

// HasTraderCode returns a boolean if a field has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) HasTraderCode() bool {
	if o != nil && !IsNil(o.TraderCode) {
		return true
	}

	return false
}

// SetTraderCode gets a reference to the given string and assigns it to the TraderCode field.
func (o *GetRfqQuotesV5RespDataInnerDataInner) SetTraderCode(v string) {
	o.TraderCode = &v
}

// GetUTime returns the UTime field value if set, zero value otherwise.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetUTime() string {
	if o == nil || IsNil(o.UTime) {
		var ret string
		return ret
	}
	return *o.UTime
}

// GetUTimeOk returns a tuple with the UTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetUTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UTime) {
		return nil, false
	}
	return o.UTime, true
}

// HasUTime returns a boolean if a field has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) HasUTime() bool {
	if o != nil && !IsNil(o.UTime) {
		return true
	}

	return false
}

// SetUTime gets a reference to the given string and assigns it to the UTime field.
func (o *GetRfqQuotesV5RespDataInnerDataInner) SetUTime(v string) {
	o.UTime = &v
}

// GetValidUntil returns the ValidUntil field value if set, zero value otherwise.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetValidUntil() string {
	if o == nil || IsNil(o.ValidUntil) {
		var ret string
		return ret
	}
	return *o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) GetValidUntilOk() (*string, bool) {
	if o == nil || IsNil(o.ValidUntil) {
		return nil, false
	}
	return o.ValidUntil, true
}

// HasValidUntil returns a boolean if a field has been set.
func (o *GetRfqQuotesV5RespDataInnerDataInner) HasValidUntil() bool {
	if o != nil && !IsNil(o.ValidUntil) {
		return true
	}

	return false
}

// SetValidUntil gets a reference to the given string and assigns it to the ValidUntil field.
func (o *GetRfqQuotesV5RespDataInnerDataInner) SetValidUntil(v string) {
	o.ValidUntil = &v
}

func (o GetRfqQuotesV5RespDataInnerDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRfqQuotesV5RespDataInnerDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CTime) {
		toSerialize["cTime"] = o.CTime
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.ClQuoteId) {
		toSerialize["clQuoteId"] = o.ClQuoteId
	}
	if !IsNil(o.ClRfqId) {
		toSerialize["clRfqId"] = o.ClRfqId
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
	if !IsNil(o.Px) {
		toSerialize["px"] = o.Px
	}
	if !IsNil(o.QuoteId) {
		toSerialize["quoteId"] = o.QuoteId
	}
	if !IsNil(o.QuoteSide) {
		toSerialize["quoteSide"] = o.QuoteSide
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
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

type NullableGetRfqQuotesV5RespDataInnerDataInner struct {
	value *GetRfqQuotesV5RespDataInnerDataInner
	isSet bool
}

func (v NullableGetRfqQuotesV5RespDataInnerDataInner) Get() *GetRfqQuotesV5RespDataInnerDataInner {
	return v.value
}

func (v *NullableGetRfqQuotesV5RespDataInnerDataInner) Set(val *GetRfqQuotesV5RespDataInnerDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRfqQuotesV5RespDataInnerDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRfqQuotesV5RespDataInnerDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRfqQuotesV5RespDataInnerDataInner(val *GetRfqQuotesV5RespDataInnerDataInner) *NullableGetRfqQuotesV5RespDataInnerDataInner {
	return &NullableGetRfqQuotesV5RespDataInnerDataInner{value: val, isSet: true}
}

func (v NullableGetRfqQuotesV5RespDataInnerDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRfqQuotesV5RespDataInnerDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


