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

// checks if the GetFiatDepositV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFiatDepositV5RespData{}

// GetFiatDepositV5RespData struct for GetFiatDepositV5RespData
type GetFiatDepositV5RespData struct {
	// Amount of the transaction
	Amt *string `json:"amt,omitempty"`
	// The creation time of the transaction, Unix timestamp format in milliseconds, e.g. `1597026383085`
	CTime *string `json:"cTime,omitempty"`
	// The currency of the transaction
	Ccy *string `json:"ccy,omitempty"`
	// The original request ID associated with the transaction. If it's a deposit, it's most likely an empty string (\"\").
	ClientId *string `json:"clientId,omitempty"`
	// The transaction fee
	Fee *string `json:"fee,omitempty"`
	// Order ID
	OrdId *string `json:"ordId,omitempty"`
	// The ID of the payment account used
	PaymentAcctId *string `json:"paymentAcctId,omitempty"`
	// Payment method, e.g.`TR_BANKS`
	PaymentMethod *string `json:"paymentMethod,omitempty"`
	// The state of the transaction  `completed`  `failed`  `pending`  `canceled`  `inqueue`  `processing`
	State *string `json:"state,omitempty"`
	// The update time of the transaction, Unix timestamp format in milliseconds, e.g. `1597026383085`
	UTime *string `json:"uTime,omitempty"`
}

// NewGetFiatDepositV5RespData instantiates a new GetFiatDepositV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFiatDepositV5RespData() *GetFiatDepositV5RespData {
	this := GetFiatDepositV5RespData{}
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

// NewGetFiatDepositV5RespDataWithDefaults instantiates a new GetFiatDepositV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFiatDepositV5RespDataWithDefaults() *GetFiatDepositV5RespData {
	this := GetFiatDepositV5RespData{}
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
func (o *GetFiatDepositV5RespData) GetAmt() string {
	if o == nil || IsNil(o.Amt) {
		var ret string
		return ret
	}
	return *o.Amt
}

// GetAmtOk returns a tuple with the Amt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositV5RespData) GetAmtOk() (*string, bool) {
	if o == nil || IsNil(o.Amt) {
		return nil, false
	}
	return o.Amt, true
}

// HasAmt returns a boolean if a field has been set.
func (o *GetFiatDepositV5RespData) HasAmt() bool {
	if o != nil && !IsNil(o.Amt) {
		return true
	}

	return false
}

// SetAmt gets a reference to the given string and assigns it to the Amt field.
func (o *GetFiatDepositV5RespData) SetAmt(v string) {
	o.Amt = &v
}

// GetCTime returns the CTime field value if set, zero value otherwise.
func (o *GetFiatDepositV5RespData) GetCTime() string {
	if o == nil || IsNil(o.CTime) {
		var ret string
		return ret
	}
	return *o.CTime
}

// GetCTimeOk returns a tuple with the CTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositV5RespData) GetCTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CTime) {
		return nil, false
	}
	return o.CTime, true
}

// HasCTime returns a boolean if a field has been set.
func (o *GetFiatDepositV5RespData) HasCTime() bool {
	if o != nil && !IsNil(o.CTime) {
		return true
	}

	return false
}

// SetCTime gets a reference to the given string and assigns it to the CTime field.
func (o *GetFiatDepositV5RespData) SetCTime(v string) {
	o.CTime = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetFiatDepositV5RespData) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositV5RespData) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetFiatDepositV5RespData) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetFiatDepositV5RespData) SetCcy(v string) {
	o.Ccy = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *GetFiatDepositV5RespData) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositV5RespData) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *GetFiatDepositV5RespData) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *GetFiatDepositV5RespData) SetClientId(v string) {
	o.ClientId = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *GetFiatDepositV5RespData) GetFee() string {
	if o == nil || IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositV5RespData) GetFeeOk() (*string, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *GetFiatDepositV5RespData) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *GetFiatDepositV5RespData) SetFee(v string) {
	o.Fee = &v
}

// GetOrdId returns the OrdId field value if set, zero value otherwise.
func (o *GetFiatDepositV5RespData) GetOrdId() string {
	if o == nil || IsNil(o.OrdId) {
		var ret string
		return ret
	}
	return *o.OrdId
}

// GetOrdIdOk returns a tuple with the OrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositV5RespData) GetOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrdId) {
		return nil, false
	}
	return o.OrdId, true
}

// HasOrdId returns a boolean if a field has been set.
func (o *GetFiatDepositV5RespData) HasOrdId() bool {
	if o != nil && !IsNil(o.OrdId) {
		return true
	}

	return false
}

// SetOrdId gets a reference to the given string and assigns it to the OrdId field.
func (o *GetFiatDepositV5RespData) SetOrdId(v string) {
	o.OrdId = &v
}

// GetPaymentAcctId returns the PaymentAcctId field value if set, zero value otherwise.
func (o *GetFiatDepositV5RespData) GetPaymentAcctId() string {
	if o == nil || IsNil(o.PaymentAcctId) {
		var ret string
		return ret
	}
	return *o.PaymentAcctId
}

// GetPaymentAcctIdOk returns a tuple with the PaymentAcctId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositV5RespData) GetPaymentAcctIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentAcctId) {
		return nil, false
	}
	return o.PaymentAcctId, true
}

// HasPaymentAcctId returns a boolean if a field has been set.
func (o *GetFiatDepositV5RespData) HasPaymentAcctId() bool {
	if o != nil && !IsNil(o.PaymentAcctId) {
		return true
	}

	return false
}

// SetPaymentAcctId gets a reference to the given string and assigns it to the PaymentAcctId field.
func (o *GetFiatDepositV5RespData) SetPaymentAcctId(v string) {
	o.PaymentAcctId = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *GetFiatDepositV5RespData) GetPaymentMethod() string {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret string
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositV5RespData) GetPaymentMethodOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *GetFiatDepositV5RespData) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given string and assigns it to the PaymentMethod field.
func (o *GetFiatDepositV5RespData) SetPaymentMethod(v string) {
	o.PaymentMethod = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GetFiatDepositV5RespData) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositV5RespData) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GetFiatDepositV5RespData) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GetFiatDepositV5RespData) SetState(v string) {
	o.State = &v
}

// GetUTime returns the UTime field value if set, zero value otherwise.
func (o *GetFiatDepositV5RespData) GetUTime() string {
	if o == nil || IsNil(o.UTime) {
		var ret string
		return ret
	}
	return *o.UTime
}

// GetUTimeOk returns a tuple with the UTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositV5RespData) GetUTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UTime) {
		return nil, false
	}
	return o.UTime, true
}

// HasUTime returns a boolean if a field has been set.
func (o *GetFiatDepositV5RespData) HasUTime() bool {
	if o != nil && !IsNil(o.UTime) {
		return true
	}

	return false
}

// SetUTime gets a reference to the given string and assigns it to the UTime field.
func (o *GetFiatDepositV5RespData) SetUTime(v string) {
	o.UTime = &v
}

func (o GetFiatDepositV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFiatDepositV5RespData) ToMap() (map[string]interface{}, error) {
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

type NullableGetFiatDepositV5RespData struct {
	value *GetFiatDepositV5RespData
	isSet bool
}

func (v NullableGetFiatDepositV5RespData) Get() *GetFiatDepositV5RespData {
	return v.value
}

func (v *NullableGetFiatDepositV5RespData) Set(val *GetFiatDepositV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFiatDepositV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFiatDepositV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFiatDepositV5RespData(val *GetFiatDepositV5RespData) *NullableGetFiatDepositV5RespData {
	return &NullableGetFiatDepositV5RespData{value: val, isSet: true}
}

func (v NullableGetFiatDepositV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFiatDepositV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


