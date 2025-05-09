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

// checks if the GetAccountQuickMarginBorrowRepayHistoryV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountQuickMarginBorrowRepayHistoryV5RespData{}

// GetAccountQuickMarginBorrowRepayHistoryV5RespData struct for GetAccountQuickMarginBorrowRepayHistoryV5RespData
type GetAccountQuickMarginBorrowRepayHistoryV5RespData struct {
	// Accumulate borrow amount
	AccBorrowed *string `json:"accBorrowed,omitempty"`
	// borrow/repay amount
	Amt *string `json:"amt,omitempty"`
	// Loan currency, e.g. `BTC`
	Ccy *string `json:"ccy,omitempty"`
	// Instrument ID, e.g. BTC-USDT
	InstId *string `json:"instId,omitempty"`
	// The ID of borrowing or repayment
	RefId *string `json:"refId,omitempty"`
	// `borrow`  `repay`
	Side *string `json:"side,omitempty"`
	// Timestamp for Borrow/Repay
	Ts *string `json:"ts,omitempty"`
}

// NewGetAccountQuickMarginBorrowRepayHistoryV5RespData instantiates a new GetAccountQuickMarginBorrowRepayHistoryV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountQuickMarginBorrowRepayHistoryV5RespData() *GetAccountQuickMarginBorrowRepayHistoryV5RespData {
	this := GetAccountQuickMarginBorrowRepayHistoryV5RespData{}
	var accBorrowed string = ""
	this.AccBorrowed = &accBorrowed
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	var instId string = ""
	this.InstId = &instId
	var refId string = ""
	this.RefId = &refId
	var side string = ""
	this.Side = &side
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewGetAccountQuickMarginBorrowRepayHistoryV5RespDataWithDefaults instantiates a new GetAccountQuickMarginBorrowRepayHistoryV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountQuickMarginBorrowRepayHistoryV5RespDataWithDefaults() *GetAccountQuickMarginBorrowRepayHistoryV5RespData {
	this := GetAccountQuickMarginBorrowRepayHistoryV5RespData{}
	var accBorrowed string = ""
	this.AccBorrowed = &accBorrowed
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	var instId string = ""
	this.InstId = &instId
	var refId string = ""
	this.RefId = &refId
	var side string = ""
	this.Side = &side
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetAccBorrowed returns the AccBorrowed field value if set, zero value otherwise.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) GetAccBorrowed() string {
	if o == nil || IsNil(o.AccBorrowed) {
		var ret string
		return ret
	}
	return *o.AccBorrowed
}

// GetAccBorrowedOk returns a tuple with the AccBorrowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) GetAccBorrowedOk() (*string, bool) {
	if o == nil || IsNil(o.AccBorrowed) {
		return nil, false
	}
	return o.AccBorrowed, true
}

// HasAccBorrowed returns a boolean if a field has been set.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) HasAccBorrowed() bool {
	if o != nil && !IsNil(o.AccBorrowed) {
		return true
	}

	return false
}

// SetAccBorrowed gets a reference to the given string and assigns it to the AccBorrowed field.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) SetAccBorrowed(v string) {
	o.AccBorrowed = &v
}

// GetAmt returns the Amt field value if set, zero value otherwise.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) GetAmt() string {
	if o == nil || IsNil(o.Amt) {
		var ret string
		return ret
	}
	return *o.Amt
}

// GetAmtOk returns a tuple with the Amt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) GetAmtOk() (*string, bool) {
	if o == nil || IsNil(o.Amt) {
		return nil, false
	}
	return o.Amt, true
}

// HasAmt returns a boolean if a field has been set.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) HasAmt() bool {
	if o != nil && !IsNil(o.Amt) {
		return true
	}

	return false
}

// SetAmt gets a reference to the given string and assigns it to the Amt field.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) SetAmt(v string) {
	o.Amt = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) SetCcy(v string) {
	o.Ccy = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) SetInstId(v string) {
	o.InstId = &v
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) GetRefId() string {
	if o == nil || IsNil(o.RefId) {
		var ret string
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) GetRefIdOk() (*string, bool) {
	if o == nil || IsNil(o.RefId) {
		return nil, false
	}
	return o.RefId, true
}

// HasRefId returns a boolean if a field has been set.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) HasRefId() bool {
	if o != nil && !IsNil(o.RefId) {
		return true
	}

	return false
}

// SetRefId gets a reference to the given string and assigns it to the RefId field.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) SetRefId(v string) {
	o.RefId = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) SetSide(v string) {
	o.Side = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespData) SetTs(v string) {
	o.Ts = &v
}

func (o GetAccountQuickMarginBorrowRepayHistoryV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountQuickMarginBorrowRepayHistoryV5RespData) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	if !IsNil(o.RefId) {
		toSerialize["refId"] = o.RefId
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableGetAccountQuickMarginBorrowRepayHistoryV5RespData struct {
	value *GetAccountQuickMarginBorrowRepayHistoryV5RespData
	isSet bool
}

func (v NullableGetAccountQuickMarginBorrowRepayHistoryV5RespData) Get() *GetAccountQuickMarginBorrowRepayHistoryV5RespData {
	return v.value
}

func (v *NullableGetAccountQuickMarginBorrowRepayHistoryV5RespData) Set(val *GetAccountQuickMarginBorrowRepayHistoryV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountQuickMarginBorrowRepayHistoryV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountQuickMarginBorrowRepayHistoryV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountQuickMarginBorrowRepayHistoryV5RespData(val *GetAccountQuickMarginBorrowRepayHistoryV5RespData) *NullableGetAccountQuickMarginBorrowRepayHistoryV5RespData {
	return &NullableGetAccountQuickMarginBorrowRepayHistoryV5RespData{value: val, isSet: true}
}

func (v NullableGetAccountQuickMarginBorrowRepayHistoryV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountQuickMarginBorrowRepayHistoryV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


