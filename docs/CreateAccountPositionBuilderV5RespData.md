# CreateAccountPositionBuilderV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctLever** | Pointer to **string** | Leverage of the account | [optional] [default to ""]
**Assets** | Pointer to [**[]CreateAccountPositionBuilderV5RespDataAssetsInner**](CreateAccountPositionBuilderV5RespDataAssetsInner.md) | Asset info | [optional] 
**BorrowMmr** | Pointer to **string** | Borrow MMR (&#x60;USD&#x60;) for the account | [optional] [default to ""]
**DerivMmr** | Pointer to **string** | Derivatives MMR (&#x60;USD&#x60;) for the account | [optional] [default to ""]
**Eq** | Pointer to **string** | Adjusted equity (&#x60;USD&#x60;) for the account | [optional] [default to ""]
**MarginRatio** | Pointer to **string** | Cross margin ratio for the account | [optional] [default to ""]
**Positions** | Pointer to [**[]CreateAccountPositionBuilderV5RespDataPositionsInner**](CreateAccountPositionBuilderV5RespDataPositionsInner.md) | Position info  Only applicable to &#x60;Multi-currency margin&#x60; | [optional] 
**RiskUnitData** | Pointer to [**[]CreateAccountPositionBuilderV5RespDataRiskUnitDataInner**](CreateAccountPositionBuilderV5RespDataRiskUnitDataInner.md) | Risk unit info | [optional] 
**TotalImr** | Pointer to **string** | Total IMR (&#x60;USD&#x60;) for the account | [optional] [default to ""]
**TotalMmr** | Pointer to **string** | Total MMR (&#x60;USD&#x60;) for the account | [optional] [default to ""]
**Ts** | Pointer to **string** | Update time for the account, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Upl** | Pointer to **string** | UPL for the account | [optional] [default to ""]

## Methods

### NewCreateAccountPositionBuilderV5RespData

`func NewCreateAccountPositionBuilderV5RespData() *CreateAccountPositionBuilderV5RespData`

NewCreateAccountPositionBuilderV5RespData instantiates a new CreateAccountPositionBuilderV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountPositionBuilderV5RespDataWithDefaults

`func NewCreateAccountPositionBuilderV5RespDataWithDefaults() *CreateAccountPositionBuilderV5RespData`

NewCreateAccountPositionBuilderV5RespDataWithDefaults instantiates a new CreateAccountPositionBuilderV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctLever

`func (o *CreateAccountPositionBuilderV5RespData) GetAcctLever() string`

GetAcctLever returns the AcctLever field if non-nil, zero value otherwise.

### GetAcctLeverOk

`func (o *CreateAccountPositionBuilderV5RespData) GetAcctLeverOk() (*string, bool)`

GetAcctLeverOk returns a tuple with the AcctLever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctLever

`func (o *CreateAccountPositionBuilderV5RespData) SetAcctLever(v string)`

SetAcctLever sets AcctLever field to given value.

### HasAcctLever

`func (o *CreateAccountPositionBuilderV5RespData) HasAcctLever() bool`

HasAcctLever returns a boolean if a field has been set.

### GetAssets

`func (o *CreateAccountPositionBuilderV5RespData) GetAssets() []CreateAccountPositionBuilderV5RespDataAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *CreateAccountPositionBuilderV5RespData) GetAssetsOk() (*[]CreateAccountPositionBuilderV5RespDataAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *CreateAccountPositionBuilderV5RespData) SetAssets(v []CreateAccountPositionBuilderV5RespDataAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *CreateAccountPositionBuilderV5RespData) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetBorrowMmr

`func (o *CreateAccountPositionBuilderV5RespData) GetBorrowMmr() string`

GetBorrowMmr returns the BorrowMmr field if non-nil, zero value otherwise.

### GetBorrowMmrOk

`func (o *CreateAccountPositionBuilderV5RespData) GetBorrowMmrOk() (*string, bool)`

GetBorrowMmrOk returns a tuple with the BorrowMmr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowMmr

`func (o *CreateAccountPositionBuilderV5RespData) SetBorrowMmr(v string)`

SetBorrowMmr sets BorrowMmr field to given value.

### HasBorrowMmr

`func (o *CreateAccountPositionBuilderV5RespData) HasBorrowMmr() bool`

HasBorrowMmr returns a boolean if a field has been set.

### GetDerivMmr

`func (o *CreateAccountPositionBuilderV5RespData) GetDerivMmr() string`

GetDerivMmr returns the DerivMmr field if non-nil, zero value otherwise.

### GetDerivMmrOk

`func (o *CreateAccountPositionBuilderV5RespData) GetDerivMmrOk() (*string, bool)`

GetDerivMmrOk returns a tuple with the DerivMmr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivMmr

`func (o *CreateAccountPositionBuilderV5RespData) SetDerivMmr(v string)`

SetDerivMmr sets DerivMmr field to given value.

### HasDerivMmr

`func (o *CreateAccountPositionBuilderV5RespData) HasDerivMmr() bool`

HasDerivMmr returns a boolean if a field has been set.

### GetEq

`func (o *CreateAccountPositionBuilderV5RespData) GetEq() string`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *CreateAccountPositionBuilderV5RespData) GetEqOk() (*string, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *CreateAccountPositionBuilderV5RespData) SetEq(v string)`

SetEq sets Eq field to given value.

### HasEq

`func (o *CreateAccountPositionBuilderV5RespData) HasEq() bool`

HasEq returns a boolean if a field has been set.

### GetMarginRatio

`func (o *CreateAccountPositionBuilderV5RespData) GetMarginRatio() string`

GetMarginRatio returns the MarginRatio field if non-nil, zero value otherwise.

### GetMarginRatioOk

`func (o *CreateAccountPositionBuilderV5RespData) GetMarginRatioOk() (*string, bool)`

GetMarginRatioOk returns a tuple with the MarginRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginRatio

`func (o *CreateAccountPositionBuilderV5RespData) SetMarginRatio(v string)`

SetMarginRatio sets MarginRatio field to given value.

### HasMarginRatio

`func (o *CreateAccountPositionBuilderV5RespData) HasMarginRatio() bool`

HasMarginRatio returns a boolean if a field has been set.

### GetPositions

`func (o *CreateAccountPositionBuilderV5RespData) GetPositions() []CreateAccountPositionBuilderV5RespDataPositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *CreateAccountPositionBuilderV5RespData) GetPositionsOk() (*[]CreateAccountPositionBuilderV5RespDataPositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *CreateAccountPositionBuilderV5RespData) SetPositions(v []CreateAccountPositionBuilderV5RespDataPositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *CreateAccountPositionBuilderV5RespData) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetRiskUnitData

`func (o *CreateAccountPositionBuilderV5RespData) GetRiskUnitData() []CreateAccountPositionBuilderV5RespDataRiskUnitDataInner`

GetRiskUnitData returns the RiskUnitData field if non-nil, zero value otherwise.

### GetRiskUnitDataOk

`func (o *CreateAccountPositionBuilderV5RespData) GetRiskUnitDataOk() (*[]CreateAccountPositionBuilderV5RespDataRiskUnitDataInner, bool)`

GetRiskUnitDataOk returns a tuple with the RiskUnitData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskUnitData

`func (o *CreateAccountPositionBuilderV5RespData) SetRiskUnitData(v []CreateAccountPositionBuilderV5RespDataRiskUnitDataInner)`

SetRiskUnitData sets RiskUnitData field to given value.

### HasRiskUnitData

`func (o *CreateAccountPositionBuilderV5RespData) HasRiskUnitData() bool`

HasRiskUnitData returns a boolean if a field has been set.

### GetTotalImr

`func (o *CreateAccountPositionBuilderV5RespData) GetTotalImr() string`

GetTotalImr returns the TotalImr field if non-nil, zero value otherwise.

### GetTotalImrOk

`func (o *CreateAccountPositionBuilderV5RespData) GetTotalImrOk() (*string, bool)`

GetTotalImrOk returns a tuple with the TotalImr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalImr

`func (o *CreateAccountPositionBuilderV5RespData) SetTotalImr(v string)`

SetTotalImr sets TotalImr field to given value.

### HasTotalImr

`func (o *CreateAccountPositionBuilderV5RespData) HasTotalImr() bool`

HasTotalImr returns a boolean if a field has been set.

### GetTotalMmr

`func (o *CreateAccountPositionBuilderV5RespData) GetTotalMmr() string`

GetTotalMmr returns the TotalMmr field if non-nil, zero value otherwise.

### GetTotalMmrOk

`func (o *CreateAccountPositionBuilderV5RespData) GetTotalMmrOk() (*string, bool)`

GetTotalMmrOk returns a tuple with the TotalMmr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMmr

`func (o *CreateAccountPositionBuilderV5RespData) SetTotalMmr(v string)`

SetTotalMmr sets TotalMmr field to given value.

### HasTotalMmr

`func (o *CreateAccountPositionBuilderV5RespData) HasTotalMmr() bool`

HasTotalMmr returns a boolean if a field has been set.

### GetTs

`func (o *CreateAccountPositionBuilderV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *CreateAccountPositionBuilderV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *CreateAccountPositionBuilderV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *CreateAccountPositionBuilderV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetUpl

`func (o *CreateAccountPositionBuilderV5RespData) GetUpl() string`

GetUpl returns the Upl field if non-nil, zero value otherwise.

### GetUplOk

`func (o *CreateAccountPositionBuilderV5RespData) GetUplOk() (*string, bool)`

GetUplOk returns a tuple with the Upl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpl

`func (o *CreateAccountPositionBuilderV5RespData) SetUpl(v string)`

SetUpl sets Upl field to given value.

### HasUpl

`func (o *CreateAccountPositionBuilderV5RespData) HasUpl() bool`

HasUpl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


