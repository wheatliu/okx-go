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

// checks if the CreateTradeCancelBatchOrdersV5RespDataDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradeCancelBatchOrdersV5RespDataDataInner{}

// CreateTradeCancelBatchOrdersV5RespDataDataInner struct for CreateTradeCancelBatchOrdersV5RespDataDataInner
type CreateTradeCancelBatchOrdersV5RespDataDataInner struct {
	// Client Order ID as assigned by the client
	ClOrdId *string `json:"clOrdId,omitempty"`
	// Order ID
	OrdId *string `json:"ordId,omitempty"`
	// The code of the event execution result, `0` means success.
	SCode *string `json:"sCode,omitempty"`
	// Rejection message if the request is unsuccessful.
	SMsg *string `json:"sMsg,omitempty"`
	// Timestamp when the order request processing is finished by our system, Unix timestamp format in milliseconds, e.g. `1597026383085`
	Ts *string `json:"ts,omitempty"`
}

// NewCreateTradeCancelBatchOrdersV5RespDataDataInner instantiates a new CreateTradeCancelBatchOrdersV5RespDataDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradeCancelBatchOrdersV5RespDataDataInner() *CreateTradeCancelBatchOrdersV5RespDataDataInner {
	this := CreateTradeCancelBatchOrdersV5RespDataDataInner{}
	var clOrdId string = ""
	this.ClOrdId = &clOrdId
	var ordId string = ""
	this.OrdId = &ordId
	var sCode string = ""
	this.SCode = &sCode
	var sMsg string = ""
	this.SMsg = &sMsg
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewCreateTradeCancelBatchOrdersV5RespDataDataInnerWithDefaults instantiates a new CreateTradeCancelBatchOrdersV5RespDataDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradeCancelBatchOrdersV5RespDataDataInnerWithDefaults() *CreateTradeCancelBatchOrdersV5RespDataDataInner {
	this := CreateTradeCancelBatchOrdersV5RespDataDataInner{}
	var clOrdId string = ""
	this.ClOrdId = &clOrdId
	var ordId string = ""
	this.OrdId = &ordId
	var sCode string = ""
	this.SCode = &sCode
	var sMsg string = ""
	this.SMsg = &sMsg
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetClOrdId returns the ClOrdId field value if set, zero value otherwise.
func (o *CreateTradeCancelBatchOrdersV5RespDataDataInner) GetClOrdId() string {
	if o == nil || IsNil(o.ClOrdId) {
		var ret string
		return ret
	}
	return *o.ClOrdId
}

// GetClOrdIdOk returns a tuple with the ClOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelBatchOrdersV5RespDataDataInner) GetClOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClOrdId) {
		return nil, false
	}
	return o.ClOrdId, true
}

// HasClOrdId returns a boolean if a field has been set.
func (o *CreateTradeCancelBatchOrdersV5RespDataDataInner) HasClOrdId() bool {
	if o != nil && !IsNil(o.ClOrdId) {
		return true
	}

	return false
}

// SetClOrdId gets a reference to the given string and assigns it to the ClOrdId field.
func (o *CreateTradeCancelBatchOrdersV5RespDataDataInner) SetClOrdId(v string) {
	o.ClOrdId = &v
}

// GetOrdId returns the OrdId field value if set, zero value otherwise.
func (o *CreateTradeCancelBatchOrdersV5RespDataDataInner) GetOrdId() string {
	if o == nil || IsNil(o.OrdId) {
		var ret string
		return ret
	}
	return *o.OrdId
}

// GetOrdIdOk returns a tuple with the OrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelBatchOrdersV5RespDataDataInner) GetOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrdId) {
		return nil, false
	}
	return o.OrdId, true
}

// HasOrdId returns a boolean if a field has been set.
func (o *CreateTradeCancelBatchOrdersV5RespDataDataInner) HasOrdId() bool {
	if o != nil && !IsNil(o.OrdId) {
		return true
	}

	return false
}

// SetOrdId gets a reference to the given string and assigns it to the OrdId field.
func (o *CreateTradeCancelBatchOrdersV5RespDataDataInner) SetOrdId(v string) {
	o.OrdId = &v
}

// GetSCode returns the SCode field value if set, zero value otherwise.
func (o *CreateTradeCancelBatchOrdersV5RespDataDataInner) GetSCode() string {
	if o == nil || IsNil(o.SCode) {
		var ret string
		return ret
	}
	return *o.SCode
}

// GetSCodeOk returns a tuple with the SCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelBatchOrdersV5RespDataDataInner) GetSCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SCode) {
		return nil, false
	}
	return o.SCode, true
}

// HasSCode returns a boolean if a field has been set.
func (o *CreateTradeCancelBatchOrdersV5RespDataDataInner) HasSCode() bool {
	if o != nil && !IsNil(o.SCode) {
		return true
	}

	return false
}

// SetSCode gets a reference to the given string and assigns it to the SCode field.
func (o *CreateTradeCancelBatchOrdersV5RespDataDataInner) SetSCode(v string) {
	o.SCode = &v
}

// GetSMsg returns the SMsg field value if set, zero value otherwise.
func (o *CreateTradeCancelBatchOrdersV5RespDataDataInner) GetSMsg() string {
	if o == nil || IsNil(o.SMsg) {
		var ret string
		return ret
	}
	return *o.SMsg
}

// GetSMsgOk returns a tuple with the SMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelBatchOrdersV5RespDataDataInner) GetSMsgOk() (*string, bool) {
	if o == nil || IsNil(o.SMsg) {
		return nil, false
	}
	return o.SMsg, true
}

// HasSMsg returns a boolean if a field has been set.
func (o *CreateTradeCancelBatchOrdersV5RespDataDataInner) HasSMsg() bool {
	if o != nil && !IsNil(o.SMsg) {
		return true
	}

	return false
}

// SetSMsg gets a reference to the given string and assigns it to the SMsg field.
func (o *CreateTradeCancelBatchOrdersV5RespDataDataInner) SetSMsg(v string) {
	o.SMsg = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *CreateTradeCancelBatchOrdersV5RespDataDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelBatchOrdersV5RespDataDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *CreateTradeCancelBatchOrdersV5RespDataDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *CreateTradeCancelBatchOrdersV5RespDataDataInner) SetTs(v string) {
	o.Ts = &v
}

func (o CreateTradeCancelBatchOrdersV5RespDataDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradeCancelBatchOrdersV5RespDataDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClOrdId) {
		toSerialize["clOrdId"] = o.ClOrdId
	}
	if !IsNil(o.OrdId) {
		toSerialize["ordId"] = o.OrdId
	}
	if !IsNil(o.SCode) {
		toSerialize["sCode"] = o.SCode
	}
	if !IsNil(o.SMsg) {
		toSerialize["sMsg"] = o.SMsg
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableCreateTradeCancelBatchOrdersV5RespDataDataInner struct {
	value *CreateTradeCancelBatchOrdersV5RespDataDataInner
	isSet bool
}

func (v NullableCreateTradeCancelBatchOrdersV5RespDataDataInner) Get() *CreateTradeCancelBatchOrdersV5RespDataDataInner {
	return v.value
}

func (v *NullableCreateTradeCancelBatchOrdersV5RespDataDataInner) Set(val *CreateTradeCancelBatchOrdersV5RespDataDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradeCancelBatchOrdersV5RespDataDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradeCancelBatchOrdersV5RespDataDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradeCancelBatchOrdersV5RespDataDataInner(val *CreateTradeCancelBatchOrdersV5RespDataDataInner) *NullableCreateTradeCancelBatchOrdersV5RespDataDataInner {
	return &NullableCreateTradeCancelBatchOrdersV5RespDataDataInner{value: val, isSet: true}
}

func (v NullableCreateTradeCancelBatchOrdersV5RespDataDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradeCancelBatchOrdersV5RespDataDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


