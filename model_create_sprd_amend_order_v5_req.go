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

// checks if the CreateSprdAmendOrderV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSprdAmendOrderV5Req{}

// CreateSprdAmendOrderV5Req struct for CreateSprdAmendOrderV5Req
type CreateSprdAmendOrderV5Req struct {
	// Client Order ID as assigned by the client
	ClOrdId *string `json:"clOrdId,omitempty"`
	// New price after amendment   Either `newSz` or `newPx` is required.
	NewPx *string `json:"newPx,omitempty"`
	// New quantity after amendment   Either `newSz` or `newPx` is required.   When amending a partially-filled order, the newSz should include the amount that has been filled.
	NewSz *string `json:"newSz,omitempty"`
	// Order ID   Either `ordId` or `clOrdId` is required. If both are passed, ordId will be used.
	OrdId *string `json:"ordId,omitempty"`
	// Client Request ID as assigned by the client for order amendment   A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters.   The response will include the corresponding reqId to help you identify the request if you provide it in the request.
	ReqId *string `json:"reqId,omitempty"`
}

// NewCreateSprdAmendOrderV5Req instantiates a new CreateSprdAmendOrderV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSprdAmendOrderV5Req() *CreateSprdAmendOrderV5Req {
	this := CreateSprdAmendOrderV5Req{}
	var clOrdId string = ""
	this.ClOrdId = &clOrdId
	var newPx string = ""
	this.NewPx = &newPx
	var newSz string = ""
	this.NewSz = &newSz
	var ordId string = ""
	this.OrdId = &ordId
	var reqId string = ""
	this.ReqId = &reqId
	return &this
}

// NewCreateSprdAmendOrderV5ReqWithDefaults instantiates a new CreateSprdAmendOrderV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSprdAmendOrderV5ReqWithDefaults() *CreateSprdAmendOrderV5Req {
	this := CreateSprdAmendOrderV5Req{}
	var clOrdId string = ""
	this.ClOrdId = &clOrdId
	var newPx string = ""
	this.NewPx = &newPx
	var newSz string = ""
	this.NewSz = &newSz
	var ordId string = ""
	this.OrdId = &ordId
	var reqId string = ""
	this.ReqId = &reqId
	return &this
}

// GetClOrdId returns the ClOrdId field value if set, zero value otherwise.
func (o *CreateSprdAmendOrderV5Req) GetClOrdId() string {
	if o == nil || IsNil(o.ClOrdId) {
		var ret string
		return ret
	}
	return *o.ClOrdId
}

// GetClOrdIdOk returns a tuple with the ClOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSprdAmendOrderV5Req) GetClOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClOrdId) {
		return nil, false
	}
	return o.ClOrdId, true
}

// HasClOrdId returns a boolean if a field has been set.
func (o *CreateSprdAmendOrderV5Req) HasClOrdId() bool {
	if o != nil && !IsNil(o.ClOrdId) {
		return true
	}

	return false
}

// SetClOrdId gets a reference to the given string and assigns it to the ClOrdId field.
func (o *CreateSprdAmendOrderV5Req) SetClOrdId(v string) {
	o.ClOrdId = &v
}

// GetNewPx returns the NewPx field value if set, zero value otherwise.
func (o *CreateSprdAmendOrderV5Req) GetNewPx() string {
	if o == nil || IsNil(o.NewPx) {
		var ret string
		return ret
	}
	return *o.NewPx
}

// GetNewPxOk returns a tuple with the NewPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSprdAmendOrderV5Req) GetNewPxOk() (*string, bool) {
	if o == nil || IsNil(o.NewPx) {
		return nil, false
	}
	return o.NewPx, true
}

// HasNewPx returns a boolean if a field has been set.
func (o *CreateSprdAmendOrderV5Req) HasNewPx() bool {
	if o != nil && !IsNil(o.NewPx) {
		return true
	}

	return false
}

// SetNewPx gets a reference to the given string and assigns it to the NewPx field.
func (o *CreateSprdAmendOrderV5Req) SetNewPx(v string) {
	o.NewPx = &v
}

// GetNewSz returns the NewSz field value if set, zero value otherwise.
func (o *CreateSprdAmendOrderV5Req) GetNewSz() string {
	if o == nil || IsNil(o.NewSz) {
		var ret string
		return ret
	}
	return *o.NewSz
}

// GetNewSzOk returns a tuple with the NewSz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSprdAmendOrderV5Req) GetNewSzOk() (*string, bool) {
	if o == nil || IsNil(o.NewSz) {
		return nil, false
	}
	return o.NewSz, true
}

// HasNewSz returns a boolean if a field has been set.
func (o *CreateSprdAmendOrderV5Req) HasNewSz() bool {
	if o != nil && !IsNil(o.NewSz) {
		return true
	}

	return false
}

// SetNewSz gets a reference to the given string and assigns it to the NewSz field.
func (o *CreateSprdAmendOrderV5Req) SetNewSz(v string) {
	o.NewSz = &v
}

// GetOrdId returns the OrdId field value if set, zero value otherwise.
func (o *CreateSprdAmendOrderV5Req) GetOrdId() string {
	if o == nil || IsNil(o.OrdId) {
		var ret string
		return ret
	}
	return *o.OrdId
}

// GetOrdIdOk returns a tuple with the OrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSprdAmendOrderV5Req) GetOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrdId) {
		return nil, false
	}
	return o.OrdId, true
}

// HasOrdId returns a boolean if a field has been set.
func (o *CreateSprdAmendOrderV5Req) HasOrdId() bool {
	if o != nil && !IsNil(o.OrdId) {
		return true
	}

	return false
}

// SetOrdId gets a reference to the given string and assigns it to the OrdId field.
func (o *CreateSprdAmendOrderV5Req) SetOrdId(v string) {
	o.OrdId = &v
}

// GetReqId returns the ReqId field value if set, zero value otherwise.
func (o *CreateSprdAmendOrderV5Req) GetReqId() string {
	if o == nil || IsNil(o.ReqId) {
		var ret string
		return ret
	}
	return *o.ReqId
}

// GetReqIdOk returns a tuple with the ReqId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSprdAmendOrderV5Req) GetReqIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReqId) {
		return nil, false
	}
	return o.ReqId, true
}

// HasReqId returns a boolean if a field has been set.
func (o *CreateSprdAmendOrderV5Req) HasReqId() bool {
	if o != nil && !IsNil(o.ReqId) {
		return true
	}

	return false
}

// SetReqId gets a reference to the given string and assigns it to the ReqId field.
func (o *CreateSprdAmendOrderV5Req) SetReqId(v string) {
	o.ReqId = &v
}

func (o CreateSprdAmendOrderV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSprdAmendOrderV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClOrdId) {
		toSerialize["clOrdId"] = o.ClOrdId
	}
	if !IsNil(o.NewPx) {
		toSerialize["newPx"] = o.NewPx
	}
	if !IsNil(o.NewSz) {
		toSerialize["newSz"] = o.NewSz
	}
	if !IsNil(o.OrdId) {
		toSerialize["ordId"] = o.OrdId
	}
	if !IsNil(o.ReqId) {
		toSerialize["reqId"] = o.ReqId
	}
	return toSerialize, nil
}

type NullableCreateSprdAmendOrderV5Req struct {
	value *CreateSprdAmendOrderV5Req
	isSet bool
}

func (v NullableCreateSprdAmendOrderV5Req) Get() *CreateSprdAmendOrderV5Req {
	return v.value
}

func (v *NullableCreateSprdAmendOrderV5Req) Set(val *CreateSprdAmendOrderV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSprdAmendOrderV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSprdAmendOrderV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSprdAmendOrderV5Req(val *CreateSprdAmendOrderV5Req) *NullableCreateSprdAmendOrderV5Req {
	return &NullableCreateSprdAmendOrderV5Req{value: val, isSet: true}
}

func (v NullableCreateSprdAmendOrderV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSprdAmendOrderV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


