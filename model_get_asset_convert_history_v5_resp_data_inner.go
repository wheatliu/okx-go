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

// checks if the GetAssetConvertHistoryV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAssetConvertHistoryV5RespDataInner{}

// GetAssetConvertHistoryV5RespDataInner struct for GetAssetConvertHistoryV5RespDataInner
type GetAssetConvertHistoryV5RespDataInner struct {
	// Base currency, e.g. `BTC` in `BTC-USDT`
	BaseCcy *string `json:"baseCcy,omitempty"`
	// Client Order ID as assigned by the client
	ClTReqId *string `json:"clTReqId,omitempty"`
	// Filled amount for base currency
	FillBaseSz *string `json:"fillBaseSz,omitempty"`
	// Filled price based on quote currency
	FillPx *string `json:"fillPx,omitempty"`
	// Filled amount for quote currency
	FillQuoteSz *string `json:"fillQuoteSz,omitempty"`
	// Currency pair, e.g. `BTC-USDT`
	InstId *string `json:"instId,omitempty"`
	// Quote currency, e.g. `USDT` in `BTC-USDT`
	QuoteCcy *string `json:"quoteCcy,omitempty"`
	// Trade side based on `baseCcy`  `buy` `sell`
	Side *string `json:"side,omitempty"`
	// Trade state  `fullyFilled` : success   `rejected` : failed
	State *string `json:"state,omitempty"`
	// Trade ID
	TradeId *string `json:"tradeId,omitempty"`
	// Convert trade time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	Ts *string `json:"ts,omitempty"`
}

// NewGetAssetConvertHistoryV5RespDataInner instantiates a new GetAssetConvertHistoryV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAssetConvertHistoryV5RespDataInner() *GetAssetConvertHistoryV5RespDataInner {
	this := GetAssetConvertHistoryV5RespDataInner{}
	var baseCcy string = ""
	this.BaseCcy = &baseCcy
	var clTReqId string = ""
	this.ClTReqId = &clTReqId
	var fillBaseSz string = ""
	this.FillBaseSz = &fillBaseSz
	var fillPx string = ""
	this.FillPx = &fillPx
	var fillQuoteSz string = ""
	this.FillQuoteSz = &fillQuoteSz
	var instId string = ""
	this.InstId = &instId
	var quoteCcy string = ""
	this.QuoteCcy = &quoteCcy
	var side string = ""
	this.Side = &side
	var state string = ""
	this.State = &state
	var tradeId string = ""
	this.TradeId = &tradeId
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewGetAssetConvertHistoryV5RespDataInnerWithDefaults instantiates a new GetAssetConvertHistoryV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAssetConvertHistoryV5RespDataInnerWithDefaults() *GetAssetConvertHistoryV5RespDataInner {
	this := GetAssetConvertHistoryV5RespDataInner{}
	var baseCcy string = ""
	this.BaseCcy = &baseCcy
	var clTReqId string = ""
	this.ClTReqId = &clTReqId
	var fillBaseSz string = ""
	this.FillBaseSz = &fillBaseSz
	var fillPx string = ""
	this.FillPx = &fillPx
	var fillQuoteSz string = ""
	this.FillQuoteSz = &fillQuoteSz
	var instId string = ""
	this.InstId = &instId
	var quoteCcy string = ""
	this.QuoteCcy = &quoteCcy
	var side string = ""
	this.Side = &side
	var state string = ""
	this.State = &state
	var tradeId string = ""
	this.TradeId = &tradeId
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetBaseCcy returns the BaseCcy field value if set, zero value otherwise.
func (o *GetAssetConvertHistoryV5RespDataInner) GetBaseCcy() string {
	if o == nil || IsNil(o.BaseCcy) {
		var ret string
		return ret
	}
	return *o.BaseCcy
}

// GetBaseCcyOk returns a tuple with the BaseCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) GetBaseCcyOk() (*string, bool) {
	if o == nil || IsNil(o.BaseCcy) {
		return nil, false
	}
	return o.BaseCcy, true
}

// HasBaseCcy returns a boolean if a field has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) HasBaseCcy() bool {
	if o != nil && !IsNil(o.BaseCcy) {
		return true
	}

	return false
}

// SetBaseCcy gets a reference to the given string and assigns it to the BaseCcy field.
func (o *GetAssetConvertHistoryV5RespDataInner) SetBaseCcy(v string) {
	o.BaseCcy = &v
}

// GetClTReqId returns the ClTReqId field value if set, zero value otherwise.
func (o *GetAssetConvertHistoryV5RespDataInner) GetClTReqId() string {
	if o == nil || IsNil(o.ClTReqId) {
		var ret string
		return ret
	}
	return *o.ClTReqId
}

// GetClTReqIdOk returns a tuple with the ClTReqId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) GetClTReqIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClTReqId) {
		return nil, false
	}
	return o.ClTReqId, true
}

// HasClTReqId returns a boolean if a field has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) HasClTReqId() bool {
	if o != nil && !IsNil(o.ClTReqId) {
		return true
	}

	return false
}

// SetClTReqId gets a reference to the given string and assigns it to the ClTReqId field.
func (o *GetAssetConvertHistoryV5RespDataInner) SetClTReqId(v string) {
	o.ClTReqId = &v
}

// GetFillBaseSz returns the FillBaseSz field value if set, zero value otherwise.
func (o *GetAssetConvertHistoryV5RespDataInner) GetFillBaseSz() string {
	if o == nil || IsNil(o.FillBaseSz) {
		var ret string
		return ret
	}
	return *o.FillBaseSz
}

// GetFillBaseSzOk returns a tuple with the FillBaseSz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) GetFillBaseSzOk() (*string, bool) {
	if o == nil || IsNil(o.FillBaseSz) {
		return nil, false
	}
	return o.FillBaseSz, true
}

// HasFillBaseSz returns a boolean if a field has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) HasFillBaseSz() bool {
	if o != nil && !IsNil(o.FillBaseSz) {
		return true
	}

	return false
}

// SetFillBaseSz gets a reference to the given string and assigns it to the FillBaseSz field.
func (o *GetAssetConvertHistoryV5RespDataInner) SetFillBaseSz(v string) {
	o.FillBaseSz = &v
}

// GetFillPx returns the FillPx field value if set, zero value otherwise.
func (o *GetAssetConvertHistoryV5RespDataInner) GetFillPx() string {
	if o == nil || IsNil(o.FillPx) {
		var ret string
		return ret
	}
	return *o.FillPx
}

// GetFillPxOk returns a tuple with the FillPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) GetFillPxOk() (*string, bool) {
	if o == nil || IsNil(o.FillPx) {
		return nil, false
	}
	return o.FillPx, true
}

// HasFillPx returns a boolean if a field has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) HasFillPx() bool {
	if o != nil && !IsNil(o.FillPx) {
		return true
	}

	return false
}

// SetFillPx gets a reference to the given string and assigns it to the FillPx field.
func (o *GetAssetConvertHistoryV5RespDataInner) SetFillPx(v string) {
	o.FillPx = &v
}

// GetFillQuoteSz returns the FillQuoteSz field value if set, zero value otherwise.
func (o *GetAssetConvertHistoryV5RespDataInner) GetFillQuoteSz() string {
	if o == nil || IsNil(o.FillQuoteSz) {
		var ret string
		return ret
	}
	return *o.FillQuoteSz
}

// GetFillQuoteSzOk returns a tuple with the FillQuoteSz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) GetFillQuoteSzOk() (*string, bool) {
	if o == nil || IsNil(o.FillQuoteSz) {
		return nil, false
	}
	return o.FillQuoteSz, true
}

// HasFillQuoteSz returns a boolean if a field has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) HasFillQuoteSz() bool {
	if o != nil && !IsNil(o.FillQuoteSz) {
		return true
	}

	return false
}

// SetFillQuoteSz gets a reference to the given string and assigns it to the FillQuoteSz field.
func (o *GetAssetConvertHistoryV5RespDataInner) SetFillQuoteSz(v string) {
	o.FillQuoteSz = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetAssetConvertHistoryV5RespDataInner) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetAssetConvertHistoryV5RespDataInner) SetInstId(v string) {
	o.InstId = &v
}

// GetQuoteCcy returns the QuoteCcy field value if set, zero value otherwise.
func (o *GetAssetConvertHistoryV5RespDataInner) GetQuoteCcy() string {
	if o == nil || IsNil(o.QuoteCcy) {
		var ret string
		return ret
	}
	return *o.QuoteCcy
}

// GetQuoteCcyOk returns a tuple with the QuoteCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) GetQuoteCcyOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteCcy) {
		return nil, false
	}
	return o.QuoteCcy, true
}

// HasQuoteCcy returns a boolean if a field has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) HasQuoteCcy() bool {
	if o != nil && !IsNil(o.QuoteCcy) {
		return true
	}

	return false
}

// SetQuoteCcy gets a reference to the given string and assigns it to the QuoteCcy field.
func (o *GetAssetConvertHistoryV5RespDataInner) SetQuoteCcy(v string) {
	o.QuoteCcy = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *GetAssetConvertHistoryV5RespDataInner) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *GetAssetConvertHistoryV5RespDataInner) SetSide(v string) {
	o.Side = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GetAssetConvertHistoryV5RespDataInner) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GetAssetConvertHistoryV5RespDataInner) SetState(v string) {
	o.State = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *GetAssetConvertHistoryV5RespDataInner) GetTradeId() string {
	if o == nil || IsNil(o.TradeId) {
		var ret string
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) GetTradeIdOk() (*string, bool) {
	if o == nil || IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) HasTradeId() bool {
	if o != nil && !IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given string and assigns it to the TradeId field.
func (o *GetAssetConvertHistoryV5RespDataInner) SetTradeId(v string) {
	o.TradeId = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetAssetConvertHistoryV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetAssetConvertHistoryV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetAssetConvertHistoryV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

func (o GetAssetConvertHistoryV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAssetConvertHistoryV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BaseCcy) {
		toSerialize["baseCcy"] = o.BaseCcy
	}
	if !IsNil(o.ClTReqId) {
		toSerialize["clTReqId"] = o.ClTReqId
	}
	if !IsNil(o.FillBaseSz) {
		toSerialize["fillBaseSz"] = o.FillBaseSz
	}
	if !IsNil(o.FillPx) {
		toSerialize["fillPx"] = o.FillPx
	}
	if !IsNil(o.FillQuoteSz) {
		toSerialize["fillQuoteSz"] = o.FillQuoteSz
	}
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	if !IsNil(o.QuoteCcy) {
		toSerialize["quoteCcy"] = o.QuoteCcy
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.TradeId) {
		toSerialize["tradeId"] = o.TradeId
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableGetAssetConvertHistoryV5RespDataInner struct {
	value *GetAssetConvertHistoryV5RespDataInner
	isSet bool
}

func (v NullableGetAssetConvertHistoryV5RespDataInner) Get() *GetAssetConvertHistoryV5RespDataInner {
	return v.value
}

func (v *NullableGetAssetConvertHistoryV5RespDataInner) Set(val *GetAssetConvertHistoryV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAssetConvertHistoryV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAssetConvertHistoryV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAssetConvertHistoryV5RespDataInner(val *GetAssetConvertHistoryV5RespDataInner) *NullableGetAssetConvertHistoryV5RespDataInner {
	return &NullableGetAssetConvertHistoryV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetAssetConvertHistoryV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAssetConvertHistoryV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


