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

// checks if the GetAccountPositionTiersV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountPositionTiersV5RespData{}

// GetAccountPositionTiersV5RespData struct for GetAccountPositionTiersV5RespData
type GetAccountPositionTiersV5RespData struct {
	// Instrument family  Applicable to `FUTURES`/`SWAP`/`OPTION`
	InstFamily *string `json:"instFamily,omitempty"`
	// Max number of positions
	MaxSz *string `json:"maxSz,omitempty"`
	// Limitation of position type, only applicable to cross `OPTION` under portfolio margin mode   `1`: Contracts of pending orders and open positions for all derivatives instruments. `2`: Contracts of pending orders for all derivatives instruments. `3`: Pending orders for all derivatives instruments. `4`: Contracts of pending orders and open positions for all derivatives instruments on the same side. `5`: Pending orders for one derivatives instrument. `6`: Contracts of pending orders and open positions for one derivatives instrument. `7`: Contracts of one pending order.
	PosType *string `json:"posType,omitempty"`
	// Underlying  Applicable to `FUTURES`/`SWAP`/`OPTION`
	Uly *string `json:"uly,omitempty"`
}

// NewGetAccountPositionTiersV5RespData instantiates a new GetAccountPositionTiersV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountPositionTiersV5RespData() *GetAccountPositionTiersV5RespData {
	this := GetAccountPositionTiersV5RespData{}
	var instFamily string = ""
	this.InstFamily = &instFamily
	var maxSz string = ""
	this.MaxSz = &maxSz
	var posType string = ""
	this.PosType = &posType
	var uly string = ""
	this.Uly = &uly
	return &this
}

// NewGetAccountPositionTiersV5RespDataWithDefaults instantiates a new GetAccountPositionTiersV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountPositionTiersV5RespDataWithDefaults() *GetAccountPositionTiersV5RespData {
	this := GetAccountPositionTiersV5RespData{}
	var instFamily string = ""
	this.InstFamily = &instFamily
	var maxSz string = ""
	this.MaxSz = &maxSz
	var posType string = ""
	this.PosType = &posType
	var uly string = ""
	this.Uly = &uly
	return &this
}

// GetInstFamily returns the InstFamily field value if set, zero value otherwise.
func (o *GetAccountPositionTiersV5RespData) GetInstFamily() string {
	if o == nil || IsNil(o.InstFamily) {
		var ret string
		return ret
	}
	return *o.InstFamily
}

// GetInstFamilyOk returns a tuple with the InstFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountPositionTiersV5RespData) GetInstFamilyOk() (*string, bool) {
	if o == nil || IsNil(o.InstFamily) {
		return nil, false
	}
	return o.InstFamily, true
}

// HasInstFamily returns a boolean if a field has been set.
func (o *GetAccountPositionTiersV5RespData) HasInstFamily() bool {
	if o != nil && !IsNil(o.InstFamily) {
		return true
	}

	return false
}

// SetInstFamily gets a reference to the given string and assigns it to the InstFamily field.
func (o *GetAccountPositionTiersV5RespData) SetInstFamily(v string) {
	o.InstFamily = &v
}

// GetMaxSz returns the MaxSz field value if set, zero value otherwise.
func (o *GetAccountPositionTiersV5RespData) GetMaxSz() string {
	if o == nil || IsNil(o.MaxSz) {
		var ret string
		return ret
	}
	return *o.MaxSz
}

// GetMaxSzOk returns a tuple with the MaxSz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountPositionTiersV5RespData) GetMaxSzOk() (*string, bool) {
	if o == nil || IsNil(o.MaxSz) {
		return nil, false
	}
	return o.MaxSz, true
}

// HasMaxSz returns a boolean if a field has been set.
func (o *GetAccountPositionTiersV5RespData) HasMaxSz() bool {
	if o != nil && !IsNil(o.MaxSz) {
		return true
	}

	return false
}

// SetMaxSz gets a reference to the given string and assigns it to the MaxSz field.
func (o *GetAccountPositionTiersV5RespData) SetMaxSz(v string) {
	o.MaxSz = &v
}

// GetPosType returns the PosType field value if set, zero value otherwise.
func (o *GetAccountPositionTiersV5RespData) GetPosType() string {
	if o == nil || IsNil(o.PosType) {
		var ret string
		return ret
	}
	return *o.PosType
}

// GetPosTypeOk returns a tuple with the PosType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountPositionTiersV5RespData) GetPosTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PosType) {
		return nil, false
	}
	return o.PosType, true
}

// HasPosType returns a boolean if a field has been set.
func (o *GetAccountPositionTiersV5RespData) HasPosType() bool {
	if o != nil && !IsNil(o.PosType) {
		return true
	}

	return false
}

// SetPosType gets a reference to the given string and assigns it to the PosType field.
func (o *GetAccountPositionTiersV5RespData) SetPosType(v string) {
	o.PosType = &v
}

// GetUly returns the Uly field value if set, zero value otherwise.
func (o *GetAccountPositionTiersV5RespData) GetUly() string {
	if o == nil || IsNil(o.Uly) {
		var ret string
		return ret
	}
	return *o.Uly
}

// GetUlyOk returns a tuple with the Uly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountPositionTiersV5RespData) GetUlyOk() (*string, bool) {
	if o == nil || IsNil(o.Uly) {
		return nil, false
	}
	return o.Uly, true
}

// HasUly returns a boolean if a field has been set.
func (o *GetAccountPositionTiersV5RespData) HasUly() bool {
	if o != nil && !IsNil(o.Uly) {
		return true
	}

	return false
}

// SetUly gets a reference to the given string and assigns it to the Uly field.
func (o *GetAccountPositionTiersV5RespData) SetUly(v string) {
	o.Uly = &v
}

func (o GetAccountPositionTiersV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountPositionTiersV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstFamily) {
		toSerialize["instFamily"] = o.InstFamily
	}
	if !IsNil(o.MaxSz) {
		toSerialize["maxSz"] = o.MaxSz
	}
	if !IsNil(o.PosType) {
		toSerialize["posType"] = o.PosType
	}
	if !IsNil(o.Uly) {
		toSerialize["uly"] = o.Uly
	}
	return toSerialize, nil
}

type NullableGetAccountPositionTiersV5RespData struct {
	value *GetAccountPositionTiersV5RespData
	isSet bool
}

func (v NullableGetAccountPositionTiersV5RespData) Get() *GetAccountPositionTiersV5RespData {
	return v.value
}

func (v *NullableGetAccountPositionTiersV5RespData) Set(val *GetAccountPositionTiersV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountPositionTiersV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountPositionTiersV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountPositionTiersV5RespData(val *GetAccountPositionTiersV5RespData) *NullableGetAccountPositionTiersV5RespData {
	return &NullableGetAccountPositionTiersV5RespData{value: val, isSet: true}
}

func (v NullableGetAccountPositionTiersV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountPositionTiersV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


