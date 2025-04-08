# CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pnl** | Pointer to **string** | MR1 stress P&amp;L (&#x60;USD&#x60;) | [optional] [default to ""]
**SpotShock** | Pointer to **string** | MR1 worst-case scenario spot shock (in percentage), e.g. &#x60;0.01&#x60; representing &#x60;1%&#x60; | [optional] [default to ""]
**VolShock** | Pointer to **string** | MR1 worst-case scenario volatility shock  &#x60;down&#x60;: volatility shock down  &#x60;unchange&#x60;: volatility unchanged  &#x60;up&#x60;: volatility shock up | [optional] [default to ""]

## Methods

### NewCreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult

`func NewCreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult() *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult`

NewCreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult instantiates a new CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResultWithDefaults

`func NewCreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResultWithDefaults() *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult`

NewCreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResultWithDefaults instantiates a new CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPnl

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult) GetPnl() string`

GetPnl returns the Pnl field if non-nil, zero value otherwise.

### GetPnlOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult) GetPnlOk() (*string, bool)`

GetPnlOk returns a tuple with the Pnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnl

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult) SetPnl(v string)`

SetPnl sets Pnl field to given value.

### HasPnl

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult) HasPnl() bool`

HasPnl returns a boolean if a field has been set.

### GetSpotShock

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult) GetSpotShock() string`

GetSpotShock returns the SpotShock field if non-nil, zero value otherwise.

### GetSpotShockOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult) GetSpotShockOk() (*string, bool)`

GetSpotShockOk returns a tuple with the SpotShock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotShock

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult) SetSpotShock(v string)`

SetSpotShock sets SpotShock field to given value.

### HasSpotShock

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult) HasSpotShock() bool`

HasSpotShock returns a boolean if a field has been set.

### GetVolShock

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult) GetVolShock() string`

GetVolShock returns the VolShock field if non-nil, zero value otherwise.

### GetVolShockOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult) GetVolShockOk() (*string, bool)`

GetVolShockOk returns a tuple with the VolShock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolShock

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult) SetVolShock(v string)`

SetVolShock sets VolShock field to given value.

### HasVolShock

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1FinalResult) HasVolShock() bool`

HasVolShock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


