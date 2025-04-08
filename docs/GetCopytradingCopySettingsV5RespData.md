# GetCopytradingCopySettingsV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | Pointer to **string** | Margin currency | [optional] [default to ""]
**CopyAmt** | Pointer to **string** | Copy amount in USDT per order. | [optional] [default to ""]
**CopyInstIdType** | Pointer to **string** | Copy contract type setted  &#x60;custom&#x60;: custom by &#x60;instId&#x60; which is required；  &#x60;copy&#x60;: Keep your contracts consistent with this trader by automatically adding or removing contracts when they do | [optional] [default to ""]
**CopyMgnMode** | Pointer to **string** | Copy margin mode  &#x60;cross&#x60;: cross  &#x60;isolated&#x60;: isolated  &#x60;copy&#x60;: Use the same margin mode as lead trader when opening positions | [optional] [default to ""]
**CopyMode** | Pointer to **string** | Copy mode  &#x60;fixed_amount&#x60; &#x60;ratio_copy&#x60; | [optional] [default to ""]
**CopyRatio** | Pointer to **string** | Copy ratio per order. | [optional] [default to ""]
**CopyState** | Pointer to **string** | Current copy state   &#x60;0&#x60;: non-copy, &#x60;1&#x60;: copy | [optional] [default to ""]
**CopyTotalAmt** | Pointer to **string** | Maximum total amount in USDT.   The maximum total amount you&#39;ll invest at any given time across all orders in this copy trade | [optional] [default to ""]
**InstIds** | Pointer to [**[]GetCopytradingCopySettingsV5RespDataInstIdsInner**](GetCopytradingCopySettingsV5RespDataInstIdsInner.md) | Instrument list. It will return all lead contracts of the lead trader | [optional] 
**SlRatio** | Pointer to **string** | Stop loss per order. 0.1 represents 10% | [optional] [default to ""]
**SlTotalAmt** | Pointer to **string** | Total stop loss in USDT for trader. | [optional] [default to ""]
**SubPosCloseType** | Pointer to **string** | Action type for open positions  &#x60;market_close&#x60;: immediately close at market price  &#x60;copy_close&#x60;：close when trader closes  &#x60;manual_close&#x60;: close manually | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**TpRatio** | Pointer to **string** | Take profit per order. 0.1 represents 10% | [optional] [default to ""]

## Methods

### NewGetCopytradingCopySettingsV5RespData

`func NewGetCopytradingCopySettingsV5RespData() *GetCopytradingCopySettingsV5RespData`

NewGetCopytradingCopySettingsV5RespData instantiates a new GetCopytradingCopySettingsV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCopytradingCopySettingsV5RespDataWithDefaults

`func NewGetCopytradingCopySettingsV5RespDataWithDefaults() *GetCopytradingCopySettingsV5RespData`

NewGetCopytradingCopySettingsV5RespDataWithDefaults instantiates a new GetCopytradingCopySettingsV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *GetCopytradingCopySettingsV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetCopytradingCopySettingsV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetCopytradingCopySettingsV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetCopytradingCopySettingsV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetCopyAmt

`func (o *GetCopytradingCopySettingsV5RespData) GetCopyAmt() string`

GetCopyAmt returns the CopyAmt field if non-nil, zero value otherwise.

### GetCopyAmtOk

`func (o *GetCopytradingCopySettingsV5RespData) GetCopyAmtOk() (*string, bool)`

GetCopyAmtOk returns a tuple with the CopyAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyAmt

`func (o *GetCopytradingCopySettingsV5RespData) SetCopyAmt(v string)`

SetCopyAmt sets CopyAmt field to given value.

### HasCopyAmt

`func (o *GetCopytradingCopySettingsV5RespData) HasCopyAmt() bool`

HasCopyAmt returns a boolean if a field has been set.

### GetCopyInstIdType

`func (o *GetCopytradingCopySettingsV5RespData) GetCopyInstIdType() string`

GetCopyInstIdType returns the CopyInstIdType field if non-nil, zero value otherwise.

### GetCopyInstIdTypeOk

`func (o *GetCopytradingCopySettingsV5RespData) GetCopyInstIdTypeOk() (*string, bool)`

GetCopyInstIdTypeOk returns a tuple with the CopyInstIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyInstIdType

`func (o *GetCopytradingCopySettingsV5RespData) SetCopyInstIdType(v string)`

SetCopyInstIdType sets CopyInstIdType field to given value.

### HasCopyInstIdType

`func (o *GetCopytradingCopySettingsV5RespData) HasCopyInstIdType() bool`

HasCopyInstIdType returns a boolean if a field has been set.

### GetCopyMgnMode

`func (o *GetCopytradingCopySettingsV5RespData) GetCopyMgnMode() string`

GetCopyMgnMode returns the CopyMgnMode field if non-nil, zero value otherwise.

### GetCopyMgnModeOk

`func (o *GetCopytradingCopySettingsV5RespData) GetCopyMgnModeOk() (*string, bool)`

GetCopyMgnModeOk returns a tuple with the CopyMgnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyMgnMode

`func (o *GetCopytradingCopySettingsV5RespData) SetCopyMgnMode(v string)`

SetCopyMgnMode sets CopyMgnMode field to given value.

### HasCopyMgnMode

`func (o *GetCopytradingCopySettingsV5RespData) HasCopyMgnMode() bool`

HasCopyMgnMode returns a boolean if a field has been set.

### GetCopyMode

`func (o *GetCopytradingCopySettingsV5RespData) GetCopyMode() string`

GetCopyMode returns the CopyMode field if non-nil, zero value otherwise.

### GetCopyModeOk

`func (o *GetCopytradingCopySettingsV5RespData) GetCopyModeOk() (*string, bool)`

GetCopyModeOk returns a tuple with the CopyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyMode

`func (o *GetCopytradingCopySettingsV5RespData) SetCopyMode(v string)`

SetCopyMode sets CopyMode field to given value.

### HasCopyMode

`func (o *GetCopytradingCopySettingsV5RespData) HasCopyMode() bool`

HasCopyMode returns a boolean if a field has been set.

### GetCopyRatio

`func (o *GetCopytradingCopySettingsV5RespData) GetCopyRatio() string`

GetCopyRatio returns the CopyRatio field if non-nil, zero value otherwise.

### GetCopyRatioOk

`func (o *GetCopytradingCopySettingsV5RespData) GetCopyRatioOk() (*string, bool)`

GetCopyRatioOk returns a tuple with the CopyRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyRatio

`func (o *GetCopytradingCopySettingsV5RespData) SetCopyRatio(v string)`

SetCopyRatio sets CopyRatio field to given value.

### HasCopyRatio

`func (o *GetCopytradingCopySettingsV5RespData) HasCopyRatio() bool`

HasCopyRatio returns a boolean if a field has been set.

### GetCopyState

`func (o *GetCopytradingCopySettingsV5RespData) GetCopyState() string`

GetCopyState returns the CopyState field if non-nil, zero value otherwise.

### GetCopyStateOk

`func (o *GetCopytradingCopySettingsV5RespData) GetCopyStateOk() (*string, bool)`

GetCopyStateOk returns a tuple with the CopyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyState

`func (o *GetCopytradingCopySettingsV5RespData) SetCopyState(v string)`

SetCopyState sets CopyState field to given value.

### HasCopyState

`func (o *GetCopytradingCopySettingsV5RespData) HasCopyState() bool`

HasCopyState returns a boolean if a field has been set.

### GetCopyTotalAmt

`func (o *GetCopytradingCopySettingsV5RespData) GetCopyTotalAmt() string`

GetCopyTotalAmt returns the CopyTotalAmt field if non-nil, zero value otherwise.

### GetCopyTotalAmtOk

`func (o *GetCopytradingCopySettingsV5RespData) GetCopyTotalAmtOk() (*string, bool)`

GetCopyTotalAmtOk returns a tuple with the CopyTotalAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTotalAmt

`func (o *GetCopytradingCopySettingsV5RespData) SetCopyTotalAmt(v string)`

SetCopyTotalAmt sets CopyTotalAmt field to given value.

### HasCopyTotalAmt

`func (o *GetCopytradingCopySettingsV5RespData) HasCopyTotalAmt() bool`

HasCopyTotalAmt returns a boolean if a field has been set.

### GetInstIds

`func (o *GetCopytradingCopySettingsV5RespData) GetInstIds() []GetCopytradingCopySettingsV5RespDataInstIdsInner`

GetInstIds returns the InstIds field if non-nil, zero value otherwise.

### GetInstIdsOk

`func (o *GetCopytradingCopySettingsV5RespData) GetInstIdsOk() (*[]GetCopytradingCopySettingsV5RespDataInstIdsInner, bool)`

GetInstIdsOk returns a tuple with the InstIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstIds

`func (o *GetCopytradingCopySettingsV5RespData) SetInstIds(v []GetCopytradingCopySettingsV5RespDataInstIdsInner)`

SetInstIds sets InstIds field to given value.

### HasInstIds

`func (o *GetCopytradingCopySettingsV5RespData) HasInstIds() bool`

HasInstIds returns a boolean if a field has been set.

### GetSlRatio

`func (o *GetCopytradingCopySettingsV5RespData) GetSlRatio() string`

GetSlRatio returns the SlRatio field if non-nil, zero value otherwise.

### GetSlRatioOk

`func (o *GetCopytradingCopySettingsV5RespData) GetSlRatioOk() (*string, bool)`

GetSlRatioOk returns a tuple with the SlRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlRatio

`func (o *GetCopytradingCopySettingsV5RespData) SetSlRatio(v string)`

SetSlRatio sets SlRatio field to given value.

### HasSlRatio

`func (o *GetCopytradingCopySettingsV5RespData) HasSlRatio() bool`

HasSlRatio returns a boolean if a field has been set.

### GetSlTotalAmt

`func (o *GetCopytradingCopySettingsV5RespData) GetSlTotalAmt() string`

GetSlTotalAmt returns the SlTotalAmt field if non-nil, zero value otherwise.

### GetSlTotalAmtOk

`func (o *GetCopytradingCopySettingsV5RespData) GetSlTotalAmtOk() (*string, bool)`

GetSlTotalAmtOk returns a tuple with the SlTotalAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTotalAmt

`func (o *GetCopytradingCopySettingsV5RespData) SetSlTotalAmt(v string)`

SetSlTotalAmt sets SlTotalAmt field to given value.

### HasSlTotalAmt

`func (o *GetCopytradingCopySettingsV5RespData) HasSlTotalAmt() bool`

HasSlTotalAmt returns a boolean if a field has been set.

### GetSubPosCloseType

`func (o *GetCopytradingCopySettingsV5RespData) GetSubPosCloseType() string`

GetSubPosCloseType returns the SubPosCloseType field if non-nil, zero value otherwise.

### GetSubPosCloseTypeOk

`func (o *GetCopytradingCopySettingsV5RespData) GetSubPosCloseTypeOk() (*string, bool)`

GetSubPosCloseTypeOk returns a tuple with the SubPosCloseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPosCloseType

`func (o *GetCopytradingCopySettingsV5RespData) SetSubPosCloseType(v string)`

SetSubPosCloseType sets SubPosCloseType field to given value.

### HasSubPosCloseType

`func (o *GetCopytradingCopySettingsV5RespData) HasSubPosCloseType() bool`

HasSubPosCloseType returns a boolean if a field has been set.

### GetTag

`func (o *GetCopytradingCopySettingsV5RespData) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetCopytradingCopySettingsV5RespData) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetCopytradingCopySettingsV5RespData) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetCopytradingCopySettingsV5RespData) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTpRatio

`func (o *GetCopytradingCopySettingsV5RespData) GetTpRatio() string`

GetTpRatio returns the TpRatio field if non-nil, zero value otherwise.

### GetTpRatioOk

`func (o *GetCopytradingCopySettingsV5RespData) GetTpRatioOk() (*string, bool)`

GetTpRatioOk returns a tuple with the TpRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpRatio

`func (o *GetCopytradingCopySettingsV5RespData) SetTpRatio(v string)`

SetTpRatio sets TpRatio field to given value.

### HasTpRatio

`func (o *GetCopytradingCopySettingsV5RespData) HasTpRatio() bool`

HasTpRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


