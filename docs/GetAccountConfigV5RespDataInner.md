# GetAccountConfigV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctLv** | Pointer to **string** | Account mode   &#x60;1&#x60;: Spot mode  &#x60;2&#x60;: Spot and futures mode  &#x60;3&#x60;: Multi-currency margin  &#x60;4&#x60;: Portfolio margin | [optional] [default to ""]
**AcctStpMode** | Pointer to **string** | Account self-trade prevention mode   &#x60;cancel_maker&#x60;   &#x60;cancel_taker&#x60;   &#x60;cancel_both&#x60;   Users can log in to the webpage through the master account to modify this configuration | [optional] [default to ""]
**AutoLoan** | Pointer to **bool** | Whether to borrow coins automatically  &#x60;true&#x60;: borrow coins automatically  &#x60;false&#x60;: not borrow coins automatically | [optional] 
**CtIsoMode** | Pointer to **string** | Contract isolated margin trading settings  &#x60;automatic&#x60;: Auto transfers  &#x60;autonomy&#x60;: Manual transfers | [optional] [default to ""]
**EnableSpotBorrow** | Pointer to **bool** | Whether borrow is allowed or not in &#x60;Spot mode&#x60;  &#x60;true&#x60;: Enabled  &#x60;false&#x60;: Disabled | [optional] 
**GreeksType** | Pointer to **string** | Current display type of Greeks  &#x60;PA&#x60;: Greeks in coins  &#x60;BS&#x60;: Black-Scholes Greeks in dollars | [optional] [default to ""]
**Ip** | Pointer to **string** | IP addresses that linked with current API key, separate with commas if more than one, e.g. &#x60;117.37.203.58,117.37.203.57&#x60;. It is an empty string \&quot;\&quot; if there is no IP bonded. | [optional] [default to ""]
**KycLv** | Pointer to **string** | Main account KYC level  &#x60;0&#x60;: No verification  &#x60;1&#x60;: level 1 completed  &#x60;2&#x60;: level 2 completed  &#x60;3&#x60;: level 3 completed  If the request originates from a subaccount, kycLv is the KYC level of the main account.   If the request originates from the main account, kycLv is the KYC level of the current account. | [optional] [default to ""]
**Label** | Pointer to **string** | API key note  of current request API key. No more than 50 letters (case sensitive) or numbers, which can be pure letters or pure numbers. | [optional] [default to ""]
**Level** | Pointer to **string** | The user level of the current real trading volume on the platform,  e.g &#x60;Lv1&#x60; | [optional] [default to ""]
**LevelTmp** | Pointer to **string** | Temporary experience user level of special users, e.g &#x60;Lv3&#x60; | [optional] [default to ""]
**LiquidationGear** | Pointer to **string** | The margin ratio level of liquidation alert    &#x60;3&#x60; and &#x60;-1&#x60; means that you will get hourly liquidation alerts on app and channel \&quot;Position risk warning\&quot; when your margin level drops to or below 300%. &#x60;-1&#x60; is the initial value which has the same effect as &#x60;-3&#x60;   &#x60;0&#x60; means that there is not alert | [optional] [default to ""]
**MainUid** | Pointer to **string** | Main Account ID of current request.   The current request account is main account if uid &#x3D; mainUid.  The current request account is sub-account if uid !&#x3D; mainUid. | [optional] [default to ""]
**MgnIsoMode** | Pointer to **string** | Margin isolated margin trading settings   &#x60;auto_transfers_ccy&#x60;: New auto transfers, enabling both base and quote currency as the margin for isolated margin trading  &#x60;automatic&#x60;: Auto transfers  &#x60;quick_margin&#x60;: Quick Margin Mode (For new accounts, including subaccounts, some defaults will be &#x60;automatic&#x60;, and others will be &#x60;quick_margin&#x60;) | [optional] [default to ""]
**OpAuth** | Pointer to **string** | Whether the optional trading was activated  &#x60;0&#x60;: not activate  &#x60;1&#x60;: activated | [optional] [default to ""]
**Perm** | Pointer to **string** | The permission of the current requesting API key or Access token  &#x60;read_only&#x60;: Read  &#x60;trade&#x60;: Trade  &#x60;withdraw&#x60;: Withdraw | [optional] [default to ""]
**PosMode** | Pointer to **string** | Position mode  &#x60;long_short_mode&#x60;: long/short, only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;  &#x60;net_mode&#x60;: net | [optional] [default to ""]
**RoleType** | Pointer to **string** | Role type  &#x60;0&#x60;: General user  &#x60;1&#x60;: Leading trader  &#x60;2&#x60;: Copy trader | [optional] [default to ""]
**SpotBorrowAutoRepay** | Pointer to **bool** | Whether auto-repay is allowed or not in &#x60;Spot mode&#x60;  &#x60;true&#x60;: Enabled  &#x60;false&#x60;: Disabled | [optional] 
**SpotOffsetType** | Pointer to **string** | Risk offset type  &#x60;1&#x60;: Spot-Derivatives(USDT) to be offsetted  &#x60;2&#x60;: Spot-Derivatives(Coin) to be offsetted  &#x60;3&#x60;: Only derivatives to be offsetted  Only applicable to &#x60;Portfolio margin&#x60;  (Deprecated) | [optional] [default to ""]
**SpotRoleType** | Pointer to **string** | SPOT copy trading role type.  &#x60;0&#x60;: General user；&#x60;1&#x60;: Leading trader；&#x60;2&#x60;: Copy trader | [optional] [default to ""]
**SpotTraderInsts** | Pointer to **[]string** | Spot lead trading instruments, only applicable to lead trader | [optional] 
**TraderInsts** | Pointer to **[]string** | Leading trade instruments, only applicable to Leading trader | [optional] 
**Type** | Pointer to **string** | Account type   &#x60;0&#x60;: Main account   &#x60;1&#x60;: Standard sub-account   &#x60;2&#x60;: Managed trading sub-account   &#x60;5&#x60;: Custody trading sub-account - Copper  &#x60;9&#x60;: Managed trading sub-account - Copper   &#x60;12&#x60;: Custody trading sub-account - Komainu | [optional] [default to ""]
**Uid** | Pointer to **string** | Account ID of current request. | [optional] [default to ""]

## Methods

### NewGetAccountConfigV5RespDataInner

`func NewGetAccountConfigV5RespDataInner() *GetAccountConfigV5RespDataInner`

NewGetAccountConfigV5RespDataInner instantiates a new GetAccountConfigV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountConfigV5RespDataInnerWithDefaults

`func NewGetAccountConfigV5RespDataInnerWithDefaults() *GetAccountConfigV5RespDataInner`

NewGetAccountConfigV5RespDataInnerWithDefaults instantiates a new GetAccountConfigV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctLv

`func (o *GetAccountConfigV5RespDataInner) GetAcctLv() string`

GetAcctLv returns the AcctLv field if non-nil, zero value otherwise.

### GetAcctLvOk

`func (o *GetAccountConfigV5RespDataInner) GetAcctLvOk() (*string, bool)`

GetAcctLvOk returns a tuple with the AcctLv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctLv

`func (o *GetAccountConfigV5RespDataInner) SetAcctLv(v string)`

SetAcctLv sets AcctLv field to given value.

### HasAcctLv

`func (o *GetAccountConfigV5RespDataInner) HasAcctLv() bool`

HasAcctLv returns a boolean if a field has been set.

### GetAcctStpMode

`func (o *GetAccountConfigV5RespDataInner) GetAcctStpMode() string`

GetAcctStpMode returns the AcctStpMode field if non-nil, zero value otherwise.

### GetAcctStpModeOk

`func (o *GetAccountConfigV5RespDataInner) GetAcctStpModeOk() (*string, bool)`

GetAcctStpModeOk returns a tuple with the AcctStpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctStpMode

`func (o *GetAccountConfigV5RespDataInner) SetAcctStpMode(v string)`

SetAcctStpMode sets AcctStpMode field to given value.

### HasAcctStpMode

`func (o *GetAccountConfigV5RespDataInner) HasAcctStpMode() bool`

HasAcctStpMode returns a boolean if a field has been set.

### GetAutoLoan

`func (o *GetAccountConfigV5RespDataInner) GetAutoLoan() bool`

GetAutoLoan returns the AutoLoan field if non-nil, zero value otherwise.

### GetAutoLoanOk

`func (o *GetAccountConfigV5RespDataInner) GetAutoLoanOk() (*bool, bool)`

GetAutoLoanOk returns a tuple with the AutoLoan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLoan

`func (o *GetAccountConfigV5RespDataInner) SetAutoLoan(v bool)`

SetAutoLoan sets AutoLoan field to given value.

### HasAutoLoan

`func (o *GetAccountConfigV5RespDataInner) HasAutoLoan() bool`

HasAutoLoan returns a boolean if a field has been set.

### GetCtIsoMode

`func (o *GetAccountConfigV5RespDataInner) GetCtIsoMode() string`

GetCtIsoMode returns the CtIsoMode field if non-nil, zero value otherwise.

### GetCtIsoModeOk

`func (o *GetAccountConfigV5RespDataInner) GetCtIsoModeOk() (*string, bool)`

GetCtIsoModeOk returns a tuple with the CtIsoMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtIsoMode

`func (o *GetAccountConfigV5RespDataInner) SetCtIsoMode(v string)`

SetCtIsoMode sets CtIsoMode field to given value.

### HasCtIsoMode

`func (o *GetAccountConfigV5RespDataInner) HasCtIsoMode() bool`

HasCtIsoMode returns a boolean if a field has been set.

### GetEnableSpotBorrow

`func (o *GetAccountConfigV5RespDataInner) GetEnableSpotBorrow() bool`

GetEnableSpotBorrow returns the EnableSpotBorrow field if non-nil, zero value otherwise.

### GetEnableSpotBorrowOk

`func (o *GetAccountConfigV5RespDataInner) GetEnableSpotBorrowOk() (*bool, bool)`

GetEnableSpotBorrowOk returns a tuple with the EnableSpotBorrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSpotBorrow

`func (o *GetAccountConfigV5RespDataInner) SetEnableSpotBorrow(v bool)`

SetEnableSpotBorrow sets EnableSpotBorrow field to given value.

### HasEnableSpotBorrow

`func (o *GetAccountConfigV5RespDataInner) HasEnableSpotBorrow() bool`

HasEnableSpotBorrow returns a boolean if a field has been set.

### GetGreeksType

`func (o *GetAccountConfigV5RespDataInner) GetGreeksType() string`

GetGreeksType returns the GreeksType field if non-nil, zero value otherwise.

### GetGreeksTypeOk

`func (o *GetAccountConfigV5RespDataInner) GetGreeksTypeOk() (*string, bool)`

GetGreeksTypeOk returns a tuple with the GreeksType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreeksType

`func (o *GetAccountConfigV5RespDataInner) SetGreeksType(v string)`

SetGreeksType sets GreeksType field to given value.

### HasGreeksType

`func (o *GetAccountConfigV5RespDataInner) HasGreeksType() bool`

HasGreeksType returns a boolean if a field has been set.

### GetIp

`func (o *GetAccountConfigV5RespDataInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetAccountConfigV5RespDataInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetAccountConfigV5RespDataInner) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GetAccountConfigV5RespDataInner) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetKycLv

`func (o *GetAccountConfigV5RespDataInner) GetKycLv() string`

GetKycLv returns the KycLv field if non-nil, zero value otherwise.

### GetKycLvOk

`func (o *GetAccountConfigV5RespDataInner) GetKycLvOk() (*string, bool)`

GetKycLvOk returns a tuple with the KycLv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycLv

`func (o *GetAccountConfigV5RespDataInner) SetKycLv(v string)`

SetKycLv sets KycLv field to given value.

### HasKycLv

`func (o *GetAccountConfigV5RespDataInner) HasKycLv() bool`

HasKycLv returns a boolean if a field has been set.

### GetLabel

`func (o *GetAccountConfigV5RespDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetAccountConfigV5RespDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetAccountConfigV5RespDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetAccountConfigV5RespDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLevel

`func (o *GetAccountConfigV5RespDataInner) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *GetAccountConfigV5RespDataInner) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *GetAccountConfigV5RespDataInner) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *GetAccountConfigV5RespDataInner) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetLevelTmp

`func (o *GetAccountConfigV5RespDataInner) GetLevelTmp() string`

GetLevelTmp returns the LevelTmp field if non-nil, zero value otherwise.

### GetLevelTmpOk

`func (o *GetAccountConfigV5RespDataInner) GetLevelTmpOk() (*string, bool)`

GetLevelTmpOk returns a tuple with the LevelTmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelTmp

`func (o *GetAccountConfigV5RespDataInner) SetLevelTmp(v string)`

SetLevelTmp sets LevelTmp field to given value.

### HasLevelTmp

`func (o *GetAccountConfigV5RespDataInner) HasLevelTmp() bool`

HasLevelTmp returns a boolean if a field has been set.

### GetLiquidationGear

`func (o *GetAccountConfigV5RespDataInner) GetLiquidationGear() string`

GetLiquidationGear returns the LiquidationGear field if non-nil, zero value otherwise.

### GetLiquidationGearOk

`func (o *GetAccountConfigV5RespDataInner) GetLiquidationGearOk() (*string, bool)`

GetLiquidationGearOk returns a tuple with the LiquidationGear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidationGear

`func (o *GetAccountConfigV5RespDataInner) SetLiquidationGear(v string)`

SetLiquidationGear sets LiquidationGear field to given value.

### HasLiquidationGear

`func (o *GetAccountConfigV5RespDataInner) HasLiquidationGear() bool`

HasLiquidationGear returns a boolean if a field has been set.

### GetMainUid

`func (o *GetAccountConfigV5RespDataInner) GetMainUid() string`

GetMainUid returns the MainUid field if non-nil, zero value otherwise.

### GetMainUidOk

`func (o *GetAccountConfigV5RespDataInner) GetMainUidOk() (*string, bool)`

GetMainUidOk returns a tuple with the MainUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainUid

`func (o *GetAccountConfigV5RespDataInner) SetMainUid(v string)`

SetMainUid sets MainUid field to given value.

### HasMainUid

`func (o *GetAccountConfigV5RespDataInner) HasMainUid() bool`

HasMainUid returns a boolean if a field has been set.

### GetMgnIsoMode

`func (o *GetAccountConfigV5RespDataInner) GetMgnIsoMode() string`

GetMgnIsoMode returns the MgnIsoMode field if non-nil, zero value otherwise.

### GetMgnIsoModeOk

`func (o *GetAccountConfigV5RespDataInner) GetMgnIsoModeOk() (*string, bool)`

GetMgnIsoModeOk returns a tuple with the MgnIsoMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnIsoMode

`func (o *GetAccountConfigV5RespDataInner) SetMgnIsoMode(v string)`

SetMgnIsoMode sets MgnIsoMode field to given value.

### HasMgnIsoMode

`func (o *GetAccountConfigV5RespDataInner) HasMgnIsoMode() bool`

HasMgnIsoMode returns a boolean if a field has been set.

### GetOpAuth

`func (o *GetAccountConfigV5RespDataInner) GetOpAuth() string`

GetOpAuth returns the OpAuth field if non-nil, zero value otherwise.

### GetOpAuthOk

`func (o *GetAccountConfigV5RespDataInner) GetOpAuthOk() (*string, bool)`

GetOpAuthOk returns a tuple with the OpAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpAuth

`func (o *GetAccountConfigV5RespDataInner) SetOpAuth(v string)`

SetOpAuth sets OpAuth field to given value.

### HasOpAuth

`func (o *GetAccountConfigV5RespDataInner) HasOpAuth() bool`

HasOpAuth returns a boolean if a field has been set.

### GetPerm

`func (o *GetAccountConfigV5RespDataInner) GetPerm() string`

GetPerm returns the Perm field if non-nil, zero value otherwise.

### GetPermOk

`func (o *GetAccountConfigV5RespDataInner) GetPermOk() (*string, bool)`

GetPermOk returns a tuple with the Perm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerm

`func (o *GetAccountConfigV5RespDataInner) SetPerm(v string)`

SetPerm sets Perm field to given value.

### HasPerm

`func (o *GetAccountConfigV5RespDataInner) HasPerm() bool`

HasPerm returns a boolean if a field has been set.

### GetPosMode

`func (o *GetAccountConfigV5RespDataInner) GetPosMode() string`

GetPosMode returns the PosMode field if non-nil, zero value otherwise.

### GetPosModeOk

`func (o *GetAccountConfigV5RespDataInner) GetPosModeOk() (*string, bool)`

GetPosModeOk returns a tuple with the PosMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosMode

`func (o *GetAccountConfigV5RespDataInner) SetPosMode(v string)`

SetPosMode sets PosMode field to given value.

### HasPosMode

`func (o *GetAccountConfigV5RespDataInner) HasPosMode() bool`

HasPosMode returns a boolean if a field has been set.

### GetRoleType

`func (o *GetAccountConfigV5RespDataInner) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *GetAccountConfigV5RespDataInner) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *GetAccountConfigV5RespDataInner) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *GetAccountConfigV5RespDataInner) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetSpotBorrowAutoRepay

`func (o *GetAccountConfigV5RespDataInner) GetSpotBorrowAutoRepay() bool`

GetSpotBorrowAutoRepay returns the SpotBorrowAutoRepay field if non-nil, zero value otherwise.

### GetSpotBorrowAutoRepayOk

`func (o *GetAccountConfigV5RespDataInner) GetSpotBorrowAutoRepayOk() (*bool, bool)`

GetSpotBorrowAutoRepayOk returns a tuple with the SpotBorrowAutoRepay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotBorrowAutoRepay

`func (o *GetAccountConfigV5RespDataInner) SetSpotBorrowAutoRepay(v bool)`

SetSpotBorrowAutoRepay sets SpotBorrowAutoRepay field to given value.

### HasSpotBorrowAutoRepay

`func (o *GetAccountConfigV5RespDataInner) HasSpotBorrowAutoRepay() bool`

HasSpotBorrowAutoRepay returns a boolean if a field has been set.

### GetSpotOffsetType

`func (o *GetAccountConfigV5RespDataInner) GetSpotOffsetType() string`

GetSpotOffsetType returns the SpotOffsetType field if non-nil, zero value otherwise.

### GetSpotOffsetTypeOk

`func (o *GetAccountConfigV5RespDataInner) GetSpotOffsetTypeOk() (*string, bool)`

GetSpotOffsetTypeOk returns a tuple with the SpotOffsetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotOffsetType

`func (o *GetAccountConfigV5RespDataInner) SetSpotOffsetType(v string)`

SetSpotOffsetType sets SpotOffsetType field to given value.

### HasSpotOffsetType

`func (o *GetAccountConfigV5RespDataInner) HasSpotOffsetType() bool`

HasSpotOffsetType returns a boolean if a field has been set.

### GetSpotRoleType

`func (o *GetAccountConfigV5RespDataInner) GetSpotRoleType() string`

GetSpotRoleType returns the SpotRoleType field if non-nil, zero value otherwise.

### GetSpotRoleTypeOk

`func (o *GetAccountConfigV5RespDataInner) GetSpotRoleTypeOk() (*string, bool)`

GetSpotRoleTypeOk returns a tuple with the SpotRoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotRoleType

`func (o *GetAccountConfigV5RespDataInner) SetSpotRoleType(v string)`

SetSpotRoleType sets SpotRoleType field to given value.

### HasSpotRoleType

`func (o *GetAccountConfigV5RespDataInner) HasSpotRoleType() bool`

HasSpotRoleType returns a boolean if a field has been set.

### GetSpotTraderInsts

`func (o *GetAccountConfigV5RespDataInner) GetSpotTraderInsts() []string`

GetSpotTraderInsts returns the SpotTraderInsts field if non-nil, zero value otherwise.

### GetSpotTraderInstsOk

`func (o *GetAccountConfigV5RespDataInner) GetSpotTraderInstsOk() (*[]string, bool)`

GetSpotTraderInstsOk returns a tuple with the SpotTraderInsts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotTraderInsts

`func (o *GetAccountConfigV5RespDataInner) SetSpotTraderInsts(v []string)`

SetSpotTraderInsts sets SpotTraderInsts field to given value.

### HasSpotTraderInsts

`func (o *GetAccountConfigV5RespDataInner) HasSpotTraderInsts() bool`

HasSpotTraderInsts returns a boolean if a field has been set.

### GetTraderInsts

`func (o *GetAccountConfigV5RespDataInner) GetTraderInsts() []string`

GetTraderInsts returns the TraderInsts field if non-nil, zero value otherwise.

### GetTraderInstsOk

`func (o *GetAccountConfigV5RespDataInner) GetTraderInstsOk() (*[]string, bool)`

GetTraderInstsOk returns a tuple with the TraderInsts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraderInsts

`func (o *GetAccountConfigV5RespDataInner) SetTraderInsts(v []string)`

SetTraderInsts sets TraderInsts field to given value.

### HasTraderInsts

`func (o *GetAccountConfigV5RespDataInner) HasTraderInsts() bool`

HasTraderInsts returns a boolean if a field has been set.

### GetType

`func (o *GetAccountConfigV5RespDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAccountConfigV5RespDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAccountConfigV5RespDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAccountConfigV5RespDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUid

`func (o *GetAccountConfigV5RespDataInner) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *GetAccountConfigV5RespDataInner) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *GetAccountConfigV5RespDataInner) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *GetAccountConfigV5RespDataInner) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


