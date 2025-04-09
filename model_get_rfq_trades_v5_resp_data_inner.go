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

// checks if the GetRfqTradesV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRfqTradesV5RespDataInner{}

// GetRfqTradesV5RespDataInner struct for GetRfqTradesV5RespDataInner
type GetRfqTradesV5RespDataInner struct {
	// Block trade ID.
	BlockTdId *string `json:"blockTdId,omitempty"`
	// The time the trade was executed. Unix timestamp in milliseconds.
	CTime *string `json:"cTime,omitempty"`
	// Client-supplied Quote ID. This attribute is treated as client sensitive information. It will not be exposed to the Taker, only return empty string.
	ClQuoteId *string `json:"clQuoteId,omitempty"`
	// Client-supplied RFQ ID. This attribute is treated as client sensitive information. It will not be exposed to the Maker, only return empty string.
	ClRfqId *string `json:"clRfqId,omitempty"`
	// Error code for unsuccessful trades.   It is \"\" for successful trade.
	ErrorCode *string `json:"errorCode,omitempty"`
	// Fee. Negative number represents the user transaction fee charged by the platform. Positive number represents rebate.
	Fee *string `json:"fee,omitempty"`
	// Fee currency
	FeeCcy *string `json:"feeCcy,omitempty"`
	// Instrument ID, e.g. `BTC-USDT-SWAP`
	InstId *string `json:"instId,omitempty"`
	// Whether the trade is filled successfully
	IsSuccessful *bool `json:"isSuccessful,omitempty"`
	// Legs of trade
	Legs []GetRfqTradesV5RespDataInnerLegsInner `json:"legs,omitempty"`
	// A unique identifier of the Maker. Empty if the anonymous parameter of the Quote is set to be `true`.
	MTraderCode *string `json:"mTraderCode,omitempty"`
	// The price the leg executed
	Px *string `json:"px,omitempty"`
	// Quote ID.
	QuoteId *string `json:"quoteId,omitempty"`
	// RFQ ID.
	RfqId *string `json:"rfqId,omitempty"`
	// The direction of the leg. Valid value can be buy or sell.
	Side *string `json:"side,omitempty"`
	// Size of the leg in contracts or spot.
	Sz *string `json:"sz,omitempty"`
	// A unique identifier of the Taker. Empty if the anonymous parameter of the RFQ is set to be `true`.
	TTraderCode *string `json:"tTraderCode,omitempty"`
	// Trade tag. The block trade will have the tag of the RFQ or Quote it corresponds to.
	Tag *string `json:"tag,omitempty"`
	// Last traded ID.
	TradeId *string `json:"tradeId,omitempty"`
}

// NewGetRfqTradesV5RespDataInner instantiates a new GetRfqTradesV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRfqTradesV5RespDataInner() *GetRfqTradesV5RespDataInner {
	this := GetRfqTradesV5RespDataInner{}
	var blockTdId string = ""
	this.BlockTdId = &blockTdId
	var cTime string = ""
	this.CTime = &cTime
	var clQuoteId string = ""
	this.ClQuoteId = &clQuoteId
	var clRfqId string = ""
	this.ClRfqId = &clRfqId
	var errorCode string = ""
	this.ErrorCode = &errorCode
	var fee string = ""
	this.Fee = &fee
	var feeCcy string = ""
	this.FeeCcy = &feeCcy
	var instId string = ""
	this.InstId = &instId
	var mTraderCode string = ""
	this.MTraderCode = &mTraderCode
	var px string = ""
	this.Px = &px
	var quoteId string = ""
	this.QuoteId = &quoteId
	var rfqId string = ""
	this.RfqId = &rfqId
	var side string = ""
	this.Side = &side
	var sz string = ""
	this.Sz = &sz
	var tTraderCode string = ""
	this.TTraderCode = &tTraderCode
	var tag string = ""
	this.Tag = &tag
	var tradeId string = ""
	this.TradeId = &tradeId
	return &this
}

// NewGetRfqTradesV5RespDataInnerWithDefaults instantiates a new GetRfqTradesV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRfqTradesV5RespDataInnerWithDefaults() *GetRfqTradesV5RespDataInner {
	this := GetRfqTradesV5RespDataInner{}
	var blockTdId string = ""
	this.BlockTdId = &blockTdId
	var cTime string = ""
	this.CTime = &cTime
	var clQuoteId string = ""
	this.ClQuoteId = &clQuoteId
	var clRfqId string = ""
	this.ClRfqId = &clRfqId
	var errorCode string = ""
	this.ErrorCode = &errorCode
	var fee string = ""
	this.Fee = &fee
	var feeCcy string = ""
	this.FeeCcy = &feeCcy
	var instId string = ""
	this.InstId = &instId
	var mTraderCode string = ""
	this.MTraderCode = &mTraderCode
	var px string = ""
	this.Px = &px
	var quoteId string = ""
	this.QuoteId = &quoteId
	var rfqId string = ""
	this.RfqId = &rfqId
	var side string = ""
	this.Side = &side
	var sz string = ""
	this.Sz = &sz
	var tTraderCode string = ""
	this.TTraderCode = &tTraderCode
	var tag string = ""
	this.Tag = &tag
	var tradeId string = ""
	this.TradeId = &tradeId
	return &this
}

// GetBlockTdId returns the BlockTdId field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInner) GetBlockTdId() string {
	if o == nil || IsNil(o.BlockTdId) {
		var ret string
		return ret
	}
	return *o.BlockTdId
}

// GetBlockTdIdOk returns a tuple with the BlockTdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInner) GetBlockTdIdOk() (*string, bool) {
	if o == nil || IsNil(o.BlockTdId) {
		return nil, false
	}
	return o.BlockTdId, true
}

// HasBlockTdId returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInner) HasBlockTdId() bool {
	if o != nil && !IsNil(o.BlockTdId) {
		return true
	}

	return false
}

// SetBlockTdId gets a reference to the given string and assigns it to the BlockTdId field.
func (o *GetRfqTradesV5RespDataInner) SetBlockTdId(v string) {
	o.BlockTdId = &v
}

// GetCTime returns the CTime field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInner) GetCTime() string {
	if o == nil || IsNil(o.CTime) {
		var ret string
		return ret
	}
	return *o.CTime
}

// GetCTimeOk returns a tuple with the CTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInner) GetCTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CTime) {
		return nil, false
	}
	return o.CTime, true
}

// HasCTime returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInner) HasCTime() bool {
	if o != nil && !IsNil(o.CTime) {
		return true
	}

	return false
}

// SetCTime gets a reference to the given string and assigns it to the CTime field.
func (o *GetRfqTradesV5RespDataInner) SetCTime(v string) {
	o.CTime = &v
}

// GetClQuoteId returns the ClQuoteId field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInner) GetClQuoteId() string {
	if o == nil || IsNil(o.ClQuoteId) {
		var ret string
		return ret
	}
	return *o.ClQuoteId
}

// GetClQuoteIdOk returns a tuple with the ClQuoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInner) GetClQuoteIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClQuoteId) {
		return nil, false
	}
	return o.ClQuoteId, true
}

// HasClQuoteId returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInner) HasClQuoteId() bool {
	if o != nil && !IsNil(o.ClQuoteId) {
		return true
	}

	return false
}

// SetClQuoteId gets a reference to the given string and assigns it to the ClQuoteId field.
func (o *GetRfqTradesV5RespDataInner) SetClQuoteId(v string) {
	o.ClQuoteId = &v
}

// GetClRfqId returns the ClRfqId field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInner) GetClRfqId() string {
	if o == nil || IsNil(o.ClRfqId) {
		var ret string
		return ret
	}
	return *o.ClRfqId
}

// GetClRfqIdOk returns a tuple with the ClRfqId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInner) GetClRfqIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClRfqId) {
		return nil, false
	}
	return o.ClRfqId, true
}

// HasClRfqId returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInner) HasClRfqId() bool {
	if o != nil && !IsNil(o.ClRfqId) {
		return true
	}

	return false
}

// SetClRfqId gets a reference to the given string and assigns it to the ClRfqId field.
func (o *GetRfqTradesV5RespDataInner) SetClRfqId(v string) {
	o.ClRfqId = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInner) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInner) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInner) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *GetRfqTradesV5RespDataInner) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInner) GetFee() string {
	if o == nil || IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInner) GetFeeOk() (*string, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInner) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *GetRfqTradesV5RespDataInner) SetFee(v string) {
	o.Fee = &v
}

// GetFeeCcy returns the FeeCcy field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInner) GetFeeCcy() string {
	if o == nil || IsNil(o.FeeCcy) {
		var ret string
		return ret
	}
	return *o.FeeCcy
}

// GetFeeCcyOk returns a tuple with the FeeCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInner) GetFeeCcyOk() (*string, bool) {
	if o == nil || IsNil(o.FeeCcy) {
		return nil, false
	}
	return o.FeeCcy, true
}

// HasFeeCcy returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInner) HasFeeCcy() bool {
	if o != nil && !IsNil(o.FeeCcy) {
		return true
	}

	return false
}

// SetFeeCcy gets a reference to the given string and assigns it to the FeeCcy field.
func (o *GetRfqTradesV5RespDataInner) SetFeeCcy(v string) {
	o.FeeCcy = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInner) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInner) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInner) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetRfqTradesV5RespDataInner) SetInstId(v string) {
	o.InstId = &v
}

// GetIsSuccessful returns the IsSuccessful field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInner) GetIsSuccessful() bool {
	if o == nil || IsNil(o.IsSuccessful) {
		var ret bool
		return ret
	}
	return *o.IsSuccessful
}

// GetIsSuccessfulOk returns a tuple with the IsSuccessful field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInner) GetIsSuccessfulOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSuccessful) {
		return nil, false
	}
	return o.IsSuccessful, true
}

// HasIsSuccessful returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInner) HasIsSuccessful() bool {
	if o != nil && !IsNil(o.IsSuccessful) {
		return true
	}

	return false
}

// SetIsSuccessful gets a reference to the given bool and assigns it to the IsSuccessful field.
func (o *GetRfqTradesV5RespDataInner) SetIsSuccessful(v bool) {
	o.IsSuccessful = &v
}

// GetLegs returns the Legs field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInner) GetLegs() []GetRfqTradesV5RespDataInnerLegsInner {
	if o == nil || IsNil(o.Legs) {
		var ret []GetRfqTradesV5RespDataInnerLegsInner
		return ret
	}
	return o.Legs
}

// GetLegsOk returns a tuple with the Legs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInner) GetLegsOk() ([]GetRfqTradesV5RespDataInnerLegsInner, bool) {
	if o == nil || IsNil(o.Legs) {
		return nil, false
	}
	return o.Legs, true
}

// HasLegs returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInner) HasLegs() bool {
	if o != nil && !IsNil(o.Legs) {
		return true
	}

	return false
}

// SetLegs gets a reference to the given []GetRfqTradesV5RespDataInnerLegsInner and assigns it to the Legs field.
func (o *GetRfqTradesV5RespDataInner) SetLegs(v []GetRfqTradesV5RespDataInnerLegsInner) {
	o.Legs = v
}

// GetMTraderCode returns the MTraderCode field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInner) GetMTraderCode() string {
	if o == nil || IsNil(o.MTraderCode) {
		var ret string
		return ret
	}
	return *o.MTraderCode
}

// GetMTraderCodeOk returns a tuple with the MTraderCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInner) GetMTraderCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MTraderCode) {
		return nil, false
	}
	return o.MTraderCode, true
}

// HasMTraderCode returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInner) HasMTraderCode() bool {
	if o != nil && !IsNil(o.MTraderCode) {
		return true
	}

	return false
}

// SetMTraderCode gets a reference to the given string and assigns it to the MTraderCode field.
func (o *GetRfqTradesV5RespDataInner) SetMTraderCode(v string) {
	o.MTraderCode = &v
}

// GetPx returns the Px field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInner) GetPx() string {
	if o == nil || IsNil(o.Px) {
		var ret string
		return ret
	}
	return *o.Px
}

// GetPxOk returns a tuple with the Px field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInner) GetPxOk() (*string, bool) {
	if o == nil || IsNil(o.Px) {
		return nil, false
	}
	return o.Px, true
}

// HasPx returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInner) HasPx() bool {
	if o != nil && !IsNil(o.Px) {
		return true
	}

	return false
}

// SetPx gets a reference to the given string and assigns it to the Px field.
func (o *GetRfqTradesV5RespDataInner) SetPx(v string) {
	o.Px = &v
}

// GetQuoteId returns the QuoteId field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInner) GetQuoteId() string {
	if o == nil || IsNil(o.QuoteId) {
		var ret string
		return ret
	}
	return *o.QuoteId
}

// GetQuoteIdOk returns a tuple with the QuoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInner) GetQuoteIdOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteId) {
		return nil, false
	}
	return o.QuoteId, true
}

// HasQuoteId returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInner) HasQuoteId() bool {
	if o != nil && !IsNil(o.QuoteId) {
		return true
	}

	return false
}

// SetQuoteId gets a reference to the given string and assigns it to the QuoteId field.
func (o *GetRfqTradesV5RespDataInner) SetQuoteId(v string) {
	o.QuoteId = &v
}

// GetRfqId returns the RfqId field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInner) GetRfqId() string {
	if o == nil || IsNil(o.RfqId) {
		var ret string
		return ret
	}
	return *o.RfqId
}

// GetRfqIdOk returns a tuple with the RfqId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInner) GetRfqIdOk() (*string, bool) {
	if o == nil || IsNil(o.RfqId) {
		return nil, false
	}
	return o.RfqId, true
}

// HasRfqId returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInner) HasRfqId() bool {
	if o != nil && !IsNil(o.RfqId) {
		return true
	}

	return false
}

// SetRfqId gets a reference to the given string and assigns it to the RfqId field.
func (o *GetRfqTradesV5RespDataInner) SetRfqId(v string) {
	o.RfqId = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInner) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInner) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInner) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *GetRfqTradesV5RespDataInner) SetSide(v string) {
	o.Side = &v
}

// GetSz returns the Sz field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInner) GetSz() string {
	if o == nil || IsNil(o.Sz) {
		var ret string
		return ret
	}
	return *o.Sz
}

// GetSzOk returns a tuple with the Sz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInner) GetSzOk() (*string, bool) {
	if o == nil || IsNil(o.Sz) {
		return nil, false
	}
	return o.Sz, true
}

// HasSz returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInner) HasSz() bool {
	if o != nil && !IsNil(o.Sz) {
		return true
	}

	return false
}

// SetSz gets a reference to the given string and assigns it to the Sz field.
func (o *GetRfqTradesV5RespDataInner) SetSz(v string) {
	o.Sz = &v
}

// GetTTraderCode returns the TTraderCode field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInner) GetTTraderCode() string {
	if o == nil || IsNil(o.TTraderCode) {
		var ret string
		return ret
	}
	return *o.TTraderCode
}

// GetTTraderCodeOk returns a tuple with the TTraderCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInner) GetTTraderCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TTraderCode) {
		return nil, false
	}
	return o.TTraderCode, true
}

// HasTTraderCode returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInner) HasTTraderCode() bool {
	if o != nil && !IsNil(o.TTraderCode) {
		return true
	}

	return false
}

// SetTTraderCode gets a reference to the given string and assigns it to the TTraderCode field.
func (o *GetRfqTradesV5RespDataInner) SetTTraderCode(v string) {
	o.TTraderCode = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInner) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInner) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInner) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *GetRfqTradesV5RespDataInner) SetTag(v string) {
	o.Tag = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInner) GetTradeId() string {
	if o == nil || IsNil(o.TradeId) {
		var ret string
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInner) GetTradeIdOk() (*string, bool) {
	if o == nil || IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInner) HasTradeId() bool {
	if o != nil && !IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given string and assigns it to the TradeId field.
func (o *GetRfqTradesV5RespDataInner) SetTradeId(v string) {
	o.TradeId = &v
}

func (o GetRfqTradesV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRfqTradesV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockTdId) {
		toSerialize["blockTdId"] = o.BlockTdId
	}
	if !IsNil(o.CTime) {
		toSerialize["cTime"] = o.CTime
	}
	if !IsNil(o.ClQuoteId) {
		toSerialize["clQuoteId"] = o.ClQuoteId
	}
	if !IsNil(o.ClRfqId) {
		toSerialize["clRfqId"] = o.ClRfqId
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
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
	if !IsNil(o.IsSuccessful) {
		toSerialize["isSuccessful"] = o.IsSuccessful
	}
	if !IsNil(o.Legs) {
		toSerialize["legs"] = o.Legs
	}
	if !IsNil(o.MTraderCode) {
		toSerialize["mTraderCode"] = o.MTraderCode
	}
	if !IsNil(o.Px) {
		toSerialize["px"] = o.Px
	}
	if !IsNil(o.QuoteId) {
		toSerialize["quoteId"] = o.QuoteId
	}
	if !IsNil(o.RfqId) {
		toSerialize["rfqId"] = o.RfqId
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !IsNil(o.Sz) {
		toSerialize["sz"] = o.Sz
	}
	if !IsNil(o.TTraderCode) {
		toSerialize["tTraderCode"] = o.TTraderCode
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.TradeId) {
		toSerialize["tradeId"] = o.TradeId
	}
	return toSerialize, nil
}

type NullableGetRfqTradesV5RespDataInner struct {
	value *GetRfqTradesV5RespDataInner
	isSet bool
}

func (v NullableGetRfqTradesV5RespDataInner) Get() *GetRfqTradesV5RespDataInner {
	return v.value
}

func (v *NullableGetRfqTradesV5RespDataInner) Set(val *GetRfqTradesV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRfqTradesV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRfqTradesV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRfqTradesV5RespDataInner(val *GetRfqTradesV5RespDataInner) *NullableGetRfqTradesV5RespDataInner {
	return &NullableGetRfqTradesV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetRfqTradesV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRfqTradesV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


