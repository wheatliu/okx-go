/*
Okx Rest API

OpenAPI specification for Okx cryptocurrency exchange - Rest API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rest

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateAccountPositionMarginBalanceV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAccountPositionMarginBalanceV5Req{}

// CreateAccountPositionMarginBalanceV5Req struct for CreateAccountPositionMarginBalanceV5Req
type CreateAccountPositionMarginBalanceV5Req struct {
	// Amount to be increased or decreased.
	Amt string `json:"amt"`
	// Currency   Applicable to `isolated` `MARGIN` orders
	Ccy *string `json:"ccy,omitempty"`
	// Instrument ID
	InstId string `json:"instId"`
	// Position side, the default is `net`  `long`   `short`   `net`
	PosSide string `json:"posSide"`
	// `add`: add margin   `reduce`: reduce margin
	Type string `json:"type"`
}

type _CreateAccountPositionMarginBalanceV5Req CreateAccountPositionMarginBalanceV5Req

// NewCreateAccountPositionMarginBalanceV5Req instantiates a new CreateAccountPositionMarginBalanceV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccountPositionMarginBalanceV5Req(amt string, instId string, posSide string, type_ string) *CreateAccountPositionMarginBalanceV5Req {
	this := CreateAccountPositionMarginBalanceV5Req{}
	this.Amt = amt
	var ccy string = ""
	this.Ccy = &ccy
	this.InstId = instId
	this.PosSide = posSide
	this.Type = type_
	return &this
}

// NewCreateAccountPositionMarginBalanceV5ReqWithDefaults instantiates a new CreateAccountPositionMarginBalanceV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountPositionMarginBalanceV5ReqWithDefaults() *CreateAccountPositionMarginBalanceV5Req {
	this := CreateAccountPositionMarginBalanceV5Req{}
	var amt string = ""
	this.Amt = amt
	var ccy string = ""
	this.Ccy = &ccy
	var instId string = ""
	this.InstId = instId
	var posSide string = ""
	this.PosSide = posSide
	var type_ string = ""
	this.Type = type_
	return &this
}

// GetAmt returns the Amt field value
func (o *CreateAccountPositionMarginBalanceV5Req) GetAmt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amt
}

// GetAmtOk returns a tuple with the Amt field value
// and a boolean to check if the value has been set.
func (o *CreateAccountPositionMarginBalanceV5Req) GetAmtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amt, true
}

// SetAmt sets field value
func (o *CreateAccountPositionMarginBalanceV5Req) SetAmt(v string) {
	o.Amt = v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *CreateAccountPositionMarginBalanceV5Req) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountPositionMarginBalanceV5Req) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *CreateAccountPositionMarginBalanceV5Req) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *CreateAccountPositionMarginBalanceV5Req) SetCcy(v string) {
	o.Ccy = &v
}

// GetInstId returns the InstId field value
func (o *CreateAccountPositionMarginBalanceV5Req) GetInstId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value
// and a boolean to check if the value has been set.
func (o *CreateAccountPositionMarginBalanceV5Req) GetInstIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstId, true
}

// SetInstId sets field value
func (o *CreateAccountPositionMarginBalanceV5Req) SetInstId(v string) {
	o.InstId = v
}

// GetPosSide returns the PosSide field value
func (o *CreateAccountPositionMarginBalanceV5Req) GetPosSide() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PosSide
}

// GetPosSideOk returns a tuple with the PosSide field value
// and a boolean to check if the value has been set.
func (o *CreateAccountPositionMarginBalanceV5Req) GetPosSideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PosSide, true
}

// SetPosSide sets field value
func (o *CreateAccountPositionMarginBalanceV5Req) SetPosSide(v string) {
	o.PosSide = v
}

// GetType returns the Type field value
func (o *CreateAccountPositionMarginBalanceV5Req) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateAccountPositionMarginBalanceV5Req) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateAccountPositionMarginBalanceV5Req) SetType(v string) {
	o.Type = v
}

func (o CreateAccountPositionMarginBalanceV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAccountPositionMarginBalanceV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amt"] = o.Amt
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	toSerialize["instId"] = o.InstId
	toSerialize["posSide"] = o.PosSide
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *CreateAccountPositionMarginBalanceV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amt",
		"instId",
		"posSide",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateAccountPositionMarginBalanceV5Req := _CreateAccountPositionMarginBalanceV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateAccountPositionMarginBalanceV5Req)

	if err != nil {
		return err
	}

	*o = CreateAccountPositionMarginBalanceV5Req(varCreateAccountPositionMarginBalanceV5Req)

	return err
}

type NullableCreateAccountPositionMarginBalanceV5Req struct {
	value *CreateAccountPositionMarginBalanceV5Req
	isSet bool
}

func (v NullableCreateAccountPositionMarginBalanceV5Req) Get() *CreateAccountPositionMarginBalanceV5Req {
	return v.value
}

func (v *NullableCreateAccountPositionMarginBalanceV5Req) Set(val *CreateAccountPositionMarginBalanceV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccountPositionMarginBalanceV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccountPositionMarginBalanceV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccountPositionMarginBalanceV5Req(val *CreateAccountPositionMarginBalanceV5Req) *NullableCreateAccountPositionMarginBalanceV5Req {
	return &NullableCreateAccountPositionMarginBalanceV5Req{value: val, isSet: true}
}

func (v NullableCreateAccountPositionMarginBalanceV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccountPositionMarginBalanceV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


