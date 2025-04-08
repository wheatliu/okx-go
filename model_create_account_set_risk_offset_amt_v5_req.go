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

// checks if the CreateAccountSetRiskOffsetAmtV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAccountSetRiskOffsetAmtV5Req{}

// CreateAccountSetRiskOffsetAmtV5Req struct for CreateAccountSetRiskOffsetAmtV5Req
type CreateAccountSetRiskOffsetAmtV5Req struct {
	// Currency
	Ccy string `json:"ccy"`
	// Spot risk offset amount defined by users
	ClSpotInUseAmt string `json:"clSpotInUseAmt"`
}

type _CreateAccountSetRiskOffsetAmtV5Req CreateAccountSetRiskOffsetAmtV5Req

// NewCreateAccountSetRiskOffsetAmtV5Req instantiates a new CreateAccountSetRiskOffsetAmtV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccountSetRiskOffsetAmtV5Req(ccy string, clSpotInUseAmt string) *CreateAccountSetRiskOffsetAmtV5Req {
	this := CreateAccountSetRiskOffsetAmtV5Req{}
	this.Ccy = ccy
	this.ClSpotInUseAmt = clSpotInUseAmt
	return &this
}

// NewCreateAccountSetRiskOffsetAmtV5ReqWithDefaults instantiates a new CreateAccountSetRiskOffsetAmtV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountSetRiskOffsetAmtV5ReqWithDefaults() *CreateAccountSetRiskOffsetAmtV5Req {
	this := CreateAccountSetRiskOffsetAmtV5Req{}
	var ccy string = ""
	this.Ccy = ccy
	var clSpotInUseAmt string = ""
	this.ClSpotInUseAmt = clSpotInUseAmt
	return &this
}

// GetCcy returns the Ccy field value
func (o *CreateAccountSetRiskOffsetAmtV5Req) GetCcy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value
// and a boolean to check if the value has been set.
func (o *CreateAccountSetRiskOffsetAmtV5Req) GetCcyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ccy, true
}

// SetCcy sets field value
func (o *CreateAccountSetRiskOffsetAmtV5Req) SetCcy(v string) {
	o.Ccy = v
}

// GetClSpotInUseAmt returns the ClSpotInUseAmt field value
func (o *CreateAccountSetRiskOffsetAmtV5Req) GetClSpotInUseAmt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClSpotInUseAmt
}

// GetClSpotInUseAmtOk returns a tuple with the ClSpotInUseAmt field value
// and a boolean to check if the value has been set.
func (o *CreateAccountSetRiskOffsetAmtV5Req) GetClSpotInUseAmtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClSpotInUseAmt, true
}

// SetClSpotInUseAmt sets field value
func (o *CreateAccountSetRiskOffsetAmtV5Req) SetClSpotInUseAmt(v string) {
	o.ClSpotInUseAmt = v
}

func (o CreateAccountSetRiskOffsetAmtV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAccountSetRiskOffsetAmtV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ccy"] = o.Ccy
	toSerialize["clSpotInUseAmt"] = o.ClSpotInUseAmt
	return toSerialize, nil
}

func (o *CreateAccountSetRiskOffsetAmtV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ccy",
		"clSpotInUseAmt",
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

	varCreateAccountSetRiskOffsetAmtV5Req := _CreateAccountSetRiskOffsetAmtV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateAccountSetRiskOffsetAmtV5Req)

	if err != nil {
		return err
	}

	*o = CreateAccountSetRiskOffsetAmtV5Req(varCreateAccountSetRiskOffsetAmtV5Req)

	return err
}

type NullableCreateAccountSetRiskOffsetAmtV5Req struct {
	value *CreateAccountSetRiskOffsetAmtV5Req
	isSet bool
}

func (v NullableCreateAccountSetRiskOffsetAmtV5Req) Get() *CreateAccountSetRiskOffsetAmtV5Req {
	return v.value
}

func (v *NullableCreateAccountSetRiskOffsetAmtV5Req) Set(val *CreateAccountSetRiskOffsetAmtV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccountSetRiskOffsetAmtV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccountSetRiskOffsetAmtV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccountSetRiskOffsetAmtV5Req(val *CreateAccountSetRiskOffsetAmtV5Req) *NullableCreateAccountSetRiskOffsetAmtV5Req {
	return &NullableCreateAccountSetRiskOffsetAmtV5Req{value: val, isSet: true}
}

func (v NullableCreateAccountSetRiskOffsetAmtV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccountSetRiskOffsetAmtV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


