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

// checks if the GetRubikStatTradingDataSupportCoinV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRubikStatTradingDataSupportCoinV5RespDataInner{}

// GetRubikStatTradingDataSupportCoinV5RespDataInner struct for GetRubikStatTradingDataSupportCoinV5RespDataInner
type GetRubikStatTradingDataSupportCoinV5RespDataInner struct {
	// Currency supported by derivatives trading data
	Contract []string `json:"contract,omitempty"`
	// Currency supported by option trading data
	Option []string `json:"option,omitempty"`
	// Currency supported by spot trading data
	Spot []string `json:"spot,omitempty"`
}

// NewGetRubikStatTradingDataSupportCoinV5RespDataInner instantiates a new GetRubikStatTradingDataSupportCoinV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRubikStatTradingDataSupportCoinV5RespDataInner() *GetRubikStatTradingDataSupportCoinV5RespDataInner {
	this := GetRubikStatTradingDataSupportCoinV5RespDataInner{}
	return &this
}

// NewGetRubikStatTradingDataSupportCoinV5RespDataInnerWithDefaults instantiates a new GetRubikStatTradingDataSupportCoinV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRubikStatTradingDataSupportCoinV5RespDataInnerWithDefaults() *GetRubikStatTradingDataSupportCoinV5RespDataInner {
	this := GetRubikStatTradingDataSupportCoinV5RespDataInner{}
	return &this
}

// GetContract returns the Contract field value if set, zero value otherwise.
func (o *GetRubikStatTradingDataSupportCoinV5RespDataInner) GetContract() []string {
	if o == nil || IsNil(o.Contract) {
		var ret []string
		return ret
	}
	return o.Contract
}

// GetContractOk returns a tuple with the Contract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRubikStatTradingDataSupportCoinV5RespDataInner) GetContractOk() ([]string, bool) {
	if o == nil || IsNil(o.Contract) {
		return nil, false
	}
	return o.Contract, true
}

// HasContract returns a boolean if a field has been set.
func (o *GetRubikStatTradingDataSupportCoinV5RespDataInner) HasContract() bool {
	if o != nil && !IsNil(o.Contract) {
		return true
	}

	return false
}

// SetContract gets a reference to the given []string and assigns it to the Contract field.
func (o *GetRubikStatTradingDataSupportCoinV5RespDataInner) SetContract(v []string) {
	o.Contract = v
}

// GetOption returns the Option field value if set, zero value otherwise.
func (o *GetRubikStatTradingDataSupportCoinV5RespDataInner) GetOption() []string {
	if o == nil || IsNil(o.Option) {
		var ret []string
		return ret
	}
	return o.Option
}

// GetOptionOk returns a tuple with the Option field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRubikStatTradingDataSupportCoinV5RespDataInner) GetOptionOk() ([]string, bool) {
	if o == nil || IsNil(o.Option) {
		return nil, false
	}
	return o.Option, true
}

// HasOption returns a boolean if a field has been set.
func (o *GetRubikStatTradingDataSupportCoinV5RespDataInner) HasOption() bool {
	if o != nil && !IsNil(o.Option) {
		return true
	}

	return false
}

// SetOption gets a reference to the given []string and assigns it to the Option field.
func (o *GetRubikStatTradingDataSupportCoinV5RespDataInner) SetOption(v []string) {
	o.Option = v
}

// GetSpot returns the Spot field value if set, zero value otherwise.
func (o *GetRubikStatTradingDataSupportCoinV5RespDataInner) GetSpot() []string {
	if o == nil || IsNil(o.Spot) {
		var ret []string
		return ret
	}
	return o.Spot
}

// GetSpotOk returns a tuple with the Spot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRubikStatTradingDataSupportCoinV5RespDataInner) GetSpotOk() ([]string, bool) {
	if o == nil || IsNil(o.Spot) {
		return nil, false
	}
	return o.Spot, true
}

// HasSpot returns a boolean if a field has been set.
func (o *GetRubikStatTradingDataSupportCoinV5RespDataInner) HasSpot() bool {
	if o != nil && !IsNil(o.Spot) {
		return true
	}

	return false
}

// SetSpot gets a reference to the given []string and assigns it to the Spot field.
func (o *GetRubikStatTradingDataSupportCoinV5RespDataInner) SetSpot(v []string) {
	o.Spot = v
}

func (o GetRubikStatTradingDataSupportCoinV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRubikStatTradingDataSupportCoinV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Contract) {
		toSerialize["contract"] = o.Contract
	}
	if !IsNil(o.Option) {
		toSerialize["option"] = o.Option
	}
	if !IsNil(o.Spot) {
		toSerialize["spot"] = o.Spot
	}
	return toSerialize, nil
}

type NullableGetRubikStatTradingDataSupportCoinV5RespDataInner struct {
	value *GetRubikStatTradingDataSupportCoinV5RespDataInner
	isSet bool
}

func (v NullableGetRubikStatTradingDataSupportCoinV5RespDataInner) Get() *GetRubikStatTradingDataSupportCoinV5RespDataInner {
	return v.value
}

func (v *NullableGetRubikStatTradingDataSupportCoinV5RespDataInner) Set(val *GetRubikStatTradingDataSupportCoinV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRubikStatTradingDataSupportCoinV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRubikStatTradingDataSupportCoinV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRubikStatTradingDataSupportCoinV5RespDataInner(val *GetRubikStatTradingDataSupportCoinV5RespDataInner) *NullableGetRubikStatTradingDataSupportCoinV5RespDataInner {
	return &NullableGetRubikStatTradingDataSupportCoinV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetRubikStatTradingDataSupportCoinV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRubikStatTradingDataSupportCoinV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


