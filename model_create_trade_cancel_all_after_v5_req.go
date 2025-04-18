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

// checks if the CreateTradeCancelAllAfterV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradeCancelAllAfterV5Req{}

// CreateTradeCancelAllAfterV5Req struct for CreateTradeCancelAllAfterV5Req
type CreateTradeCancelAllAfterV5Req struct {
	// CAA order tag   A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 16 characters.
	Tag *string `json:"tag,omitempty"`
	// The countdown for order cancellation, with second as the unit.  Range of value can be 0, [10, 120].   Setting timeOut to 0 disables Cancel All After.
	TimeOut string `json:"timeOut"`
}

type _CreateTradeCancelAllAfterV5Req CreateTradeCancelAllAfterV5Req

// NewCreateTradeCancelAllAfterV5Req instantiates a new CreateTradeCancelAllAfterV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradeCancelAllAfterV5Req(timeOut string) *CreateTradeCancelAllAfterV5Req {
	this := CreateTradeCancelAllAfterV5Req{}
	var tag string = ""
	this.Tag = &tag
	this.TimeOut = timeOut
	return &this
}

// NewCreateTradeCancelAllAfterV5ReqWithDefaults instantiates a new CreateTradeCancelAllAfterV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradeCancelAllAfterV5ReqWithDefaults() *CreateTradeCancelAllAfterV5Req {
	this := CreateTradeCancelAllAfterV5Req{}
	var tag string = ""
	this.Tag = &tag
	var timeOut string = ""
	this.TimeOut = timeOut
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *CreateTradeCancelAllAfterV5Req) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelAllAfterV5Req) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *CreateTradeCancelAllAfterV5Req) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *CreateTradeCancelAllAfterV5Req) SetTag(v string) {
	o.Tag = &v
}

// GetTimeOut returns the TimeOut field value
func (o *CreateTradeCancelAllAfterV5Req) GetTimeOut() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeOut
}

// GetTimeOutOk returns a tuple with the TimeOut field value
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelAllAfterV5Req) GetTimeOutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeOut, true
}

// SetTimeOut sets field value
func (o *CreateTradeCancelAllAfterV5Req) SetTimeOut(v string) {
	o.TimeOut = v
}

func (o CreateTradeCancelAllAfterV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradeCancelAllAfterV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	toSerialize["timeOut"] = o.TimeOut
	return toSerialize, nil
}

func (o *CreateTradeCancelAllAfterV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"timeOut",
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

	varCreateTradeCancelAllAfterV5Req := _CreateTradeCancelAllAfterV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTradeCancelAllAfterV5Req)

	if err != nil {
		return err
	}

	*o = CreateTradeCancelAllAfterV5Req(varCreateTradeCancelAllAfterV5Req)

	return err
}

type NullableCreateTradeCancelAllAfterV5Req struct {
	value *CreateTradeCancelAllAfterV5Req
	isSet bool
}

func (v NullableCreateTradeCancelAllAfterV5Req) Get() *CreateTradeCancelAllAfterV5Req {
	return v.value
}

func (v *NullableCreateTradeCancelAllAfterV5Req) Set(val *CreateTradeCancelAllAfterV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradeCancelAllAfterV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradeCancelAllAfterV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradeCancelAllAfterV5Req(val *CreateTradeCancelAllAfterV5Req) *NullableCreateTradeCancelAllAfterV5Req {
	return &NullableCreateTradeCancelAllAfterV5Req{value: val, isSet: true}
}

func (v NullableCreateTradeCancelAllAfterV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradeCancelAllAfterV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


