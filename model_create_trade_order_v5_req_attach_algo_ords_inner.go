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

// checks if the CreateTradeOrderV5ReqAttachAlgoOrdsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradeOrderV5ReqAttachAlgoOrdsInner{}

// CreateTradeOrderV5ReqAttachAlgoOrdsInner struct for CreateTradeOrderV5ReqAttachAlgoOrdsInner
type CreateTradeOrderV5ReqAttachAlgoOrdsInner struct {
	// Whether to enable Cost-price SL. Only applicable to SL order of split TPs. Whether `slTriggerPx` will move to `avgPx` when the first TP order is triggered  `0`: disable, the default value   `1`: Enable
	AmendPxOnTriggerType *string `json:"amendPxOnTriggerType,omitempty"`
	// Client-supplied Algo ID when placing order attaching TP/SL  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters.  It will be posted to `algoClOrdId` when placing TP/SL order once the general order is filled completely.
	AttachAlgoClOrdId *string `json:"attachAlgoClOrdId,omitempty"`
	// Stop-loss order price  If you fill in this parameter, you should fill in the stop-loss trigger price.  If the price is -1, stop-loss will be executed at the market price.
	SlOrdPx *string `json:"slOrdPx,omitempty"`
	// Stop-loss trigger price  If you fill in this parameter, you should fill in the stop-loss order price.
	SlTriggerPx *string `json:"slTriggerPx,omitempty"`
	// Stop-loss trigger price type  `last`: last price   `index`: index price   `mark`: mark price   The default is last
	SlTriggerPxType *string `json:"slTriggerPxType,omitempty"`
	// Size. Only applicable to TP order of split TPs, and it is required for TP order of split TPs
	Sz *string `json:"sz,omitempty"`
	// TP order kind  `condition`  `limit`   The default is `condition`
	TpOrdKind *string `json:"tpOrdKind,omitempty"`
	// Take-profit order price     For condition TP order, if you fill in this parameter, you should fill in the take-profit trigger price as well.   For limit TP order, you need to fill in this parameter, take-profit trigger needn‘t to be filled.   If the price is -1, take-profit will be executed at the market price.
	TpOrdPx *string `json:"tpOrdPx,omitempty"`
	// Take-profit trigger price  For condition TP order, if you fill in this parameter, you should fill in the take-profit order price as well.
	TpTriggerPx *string `json:"tpTriggerPx,omitempty"`
	// Take-profit trigger price type  `last`: last price   `index`: index price   `mark`: mark price   The default is last
	TpTriggerPxType *string `json:"tpTriggerPxType,omitempty"`
}

// NewCreateTradeOrderV5ReqAttachAlgoOrdsInner instantiates a new CreateTradeOrderV5ReqAttachAlgoOrdsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradeOrderV5ReqAttachAlgoOrdsInner() *CreateTradeOrderV5ReqAttachAlgoOrdsInner {
	this := CreateTradeOrderV5ReqAttachAlgoOrdsInner{}
	var amendPxOnTriggerType string = ""
	this.AmendPxOnTriggerType = &amendPxOnTriggerType
	var attachAlgoClOrdId string = ""
	this.AttachAlgoClOrdId = &attachAlgoClOrdId
	var slOrdPx string = ""
	this.SlOrdPx = &slOrdPx
	var slTriggerPx string = ""
	this.SlTriggerPx = &slTriggerPx
	var slTriggerPxType string = ""
	this.SlTriggerPxType = &slTriggerPxType
	var sz string = ""
	this.Sz = &sz
	var tpOrdKind string = ""
	this.TpOrdKind = &tpOrdKind
	var tpOrdPx string = ""
	this.TpOrdPx = &tpOrdPx
	var tpTriggerPx string = ""
	this.TpTriggerPx = &tpTriggerPx
	var tpTriggerPxType string = ""
	this.TpTriggerPxType = &tpTriggerPxType
	return &this
}

// NewCreateTradeOrderV5ReqAttachAlgoOrdsInnerWithDefaults instantiates a new CreateTradeOrderV5ReqAttachAlgoOrdsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradeOrderV5ReqAttachAlgoOrdsInnerWithDefaults() *CreateTradeOrderV5ReqAttachAlgoOrdsInner {
	this := CreateTradeOrderV5ReqAttachAlgoOrdsInner{}
	var amendPxOnTriggerType string = ""
	this.AmendPxOnTriggerType = &amendPxOnTriggerType
	var attachAlgoClOrdId string = ""
	this.AttachAlgoClOrdId = &attachAlgoClOrdId
	var slOrdPx string = ""
	this.SlOrdPx = &slOrdPx
	var slTriggerPx string = ""
	this.SlTriggerPx = &slTriggerPx
	var slTriggerPxType string = ""
	this.SlTriggerPxType = &slTriggerPxType
	var sz string = ""
	this.Sz = &sz
	var tpOrdKind string = ""
	this.TpOrdKind = &tpOrdKind
	var tpOrdPx string = ""
	this.TpOrdPx = &tpOrdPx
	var tpTriggerPx string = ""
	this.TpTriggerPx = &tpTriggerPx
	var tpTriggerPxType string = ""
	this.TpTriggerPxType = &tpTriggerPxType
	return &this
}

// GetAmendPxOnTriggerType returns the AmendPxOnTriggerType field value if set, zero value otherwise.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetAmendPxOnTriggerType() string {
	if o == nil || IsNil(o.AmendPxOnTriggerType) {
		var ret string
		return ret
	}
	return *o.AmendPxOnTriggerType
}

// GetAmendPxOnTriggerTypeOk returns a tuple with the AmendPxOnTriggerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetAmendPxOnTriggerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AmendPxOnTriggerType) {
		return nil, false
	}
	return o.AmendPxOnTriggerType, true
}

// HasAmendPxOnTriggerType returns a boolean if a field has been set.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) HasAmendPxOnTriggerType() bool {
	if o != nil && !IsNil(o.AmendPxOnTriggerType) {
		return true
	}

	return false
}

// SetAmendPxOnTriggerType gets a reference to the given string and assigns it to the AmendPxOnTriggerType field.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) SetAmendPxOnTriggerType(v string) {
	o.AmendPxOnTriggerType = &v
}

// GetAttachAlgoClOrdId returns the AttachAlgoClOrdId field value if set, zero value otherwise.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetAttachAlgoClOrdId() string {
	if o == nil || IsNil(o.AttachAlgoClOrdId) {
		var ret string
		return ret
	}
	return *o.AttachAlgoClOrdId
}

// GetAttachAlgoClOrdIdOk returns a tuple with the AttachAlgoClOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetAttachAlgoClOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.AttachAlgoClOrdId) {
		return nil, false
	}
	return o.AttachAlgoClOrdId, true
}

// HasAttachAlgoClOrdId returns a boolean if a field has been set.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) HasAttachAlgoClOrdId() bool {
	if o != nil && !IsNil(o.AttachAlgoClOrdId) {
		return true
	}

	return false
}

// SetAttachAlgoClOrdId gets a reference to the given string and assigns it to the AttachAlgoClOrdId field.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) SetAttachAlgoClOrdId(v string) {
	o.AttachAlgoClOrdId = &v
}

// GetSlOrdPx returns the SlOrdPx field value if set, zero value otherwise.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetSlOrdPx() string {
	if o == nil || IsNil(o.SlOrdPx) {
		var ret string
		return ret
	}
	return *o.SlOrdPx
}

// GetSlOrdPxOk returns a tuple with the SlOrdPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetSlOrdPxOk() (*string, bool) {
	if o == nil || IsNil(o.SlOrdPx) {
		return nil, false
	}
	return o.SlOrdPx, true
}

// HasSlOrdPx returns a boolean if a field has been set.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) HasSlOrdPx() bool {
	if o != nil && !IsNil(o.SlOrdPx) {
		return true
	}

	return false
}

// SetSlOrdPx gets a reference to the given string and assigns it to the SlOrdPx field.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) SetSlOrdPx(v string) {
	o.SlOrdPx = &v
}

// GetSlTriggerPx returns the SlTriggerPx field value if set, zero value otherwise.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetSlTriggerPx() string {
	if o == nil || IsNil(o.SlTriggerPx) {
		var ret string
		return ret
	}
	return *o.SlTriggerPx
}

// GetSlTriggerPxOk returns a tuple with the SlTriggerPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetSlTriggerPxOk() (*string, bool) {
	if o == nil || IsNil(o.SlTriggerPx) {
		return nil, false
	}
	return o.SlTriggerPx, true
}

// HasSlTriggerPx returns a boolean if a field has been set.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) HasSlTriggerPx() bool {
	if o != nil && !IsNil(o.SlTriggerPx) {
		return true
	}

	return false
}

// SetSlTriggerPx gets a reference to the given string and assigns it to the SlTriggerPx field.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) SetSlTriggerPx(v string) {
	o.SlTriggerPx = &v
}

// GetSlTriggerPxType returns the SlTriggerPxType field value if set, zero value otherwise.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetSlTriggerPxType() string {
	if o == nil || IsNil(o.SlTriggerPxType) {
		var ret string
		return ret
	}
	return *o.SlTriggerPxType
}

// GetSlTriggerPxTypeOk returns a tuple with the SlTriggerPxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetSlTriggerPxTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SlTriggerPxType) {
		return nil, false
	}
	return o.SlTriggerPxType, true
}

// HasSlTriggerPxType returns a boolean if a field has been set.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) HasSlTriggerPxType() bool {
	if o != nil && !IsNil(o.SlTriggerPxType) {
		return true
	}

	return false
}

// SetSlTriggerPxType gets a reference to the given string and assigns it to the SlTriggerPxType field.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) SetSlTriggerPxType(v string) {
	o.SlTriggerPxType = &v
}

// GetSz returns the Sz field value if set, zero value otherwise.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetSz() string {
	if o == nil || IsNil(o.Sz) {
		var ret string
		return ret
	}
	return *o.Sz
}

// GetSzOk returns a tuple with the Sz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetSzOk() (*string, bool) {
	if o == nil || IsNil(o.Sz) {
		return nil, false
	}
	return o.Sz, true
}

// HasSz returns a boolean if a field has been set.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) HasSz() bool {
	if o != nil && !IsNil(o.Sz) {
		return true
	}

	return false
}

// SetSz gets a reference to the given string and assigns it to the Sz field.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) SetSz(v string) {
	o.Sz = &v
}

// GetTpOrdKind returns the TpOrdKind field value if set, zero value otherwise.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetTpOrdKind() string {
	if o == nil || IsNil(o.TpOrdKind) {
		var ret string
		return ret
	}
	return *o.TpOrdKind
}

// GetTpOrdKindOk returns a tuple with the TpOrdKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetTpOrdKindOk() (*string, bool) {
	if o == nil || IsNil(o.TpOrdKind) {
		return nil, false
	}
	return o.TpOrdKind, true
}

// HasTpOrdKind returns a boolean if a field has been set.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) HasTpOrdKind() bool {
	if o != nil && !IsNil(o.TpOrdKind) {
		return true
	}

	return false
}

// SetTpOrdKind gets a reference to the given string and assigns it to the TpOrdKind field.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) SetTpOrdKind(v string) {
	o.TpOrdKind = &v
}

// GetTpOrdPx returns the TpOrdPx field value if set, zero value otherwise.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetTpOrdPx() string {
	if o == nil || IsNil(o.TpOrdPx) {
		var ret string
		return ret
	}
	return *o.TpOrdPx
}

// GetTpOrdPxOk returns a tuple with the TpOrdPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetTpOrdPxOk() (*string, bool) {
	if o == nil || IsNil(o.TpOrdPx) {
		return nil, false
	}
	return o.TpOrdPx, true
}

// HasTpOrdPx returns a boolean if a field has been set.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) HasTpOrdPx() bool {
	if o != nil && !IsNil(o.TpOrdPx) {
		return true
	}

	return false
}

// SetTpOrdPx gets a reference to the given string and assigns it to the TpOrdPx field.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) SetTpOrdPx(v string) {
	o.TpOrdPx = &v
}

// GetTpTriggerPx returns the TpTriggerPx field value if set, zero value otherwise.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetTpTriggerPx() string {
	if o == nil || IsNil(o.TpTriggerPx) {
		var ret string
		return ret
	}
	return *o.TpTriggerPx
}

// GetTpTriggerPxOk returns a tuple with the TpTriggerPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetTpTriggerPxOk() (*string, bool) {
	if o == nil || IsNil(o.TpTriggerPx) {
		return nil, false
	}
	return o.TpTriggerPx, true
}

// HasTpTriggerPx returns a boolean if a field has been set.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) HasTpTriggerPx() bool {
	if o != nil && !IsNil(o.TpTriggerPx) {
		return true
	}

	return false
}

// SetTpTriggerPx gets a reference to the given string and assigns it to the TpTriggerPx field.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) SetTpTriggerPx(v string) {
	o.TpTriggerPx = &v
}

// GetTpTriggerPxType returns the TpTriggerPxType field value if set, zero value otherwise.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetTpTriggerPxType() string {
	if o == nil || IsNil(o.TpTriggerPxType) {
		var ret string
		return ret
	}
	return *o.TpTriggerPxType
}

// GetTpTriggerPxTypeOk returns a tuple with the TpTriggerPxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetTpTriggerPxTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TpTriggerPxType) {
		return nil, false
	}
	return o.TpTriggerPxType, true
}

// HasTpTriggerPxType returns a boolean if a field has been set.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) HasTpTriggerPxType() bool {
	if o != nil && !IsNil(o.TpTriggerPxType) {
		return true
	}

	return false
}

// SetTpTriggerPxType gets a reference to the given string and assigns it to the TpTriggerPxType field.
func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) SetTpTriggerPxType(v string) {
	o.TpTriggerPxType = &v
}

func (o CreateTradeOrderV5ReqAttachAlgoOrdsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradeOrderV5ReqAttachAlgoOrdsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AmendPxOnTriggerType) {
		toSerialize["amendPxOnTriggerType"] = o.AmendPxOnTriggerType
	}
	if !IsNil(o.AttachAlgoClOrdId) {
		toSerialize["attachAlgoClOrdId"] = o.AttachAlgoClOrdId
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
	if !IsNil(o.Sz) {
		toSerialize["sz"] = o.Sz
	}
	if !IsNil(o.TpOrdKind) {
		toSerialize["tpOrdKind"] = o.TpOrdKind
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

type NullableCreateTradeOrderV5ReqAttachAlgoOrdsInner struct {
	value *CreateTradeOrderV5ReqAttachAlgoOrdsInner
	isSet bool
}

func (v NullableCreateTradeOrderV5ReqAttachAlgoOrdsInner) Get() *CreateTradeOrderV5ReqAttachAlgoOrdsInner {
	return v.value
}

func (v *NullableCreateTradeOrderV5ReqAttachAlgoOrdsInner) Set(val *CreateTradeOrderV5ReqAttachAlgoOrdsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradeOrderV5ReqAttachAlgoOrdsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradeOrderV5ReqAttachAlgoOrdsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradeOrderV5ReqAttachAlgoOrdsInner(val *CreateTradeOrderV5ReqAttachAlgoOrdsInner) *NullableCreateTradeOrderV5ReqAttachAlgoOrdsInner {
	return &NullableCreateTradeOrderV5ReqAttachAlgoOrdsInner{value: val, isSet: true}
}

func (v NullableCreateTradeOrderV5ReqAttachAlgoOrdsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradeOrderV5ReqAttachAlgoOrdsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


