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

// checks if the GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner{}

// GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner struct for GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner
type GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner struct {
	// Total call open interest (`coin` as the unit)
	CallOI *string `json:"callOI,omitempty"`
	// Total call trading volume (`coin` as the unit)
	CallVol *string `json:"callVol,omitempty"`
	// Total put open interest (`coin` as the unit)
	PutOI *string `json:"putOI,omitempty"`
	// Total put trading volume (`coin` as the unit)
	PutVol *string `json:"putVol,omitempty"`
	// Strike price
	Strike *string `json:"strike,omitempty"`
	// Timestamp
	Ts *string `json:"ts,omitempty"`
}

// NewGetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner instantiates a new GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner() *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner {
	this := GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner{}
	var callOI string = ""
	this.CallOI = &callOI
	var callVol string = ""
	this.CallVol = &callVol
	var putOI string = ""
	this.PutOI = &putOI
	var putVol string = ""
	this.PutVol = &putVol
	var strike string = ""
	this.Strike = &strike
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewGetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInnerWithDefaults instantiates a new GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInnerWithDefaults() *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner {
	this := GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner{}
	var callOI string = ""
	this.CallOI = &callOI
	var callVol string = ""
	this.CallVol = &callVol
	var putOI string = ""
	this.PutOI = &putOI
	var putVol string = ""
	this.PutVol = &putVol
	var strike string = ""
	this.Strike = &strike
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetCallOI returns the CallOI field value if set, zero value otherwise.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetCallOI() string {
	if o == nil || IsNil(o.CallOI) {
		var ret string
		return ret
	}
	return *o.CallOI
}

// GetCallOIOk returns a tuple with the CallOI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetCallOIOk() (*string, bool) {
	if o == nil || IsNil(o.CallOI) {
		return nil, false
	}
	return o.CallOI, true
}

// HasCallOI returns a boolean if a field has been set.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) HasCallOI() bool {
	if o != nil && !IsNil(o.CallOI) {
		return true
	}

	return false
}

// SetCallOI gets a reference to the given string and assigns it to the CallOI field.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) SetCallOI(v string) {
	o.CallOI = &v
}

// GetCallVol returns the CallVol field value if set, zero value otherwise.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetCallVol() string {
	if o == nil || IsNil(o.CallVol) {
		var ret string
		return ret
	}
	return *o.CallVol
}

// GetCallVolOk returns a tuple with the CallVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetCallVolOk() (*string, bool) {
	if o == nil || IsNil(o.CallVol) {
		return nil, false
	}
	return o.CallVol, true
}

// HasCallVol returns a boolean if a field has been set.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) HasCallVol() bool {
	if o != nil && !IsNil(o.CallVol) {
		return true
	}

	return false
}

// SetCallVol gets a reference to the given string and assigns it to the CallVol field.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) SetCallVol(v string) {
	o.CallVol = &v
}

// GetPutOI returns the PutOI field value if set, zero value otherwise.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetPutOI() string {
	if o == nil || IsNil(o.PutOI) {
		var ret string
		return ret
	}
	return *o.PutOI
}

// GetPutOIOk returns a tuple with the PutOI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetPutOIOk() (*string, bool) {
	if o == nil || IsNil(o.PutOI) {
		return nil, false
	}
	return o.PutOI, true
}

// HasPutOI returns a boolean if a field has been set.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) HasPutOI() bool {
	if o != nil && !IsNil(o.PutOI) {
		return true
	}

	return false
}

// SetPutOI gets a reference to the given string and assigns it to the PutOI field.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) SetPutOI(v string) {
	o.PutOI = &v
}

// GetPutVol returns the PutVol field value if set, zero value otherwise.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetPutVol() string {
	if o == nil || IsNil(o.PutVol) {
		var ret string
		return ret
	}
	return *o.PutVol
}

// GetPutVolOk returns a tuple with the PutVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetPutVolOk() (*string, bool) {
	if o == nil || IsNil(o.PutVol) {
		return nil, false
	}
	return o.PutVol, true
}

// HasPutVol returns a boolean if a field has been set.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) HasPutVol() bool {
	if o != nil && !IsNil(o.PutVol) {
		return true
	}

	return false
}

// SetPutVol gets a reference to the given string and assigns it to the PutVol field.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) SetPutVol(v string) {
	o.PutVol = &v
}

// GetStrike returns the Strike field value if set, zero value otherwise.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetStrike() string {
	if o == nil || IsNil(o.Strike) {
		var ret string
		return ret
	}
	return *o.Strike
}

// GetStrikeOk returns a tuple with the Strike field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetStrikeOk() (*string, bool) {
	if o == nil || IsNil(o.Strike) {
		return nil, false
	}
	return o.Strike, true
}

// HasStrike returns a boolean if a field has been set.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) HasStrike() bool {
	if o != nil && !IsNil(o.Strike) {
		return true
	}

	return false
}

// SetStrike gets a reference to the given string and assigns it to the Strike field.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) SetStrike(v string) {
	o.Strike = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

func (o GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallOI) {
		toSerialize["callOI"] = o.CallOI
	}
	if !IsNil(o.CallVol) {
		toSerialize["callVol"] = o.CallVol
	}
	if !IsNil(o.PutOI) {
		toSerialize["putOI"] = o.PutOI
	}
	if !IsNil(o.PutVol) {
		toSerialize["putVol"] = o.PutVol
	}
	if !IsNil(o.Strike) {
		toSerialize["strike"] = o.Strike
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableGetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner struct {
	value *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner
	isSet bool
}

func (v NullableGetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) Get() *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner {
	return v.value
}

func (v *NullableGetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) Set(val *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner(val *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) *NullableGetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner {
	return &NullableGetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


