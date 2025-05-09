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

// checks if the GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam{}

// GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam Entry setting
type GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam struct {
	// Whether or not allow multiple entries in the same direction for the same trading pairs
	AllowMultipleEntry *bool `json:"allowMultipleEntry,omitempty"`
	// Amount per order  Only applicable to `entryType` in `2`/`3`
	Amt *string `json:"amt,omitempty"`
	// Entry type  `1`: TradingView signal  `2`: Fixed margin  `3`: Contracts  `4`: Percentage of free margin  `5`: Percentage of the initial invested margin
	EntryType *string `json:"entryType,omitempty"`
	// Amount ratio per order  Only applicable to `entryType` in `4`/`5`
	Ratio *string `json:"ratio,omitempty"`
}

// NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam instantiates a new GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam() *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam {
	this := GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam{}
	var amt string = ""
	this.Amt = &amt
	var entryType string = ""
	this.EntryType = &entryType
	var ratio string = ""
	this.Ratio = &ratio
	return &this
}

// NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParamWithDefaults instantiates a new GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParamWithDefaults() *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam {
	this := GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam{}
	var amt string = ""
	this.Amt = &amt
	var entryType string = ""
	this.EntryType = &entryType
	var ratio string = ""
	this.Ratio = &ratio
	return &this
}

// GetAllowMultipleEntry returns the AllowMultipleEntry field value if set, zero value otherwise.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) GetAllowMultipleEntry() bool {
	if o == nil || IsNil(o.AllowMultipleEntry) {
		var ret bool
		return ret
	}
	return *o.AllowMultipleEntry
}

// GetAllowMultipleEntryOk returns a tuple with the AllowMultipleEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) GetAllowMultipleEntryOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowMultipleEntry) {
		return nil, false
	}
	return o.AllowMultipleEntry, true
}

// HasAllowMultipleEntry returns a boolean if a field has been set.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) HasAllowMultipleEntry() bool {
	if o != nil && !IsNil(o.AllowMultipleEntry) {
		return true
	}

	return false
}

// SetAllowMultipleEntry gets a reference to the given bool and assigns it to the AllowMultipleEntry field.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) SetAllowMultipleEntry(v bool) {
	o.AllowMultipleEntry = &v
}

// GetAmt returns the Amt field value if set, zero value otherwise.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) GetAmt() string {
	if o == nil || IsNil(o.Amt) {
		var ret string
		return ret
	}
	return *o.Amt
}

// GetAmtOk returns a tuple with the Amt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) GetAmtOk() (*string, bool) {
	if o == nil || IsNil(o.Amt) {
		return nil, false
	}
	return o.Amt, true
}

// HasAmt returns a boolean if a field has been set.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) HasAmt() bool {
	if o != nil && !IsNil(o.Amt) {
		return true
	}

	return false
}

// SetAmt gets a reference to the given string and assigns it to the Amt field.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) SetAmt(v string) {
	o.Amt = &v
}

// GetEntryType returns the EntryType field value if set, zero value otherwise.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) GetEntryType() string {
	if o == nil || IsNil(o.EntryType) {
		var ret string
		return ret
	}
	return *o.EntryType
}

// GetEntryTypeOk returns a tuple with the EntryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) GetEntryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntryType) {
		return nil, false
	}
	return o.EntryType, true
}

// HasEntryType returns a boolean if a field has been set.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) HasEntryType() bool {
	if o != nil && !IsNil(o.EntryType) {
		return true
	}

	return false
}

// SetEntryType gets a reference to the given string and assigns it to the EntryType field.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) SetEntryType(v string) {
	o.EntryType = &v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) GetRatio() string {
	if o == nil || IsNil(o.Ratio) {
		var ret string
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) GetRatioOk() (*string, bool) {
	if o == nil || IsNil(o.Ratio) {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) HasRatio() bool {
	if o != nil && !IsNil(o.Ratio) {
		return true
	}

	return false
}

// SetRatio gets a reference to the given string and assigns it to the Ratio field.
func (o *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) SetRatio(v string) {
	o.Ratio = &v
}

func (o GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowMultipleEntry) {
		toSerialize["allowMultipleEntry"] = o.AllowMultipleEntry
	}
	if !IsNil(o.Amt) {
		toSerialize["amt"] = o.Amt
	}
	if !IsNil(o.EntryType) {
		toSerialize["entryType"] = o.EntryType
	}
	if !IsNil(o.Ratio) {
		toSerialize["ratio"] = o.Ratio
	}
	return toSerialize, nil
}

type NullableGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam struct {
	value *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam
	isSet bool
}

func (v NullableGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) Get() *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam {
	return v.value
}

func (v *NullableGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) Set(val *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam(val *GetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) *NullableGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam {
	return &NullableGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam{value: val, isSet: true}
}

func (v NullableGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTradingBotSignalOrdersAlgoDetailsV5RespDataInnerEntrySettingParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


