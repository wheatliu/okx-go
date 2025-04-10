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

// checks if the GetRfqPublicTradesV5RespDataInnerDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRfqPublicTradesV5RespDataInnerDataInner{}

// GetRfqPublicTradesV5RespDataInnerDataInner struct for GetRfqPublicTradesV5RespDataInnerDataInner
type GetRfqPublicTradesV5RespDataInnerDataInner struct {
	// Block trade ID.
	BlockTdId *string `json:"blockTdId,omitempty"`
	// The time the trade was executed. Unix timestamp in milliseconds.
	CTime *string `json:"cTime,omitempty"`
	// Instrument ID, e.g. BTC-USDT-SWAP
	InstId *string `json:"instId,omitempty"`
	// Legs of trade
	Legs []GetRfqPublicTradesV5RespDataInnerDataInnerLegsInner `json:"legs,omitempty"`
	// The price the leg executed
	Px *string `json:"px,omitempty"`
	// The direction of the leg from the Takers perspective. Valid value can be buy or sell.
	Side *string `json:"side,omitempty"`
	// Option strategy, e.g. CALL_CALENDAR_SPREAD
	Strategy *string `json:"strategy,omitempty"`
	// Trade quantity   For spot trading, the unit is base currency  For `FUTURES`/`SWAP`/`OPTION`, the unit is contract.
	Sz *string `json:"sz,omitempty"`
	// Last traded ID.
	TradeId *string `json:"tradeId,omitempty"`
}

// NewGetRfqPublicTradesV5RespDataInnerDataInner instantiates a new GetRfqPublicTradesV5RespDataInnerDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRfqPublicTradesV5RespDataInnerDataInner() *GetRfqPublicTradesV5RespDataInnerDataInner {
	this := GetRfqPublicTradesV5RespDataInnerDataInner{}
	var blockTdId string = ""
	this.BlockTdId = &blockTdId
	var cTime string = ""
	this.CTime = &cTime
	var instId string = ""
	this.InstId = &instId
	var px string = ""
	this.Px = &px
	var side string = ""
	this.Side = &side
	var strategy string = ""
	this.Strategy = &strategy
	var sz string = ""
	this.Sz = &sz
	var tradeId string = ""
	this.TradeId = &tradeId
	return &this
}

// NewGetRfqPublicTradesV5RespDataInnerDataInnerWithDefaults instantiates a new GetRfqPublicTradesV5RespDataInnerDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRfqPublicTradesV5RespDataInnerDataInnerWithDefaults() *GetRfqPublicTradesV5RespDataInnerDataInner {
	this := GetRfqPublicTradesV5RespDataInnerDataInner{}
	var blockTdId string = ""
	this.BlockTdId = &blockTdId
	var cTime string = ""
	this.CTime = &cTime
	var instId string = ""
	this.InstId = &instId
	var px string = ""
	this.Px = &px
	var side string = ""
	this.Side = &side
	var strategy string = ""
	this.Strategy = &strategy
	var sz string = ""
	this.Sz = &sz
	var tradeId string = ""
	this.TradeId = &tradeId
	return &this
}

// GetBlockTdId returns the BlockTdId field value if set, zero value otherwise.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetBlockTdId() string {
	if o == nil || IsNil(o.BlockTdId) {
		var ret string
		return ret
	}
	return *o.BlockTdId
}

// GetBlockTdIdOk returns a tuple with the BlockTdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetBlockTdIdOk() (*string, bool) {
	if o == nil || IsNil(o.BlockTdId) {
		return nil, false
	}
	return o.BlockTdId, true
}

// HasBlockTdId returns a boolean if a field has been set.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) HasBlockTdId() bool {
	if o != nil && !IsNil(o.BlockTdId) {
		return true
	}

	return false
}

// SetBlockTdId gets a reference to the given string and assigns it to the BlockTdId field.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) SetBlockTdId(v string) {
	o.BlockTdId = &v
}

// GetCTime returns the CTime field value if set, zero value otherwise.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetCTime() string {
	if o == nil || IsNil(o.CTime) {
		var ret string
		return ret
	}
	return *o.CTime
}

// GetCTimeOk returns a tuple with the CTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetCTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CTime) {
		return nil, false
	}
	return o.CTime, true
}

// HasCTime returns a boolean if a field has been set.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) HasCTime() bool {
	if o != nil && !IsNil(o.CTime) {
		return true
	}

	return false
}

// SetCTime gets a reference to the given string and assigns it to the CTime field.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) SetCTime(v string) {
	o.CTime = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) SetInstId(v string) {
	o.InstId = &v
}

// GetLegs returns the Legs field value if set, zero value otherwise.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetLegs() []GetRfqPublicTradesV5RespDataInnerDataInnerLegsInner {
	if o == nil || IsNil(o.Legs) {
		var ret []GetRfqPublicTradesV5RespDataInnerDataInnerLegsInner
		return ret
	}
	return o.Legs
}

// GetLegsOk returns a tuple with the Legs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetLegsOk() ([]GetRfqPublicTradesV5RespDataInnerDataInnerLegsInner, bool) {
	if o == nil || IsNil(o.Legs) {
		return nil, false
	}
	return o.Legs, true
}

// HasLegs returns a boolean if a field has been set.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) HasLegs() bool {
	if o != nil && !IsNil(o.Legs) {
		return true
	}

	return false
}

// SetLegs gets a reference to the given []GetRfqPublicTradesV5RespDataInnerDataInnerLegsInner and assigns it to the Legs field.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) SetLegs(v []GetRfqPublicTradesV5RespDataInnerDataInnerLegsInner) {
	o.Legs = v
}

// GetPx returns the Px field value if set, zero value otherwise.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetPx() string {
	if o == nil || IsNil(o.Px) {
		var ret string
		return ret
	}
	return *o.Px
}

// GetPxOk returns a tuple with the Px field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetPxOk() (*string, bool) {
	if o == nil || IsNil(o.Px) {
		return nil, false
	}
	return o.Px, true
}

// HasPx returns a boolean if a field has been set.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) HasPx() bool {
	if o != nil && !IsNil(o.Px) {
		return true
	}

	return false
}

// SetPx gets a reference to the given string and assigns it to the Px field.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) SetPx(v string) {
	o.Px = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) SetSide(v string) {
	o.Side = &v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetStrategy() string {
	if o == nil || IsNil(o.Strategy) {
		var ret string
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.Strategy) {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) HasStrategy() bool {
	if o != nil && !IsNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given string and assigns it to the Strategy field.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) SetStrategy(v string) {
	o.Strategy = &v
}

// GetSz returns the Sz field value if set, zero value otherwise.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetSz() string {
	if o == nil || IsNil(o.Sz) {
		var ret string
		return ret
	}
	return *o.Sz
}

// GetSzOk returns a tuple with the Sz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetSzOk() (*string, bool) {
	if o == nil || IsNil(o.Sz) {
		return nil, false
	}
	return o.Sz, true
}

// HasSz returns a boolean if a field has been set.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) HasSz() bool {
	if o != nil && !IsNil(o.Sz) {
		return true
	}

	return false
}

// SetSz gets a reference to the given string and assigns it to the Sz field.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) SetSz(v string) {
	o.Sz = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetTradeId() string {
	if o == nil || IsNil(o.TradeId) {
		var ret string
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) GetTradeIdOk() (*string, bool) {
	if o == nil || IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) HasTradeId() bool {
	if o != nil && !IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given string and assigns it to the TradeId field.
func (o *GetRfqPublicTradesV5RespDataInnerDataInner) SetTradeId(v string) {
	o.TradeId = &v
}

func (o GetRfqPublicTradesV5RespDataInnerDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRfqPublicTradesV5RespDataInnerDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockTdId) {
		toSerialize["blockTdId"] = o.BlockTdId
	}
	if !IsNil(o.CTime) {
		toSerialize["cTime"] = o.CTime
	}
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	if !IsNil(o.Legs) {
		toSerialize["legs"] = o.Legs
	}
	if !IsNil(o.Px) {
		toSerialize["px"] = o.Px
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !IsNil(o.Strategy) {
		toSerialize["strategy"] = o.Strategy
	}
	if !IsNil(o.Sz) {
		toSerialize["sz"] = o.Sz
	}
	if !IsNil(o.TradeId) {
		toSerialize["tradeId"] = o.TradeId
	}
	return toSerialize, nil
}

type NullableGetRfqPublicTradesV5RespDataInnerDataInner struct {
	value *GetRfqPublicTradesV5RespDataInnerDataInner
	isSet bool
}

func (v NullableGetRfqPublicTradesV5RespDataInnerDataInner) Get() *GetRfqPublicTradesV5RespDataInnerDataInner {
	return v.value
}

func (v *NullableGetRfqPublicTradesV5RespDataInnerDataInner) Set(val *GetRfqPublicTradesV5RespDataInnerDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRfqPublicTradesV5RespDataInnerDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRfqPublicTradesV5RespDataInnerDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRfqPublicTradesV5RespDataInnerDataInner(val *GetRfqPublicTradesV5RespDataInnerDataInner) *NullableGetRfqPublicTradesV5RespDataInnerDataInner {
	return &NullableGetRfqPublicTradesV5RespDataInnerDataInner{value: val, isSet: true}
}

func (v NullableGetRfqPublicTradesV5RespDataInnerDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRfqPublicTradesV5RespDataInnerDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


