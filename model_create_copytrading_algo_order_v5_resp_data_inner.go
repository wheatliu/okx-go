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

// checks if the CreateCopytradingAlgoOrderV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCopytradingAlgoOrderV5RespDataInner{}

// CreateCopytradingAlgoOrderV5RespDataInner struct for CreateCopytradingAlgoOrderV5RespDataInner
type CreateCopytradingAlgoOrderV5RespDataInner struct {
	// Lead position ID
	SubPosId *string `json:"subPosId,omitempty"`
	// Order tag
	Tag *string `json:"tag,omitempty"`
}

// NewCreateCopytradingAlgoOrderV5RespDataInner instantiates a new CreateCopytradingAlgoOrderV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCopytradingAlgoOrderV5RespDataInner() *CreateCopytradingAlgoOrderV5RespDataInner {
	this := CreateCopytradingAlgoOrderV5RespDataInner{}
	var subPosId string = ""
	this.SubPosId = &subPosId
	var tag string = ""
	this.Tag = &tag
	return &this
}

// NewCreateCopytradingAlgoOrderV5RespDataInnerWithDefaults instantiates a new CreateCopytradingAlgoOrderV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCopytradingAlgoOrderV5RespDataInnerWithDefaults() *CreateCopytradingAlgoOrderV5RespDataInner {
	this := CreateCopytradingAlgoOrderV5RespDataInner{}
	var subPosId string = ""
	this.SubPosId = &subPosId
	var tag string = ""
	this.Tag = &tag
	return &this
}

// GetSubPosId returns the SubPosId field value if set, zero value otherwise.
func (o *CreateCopytradingAlgoOrderV5RespDataInner) GetSubPosId() string {
	if o == nil || IsNil(o.SubPosId) {
		var ret string
		return ret
	}
	return *o.SubPosId
}

// GetSubPosIdOk returns a tuple with the SubPosId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAlgoOrderV5RespDataInner) GetSubPosIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubPosId) {
		return nil, false
	}
	return o.SubPosId, true
}

// HasSubPosId returns a boolean if a field has been set.
func (o *CreateCopytradingAlgoOrderV5RespDataInner) HasSubPosId() bool {
	if o != nil && !IsNil(o.SubPosId) {
		return true
	}

	return false
}

// SetSubPosId gets a reference to the given string and assigns it to the SubPosId field.
func (o *CreateCopytradingAlgoOrderV5RespDataInner) SetSubPosId(v string) {
	o.SubPosId = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *CreateCopytradingAlgoOrderV5RespDataInner) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAlgoOrderV5RespDataInner) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *CreateCopytradingAlgoOrderV5RespDataInner) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *CreateCopytradingAlgoOrderV5RespDataInner) SetTag(v string) {
	o.Tag = &v
}

func (o CreateCopytradingAlgoOrderV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCopytradingAlgoOrderV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubPosId) {
		toSerialize["subPosId"] = o.SubPosId
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return toSerialize, nil
}

type NullableCreateCopytradingAlgoOrderV5RespDataInner struct {
	value *CreateCopytradingAlgoOrderV5RespDataInner
	isSet bool
}

func (v NullableCreateCopytradingAlgoOrderV5RespDataInner) Get() *CreateCopytradingAlgoOrderV5RespDataInner {
	return v.value
}

func (v *NullableCreateCopytradingAlgoOrderV5RespDataInner) Set(val *CreateCopytradingAlgoOrderV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCopytradingAlgoOrderV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCopytradingAlgoOrderV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCopytradingAlgoOrderV5RespDataInner(val *CreateCopytradingAlgoOrderV5RespDataInner) *NullableCreateCopytradingAlgoOrderV5RespDataInner {
	return &NullableCreateCopytradingAlgoOrderV5RespDataInner{value: val, isSet: true}
}

func (v NullableCreateCopytradingAlgoOrderV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCopytradingAlgoOrderV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


