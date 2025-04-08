# GetPublicDiscountRateInterestFreeQuotaV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | Pointer to **string** | Interest-free quota | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency | [optional] [default to ""]
**Details** | Pointer to [**[]GetPublicDiscountRateInterestFreeQuotaV5RespDataDetailsInner**](GetPublicDiscountRateInterestFreeQuotaV5RespDataDetailsInner.md) | New discount details. | [optional] 
**DiscountLv** | Pointer to **string** | Discount rate level.(Deprecated) | [optional] [default to ""]
**MinDiscountRate** | Pointer to **string** | Minimum discount rate when it exceeds the maximum amount of the last tier. | [optional] [default to ""]

## Methods

### NewGetPublicDiscountRateInterestFreeQuotaV5RespData

`func NewGetPublicDiscountRateInterestFreeQuotaV5RespData() *GetPublicDiscountRateInterestFreeQuotaV5RespData`

NewGetPublicDiscountRateInterestFreeQuotaV5RespData instantiates a new GetPublicDiscountRateInterestFreeQuotaV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicDiscountRateInterestFreeQuotaV5RespDataWithDefaults

`func NewGetPublicDiscountRateInterestFreeQuotaV5RespDataWithDefaults() *GetPublicDiscountRateInterestFreeQuotaV5RespData`

NewGetPublicDiscountRateInterestFreeQuotaV5RespDataWithDefaults instantiates a new GetPublicDiscountRateInterestFreeQuotaV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespData) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespData) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespData) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespData) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetCcy

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetDetails

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespData) GetDetails() []GetPublicDiscountRateInterestFreeQuotaV5RespDataDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespData) GetDetailsOk() (*[]GetPublicDiscountRateInterestFreeQuotaV5RespDataDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespData) SetDetails(v []GetPublicDiscountRateInterestFreeQuotaV5RespDataDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespData) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetDiscountLv

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespData) GetDiscountLv() string`

GetDiscountLv returns the DiscountLv field if non-nil, zero value otherwise.

### GetDiscountLvOk

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespData) GetDiscountLvOk() (*string, bool)`

GetDiscountLvOk returns a tuple with the DiscountLv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountLv

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespData) SetDiscountLv(v string)`

SetDiscountLv sets DiscountLv field to given value.

### HasDiscountLv

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespData) HasDiscountLv() bool`

HasDiscountLv returns a boolean if a field has been set.

### GetMinDiscountRate

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespData) GetMinDiscountRate() string`

GetMinDiscountRate returns the MinDiscountRate field if non-nil, zero value otherwise.

### GetMinDiscountRateOk

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespData) GetMinDiscountRateOk() (*string, bool)`

GetMinDiscountRateOk returns a tuple with the MinDiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDiscountRate

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespData) SetMinDiscountRate(v string)`

SetMinDiscountRate sets MinDiscountRate field to given value.

### HasMinDiscountRate

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespData) HasMinDiscountRate() bool`

HasMinDiscountRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


