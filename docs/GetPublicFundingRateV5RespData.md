# GetPublicFundingRateV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormulaType** | Pointer to **string** | Formula type  &#x60;noRate&#x60;: old funding rate formula  &#x60;withRate&#x60;: new funding rate formula | [optional] [default to ""]
**FundingRate** | Pointer to **string** | Current funding rate | [optional] [default to ""]
**FundingTime** | Pointer to **string** | Settlement time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**ImpactValue** | Pointer to **string** | Depth weighted amount (in the unit of quote currency) | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. &#x60;BTC-USD-SWAP&#x60; | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type  &#x60;SWAP&#x60; | [optional] [default to ""]
**InterestRate** | Pointer to **string** | Interest rate | [optional] [default to ""]
**MaxFundingRate** | Pointer to **string** | The upper limit of the predicted funding rate of the next cycle | [optional] [default to ""]
**Method** | Pointer to **string** | Funding rate mechanism   &#x60;current_period&#x60;   &#x60;next_period&#x60;(no longer supported) | [optional] [default to ""]
**MinFundingRate** | Pointer to **string** | The lower limit of the predicted funding rate of the next cycle | [optional] [default to ""]
**NextFundingRate** | Pointer to **string** | Forecasted funding rate for the next period   The nextFundingRate will be \&quot;\&quot; if the method is &#x60;current_period&#x60;(no longer supported) | [optional] [default to ""]
**NextFundingTime** | Pointer to **string** | Forecasted funding time for the next period , Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Premium** | Pointer to **string** | Premium between the mid price of perps market and the index price | [optional] [default to ""]
**SettFundingRate** | Pointer to **string** | If settState &#x3D; &#x60;processing&#x60;, it is the funding rate that is being used for current settlement cycle.   If settState &#x3D; &#x60;settled&#x60;, it is the funding rate that is being used for previous settlement cycle | [optional] [default to ""]
**SettState** | Pointer to **string** | Settlement state of funding rate   &#x60;processing&#x60;   &#x60;settled&#x60; | [optional] [default to ""]
**Ts** | Pointer to **string** | Data return time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetPublicFundingRateV5RespData

`func NewGetPublicFundingRateV5RespData() *GetPublicFundingRateV5RespData`

NewGetPublicFundingRateV5RespData instantiates a new GetPublicFundingRateV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicFundingRateV5RespDataWithDefaults

`func NewGetPublicFundingRateV5RespDataWithDefaults() *GetPublicFundingRateV5RespData`

NewGetPublicFundingRateV5RespDataWithDefaults instantiates a new GetPublicFundingRateV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormulaType

`func (o *GetPublicFundingRateV5RespData) GetFormulaType() string`

GetFormulaType returns the FormulaType field if non-nil, zero value otherwise.

### GetFormulaTypeOk

`func (o *GetPublicFundingRateV5RespData) GetFormulaTypeOk() (*string, bool)`

GetFormulaTypeOk returns a tuple with the FormulaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulaType

`func (o *GetPublicFundingRateV5RespData) SetFormulaType(v string)`

SetFormulaType sets FormulaType field to given value.

### HasFormulaType

`func (o *GetPublicFundingRateV5RespData) HasFormulaType() bool`

HasFormulaType returns a boolean if a field has been set.

### GetFundingRate

`func (o *GetPublicFundingRateV5RespData) GetFundingRate() string`

GetFundingRate returns the FundingRate field if non-nil, zero value otherwise.

### GetFundingRateOk

`func (o *GetPublicFundingRateV5RespData) GetFundingRateOk() (*string, bool)`

GetFundingRateOk returns a tuple with the FundingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingRate

`func (o *GetPublicFundingRateV5RespData) SetFundingRate(v string)`

SetFundingRate sets FundingRate field to given value.

### HasFundingRate

`func (o *GetPublicFundingRateV5RespData) HasFundingRate() bool`

HasFundingRate returns a boolean if a field has been set.

### GetFundingTime

`func (o *GetPublicFundingRateV5RespData) GetFundingTime() string`

GetFundingTime returns the FundingTime field if non-nil, zero value otherwise.

### GetFundingTimeOk

`func (o *GetPublicFundingRateV5RespData) GetFundingTimeOk() (*string, bool)`

GetFundingTimeOk returns a tuple with the FundingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingTime

`func (o *GetPublicFundingRateV5RespData) SetFundingTime(v string)`

SetFundingTime sets FundingTime field to given value.

### HasFundingTime

`func (o *GetPublicFundingRateV5RespData) HasFundingTime() bool`

HasFundingTime returns a boolean if a field has been set.

### GetImpactValue

`func (o *GetPublicFundingRateV5RespData) GetImpactValue() string`

GetImpactValue returns the ImpactValue field if non-nil, zero value otherwise.

### GetImpactValueOk

`func (o *GetPublicFundingRateV5RespData) GetImpactValueOk() (*string, bool)`

GetImpactValueOk returns a tuple with the ImpactValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactValue

`func (o *GetPublicFundingRateV5RespData) SetImpactValue(v string)`

SetImpactValue sets ImpactValue field to given value.

### HasImpactValue

`func (o *GetPublicFundingRateV5RespData) HasImpactValue() bool`

HasImpactValue returns a boolean if a field has been set.

### GetInstId

`func (o *GetPublicFundingRateV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetPublicFundingRateV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetPublicFundingRateV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetPublicFundingRateV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetPublicFundingRateV5RespData) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetPublicFundingRateV5RespData) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetPublicFundingRateV5RespData) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetPublicFundingRateV5RespData) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetInterestRate

`func (o *GetPublicFundingRateV5RespData) GetInterestRate() string`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *GetPublicFundingRateV5RespData) GetInterestRateOk() (*string, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *GetPublicFundingRateV5RespData) SetInterestRate(v string)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *GetPublicFundingRateV5RespData) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetMaxFundingRate

`func (o *GetPublicFundingRateV5RespData) GetMaxFundingRate() string`

GetMaxFundingRate returns the MaxFundingRate field if non-nil, zero value otherwise.

### GetMaxFundingRateOk

`func (o *GetPublicFundingRateV5RespData) GetMaxFundingRateOk() (*string, bool)`

GetMaxFundingRateOk returns a tuple with the MaxFundingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFundingRate

`func (o *GetPublicFundingRateV5RespData) SetMaxFundingRate(v string)`

SetMaxFundingRate sets MaxFundingRate field to given value.

### HasMaxFundingRate

`func (o *GetPublicFundingRateV5RespData) HasMaxFundingRate() bool`

HasMaxFundingRate returns a boolean if a field has been set.

### GetMethod

`func (o *GetPublicFundingRateV5RespData) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *GetPublicFundingRateV5RespData) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *GetPublicFundingRateV5RespData) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *GetPublicFundingRateV5RespData) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMinFundingRate

`func (o *GetPublicFundingRateV5RespData) GetMinFundingRate() string`

GetMinFundingRate returns the MinFundingRate field if non-nil, zero value otherwise.

### GetMinFundingRateOk

`func (o *GetPublicFundingRateV5RespData) GetMinFundingRateOk() (*string, bool)`

GetMinFundingRateOk returns a tuple with the MinFundingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFundingRate

`func (o *GetPublicFundingRateV5RespData) SetMinFundingRate(v string)`

SetMinFundingRate sets MinFundingRate field to given value.

### HasMinFundingRate

`func (o *GetPublicFundingRateV5RespData) HasMinFundingRate() bool`

HasMinFundingRate returns a boolean if a field has been set.

### GetNextFundingRate

`func (o *GetPublicFundingRateV5RespData) GetNextFundingRate() string`

GetNextFundingRate returns the NextFundingRate field if non-nil, zero value otherwise.

### GetNextFundingRateOk

`func (o *GetPublicFundingRateV5RespData) GetNextFundingRateOk() (*string, bool)`

GetNextFundingRateOk returns a tuple with the NextFundingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFundingRate

`func (o *GetPublicFundingRateV5RespData) SetNextFundingRate(v string)`

SetNextFundingRate sets NextFundingRate field to given value.

### HasNextFundingRate

`func (o *GetPublicFundingRateV5RespData) HasNextFundingRate() bool`

HasNextFundingRate returns a boolean if a field has been set.

### GetNextFundingTime

`func (o *GetPublicFundingRateV5RespData) GetNextFundingTime() string`

GetNextFundingTime returns the NextFundingTime field if non-nil, zero value otherwise.

### GetNextFundingTimeOk

`func (o *GetPublicFundingRateV5RespData) GetNextFundingTimeOk() (*string, bool)`

GetNextFundingTimeOk returns a tuple with the NextFundingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFundingTime

`func (o *GetPublicFundingRateV5RespData) SetNextFundingTime(v string)`

SetNextFundingTime sets NextFundingTime field to given value.

### HasNextFundingTime

`func (o *GetPublicFundingRateV5RespData) HasNextFundingTime() bool`

HasNextFundingTime returns a boolean if a field has been set.

### GetPremium

`func (o *GetPublicFundingRateV5RespData) GetPremium() string`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *GetPublicFundingRateV5RespData) GetPremiumOk() (*string, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *GetPublicFundingRateV5RespData) SetPremium(v string)`

SetPremium sets Premium field to given value.

### HasPremium

`func (o *GetPublicFundingRateV5RespData) HasPremium() bool`

HasPremium returns a boolean if a field has been set.

### GetSettFundingRate

`func (o *GetPublicFundingRateV5RespData) GetSettFundingRate() string`

GetSettFundingRate returns the SettFundingRate field if non-nil, zero value otherwise.

### GetSettFundingRateOk

`func (o *GetPublicFundingRateV5RespData) GetSettFundingRateOk() (*string, bool)`

GetSettFundingRateOk returns a tuple with the SettFundingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettFundingRate

`func (o *GetPublicFundingRateV5RespData) SetSettFundingRate(v string)`

SetSettFundingRate sets SettFundingRate field to given value.

### HasSettFundingRate

`func (o *GetPublicFundingRateV5RespData) HasSettFundingRate() bool`

HasSettFundingRate returns a boolean if a field has been set.

### GetSettState

`func (o *GetPublicFundingRateV5RespData) GetSettState() string`

GetSettState returns the SettState field if non-nil, zero value otherwise.

### GetSettStateOk

`func (o *GetPublicFundingRateV5RespData) GetSettStateOk() (*string, bool)`

GetSettStateOk returns a tuple with the SettState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettState

`func (o *GetPublicFundingRateV5RespData) SetSettState(v string)`

SetSettState sets SettState field to given value.

### HasSettState

`func (o *GetPublicFundingRateV5RespData) HasSettState() bool`

HasSettState returns a boolean if a field has been set.

### GetTs

`func (o *GetPublicFundingRateV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetPublicFundingRateV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetPublicFundingRateV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetPublicFundingRateV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


