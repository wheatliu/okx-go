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

// checks if the GetCopytradingSubpositionsHistoryV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCopytradingSubpositionsHistoryV5RespDataInner{}

// GetCopytradingSubpositionsHistoryV5RespDataInner struct for GetCopytradingSubpositionsHistoryV5RespDataInner
type GetCopytradingSubpositionsHistoryV5RespDataInner struct {
	// Currency
	Ccy *string `json:"ccy,omitempty"`
	// Average price of closing position
	CloseAvgPx *string `json:"closeAvgPx,omitempty"`
	// Quantity of positions that is already closed
	CloseSubPos *string `json:"closeSubPos,omitempty"`
	// Time of closing position
	CloseTime *string `json:"closeTime,omitempty"`
	// Instrument ID, e.g. BTC-USDT-SWAP
	InstId *string `json:"instId,omitempty"`
	// Instrument type
	InstType *string `json:"instType,omitempty"`
	// Leverage
	Lever *string `json:"lever,omitempty"`
	// Margin
	Margin *string `json:"margin,omitempty"`
	// Latest mark price, only applicable to contract
	MarkPx *string `json:"markPx,omitempty"`
	// Margin mode. `cross` `isolated`
	MgnMode *string `json:"mgnMode,omitempty"`
	// Average open price
	OpenAvgPx *string `json:"openAvgPx,omitempty"`
	// Order ID for opening position, only applicable to lead position
	OpenOrdId *string `json:"openOrdId,omitempty"`
	// Time of opening
	OpenTime *string `json:"openTime,omitempty"`
	// Profit and loss
	Pnl *string `json:"pnl,omitempty"`
	// P&L ratio
	PnlRatio *string `json:"pnlRatio,omitempty"`
	// Position side  `long`   `short`   `net`  (long position has positive subPos; short position has negative subPos)
	PosSide *string `json:"posSide,omitempty"`
	// Profit sharing amount, only applicable to copy trading. Note: this parameter is already deprecated.
	// Deprecated
	ProfitSharingAmt *string `json:"profitSharingAmt,omitempty"`
	// Quantity of positions
	SubPos *string `json:"subPos,omitempty"`
	// Lead position ID
	SubPosId *string `json:"subPosId,omitempty"`
	// The type of closing position  `1`：Close position partially;  `2`：Close all
	Type *string `json:"type,omitempty"`
	// Lead trader unique code
	UniqueCode *string `json:"uniqueCode,omitempty"`
}

// NewGetCopytradingSubpositionsHistoryV5RespDataInner instantiates a new GetCopytradingSubpositionsHistoryV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCopytradingSubpositionsHistoryV5RespDataInner() *GetCopytradingSubpositionsHistoryV5RespDataInner {
	this := GetCopytradingSubpositionsHistoryV5RespDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var closeAvgPx string = ""
	this.CloseAvgPx = &closeAvgPx
	var closeSubPos string = ""
	this.CloseSubPos = &closeSubPos
	var closeTime string = ""
	this.CloseTime = &closeTime
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var lever string = ""
	this.Lever = &lever
	var margin string = ""
	this.Margin = &margin
	var markPx string = ""
	this.MarkPx = &markPx
	var mgnMode string = ""
	this.MgnMode = &mgnMode
	var openAvgPx string = ""
	this.OpenAvgPx = &openAvgPx
	var openOrdId string = ""
	this.OpenOrdId = &openOrdId
	var openTime string = ""
	this.OpenTime = &openTime
	var pnl string = ""
	this.Pnl = &pnl
	var pnlRatio string = ""
	this.PnlRatio = &pnlRatio
	var posSide string = ""
	this.PosSide = &posSide
	var profitSharingAmt string = ""
	this.ProfitSharingAmt = &profitSharingAmt
	var subPos string = ""
	this.SubPos = &subPos
	var subPosId string = ""
	this.SubPosId = &subPosId
	var type_ string = ""
	this.Type = &type_
	var uniqueCode string = ""
	this.UniqueCode = &uniqueCode
	return &this
}

// NewGetCopytradingSubpositionsHistoryV5RespDataInnerWithDefaults instantiates a new GetCopytradingSubpositionsHistoryV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCopytradingSubpositionsHistoryV5RespDataInnerWithDefaults() *GetCopytradingSubpositionsHistoryV5RespDataInner {
	this := GetCopytradingSubpositionsHistoryV5RespDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var closeAvgPx string = ""
	this.CloseAvgPx = &closeAvgPx
	var closeSubPos string = ""
	this.CloseSubPos = &closeSubPos
	var closeTime string = ""
	this.CloseTime = &closeTime
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var lever string = ""
	this.Lever = &lever
	var margin string = ""
	this.Margin = &margin
	var markPx string = ""
	this.MarkPx = &markPx
	var mgnMode string = ""
	this.MgnMode = &mgnMode
	var openAvgPx string = ""
	this.OpenAvgPx = &openAvgPx
	var openOrdId string = ""
	this.OpenOrdId = &openOrdId
	var openTime string = ""
	this.OpenTime = &openTime
	var pnl string = ""
	this.Pnl = &pnl
	var pnlRatio string = ""
	this.PnlRatio = &pnlRatio
	var posSide string = ""
	this.PosSide = &posSide
	var profitSharingAmt string = ""
	this.ProfitSharingAmt = &profitSharingAmt
	var subPos string = ""
	this.SubPos = &subPos
	var subPosId string = ""
	this.SubPosId = &subPosId
	var type_ string = ""
	this.Type = &type_
	var uniqueCode string = ""
	this.UniqueCode = &uniqueCode
	return &this
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetCloseAvgPx returns the CloseAvgPx field value if set, zero value otherwise.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetCloseAvgPx() string {
	if o == nil || IsNil(o.CloseAvgPx) {
		var ret string
		return ret
	}
	return *o.CloseAvgPx
}

// GetCloseAvgPxOk returns a tuple with the CloseAvgPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetCloseAvgPxOk() (*string, bool) {
	if o == nil || IsNil(o.CloseAvgPx) {
		return nil, false
	}
	return o.CloseAvgPx, true
}

// HasCloseAvgPx returns a boolean if a field has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasCloseAvgPx() bool {
	if o != nil && !IsNil(o.CloseAvgPx) {
		return true
	}

	return false
}

// SetCloseAvgPx gets a reference to the given string and assigns it to the CloseAvgPx field.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetCloseAvgPx(v string) {
	o.CloseAvgPx = &v
}

// GetCloseSubPos returns the CloseSubPos field value if set, zero value otherwise.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetCloseSubPos() string {
	if o == nil || IsNil(o.CloseSubPos) {
		var ret string
		return ret
	}
	return *o.CloseSubPos
}

// GetCloseSubPosOk returns a tuple with the CloseSubPos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetCloseSubPosOk() (*string, bool) {
	if o == nil || IsNil(o.CloseSubPos) {
		return nil, false
	}
	return o.CloseSubPos, true
}

// HasCloseSubPos returns a boolean if a field has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasCloseSubPos() bool {
	if o != nil && !IsNil(o.CloseSubPos) {
		return true
	}

	return false
}

// SetCloseSubPos gets a reference to the given string and assigns it to the CloseSubPos field.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetCloseSubPos(v string) {
	o.CloseSubPos = &v
}

// GetCloseTime returns the CloseTime field value if set, zero value otherwise.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetCloseTime() string {
	if o == nil || IsNil(o.CloseTime) {
		var ret string
		return ret
	}
	return *o.CloseTime
}

// GetCloseTimeOk returns a tuple with the CloseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetCloseTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CloseTime) {
		return nil, false
	}
	return o.CloseTime, true
}

// HasCloseTime returns a boolean if a field has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasCloseTime() bool {
	if o != nil && !IsNil(o.CloseTime) {
		return true
	}

	return false
}

// SetCloseTime gets a reference to the given string and assigns it to the CloseTime field.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetCloseTime(v string) {
	o.CloseTime = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetInstId(v string) {
	o.InstId = &v
}

// GetInstType returns the InstType field value if set, zero value otherwise.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetInstType() string {
	if o == nil || IsNil(o.InstType) {
		var ret string
		return ret
	}
	return *o.InstType
}

// GetInstTypeOk returns a tuple with the InstType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetInstTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstType) {
		return nil, false
	}
	return o.InstType, true
}

// HasInstType returns a boolean if a field has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasInstType() bool {
	if o != nil && !IsNil(o.InstType) {
		return true
	}

	return false
}

// SetInstType gets a reference to the given string and assigns it to the InstType field.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetInstType(v string) {
	o.InstType = &v
}

// GetLever returns the Lever field value if set, zero value otherwise.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetLever() string {
	if o == nil || IsNil(o.Lever) {
		var ret string
		return ret
	}
	return *o.Lever
}

// GetLeverOk returns a tuple with the Lever field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetLeverOk() (*string, bool) {
	if o == nil || IsNil(o.Lever) {
		return nil, false
	}
	return o.Lever, true
}

// HasLever returns a boolean if a field has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasLever() bool {
	if o != nil && !IsNil(o.Lever) {
		return true
	}

	return false
}

// SetLever gets a reference to the given string and assigns it to the Lever field.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetLever(v string) {
	o.Lever = &v
}

// GetMargin returns the Margin field value if set, zero value otherwise.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetMargin() string {
	if o == nil || IsNil(o.Margin) {
		var ret string
		return ret
	}
	return *o.Margin
}

// GetMarginOk returns a tuple with the Margin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetMarginOk() (*string, bool) {
	if o == nil || IsNil(o.Margin) {
		return nil, false
	}
	return o.Margin, true
}

// HasMargin returns a boolean if a field has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasMargin() bool {
	if o != nil && !IsNil(o.Margin) {
		return true
	}

	return false
}

// SetMargin gets a reference to the given string and assigns it to the Margin field.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetMargin(v string) {
	o.Margin = &v
}

// GetMarkPx returns the MarkPx field value if set, zero value otherwise.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetMarkPx() string {
	if o == nil || IsNil(o.MarkPx) {
		var ret string
		return ret
	}
	return *o.MarkPx
}

// GetMarkPxOk returns a tuple with the MarkPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetMarkPxOk() (*string, bool) {
	if o == nil || IsNil(o.MarkPx) {
		return nil, false
	}
	return o.MarkPx, true
}

// HasMarkPx returns a boolean if a field has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasMarkPx() bool {
	if o != nil && !IsNil(o.MarkPx) {
		return true
	}

	return false
}

// SetMarkPx gets a reference to the given string and assigns it to the MarkPx field.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetMarkPx(v string) {
	o.MarkPx = &v
}

// GetMgnMode returns the MgnMode field value if set, zero value otherwise.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetMgnMode() string {
	if o == nil || IsNil(o.MgnMode) {
		var ret string
		return ret
	}
	return *o.MgnMode
}

// GetMgnModeOk returns a tuple with the MgnMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetMgnModeOk() (*string, bool) {
	if o == nil || IsNil(o.MgnMode) {
		return nil, false
	}
	return o.MgnMode, true
}

// HasMgnMode returns a boolean if a field has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasMgnMode() bool {
	if o != nil && !IsNil(o.MgnMode) {
		return true
	}

	return false
}

// SetMgnMode gets a reference to the given string and assigns it to the MgnMode field.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetMgnMode(v string) {
	o.MgnMode = &v
}

// GetOpenAvgPx returns the OpenAvgPx field value if set, zero value otherwise.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetOpenAvgPx() string {
	if o == nil || IsNil(o.OpenAvgPx) {
		var ret string
		return ret
	}
	return *o.OpenAvgPx
}

// GetOpenAvgPxOk returns a tuple with the OpenAvgPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetOpenAvgPxOk() (*string, bool) {
	if o == nil || IsNil(o.OpenAvgPx) {
		return nil, false
	}
	return o.OpenAvgPx, true
}

// HasOpenAvgPx returns a boolean if a field has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasOpenAvgPx() bool {
	if o != nil && !IsNil(o.OpenAvgPx) {
		return true
	}

	return false
}

// SetOpenAvgPx gets a reference to the given string and assigns it to the OpenAvgPx field.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetOpenAvgPx(v string) {
	o.OpenAvgPx = &v
}

// GetOpenOrdId returns the OpenOrdId field value if set, zero value otherwise.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetOpenOrdId() string {
	if o == nil || IsNil(o.OpenOrdId) {
		var ret string
		return ret
	}
	return *o.OpenOrdId
}

// GetOpenOrdIdOk returns a tuple with the OpenOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetOpenOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenOrdId) {
		return nil, false
	}
	return o.OpenOrdId, true
}

// HasOpenOrdId returns a boolean if a field has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasOpenOrdId() bool {
	if o != nil && !IsNil(o.OpenOrdId) {
		return true
	}

	return false
}

// SetOpenOrdId gets a reference to the given string and assigns it to the OpenOrdId field.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetOpenOrdId(v string) {
	o.OpenOrdId = &v
}

// GetOpenTime returns the OpenTime field value if set, zero value otherwise.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetOpenTime() string {
	if o == nil || IsNil(o.OpenTime) {
		var ret string
		return ret
	}
	return *o.OpenTime
}

// GetOpenTimeOk returns a tuple with the OpenTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetOpenTimeOk() (*string, bool) {
	if o == nil || IsNil(o.OpenTime) {
		return nil, false
	}
	return o.OpenTime, true
}

// HasOpenTime returns a boolean if a field has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasOpenTime() bool {
	if o != nil && !IsNil(o.OpenTime) {
		return true
	}

	return false
}

// SetOpenTime gets a reference to the given string and assigns it to the OpenTime field.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetOpenTime(v string) {
	o.OpenTime = &v
}

// GetPnl returns the Pnl field value if set, zero value otherwise.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetPnl() string {
	if o == nil || IsNil(o.Pnl) {
		var ret string
		return ret
	}
	return *o.Pnl
}

// GetPnlOk returns a tuple with the Pnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetPnlOk() (*string, bool) {
	if o == nil || IsNil(o.Pnl) {
		return nil, false
	}
	return o.Pnl, true
}

// HasPnl returns a boolean if a field has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasPnl() bool {
	if o != nil && !IsNil(o.Pnl) {
		return true
	}

	return false
}

// SetPnl gets a reference to the given string and assigns it to the Pnl field.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetPnl(v string) {
	o.Pnl = &v
}

// GetPnlRatio returns the PnlRatio field value if set, zero value otherwise.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetPnlRatio() string {
	if o == nil || IsNil(o.PnlRatio) {
		var ret string
		return ret
	}
	return *o.PnlRatio
}

// GetPnlRatioOk returns a tuple with the PnlRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetPnlRatioOk() (*string, bool) {
	if o == nil || IsNil(o.PnlRatio) {
		return nil, false
	}
	return o.PnlRatio, true
}

// HasPnlRatio returns a boolean if a field has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasPnlRatio() bool {
	if o != nil && !IsNil(o.PnlRatio) {
		return true
	}

	return false
}

// SetPnlRatio gets a reference to the given string and assigns it to the PnlRatio field.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetPnlRatio(v string) {
	o.PnlRatio = &v
}

// GetPosSide returns the PosSide field value if set, zero value otherwise.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetPosSide() string {
	if o == nil || IsNil(o.PosSide) {
		var ret string
		return ret
	}
	return *o.PosSide
}

// GetPosSideOk returns a tuple with the PosSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetPosSideOk() (*string, bool) {
	if o == nil || IsNil(o.PosSide) {
		return nil, false
	}
	return o.PosSide, true
}

// HasPosSide returns a boolean if a field has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasPosSide() bool {
	if o != nil && !IsNil(o.PosSide) {
		return true
	}

	return false
}

// SetPosSide gets a reference to the given string and assigns it to the PosSide field.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetPosSide(v string) {
	o.PosSide = &v
}

// GetProfitSharingAmt returns the ProfitSharingAmt field value if set, zero value otherwise.
// Deprecated
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetProfitSharingAmt() string {
	if o == nil || IsNil(o.ProfitSharingAmt) {
		var ret string
		return ret
	}
	return *o.ProfitSharingAmt
}

// GetProfitSharingAmtOk returns a tuple with the ProfitSharingAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetProfitSharingAmtOk() (*string, bool) {
	if o == nil || IsNil(o.ProfitSharingAmt) {
		return nil, false
	}
	return o.ProfitSharingAmt, true
}

// HasProfitSharingAmt returns a boolean if a field has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasProfitSharingAmt() bool {
	if o != nil && !IsNil(o.ProfitSharingAmt) {
		return true
	}

	return false
}

// SetProfitSharingAmt gets a reference to the given string and assigns it to the ProfitSharingAmt field.
// Deprecated
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetProfitSharingAmt(v string) {
	o.ProfitSharingAmt = &v
}

// GetSubPos returns the SubPos field value if set, zero value otherwise.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetSubPos() string {
	if o == nil || IsNil(o.SubPos) {
		var ret string
		return ret
	}
	return *o.SubPos
}

// GetSubPosOk returns a tuple with the SubPos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetSubPosOk() (*string, bool) {
	if o == nil || IsNil(o.SubPos) {
		return nil, false
	}
	return o.SubPos, true
}

// HasSubPos returns a boolean if a field has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasSubPos() bool {
	if o != nil && !IsNil(o.SubPos) {
		return true
	}

	return false
}

// SetSubPos gets a reference to the given string and assigns it to the SubPos field.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetSubPos(v string) {
	o.SubPos = &v
}

// GetSubPosId returns the SubPosId field value if set, zero value otherwise.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetSubPosId() string {
	if o == nil || IsNil(o.SubPosId) {
		var ret string
		return ret
	}
	return *o.SubPosId
}

// GetSubPosIdOk returns a tuple with the SubPosId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetSubPosIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubPosId) {
		return nil, false
	}
	return o.SubPosId, true
}

// HasSubPosId returns a boolean if a field has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasSubPosId() bool {
	if o != nil && !IsNil(o.SubPosId) {
		return true
	}

	return false
}

// SetSubPosId gets a reference to the given string and assigns it to the SubPosId field.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetSubPosId(v string) {
	o.SubPosId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetType(v string) {
	o.Type = &v
}

// GetUniqueCode returns the UniqueCode field value if set, zero value otherwise.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetUniqueCode() string {
	if o == nil || IsNil(o.UniqueCode) {
		var ret string
		return ret
	}
	return *o.UniqueCode
}

// GetUniqueCodeOk returns a tuple with the UniqueCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetUniqueCodeOk() (*string, bool) {
	if o == nil || IsNil(o.UniqueCode) {
		return nil, false
	}
	return o.UniqueCode, true
}

// HasUniqueCode returns a boolean if a field has been set.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasUniqueCode() bool {
	if o != nil && !IsNil(o.UniqueCode) {
		return true
	}

	return false
}

// SetUniqueCode gets a reference to the given string and assigns it to the UniqueCode field.
func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetUniqueCode(v string) {
	o.UniqueCode = &v
}

func (o GetCopytradingSubpositionsHistoryV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCopytradingSubpositionsHistoryV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.CloseAvgPx) {
		toSerialize["closeAvgPx"] = o.CloseAvgPx
	}
	if !IsNil(o.CloseSubPos) {
		toSerialize["closeSubPos"] = o.CloseSubPos
	}
	if !IsNil(o.CloseTime) {
		toSerialize["closeTime"] = o.CloseTime
	}
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	if !IsNil(o.InstType) {
		toSerialize["instType"] = o.InstType
	}
	if !IsNil(o.Lever) {
		toSerialize["lever"] = o.Lever
	}
	if !IsNil(o.Margin) {
		toSerialize["margin"] = o.Margin
	}
	if !IsNil(o.MarkPx) {
		toSerialize["markPx"] = o.MarkPx
	}
	if !IsNil(o.MgnMode) {
		toSerialize["mgnMode"] = o.MgnMode
	}
	if !IsNil(o.OpenAvgPx) {
		toSerialize["openAvgPx"] = o.OpenAvgPx
	}
	if !IsNil(o.OpenOrdId) {
		toSerialize["openOrdId"] = o.OpenOrdId
	}
	if !IsNil(o.OpenTime) {
		toSerialize["openTime"] = o.OpenTime
	}
	if !IsNil(o.Pnl) {
		toSerialize["pnl"] = o.Pnl
	}
	if !IsNil(o.PnlRatio) {
		toSerialize["pnlRatio"] = o.PnlRatio
	}
	if !IsNil(o.PosSide) {
		toSerialize["posSide"] = o.PosSide
	}
	if !IsNil(o.ProfitSharingAmt) {
		toSerialize["profitSharingAmt"] = o.ProfitSharingAmt
	}
	if !IsNil(o.SubPos) {
		toSerialize["subPos"] = o.SubPos
	}
	if !IsNil(o.SubPosId) {
		toSerialize["subPosId"] = o.SubPosId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UniqueCode) {
		toSerialize["uniqueCode"] = o.UniqueCode
	}
	return toSerialize, nil
}

type NullableGetCopytradingSubpositionsHistoryV5RespDataInner struct {
	value *GetCopytradingSubpositionsHistoryV5RespDataInner
	isSet bool
}

func (v NullableGetCopytradingSubpositionsHistoryV5RespDataInner) Get() *GetCopytradingSubpositionsHistoryV5RespDataInner {
	return v.value
}

func (v *NullableGetCopytradingSubpositionsHistoryV5RespDataInner) Set(val *GetCopytradingSubpositionsHistoryV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCopytradingSubpositionsHistoryV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCopytradingSubpositionsHistoryV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCopytradingSubpositionsHistoryV5RespDataInner(val *GetCopytradingSubpositionsHistoryV5RespDataInner) *NullableGetCopytradingSubpositionsHistoryV5RespDataInner {
	return &NullableGetCopytradingSubpositionsHistoryV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetCopytradingSubpositionsHistoryV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCopytradingSubpositionsHistoryV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


