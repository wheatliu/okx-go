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

// checks if the GetFinanceFlexibleLoanLoanHistoryV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinanceFlexibleLoanLoanHistoryV5RespDataInner{}

// GetFinanceFlexibleLoanLoanHistoryV5RespDataInner struct for GetFinanceFlexibleLoanLoanHistoryV5RespDataInner
type GetFinanceFlexibleLoanLoanHistoryV5RespDataInner struct {
	// Amount
	Amt *string `json:"amt,omitempty"`
	// Currency, e.g. `BTC`
	Ccy *string `json:"ccy,omitempty"`
	// Reference ID
	RefId *string `json:"refId,omitempty"`
	// Timestamp for the action, Unix timestamp format in milliseconds, e.g. `1597026383085`
	Ts *string `json:"ts,omitempty"`
	// Action type
	Type *string `json:"type,omitempty"`
}

// NewGetFinanceFlexibleLoanLoanHistoryV5RespDataInner instantiates a new GetFinanceFlexibleLoanLoanHistoryV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinanceFlexibleLoanLoanHistoryV5RespDataInner() *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner {
	this := GetFinanceFlexibleLoanLoanHistoryV5RespDataInner{}
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	var refId string = ""
	this.RefId = &refId
	var ts string = ""
	this.Ts = &ts
	var type_ string = ""
	this.Type = &type_
	return &this
}

// NewGetFinanceFlexibleLoanLoanHistoryV5RespDataInnerWithDefaults instantiates a new GetFinanceFlexibleLoanLoanHistoryV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinanceFlexibleLoanLoanHistoryV5RespDataInnerWithDefaults() *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner {
	this := GetFinanceFlexibleLoanLoanHistoryV5RespDataInner{}
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	var refId string = ""
	this.RefId = &refId
	var ts string = ""
	this.Ts = &ts
	var type_ string = ""
	this.Type = &type_
	return &this
}

// GetAmt returns the Amt field value if set, zero value otherwise.
func (o *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) GetAmt() string {
	if o == nil || IsNil(o.Amt) {
		var ret string
		return ret
	}
	return *o.Amt
}

// GetAmtOk returns a tuple with the Amt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) GetAmtOk() (*string, bool) {
	if o == nil || IsNil(o.Amt) {
		return nil, false
	}
	return o.Amt, true
}

// HasAmt returns a boolean if a field has been set.
func (o *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) HasAmt() bool {
	if o != nil && !IsNil(o.Amt) {
		return true
	}

	return false
}

// SetAmt gets a reference to the given string and assigns it to the Amt field.
func (o *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) SetAmt(v string) {
	o.Amt = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) GetRefId() string {
	if o == nil || IsNil(o.RefId) {
		var ret string
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) GetRefIdOk() (*string, bool) {
	if o == nil || IsNil(o.RefId) {
		return nil, false
	}
	return o.RefId, true
}

// HasRefId returns a boolean if a field has been set.
func (o *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) HasRefId() bool {
	if o != nil && !IsNil(o.RefId) {
		return true
	}

	return false
}

// SetRefId gets a reference to the given string and assigns it to the RefId field.
func (o *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) SetRefId(v string) {
	o.RefId = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) SetType(v string) {
	o.Type = &v
}

func (o GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amt) {
		toSerialize["amt"] = o.Amt
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.RefId) {
		toSerialize["refId"] = o.RefId
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableGetFinanceFlexibleLoanLoanHistoryV5RespDataInner struct {
	value *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner
	isSet bool
}

func (v NullableGetFinanceFlexibleLoanLoanHistoryV5RespDataInner) Get() *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner {
	return v.value
}

func (v *NullableGetFinanceFlexibleLoanLoanHistoryV5RespDataInner) Set(val *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinanceFlexibleLoanLoanHistoryV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinanceFlexibleLoanLoanHistoryV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinanceFlexibleLoanLoanHistoryV5RespDataInner(val *GetFinanceFlexibleLoanLoanHistoryV5RespDataInner) *NullableGetFinanceFlexibleLoanLoanHistoryV5RespDataInner {
	return &NullableGetFinanceFlexibleLoanLoanHistoryV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetFinanceFlexibleLoanLoanHistoryV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinanceFlexibleLoanLoanHistoryV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


