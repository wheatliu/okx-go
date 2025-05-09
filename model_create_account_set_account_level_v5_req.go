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

// checks if the CreateAccountSetAccountLevelV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAccountSetAccountLevelV5Req{}

// CreateAccountSetAccountLevelV5Req struct for CreateAccountSetAccountLevelV5Req
type CreateAccountSetAccountLevelV5Req struct {
	// Account mode  `1`: Spot mode  `2`: Spot and futures mode   `3`: Multi-currency margin code   `4`: Portfolio margin mode
	AcctLv string `json:"acctLv"`
}

type _CreateAccountSetAccountLevelV5Req CreateAccountSetAccountLevelV5Req

// NewCreateAccountSetAccountLevelV5Req instantiates a new CreateAccountSetAccountLevelV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccountSetAccountLevelV5Req(acctLv string) *CreateAccountSetAccountLevelV5Req {
	this := CreateAccountSetAccountLevelV5Req{}
	this.AcctLv = acctLv
	return &this
}

// NewCreateAccountSetAccountLevelV5ReqWithDefaults instantiates a new CreateAccountSetAccountLevelV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountSetAccountLevelV5ReqWithDefaults() *CreateAccountSetAccountLevelV5Req {
	this := CreateAccountSetAccountLevelV5Req{}
	var acctLv string = ""
	this.AcctLv = acctLv
	return &this
}

// GetAcctLv returns the AcctLv field value
func (o *CreateAccountSetAccountLevelV5Req) GetAcctLv() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AcctLv
}

// GetAcctLvOk returns a tuple with the AcctLv field value
// and a boolean to check if the value has been set.
func (o *CreateAccountSetAccountLevelV5Req) GetAcctLvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcctLv, true
}

// SetAcctLv sets field value
func (o *CreateAccountSetAccountLevelV5Req) SetAcctLv(v string) {
	o.AcctLv = v
}

func (o CreateAccountSetAccountLevelV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAccountSetAccountLevelV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["acctLv"] = o.AcctLv
	return toSerialize, nil
}

func (o *CreateAccountSetAccountLevelV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"acctLv",
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

	varCreateAccountSetAccountLevelV5Req := _CreateAccountSetAccountLevelV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateAccountSetAccountLevelV5Req)

	if err != nil {
		return err
	}

	*o = CreateAccountSetAccountLevelV5Req(varCreateAccountSetAccountLevelV5Req)

	return err
}

type NullableCreateAccountSetAccountLevelV5Req struct {
	value *CreateAccountSetAccountLevelV5Req
	isSet bool
}

func (v NullableCreateAccountSetAccountLevelV5Req) Get() *CreateAccountSetAccountLevelV5Req {
	return v.value
}

func (v *NullableCreateAccountSetAccountLevelV5Req) Set(val *CreateAccountSetAccountLevelV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccountSetAccountLevelV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccountSetAccountLevelV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccountSetAccountLevelV5Req(val *CreateAccountSetAccountLevelV5Req) *NullableCreateAccountSetAccountLevelV5Req {
	return &NullableCreateAccountSetAccountLevelV5Req{value: val, isSet: true}
}

func (v NullableCreateAccountSetAccountLevelV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccountSetAccountLevelV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


