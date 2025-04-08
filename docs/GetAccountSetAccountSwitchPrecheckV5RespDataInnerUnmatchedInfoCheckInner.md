# GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PosList** | Pointer to **[]string** | Unmatched position list (posId)   Applicable when type is related to positions, return [] for other scenarios | [optional] 
**TotalAsset** | Pointer to **string** | Total assets   Only applicable when type is &#x60;asset_validation&#x60;, return \&quot;\&quot; for other scenarios | [optional] [default to ""]
**Type** | Pointer to **string** | Unmatched information type   &#x60;asset_validation&#x60;: asset validation   &#x60;pending_orders&#x60;: order book pending orders   &#x60;pending_algos&#x60;: pending algo orders and trading bots, such as iceberg, recurring buy and twap   &#x60;isolated_margin&#x60;: isolated margin (quick margin and manual transfers)   &#x60;isolated_contract&#x60;: isolated contract (manual transfers)   &#x60;contract_long_short&#x60;: contract positions in hedge mode   &#x60;cross_margin&#x60;: cross margin positions   &#x60;cross_option_buyer&#x60;: cross options buyer   &#x60;isolated_option&#x60;: isolated options (only applicable to spot mode)   &#x60;growth_fund&#x60;: positions with trial funds   &#x60;all_positions&#x60;: all positions   &#x60;spot_lead_copy_only_simple_single&#x60;: copy trader and customize lead trader can only use spot mode or spot and futures mode   &#x60;stop_spot_custom&#x60;: spot customize copy trading   &#x60;stop_futures_custom&#x60;: contract customize copy trading   &#x60;lead_portfolio&#x60;: lead trader can not switch to portfolio margin mode   &#x60;futures_smart_sync&#x60;: you can not switch to spot mode when having smart contract sync   &#x60;vip_fixed_loan&#x60;: vip loan   &#x60;repay_borrowings&#x60;: borrowings   &#x60;compliance_restriction&#x60;: due to compliance restrictions, margin trading services are unavailable   &#x60;compliance_kyc2&#x60;: Due to compliance restrictions, margin trading services are unavailable. If you are not a resident of this region, please complete kyc2 identity verification. | [optional] [default to ""]

## Methods

### NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner

`func NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner() *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner`

NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner instantiates a new GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInnerWithDefaults

`func NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInnerWithDefaults() *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner`

NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInnerWithDefaults instantiates a new GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosList

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) GetPosList() []string`

GetPosList returns the PosList field if non-nil, zero value otherwise.

### GetPosListOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) GetPosListOk() (*[]string, bool)`

GetPosListOk returns a tuple with the PosList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosList

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) SetPosList(v []string)`

SetPosList sets PosList field to given value.

### HasPosList

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) HasPosList() bool`

HasPosList returns a boolean if a field has been set.

### GetTotalAsset

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) GetTotalAsset() string`

GetTotalAsset returns the TotalAsset field if non-nil, zero value otherwise.

### GetTotalAssetOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) GetTotalAssetOk() (*string, bool)`

GetTotalAssetOk returns a tuple with the TotalAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAsset

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) SetTotalAsset(v string)`

SetTotalAsset sets TotalAsset field to given value.

### HasTotalAsset

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) HasTotalAsset() bool`

HasTotalAsset returns a boolean if a field has been set.

### GetType

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


