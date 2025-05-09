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

// checks if the CreateRfqExecuteQuoteV5RespDataInnerLegsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRfqExecuteQuoteV5RespDataInnerLegsInner{}

// CreateRfqExecuteQuoteV5RespDataInnerLegsInner struct for CreateRfqExecuteQuoteV5RespDataInnerLegsInner
type CreateRfqExecuteQuoteV5RespDataInnerLegsInner struct {
	// Fee for the individual leg.   Negative fee represents the user transaction fee charged by the platform. Positive fee represents rebate.
	Fee *string `json:"fee,omitempty"`
	// Fee currency. To be read in conjunction with fee
	FeeCcy *string `json:"feeCcy,omitempty"`
	// Instrument ID, e.g. BTC-USDT-SWAP
	InstId *string `json:"instId,omitempty"`
	// The price the leg executed
	Px *string `json:"px,omitempty"`
	// The direction of the leg from the Takers perspective. Valid value can be buy or sell.
	Side *string `json:"side,omitempty"`
	// Size of the leg in contracts or spot.
	Sz *string `json:"sz,omitempty"`
	// Last traded ID.
	TradeId *string `json:"tradeId,omitempty"`
}

// NewCreateRfqExecuteQuoteV5RespDataInnerLegsInner instantiates a new CreateRfqExecuteQuoteV5RespDataInnerLegsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRfqExecuteQuoteV5RespDataInnerLegsInner() *CreateRfqExecuteQuoteV5RespDataInnerLegsInner {
	this := CreateRfqExecuteQuoteV5RespDataInnerLegsInner{}
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

// NewCreateRfqExecuteQuoteV5RespDataInnerLegsInnerWithDefaults instantiates a new CreateRfqExecuteQuoteV5RespDataInnerLegsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRfqExecuteQuoteV5RespDataInnerLegsInnerWithDefaults() *CreateRfqExecuteQuoteV5RespDataInnerLegsInner {
	this := CreateRfqExecuteQuoteV5RespDataInnerLegsInner{}
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
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetFee() string {
	if o == nil || IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetFeeOk() (*string, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) SetFee(v string) {
	o.Fee = &v
}

// GetFeeCcy returns the FeeCcy field value if set, zero value otherwise.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetFeeCcy() string {
	if o == nil || IsNil(o.FeeCcy) {
		var ret string
		return ret
	}
	return *o.FeeCcy
}

// GetFeeCcyOk returns a tuple with the FeeCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetFeeCcyOk() (*string, bool) {
	if o == nil || IsNil(o.FeeCcy) {
		return nil, false
	}
	return o.FeeCcy, true
}

// HasFeeCcy returns a boolean if a field has been set.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) HasFeeCcy() bool {
	if o != nil && !IsNil(o.FeeCcy) {
		return true
	}

	return false
}

// SetFeeCcy gets a reference to the given string and assigns it to the FeeCcy field.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) SetFeeCcy(v string) {
	o.FeeCcy = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) SetInstId(v string) {
	o.InstId = &v
}

// GetPx returns the Px field value if set, zero value otherwise.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetPx() string {
	if o == nil || IsNil(o.Px) {
		var ret string
		return ret
	}
	return *o.Px
}

// GetPxOk returns a tuple with the Px field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetPxOk() (*string, bool) {
	if o == nil || IsNil(o.Px) {
		return nil, false
	}
	return o.Px, true
}

// HasPx returns a boolean if a field has been set.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) HasPx() bool {
	if o != nil && !IsNil(o.Px) {
		return true
	}

	return false
}

// SetPx gets a reference to the given string and assigns it to the Px field.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) SetPx(v string) {
	o.Px = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) SetSide(v string) {
	o.Side = &v
}

// GetSz returns the Sz field value if set, zero value otherwise.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetSz() string {
	if o == nil || IsNil(o.Sz) {
		var ret string
		return ret
	}
	return *o.Sz
}

// GetSzOk returns a tuple with the Sz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetSzOk() (*string, bool) {
	if o == nil || IsNil(o.Sz) {
		return nil, false
	}
	return o.Sz, true
}

// HasSz returns a boolean if a field has been set.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) HasSz() bool {
	if o != nil && !IsNil(o.Sz) {
		return true
	}

	return false
}

// SetSz gets a reference to the given string and assigns it to the Sz field.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) SetSz(v string) {
	o.Sz = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetTradeId() string {
	if o == nil || IsNil(o.TradeId) {
		var ret string
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) GetTradeIdOk() (*string, bool) {
	if o == nil || IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) HasTradeId() bool {
	if o != nil && !IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given string and assigns it to the TradeId field.
func (o *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) SetTradeId(v string) {
	o.TradeId = &v
}

func (o CreateRfqExecuteQuoteV5RespDataInnerLegsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRfqExecuteQuoteV5RespDataInnerLegsInner) ToMap() (map[string]interface{}, error) {
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

type NullableCreateRfqExecuteQuoteV5RespDataInnerLegsInner struct {
	value *CreateRfqExecuteQuoteV5RespDataInnerLegsInner
	isSet bool
}

func (v NullableCreateRfqExecuteQuoteV5RespDataInnerLegsInner) Get() *CreateRfqExecuteQuoteV5RespDataInnerLegsInner {
	return v.value
}

func (v *NullableCreateRfqExecuteQuoteV5RespDataInnerLegsInner) Set(val *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRfqExecuteQuoteV5RespDataInnerLegsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRfqExecuteQuoteV5RespDataInnerLegsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRfqExecuteQuoteV5RespDataInnerLegsInner(val *CreateRfqExecuteQuoteV5RespDataInnerLegsInner) *NullableCreateRfqExecuteQuoteV5RespDataInnerLegsInner {
	return &NullableCreateRfqExecuteQuoteV5RespDataInnerLegsInner{value: val, isSet: true}
}

func (v NullableCreateRfqExecuteQuoteV5RespDataInnerLegsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRfqExecuteQuoteV5RespDataInnerLegsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


