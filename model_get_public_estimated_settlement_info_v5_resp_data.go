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

// checks if the GetPublicEstimatedSettlementInfoV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPublicEstimatedSettlementInfoV5RespData{}

// GetPublicEstimatedSettlementInfoV5RespData struct for GetPublicEstimatedSettlementInfoV5RespData
type GetPublicEstimatedSettlementInfoV5RespData struct {
	// Estimated settlement price
	EstSettlePx *string `json:"estSettlePx,omitempty"`
	// Instrument ID, e.g. `XRP-USDT-250307`
	InstId *string `json:"instId,omitempty"`
	// Next settlement time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	NextSettleTime *string `json:"nextSettleTime,omitempty"`
	// Data return time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	Ts *string `json:"ts,omitempty"`
}

// NewGetPublicEstimatedSettlementInfoV5RespData instantiates a new GetPublicEstimatedSettlementInfoV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPublicEstimatedSettlementInfoV5RespData() *GetPublicEstimatedSettlementInfoV5RespData {
	this := GetPublicEstimatedSettlementInfoV5RespData{}
	var estSettlePx string = ""
	this.EstSettlePx = &estSettlePx
	var instId string = ""
	this.InstId = &instId
	var nextSettleTime string = ""
	this.NextSettleTime = &nextSettleTime
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewGetPublicEstimatedSettlementInfoV5RespDataWithDefaults instantiates a new GetPublicEstimatedSettlementInfoV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPublicEstimatedSettlementInfoV5RespDataWithDefaults() *GetPublicEstimatedSettlementInfoV5RespData {
	this := GetPublicEstimatedSettlementInfoV5RespData{}
	var estSettlePx string = ""
	this.EstSettlePx = &estSettlePx
	var instId string = ""
	this.InstId = &instId
	var nextSettleTime string = ""
	this.NextSettleTime = &nextSettleTime
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetEstSettlePx returns the EstSettlePx field value if set, zero value otherwise.
func (o *GetPublicEstimatedSettlementInfoV5RespData) GetEstSettlePx() string {
	if o == nil || IsNil(o.EstSettlePx) {
		var ret string
		return ret
	}
	return *o.EstSettlePx
}

// GetEstSettlePxOk returns a tuple with the EstSettlePx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicEstimatedSettlementInfoV5RespData) GetEstSettlePxOk() (*string, bool) {
	if o == nil || IsNil(o.EstSettlePx) {
		return nil, false
	}
	return o.EstSettlePx, true
}

// HasEstSettlePx returns a boolean if a field has been set.
func (o *GetPublicEstimatedSettlementInfoV5RespData) HasEstSettlePx() bool {
	if o != nil && !IsNil(o.EstSettlePx) {
		return true
	}

	return false
}

// SetEstSettlePx gets a reference to the given string and assigns it to the EstSettlePx field.
func (o *GetPublicEstimatedSettlementInfoV5RespData) SetEstSettlePx(v string) {
	o.EstSettlePx = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetPublicEstimatedSettlementInfoV5RespData) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicEstimatedSettlementInfoV5RespData) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetPublicEstimatedSettlementInfoV5RespData) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetPublicEstimatedSettlementInfoV5RespData) SetInstId(v string) {
	o.InstId = &v
}

// GetNextSettleTime returns the NextSettleTime field value if set, zero value otherwise.
func (o *GetPublicEstimatedSettlementInfoV5RespData) GetNextSettleTime() string {
	if o == nil || IsNil(o.NextSettleTime) {
		var ret string
		return ret
	}
	return *o.NextSettleTime
}

// GetNextSettleTimeOk returns a tuple with the NextSettleTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicEstimatedSettlementInfoV5RespData) GetNextSettleTimeOk() (*string, bool) {
	if o == nil || IsNil(o.NextSettleTime) {
		return nil, false
	}
	return o.NextSettleTime, true
}

// HasNextSettleTime returns a boolean if a field has been set.
func (o *GetPublicEstimatedSettlementInfoV5RespData) HasNextSettleTime() bool {
	if o != nil && !IsNil(o.NextSettleTime) {
		return true
	}

	return false
}

// SetNextSettleTime gets a reference to the given string and assigns it to the NextSettleTime field.
func (o *GetPublicEstimatedSettlementInfoV5RespData) SetNextSettleTime(v string) {
	o.NextSettleTime = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetPublicEstimatedSettlementInfoV5RespData) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicEstimatedSettlementInfoV5RespData) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetPublicEstimatedSettlementInfoV5RespData) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetPublicEstimatedSettlementInfoV5RespData) SetTs(v string) {
	o.Ts = &v
}

func (o GetPublicEstimatedSettlementInfoV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPublicEstimatedSettlementInfoV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EstSettlePx) {
		toSerialize["estSettlePx"] = o.EstSettlePx
	}
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	if !IsNil(o.NextSettleTime) {
		toSerialize["nextSettleTime"] = o.NextSettleTime
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableGetPublicEstimatedSettlementInfoV5RespData struct {
	value *GetPublicEstimatedSettlementInfoV5RespData
	isSet bool
}

func (v NullableGetPublicEstimatedSettlementInfoV5RespData) Get() *GetPublicEstimatedSettlementInfoV5RespData {
	return v.value
}

func (v *NullableGetPublicEstimatedSettlementInfoV5RespData) Set(val *GetPublicEstimatedSettlementInfoV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPublicEstimatedSettlementInfoV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPublicEstimatedSettlementInfoV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPublicEstimatedSettlementInfoV5RespData(val *GetPublicEstimatedSettlementInfoV5RespData) *NullableGetPublicEstimatedSettlementInfoV5RespData {
	return &NullableGetPublicEstimatedSettlementInfoV5RespData{value: val, isSet: true}
}

func (v NullableGetPublicEstimatedSettlementInfoV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPublicEstimatedSettlementInfoV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


