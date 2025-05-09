/*
OKX v5 API

OpenAPI specification for Okx exchange - Rest API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rest

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateAccountSetIsolatedModeV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAccountSetIsolatedModeV5Req{}

// CreateAccountSetIsolatedModeV5Req struct for CreateAccountSetIsolatedModeV5Req
type CreateAccountSetIsolatedModeV5Req struct {
	// Isolated margin trading settings   `auto_transfers_ccy`: New auto transfers, enabling both base and quote currency as the margin for isolated margin trading. Only applicable to `MARGIN`.  `automatic`: Auto transfers
	IsoMode string `json:"isoMode"`
	// Instrument type  `MARGIN`  `CONTRACTS`
	Type string `json:"type"`
}

type _CreateAccountSetIsolatedModeV5Req CreateAccountSetIsolatedModeV5Req

// NewCreateAccountSetIsolatedModeV5Req instantiates a new CreateAccountSetIsolatedModeV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccountSetIsolatedModeV5Req(isoMode string, type_ string) *CreateAccountSetIsolatedModeV5Req {
	this := CreateAccountSetIsolatedModeV5Req{}
	this.IsoMode = isoMode
	this.Type = type_
	return &this
}

// NewCreateAccountSetIsolatedModeV5ReqWithDefaults instantiates a new CreateAccountSetIsolatedModeV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountSetIsolatedModeV5ReqWithDefaults() *CreateAccountSetIsolatedModeV5Req {
	this := CreateAccountSetIsolatedModeV5Req{}
	var isoMode string = ""
	this.IsoMode = isoMode
	var type_ string = ""
	this.Type = type_
	return &this
}

// GetIsoMode returns the IsoMode field value
func (o *CreateAccountSetIsolatedModeV5Req) GetIsoMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoMode
}

// GetIsoModeOk returns a tuple with the IsoMode field value
// and a boolean to check if the value has been set.
func (o *CreateAccountSetIsolatedModeV5Req) GetIsoModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsoMode, true
}

// SetIsoMode sets field value
func (o *CreateAccountSetIsolatedModeV5Req) SetIsoMode(v string) {
	o.IsoMode = v
}

// GetType returns the Type field value
func (o *CreateAccountSetIsolatedModeV5Req) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateAccountSetIsolatedModeV5Req) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateAccountSetIsolatedModeV5Req) SetType(v string) {
	o.Type = v
}

func (o CreateAccountSetIsolatedModeV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAccountSetIsolatedModeV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isoMode"] = o.IsoMode
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *CreateAccountSetIsolatedModeV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isoMode",
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

	varCreateAccountSetIsolatedModeV5Req := _CreateAccountSetIsolatedModeV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateAccountSetIsolatedModeV5Req)

	if err != nil {
		return err
	}

	*o = CreateAccountSetIsolatedModeV5Req(varCreateAccountSetIsolatedModeV5Req)

	return err
}

type NullableCreateAccountSetIsolatedModeV5Req struct {
	value *CreateAccountSetIsolatedModeV5Req
	isSet bool
}

func (v NullableCreateAccountSetIsolatedModeV5Req) Get() *CreateAccountSetIsolatedModeV5Req {
	return v.value
}

func (v *NullableCreateAccountSetIsolatedModeV5Req) Set(val *CreateAccountSetIsolatedModeV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccountSetIsolatedModeV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccountSetIsolatedModeV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccountSetIsolatedModeV5Req(val *CreateAccountSetIsolatedModeV5Req) *NullableCreateAccountSetIsolatedModeV5Req {
	return &NullableCreateAccountSetIsolatedModeV5Req{value: val, isSet: true}
}

func (v NullableCreateAccountSetIsolatedModeV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccountSetIsolatedModeV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


