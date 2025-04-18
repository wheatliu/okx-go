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

// checks if the GetCopytradingTotalUnrealizedProfitSharingV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCopytradingTotalUnrealizedProfitSharingV5RespData{}

// GetCopytradingTotalUnrealizedProfitSharingV5RespData struct for GetCopytradingTotalUnrealizedProfitSharingV5RespData
type GetCopytradingTotalUnrealizedProfitSharingV5RespData struct {
	// The settlement time for the total unrealized profit sharing amount. Unix timestamp format in milliseconds, e.g.1597026383085
	ProfitSharingTs *string `json:"profitSharingTs,omitempty"`
	// Total unrealized profit sharing amount
	TotalUnrealizedProfitSharingAmt *string `json:"totalUnrealizedProfitSharingAmt,omitempty"`
}

// NewGetCopytradingTotalUnrealizedProfitSharingV5RespData instantiates a new GetCopytradingTotalUnrealizedProfitSharingV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCopytradingTotalUnrealizedProfitSharingV5RespData() *GetCopytradingTotalUnrealizedProfitSharingV5RespData {
	this := GetCopytradingTotalUnrealizedProfitSharingV5RespData{}
	var profitSharingTs string = ""
	this.ProfitSharingTs = &profitSharingTs
	var totalUnrealizedProfitSharingAmt string = ""
	this.TotalUnrealizedProfitSharingAmt = &totalUnrealizedProfitSharingAmt
	return &this
}

// NewGetCopytradingTotalUnrealizedProfitSharingV5RespDataWithDefaults instantiates a new GetCopytradingTotalUnrealizedProfitSharingV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCopytradingTotalUnrealizedProfitSharingV5RespDataWithDefaults() *GetCopytradingTotalUnrealizedProfitSharingV5RespData {
	this := GetCopytradingTotalUnrealizedProfitSharingV5RespData{}
	var profitSharingTs string = ""
	this.ProfitSharingTs = &profitSharingTs
	var totalUnrealizedProfitSharingAmt string = ""
	this.TotalUnrealizedProfitSharingAmt = &totalUnrealizedProfitSharingAmt
	return &this
}

// GetProfitSharingTs returns the ProfitSharingTs field value if set, zero value otherwise.
func (o *GetCopytradingTotalUnrealizedProfitSharingV5RespData) GetProfitSharingTs() string {
	if o == nil || IsNil(o.ProfitSharingTs) {
		var ret string
		return ret
	}
	return *o.ProfitSharingTs
}

// GetProfitSharingTsOk returns a tuple with the ProfitSharingTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingTotalUnrealizedProfitSharingV5RespData) GetProfitSharingTsOk() (*string, bool) {
	if o == nil || IsNil(o.ProfitSharingTs) {
		return nil, false
	}
	return o.ProfitSharingTs, true
}

// HasProfitSharingTs returns a boolean if a field has been set.
func (o *GetCopytradingTotalUnrealizedProfitSharingV5RespData) HasProfitSharingTs() bool {
	if o != nil && !IsNil(o.ProfitSharingTs) {
		return true
	}

	return false
}

// SetProfitSharingTs gets a reference to the given string and assigns it to the ProfitSharingTs field.
func (o *GetCopytradingTotalUnrealizedProfitSharingV5RespData) SetProfitSharingTs(v string) {
	o.ProfitSharingTs = &v
}

// GetTotalUnrealizedProfitSharingAmt returns the TotalUnrealizedProfitSharingAmt field value if set, zero value otherwise.
func (o *GetCopytradingTotalUnrealizedProfitSharingV5RespData) GetTotalUnrealizedProfitSharingAmt() string {
	if o == nil || IsNil(o.TotalUnrealizedProfitSharingAmt) {
		var ret string
		return ret
	}
	return *o.TotalUnrealizedProfitSharingAmt
}

// GetTotalUnrealizedProfitSharingAmtOk returns a tuple with the TotalUnrealizedProfitSharingAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingTotalUnrealizedProfitSharingV5RespData) GetTotalUnrealizedProfitSharingAmtOk() (*string, bool) {
	if o == nil || IsNil(o.TotalUnrealizedProfitSharingAmt) {
		return nil, false
	}
	return o.TotalUnrealizedProfitSharingAmt, true
}

// HasTotalUnrealizedProfitSharingAmt returns a boolean if a field has been set.
func (o *GetCopytradingTotalUnrealizedProfitSharingV5RespData) HasTotalUnrealizedProfitSharingAmt() bool {
	if o != nil && !IsNil(o.TotalUnrealizedProfitSharingAmt) {
		return true
	}

	return false
}

// SetTotalUnrealizedProfitSharingAmt gets a reference to the given string and assigns it to the TotalUnrealizedProfitSharingAmt field.
func (o *GetCopytradingTotalUnrealizedProfitSharingV5RespData) SetTotalUnrealizedProfitSharingAmt(v string) {
	o.TotalUnrealizedProfitSharingAmt = &v
}

func (o GetCopytradingTotalUnrealizedProfitSharingV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCopytradingTotalUnrealizedProfitSharingV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProfitSharingTs) {
		toSerialize["profitSharingTs"] = o.ProfitSharingTs
	}
	if !IsNil(o.TotalUnrealizedProfitSharingAmt) {
		toSerialize["totalUnrealizedProfitSharingAmt"] = o.TotalUnrealizedProfitSharingAmt
	}
	return toSerialize, nil
}

type NullableGetCopytradingTotalUnrealizedProfitSharingV5RespData struct {
	value *GetCopytradingTotalUnrealizedProfitSharingV5RespData
	isSet bool
}

func (v NullableGetCopytradingTotalUnrealizedProfitSharingV5RespData) Get() *GetCopytradingTotalUnrealizedProfitSharingV5RespData {
	return v.value
}

func (v *NullableGetCopytradingTotalUnrealizedProfitSharingV5RespData) Set(val *GetCopytradingTotalUnrealizedProfitSharingV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCopytradingTotalUnrealizedProfitSharingV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCopytradingTotalUnrealizedProfitSharingV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCopytradingTotalUnrealizedProfitSharingV5RespData(val *GetCopytradingTotalUnrealizedProfitSharingV5RespData) *NullableGetCopytradingTotalUnrealizedProfitSharingV5RespData {
	return &NullableGetCopytradingTotalUnrealizedProfitSharingV5RespData{value: val, isSet: true}
}

func (v NullableGetCopytradingTotalUnrealizedProfitSharingV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCopytradingTotalUnrealizedProfitSharingV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


