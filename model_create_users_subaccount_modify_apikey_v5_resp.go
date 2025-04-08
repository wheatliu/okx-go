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

// checks if the CreateUsersSubaccountModifyApikeyV5Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUsersSubaccountModifyApikeyV5Resp{}

// CreateUsersSubaccountModifyApikeyV5Resp struct for CreateUsersSubaccountModifyApikeyV5Resp
type CreateUsersSubaccountModifyApikeyV5Resp struct {
	Code *string `json:"code,omitempty"`
	Data *CreateUsersSubaccountModifyApikeyV5RespData `json:"data,omitempty"`
	Msg *string `json:"msg,omitempty"`
}

// NewCreateUsersSubaccountModifyApikeyV5Resp instantiates a new CreateUsersSubaccountModifyApikeyV5Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUsersSubaccountModifyApikeyV5Resp() *CreateUsersSubaccountModifyApikeyV5Resp {
	this := CreateUsersSubaccountModifyApikeyV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// NewCreateUsersSubaccountModifyApikeyV5RespWithDefaults instantiates a new CreateUsersSubaccountModifyApikeyV5Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUsersSubaccountModifyApikeyV5RespWithDefaults() *CreateUsersSubaccountModifyApikeyV5Resp {
	this := CreateUsersSubaccountModifyApikeyV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CreateUsersSubaccountModifyApikeyV5Resp) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUsersSubaccountModifyApikeyV5Resp) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CreateUsersSubaccountModifyApikeyV5Resp) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CreateUsersSubaccountModifyApikeyV5Resp) SetCode(v string) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateUsersSubaccountModifyApikeyV5Resp) GetData() CreateUsersSubaccountModifyApikeyV5RespData {
	if o == nil || IsNil(o.Data) {
		var ret CreateUsersSubaccountModifyApikeyV5RespData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUsersSubaccountModifyApikeyV5Resp) GetDataOk() (*CreateUsersSubaccountModifyApikeyV5RespData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateUsersSubaccountModifyApikeyV5Resp) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CreateUsersSubaccountModifyApikeyV5RespData and assigns it to the Data field.
func (o *CreateUsersSubaccountModifyApikeyV5Resp) SetData(v CreateUsersSubaccountModifyApikeyV5RespData) {
	o.Data = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CreateUsersSubaccountModifyApikeyV5Resp) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUsersSubaccountModifyApikeyV5Resp) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *CreateUsersSubaccountModifyApikeyV5Resp) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CreateUsersSubaccountModifyApikeyV5Resp) SetMsg(v string) {
	o.Msg = &v
}

func (o CreateUsersSubaccountModifyApikeyV5Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUsersSubaccountModifyApikeyV5Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	return toSerialize, nil
}

type NullableCreateUsersSubaccountModifyApikeyV5Resp struct {
	value *CreateUsersSubaccountModifyApikeyV5Resp
	isSet bool
}

func (v NullableCreateUsersSubaccountModifyApikeyV5Resp) Get() *CreateUsersSubaccountModifyApikeyV5Resp {
	return v.value
}

func (v *NullableCreateUsersSubaccountModifyApikeyV5Resp) Set(val *CreateUsersSubaccountModifyApikeyV5Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUsersSubaccountModifyApikeyV5Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUsersSubaccountModifyApikeyV5Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUsersSubaccountModifyApikeyV5Resp(val *CreateUsersSubaccountModifyApikeyV5Resp) *NullableCreateUsersSubaccountModifyApikeyV5Resp {
	return &NullableCreateUsersSubaccountModifyApikeyV5Resp{value: val, isSet: true}
}

func (v NullableCreateUsersSubaccountModifyApikeyV5Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUsersSubaccountModifyApikeyV5Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


