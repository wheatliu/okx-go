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

// checks if the GetMarketSprdTickerV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMarketSprdTickerV5RespDataInner{}

// GetMarketSprdTickerV5RespDataInner struct for GetMarketSprdTickerV5RespDataInner
type GetMarketSprdTickerV5RespDataInner struct {
	// Best ask price
	AskPx *string `json:"askPx,omitempty"`
	// Best ask size
	AskSz *string `json:"askSz,omitempty"`
	// Best bid price
	BidPx *string `json:"bidPx,omitempty"`
	// Best bid size
	BidSz *string `json:"bidSz,omitempty"`
	// Highest price in the past 24 hours
	High24h *string `json:"high24h,omitempty"`
	// Last traded price
	Last *string `json:"last,omitempty"`
	// Last traded size
	LastSz *string `json:"lastSz,omitempty"`
	// Lowest price in the past 24 hours
	Low24h *string `json:"low24h,omitempty"`
	// Open price in the past 24 hours
	Open24h *string `json:"open24h,omitempty"`
	// spread ID
	SprdId *string `json:"sprdId,omitempty"`
	// Ticker data generation time, Unix timestamp format in milliseconds, e.g. 1597026383085.
	Ts *string `json:"ts,omitempty"`
	// 24h trading volume   The unit is USD for inverse spreads, and the corresponding baseCcy for linear and hybrid spreads.
	Vol24h *string `json:"vol24h,omitempty"`
}

// NewGetMarketSprdTickerV5RespDataInner instantiates a new GetMarketSprdTickerV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarketSprdTickerV5RespDataInner() *GetMarketSprdTickerV5RespDataInner {
	this := GetMarketSprdTickerV5RespDataInner{}
	var askPx string = ""
	this.AskPx = &askPx
	var askSz string = ""
	this.AskSz = &askSz
	var bidPx string = ""
	this.BidPx = &bidPx
	var bidSz string = ""
	this.BidSz = &bidSz
	var high24h string = ""
	this.High24h = &high24h
	var last string = ""
	this.Last = &last
	var lastSz string = ""
	this.LastSz = &lastSz
	var low24h string = ""
	this.Low24h = &low24h
	var open24h string = ""
	this.Open24h = &open24h
	var sprdId string = ""
	this.SprdId = &sprdId
	var ts string = ""
	this.Ts = &ts
	var vol24h string = ""
	this.Vol24h = &vol24h
	return &this
}

// NewGetMarketSprdTickerV5RespDataInnerWithDefaults instantiates a new GetMarketSprdTickerV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarketSprdTickerV5RespDataInnerWithDefaults() *GetMarketSprdTickerV5RespDataInner {
	this := GetMarketSprdTickerV5RespDataInner{}
	var askPx string = ""
	this.AskPx = &askPx
	var askSz string = ""
	this.AskSz = &askSz
	var bidPx string = ""
	this.BidPx = &bidPx
	var bidSz string = ""
	this.BidSz = &bidSz
	var high24h string = ""
	this.High24h = &high24h
	var last string = ""
	this.Last = &last
	var lastSz string = ""
	this.LastSz = &lastSz
	var low24h string = ""
	this.Low24h = &low24h
	var open24h string = ""
	this.Open24h = &open24h
	var sprdId string = ""
	this.SprdId = &sprdId
	var ts string = ""
	this.Ts = &ts
	var vol24h string = ""
	this.Vol24h = &vol24h
	return &this
}

// GetAskPx returns the AskPx field value if set, zero value otherwise.
func (o *GetMarketSprdTickerV5RespDataInner) GetAskPx() string {
	if o == nil || IsNil(o.AskPx) {
		var ret string
		return ret
	}
	return *o.AskPx
}

// GetAskPxOk returns a tuple with the AskPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketSprdTickerV5RespDataInner) GetAskPxOk() (*string, bool) {
	if o == nil || IsNil(o.AskPx) {
		return nil, false
	}
	return o.AskPx, true
}

// HasAskPx returns a boolean if a field has been set.
func (o *GetMarketSprdTickerV5RespDataInner) HasAskPx() bool {
	if o != nil && !IsNil(o.AskPx) {
		return true
	}

	return false
}

// SetAskPx gets a reference to the given string and assigns it to the AskPx field.
func (o *GetMarketSprdTickerV5RespDataInner) SetAskPx(v string) {
	o.AskPx = &v
}

// GetAskSz returns the AskSz field value if set, zero value otherwise.
func (o *GetMarketSprdTickerV5RespDataInner) GetAskSz() string {
	if o == nil || IsNil(o.AskSz) {
		var ret string
		return ret
	}
	return *o.AskSz
}

// GetAskSzOk returns a tuple with the AskSz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketSprdTickerV5RespDataInner) GetAskSzOk() (*string, bool) {
	if o == nil || IsNil(o.AskSz) {
		return nil, false
	}
	return o.AskSz, true
}

// HasAskSz returns a boolean if a field has been set.
func (o *GetMarketSprdTickerV5RespDataInner) HasAskSz() bool {
	if o != nil && !IsNil(o.AskSz) {
		return true
	}

	return false
}

// SetAskSz gets a reference to the given string and assigns it to the AskSz field.
func (o *GetMarketSprdTickerV5RespDataInner) SetAskSz(v string) {
	o.AskSz = &v
}

// GetBidPx returns the BidPx field value if set, zero value otherwise.
func (o *GetMarketSprdTickerV5RespDataInner) GetBidPx() string {
	if o == nil || IsNil(o.BidPx) {
		var ret string
		return ret
	}
	return *o.BidPx
}

// GetBidPxOk returns a tuple with the BidPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketSprdTickerV5RespDataInner) GetBidPxOk() (*string, bool) {
	if o == nil || IsNil(o.BidPx) {
		return nil, false
	}
	return o.BidPx, true
}

// HasBidPx returns a boolean if a field has been set.
func (o *GetMarketSprdTickerV5RespDataInner) HasBidPx() bool {
	if o != nil && !IsNil(o.BidPx) {
		return true
	}

	return false
}

// SetBidPx gets a reference to the given string and assigns it to the BidPx field.
func (o *GetMarketSprdTickerV5RespDataInner) SetBidPx(v string) {
	o.BidPx = &v
}

// GetBidSz returns the BidSz field value if set, zero value otherwise.
func (o *GetMarketSprdTickerV5RespDataInner) GetBidSz() string {
	if o == nil || IsNil(o.BidSz) {
		var ret string
		return ret
	}
	return *o.BidSz
}

// GetBidSzOk returns a tuple with the BidSz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketSprdTickerV5RespDataInner) GetBidSzOk() (*string, bool) {
	if o == nil || IsNil(o.BidSz) {
		return nil, false
	}
	return o.BidSz, true
}

// HasBidSz returns a boolean if a field has been set.
func (o *GetMarketSprdTickerV5RespDataInner) HasBidSz() bool {
	if o != nil && !IsNil(o.BidSz) {
		return true
	}

	return false
}

// SetBidSz gets a reference to the given string and assigns it to the BidSz field.
func (o *GetMarketSprdTickerV5RespDataInner) SetBidSz(v string) {
	o.BidSz = &v
}

// GetHigh24h returns the High24h field value if set, zero value otherwise.
func (o *GetMarketSprdTickerV5RespDataInner) GetHigh24h() string {
	if o == nil || IsNil(o.High24h) {
		var ret string
		return ret
	}
	return *o.High24h
}

// GetHigh24hOk returns a tuple with the High24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketSprdTickerV5RespDataInner) GetHigh24hOk() (*string, bool) {
	if o == nil || IsNil(o.High24h) {
		return nil, false
	}
	return o.High24h, true
}

// HasHigh24h returns a boolean if a field has been set.
func (o *GetMarketSprdTickerV5RespDataInner) HasHigh24h() bool {
	if o != nil && !IsNil(o.High24h) {
		return true
	}

	return false
}

// SetHigh24h gets a reference to the given string and assigns it to the High24h field.
func (o *GetMarketSprdTickerV5RespDataInner) SetHigh24h(v string) {
	o.High24h = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *GetMarketSprdTickerV5RespDataInner) GetLast() string {
	if o == nil || IsNil(o.Last) {
		var ret string
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketSprdTickerV5RespDataInner) GetLastOk() (*string, bool) {
	if o == nil || IsNil(o.Last) {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *GetMarketSprdTickerV5RespDataInner) HasLast() bool {
	if o != nil && !IsNil(o.Last) {
		return true
	}

	return false
}

// SetLast gets a reference to the given string and assigns it to the Last field.
func (o *GetMarketSprdTickerV5RespDataInner) SetLast(v string) {
	o.Last = &v
}

// GetLastSz returns the LastSz field value if set, zero value otherwise.
func (o *GetMarketSprdTickerV5RespDataInner) GetLastSz() string {
	if o == nil || IsNil(o.LastSz) {
		var ret string
		return ret
	}
	return *o.LastSz
}

// GetLastSzOk returns a tuple with the LastSz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketSprdTickerV5RespDataInner) GetLastSzOk() (*string, bool) {
	if o == nil || IsNil(o.LastSz) {
		return nil, false
	}
	return o.LastSz, true
}

// HasLastSz returns a boolean if a field has been set.
func (o *GetMarketSprdTickerV5RespDataInner) HasLastSz() bool {
	if o != nil && !IsNil(o.LastSz) {
		return true
	}

	return false
}

// SetLastSz gets a reference to the given string and assigns it to the LastSz field.
func (o *GetMarketSprdTickerV5RespDataInner) SetLastSz(v string) {
	o.LastSz = &v
}

// GetLow24h returns the Low24h field value if set, zero value otherwise.
func (o *GetMarketSprdTickerV5RespDataInner) GetLow24h() string {
	if o == nil || IsNil(o.Low24h) {
		var ret string
		return ret
	}
	return *o.Low24h
}

// GetLow24hOk returns a tuple with the Low24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketSprdTickerV5RespDataInner) GetLow24hOk() (*string, bool) {
	if o == nil || IsNil(o.Low24h) {
		return nil, false
	}
	return o.Low24h, true
}

// HasLow24h returns a boolean if a field has been set.
func (o *GetMarketSprdTickerV5RespDataInner) HasLow24h() bool {
	if o != nil && !IsNil(o.Low24h) {
		return true
	}

	return false
}

// SetLow24h gets a reference to the given string and assigns it to the Low24h field.
func (o *GetMarketSprdTickerV5RespDataInner) SetLow24h(v string) {
	o.Low24h = &v
}

// GetOpen24h returns the Open24h field value if set, zero value otherwise.
func (o *GetMarketSprdTickerV5RespDataInner) GetOpen24h() string {
	if o == nil || IsNil(o.Open24h) {
		var ret string
		return ret
	}
	return *o.Open24h
}

// GetOpen24hOk returns a tuple with the Open24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketSprdTickerV5RespDataInner) GetOpen24hOk() (*string, bool) {
	if o == nil || IsNil(o.Open24h) {
		return nil, false
	}
	return o.Open24h, true
}

// HasOpen24h returns a boolean if a field has been set.
func (o *GetMarketSprdTickerV5RespDataInner) HasOpen24h() bool {
	if o != nil && !IsNil(o.Open24h) {
		return true
	}

	return false
}

// SetOpen24h gets a reference to the given string and assigns it to the Open24h field.
func (o *GetMarketSprdTickerV5RespDataInner) SetOpen24h(v string) {
	o.Open24h = &v
}

// GetSprdId returns the SprdId field value if set, zero value otherwise.
func (o *GetMarketSprdTickerV5RespDataInner) GetSprdId() string {
	if o == nil || IsNil(o.SprdId) {
		var ret string
		return ret
	}
	return *o.SprdId
}

// GetSprdIdOk returns a tuple with the SprdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketSprdTickerV5RespDataInner) GetSprdIdOk() (*string, bool) {
	if o == nil || IsNil(o.SprdId) {
		return nil, false
	}
	return o.SprdId, true
}

// HasSprdId returns a boolean if a field has been set.
func (o *GetMarketSprdTickerV5RespDataInner) HasSprdId() bool {
	if o != nil && !IsNil(o.SprdId) {
		return true
	}

	return false
}

// SetSprdId gets a reference to the given string and assigns it to the SprdId field.
func (o *GetMarketSprdTickerV5RespDataInner) SetSprdId(v string) {
	o.SprdId = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetMarketSprdTickerV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketSprdTickerV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetMarketSprdTickerV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetMarketSprdTickerV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

// GetVol24h returns the Vol24h field value if set, zero value otherwise.
func (o *GetMarketSprdTickerV5RespDataInner) GetVol24h() string {
	if o == nil || IsNil(o.Vol24h) {
		var ret string
		return ret
	}
	return *o.Vol24h
}

// GetVol24hOk returns a tuple with the Vol24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketSprdTickerV5RespDataInner) GetVol24hOk() (*string, bool) {
	if o == nil || IsNil(o.Vol24h) {
		return nil, false
	}
	return o.Vol24h, true
}

// HasVol24h returns a boolean if a field has been set.
func (o *GetMarketSprdTickerV5RespDataInner) HasVol24h() bool {
	if o != nil && !IsNil(o.Vol24h) {
		return true
	}

	return false
}

// SetVol24h gets a reference to the given string and assigns it to the Vol24h field.
func (o *GetMarketSprdTickerV5RespDataInner) SetVol24h(v string) {
	o.Vol24h = &v
}

func (o GetMarketSprdTickerV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarketSprdTickerV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AskPx) {
		toSerialize["askPx"] = o.AskPx
	}
	if !IsNil(o.AskSz) {
		toSerialize["askSz"] = o.AskSz
	}
	if !IsNil(o.BidPx) {
		toSerialize["bidPx"] = o.BidPx
	}
	if !IsNil(o.BidSz) {
		toSerialize["bidSz"] = o.BidSz
	}
	if !IsNil(o.High24h) {
		toSerialize["high24h"] = o.High24h
	}
	if !IsNil(o.Last) {
		toSerialize["last"] = o.Last
	}
	if !IsNil(o.LastSz) {
		toSerialize["lastSz"] = o.LastSz
	}
	if !IsNil(o.Low24h) {
		toSerialize["low24h"] = o.Low24h
	}
	if !IsNil(o.Open24h) {
		toSerialize["open24h"] = o.Open24h
	}
	if !IsNil(o.SprdId) {
		toSerialize["sprdId"] = o.SprdId
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.Vol24h) {
		toSerialize["vol24h"] = o.Vol24h
	}
	return toSerialize, nil
}

type NullableGetMarketSprdTickerV5RespDataInner struct {
	value *GetMarketSprdTickerV5RespDataInner
	isSet bool
}

func (v NullableGetMarketSprdTickerV5RespDataInner) Get() *GetMarketSprdTickerV5RespDataInner {
	return v.value
}

func (v *NullableGetMarketSprdTickerV5RespDataInner) Set(val *GetMarketSprdTickerV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarketSprdTickerV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarketSprdTickerV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarketSprdTickerV5RespDataInner(val *GetMarketSprdTickerV5RespDataInner) *NullableGetMarketSprdTickerV5RespDataInner {
	return &NullableGetMarketSprdTickerV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetMarketSprdTickerV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarketSprdTickerV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


