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

// checks if the CreateFinanceStakingDefiEthPurchaseV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFinanceStakingDefiEthPurchaseV5Req{}

// CreateFinanceStakingDefiEthPurchaseV5Req struct for CreateFinanceStakingDefiEthPurchaseV5Req
type CreateFinanceStakingDefiEthPurchaseV5Req struct {
	// Investment amount
	Amt string `json:"amt"`
}

type _CreateFinanceStakingDefiEthPurchaseV5Req CreateFinanceStakingDefiEthPurchaseV5Req

// NewCreateFinanceStakingDefiEthPurchaseV5Req instantiates a new CreateFinanceStakingDefiEthPurchaseV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFinanceStakingDefiEthPurchaseV5Req(amt string) *CreateFinanceStakingDefiEthPurchaseV5Req {
	this := CreateFinanceStakingDefiEthPurchaseV5Req{}
	this.Amt = amt
	return &this
}

// NewCreateFinanceStakingDefiEthPurchaseV5ReqWithDefaults instantiates a new CreateFinanceStakingDefiEthPurchaseV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFinanceStakingDefiEthPurchaseV5ReqWithDefaults() *CreateFinanceStakingDefiEthPurchaseV5Req {
	this := CreateFinanceStakingDefiEthPurchaseV5Req{}
	var amt string = ""
	this.Amt = amt
	return &this
}

// GetAmt returns the Amt field value
func (o *CreateFinanceStakingDefiEthPurchaseV5Req) GetAmt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amt
}

// GetAmtOk returns a tuple with the Amt field value
// and a boolean to check if the value has been set.
func (o *CreateFinanceStakingDefiEthPurchaseV5Req) GetAmtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amt, true
}

// SetAmt sets field value
func (o *CreateFinanceStakingDefiEthPurchaseV5Req) SetAmt(v string) {
	o.Amt = v
}

func (o CreateFinanceStakingDefiEthPurchaseV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFinanceStakingDefiEthPurchaseV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amt"] = o.Amt
	return toSerialize, nil
}

func (o *CreateFinanceStakingDefiEthPurchaseV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amt",
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

	varCreateFinanceStakingDefiEthPurchaseV5Req := _CreateFinanceStakingDefiEthPurchaseV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateFinanceStakingDefiEthPurchaseV5Req)

	if err != nil {
		return err
	}

	*o = CreateFinanceStakingDefiEthPurchaseV5Req(varCreateFinanceStakingDefiEthPurchaseV5Req)

	return err
}

type NullableCreateFinanceStakingDefiEthPurchaseV5Req struct {
	value *CreateFinanceStakingDefiEthPurchaseV5Req
	isSet bool
}

func (v NullableCreateFinanceStakingDefiEthPurchaseV5Req) Get() *CreateFinanceStakingDefiEthPurchaseV5Req {
	return v.value
}

func (v *NullableCreateFinanceStakingDefiEthPurchaseV5Req) Set(val *CreateFinanceStakingDefiEthPurchaseV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFinanceStakingDefiEthPurchaseV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFinanceStakingDefiEthPurchaseV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFinanceStakingDefiEthPurchaseV5Req(val *CreateFinanceStakingDefiEthPurchaseV5Req) *NullableCreateFinanceStakingDefiEthPurchaseV5Req {
	return &NullableCreateFinanceStakingDefiEthPurchaseV5Req{value: val, isSet: true}
}

func (v NullableCreateFinanceStakingDefiEthPurchaseV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFinanceStakingDefiEthPurchaseV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


