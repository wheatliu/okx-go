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

// checks if the GetRfqRfqsV5RespDataInnerLegsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRfqRfqsV5RespDataInnerLegsInner{}

// GetRfqRfqsV5RespDataInnerLegsInner struct for GetRfqRfqsV5RespDataInnerLegsInner
type GetRfqRfqsV5RespDataInnerLegsInner struct {
	// Margin currency.   Only applicable to `cross` `MARGIN` orders in `Spot and futures mode`. The parameter will be ignored in other scenarios.
	Ccy *string `json:"ccy,omitempty"`
	// Instrument ID, e.g. BTC-USDT-SWAP
	InstId *string `json:"instId,omitempty"`
	// Position side.   The default is `net` in the net mode. If not specified, return \"\", which is equivalent to net.   It can only be `long` or `short` in the long/short mode. If not specified, return \"\", which corresponds to the direction that opens new positions for the trade (buy => long, sell => short).   Only applicable to `FUTURES`/`SWAP`.
	PosSide *string `json:"posSide,omitempty"`
	// The direction of the leg. Valid values can be buy or sell.
	Side *string `json:"side,omitempty"`
	// Size of the leg in contracts or spot.
	Sz *string `json:"sz,omitempty"`
	// Trade mode   Margin mode: `cross` `isolated`   Non-Margin mode: `cash`.   If not provided, tdMode will inherit default values set by the system shown below:   Spot and futures mode & SPOT: `cash`   Buy options in Spot and futures mode and Multi-currency Margin: `isolated`   Other cases: `cross`
	TdMode *string `json:"tdMode,omitempty"`
	// Defines the unit of the “sz” attribute.   Only applicable to instType = SPOT.   The valid enumerations are base_ccy and quote_ccy. When not specified this is equal to base_ccy by default.
	TgtCcy *string `json:"tgtCcy,omitempty"`
}

// NewGetRfqRfqsV5RespDataInnerLegsInner instantiates a new GetRfqRfqsV5RespDataInnerLegsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRfqRfqsV5RespDataInnerLegsInner() *GetRfqRfqsV5RespDataInnerLegsInner {
	this := GetRfqRfqsV5RespDataInnerLegsInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var instId string = ""
	this.InstId = &instId
	var posSide string = ""
	this.PosSide = &posSide
	var side string = ""
	this.Side = &side
	var sz string = ""
	this.Sz = &sz
	var tdMode string = ""
	this.TdMode = &tdMode
	var tgtCcy string = ""
	this.TgtCcy = &tgtCcy
	return &this
}

// NewGetRfqRfqsV5RespDataInnerLegsInnerWithDefaults instantiates a new GetRfqRfqsV5RespDataInnerLegsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRfqRfqsV5RespDataInnerLegsInnerWithDefaults() *GetRfqRfqsV5RespDataInnerLegsInner {
	this := GetRfqRfqsV5RespDataInnerLegsInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var instId string = ""
	this.InstId = &instId
	var posSide string = ""
	this.PosSide = &posSide
	var side string = ""
	this.Side = &side
	var sz string = ""
	this.Sz = &sz
	var tdMode string = ""
	this.TdMode = &tdMode
	var tgtCcy string = ""
	this.TgtCcy = &tgtCcy
	return &this
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) SetInstId(v string) {
	o.InstId = &v
}

// GetPosSide returns the PosSide field value if set, zero value otherwise.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) GetPosSide() string {
	if o == nil || IsNil(o.PosSide) {
		var ret string
		return ret
	}
	return *o.PosSide
}

// GetPosSideOk returns a tuple with the PosSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) GetPosSideOk() (*string, bool) {
	if o == nil || IsNil(o.PosSide) {
		return nil, false
	}
	return o.PosSide, true
}

// HasPosSide returns a boolean if a field has been set.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) HasPosSide() bool {
	if o != nil && !IsNil(o.PosSide) {
		return true
	}

	return false
}

// SetPosSide gets a reference to the given string and assigns it to the PosSide field.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) SetPosSide(v string) {
	o.PosSide = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) SetSide(v string) {
	o.Side = &v
}

// GetSz returns the Sz field value if set, zero value otherwise.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) GetSz() string {
	if o == nil || IsNil(o.Sz) {
		var ret string
		return ret
	}
	return *o.Sz
}

// GetSzOk returns a tuple with the Sz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) GetSzOk() (*string, bool) {
	if o == nil || IsNil(o.Sz) {
		return nil, false
	}
	return o.Sz, true
}

// HasSz returns a boolean if a field has been set.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) HasSz() bool {
	if o != nil && !IsNil(o.Sz) {
		return true
	}

	return false
}

// SetSz gets a reference to the given string and assigns it to the Sz field.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) SetSz(v string) {
	o.Sz = &v
}

// GetTdMode returns the TdMode field value if set, zero value otherwise.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) GetTdMode() string {
	if o == nil || IsNil(o.TdMode) {
		var ret string
		return ret
	}
	return *o.TdMode
}

// GetTdModeOk returns a tuple with the TdMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) GetTdModeOk() (*string, bool) {
	if o == nil || IsNil(o.TdMode) {
		return nil, false
	}
	return o.TdMode, true
}

// HasTdMode returns a boolean if a field has been set.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) HasTdMode() bool {
	if o != nil && !IsNil(o.TdMode) {
		return true
	}

	return false
}

// SetTdMode gets a reference to the given string and assigns it to the TdMode field.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) SetTdMode(v string) {
	o.TdMode = &v
}

// GetTgtCcy returns the TgtCcy field value if set, zero value otherwise.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) GetTgtCcy() string {
	if o == nil || IsNil(o.TgtCcy) {
		var ret string
		return ret
	}
	return *o.TgtCcy
}

// GetTgtCcyOk returns a tuple with the TgtCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) GetTgtCcyOk() (*string, bool) {
	if o == nil || IsNil(o.TgtCcy) {
		return nil, false
	}
	return o.TgtCcy, true
}

// HasTgtCcy returns a boolean if a field has been set.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) HasTgtCcy() bool {
	if o != nil && !IsNil(o.TgtCcy) {
		return true
	}

	return false
}

// SetTgtCcy gets a reference to the given string and assigns it to the TgtCcy field.
func (o *GetRfqRfqsV5RespDataInnerLegsInner) SetTgtCcy(v string) {
	o.TgtCcy = &v
}

func (o GetRfqRfqsV5RespDataInnerLegsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRfqRfqsV5RespDataInnerLegsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	if !IsNil(o.PosSide) {
		toSerialize["posSide"] = o.PosSide
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !IsNil(o.Sz) {
		toSerialize["sz"] = o.Sz
	}
	if !IsNil(o.TdMode) {
		toSerialize["tdMode"] = o.TdMode
	}
	if !IsNil(o.TgtCcy) {
		toSerialize["tgtCcy"] = o.TgtCcy
	}
	return toSerialize, nil
}

type NullableGetRfqRfqsV5RespDataInnerLegsInner struct {
	value *GetRfqRfqsV5RespDataInnerLegsInner
	isSet bool
}

func (v NullableGetRfqRfqsV5RespDataInnerLegsInner) Get() *GetRfqRfqsV5RespDataInnerLegsInner {
	return v.value
}

func (v *NullableGetRfqRfqsV5RespDataInnerLegsInner) Set(val *GetRfqRfqsV5RespDataInnerLegsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRfqRfqsV5RespDataInnerLegsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRfqRfqsV5RespDataInnerLegsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRfqRfqsV5RespDataInnerLegsInner(val *GetRfqRfqsV5RespDataInnerLegsInner) *NullableGetRfqRfqsV5RespDataInnerLegsInner {
	return &NullableGetRfqRfqsV5RespDataInnerLegsInner{value: val, isSet: true}
}

func (v NullableGetRfqRfqsV5RespDataInnerLegsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRfqRfqsV5RespDataInnerLegsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


