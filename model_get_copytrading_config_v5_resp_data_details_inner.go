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

// checks if the GetCopytradingConfigV5RespDataDetailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCopytradingConfigV5RespDataDetailsInner{}

// GetCopytradingConfigV5RespDataDetailsInner struct for GetCopytradingConfigV5RespDataDetailsInner
type GetCopytradingConfigV5RespDataDetailsInner struct {
	// Current number of copy traders
	CopyTraderNum *string `json:"copyTraderNum,omitempty"`
	// Instrument type  `SPOT`  `SWAP`
	InstType *string `json:"instType,omitempty"`
	// Maximum number of copy traders
	MaxCopyTraderNum *string `json:"maxCopyTraderNum,omitempty"`
	// Profit sharing ratio.   Only applicable to lead trader, or it will be \"\". 0.1 represents 10%
	ProfitSharingRatio *string `json:"profitSharingRatio,omitempty"`
	// Role type  `0`: General user  `1`: Leading trader  `2`: Copy trader
	RoleType *string `json:"roleType,omitempty"`
}

// NewGetCopytradingConfigV5RespDataDetailsInner instantiates a new GetCopytradingConfigV5RespDataDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCopytradingConfigV5RespDataDetailsInner() *GetCopytradingConfigV5RespDataDetailsInner {
	this := GetCopytradingConfigV5RespDataDetailsInner{}
	var copyTraderNum string = ""
	this.CopyTraderNum = &copyTraderNum
	var instType string = ""
	this.InstType = &instType
	var maxCopyTraderNum string = ""
	this.MaxCopyTraderNum = &maxCopyTraderNum
	var profitSharingRatio string = ""
	this.ProfitSharingRatio = &profitSharingRatio
	var roleType string = ""
	this.RoleType = &roleType
	return &this
}

// NewGetCopytradingConfigV5RespDataDetailsInnerWithDefaults instantiates a new GetCopytradingConfigV5RespDataDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCopytradingConfigV5RespDataDetailsInnerWithDefaults() *GetCopytradingConfigV5RespDataDetailsInner {
	this := GetCopytradingConfigV5RespDataDetailsInner{}
	var copyTraderNum string = ""
	this.CopyTraderNum = &copyTraderNum
	var instType string = ""
	this.InstType = &instType
	var maxCopyTraderNum string = ""
	this.MaxCopyTraderNum = &maxCopyTraderNum
	var profitSharingRatio string = ""
	this.ProfitSharingRatio = &profitSharingRatio
	var roleType string = ""
	this.RoleType = &roleType
	return &this
}

// GetCopyTraderNum returns the CopyTraderNum field value if set, zero value otherwise.
func (o *GetCopytradingConfigV5RespDataDetailsInner) GetCopyTraderNum() string {
	if o == nil || IsNil(o.CopyTraderNum) {
		var ret string
		return ret
	}
	return *o.CopyTraderNum
}

// GetCopyTraderNumOk returns a tuple with the CopyTraderNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingConfigV5RespDataDetailsInner) GetCopyTraderNumOk() (*string, bool) {
	if o == nil || IsNil(o.CopyTraderNum) {
		return nil, false
	}
	return o.CopyTraderNum, true
}

// HasCopyTraderNum returns a boolean if a field has been set.
func (o *GetCopytradingConfigV5RespDataDetailsInner) HasCopyTraderNum() bool {
	if o != nil && !IsNil(o.CopyTraderNum) {
		return true
	}

	return false
}

// SetCopyTraderNum gets a reference to the given string and assigns it to the CopyTraderNum field.
func (o *GetCopytradingConfigV5RespDataDetailsInner) SetCopyTraderNum(v string) {
	o.CopyTraderNum = &v
}

// GetInstType returns the InstType field value if set, zero value otherwise.
func (o *GetCopytradingConfigV5RespDataDetailsInner) GetInstType() string {
	if o == nil || IsNil(o.InstType) {
		var ret string
		return ret
	}
	return *o.InstType
}

// GetInstTypeOk returns a tuple with the InstType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingConfigV5RespDataDetailsInner) GetInstTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstType) {
		return nil, false
	}
	return o.InstType, true
}

// HasInstType returns a boolean if a field has been set.
func (o *GetCopytradingConfigV5RespDataDetailsInner) HasInstType() bool {
	if o != nil && !IsNil(o.InstType) {
		return true
	}

	return false
}

// SetInstType gets a reference to the given string and assigns it to the InstType field.
func (o *GetCopytradingConfigV5RespDataDetailsInner) SetInstType(v string) {
	o.InstType = &v
}

// GetMaxCopyTraderNum returns the MaxCopyTraderNum field value if set, zero value otherwise.
func (o *GetCopytradingConfigV5RespDataDetailsInner) GetMaxCopyTraderNum() string {
	if o == nil || IsNil(o.MaxCopyTraderNum) {
		var ret string
		return ret
	}
	return *o.MaxCopyTraderNum
}

// GetMaxCopyTraderNumOk returns a tuple with the MaxCopyTraderNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingConfigV5RespDataDetailsInner) GetMaxCopyTraderNumOk() (*string, bool) {
	if o == nil || IsNil(o.MaxCopyTraderNum) {
		return nil, false
	}
	return o.MaxCopyTraderNum, true
}

// HasMaxCopyTraderNum returns a boolean if a field has been set.
func (o *GetCopytradingConfigV5RespDataDetailsInner) HasMaxCopyTraderNum() bool {
	if o != nil && !IsNil(o.MaxCopyTraderNum) {
		return true
	}

	return false
}

// SetMaxCopyTraderNum gets a reference to the given string and assigns it to the MaxCopyTraderNum field.
func (o *GetCopytradingConfigV5RespDataDetailsInner) SetMaxCopyTraderNum(v string) {
	o.MaxCopyTraderNum = &v
}

// GetProfitSharingRatio returns the ProfitSharingRatio field value if set, zero value otherwise.
func (o *GetCopytradingConfigV5RespDataDetailsInner) GetProfitSharingRatio() string {
	if o == nil || IsNil(o.ProfitSharingRatio) {
		var ret string
		return ret
	}
	return *o.ProfitSharingRatio
}

// GetProfitSharingRatioOk returns a tuple with the ProfitSharingRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingConfigV5RespDataDetailsInner) GetProfitSharingRatioOk() (*string, bool) {
	if o == nil || IsNil(o.ProfitSharingRatio) {
		return nil, false
	}
	return o.ProfitSharingRatio, true
}

// HasProfitSharingRatio returns a boolean if a field has been set.
func (o *GetCopytradingConfigV5RespDataDetailsInner) HasProfitSharingRatio() bool {
	if o != nil && !IsNil(o.ProfitSharingRatio) {
		return true
	}

	return false
}

// SetProfitSharingRatio gets a reference to the given string and assigns it to the ProfitSharingRatio field.
func (o *GetCopytradingConfigV5RespDataDetailsInner) SetProfitSharingRatio(v string) {
	o.ProfitSharingRatio = &v
}

// GetRoleType returns the RoleType field value if set, zero value otherwise.
func (o *GetCopytradingConfigV5RespDataDetailsInner) GetRoleType() string {
	if o == nil || IsNil(o.RoleType) {
		var ret string
		return ret
	}
	return *o.RoleType
}

// GetRoleTypeOk returns a tuple with the RoleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingConfigV5RespDataDetailsInner) GetRoleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RoleType) {
		return nil, false
	}
	return o.RoleType, true
}

// HasRoleType returns a boolean if a field has been set.
func (o *GetCopytradingConfigV5RespDataDetailsInner) HasRoleType() bool {
	if o != nil && !IsNil(o.RoleType) {
		return true
	}

	return false
}

// SetRoleType gets a reference to the given string and assigns it to the RoleType field.
func (o *GetCopytradingConfigV5RespDataDetailsInner) SetRoleType(v string) {
	o.RoleType = &v
}

func (o GetCopytradingConfigV5RespDataDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCopytradingConfigV5RespDataDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CopyTraderNum) {
		toSerialize["copyTraderNum"] = o.CopyTraderNum
	}
	if !IsNil(o.InstType) {
		toSerialize["instType"] = o.InstType
	}
	if !IsNil(o.MaxCopyTraderNum) {
		toSerialize["maxCopyTraderNum"] = o.MaxCopyTraderNum
	}
	if !IsNil(o.ProfitSharingRatio) {
		toSerialize["profitSharingRatio"] = o.ProfitSharingRatio
	}
	if !IsNil(o.RoleType) {
		toSerialize["roleType"] = o.RoleType
	}
	return toSerialize, nil
}

type NullableGetCopytradingConfigV5RespDataDetailsInner struct {
	value *GetCopytradingConfigV5RespDataDetailsInner
	isSet bool
}

func (v NullableGetCopytradingConfigV5RespDataDetailsInner) Get() *GetCopytradingConfigV5RespDataDetailsInner {
	return v.value
}

func (v *NullableGetCopytradingConfigV5RespDataDetailsInner) Set(val *GetCopytradingConfigV5RespDataDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCopytradingConfigV5RespDataDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCopytradingConfigV5RespDataDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCopytradingConfigV5RespDataDetailsInner(val *GetCopytradingConfigV5RespDataDetailsInner) *NullableGetCopytradingConfigV5RespDataDetailsInner {
	return &NullableGetCopytradingConfigV5RespDataDetailsInner{value: val, isSet: true}
}

func (v NullableGetCopytradingConfigV5RespDataDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCopytradingConfigV5RespDataDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


