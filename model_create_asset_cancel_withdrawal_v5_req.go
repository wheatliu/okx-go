/*
Okx Rest API

OpenAPI specification for Okx cryptocurrency exchange - Rest API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rest

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateAssetCancelWithdrawalV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAssetCancelWithdrawalV5Req{}

// CreateAssetCancelWithdrawalV5Req struct for CreateAssetCancelWithdrawalV5Req
type CreateAssetCancelWithdrawalV5Req struct {
	// Withdrawal ID
	WdId string `json:"wdId"`
}

type _CreateAssetCancelWithdrawalV5Req CreateAssetCancelWithdrawalV5Req

// NewCreateAssetCancelWithdrawalV5Req instantiates a new CreateAssetCancelWithdrawalV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAssetCancelWithdrawalV5Req(wdId string) *CreateAssetCancelWithdrawalV5Req {
	this := CreateAssetCancelWithdrawalV5Req{}
	this.WdId = wdId
	return &this
}

// NewCreateAssetCancelWithdrawalV5ReqWithDefaults instantiates a new CreateAssetCancelWithdrawalV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAssetCancelWithdrawalV5ReqWithDefaults() *CreateAssetCancelWithdrawalV5Req {
	this := CreateAssetCancelWithdrawalV5Req{}
	var wdId string = ""
	this.WdId = wdId
	return &this
}

// GetWdId returns the WdId field value
func (o *CreateAssetCancelWithdrawalV5Req) GetWdId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WdId
}

// GetWdIdOk returns a tuple with the WdId field value
// and a boolean to check if the value has been set.
func (o *CreateAssetCancelWithdrawalV5Req) GetWdIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WdId, true
}

// SetWdId sets field value
func (o *CreateAssetCancelWithdrawalV5Req) SetWdId(v string) {
	o.WdId = v
}

func (o CreateAssetCancelWithdrawalV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAssetCancelWithdrawalV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["wdId"] = o.WdId
	return toSerialize, nil
}

func (o *CreateAssetCancelWithdrawalV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"wdId",
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

	varCreateAssetCancelWithdrawalV5Req := _CreateAssetCancelWithdrawalV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateAssetCancelWithdrawalV5Req)

	if err != nil {
		return err
	}

	*o = CreateAssetCancelWithdrawalV5Req(varCreateAssetCancelWithdrawalV5Req)

	return err
}

type NullableCreateAssetCancelWithdrawalV5Req struct {
	value *CreateAssetCancelWithdrawalV5Req
	isSet bool
}

func (v NullableCreateAssetCancelWithdrawalV5Req) Get() *CreateAssetCancelWithdrawalV5Req {
	return v.value
}

func (v *NullableCreateAssetCancelWithdrawalV5Req) Set(val *CreateAssetCancelWithdrawalV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAssetCancelWithdrawalV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAssetCancelWithdrawalV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAssetCancelWithdrawalV5Req(val *CreateAssetCancelWithdrawalV5Req) *NullableCreateAssetCancelWithdrawalV5Req {
	return &NullableCreateAssetCancelWithdrawalV5Req{value: val, isSet: true}
}

func (v NullableCreateAssetCancelWithdrawalV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAssetCancelWithdrawalV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


