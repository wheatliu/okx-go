# GetPublicFundingRateHistoryV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormulaType** | Pointer to **string** | Formula type  &#x60;noRate&#x60;: old funding rate formula  &#x60;withRate&#x60;: new funding rate formula | [optional] [default to ""]
**FundingRate** | Pointer to **string** | Predicted funding rate | [optional] [default to ""]
**FundingTime** | Pointer to **string** | Settlement time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. &#x60;BTC-USD-SWAP&#x60; | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type  &#x60;SWAP&#x60; | [optional] [default to ""]
**Method** | Pointer to **string** | Funding rate mechanism   &#x60;current_period&#x60;   &#x60;next_period&#x60; | [optional] [default to ""]
**RealizedRate** | Pointer to **string** | Actual funding rate | [optional] [default to ""]

## Methods

### NewGetPublicFundingRateHistoryV5RespData

`func NewGetPublicFundingRateHistoryV5RespData() *GetPublicFundingRateHistoryV5RespData`

NewGetPublicFundingRateHistoryV5RespData instantiates a new GetPublicFundingRateHistoryV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicFundingRateHistoryV5RespDataWithDefaults

`func NewGetPublicFundingRateHistoryV5RespDataWithDefaults() *GetPublicFundingRateHistoryV5RespData`

NewGetPublicFundingRateHistoryV5RespDataWithDefaults instantiates a new GetPublicFundingRateHistoryV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormulaType

`func (o *GetPublicFundingRateHistoryV5RespData) GetFormulaType() string`

GetFormulaType returns the FormulaType field if non-nil, zero value otherwise.

### GetFormulaTypeOk

`func (o *GetPublicFundingRateHistoryV5RespData) GetFormulaTypeOk() (*string, bool)`

GetFormulaTypeOk returns a tuple with the FormulaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulaType

`func (o *GetPublicFundingRateHistoryV5RespData) SetFormulaType(v string)`

SetFormulaType sets FormulaType field to given value.

### HasFormulaType

`func (o *GetPublicFundingRateHistoryV5RespData) HasFormulaType() bool`

HasFormulaType returns a boolean if a field has been set.

### GetFundingRate

`func (o *GetPublicFundingRateHistoryV5RespData) GetFundingRate() string`

GetFundingRate returns the FundingRate field if non-nil, zero value otherwise.

### GetFundingRateOk

`func (o *GetPublicFundingRateHistoryV5RespData) GetFundingRateOk() (*string, bool)`

GetFundingRateOk returns a tuple with the FundingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingRate

`func (o *GetPublicFundingRateHistoryV5RespData) SetFundingRate(v string)`

SetFundingRate sets FundingRate field to given value.

### HasFundingRate

`func (o *GetPublicFundingRateHistoryV5RespData) HasFundingRate() bool`

HasFundingRate returns a boolean if a field has been set.

### GetFundingTime

`func (o *GetPublicFundingRateHistoryV5RespData) GetFundingTime() string`

GetFundingTime returns the FundingTime field if non-nil, zero value otherwise.

### GetFundingTimeOk

`func (o *GetPublicFundingRateHistoryV5RespData) GetFundingTimeOk() (*string, bool)`

GetFundingTimeOk returns a tuple with the FundingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingTime

`func (o *GetPublicFundingRateHistoryV5RespData) SetFundingTime(v string)`

SetFundingTime sets FundingTime field to given value.

### HasFundingTime

`func (o *GetPublicFundingRateHistoryV5RespData) HasFundingTime() bool`

HasFundingTime returns a boolean if a field has been set.

### GetInstId

`func (o *GetPublicFundingRateHistoryV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetPublicFundingRateHistoryV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetPublicFundingRateHistoryV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetPublicFundingRateHistoryV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetPublicFundingRateHistoryV5RespData) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetPublicFundingRateHistoryV5RespData) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetPublicFundingRateHistoryV5RespData) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetPublicFundingRateHistoryV5RespData) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetMethod

`func (o *GetPublicFundingRateHistoryV5RespData) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *GetPublicFundingRateHistoryV5RespData) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *GetPublicFundingRateHistoryV5RespData) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *GetPublicFundingRateHistoryV5RespData) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetRealizedRate

`func (o *GetPublicFundingRateHistoryV5RespData) GetRealizedRate() string`

GetRealizedRate returns the RealizedRate field if non-nil, zero value otherwise.

### GetRealizedRateOk

`func (o *GetPublicFundingRateHistoryV5RespData) GetRealizedRateOk() (*string, bool)`

GetRealizedRateOk returns a tuple with the RealizedRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedRate

`func (o *GetPublicFundingRateHistoryV5RespData) SetRealizedRate(v string)`

SetRealizedRate sets RealizedRate field to given value.

### HasRealizedRate

`func (o *GetPublicFundingRateHistoryV5RespData) HasRealizedRate() bool`

HasRealizedRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


