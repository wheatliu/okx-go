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

// checks if the CreateFiatCreateWithdrawalV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFiatCreateWithdrawalV5RespDataInner{}

// CreateFiatCreateWithdrawalV5RespDataInner struct for CreateFiatCreateWithdrawalV5RespDataInner
type CreateFiatCreateWithdrawalV5RespDataInner struct {
	// The requested amount for the transaction
	Amt *string `json:"amt,omitempty"`
	// The creation time of the transaction, Unix timestamp format in milliseconds, e.g. `1597026383085`
	CTime *string `json:"cTime,omitempty"`
	// The currency of the transaction
	Ccy *string `json:"ccy,omitempty"`
	// The client ID associated with the transaction
	ClientId *string `json:"clientId,omitempty"`
	// The transaction fee
	Fee *string `json:"fee,omitempty"`
	// The unique order Id
	OrdId *string `json:"ordId,omitempty"`
	// The Id of the payment account used
	PaymentAcctId *string `json:"paymentAcctId,omitempty"`
	// Payment Method  `TR_BANKS`  `PIX`  `SEPA`
	PaymentMethod *string `json:"paymentMethod,omitempty"`
	// The State of the transaction  `processing`  `completed`
	State *string `json:"state,omitempty"`
	// The update time of the transaction, Unix timestamp format in milliseconds, e.g. `1597026383085`
	UTime *string `json:"uTime,omitempty"`
}

// NewCreateFiatCreateWithdrawalV5RespDataInner instantiates a new CreateFiatCreateWithdrawalV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFiatCreateWithdrawalV5RespDataInner() *CreateFiatCreateWithdrawalV5RespDataInner {
	this := CreateFiatCreateWithdrawalV5RespDataInner{}
	var amt string = ""
	this.Amt = &amt
	var cTime string = ""
	this.CTime = &cTime
	var ccy string = ""
	this.Ccy = &ccy
	var clientId string = ""
	this.ClientId = &clientId
	var fee string = ""
	this.Fee = &fee
	var ordId string = ""
	this.OrdId = &ordId
	var paymentAcctId string = ""
	this.PaymentAcctId = &paymentAcctId
	var paymentMethod string = ""
	this.PaymentMethod = &paymentMethod
	var state string = ""
	this.State = &state
	var uTime string = ""
	this.UTime = &uTime
	return &this
}

// NewCreateFiatCreateWithdrawalV5RespDataInnerWithDefaults instantiates a new CreateFiatCreateWithdrawalV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFiatCreateWithdrawalV5RespDataInnerWithDefaults() *CreateFiatCreateWithdrawalV5RespDataInner {
	this := CreateFiatCreateWithdrawalV5RespDataInner{}
	var amt string = ""
	this.Amt = &amt
	var cTime string = ""
	this.CTime = &cTime
	var ccy string = ""
	this.Ccy = &ccy
	var clientId string = ""
	this.ClientId = &clientId
	var fee string = ""
	this.Fee = &fee
	var ordId string = ""
	this.OrdId = &ordId
	var paymentAcctId string = ""
	this.PaymentAcctId = &paymentAcctId
	var paymentMethod string = ""
	this.PaymentMethod = &paymentMethod
	var state string = ""
	this.State = &state
	var uTime string = ""
	this.UTime = &uTime
	return &this
}

// GetAmt returns the Amt field value if set, zero value otherwise.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) GetAmt() string {
	if o == nil || IsNil(o.Amt) {
		var ret string
		return ret
	}
	return *o.Amt
}

// GetAmtOk returns a tuple with the Amt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) GetAmtOk() (*string, bool) {
	if o == nil || IsNil(o.Amt) {
		return nil, false
	}
	return o.Amt, true
}

// HasAmt returns a boolean if a field has been set.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) HasAmt() bool {
	if o != nil && !IsNil(o.Amt) {
		return true
	}

	return false
}

// SetAmt gets a reference to the given string and assigns it to the Amt field.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) SetAmt(v string) {
	o.Amt = &v
}

// GetCTime returns the CTime field value if set, zero value otherwise.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) GetCTime() string {
	if o == nil || IsNil(o.CTime) {
		var ret string
		return ret
	}
	return *o.CTime
}

// GetCTimeOk returns a tuple with the CTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) GetCTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CTime) {
		return nil, false
	}
	return o.CTime, true
}

// HasCTime returns a boolean if a field has been set.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) HasCTime() bool {
	if o != nil && !IsNil(o.CTime) {
		return true
	}

	return false
}

// SetCTime gets a reference to the given string and assigns it to the CTime field.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) SetCTime(v string) {
	o.CTime = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) SetClientId(v string) {
	o.ClientId = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) GetFee() string {
	if o == nil || IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) GetFeeOk() (*string, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) SetFee(v string) {
	o.Fee = &v
}

// GetOrdId returns the OrdId field value if set, zero value otherwise.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) GetOrdId() string {
	if o == nil || IsNil(o.OrdId) {
		var ret string
		return ret
	}
	return *o.OrdId
}

// GetOrdIdOk returns a tuple with the OrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) GetOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrdId) {
		return nil, false
	}
	return o.OrdId, true
}

// HasOrdId returns a boolean if a field has been set.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) HasOrdId() bool {
	if o != nil && !IsNil(o.OrdId) {
		return true
	}

	return false
}

// SetOrdId gets a reference to the given string and assigns it to the OrdId field.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) SetOrdId(v string) {
	o.OrdId = &v
}

// GetPaymentAcctId returns the PaymentAcctId field value if set, zero value otherwise.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) GetPaymentAcctId() string {
	if o == nil || IsNil(o.PaymentAcctId) {
		var ret string
		return ret
	}
	return *o.PaymentAcctId
}

// GetPaymentAcctIdOk returns a tuple with the PaymentAcctId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) GetPaymentAcctIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentAcctId) {
		return nil, false
	}
	return o.PaymentAcctId, true
}

// HasPaymentAcctId returns a boolean if a field has been set.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) HasPaymentAcctId() bool {
	if o != nil && !IsNil(o.PaymentAcctId) {
		return true
	}

	return false
}

// SetPaymentAcctId gets a reference to the given string and assigns it to the PaymentAcctId field.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) SetPaymentAcctId(v string) {
	o.PaymentAcctId = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) GetPaymentMethod() string {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret string
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) GetPaymentMethodOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given string and assigns it to the PaymentMethod field.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) SetPaymentMethod(v string) {
	o.PaymentMethod = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) SetState(v string) {
	o.State = &v
}

// GetUTime returns the UTime field value if set, zero value otherwise.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) GetUTime() string {
	if o == nil || IsNil(o.UTime) {
		var ret string
		return ret
	}
	return *o.UTime
}

// GetUTimeOk returns a tuple with the UTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) GetUTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UTime) {
		return nil, false
	}
	return o.UTime, true
}

// HasUTime returns a boolean if a field has been set.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) HasUTime() bool {
	if o != nil && !IsNil(o.UTime) {
		return true
	}

	return false
}

// SetUTime gets a reference to the given string and assigns it to the UTime field.
func (o *CreateFiatCreateWithdrawalV5RespDataInner) SetUTime(v string) {
	o.UTime = &v
}

func (o CreateFiatCreateWithdrawalV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFiatCreateWithdrawalV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amt) {
		toSerialize["amt"] = o.Amt
	}
	if !IsNil(o.CTime) {
		toSerialize["cTime"] = o.CTime
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.OrdId) {
		toSerialize["ordId"] = o.OrdId
	}
	if !IsNil(o.PaymentAcctId) {
		toSerialize["paymentAcctId"] = o.PaymentAcctId
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.UTime) {
		toSerialize["uTime"] = o.UTime
	}
	return toSerialize, nil
}

type NullableCreateFiatCreateWithdrawalV5RespDataInner struct {
	value *CreateFiatCreateWithdrawalV5RespDataInner
	isSet bool
}

func (v NullableCreateFiatCreateWithdrawalV5RespDataInner) Get() *CreateFiatCreateWithdrawalV5RespDataInner {
	return v.value
}

func (v *NullableCreateFiatCreateWithdrawalV5RespDataInner) Set(val *CreateFiatCreateWithdrawalV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFiatCreateWithdrawalV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFiatCreateWithdrawalV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFiatCreateWithdrawalV5RespDataInner(val *CreateFiatCreateWithdrawalV5RespDataInner) *NullableCreateFiatCreateWithdrawalV5RespDataInner {
	return &NullableCreateFiatCreateWithdrawalV5RespDataInner{value: val, isSet: true}
}

func (v NullableCreateFiatCreateWithdrawalV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFiatCreateWithdrawalV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


