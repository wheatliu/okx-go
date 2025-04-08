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

// checks if the CreateTradeCancelAlgosV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradeCancelAlgosV5RespData{}

// CreateTradeCancelAlgosV5RespData struct for CreateTradeCancelAlgosV5RespData
type CreateTradeCancelAlgosV5RespData struct {
	// Client-supplied Algo ID(Deprecated)
	// Deprecated
	AlgoClOrdId *string `json:"algoClOrdId,omitempty"`
	// Algo ID
	AlgoId *string `json:"algoId,omitempty"`
	// Client Order ID as assigned by the client(Deprecated)
	// Deprecated
	ClOrdId *string `json:"clOrdId,omitempty"`
	// The code of the event execution result, `0` means success.
	SCode *string `json:"sCode,omitempty"`
	// Rejection message if the request is unsuccessful.
	SMsg *string `json:"sMsg,omitempty"`
	// Order tag  (Deprecated)
	// Deprecated
	Tag *string `json:"tag,omitempty"`
}

// NewCreateTradeCancelAlgosV5RespData instantiates a new CreateTradeCancelAlgosV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradeCancelAlgosV5RespData() *CreateTradeCancelAlgosV5RespData {
	this := CreateTradeCancelAlgosV5RespData{}
	var algoClOrdId string = ""
	this.AlgoClOrdId = &algoClOrdId
	var algoId string = ""
	this.AlgoId = &algoId
	var clOrdId string = ""
	this.ClOrdId = &clOrdId
	var sCode string = ""
	this.SCode = &sCode
	var sMsg string = ""
	this.SMsg = &sMsg
	var tag string = ""
	this.Tag = &tag
	return &this
}

// NewCreateTradeCancelAlgosV5RespDataWithDefaults instantiates a new CreateTradeCancelAlgosV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradeCancelAlgosV5RespDataWithDefaults() *CreateTradeCancelAlgosV5RespData {
	this := CreateTradeCancelAlgosV5RespData{}
	var algoClOrdId string = ""
	this.AlgoClOrdId = &algoClOrdId
	var algoId string = ""
	this.AlgoId = &algoId
	var clOrdId string = ""
	this.ClOrdId = &clOrdId
	var sCode string = ""
	this.SCode = &sCode
	var sMsg string = ""
	this.SMsg = &sMsg
	var tag string = ""
	this.Tag = &tag
	return &this
}

// GetAlgoClOrdId returns the AlgoClOrdId field value if set, zero value otherwise.
// Deprecated
func (o *CreateTradeCancelAlgosV5RespData) GetAlgoClOrdId() string {
	if o == nil || IsNil(o.AlgoClOrdId) {
		var ret string
		return ret
	}
	return *o.AlgoClOrdId
}

// GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateTradeCancelAlgosV5RespData) GetAlgoClOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoClOrdId) {
		return nil, false
	}
	return o.AlgoClOrdId, true
}

// HasAlgoClOrdId returns a boolean if a field has been set.
func (o *CreateTradeCancelAlgosV5RespData) HasAlgoClOrdId() bool {
	if o != nil && !IsNil(o.AlgoClOrdId) {
		return true
	}

	return false
}

// SetAlgoClOrdId gets a reference to the given string and assigns it to the AlgoClOrdId field.
// Deprecated
func (o *CreateTradeCancelAlgosV5RespData) SetAlgoClOrdId(v string) {
	o.AlgoClOrdId = &v
}

// GetAlgoId returns the AlgoId field value if set, zero value otherwise.
func (o *CreateTradeCancelAlgosV5RespData) GetAlgoId() string {
	if o == nil || IsNil(o.AlgoId) {
		var ret string
		return ret
	}
	return *o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelAlgosV5RespData) GetAlgoIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoId) {
		return nil, false
	}
	return o.AlgoId, true
}

// HasAlgoId returns a boolean if a field has been set.
func (o *CreateTradeCancelAlgosV5RespData) HasAlgoId() bool {
	if o != nil && !IsNil(o.AlgoId) {
		return true
	}

	return false
}

// SetAlgoId gets a reference to the given string and assigns it to the AlgoId field.
func (o *CreateTradeCancelAlgosV5RespData) SetAlgoId(v string) {
	o.AlgoId = &v
}

// GetClOrdId returns the ClOrdId field value if set, zero value otherwise.
// Deprecated
func (o *CreateTradeCancelAlgosV5RespData) GetClOrdId() string {
	if o == nil || IsNil(o.ClOrdId) {
		var ret string
		return ret
	}
	return *o.ClOrdId
}

// GetClOrdIdOk returns a tuple with the ClOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateTradeCancelAlgosV5RespData) GetClOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClOrdId) {
		return nil, false
	}
	return o.ClOrdId, true
}

// HasClOrdId returns a boolean if a field has been set.
func (o *CreateTradeCancelAlgosV5RespData) HasClOrdId() bool {
	if o != nil && !IsNil(o.ClOrdId) {
		return true
	}

	return false
}

// SetClOrdId gets a reference to the given string and assigns it to the ClOrdId field.
// Deprecated
func (o *CreateTradeCancelAlgosV5RespData) SetClOrdId(v string) {
	o.ClOrdId = &v
}

// GetSCode returns the SCode field value if set, zero value otherwise.
func (o *CreateTradeCancelAlgosV5RespData) GetSCode() string {
	if o == nil || IsNil(o.SCode) {
		var ret string
		return ret
	}
	return *o.SCode
}

// GetSCodeOk returns a tuple with the SCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelAlgosV5RespData) GetSCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SCode) {
		return nil, false
	}
	return o.SCode, true
}

// HasSCode returns a boolean if a field has been set.
func (o *CreateTradeCancelAlgosV5RespData) HasSCode() bool {
	if o != nil && !IsNil(o.SCode) {
		return true
	}

	return false
}

// SetSCode gets a reference to the given string and assigns it to the SCode field.
func (o *CreateTradeCancelAlgosV5RespData) SetSCode(v string) {
	o.SCode = &v
}

// GetSMsg returns the SMsg field value if set, zero value otherwise.
func (o *CreateTradeCancelAlgosV5RespData) GetSMsg() string {
	if o == nil || IsNil(o.SMsg) {
		var ret string
		return ret
	}
	return *o.SMsg
}

// GetSMsgOk returns a tuple with the SMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelAlgosV5RespData) GetSMsgOk() (*string, bool) {
	if o == nil || IsNil(o.SMsg) {
		return nil, false
	}
	return o.SMsg, true
}

// HasSMsg returns a boolean if a field has been set.
func (o *CreateTradeCancelAlgosV5RespData) HasSMsg() bool {
	if o != nil && !IsNil(o.SMsg) {
		return true
	}

	return false
}

// SetSMsg gets a reference to the given string and assigns it to the SMsg field.
func (o *CreateTradeCancelAlgosV5RespData) SetSMsg(v string) {
	o.SMsg = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
// Deprecated
func (o *CreateTradeCancelAlgosV5RespData) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateTradeCancelAlgosV5RespData) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *CreateTradeCancelAlgosV5RespData) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
// Deprecated
func (o *CreateTradeCancelAlgosV5RespData) SetTag(v string) {
	o.Tag = &v
}

func (o CreateTradeCancelAlgosV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradeCancelAlgosV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlgoClOrdId) {
		toSerialize["algoClOrdId"] = o.AlgoClOrdId
	}
	if !IsNil(o.AlgoId) {
		toSerialize["algoId"] = o.AlgoId
	}
	if !IsNil(o.ClOrdId) {
		toSerialize["clOrdId"] = o.ClOrdId
	}
	if !IsNil(o.SCode) {
		toSerialize["sCode"] = o.SCode
	}
	if !IsNil(o.SMsg) {
		toSerialize["sMsg"] = o.SMsg
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return toSerialize, nil
}

type NullableCreateTradeCancelAlgosV5RespData struct {
	value *CreateTradeCancelAlgosV5RespData
	isSet bool
}

func (v NullableCreateTradeCancelAlgosV5RespData) Get() *CreateTradeCancelAlgosV5RespData {
	return v.value
}

func (v *NullableCreateTradeCancelAlgosV5RespData) Set(val *CreateTradeCancelAlgosV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradeCancelAlgosV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradeCancelAlgosV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradeCancelAlgosV5RespData(val *CreateTradeCancelAlgosV5RespData) *NullableCreateTradeCancelAlgosV5RespData {
	return &NullableCreateTradeCancelAlgosV5RespData{value: val, isSet: true}
}

func (v NullableCreateTradeCancelAlgosV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradeCancelAlgosV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


