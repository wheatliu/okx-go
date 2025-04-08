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

// checks if the GetAssetConvertCurrencyPairV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAssetConvertCurrencyPairV5RespData{}

// GetAssetConvertCurrencyPairV5RespData struct for GetAssetConvertCurrencyPairV5RespData
type GetAssetConvertCurrencyPairV5RespData struct {
	// Currency to convert from, e.g. `USDT`
	FromCcy *string `json:"fromCcy,omitempty"`
	// Currency to convert to, e.g. `BTC`
	ToCcy *string `json:"toCcy,omitempty"`
}

// NewGetAssetConvertCurrencyPairV5RespData instantiates a new GetAssetConvertCurrencyPairV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAssetConvertCurrencyPairV5RespData() *GetAssetConvertCurrencyPairV5RespData {
	this := GetAssetConvertCurrencyPairV5RespData{}
	var fromCcy string = ""
	this.FromCcy = &fromCcy
	var toCcy string = ""
	this.ToCcy = &toCcy
	return &this
}

// NewGetAssetConvertCurrencyPairV5RespDataWithDefaults instantiates a new GetAssetConvertCurrencyPairV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAssetConvertCurrencyPairV5RespDataWithDefaults() *GetAssetConvertCurrencyPairV5RespData {
	this := GetAssetConvertCurrencyPairV5RespData{}
	var fromCcy string = ""
	this.FromCcy = &fromCcy
	var toCcy string = ""
	this.ToCcy = &toCcy
	return &this
}

// GetFromCcy returns the FromCcy field value if set, zero value otherwise.
func (o *GetAssetConvertCurrencyPairV5RespData) GetFromCcy() string {
	if o == nil || IsNil(o.FromCcy) {
		var ret string
		return ret
	}
	return *o.FromCcy
}

// GetFromCcyOk returns a tuple with the FromCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetConvertCurrencyPairV5RespData) GetFromCcyOk() (*string, bool) {
	if o == nil || IsNil(o.FromCcy) {
		return nil, false
	}
	return o.FromCcy, true
}

// HasFromCcy returns a boolean if a field has been set.
func (o *GetAssetConvertCurrencyPairV5RespData) HasFromCcy() bool {
	if o != nil && !IsNil(o.FromCcy) {
		return true
	}

	return false
}

// SetFromCcy gets a reference to the given string and assigns it to the FromCcy field.
func (o *GetAssetConvertCurrencyPairV5RespData) SetFromCcy(v string) {
	o.FromCcy = &v
}

// GetToCcy returns the ToCcy field value if set, zero value otherwise.
func (o *GetAssetConvertCurrencyPairV5RespData) GetToCcy() string {
	if o == nil || IsNil(o.ToCcy) {
		var ret string
		return ret
	}
	return *o.ToCcy
}

// GetToCcyOk returns a tuple with the ToCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetConvertCurrencyPairV5RespData) GetToCcyOk() (*string, bool) {
	if o == nil || IsNil(o.ToCcy) {
		return nil, false
	}
	return o.ToCcy, true
}

// HasToCcy returns a boolean if a field has been set.
func (o *GetAssetConvertCurrencyPairV5RespData) HasToCcy() bool {
	if o != nil && !IsNil(o.ToCcy) {
		return true
	}

	return false
}

// SetToCcy gets a reference to the given string and assigns it to the ToCcy field.
func (o *GetAssetConvertCurrencyPairV5RespData) SetToCcy(v string) {
	o.ToCcy = &v
}

func (o GetAssetConvertCurrencyPairV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAssetConvertCurrencyPairV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FromCcy) {
		toSerialize["fromCcy"] = o.FromCcy
	}
	if !IsNil(o.ToCcy) {
		toSerialize["toCcy"] = o.ToCcy
	}
	return toSerialize, nil
}

type NullableGetAssetConvertCurrencyPairV5RespData struct {
	value *GetAssetConvertCurrencyPairV5RespData
	isSet bool
}

func (v NullableGetAssetConvertCurrencyPairV5RespData) Get() *GetAssetConvertCurrencyPairV5RespData {
	return v.value
}

func (v *NullableGetAssetConvertCurrencyPairV5RespData) Set(val *GetAssetConvertCurrencyPairV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAssetConvertCurrencyPairV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAssetConvertCurrencyPairV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAssetConvertCurrencyPairV5RespData(val *GetAssetConvertCurrencyPairV5RespData) *NullableGetAssetConvertCurrencyPairV5RespData {
	return &NullableGetAssetConvertCurrencyPairV5RespData{value: val, isSet: true}
}

func (v NullableGetAssetConvertCurrencyPairV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAssetConvertCurrencyPairV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


