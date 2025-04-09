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

// checks if the CreateTradeOneClickRepayV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradeOneClickRepayV5Req{}

// CreateTradeOneClickRepayV5Req struct for CreateTradeOneClickRepayV5Req
type CreateTradeOneClickRepayV5Req struct {
	// Debt currency type   Maximum 5 currencies can be selected in one order. If there are multiple currencies, separate them with commas.
	DebtCcy []string `json:"debtCcy"`
	// Repay currency type   Only one receiving currency type can be selected in one order and cannot be the same as the small payment currencies.
	RepayCcy string `json:"repayCcy"`
}

type _CreateTradeOneClickRepayV5Req CreateTradeOneClickRepayV5Req

// NewCreateTradeOneClickRepayV5Req instantiates a new CreateTradeOneClickRepayV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradeOneClickRepayV5Req(debtCcy []string, repayCcy string) *CreateTradeOneClickRepayV5Req {
	this := CreateTradeOneClickRepayV5Req{}
	this.DebtCcy = debtCcy
	this.RepayCcy = repayCcy
	return &this
}

// NewCreateTradeOneClickRepayV5ReqWithDefaults instantiates a new CreateTradeOneClickRepayV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradeOneClickRepayV5ReqWithDefaults() *CreateTradeOneClickRepayV5Req {
	this := CreateTradeOneClickRepayV5Req{}
	var repayCcy string = ""
	this.RepayCcy = repayCcy
	return &this
}

// GetDebtCcy returns the DebtCcy field value
func (o *CreateTradeOneClickRepayV5Req) GetDebtCcy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DebtCcy
}

// GetDebtCcyOk returns a tuple with the DebtCcy field value
// and a boolean to check if the value has been set.
func (o *CreateTradeOneClickRepayV5Req) GetDebtCcyOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DebtCcy, true
}

// SetDebtCcy sets field value
func (o *CreateTradeOneClickRepayV5Req) SetDebtCcy(v []string) {
	o.DebtCcy = v
}

// GetRepayCcy returns the RepayCcy field value
func (o *CreateTradeOneClickRepayV5Req) GetRepayCcy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepayCcy
}

// GetRepayCcyOk returns a tuple with the RepayCcy field value
// and a boolean to check if the value has been set.
func (o *CreateTradeOneClickRepayV5Req) GetRepayCcyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepayCcy, true
}

// SetRepayCcy sets field value
func (o *CreateTradeOneClickRepayV5Req) SetRepayCcy(v string) {
	o.RepayCcy = v
}

func (o CreateTradeOneClickRepayV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradeOneClickRepayV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["debtCcy"] = o.DebtCcy
	toSerialize["repayCcy"] = o.RepayCcy
	return toSerialize, nil
}

func (o *CreateTradeOneClickRepayV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"debtCcy",
		"repayCcy",
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

	varCreateTradeOneClickRepayV5Req := _CreateTradeOneClickRepayV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTradeOneClickRepayV5Req)

	if err != nil {
		return err
	}

	*o = CreateTradeOneClickRepayV5Req(varCreateTradeOneClickRepayV5Req)

	return err
}

type NullableCreateTradeOneClickRepayV5Req struct {
	value *CreateTradeOneClickRepayV5Req
	isSet bool
}

func (v NullableCreateTradeOneClickRepayV5Req) Get() *CreateTradeOneClickRepayV5Req {
	return v.value
}

func (v *NullableCreateTradeOneClickRepayV5Req) Set(val *CreateTradeOneClickRepayV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradeOneClickRepayV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradeOneClickRepayV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradeOneClickRepayV5Req(val *CreateTradeOneClickRepayV5Req) *NullableCreateTradeOneClickRepayV5Req {
	return &NullableCreateTradeOneClickRepayV5Req{value: val, isSet: true}
}

func (v NullableCreateTradeOneClickRepayV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradeOneClickRepayV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


