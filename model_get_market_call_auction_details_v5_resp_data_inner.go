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

// checks if the GetMarketCallAuctionDetailsV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMarketCallAuctionDetailsV5RespDataInner{}

// GetMarketCallAuctionDetailsV5RespDataInner struct for GetMarketCallAuctionDetailsV5RespDataInner
type GetMarketCallAuctionDetailsV5RespDataInner struct {
	// Call auction end time. Unix timestamp in milliseconds.
	AuctionEndTime *string `json:"auctionEndTime,omitempty"`
	// Equilibrium price
	EqPx *string `json:"eqPx,omitempty"`
	// Instrument ID
	InstId *string `json:"instId,omitempty"`
	// Matched size for both buy and sell  The unit is in base currency
	MatchedSz *string `json:"matchedSz,omitempty"`
	// Trading state of the symbol  `call_auction`  `continuous_trading`
	State *string `json:"state,omitempty"`
	// Data generation time. Unix timestamp in millieseconds.
	Ts *string `json:"ts,omitempty"`
	// Unmatched size
	UnmatchedSz *string `json:"unmatchedSz,omitempty"`
}

// NewGetMarketCallAuctionDetailsV5RespDataInner instantiates a new GetMarketCallAuctionDetailsV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarketCallAuctionDetailsV5RespDataInner() *GetMarketCallAuctionDetailsV5RespDataInner {
	this := GetMarketCallAuctionDetailsV5RespDataInner{}
	var auctionEndTime string = ""
	this.AuctionEndTime = &auctionEndTime
	var eqPx string = ""
	this.EqPx = &eqPx
	var instId string = ""
	this.InstId = &instId
	var matchedSz string = ""
	this.MatchedSz = &matchedSz
	var state string = ""
	this.State = &state
	var ts string = ""
	this.Ts = &ts
	var unmatchedSz string = ""
	this.UnmatchedSz = &unmatchedSz
	return &this
}

// NewGetMarketCallAuctionDetailsV5RespDataInnerWithDefaults instantiates a new GetMarketCallAuctionDetailsV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarketCallAuctionDetailsV5RespDataInnerWithDefaults() *GetMarketCallAuctionDetailsV5RespDataInner {
	this := GetMarketCallAuctionDetailsV5RespDataInner{}
	var auctionEndTime string = ""
	this.AuctionEndTime = &auctionEndTime
	var eqPx string = ""
	this.EqPx = &eqPx
	var instId string = ""
	this.InstId = &instId
	var matchedSz string = ""
	this.MatchedSz = &matchedSz
	var state string = ""
	this.State = &state
	var ts string = ""
	this.Ts = &ts
	var unmatchedSz string = ""
	this.UnmatchedSz = &unmatchedSz
	return &this
}

// GetAuctionEndTime returns the AuctionEndTime field value if set, zero value otherwise.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetAuctionEndTime() string {
	if o == nil || IsNil(o.AuctionEndTime) {
		var ret string
		return ret
	}
	return *o.AuctionEndTime
}

// GetAuctionEndTimeOk returns a tuple with the AuctionEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetAuctionEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AuctionEndTime) {
		return nil, false
	}
	return o.AuctionEndTime, true
}

// HasAuctionEndTime returns a boolean if a field has been set.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) HasAuctionEndTime() bool {
	if o != nil && !IsNil(o.AuctionEndTime) {
		return true
	}

	return false
}

// SetAuctionEndTime gets a reference to the given string and assigns it to the AuctionEndTime field.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) SetAuctionEndTime(v string) {
	o.AuctionEndTime = &v
}

// GetEqPx returns the EqPx field value if set, zero value otherwise.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetEqPx() string {
	if o == nil || IsNil(o.EqPx) {
		var ret string
		return ret
	}
	return *o.EqPx
}

// GetEqPxOk returns a tuple with the EqPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetEqPxOk() (*string, bool) {
	if o == nil || IsNil(o.EqPx) {
		return nil, false
	}
	return o.EqPx, true
}

// HasEqPx returns a boolean if a field has been set.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) HasEqPx() bool {
	if o != nil && !IsNil(o.EqPx) {
		return true
	}

	return false
}

// SetEqPx gets a reference to the given string and assigns it to the EqPx field.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) SetEqPx(v string) {
	o.EqPx = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) SetInstId(v string) {
	o.InstId = &v
}

// GetMatchedSz returns the MatchedSz field value if set, zero value otherwise.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetMatchedSz() string {
	if o == nil || IsNil(o.MatchedSz) {
		var ret string
		return ret
	}
	return *o.MatchedSz
}

// GetMatchedSzOk returns a tuple with the MatchedSz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetMatchedSzOk() (*string, bool) {
	if o == nil || IsNil(o.MatchedSz) {
		return nil, false
	}
	return o.MatchedSz, true
}

// HasMatchedSz returns a boolean if a field has been set.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) HasMatchedSz() bool {
	if o != nil && !IsNil(o.MatchedSz) {
		return true
	}

	return false
}

// SetMatchedSz gets a reference to the given string and assigns it to the MatchedSz field.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) SetMatchedSz(v string) {
	o.MatchedSz = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) SetState(v string) {
	o.State = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

// GetUnmatchedSz returns the UnmatchedSz field value if set, zero value otherwise.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetUnmatchedSz() string {
	if o == nil || IsNil(o.UnmatchedSz) {
		var ret string
		return ret
	}
	return *o.UnmatchedSz
}

// GetUnmatchedSzOk returns a tuple with the UnmatchedSz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetUnmatchedSzOk() (*string, bool) {
	if o == nil || IsNil(o.UnmatchedSz) {
		return nil, false
	}
	return o.UnmatchedSz, true
}

// HasUnmatchedSz returns a boolean if a field has been set.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) HasUnmatchedSz() bool {
	if o != nil && !IsNil(o.UnmatchedSz) {
		return true
	}

	return false
}

// SetUnmatchedSz gets a reference to the given string and assigns it to the UnmatchedSz field.
func (o *GetMarketCallAuctionDetailsV5RespDataInner) SetUnmatchedSz(v string) {
	o.UnmatchedSz = &v
}

func (o GetMarketCallAuctionDetailsV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarketCallAuctionDetailsV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuctionEndTime) {
		toSerialize["auctionEndTime"] = o.AuctionEndTime
	}
	if !IsNil(o.EqPx) {
		toSerialize["eqPx"] = o.EqPx
	}
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	if !IsNil(o.MatchedSz) {
		toSerialize["matchedSz"] = o.MatchedSz
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.UnmatchedSz) {
		toSerialize["unmatchedSz"] = o.UnmatchedSz
	}
	return toSerialize, nil
}

type NullableGetMarketCallAuctionDetailsV5RespDataInner struct {
	value *GetMarketCallAuctionDetailsV5RespDataInner
	isSet bool
}

func (v NullableGetMarketCallAuctionDetailsV5RespDataInner) Get() *GetMarketCallAuctionDetailsV5RespDataInner {
	return v.value
}

func (v *NullableGetMarketCallAuctionDetailsV5RespDataInner) Set(val *GetMarketCallAuctionDetailsV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarketCallAuctionDetailsV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarketCallAuctionDetailsV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarketCallAuctionDetailsV5RespDataInner(val *GetMarketCallAuctionDetailsV5RespDataInner) *NullableGetMarketCallAuctionDetailsV5RespDataInner {
	return &NullableGetMarketCallAuctionDetailsV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetMarketCallAuctionDetailsV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarketCallAuctionDetailsV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


