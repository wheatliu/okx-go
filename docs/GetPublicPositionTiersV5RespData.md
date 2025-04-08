# GetPublicPositionTiersV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseMaxLoan** | Pointer to **string** | Base currency borrowing amount (only applicable to leverage and the case when &#x60;instId&#x60; takes effect) | [optional] [default to ""]
**Imr** | Pointer to **string** | Initial margin requirement rate | [optional] [default to ""]
**InstFamily** | Pointer to **string** | Instrument family  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**MaxLever** | Pointer to **string** | Maximum available leverage | [optional] [default to ""]
**MaxSz** | Pointer to **string** | The maximum borrowing amount or number of positions held in this position is only applicable to margin/options/perpetual/delivery  It will return the maximum borrowing amount when &#x60;ccy&#x60; takes effect. | [optional] [default to ""]
**MinSz** | Pointer to **string** | The minimum borrowing amount or position of this gear is only applicable to margin/options/perpetual/delivery, the minimum position is 0 by default  It will return the minimum borrowing amount when &#x60;ccy&#x60; takes effect. | [optional] [default to ""]
**Mmr** | Pointer to **string** | Maintenance margin requirement rate | [optional] [default to ""]
**OptMgnFactor** | Pointer to **string** | Option Margin Coefficient (only applicable to options) | [optional] [default to ""]
**QuoteMaxLoan** | Pointer to **string** | Quote currency borrowing amount (only applicable to leverage and the case when &#x60;instId&#x60; takes effect) | [optional] [default to ""]
**Tier** | Pointer to **string** | Tiers | [optional] [default to ""]
**Uly** | Pointer to **string** | Underlying  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [optional] [default to ""]

## Methods

### NewGetPublicPositionTiersV5RespData

`func NewGetPublicPositionTiersV5RespData() *GetPublicPositionTiersV5RespData`

NewGetPublicPositionTiersV5RespData instantiates a new GetPublicPositionTiersV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicPositionTiersV5RespDataWithDefaults

`func NewGetPublicPositionTiersV5RespDataWithDefaults() *GetPublicPositionTiersV5RespData`

NewGetPublicPositionTiersV5RespDataWithDefaults instantiates a new GetPublicPositionTiersV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseMaxLoan

`func (o *GetPublicPositionTiersV5RespData) GetBaseMaxLoan() string`

GetBaseMaxLoan returns the BaseMaxLoan field if non-nil, zero value otherwise.

### GetBaseMaxLoanOk

`func (o *GetPublicPositionTiersV5RespData) GetBaseMaxLoanOk() (*string, bool)`

GetBaseMaxLoanOk returns a tuple with the BaseMaxLoan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseMaxLoan

`func (o *GetPublicPositionTiersV5RespData) SetBaseMaxLoan(v string)`

SetBaseMaxLoan sets BaseMaxLoan field to given value.

### HasBaseMaxLoan

`func (o *GetPublicPositionTiersV5RespData) HasBaseMaxLoan() bool`

HasBaseMaxLoan returns a boolean if a field has been set.

### GetImr

`func (o *GetPublicPositionTiersV5RespData) GetImr() string`

GetImr returns the Imr field if non-nil, zero value otherwise.

### GetImrOk

`func (o *GetPublicPositionTiersV5RespData) GetImrOk() (*string, bool)`

GetImrOk returns a tuple with the Imr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImr

`func (o *GetPublicPositionTiersV5RespData) SetImr(v string)`

SetImr sets Imr field to given value.

### HasImr

`func (o *GetPublicPositionTiersV5RespData) HasImr() bool`

HasImr returns a boolean if a field has been set.

### GetInstFamily

`func (o *GetPublicPositionTiersV5RespData) GetInstFamily() string`

GetInstFamily returns the InstFamily field if non-nil, zero value otherwise.

### GetInstFamilyOk

`func (o *GetPublicPositionTiersV5RespData) GetInstFamilyOk() (*string, bool)`

GetInstFamilyOk returns a tuple with the InstFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstFamily

`func (o *GetPublicPositionTiersV5RespData) SetInstFamily(v string)`

SetInstFamily sets InstFamily field to given value.

### HasInstFamily

`func (o *GetPublicPositionTiersV5RespData) HasInstFamily() bool`

HasInstFamily returns a boolean if a field has been set.

### GetInstId

`func (o *GetPublicPositionTiersV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetPublicPositionTiersV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetPublicPositionTiersV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetPublicPositionTiersV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetMaxLever

`func (o *GetPublicPositionTiersV5RespData) GetMaxLever() string`

GetMaxLever returns the MaxLever field if non-nil, zero value otherwise.

### GetMaxLeverOk

`func (o *GetPublicPositionTiersV5RespData) GetMaxLeverOk() (*string, bool)`

GetMaxLeverOk returns a tuple with the MaxLever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLever

`func (o *GetPublicPositionTiersV5RespData) SetMaxLever(v string)`

SetMaxLever sets MaxLever field to given value.

### HasMaxLever

`func (o *GetPublicPositionTiersV5RespData) HasMaxLever() bool`

HasMaxLever returns a boolean if a field has been set.

### GetMaxSz

`func (o *GetPublicPositionTiersV5RespData) GetMaxSz() string`

GetMaxSz returns the MaxSz field if non-nil, zero value otherwise.

### GetMaxSzOk

`func (o *GetPublicPositionTiersV5RespData) GetMaxSzOk() (*string, bool)`

GetMaxSzOk returns a tuple with the MaxSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSz

`func (o *GetPublicPositionTiersV5RespData) SetMaxSz(v string)`

SetMaxSz sets MaxSz field to given value.

### HasMaxSz

`func (o *GetPublicPositionTiersV5RespData) HasMaxSz() bool`

HasMaxSz returns a boolean if a field has been set.

### GetMinSz

`func (o *GetPublicPositionTiersV5RespData) GetMinSz() string`

GetMinSz returns the MinSz field if non-nil, zero value otherwise.

### GetMinSzOk

`func (o *GetPublicPositionTiersV5RespData) GetMinSzOk() (*string, bool)`

GetMinSzOk returns a tuple with the MinSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSz

`func (o *GetPublicPositionTiersV5RespData) SetMinSz(v string)`

SetMinSz sets MinSz field to given value.

### HasMinSz

`func (o *GetPublicPositionTiersV5RespData) HasMinSz() bool`

HasMinSz returns a boolean if a field has been set.

### GetMmr

`func (o *GetPublicPositionTiersV5RespData) GetMmr() string`

GetMmr returns the Mmr field if non-nil, zero value otherwise.

### GetMmrOk

`func (o *GetPublicPositionTiersV5RespData) GetMmrOk() (*string, bool)`

GetMmrOk returns a tuple with the Mmr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmr

`func (o *GetPublicPositionTiersV5RespData) SetMmr(v string)`

SetMmr sets Mmr field to given value.

### HasMmr

`func (o *GetPublicPositionTiersV5RespData) HasMmr() bool`

HasMmr returns a boolean if a field has been set.

### GetOptMgnFactor

`func (o *GetPublicPositionTiersV5RespData) GetOptMgnFactor() string`

GetOptMgnFactor returns the OptMgnFactor field if non-nil, zero value otherwise.

### GetOptMgnFactorOk

`func (o *GetPublicPositionTiersV5RespData) GetOptMgnFactorOk() (*string, bool)`

GetOptMgnFactorOk returns a tuple with the OptMgnFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptMgnFactor

`func (o *GetPublicPositionTiersV5RespData) SetOptMgnFactor(v string)`

SetOptMgnFactor sets OptMgnFactor field to given value.

### HasOptMgnFactor

`func (o *GetPublicPositionTiersV5RespData) HasOptMgnFactor() bool`

HasOptMgnFactor returns a boolean if a field has been set.

### GetQuoteMaxLoan

`func (o *GetPublicPositionTiersV5RespData) GetQuoteMaxLoan() string`

GetQuoteMaxLoan returns the QuoteMaxLoan field if non-nil, zero value otherwise.

### GetQuoteMaxLoanOk

`func (o *GetPublicPositionTiersV5RespData) GetQuoteMaxLoanOk() (*string, bool)`

GetQuoteMaxLoanOk returns a tuple with the QuoteMaxLoan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteMaxLoan

`func (o *GetPublicPositionTiersV5RespData) SetQuoteMaxLoan(v string)`

SetQuoteMaxLoan sets QuoteMaxLoan field to given value.

### HasQuoteMaxLoan

`func (o *GetPublicPositionTiersV5RespData) HasQuoteMaxLoan() bool`

HasQuoteMaxLoan returns a boolean if a field has been set.

### GetTier

`func (o *GetPublicPositionTiersV5RespData) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *GetPublicPositionTiersV5RespData) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *GetPublicPositionTiersV5RespData) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *GetPublicPositionTiersV5RespData) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetUly

`func (o *GetPublicPositionTiersV5RespData) GetUly() string`

GetUly returns the Uly field if non-nil, zero value otherwise.

### GetUlyOk

`func (o *GetPublicPositionTiersV5RespData) GetUlyOk() (*string, bool)`

GetUlyOk returns a tuple with the Uly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUly

`func (o *GetPublicPositionTiersV5RespData) SetUly(v string)`

SetUly sets Uly field to given value.

### HasUly

`func (o *GetPublicPositionTiersV5RespData) HasUly() bool`

HasUly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


