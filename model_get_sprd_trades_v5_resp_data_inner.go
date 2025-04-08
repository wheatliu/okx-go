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

// checks if the GetSprdTradesV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSprdTradesV5RespDataInner{}

// GetSprdTradesV5RespDataInner struct for GetSprdTradesV5RespDataInner
type GetSprdTradesV5RespDataInner struct {
	// Client Order ID as assigned by the client
	ClOrdId *string `json:"clOrdId,omitempty"`
	// Error Code, the default is 0
	Code *string `json:"code,omitempty"`
	// Liquidity taker or maker, `T`: taker  `M`: maker
	ExecType *string `json:"execType,omitempty"`
	// Filled price
	FillPx *string `json:"fillPx,omitempty"`
	// Filled quantity
	FillSz *string `json:"fillSz,omitempty"`
	// Legs of trade
	Legs []GetSprdTradesV5RespDataInnerLegsInner `json:"legs,omitempty"`
	// Error Message, the default is \"\"
	Msg *string `json:"msg,omitempty"`
	// Order ID
	OrdId *string `json:"ordId,omitempty"`
	// Order side, `buy` `sell`
	Side *string `json:"side,omitempty"`
	// spread ID
	SprdId *string `json:"sprdId,omitempty"`
	// Trade state.   Valid values are `filled` and `rejected`
	State *string `json:"state,omitempty"`
	// Order tag
	Tag *string `json:"tag,omitempty"`
	// Trade ID
	TradeId *string `json:"tradeId,omitempty"`
	// Data generation time, Unix timestamp format in milliseconds, e.g. `1597026383085`.
	Ts *string `json:"ts,omitempty"`
}

// NewGetSprdTradesV5RespDataInner instantiates a new GetSprdTradesV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSprdTradesV5RespDataInner() *GetSprdTradesV5RespDataInner {
	this := GetSprdTradesV5RespDataInner{}
	var clOrdId string = ""
	this.ClOrdId = &clOrdId
	var code string = ""
	this.Code = &code
	var execType string = ""
	this.ExecType = &execType
	var fillPx string = ""
	this.FillPx = &fillPx
	var fillSz string = ""
	this.FillSz = &fillSz
	var msg string = ""
	this.Msg = &msg
	var ordId string = ""
	this.OrdId = &ordId
	var side string = ""
	this.Side = &side
	var sprdId string = ""
	this.SprdId = &sprdId
	var state string = ""
	this.State = &state
	var tag string = ""
	this.Tag = &tag
	var tradeId string = ""
	this.TradeId = &tradeId
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewGetSprdTradesV5RespDataInnerWithDefaults instantiates a new GetSprdTradesV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSprdTradesV5RespDataInnerWithDefaults() *GetSprdTradesV5RespDataInner {
	this := GetSprdTradesV5RespDataInner{}
	var clOrdId string = ""
	this.ClOrdId = &clOrdId
	var code string = ""
	this.Code = &code
	var execType string = ""
	this.ExecType = &execType
	var fillPx string = ""
	this.FillPx = &fillPx
	var fillSz string = ""
	this.FillSz = &fillSz
	var msg string = ""
	this.Msg = &msg
	var ordId string = ""
	this.OrdId = &ordId
	var side string = ""
	this.Side = &side
	var sprdId string = ""
	this.SprdId = &sprdId
	var state string = ""
	this.State = &state
	var tag string = ""
	this.Tag = &tag
	var tradeId string = ""
	this.TradeId = &tradeId
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetClOrdId returns the ClOrdId field value if set, zero value otherwise.
func (o *GetSprdTradesV5RespDataInner) GetClOrdId() string {
	if o == nil || IsNil(o.ClOrdId) {
		var ret string
		return ret
	}
	return *o.ClOrdId
}

// GetClOrdIdOk returns a tuple with the ClOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSprdTradesV5RespDataInner) GetClOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClOrdId) {
		return nil, false
	}
	return o.ClOrdId, true
}

// HasClOrdId returns a boolean if a field has been set.
func (o *GetSprdTradesV5RespDataInner) HasClOrdId() bool {
	if o != nil && !IsNil(o.ClOrdId) {
		return true
	}

	return false
}

// SetClOrdId gets a reference to the given string and assigns it to the ClOrdId field.
func (o *GetSprdTradesV5RespDataInner) SetClOrdId(v string) {
	o.ClOrdId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetSprdTradesV5RespDataInner) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSprdTradesV5RespDataInner) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetSprdTradesV5RespDataInner) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetSprdTradesV5RespDataInner) SetCode(v string) {
	o.Code = &v
}

// GetExecType returns the ExecType field value if set, zero value otherwise.
func (o *GetSprdTradesV5RespDataInner) GetExecType() string {
	if o == nil || IsNil(o.ExecType) {
		var ret string
		return ret
	}
	return *o.ExecType
}

// GetExecTypeOk returns a tuple with the ExecType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSprdTradesV5RespDataInner) GetExecTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExecType) {
		return nil, false
	}
	return o.ExecType, true
}

// HasExecType returns a boolean if a field has been set.
func (o *GetSprdTradesV5RespDataInner) HasExecType() bool {
	if o != nil && !IsNil(o.ExecType) {
		return true
	}

	return false
}

// SetExecType gets a reference to the given string and assigns it to the ExecType field.
func (o *GetSprdTradesV5RespDataInner) SetExecType(v string) {
	o.ExecType = &v
}

// GetFillPx returns the FillPx field value if set, zero value otherwise.
func (o *GetSprdTradesV5RespDataInner) GetFillPx() string {
	if o == nil || IsNil(o.FillPx) {
		var ret string
		return ret
	}
	return *o.FillPx
}

// GetFillPxOk returns a tuple with the FillPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSprdTradesV5RespDataInner) GetFillPxOk() (*string, bool) {
	if o == nil || IsNil(o.FillPx) {
		return nil, false
	}
	return o.FillPx, true
}

// HasFillPx returns a boolean if a field has been set.
func (o *GetSprdTradesV5RespDataInner) HasFillPx() bool {
	if o != nil && !IsNil(o.FillPx) {
		return true
	}

	return false
}

// SetFillPx gets a reference to the given string and assigns it to the FillPx field.
func (o *GetSprdTradesV5RespDataInner) SetFillPx(v string) {
	o.FillPx = &v
}

// GetFillSz returns the FillSz field value if set, zero value otherwise.
func (o *GetSprdTradesV5RespDataInner) GetFillSz() string {
	if o == nil || IsNil(o.FillSz) {
		var ret string
		return ret
	}
	return *o.FillSz
}

// GetFillSzOk returns a tuple with the FillSz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSprdTradesV5RespDataInner) GetFillSzOk() (*string, bool) {
	if o == nil || IsNil(o.FillSz) {
		return nil, false
	}
	return o.FillSz, true
}

// HasFillSz returns a boolean if a field has been set.
func (o *GetSprdTradesV5RespDataInner) HasFillSz() bool {
	if o != nil && !IsNil(o.FillSz) {
		return true
	}

	return false
}

// SetFillSz gets a reference to the given string and assigns it to the FillSz field.
func (o *GetSprdTradesV5RespDataInner) SetFillSz(v string) {
	o.FillSz = &v
}

// GetLegs returns the Legs field value if set, zero value otherwise.
func (o *GetSprdTradesV5RespDataInner) GetLegs() []GetSprdTradesV5RespDataInnerLegsInner {
	if o == nil || IsNil(o.Legs) {
		var ret []GetSprdTradesV5RespDataInnerLegsInner
		return ret
	}
	return o.Legs
}

// GetLegsOk returns a tuple with the Legs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSprdTradesV5RespDataInner) GetLegsOk() ([]GetSprdTradesV5RespDataInnerLegsInner, bool) {
	if o == nil || IsNil(o.Legs) {
		return nil, false
	}
	return o.Legs, true
}

// HasLegs returns a boolean if a field has been set.
func (o *GetSprdTradesV5RespDataInner) HasLegs() bool {
	if o != nil && !IsNil(o.Legs) {
		return true
	}

	return false
}

// SetLegs gets a reference to the given []GetSprdTradesV5RespDataInnerLegsInner and assigns it to the Legs field.
func (o *GetSprdTradesV5RespDataInner) SetLegs(v []GetSprdTradesV5RespDataInnerLegsInner) {
	o.Legs = v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *GetSprdTradesV5RespDataInner) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSprdTradesV5RespDataInner) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *GetSprdTradesV5RespDataInner) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *GetSprdTradesV5RespDataInner) SetMsg(v string) {
	o.Msg = &v
}

// GetOrdId returns the OrdId field value if set, zero value otherwise.
func (o *GetSprdTradesV5RespDataInner) GetOrdId() string {
	if o == nil || IsNil(o.OrdId) {
		var ret string
		return ret
	}
	return *o.OrdId
}

// GetOrdIdOk returns a tuple with the OrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSprdTradesV5RespDataInner) GetOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrdId) {
		return nil, false
	}
	return o.OrdId, true
}

// HasOrdId returns a boolean if a field has been set.
func (o *GetSprdTradesV5RespDataInner) HasOrdId() bool {
	if o != nil && !IsNil(o.OrdId) {
		return true
	}

	return false
}

// SetOrdId gets a reference to the given string and assigns it to the OrdId field.
func (o *GetSprdTradesV5RespDataInner) SetOrdId(v string) {
	o.OrdId = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *GetSprdTradesV5RespDataInner) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSprdTradesV5RespDataInner) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *GetSprdTradesV5RespDataInner) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *GetSprdTradesV5RespDataInner) SetSide(v string) {
	o.Side = &v
}

// GetSprdId returns the SprdId field value if set, zero value otherwise.
func (o *GetSprdTradesV5RespDataInner) GetSprdId() string {
	if o == nil || IsNil(o.SprdId) {
		var ret string
		return ret
	}
	return *o.SprdId
}

// GetSprdIdOk returns a tuple with the SprdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSprdTradesV5RespDataInner) GetSprdIdOk() (*string, bool) {
	if o == nil || IsNil(o.SprdId) {
		return nil, false
	}
	return o.SprdId, true
}

// HasSprdId returns a boolean if a field has been set.
func (o *GetSprdTradesV5RespDataInner) HasSprdId() bool {
	if o != nil && !IsNil(o.SprdId) {
		return true
	}

	return false
}

// SetSprdId gets a reference to the given string and assigns it to the SprdId field.
func (o *GetSprdTradesV5RespDataInner) SetSprdId(v string) {
	o.SprdId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GetSprdTradesV5RespDataInner) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSprdTradesV5RespDataInner) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GetSprdTradesV5RespDataInner) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GetSprdTradesV5RespDataInner) SetState(v string) {
	o.State = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *GetSprdTradesV5RespDataInner) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSprdTradesV5RespDataInner) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *GetSprdTradesV5RespDataInner) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *GetSprdTradesV5RespDataInner) SetTag(v string) {
	o.Tag = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *GetSprdTradesV5RespDataInner) GetTradeId() string {
	if o == nil || IsNil(o.TradeId) {
		var ret string
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSprdTradesV5RespDataInner) GetTradeIdOk() (*string, bool) {
	if o == nil || IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *GetSprdTradesV5RespDataInner) HasTradeId() bool {
	if o != nil && !IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given string and assigns it to the TradeId field.
func (o *GetSprdTradesV5RespDataInner) SetTradeId(v string) {
	o.TradeId = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetSprdTradesV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSprdTradesV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetSprdTradesV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetSprdTradesV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

func (o GetSprdTradesV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSprdTradesV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClOrdId) {
		toSerialize["clOrdId"] = o.ClOrdId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.ExecType) {
		toSerialize["execType"] = o.ExecType
	}
	if !IsNil(o.FillPx) {
		toSerialize["fillPx"] = o.FillPx
	}
	if !IsNil(o.FillSz) {
		toSerialize["fillSz"] = o.FillSz
	}
	if !IsNil(o.Legs) {
		toSerialize["legs"] = o.Legs
	}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !IsNil(o.OrdId) {
		toSerialize["ordId"] = o.OrdId
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !IsNil(o.SprdId) {
		toSerialize["sprdId"] = o.SprdId
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.TradeId) {
		toSerialize["tradeId"] = o.TradeId
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableGetSprdTradesV5RespDataInner struct {
	value *GetSprdTradesV5RespDataInner
	isSet bool
}

func (v NullableGetSprdTradesV5RespDataInner) Get() *GetSprdTradesV5RespDataInner {
	return v.value
}

func (v *NullableGetSprdTradesV5RespDataInner) Set(val *GetSprdTradesV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSprdTradesV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSprdTradesV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSprdTradesV5RespDataInner(val *GetSprdTradesV5RespDataInner) *NullableGetSprdTradesV5RespDataInner {
	return &NullableGetSprdTradesV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetSprdTradesV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSprdTradesV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


