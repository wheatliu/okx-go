/*
OKX v5 API

OpenAPI specification for Okx exchange - Rest API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rest

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateTradeCancelBatchOrdersV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradeCancelBatchOrdersV5Req{}

// CreateTradeCancelBatchOrdersV5Req struct for CreateTradeCancelBatchOrdersV5Req
type CreateTradeCancelBatchOrdersV5Req struct {
	// Client Order ID as assigned by the client
	ClOrdId *string `json:"clOrdId,omitempty"`
	// Instrument ID, e.g. `BTC-USDT`
	InstId string `json:"instId"`
	// Order ID  Either `ordId` or `clOrdId` is required. If both are passed, `ordId` will be used.
	OrdId *string `json:"ordId,omitempty"`
}

type _CreateTradeCancelBatchOrdersV5Req CreateTradeCancelBatchOrdersV5Req

// NewCreateTradeCancelBatchOrdersV5Req instantiates a new CreateTradeCancelBatchOrdersV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradeCancelBatchOrdersV5Req(instId string) *CreateTradeCancelBatchOrdersV5Req {
	this := CreateTradeCancelBatchOrdersV5Req{}
	var clOrdId string = ""
	this.ClOrdId = &clOrdId
	this.InstId = instId
	var ordId string = ""
	this.OrdId = &ordId
	return &this
}

// NewCreateTradeCancelBatchOrdersV5ReqWithDefaults instantiates a new CreateTradeCancelBatchOrdersV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradeCancelBatchOrdersV5ReqWithDefaults() *CreateTradeCancelBatchOrdersV5Req {
	this := CreateTradeCancelBatchOrdersV5Req{}
	var clOrdId string = ""
	this.ClOrdId = &clOrdId
	var instId string = ""
	this.InstId = instId
	var ordId string = ""
	this.OrdId = &ordId
	return &this
}

// GetClOrdId returns the ClOrdId field value if set, zero value otherwise.
func (o *CreateTradeCancelBatchOrdersV5Req) GetClOrdId() string {
	if o == nil || IsNil(o.ClOrdId) {
		var ret string
		return ret
	}
	return *o.ClOrdId
}

// GetClOrdIdOk returns a tuple with the ClOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelBatchOrdersV5Req) GetClOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClOrdId) {
		return nil, false
	}
	return o.ClOrdId, true
}

// HasClOrdId returns a boolean if a field has been set.
func (o *CreateTradeCancelBatchOrdersV5Req) HasClOrdId() bool {
	if o != nil && !IsNil(o.ClOrdId) {
		return true
	}

	return false
}

// SetClOrdId gets a reference to the given string and assigns it to the ClOrdId field.
func (o *CreateTradeCancelBatchOrdersV5Req) SetClOrdId(v string) {
	o.ClOrdId = &v
}

// GetInstId returns the InstId field value
func (o *CreateTradeCancelBatchOrdersV5Req) GetInstId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelBatchOrdersV5Req) GetInstIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstId, true
}

// SetInstId sets field value
func (o *CreateTradeCancelBatchOrdersV5Req) SetInstId(v string) {
	o.InstId = v
}

// GetOrdId returns the OrdId field value if set, zero value otherwise.
func (o *CreateTradeCancelBatchOrdersV5Req) GetOrdId() string {
	if o == nil || IsNil(o.OrdId) {
		var ret string
		return ret
	}
	return *o.OrdId
}

// GetOrdIdOk returns a tuple with the OrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelBatchOrdersV5Req) GetOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrdId) {
		return nil, false
	}
	return o.OrdId, true
}

// HasOrdId returns a boolean if a field has been set.
func (o *CreateTradeCancelBatchOrdersV5Req) HasOrdId() bool {
	if o != nil && !IsNil(o.OrdId) {
		return true
	}

	return false
}

// SetOrdId gets a reference to the given string and assigns it to the OrdId field.
func (o *CreateTradeCancelBatchOrdersV5Req) SetOrdId(v string) {
	o.OrdId = &v
}

func (o CreateTradeCancelBatchOrdersV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradeCancelBatchOrdersV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClOrdId) {
		toSerialize["clOrdId"] = o.ClOrdId
	}
	toSerialize["instId"] = o.InstId
	if !IsNil(o.OrdId) {
		toSerialize["ordId"] = o.OrdId
	}
	return toSerialize, nil
}

func (o *CreateTradeCancelBatchOrdersV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"instId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateTradeCancelBatchOrdersV5Req := _CreateTradeCancelBatchOrdersV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTradeCancelBatchOrdersV5Req)

	if err != nil {
		return err
	}

	*o = CreateTradeCancelBatchOrdersV5Req(varCreateTradeCancelBatchOrdersV5Req)

	return err
}

type NullableCreateTradeCancelBatchOrdersV5Req struct {
	value *CreateTradeCancelBatchOrdersV5Req
	isSet bool
}

func (v NullableCreateTradeCancelBatchOrdersV5Req) Get() *CreateTradeCancelBatchOrdersV5Req {
	return v.value
}

func (v *NullableCreateTradeCancelBatchOrdersV5Req) Set(val *CreateTradeCancelBatchOrdersV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradeCancelBatchOrdersV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradeCancelBatchOrdersV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradeCancelBatchOrdersV5Req(val *CreateTradeCancelBatchOrdersV5Req) *NullableCreateTradeCancelBatchOrdersV5Req {
	return &NullableCreateTradeCancelBatchOrdersV5Req{value: val, isSet: true}
}

func (v NullableCreateTradeCancelBatchOrdersV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradeCancelBatchOrdersV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


