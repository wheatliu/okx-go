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

// checks if the APIError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIError{}

// APIError OKX API error
type APIError struct {
	Code *string `json:"code,omitempty"`
	Msg *string `json:"msg,omitempty"`
}

// NewAPIError instantiates a new APIError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIError() *APIError {
	this := APIError{}
	return &this
}

// NewAPIErrorWithDefaults instantiates a new APIError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIErrorWithDefaults() *APIError {
	this := APIError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *APIError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *APIError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *APIError) SetCode(v string) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *APIError) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIError) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *APIError) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *APIError) SetMsg(v string) {
	o.Msg = &v
}

func (o APIError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	return toSerialize, nil
}

type NullableAPIError struct {
	value *APIError
	isSet bool
}

func (v NullableAPIError) Get() *APIError {
	return v.value
}

func (v *NullableAPIError) Set(val *APIError) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIError) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIError(val *APIError) *NullableAPIError {
	return &NullableAPIError{value: val, isSet: true}
}

func (v NullableAPIError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


