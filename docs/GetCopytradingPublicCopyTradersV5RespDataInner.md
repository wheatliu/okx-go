# GetCopytradingPublicCopyTradersV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | Pointer to **string** | The currency name of profit and loss | [optional] [default to ""]
**CopyTotalPnl** | Pointer to **string** | Total copy trader profit and loss | [optional] [default to ""]
**CopyTraderNumChg** | Pointer to **string** | Number change in last 7 days | [optional] [default to ""]
**CopyTraderNumChgRatio** | Pointer to **string** | Ratio change in last 7 days | [optional] [default to ""]
**CopyTraders** | Pointer to [**[]GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner**](GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner.md) | Copy trader information | [optional] 

## Methods

### NewGetCopytradingPublicCopyTradersV5RespDataInner

`func NewGetCopytradingPublicCopyTradersV5RespDataInner() *GetCopytradingPublicCopyTradersV5RespDataInner`

NewGetCopytradingPublicCopyTradersV5RespDataInner instantiates a new GetCopytradingPublicCopyTradersV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCopytradingPublicCopyTradersV5RespDataInnerWithDefaults

`func NewGetCopytradingPublicCopyTradersV5RespDataInnerWithDefaults() *GetCopytradingPublicCopyTradersV5RespDataInner`

NewGetCopytradingPublicCopyTradersV5RespDataInnerWithDefaults instantiates a new GetCopytradingPublicCopyTradersV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *GetCopytradingPublicCopyTradersV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetCopytradingPublicCopyTradersV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetCopytradingPublicCopyTradersV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetCopytradingPublicCopyTradersV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetCopyTotalPnl

`func (o *GetCopytradingPublicCopyTradersV5RespDataInner) GetCopyTotalPnl() string`

GetCopyTotalPnl returns the CopyTotalPnl field if non-nil, zero value otherwise.

### GetCopyTotalPnlOk

`func (o *GetCopytradingPublicCopyTradersV5RespDataInner) GetCopyTotalPnlOk() (*string, bool)`

GetCopyTotalPnlOk returns a tuple with the CopyTotalPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTotalPnl

`func (o *GetCopytradingPublicCopyTradersV5RespDataInner) SetCopyTotalPnl(v string)`

SetCopyTotalPnl sets CopyTotalPnl field to given value.

### HasCopyTotalPnl

`func (o *GetCopytradingPublicCopyTradersV5RespDataInner) HasCopyTotalPnl() bool`

HasCopyTotalPnl returns a boolean if a field has been set.

### GetCopyTraderNumChg

`func (o *GetCopytradingPublicCopyTradersV5RespDataInner) GetCopyTraderNumChg() string`

GetCopyTraderNumChg returns the CopyTraderNumChg field if non-nil, zero value otherwise.

### GetCopyTraderNumChgOk

`func (o *GetCopytradingPublicCopyTradersV5RespDataInner) GetCopyTraderNumChgOk() (*string, bool)`

GetCopyTraderNumChgOk returns a tuple with the CopyTraderNumChg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTraderNumChg

`func (o *GetCopytradingPublicCopyTradersV5RespDataInner) SetCopyTraderNumChg(v string)`

SetCopyTraderNumChg sets CopyTraderNumChg field to given value.

### HasCopyTraderNumChg

`func (o *GetCopytradingPublicCopyTradersV5RespDataInner) HasCopyTraderNumChg() bool`

HasCopyTraderNumChg returns a boolean if a field has been set.

### GetCopyTraderNumChgRatio

`func (o *GetCopytradingPublicCopyTradersV5RespDataInner) GetCopyTraderNumChgRatio() string`

GetCopyTraderNumChgRatio returns the CopyTraderNumChgRatio field if non-nil, zero value otherwise.

### GetCopyTraderNumChgRatioOk

`func (o *GetCopytradingPublicCopyTradersV5RespDataInner) GetCopyTraderNumChgRatioOk() (*string, bool)`

GetCopyTraderNumChgRatioOk returns a tuple with the CopyTraderNumChgRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTraderNumChgRatio

`func (o *GetCopytradingPublicCopyTradersV5RespDataInner) SetCopyTraderNumChgRatio(v string)`

SetCopyTraderNumChgRatio sets CopyTraderNumChgRatio field to given value.

### HasCopyTraderNumChgRatio

`func (o *GetCopytradingPublicCopyTradersV5RespDataInner) HasCopyTraderNumChgRatio() bool`

HasCopyTraderNumChgRatio returns a boolean if a field has been set.

### GetCopyTraders

`func (o *GetCopytradingPublicCopyTradersV5RespDataInner) GetCopyTraders() []GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner`

GetCopyTraders returns the CopyTraders field if non-nil, zero value otherwise.

### GetCopyTradersOk

`func (o *GetCopytradingPublicCopyTradersV5RespDataInner) GetCopyTradersOk() (*[]GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner, bool)`

GetCopyTradersOk returns a tuple with the CopyTraders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTraders

`func (o *GetCopytradingPublicCopyTradersV5RespDataInner) SetCopyTraders(v []GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner)`

SetCopyTraders sets CopyTraders field to given value.

### HasCopyTraders

`func (o *GetCopytradingPublicCopyTradersV5RespDataInner) HasCopyTraders() bool`

HasCopyTraders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


