# GetAccountInterestRateV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | Pointer to **string** | currency | [optional] [default to ""]
**InterestRate** | Pointer to **string** | interest rate(the current hour) | [optional] [default to ""]

## Methods

### NewGetAccountInterestRateV5RespData

`func NewGetAccountInterestRateV5RespData() *GetAccountInterestRateV5RespData`

NewGetAccountInterestRateV5RespData instantiates a new GetAccountInterestRateV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountInterestRateV5RespDataWithDefaults

`func NewGetAccountInterestRateV5RespDataWithDefaults() *GetAccountInterestRateV5RespData`

NewGetAccountInterestRateV5RespDataWithDefaults instantiates a new GetAccountInterestRateV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *GetAccountInterestRateV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAccountInterestRateV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAccountInterestRateV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAccountInterestRateV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetInterestRate

`func (o *GetAccountInterestRateV5RespData) GetInterestRate() string`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *GetAccountInterestRateV5RespData) GetInterestRateOk() (*string, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *GetAccountInterestRateV5RespData) SetInterestRate(v string)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *GetAccountInterestRateV5RespData) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


