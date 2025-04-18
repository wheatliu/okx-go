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

// checks if the GetUsersPartnerIfRebateV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUsersPartnerIfRebateV5RespDataInner{}

// GetUsersPartnerIfRebateV5RespDataInner struct for GetUsersPartnerIfRebateV5RespDataInner
type GetUsersPartnerIfRebateV5RespDataInner struct {
	// Whether the user is invited by the current affiliate. `true`, `false`
	Result *bool `json:"result,omitempty"`
	// Whether there is affiliate rebate.  `0` There is affiliate rebate  `1` There is no affiliate rebate. Because the account which is requesting this endpoint is not affiliate   `2` There is no affiliate rebate. Because there is no relationship of invitation or recall, e.g. api key does not exist   `4` There is no affiliate rebate. Because the user level is equal to or more than VIP6
	Type *string `json:"type,omitempty"`
}

// NewGetUsersPartnerIfRebateV5RespDataInner instantiates a new GetUsersPartnerIfRebateV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUsersPartnerIfRebateV5RespDataInner() *GetUsersPartnerIfRebateV5RespDataInner {
	this := GetUsersPartnerIfRebateV5RespDataInner{}
	var type_ string = ""
	this.Type = &type_
	return &this
}

// NewGetUsersPartnerIfRebateV5RespDataInnerWithDefaults instantiates a new GetUsersPartnerIfRebateV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUsersPartnerIfRebateV5RespDataInnerWithDefaults() *GetUsersPartnerIfRebateV5RespDataInner {
	this := GetUsersPartnerIfRebateV5RespDataInner{}
	var type_ string = ""
	this.Type = &type_
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetUsersPartnerIfRebateV5RespDataInner) GetResult() bool {
	if o == nil || IsNil(o.Result) {
		var ret bool
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsersPartnerIfRebateV5RespDataInner) GetResultOk() (*bool, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetUsersPartnerIfRebateV5RespDataInner) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given bool and assigns it to the Result field.
func (o *GetUsersPartnerIfRebateV5RespDataInner) SetResult(v bool) {
	o.Result = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetUsersPartnerIfRebateV5RespDataInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsersPartnerIfRebateV5RespDataInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetUsersPartnerIfRebateV5RespDataInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetUsersPartnerIfRebateV5RespDataInner) SetType(v string) {
	o.Type = &v
}

func (o GetUsersPartnerIfRebateV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUsersPartnerIfRebateV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableGetUsersPartnerIfRebateV5RespDataInner struct {
	value *GetUsersPartnerIfRebateV5RespDataInner
	isSet bool
}

func (v NullableGetUsersPartnerIfRebateV5RespDataInner) Get() *GetUsersPartnerIfRebateV5RespDataInner {
	return v.value
}

func (v *NullableGetUsersPartnerIfRebateV5RespDataInner) Set(val *GetUsersPartnerIfRebateV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUsersPartnerIfRebateV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUsersPartnerIfRebateV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUsersPartnerIfRebateV5RespDataInner(val *GetUsersPartnerIfRebateV5RespDataInner) *NullableGetUsersPartnerIfRebateV5RespDataInner {
	return &NullableGetUsersPartnerIfRebateV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetUsersPartnerIfRebateV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUsersPartnerIfRebateV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


