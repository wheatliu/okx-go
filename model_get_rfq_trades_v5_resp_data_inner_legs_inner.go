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

// checks if the GetRfqTradesV5RespDataInnerLegsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRfqTradesV5RespDataInnerLegsInner{}

// GetRfqTradesV5RespDataInnerLegsInner struct for GetRfqTradesV5RespDataInnerLegsInner
type GetRfqTradesV5RespDataInnerLegsInner struct {
	// Fee. Negative number represents the user transaction fee charged by the platform. Positive number represents rebate.
	Fee *string `json:"fee,omitempty"`
	// Fee currency
	FeeCcy *string `json:"feeCcy,omitempty"`
	// Instrument ID, e.g. `BTC-USDT-SWAP`
	InstId *string `json:"instId,omitempty"`
	// The price the leg executed
	Px *string `json:"px,omitempty"`
	// The direction of the leg. Valid value can be buy or sell.
	Side *string `json:"side,omitempty"`
	// Size of the leg in contracts or spot.
	Sz *string `json:"sz,omitempty"`
	// Last traded ID.
	TradeId *string `json:"tradeId,omitempty"`
}

// NewGetRfqTradesV5RespDataInnerLegsInner instantiates a new GetRfqTradesV5RespDataInnerLegsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRfqTradesV5RespDataInnerLegsInner() *GetRfqTradesV5RespDataInnerLegsInner {
	this := GetRfqTradesV5RespDataInnerLegsInner{}
	var fee string = ""
	this.Fee = &fee
	var feeCcy string = ""
	this.FeeCcy = &feeCcy
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
	return &this
}

// NewGetRfqTradesV5RespDataInnerLegsInnerWithDefaults instantiates a new GetRfqTradesV5RespDataInnerLegsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRfqTradesV5RespDataInnerLegsInnerWithDefaults() *GetRfqTradesV5RespDataInnerLegsInner {
	this := GetRfqTradesV5RespDataInnerLegsInner{}
	var fee string = ""
	this.Fee = &fee
	var feeCcy string = ""
	this.FeeCcy = &feeCcy
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
	return &this
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInnerLegsInner) GetFee() string {
	if o == nil || IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInnerLegsInner) GetFeeOk() (*string, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInnerLegsInner) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *GetRfqTradesV5RespDataInnerLegsInner) SetFee(v string) {
	o.Fee = &v
}

// GetFeeCcy returns the FeeCcy field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInnerLegsInner) GetFeeCcy() string {
	if o == nil || IsNil(o.FeeCcy) {
		var ret string
		return ret
	}
	return *o.FeeCcy
}

// GetFeeCcyOk returns a tuple with the FeeCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInnerLegsInner) GetFeeCcyOk() (*string, bool) {
	if o == nil || IsNil(o.FeeCcy) {
		return nil, false
	}
	return o.FeeCcy, true
}

// HasFeeCcy returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInnerLegsInner) HasFeeCcy() bool {
	if o != nil && !IsNil(o.FeeCcy) {
		return true
	}

	return false
}

// SetFeeCcy gets a reference to the given string and assigns it to the FeeCcy field.
func (o *GetRfqTradesV5RespDataInnerLegsInner) SetFeeCcy(v string) {
	o.FeeCcy = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInnerLegsInner) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInnerLegsInner) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInnerLegsInner) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetRfqTradesV5RespDataInnerLegsInner) SetInstId(v string) {
	o.InstId = &v
}

// GetPx returns the Px field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInnerLegsInner) GetPx() string {
	if o == nil || IsNil(o.Px) {
		var ret string
		return ret
	}
	return *o.Px
}

// GetPxOk returns a tuple with the Px field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInnerLegsInner) GetPxOk() (*string, bool) {
	if o == nil || IsNil(o.Px) {
		return nil, false
	}
	return o.Px, true
}

// HasPx returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInnerLegsInner) HasPx() bool {
	if o != nil && !IsNil(o.Px) {
		return true
	}

	return false
}

// SetPx gets a reference to the given string and assigns it to the Px field.
func (o *GetRfqTradesV5RespDataInnerLegsInner) SetPx(v string) {
	o.Px = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInnerLegsInner) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInnerLegsInner) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInnerLegsInner) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *GetRfqTradesV5RespDataInnerLegsInner) SetSide(v string) {
	o.Side = &v
}

// GetSz returns the Sz field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInnerLegsInner) GetSz() string {
	if o == nil || IsNil(o.Sz) {
		var ret string
		return ret
	}
	return *o.Sz
}

// GetSzOk returns a tuple with the Sz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInnerLegsInner) GetSzOk() (*string, bool) {
	if o == nil || IsNil(o.Sz) {
		return nil, false
	}
	return o.Sz, true
}

// HasSz returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInnerLegsInner) HasSz() bool {
	if o != nil && !IsNil(o.Sz) {
		return true
	}

	return false
}

// SetSz gets a reference to the given string and assigns it to the Sz field.
func (o *GetRfqTradesV5RespDataInnerLegsInner) SetSz(v string) {
	o.Sz = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *GetRfqTradesV5RespDataInnerLegsInner) GetTradeId() string {
	if o == nil || IsNil(o.TradeId) {
		var ret string
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqTradesV5RespDataInnerLegsInner) GetTradeIdOk() (*string, bool) {
	if o == nil || IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *GetRfqTradesV5RespDataInnerLegsInner) HasTradeId() bool {
	if o != nil && !IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given string and assigns it to the TradeId field.
func (o *GetRfqTradesV5RespDataInnerLegsInner) SetTradeId(v string) {
	o.TradeId = &v
}

func (o GetRfqTradesV5RespDataInnerLegsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRfqTradesV5RespDataInnerLegsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.FeeCcy) {
		toSerialize["feeCcy"] = o.FeeCcy
	}
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
	return toSerialize, nil
}

type NullableGetRfqTradesV5RespDataInnerLegsInner struct {
	value *GetRfqTradesV5RespDataInnerLegsInner
	isSet bool
}

func (v NullableGetRfqTradesV5RespDataInnerLegsInner) Get() *GetRfqTradesV5RespDataInnerLegsInner {
	return v.value
}

func (v *NullableGetRfqTradesV5RespDataInnerLegsInner) Set(val *GetRfqTradesV5RespDataInnerLegsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRfqTradesV5RespDataInnerLegsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRfqTradesV5RespDataInnerLegsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRfqTradesV5RespDataInnerLegsInner(val *GetRfqTradesV5RespDataInnerLegsInner) *NullableGetRfqTradesV5RespDataInnerLegsInner {
	return &NullableGetRfqTradesV5RespDataInnerLegsInner{value: val, isSet: true}
}

func (v NullableGetRfqTradesV5RespDataInnerLegsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRfqTradesV5RespDataInnerLegsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


