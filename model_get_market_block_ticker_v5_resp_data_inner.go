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

// checks if the GetMarketBlockTickerV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMarketBlockTickerV5RespDataInner{}

// GetMarketBlockTickerV5RespDataInner struct for GetMarketBlockTickerV5RespDataInner
type GetMarketBlockTickerV5RespDataInner struct {
	// Instrument ID
	InstId *string `json:"instId,omitempty"`
	// Instrument type
	InstType *string `json:"instType,omitempty"`
	// Block ticker data generation time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	Ts *string `json:"ts,omitempty"`
	// 24h trading volume, with a unit of `contract`.   If it is a `derivatives` contract, the value is the number of contracts.   If it is `SPOT`/`MARGIN`, the value is the quantity in base currency.
	Vol24h *string `json:"vol24h,omitempty"`
	// 24h trading volume, with a unit of `currency`.   If it is a `derivatives` contract, the value is the number of base currency.   If it is `SPOT`/`MARGIN`, the value is the quantity in quote currency.
	VolCcy24h *string `json:"volCcy24h,omitempty"`
}

// NewGetMarketBlockTickerV5RespDataInner instantiates a new GetMarketBlockTickerV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarketBlockTickerV5RespDataInner() *GetMarketBlockTickerV5RespDataInner {
	this := GetMarketBlockTickerV5RespDataInner{}
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var ts string = ""
	this.Ts = &ts
	var vol24h string = ""
	this.Vol24h = &vol24h
	var volCcy24h string = ""
	this.VolCcy24h = &volCcy24h
	return &this
}

// NewGetMarketBlockTickerV5RespDataInnerWithDefaults instantiates a new GetMarketBlockTickerV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarketBlockTickerV5RespDataInnerWithDefaults() *GetMarketBlockTickerV5RespDataInner {
	this := GetMarketBlockTickerV5RespDataInner{}
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var ts string = ""
	this.Ts = &ts
	var vol24h string = ""
	this.Vol24h = &vol24h
	var volCcy24h string = ""
	this.VolCcy24h = &volCcy24h
	return &this
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetMarketBlockTickerV5RespDataInner) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketBlockTickerV5RespDataInner) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetMarketBlockTickerV5RespDataInner) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetMarketBlockTickerV5RespDataInner) SetInstId(v string) {
	o.InstId = &v
}

// GetInstType returns the InstType field value if set, zero value otherwise.
func (o *GetMarketBlockTickerV5RespDataInner) GetInstType() string {
	if o == nil || IsNil(o.InstType) {
		var ret string
		return ret
	}
	return *o.InstType
}

// GetInstTypeOk returns a tuple with the InstType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketBlockTickerV5RespDataInner) GetInstTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstType) {
		return nil, false
	}
	return o.InstType, true
}

// HasInstType returns a boolean if a field has been set.
func (o *GetMarketBlockTickerV5RespDataInner) HasInstType() bool {
	if o != nil && !IsNil(o.InstType) {
		return true
	}

	return false
}

// SetInstType gets a reference to the given string and assigns it to the InstType field.
func (o *GetMarketBlockTickerV5RespDataInner) SetInstType(v string) {
	o.InstType = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetMarketBlockTickerV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketBlockTickerV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetMarketBlockTickerV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetMarketBlockTickerV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

// GetVol24h returns the Vol24h field value if set, zero value otherwise.
func (o *GetMarketBlockTickerV5RespDataInner) GetVol24h() string {
	if o == nil || IsNil(o.Vol24h) {
		var ret string
		return ret
	}
	return *o.Vol24h
}

// GetVol24hOk returns a tuple with the Vol24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketBlockTickerV5RespDataInner) GetVol24hOk() (*string, bool) {
	if o == nil || IsNil(o.Vol24h) {
		return nil, false
	}
	return o.Vol24h, true
}

// HasVol24h returns a boolean if a field has been set.
func (o *GetMarketBlockTickerV5RespDataInner) HasVol24h() bool {
	if o != nil && !IsNil(o.Vol24h) {
		return true
	}

	return false
}

// SetVol24h gets a reference to the given string and assigns it to the Vol24h field.
func (o *GetMarketBlockTickerV5RespDataInner) SetVol24h(v string) {
	o.Vol24h = &v
}

// GetVolCcy24h returns the VolCcy24h field value if set, zero value otherwise.
func (o *GetMarketBlockTickerV5RespDataInner) GetVolCcy24h() string {
	if o == nil || IsNil(o.VolCcy24h) {
		var ret string
		return ret
	}
	return *o.VolCcy24h
}

// GetVolCcy24hOk returns a tuple with the VolCcy24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketBlockTickerV5RespDataInner) GetVolCcy24hOk() (*string, bool) {
	if o == nil || IsNil(o.VolCcy24h) {
		return nil, false
	}
	return o.VolCcy24h, true
}

// HasVolCcy24h returns a boolean if a field has been set.
func (o *GetMarketBlockTickerV5RespDataInner) HasVolCcy24h() bool {
	if o != nil && !IsNil(o.VolCcy24h) {
		return true
	}

	return false
}

// SetVolCcy24h gets a reference to the given string and assigns it to the VolCcy24h field.
func (o *GetMarketBlockTickerV5RespDataInner) SetVolCcy24h(v string) {
	o.VolCcy24h = &v
}

func (o GetMarketBlockTickerV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarketBlockTickerV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	if !IsNil(o.InstType) {
		toSerialize["instType"] = o.InstType
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.Vol24h) {
		toSerialize["vol24h"] = o.Vol24h
	}
	if !IsNil(o.VolCcy24h) {
		toSerialize["volCcy24h"] = o.VolCcy24h
	}
	return toSerialize, nil
}

type NullableGetMarketBlockTickerV5RespDataInner struct {
	value *GetMarketBlockTickerV5RespDataInner
	isSet bool
}

func (v NullableGetMarketBlockTickerV5RespDataInner) Get() *GetMarketBlockTickerV5RespDataInner {
	return v.value
}

func (v *NullableGetMarketBlockTickerV5RespDataInner) Set(val *GetMarketBlockTickerV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarketBlockTickerV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarketBlockTickerV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarketBlockTickerV5RespDataInner(val *GetMarketBlockTickerV5RespDataInner) *NullableGetMarketBlockTickerV5RespDataInner {
	return &NullableGetMarketBlockTickerV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetMarketBlockTickerV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarketBlockTickerV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


