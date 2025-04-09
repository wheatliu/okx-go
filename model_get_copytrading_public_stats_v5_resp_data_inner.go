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

// checks if the GetCopytradingPublicStatsV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCopytradingPublicStatsV5RespDataInner{}

// GetCopytradingPublicStatsV5RespDataInner struct for GetCopytradingPublicStatsV5RespDataInner
type GetCopytradingPublicStatsV5RespDataInner struct {
	// Average lead position notional (USDT)
	AvgSubPosNotional *string `json:"avgSubPosNotional,omitempty"`
	// Margin currency
	Ccy *string `json:"ccy,omitempty"`
	// Current copy trader pnl (USDT)
	CurCopyTraderPnl *string `json:"curCopyTraderPnl,omitempty"`
	// Investment amount (USDT)
	InvestAmt *string `json:"investAmt,omitempty"`
	// Loss days
	LossDays *string `json:"lossDays,omitempty"`
	// Profit days
	ProfitDays *string `json:"profitDays,omitempty"`
	// Win ratio
	WinRatio *string `json:"winRatio,omitempty"`
}

// NewGetCopytradingPublicStatsV5RespDataInner instantiates a new GetCopytradingPublicStatsV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCopytradingPublicStatsV5RespDataInner() *GetCopytradingPublicStatsV5RespDataInner {
	this := GetCopytradingPublicStatsV5RespDataInner{}
	var avgSubPosNotional string = ""
	this.AvgSubPosNotional = &avgSubPosNotional
	var ccy string = ""
	this.Ccy = &ccy
	var curCopyTraderPnl string = ""
	this.CurCopyTraderPnl = &curCopyTraderPnl
	var investAmt string = ""
	this.InvestAmt = &investAmt
	var lossDays string = ""
	this.LossDays = &lossDays
	var profitDays string = ""
	this.ProfitDays = &profitDays
	var winRatio string = ""
	this.WinRatio = &winRatio
	return &this
}

// NewGetCopytradingPublicStatsV5RespDataInnerWithDefaults instantiates a new GetCopytradingPublicStatsV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCopytradingPublicStatsV5RespDataInnerWithDefaults() *GetCopytradingPublicStatsV5RespDataInner {
	this := GetCopytradingPublicStatsV5RespDataInner{}
	var avgSubPosNotional string = ""
	this.AvgSubPosNotional = &avgSubPosNotional
	var ccy string = ""
	this.Ccy = &ccy
	var curCopyTraderPnl string = ""
	this.CurCopyTraderPnl = &curCopyTraderPnl
	var investAmt string = ""
	this.InvestAmt = &investAmt
	var lossDays string = ""
	this.LossDays = &lossDays
	var profitDays string = ""
	this.ProfitDays = &profitDays
	var winRatio string = ""
	this.WinRatio = &winRatio
	return &this
}

// GetAvgSubPosNotional returns the AvgSubPosNotional field value if set, zero value otherwise.
func (o *GetCopytradingPublicStatsV5RespDataInner) GetAvgSubPosNotional() string {
	if o == nil || IsNil(o.AvgSubPosNotional) {
		var ret string
		return ret
	}
	return *o.AvgSubPosNotional
}

// GetAvgSubPosNotionalOk returns a tuple with the AvgSubPosNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicStatsV5RespDataInner) GetAvgSubPosNotionalOk() (*string, bool) {
	if o == nil || IsNil(o.AvgSubPosNotional) {
		return nil, false
	}
	return o.AvgSubPosNotional, true
}

// HasAvgSubPosNotional returns a boolean if a field has been set.
func (o *GetCopytradingPublicStatsV5RespDataInner) HasAvgSubPosNotional() bool {
	if o != nil && !IsNil(o.AvgSubPosNotional) {
		return true
	}

	return false
}

// SetAvgSubPosNotional gets a reference to the given string and assigns it to the AvgSubPosNotional field.
func (o *GetCopytradingPublicStatsV5RespDataInner) SetAvgSubPosNotional(v string) {
	o.AvgSubPosNotional = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetCopytradingPublicStatsV5RespDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicStatsV5RespDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetCopytradingPublicStatsV5RespDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetCopytradingPublicStatsV5RespDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetCurCopyTraderPnl returns the CurCopyTraderPnl field value if set, zero value otherwise.
func (o *GetCopytradingPublicStatsV5RespDataInner) GetCurCopyTraderPnl() string {
	if o == nil || IsNil(o.CurCopyTraderPnl) {
		var ret string
		return ret
	}
	return *o.CurCopyTraderPnl
}

// GetCurCopyTraderPnlOk returns a tuple with the CurCopyTraderPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicStatsV5RespDataInner) GetCurCopyTraderPnlOk() (*string, bool) {
	if o == nil || IsNil(o.CurCopyTraderPnl) {
		return nil, false
	}
	return o.CurCopyTraderPnl, true
}

// HasCurCopyTraderPnl returns a boolean if a field has been set.
func (o *GetCopytradingPublicStatsV5RespDataInner) HasCurCopyTraderPnl() bool {
	if o != nil && !IsNil(o.CurCopyTraderPnl) {
		return true
	}

	return false
}

// SetCurCopyTraderPnl gets a reference to the given string and assigns it to the CurCopyTraderPnl field.
func (o *GetCopytradingPublicStatsV5RespDataInner) SetCurCopyTraderPnl(v string) {
	o.CurCopyTraderPnl = &v
}

// GetInvestAmt returns the InvestAmt field value if set, zero value otherwise.
func (o *GetCopytradingPublicStatsV5RespDataInner) GetInvestAmt() string {
	if o == nil || IsNil(o.InvestAmt) {
		var ret string
		return ret
	}
	return *o.InvestAmt
}

// GetInvestAmtOk returns a tuple with the InvestAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicStatsV5RespDataInner) GetInvestAmtOk() (*string, bool) {
	if o == nil || IsNil(o.InvestAmt) {
		return nil, false
	}
	return o.InvestAmt, true
}

// HasInvestAmt returns a boolean if a field has been set.
func (o *GetCopytradingPublicStatsV5RespDataInner) HasInvestAmt() bool {
	if o != nil && !IsNil(o.InvestAmt) {
		return true
	}

	return false
}

// SetInvestAmt gets a reference to the given string and assigns it to the InvestAmt field.
func (o *GetCopytradingPublicStatsV5RespDataInner) SetInvestAmt(v string) {
	o.InvestAmt = &v
}

// GetLossDays returns the LossDays field value if set, zero value otherwise.
func (o *GetCopytradingPublicStatsV5RespDataInner) GetLossDays() string {
	if o == nil || IsNil(o.LossDays) {
		var ret string
		return ret
	}
	return *o.LossDays
}

// GetLossDaysOk returns a tuple with the LossDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicStatsV5RespDataInner) GetLossDaysOk() (*string, bool) {
	if o == nil || IsNil(o.LossDays) {
		return nil, false
	}
	return o.LossDays, true
}

// HasLossDays returns a boolean if a field has been set.
func (o *GetCopytradingPublicStatsV5RespDataInner) HasLossDays() bool {
	if o != nil && !IsNil(o.LossDays) {
		return true
	}

	return false
}

// SetLossDays gets a reference to the given string and assigns it to the LossDays field.
func (o *GetCopytradingPublicStatsV5RespDataInner) SetLossDays(v string) {
	o.LossDays = &v
}

// GetProfitDays returns the ProfitDays field value if set, zero value otherwise.
func (o *GetCopytradingPublicStatsV5RespDataInner) GetProfitDays() string {
	if o == nil || IsNil(o.ProfitDays) {
		var ret string
		return ret
	}
	return *o.ProfitDays
}

// GetProfitDaysOk returns a tuple with the ProfitDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicStatsV5RespDataInner) GetProfitDaysOk() (*string, bool) {
	if o == nil || IsNil(o.ProfitDays) {
		return nil, false
	}
	return o.ProfitDays, true
}

// HasProfitDays returns a boolean if a field has been set.
func (o *GetCopytradingPublicStatsV5RespDataInner) HasProfitDays() bool {
	if o != nil && !IsNil(o.ProfitDays) {
		return true
	}

	return false
}

// SetProfitDays gets a reference to the given string and assigns it to the ProfitDays field.
func (o *GetCopytradingPublicStatsV5RespDataInner) SetProfitDays(v string) {
	o.ProfitDays = &v
}

// GetWinRatio returns the WinRatio field value if set, zero value otherwise.
func (o *GetCopytradingPublicStatsV5RespDataInner) GetWinRatio() string {
	if o == nil || IsNil(o.WinRatio) {
		var ret string
		return ret
	}
	return *o.WinRatio
}

// GetWinRatioOk returns a tuple with the WinRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicStatsV5RespDataInner) GetWinRatioOk() (*string, bool) {
	if o == nil || IsNil(o.WinRatio) {
		return nil, false
	}
	return o.WinRatio, true
}

// HasWinRatio returns a boolean if a field has been set.
func (o *GetCopytradingPublicStatsV5RespDataInner) HasWinRatio() bool {
	if o != nil && !IsNil(o.WinRatio) {
		return true
	}

	return false
}

// SetWinRatio gets a reference to the given string and assigns it to the WinRatio field.
func (o *GetCopytradingPublicStatsV5RespDataInner) SetWinRatio(v string) {
	o.WinRatio = &v
}

func (o GetCopytradingPublicStatsV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCopytradingPublicStatsV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvgSubPosNotional) {
		toSerialize["avgSubPosNotional"] = o.AvgSubPosNotional
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.CurCopyTraderPnl) {
		toSerialize["curCopyTraderPnl"] = o.CurCopyTraderPnl
	}
	if !IsNil(o.InvestAmt) {
		toSerialize["investAmt"] = o.InvestAmt
	}
	if !IsNil(o.LossDays) {
		toSerialize["lossDays"] = o.LossDays
	}
	if !IsNil(o.ProfitDays) {
		toSerialize["profitDays"] = o.ProfitDays
	}
	if !IsNil(o.WinRatio) {
		toSerialize["winRatio"] = o.WinRatio
	}
	return toSerialize, nil
}

type NullableGetCopytradingPublicStatsV5RespDataInner struct {
	value *GetCopytradingPublicStatsV5RespDataInner
	isSet bool
}

func (v NullableGetCopytradingPublicStatsV5RespDataInner) Get() *GetCopytradingPublicStatsV5RespDataInner {
	return v.value
}

func (v *NullableGetCopytradingPublicStatsV5RespDataInner) Set(val *GetCopytradingPublicStatsV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCopytradingPublicStatsV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCopytradingPublicStatsV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCopytradingPublicStatsV5RespDataInner(val *GetCopytradingPublicStatsV5RespDataInner) *NullableGetCopytradingPublicStatsV5RespDataInner {
	return &NullableGetCopytradingPublicStatsV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetCopytradingPublicStatsV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCopytradingPublicStatsV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


