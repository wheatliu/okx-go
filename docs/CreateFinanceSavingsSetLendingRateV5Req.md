# CreateFinanceSavingsSetLendingRateV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | **string** | Currency, e.g. &#x60;BTC&#x60; | [default to ""]
**Rate** | **string** | Annual lending rate  The rate value range is between 1% and 365% | [default to ""]

## Methods

### NewCreateFinanceSavingsSetLendingRateV5Req

`func NewCreateFinanceSavingsSetLendingRateV5Req(ccy string, rate string, ) *CreateFinanceSavingsSetLendingRateV5Req`

NewCreateFinanceSavingsSetLendingRateV5Req instantiates a new CreateFinanceSavingsSetLendingRateV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFinanceSavingsSetLendingRateV5ReqWithDefaults

`func NewCreateFinanceSavingsSetLendingRateV5ReqWithDefaults() *CreateFinanceSavingsSetLendingRateV5Req`

NewCreateFinanceSavingsSetLendingRateV5ReqWithDefaults instantiates a new CreateFinanceSavingsSetLendingRateV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *CreateFinanceSavingsSetLendingRateV5Req) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateFinanceSavingsSetLendingRateV5Req) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateFinanceSavingsSetLendingRateV5Req) SetCcy(v string)`

SetCcy sets Ccy field to given value.


### GetRate

`func (o *CreateFinanceSavingsSetLendingRateV5Req) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *CreateFinanceSavingsSetLendingRateV5Req) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *CreateFinanceSavingsSetLendingRateV5Req) SetRate(v string)`

SetRate sets Rate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


