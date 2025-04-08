# GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | Pointer to **string** | Purchase/Redeem amount | [optional] [default to ""]
**CompletedTime** | Pointer to **string** | Completed time of redeem settlement, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**EstCompletedTime** | Pointer to **string** | Estimated completed time of redeem settlement, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**RedeemingAmt** | Pointer to **string** | Redeeming amount | [optional] [default to ""]
**RequestTime** | Pointer to **string** | Request time of make purchase/redeem, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Status** | Pointer to **string** | Status  &#x60;pending&#x60;  &#x60;success&#x60;  &#x60;failed&#x60; | [optional] [default to ""]
**Type** | Pointer to **string** | Type  &#x60;purchase&#x60;  &#x60;redeem&#x60; | [optional] [default to ""]

## Methods

### NewGetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner

`func NewGetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner() *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner`

NewGetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner instantiates a new GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInnerWithDefaults

`func NewGetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInnerWithDefaults() *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner`

NewGetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInnerWithDefaults instantiates a new GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetCompletedTime

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) GetCompletedTime() string`

GetCompletedTime returns the CompletedTime field if non-nil, zero value otherwise.

### GetCompletedTimeOk

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) GetCompletedTimeOk() (*string, bool)`

GetCompletedTimeOk returns a tuple with the CompletedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedTime

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) SetCompletedTime(v string)`

SetCompletedTime sets CompletedTime field to given value.

### HasCompletedTime

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) HasCompletedTime() bool`

HasCompletedTime returns a boolean if a field has been set.

### GetEstCompletedTime

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) GetEstCompletedTime() string`

GetEstCompletedTime returns the EstCompletedTime field if non-nil, zero value otherwise.

### GetEstCompletedTimeOk

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) GetEstCompletedTimeOk() (*string, bool)`

GetEstCompletedTimeOk returns a tuple with the EstCompletedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstCompletedTime

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) SetEstCompletedTime(v string)`

SetEstCompletedTime sets EstCompletedTime field to given value.

### HasEstCompletedTime

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) HasEstCompletedTime() bool`

HasEstCompletedTime returns a boolean if a field has been set.

### GetRedeemingAmt

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) GetRedeemingAmt() string`

GetRedeemingAmt returns the RedeemingAmt field if non-nil, zero value otherwise.

### GetRedeemingAmtOk

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) GetRedeemingAmtOk() (*string, bool)`

GetRedeemingAmtOk returns a tuple with the RedeemingAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemingAmt

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) SetRedeemingAmt(v string)`

SetRedeemingAmt sets RedeemingAmt field to given value.

### HasRedeemingAmt

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) HasRedeemingAmt() bool`

HasRedeemingAmt returns a boolean if a field has been set.

### GetRequestTime

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) GetRequestTime() string`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) GetRequestTimeOk() (*string, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) SetRequestTime(v string)`

SetRequestTime sets RequestTime field to given value.

### HasRequestTime

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) HasRequestTime() bool`

HasRequestTime returns a boolean if a field has been set.

### GetStatus

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5RespDataInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


