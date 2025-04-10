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

// checks if the GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner{}

// GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner struct for GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner
type GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner struct {
	// Amount of small payment currency convert from
	FromAmt *string `json:"fromAmt,omitempty"`
	// Type of small payment currency convert from, e.g. `BTC`
	FromCcy *string `json:"fromCcy,omitempty"`
}

// NewGetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner instantiates a new GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner() *GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner {
	this := GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner{}
	var fromAmt string = ""
	this.FromAmt = &fromAmt
	var fromCcy string = ""
	this.FromCcy = &fromCcy
	return &this
}

// NewGetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInnerWithDefaults instantiates a new GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInnerWithDefaults() *GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner {
	this := GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner{}
	var fromAmt string = ""
	this.FromAmt = &fromAmt
	var fromCcy string = ""
	this.FromCcy = &fromCcy
	return &this
}

// GetFromAmt returns the FromAmt field value if set, zero value otherwise.
func (o *GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner) GetFromAmt() string {
	if o == nil || IsNil(o.FromAmt) {
		var ret string
		return ret
	}
	return *o.FromAmt
}

// GetFromAmtOk returns a tuple with the FromAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner) GetFromAmtOk() (*string, bool) {
	if o == nil || IsNil(o.FromAmt) {
		return nil, false
	}
	return o.FromAmt, true
}

// HasFromAmt returns a boolean if a field has been set.
func (o *GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner) HasFromAmt() bool {
	if o != nil && !IsNil(o.FromAmt) {
		return true
	}

	return false
}

// SetFromAmt gets a reference to the given string and assigns it to the FromAmt field.
func (o *GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner) SetFromAmt(v string) {
	o.FromAmt = &v
}

// GetFromCcy returns the FromCcy field value if set, zero value otherwise.
func (o *GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner) GetFromCcy() string {
	if o == nil || IsNil(o.FromCcy) {
		var ret string
		return ret
	}
	return *o.FromCcy
}

// GetFromCcyOk returns a tuple with the FromCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner) GetFromCcyOk() (*string, bool) {
	if o == nil || IsNil(o.FromCcy) {
		return nil, false
	}
	return o.FromCcy, true
}

// HasFromCcy returns a boolean if a field has been set.
func (o *GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner) HasFromCcy() bool {
	if o != nil && !IsNil(o.FromCcy) {
		return true
	}

	return false
}

// SetFromCcy gets a reference to the given string and assigns it to the FromCcy field.
func (o *GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner) SetFromCcy(v string) {
	o.FromCcy = &v
}

func (o GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FromAmt) {
		toSerialize["fromAmt"] = o.FromAmt
	}
	if !IsNil(o.FromCcy) {
		toSerialize["fromCcy"] = o.FromCcy
	}
	return toSerialize, nil
}

type NullableGetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner struct {
	value *GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner
	isSet bool
}

func (v NullableGetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner) Get() *GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner {
	return v.value
}

func (v *NullableGetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner) Set(val *GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner(val *GetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner) *NullableGetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner {
	return &NullableGetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner{value: val, isSet: true}
}

func (v NullableGetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTradeEasyConvertCurrencyListV5RespDataInnerFromDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


