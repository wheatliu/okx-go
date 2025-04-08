# GetFinanceStakingDefiOrdersActiveV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apy** | Pointer to **string** | Estimated APY  If the estimated APY is 7% , this field is 0.07  Retain to 4 decimal places (truncated) | [optional] [default to ""]
**CancelRedemptionDeadline** | Pointer to **string** | Deadline for cancellation of redemption application | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency, e.g. &#x60;BTC&#x60; | [optional] [default to ""]
**EarningData** | Pointer to [**[]GetFinanceStakingDefiOrdersActiveV5RespDataInnerEarningDataInner**](GetFinanceStakingDefiOrdersActiveV5RespDataInnerEarningDataInner.md) | Earning data | [optional] 
**EstSettlementTime** | Pointer to **string** | Estimated redemption settlement time | [optional] [default to ""]
**FastRedemptionData** | Pointer to [**[]GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner**](GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner.md) | Fast redemption data | [optional] 
**InvestData** | Pointer to [**[]GetFinanceStakingDefiOrdersActiveV5RespDataInnerInvestDataInner**](GetFinanceStakingDefiOrdersActiveV5RespDataInnerInvestDataInner.md) | Investment data | [optional] 
**OrdId** | Pointer to **string** | Order ID | [optional] [default to ""]
**ProductId** | Pointer to **string** | Product ID | [optional] [default to ""]
**Protocol** | Pointer to **string** | Protocol | [optional] [default to ""]
**ProtocolType** | Pointer to **string** | Protocol type  &#x60;defi&#x60;: on-chain earn | [optional] [default to ""]
**PurchasedTime** | Pointer to **string** | Order purchased time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**State** | Pointer to **string** | Order state  8: Pending   13: Cancelling   9: Onchain   1: Earning   2: Redeeming | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**Term** | Pointer to **string** | Protocol term  It will return the days of fixed term and will return &#x60;0&#x60; for flexible product | [optional] [default to ""]

## Methods

### NewGetFinanceStakingDefiOrdersActiveV5RespDataInner

`func NewGetFinanceStakingDefiOrdersActiveV5RespDataInner() *GetFinanceStakingDefiOrdersActiveV5RespDataInner`

NewGetFinanceStakingDefiOrdersActiveV5RespDataInner instantiates a new GetFinanceStakingDefiOrdersActiveV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinanceStakingDefiOrdersActiveV5RespDataInnerWithDefaults

`func NewGetFinanceStakingDefiOrdersActiveV5RespDataInnerWithDefaults() *GetFinanceStakingDefiOrdersActiveV5RespDataInner`

NewGetFinanceStakingDefiOrdersActiveV5RespDataInnerWithDefaults instantiates a new GetFinanceStakingDefiOrdersActiveV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApy

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetApy() string`

GetApy returns the Apy field if non-nil, zero value otherwise.

### GetApyOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetApyOk() (*string, bool)`

GetApyOk returns a tuple with the Apy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApy

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) SetApy(v string)`

SetApy sets Apy field to given value.

### HasApy

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) HasApy() bool`

HasApy returns a boolean if a field has been set.

### GetCancelRedemptionDeadline

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetCancelRedemptionDeadline() string`

GetCancelRedemptionDeadline returns the CancelRedemptionDeadline field if non-nil, zero value otherwise.

### GetCancelRedemptionDeadlineOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetCancelRedemptionDeadlineOk() (*string, bool)`

GetCancelRedemptionDeadlineOk returns a tuple with the CancelRedemptionDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelRedemptionDeadline

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) SetCancelRedemptionDeadline(v string)`

SetCancelRedemptionDeadline sets CancelRedemptionDeadline field to given value.

### HasCancelRedemptionDeadline

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) HasCancelRedemptionDeadline() bool`

HasCancelRedemptionDeadline returns a boolean if a field has been set.

### GetCcy

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetEarningData

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetEarningData() []GetFinanceStakingDefiOrdersActiveV5RespDataInnerEarningDataInner`

GetEarningData returns the EarningData field if non-nil, zero value otherwise.

### GetEarningDataOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetEarningDataOk() (*[]GetFinanceStakingDefiOrdersActiveV5RespDataInnerEarningDataInner, bool)`

GetEarningDataOk returns a tuple with the EarningData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarningData

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) SetEarningData(v []GetFinanceStakingDefiOrdersActiveV5RespDataInnerEarningDataInner)`

SetEarningData sets EarningData field to given value.

### HasEarningData

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) HasEarningData() bool`

HasEarningData returns a boolean if a field has been set.

### GetEstSettlementTime

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetEstSettlementTime() string`

GetEstSettlementTime returns the EstSettlementTime field if non-nil, zero value otherwise.

### GetEstSettlementTimeOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetEstSettlementTimeOk() (*string, bool)`

GetEstSettlementTimeOk returns a tuple with the EstSettlementTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstSettlementTime

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) SetEstSettlementTime(v string)`

SetEstSettlementTime sets EstSettlementTime field to given value.

### HasEstSettlementTime

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) HasEstSettlementTime() bool`

HasEstSettlementTime returns a boolean if a field has been set.

### GetFastRedemptionData

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetFastRedemptionData() []GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner`

GetFastRedemptionData returns the FastRedemptionData field if non-nil, zero value otherwise.

### GetFastRedemptionDataOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetFastRedemptionDataOk() (*[]GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner, bool)`

GetFastRedemptionDataOk returns a tuple with the FastRedemptionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastRedemptionData

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) SetFastRedemptionData(v []GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner)`

SetFastRedemptionData sets FastRedemptionData field to given value.

### HasFastRedemptionData

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) HasFastRedemptionData() bool`

HasFastRedemptionData returns a boolean if a field has been set.

### GetInvestData

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetInvestData() []GetFinanceStakingDefiOrdersActiveV5RespDataInnerInvestDataInner`

GetInvestData returns the InvestData field if non-nil, zero value otherwise.

### GetInvestDataOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetInvestDataOk() (*[]GetFinanceStakingDefiOrdersActiveV5RespDataInnerInvestDataInner, bool)`

GetInvestDataOk returns a tuple with the InvestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestData

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) SetInvestData(v []GetFinanceStakingDefiOrdersActiveV5RespDataInnerInvestDataInner)`

SetInvestData sets InvestData field to given value.

### HasInvestData

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) HasInvestData() bool`

HasInvestData returns a boolean if a field has been set.

### GetOrdId

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetProductId

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProtocol

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetProtocolType

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetPurchasedTime

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetPurchasedTime() string`

GetPurchasedTime returns the PurchasedTime field if non-nil, zero value otherwise.

### GetPurchasedTimeOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetPurchasedTimeOk() (*string, bool)`

GetPurchasedTimeOk returns a tuple with the PurchasedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedTime

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) SetPurchasedTime(v string)`

SetPurchasedTime sets PurchasedTime field to given value.

### HasPurchasedTime

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) HasPurchasedTime() bool`

HasPurchasedTime returns a boolean if a field has been set.

### GetState

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTag

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTerm

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetTerm() string`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) GetTermOk() (*string, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) SetTerm(v string)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInner) HasTerm() bool`

HasTerm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


