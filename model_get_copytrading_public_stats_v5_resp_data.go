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

// checks if the GetCopytradingPublicStatsV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCopytradingPublicStatsV5RespData{}

// GetCopytradingPublicStatsV5RespData struct for GetCopytradingPublicStatsV5RespData
type GetCopytradingPublicStatsV5RespData struct {
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

// NewGetCopytradingPublicStatsV5RespData instantiates a new GetCopytradingPublicStatsV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCopytradingPublicStatsV5RespData() *GetCopytradingPublicStatsV5RespData {
	this := GetCopytradingPublicStatsV5RespData{}
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

// NewGetCopytradingPublicStatsV5RespDataWithDefaults instantiates a new GetCopytradingPublicStatsV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCopytradingPublicStatsV5RespDataWithDefaults() *GetCopytradingPublicStatsV5RespData {
	this := GetCopytradingPublicStatsV5RespData{}
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
func (o *GetCopytradingPublicStatsV5RespData) GetAvgSubPosNotional() string {
	if o == nil || IsNil(o.AvgSubPosNotional) {
		var ret string
		return ret
	}
	return *o.AvgSubPosNotional
}

// GetAvgSubPosNotionalOk returns a tuple with the AvgSubPosNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicStatsV5RespData) GetAvgSubPosNotionalOk() (*string, bool) {
	if o == nil || IsNil(o.AvgSubPosNotional) {
		return nil, false
	}
	return o.AvgSubPosNotional, true
}

// HasAvgSubPosNotional returns a boolean if a field has been set.
func (o *GetCopytradingPublicStatsV5RespData) HasAvgSubPosNotional() bool {
	if o != nil && !IsNil(o.AvgSubPosNotional) {
		return true
	}

	return false
}

// SetAvgSubPosNotional gets a reference to the given string and assigns it to the AvgSubPosNotional field.
func (o *GetCopytradingPublicStatsV5RespData) SetAvgSubPosNotional(v string) {
	o.AvgSubPosNotional = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetCopytradingPublicStatsV5RespData) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicStatsV5RespData) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetCopytradingPublicStatsV5RespData) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetCopytradingPublicStatsV5RespData) SetCcy(v string) {
	o.Ccy = &v
}

// GetCurCopyTraderPnl returns the CurCopyTraderPnl field value if set, zero value otherwise.
func (o *GetCopytradingPublicStatsV5RespData) GetCurCopyTraderPnl() string {
	if o == nil || IsNil(o.CurCopyTraderPnl) {
		var ret string
		return ret
	}
	return *o.CurCopyTraderPnl
}

// GetCurCopyTraderPnlOk returns a tuple with the CurCopyTraderPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicStatsV5RespData) GetCurCopyTraderPnlOk() (*string, bool) {
	if o == nil || IsNil(o.CurCopyTraderPnl) {
		return nil, false
	}
	return o.CurCopyTraderPnl, true
}

// HasCurCopyTraderPnl returns a boolean if a field has been set.
func (o *GetCopytradingPublicStatsV5RespData) HasCurCopyTraderPnl() bool {
	if o != nil && !IsNil(o.CurCopyTraderPnl) {
		return true
	}

	return false
}

// SetCurCopyTraderPnl gets a reference to the given string and assigns it to the CurCopyTraderPnl field.
func (o *GetCopytradingPublicStatsV5RespData) SetCurCopyTraderPnl(v string) {
	o.CurCopyTraderPnl = &v
}

// GetInvestAmt returns the InvestAmt field value if set, zero value otherwise.
func (o *GetCopytradingPublicStatsV5RespData) GetInvestAmt() string {
	if o == nil || IsNil(o.InvestAmt) {
		var ret string
		return ret
	}
	return *o.InvestAmt
}

// GetInvestAmtOk returns a tuple with the InvestAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicStatsV5RespData) GetInvestAmtOk() (*string, bool) {
	if o == nil || IsNil(o.InvestAmt) {
		return nil, false
	}
	return o.InvestAmt, true
}

// HasInvestAmt returns a boolean if a field has been set.
func (o *GetCopytradingPublicStatsV5RespData) HasInvestAmt() bool {
	if o != nil && !IsNil(o.InvestAmt) {
		return true
	}

	return false
}

// SetInvestAmt gets a reference to the given string and assigns it to the InvestAmt field.
func (o *GetCopytradingPublicStatsV5RespData) SetInvestAmt(v string) {
	o.InvestAmt = &v
}

// GetLossDays returns the LossDays field value if set, zero value otherwise.
func (o *GetCopytradingPublicStatsV5RespData) GetLossDays() string {
	if o == nil || IsNil(o.LossDays) {
		var ret string
		return ret
	}
	return *o.LossDays
}

// GetLossDaysOk returns a tuple with the LossDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicStatsV5RespData) GetLossDaysOk() (*string, bool) {
	if o == nil || IsNil(o.LossDays) {
		return nil, false
	}
	return o.LossDays, true
}

// HasLossDays returns a boolean if a field has been set.
func (o *GetCopytradingPublicStatsV5RespData) HasLossDays() bool {
	if o != nil && !IsNil(o.LossDays) {
		return true
	}

	return false
}

// SetLossDays gets a reference to the given string and assigns it to the LossDays field.
func (o *GetCopytradingPublicStatsV5RespData) SetLossDays(v string) {
	o.LossDays = &v
}

// GetProfitDays returns the ProfitDays field value if set, zero value otherwise.
func (o *GetCopytradingPublicStatsV5RespData) GetProfitDays() string {
	if o == nil || IsNil(o.ProfitDays) {
		var ret string
		return ret
	}
	return *o.ProfitDays
}

// GetProfitDaysOk returns a tuple with the ProfitDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicStatsV5RespData) GetProfitDaysOk() (*string, bool) {
	if o == nil || IsNil(o.ProfitDays) {
		return nil, false
	}
	return o.ProfitDays, true
}

// HasProfitDays returns a boolean if a field has been set.
func (o *GetCopytradingPublicStatsV5RespData) HasProfitDays() bool {
	if o != nil && !IsNil(o.ProfitDays) {
		return true
	}

	return false
}

// SetProfitDays gets a reference to the given string and assigns it to the ProfitDays field.
func (o *GetCopytradingPublicStatsV5RespData) SetProfitDays(v string) {
	o.ProfitDays = &v
}

// GetWinRatio returns the WinRatio field value if set, zero value otherwise.
func (o *GetCopytradingPublicStatsV5RespData) GetWinRatio() string {
	if o == nil || IsNil(o.WinRatio) {
		var ret string
		return ret
	}
	return *o.WinRatio
}

// GetWinRatioOk returns a tuple with the WinRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicStatsV5RespData) GetWinRatioOk() (*string, bool) {
	if o == nil || IsNil(o.WinRatio) {
		return nil, false
	}
	return o.WinRatio, true
}

// HasWinRatio returns a boolean if a field has been set.
func (o *GetCopytradingPublicStatsV5RespData) HasWinRatio() bool {
	if o != nil && !IsNil(o.WinRatio) {
		return true
	}

	return false
}

// SetWinRatio gets a reference to the given string and assigns it to the WinRatio field.
func (o *GetCopytradingPublicStatsV5RespData) SetWinRatio(v string) {
	o.WinRatio = &v
}

func (o GetCopytradingPublicStatsV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCopytradingPublicStatsV5RespData) ToMap() (map[string]interface{}, error) {
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

type NullableGetCopytradingPublicStatsV5RespData struct {
	value *GetCopytradingPublicStatsV5RespData
	isSet bool
}

func (v NullableGetCopytradingPublicStatsV5RespData) Get() *GetCopytradingPublicStatsV5RespData {
	return v.value
}

func (v *NullableGetCopytradingPublicStatsV5RespData) Set(val *GetCopytradingPublicStatsV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCopytradingPublicStatsV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCopytradingPublicStatsV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCopytradingPublicStatsV5RespData(val *GetCopytradingPublicStatsV5RespData) *NullableGetCopytradingPublicStatsV5RespData {
	return &NullableGetCopytradingPublicStatsV5RespData{value: val, isSet: true}
}

func (v NullableGetCopytradingPublicStatsV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCopytradingPublicStatsV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


