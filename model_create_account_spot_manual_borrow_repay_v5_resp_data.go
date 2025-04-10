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

// checks if the CreateAccountSpotManualBorrowRepayV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAccountSpotManualBorrowRepayV5RespData{}

// CreateAccountSpotManualBorrowRepayV5RespData struct for CreateAccountSpotManualBorrowRepayV5RespData
type CreateAccountSpotManualBorrowRepayV5RespData struct {
	// Actual amount
	Amt *string `json:"amt,omitempty"`
	// Currency, e.g. `BTC`
	Ccy *string `json:"ccy,omitempty"`
	// Side  `borrow`  `repay`
	Side *string `json:"side,omitempty"`
}

// NewCreateAccountSpotManualBorrowRepayV5RespData instantiates a new CreateAccountSpotManualBorrowRepayV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccountSpotManualBorrowRepayV5RespData() *CreateAccountSpotManualBorrowRepayV5RespData {
	this := CreateAccountSpotManualBorrowRepayV5RespData{}
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	var side string = ""
	this.Side = &side
	return &this
}

// NewCreateAccountSpotManualBorrowRepayV5RespDataWithDefaults instantiates a new CreateAccountSpotManualBorrowRepayV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountSpotManualBorrowRepayV5RespDataWithDefaults() *CreateAccountSpotManualBorrowRepayV5RespData {
	this := CreateAccountSpotManualBorrowRepayV5RespData{}
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	var side string = ""
	this.Side = &side
	return &this
}

// GetAmt returns the Amt field value if set, zero value otherwise.
func (o *CreateAccountSpotManualBorrowRepayV5RespData) GetAmt() string {
	if o == nil || IsNil(o.Amt) {
		var ret string
		return ret
	}
	return *o.Amt
}

// GetAmtOk returns a tuple with the Amt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountSpotManualBorrowRepayV5RespData) GetAmtOk() (*string, bool) {
	if o == nil || IsNil(o.Amt) {
		return nil, false
	}
	return o.Amt, true
}

// HasAmt returns a boolean if a field has been set.
func (o *CreateAccountSpotManualBorrowRepayV5RespData) HasAmt() bool {
	if o != nil && !IsNil(o.Amt) {
		return true
	}

	return false
}

// SetAmt gets a reference to the given string and assigns it to the Amt field.
func (o *CreateAccountSpotManualBorrowRepayV5RespData) SetAmt(v string) {
	o.Amt = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *CreateAccountSpotManualBorrowRepayV5RespData) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountSpotManualBorrowRepayV5RespData) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *CreateAccountSpotManualBorrowRepayV5RespData) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *CreateAccountSpotManualBorrowRepayV5RespData) SetCcy(v string) {
	o.Ccy = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *CreateAccountSpotManualBorrowRepayV5RespData) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountSpotManualBorrowRepayV5RespData) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *CreateAccountSpotManualBorrowRepayV5RespData) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *CreateAccountSpotManualBorrowRepayV5RespData) SetSide(v string) {
	o.Side = &v
}

func (o CreateAccountSpotManualBorrowRepayV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAccountSpotManualBorrowRepayV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amt) {
		toSerialize["amt"] = o.Amt
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	return toSerialize, nil
}

type NullableCreateAccountSpotManualBorrowRepayV5RespData struct {
	value *CreateAccountSpotManualBorrowRepayV5RespData
	isSet bool
}

func (v NullableCreateAccountSpotManualBorrowRepayV5RespData) Get() *CreateAccountSpotManualBorrowRepayV5RespData {
	return v.value
}

func (v *NullableCreateAccountSpotManualBorrowRepayV5RespData) Set(val *CreateAccountSpotManualBorrowRepayV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccountSpotManualBorrowRepayV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccountSpotManualBorrowRepayV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccountSpotManualBorrowRepayV5RespData(val *CreateAccountSpotManualBorrowRepayV5RespData) *NullableCreateAccountSpotManualBorrowRepayV5RespData {
	return &NullableCreateAccountSpotManualBorrowRepayV5RespData{value: val, isSet: true}
}

func (v NullableCreateAccountSpotManualBorrowRepayV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccountSpotManualBorrowRepayV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


