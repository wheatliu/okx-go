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

// checks if the GetAccountMmpConfigV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountMmpConfigV5RespData{}

// GetAccountMmpConfigV5RespData struct for GetAccountMmpConfigV5RespData
type GetAccountMmpConfigV5RespData struct {
	// Frozen period (ms). If it is \"0\", the trade will remain frozen until manually reset and `mmpFrozenUntil` will be \"\".
	FrozenInterval *string `json:"frozenInterval,omitempty"`
	// Instrument Family
	InstFamily *string `json:"instFamily,omitempty"`
	// Whether MMP is currently triggered. `true` or `false`
	MmpFrozen *bool `json:"mmpFrozen,omitempty"`
	// If frozenInterval is configured and mmpFrozen = True, it is the time interval (in ms) when MMP is no longer triggered, otherwise \"\".
	MmpFrozenUntil *string `json:"mmpFrozenUntil,omitempty"`
	// Trade qty limit in number of contracts
	QtyLimit *string `json:"qtyLimit,omitempty"`
	// Time window (ms). MMP interval where monitoring is done
	TimeInterval *string `json:"timeInterval,omitempty"`
}

// NewGetAccountMmpConfigV5RespData instantiates a new GetAccountMmpConfigV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountMmpConfigV5RespData() *GetAccountMmpConfigV5RespData {
	this := GetAccountMmpConfigV5RespData{}
	var frozenInterval string = ""
	this.FrozenInterval = &frozenInterval
	var instFamily string = ""
	this.InstFamily = &instFamily
	var mmpFrozenUntil string = ""
	this.MmpFrozenUntil = &mmpFrozenUntil
	var qtyLimit string = ""
	this.QtyLimit = &qtyLimit
	var timeInterval string = ""
	this.TimeInterval = &timeInterval
	return &this
}

// NewGetAccountMmpConfigV5RespDataWithDefaults instantiates a new GetAccountMmpConfigV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountMmpConfigV5RespDataWithDefaults() *GetAccountMmpConfigV5RespData {
	this := GetAccountMmpConfigV5RespData{}
	var frozenInterval string = ""
	this.FrozenInterval = &frozenInterval
	var instFamily string = ""
	this.InstFamily = &instFamily
	var mmpFrozenUntil string = ""
	this.MmpFrozenUntil = &mmpFrozenUntil
	var qtyLimit string = ""
	this.QtyLimit = &qtyLimit
	var timeInterval string = ""
	this.TimeInterval = &timeInterval
	return &this
}

// GetFrozenInterval returns the FrozenInterval field value if set, zero value otherwise.
func (o *GetAccountMmpConfigV5RespData) GetFrozenInterval() string {
	if o == nil || IsNil(o.FrozenInterval) {
		var ret string
		return ret
	}
	return *o.FrozenInterval
}

// GetFrozenIntervalOk returns a tuple with the FrozenInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountMmpConfigV5RespData) GetFrozenIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.FrozenInterval) {
		return nil, false
	}
	return o.FrozenInterval, true
}

// HasFrozenInterval returns a boolean if a field has been set.
func (o *GetAccountMmpConfigV5RespData) HasFrozenInterval() bool {
	if o != nil && !IsNil(o.FrozenInterval) {
		return true
	}

	return false
}

// SetFrozenInterval gets a reference to the given string and assigns it to the FrozenInterval field.
func (o *GetAccountMmpConfigV5RespData) SetFrozenInterval(v string) {
	o.FrozenInterval = &v
}

// GetInstFamily returns the InstFamily field value if set, zero value otherwise.
func (o *GetAccountMmpConfigV5RespData) GetInstFamily() string {
	if o == nil || IsNil(o.InstFamily) {
		var ret string
		return ret
	}
	return *o.InstFamily
}

// GetInstFamilyOk returns a tuple with the InstFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountMmpConfigV5RespData) GetInstFamilyOk() (*string, bool) {
	if o == nil || IsNil(o.InstFamily) {
		return nil, false
	}
	return o.InstFamily, true
}

// HasInstFamily returns a boolean if a field has been set.
func (o *GetAccountMmpConfigV5RespData) HasInstFamily() bool {
	if o != nil && !IsNil(o.InstFamily) {
		return true
	}

	return false
}

// SetInstFamily gets a reference to the given string and assigns it to the InstFamily field.
func (o *GetAccountMmpConfigV5RespData) SetInstFamily(v string) {
	o.InstFamily = &v
}

// GetMmpFrozen returns the MmpFrozen field value if set, zero value otherwise.
func (o *GetAccountMmpConfigV5RespData) GetMmpFrozen() bool {
	if o == nil || IsNil(o.MmpFrozen) {
		var ret bool
		return ret
	}
	return *o.MmpFrozen
}

// GetMmpFrozenOk returns a tuple with the MmpFrozen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountMmpConfigV5RespData) GetMmpFrozenOk() (*bool, bool) {
	if o == nil || IsNil(o.MmpFrozen) {
		return nil, false
	}
	return o.MmpFrozen, true
}

// HasMmpFrozen returns a boolean if a field has been set.
func (o *GetAccountMmpConfigV5RespData) HasMmpFrozen() bool {
	if o != nil && !IsNil(o.MmpFrozen) {
		return true
	}

	return false
}

// SetMmpFrozen gets a reference to the given bool and assigns it to the MmpFrozen field.
func (o *GetAccountMmpConfigV5RespData) SetMmpFrozen(v bool) {
	o.MmpFrozen = &v
}

// GetMmpFrozenUntil returns the MmpFrozenUntil field value if set, zero value otherwise.
func (o *GetAccountMmpConfigV5RespData) GetMmpFrozenUntil() string {
	if o == nil || IsNil(o.MmpFrozenUntil) {
		var ret string
		return ret
	}
	return *o.MmpFrozenUntil
}

// GetMmpFrozenUntilOk returns a tuple with the MmpFrozenUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountMmpConfigV5RespData) GetMmpFrozenUntilOk() (*string, bool) {
	if o == nil || IsNil(o.MmpFrozenUntil) {
		return nil, false
	}
	return o.MmpFrozenUntil, true
}

// HasMmpFrozenUntil returns a boolean if a field has been set.
func (o *GetAccountMmpConfigV5RespData) HasMmpFrozenUntil() bool {
	if o != nil && !IsNil(o.MmpFrozenUntil) {
		return true
	}

	return false
}

// SetMmpFrozenUntil gets a reference to the given string and assigns it to the MmpFrozenUntil field.
func (o *GetAccountMmpConfigV5RespData) SetMmpFrozenUntil(v string) {
	o.MmpFrozenUntil = &v
}

// GetQtyLimit returns the QtyLimit field value if set, zero value otherwise.
func (o *GetAccountMmpConfigV5RespData) GetQtyLimit() string {
	if o == nil || IsNil(o.QtyLimit) {
		var ret string
		return ret
	}
	return *o.QtyLimit
}

// GetQtyLimitOk returns a tuple with the QtyLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountMmpConfigV5RespData) GetQtyLimitOk() (*string, bool) {
	if o == nil || IsNil(o.QtyLimit) {
		return nil, false
	}
	return o.QtyLimit, true
}

// HasQtyLimit returns a boolean if a field has been set.
func (o *GetAccountMmpConfigV5RespData) HasQtyLimit() bool {
	if o != nil && !IsNil(o.QtyLimit) {
		return true
	}

	return false
}

// SetQtyLimit gets a reference to the given string and assigns it to the QtyLimit field.
func (o *GetAccountMmpConfigV5RespData) SetQtyLimit(v string) {
	o.QtyLimit = &v
}

// GetTimeInterval returns the TimeInterval field value if set, zero value otherwise.
func (o *GetAccountMmpConfigV5RespData) GetTimeInterval() string {
	if o == nil || IsNil(o.TimeInterval) {
		var ret string
		return ret
	}
	return *o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountMmpConfigV5RespData) GetTimeIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.TimeInterval) {
		return nil, false
	}
	return o.TimeInterval, true
}

// HasTimeInterval returns a boolean if a field has been set.
func (o *GetAccountMmpConfigV5RespData) HasTimeInterval() bool {
	if o != nil && !IsNil(o.TimeInterval) {
		return true
	}

	return false
}

// SetTimeInterval gets a reference to the given string and assigns it to the TimeInterval field.
func (o *GetAccountMmpConfigV5RespData) SetTimeInterval(v string) {
	o.TimeInterval = &v
}

func (o GetAccountMmpConfigV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountMmpConfigV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FrozenInterval) {
		toSerialize["frozenInterval"] = o.FrozenInterval
	}
	if !IsNil(o.InstFamily) {
		toSerialize["instFamily"] = o.InstFamily
	}
	if !IsNil(o.MmpFrozen) {
		toSerialize["mmpFrozen"] = o.MmpFrozen
	}
	if !IsNil(o.MmpFrozenUntil) {
		toSerialize["mmpFrozenUntil"] = o.MmpFrozenUntil
	}
	if !IsNil(o.QtyLimit) {
		toSerialize["qtyLimit"] = o.QtyLimit
	}
	if !IsNil(o.TimeInterval) {
		toSerialize["timeInterval"] = o.TimeInterval
	}
	return toSerialize, nil
}

type NullableGetAccountMmpConfigV5RespData struct {
	value *GetAccountMmpConfigV5RespData
	isSet bool
}

func (v NullableGetAccountMmpConfigV5RespData) Get() *GetAccountMmpConfigV5RespData {
	return v.value
}

func (v *NullableGetAccountMmpConfigV5RespData) Set(val *GetAccountMmpConfigV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountMmpConfigV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountMmpConfigV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountMmpConfigV5RespData(val *GetAccountMmpConfigV5RespData) *NullableGetAccountMmpConfigV5RespData {
	return &NullableGetAccountMmpConfigV5RespData{value: val, isSet: true}
}

func (v NullableGetAccountMmpConfigV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountMmpConfigV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


