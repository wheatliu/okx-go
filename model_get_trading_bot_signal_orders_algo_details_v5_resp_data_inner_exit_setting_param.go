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

// checks if the GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam{}

// GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam Exit setting
type GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam struct {
	// Stop-loss percentage
	SlPct *string `json:"slPct,omitempty"`
	// Take-profit percentage
	TpPct *string `json:"tpPct,omitempty"`
	// Type of set the take-profit and stop-loss trigger price  `pnl`: Based on the estimated profit and loss percentage from the entry point  `price`: Based on price increase or decrease from the crypto’s entry price
	TpSlType *string `json:"tpSlType,omitempty"`
}

// NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam instantiates a new GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam() *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam {
	this := GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam{}
	var slPct string = ""
	this.SlPct = &slPct
	var tpPct string = ""
	this.TpPct = &tpPct
	var tpSlType string = ""
	this.TpSlType = &tpSlType
	return &this
}

// NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParamWithDefaults instantiates a new GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParamWithDefaults() *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam {
	this := GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam{}
	var slPct string = ""
	this.SlPct = &slPct
	var tpPct string = ""
	this.TpPct = &tpPct
	var tpSlType string = ""
	this.TpSlType = &tpSlType
	return &this
}

// GetSlPct returns the SlPct field value if set, zero value otherwise.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) GetSlPct() string {
	if o == nil || IsNil(o.SlPct) {
		var ret string
		return ret
	}
	return *o.SlPct
}

// GetSlPctOk returns a tuple with the SlPct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) GetSlPctOk() (*string, bool) {
	if o == nil || IsNil(o.SlPct) {
		return nil, false
	}
	return o.SlPct, true
}

// HasSlPct returns a boolean if a field has been set.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) HasSlPct() bool {
	if o != nil && !IsNil(o.SlPct) {
		return true
	}

	return false
}

// SetSlPct gets a reference to the given string and assigns it to the SlPct field.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) SetSlPct(v string) {
	o.SlPct = &v
}

// GetTpPct returns the TpPct field value if set, zero value otherwise.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) GetTpPct() string {
	if o == nil || IsNil(o.TpPct) {
		var ret string
		return ret
	}
	return *o.TpPct
}

// GetTpPctOk returns a tuple with the TpPct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) GetTpPctOk() (*string, bool) {
	if o == nil || IsNil(o.TpPct) {
		return nil, false
	}
	return o.TpPct, true
}

// HasTpPct returns a boolean if a field has been set.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) HasTpPct() bool {
	if o != nil && !IsNil(o.TpPct) {
		return true
	}

	return false
}

// SetTpPct gets a reference to the given string and assigns it to the TpPct field.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) SetTpPct(v string) {
	o.TpPct = &v
}

// GetTpSlType returns the TpSlType field value if set, zero value otherwise.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) GetTpSlType() string {
	if o == nil || IsNil(o.TpSlType) {
		var ret string
		return ret
	}
	return *o.TpSlType
}

// GetTpSlTypeOk returns a tuple with the TpSlType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) GetTpSlTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TpSlType) {
		return nil, false
	}
	return o.TpSlType, true
}

// HasTpSlType returns a boolean if a field has been set.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) HasTpSlType() bool {
	if o != nil && !IsNil(o.TpSlType) {
		return true
	}

	return false
}

// SetTpSlType gets a reference to the given string and assigns it to the TpSlType field.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) SetTpSlType(v string) {
	o.TpSlType = &v
}

func (o GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SlPct) {
		toSerialize["slPct"] = o.SlPct
	}
	if !IsNil(o.TpPct) {
		toSerialize["tpPct"] = o.TpPct
	}
	if !IsNil(o.TpSlType) {
		toSerialize["tpSlType"] = o.TpSlType
	}
	return toSerialize, nil
}

type NullableGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam struct {
	value *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam
	isSet bool
}

func (v NullableGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) Get() *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam {
	return v.value
}

func (v *NullableGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) Set(val *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam(val *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) *NullableGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam {
	return &NullableGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam{value: val, isSet: true}
}

func (v NullableGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerExitSettingParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


