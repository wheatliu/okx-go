/*
Okx Rest API

OpenAPI specification for Okx cryptocurrency exchange - Rest API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rest

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateAssetSubaccountTransferV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAssetSubaccountTransferV5Req{}

// CreateAssetSubaccountTransferV5Req struct for CreateAssetSubaccountTransferV5Req
type CreateAssetSubaccountTransferV5Req struct {
	// Transfer amount
	Amt string `json:"amt"`
	// Currency
	Ccy string `json:"ccy"`
	// Account type of transfer from sub-account  `6`: Funding Account  `18`: Trading account
	From string `json:"from"`
	// Sub-account name of the account that transfers funds out.
	FromSubAccount string `json:"fromSubAccount"`
	// Whether or not borrowed coins can be transferred out under `Multi-currency margin`/`Portfolio margin`  The default is `false`
	LoanTrans *bool `json:"loanTrans,omitempty"`
	// Ignore position risk  Default is `false`  Applicable to `Portfolio margin`
	OmitPosRisk *string `json:"omitPosRisk,omitempty"`
	// Account type of transfer to sub-account  `6`: Funding Account  `18`: Trading account
	To string `json:"to"`
	// Sub-account name of the account that transfers funds in.
	ToSubAccount string `json:"toSubAccount"`
}

type _CreateAssetSubaccountTransferV5Req CreateAssetSubaccountTransferV5Req

// NewCreateAssetSubaccountTransferV5Req instantiates a new CreateAssetSubaccountTransferV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAssetSubaccountTransferV5Req(amt string, ccy string, from string, fromSubAccount string, to string, toSubAccount string) *CreateAssetSubaccountTransferV5Req {
	this := CreateAssetSubaccountTransferV5Req{}
	this.Amt = amt
	this.Ccy = ccy
	this.From = from
	this.FromSubAccount = fromSubAccount
	var omitPosRisk string = ""
	this.OmitPosRisk = &omitPosRisk
	this.To = to
	this.ToSubAccount = toSubAccount
	return &this
}

// NewCreateAssetSubaccountTransferV5ReqWithDefaults instantiates a new CreateAssetSubaccountTransferV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAssetSubaccountTransferV5ReqWithDefaults() *CreateAssetSubaccountTransferV5Req {
	this := CreateAssetSubaccountTransferV5Req{}
	var amt string = ""
	this.Amt = amt
	var ccy string = ""
	this.Ccy = ccy
	var from string = ""
	this.From = from
	var fromSubAccount string = ""
	this.FromSubAccount = fromSubAccount
	var omitPosRisk string = ""
	this.OmitPosRisk = &omitPosRisk
	var to string = ""
	this.To = to
	var toSubAccount string = ""
	this.ToSubAccount = toSubAccount
	return &this
}

// GetAmt returns the Amt field value
func (o *CreateAssetSubaccountTransferV5Req) GetAmt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amt
}

// GetAmtOk returns a tuple with the Amt field value
// and a boolean to check if the value has been set.
func (o *CreateAssetSubaccountTransferV5Req) GetAmtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amt, true
}

// SetAmt sets field value
func (o *CreateAssetSubaccountTransferV5Req) SetAmt(v string) {
	o.Amt = v
}

// GetCcy returns the Ccy field value
func (o *CreateAssetSubaccountTransferV5Req) GetCcy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value
// and a boolean to check if the value has been set.
func (o *CreateAssetSubaccountTransferV5Req) GetCcyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ccy, true
}

// SetCcy sets field value
func (o *CreateAssetSubaccountTransferV5Req) SetCcy(v string) {
	o.Ccy = v
}

// GetFrom returns the From field value
func (o *CreateAssetSubaccountTransferV5Req) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *CreateAssetSubaccountTransferV5Req) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *CreateAssetSubaccountTransferV5Req) SetFrom(v string) {
	o.From = v
}

// GetFromSubAccount returns the FromSubAccount field value
func (o *CreateAssetSubaccountTransferV5Req) GetFromSubAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromSubAccount
}

// GetFromSubAccountOk returns a tuple with the FromSubAccount field value
// and a boolean to check if the value has been set.
func (o *CreateAssetSubaccountTransferV5Req) GetFromSubAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromSubAccount, true
}

// SetFromSubAccount sets field value
func (o *CreateAssetSubaccountTransferV5Req) SetFromSubAccount(v string) {
	o.FromSubAccount = v
}

// GetLoanTrans returns the LoanTrans field value if set, zero value otherwise.
func (o *CreateAssetSubaccountTransferV5Req) GetLoanTrans() bool {
	if o == nil || IsNil(o.LoanTrans) {
		var ret bool
		return ret
	}
	return *o.LoanTrans
}

// GetLoanTransOk returns a tuple with the LoanTrans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetSubaccountTransferV5Req) GetLoanTransOk() (*bool, bool) {
	if o == nil || IsNil(o.LoanTrans) {
		return nil, false
	}
	return o.LoanTrans, true
}

// HasLoanTrans returns a boolean if a field has been set.
func (o *CreateAssetSubaccountTransferV5Req) HasLoanTrans() bool {
	if o != nil && !IsNil(o.LoanTrans) {
		return true
	}

	return false
}

// SetLoanTrans gets a reference to the given bool and assigns it to the LoanTrans field.
func (o *CreateAssetSubaccountTransferV5Req) SetLoanTrans(v bool) {
	o.LoanTrans = &v
}

// GetOmitPosRisk returns the OmitPosRisk field value if set, zero value otherwise.
func (o *CreateAssetSubaccountTransferV5Req) GetOmitPosRisk() string {
	if o == nil || IsNil(o.OmitPosRisk) {
		var ret string
		return ret
	}
	return *o.OmitPosRisk
}

// GetOmitPosRiskOk returns a tuple with the OmitPosRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetSubaccountTransferV5Req) GetOmitPosRiskOk() (*string, bool) {
	if o == nil || IsNil(o.OmitPosRisk) {
		return nil, false
	}
	return o.OmitPosRisk, true
}

// HasOmitPosRisk returns a boolean if a field has been set.
func (o *CreateAssetSubaccountTransferV5Req) HasOmitPosRisk() bool {
	if o != nil && !IsNil(o.OmitPosRisk) {
		return true
	}

	return false
}

// SetOmitPosRisk gets a reference to the given string and assigns it to the OmitPosRisk field.
func (o *CreateAssetSubaccountTransferV5Req) SetOmitPosRisk(v string) {
	o.OmitPosRisk = &v
}

// GetTo returns the To field value
func (o *CreateAssetSubaccountTransferV5Req) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *CreateAssetSubaccountTransferV5Req) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *CreateAssetSubaccountTransferV5Req) SetTo(v string) {
	o.To = v
}

// GetToSubAccount returns the ToSubAccount field value
func (o *CreateAssetSubaccountTransferV5Req) GetToSubAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToSubAccount
}

// GetToSubAccountOk returns a tuple with the ToSubAccount field value
// and a boolean to check if the value has been set.
func (o *CreateAssetSubaccountTransferV5Req) GetToSubAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToSubAccount, true
}

// SetToSubAccount sets field value
func (o *CreateAssetSubaccountTransferV5Req) SetToSubAccount(v string) {
	o.ToSubAccount = v
}

func (o CreateAssetSubaccountTransferV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAssetSubaccountTransferV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amt"] = o.Amt
	toSerialize["ccy"] = o.Ccy
	toSerialize["from"] = o.From
	toSerialize["fromSubAccount"] = o.FromSubAccount
	if !IsNil(o.LoanTrans) {
		toSerialize["loanTrans"] = o.LoanTrans
	}
	if !IsNil(o.OmitPosRisk) {
		toSerialize["omitPosRisk"] = o.OmitPosRisk
	}
	toSerialize["to"] = o.To
	toSerialize["toSubAccount"] = o.ToSubAccount
	return toSerialize, nil
}

func (o *CreateAssetSubaccountTransferV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amt",
		"ccy",
		"from",
		"fromSubAccount",
		"to",
		"toSubAccount",
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

	varCreateAssetSubaccountTransferV5Req := _CreateAssetSubaccountTransferV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateAssetSubaccountTransferV5Req)

	if err != nil {
		return err
	}

	*o = CreateAssetSubaccountTransferV5Req(varCreateAssetSubaccountTransferV5Req)

	return err
}

type NullableCreateAssetSubaccountTransferV5Req struct {
	value *CreateAssetSubaccountTransferV5Req
	isSet bool
}

func (v NullableCreateAssetSubaccountTransferV5Req) Get() *CreateAssetSubaccountTransferV5Req {
	return v.value
}

func (v *NullableCreateAssetSubaccountTransferV5Req) Set(val *CreateAssetSubaccountTransferV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAssetSubaccountTransferV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAssetSubaccountTransferV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAssetSubaccountTransferV5Req(val *CreateAssetSubaccountTransferV5Req) *NullableCreateAssetSubaccountTransferV5Req {
	return &NullableCreateAssetSubaccountTransferV5Req{value: val, isSet: true}
}

func (v NullableCreateAssetSubaccountTransferV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAssetSubaccountTransferV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


