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

// checks if the GetRfqMakerInstrumentSettingsV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRfqMakerInstrumentSettingsV5RespDataInner{}

// GetRfqMakerInstrumentSettingsV5RespDataInner struct for GetRfqMakerInstrumentSettingsV5RespDataInner
type GetRfqMakerInstrumentSettingsV5RespDataInner struct {
	// Elements of the instType.
	Data []CreateRfqMakerInstrumentSettingsV5ReqDataInner `json:"data,omitempty"`
	// Receive all instruments or not under specific instType setting.   Valid value can be boolean (`True`/`False`). By default, the value will be `false`.
	IncludeAll *bool `json:"includeAll,omitempty"`
	// Instrument family. Required for `FUTURES`, `OPTION` and `SWAP` only.
	InstFamily *string `json:"instFamily,omitempty"`
	// Instrument ID. Required for `SPOT` only.
	InstId *string `json:"instId,omitempty"`
	// Type of instrument. Valid value can be `FUTURES`, `OPTION`, `SWAP` or `SPOT`.
	InstType *string `json:"instType,omitempty"`
	// Price bands in unit of ticks, measured against mark price.   Setting makerPxBand to 1 tick means:   If Bid price > Mark + 1 tick, it will be stopped   If Ask price < Mark - 1 tick, It will be stopped
	MakerPxBand *string `json:"makerPxBand,omitempty"`
	// Max trade quantity for the product(s).   For `FUTURES`, `OPTION` and `SWAP`, the max quantity of the RFQ/Quote is in unit of contracts. For `SPOT`, this parameter is in base currency.
	MaxBlockSz *string `json:"maxBlockSz,omitempty"`
}

// NewGetRfqMakerInstrumentSettingsV5RespDataInner instantiates a new GetRfqMakerInstrumentSettingsV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRfqMakerInstrumentSettingsV5RespDataInner() *GetRfqMakerInstrumentSettingsV5RespDataInner {
	this := GetRfqMakerInstrumentSettingsV5RespDataInner{}
	var instFamily string = ""
	this.InstFamily = &instFamily
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var makerPxBand string = ""
	this.MakerPxBand = &makerPxBand
	var maxBlockSz string = ""
	this.MaxBlockSz = &maxBlockSz
	return &this
}

// NewGetRfqMakerInstrumentSettingsV5RespDataInnerWithDefaults instantiates a new GetRfqMakerInstrumentSettingsV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRfqMakerInstrumentSettingsV5RespDataInnerWithDefaults() *GetRfqMakerInstrumentSettingsV5RespDataInner {
	this := GetRfqMakerInstrumentSettingsV5RespDataInner{}
	var instFamily string = ""
	this.InstFamily = &instFamily
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var makerPxBand string = ""
	this.MakerPxBand = &makerPxBand
	var maxBlockSz string = ""
	this.MaxBlockSz = &maxBlockSz
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) GetData() []CreateRfqMakerInstrumentSettingsV5ReqDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []CreateRfqMakerInstrumentSettingsV5ReqDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) GetDataOk() ([]CreateRfqMakerInstrumentSettingsV5ReqDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CreateRfqMakerInstrumentSettingsV5ReqDataInner and assigns it to the Data field.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) SetData(v []CreateRfqMakerInstrumentSettingsV5ReqDataInner) {
	o.Data = v
}

// GetIncludeAll returns the IncludeAll field value if set, zero value otherwise.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) GetIncludeAll() bool {
	if o == nil || IsNil(o.IncludeAll) {
		var ret bool
		return ret
	}
	return *o.IncludeAll
}

// GetIncludeAllOk returns a tuple with the IncludeAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) GetIncludeAllOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeAll) {
		return nil, false
	}
	return o.IncludeAll, true
}

// HasIncludeAll returns a boolean if a field has been set.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) HasIncludeAll() bool {
	if o != nil && !IsNil(o.IncludeAll) {
		return true
	}

	return false
}

// SetIncludeAll gets a reference to the given bool and assigns it to the IncludeAll field.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) SetIncludeAll(v bool) {
	o.IncludeAll = &v
}

// GetInstFamily returns the InstFamily field value if set, zero value otherwise.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) GetInstFamily() string {
	if o == nil || IsNil(o.InstFamily) {
		var ret string
		return ret
	}
	return *o.InstFamily
}

// GetInstFamilyOk returns a tuple with the InstFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) GetInstFamilyOk() (*string, bool) {
	if o == nil || IsNil(o.InstFamily) {
		return nil, false
	}
	return o.InstFamily, true
}

// HasInstFamily returns a boolean if a field has been set.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) HasInstFamily() bool {
	if o != nil && !IsNil(o.InstFamily) {
		return true
	}

	return false
}

// SetInstFamily gets a reference to the given string and assigns it to the InstFamily field.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) SetInstFamily(v string) {
	o.InstFamily = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) SetInstId(v string) {
	o.InstId = &v
}

// GetInstType returns the InstType field value if set, zero value otherwise.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) GetInstType() string {
	if o == nil || IsNil(o.InstType) {
		var ret string
		return ret
	}
	return *o.InstType
}

// GetInstTypeOk returns a tuple with the InstType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) GetInstTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstType) {
		return nil, false
	}
	return o.InstType, true
}

// HasInstType returns a boolean if a field has been set.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) HasInstType() bool {
	if o != nil && !IsNil(o.InstType) {
		return true
	}

	return false
}

// SetInstType gets a reference to the given string and assigns it to the InstType field.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) SetInstType(v string) {
	o.InstType = &v
}

// GetMakerPxBand returns the MakerPxBand field value if set, zero value otherwise.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) GetMakerPxBand() string {
	if o == nil || IsNil(o.MakerPxBand) {
		var ret string
		return ret
	}
	return *o.MakerPxBand
}

// GetMakerPxBandOk returns a tuple with the MakerPxBand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) GetMakerPxBandOk() (*string, bool) {
	if o == nil || IsNil(o.MakerPxBand) {
		return nil, false
	}
	return o.MakerPxBand, true
}

// HasMakerPxBand returns a boolean if a field has been set.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) HasMakerPxBand() bool {
	if o != nil && !IsNil(o.MakerPxBand) {
		return true
	}

	return false
}

// SetMakerPxBand gets a reference to the given string and assigns it to the MakerPxBand field.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) SetMakerPxBand(v string) {
	o.MakerPxBand = &v
}

// GetMaxBlockSz returns the MaxBlockSz field value if set, zero value otherwise.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) GetMaxBlockSz() string {
	if o == nil || IsNil(o.MaxBlockSz) {
		var ret string
		return ret
	}
	return *o.MaxBlockSz
}

// GetMaxBlockSzOk returns a tuple with the MaxBlockSz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) GetMaxBlockSzOk() (*string, bool) {
	if o == nil || IsNil(o.MaxBlockSz) {
		return nil, false
	}
	return o.MaxBlockSz, true
}

// HasMaxBlockSz returns a boolean if a field has been set.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) HasMaxBlockSz() bool {
	if o != nil && !IsNil(o.MaxBlockSz) {
		return true
	}

	return false
}

// SetMaxBlockSz gets a reference to the given string and assigns it to the MaxBlockSz field.
func (o *GetRfqMakerInstrumentSettingsV5RespDataInner) SetMaxBlockSz(v string) {
	o.MaxBlockSz = &v
}

func (o GetRfqMakerInstrumentSettingsV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRfqMakerInstrumentSettingsV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.IncludeAll) {
		toSerialize["includeAll"] = o.IncludeAll
	}
	if !IsNil(o.InstFamily) {
		toSerialize["instFamily"] = o.InstFamily
	}
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	if !IsNil(o.InstType) {
		toSerialize["instType"] = o.InstType
	}
	if !IsNil(o.MakerPxBand) {
		toSerialize["makerPxBand"] = o.MakerPxBand
	}
	if !IsNil(o.MaxBlockSz) {
		toSerialize["maxBlockSz"] = o.MaxBlockSz
	}
	return toSerialize, nil
}

type NullableGetRfqMakerInstrumentSettingsV5RespDataInner struct {
	value *GetRfqMakerInstrumentSettingsV5RespDataInner
	isSet bool
}

func (v NullableGetRfqMakerInstrumentSettingsV5RespDataInner) Get() *GetRfqMakerInstrumentSettingsV5RespDataInner {
	return v.value
}

func (v *NullableGetRfqMakerInstrumentSettingsV5RespDataInner) Set(val *GetRfqMakerInstrumentSettingsV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRfqMakerInstrumentSettingsV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRfqMakerInstrumentSettingsV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRfqMakerInstrumentSettingsV5RespDataInner(val *GetRfqMakerInstrumentSettingsV5RespDataInner) *NullableGetRfqMakerInstrumentSettingsV5RespDataInner {
	return &NullableGetRfqMakerInstrumentSettingsV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetRfqMakerInstrumentSettingsV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRfqMakerInstrumentSettingsV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


