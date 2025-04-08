# GetFinanceStakingDefiOffersV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apy** | Pointer to **string** | Estimated annualization  If the annualization is 7% , this field is 0.07 | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency type, e.g. &#x60;BTC&#x60; | [optional] [default to ""]
**EarlyRedeem** | Pointer to **bool** | Whether the protocol supports early redemption | [optional] 
**EarningData** | Pointer to [**[]GetFinanceStakingDefiOffersV5RespDataEarningDataInner**](GetFinanceStakingDefiOffersV5RespDataEarningDataInner.md) | Earning data | [optional] 
**FastRedemptionDailyLimit** | Pointer to **string** | Fast redemption daily limit  If fast redemption is not supported, it will return &#39;&#39;. | [optional] [default to ""]
**InvestData** | Pointer to [**[]GetFinanceStakingDefiOffersV5RespDataInvestDataInner**](GetFinanceStakingDefiOffersV5RespDataInvestDataInner.md) | Current target currency information available for investment | [optional] 
**ProductId** | Pointer to **string** | Product ID | [optional] [default to ""]
**Protocol** | Pointer to **string** | Protocol | [optional] [default to ""]
**ProtocolType** | Pointer to **string** | Protocol type  &#x60;defi&#x60;: on-chain earn | [optional] [default to ""]
**RedeemPeriod** | Pointer to **[]string** | Redemption Period, format in [min time,max time]  &#x60;H&#x60;: Hour, &#x60;D&#x60;: Day  e.g. [\&quot;1H\&quot;,\&quot;24H\&quot;] represents redemption period is between 1 Hour and 24 Hours.  [\&quot;14D\&quot;,\&quot;14D\&quot;] represents redemption period is 14 days. | [optional] 
**State** | Pointer to **string** | Product state  &#x60;purchasable&#x60;: Purchasable  &#x60;sold_out&#x60;: Sold out  &#x60;Stop&#x60;: Suspension of subscription | [optional] [default to ""]
**Term** | Pointer to **string** | Protocol term  It will return the days of fixed term and will return &#x60;0&#x60; for flexible product | [optional] [default to ""]

## Methods

### NewGetFinanceStakingDefiOffersV5RespData

`func NewGetFinanceStakingDefiOffersV5RespData() *GetFinanceStakingDefiOffersV5RespData`

NewGetFinanceStakingDefiOffersV5RespData instantiates a new GetFinanceStakingDefiOffersV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinanceStakingDefiOffersV5RespDataWithDefaults

`func NewGetFinanceStakingDefiOffersV5RespDataWithDefaults() *GetFinanceStakingDefiOffersV5RespData`

NewGetFinanceStakingDefiOffersV5RespDataWithDefaults instantiates a new GetFinanceStakingDefiOffersV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApy

`func (o *GetFinanceStakingDefiOffersV5RespData) GetApy() string`

GetApy returns the Apy field if non-nil, zero value otherwise.

### GetApyOk

`func (o *GetFinanceStakingDefiOffersV5RespData) GetApyOk() (*string, bool)`

GetApyOk returns a tuple with the Apy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApy

`func (o *GetFinanceStakingDefiOffersV5RespData) SetApy(v string)`

SetApy sets Apy field to given value.

### HasApy

`func (o *GetFinanceStakingDefiOffersV5RespData) HasApy() bool`

HasApy returns a boolean if a field has been set.

### GetCcy

`func (o *GetFinanceStakingDefiOffersV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetFinanceStakingDefiOffersV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetFinanceStakingDefiOffersV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetFinanceStakingDefiOffersV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetEarlyRedeem

`func (o *GetFinanceStakingDefiOffersV5RespData) GetEarlyRedeem() bool`

GetEarlyRedeem returns the EarlyRedeem field if non-nil, zero value otherwise.

### GetEarlyRedeemOk

`func (o *GetFinanceStakingDefiOffersV5RespData) GetEarlyRedeemOk() (*bool, bool)`

GetEarlyRedeemOk returns a tuple with the EarlyRedeem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyRedeem

`func (o *GetFinanceStakingDefiOffersV5RespData) SetEarlyRedeem(v bool)`

SetEarlyRedeem sets EarlyRedeem field to given value.

### HasEarlyRedeem

`func (o *GetFinanceStakingDefiOffersV5RespData) HasEarlyRedeem() bool`

HasEarlyRedeem returns a boolean if a field has been set.

### GetEarningData

`func (o *GetFinanceStakingDefiOffersV5RespData) GetEarningData() []GetFinanceStakingDefiOffersV5RespDataEarningDataInner`

GetEarningData returns the EarningData field if non-nil, zero value otherwise.

### GetEarningDataOk

`func (o *GetFinanceStakingDefiOffersV5RespData) GetEarningDataOk() (*[]GetFinanceStakingDefiOffersV5RespDataEarningDataInner, bool)`

GetEarningDataOk returns a tuple with the EarningData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarningData

`func (o *GetFinanceStakingDefiOffersV5RespData) SetEarningData(v []GetFinanceStakingDefiOffersV5RespDataEarningDataInner)`

SetEarningData sets EarningData field to given value.

### HasEarningData

`func (o *GetFinanceStakingDefiOffersV5RespData) HasEarningData() bool`

HasEarningData returns a boolean if a field has been set.

### GetFastRedemptionDailyLimit

`func (o *GetFinanceStakingDefiOffersV5RespData) GetFastRedemptionDailyLimit() string`

GetFastRedemptionDailyLimit returns the FastRedemptionDailyLimit field if non-nil, zero value otherwise.

### GetFastRedemptionDailyLimitOk

`func (o *GetFinanceStakingDefiOffersV5RespData) GetFastRedemptionDailyLimitOk() (*string, bool)`

GetFastRedemptionDailyLimitOk returns a tuple with the FastRedemptionDailyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastRedemptionDailyLimit

`func (o *GetFinanceStakingDefiOffersV5RespData) SetFastRedemptionDailyLimit(v string)`

SetFastRedemptionDailyLimit sets FastRedemptionDailyLimit field to given value.

### HasFastRedemptionDailyLimit

`func (o *GetFinanceStakingDefiOffersV5RespData) HasFastRedemptionDailyLimit() bool`

HasFastRedemptionDailyLimit returns a boolean if a field has been set.

### GetInvestData

`func (o *GetFinanceStakingDefiOffersV5RespData) GetInvestData() []GetFinanceStakingDefiOffersV5RespDataInvestDataInner`

GetInvestData returns the InvestData field if non-nil, zero value otherwise.

### GetInvestDataOk

`func (o *GetFinanceStakingDefiOffersV5RespData) GetInvestDataOk() (*[]GetFinanceStakingDefiOffersV5RespDataInvestDataInner, bool)`

GetInvestDataOk returns a tuple with the InvestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestData

`func (o *GetFinanceStakingDefiOffersV5RespData) SetInvestData(v []GetFinanceStakingDefiOffersV5RespDataInvestDataInner)`

SetInvestData sets InvestData field to given value.

### HasInvestData

`func (o *GetFinanceStakingDefiOffersV5RespData) HasInvestData() bool`

HasInvestData returns a boolean if a field has been set.

### GetProductId

`func (o *GetFinanceStakingDefiOffersV5RespData) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *GetFinanceStakingDefiOffersV5RespData) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *GetFinanceStakingDefiOffersV5RespData) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *GetFinanceStakingDefiOffersV5RespData) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProtocol

`func (o *GetFinanceStakingDefiOffersV5RespData) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetFinanceStakingDefiOffersV5RespData) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetFinanceStakingDefiOffersV5RespData) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetFinanceStakingDefiOffersV5RespData) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetProtocolType

`func (o *GetFinanceStakingDefiOffersV5RespData) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *GetFinanceStakingDefiOffersV5RespData) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *GetFinanceStakingDefiOffersV5RespData) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *GetFinanceStakingDefiOffersV5RespData) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetRedeemPeriod

`func (o *GetFinanceStakingDefiOffersV5RespData) GetRedeemPeriod() []string`

GetRedeemPeriod returns the RedeemPeriod field if non-nil, zero value otherwise.

### GetRedeemPeriodOk

`func (o *GetFinanceStakingDefiOffersV5RespData) GetRedeemPeriodOk() (*[]string, bool)`

GetRedeemPeriodOk returns a tuple with the RedeemPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemPeriod

`func (o *GetFinanceStakingDefiOffersV5RespData) SetRedeemPeriod(v []string)`

SetRedeemPeriod sets RedeemPeriod field to given value.

### HasRedeemPeriod

`func (o *GetFinanceStakingDefiOffersV5RespData) HasRedeemPeriod() bool`

HasRedeemPeriod returns a boolean if a field has been set.

### GetState

`func (o *GetFinanceStakingDefiOffersV5RespData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetFinanceStakingDefiOffersV5RespData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetFinanceStakingDefiOffersV5RespData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetFinanceStakingDefiOffersV5RespData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTerm

`func (o *GetFinanceStakingDefiOffersV5RespData) GetTerm() string`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *GetFinanceStakingDefiOffersV5RespData) GetTermOk() (*string, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *GetFinanceStakingDefiOffersV5RespData) SetTerm(v string)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *GetFinanceStakingDefiOffersV5RespData) HasTerm() bool`

HasTerm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


