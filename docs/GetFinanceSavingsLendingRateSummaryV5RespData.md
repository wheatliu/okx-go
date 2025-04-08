# GetFinanceSavingsLendingRateSummaryV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgAmt** | Pointer to **string** | 24H average borrowing amount | [optional] [default to ""]
**AvgAmtUsd** | Pointer to **string** | 24H average borrowing amount in &#x60;USD&#x60; value | [optional] [default to ""]
**AvgRate** | Pointer to **string** | 24H average lending rate | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency, e.g. &#x60;BTC&#x60; | [optional] [default to ""]
**EstRate** | Pointer to **string** | Next estimate annual interest rate | [optional] [default to ""]
**PreRate** | Pointer to **string** | Last annual interest rate | [optional] [default to ""]

## Methods

### NewGetFinanceSavingsLendingRateSummaryV5RespData

`func NewGetFinanceSavingsLendingRateSummaryV5RespData() *GetFinanceSavingsLendingRateSummaryV5RespData`

NewGetFinanceSavingsLendingRateSummaryV5RespData instantiates a new GetFinanceSavingsLendingRateSummaryV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinanceSavingsLendingRateSummaryV5RespDataWithDefaults

`func NewGetFinanceSavingsLendingRateSummaryV5RespDataWithDefaults() *GetFinanceSavingsLendingRateSummaryV5RespData`

NewGetFinanceSavingsLendingRateSummaryV5RespDataWithDefaults instantiates a new GetFinanceSavingsLendingRateSummaryV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgAmt

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) GetAvgAmt() string`

GetAvgAmt returns the AvgAmt field if non-nil, zero value otherwise.

### GetAvgAmtOk

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) GetAvgAmtOk() (*string, bool)`

GetAvgAmtOk returns a tuple with the AvgAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgAmt

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) SetAvgAmt(v string)`

SetAvgAmt sets AvgAmt field to given value.

### HasAvgAmt

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) HasAvgAmt() bool`

HasAvgAmt returns a boolean if a field has been set.

### GetAvgAmtUsd

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) GetAvgAmtUsd() string`

GetAvgAmtUsd returns the AvgAmtUsd field if non-nil, zero value otherwise.

### GetAvgAmtUsdOk

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) GetAvgAmtUsdOk() (*string, bool)`

GetAvgAmtUsdOk returns a tuple with the AvgAmtUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgAmtUsd

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) SetAvgAmtUsd(v string)`

SetAvgAmtUsd sets AvgAmtUsd field to given value.

### HasAvgAmtUsd

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) HasAvgAmtUsd() bool`

HasAvgAmtUsd returns a boolean if a field has been set.

### GetAvgRate

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) GetAvgRate() string`

GetAvgRate returns the AvgRate field if non-nil, zero value otherwise.

### GetAvgRateOk

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) GetAvgRateOk() (*string, bool)`

GetAvgRateOk returns a tuple with the AvgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgRate

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) SetAvgRate(v string)`

SetAvgRate sets AvgRate field to given value.

### HasAvgRate

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) HasAvgRate() bool`

HasAvgRate returns a boolean if a field has been set.

### GetCcy

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetEstRate

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) GetEstRate() string`

GetEstRate returns the EstRate field if non-nil, zero value otherwise.

### GetEstRateOk

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) GetEstRateOk() (*string, bool)`

GetEstRateOk returns a tuple with the EstRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstRate

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) SetEstRate(v string)`

SetEstRate sets EstRate field to given value.

### HasEstRate

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) HasEstRate() bool`

HasEstRate returns a boolean if a field has been set.

### GetPreRate

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) GetPreRate() string`

GetPreRate returns the PreRate field if non-nil, zero value otherwise.

### GetPreRateOk

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) GetPreRateOk() (*string, bool)`

GetPreRateOk returns a tuple with the PreRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreRate

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) SetPreRate(v string)`

SetPreRate sets PreRate field to given value.

### HasPreRate

`func (o *GetFinanceSavingsLendingRateSummaryV5RespData) HasPreRate() bool`

HasPreRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


