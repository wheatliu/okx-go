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

// checks if the GetFinanceSavingsLendingRateSummaryV5Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinanceSavingsLendingRateSummaryV5Resp{}

// GetFinanceSavingsLendingRateSummaryV5Resp struct for GetFinanceSavingsLendingRateSummaryV5Resp
type GetFinanceSavingsLendingRateSummaryV5Resp struct {
	Code *string `json:"code,omitempty"`
	Data []GetFinanceSavingsLendingRateSummaryV5RespDataInner `json:"data,omitempty"`
	Msg *string `json:"msg,omitempty"`
}

// NewGetFinanceSavingsLendingRateSummaryV5Resp instantiates a new GetFinanceSavingsLendingRateSummaryV5Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinanceSavingsLendingRateSummaryV5Resp() *GetFinanceSavingsLendingRateSummaryV5Resp {
	this := GetFinanceSavingsLendingRateSummaryV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// NewGetFinanceSavingsLendingRateSummaryV5RespWithDefaults instantiates a new GetFinanceSavingsLendingRateSummaryV5Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinanceSavingsLendingRateSummaryV5RespWithDefaults() *GetFinanceSavingsLendingRateSummaryV5Resp {
	this := GetFinanceSavingsLendingRateSummaryV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetFinanceSavingsLendingRateSummaryV5Resp) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceSavingsLendingRateSummaryV5Resp) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetFinanceSavingsLendingRateSummaryV5Resp) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetFinanceSavingsLendingRateSummaryV5Resp) SetCode(v string) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetFinanceSavingsLendingRateSummaryV5Resp) GetData() []GetFinanceSavingsLendingRateSummaryV5RespDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetFinanceSavingsLendingRateSummaryV5RespDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceSavingsLendingRateSummaryV5Resp) GetDataOk() ([]GetFinanceSavingsLendingRateSummaryV5RespDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetFinanceSavingsLendingRateSummaryV5Resp) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetFinanceSavingsLendingRateSummaryV5RespDataInner and assigns it to the Data field.
func (o *GetFinanceSavingsLendingRateSummaryV5Resp) SetData(v []GetFinanceSavingsLendingRateSummaryV5RespDataInner) {
	o.Data = v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *GetFinanceSavingsLendingRateSummaryV5Resp) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceSavingsLendingRateSummaryV5Resp) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *GetFinanceSavingsLendingRateSummaryV5Resp) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *GetFinanceSavingsLendingRateSummaryV5Resp) SetMsg(v string) {
	o.Msg = &v
}

func (o GetFinanceSavingsLendingRateSummaryV5Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinanceSavingsLendingRateSummaryV5Resp) ToMap() (map[string]interface{}, error) {
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

type NullableGetFinanceSavingsLendingRateSummaryV5Resp struct {
	value *GetFinanceSavingsLendingRateSummaryV5Resp
	isSet bool
}

func (v NullableGetFinanceSavingsLendingRateSummaryV5Resp) Get() *GetFinanceSavingsLendingRateSummaryV5Resp {
	return v.value
}

func (v *NullableGetFinanceSavingsLendingRateSummaryV5Resp) Set(val *GetFinanceSavingsLendingRateSummaryV5Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinanceSavingsLendingRateSummaryV5Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinanceSavingsLendingRateSummaryV5Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinanceSavingsLendingRateSummaryV5Resp(val *GetFinanceSavingsLendingRateSummaryV5Resp) *NullableGetFinanceSavingsLendingRateSummaryV5Resp {
	return &NullableGetFinanceSavingsLendingRateSummaryV5Resp{value: val, isSet: true}
}

func (v NullableGetFinanceSavingsLendingRateSummaryV5Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinanceSavingsLendingRateSummaryV5Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


