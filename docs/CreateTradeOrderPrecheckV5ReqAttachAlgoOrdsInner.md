# CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachAlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID when placing order attaching TP/SL  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters.  It will be posted to &#x60;algoClOrdId&#x60; when placing TP/SL order once the general order is filled completely. | [optional] [default to ""]
**SlOrdPx** | Pointer to **string** | Stop-loss order price  If you fill in this parameter, you should fill in the stop-loss trigger price.  If the price is -1, stop-loss will be executed at the market price. | [optional] [default to ""]
**SlTriggerPx** | Pointer to **string** | Stop-loss trigger price  If you fill in this parameter, you should fill in the stop-loss order price. | [optional] [default to ""]
**SlTriggerPxType** | Pointer to **string** | Stop-loss trigger price type  &#x60;last&#x60;: last price   &#x60;index&#x60;: index price   &#x60;mark&#x60;: mark price   The default is last | [optional] [default to ""]
**Sz** | Pointer to **string** | Size. Only applicable to TP order of split TPs, and it is required for TP order of split TPs | [optional] [default to ""]
**TpOrdKind** | Pointer to **string** | TP order kind  &#x60;condition&#x60;  &#x60;limit&#x60;   The default is &#x60;condition&#x60; | [optional] [default to ""]
**TpOrdPx** | Pointer to **string** | Take-profit order price     For condition TP order, if you fill in this parameter, you should fill in the take-profit trigger price as well.   For limit TP order, you need to fill in this parameter, take-profit trigger needn‘t to be filled.   If the price is -1, take-profit will be executed at the market price. | [optional] [default to ""]
**TpTriggerPx** | Pointer to **string** | Take-profit trigger price  For condition TP order, if you fill in this parameter, you should fill in the take-profit order price as well. | [optional] [default to ""]
**TpTriggerPxType** | Pointer to **string** | Take-profit trigger price type  &#x60;last&#x60;: last price   &#x60;index&#x60;: index price   &#x60;mark&#x60;: mark price   The default is last | [optional] [default to ""]

## Methods

### NewCreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner

`func NewCreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner() *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner`

NewCreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner instantiates a new CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInnerWithDefaults

`func NewCreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInnerWithDefaults() *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner`

NewCreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInnerWithDefaults instantiates a new CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachAlgoClOrdId

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) GetAttachAlgoClOrdId() string`

GetAttachAlgoClOrdId returns the AttachAlgoClOrdId field if non-nil, zero value otherwise.

### GetAttachAlgoClOrdIdOk

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) GetAttachAlgoClOrdIdOk() (*string, bool)`

GetAttachAlgoClOrdIdOk returns a tuple with the AttachAlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachAlgoClOrdId

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) SetAttachAlgoClOrdId(v string)`

SetAttachAlgoClOrdId sets AttachAlgoClOrdId field to given value.

### HasAttachAlgoClOrdId

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) HasAttachAlgoClOrdId() bool`

HasAttachAlgoClOrdId returns a boolean if a field has been set.

### GetSlOrdPx

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) GetSlOrdPx() string`

GetSlOrdPx returns the SlOrdPx field if non-nil, zero value otherwise.

### GetSlOrdPxOk

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) GetSlOrdPxOk() (*string, bool)`

GetSlOrdPxOk returns a tuple with the SlOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlOrdPx

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) SetSlOrdPx(v string)`

SetSlOrdPx sets SlOrdPx field to given value.

### HasSlOrdPx

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) HasSlOrdPx() bool`

HasSlOrdPx returns a boolean if a field has been set.

### GetSlTriggerPx

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) GetSlTriggerPx() string`

GetSlTriggerPx returns the SlTriggerPx field if non-nil, zero value otherwise.

### GetSlTriggerPxOk

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) GetSlTriggerPxOk() (*string, bool)`

GetSlTriggerPxOk returns a tuple with the SlTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPx

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) SetSlTriggerPx(v string)`

SetSlTriggerPx sets SlTriggerPx field to given value.

### HasSlTriggerPx

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) HasSlTriggerPx() bool`

HasSlTriggerPx returns a boolean if a field has been set.

### GetSlTriggerPxType

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) GetSlTriggerPxType() string`

GetSlTriggerPxType returns the SlTriggerPxType field if non-nil, zero value otherwise.

### GetSlTriggerPxTypeOk

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) GetSlTriggerPxTypeOk() (*string, bool)`

GetSlTriggerPxTypeOk returns a tuple with the SlTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPxType

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) SetSlTriggerPxType(v string)`

SetSlTriggerPxType sets SlTriggerPxType field to given value.

### HasSlTriggerPxType

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) HasSlTriggerPxType() bool`

HasSlTriggerPxType returns a boolean if a field has been set.

### GetSz

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTpOrdKind

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) GetTpOrdKind() string`

GetTpOrdKind returns the TpOrdKind field if non-nil, zero value otherwise.

### GetTpOrdKindOk

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) GetTpOrdKindOk() (*string, bool)`

GetTpOrdKindOk returns a tuple with the TpOrdKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpOrdKind

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) SetTpOrdKind(v string)`

SetTpOrdKind sets TpOrdKind field to given value.

### HasTpOrdKind

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) HasTpOrdKind() bool`

HasTpOrdKind returns a boolean if a field has been set.

### GetTpOrdPx

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) GetTpOrdPx() string`

GetTpOrdPx returns the TpOrdPx field if non-nil, zero value otherwise.

### GetTpOrdPxOk

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) GetTpOrdPxOk() (*string, bool)`

GetTpOrdPxOk returns a tuple with the TpOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpOrdPx

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) SetTpOrdPx(v string)`

SetTpOrdPx sets TpOrdPx field to given value.

### HasTpOrdPx

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) HasTpOrdPx() bool`

HasTpOrdPx returns a boolean if a field has been set.

### GetTpTriggerPx

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) GetTpTriggerPx() string`

GetTpTriggerPx returns the TpTriggerPx field if non-nil, zero value otherwise.

### GetTpTriggerPxOk

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) GetTpTriggerPxOk() (*string, bool)`

GetTpTriggerPxOk returns a tuple with the TpTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPx

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) SetTpTriggerPx(v string)`

SetTpTriggerPx sets TpTriggerPx field to given value.

### HasTpTriggerPx

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) HasTpTriggerPx() bool`

HasTpTriggerPx returns a boolean if a field has been set.

### GetTpTriggerPxType

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) GetTpTriggerPxType() string`

GetTpTriggerPxType returns the TpTriggerPxType field if non-nil, zero value otherwise.

### GetTpTriggerPxTypeOk

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) GetTpTriggerPxTypeOk() (*string, bool)`

GetTpTriggerPxTypeOk returns a tuple with the TpTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPxType

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) SetTpTriggerPxType(v string)`

SetTpTriggerPxType sets TpTriggerPxType field to given value.

### HasTpTriggerPxType

`func (o *CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner) HasTpTriggerPxType() bool`

HasTpTriggerPxType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


