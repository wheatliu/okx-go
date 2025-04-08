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

// checks if the GetCopytradingPublicConfigV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCopytradingPublicConfigV5RespData{}

// GetCopytradingPublicConfigV5RespData struct for GetCopytradingPublicConfigV5RespData
type GetCopytradingPublicConfigV5RespData struct {
	// Maximum copy amount per order in USDT when you are using copy mode `fixed_amount`
	MaxCopyAmt *string `json:"maxCopyAmt,omitempty"`
	// Maximum ratio per order when you are using copy mode `ratio_copy`
	MaxCopyRatio *string `json:"maxCopyRatio,omitempty"`
	// Maximum copy total amount under the certain lead trader, the minimum is the same with `minCopyAmt`
	MaxCopyTotalAmt *string `json:"maxCopyTotalAmt,omitempty"`
	// Maximum ratio of stopping loss per order, the minimum is 0
	MaxSlRatio *string `json:"maxSlRatio,omitempty"`
	// Maximum ratio of taking profit per order, the minimum is 0
	MaxTpRatio *string `json:"maxTpRatio,omitempty"`
	// Minimum copy amount per order in USDT when you are using copy mode `fixed_amount`
	MinCopyAmt *string `json:"minCopyAmt,omitempty"`
	// Minimum ratio per order when you are using copy mode `ratio_copy`
	MinCopyRatio *string `json:"minCopyRatio,omitempty"`
}

// NewGetCopytradingPublicConfigV5RespData instantiates a new GetCopytradingPublicConfigV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCopytradingPublicConfigV5RespData() *GetCopytradingPublicConfigV5RespData {
	this := GetCopytradingPublicConfigV5RespData{}
	var maxCopyAmt string = ""
	this.MaxCopyAmt = &maxCopyAmt
	var maxCopyRatio string = ""
	this.MaxCopyRatio = &maxCopyRatio
	var maxCopyTotalAmt string = ""
	this.MaxCopyTotalAmt = &maxCopyTotalAmt
	var maxSlRatio string = ""
	this.MaxSlRatio = &maxSlRatio
	var maxTpRatio string = ""
	this.MaxTpRatio = &maxTpRatio
	var minCopyAmt string = ""
	this.MinCopyAmt = &minCopyAmt
	var minCopyRatio string = ""
	this.MinCopyRatio = &minCopyRatio
	return &this
}

// NewGetCopytradingPublicConfigV5RespDataWithDefaults instantiates a new GetCopytradingPublicConfigV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCopytradingPublicConfigV5RespDataWithDefaults() *GetCopytradingPublicConfigV5RespData {
	this := GetCopytradingPublicConfigV5RespData{}
	var maxCopyAmt string = ""
	this.MaxCopyAmt = &maxCopyAmt
	var maxCopyRatio string = ""
	this.MaxCopyRatio = &maxCopyRatio
	var maxCopyTotalAmt string = ""
	this.MaxCopyTotalAmt = &maxCopyTotalAmt
	var maxSlRatio string = ""
	this.MaxSlRatio = &maxSlRatio
	var maxTpRatio string = ""
	this.MaxTpRatio = &maxTpRatio
	var minCopyAmt string = ""
	this.MinCopyAmt = &minCopyAmt
	var minCopyRatio string = ""
	this.MinCopyRatio = &minCopyRatio
	return &this
}

// GetMaxCopyAmt returns the MaxCopyAmt field value if set, zero value otherwise.
func (o *GetCopytradingPublicConfigV5RespData) GetMaxCopyAmt() string {
	if o == nil || IsNil(o.MaxCopyAmt) {
		var ret string
		return ret
	}
	return *o.MaxCopyAmt
}

// GetMaxCopyAmtOk returns a tuple with the MaxCopyAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicConfigV5RespData) GetMaxCopyAmtOk() (*string, bool) {
	if o == nil || IsNil(o.MaxCopyAmt) {
		return nil, false
	}
	return o.MaxCopyAmt, true
}

// HasMaxCopyAmt returns a boolean if a field has been set.
func (o *GetCopytradingPublicConfigV5RespData) HasMaxCopyAmt() bool {
	if o != nil && !IsNil(o.MaxCopyAmt) {
		return true
	}

	return false
}

// SetMaxCopyAmt gets a reference to the given string and assigns it to the MaxCopyAmt field.
func (o *GetCopytradingPublicConfigV5RespData) SetMaxCopyAmt(v string) {
	o.MaxCopyAmt = &v
}

// GetMaxCopyRatio returns the MaxCopyRatio field value if set, zero value otherwise.
func (o *GetCopytradingPublicConfigV5RespData) GetMaxCopyRatio() string {
	if o == nil || IsNil(o.MaxCopyRatio) {
		var ret string
		return ret
	}
	return *o.MaxCopyRatio
}

// GetMaxCopyRatioOk returns a tuple with the MaxCopyRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicConfigV5RespData) GetMaxCopyRatioOk() (*string, bool) {
	if o == nil || IsNil(o.MaxCopyRatio) {
		return nil, false
	}
	return o.MaxCopyRatio, true
}

// HasMaxCopyRatio returns a boolean if a field has been set.
func (o *GetCopytradingPublicConfigV5RespData) HasMaxCopyRatio() bool {
	if o != nil && !IsNil(o.MaxCopyRatio) {
		return true
	}

	return false
}

// SetMaxCopyRatio gets a reference to the given string and assigns it to the MaxCopyRatio field.
func (o *GetCopytradingPublicConfigV5RespData) SetMaxCopyRatio(v string) {
	o.MaxCopyRatio = &v
}

// GetMaxCopyTotalAmt returns the MaxCopyTotalAmt field value if set, zero value otherwise.
func (o *GetCopytradingPublicConfigV5RespData) GetMaxCopyTotalAmt() string {
	if o == nil || IsNil(o.MaxCopyTotalAmt) {
		var ret string
		return ret
	}
	return *o.MaxCopyTotalAmt
}

// GetMaxCopyTotalAmtOk returns a tuple with the MaxCopyTotalAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicConfigV5RespData) GetMaxCopyTotalAmtOk() (*string, bool) {
	if o == nil || IsNil(o.MaxCopyTotalAmt) {
		return nil, false
	}
	return o.MaxCopyTotalAmt, true
}

// HasMaxCopyTotalAmt returns a boolean if a field has been set.
func (o *GetCopytradingPublicConfigV5RespData) HasMaxCopyTotalAmt() bool {
	if o != nil && !IsNil(o.MaxCopyTotalAmt) {
		return true
	}

	return false
}

// SetMaxCopyTotalAmt gets a reference to the given string and assigns it to the MaxCopyTotalAmt field.
func (o *GetCopytradingPublicConfigV5RespData) SetMaxCopyTotalAmt(v string) {
	o.MaxCopyTotalAmt = &v
}

// GetMaxSlRatio returns the MaxSlRatio field value if set, zero value otherwise.
func (o *GetCopytradingPublicConfigV5RespData) GetMaxSlRatio() string {
	if o == nil || IsNil(o.MaxSlRatio) {
		var ret string
		return ret
	}
	return *o.MaxSlRatio
}

// GetMaxSlRatioOk returns a tuple with the MaxSlRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicConfigV5RespData) GetMaxSlRatioOk() (*string, bool) {
	if o == nil || IsNil(o.MaxSlRatio) {
		return nil, false
	}
	return o.MaxSlRatio, true
}

// HasMaxSlRatio returns a boolean if a field has been set.
func (o *GetCopytradingPublicConfigV5RespData) HasMaxSlRatio() bool {
	if o != nil && !IsNil(o.MaxSlRatio) {
		return true
	}

	return false
}

// SetMaxSlRatio gets a reference to the given string and assigns it to the MaxSlRatio field.
func (o *GetCopytradingPublicConfigV5RespData) SetMaxSlRatio(v string) {
	o.MaxSlRatio = &v
}

// GetMaxTpRatio returns the MaxTpRatio field value if set, zero value otherwise.
func (o *GetCopytradingPublicConfigV5RespData) GetMaxTpRatio() string {
	if o == nil || IsNil(o.MaxTpRatio) {
		var ret string
		return ret
	}
	return *o.MaxTpRatio
}

// GetMaxTpRatioOk returns a tuple with the MaxTpRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicConfigV5RespData) GetMaxTpRatioOk() (*string, bool) {
	if o == nil || IsNil(o.MaxTpRatio) {
		return nil, false
	}
	return o.MaxTpRatio, true
}

// HasMaxTpRatio returns a boolean if a field has been set.
func (o *GetCopytradingPublicConfigV5RespData) HasMaxTpRatio() bool {
	if o != nil && !IsNil(o.MaxTpRatio) {
		return true
	}

	return false
}

// SetMaxTpRatio gets a reference to the given string and assigns it to the MaxTpRatio field.
func (o *GetCopytradingPublicConfigV5RespData) SetMaxTpRatio(v string) {
	o.MaxTpRatio = &v
}

// GetMinCopyAmt returns the MinCopyAmt field value if set, zero value otherwise.
func (o *GetCopytradingPublicConfigV5RespData) GetMinCopyAmt() string {
	if o == nil || IsNil(o.MinCopyAmt) {
		var ret string
		return ret
	}
	return *o.MinCopyAmt
}

// GetMinCopyAmtOk returns a tuple with the MinCopyAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicConfigV5RespData) GetMinCopyAmtOk() (*string, bool) {
	if o == nil || IsNil(o.MinCopyAmt) {
		return nil, false
	}
	return o.MinCopyAmt, true
}

// HasMinCopyAmt returns a boolean if a field has been set.
func (o *GetCopytradingPublicConfigV5RespData) HasMinCopyAmt() bool {
	if o != nil && !IsNil(o.MinCopyAmt) {
		return true
	}

	return false
}

// SetMinCopyAmt gets a reference to the given string and assigns it to the MinCopyAmt field.
func (o *GetCopytradingPublicConfigV5RespData) SetMinCopyAmt(v string) {
	o.MinCopyAmt = &v
}

// GetMinCopyRatio returns the MinCopyRatio field value if set, zero value otherwise.
func (o *GetCopytradingPublicConfigV5RespData) GetMinCopyRatio() string {
	if o == nil || IsNil(o.MinCopyRatio) {
		var ret string
		return ret
	}
	return *o.MinCopyRatio
}

// GetMinCopyRatioOk returns a tuple with the MinCopyRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicConfigV5RespData) GetMinCopyRatioOk() (*string, bool) {
	if o == nil || IsNil(o.MinCopyRatio) {
		return nil, false
	}
	return o.MinCopyRatio, true
}

// HasMinCopyRatio returns a boolean if a field has been set.
func (o *GetCopytradingPublicConfigV5RespData) HasMinCopyRatio() bool {
	if o != nil && !IsNil(o.MinCopyRatio) {
		return true
	}

	return false
}

// SetMinCopyRatio gets a reference to the given string and assigns it to the MinCopyRatio field.
func (o *GetCopytradingPublicConfigV5RespData) SetMinCopyRatio(v string) {
	o.MinCopyRatio = &v
}

func (o GetCopytradingPublicConfigV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCopytradingPublicConfigV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxCopyAmt) {
		toSerialize["maxCopyAmt"] = o.MaxCopyAmt
	}
	if !IsNil(o.MaxCopyRatio) {
		toSerialize["maxCopyRatio"] = o.MaxCopyRatio
	}
	if !IsNil(o.MaxCopyTotalAmt) {
		toSerialize["maxCopyTotalAmt"] = o.MaxCopyTotalAmt
	}
	if !IsNil(o.MaxSlRatio) {
		toSerialize["maxSlRatio"] = o.MaxSlRatio
	}
	if !IsNil(o.MaxTpRatio) {
		toSerialize["maxTpRatio"] = o.MaxTpRatio
	}
	if !IsNil(o.MinCopyAmt) {
		toSerialize["minCopyAmt"] = o.MinCopyAmt
	}
	if !IsNil(o.MinCopyRatio) {
		toSerialize["minCopyRatio"] = o.MinCopyRatio
	}
	return toSerialize, nil
}

type NullableGetCopytradingPublicConfigV5RespData struct {
	value *GetCopytradingPublicConfigV5RespData
	isSet bool
}

func (v NullableGetCopytradingPublicConfigV5RespData) Get() *GetCopytradingPublicConfigV5RespData {
	return v.value
}

func (v *NullableGetCopytradingPublicConfigV5RespData) Set(val *GetCopytradingPublicConfigV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCopytradingPublicConfigV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCopytradingPublicConfigV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCopytradingPublicConfigV5RespData(val *GetCopytradingPublicConfigV5RespData) *NullableGetCopytradingPublicConfigV5RespData {
	return &NullableGetCopytradingPublicConfigV5RespData{value: val, isSet: true}
}

func (v NullableGetCopytradingPublicConfigV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCopytradingPublicConfigV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


