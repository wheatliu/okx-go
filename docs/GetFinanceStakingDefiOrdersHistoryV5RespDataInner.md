# GetFinanceStakingDefiOrdersHistoryV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apy** | Pointer to **string** | Estimated APY  If the estimated APY is 7% , this field is &#x60;0.07&#x60;  Retain to 4 decimal places (truncated) | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency, e.g. &#x60;BTC&#x60; | [optional] [default to ""]
**EarningData** | Pointer to [**[]GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner**](GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner.md) | Earning data | [optional] 
**InvestData** | Pointer to [**[]GetFinanceStakingDefiOrdersActiveV5RespDataInnerInvestDataInner**](GetFinanceStakingDefiOrdersActiveV5RespDataInnerInvestDataInner.md) | Investment data | [optional] 
**OrdId** | Pointer to **string** | Order ID | [optional] [default to ""]
**ProductId** | Pointer to **string** | Product ID | [optional] [default to ""]
**Protocol** | Pointer to **string** | Protocol | [optional] [default to ""]
**ProtocolType** | Pointer to **string** | Protocol type  &#x60;defi&#x60;: on-chain earn | [optional] [default to ""]
**PurchasedTime** | Pointer to **string** | Order purchased time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**RedeemedTime** | Pointer to **string** | Order redeemed time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**State** | Pointer to **string** | Order state  &#x60;3&#x60;: Completed (including canceled and redeemed) | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**Term** | Pointer to **string** | Protocol term  It will return the days of fixed term and will return &#x60;0&#x60; for flexible product | [optional] [default to ""]

## Methods

### NewGetFinanceStakingDefiOrdersHistoryV5RespDataInner

`func NewGetFinanceStakingDefiOrdersHistoryV5RespDataInner() *GetFinanceStakingDefiOrdersHistoryV5RespDataInner`

NewGetFinanceStakingDefiOrdersHistoryV5RespDataInner instantiates a new GetFinanceStakingDefiOrdersHistoryV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinanceStakingDefiOrdersHistoryV5RespDataInnerWithDefaults

`func NewGetFinanceStakingDefiOrdersHistoryV5RespDataInnerWithDefaults() *GetFinanceStakingDefiOrdersHistoryV5RespDataInner`

NewGetFinanceStakingDefiOrdersHistoryV5RespDataInnerWithDefaults instantiates a new GetFinanceStakingDefiOrdersHistoryV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApy

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetApy() string`

GetApy returns the Apy field if non-nil, zero value otherwise.

### GetApyOk

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetApyOk() (*string, bool)`

GetApyOk returns a tuple with the Apy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApy

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) SetApy(v string)`

SetApy sets Apy field to given value.

### HasApy

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) HasApy() bool`

HasApy returns a boolean if a field has been set.

### GetCcy

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetEarningData

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetEarningData() []GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner`

GetEarningData returns the EarningData field if non-nil, zero value otherwise.

### GetEarningDataOk

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetEarningDataOk() (*[]GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner, bool)`

GetEarningDataOk returns a tuple with the EarningData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarningData

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) SetEarningData(v []GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner)`

SetEarningData sets EarningData field to given value.

### HasEarningData

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) HasEarningData() bool`

HasEarningData returns a boolean if a field has been set.

### GetInvestData

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetInvestData() []GetFinanceStakingDefiOrdersActiveV5RespDataInnerInvestDataInner`

GetInvestData returns the InvestData field if non-nil, zero value otherwise.

### GetInvestDataOk

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetInvestDataOk() (*[]GetFinanceStakingDefiOrdersActiveV5RespDataInnerInvestDataInner, bool)`

GetInvestDataOk returns a tuple with the InvestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestData

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) SetInvestData(v []GetFinanceStakingDefiOrdersActiveV5RespDataInnerInvestDataInner)`

SetInvestData sets InvestData field to given value.

### HasInvestData

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) HasInvestData() bool`

HasInvestData returns a boolean if a field has been set.

### GetOrdId

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetProductId

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProtocol

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetProtocolType

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetPurchasedTime

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetPurchasedTime() string`

GetPurchasedTime returns the PurchasedTime field if non-nil, zero value otherwise.

### GetPurchasedTimeOk

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetPurchasedTimeOk() (*string, bool)`

GetPurchasedTimeOk returns a tuple with the PurchasedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedTime

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) SetPurchasedTime(v string)`

SetPurchasedTime sets PurchasedTime field to given value.

### HasPurchasedTime

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) HasPurchasedTime() bool`

HasPurchasedTime returns a boolean if a field has been set.

### GetRedeemedTime

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetRedeemedTime() string`

GetRedeemedTime returns the RedeemedTime field if non-nil, zero value otherwise.

### GetRedeemedTimeOk

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetRedeemedTimeOk() (*string, bool)`

GetRedeemedTimeOk returns a tuple with the RedeemedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemedTime

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) SetRedeemedTime(v string)`

SetRedeemedTime sets RedeemedTime field to given value.

### HasRedeemedTime

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) HasRedeemedTime() bool`

HasRedeemedTime returns a boolean if a field has been set.

### GetState

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTag

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTerm

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetTerm() string`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) GetTermOk() (*string, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) SetTerm(v string)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInner) HasTerm() bool`

HasTerm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


