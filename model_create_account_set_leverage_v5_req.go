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

// checks if the CreateAccountSetLeverageV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAccountSetLeverageV5Req{}

// CreateAccountSetLeverageV5Req struct for CreateAccountSetLeverageV5Req
type CreateAccountSetLeverageV5Req struct {
	// Currency used for margin, used for the leverage setting for the currency in auto borrow.  Only applicable to `cross` `MARGIN` of `Spot mode`/`Multi-currency margin`/`Portfolio margin`  Required when setting the leverage for automatically borrowing coin.
	Ccy *string `json:"ccy,omitempty"`
	// Instrument ID  Under cross mode, either `instId` or `ccy` is required; if both are passed, `instId` will be used by default.
	InstId *string `json:"instId,omitempty"`
	// Leverage
	Lever string `json:"lever"`
	// Margin mode  `isolated` `cross`   Can only be `cross` if `ccy` is passed.
	MgnMode string `json:"mgnMode"`
	// Position side  `long` `short`  Only required when margin mode is `isolated` in `long/short` mode for `FUTURES`/`SWAP`.
	PosSide *string `json:"posSide,omitempty"`
}

type _CreateAccountSetLeverageV5Req CreateAccountSetLeverageV5Req

// NewCreateAccountSetLeverageV5Req instantiates a new CreateAccountSetLeverageV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccountSetLeverageV5Req(lever string, mgnMode string) *CreateAccountSetLeverageV5Req {
	this := CreateAccountSetLeverageV5Req{}
	var ccy string = ""
	this.Ccy = &ccy
	var instId string = ""
	this.InstId = &instId
	this.Lever = lever
	this.MgnMode = mgnMode
	var posSide string = ""
	this.PosSide = &posSide
	return &this
}

// NewCreateAccountSetLeverageV5ReqWithDefaults instantiates a new CreateAccountSetLeverageV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountSetLeverageV5ReqWithDefaults() *CreateAccountSetLeverageV5Req {
	this := CreateAccountSetLeverageV5Req{}
	var ccy string = ""
	this.Ccy = &ccy
	var instId string = ""
	this.InstId = &instId
	var lever string = ""
	this.Lever = lever
	var mgnMode string = ""
	this.MgnMode = mgnMode
	var posSide string = ""
	this.PosSide = &posSide
	return &this
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *CreateAccountSetLeverageV5Req) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountSetLeverageV5Req) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *CreateAccountSetLeverageV5Req) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *CreateAccountSetLeverageV5Req) SetCcy(v string) {
	o.Ccy = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *CreateAccountSetLeverageV5Req) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountSetLeverageV5Req) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *CreateAccountSetLeverageV5Req) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *CreateAccountSetLeverageV5Req) SetInstId(v string) {
	o.InstId = &v
}

// GetLever returns the Lever field value
func (o *CreateAccountSetLeverageV5Req) GetLever() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Lever
}

// GetLeverOk returns a tuple with the Lever field value
// and a boolean to check if the value has been set.
func (o *CreateAccountSetLeverageV5Req) GetLeverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lever, true
}

// SetLever sets field value
func (o *CreateAccountSetLeverageV5Req) SetLever(v string) {
	o.Lever = v
}

// GetMgnMode returns the MgnMode field value
func (o *CreateAccountSetLeverageV5Req) GetMgnMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MgnMode
}

// GetMgnModeOk returns a tuple with the MgnMode field value
// and a boolean to check if the value has been set.
func (o *CreateAccountSetLeverageV5Req) GetMgnModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MgnMode, true
}

// SetMgnMode sets field value
func (o *CreateAccountSetLeverageV5Req) SetMgnMode(v string) {
	o.MgnMode = v
}

// GetPosSide returns the PosSide field value if set, zero value otherwise.
func (o *CreateAccountSetLeverageV5Req) GetPosSide() string {
	if o == nil || IsNil(o.PosSide) {
		var ret string
		return ret
	}
	return *o.PosSide
}

// GetPosSideOk returns a tuple with the PosSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountSetLeverageV5Req) GetPosSideOk() (*string, bool) {
	if o == nil || IsNil(o.PosSide) {
		return nil, false
	}
	return o.PosSide, true
}

// HasPosSide returns a boolean if a field has been set.
func (o *CreateAccountSetLeverageV5Req) HasPosSide() bool {
	if o != nil && !IsNil(o.PosSide) {
		return true
	}

	return false
}

// SetPosSide gets a reference to the given string and assigns it to the PosSide field.
func (o *CreateAccountSetLeverageV5Req) SetPosSide(v string) {
	o.PosSide = &v
}

func (o CreateAccountSetLeverageV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAccountSetLeverageV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	toSerialize["lever"] = o.Lever
	toSerialize["mgnMode"] = o.MgnMode
	if !IsNil(o.PosSide) {
		toSerialize["posSide"] = o.PosSide
	}
	return toSerialize, nil
}

func (o *CreateAccountSetLeverageV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lever",
		"mgnMode",
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

	varCreateAccountSetLeverageV5Req := _CreateAccountSetLeverageV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateAccountSetLeverageV5Req)

	if err != nil {
		return err
	}

	*o = CreateAccountSetLeverageV5Req(varCreateAccountSetLeverageV5Req)

	return err
}

type NullableCreateAccountSetLeverageV5Req struct {
	value *CreateAccountSetLeverageV5Req
	isSet bool
}

func (v NullableCreateAccountSetLeverageV5Req) Get() *CreateAccountSetLeverageV5Req {
	return v.value
}

func (v *NullableCreateAccountSetLeverageV5Req) Set(val *CreateAccountSetLeverageV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccountSetLeverageV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccountSetLeverageV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccountSetLeverageV5Req(val *CreateAccountSetLeverageV5Req) *NullableCreateAccountSetLeverageV5Req {
	return &NullableCreateAccountSetLeverageV5Req{value: val, isSet: true}
}

func (v NullableCreateAccountSetLeverageV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccountSetLeverageV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


