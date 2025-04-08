# GetAccountAccountPositionRiskV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdjEq** | Pointer to **string** | Adjusted / Effective equity in &#x60;USD&#x60;  Applicable to &#x60;Multi-currency margin&#x60; and &#x60;Portfolio margin&#x60; | [optional] [default to ""]
**BalData** | Pointer to [**[]GetAccountAccountPositionRiskV5RespDataInnerBalDataInner**](GetAccountAccountPositionRiskV5RespDataInnerBalDataInner.md) | Detailed asset information in all currencies | [optional] 
**PosData** | Pointer to [**[]GetAccountAccountPositionRiskV5RespDataInnerPosDataInner**](GetAccountAccountPositionRiskV5RespDataInnerPosDataInner.md) | Detailed position information in all currencies | [optional] 
**Ts** | Pointer to **string** | Update time of account information, millisecond format of Unix timestamp, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetAccountAccountPositionRiskV5RespDataInner

`func NewGetAccountAccountPositionRiskV5RespDataInner() *GetAccountAccountPositionRiskV5RespDataInner`

NewGetAccountAccountPositionRiskV5RespDataInner instantiates a new GetAccountAccountPositionRiskV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountAccountPositionRiskV5RespDataInnerWithDefaults

`func NewGetAccountAccountPositionRiskV5RespDataInnerWithDefaults() *GetAccountAccountPositionRiskV5RespDataInner`

NewGetAccountAccountPositionRiskV5RespDataInnerWithDefaults instantiates a new GetAccountAccountPositionRiskV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdjEq

`func (o *GetAccountAccountPositionRiskV5RespDataInner) GetAdjEq() string`

GetAdjEq returns the AdjEq field if non-nil, zero value otherwise.

### GetAdjEqOk

`func (o *GetAccountAccountPositionRiskV5RespDataInner) GetAdjEqOk() (*string, bool)`

GetAdjEqOk returns a tuple with the AdjEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjEq

`func (o *GetAccountAccountPositionRiskV5RespDataInner) SetAdjEq(v string)`

SetAdjEq sets AdjEq field to given value.

### HasAdjEq

`func (o *GetAccountAccountPositionRiskV5RespDataInner) HasAdjEq() bool`

HasAdjEq returns a boolean if a field has been set.

### GetBalData

`func (o *GetAccountAccountPositionRiskV5RespDataInner) GetBalData() []GetAccountAccountPositionRiskV5RespDataInnerBalDataInner`

GetBalData returns the BalData field if non-nil, zero value otherwise.

### GetBalDataOk

`func (o *GetAccountAccountPositionRiskV5RespDataInner) GetBalDataOk() (*[]GetAccountAccountPositionRiskV5RespDataInnerBalDataInner, bool)`

GetBalDataOk returns a tuple with the BalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalData

`func (o *GetAccountAccountPositionRiskV5RespDataInner) SetBalData(v []GetAccountAccountPositionRiskV5RespDataInnerBalDataInner)`

SetBalData sets BalData field to given value.

### HasBalData

`func (o *GetAccountAccountPositionRiskV5RespDataInner) HasBalData() bool`

HasBalData returns a boolean if a field has been set.

### GetPosData

`func (o *GetAccountAccountPositionRiskV5RespDataInner) GetPosData() []GetAccountAccountPositionRiskV5RespDataInnerPosDataInner`

GetPosData returns the PosData field if non-nil, zero value otherwise.

### GetPosDataOk

`func (o *GetAccountAccountPositionRiskV5RespDataInner) GetPosDataOk() (*[]GetAccountAccountPositionRiskV5RespDataInnerPosDataInner, bool)`

GetPosDataOk returns a tuple with the PosData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosData

`func (o *GetAccountAccountPositionRiskV5RespDataInner) SetPosData(v []GetAccountAccountPositionRiskV5RespDataInnerPosDataInner)`

SetPosData sets PosData field to given value.

### HasPosData

`func (o *GetAccountAccountPositionRiskV5RespDataInner) HasPosData() bool`

HasPosData returns a boolean if a field has been set.

### GetTs

`func (o *GetAccountAccountPositionRiskV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetAccountAccountPositionRiskV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetAccountAccountPositionRiskV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetAccountAccountPositionRiskV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


