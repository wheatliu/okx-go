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

// checks if the GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner{}

// GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner struct for GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner
type GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner struct {
	// Accumulated number of copy traders
	AccCopyTraderNum *string `json:"accCopyTraderNum,omitempty"`
	// assets under management
	Aum *string `json:"aum,omitempty"`
	// Begin time of pnl ratio on that day
	BeginTs *string `json:"beginTs,omitempty"`
	// Margin currency
	Ccy *string `json:"ccy,omitempty"`
	// Current copy state   `0`: non-copy, `1`: copy
	CopyState *string `json:"copyState,omitempty"`
	// Current number of copy traders
	CopyTraderNum *string `json:"copyTraderNum,omitempty"`
	// Lead days
	LeadDays *string `json:"leadDays,omitempty"`
	// Maximum number of copy traders
	MaxCopyTraderNum *string `json:"maxCopyTraderNum,omitempty"`
	// Nick name
	NickName *string `json:"nickName,omitempty"`
	// Pnl (in USDT) of last 90 days.
	Pnl *string `json:"pnl,omitempty"`
	// Pnl ratio on that day
	PnlRatio *string `json:"pnlRatio,omitempty"`
	// Pnl ratios
	PnlRatios []GetCopytradingPublicLeadTradersV5RespDataInnerRanksInnerPnlRatiosInner `json:"pnlRatios,omitempty"`
	// Portrait link
	PortLink *string `json:"portLink,omitempty"`
	// Contract list which lead trader is leading
	TraderInsts []string `json:"traderInsts,omitempty"`
	// Lead trader unique code
	UniqueCode *string `json:"uniqueCode,omitempty"`
	// Win ratio, 0.1 represents 10%
	WinRatio *string `json:"winRatio,omitempty"`
}

// NewGetCopytradingPublicLeadTradersV5RespDataInnerRanksInner instantiates a new GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCopytradingPublicLeadTradersV5RespDataInnerRanksInner() *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner {
	this := GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner{}
	var accCopyTraderNum string = ""
	this.AccCopyTraderNum = &accCopyTraderNum
	var aum string = ""
	this.Aum = &aum
	var beginTs string = ""
	this.BeginTs = &beginTs
	var ccy string = ""
	this.Ccy = &ccy
	var copyState string = ""
	this.CopyState = &copyState
	var copyTraderNum string = ""
	this.CopyTraderNum = &copyTraderNum
	var leadDays string = ""
	this.LeadDays = &leadDays
	var maxCopyTraderNum string = ""
	this.MaxCopyTraderNum = &maxCopyTraderNum
	var nickName string = ""
	this.NickName = &nickName
	var pnl string = ""
	this.Pnl = &pnl
	var pnlRatio string = ""
	this.PnlRatio = &pnlRatio
	var portLink string = ""
	this.PortLink = &portLink
	var uniqueCode string = ""
	this.UniqueCode = &uniqueCode
	var winRatio string = ""
	this.WinRatio = &winRatio
	return &this
}

// NewGetCopytradingPublicLeadTradersV5RespDataInnerRanksInnerWithDefaults instantiates a new GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCopytradingPublicLeadTradersV5RespDataInnerRanksInnerWithDefaults() *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner {
	this := GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner{}
	var accCopyTraderNum string = ""
	this.AccCopyTraderNum = &accCopyTraderNum
	var aum string = ""
	this.Aum = &aum
	var beginTs string = ""
	this.BeginTs = &beginTs
	var ccy string = ""
	this.Ccy = &ccy
	var copyState string = ""
	this.CopyState = &copyState
	var copyTraderNum string = ""
	this.CopyTraderNum = &copyTraderNum
	var leadDays string = ""
	this.LeadDays = &leadDays
	var maxCopyTraderNum string = ""
	this.MaxCopyTraderNum = &maxCopyTraderNum
	var nickName string = ""
	this.NickName = &nickName
	var pnl string = ""
	this.Pnl = &pnl
	var pnlRatio string = ""
	this.PnlRatio = &pnlRatio
	var portLink string = ""
	this.PortLink = &portLink
	var uniqueCode string = ""
	this.UniqueCode = &uniqueCode
	var winRatio string = ""
	this.WinRatio = &winRatio
	return &this
}

// GetAccCopyTraderNum returns the AccCopyTraderNum field value if set, zero value otherwise.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetAccCopyTraderNum() string {
	if o == nil || IsNil(o.AccCopyTraderNum) {
		var ret string
		return ret
	}
	return *o.AccCopyTraderNum
}

// GetAccCopyTraderNumOk returns a tuple with the AccCopyTraderNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetAccCopyTraderNumOk() (*string, bool) {
	if o == nil || IsNil(o.AccCopyTraderNum) {
		return nil, false
	}
	return o.AccCopyTraderNum, true
}

// HasAccCopyTraderNum returns a boolean if a field has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasAccCopyTraderNum() bool {
	if o != nil && !IsNil(o.AccCopyTraderNum) {
		return true
	}

	return false
}

// SetAccCopyTraderNum gets a reference to the given string and assigns it to the AccCopyTraderNum field.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetAccCopyTraderNum(v string) {
	o.AccCopyTraderNum = &v
}

// GetAum returns the Aum field value if set, zero value otherwise.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetAum() string {
	if o == nil || IsNil(o.Aum) {
		var ret string
		return ret
	}
	return *o.Aum
}

// GetAumOk returns a tuple with the Aum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetAumOk() (*string, bool) {
	if o == nil || IsNil(o.Aum) {
		return nil, false
	}
	return o.Aum, true
}

// HasAum returns a boolean if a field has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasAum() bool {
	if o != nil && !IsNil(o.Aum) {
		return true
	}

	return false
}

// SetAum gets a reference to the given string and assigns it to the Aum field.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetAum(v string) {
	o.Aum = &v
}

// GetBeginTs returns the BeginTs field value if set, zero value otherwise.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetBeginTs() string {
	if o == nil || IsNil(o.BeginTs) {
		var ret string
		return ret
	}
	return *o.BeginTs
}

// GetBeginTsOk returns a tuple with the BeginTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetBeginTsOk() (*string, bool) {
	if o == nil || IsNil(o.BeginTs) {
		return nil, false
	}
	return o.BeginTs, true
}

// HasBeginTs returns a boolean if a field has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasBeginTs() bool {
	if o != nil && !IsNil(o.BeginTs) {
		return true
	}

	return false
}

// SetBeginTs gets a reference to the given string and assigns it to the BeginTs field.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetBeginTs(v string) {
	o.BeginTs = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetCopyState returns the CopyState field value if set, zero value otherwise.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetCopyState() string {
	if o == nil || IsNil(o.CopyState) {
		var ret string
		return ret
	}
	return *o.CopyState
}

// GetCopyStateOk returns a tuple with the CopyState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetCopyStateOk() (*string, bool) {
	if o == nil || IsNil(o.CopyState) {
		return nil, false
	}
	return o.CopyState, true
}

// HasCopyState returns a boolean if a field has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasCopyState() bool {
	if o != nil && !IsNil(o.CopyState) {
		return true
	}

	return false
}

// SetCopyState gets a reference to the given string and assigns it to the CopyState field.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetCopyState(v string) {
	o.CopyState = &v
}

// GetCopyTraderNum returns the CopyTraderNum field value if set, zero value otherwise.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetCopyTraderNum() string {
	if o == nil || IsNil(o.CopyTraderNum) {
		var ret string
		return ret
	}
	return *o.CopyTraderNum
}

// GetCopyTraderNumOk returns a tuple with the CopyTraderNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetCopyTraderNumOk() (*string, bool) {
	if o == nil || IsNil(o.CopyTraderNum) {
		return nil, false
	}
	return o.CopyTraderNum, true
}

// HasCopyTraderNum returns a boolean if a field has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasCopyTraderNum() bool {
	if o != nil && !IsNil(o.CopyTraderNum) {
		return true
	}

	return false
}

// SetCopyTraderNum gets a reference to the given string and assigns it to the CopyTraderNum field.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetCopyTraderNum(v string) {
	o.CopyTraderNum = &v
}

// GetLeadDays returns the LeadDays field value if set, zero value otherwise.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetLeadDays() string {
	if o == nil || IsNil(o.LeadDays) {
		var ret string
		return ret
	}
	return *o.LeadDays
}

// GetLeadDaysOk returns a tuple with the LeadDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetLeadDaysOk() (*string, bool) {
	if o == nil || IsNil(o.LeadDays) {
		return nil, false
	}
	return o.LeadDays, true
}

// HasLeadDays returns a boolean if a field has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasLeadDays() bool {
	if o != nil && !IsNil(o.LeadDays) {
		return true
	}

	return false
}

// SetLeadDays gets a reference to the given string and assigns it to the LeadDays field.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetLeadDays(v string) {
	o.LeadDays = &v
}

// GetMaxCopyTraderNum returns the MaxCopyTraderNum field value if set, zero value otherwise.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetMaxCopyTraderNum() string {
	if o == nil || IsNil(o.MaxCopyTraderNum) {
		var ret string
		return ret
	}
	return *o.MaxCopyTraderNum
}

// GetMaxCopyTraderNumOk returns a tuple with the MaxCopyTraderNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetMaxCopyTraderNumOk() (*string, bool) {
	if o == nil || IsNil(o.MaxCopyTraderNum) {
		return nil, false
	}
	return o.MaxCopyTraderNum, true
}

// HasMaxCopyTraderNum returns a boolean if a field has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasMaxCopyTraderNum() bool {
	if o != nil && !IsNil(o.MaxCopyTraderNum) {
		return true
	}

	return false
}

// SetMaxCopyTraderNum gets a reference to the given string and assigns it to the MaxCopyTraderNum field.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetMaxCopyTraderNum(v string) {
	o.MaxCopyTraderNum = &v
}

// GetNickName returns the NickName field value if set, zero value otherwise.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetNickName() string {
	if o == nil || IsNil(o.NickName) {
		var ret string
		return ret
	}
	return *o.NickName
}

// GetNickNameOk returns a tuple with the NickName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetNickNameOk() (*string, bool) {
	if o == nil || IsNil(o.NickName) {
		return nil, false
	}
	return o.NickName, true
}

// HasNickName returns a boolean if a field has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasNickName() bool {
	if o != nil && !IsNil(o.NickName) {
		return true
	}

	return false
}

// SetNickName gets a reference to the given string and assigns it to the NickName field.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetNickName(v string) {
	o.NickName = &v
}

// GetPnl returns the Pnl field value if set, zero value otherwise.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetPnl() string {
	if o == nil || IsNil(o.Pnl) {
		var ret string
		return ret
	}
	return *o.Pnl
}

// GetPnlOk returns a tuple with the Pnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetPnlOk() (*string, bool) {
	if o == nil || IsNil(o.Pnl) {
		return nil, false
	}
	return o.Pnl, true
}

// HasPnl returns a boolean if a field has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasPnl() bool {
	if o != nil && !IsNil(o.Pnl) {
		return true
	}

	return false
}

// SetPnl gets a reference to the given string and assigns it to the Pnl field.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetPnl(v string) {
	o.Pnl = &v
}

// GetPnlRatio returns the PnlRatio field value if set, zero value otherwise.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetPnlRatio() string {
	if o == nil || IsNil(o.PnlRatio) {
		var ret string
		return ret
	}
	return *o.PnlRatio
}

// GetPnlRatioOk returns a tuple with the PnlRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetPnlRatioOk() (*string, bool) {
	if o == nil || IsNil(o.PnlRatio) {
		return nil, false
	}
	return o.PnlRatio, true
}

// HasPnlRatio returns a boolean if a field has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasPnlRatio() bool {
	if o != nil && !IsNil(o.PnlRatio) {
		return true
	}

	return false
}

// SetPnlRatio gets a reference to the given string and assigns it to the PnlRatio field.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetPnlRatio(v string) {
	o.PnlRatio = &v
}

// GetPnlRatios returns the PnlRatios field value if set, zero value otherwise.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetPnlRatios() []GetCopytradingPublicLeadTradersV5RespDataInnerRanksInnerPnlRatiosInner {
	if o == nil || IsNil(o.PnlRatios) {
		var ret []GetCopytradingPublicLeadTradersV5RespDataInnerRanksInnerPnlRatiosInner
		return ret
	}
	return o.PnlRatios
}

// GetPnlRatiosOk returns a tuple with the PnlRatios field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetPnlRatiosOk() ([]GetCopytradingPublicLeadTradersV5RespDataInnerRanksInnerPnlRatiosInner, bool) {
	if o == nil || IsNil(o.PnlRatios) {
		return nil, false
	}
	return o.PnlRatios, true
}

// HasPnlRatios returns a boolean if a field has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasPnlRatios() bool {
	if o != nil && !IsNil(o.PnlRatios) {
		return true
	}

	return false
}

// SetPnlRatios gets a reference to the given []GetCopytradingPublicLeadTradersV5RespDataInnerRanksInnerPnlRatiosInner and assigns it to the PnlRatios field.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetPnlRatios(v []GetCopytradingPublicLeadTradersV5RespDataInnerRanksInnerPnlRatiosInner) {
	o.PnlRatios = v
}

// GetPortLink returns the PortLink field value if set, zero value otherwise.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetPortLink() string {
	if o == nil || IsNil(o.PortLink) {
		var ret string
		return ret
	}
	return *o.PortLink
}

// GetPortLinkOk returns a tuple with the PortLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetPortLinkOk() (*string, bool) {
	if o == nil || IsNil(o.PortLink) {
		return nil, false
	}
	return o.PortLink, true
}

// HasPortLink returns a boolean if a field has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasPortLink() bool {
	if o != nil && !IsNil(o.PortLink) {
		return true
	}

	return false
}

// SetPortLink gets a reference to the given string and assigns it to the PortLink field.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetPortLink(v string) {
	o.PortLink = &v
}

// GetTraderInsts returns the TraderInsts field value if set, zero value otherwise.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetTraderInsts() []string {
	if o == nil || IsNil(o.TraderInsts) {
		var ret []string
		return ret
	}
	return o.TraderInsts
}

// GetTraderInstsOk returns a tuple with the TraderInsts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetTraderInstsOk() ([]string, bool) {
	if o == nil || IsNil(o.TraderInsts) {
		return nil, false
	}
	return o.TraderInsts, true
}

// HasTraderInsts returns a boolean if a field has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasTraderInsts() bool {
	if o != nil && !IsNil(o.TraderInsts) {
		return true
	}

	return false
}

// SetTraderInsts gets a reference to the given []string and assigns it to the TraderInsts field.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetTraderInsts(v []string) {
	o.TraderInsts = v
}

// GetUniqueCode returns the UniqueCode field value if set, zero value otherwise.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetUniqueCode() string {
	if o == nil || IsNil(o.UniqueCode) {
		var ret string
		return ret
	}
	return *o.UniqueCode
}

// GetUniqueCodeOk returns a tuple with the UniqueCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetUniqueCodeOk() (*string, bool) {
	if o == nil || IsNil(o.UniqueCode) {
		return nil, false
	}
	return o.UniqueCode, true
}

// HasUniqueCode returns a boolean if a field has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasUniqueCode() bool {
	if o != nil && !IsNil(o.UniqueCode) {
		return true
	}

	return false
}

// SetUniqueCode gets a reference to the given string and assigns it to the UniqueCode field.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetUniqueCode(v string) {
	o.UniqueCode = &v
}

// GetWinRatio returns the WinRatio field value if set, zero value otherwise.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetWinRatio() string {
	if o == nil || IsNil(o.WinRatio) {
		var ret string
		return ret
	}
	return *o.WinRatio
}

// GetWinRatioOk returns a tuple with the WinRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetWinRatioOk() (*string, bool) {
	if o == nil || IsNil(o.WinRatio) {
		return nil, false
	}
	return o.WinRatio, true
}

// HasWinRatio returns a boolean if a field has been set.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasWinRatio() bool {
	if o != nil && !IsNil(o.WinRatio) {
		return true
	}

	return false
}

// SetWinRatio gets a reference to the given string and assigns it to the WinRatio field.
func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetWinRatio(v string) {
	o.WinRatio = &v
}

func (o GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccCopyTraderNum) {
		toSerialize["accCopyTraderNum"] = o.AccCopyTraderNum
	}
	if !IsNil(o.Aum) {
		toSerialize["aum"] = o.Aum
	}
	if !IsNil(o.BeginTs) {
		toSerialize["beginTs"] = o.BeginTs
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.CopyState) {
		toSerialize["copyState"] = o.CopyState
	}
	if !IsNil(o.CopyTraderNum) {
		toSerialize["copyTraderNum"] = o.CopyTraderNum
	}
	if !IsNil(o.LeadDays) {
		toSerialize["leadDays"] = o.LeadDays
	}
	if !IsNil(o.MaxCopyTraderNum) {
		toSerialize["maxCopyTraderNum"] = o.MaxCopyTraderNum
	}
	if !IsNil(o.NickName) {
		toSerialize["nickName"] = o.NickName
	}
	if !IsNil(o.Pnl) {
		toSerialize["pnl"] = o.Pnl
	}
	if !IsNil(o.PnlRatio) {
		toSerialize["pnlRatio"] = o.PnlRatio
	}
	if !IsNil(o.PnlRatios) {
		toSerialize["pnlRatios"] = o.PnlRatios
	}
	if !IsNil(o.PortLink) {
		toSerialize["portLink"] = o.PortLink
	}
	if !IsNil(o.TraderInsts) {
		toSerialize["traderInsts"] = o.TraderInsts
	}
	if !IsNil(o.UniqueCode) {
		toSerialize["uniqueCode"] = o.UniqueCode
	}
	if !IsNil(o.WinRatio) {
		toSerialize["winRatio"] = o.WinRatio
	}
	return toSerialize, nil
}

type NullableGetCopytradingPublicLeadTradersV5RespDataInnerRanksInner struct {
	value *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner
	isSet bool
}

func (v NullableGetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) Get() *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner {
	return v.value
}

func (v *NullableGetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) Set(val *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCopytradingPublicLeadTradersV5RespDataInnerRanksInner(val *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) *NullableGetCopytradingPublicLeadTradersV5RespDataInnerRanksInner {
	return &NullableGetCopytradingPublicLeadTradersV5RespDataInnerRanksInner{value: val, isSet: true}
}

func (v NullableGetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


