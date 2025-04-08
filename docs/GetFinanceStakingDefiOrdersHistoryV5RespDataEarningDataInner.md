# GetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | Pointer to **string** | Earning currency, e.g. &#x60;BTC&#x60; | [optional] [default to ""]
**EarningType** | Pointer to **string** | Earning type  &#x60;0&#x60;: Estimated earning  &#x60;1&#x60;: Cumulative earning | [optional] [default to ""]
**RealizedEarnings** | Pointer to **string** | Cumulative earning of redeemed orders  This field is just valid when the order is in redemption state | [optional] [default to ""]

## Methods

### NewGetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInner

`func NewGetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInner() *GetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInner`

NewGetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInner instantiates a new GetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInnerWithDefaults

`func NewGetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInnerWithDefaults() *GetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInner`

NewGetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInnerWithDefaults instantiates a new GetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetEarningType

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInner) GetEarningType() string`

GetEarningType returns the EarningType field if non-nil, zero value otherwise.

### GetEarningTypeOk

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInner) GetEarningTypeOk() (*string, bool)`

GetEarningTypeOk returns a tuple with the EarningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarningType

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInner) SetEarningType(v string)`

SetEarningType sets EarningType field to given value.

### HasEarningType

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInner) HasEarningType() bool`

HasEarningType returns a boolean if a field has been set.

### GetRealizedEarnings

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInner) GetRealizedEarnings() string`

GetRealizedEarnings returns the RealizedEarnings field if non-nil, zero value otherwise.

### GetRealizedEarningsOk

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInner) GetRealizedEarningsOk() (*string, bool)`

GetRealizedEarningsOk returns a tuple with the RealizedEarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedEarnings

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInner) SetRealizedEarnings(v string)`

SetRealizedEarnings sets RealizedEarnings field to given value.

### HasRealizedEarnings

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataEarningDataInner) HasRealizedEarnings() bool`

HasRealizedEarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


