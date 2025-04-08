# CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmendPxOnTriggerType** | Pointer to **string** | Whether to enable Cost-price SL. Only applicable to SL order of split TPs.   &#x60;0&#x60;: disable, the default value   &#x60;1&#x60;: Enable | [optional] [default to ""]
**AttachAlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID when placing order attaching TP/SL  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters.  It will be posted to &#x60;algoClOrdId&#x60; when placing TP/SL order once the general order is filled completely. | [optional] [default to ""]
**AttachAlgoId** | Pointer to **string** | The order ID of attached TP/SL order. It can be used to identity the TP/SL order when amending. It will not be posted to algoId when placing TP/SL order after the general order is filled completely. | [optional] [default to ""]
**NewSlOrdPx** | Pointer to **string** | Stop-loss order price  If the price is -1, stop-loss will be executed at the market price. | [optional] [default to ""]
**NewSlTriggerPx** | Pointer to **string** | Stop-loss trigger price  Either the stop loss trigger price or order price is 0, it means that the stop loss is deleted. | [optional] [default to ""]
**NewSlTriggerPxType** | Pointer to **string** | Stop-loss trigger price type  &#x60;last&#x60;: last price   &#x60;index&#x60;: index price   &#x60;mark&#x60;: mark price  Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;  If you want to add the stop-loss, this parameter is required | [optional] [default to ""]
**NewTpOrdKind** | Pointer to **string** | TP order kind  &#x60;condition&#x60;  &#x60;limit&#x60; | [optional] [default to ""]
**NewTpOrdPx** | Pointer to **string** | Take-profit order price  If the price is -1, take-profit will be executed at the market price. | [optional] [default to ""]
**NewTpTriggerPx** | Pointer to **string** | Take-profit trigger price.   Either the take profit trigger price or order price is 0, it means that the take profit is deleted. | [optional] [default to ""]
**NewTpTriggerPxType** | Pointer to **string** | Take-profit trigger price type  &#x60;last&#x60;: last price   &#x60;index&#x60;: index price   &#x60;mark&#x60;: mark price  Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;  If you want to add the take-profit, this parameter is required | [optional] [default to ""]
**Sz** | Pointer to **string** | New size. Only applicable to TP order of split TPs, and it is required for TP order of split TPs | [optional] [default to ""]

## Methods

### NewCreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner

`func NewCreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner() *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner`

NewCreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner instantiates a new CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInnerWithDefaults

`func NewCreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInnerWithDefaults() *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner`

NewCreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInnerWithDefaults instantiates a new CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmendPxOnTriggerType

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetAmendPxOnTriggerType() string`

GetAmendPxOnTriggerType returns the AmendPxOnTriggerType field if non-nil, zero value otherwise.

### GetAmendPxOnTriggerTypeOk

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetAmendPxOnTriggerTypeOk() (*string, bool)`

GetAmendPxOnTriggerTypeOk returns a tuple with the AmendPxOnTriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendPxOnTriggerType

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) SetAmendPxOnTriggerType(v string)`

SetAmendPxOnTriggerType sets AmendPxOnTriggerType field to given value.

### HasAmendPxOnTriggerType

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) HasAmendPxOnTriggerType() bool`

HasAmendPxOnTriggerType returns a boolean if a field has been set.

### GetAttachAlgoClOrdId

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetAttachAlgoClOrdId() string`

GetAttachAlgoClOrdId returns the AttachAlgoClOrdId field if non-nil, zero value otherwise.

### GetAttachAlgoClOrdIdOk

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetAttachAlgoClOrdIdOk() (*string, bool)`

GetAttachAlgoClOrdIdOk returns a tuple with the AttachAlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachAlgoClOrdId

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) SetAttachAlgoClOrdId(v string)`

SetAttachAlgoClOrdId sets AttachAlgoClOrdId field to given value.

### HasAttachAlgoClOrdId

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) HasAttachAlgoClOrdId() bool`

HasAttachAlgoClOrdId returns a boolean if a field has been set.

### GetAttachAlgoId

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetAttachAlgoId() string`

GetAttachAlgoId returns the AttachAlgoId field if non-nil, zero value otherwise.

### GetAttachAlgoIdOk

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetAttachAlgoIdOk() (*string, bool)`

GetAttachAlgoIdOk returns a tuple with the AttachAlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachAlgoId

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) SetAttachAlgoId(v string)`

SetAttachAlgoId sets AttachAlgoId field to given value.

### HasAttachAlgoId

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) HasAttachAlgoId() bool`

HasAttachAlgoId returns a boolean if a field has been set.

### GetNewSlOrdPx

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetNewSlOrdPx() string`

GetNewSlOrdPx returns the NewSlOrdPx field if non-nil, zero value otherwise.

### GetNewSlOrdPxOk

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetNewSlOrdPxOk() (*string, bool)`

GetNewSlOrdPxOk returns a tuple with the NewSlOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSlOrdPx

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) SetNewSlOrdPx(v string)`

SetNewSlOrdPx sets NewSlOrdPx field to given value.

### HasNewSlOrdPx

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) HasNewSlOrdPx() bool`

HasNewSlOrdPx returns a boolean if a field has been set.

### GetNewSlTriggerPx

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetNewSlTriggerPx() string`

GetNewSlTriggerPx returns the NewSlTriggerPx field if non-nil, zero value otherwise.

### GetNewSlTriggerPxOk

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetNewSlTriggerPxOk() (*string, bool)`

GetNewSlTriggerPxOk returns a tuple with the NewSlTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSlTriggerPx

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) SetNewSlTriggerPx(v string)`

SetNewSlTriggerPx sets NewSlTriggerPx field to given value.

### HasNewSlTriggerPx

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) HasNewSlTriggerPx() bool`

HasNewSlTriggerPx returns a boolean if a field has been set.

### GetNewSlTriggerPxType

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetNewSlTriggerPxType() string`

GetNewSlTriggerPxType returns the NewSlTriggerPxType field if non-nil, zero value otherwise.

### GetNewSlTriggerPxTypeOk

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetNewSlTriggerPxTypeOk() (*string, bool)`

GetNewSlTriggerPxTypeOk returns a tuple with the NewSlTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSlTriggerPxType

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) SetNewSlTriggerPxType(v string)`

SetNewSlTriggerPxType sets NewSlTriggerPxType field to given value.

### HasNewSlTriggerPxType

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) HasNewSlTriggerPxType() bool`

HasNewSlTriggerPxType returns a boolean if a field has been set.

### GetNewTpOrdKind

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetNewTpOrdKind() string`

GetNewTpOrdKind returns the NewTpOrdKind field if non-nil, zero value otherwise.

### GetNewTpOrdKindOk

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetNewTpOrdKindOk() (*string, bool)`

GetNewTpOrdKindOk returns a tuple with the NewTpOrdKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTpOrdKind

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) SetNewTpOrdKind(v string)`

SetNewTpOrdKind sets NewTpOrdKind field to given value.

### HasNewTpOrdKind

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) HasNewTpOrdKind() bool`

HasNewTpOrdKind returns a boolean if a field has been set.

### GetNewTpOrdPx

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetNewTpOrdPx() string`

GetNewTpOrdPx returns the NewTpOrdPx field if non-nil, zero value otherwise.

### GetNewTpOrdPxOk

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetNewTpOrdPxOk() (*string, bool)`

GetNewTpOrdPxOk returns a tuple with the NewTpOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTpOrdPx

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) SetNewTpOrdPx(v string)`

SetNewTpOrdPx sets NewTpOrdPx field to given value.

### HasNewTpOrdPx

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) HasNewTpOrdPx() bool`

HasNewTpOrdPx returns a boolean if a field has been set.

### GetNewTpTriggerPx

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetNewTpTriggerPx() string`

GetNewTpTriggerPx returns the NewTpTriggerPx field if non-nil, zero value otherwise.

### GetNewTpTriggerPxOk

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetNewTpTriggerPxOk() (*string, bool)`

GetNewTpTriggerPxOk returns a tuple with the NewTpTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTpTriggerPx

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) SetNewTpTriggerPx(v string)`

SetNewTpTriggerPx sets NewTpTriggerPx field to given value.

### HasNewTpTriggerPx

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) HasNewTpTriggerPx() bool`

HasNewTpTriggerPx returns a boolean if a field has been set.

### GetNewTpTriggerPxType

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetNewTpTriggerPxType() string`

GetNewTpTriggerPxType returns the NewTpTriggerPxType field if non-nil, zero value otherwise.

### GetNewTpTriggerPxTypeOk

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetNewTpTriggerPxTypeOk() (*string, bool)`

GetNewTpTriggerPxTypeOk returns a tuple with the NewTpTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTpTriggerPxType

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) SetNewTpTriggerPxType(v string)`

SetNewTpTriggerPxType sets NewTpTriggerPxType field to given value.

### HasNewTpTriggerPxType

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) HasNewTpTriggerPxType() bool`

HasNewTpTriggerPxType returns a boolean if a field has been set.

### GetSz

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner) HasSz() bool`

HasSz returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


