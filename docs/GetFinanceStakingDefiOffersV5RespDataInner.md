# GetFinanceStakingDefiOffersV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apy** | Pointer to **string** | Estimated annualization  If the annualization is 7% , this field is 0.07 | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency type, e.g. &#x60;BTC&#x60; | [optional] [default to ""]
**EarlyRedeem** | Pointer to **bool** | Whether the protocol supports early redemption | [optional] 
**EarningData** | Pointer to [**[]GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner**](GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner.md) | Earning data | [optional] 
**FastRedemptionDailyLimit** | Pointer to **string** | Fast redemption daily limit  If fast redemption is not supported, it will return &#39;&#39;. | [optional] [default to ""]
**InvestData** | Pointer to [**[]GetFinanceStakingDefiOffersV5RespDataInnerInvestDataInner**](GetFinanceStakingDefiOffersV5RespDataInnerInvestDataInner.md) | Current target currency information available for investment | [optional] 
**ProductId** | Pointer to **string** | Product ID | [optional] [default to ""]
**Protocol** | Pointer to **string** | Protocol | [optional] [default to ""]
**ProtocolType** | Pointer to **string** | Protocol type  &#x60;defi&#x60;: on-chain earn | [optional] [default to ""]
**RedeemPeriod** | Pointer to **[]string** | Redemption Period, format in [min time,max time]  &#x60;H&#x60;: Hour, &#x60;D&#x60;: Day  e.g. [\&quot;1H\&quot;,\&quot;24H\&quot;] represents redemption period is between 1 Hour and 24 Hours.  [\&quot;14D\&quot;,\&quot;14D\&quot;] represents redemption period is 14 days. | [optional] 
**State** | Pointer to **string** | Product state  &#x60;purchasable&#x60;: Purchasable  &#x60;sold_out&#x60;: Sold out  &#x60;Stop&#x60;: Suspension of subscription | [optional] [default to ""]
**Term** | Pointer to **string** | Protocol term  It will return the days of fixed term and will return &#x60;0&#x60; for flexible product | [optional] [default to ""]

## Methods

### NewGetFinanceStakingDefiOffersV5RespDataInner

`func NewGetFinanceStakingDefiOffersV5RespDataInner() *GetFinanceStakingDefiOffersV5RespDataInner`

NewGetFinanceStakingDefiOffersV5RespDataInner instantiates a new GetFinanceStakingDefiOffersV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinanceStakingDefiOffersV5RespDataInnerWithDefaults

`func NewGetFinanceStakingDefiOffersV5RespDataInnerWithDefaults() *GetFinanceStakingDefiOffersV5RespDataInner`

NewGetFinanceStakingDefiOffersV5RespDataInnerWithDefaults instantiates a new GetFinanceStakingDefiOffersV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApy

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetApy() string`

GetApy returns the Apy field if non-nil, zero value otherwise.

### GetApyOk

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetApyOk() (*string, bool)`

GetApyOk returns a tuple with the Apy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApy

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetApy(v string)`

SetApy sets Apy field to given value.

### HasApy

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasApy() bool`

HasApy returns a boolean if a field has been set.

### GetCcy

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetEarlyRedeem

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetEarlyRedeem() bool`

GetEarlyRedeem returns the EarlyRedeem field if non-nil, zero value otherwise.

### GetEarlyRedeemOk

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetEarlyRedeemOk() (*bool, bool)`

GetEarlyRedeemOk returns a tuple with the EarlyRedeem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyRedeem

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetEarlyRedeem(v bool)`

SetEarlyRedeem sets EarlyRedeem field to given value.

### HasEarlyRedeem

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasEarlyRedeem() bool`

HasEarlyRedeem returns a boolean if a field has been set.

### GetEarningData

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetEarningData() []GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner`

GetEarningData returns the EarningData field if non-nil, zero value otherwise.

### GetEarningDataOk

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetEarningDataOk() (*[]GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner, bool)`

GetEarningDataOk returns a tuple with the EarningData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarningData

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetEarningData(v []GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner)`

SetEarningData sets EarningData field to given value.

### HasEarningData

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasEarningData() bool`

HasEarningData returns a boolean if a field has been set.

### GetFastRedemptionDailyLimit

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetFastRedemptionDailyLimit() string`

GetFastRedemptionDailyLimit returns the FastRedemptionDailyLimit field if non-nil, zero value otherwise.

### GetFastRedemptionDailyLimitOk

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetFastRedemptionDailyLimitOk() (*string, bool)`

GetFastRedemptionDailyLimitOk returns a tuple with the FastRedemptionDailyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastRedemptionDailyLimit

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetFastRedemptionDailyLimit(v string)`

SetFastRedemptionDailyLimit sets FastRedemptionDailyLimit field to given value.

### HasFastRedemptionDailyLimit

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasFastRedemptionDailyLimit() bool`

HasFastRedemptionDailyLimit returns a boolean if a field has been set.

### GetInvestData

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetInvestData() []GetFinanceStakingDefiOffersV5RespDataInnerInvestDataInner`

GetInvestData returns the InvestData field if non-nil, zero value otherwise.

### GetInvestDataOk

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetInvestDataOk() (*[]GetFinanceStakingDefiOffersV5RespDataInnerInvestDataInner, bool)`

GetInvestDataOk returns a tuple with the InvestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestData

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetInvestData(v []GetFinanceStakingDefiOffersV5RespDataInnerInvestDataInner)`

SetInvestData sets InvestData field to given value.

### HasInvestData

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasInvestData() bool`

HasInvestData returns a boolean if a field has been set.

### GetProductId

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProtocol

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetProtocolType

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetRedeemPeriod

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetRedeemPeriod() []string`

GetRedeemPeriod returns the RedeemPeriod field if non-nil, zero value otherwise.

### GetRedeemPeriodOk

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetRedeemPeriodOk() (*[]string, bool)`

GetRedeemPeriodOk returns a tuple with the RedeemPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemPeriod

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetRedeemPeriod(v []string)`

SetRedeemPeriod sets RedeemPeriod field to given value.

### HasRedeemPeriod

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasRedeemPeriod() bool`

HasRedeemPeriod returns a boolean if a field has been set.

### GetState

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTerm

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetTerm() string`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetTermOk() (*string, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetTerm(v string)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasTerm() bool`

HasTerm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


