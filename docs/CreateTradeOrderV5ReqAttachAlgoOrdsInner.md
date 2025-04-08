# CreateTradeOrderV5ReqAttachAlgoOrdsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmendPxOnTriggerType** | Pointer to **string** | Whether to enable Cost-price SL. Only applicable to SL order of split TPs. Whether &#x60;slTriggerPx&#x60; will move to &#x60;avgPx&#x60; when the first TP order is triggered  &#x60;0&#x60;: disable, the default value   &#x60;1&#x60;: Enable | [optional] [default to ""]
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

### NewCreateTradeOrderV5ReqAttachAlgoOrdsInner

`func NewCreateTradeOrderV5ReqAttachAlgoOrdsInner() *CreateTradeOrderV5ReqAttachAlgoOrdsInner`

NewCreateTradeOrderV5ReqAttachAlgoOrdsInner instantiates a new CreateTradeOrderV5ReqAttachAlgoOrdsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeOrderV5ReqAttachAlgoOrdsInnerWithDefaults

`func NewCreateTradeOrderV5ReqAttachAlgoOrdsInnerWithDefaults() *CreateTradeOrderV5ReqAttachAlgoOrdsInner`

NewCreateTradeOrderV5ReqAttachAlgoOrdsInnerWithDefaults instantiates a new CreateTradeOrderV5ReqAttachAlgoOrdsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmendPxOnTriggerType

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetAmendPxOnTriggerType() string`

GetAmendPxOnTriggerType returns the AmendPxOnTriggerType field if non-nil, zero value otherwise.

### GetAmendPxOnTriggerTypeOk

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetAmendPxOnTriggerTypeOk() (*string, bool)`

GetAmendPxOnTriggerTypeOk returns a tuple with the AmendPxOnTriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendPxOnTriggerType

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) SetAmendPxOnTriggerType(v string)`

SetAmendPxOnTriggerType sets AmendPxOnTriggerType field to given value.

### HasAmendPxOnTriggerType

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) HasAmendPxOnTriggerType() bool`

HasAmendPxOnTriggerType returns a boolean if a field has been set.

### GetAttachAlgoClOrdId

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetAttachAlgoClOrdId() string`

GetAttachAlgoClOrdId returns the AttachAlgoClOrdId field if non-nil, zero value otherwise.

### GetAttachAlgoClOrdIdOk

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetAttachAlgoClOrdIdOk() (*string, bool)`

GetAttachAlgoClOrdIdOk returns a tuple with the AttachAlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachAlgoClOrdId

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) SetAttachAlgoClOrdId(v string)`

SetAttachAlgoClOrdId sets AttachAlgoClOrdId field to given value.

### HasAttachAlgoClOrdId

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) HasAttachAlgoClOrdId() bool`

HasAttachAlgoClOrdId returns a boolean if a field has been set.

### GetSlOrdPx

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetSlOrdPx() string`

GetSlOrdPx returns the SlOrdPx field if non-nil, zero value otherwise.

### GetSlOrdPxOk

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetSlOrdPxOk() (*string, bool)`

GetSlOrdPxOk returns a tuple with the SlOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlOrdPx

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) SetSlOrdPx(v string)`

SetSlOrdPx sets SlOrdPx field to given value.

### HasSlOrdPx

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) HasSlOrdPx() bool`

HasSlOrdPx returns a boolean if a field has been set.

### GetSlTriggerPx

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetSlTriggerPx() string`

GetSlTriggerPx returns the SlTriggerPx field if non-nil, zero value otherwise.

### GetSlTriggerPxOk

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetSlTriggerPxOk() (*string, bool)`

GetSlTriggerPxOk returns a tuple with the SlTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPx

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) SetSlTriggerPx(v string)`

SetSlTriggerPx sets SlTriggerPx field to given value.

### HasSlTriggerPx

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) HasSlTriggerPx() bool`

HasSlTriggerPx returns a boolean if a field has been set.

### GetSlTriggerPxType

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetSlTriggerPxType() string`

GetSlTriggerPxType returns the SlTriggerPxType field if non-nil, zero value otherwise.

### GetSlTriggerPxTypeOk

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetSlTriggerPxTypeOk() (*string, bool)`

GetSlTriggerPxTypeOk returns a tuple with the SlTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPxType

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) SetSlTriggerPxType(v string)`

SetSlTriggerPxType sets SlTriggerPxType field to given value.

### HasSlTriggerPxType

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) HasSlTriggerPxType() bool`

HasSlTriggerPxType returns a boolean if a field has been set.

### GetSz

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTpOrdKind

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetTpOrdKind() string`

GetTpOrdKind returns the TpOrdKind field if non-nil, zero value otherwise.

### GetTpOrdKindOk

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetTpOrdKindOk() (*string, bool)`

GetTpOrdKindOk returns a tuple with the TpOrdKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpOrdKind

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) SetTpOrdKind(v string)`

SetTpOrdKind sets TpOrdKind field to given value.

### HasTpOrdKind

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) HasTpOrdKind() bool`

HasTpOrdKind returns a boolean if a field has been set.

### GetTpOrdPx

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetTpOrdPx() string`

GetTpOrdPx returns the TpOrdPx field if non-nil, zero value otherwise.

### GetTpOrdPxOk

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetTpOrdPxOk() (*string, bool)`

GetTpOrdPxOk returns a tuple with the TpOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpOrdPx

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) SetTpOrdPx(v string)`

SetTpOrdPx sets TpOrdPx field to given value.

### HasTpOrdPx

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) HasTpOrdPx() bool`

HasTpOrdPx returns a boolean if a field has been set.

### GetTpTriggerPx

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetTpTriggerPx() string`

GetTpTriggerPx returns the TpTriggerPx field if non-nil, zero value otherwise.

### GetTpTriggerPxOk

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetTpTriggerPxOk() (*string, bool)`

GetTpTriggerPxOk returns a tuple with the TpTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPx

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) SetTpTriggerPx(v string)`

SetTpTriggerPx sets TpTriggerPx field to given value.

### HasTpTriggerPx

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) HasTpTriggerPx() bool`

HasTpTriggerPx returns a boolean if a field has been set.

### GetTpTriggerPxType

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetTpTriggerPxType() string`

GetTpTriggerPxType returns the TpTriggerPxType field if non-nil, zero value otherwise.

### GetTpTriggerPxTypeOk

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) GetTpTriggerPxTypeOk() (*string, bool)`

GetTpTriggerPxTypeOk returns a tuple with the TpTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPxType

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) SetTpTriggerPxType(v string)`

SetTpTriggerPxType sets TpTriggerPxType field to given value.

### HasTpTriggerPxType

`func (o *CreateTradeOrderV5ReqAttachAlgoOrdsInner) HasTpTriggerPxType() bool`

HasTpTriggerPxType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


