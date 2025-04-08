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

// checks if the GetAccountMaxWithdrawalV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountMaxWithdrawalV5RespData{}

// GetAccountMaxWithdrawalV5RespData struct for GetAccountMaxWithdrawalV5RespData
type GetAccountMaxWithdrawalV5RespData struct {
	// Currency
	Ccy *string `json:"ccy,omitempty"`
	// Max withdrawal (excluding borrowed assets under `Spot mode`/`Multi-currency margin`/`Portfolio margin`)
	MaxWd *string `json:"maxWd,omitempty"`
	// Max withdrawal (including borrowed assets under `Spot mode`/`Multi-currency margin`/`Portfolio margin`)
	MaxWdEx *string `json:"maxWdEx,omitempty"`
	// Max withdrawal under Spot-Derivatives risk offset mode (excluding borrowed assets under `Portfolio margin`)  Applicable to `Portfolio margin`
	SpotOffsetMaxWd *string `json:"spotOffsetMaxWd,omitempty"`
	// Max withdrawal under Spot-Derivatives risk offset mode (including borrowed assets under `Portfolio margin`)  Applicable to `Portfolio margin`
	SpotOffsetMaxWdEx *string `json:"spotOffsetMaxWdEx,omitempty"`
}

// NewGetAccountMaxWithdrawalV5RespData instantiates a new GetAccountMaxWithdrawalV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountMaxWithdrawalV5RespData() *GetAccountMaxWithdrawalV5RespData {
	this := GetAccountMaxWithdrawalV5RespData{}
	var ccy string = ""
	this.Ccy = &ccy
	var maxWd string = ""
	this.MaxWd = &maxWd
	var maxWdEx string = ""
	this.MaxWdEx = &maxWdEx
	var spotOffsetMaxWd string = ""
	this.SpotOffsetMaxWd = &spotOffsetMaxWd
	var spotOffsetMaxWdEx string = ""
	this.SpotOffsetMaxWdEx = &spotOffsetMaxWdEx
	return &this
}

// NewGetAccountMaxWithdrawalV5RespDataWithDefaults instantiates a new GetAccountMaxWithdrawalV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountMaxWithdrawalV5RespDataWithDefaults() *GetAccountMaxWithdrawalV5RespData {
	this := GetAccountMaxWithdrawalV5RespData{}
	var ccy string = ""
	this.Ccy = &ccy
	var maxWd string = ""
	this.MaxWd = &maxWd
	var maxWdEx string = ""
	this.MaxWdEx = &maxWdEx
	var spotOffsetMaxWd string = ""
	this.SpotOffsetMaxWd = &spotOffsetMaxWd
	var spotOffsetMaxWdEx string = ""
	this.SpotOffsetMaxWdEx = &spotOffsetMaxWdEx
	return &this
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetAccountMaxWithdrawalV5RespData) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountMaxWithdrawalV5RespData) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetAccountMaxWithdrawalV5RespData) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetAccountMaxWithdrawalV5RespData) SetCcy(v string) {
	o.Ccy = &v
}

// GetMaxWd returns the MaxWd field value if set, zero value otherwise.
func (o *GetAccountMaxWithdrawalV5RespData) GetMaxWd() string {
	if o == nil || IsNil(o.MaxWd) {
		var ret string
		return ret
	}
	return *o.MaxWd
}

// GetMaxWdOk returns a tuple with the MaxWd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountMaxWithdrawalV5RespData) GetMaxWdOk() (*string, bool) {
	if o == nil || IsNil(o.MaxWd) {
		return nil, false
	}
	return o.MaxWd, true
}

// HasMaxWd returns a boolean if a field has been set.
func (o *GetAccountMaxWithdrawalV5RespData) HasMaxWd() bool {
	if o != nil && !IsNil(o.MaxWd) {
		return true
	}

	return false
}

// SetMaxWd gets a reference to the given string and assigns it to the MaxWd field.
func (o *GetAccountMaxWithdrawalV5RespData) SetMaxWd(v string) {
	o.MaxWd = &v
}

// GetMaxWdEx returns the MaxWdEx field value if set, zero value otherwise.
func (o *GetAccountMaxWithdrawalV5RespData) GetMaxWdEx() string {
	if o == nil || IsNil(o.MaxWdEx) {
		var ret string
		return ret
	}
	return *o.MaxWdEx
}

// GetMaxWdExOk returns a tuple with the MaxWdEx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountMaxWithdrawalV5RespData) GetMaxWdExOk() (*string, bool) {
	if o == nil || IsNil(o.MaxWdEx) {
		return nil, false
	}
	return o.MaxWdEx, true
}

// HasMaxWdEx returns a boolean if a field has been set.
func (o *GetAccountMaxWithdrawalV5RespData) HasMaxWdEx() bool {
	if o != nil && !IsNil(o.MaxWdEx) {
		return true
	}

	return false
}

// SetMaxWdEx gets a reference to the given string and assigns it to the MaxWdEx field.
func (o *GetAccountMaxWithdrawalV5RespData) SetMaxWdEx(v string) {
	o.MaxWdEx = &v
}

// GetSpotOffsetMaxWd returns the SpotOffsetMaxWd field value if set, zero value otherwise.
func (o *GetAccountMaxWithdrawalV5RespData) GetSpotOffsetMaxWd() string {
	if o == nil || IsNil(o.SpotOffsetMaxWd) {
		var ret string
		return ret
	}
	return *o.SpotOffsetMaxWd
}

// GetSpotOffsetMaxWdOk returns a tuple with the SpotOffsetMaxWd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountMaxWithdrawalV5RespData) GetSpotOffsetMaxWdOk() (*string, bool) {
	if o == nil || IsNil(o.SpotOffsetMaxWd) {
		return nil, false
	}
	return o.SpotOffsetMaxWd, true
}

// HasSpotOffsetMaxWd returns a boolean if a field has been set.
func (o *GetAccountMaxWithdrawalV5RespData) HasSpotOffsetMaxWd() bool {
	if o != nil && !IsNil(o.SpotOffsetMaxWd) {
		return true
	}

	return false
}

// SetSpotOffsetMaxWd gets a reference to the given string and assigns it to the SpotOffsetMaxWd field.
func (o *GetAccountMaxWithdrawalV5RespData) SetSpotOffsetMaxWd(v string) {
	o.SpotOffsetMaxWd = &v
}

// GetSpotOffsetMaxWdEx returns the SpotOffsetMaxWdEx field value if set, zero value otherwise.
func (o *GetAccountMaxWithdrawalV5RespData) GetSpotOffsetMaxWdEx() string {
	if o == nil || IsNil(o.SpotOffsetMaxWdEx) {
		var ret string
		return ret
	}
	return *o.SpotOffsetMaxWdEx
}

// GetSpotOffsetMaxWdExOk returns a tuple with the SpotOffsetMaxWdEx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountMaxWithdrawalV5RespData) GetSpotOffsetMaxWdExOk() (*string, bool) {
	if o == nil || IsNil(o.SpotOffsetMaxWdEx) {
		return nil, false
	}
	return o.SpotOffsetMaxWdEx, true
}

// HasSpotOffsetMaxWdEx returns a boolean if a field has been set.
func (o *GetAccountMaxWithdrawalV5RespData) HasSpotOffsetMaxWdEx() bool {
	if o != nil && !IsNil(o.SpotOffsetMaxWdEx) {
		return true
	}

	return false
}

// SetSpotOffsetMaxWdEx gets a reference to the given string and assigns it to the SpotOffsetMaxWdEx field.
func (o *GetAccountMaxWithdrawalV5RespData) SetSpotOffsetMaxWdEx(v string) {
	o.SpotOffsetMaxWdEx = &v
}

func (o GetAccountMaxWithdrawalV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountMaxWithdrawalV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.MaxWd) {
		toSerialize["maxWd"] = o.MaxWd
	}
	if !IsNil(o.MaxWdEx) {
		toSerialize["maxWdEx"] = o.MaxWdEx
	}
	if !IsNil(o.SpotOffsetMaxWd) {
		toSerialize["spotOffsetMaxWd"] = o.SpotOffsetMaxWd
	}
	if !IsNil(o.SpotOffsetMaxWdEx) {
		toSerialize["spotOffsetMaxWdEx"] = o.SpotOffsetMaxWdEx
	}
	return toSerialize, nil
}

type NullableGetAccountMaxWithdrawalV5RespData struct {
	value *GetAccountMaxWithdrawalV5RespData
	isSet bool
}

func (v NullableGetAccountMaxWithdrawalV5RespData) Get() *GetAccountMaxWithdrawalV5RespData {
	return v.value
}

func (v *NullableGetAccountMaxWithdrawalV5RespData) Set(val *GetAccountMaxWithdrawalV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountMaxWithdrawalV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountMaxWithdrawalV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountMaxWithdrawalV5RespData(val *GetAccountMaxWithdrawalV5RespData) *NullableGetAccountMaxWithdrawalV5RespData {
	return &NullableGetAccountMaxWithdrawalV5RespData{value: val, isSet: true}
}

func (v NullableGetAccountMaxWithdrawalV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountMaxWithdrawalV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


