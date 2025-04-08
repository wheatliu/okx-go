# GetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstId** | Pointer to **string** | Liquidation instrument ID, e.g. &#x60;BTC-USDT&#x60;  This field is only valid when there is only one type of collateral and one type of borrowed currency. In other cases, it returns \&quot;\&quot;. | [optional] [default to ""]
**LiqPx** | Pointer to **string** | Liquidation price  The unit of the liquidation price is the quote currency of the instrument, e.g. &#x60;USDT&#x60; in &#x60;BTC-USDT&#x60;.  This field is only valid when there is only one type of collateral and one type of borrowed currency. In other cases, it returns \&quot;\&quot;. | [optional] [default to ""]

## Methods

### NewGetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningData

`func NewGetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningData() *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningData`

NewGetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningData instantiates a new GetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningDataWithDefaults

`func NewGetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningDataWithDefaults() *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningData`

NewGetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningDataWithDefaults instantiates a new GetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstId

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetLiqPx

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningData) GetLiqPx() string`

GetLiqPx returns the LiqPx field if non-nil, zero value otherwise.

### GetLiqPxOk

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningData) GetLiqPxOk() (*string, bool)`

GetLiqPxOk returns a tuple with the LiqPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiqPx

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningData) SetLiqPx(v string)`

SetLiqPx sets LiqPx field to given value.

### HasLiqPx

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningData) HasLiqPx() bool`

HasLiqPx returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


