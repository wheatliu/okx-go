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

// checks if the CreateTradeAmendBatchOrdersV5RespDataInnerDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradeAmendBatchOrdersV5RespDataInnerDataInner{}

// CreateTradeAmendBatchOrdersV5RespDataInnerDataInner struct for CreateTradeAmendBatchOrdersV5RespDataInnerDataInner
type CreateTradeAmendBatchOrdersV5RespDataInnerDataInner struct {
	// Client Order ID as assigned by the client
	ClOrdId *string `json:"clOrdId,omitempty"`
	// Order ID
	OrdId *string `json:"ordId,omitempty"`
	// Client Request ID as assigned by the client for order amendment.
	ReqId *string `json:"reqId,omitempty"`
	// The code of the event execution result, `0` means success.
	SCode *string `json:"sCode,omitempty"`
	// Rejection message if the request is unsuccessful.
	SMsg *string `json:"sMsg,omitempty"`
	// Timestamp when the order request processing is finished by our system, Unix timestamp format in milliseconds, e.g. `1597026383085`
	Ts *string `json:"ts,omitempty"`
}

// NewCreateTradeAmendBatchOrdersV5RespDataInnerDataInner instantiates a new CreateTradeAmendBatchOrdersV5RespDataInnerDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradeAmendBatchOrdersV5RespDataInnerDataInner() *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner {
	this := CreateTradeAmendBatchOrdersV5RespDataInnerDataInner{}
	var clOrdId string = ""
	this.ClOrdId = &clOrdId
	var ordId string = ""
	this.OrdId = &ordId
	var reqId string = ""
	this.ReqId = &reqId
	var sCode string = ""
	this.SCode = &sCode
	var sMsg string = ""
	this.SMsg = &sMsg
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewCreateTradeAmendBatchOrdersV5RespDataInnerDataInnerWithDefaults instantiates a new CreateTradeAmendBatchOrdersV5RespDataInnerDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradeAmendBatchOrdersV5RespDataInnerDataInnerWithDefaults() *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner {
	this := CreateTradeAmendBatchOrdersV5RespDataInnerDataInner{}
	var clOrdId string = ""
	this.ClOrdId = &clOrdId
	var ordId string = ""
	this.OrdId = &ordId
	var reqId string = ""
	this.ReqId = &reqId
	var sCode string = ""
	this.SCode = &sCode
	var sMsg string = ""
	this.SMsg = &sMsg
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetClOrdId returns the ClOrdId field value if set, zero value otherwise.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) GetClOrdId() string {
	if o == nil || IsNil(o.ClOrdId) {
		var ret string
		return ret
	}
	return *o.ClOrdId
}

// GetClOrdIdOk returns a tuple with the ClOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) GetClOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClOrdId) {
		return nil, false
	}
	return o.ClOrdId, true
}

// HasClOrdId returns a boolean if a field has been set.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) HasClOrdId() bool {
	if o != nil && !IsNil(o.ClOrdId) {
		return true
	}

	return false
}

// SetClOrdId gets a reference to the given string and assigns it to the ClOrdId field.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) SetClOrdId(v string) {
	o.ClOrdId = &v
}

// GetOrdId returns the OrdId field value if set, zero value otherwise.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) GetOrdId() string {
	if o == nil || IsNil(o.OrdId) {
		var ret string
		return ret
	}
	return *o.OrdId
}

// GetOrdIdOk returns a tuple with the OrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) GetOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrdId) {
		return nil, false
	}
	return o.OrdId, true
}

// HasOrdId returns a boolean if a field has been set.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) HasOrdId() bool {
	if o != nil && !IsNil(o.OrdId) {
		return true
	}

	return false
}

// SetOrdId gets a reference to the given string and assigns it to the OrdId field.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) SetOrdId(v string) {
	o.OrdId = &v
}

// GetReqId returns the ReqId field value if set, zero value otherwise.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) GetReqId() string {
	if o == nil || IsNil(o.ReqId) {
		var ret string
		return ret
	}
	return *o.ReqId
}

// GetReqIdOk returns a tuple with the ReqId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) GetReqIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReqId) {
		return nil, false
	}
	return o.ReqId, true
}

// HasReqId returns a boolean if a field has been set.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) HasReqId() bool {
	if o != nil && !IsNil(o.ReqId) {
		return true
	}

	return false
}

// SetReqId gets a reference to the given string and assigns it to the ReqId field.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) SetReqId(v string) {
	o.ReqId = &v
}

// GetSCode returns the SCode field value if set, zero value otherwise.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) GetSCode() string {
	if o == nil || IsNil(o.SCode) {
		var ret string
		return ret
	}
	return *o.SCode
}

// GetSCodeOk returns a tuple with the SCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) GetSCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SCode) {
		return nil, false
	}
	return o.SCode, true
}

// HasSCode returns a boolean if a field has been set.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) HasSCode() bool {
	if o != nil && !IsNil(o.SCode) {
		return true
	}

	return false
}

// SetSCode gets a reference to the given string and assigns it to the SCode field.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) SetSCode(v string) {
	o.SCode = &v
}

// GetSMsg returns the SMsg field value if set, zero value otherwise.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) GetSMsg() string {
	if o == nil || IsNil(o.SMsg) {
		var ret string
		return ret
	}
	return *o.SMsg
}

// GetSMsgOk returns a tuple with the SMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) GetSMsgOk() (*string, bool) {
	if o == nil || IsNil(o.SMsg) {
		return nil, false
	}
	return o.SMsg, true
}

// HasSMsg returns a boolean if a field has been set.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) HasSMsg() bool {
	if o != nil && !IsNil(o.SMsg) {
		return true
	}

	return false
}

// SetSMsg gets a reference to the given string and assigns it to the SMsg field.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) SetSMsg(v string) {
	o.SMsg = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) SetTs(v string) {
	o.Ts = &v
}

func (o CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClOrdId) {
		toSerialize["clOrdId"] = o.ClOrdId
	}
	if !IsNil(o.OrdId) {
		toSerialize["ordId"] = o.OrdId
	}
	if !IsNil(o.ReqId) {
		toSerialize["reqId"] = o.ReqId
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

type NullableCreateTradeAmendBatchOrdersV5RespDataInnerDataInner struct {
	value *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner
	isSet bool
}

func (v NullableCreateTradeAmendBatchOrdersV5RespDataInnerDataInner) Get() *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner {
	return v.value
}

func (v *NullableCreateTradeAmendBatchOrdersV5RespDataInnerDataInner) Set(val *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradeAmendBatchOrdersV5RespDataInnerDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradeAmendBatchOrdersV5RespDataInnerDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradeAmendBatchOrdersV5RespDataInnerDataInner(val *CreateTradeAmendBatchOrdersV5RespDataInnerDataInner) *NullableCreateTradeAmendBatchOrdersV5RespDataInnerDataInner {
	return &NullableCreateTradeAmendBatchOrdersV5RespDataInnerDataInner{value: val, isSet: true}
}

func (v NullableCreateTradeAmendBatchOrdersV5RespDataInnerDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradeAmendBatchOrdersV5RespDataInnerDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


