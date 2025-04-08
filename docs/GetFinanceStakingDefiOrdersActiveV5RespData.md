# GetFinanceStakingDefiOrdersActiveV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apy** | Pointer to **string** | Estimated APY  If the estimated APY is 7% , this field is 0.07  Retain to 4 decimal places (truncated) | [optional] [default to ""]
**CancelRedemptionDeadline** | Pointer to **string** | Deadline for cancellation of redemption application | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency, e.g. &#x60;BTC&#x60; | [optional] [default to ""]
**EarningData** | Pointer to [**[]GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner**](GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner.md) | Earning data | [optional] 
**EstSettlementTime** | Pointer to **string** | Estimated redemption settlement time | [optional] [default to ""]
**FastRedemptionData** | Pointer to [**[]GetFinanceStakingDefiOrdersActiveV5RespDataFastRedemptionDataInner**](GetFinanceStakingDefiOrdersActiveV5RespDataFastRedemptionDataInner.md) | Fast redemption data | [optional] 
**InvestData** | Pointer to [**[]GetFinanceStakingDefiOrdersActiveV5RespDataInvestDataInner**](GetFinanceStakingDefiOrdersActiveV5RespDataInvestDataInner.md) | Investment data | [optional] 
**OrdId** | Pointer to **string** | Order ID | [optional] [default to ""]
**ProductId** | Pointer to **string** | Product ID | [optional] [default to ""]
**Protocol** | Pointer to **string** | Protocol | [optional] [default to ""]
**ProtocolType** | Pointer to **string** | Protocol type  &#x60;defi&#x60;: on-chain earn | [optional] [default to ""]
**PurchasedTime** | Pointer to **string** | Order purchased time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**State** | Pointer to **string** | Order state  8: Pending   13: Cancelling   9: Onchain   1: Earning   2: Redeeming | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**Term** | Pointer to **string** | Protocol term  It will return the days of fixed term and will return &#x60;0&#x60; for flexible product | [optional] [default to ""]

## Methods

### NewGetFinanceStakingDefiOrdersActiveV5RespData

`func NewGetFinanceStakingDefiOrdersActiveV5RespData() *GetFinanceStakingDefiOrdersActiveV5RespData`

NewGetFinanceStakingDefiOrdersActiveV5RespData instantiates a new GetFinanceStakingDefiOrdersActiveV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinanceStakingDefiOrdersActiveV5RespDataWithDefaults

`func NewGetFinanceStakingDefiOrdersActiveV5RespDataWithDefaults() *GetFinanceStakingDefiOrdersActiveV5RespData`

NewGetFinanceStakingDefiOrdersActiveV5RespDataWithDefaults instantiates a new GetFinanceStakingDefiOrdersActiveV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApy

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetApy() string`

GetApy returns the Apy field if non-nil, zero value otherwise.

### GetApyOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetApyOk() (*string, bool)`

GetApyOk returns a tuple with the Apy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApy

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetApy(v string)`

SetApy sets Apy field to given value.

### HasApy

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasApy() bool`

HasApy returns a boolean if a field has been set.

### GetCancelRedemptionDeadline

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetCancelRedemptionDeadline() string`

GetCancelRedemptionDeadline returns the CancelRedemptionDeadline field if non-nil, zero value otherwise.

### GetCancelRedemptionDeadlineOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetCancelRedemptionDeadlineOk() (*string, bool)`

GetCancelRedemptionDeadlineOk returns a tuple with the CancelRedemptionDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelRedemptionDeadline

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetCancelRedemptionDeadline(v string)`

SetCancelRedemptionDeadline sets CancelRedemptionDeadline field to given value.

### HasCancelRedemptionDeadline

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasCancelRedemptionDeadline() bool`

HasCancelRedemptionDeadline returns a boolean if a field has been set.

### GetCcy

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetEarningData

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetEarningData() []GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner`

GetEarningData returns the EarningData field if non-nil, zero value otherwise.

### GetEarningDataOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetEarningDataOk() (*[]GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner, bool)`

GetEarningDataOk returns a tuple with the EarningData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarningData

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetEarningData(v []GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner)`

SetEarningData sets EarningData field to given value.

### HasEarningData

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasEarningData() bool`

HasEarningData returns a boolean if a field has been set.

### GetEstSettlementTime

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetEstSettlementTime() string`

GetEstSettlementTime returns the EstSettlementTime field if non-nil, zero value otherwise.

### GetEstSettlementTimeOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetEstSettlementTimeOk() (*string, bool)`

GetEstSettlementTimeOk returns a tuple with the EstSettlementTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstSettlementTime

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetEstSettlementTime(v string)`

SetEstSettlementTime sets EstSettlementTime field to given value.

### HasEstSettlementTime

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasEstSettlementTime() bool`

HasEstSettlementTime returns a boolean if a field has been set.

### GetFastRedemptionData

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetFastRedemptionData() []GetFinanceStakingDefiOrdersActiveV5RespDataFastRedemptionDataInner`

GetFastRedemptionData returns the FastRedemptionData field if non-nil, zero value otherwise.

### GetFastRedemptionDataOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetFastRedemptionDataOk() (*[]GetFinanceStakingDefiOrdersActiveV5RespDataFastRedemptionDataInner, bool)`

GetFastRedemptionDataOk returns a tuple with the FastRedemptionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastRedemptionData

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetFastRedemptionData(v []GetFinanceStakingDefiOrdersActiveV5RespDataFastRedemptionDataInner)`

SetFastRedemptionData sets FastRedemptionData field to given value.

### HasFastRedemptionData

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasFastRedemptionData() bool`

HasFastRedemptionData returns a boolean if a field has been set.

### GetInvestData

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetInvestData() []GetFinanceStakingDefiOrdersActiveV5RespDataInvestDataInner`

GetInvestData returns the InvestData field if non-nil, zero value otherwise.

### GetInvestDataOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetInvestDataOk() (*[]GetFinanceStakingDefiOrdersActiveV5RespDataInvestDataInner, bool)`

GetInvestDataOk returns a tuple with the InvestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestData

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetInvestData(v []GetFinanceStakingDefiOrdersActiveV5RespDataInvestDataInner)`

SetInvestData sets InvestData field to given value.

### HasInvestData

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasInvestData() bool`

HasInvestData returns a boolean if a field has been set.

### GetOrdId

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetProductId

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProtocol

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetProtocolType

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetPurchasedTime

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetPurchasedTime() string`

GetPurchasedTime returns the PurchasedTime field if non-nil, zero value otherwise.

### GetPurchasedTimeOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetPurchasedTimeOk() (*string, bool)`

GetPurchasedTimeOk returns a tuple with the PurchasedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedTime

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetPurchasedTime(v string)`

SetPurchasedTime sets PurchasedTime field to given value.

### HasPurchasedTime

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasPurchasedTime() bool`

HasPurchasedTime returns a boolean if a field has been set.

### GetState

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTag

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTerm

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetTerm() string`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetTermOk() (*string, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetTerm(v string)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasTerm() bool`

HasTerm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


