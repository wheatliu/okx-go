# GetAffiliateInviteeDetailV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccFee** | Pointer to **string** | Accumulated Amount of trading fee in USDT  If there is no any fee, 0 will be returned | [optional] [default to ""]
**AffiliateCode** | Pointer to **string** | Affiliate invite code that the invitee registered/recalled via | [optional] [default to ""]
**DepAmt** | Pointer to **string** | Accumulated amount of deposit in USDT  If user has not deposited, 0 will be returned | [optional] [default to ""]
**FirstTradeTime** | Pointer to **string** | Timestamp that the first trade is completed after the latest rebate relationship is established with the parent affiliate   Unix timestamp in millisecond format, e.g. 1597026383085  If user has not traded, \&quot;\&quot; will be returned | [optional] [default to ""]
**InviteeLevel** | Pointer to **string** | Invitee&#39;s relative level to the affiliate  If the user is a invitee, the level will be &#x60;2&#x60;. | [optional] [default to ""]
**InviteeRebateRate** | Pointer to **string** | Self rebate rate of the invitee (in decimal), e.g. &#x60;0.01&#x60; represents &#x60;10%&#x60; | [optional] [default to ""]
**JoinTime** | Pointer to **string** | Timestamp that the rebate relationship is established, Unix timestamp in millisecond format, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**KycTime** | Pointer to **string** | KYC2 verification time. Unix timestamp in millisecond format and the precision is in day  If user has not passed KYC2, \&quot;\&quot; will be returned | [optional] [default to ""]
**Level** | Pointer to **string** | Invitee trading fee level, e.g. Lv1 | [optional] [default to ""]
**Region** | Pointer to **string** | User country or region. e.g. \&quot;United Kingdom\&quot; | [optional] [default to ""]
**TotalCommission** | Pointer to **string** | Total commission earned from the invitee, unit in &#x60;USDT&#x60; | [optional] [default to ""]
**VolMonth** | Pointer to **string** | Accumulated Trading volume in the current month in USDT  If user has not traded, 0 will be returned | [optional] [default to ""]

## Methods

### NewGetAffiliateInviteeDetailV5RespDataInner

`func NewGetAffiliateInviteeDetailV5RespDataInner() *GetAffiliateInviteeDetailV5RespDataInner`

NewGetAffiliateInviteeDetailV5RespDataInner instantiates a new GetAffiliateInviteeDetailV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAffiliateInviteeDetailV5RespDataInnerWithDefaults

`func NewGetAffiliateInviteeDetailV5RespDataInnerWithDefaults() *GetAffiliateInviteeDetailV5RespDataInner`

NewGetAffiliateInviteeDetailV5RespDataInnerWithDefaults instantiates a new GetAffiliateInviteeDetailV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccFee

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetAccFee() string`

GetAccFee returns the AccFee field if non-nil, zero value otherwise.

### GetAccFeeOk

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetAccFeeOk() (*string, bool)`

GetAccFeeOk returns a tuple with the AccFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccFee

`func (o *GetAffiliateInviteeDetailV5RespDataInner) SetAccFee(v string)`

SetAccFee sets AccFee field to given value.

### HasAccFee

`func (o *GetAffiliateInviteeDetailV5RespDataInner) HasAccFee() bool`

HasAccFee returns a boolean if a field has been set.

### GetAffiliateCode

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetAffiliateCode() string`

GetAffiliateCode returns the AffiliateCode field if non-nil, zero value otherwise.

### GetAffiliateCodeOk

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetAffiliateCodeOk() (*string, bool)`

GetAffiliateCodeOk returns a tuple with the AffiliateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateCode

`func (o *GetAffiliateInviteeDetailV5RespDataInner) SetAffiliateCode(v string)`

SetAffiliateCode sets AffiliateCode field to given value.

### HasAffiliateCode

`func (o *GetAffiliateInviteeDetailV5RespDataInner) HasAffiliateCode() bool`

HasAffiliateCode returns a boolean if a field has been set.

### GetDepAmt

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetDepAmt() string`

GetDepAmt returns the DepAmt field if non-nil, zero value otherwise.

### GetDepAmtOk

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetDepAmtOk() (*string, bool)`

GetDepAmtOk returns a tuple with the DepAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepAmt

`func (o *GetAffiliateInviteeDetailV5RespDataInner) SetDepAmt(v string)`

SetDepAmt sets DepAmt field to given value.

### HasDepAmt

`func (o *GetAffiliateInviteeDetailV5RespDataInner) HasDepAmt() bool`

HasDepAmt returns a boolean if a field has been set.

### GetFirstTradeTime

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetFirstTradeTime() string`

GetFirstTradeTime returns the FirstTradeTime field if non-nil, zero value otherwise.

### GetFirstTradeTimeOk

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetFirstTradeTimeOk() (*string, bool)`

GetFirstTradeTimeOk returns a tuple with the FirstTradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTradeTime

`func (o *GetAffiliateInviteeDetailV5RespDataInner) SetFirstTradeTime(v string)`

SetFirstTradeTime sets FirstTradeTime field to given value.

### HasFirstTradeTime

`func (o *GetAffiliateInviteeDetailV5RespDataInner) HasFirstTradeTime() bool`

HasFirstTradeTime returns a boolean if a field has been set.

### GetInviteeLevel

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetInviteeLevel() string`

GetInviteeLevel returns the InviteeLevel field if non-nil, zero value otherwise.

### GetInviteeLevelOk

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetInviteeLevelOk() (*string, bool)`

GetInviteeLevelOk returns a tuple with the InviteeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteeLevel

`func (o *GetAffiliateInviteeDetailV5RespDataInner) SetInviteeLevel(v string)`

SetInviteeLevel sets InviteeLevel field to given value.

### HasInviteeLevel

`func (o *GetAffiliateInviteeDetailV5RespDataInner) HasInviteeLevel() bool`

HasInviteeLevel returns a boolean if a field has been set.

### GetInviteeRebateRate

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetInviteeRebateRate() string`

GetInviteeRebateRate returns the InviteeRebateRate field if non-nil, zero value otherwise.

### GetInviteeRebateRateOk

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetInviteeRebateRateOk() (*string, bool)`

GetInviteeRebateRateOk returns a tuple with the InviteeRebateRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteeRebateRate

`func (o *GetAffiliateInviteeDetailV5RespDataInner) SetInviteeRebateRate(v string)`

SetInviteeRebateRate sets InviteeRebateRate field to given value.

### HasInviteeRebateRate

`func (o *GetAffiliateInviteeDetailV5RespDataInner) HasInviteeRebateRate() bool`

HasInviteeRebateRate returns a boolean if a field has been set.

### GetJoinTime

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetJoinTime() string`

GetJoinTime returns the JoinTime field if non-nil, zero value otherwise.

### GetJoinTimeOk

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetJoinTimeOk() (*string, bool)`

GetJoinTimeOk returns a tuple with the JoinTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinTime

`func (o *GetAffiliateInviteeDetailV5RespDataInner) SetJoinTime(v string)`

SetJoinTime sets JoinTime field to given value.

### HasJoinTime

`func (o *GetAffiliateInviteeDetailV5RespDataInner) HasJoinTime() bool`

HasJoinTime returns a boolean if a field has been set.

### GetKycTime

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetKycTime() string`

GetKycTime returns the KycTime field if non-nil, zero value otherwise.

### GetKycTimeOk

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetKycTimeOk() (*string, bool)`

GetKycTimeOk returns a tuple with the KycTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycTime

`func (o *GetAffiliateInviteeDetailV5RespDataInner) SetKycTime(v string)`

SetKycTime sets KycTime field to given value.

### HasKycTime

`func (o *GetAffiliateInviteeDetailV5RespDataInner) HasKycTime() bool`

HasKycTime returns a boolean if a field has been set.

### GetLevel

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *GetAffiliateInviteeDetailV5RespDataInner) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *GetAffiliateInviteeDetailV5RespDataInner) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetRegion

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetAffiliateInviteeDetailV5RespDataInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetAffiliateInviteeDetailV5RespDataInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetTotalCommission

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetTotalCommission() string`

GetTotalCommission returns the TotalCommission field if non-nil, zero value otherwise.

### GetTotalCommissionOk

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetTotalCommissionOk() (*string, bool)`

GetTotalCommissionOk returns a tuple with the TotalCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCommission

`func (o *GetAffiliateInviteeDetailV5RespDataInner) SetTotalCommission(v string)`

SetTotalCommission sets TotalCommission field to given value.

### HasTotalCommission

`func (o *GetAffiliateInviteeDetailV5RespDataInner) HasTotalCommission() bool`

HasTotalCommission returns a boolean if a field has been set.

### GetVolMonth

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetVolMonth() string`

GetVolMonth returns the VolMonth field if non-nil, zero value otherwise.

### GetVolMonthOk

`func (o *GetAffiliateInviteeDetailV5RespDataInner) GetVolMonthOk() (*string, bool)`

GetVolMonthOk returns a tuple with the VolMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolMonth

`func (o *GetAffiliateInviteeDetailV5RespDataInner) SetVolMonth(v string)`

SetVolMonth sets VolMonth field to given value.

### HasVolMonth

`func (o *GetAffiliateInviteeDetailV5RespDataInner) HasVolMonth() bool`

HasVolMonth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


