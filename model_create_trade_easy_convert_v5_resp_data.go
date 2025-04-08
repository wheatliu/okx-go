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

// checks if the CreateTradeEasyConvertV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradeEasyConvertV5RespData{}

// CreateTradeEasyConvertV5RespData struct for CreateTradeEasyConvertV5RespData
type CreateTradeEasyConvertV5RespData struct {
	// Filled amount of small payment currency convert from
	FillFromSz *string `json:"fillFromSz,omitempty"`
	// Filled amount of mainstream currency convert to
	FillToSz *string `json:"fillToSz,omitempty"`
	// Type of small payment currency convert from
	FromCcy *string `json:"fromCcy,omitempty"`
	// Current status of easy convert   `running`: Running   `filled`: Filled   `failed`: Failed
	Status *string `json:"status,omitempty"`
	// Type of mainstream currency convert to
	ToCcy *string `json:"toCcy,omitempty"`
	// Trade time, Unix timestamp format in milliseconds, e.g. 1597026383085
	UTime *string `json:"uTime,omitempty"`
}

// NewCreateTradeEasyConvertV5RespData instantiates a new CreateTradeEasyConvertV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradeEasyConvertV5RespData() *CreateTradeEasyConvertV5RespData {
	this := CreateTradeEasyConvertV5RespData{}
	var fillFromSz string = ""
	this.FillFromSz = &fillFromSz
	var fillToSz string = ""
	this.FillToSz = &fillToSz
	var fromCcy string = ""
	this.FromCcy = &fromCcy
	var status string = ""
	this.Status = &status
	var toCcy string = ""
	this.ToCcy = &toCcy
	var uTime string = ""
	this.UTime = &uTime
	return &this
}

// NewCreateTradeEasyConvertV5RespDataWithDefaults instantiates a new CreateTradeEasyConvertV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradeEasyConvertV5RespDataWithDefaults() *CreateTradeEasyConvertV5RespData {
	this := CreateTradeEasyConvertV5RespData{}
	var fillFromSz string = ""
	this.FillFromSz = &fillFromSz
	var fillToSz string = ""
	this.FillToSz = &fillToSz
	var fromCcy string = ""
	this.FromCcy = &fromCcy
	var status string = ""
	this.Status = &status
	var toCcy string = ""
	this.ToCcy = &toCcy
	var uTime string = ""
	this.UTime = &uTime
	return &this
}

// GetFillFromSz returns the FillFromSz field value if set, zero value otherwise.
func (o *CreateTradeEasyConvertV5RespData) GetFillFromSz() string {
	if o == nil || IsNil(o.FillFromSz) {
		var ret string
		return ret
	}
	return *o.FillFromSz
}

// GetFillFromSzOk returns a tuple with the FillFromSz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeEasyConvertV5RespData) GetFillFromSzOk() (*string, bool) {
	if o == nil || IsNil(o.FillFromSz) {
		return nil, false
	}
	return o.FillFromSz, true
}

// HasFillFromSz returns a boolean if a field has been set.
func (o *CreateTradeEasyConvertV5RespData) HasFillFromSz() bool {
	if o != nil && !IsNil(o.FillFromSz) {
		return true
	}

	return false
}

// SetFillFromSz gets a reference to the given string and assigns it to the FillFromSz field.
func (o *CreateTradeEasyConvertV5RespData) SetFillFromSz(v string) {
	o.FillFromSz = &v
}

// GetFillToSz returns the FillToSz field value if set, zero value otherwise.
func (o *CreateTradeEasyConvertV5RespData) GetFillToSz() string {
	if o == nil || IsNil(o.FillToSz) {
		var ret string
		return ret
	}
	return *o.FillToSz
}

// GetFillToSzOk returns a tuple with the FillToSz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeEasyConvertV5RespData) GetFillToSzOk() (*string, bool) {
	if o == nil || IsNil(o.FillToSz) {
		return nil, false
	}
	return o.FillToSz, true
}

// HasFillToSz returns a boolean if a field has been set.
func (o *CreateTradeEasyConvertV5RespData) HasFillToSz() bool {
	if o != nil && !IsNil(o.FillToSz) {
		return true
	}

	return false
}

// SetFillToSz gets a reference to the given string and assigns it to the FillToSz field.
func (o *CreateTradeEasyConvertV5RespData) SetFillToSz(v string) {
	o.FillToSz = &v
}

// GetFromCcy returns the FromCcy field value if set, zero value otherwise.
func (o *CreateTradeEasyConvertV5RespData) GetFromCcy() string {
	if o == nil || IsNil(o.FromCcy) {
		var ret string
		return ret
	}
	return *o.FromCcy
}

// GetFromCcyOk returns a tuple with the FromCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeEasyConvertV5RespData) GetFromCcyOk() (*string, bool) {
	if o == nil || IsNil(o.FromCcy) {
		return nil, false
	}
	return o.FromCcy, true
}

// HasFromCcy returns a boolean if a field has been set.
func (o *CreateTradeEasyConvertV5RespData) HasFromCcy() bool {
	if o != nil && !IsNil(o.FromCcy) {
		return true
	}

	return false
}

// SetFromCcy gets a reference to the given string and assigns it to the FromCcy field.
func (o *CreateTradeEasyConvertV5RespData) SetFromCcy(v string) {
	o.FromCcy = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateTradeEasyConvertV5RespData) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeEasyConvertV5RespData) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateTradeEasyConvertV5RespData) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateTradeEasyConvertV5RespData) SetStatus(v string) {
	o.Status = &v
}

// GetToCcy returns the ToCcy field value if set, zero value otherwise.
func (o *CreateTradeEasyConvertV5RespData) GetToCcy() string {
	if o == nil || IsNil(o.ToCcy) {
		var ret string
		return ret
	}
	return *o.ToCcy
}

// GetToCcyOk returns a tuple with the ToCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeEasyConvertV5RespData) GetToCcyOk() (*string, bool) {
	if o == nil || IsNil(o.ToCcy) {
		return nil, false
	}
	return o.ToCcy, true
}

// HasToCcy returns a boolean if a field has been set.
func (o *CreateTradeEasyConvertV5RespData) HasToCcy() bool {
	if o != nil && !IsNil(o.ToCcy) {
		return true
	}

	return false
}

// SetToCcy gets a reference to the given string and assigns it to the ToCcy field.
func (o *CreateTradeEasyConvertV5RespData) SetToCcy(v string) {
	o.ToCcy = &v
}

// GetUTime returns the UTime field value if set, zero value otherwise.
func (o *CreateTradeEasyConvertV5RespData) GetUTime() string {
	if o == nil || IsNil(o.UTime) {
		var ret string
		return ret
	}
	return *o.UTime
}

// GetUTimeOk returns a tuple with the UTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeEasyConvertV5RespData) GetUTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UTime) {
		return nil, false
	}
	return o.UTime, true
}

// HasUTime returns a boolean if a field has been set.
func (o *CreateTradeEasyConvertV5RespData) HasUTime() bool {
	if o != nil && !IsNil(o.UTime) {
		return true
	}

	return false
}

// SetUTime gets a reference to the given string and assigns it to the UTime field.
func (o *CreateTradeEasyConvertV5RespData) SetUTime(v string) {
	o.UTime = &v
}

func (o CreateTradeEasyConvertV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradeEasyConvertV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FillFromSz) {
		toSerialize["fillFromSz"] = o.FillFromSz
	}
	if !IsNil(o.FillToSz) {
		toSerialize["fillToSz"] = o.FillToSz
	}
	if !IsNil(o.FromCcy) {
		toSerialize["fromCcy"] = o.FromCcy
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ToCcy) {
		toSerialize["toCcy"] = o.ToCcy
	}
	if !IsNil(o.UTime) {
		toSerialize["uTime"] = o.UTime
	}
	return toSerialize, nil
}

type NullableCreateTradeEasyConvertV5RespData struct {
	value *CreateTradeEasyConvertV5RespData
	isSet bool
}

func (v NullableCreateTradeEasyConvertV5RespData) Get() *CreateTradeEasyConvertV5RespData {
	return v.value
}

func (v *NullableCreateTradeEasyConvertV5RespData) Set(val *CreateTradeEasyConvertV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradeEasyConvertV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradeEasyConvertV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradeEasyConvertV5RespData(val *CreateTradeEasyConvertV5RespData) *NullableCreateTradeEasyConvertV5RespData {
	return &NullableCreateTradeEasyConvertV5RespData{value: val, isSet: true}
}

func (v NullableCreateTradeEasyConvertV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradeEasyConvertV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


