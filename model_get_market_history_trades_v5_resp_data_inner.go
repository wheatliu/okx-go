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

// checks if the GetMarketHistoryTradesV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMarketHistoryTradesV5RespDataInner{}

// GetMarketHistoryTradesV5RespDataInner struct for GetMarketHistoryTradesV5RespDataInner
type GetMarketHistoryTradesV5RespDataInner struct {
	// Instrument ID
	InstId *string `json:"instId,omitempty"`
	// Trade price
	Px *string `json:"px,omitempty"`
	// Trade side   `buy`    `sell`
	Side *string `json:"side,omitempty"`
	// Trade quantity   For spot trading, the unit is base currency  For `FUTURES`/`SWAP`/`OPTION`, the unit is contract.
	Sz *string `json:"sz,omitempty"`
	// Trade ID
	TradeId *string `json:"tradeId,omitempty"`
	// Trade time, Unix timestamp format in milliseconds, e.g. `1597026383085`.
	Ts *string `json:"ts,omitempty"`
}

// NewGetMarketHistoryTradesV5RespDataInner instantiates a new GetMarketHistoryTradesV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarketHistoryTradesV5RespDataInner() *GetMarketHistoryTradesV5RespDataInner {
	this := GetMarketHistoryTradesV5RespDataInner{}
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

// NewGetMarketHistoryTradesV5RespDataInnerWithDefaults instantiates a new GetMarketHistoryTradesV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarketHistoryTradesV5RespDataInnerWithDefaults() *GetMarketHistoryTradesV5RespDataInner {
	this := GetMarketHistoryTradesV5RespDataInner{}
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
func (o *GetMarketHistoryTradesV5RespDataInner) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketHistoryTradesV5RespDataInner) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetMarketHistoryTradesV5RespDataInner) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetMarketHistoryTradesV5RespDataInner) SetInstId(v string) {
	o.InstId = &v
}

// GetPx returns the Px field value if set, zero value otherwise.
func (o *GetMarketHistoryTradesV5RespDataInner) GetPx() string {
	if o == nil || IsNil(o.Px) {
		var ret string
		return ret
	}
	return *o.Px
}

// GetPxOk returns a tuple with the Px field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketHistoryTradesV5RespDataInner) GetPxOk() (*string, bool) {
	if o == nil || IsNil(o.Px) {
		return nil, false
	}
	return o.Px, true
}

// HasPx returns a boolean if a field has been set.
func (o *GetMarketHistoryTradesV5RespDataInner) HasPx() bool {
	if o != nil && !IsNil(o.Px) {
		return true
	}

	return false
}

// SetPx gets a reference to the given string and assigns it to the Px field.
func (o *GetMarketHistoryTradesV5RespDataInner) SetPx(v string) {
	o.Px = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *GetMarketHistoryTradesV5RespDataInner) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketHistoryTradesV5RespDataInner) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *GetMarketHistoryTradesV5RespDataInner) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *GetMarketHistoryTradesV5RespDataInner) SetSide(v string) {
	o.Side = &v
}

// GetSz returns the Sz field value if set, zero value otherwise.
func (o *GetMarketHistoryTradesV5RespDataInner) GetSz() string {
	if o == nil || IsNil(o.Sz) {
		var ret string
		return ret
	}
	return *o.Sz
}

// GetSzOk returns a tuple with the Sz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketHistoryTradesV5RespDataInner) GetSzOk() (*string, bool) {
	if o == nil || IsNil(o.Sz) {
		return nil, false
	}
	return o.Sz, true
}

// HasSz returns a boolean if a field has been set.
func (o *GetMarketHistoryTradesV5RespDataInner) HasSz() bool {
	if o != nil && !IsNil(o.Sz) {
		return true
	}

	return false
}

// SetSz gets a reference to the given string and assigns it to the Sz field.
func (o *GetMarketHistoryTradesV5RespDataInner) SetSz(v string) {
	o.Sz = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *GetMarketHistoryTradesV5RespDataInner) GetTradeId() string {
	if o == nil || IsNil(o.TradeId) {
		var ret string
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketHistoryTradesV5RespDataInner) GetTradeIdOk() (*string, bool) {
	if o == nil || IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *GetMarketHistoryTradesV5RespDataInner) HasTradeId() bool {
	if o != nil && !IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given string and assigns it to the TradeId field.
func (o *GetMarketHistoryTradesV5RespDataInner) SetTradeId(v string) {
	o.TradeId = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetMarketHistoryTradesV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketHistoryTradesV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetMarketHistoryTradesV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetMarketHistoryTradesV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

func (o GetMarketHistoryTradesV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarketHistoryTradesV5RespDataInner) ToMap() (map[string]interface{}, error) {
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

type NullableGetMarketHistoryTradesV5RespDataInner struct {
	value *GetMarketHistoryTradesV5RespDataInner
	isSet bool
}

func (v NullableGetMarketHistoryTradesV5RespDataInner) Get() *GetMarketHistoryTradesV5RespDataInner {
	return v.value
}

func (v *NullableGetMarketHistoryTradesV5RespDataInner) Set(val *GetMarketHistoryTradesV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarketHistoryTradesV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarketHistoryTradesV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarketHistoryTradesV5RespDataInner(val *GetMarketHistoryTradesV5RespDataInner) *NullableGetMarketHistoryTradesV5RespDataInner {
	return &NullableGetMarketHistoryTradesV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetMarketHistoryTradesV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarketHistoryTradesV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


