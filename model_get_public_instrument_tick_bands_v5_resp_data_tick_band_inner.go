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

// checks if the GetPublicInstrumentTickBandsV5RespDataTickBandInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPublicInstrumentTickBandsV5RespDataTickBandInner{}

// GetPublicInstrumentTickBandsV5RespDataTickBandInner struct for GetPublicInstrumentTickBandsV5RespDataTickBandInner
type GetPublicInstrumentTickBandsV5RespDataTickBandInner struct {
	// Maximum price while placing an order
	MaxPx *string `json:"maxPx,omitempty"`
	// Minimum price while placing an order
	MinPx *string `json:"minPx,omitempty"`
	// Tick size, e.g. 0.0001
	TickSz *string `json:"tickSz,omitempty"`
}

// NewGetPublicInstrumentTickBandsV5RespDataTickBandInner instantiates a new GetPublicInstrumentTickBandsV5RespDataTickBandInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPublicInstrumentTickBandsV5RespDataTickBandInner() *GetPublicInstrumentTickBandsV5RespDataTickBandInner {
	this := GetPublicInstrumentTickBandsV5RespDataTickBandInner{}
	var maxPx string = ""
	this.MaxPx = &maxPx
	var minPx string = ""
	this.MinPx = &minPx
	var tickSz string = ""
	this.TickSz = &tickSz
	return &this
}

// NewGetPublicInstrumentTickBandsV5RespDataTickBandInnerWithDefaults instantiates a new GetPublicInstrumentTickBandsV5RespDataTickBandInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPublicInstrumentTickBandsV5RespDataTickBandInnerWithDefaults() *GetPublicInstrumentTickBandsV5RespDataTickBandInner {
	this := GetPublicInstrumentTickBandsV5RespDataTickBandInner{}
	var maxPx string = ""
	this.MaxPx = &maxPx
	var minPx string = ""
	this.MinPx = &minPx
	var tickSz string = ""
	this.TickSz = &tickSz
	return &this
}

// GetMaxPx returns the MaxPx field value if set, zero value otherwise.
func (o *GetPublicInstrumentTickBandsV5RespDataTickBandInner) GetMaxPx() string {
	if o == nil || IsNil(o.MaxPx) {
		var ret string
		return ret
	}
	return *o.MaxPx
}

// GetMaxPxOk returns a tuple with the MaxPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicInstrumentTickBandsV5RespDataTickBandInner) GetMaxPxOk() (*string, bool) {
	if o == nil || IsNil(o.MaxPx) {
		return nil, false
	}
	return o.MaxPx, true
}

// HasMaxPx returns a boolean if a field has been set.
func (o *GetPublicInstrumentTickBandsV5RespDataTickBandInner) HasMaxPx() bool {
	if o != nil && !IsNil(o.MaxPx) {
		return true
	}

	return false
}

// SetMaxPx gets a reference to the given string and assigns it to the MaxPx field.
func (o *GetPublicInstrumentTickBandsV5RespDataTickBandInner) SetMaxPx(v string) {
	o.MaxPx = &v
}

// GetMinPx returns the MinPx field value if set, zero value otherwise.
func (o *GetPublicInstrumentTickBandsV5RespDataTickBandInner) GetMinPx() string {
	if o == nil || IsNil(o.MinPx) {
		var ret string
		return ret
	}
	return *o.MinPx
}

// GetMinPxOk returns a tuple with the MinPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicInstrumentTickBandsV5RespDataTickBandInner) GetMinPxOk() (*string, bool) {
	if o == nil || IsNil(o.MinPx) {
		return nil, false
	}
	return o.MinPx, true
}

// HasMinPx returns a boolean if a field has been set.
func (o *GetPublicInstrumentTickBandsV5RespDataTickBandInner) HasMinPx() bool {
	if o != nil && !IsNil(o.MinPx) {
		return true
	}

	return false
}

// SetMinPx gets a reference to the given string and assigns it to the MinPx field.
func (o *GetPublicInstrumentTickBandsV5RespDataTickBandInner) SetMinPx(v string) {
	o.MinPx = &v
}

// GetTickSz returns the TickSz field value if set, zero value otherwise.
func (o *GetPublicInstrumentTickBandsV5RespDataTickBandInner) GetTickSz() string {
	if o == nil || IsNil(o.TickSz) {
		var ret string
		return ret
	}
	return *o.TickSz
}

// GetTickSzOk returns a tuple with the TickSz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicInstrumentTickBandsV5RespDataTickBandInner) GetTickSzOk() (*string, bool) {
	if o == nil || IsNil(o.TickSz) {
		return nil, false
	}
	return o.TickSz, true
}

// HasTickSz returns a boolean if a field has been set.
func (o *GetPublicInstrumentTickBandsV5RespDataTickBandInner) HasTickSz() bool {
	if o != nil && !IsNil(o.TickSz) {
		return true
	}

	return false
}

// SetTickSz gets a reference to the given string and assigns it to the TickSz field.
func (o *GetPublicInstrumentTickBandsV5RespDataTickBandInner) SetTickSz(v string) {
	o.TickSz = &v
}

func (o GetPublicInstrumentTickBandsV5RespDataTickBandInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPublicInstrumentTickBandsV5RespDataTickBandInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxPx) {
		toSerialize["maxPx"] = o.MaxPx
	}
	if !IsNil(o.MinPx) {
		toSerialize["minPx"] = o.MinPx
	}
	if !IsNil(o.TickSz) {
		toSerialize["tickSz"] = o.TickSz
	}
	return toSerialize, nil
}

type NullableGetPublicInstrumentTickBandsV5RespDataTickBandInner struct {
	value *GetPublicInstrumentTickBandsV5RespDataTickBandInner
	isSet bool
}

func (v NullableGetPublicInstrumentTickBandsV5RespDataTickBandInner) Get() *GetPublicInstrumentTickBandsV5RespDataTickBandInner {
	return v.value
}

func (v *NullableGetPublicInstrumentTickBandsV5RespDataTickBandInner) Set(val *GetPublicInstrumentTickBandsV5RespDataTickBandInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPublicInstrumentTickBandsV5RespDataTickBandInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPublicInstrumentTickBandsV5RespDataTickBandInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPublicInstrumentTickBandsV5RespDataTickBandInner(val *GetPublicInstrumentTickBandsV5RespDataTickBandInner) *NullableGetPublicInstrumentTickBandsV5RespDataTickBandInner {
	return &NullableGetPublicInstrumentTickBandsV5RespDataTickBandInner{value: val, isSet: true}
}

func (v NullableGetPublicInstrumentTickBandsV5RespDataTickBandInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPublicInstrumentTickBandsV5RespDataTickBandInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


