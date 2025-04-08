# CreateFinanceSavingsPurchaseRedemptV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | **string** | Purchase/redemption amount | [default to ""]
**Ccy** | **string** | Currency, e.g. &#x60;BTC&#x60; | [default to ""]
**Rate** | **string** | Annual purchase rate  Only applicable to purchase saving shares  The interest rate of the new subscription will cover the interest rate of the last subscription  The rate value range is between 1% and 365% | [default to ""]
**Side** | **string** | Action type.   &#x60;purchase&#x60;: purchase saving shares   &#x60;redempt&#x60;: redeem saving shares | [default to ""]

## Methods

### NewCreateFinanceSavingsPurchaseRedemptV5Req

`func NewCreateFinanceSavingsPurchaseRedemptV5Req(amt string, ccy string, rate string, side string, ) *CreateFinanceSavingsPurchaseRedemptV5Req`

NewCreateFinanceSavingsPurchaseRedemptV5Req instantiates a new CreateFinanceSavingsPurchaseRedemptV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFinanceSavingsPurchaseRedemptV5ReqWithDefaults

`func NewCreateFinanceSavingsPurchaseRedemptV5ReqWithDefaults() *CreateFinanceSavingsPurchaseRedemptV5Req`

NewCreateFinanceSavingsPurchaseRedemptV5ReqWithDefaults instantiates a new CreateFinanceSavingsPurchaseRedemptV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *CreateFinanceSavingsPurchaseRedemptV5Req) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *CreateFinanceSavingsPurchaseRedemptV5Req) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *CreateFinanceSavingsPurchaseRedemptV5Req) SetAmt(v string)`

SetAmt sets Amt field to given value.


### GetCcy

`func (o *CreateFinanceSavingsPurchaseRedemptV5Req) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateFinanceSavingsPurchaseRedemptV5Req) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateFinanceSavingsPurchaseRedemptV5Req) SetCcy(v string)`

SetCcy sets Ccy field to given value.


### GetRate

`func (o *CreateFinanceSavingsPurchaseRedemptV5Req) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *CreateFinanceSavingsPurchaseRedemptV5Req) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *CreateFinanceSavingsPurchaseRedemptV5Req) SetRate(v string)`

SetRate sets Rate field to given value.


### GetSide

`func (o *CreateFinanceSavingsPurchaseRedemptV5Req) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateFinanceSavingsPurchaseRedemptV5Req) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateFinanceSavingsPurchaseRedemptV5Req) SetSide(v string)`

SetSide sets Side field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


