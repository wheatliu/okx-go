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

// checks if the GetTradeFillsV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTradeFillsV5RespData{}

// GetTradeFillsV5RespData struct for GetTradeFillsV5RespData
type GetTradeFillsV5RespData struct {
	// Bill ID
	BillId *string `json:"billId,omitempty"`
	// Client Order ID as assigned by the client
	ClOrdId *string `json:"clOrdId,omitempty"`
	// Liquidity taker or maker  `T`: taker  `M`: maker  Not applicable to system orders such as ADL and liquidation
	ExecType *string `json:"execType,omitempty"`
	// The amount of trading fee or rebate. The trading fee deduction is negative, such as '-0.01'; the rebate is positive, such as '0.01'.
	Fee *string `json:"fee,omitempty"`
	// Trading fee or rebate currency
	FeeCcy *string `json:"feeCcy,omitempty"`
	// Fee rate. This field is returned for `SPOT` and `MARGIN` only
	FeeRate *string `json:"feeRate,omitempty"`
	// Forward price when filled   Only applicable to options; return \"\" for other instrument types
	FillFwdPx *string `json:"fillFwdPx,omitempty"`
	// Index price at the moment of trade execution   For cross currency spot pairs, it returns baseCcy-USDT index price. For example, for LTC-ETH, this field returns the index price of LTC-USDT.
	FillIdxPx *string `json:"fillIdxPx,omitempty"`
	// Mark price when filled   Applicable to `FUTURES`, `SWAP`, `OPTION`
	FillMarkPx *string `json:"fillMarkPx,omitempty"`
	// Mark volatility when filled   Only applicable to options; return \"\" for other instrument types
	FillMarkVol *string `json:"fillMarkVol,omitempty"`
	// Last filled profit and loss, applicable to orders which have a trade and aim to close position. It always is 0 in other conditions
	FillPnl *string `json:"fillPnl,omitempty"`
	// Last filled price. It is the same as the px from \"Get bills details\".
	FillPx *string `json:"fillPx,omitempty"`
	// Options price when filled, in the unit of USD   Only applicable to options; return \"\" for other instrument types
	FillPxUsd *string `json:"fillPxUsd,omitempty"`
	// Implied volatility when filled   Only applicable to options; return \"\" for other instrument types
	FillPxVol *string `json:"fillPxVol,omitempty"`
	// Last filled quantity
	FillSz *string `json:"fillSz,omitempty"`
	// Trade time which is the same as `fillTime` for the order channel.
	FillTime *string `json:"fillTime,omitempty"`
	// Instrument ID
	InstId *string `json:"instId,omitempty"`
	// Instrument type
	InstType *string `json:"instType,omitempty"`
	// Order ID
	OrdId *string `json:"ordId,omitempty"`
	// Position side   `long`  `short`     it returns `net` in`net` mode.
	PosSide *string `json:"posSide,omitempty"`
	// Order side,  `buy`  `sell`
	Side *string `json:"side,omitempty"`
	// Transaction type
	SubType *string `json:"subType,omitempty"`
	// Order tag
	Tag *string `json:"tag,omitempty"`
	// Last trade ID
	TradeId *string `json:"tradeId,omitempty"`
	// Data generation time,  Unix timestamp format in milliseconds, e.g. `1597026383085`.
	Ts *string `json:"ts,omitempty"`
}

// NewGetTradeFillsV5RespData instantiates a new GetTradeFillsV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTradeFillsV5RespData() *GetTradeFillsV5RespData {
	this := GetTradeFillsV5RespData{}
	var billId string = ""
	this.BillId = &billId
	var clOrdId string = ""
	this.ClOrdId = &clOrdId
	var execType string = ""
	this.ExecType = &execType
	var fee string = ""
	this.Fee = &fee
	var feeCcy string = ""
	this.FeeCcy = &feeCcy
	var feeRate string = ""
	this.FeeRate = &feeRate
	var fillFwdPx string = ""
	this.FillFwdPx = &fillFwdPx
	var fillIdxPx string = ""
	this.FillIdxPx = &fillIdxPx
	var fillMarkPx string = ""
	this.FillMarkPx = &fillMarkPx
	var fillMarkVol string = ""
	this.FillMarkVol = &fillMarkVol
	var fillPnl string = ""
	this.FillPnl = &fillPnl
	var fillPx string = ""
	this.FillPx = &fillPx
	var fillPxUsd string = ""
	this.FillPxUsd = &fillPxUsd
	var fillPxVol string = ""
	this.FillPxVol = &fillPxVol
	var fillSz string = ""
	this.FillSz = &fillSz
	var fillTime string = ""
	this.FillTime = &fillTime
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var ordId string = ""
	this.OrdId = &ordId
	var posSide string = ""
	this.PosSide = &posSide
	var side string = ""
	this.Side = &side
	var subType string = ""
	this.SubType = &subType
	var tag string = ""
	this.Tag = &tag
	var tradeId string = ""
	this.TradeId = &tradeId
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewGetTradeFillsV5RespDataWithDefaults instantiates a new GetTradeFillsV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTradeFillsV5RespDataWithDefaults() *GetTradeFillsV5RespData {
	this := GetTradeFillsV5RespData{}
	var billId string = ""
	this.BillId = &billId
	var clOrdId string = ""
	this.ClOrdId = &clOrdId
	var execType string = ""
	this.ExecType = &execType
	var fee string = ""
	this.Fee = &fee
	var feeCcy string = ""
	this.FeeCcy = &feeCcy
	var feeRate string = ""
	this.FeeRate = &feeRate
	var fillFwdPx string = ""
	this.FillFwdPx = &fillFwdPx
	var fillIdxPx string = ""
	this.FillIdxPx = &fillIdxPx
	var fillMarkPx string = ""
	this.FillMarkPx = &fillMarkPx
	var fillMarkVol string = ""
	this.FillMarkVol = &fillMarkVol
	var fillPnl string = ""
	this.FillPnl = &fillPnl
	var fillPx string = ""
	this.FillPx = &fillPx
	var fillPxUsd string = ""
	this.FillPxUsd = &fillPxUsd
	var fillPxVol string = ""
	this.FillPxVol = &fillPxVol
	var fillSz string = ""
	this.FillSz = &fillSz
	var fillTime string = ""
	this.FillTime = &fillTime
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var ordId string = ""
	this.OrdId = &ordId
	var posSide string = ""
	this.PosSide = &posSide
	var side string = ""
	this.Side = &side
	var subType string = ""
	this.SubType = &subType
	var tag string = ""
	this.Tag = &tag
	var tradeId string = ""
	this.TradeId = &tradeId
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetBillId returns the BillId field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetBillId() string {
	if o == nil || IsNil(o.BillId) {
		var ret string
		return ret
	}
	return *o.BillId
}

// GetBillIdOk returns a tuple with the BillId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetBillIdOk() (*string, bool) {
	if o == nil || IsNil(o.BillId) {
		return nil, false
	}
	return o.BillId, true
}

// HasBillId returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasBillId() bool {
	if o != nil && !IsNil(o.BillId) {
		return true
	}

	return false
}

// SetBillId gets a reference to the given string and assigns it to the BillId field.
func (o *GetTradeFillsV5RespData) SetBillId(v string) {
	o.BillId = &v
}

// GetClOrdId returns the ClOrdId field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetClOrdId() string {
	if o == nil || IsNil(o.ClOrdId) {
		var ret string
		return ret
	}
	return *o.ClOrdId
}

// GetClOrdIdOk returns a tuple with the ClOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetClOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClOrdId) {
		return nil, false
	}
	return o.ClOrdId, true
}

// HasClOrdId returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasClOrdId() bool {
	if o != nil && !IsNil(o.ClOrdId) {
		return true
	}

	return false
}

// SetClOrdId gets a reference to the given string and assigns it to the ClOrdId field.
func (o *GetTradeFillsV5RespData) SetClOrdId(v string) {
	o.ClOrdId = &v
}

// GetExecType returns the ExecType field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetExecType() string {
	if o == nil || IsNil(o.ExecType) {
		var ret string
		return ret
	}
	return *o.ExecType
}

// GetExecTypeOk returns a tuple with the ExecType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetExecTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExecType) {
		return nil, false
	}
	return o.ExecType, true
}

// HasExecType returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasExecType() bool {
	if o != nil && !IsNil(o.ExecType) {
		return true
	}

	return false
}

// SetExecType gets a reference to the given string and assigns it to the ExecType field.
func (o *GetTradeFillsV5RespData) SetExecType(v string) {
	o.ExecType = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetFee() string {
	if o == nil || IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetFeeOk() (*string, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *GetTradeFillsV5RespData) SetFee(v string) {
	o.Fee = &v
}

// GetFeeCcy returns the FeeCcy field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetFeeCcy() string {
	if o == nil || IsNil(o.FeeCcy) {
		var ret string
		return ret
	}
	return *o.FeeCcy
}

// GetFeeCcyOk returns a tuple with the FeeCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetFeeCcyOk() (*string, bool) {
	if o == nil || IsNil(o.FeeCcy) {
		return nil, false
	}
	return o.FeeCcy, true
}

// HasFeeCcy returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasFeeCcy() bool {
	if o != nil && !IsNil(o.FeeCcy) {
		return true
	}

	return false
}

// SetFeeCcy gets a reference to the given string and assigns it to the FeeCcy field.
func (o *GetTradeFillsV5RespData) SetFeeCcy(v string) {
	o.FeeCcy = &v
}

// GetFeeRate returns the FeeRate field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetFeeRate() string {
	if o == nil || IsNil(o.FeeRate) {
		var ret string
		return ret
	}
	return *o.FeeRate
}

// GetFeeRateOk returns a tuple with the FeeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetFeeRateOk() (*string, bool) {
	if o == nil || IsNil(o.FeeRate) {
		return nil, false
	}
	return o.FeeRate, true
}

// HasFeeRate returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasFeeRate() bool {
	if o != nil && !IsNil(o.FeeRate) {
		return true
	}

	return false
}

// SetFeeRate gets a reference to the given string and assigns it to the FeeRate field.
func (o *GetTradeFillsV5RespData) SetFeeRate(v string) {
	o.FeeRate = &v
}

// GetFillFwdPx returns the FillFwdPx field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetFillFwdPx() string {
	if o == nil || IsNil(o.FillFwdPx) {
		var ret string
		return ret
	}
	return *o.FillFwdPx
}

// GetFillFwdPxOk returns a tuple with the FillFwdPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetFillFwdPxOk() (*string, bool) {
	if o == nil || IsNil(o.FillFwdPx) {
		return nil, false
	}
	return o.FillFwdPx, true
}

// HasFillFwdPx returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasFillFwdPx() bool {
	if o != nil && !IsNil(o.FillFwdPx) {
		return true
	}

	return false
}

// SetFillFwdPx gets a reference to the given string and assigns it to the FillFwdPx field.
func (o *GetTradeFillsV5RespData) SetFillFwdPx(v string) {
	o.FillFwdPx = &v
}

// GetFillIdxPx returns the FillIdxPx field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetFillIdxPx() string {
	if o == nil || IsNil(o.FillIdxPx) {
		var ret string
		return ret
	}
	return *o.FillIdxPx
}

// GetFillIdxPxOk returns a tuple with the FillIdxPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetFillIdxPxOk() (*string, bool) {
	if o == nil || IsNil(o.FillIdxPx) {
		return nil, false
	}
	return o.FillIdxPx, true
}

// HasFillIdxPx returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasFillIdxPx() bool {
	if o != nil && !IsNil(o.FillIdxPx) {
		return true
	}

	return false
}

// SetFillIdxPx gets a reference to the given string and assigns it to the FillIdxPx field.
func (o *GetTradeFillsV5RespData) SetFillIdxPx(v string) {
	o.FillIdxPx = &v
}

// GetFillMarkPx returns the FillMarkPx field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetFillMarkPx() string {
	if o == nil || IsNil(o.FillMarkPx) {
		var ret string
		return ret
	}
	return *o.FillMarkPx
}

// GetFillMarkPxOk returns a tuple with the FillMarkPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetFillMarkPxOk() (*string, bool) {
	if o == nil || IsNil(o.FillMarkPx) {
		return nil, false
	}
	return o.FillMarkPx, true
}

// HasFillMarkPx returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasFillMarkPx() bool {
	if o != nil && !IsNil(o.FillMarkPx) {
		return true
	}

	return false
}

// SetFillMarkPx gets a reference to the given string and assigns it to the FillMarkPx field.
func (o *GetTradeFillsV5RespData) SetFillMarkPx(v string) {
	o.FillMarkPx = &v
}

// GetFillMarkVol returns the FillMarkVol field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetFillMarkVol() string {
	if o == nil || IsNil(o.FillMarkVol) {
		var ret string
		return ret
	}
	return *o.FillMarkVol
}

// GetFillMarkVolOk returns a tuple with the FillMarkVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetFillMarkVolOk() (*string, bool) {
	if o == nil || IsNil(o.FillMarkVol) {
		return nil, false
	}
	return o.FillMarkVol, true
}

// HasFillMarkVol returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasFillMarkVol() bool {
	if o != nil && !IsNil(o.FillMarkVol) {
		return true
	}

	return false
}

// SetFillMarkVol gets a reference to the given string and assigns it to the FillMarkVol field.
func (o *GetTradeFillsV5RespData) SetFillMarkVol(v string) {
	o.FillMarkVol = &v
}

// GetFillPnl returns the FillPnl field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetFillPnl() string {
	if o == nil || IsNil(o.FillPnl) {
		var ret string
		return ret
	}
	return *o.FillPnl
}

// GetFillPnlOk returns a tuple with the FillPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetFillPnlOk() (*string, bool) {
	if o == nil || IsNil(o.FillPnl) {
		return nil, false
	}
	return o.FillPnl, true
}

// HasFillPnl returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasFillPnl() bool {
	if o != nil && !IsNil(o.FillPnl) {
		return true
	}

	return false
}

// SetFillPnl gets a reference to the given string and assigns it to the FillPnl field.
func (o *GetTradeFillsV5RespData) SetFillPnl(v string) {
	o.FillPnl = &v
}

// GetFillPx returns the FillPx field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetFillPx() string {
	if o == nil || IsNil(o.FillPx) {
		var ret string
		return ret
	}
	return *o.FillPx
}

// GetFillPxOk returns a tuple with the FillPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetFillPxOk() (*string, bool) {
	if o == nil || IsNil(o.FillPx) {
		return nil, false
	}
	return o.FillPx, true
}

// HasFillPx returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasFillPx() bool {
	if o != nil && !IsNil(o.FillPx) {
		return true
	}

	return false
}

// SetFillPx gets a reference to the given string and assigns it to the FillPx field.
func (o *GetTradeFillsV5RespData) SetFillPx(v string) {
	o.FillPx = &v
}

// GetFillPxUsd returns the FillPxUsd field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetFillPxUsd() string {
	if o == nil || IsNil(o.FillPxUsd) {
		var ret string
		return ret
	}
	return *o.FillPxUsd
}

// GetFillPxUsdOk returns a tuple with the FillPxUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetFillPxUsdOk() (*string, bool) {
	if o == nil || IsNil(o.FillPxUsd) {
		return nil, false
	}
	return o.FillPxUsd, true
}

// HasFillPxUsd returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasFillPxUsd() bool {
	if o != nil && !IsNil(o.FillPxUsd) {
		return true
	}

	return false
}

// SetFillPxUsd gets a reference to the given string and assigns it to the FillPxUsd field.
func (o *GetTradeFillsV5RespData) SetFillPxUsd(v string) {
	o.FillPxUsd = &v
}

// GetFillPxVol returns the FillPxVol field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetFillPxVol() string {
	if o == nil || IsNil(o.FillPxVol) {
		var ret string
		return ret
	}
	return *o.FillPxVol
}

// GetFillPxVolOk returns a tuple with the FillPxVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetFillPxVolOk() (*string, bool) {
	if o == nil || IsNil(o.FillPxVol) {
		return nil, false
	}
	return o.FillPxVol, true
}

// HasFillPxVol returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasFillPxVol() bool {
	if o != nil && !IsNil(o.FillPxVol) {
		return true
	}

	return false
}

// SetFillPxVol gets a reference to the given string and assigns it to the FillPxVol field.
func (o *GetTradeFillsV5RespData) SetFillPxVol(v string) {
	o.FillPxVol = &v
}

// GetFillSz returns the FillSz field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetFillSz() string {
	if o == nil || IsNil(o.FillSz) {
		var ret string
		return ret
	}
	return *o.FillSz
}

// GetFillSzOk returns a tuple with the FillSz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetFillSzOk() (*string, bool) {
	if o == nil || IsNil(o.FillSz) {
		return nil, false
	}
	return o.FillSz, true
}

// HasFillSz returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasFillSz() bool {
	if o != nil && !IsNil(o.FillSz) {
		return true
	}

	return false
}

// SetFillSz gets a reference to the given string and assigns it to the FillSz field.
func (o *GetTradeFillsV5RespData) SetFillSz(v string) {
	o.FillSz = &v
}

// GetFillTime returns the FillTime field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetFillTime() string {
	if o == nil || IsNil(o.FillTime) {
		var ret string
		return ret
	}
	return *o.FillTime
}

// GetFillTimeOk returns a tuple with the FillTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetFillTimeOk() (*string, bool) {
	if o == nil || IsNil(o.FillTime) {
		return nil, false
	}
	return o.FillTime, true
}

// HasFillTime returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasFillTime() bool {
	if o != nil && !IsNil(o.FillTime) {
		return true
	}

	return false
}

// SetFillTime gets a reference to the given string and assigns it to the FillTime field.
func (o *GetTradeFillsV5RespData) SetFillTime(v string) {
	o.FillTime = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetTradeFillsV5RespData) SetInstId(v string) {
	o.InstId = &v
}

// GetInstType returns the InstType field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetInstType() string {
	if o == nil || IsNil(o.InstType) {
		var ret string
		return ret
	}
	return *o.InstType
}

// GetInstTypeOk returns a tuple with the InstType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetInstTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstType) {
		return nil, false
	}
	return o.InstType, true
}

// HasInstType returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasInstType() bool {
	if o != nil && !IsNil(o.InstType) {
		return true
	}

	return false
}

// SetInstType gets a reference to the given string and assigns it to the InstType field.
func (o *GetTradeFillsV5RespData) SetInstType(v string) {
	o.InstType = &v
}

// GetOrdId returns the OrdId field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetOrdId() string {
	if o == nil || IsNil(o.OrdId) {
		var ret string
		return ret
	}
	return *o.OrdId
}

// GetOrdIdOk returns a tuple with the OrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrdId) {
		return nil, false
	}
	return o.OrdId, true
}

// HasOrdId returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasOrdId() bool {
	if o != nil && !IsNil(o.OrdId) {
		return true
	}

	return false
}

// SetOrdId gets a reference to the given string and assigns it to the OrdId field.
func (o *GetTradeFillsV5RespData) SetOrdId(v string) {
	o.OrdId = &v
}

// GetPosSide returns the PosSide field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetPosSide() string {
	if o == nil || IsNil(o.PosSide) {
		var ret string
		return ret
	}
	return *o.PosSide
}

// GetPosSideOk returns a tuple with the PosSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetPosSideOk() (*string, bool) {
	if o == nil || IsNil(o.PosSide) {
		return nil, false
	}
	return o.PosSide, true
}

// HasPosSide returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasPosSide() bool {
	if o != nil && !IsNil(o.PosSide) {
		return true
	}

	return false
}

// SetPosSide gets a reference to the given string and assigns it to the PosSide field.
func (o *GetTradeFillsV5RespData) SetPosSide(v string) {
	o.PosSide = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *GetTradeFillsV5RespData) SetSide(v string) {
	o.Side = &v
}

// GetSubType returns the SubType field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetSubType() string {
	if o == nil || IsNil(o.SubType) {
		var ret string
		return ret
	}
	return *o.SubType
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetSubTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SubType) {
		return nil, false
	}
	return o.SubType, true
}

// HasSubType returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasSubType() bool {
	if o != nil && !IsNil(o.SubType) {
		return true
	}

	return false
}

// SetSubType gets a reference to the given string and assigns it to the SubType field.
func (o *GetTradeFillsV5RespData) SetSubType(v string) {
	o.SubType = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *GetTradeFillsV5RespData) SetTag(v string) {
	o.Tag = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetTradeId() string {
	if o == nil || IsNil(o.TradeId) {
		var ret string
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetTradeIdOk() (*string, bool) {
	if o == nil || IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasTradeId() bool {
	if o != nil && !IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given string and assigns it to the TradeId field.
func (o *GetTradeFillsV5RespData) SetTradeId(v string) {
	o.TradeId = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetTradeFillsV5RespData) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeFillsV5RespData) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetTradeFillsV5RespData) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetTradeFillsV5RespData) SetTs(v string) {
	o.Ts = &v
}

func (o GetTradeFillsV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTradeFillsV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillId) {
		toSerialize["billId"] = o.BillId
	}
	if !IsNil(o.ClOrdId) {
		toSerialize["clOrdId"] = o.ClOrdId
	}
	if !IsNil(o.ExecType) {
		toSerialize["execType"] = o.ExecType
	}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.FeeCcy) {
		toSerialize["feeCcy"] = o.FeeCcy
	}
	if !IsNil(o.FeeRate) {
		toSerialize["feeRate"] = o.FeeRate
	}
	if !IsNil(o.FillFwdPx) {
		toSerialize["fillFwdPx"] = o.FillFwdPx
	}
	if !IsNil(o.FillIdxPx) {
		toSerialize["fillIdxPx"] = o.FillIdxPx
	}
	if !IsNil(o.FillMarkPx) {
		toSerialize["fillMarkPx"] = o.FillMarkPx
	}
	if !IsNil(o.FillMarkVol) {
		toSerialize["fillMarkVol"] = o.FillMarkVol
	}
	if !IsNil(o.FillPnl) {
		toSerialize["fillPnl"] = o.FillPnl
	}
	if !IsNil(o.FillPx) {
		toSerialize["fillPx"] = o.FillPx
	}
	if !IsNil(o.FillPxUsd) {
		toSerialize["fillPxUsd"] = o.FillPxUsd
	}
	if !IsNil(o.FillPxVol) {
		toSerialize["fillPxVol"] = o.FillPxVol
	}
	if !IsNil(o.FillSz) {
		toSerialize["fillSz"] = o.FillSz
	}
	if !IsNil(o.FillTime) {
		toSerialize["fillTime"] = o.FillTime
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
	if !IsNil(o.PosSide) {
		toSerialize["posSide"] = o.PosSide
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !IsNil(o.SubType) {
		toSerialize["subType"] = o.SubType
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.TradeId) {
		toSerialize["tradeId"] = o.TradeId
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableGetTradeFillsV5RespData struct {
	value *GetTradeFillsV5RespData
	isSet bool
}

func (v NullableGetTradeFillsV5RespData) Get() *GetTradeFillsV5RespData {
	return v.value
}

func (v *NullableGetTradeFillsV5RespData) Set(val *GetTradeFillsV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTradeFillsV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTradeFillsV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTradeFillsV5RespData(val *GetTradeFillsV5RespData) *NullableGetTradeFillsV5RespData {
	return &NullableGetTradeFillsV5RespData{value: val, isSet: true}
}

func (v NullableGetTradeFillsV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTradeFillsV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


