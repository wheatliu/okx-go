# GetTradeOneClickRepayHistoryV2V5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DebtCcy** | Pointer to **string** | Debt currency | [optional] [default to ""]
**FillDebtSz** | Pointer to **string** | Amount of debt currency transacted | [optional] [default to ""]
**OrdIdInfo** | Pointer to [**[]GetTradeOneClickRepayHistoryV2V5RespDataOrdIdInfoInner**](GetTradeOneClickRepayHistoryV2V5RespDataOrdIdInfoInner.md) | Order info | [optional] 
**RepayCcyList** | Pointer to **[]string** | Repay currency list, e.g. [\&quot;USDC\&quot;,\&quot;BTC\&quot;] | [optional] 
**Status** | Pointer to **string** | Current status of one-click repay   &#x60;running&#x60;: Running   &#x60;filled&#x60;: Filled   &#x60;failed&#x60;: Failed | [optional] [default to ""]
**Ts** | Pointer to **string** | Request time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetTradeOneClickRepayHistoryV2V5RespData

`func NewGetTradeOneClickRepayHistoryV2V5RespData() *GetTradeOneClickRepayHistoryV2V5RespData`

NewGetTradeOneClickRepayHistoryV2V5RespData instantiates a new GetTradeOneClickRepayHistoryV2V5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradeOneClickRepayHistoryV2V5RespDataWithDefaults

`func NewGetTradeOneClickRepayHistoryV2V5RespDataWithDefaults() *GetTradeOneClickRepayHistoryV2V5RespData`

NewGetTradeOneClickRepayHistoryV2V5RespDataWithDefaults instantiates a new GetTradeOneClickRepayHistoryV2V5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebtCcy

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) GetDebtCcy() string`

GetDebtCcy returns the DebtCcy field if non-nil, zero value otherwise.

### GetDebtCcyOk

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) GetDebtCcyOk() (*string, bool)`

GetDebtCcyOk returns a tuple with the DebtCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebtCcy

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) SetDebtCcy(v string)`

SetDebtCcy sets DebtCcy field to given value.

### HasDebtCcy

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) HasDebtCcy() bool`

HasDebtCcy returns a boolean if a field has been set.

### GetFillDebtSz

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) GetFillDebtSz() string`

GetFillDebtSz returns the FillDebtSz field if non-nil, zero value otherwise.

### GetFillDebtSzOk

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) GetFillDebtSzOk() (*string, bool)`

GetFillDebtSzOk returns a tuple with the FillDebtSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillDebtSz

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) SetFillDebtSz(v string)`

SetFillDebtSz sets FillDebtSz field to given value.

### HasFillDebtSz

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) HasFillDebtSz() bool`

HasFillDebtSz returns a boolean if a field has been set.

### GetOrdIdInfo

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) GetOrdIdInfo() []GetTradeOneClickRepayHistoryV2V5RespDataOrdIdInfoInner`

GetOrdIdInfo returns the OrdIdInfo field if non-nil, zero value otherwise.

### GetOrdIdInfoOk

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) GetOrdIdInfoOk() (*[]GetTradeOneClickRepayHistoryV2V5RespDataOrdIdInfoInner, bool)`

GetOrdIdInfoOk returns a tuple with the OrdIdInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdIdInfo

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) SetOrdIdInfo(v []GetTradeOneClickRepayHistoryV2V5RespDataOrdIdInfoInner)`

SetOrdIdInfo sets OrdIdInfo field to given value.

### HasOrdIdInfo

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) HasOrdIdInfo() bool`

HasOrdIdInfo returns a boolean if a field has been set.

### GetRepayCcyList

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) GetRepayCcyList() []string`

GetRepayCcyList returns the RepayCcyList field if non-nil, zero value otherwise.

### GetRepayCcyListOk

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) GetRepayCcyListOk() (*[]string, bool)`

GetRepayCcyListOk returns a tuple with the RepayCcyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepayCcyList

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) SetRepayCcyList(v []string)`

SetRepayCcyList sets RepayCcyList field to given value.

### HasRepayCcyList

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) HasRepayCcyList() bool`

HasRepayCcyList returns a boolean if a field has been set.

### GetStatus

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTs

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetTradeOneClickRepayHistoryV2V5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


