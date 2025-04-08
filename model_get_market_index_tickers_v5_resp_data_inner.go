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

// checks if the GetMarketIndexTickersV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMarketIndexTickersV5RespDataInner{}

// GetMarketIndexTickersV5RespDataInner struct for GetMarketIndexTickersV5RespDataInner
type GetMarketIndexTickersV5RespDataInner struct {
	// Highest price in the past 24 hours
	High24h *string `json:"high24h,omitempty"`
	// Latest index price
	IdxPx *string `json:"idxPx,omitempty"`
	// Index
	InstId *string `json:"instId,omitempty"`
	// Lowest price in the past 24 hours
	Low24h *string `json:"low24h,omitempty"`
	// Open price in the past 24 hours
	Open24h *string `json:"open24h,omitempty"`
	// Open price in the UTC 0
	SodUtc0 *string `json:"sodUtc0,omitempty"`
	// Open price in the UTC 8
	SodUtc8 *string `json:"sodUtc8,omitempty"`
	// Index price update time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	Ts *string `json:"ts,omitempty"`
}

// NewGetMarketIndexTickersV5RespDataInner instantiates a new GetMarketIndexTickersV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarketIndexTickersV5RespDataInner() *GetMarketIndexTickersV5RespDataInner {
	this := GetMarketIndexTickersV5RespDataInner{}
	var high24h string = ""
	this.High24h = &high24h
	var idxPx string = ""
	this.IdxPx = &idxPx
	var instId string = ""
	this.InstId = &instId
	var low24h string = ""
	this.Low24h = &low24h
	var open24h string = ""
	this.Open24h = &open24h
	var sodUtc0 string = ""
	this.SodUtc0 = &sodUtc0
	var sodUtc8 string = ""
	this.SodUtc8 = &sodUtc8
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewGetMarketIndexTickersV5RespDataInnerWithDefaults instantiates a new GetMarketIndexTickersV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarketIndexTickersV5RespDataInnerWithDefaults() *GetMarketIndexTickersV5RespDataInner {
	this := GetMarketIndexTickersV5RespDataInner{}
	var high24h string = ""
	this.High24h = &high24h
	var idxPx string = ""
	this.IdxPx = &idxPx
	var instId string = ""
	this.InstId = &instId
	var low24h string = ""
	this.Low24h = &low24h
	var open24h string = ""
	this.Open24h = &open24h
	var sodUtc0 string = ""
	this.SodUtc0 = &sodUtc0
	var sodUtc8 string = ""
	this.SodUtc8 = &sodUtc8
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetHigh24h returns the High24h field value if set, zero value otherwise.
func (o *GetMarketIndexTickersV5RespDataInner) GetHigh24h() string {
	if o == nil || IsNil(o.High24h) {
		var ret string
		return ret
	}
	return *o.High24h
}

// GetHigh24hOk returns a tuple with the High24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketIndexTickersV5RespDataInner) GetHigh24hOk() (*string, bool) {
	if o == nil || IsNil(o.High24h) {
		return nil, false
	}
	return o.High24h, true
}

// HasHigh24h returns a boolean if a field has been set.
func (o *GetMarketIndexTickersV5RespDataInner) HasHigh24h() bool {
	if o != nil && !IsNil(o.High24h) {
		return true
	}

	return false
}

// SetHigh24h gets a reference to the given string and assigns it to the High24h field.
func (o *GetMarketIndexTickersV5RespDataInner) SetHigh24h(v string) {
	o.High24h = &v
}

// GetIdxPx returns the IdxPx field value if set, zero value otherwise.
func (o *GetMarketIndexTickersV5RespDataInner) GetIdxPx() string {
	if o == nil || IsNil(o.IdxPx) {
		var ret string
		return ret
	}
	return *o.IdxPx
}

// GetIdxPxOk returns a tuple with the IdxPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketIndexTickersV5RespDataInner) GetIdxPxOk() (*string, bool) {
	if o == nil || IsNil(o.IdxPx) {
		return nil, false
	}
	return o.IdxPx, true
}

// HasIdxPx returns a boolean if a field has been set.
func (o *GetMarketIndexTickersV5RespDataInner) HasIdxPx() bool {
	if o != nil && !IsNil(o.IdxPx) {
		return true
	}

	return false
}

// SetIdxPx gets a reference to the given string and assigns it to the IdxPx field.
func (o *GetMarketIndexTickersV5RespDataInner) SetIdxPx(v string) {
	o.IdxPx = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetMarketIndexTickersV5RespDataInner) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketIndexTickersV5RespDataInner) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetMarketIndexTickersV5RespDataInner) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetMarketIndexTickersV5RespDataInner) SetInstId(v string) {
	o.InstId = &v
}

// GetLow24h returns the Low24h field value if set, zero value otherwise.
func (o *GetMarketIndexTickersV5RespDataInner) GetLow24h() string {
	if o == nil || IsNil(o.Low24h) {
		var ret string
		return ret
	}
	return *o.Low24h
}

// GetLow24hOk returns a tuple with the Low24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketIndexTickersV5RespDataInner) GetLow24hOk() (*string, bool) {
	if o == nil || IsNil(o.Low24h) {
		return nil, false
	}
	return o.Low24h, true
}

// HasLow24h returns a boolean if a field has been set.
func (o *GetMarketIndexTickersV5RespDataInner) HasLow24h() bool {
	if o != nil && !IsNil(o.Low24h) {
		return true
	}

	return false
}

// SetLow24h gets a reference to the given string and assigns it to the Low24h field.
func (o *GetMarketIndexTickersV5RespDataInner) SetLow24h(v string) {
	o.Low24h = &v
}

// GetOpen24h returns the Open24h field value if set, zero value otherwise.
func (o *GetMarketIndexTickersV5RespDataInner) GetOpen24h() string {
	if o == nil || IsNil(o.Open24h) {
		var ret string
		return ret
	}
	return *o.Open24h
}

// GetOpen24hOk returns a tuple with the Open24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketIndexTickersV5RespDataInner) GetOpen24hOk() (*string, bool) {
	if o == nil || IsNil(o.Open24h) {
		return nil, false
	}
	return o.Open24h, true
}

// HasOpen24h returns a boolean if a field has been set.
func (o *GetMarketIndexTickersV5RespDataInner) HasOpen24h() bool {
	if o != nil && !IsNil(o.Open24h) {
		return true
	}

	return false
}

// SetOpen24h gets a reference to the given string and assigns it to the Open24h field.
func (o *GetMarketIndexTickersV5RespDataInner) SetOpen24h(v string) {
	o.Open24h = &v
}

// GetSodUtc0 returns the SodUtc0 field value if set, zero value otherwise.
func (o *GetMarketIndexTickersV5RespDataInner) GetSodUtc0() string {
	if o == nil || IsNil(o.SodUtc0) {
		var ret string
		return ret
	}
	return *o.SodUtc0
}

// GetSodUtc0Ok returns a tuple with the SodUtc0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketIndexTickersV5RespDataInner) GetSodUtc0Ok() (*string, bool) {
	if o == nil || IsNil(o.SodUtc0) {
		return nil, false
	}
	return o.SodUtc0, true
}

// HasSodUtc0 returns a boolean if a field has been set.
func (o *GetMarketIndexTickersV5RespDataInner) HasSodUtc0() bool {
	if o != nil && !IsNil(o.SodUtc0) {
		return true
	}

	return false
}

// SetSodUtc0 gets a reference to the given string and assigns it to the SodUtc0 field.
func (o *GetMarketIndexTickersV5RespDataInner) SetSodUtc0(v string) {
	o.SodUtc0 = &v
}

// GetSodUtc8 returns the SodUtc8 field value if set, zero value otherwise.
func (o *GetMarketIndexTickersV5RespDataInner) GetSodUtc8() string {
	if o == nil || IsNil(o.SodUtc8) {
		var ret string
		return ret
	}
	return *o.SodUtc8
}

// GetSodUtc8Ok returns a tuple with the SodUtc8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketIndexTickersV5RespDataInner) GetSodUtc8Ok() (*string, bool) {
	if o == nil || IsNil(o.SodUtc8) {
		return nil, false
	}
	return o.SodUtc8, true
}

// HasSodUtc8 returns a boolean if a field has been set.
func (o *GetMarketIndexTickersV5RespDataInner) HasSodUtc8() bool {
	if o != nil && !IsNil(o.SodUtc8) {
		return true
	}

	return false
}

// SetSodUtc8 gets a reference to the given string and assigns it to the SodUtc8 field.
func (o *GetMarketIndexTickersV5RespDataInner) SetSodUtc8(v string) {
	o.SodUtc8 = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetMarketIndexTickersV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketIndexTickersV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetMarketIndexTickersV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetMarketIndexTickersV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

func (o GetMarketIndexTickersV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarketIndexTickersV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.High24h) {
		toSerialize["high24h"] = o.High24h
	}
	if !IsNil(o.IdxPx) {
		toSerialize["idxPx"] = o.IdxPx
	}
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	if !IsNil(o.Low24h) {
		toSerialize["low24h"] = o.Low24h
	}
	if !IsNil(o.Open24h) {
		toSerialize["open24h"] = o.Open24h
	}
	if !IsNil(o.SodUtc0) {
		toSerialize["sodUtc0"] = o.SodUtc0
	}
	if !IsNil(o.SodUtc8) {
		toSerialize["sodUtc8"] = o.SodUtc8
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableGetMarketIndexTickersV5RespDataInner struct {
	value *GetMarketIndexTickersV5RespDataInner
	isSet bool
}

func (v NullableGetMarketIndexTickersV5RespDataInner) Get() *GetMarketIndexTickersV5RespDataInner {
	return v.value
}

func (v *NullableGetMarketIndexTickersV5RespDataInner) Set(val *GetMarketIndexTickersV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarketIndexTickersV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarketIndexTickersV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarketIndexTickersV5RespDataInner(val *GetMarketIndexTickersV5RespDataInner) *NullableGetMarketIndexTickersV5RespDataInner {
	return &NullableGetMarketIndexTickersV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetMarketIndexTickersV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarketIndexTickersV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


