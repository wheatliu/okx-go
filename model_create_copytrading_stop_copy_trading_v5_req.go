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

// checks if the CreateCopytradingStopCopyTradingV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCopytradingStopCopyTradingV5Req{}

// CreateCopytradingStopCopyTradingV5Req struct for CreateCopytradingStopCopyTradingV5Req
type CreateCopytradingStopCopyTradingV5Req struct {
	// Instrument type  `SWAP`
	InstType *string `json:"instType,omitempty"`
	// Action type for open positions, it is required if you have related copy position  `market_close`: immediately close at market price  `copy_close`：close when trader closes  `manual_close`: close manually
	SubPosCloseType string `json:"subPosCloseType"`
	// Lead trader unique code  A combination of case-sensitive alphanumerics, all numbers and the length is 16 characters, e.g. 213E8C92DC61EFAC
	UniqueCode string `json:"uniqueCode"`
}

type _CreateCopytradingStopCopyTradingV5Req CreateCopytradingStopCopyTradingV5Req

// NewCreateCopytradingStopCopyTradingV5Req instantiates a new CreateCopytradingStopCopyTradingV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCopytradingStopCopyTradingV5Req(subPosCloseType string, uniqueCode string) *CreateCopytradingStopCopyTradingV5Req {
	this := CreateCopytradingStopCopyTradingV5Req{}
	var instType string = ""
	this.InstType = &instType
	this.SubPosCloseType = subPosCloseType
	this.UniqueCode = uniqueCode
	return &this
}

// NewCreateCopytradingStopCopyTradingV5ReqWithDefaults instantiates a new CreateCopytradingStopCopyTradingV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCopytradingStopCopyTradingV5ReqWithDefaults() *CreateCopytradingStopCopyTradingV5Req {
	this := CreateCopytradingStopCopyTradingV5Req{}
	var instType string = ""
	this.InstType = &instType
	var subPosCloseType string = ""
	this.SubPosCloseType = subPosCloseType
	var uniqueCode string = ""
	this.UniqueCode = uniqueCode
	return &this
}

// GetInstType returns the InstType field value if set, zero value otherwise.
func (o *CreateCopytradingStopCopyTradingV5Req) GetInstType() string {
	if o == nil || IsNil(o.InstType) {
		var ret string
		return ret
	}
	return *o.InstType
}

// GetInstTypeOk returns a tuple with the InstType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingStopCopyTradingV5Req) GetInstTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstType) {
		return nil, false
	}
	return o.InstType, true
}

// HasInstType returns a boolean if a field has been set.
func (o *CreateCopytradingStopCopyTradingV5Req) HasInstType() bool {
	if o != nil && !IsNil(o.InstType) {
		return true
	}

	return false
}

// SetInstType gets a reference to the given string and assigns it to the InstType field.
func (o *CreateCopytradingStopCopyTradingV5Req) SetInstType(v string) {
	o.InstType = &v
}

// GetSubPosCloseType returns the SubPosCloseType field value
func (o *CreateCopytradingStopCopyTradingV5Req) GetSubPosCloseType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubPosCloseType
}

// GetSubPosCloseTypeOk returns a tuple with the SubPosCloseType field value
// and a boolean to check if the value has been set.
func (o *CreateCopytradingStopCopyTradingV5Req) GetSubPosCloseTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubPosCloseType, true
}

// SetSubPosCloseType sets field value
func (o *CreateCopytradingStopCopyTradingV5Req) SetSubPosCloseType(v string) {
	o.SubPosCloseType = v
}

// GetUniqueCode returns the UniqueCode field value
func (o *CreateCopytradingStopCopyTradingV5Req) GetUniqueCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UniqueCode
}

// GetUniqueCodeOk returns a tuple with the UniqueCode field value
// and a boolean to check if the value has been set.
func (o *CreateCopytradingStopCopyTradingV5Req) GetUniqueCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UniqueCode, true
}

// SetUniqueCode sets field value
func (o *CreateCopytradingStopCopyTradingV5Req) SetUniqueCode(v string) {
	o.UniqueCode = v
}

func (o CreateCopytradingStopCopyTradingV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCopytradingStopCopyTradingV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstType) {
		toSerialize["instType"] = o.InstType
	}
	toSerialize["subPosCloseType"] = o.SubPosCloseType
	toSerialize["uniqueCode"] = o.UniqueCode
	return toSerialize, nil
}

func (o *CreateCopytradingStopCopyTradingV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subPosCloseType",
		"uniqueCode",
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

	varCreateCopytradingStopCopyTradingV5Req := _CreateCopytradingStopCopyTradingV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateCopytradingStopCopyTradingV5Req)

	if err != nil {
		return err
	}

	*o = CreateCopytradingStopCopyTradingV5Req(varCreateCopytradingStopCopyTradingV5Req)

	return err
}

type NullableCreateCopytradingStopCopyTradingV5Req struct {
	value *CreateCopytradingStopCopyTradingV5Req
	isSet bool
}

func (v NullableCreateCopytradingStopCopyTradingV5Req) Get() *CreateCopytradingStopCopyTradingV5Req {
	return v.value
}

func (v *NullableCreateCopytradingStopCopyTradingV5Req) Set(val *CreateCopytradingStopCopyTradingV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCopytradingStopCopyTradingV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCopytradingStopCopyTradingV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCopytradingStopCopyTradingV5Req(val *CreateCopytradingStopCopyTradingV5Req) *NullableCreateCopytradingStopCopyTradingV5Req {
	return &NullableCreateCopytradingStopCopyTradingV5Req{value: val, isSet: true}
}

func (v NullableCreateCopytradingStopCopyTradingV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCopytradingStopCopyTradingV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


