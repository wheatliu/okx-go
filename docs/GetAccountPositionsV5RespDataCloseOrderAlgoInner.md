# GetAccountPositionsV5RespDataCloseOrderAlgoInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoId** | Pointer to **string** | Algo ID | [optional] [default to ""]
**CloseFraction** | Pointer to **string** | Fraction of position to be closed when the algo order is triggered. | [optional] [default to ""]
**SlTriggerPx** | Pointer to **string** | Stop-loss trigger price. | [optional] [default to ""]
**SlTriggerPxType** | Pointer to **string** | Stop-loss trigger price type.   &#x60;last&#x60;: last price  &#x60;index&#x60;: index price  &#x60;mark&#x60;: mark price | [optional] [default to ""]
**TpTriggerPx** | Pointer to **string** | Take-profit trigger price. | [optional] [default to ""]
**TpTriggerPxType** | Pointer to **string** | Take-profit trigger price type.   &#x60;last&#x60;: last price  &#x60;index&#x60;: index price  &#x60;mark&#x60;: mark price | [optional] [default to ""]

## Methods

### NewGetAccountPositionsV5RespDataCloseOrderAlgoInner

`func NewGetAccountPositionsV5RespDataCloseOrderAlgoInner() *GetAccountPositionsV5RespDataCloseOrderAlgoInner`

NewGetAccountPositionsV5RespDataCloseOrderAlgoInner instantiates a new GetAccountPositionsV5RespDataCloseOrderAlgoInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountPositionsV5RespDataCloseOrderAlgoInnerWithDefaults

`func NewGetAccountPositionsV5RespDataCloseOrderAlgoInnerWithDefaults() *GetAccountPositionsV5RespDataCloseOrderAlgoInner`

NewGetAccountPositionsV5RespDataCloseOrderAlgoInnerWithDefaults instantiates a new GetAccountPositionsV5RespDataCloseOrderAlgoInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoId

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetCloseFraction

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) GetCloseFraction() string`

GetCloseFraction returns the CloseFraction field if non-nil, zero value otherwise.

### GetCloseFractionOk

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) GetCloseFractionOk() (*string, bool)`

GetCloseFractionOk returns a tuple with the CloseFraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseFraction

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) SetCloseFraction(v string)`

SetCloseFraction sets CloseFraction field to given value.

### HasCloseFraction

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) HasCloseFraction() bool`

HasCloseFraction returns a boolean if a field has been set.

### GetSlTriggerPx

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) GetSlTriggerPx() string`

GetSlTriggerPx returns the SlTriggerPx field if non-nil, zero value otherwise.

### GetSlTriggerPxOk

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) GetSlTriggerPxOk() (*string, bool)`

GetSlTriggerPxOk returns a tuple with the SlTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPx

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) SetSlTriggerPx(v string)`

SetSlTriggerPx sets SlTriggerPx field to given value.

### HasSlTriggerPx

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) HasSlTriggerPx() bool`

HasSlTriggerPx returns a boolean if a field has been set.

### GetSlTriggerPxType

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) GetSlTriggerPxType() string`

GetSlTriggerPxType returns the SlTriggerPxType field if non-nil, zero value otherwise.

### GetSlTriggerPxTypeOk

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) GetSlTriggerPxTypeOk() (*string, bool)`

GetSlTriggerPxTypeOk returns a tuple with the SlTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPxType

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) SetSlTriggerPxType(v string)`

SetSlTriggerPxType sets SlTriggerPxType field to given value.

### HasSlTriggerPxType

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) HasSlTriggerPxType() bool`

HasSlTriggerPxType returns a boolean if a field has been set.

### GetTpTriggerPx

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) GetTpTriggerPx() string`

GetTpTriggerPx returns the TpTriggerPx field if non-nil, zero value otherwise.

### GetTpTriggerPxOk

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) GetTpTriggerPxOk() (*string, bool)`

GetTpTriggerPxOk returns a tuple with the TpTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPx

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) SetTpTriggerPx(v string)`

SetTpTriggerPx sets TpTriggerPx field to given value.

### HasTpTriggerPx

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) HasTpTriggerPx() bool`

HasTpTriggerPx returns a boolean if a field has been set.

### GetTpTriggerPxType

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) GetTpTriggerPxType() string`

GetTpTriggerPxType returns the TpTriggerPxType field if non-nil, zero value otherwise.

### GetTpTriggerPxTypeOk

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) GetTpTriggerPxTypeOk() (*string, bool)`

GetTpTriggerPxTypeOk returns a tuple with the TpTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPxType

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) SetTpTriggerPxType(v string)`

SetTpTriggerPxType sets TpTriggerPxType field to given value.

### HasTpTriggerPxType

`func (o *GetAccountPositionsV5RespDataCloseOrderAlgoInner) HasTpTriggerPxType() bool`

HasTpTriggerPxType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


