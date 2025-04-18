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

// checks if the GetAccountSpotBorrowRepayHistoryV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountSpotBorrowRepayHistoryV5RespDataInner{}

// GetAccountSpotBorrowRepayHistoryV5RespDataInner struct for GetAccountSpotBorrowRepayHistoryV5RespDataInner
type GetAccountSpotBorrowRepayHistoryV5RespDataInner struct {
	// Accumulated borrow amount
	AccBorrowed *string `json:"accBorrowed,omitempty"`
	// Amount
	Amt *string `json:"amt,omitempty"`
	// Currency, e.g. `BTC`
	Ccy *string `json:"ccy,omitempty"`
	// Timestamp for the event, Unix timestamp format in milliseconds, e.g. `1597026383085`
	Ts *string `json:"ts,omitempty"`
	// Event type  `auto_borrow`  `auto_repay`  `manual_borrow`  `manual_repay`
	Type *string `json:"type,omitempty"`
}

// NewGetAccountSpotBorrowRepayHistoryV5RespDataInner instantiates a new GetAccountSpotBorrowRepayHistoryV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountSpotBorrowRepayHistoryV5RespDataInner() *GetAccountSpotBorrowRepayHistoryV5RespDataInner {
	this := GetAccountSpotBorrowRepayHistoryV5RespDataInner{}
	var accBorrowed string = ""
	this.AccBorrowed = &accBorrowed
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	var ts string = ""
	this.Ts = &ts
	var type_ string = ""
	this.Type = &type_
	return &this
}

// NewGetAccountSpotBorrowRepayHistoryV5RespDataInnerWithDefaults instantiates a new GetAccountSpotBorrowRepayHistoryV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountSpotBorrowRepayHistoryV5RespDataInnerWithDefaults() *GetAccountSpotBorrowRepayHistoryV5RespDataInner {
	this := GetAccountSpotBorrowRepayHistoryV5RespDataInner{}
	var accBorrowed string = ""
	this.AccBorrowed = &accBorrowed
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	var ts string = ""
	this.Ts = &ts
	var type_ string = ""
	this.Type = &type_
	return &this
}

// GetAccBorrowed returns the AccBorrowed field value if set, zero value otherwise.
func (o *GetAccountSpotBorrowRepayHistoryV5RespDataInner) GetAccBorrowed() string {
	if o == nil || IsNil(o.AccBorrowed) {
		var ret string
		return ret
	}
	return *o.AccBorrowed
}

// GetAccBorrowedOk returns a tuple with the AccBorrowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountSpotBorrowRepayHistoryV5RespDataInner) GetAccBorrowedOk() (*string, bool) {
	if o == nil || IsNil(o.AccBorrowed) {
		return nil, false
	}
	return o.AccBorrowed, true
}

// HasAccBorrowed returns a boolean if a field has been set.
func (o *GetAccountSpotBorrowRepayHistoryV5RespDataInner) HasAccBorrowed() bool {
	if o != nil && !IsNil(o.AccBorrowed) {
		return true
	}

	return false
}

// SetAccBorrowed gets a reference to the given string and assigns it to the AccBorrowed field.
func (o *GetAccountSpotBorrowRepayHistoryV5RespDataInner) SetAccBorrowed(v string) {
	o.AccBorrowed = &v
}

// GetAmt returns the Amt field value if set, zero value otherwise.
func (o *GetAccountSpotBorrowRepayHistoryV5RespDataInner) GetAmt() string {
	if o == nil || IsNil(o.Amt) {
		var ret string
		return ret
	}
	return *o.Amt
}

// GetAmtOk returns a tuple with the Amt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountSpotBorrowRepayHistoryV5RespDataInner) GetAmtOk() (*string, bool) {
	if o == nil || IsNil(o.Amt) {
		return nil, false
	}
	return o.Amt, true
}

// HasAmt returns a boolean if a field has been set.
func (o *GetAccountSpotBorrowRepayHistoryV5RespDataInner) HasAmt() bool {
	if o != nil && !IsNil(o.Amt) {
		return true
	}

	return false
}

// SetAmt gets a reference to the given string and assigns it to the Amt field.
func (o *GetAccountSpotBorrowRepayHistoryV5RespDataInner) SetAmt(v string) {
	o.Amt = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetAccountSpotBorrowRepayHistoryV5RespDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountSpotBorrowRepayHistoryV5RespDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetAccountSpotBorrowRepayHistoryV5RespDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetAccountSpotBorrowRepayHistoryV5RespDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetAccountSpotBorrowRepayHistoryV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountSpotBorrowRepayHistoryV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetAccountSpotBorrowRepayHistoryV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetAccountSpotBorrowRepayHistoryV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetAccountSpotBorrowRepayHistoryV5RespDataInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountSpotBorrowRepayHistoryV5RespDataInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetAccountSpotBorrowRepayHistoryV5RespDataInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetAccountSpotBorrowRepayHistoryV5RespDataInner) SetType(v string) {
	o.Type = &v
}

func (o GetAccountSpotBorrowRepayHistoryV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountSpotBorrowRepayHistoryV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccBorrowed) {
		toSerialize["accBorrowed"] = o.AccBorrowed
	}
	if !IsNil(o.Amt) {
		toSerialize["amt"] = o.Amt
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableGetAccountSpotBorrowRepayHistoryV5RespDataInner struct {
	value *GetAccountSpotBorrowRepayHistoryV5RespDataInner
	isSet bool
}

func (v NullableGetAccountSpotBorrowRepayHistoryV5RespDataInner) Get() *GetAccountSpotBorrowRepayHistoryV5RespDataInner {
	return v.value
}

func (v *NullableGetAccountSpotBorrowRepayHistoryV5RespDataInner) Set(val *GetAccountSpotBorrowRepayHistoryV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountSpotBorrowRepayHistoryV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountSpotBorrowRepayHistoryV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountSpotBorrowRepayHistoryV5RespDataInner(val *GetAccountSpotBorrowRepayHistoryV5RespDataInner) *NullableGetAccountSpotBorrowRepayHistoryV5RespDataInner {
	return &NullableGetAccountSpotBorrowRepayHistoryV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetAccountSpotBorrowRepayHistoryV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountSpotBorrowRepayHistoryV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


