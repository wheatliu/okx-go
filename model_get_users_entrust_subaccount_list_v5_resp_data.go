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

// checks if the GetUsersEntrustSubaccountListV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUsersEntrustSubaccountListV5RespData{}

// GetUsersEntrustSubaccountListV5RespData struct for GetUsersEntrustSubaccountListV5RespData
type GetUsersEntrustSubaccountListV5RespData struct {
	// Sub-account name
	SubAcct *string `json:"subAcct,omitempty"`
}

// NewGetUsersEntrustSubaccountListV5RespData instantiates a new GetUsersEntrustSubaccountListV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUsersEntrustSubaccountListV5RespData() *GetUsersEntrustSubaccountListV5RespData {
	this := GetUsersEntrustSubaccountListV5RespData{}
	var subAcct string = ""
	this.SubAcct = &subAcct
	return &this
}

// NewGetUsersEntrustSubaccountListV5RespDataWithDefaults instantiates a new GetUsersEntrustSubaccountListV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUsersEntrustSubaccountListV5RespDataWithDefaults() *GetUsersEntrustSubaccountListV5RespData {
	this := GetUsersEntrustSubaccountListV5RespData{}
	var subAcct string = ""
	this.SubAcct = &subAcct
	return &this
}

// GetSubAcct returns the SubAcct field value if set, zero value otherwise.
func (o *GetUsersEntrustSubaccountListV5RespData) GetSubAcct() string {
	if o == nil || IsNil(o.SubAcct) {
		var ret string
		return ret
	}
	return *o.SubAcct
}

// GetSubAcctOk returns a tuple with the SubAcct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsersEntrustSubaccountListV5RespData) GetSubAcctOk() (*string, bool) {
	if o == nil || IsNil(o.SubAcct) {
		return nil, false
	}
	return o.SubAcct, true
}

// HasSubAcct returns a boolean if a field has been set.
func (o *GetUsersEntrustSubaccountListV5RespData) HasSubAcct() bool {
	if o != nil && !IsNil(o.SubAcct) {
		return true
	}

	return false
}

// SetSubAcct gets a reference to the given string and assigns it to the SubAcct field.
func (o *GetUsersEntrustSubaccountListV5RespData) SetSubAcct(v string) {
	o.SubAcct = &v
}

func (o GetUsersEntrustSubaccountListV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUsersEntrustSubaccountListV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubAcct) {
		toSerialize["subAcct"] = o.SubAcct
	}
	return toSerialize, nil
}

type NullableGetUsersEntrustSubaccountListV5RespData struct {
	value *GetUsersEntrustSubaccountListV5RespData
	isSet bool
}

func (v NullableGetUsersEntrustSubaccountListV5RespData) Get() *GetUsersEntrustSubaccountListV5RespData {
	return v.value
}

func (v *NullableGetUsersEntrustSubaccountListV5RespData) Set(val *GetUsersEntrustSubaccountListV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUsersEntrustSubaccountListV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUsersEntrustSubaccountListV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUsersEntrustSubaccountListV5RespData(val *GetUsersEntrustSubaccountListV5RespData) *NullableGetUsersEntrustSubaccountListV5RespData {
	return &NullableGetUsersEntrustSubaccountListV5RespData{value: val, isSet: true}
}

func (v NullableGetUsersEntrustSubaccountListV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUsersEntrustSubaccountListV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


