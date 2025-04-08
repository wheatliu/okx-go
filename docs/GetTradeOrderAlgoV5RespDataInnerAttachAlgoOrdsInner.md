# GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachAlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID when placing order attaching TP/SL.  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters.  It will be posted to algoClOrdId when placing TP/SL order once the general order is filled completely. | [optional] [default to ""]
**SlOrdPx** | Pointer to **string** | Stop-loss order price   If you fill in this parameter, you should fill in the stop-loss trigger price.   If the price is &#x60;-1&#x60;, stop-loss will be executed at the market price. | [optional] [default to ""]
**SlTriggerPx** | Pointer to **string** | Stop-loss trigger price  If you fill in this parameter, you should fill in the stop-loss order price. | [optional] [default to ""]
**SlTriggerPxType** | Pointer to **string** | Stop-loss trigger price type  &#x60;last&#x60;: last price  &#x60;index&#x60;: index price  &#x60;mark&#x60;: mark price | [optional] [default to ""]
**TpOrdPx** | Pointer to **string** | Take-profit order price  If you fill in this parameter, you should fill in the take-profit trigger price as well.   If the price is &#x60;-1&#x60;, take-profit will be executed at the market price. | [optional] [default to ""]
**TpTriggerPx** | Pointer to **string** | Take-profit trigger price  If you fill in this parameter, you should fill in the take-profit order price as well. | [optional] [default to ""]
**TpTriggerPxType** | Pointer to **string** | Take-profit trigger price type  &#x60;last&#x60;: last price  &#x60;index&#x60;: index price  &#x60;mark&#x60;: mark price | [optional] [default to ""]

## Methods

### NewGetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner

`func NewGetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner() *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner`

NewGetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner instantiates a new GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInnerWithDefaults

`func NewGetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInnerWithDefaults() *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner`

NewGetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInnerWithDefaults instantiates a new GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachAlgoClOrdId

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) GetAttachAlgoClOrdId() string`

GetAttachAlgoClOrdId returns the AttachAlgoClOrdId field if non-nil, zero value otherwise.

### GetAttachAlgoClOrdIdOk

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) GetAttachAlgoClOrdIdOk() (*string, bool)`

GetAttachAlgoClOrdIdOk returns a tuple with the AttachAlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachAlgoClOrdId

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) SetAttachAlgoClOrdId(v string)`

SetAttachAlgoClOrdId sets AttachAlgoClOrdId field to given value.

### HasAttachAlgoClOrdId

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) HasAttachAlgoClOrdId() bool`

HasAttachAlgoClOrdId returns a boolean if a field has been set.

### GetSlOrdPx

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) GetSlOrdPx() string`

GetSlOrdPx returns the SlOrdPx field if non-nil, zero value otherwise.

### GetSlOrdPxOk

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) GetSlOrdPxOk() (*string, bool)`

GetSlOrdPxOk returns a tuple with the SlOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlOrdPx

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) SetSlOrdPx(v string)`

SetSlOrdPx sets SlOrdPx field to given value.

### HasSlOrdPx

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) HasSlOrdPx() bool`

HasSlOrdPx returns a boolean if a field has been set.

### GetSlTriggerPx

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) GetSlTriggerPx() string`

GetSlTriggerPx returns the SlTriggerPx field if non-nil, zero value otherwise.

### GetSlTriggerPxOk

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) GetSlTriggerPxOk() (*string, bool)`

GetSlTriggerPxOk returns a tuple with the SlTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPx

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) SetSlTriggerPx(v string)`

SetSlTriggerPx sets SlTriggerPx field to given value.

### HasSlTriggerPx

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) HasSlTriggerPx() bool`

HasSlTriggerPx returns a boolean if a field has been set.

### GetSlTriggerPxType

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) GetSlTriggerPxType() string`

GetSlTriggerPxType returns the SlTriggerPxType field if non-nil, zero value otherwise.

### GetSlTriggerPxTypeOk

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) GetSlTriggerPxTypeOk() (*string, bool)`

GetSlTriggerPxTypeOk returns a tuple with the SlTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPxType

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) SetSlTriggerPxType(v string)`

SetSlTriggerPxType sets SlTriggerPxType field to given value.

### HasSlTriggerPxType

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) HasSlTriggerPxType() bool`

HasSlTriggerPxType returns a boolean if a field has been set.

### GetTpOrdPx

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) GetTpOrdPx() string`

GetTpOrdPx returns the TpOrdPx field if non-nil, zero value otherwise.

### GetTpOrdPxOk

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) GetTpOrdPxOk() (*string, bool)`

GetTpOrdPxOk returns a tuple with the TpOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpOrdPx

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) SetTpOrdPx(v string)`

SetTpOrdPx sets TpOrdPx field to given value.

### HasTpOrdPx

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) HasTpOrdPx() bool`

HasTpOrdPx returns a boolean if a field has been set.

### GetTpTriggerPx

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) GetTpTriggerPx() string`

GetTpTriggerPx returns the TpTriggerPx field if non-nil, zero value otherwise.

### GetTpTriggerPxOk

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) GetTpTriggerPxOk() (*string, bool)`

GetTpTriggerPxOk returns a tuple with the TpTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPx

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) SetTpTriggerPx(v string)`

SetTpTriggerPx sets TpTriggerPx field to given value.

### HasTpTriggerPx

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) HasTpTriggerPx() bool`

HasTpTriggerPx returns a boolean if a field has been set.

### GetTpTriggerPxType

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) GetTpTriggerPxType() string`

GetTpTriggerPxType returns the TpTriggerPxType field if non-nil, zero value otherwise.

### GetTpTriggerPxTypeOk

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) GetTpTriggerPxTypeOk() (*string, bool)`

GetTpTriggerPxTypeOk returns a tuple with the TpTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPxType

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) SetTpTriggerPxType(v string)`

SetTpTriggerPxType sets TpTriggerPxType field to given value.

### HasTpTriggerPxType

`func (o *GetTradeOrderAlgoV5RespDataInnerAttachAlgoOrdsInner) HasTpTriggerPxType() bool`

HasTpTriggerPxType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


