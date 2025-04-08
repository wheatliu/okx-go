# GetFinanceFlexibleLoanLoanInfoV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollateralData** | Pointer to [**[]GetFinanceFlexibleLoanLoanInfoV5RespDataInnerCollateralDataInner**](GetFinanceFlexibleLoanLoanInfoV5RespDataInnerCollateralDataInner.md) | Collateral data | [optional] 
**CollateralNotionalUsd** | Pointer to **string** | Collateral value in &#x60;USD&#x60; | [optional] [default to ""]
**CurLTV** | Pointer to **string** | Current LTV, e.g. &#x60;0.1&#x60; represents &#x60;10%&#x60;  Note: LTV &#x3D; Loan to Value | [optional] [default to ""]
**LiqLTV** | Pointer to **string** | Liquidation LTV, e.g. &#x60;0.1&#x60; represents &#x60;10%&#x60;  If your loan reaches liquidation LTV, it&#39;ll trigger forced liquidation. When this happens, you&#39;ll lose access to your collateral and any repayments made. | [optional] [default to ""]
**LoanData** | Pointer to [**[]GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner**](GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner.md) | Loan data | [optional] 
**LoanNotionalUsd** | Pointer to **string** | Loan value in &#x60;USD&#x60; | [optional] [default to ""]
**MarginCallLTV** | Pointer to **string** | Margin call LTV, e.g. &#x60;0.1&#x60; represents &#x60;10%&#x60;  If your loan hits the margin call LTV, our system will automatically warn you that your loan is getting close to forced liquidation. | [optional] [default to ""]
**RiskWarningData** | Pointer to [**GetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningData**](GetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningData.md) |  | [optional] 

## Methods

### NewGetFinanceFlexibleLoanLoanInfoV5RespDataInner

`func NewGetFinanceFlexibleLoanLoanInfoV5RespDataInner() *GetFinanceFlexibleLoanLoanInfoV5RespDataInner`

NewGetFinanceFlexibleLoanLoanInfoV5RespDataInner instantiates a new GetFinanceFlexibleLoanLoanInfoV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinanceFlexibleLoanLoanInfoV5RespDataInnerWithDefaults

`func NewGetFinanceFlexibleLoanLoanInfoV5RespDataInnerWithDefaults() *GetFinanceFlexibleLoanLoanInfoV5RespDataInner`

NewGetFinanceFlexibleLoanLoanInfoV5RespDataInnerWithDefaults instantiates a new GetFinanceFlexibleLoanLoanInfoV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollateralData

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) GetCollateralData() []GetFinanceFlexibleLoanLoanInfoV5RespDataInnerCollateralDataInner`

GetCollateralData returns the CollateralData field if non-nil, zero value otherwise.

### GetCollateralDataOk

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) GetCollateralDataOk() (*[]GetFinanceFlexibleLoanLoanInfoV5RespDataInnerCollateralDataInner, bool)`

GetCollateralDataOk returns a tuple with the CollateralData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralData

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) SetCollateralData(v []GetFinanceFlexibleLoanLoanInfoV5RespDataInnerCollateralDataInner)`

SetCollateralData sets CollateralData field to given value.

### HasCollateralData

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) HasCollateralData() bool`

HasCollateralData returns a boolean if a field has been set.

### GetCollateralNotionalUsd

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) GetCollateralNotionalUsd() string`

GetCollateralNotionalUsd returns the CollateralNotionalUsd field if non-nil, zero value otherwise.

### GetCollateralNotionalUsdOk

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) GetCollateralNotionalUsdOk() (*string, bool)`

GetCollateralNotionalUsdOk returns a tuple with the CollateralNotionalUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralNotionalUsd

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) SetCollateralNotionalUsd(v string)`

SetCollateralNotionalUsd sets CollateralNotionalUsd field to given value.

### HasCollateralNotionalUsd

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) HasCollateralNotionalUsd() bool`

HasCollateralNotionalUsd returns a boolean if a field has been set.

### GetCurLTV

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) GetCurLTV() string`

GetCurLTV returns the CurLTV field if non-nil, zero value otherwise.

### GetCurLTVOk

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) GetCurLTVOk() (*string, bool)`

GetCurLTVOk returns a tuple with the CurLTV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurLTV

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) SetCurLTV(v string)`

SetCurLTV sets CurLTV field to given value.

### HasCurLTV

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) HasCurLTV() bool`

HasCurLTV returns a boolean if a field has been set.

### GetLiqLTV

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) GetLiqLTV() string`

GetLiqLTV returns the LiqLTV field if non-nil, zero value otherwise.

### GetLiqLTVOk

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) GetLiqLTVOk() (*string, bool)`

GetLiqLTVOk returns a tuple with the LiqLTV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiqLTV

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) SetLiqLTV(v string)`

SetLiqLTV sets LiqLTV field to given value.

### HasLiqLTV

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) HasLiqLTV() bool`

HasLiqLTV returns a boolean if a field has been set.

### GetLoanData

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) GetLoanData() []GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner`

GetLoanData returns the LoanData field if non-nil, zero value otherwise.

### GetLoanDataOk

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) GetLoanDataOk() (*[]GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner, bool)`

GetLoanDataOk returns a tuple with the LoanData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanData

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) SetLoanData(v []GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner)`

SetLoanData sets LoanData field to given value.

### HasLoanData

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) HasLoanData() bool`

HasLoanData returns a boolean if a field has been set.

### GetLoanNotionalUsd

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) GetLoanNotionalUsd() string`

GetLoanNotionalUsd returns the LoanNotionalUsd field if non-nil, zero value otherwise.

### GetLoanNotionalUsdOk

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) GetLoanNotionalUsdOk() (*string, bool)`

GetLoanNotionalUsdOk returns a tuple with the LoanNotionalUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanNotionalUsd

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) SetLoanNotionalUsd(v string)`

SetLoanNotionalUsd sets LoanNotionalUsd field to given value.

### HasLoanNotionalUsd

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) HasLoanNotionalUsd() bool`

HasLoanNotionalUsd returns a boolean if a field has been set.

### GetMarginCallLTV

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) GetMarginCallLTV() string`

GetMarginCallLTV returns the MarginCallLTV field if non-nil, zero value otherwise.

### GetMarginCallLTVOk

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) GetMarginCallLTVOk() (*string, bool)`

GetMarginCallLTVOk returns a tuple with the MarginCallLTV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginCallLTV

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) SetMarginCallLTV(v string)`

SetMarginCallLTV sets MarginCallLTV field to given value.

### HasMarginCallLTV

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) HasMarginCallLTV() bool`

HasMarginCallLTV returns a boolean if a field has been set.

### GetRiskWarningData

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) GetRiskWarningData() GetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningData`

GetRiskWarningData returns the RiskWarningData field if non-nil, zero value otherwise.

### GetRiskWarningDataOk

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) GetRiskWarningDataOk() (*GetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningData, bool)`

GetRiskWarningDataOk returns a tuple with the RiskWarningData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskWarningData

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) SetRiskWarningData(v GetFinanceFlexibleLoanLoanInfoV5RespDataInnerRiskWarningData)`

SetRiskWarningData sets RiskWarningData field to given value.

### HasRiskWarningData

`func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInner) HasRiskWarningData() bool`

HasRiskWarningData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


