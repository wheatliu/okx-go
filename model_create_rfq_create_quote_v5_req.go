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

// checks if the CreateRfqCreateQuoteV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRfqCreateQuoteV5Req{}

// CreateRfqCreateQuoteV5Req struct for CreateRfqCreateQuoteV5Req
type CreateRfqCreateQuoteV5Req struct {
	// Submit Quote on a disclosed or anonymous basis.   Valid value is `true` or `false`. `false` by default.
	Anonymous *bool `json:"anonymous,omitempty"`
	// Client-supplied Quote ID.   A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters.
	ClQuoteId *string `json:"clQuoteId,omitempty"`
	// Seconds that a quote expires in.   Must be an integer between 10-120. Default is 60.
	ExpiresIn *string `json:"expiresIn,omitempty"`
	// The legs of the Quote.
	Legs []CreateRfqCreateQuoteV5ReqLegsInner `json:"legs"`
	// The trading direction of the Quote. Its value can be `buy` or `sell`.   For example, if quoteSide is `buy`, all the legs are executed in their leg sides; otherwise, all the legs are executed in the opposite of their leg sides.
	QuoteSide string `json:"quoteSide"`
	// RFQ ID .
	RfqId string `json:"rfqId"`
	// Quote tag.   The block trade associated with the Quote will have the same tag.   A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 16 characters.
	Tag *string `json:"tag,omitempty"`
}

type _CreateRfqCreateQuoteV5Req CreateRfqCreateQuoteV5Req

// NewCreateRfqCreateQuoteV5Req instantiates a new CreateRfqCreateQuoteV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRfqCreateQuoteV5Req(legs []CreateRfqCreateQuoteV5ReqLegsInner, quoteSide string, rfqId string) *CreateRfqCreateQuoteV5Req {
	this := CreateRfqCreateQuoteV5Req{}
	var clQuoteId string = ""
	this.ClQuoteId = &clQuoteId
	var expiresIn string = ""
	this.ExpiresIn = &expiresIn
	this.Legs = legs
	this.QuoteSide = quoteSide
	this.RfqId = rfqId
	var tag string = ""
	this.Tag = &tag
	return &this
}

// NewCreateRfqCreateQuoteV5ReqWithDefaults instantiates a new CreateRfqCreateQuoteV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRfqCreateQuoteV5ReqWithDefaults() *CreateRfqCreateQuoteV5Req {
	this := CreateRfqCreateQuoteV5Req{}
	var clQuoteId string = ""
	this.ClQuoteId = &clQuoteId
	var expiresIn string = ""
	this.ExpiresIn = &expiresIn
	var quoteSide string = ""
	this.QuoteSide = quoteSide
	var rfqId string = ""
	this.RfqId = rfqId
	var tag string = ""
	this.Tag = &tag
	return &this
}

// GetAnonymous returns the Anonymous field value if set, zero value otherwise.
func (o *CreateRfqCreateQuoteV5Req) GetAnonymous() bool {
	if o == nil || IsNil(o.Anonymous) {
		var ret bool
		return ret
	}
	return *o.Anonymous
}

// GetAnonymousOk returns a tuple with the Anonymous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateQuoteV5Req) GetAnonymousOk() (*bool, bool) {
	if o == nil || IsNil(o.Anonymous) {
		return nil, false
	}
	return o.Anonymous, true
}

// HasAnonymous returns a boolean if a field has been set.
func (o *CreateRfqCreateQuoteV5Req) HasAnonymous() bool {
	if o != nil && !IsNil(o.Anonymous) {
		return true
	}

	return false
}

// SetAnonymous gets a reference to the given bool and assigns it to the Anonymous field.
func (o *CreateRfqCreateQuoteV5Req) SetAnonymous(v bool) {
	o.Anonymous = &v
}

// GetClQuoteId returns the ClQuoteId field value if set, zero value otherwise.
func (o *CreateRfqCreateQuoteV5Req) GetClQuoteId() string {
	if o == nil || IsNil(o.ClQuoteId) {
		var ret string
		return ret
	}
	return *o.ClQuoteId
}

// GetClQuoteIdOk returns a tuple with the ClQuoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateQuoteV5Req) GetClQuoteIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClQuoteId) {
		return nil, false
	}
	return o.ClQuoteId, true
}

// HasClQuoteId returns a boolean if a field has been set.
func (o *CreateRfqCreateQuoteV5Req) HasClQuoteId() bool {
	if o != nil && !IsNil(o.ClQuoteId) {
		return true
	}

	return false
}

// SetClQuoteId gets a reference to the given string and assigns it to the ClQuoteId field.
func (o *CreateRfqCreateQuoteV5Req) SetClQuoteId(v string) {
	o.ClQuoteId = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *CreateRfqCreateQuoteV5Req) GetExpiresIn() string {
	if o == nil || IsNil(o.ExpiresIn) {
		var ret string
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateQuoteV5Req) GetExpiresInOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *CreateRfqCreateQuoteV5Req) HasExpiresIn() bool {
	if o != nil && !IsNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given string and assigns it to the ExpiresIn field.
func (o *CreateRfqCreateQuoteV5Req) SetExpiresIn(v string) {
	o.ExpiresIn = &v
}

// GetLegs returns the Legs field value
func (o *CreateRfqCreateQuoteV5Req) GetLegs() []CreateRfqCreateQuoteV5ReqLegsInner {
	if o == nil {
		var ret []CreateRfqCreateQuoteV5ReqLegsInner
		return ret
	}

	return o.Legs
}

// GetLegsOk returns a tuple with the Legs field value
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateQuoteV5Req) GetLegsOk() ([]CreateRfqCreateQuoteV5ReqLegsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Legs, true
}

// SetLegs sets field value
func (o *CreateRfqCreateQuoteV5Req) SetLegs(v []CreateRfqCreateQuoteV5ReqLegsInner) {
	o.Legs = v
}

// GetQuoteSide returns the QuoteSide field value
func (o *CreateRfqCreateQuoteV5Req) GetQuoteSide() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QuoteSide
}

// GetQuoteSideOk returns a tuple with the QuoteSide field value
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateQuoteV5Req) GetQuoteSideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuoteSide, true
}

// SetQuoteSide sets field value
func (o *CreateRfqCreateQuoteV5Req) SetQuoteSide(v string) {
	o.QuoteSide = v
}

// GetRfqId returns the RfqId field value
func (o *CreateRfqCreateQuoteV5Req) GetRfqId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RfqId
}

// GetRfqIdOk returns a tuple with the RfqId field value
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateQuoteV5Req) GetRfqIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RfqId, true
}

// SetRfqId sets field value
func (o *CreateRfqCreateQuoteV5Req) SetRfqId(v string) {
	o.RfqId = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *CreateRfqCreateQuoteV5Req) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateQuoteV5Req) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *CreateRfqCreateQuoteV5Req) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *CreateRfqCreateQuoteV5Req) SetTag(v string) {
	o.Tag = &v
}

func (o CreateRfqCreateQuoteV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRfqCreateQuoteV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Anonymous) {
		toSerialize["anonymous"] = o.Anonymous
	}
	if !IsNil(o.ClQuoteId) {
		toSerialize["clQuoteId"] = o.ClQuoteId
	}
	if !IsNil(o.ExpiresIn) {
		toSerialize["expiresIn"] = o.ExpiresIn
	}
	toSerialize["legs"] = o.Legs
	toSerialize["quoteSide"] = o.QuoteSide
	toSerialize["rfqId"] = o.RfqId
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return toSerialize, nil
}

func (o *CreateRfqCreateQuoteV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"legs",
		"quoteSide",
		"rfqId",
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

	varCreateRfqCreateQuoteV5Req := _CreateRfqCreateQuoteV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateRfqCreateQuoteV5Req)

	if err != nil {
		return err
	}

	*o = CreateRfqCreateQuoteV5Req(varCreateRfqCreateQuoteV5Req)

	return err
}

type NullableCreateRfqCreateQuoteV5Req struct {
	value *CreateRfqCreateQuoteV5Req
	isSet bool
}

func (v NullableCreateRfqCreateQuoteV5Req) Get() *CreateRfqCreateQuoteV5Req {
	return v.value
}

func (v *NullableCreateRfqCreateQuoteV5Req) Set(val *CreateRfqCreateQuoteV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRfqCreateQuoteV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRfqCreateQuoteV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRfqCreateQuoteV5Req(val *CreateRfqCreateQuoteV5Req) *NullableCreateRfqCreateQuoteV5Req {
	return &NullableCreateRfqCreateQuoteV5Req{value: val, isSet: true}
}

func (v NullableCreateRfqCreateQuoteV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRfqCreateQuoteV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


