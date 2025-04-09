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

// checks if the CreateAccountBillsHistoryArchiveV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAccountBillsHistoryArchiveV5Req{}

// CreateAccountBillsHistoryArchiveV5Req struct for CreateAccountBillsHistoryArchiveV5Req
type CreateAccountBillsHistoryArchiveV5Req struct {
	// Quarter, valid value is `Q1`, `Q2`, `Q3`, `Q4`
	Quarter string `json:"quarter"`
	// 4 digits year
	Year string `json:"year"`
}

type _CreateAccountBillsHistoryArchiveV5Req CreateAccountBillsHistoryArchiveV5Req

// NewCreateAccountBillsHistoryArchiveV5Req instantiates a new CreateAccountBillsHistoryArchiveV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccountBillsHistoryArchiveV5Req(quarter string, year string) *CreateAccountBillsHistoryArchiveV5Req {
	this := CreateAccountBillsHistoryArchiveV5Req{}
	this.Quarter = quarter
	this.Year = year
	return &this
}

// NewCreateAccountBillsHistoryArchiveV5ReqWithDefaults instantiates a new CreateAccountBillsHistoryArchiveV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountBillsHistoryArchiveV5ReqWithDefaults() *CreateAccountBillsHistoryArchiveV5Req {
	this := CreateAccountBillsHistoryArchiveV5Req{}
	var quarter string = ""
	this.Quarter = quarter
	var year string = ""
	this.Year = year
	return &this
}

// GetQuarter returns the Quarter field value
func (o *CreateAccountBillsHistoryArchiveV5Req) GetQuarter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quarter
}

// GetQuarterOk returns a tuple with the Quarter field value
// and a boolean to check if the value has been set.
func (o *CreateAccountBillsHistoryArchiveV5Req) GetQuarterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quarter, true
}

// SetQuarter sets field value
func (o *CreateAccountBillsHistoryArchiveV5Req) SetQuarter(v string) {
	o.Quarter = v
}

// GetYear returns the Year field value
func (o *CreateAccountBillsHistoryArchiveV5Req) GetYear() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Year
}

// GetYearOk returns a tuple with the Year field value
// and a boolean to check if the value has been set.
func (o *CreateAccountBillsHistoryArchiveV5Req) GetYearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Year, true
}

// SetYear sets field value
func (o *CreateAccountBillsHistoryArchiveV5Req) SetYear(v string) {
	o.Year = v
}

func (o CreateAccountBillsHistoryArchiveV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAccountBillsHistoryArchiveV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["quarter"] = o.Quarter
	toSerialize["year"] = o.Year
	return toSerialize, nil
}

func (o *CreateAccountBillsHistoryArchiveV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"quarter",
		"year",
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

	varCreateAccountBillsHistoryArchiveV5Req := _CreateAccountBillsHistoryArchiveV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateAccountBillsHistoryArchiveV5Req)

	if err != nil {
		return err
	}

	*o = CreateAccountBillsHistoryArchiveV5Req(varCreateAccountBillsHistoryArchiveV5Req)

	return err
}

type NullableCreateAccountBillsHistoryArchiveV5Req struct {
	value *CreateAccountBillsHistoryArchiveV5Req
	isSet bool
}

func (v NullableCreateAccountBillsHistoryArchiveV5Req) Get() *CreateAccountBillsHistoryArchiveV5Req {
	return v.value
}

func (v *NullableCreateAccountBillsHistoryArchiveV5Req) Set(val *CreateAccountBillsHistoryArchiveV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccountBillsHistoryArchiveV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccountBillsHistoryArchiveV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccountBillsHistoryArchiveV5Req(val *CreateAccountBillsHistoryArchiveV5Req) *NullableCreateAccountBillsHistoryArchiveV5Req {
	return &NullableCreateAccountBillsHistoryArchiveV5Req{value: val, isSet: true}
}

func (v NullableCreateAccountBillsHistoryArchiveV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccountBillsHistoryArchiveV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


