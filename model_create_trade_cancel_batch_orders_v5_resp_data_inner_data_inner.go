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

// checks if the CreateTradeCancelBatchOrdersV5RespDataInnerDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradeCancelBatchOrdersV5RespDataInnerDataInner{}

// CreateTradeCancelBatchOrdersV5RespDataInnerDataInner struct for CreateTradeCancelBatchOrdersV5RespDataInnerDataInner
type CreateTradeCancelBatchOrdersV5RespDataInnerDataInner struct {
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

// NewCreateTradeCancelBatchOrdersV5RespDataInnerDataInner instantiates a new CreateTradeCancelBatchOrdersV5RespDataInnerDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradeCancelBatchOrdersV5RespDataInnerDataInner() *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner {
	this := CreateTradeCancelBatchOrdersV5RespDataInnerDataInner{}
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

// NewCreateTradeCancelBatchOrdersV5RespDataInnerDataInnerWithDefaults instantiates a new CreateTradeCancelBatchOrdersV5RespDataInnerDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradeCancelBatchOrdersV5RespDataInnerDataInnerWithDefaults() *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner {
	this := CreateTradeCancelBatchOrdersV5RespDataInnerDataInner{}
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
func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) GetClOrdId() string {
	if o == nil || IsNil(o.ClOrdId) {
		var ret string
		return ret
	}
	return *o.ClOrdId
}

// GetClOrdIdOk returns a tuple with the ClOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) GetClOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClOrdId) {
		return nil, false
	}
	return o.ClOrdId, true
}

// HasClOrdId returns a boolean if a field has been set.
func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) HasClOrdId() bool {
	if o != nil && !IsNil(o.ClOrdId) {
		return true
	}

	return false
}

// SetClOrdId gets a reference to the given string and assigns it to the ClOrdId field.
func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) SetClOrdId(v string) {
	o.ClOrdId = &v
}

// GetOrdId returns the OrdId field value if set, zero value otherwise.
func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) GetOrdId() string {
	if o == nil || IsNil(o.OrdId) {
		var ret string
		return ret
	}
	return *o.OrdId
}

// GetOrdIdOk returns a tuple with the OrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) GetOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrdId) {
		return nil, false
	}
	return o.OrdId, true
}

// HasOrdId returns a boolean if a field has been set.
func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) HasOrdId() bool {
	if o != nil && !IsNil(o.OrdId) {
		return true
	}

	return false
}

// SetOrdId gets a reference to the given string and assigns it to the OrdId field.
func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) SetOrdId(v string) {
	o.OrdId = &v
}

// GetSCode returns the SCode field value if set, zero value otherwise.
func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) GetSCode() string {
	if o == nil || IsNil(o.SCode) {
		var ret string
		return ret
	}
	return *o.SCode
}

// GetSCodeOk returns a tuple with the SCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) GetSCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SCode) {
		return nil, false
	}
	return o.SCode, true
}

// HasSCode returns a boolean if a field has been set.
func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) HasSCode() bool {
	if o != nil && !IsNil(o.SCode) {
		return true
	}

	return false
}

// SetSCode gets a reference to the given string and assigns it to the SCode field.
func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) SetSCode(v string) {
	o.SCode = &v
}

// GetSMsg returns the SMsg field value if set, zero value otherwise.
func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) GetSMsg() string {
	if o == nil || IsNil(o.SMsg) {
		var ret string
		return ret
	}
	return *o.SMsg
}

// GetSMsgOk returns a tuple with the SMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) GetSMsgOk() (*string, bool) {
	if o == nil || IsNil(o.SMsg) {
		return nil, false
	}
	return o.SMsg, true
}

// HasSMsg returns a boolean if a field has been set.
func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) HasSMsg() bool {
	if o != nil && !IsNil(o.SMsg) {
		return true
	}

	return false
}

// SetSMsg gets a reference to the given string and assigns it to the SMsg field.
func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) SetSMsg(v string) {
	o.SMsg = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) SetTs(v string) {
	o.Ts = &v
}

func (o CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) ToMap() (map[string]interface{}, error) {
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

type NullableCreateTradeCancelBatchOrdersV5RespDataInnerDataInner struct {
	value *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner
	isSet bool
}

func (v NullableCreateTradeCancelBatchOrdersV5RespDataInnerDataInner) Get() *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner {
	return v.value
}

func (v *NullableCreateTradeCancelBatchOrdersV5RespDataInnerDataInner) Set(val *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradeCancelBatchOrdersV5RespDataInnerDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradeCancelBatchOrdersV5RespDataInnerDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradeCancelBatchOrdersV5RespDataInnerDataInner(val *CreateTradeCancelBatchOrdersV5RespDataInnerDataInner) *NullableCreateTradeCancelBatchOrdersV5RespDataInnerDataInner {
	return &NullableCreateTradeCancelBatchOrdersV5RespDataInnerDataInner{value: val, isSet: true}
}

func (v NullableCreateTradeCancelBatchOrdersV5RespDataInnerDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradeCancelBatchOrdersV5RespDataInnerDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


