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

// checks if the GetTradingBotGridSubOrdersV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTradingBotGridSubOrdersV5RespDataInner{}

// GetTradingBotGridSubOrdersV5RespDataInner struct for GetTradingBotGridSubOrdersV5RespDataInner
type GetTradingBotGridSubOrdersV5RespDataInner struct {
	// Sub order accumulated fill quantity
	AccFillSz *string `json:"accFillSz,omitempty"`
	// Client-supplied Algo ID
	AlgoClOrdId *string `json:"algoClOrdId,omitempty"`
	// Algo ID
	AlgoId *string `json:"algoId,omitempty"`
	// Algo order type  `grid`: Spot grid  `contract_grid`: Contract grid
	AlgoOrdType *string `json:"algoOrdType,omitempty"`
	// Sub order average filled price
	AvgPx *string `json:"avgPx,omitempty"`
	// Sub order created time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	CTime *string `json:"cTime,omitempty"`
	// Margin currency  Only applicable to cross MARGIN orders in `Spot and futures mode`.
	Ccy *string `json:"ccy,omitempty"`
	// Contract value  Only applicable to `FUTURES`/`SWAP`
	CtVal *string `json:"ctVal,omitempty"`
	// Sub order fee amount
	Fee *string `json:"fee,omitempty"`
	// Sub order fee currency
	FeeCcy *string `json:"feeCcy,omitempty"`
	// Group ID
	GroupId *string `json:"groupId,omitempty"`
	// Instrument ID
	InstId *string `json:"instId,omitempty"`
	// Instrument type
	InstType *string `json:"instType,omitempty"`
	// Leverage
	Lever *string `json:"lever,omitempty"`
	// Sub order ID
	OrdId *string `json:"ordId,omitempty"`
	// Sub order type  `market`: Market order  `limit`: Limit order  `ioc`: Immediate-or-cancel order
	OrdType *string `json:"ordType,omitempty"`
	// Sub order profit and loss
	Pnl *string `json:"pnl,omitempty"`
	// Sub order position side  `net`
	PosSide *string `json:"posSide,omitempty"`
	// Sub order price
	Px *string `json:"px,omitempty"`
	// Sub order rebate amount
	Rebate *string `json:"rebate,omitempty"`
	// Sub order rebate currency
	RebateCcy *string `json:"rebateCcy,omitempty"`
	// Sub order side  `buy` `sell`
	Side *string `json:"side,omitempty"`
	// Sub order state  `canceled`  `live`  `partially_filled`  `filled`  `cancelling`
	State *string `json:"state,omitempty"`
	// Sub order quantity to buy or sell
	Sz *string `json:"sz,omitempty"`
	// Order tag
	Tag *string `json:"tag,omitempty"`
	// Sub order trade mode  Margin mode: `cross`/`isolated`  Non-Margin mode: `cash`
	TdMode *string `json:"tdMode,omitempty"`
	// Sub order updated time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	UTime *string `json:"uTime,omitempty"`
}

// NewGetTradingBotGridSubOrdersV5RespDataInner instantiates a new GetTradingBotGridSubOrdersV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTradingBotGridSubOrdersV5RespDataInner() *GetTradingBotGridSubOrdersV5RespDataInner {
	this := GetTradingBotGridSubOrdersV5RespDataInner{}
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
	var ccy string = ""
	this.Ccy = &ccy
	var ctVal string = ""
	this.CtVal = &ctVal
	var fee string = ""
	this.Fee = &fee
	var feeCcy string = ""
	this.FeeCcy = &feeCcy
	var groupId string = ""
	this.GroupId = &groupId
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var lever string = ""
	this.Lever = &lever
	var ordId string = ""
	this.OrdId = &ordId
	var ordType string = ""
	this.OrdType = &ordType
	var pnl string = ""
	this.Pnl = &pnl
	var posSide string = ""
	this.PosSide = &posSide
	var px string = ""
	this.Px = &px
	var rebate string = ""
	this.Rebate = &rebate
	var rebateCcy string = ""
	this.RebateCcy = &rebateCcy
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

// NewGetTradingBotGridSubOrdersV5RespDataInnerWithDefaults instantiates a new GetTradingBotGridSubOrdersV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTradingBotGridSubOrdersV5RespDataInnerWithDefaults() *GetTradingBotGridSubOrdersV5RespDataInner {
	this := GetTradingBotGridSubOrdersV5RespDataInner{}
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
	var ccy string = ""
	this.Ccy = &ccy
	var ctVal string = ""
	this.CtVal = &ctVal
	var fee string = ""
	this.Fee = &fee
	var feeCcy string = ""
	this.FeeCcy = &feeCcy
	var groupId string = ""
	this.GroupId = &groupId
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var lever string = ""
	this.Lever = &lever
	var ordId string = ""
	this.OrdId = &ordId
	var ordType string = ""
	this.OrdType = &ordType
	var pnl string = ""
	this.Pnl = &pnl
	var posSide string = ""
	this.PosSide = &posSide
	var px string = ""
	this.Px = &px
	var rebate string = ""
	this.Rebate = &rebate
	var rebateCcy string = ""
	this.RebateCcy = &rebateCcy
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
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetAccFillSz() string {
	if o == nil || IsNil(o.AccFillSz) {
		var ret string
		return ret
	}
	return *o.AccFillSz
}

// GetAccFillSzOk returns a tuple with the AccFillSz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetAccFillSzOk() (*string, bool) {
	if o == nil || IsNil(o.AccFillSz) {
		return nil, false
	}
	return o.AccFillSz, true
}

// HasAccFillSz returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasAccFillSz() bool {
	if o != nil && !IsNil(o.AccFillSz) {
		return true
	}

	return false
}

// SetAccFillSz gets a reference to the given string and assigns it to the AccFillSz field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetAccFillSz(v string) {
	o.AccFillSz = &v
}

// GetAlgoClOrdId returns the AlgoClOrdId field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetAlgoClOrdId() string {
	if o == nil || IsNil(o.AlgoClOrdId) {
		var ret string
		return ret
	}
	return *o.AlgoClOrdId
}

// GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetAlgoClOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoClOrdId) {
		return nil, false
	}
	return o.AlgoClOrdId, true
}

// HasAlgoClOrdId returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasAlgoClOrdId() bool {
	if o != nil && !IsNil(o.AlgoClOrdId) {
		return true
	}

	return false
}

// SetAlgoClOrdId gets a reference to the given string and assigns it to the AlgoClOrdId field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetAlgoClOrdId(v string) {
	o.AlgoClOrdId = &v
}

// GetAlgoId returns the AlgoId field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetAlgoId() string {
	if o == nil || IsNil(o.AlgoId) {
		var ret string
		return ret
	}
	return *o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetAlgoIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoId) {
		return nil, false
	}
	return o.AlgoId, true
}

// HasAlgoId returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasAlgoId() bool {
	if o != nil && !IsNil(o.AlgoId) {
		return true
	}

	return false
}

// SetAlgoId gets a reference to the given string and assigns it to the AlgoId field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetAlgoId(v string) {
	o.AlgoId = &v
}

// GetAlgoOrdType returns the AlgoOrdType field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetAlgoOrdType() string {
	if o == nil || IsNil(o.AlgoOrdType) {
		var ret string
		return ret
	}
	return *o.AlgoOrdType
}

// GetAlgoOrdTypeOk returns a tuple with the AlgoOrdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetAlgoOrdTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoOrdType) {
		return nil, false
	}
	return o.AlgoOrdType, true
}

// HasAlgoOrdType returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasAlgoOrdType() bool {
	if o != nil && !IsNil(o.AlgoOrdType) {
		return true
	}

	return false
}

// SetAlgoOrdType gets a reference to the given string and assigns it to the AlgoOrdType field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetAlgoOrdType(v string) {
	o.AlgoOrdType = &v
}

// GetAvgPx returns the AvgPx field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetAvgPx() string {
	if o == nil || IsNil(o.AvgPx) {
		var ret string
		return ret
	}
	return *o.AvgPx
}

// GetAvgPxOk returns a tuple with the AvgPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetAvgPxOk() (*string, bool) {
	if o == nil || IsNil(o.AvgPx) {
		return nil, false
	}
	return o.AvgPx, true
}

// HasAvgPx returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasAvgPx() bool {
	if o != nil && !IsNil(o.AvgPx) {
		return true
	}

	return false
}

// SetAvgPx gets a reference to the given string and assigns it to the AvgPx field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetAvgPx(v string) {
	o.AvgPx = &v
}

// GetCTime returns the CTime field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetCTime() string {
	if o == nil || IsNil(o.CTime) {
		var ret string
		return ret
	}
	return *o.CTime
}

// GetCTimeOk returns a tuple with the CTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetCTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CTime) {
		return nil, false
	}
	return o.CTime, true
}

// HasCTime returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasCTime() bool {
	if o != nil && !IsNil(o.CTime) {
		return true
	}

	return false
}

// SetCTime gets a reference to the given string and assigns it to the CTime field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetCTime(v string) {
	o.CTime = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetCtVal returns the CtVal field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetCtVal() string {
	if o == nil || IsNil(o.CtVal) {
		var ret string
		return ret
	}
	return *o.CtVal
}

// GetCtValOk returns a tuple with the CtVal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetCtValOk() (*string, bool) {
	if o == nil || IsNil(o.CtVal) {
		return nil, false
	}
	return o.CtVal, true
}

// HasCtVal returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasCtVal() bool {
	if o != nil && !IsNil(o.CtVal) {
		return true
	}

	return false
}

// SetCtVal gets a reference to the given string and assigns it to the CtVal field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetCtVal(v string) {
	o.CtVal = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetFee() string {
	if o == nil || IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetFeeOk() (*string, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetFee(v string) {
	o.Fee = &v
}

// GetFeeCcy returns the FeeCcy field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetFeeCcy() string {
	if o == nil || IsNil(o.FeeCcy) {
		var ret string
		return ret
	}
	return *o.FeeCcy
}

// GetFeeCcyOk returns a tuple with the FeeCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetFeeCcyOk() (*string, bool) {
	if o == nil || IsNil(o.FeeCcy) {
		return nil, false
	}
	return o.FeeCcy, true
}

// HasFeeCcy returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasFeeCcy() bool {
	if o != nil && !IsNil(o.FeeCcy) {
		return true
	}

	return false
}

// SetFeeCcy gets a reference to the given string and assigns it to the FeeCcy field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetFeeCcy(v string) {
	o.FeeCcy = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetGroupId(v string) {
	o.GroupId = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetInstId(v string) {
	o.InstId = &v
}

// GetInstType returns the InstType field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetInstType() string {
	if o == nil || IsNil(o.InstType) {
		var ret string
		return ret
	}
	return *o.InstType
}

// GetInstTypeOk returns a tuple with the InstType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetInstTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstType) {
		return nil, false
	}
	return o.InstType, true
}

// HasInstType returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasInstType() bool {
	if o != nil && !IsNil(o.InstType) {
		return true
	}

	return false
}

// SetInstType gets a reference to the given string and assigns it to the InstType field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetInstType(v string) {
	o.InstType = &v
}

// GetLever returns the Lever field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetLever() string {
	if o == nil || IsNil(o.Lever) {
		var ret string
		return ret
	}
	return *o.Lever
}

// GetLeverOk returns a tuple with the Lever field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetLeverOk() (*string, bool) {
	if o == nil || IsNil(o.Lever) {
		return nil, false
	}
	return o.Lever, true
}

// HasLever returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasLever() bool {
	if o != nil && !IsNil(o.Lever) {
		return true
	}

	return false
}

// SetLever gets a reference to the given string and assigns it to the Lever field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetLever(v string) {
	o.Lever = &v
}

// GetOrdId returns the OrdId field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetOrdId() string {
	if o == nil || IsNil(o.OrdId) {
		var ret string
		return ret
	}
	return *o.OrdId
}

// GetOrdIdOk returns a tuple with the OrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrdId) {
		return nil, false
	}
	return o.OrdId, true
}

// HasOrdId returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasOrdId() bool {
	if o != nil && !IsNil(o.OrdId) {
		return true
	}

	return false
}

// SetOrdId gets a reference to the given string and assigns it to the OrdId field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetOrdId(v string) {
	o.OrdId = &v
}

// GetOrdType returns the OrdType field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetOrdType() string {
	if o == nil || IsNil(o.OrdType) {
		var ret string
		return ret
	}
	return *o.OrdType
}

// GetOrdTypeOk returns a tuple with the OrdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetOrdTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OrdType) {
		return nil, false
	}
	return o.OrdType, true
}

// HasOrdType returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasOrdType() bool {
	if o != nil && !IsNil(o.OrdType) {
		return true
	}

	return false
}

// SetOrdType gets a reference to the given string and assigns it to the OrdType field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetOrdType(v string) {
	o.OrdType = &v
}

// GetPnl returns the Pnl field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetPnl() string {
	if o == nil || IsNil(o.Pnl) {
		var ret string
		return ret
	}
	return *o.Pnl
}

// GetPnlOk returns a tuple with the Pnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetPnlOk() (*string, bool) {
	if o == nil || IsNil(o.Pnl) {
		return nil, false
	}
	return o.Pnl, true
}

// HasPnl returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasPnl() bool {
	if o != nil && !IsNil(o.Pnl) {
		return true
	}

	return false
}

// SetPnl gets a reference to the given string and assigns it to the Pnl field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetPnl(v string) {
	o.Pnl = &v
}

// GetPosSide returns the PosSide field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetPosSide() string {
	if o == nil || IsNil(o.PosSide) {
		var ret string
		return ret
	}
	return *o.PosSide
}

// GetPosSideOk returns a tuple with the PosSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetPosSideOk() (*string, bool) {
	if o == nil || IsNil(o.PosSide) {
		return nil, false
	}
	return o.PosSide, true
}

// HasPosSide returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasPosSide() bool {
	if o != nil && !IsNil(o.PosSide) {
		return true
	}

	return false
}

// SetPosSide gets a reference to the given string and assigns it to the PosSide field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetPosSide(v string) {
	o.PosSide = &v
}

// GetPx returns the Px field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetPx() string {
	if o == nil || IsNil(o.Px) {
		var ret string
		return ret
	}
	return *o.Px
}

// GetPxOk returns a tuple with the Px field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetPxOk() (*string, bool) {
	if o == nil || IsNil(o.Px) {
		return nil, false
	}
	return o.Px, true
}

// HasPx returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasPx() bool {
	if o != nil && !IsNil(o.Px) {
		return true
	}

	return false
}

// SetPx gets a reference to the given string and assigns it to the Px field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetPx(v string) {
	o.Px = &v
}

// GetRebate returns the Rebate field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetRebate() string {
	if o == nil || IsNil(o.Rebate) {
		var ret string
		return ret
	}
	return *o.Rebate
}

// GetRebateOk returns a tuple with the Rebate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetRebateOk() (*string, bool) {
	if o == nil || IsNil(o.Rebate) {
		return nil, false
	}
	return o.Rebate, true
}

// HasRebate returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasRebate() bool {
	if o != nil && !IsNil(o.Rebate) {
		return true
	}

	return false
}

// SetRebate gets a reference to the given string and assigns it to the Rebate field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetRebate(v string) {
	o.Rebate = &v
}

// GetRebateCcy returns the RebateCcy field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetRebateCcy() string {
	if o == nil || IsNil(o.RebateCcy) {
		var ret string
		return ret
	}
	return *o.RebateCcy
}

// GetRebateCcyOk returns a tuple with the RebateCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetRebateCcyOk() (*string, bool) {
	if o == nil || IsNil(o.RebateCcy) {
		return nil, false
	}
	return o.RebateCcy, true
}

// HasRebateCcy returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasRebateCcy() bool {
	if o != nil && !IsNil(o.RebateCcy) {
		return true
	}

	return false
}

// SetRebateCcy gets a reference to the given string and assigns it to the RebateCcy field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetRebateCcy(v string) {
	o.RebateCcy = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetSide(v string) {
	o.Side = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetState(v string) {
	o.State = &v
}

// GetSz returns the Sz field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetSz() string {
	if o == nil || IsNil(o.Sz) {
		var ret string
		return ret
	}
	return *o.Sz
}

// GetSzOk returns a tuple with the Sz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetSzOk() (*string, bool) {
	if o == nil || IsNil(o.Sz) {
		return nil, false
	}
	return o.Sz, true
}

// HasSz returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasSz() bool {
	if o != nil && !IsNil(o.Sz) {
		return true
	}

	return false
}

// SetSz gets a reference to the given string and assigns it to the Sz field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetSz(v string) {
	o.Sz = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetTag(v string) {
	o.Tag = &v
}

// GetTdMode returns the TdMode field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetTdMode() string {
	if o == nil || IsNil(o.TdMode) {
		var ret string
		return ret
	}
	return *o.TdMode
}

// GetTdModeOk returns a tuple with the TdMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetTdModeOk() (*string, bool) {
	if o == nil || IsNil(o.TdMode) {
		return nil, false
	}
	return o.TdMode, true
}

// HasTdMode returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasTdMode() bool {
	if o != nil && !IsNil(o.TdMode) {
		return true
	}

	return false
}

// SetTdMode gets a reference to the given string and assigns it to the TdMode field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetTdMode(v string) {
	o.TdMode = &v
}

// GetUTime returns the UTime field value if set, zero value otherwise.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetUTime() string {
	if o == nil || IsNil(o.UTime) {
		var ret string
		return ret
	}
	return *o.UTime
}

// GetUTimeOk returns a tuple with the UTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) GetUTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UTime) {
		return nil, false
	}
	return o.UTime, true
}

// HasUTime returns a boolean if a field has been set.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) HasUTime() bool {
	if o != nil && !IsNil(o.UTime) {
		return true
	}

	return false
}

// SetUTime gets a reference to the given string and assigns it to the UTime field.
func (o *GetTradingBotGridSubOrdersV5RespDataInner) SetUTime(v string) {
	o.UTime = &v
}

func (o GetTradingBotGridSubOrdersV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTradingBotGridSubOrdersV5RespDataInner) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.CtVal) {
		toSerialize["ctVal"] = o.CtVal
	}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.FeeCcy) {
		toSerialize["feeCcy"] = o.FeeCcy
	}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
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
	if !IsNil(o.OrdId) {
		toSerialize["ordId"] = o.OrdId
	}
	if !IsNil(o.OrdType) {
		toSerialize["ordType"] = o.OrdType
	}
	if !IsNil(o.Pnl) {
		toSerialize["pnl"] = o.Pnl
	}
	if !IsNil(o.PosSide) {
		toSerialize["posSide"] = o.PosSide
	}
	if !IsNil(o.Px) {
		toSerialize["px"] = o.Px
	}
	if !IsNil(o.Rebate) {
		toSerialize["rebate"] = o.Rebate
	}
	if !IsNil(o.RebateCcy) {
		toSerialize["rebateCcy"] = o.RebateCcy
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

type NullableGetTradingBotGridSubOrdersV5RespDataInner struct {
	value *GetTradingBotGridSubOrdersV5RespDataInner
	isSet bool
}

func (v NullableGetTradingBotGridSubOrdersV5RespDataInner) Get() *GetTradingBotGridSubOrdersV5RespDataInner {
	return v.value
}

func (v *NullableGetTradingBotGridSubOrdersV5RespDataInner) Set(val *GetTradingBotGridSubOrdersV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTradingBotGridSubOrdersV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTradingBotGridSubOrdersV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTradingBotGridSubOrdersV5RespDataInner(val *GetTradingBotGridSubOrdersV5RespDataInner) *NullableGetTradingBotGridSubOrdersV5RespDataInner {
	return &NullableGetTradingBotGridSubOrdersV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetTradingBotGridSubOrdersV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTradingBotGridSubOrdersV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


