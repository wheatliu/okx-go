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

// checks if the GetFinanceStakingDefiOrdersHistoryV5Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinanceStakingDefiOrdersHistoryV5Resp{}

// GetFinanceStakingDefiOrdersHistoryV5Resp struct for GetFinanceStakingDefiOrdersHistoryV5Resp
type GetFinanceStakingDefiOrdersHistoryV5Resp struct {
	Code *string `json:"code,omitempty"`
	Data *GetFinanceStakingDefiOrdersHistoryV5RespData `json:"data,omitempty"`
	Msg *string `json:"msg,omitempty"`
}

// NewGetFinanceStakingDefiOrdersHistoryV5Resp instantiates a new GetFinanceStakingDefiOrdersHistoryV5Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinanceStakingDefiOrdersHistoryV5Resp() *GetFinanceStakingDefiOrdersHistoryV5Resp {
	this := GetFinanceStakingDefiOrdersHistoryV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// NewGetFinanceStakingDefiOrdersHistoryV5RespWithDefaults instantiates a new GetFinanceStakingDefiOrdersHistoryV5Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinanceStakingDefiOrdersHistoryV5RespWithDefaults() *GetFinanceStakingDefiOrdersHistoryV5Resp {
	this := GetFinanceStakingDefiOrdersHistoryV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersHistoryV5Resp) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersHistoryV5Resp) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersHistoryV5Resp) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetFinanceStakingDefiOrdersHistoryV5Resp) SetCode(v string) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersHistoryV5Resp) GetData() GetFinanceStakingDefiOrdersHistoryV5RespData {
	if o == nil || IsNil(o.Data) {
		var ret GetFinanceStakingDefiOrdersHistoryV5RespData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersHistoryV5Resp) GetDataOk() (*GetFinanceStakingDefiOrdersHistoryV5RespData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersHistoryV5Resp) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetFinanceStakingDefiOrdersHistoryV5RespData and assigns it to the Data field.
func (o *GetFinanceStakingDefiOrdersHistoryV5Resp) SetData(v GetFinanceStakingDefiOrdersHistoryV5RespData) {
	o.Data = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersHistoryV5Resp) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersHistoryV5Resp) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersHistoryV5Resp) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *GetFinanceStakingDefiOrdersHistoryV5Resp) SetMsg(v string) {
	o.Msg = &v
}

func (o GetFinanceStakingDefiOrdersHistoryV5Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinanceStakingDefiOrdersHistoryV5Resp) ToMap() (map[string]interface{}, error) {
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

type NullableGetFinanceStakingDefiOrdersHistoryV5Resp struct {
	value *GetFinanceStakingDefiOrdersHistoryV5Resp
	isSet bool
}

func (v NullableGetFinanceStakingDefiOrdersHistoryV5Resp) Get() *GetFinanceStakingDefiOrdersHistoryV5Resp {
	return v.value
}

func (v *NullableGetFinanceStakingDefiOrdersHistoryV5Resp) Set(val *GetFinanceStakingDefiOrdersHistoryV5Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinanceStakingDefiOrdersHistoryV5Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinanceStakingDefiOrdersHistoryV5Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinanceStakingDefiOrdersHistoryV5Resp(val *GetFinanceStakingDefiOrdersHistoryV5Resp) *NullableGetFinanceStakingDefiOrdersHistoryV5Resp {
	return &NullableGetFinanceStakingDefiOrdersHistoryV5Resp{value: val, isSet: true}
}

func (v NullableGetFinanceStakingDefiOrdersHistoryV5Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinanceStakingDefiOrdersHistoryV5Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


