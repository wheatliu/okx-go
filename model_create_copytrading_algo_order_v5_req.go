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

// checks if the CreateCopytradingAlgoOrderV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCopytradingAlgoOrderV5Req{}

// CreateCopytradingAlgoOrderV5Req struct for CreateCopytradingAlgoOrderV5Req
type CreateCopytradingAlgoOrderV5Req struct {
	// Instrument type  `SPOT`  `SWAP`, the default value
	InstType *string `json:"instType,omitempty"`
	// Stop-loss order price  If the price is -1, stop-loss will be executed at the market price, the default is `-1`  Only applicable to `SPOT` lead trader
	SlOrdPx *string `json:"slOrdPx,omitempty"`
	// Stop-loss trigger price. Stop-loss order price will be the market price after triggering. The stop loss order will be deleted if it is 0
	SlTriggerPx *string `json:"slTriggerPx,omitempty"`
	// Stop-loss trigger price type  `last`: last price   `index`: index price   `mark`: mark price   Default is last
	SlTriggerPxType *string `json:"slTriggerPxType,omitempty"`
	// Lead position ID
	SubPosId string `json:"subPosId"`
	// Order tag  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 16 characters.
	Tag *string `json:"tag,omitempty"`
	// Take-profit order price  If the price is -1, take-profit will be executed at the market price, the default is `-1`  Only applicable to `SPOT` lead trader
	TpOrdPx *string `json:"tpOrdPx,omitempty"`
	// Take-profit trigger price. Take-profit order price will be the market price after triggering. At least one of tpTriggerPx and slTriggerPx must be filled  The take profit order will be deleted if it is 0
	TpTriggerPx *string `json:"tpTriggerPx,omitempty"`
	// Take-profit trigger price type     `last`: last price  `index`: index price  `mark`: mark price   Default is `last`
	TpTriggerPxType *string `json:"tpTriggerPxType,omitempty"`
}

type _CreateCopytradingAlgoOrderV5Req CreateCopytradingAlgoOrderV5Req

// NewCreateCopytradingAlgoOrderV5Req instantiates a new CreateCopytradingAlgoOrderV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCopytradingAlgoOrderV5Req(subPosId string) *CreateCopytradingAlgoOrderV5Req {
	this := CreateCopytradingAlgoOrderV5Req{}
	var instType string = ""
	this.InstType = &instType
	var slOrdPx string = ""
	this.SlOrdPx = &slOrdPx
	var slTriggerPx string = ""
	this.SlTriggerPx = &slTriggerPx
	var slTriggerPxType string = ""
	this.SlTriggerPxType = &slTriggerPxType
	this.SubPosId = subPosId
	var tag string = ""
	this.Tag = &tag
	var tpOrdPx string = ""
	this.TpOrdPx = &tpOrdPx
	var tpTriggerPx string = ""
	this.TpTriggerPx = &tpTriggerPx
	var tpTriggerPxType string = ""
	this.TpTriggerPxType = &tpTriggerPxType
	return &this
}

// NewCreateCopytradingAlgoOrderV5ReqWithDefaults instantiates a new CreateCopytradingAlgoOrderV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCopytradingAlgoOrderV5ReqWithDefaults() *CreateCopytradingAlgoOrderV5Req {
	this := CreateCopytradingAlgoOrderV5Req{}
	var instType string = ""
	this.InstType = &instType
	var slOrdPx string = ""
	this.SlOrdPx = &slOrdPx
	var slTriggerPx string = ""
	this.SlTriggerPx = &slTriggerPx
	var slTriggerPxType string = ""
	this.SlTriggerPxType = &slTriggerPxType
	var subPosId string = ""
	this.SubPosId = subPosId
	var tag string = ""
	this.Tag = &tag
	var tpOrdPx string = ""
	this.TpOrdPx = &tpOrdPx
	var tpTriggerPx string = ""
	this.TpTriggerPx = &tpTriggerPx
	var tpTriggerPxType string = ""
	this.TpTriggerPxType = &tpTriggerPxType
	return &this
}

// GetInstType returns the InstType field value if set, zero value otherwise.
func (o *CreateCopytradingAlgoOrderV5Req) GetInstType() string {
	if o == nil || IsNil(o.InstType) {
		var ret string
		return ret
	}
	return *o.InstType
}

// GetInstTypeOk returns a tuple with the InstType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAlgoOrderV5Req) GetInstTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstType) {
		return nil, false
	}
	return o.InstType, true
}

// HasInstType returns a boolean if a field has been set.
func (o *CreateCopytradingAlgoOrderV5Req) HasInstType() bool {
	if o != nil && !IsNil(o.InstType) {
		return true
	}

	return false
}

// SetInstType gets a reference to the given string and assigns it to the InstType field.
func (o *CreateCopytradingAlgoOrderV5Req) SetInstType(v string) {
	o.InstType = &v
}

// GetSlOrdPx returns the SlOrdPx field value if set, zero value otherwise.
func (o *CreateCopytradingAlgoOrderV5Req) GetSlOrdPx() string {
	if o == nil || IsNil(o.SlOrdPx) {
		var ret string
		return ret
	}
	return *o.SlOrdPx
}

// GetSlOrdPxOk returns a tuple with the SlOrdPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAlgoOrderV5Req) GetSlOrdPxOk() (*string, bool) {
	if o == nil || IsNil(o.SlOrdPx) {
		return nil, false
	}
	return o.SlOrdPx, true
}

// HasSlOrdPx returns a boolean if a field has been set.
func (o *CreateCopytradingAlgoOrderV5Req) HasSlOrdPx() bool {
	if o != nil && !IsNil(o.SlOrdPx) {
		return true
	}

	return false
}

// SetSlOrdPx gets a reference to the given string and assigns it to the SlOrdPx field.
func (o *CreateCopytradingAlgoOrderV5Req) SetSlOrdPx(v string) {
	o.SlOrdPx = &v
}

// GetSlTriggerPx returns the SlTriggerPx field value if set, zero value otherwise.
func (o *CreateCopytradingAlgoOrderV5Req) GetSlTriggerPx() string {
	if o == nil || IsNil(o.SlTriggerPx) {
		var ret string
		return ret
	}
	return *o.SlTriggerPx
}

// GetSlTriggerPxOk returns a tuple with the SlTriggerPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAlgoOrderV5Req) GetSlTriggerPxOk() (*string, bool) {
	if o == nil || IsNil(o.SlTriggerPx) {
		return nil, false
	}
	return o.SlTriggerPx, true
}

// HasSlTriggerPx returns a boolean if a field has been set.
func (o *CreateCopytradingAlgoOrderV5Req) HasSlTriggerPx() bool {
	if o != nil && !IsNil(o.SlTriggerPx) {
		return true
	}

	return false
}

// SetSlTriggerPx gets a reference to the given string and assigns it to the SlTriggerPx field.
func (o *CreateCopytradingAlgoOrderV5Req) SetSlTriggerPx(v string) {
	o.SlTriggerPx = &v
}

// GetSlTriggerPxType returns the SlTriggerPxType field value if set, zero value otherwise.
func (o *CreateCopytradingAlgoOrderV5Req) GetSlTriggerPxType() string {
	if o == nil || IsNil(o.SlTriggerPxType) {
		var ret string
		return ret
	}
	return *o.SlTriggerPxType
}

// GetSlTriggerPxTypeOk returns a tuple with the SlTriggerPxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAlgoOrderV5Req) GetSlTriggerPxTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SlTriggerPxType) {
		return nil, false
	}
	return o.SlTriggerPxType, true
}

// HasSlTriggerPxType returns a boolean if a field has been set.
func (o *CreateCopytradingAlgoOrderV5Req) HasSlTriggerPxType() bool {
	if o != nil && !IsNil(o.SlTriggerPxType) {
		return true
	}

	return false
}

// SetSlTriggerPxType gets a reference to the given string and assigns it to the SlTriggerPxType field.
func (o *CreateCopytradingAlgoOrderV5Req) SetSlTriggerPxType(v string) {
	o.SlTriggerPxType = &v
}

// GetSubPosId returns the SubPosId field value
func (o *CreateCopytradingAlgoOrderV5Req) GetSubPosId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubPosId
}

// GetSubPosIdOk returns a tuple with the SubPosId field value
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAlgoOrderV5Req) GetSubPosIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubPosId, true
}

// SetSubPosId sets field value
func (o *CreateCopytradingAlgoOrderV5Req) SetSubPosId(v string) {
	o.SubPosId = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *CreateCopytradingAlgoOrderV5Req) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAlgoOrderV5Req) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *CreateCopytradingAlgoOrderV5Req) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *CreateCopytradingAlgoOrderV5Req) SetTag(v string) {
	o.Tag = &v
}

// GetTpOrdPx returns the TpOrdPx field value if set, zero value otherwise.
func (o *CreateCopytradingAlgoOrderV5Req) GetTpOrdPx() string {
	if o == nil || IsNil(o.TpOrdPx) {
		var ret string
		return ret
	}
	return *o.TpOrdPx
}

// GetTpOrdPxOk returns a tuple with the TpOrdPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAlgoOrderV5Req) GetTpOrdPxOk() (*string, bool) {
	if o == nil || IsNil(o.TpOrdPx) {
		return nil, false
	}
	return o.TpOrdPx, true
}

// HasTpOrdPx returns a boolean if a field has been set.
func (o *CreateCopytradingAlgoOrderV5Req) HasTpOrdPx() bool {
	if o != nil && !IsNil(o.TpOrdPx) {
		return true
	}

	return false
}

// SetTpOrdPx gets a reference to the given string and assigns it to the TpOrdPx field.
func (o *CreateCopytradingAlgoOrderV5Req) SetTpOrdPx(v string) {
	o.TpOrdPx = &v
}

// GetTpTriggerPx returns the TpTriggerPx field value if set, zero value otherwise.
func (o *CreateCopytradingAlgoOrderV5Req) GetTpTriggerPx() string {
	if o == nil || IsNil(o.TpTriggerPx) {
		var ret string
		return ret
	}
	return *o.TpTriggerPx
}

// GetTpTriggerPxOk returns a tuple with the TpTriggerPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAlgoOrderV5Req) GetTpTriggerPxOk() (*string, bool) {
	if o == nil || IsNil(o.TpTriggerPx) {
		return nil, false
	}
	return o.TpTriggerPx, true
}

// HasTpTriggerPx returns a boolean if a field has been set.
func (o *CreateCopytradingAlgoOrderV5Req) HasTpTriggerPx() bool {
	if o != nil && !IsNil(o.TpTriggerPx) {
		return true
	}

	return false
}

// SetTpTriggerPx gets a reference to the given string and assigns it to the TpTriggerPx field.
func (o *CreateCopytradingAlgoOrderV5Req) SetTpTriggerPx(v string) {
	o.TpTriggerPx = &v
}

// GetTpTriggerPxType returns the TpTriggerPxType field value if set, zero value otherwise.
func (o *CreateCopytradingAlgoOrderV5Req) GetTpTriggerPxType() string {
	if o == nil || IsNil(o.TpTriggerPxType) {
		var ret string
		return ret
	}
	return *o.TpTriggerPxType
}

// GetTpTriggerPxTypeOk returns a tuple with the TpTriggerPxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAlgoOrderV5Req) GetTpTriggerPxTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TpTriggerPxType) {
		return nil, false
	}
	return o.TpTriggerPxType, true
}

// HasTpTriggerPxType returns a boolean if a field has been set.
func (o *CreateCopytradingAlgoOrderV5Req) HasTpTriggerPxType() bool {
	if o != nil && !IsNil(o.TpTriggerPxType) {
		return true
	}

	return false
}

// SetTpTriggerPxType gets a reference to the given string and assigns it to the TpTriggerPxType field.
func (o *CreateCopytradingAlgoOrderV5Req) SetTpTriggerPxType(v string) {
	o.TpTriggerPxType = &v
}

func (o CreateCopytradingAlgoOrderV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCopytradingAlgoOrderV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstType) {
		toSerialize["instType"] = o.InstType
	}
	if !IsNil(o.SlOrdPx) {
		toSerialize["slOrdPx"] = o.SlOrdPx
	}
	if !IsNil(o.SlTriggerPx) {
		toSerialize["slTriggerPx"] = o.SlTriggerPx
	}
	if !IsNil(o.SlTriggerPxType) {
		toSerialize["slTriggerPxType"] = o.SlTriggerPxType
	}
	toSerialize["subPosId"] = o.SubPosId
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.TpOrdPx) {
		toSerialize["tpOrdPx"] = o.TpOrdPx
	}
	if !IsNil(o.TpTriggerPx) {
		toSerialize["tpTriggerPx"] = o.TpTriggerPx
	}
	if !IsNil(o.TpTriggerPxType) {
		toSerialize["tpTriggerPxType"] = o.TpTriggerPxType
	}
	return toSerialize, nil
}

func (o *CreateCopytradingAlgoOrderV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subPosId",
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

	varCreateCopytradingAlgoOrderV5Req := _CreateCopytradingAlgoOrderV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateCopytradingAlgoOrderV5Req)

	if err != nil {
		return err
	}

	*o = CreateCopytradingAlgoOrderV5Req(varCreateCopytradingAlgoOrderV5Req)

	return err
}

type NullableCreateCopytradingAlgoOrderV5Req struct {
	value *CreateCopytradingAlgoOrderV5Req
	isSet bool
}

func (v NullableCreateCopytradingAlgoOrderV5Req) Get() *CreateCopytradingAlgoOrderV5Req {
	return v.value
}

func (v *NullableCreateCopytradingAlgoOrderV5Req) Set(val *CreateCopytradingAlgoOrderV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCopytradingAlgoOrderV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCopytradingAlgoOrderV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCopytradingAlgoOrderV5Req(val *CreateCopytradingAlgoOrderV5Req) *NullableCreateCopytradingAlgoOrderV5Req {
	return &NullableCreateCopytradingAlgoOrderV5Req{value: val, isSet: true}
}

func (v NullableCreateCopytradingAlgoOrderV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCopytradingAlgoOrderV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


