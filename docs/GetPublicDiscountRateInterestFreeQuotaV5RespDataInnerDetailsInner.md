# GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisCcyEq** | Pointer to **string** | Discount equity in currency for quick calculation if your equity is the&#x60;maxAmt&#x60; | [optional] [default to ""]
**DiscountRate** | Pointer to **string** | Discount rate | [optional] [default to ""]
**LiqPenaltyRate** | Pointer to **string** | Liquidation penalty rate | [optional] [default to ""]
**MaxAmt** | Pointer to **string** | Tier - upper bound.   The unit is the currency like BTC. \&quot;\&quot; means positive infinity | [optional] [default to ""]
**MinAmt** | Pointer to **string** | Tier - lower bound.   The unit is the currency like BTC. The minimum is 0 | [optional] [default to ""]
**Tier** | Pointer to **string** | Tiers | [optional] [default to ""]

## Methods

### NewGetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner

`func NewGetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner() *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner`

NewGetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner instantiates a new GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInnerWithDefaults

`func NewGetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInnerWithDefaults() *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner`

NewGetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInnerWithDefaults instantiates a new GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisCcyEq

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) GetDisCcyEq() string`

GetDisCcyEq returns the DisCcyEq field if non-nil, zero value otherwise.

### GetDisCcyEqOk

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) GetDisCcyEqOk() (*string, bool)`

GetDisCcyEqOk returns a tuple with the DisCcyEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisCcyEq

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) SetDisCcyEq(v string)`

SetDisCcyEq sets DisCcyEq field to given value.

### HasDisCcyEq

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) HasDisCcyEq() bool`

HasDisCcyEq returns a boolean if a field has been set.

### GetDiscountRate

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) GetDiscountRate() string`

GetDiscountRate returns the DiscountRate field if non-nil, zero value otherwise.

### GetDiscountRateOk

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) GetDiscountRateOk() (*string, bool)`

GetDiscountRateOk returns a tuple with the DiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRate

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) SetDiscountRate(v string)`

SetDiscountRate sets DiscountRate field to given value.

### HasDiscountRate

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) HasDiscountRate() bool`

HasDiscountRate returns a boolean if a field has been set.

### GetLiqPenaltyRate

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) GetLiqPenaltyRate() string`

GetLiqPenaltyRate returns the LiqPenaltyRate field if non-nil, zero value otherwise.

### GetLiqPenaltyRateOk

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) GetLiqPenaltyRateOk() (*string, bool)`

GetLiqPenaltyRateOk returns a tuple with the LiqPenaltyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiqPenaltyRate

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) SetLiqPenaltyRate(v string)`

SetLiqPenaltyRate sets LiqPenaltyRate field to given value.

### HasLiqPenaltyRate

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) HasLiqPenaltyRate() bool`

HasLiqPenaltyRate returns a boolean if a field has been set.

### GetMaxAmt

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) GetMaxAmt() string`

GetMaxAmt returns the MaxAmt field if non-nil, zero value otherwise.

### GetMaxAmtOk

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) GetMaxAmtOk() (*string, bool)`

GetMaxAmtOk returns a tuple with the MaxAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmt

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) SetMaxAmt(v string)`

SetMaxAmt sets MaxAmt field to given value.

### HasMaxAmt

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) HasMaxAmt() bool`

HasMaxAmt returns a boolean if a field has been set.

### GetMinAmt

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) GetMinAmt() string`

GetMinAmt returns the MinAmt field if non-nil, zero value otherwise.

### GetMinAmtOk

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) GetMinAmtOk() (*string, bool)`

GetMinAmtOk returns a tuple with the MinAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmt

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) SetMinAmt(v string)`

SetMinAmt sets MinAmt field to given value.

### HasMinAmt

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) HasMinAmt() bool`

HasMinAmt returns a boolean if a field has been set.

### GetTier

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) HasTier() bool`

HasTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


