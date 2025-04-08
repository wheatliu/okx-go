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

// checks if the GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner{}

// GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner struct for GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner
type GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner struct {
	// The Instrument ID
	InstId *string `json:"instId,omitempty"`
	// Trade price
	Px *string `json:"px,omitempty"`
	// Trade side  `buy`  `sell`
	Side *string `json:"side,omitempty"`
	// Trade quantity. The unit is contract.
	Sz *string `json:"sz,omitempty"`
	// Trade ID
	TradeId *string `json:"tradeId,omitempty"`
	// Trade time, Unix timestamp format in milliseconds, e.g. 1597026383085.
	Ts *string `json:"ts,omitempty"`
}

// NewGetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner instantiates a new GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner() *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner {
	this := GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner{}
	var instId string = ""
	this.InstId = &instId
	var px string = ""
	this.Px = &px
	var side string = ""
	this.Side = &side
	var sz string = ""
	this.Sz = &sz
	var tradeId string = ""
	this.TradeId = &tradeId
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewGetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInnerWithDefaults instantiates a new GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInnerWithDefaults() *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner {
	this := GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner{}
	var instId string = ""
	this.InstId = &instId
	var px string = ""
	this.Px = &px
	var side string = ""
	this.Side = &side
	var sz string = ""
	this.Sz = &sz
	var tradeId string = ""
	this.TradeId = &tradeId
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) SetInstId(v string) {
	o.InstId = &v
}

// GetPx returns the Px field value if set, zero value otherwise.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetPx() string {
	if o == nil || IsNil(o.Px) {
		var ret string
		return ret
	}
	return *o.Px
}

// GetPxOk returns a tuple with the Px field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetPxOk() (*string, bool) {
	if o == nil || IsNil(o.Px) {
		return nil, false
	}
	return o.Px, true
}

// HasPx returns a boolean if a field has been set.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) HasPx() bool {
	if o != nil && !IsNil(o.Px) {
		return true
	}

	return false
}

// SetPx gets a reference to the given string and assigns it to the Px field.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) SetPx(v string) {
	o.Px = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) SetSide(v string) {
	o.Side = &v
}

// GetSz returns the Sz field value if set, zero value otherwise.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetSz() string {
	if o == nil || IsNil(o.Sz) {
		var ret string
		return ret
	}
	return *o.Sz
}

// GetSzOk returns a tuple with the Sz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetSzOk() (*string, bool) {
	if o == nil || IsNil(o.Sz) {
		return nil, false
	}
	return o.Sz, true
}

// HasSz returns a boolean if a field has been set.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) HasSz() bool {
	if o != nil && !IsNil(o.Sz) {
		return true
	}

	return false
}

// SetSz gets a reference to the given string and assigns it to the Sz field.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) SetSz(v string) {
	o.Sz = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetTradeId() string {
	if o == nil || IsNil(o.TradeId) {
		var ret string
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetTradeIdOk() (*string, bool) {
	if o == nil || IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) HasTradeId() bool {
	if o != nil && !IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given string and assigns it to the TradeId field.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) SetTradeId(v string) {
	o.TradeId = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) SetTs(v string) {
	o.Ts = &v
}

func (o GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	if !IsNil(o.Px) {
		toSerialize["px"] = o.Px
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !IsNil(o.Sz) {
		toSerialize["sz"] = o.Sz
	}
	if !IsNil(o.TradeId) {
		toSerialize["tradeId"] = o.TradeId
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableGetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner struct {
	value *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner
	isSet bool
}

func (v NullableGetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) Get() *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner {
	return v.value
}

func (v *NullableGetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) Set(val *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner(val *GetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) *NullableGetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner {
	return &NullableGetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner{value: val, isSet: true}
}

func (v NullableGetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarketOptionInstrumentFamilyTradesV5RespDataTradeInfoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


