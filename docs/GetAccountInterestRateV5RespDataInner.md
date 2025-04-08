# GetAccountInterestRateV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | Pointer to **string** | currency | [optional] [default to ""]
**InterestRate** | Pointer to **string** | interest rate(the current hour) | [optional] [default to ""]

## Methods

### NewGetAccountInterestRateV5RespDataInner

`func NewGetAccountInterestRateV5RespDataInner() *GetAccountInterestRateV5RespDataInner`

NewGetAccountInterestRateV5RespDataInner instantiates a new GetAccountInterestRateV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountInterestRateV5RespDataInnerWithDefaults

`func NewGetAccountInterestRateV5RespDataInnerWithDefaults() *GetAccountInterestRateV5RespDataInner`

NewGetAccountInterestRateV5RespDataInnerWithDefaults instantiates a new GetAccountInterestRateV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *GetAccountInterestRateV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAccountInterestRateV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAccountInterestRateV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAccountInterestRateV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetInterestRate

`func (o *GetAccountInterestRateV5RespDataInner) GetInterestRate() string`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *GetAccountInterestRateV5RespDataInner) GetInterestRateOk() (*string, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *GetAccountInterestRateV5RespDataInner) SetInterestRate(v string)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *GetAccountInterestRateV5RespDataInner) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


