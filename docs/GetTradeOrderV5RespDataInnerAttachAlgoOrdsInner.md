# GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmendPxOnTriggerType** | Pointer to **string** | Whether to enable Cost-price SL. Only applicable to SL order of split TPs.   &#x60;0&#x60;: disable, the default value   &#x60;1&#x60;: Enable | [optional] [default to ""]
**AttachAlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID when placing order attaching TP/SL  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters.  It will be posted to &#x60;algoClOrdId&#x60; when placing TP/SL order once the general order is filled completely. | [optional] [default to ""]
**AttachAlgoId** | Pointer to **string** | The order ID of attached TP/SL order. It can be used to identity the TP/SL order when amending. It will not be posted to algoId when placing TP/SL order after the general order is filled completely. | [optional] [default to ""]
**FailCode** | Pointer to **string** | The error code when failing to place TP/SL order, e.g. 51020   The default is \&quot;\&quot; | [optional] [default to ""]
**FailReason** | Pointer to **string** | The error reason when failing to place TP/SL order.   The default is \&quot;\&quot; | [optional] [default to ""]
**SlOrdPx** | Pointer to **string** | Stop-loss order price. | [optional] [default to ""]
**SlTriggerPx** | Pointer to **string** | Stop-loss trigger price. | [optional] [default to ""]
**SlTriggerPxType** | Pointer to **string** | Stop-loss trigger price type.   &#x60;last&#x60;: last price  &#x60;index&#x60;: index price  &#x60;mark&#x60;: mark price | [optional] [default to ""]
**Sz** | Pointer to **string** | Size. Only applicable to TP order of split TPs | [optional] [default to ""]
**TpOrdKind** | Pointer to **string** | TP order kind  &#x60;condition&#x60;  &#x60;limit&#x60; | [optional] [default to ""]
**TpOrdPx** | Pointer to **string** | Take-profit order price. | [optional] [default to ""]
**TpTriggerPx** | Pointer to **string** | Take-profit trigger price. | [optional] [default to ""]
**TpTriggerPxType** | Pointer to **string** | Take-profit trigger price type.   &#x60;last&#x60;: last price  &#x60;index&#x60;: index price  &#x60;mark&#x60;: mark price | [optional] [default to ""]

## Methods

### NewGetTradeOrderV5RespDataInnerAttachAlgoOrdsInner

`func NewGetTradeOrderV5RespDataInnerAttachAlgoOrdsInner() *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner`

NewGetTradeOrderV5RespDataInnerAttachAlgoOrdsInner instantiates a new GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradeOrderV5RespDataInnerAttachAlgoOrdsInnerWithDefaults

`func NewGetTradeOrderV5RespDataInnerAttachAlgoOrdsInnerWithDefaults() *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner`

NewGetTradeOrderV5RespDataInnerAttachAlgoOrdsInnerWithDefaults instantiates a new GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmendPxOnTriggerType

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetAmendPxOnTriggerType() string`

GetAmendPxOnTriggerType returns the AmendPxOnTriggerType field if non-nil, zero value otherwise.

### GetAmendPxOnTriggerTypeOk

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetAmendPxOnTriggerTypeOk() (*string, bool)`

GetAmendPxOnTriggerTypeOk returns a tuple with the AmendPxOnTriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendPxOnTriggerType

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) SetAmendPxOnTriggerType(v string)`

SetAmendPxOnTriggerType sets AmendPxOnTriggerType field to given value.

### HasAmendPxOnTriggerType

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) HasAmendPxOnTriggerType() bool`

HasAmendPxOnTriggerType returns a boolean if a field has been set.

### GetAttachAlgoClOrdId

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetAttachAlgoClOrdId() string`

GetAttachAlgoClOrdId returns the AttachAlgoClOrdId field if non-nil, zero value otherwise.

### GetAttachAlgoClOrdIdOk

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetAttachAlgoClOrdIdOk() (*string, bool)`

GetAttachAlgoClOrdIdOk returns a tuple with the AttachAlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachAlgoClOrdId

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) SetAttachAlgoClOrdId(v string)`

SetAttachAlgoClOrdId sets AttachAlgoClOrdId field to given value.

### HasAttachAlgoClOrdId

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) HasAttachAlgoClOrdId() bool`

HasAttachAlgoClOrdId returns a boolean if a field has been set.

### GetAttachAlgoId

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetAttachAlgoId() string`

GetAttachAlgoId returns the AttachAlgoId field if non-nil, zero value otherwise.

### GetAttachAlgoIdOk

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetAttachAlgoIdOk() (*string, bool)`

GetAttachAlgoIdOk returns a tuple with the AttachAlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachAlgoId

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) SetAttachAlgoId(v string)`

SetAttachAlgoId sets AttachAlgoId field to given value.

### HasAttachAlgoId

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) HasAttachAlgoId() bool`

HasAttachAlgoId returns a boolean if a field has been set.

### GetFailCode

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetFailCode() string`

GetFailCode returns the FailCode field if non-nil, zero value otherwise.

### GetFailCodeOk

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetFailCodeOk() (*string, bool)`

GetFailCodeOk returns a tuple with the FailCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailCode

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) SetFailCode(v string)`

SetFailCode sets FailCode field to given value.

### HasFailCode

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) HasFailCode() bool`

HasFailCode returns a boolean if a field has been set.

### GetFailReason

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetFailReason() string`

GetFailReason returns the FailReason field if non-nil, zero value otherwise.

### GetFailReasonOk

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetFailReasonOk() (*string, bool)`

GetFailReasonOk returns a tuple with the FailReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailReason

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) SetFailReason(v string)`

SetFailReason sets FailReason field to given value.

### HasFailReason

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) HasFailReason() bool`

HasFailReason returns a boolean if a field has been set.

### GetSlOrdPx

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetSlOrdPx() string`

GetSlOrdPx returns the SlOrdPx field if non-nil, zero value otherwise.

### GetSlOrdPxOk

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetSlOrdPxOk() (*string, bool)`

GetSlOrdPxOk returns a tuple with the SlOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlOrdPx

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) SetSlOrdPx(v string)`

SetSlOrdPx sets SlOrdPx field to given value.

### HasSlOrdPx

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) HasSlOrdPx() bool`

HasSlOrdPx returns a boolean if a field has been set.

### GetSlTriggerPx

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetSlTriggerPx() string`

GetSlTriggerPx returns the SlTriggerPx field if non-nil, zero value otherwise.

### GetSlTriggerPxOk

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetSlTriggerPxOk() (*string, bool)`

GetSlTriggerPxOk returns a tuple with the SlTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPx

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) SetSlTriggerPx(v string)`

SetSlTriggerPx sets SlTriggerPx field to given value.

### HasSlTriggerPx

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) HasSlTriggerPx() bool`

HasSlTriggerPx returns a boolean if a field has been set.

### GetSlTriggerPxType

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetSlTriggerPxType() string`

GetSlTriggerPxType returns the SlTriggerPxType field if non-nil, zero value otherwise.

### GetSlTriggerPxTypeOk

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetSlTriggerPxTypeOk() (*string, bool)`

GetSlTriggerPxTypeOk returns a tuple with the SlTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPxType

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) SetSlTriggerPxType(v string)`

SetSlTriggerPxType sets SlTriggerPxType field to given value.

### HasSlTriggerPxType

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) HasSlTriggerPxType() bool`

HasSlTriggerPxType returns a boolean if a field has been set.

### GetSz

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTpOrdKind

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetTpOrdKind() string`

GetTpOrdKind returns the TpOrdKind field if non-nil, zero value otherwise.

### GetTpOrdKindOk

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetTpOrdKindOk() (*string, bool)`

GetTpOrdKindOk returns a tuple with the TpOrdKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpOrdKind

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) SetTpOrdKind(v string)`

SetTpOrdKind sets TpOrdKind field to given value.

### HasTpOrdKind

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) HasTpOrdKind() bool`

HasTpOrdKind returns a boolean if a field has been set.

### GetTpOrdPx

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetTpOrdPx() string`

GetTpOrdPx returns the TpOrdPx field if non-nil, zero value otherwise.

### GetTpOrdPxOk

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetTpOrdPxOk() (*string, bool)`

GetTpOrdPxOk returns a tuple with the TpOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpOrdPx

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) SetTpOrdPx(v string)`

SetTpOrdPx sets TpOrdPx field to given value.

### HasTpOrdPx

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) HasTpOrdPx() bool`

HasTpOrdPx returns a boolean if a field has been set.

### GetTpTriggerPx

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetTpTriggerPx() string`

GetTpTriggerPx returns the TpTriggerPx field if non-nil, zero value otherwise.

### GetTpTriggerPxOk

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetTpTriggerPxOk() (*string, bool)`

GetTpTriggerPxOk returns a tuple with the TpTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPx

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) SetTpTriggerPx(v string)`

SetTpTriggerPx sets TpTriggerPx field to given value.

### HasTpTriggerPx

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) HasTpTriggerPx() bool`

HasTpTriggerPx returns a boolean if a field has been set.

### GetTpTriggerPxType

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetTpTriggerPxType() string`

GetTpTriggerPxType returns the TpTriggerPxType field if non-nil, zero value otherwise.

### GetTpTriggerPxTypeOk

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) GetTpTriggerPxTypeOk() (*string, bool)`

GetTpTriggerPxTypeOk returns a tuple with the TpTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPxType

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) SetTpTriggerPxType(v string)`

SetTpTriggerPxType sets TpTriggerPxType field to given value.

### HasTpTriggerPxType

`func (o *GetTradeOrderV5RespDataInnerAttachAlgoOrdsInner) HasTpTriggerPxType() bool`

HasTpTriggerPxType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


