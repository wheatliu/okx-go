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

// checks if the CreateAssetConvertEstimateQuoteV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAssetConvertEstimateQuoteV5Req{}

// CreateAssetConvertEstimateQuoteV5Req struct for CreateAssetConvertEstimateQuoteV5Req
type CreateAssetConvertEstimateQuoteV5Req struct {
	// Base currency, e.g. `BTC` in `BTC-USDT`
	BaseCcy string `json:"baseCcy"`
	// Client Order ID as assigned by the client  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters.
	ClQReqId *string `json:"clQReqId,omitempty"`
	// Quote currency, e.g. `USDT` in `BTC-USDT`
	QuoteCcy string `json:"quoteCcy"`
	// RFQ amount
	RfqSz string `json:"rfqSz"`
	// RFQ currency
	RfqSzCcy string `json:"rfqSzCcy"`
	// Trade side based on `baseCcy`  `buy` `sell`
	Side string `json:"side"`
	// Order tag  Applicable to broker user
	Tag *string `json:"tag,omitempty"`
}

type _CreateAssetConvertEstimateQuoteV5Req CreateAssetConvertEstimateQuoteV5Req

// NewCreateAssetConvertEstimateQuoteV5Req instantiates a new CreateAssetConvertEstimateQuoteV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAssetConvertEstimateQuoteV5Req(baseCcy string, quoteCcy string, rfqSz string, rfqSzCcy string, side string) *CreateAssetConvertEstimateQuoteV5Req {
	this := CreateAssetConvertEstimateQuoteV5Req{}
	this.BaseCcy = baseCcy
	var clQReqId string = ""
	this.ClQReqId = &clQReqId
	this.QuoteCcy = quoteCcy
	this.RfqSz = rfqSz
	this.RfqSzCcy = rfqSzCcy
	this.Side = side
	var tag string = ""
	this.Tag = &tag
	return &this
}

// NewCreateAssetConvertEstimateQuoteV5ReqWithDefaults instantiates a new CreateAssetConvertEstimateQuoteV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAssetConvertEstimateQuoteV5ReqWithDefaults() *CreateAssetConvertEstimateQuoteV5Req {
	this := CreateAssetConvertEstimateQuoteV5Req{}
	var baseCcy string = ""
	this.BaseCcy = baseCcy
	var clQReqId string = ""
	this.ClQReqId = &clQReqId
	var quoteCcy string = ""
	this.QuoteCcy = quoteCcy
	var rfqSz string = ""
	this.RfqSz = rfqSz
	var rfqSzCcy string = ""
	this.RfqSzCcy = rfqSzCcy
	var side string = ""
	this.Side = side
	var tag string = ""
	this.Tag = &tag
	return &this
}

// GetBaseCcy returns the BaseCcy field value
func (o *CreateAssetConvertEstimateQuoteV5Req) GetBaseCcy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseCcy
}

// GetBaseCcyOk returns a tuple with the BaseCcy field value
// and a boolean to check if the value has been set.
func (o *CreateAssetConvertEstimateQuoteV5Req) GetBaseCcyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseCcy, true
}

// SetBaseCcy sets field value
func (o *CreateAssetConvertEstimateQuoteV5Req) SetBaseCcy(v string) {
	o.BaseCcy = v
}

// GetClQReqId returns the ClQReqId field value if set, zero value otherwise.
func (o *CreateAssetConvertEstimateQuoteV5Req) GetClQReqId() string {
	if o == nil || IsNil(o.ClQReqId) {
		var ret string
		return ret
	}
	return *o.ClQReqId
}

// GetClQReqIdOk returns a tuple with the ClQReqId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetConvertEstimateQuoteV5Req) GetClQReqIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClQReqId) {
		return nil, false
	}
	return o.ClQReqId, true
}

// HasClQReqId returns a boolean if a field has been set.
func (o *CreateAssetConvertEstimateQuoteV5Req) HasClQReqId() bool {
	if o != nil && !IsNil(o.ClQReqId) {
		return true
	}

	return false
}

// SetClQReqId gets a reference to the given string and assigns it to the ClQReqId field.
func (o *CreateAssetConvertEstimateQuoteV5Req) SetClQReqId(v string) {
	o.ClQReqId = &v
}

// GetQuoteCcy returns the QuoteCcy field value
func (o *CreateAssetConvertEstimateQuoteV5Req) GetQuoteCcy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QuoteCcy
}

// GetQuoteCcyOk returns a tuple with the QuoteCcy field value
// and a boolean to check if the value has been set.
func (o *CreateAssetConvertEstimateQuoteV5Req) GetQuoteCcyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuoteCcy, true
}

// SetQuoteCcy sets field value
func (o *CreateAssetConvertEstimateQuoteV5Req) SetQuoteCcy(v string) {
	o.QuoteCcy = v
}

// GetRfqSz returns the RfqSz field value
func (o *CreateAssetConvertEstimateQuoteV5Req) GetRfqSz() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RfqSz
}

// GetRfqSzOk returns a tuple with the RfqSz field value
// and a boolean to check if the value has been set.
func (o *CreateAssetConvertEstimateQuoteV5Req) GetRfqSzOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RfqSz, true
}

// SetRfqSz sets field value
func (o *CreateAssetConvertEstimateQuoteV5Req) SetRfqSz(v string) {
	o.RfqSz = v
}

// GetRfqSzCcy returns the RfqSzCcy field value
func (o *CreateAssetConvertEstimateQuoteV5Req) GetRfqSzCcy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RfqSzCcy
}

// GetRfqSzCcyOk returns a tuple with the RfqSzCcy field value
// and a boolean to check if the value has been set.
func (o *CreateAssetConvertEstimateQuoteV5Req) GetRfqSzCcyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RfqSzCcy, true
}

// SetRfqSzCcy sets field value
func (o *CreateAssetConvertEstimateQuoteV5Req) SetRfqSzCcy(v string) {
	o.RfqSzCcy = v
}

// GetSide returns the Side field value
func (o *CreateAssetConvertEstimateQuoteV5Req) GetSide() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Side
}

// GetSideOk returns a tuple with the Side field value
// and a boolean to check if the value has been set.
func (o *CreateAssetConvertEstimateQuoteV5Req) GetSideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Side, true
}

// SetSide sets field value
func (o *CreateAssetConvertEstimateQuoteV5Req) SetSide(v string) {
	o.Side = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *CreateAssetConvertEstimateQuoteV5Req) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetConvertEstimateQuoteV5Req) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *CreateAssetConvertEstimateQuoteV5Req) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *CreateAssetConvertEstimateQuoteV5Req) SetTag(v string) {
	o.Tag = &v
}

func (o CreateAssetConvertEstimateQuoteV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAssetConvertEstimateQuoteV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["baseCcy"] = o.BaseCcy
	if !IsNil(o.ClQReqId) {
		toSerialize["clQReqId"] = o.ClQReqId
	}
	toSerialize["quoteCcy"] = o.QuoteCcy
	toSerialize["rfqSz"] = o.RfqSz
	toSerialize["rfqSzCcy"] = o.RfqSzCcy
	toSerialize["side"] = o.Side
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return toSerialize, nil
}

func (o *CreateAssetConvertEstimateQuoteV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"baseCcy",
		"quoteCcy",
		"rfqSz",
		"rfqSzCcy",
		"side",
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

	varCreateAssetConvertEstimateQuoteV5Req := _CreateAssetConvertEstimateQuoteV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateAssetConvertEstimateQuoteV5Req)

	if err != nil {
		return err
	}

	*o = CreateAssetConvertEstimateQuoteV5Req(varCreateAssetConvertEstimateQuoteV5Req)

	return err
}

type NullableCreateAssetConvertEstimateQuoteV5Req struct {
	value *CreateAssetConvertEstimateQuoteV5Req
	isSet bool
}

func (v NullableCreateAssetConvertEstimateQuoteV5Req) Get() *CreateAssetConvertEstimateQuoteV5Req {
	return v.value
}

func (v *NullableCreateAssetConvertEstimateQuoteV5Req) Set(val *CreateAssetConvertEstimateQuoteV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAssetConvertEstimateQuoteV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAssetConvertEstimateQuoteV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAssetConvertEstimateQuoteV5Req(val *CreateAssetConvertEstimateQuoteV5Req) *NullableCreateAssetConvertEstimateQuoteV5Req {
	return &NullableCreateAssetConvertEstimateQuoteV5Req{value: val, isSet: true}
}

func (v NullableCreateAssetConvertEstimateQuoteV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAssetConvertEstimateQuoteV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


