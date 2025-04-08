# CreateAccountPositionBuilderV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctLever** | Pointer to **string** | Leverage of the account | [optional] [default to ""]
**Assets** | Pointer to [**[]CreateAccountPositionBuilderV5RespDataInnerAssetsInner**](CreateAccountPositionBuilderV5RespDataInnerAssetsInner.md) | Asset info | [optional] 
**BorrowMmr** | Pointer to **string** | Borrow MMR (&#x60;USD&#x60;) for the account | [optional] [default to ""]
**DerivMmr** | Pointer to **string** | Derivatives MMR (&#x60;USD&#x60;) for the account | [optional] [default to ""]
**Eq** | Pointer to **string** | Adjusted equity (&#x60;USD&#x60;) for the account | [optional] [default to ""]
**MarginRatio** | Pointer to **string** | Cross margin ratio for the account | [optional] [default to ""]
**Positions** | Pointer to [**[]CreateAccountPositionBuilderV5RespDataInnerPositionsInner**](CreateAccountPositionBuilderV5RespDataInnerPositionsInner.md) | Position info  Only applicable to &#x60;Multi-currency margin&#x60; | [optional] 
**RiskUnitData** | Pointer to [**[]CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner**](CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner.md) | Risk unit info | [optional] 
**TotalImr** | Pointer to **string** | Total IMR (&#x60;USD&#x60;) for the account | [optional] [default to ""]
**TotalMmr** | Pointer to **string** | Total MMR (&#x60;USD&#x60;) for the account | [optional] [default to ""]
**Ts** | Pointer to **string** | Update time for the account, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Upl** | Pointer to **string** | UPL for the account | [optional] [default to ""]

## Methods

### NewCreateAccountPositionBuilderV5RespDataInner

`func NewCreateAccountPositionBuilderV5RespDataInner() *CreateAccountPositionBuilderV5RespDataInner`

NewCreateAccountPositionBuilderV5RespDataInner instantiates a new CreateAccountPositionBuilderV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountPositionBuilderV5RespDataInnerWithDefaults

`func NewCreateAccountPositionBuilderV5RespDataInnerWithDefaults() *CreateAccountPositionBuilderV5RespDataInner`

NewCreateAccountPositionBuilderV5RespDataInnerWithDefaults instantiates a new CreateAccountPositionBuilderV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctLever

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetAcctLever() string`

GetAcctLever returns the AcctLever field if non-nil, zero value otherwise.

### GetAcctLeverOk

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetAcctLeverOk() (*string, bool)`

GetAcctLeverOk returns a tuple with the AcctLever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctLever

`func (o *CreateAccountPositionBuilderV5RespDataInner) SetAcctLever(v string)`

SetAcctLever sets AcctLever field to given value.

### HasAcctLever

`func (o *CreateAccountPositionBuilderV5RespDataInner) HasAcctLever() bool`

HasAcctLever returns a boolean if a field has been set.

### GetAssets

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetAssets() []CreateAccountPositionBuilderV5RespDataInnerAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetAssetsOk() (*[]CreateAccountPositionBuilderV5RespDataInnerAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *CreateAccountPositionBuilderV5RespDataInner) SetAssets(v []CreateAccountPositionBuilderV5RespDataInnerAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *CreateAccountPositionBuilderV5RespDataInner) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetBorrowMmr

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetBorrowMmr() string`

GetBorrowMmr returns the BorrowMmr field if non-nil, zero value otherwise.

### GetBorrowMmrOk

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetBorrowMmrOk() (*string, bool)`

GetBorrowMmrOk returns a tuple with the BorrowMmr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowMmr

`func (o *CreateAccountPositionBuilderV5RespDataInner) SetBorrowMmr(v string)`

SetBorrowMmr sets BorrowMmr field to given value.

### HasBorrowMmr

`func (o *CreateAccountPositionBuilderV5RespDataInner) HasBorrowMmr() bool`

HasBorrowMmr returns a boolean if a field has been set.

### GetDerivMmr

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetDerivMmr() string`

GetDerivMmr returns the DerivMmr field if non-nil, zero value otherwise.

### GetDerivMmrOk

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetDerivMmrOk() (*string, bool)`

GetDerivMmrOk returns a tuple with the DerivMmr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivMmr

`func (o *CreateAccountPositionBuilderV5RespDataInner) SetDerivMmr(v string)`

SetDerivMmr sets DerivMmr field to given value.

### HasDerivMmr

`func (o *CreateAccountPositionBuilderV5RespDataInner) HasDerivMmr() bool`

HasDerivMmr returns a boolean if a field has been set.

### GetEq

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetEq() string`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetEqOk() (*string, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *CreateAccountPositionBuilderV5RespDataInner) SetEq(v string)`

SetEq sets Eq field to given value.

### HasEq

`func (o *CreateAccountPositionBuilderV5RespDataInner) HasEq() bool`

HasEq returns a boolean if a field has been set.

### GetMarginRatio

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetMarginRatio() string`

GetMarginRatio returns the MarginRatio field if non-nil, zero value otherwise.

### GetMarginRatioOk

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetMarginRatioOk() (*string, bool)`

GetMarginRatioOk returns a tuple with the MarginRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginRatio

`func (o *CreateAccountPositionBuilderV5RespDataInner) SetMarginRatio(v string)`

SetMarginRatio sets MarginRatio field to given value.

### HasMarginRatio

`func (o *CreateAccountPositionBuilderV5RespDataInner) HasMarginRatio() bool`

HasMarginRatio returns a boolean if a field has been set.

### GetPositions

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetPositions() []CreateAccountPositionBuilderV5RespDataInnerPositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetPositionsOk() (*[]CreateAccountPositionBuilderV5RespDataInnerPositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *CreateAccountPositionBuilderV5RespDataInner) SetPositions(v []CreateAccountPositionBuilderV5RespDataInnerPositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *CreateAccountPositionBuilderV5RespDataInner) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetRiskUnitData

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetRiskUnitData() []CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner`

GetRiskUnitData returns the RiskUnitData field if non-nil, zero value otherwise.

### GetRiskUnitDataOk

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetRiskUnitDataOk() (*[]CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner, bool)`

GetRiskUnitDataOk returns a tuple with the RiskUnitData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskUnitData

`func (o *CreateAccountPositionBuilderV5RespDataInner) SetRiskUnitData(v []CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInner)`

SetRiskUnitData sets RiskUnitData field to given value.

### HasRiskUnitData

`func (o *CreateAccountPositionBuilderV5RespDataInner) HasRiskUnitData() bool`

HasRiskUnitData returns a boolean if a field has been set.

### GetTotalImr

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetTotalImr() string`

GetTotalImr returns the TotalImr field if non-nil, zero value otherwise.

### GetTotalImrOk

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetTotalImrOk() (*string, bool)`

GetTotalImrOk returns a tuple with the TotalImr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalImr

`func (o *CreateAccountPositionBuilderV5RespDataInner) SetTotalImr(v string)`

SetTotalImr sets TotalImr field to given value.

### HasTotalImr

`func (o *CreateAccountPositionBuilderV5RespDataInner) HasTotalImr() bool`

HasTotalImr returns a boolean if a field has been set.

### GetTotalMmr

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetTotalMmr() string`

GetTotalMmr returns the TotalMmr field if non-nil, zero value otherwise.

### GetTotalMmrOk

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetTotalMmrOk() (*string, bool)`

GetTotalMmrOk returns a tuple with the TotalMmr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMmr

`func (o *CreateAccountPositionBuilderV5RespDataInner) SetTotalMmr(v string)`

SetTotalMmr sets TotalMmr field to given value.

### HasTotalMmr

`func (o *CreateAccountPositionBuilderV5RespDataInner) HasTotalMmr() bool`

HasTotalMmr returns a boolean if a field has been set.

### GetTs

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *CreateAccountPositionBuilderV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *CreateAccountPositionBuilderV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetUpl

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetUpl() string`

GetUpl returns the Upl field if non-nil, zero value otherwise.

### GetUplOk

`func (o *CreateAccountPositionBuilderV5RespDataInner) GetUplOk() (*string, bool)`

GetUplOk returns a tuple with the Upl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpl

`func (o *CreateAccountPositionBuilderV5RespDataInner) SetUpl(v string)`

SetUpl sets Upl field to given value.

### HasUpl

`func (o *CreateAccountPositionBuilderV5RespDataInner) HasUpl() bool`

HasUpl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


