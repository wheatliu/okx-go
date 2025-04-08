# GetAccountAdjustLeverageInfoV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstAvailQuoteTrans** | Pointer to **string** | The estimated margin(in quote currency) can be transferred out under the corresponding leverage  For cross, it is the maximum quantity that can be transferred from the trading account.  For isolated, it is the maximum quantity that can be transferred from the isolated position  Only applicable to &#x60;MARGIN&#x60; | [optional] [default to ""]
**EstAvailTrans** | Pointer to **string** | The estimated margin can be transferred out under the corresponding leverage.  For cross, it is the maximum quantity that can be transferred from the trading account.  For isolated, it is the maximum quantity that can be transferred from the isolated position   The unit is base currency for &#x60;MARGIN&#x60;   It is not applicable to the scenario when increasing leverage for isolated position under &#x60;FUTURES&#x60; and &#x60;SWAP&#x60; | [optional] [default to ""]
**EstLiqPx** | Pointer to **string** | The estimated liquidation price under the corresponding leverage. Only return when there is a position. | [optional] [default to ""]
**EstMaxAmt** | Pointer to **string** | For &#x60;MARGIN&#x60;, it is the estimated maximum loan in base currency under the corresponding leverage  For &#x60;SWAP&#x60; and &#x60;FUTURES&#x60;, it is the estimated maximum quantity of contracts that can be opened under the corresponding leverage | [optional] [default to ""]
**EstMgn** | Pointer to **string** | The estimated margin needed by position under the corresponding leverage.  For the &#x60;MARGIN&#x60; position, it is margin in base currency | [optional] [default to ""]
**EstQuoteMaxAmt** | Pointer to **string** | The &#x60;MARGIN&#x60; estimated maximum loan in quote currency under the corresponding leverage. | [optional] [default to ""]
**EstQuoteMgn** | Pointer to **string** | The estimated margin (in quote currency) needed by position under the corresponding leverage | [optional] [default to ""]
**ExistOrd** | Pointer to **bool** | Whether there is pending orders   &#x60;true&#x60;  &#x60;false&#x60; | [optional] 
**MaxLever** | Pointer to **string** | Maximum leverage | [optional] [default to ""]
**MinLever** | Pointer to **string** | Minimum leverage | [optional] [default to ""]

## Methods

### NewGetAccountAdjustLeverageInfoV5RespDataInner

`func NewGetAccountAdjustLeverageInfoV5RespDataInner() *GetAccountAdjustLeverageInfoV5RespDataInner`

NewGetAccountAdjustLeverageInfoV5RespDataInner instantiates a new GetAccountAdjustLeverageInfoV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountAdjustLeverageInfoV5RespDataInnerWithDefaults

`func NewGetAccountAdjustLeverageInfoV5RespDataInnerWithDefaults() *GetAccountAdjustLeverageInfoV5RespDataInner`

NewGetAccountAdjustLeverageInfoV5RespDataInnerWithDefaults instantiates a new GetAccountAdjustLeverageInfoV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstAvailQuoteTrans

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) GetEstAvailQuoteTrans() string`

GetEstAvailQuoteTrans returns the EstAvailQuoteTrans field if non-nil, zero value otherwise.

### GetEstAvailQuoteTransOk

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) GetEstAvailQuoteTransOk() (*string, bool)`

GetEstAvailQuoteTransOk returns a tuple with the EstAvailQuoteTrans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstAvailQuoteTrans

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) SetEstAvailQuoteTrans(v string)`

SetEstAvailQuoteTrans sets EstAvailQuoteTrans field to given value.

### HasEstAvailQuoteTrans

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) HasEstAvailQuoteTrans() bool`

HasEstAvailQuoteTrans returns a boolean if a field has been set.

### GetEstAvailTrans

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) GetEstAvailTrans() string`

GetEstAvailTrans returns the EstAvailTrans field if non-nil, zero value otherwise.

### GetEstAvailTransOk

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) GetEstAvailTransOk() (*string, bool)`

GetEstAvailTransOk returns a tuple with the EstAvailTrans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstAvailTrans

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) SetEstAvailTrans(v string)`

SetEstAvailTrans sets EstAvailTrans field to given value.

### HasEstAvailTrans

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) HasEstAvailTrans() bool`

HasEstAvailTrans returns a boolean if a field has been set.

### GetEstLiqPx

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) GetEstLiqPx() string`

GetEstLiqPx returns the EstLiqPx field if non-nil, zero value otherwise.

### GetEstLiqPxOk

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) GetEstLiqPxOk() (*string, bool)`

GetEstLiqPxOk returns a tuple with the EstLiqPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstLiqPx

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) SetEstLiqPx(v string)`

SetEstLiqPx sets EstLiqPx field to given value.

### HasEstLiqPx

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) HasEstLiqPx() bool`

HasEstLiqPx returns a boolean if a field has been set.

### GetEstMaxAmt

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) GetEstMaxAmt() string`

GetEstMaxAmt returns the EstMaxAmt field if non-nil, zero value otherwise.

### GetEstMaxAmtOk

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) GetEstMaxAmtOk() (*string, bool)`

GetEstMaxAmtOk returns a tuple with the EstMaxAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstMaxAmt

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) SetEstMaxAmt(v string)`

SetEstMaxAmt sets EstMaxAmt field to given value.

### HasEstMaxAmt

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) HasEstMaxAmt() bool`

HasEstMaxAmt returns a boolean if a field has been set.

### GetEstMgn

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) GetEstMgn() string`

GetEstMgn returns the EstMgn field if non-nil, zero value otherwise.

### GetEstMgnOk

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) GetEstMgnOk() (*string, bool)`

GetEstMgnOk returns a tuple with the EstMgn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstMgn

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) SetEstMgn(v string)`

SetEstMgn sets EstMgn field to given value.

### HasEstMgn

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) HasEstMgn() bool`

HasEstMgn returns a boolean if a field has been set.

### GetEstQuoteMaxAmt

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) GetEstQuoteMaxAmt() string`

GetEstQuoteMaxAmt returns the EstQuoteMaxAmt field if non-nil, zero value otherwise.

### GetEstQuoteMaxAmtOk

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) GetEstQuoteMaxAmtOk() (*string, bool)`

GetEstQuoteMaxAmtOk returns a tuple with the EstQuoteMaxAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstQuoteMaxAmt

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) SetEstQuoteMaxAmt(v string)`

SetEstQuoteMaxAmt sets EstQuoteMaxAmt field to given value.

### HasEstQuoteMaxAmt

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) HasEstQuoteMaxAmt() bool`

HasEstQuoteMaxAmt returns a boolean if a field has been set.

### GetEstQuoteMgn

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) GetEstQuoteMgn() string`

GetEstQuoteMgn returns the EstQuoteMgn field if non-nil, zero value otherwise.

### GetEstQuoteMgnOk

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) GetEstQuoteMgnOk() (*string, bool)`

GetEstQuoteMgnOk returns a tuple with the EstQuoteMgn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstQuoteMgn

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) SetEstQuoteMgn(v string)`

SetEstQuoteMgn sets EstQuoteMgn field to given value.

### HasEstQuoteMgn

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) HasEstQuoteMgn() bool`

HasEstQuoteMgn returns a boolean if a field has been set.

### GetExistOrd

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) GetExistOrd() bool`

GetExistOrd returns the ExistOrd field if non-nil, zero value otherwise.

### GetExistOrdOk

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) GetExistOrdOk() (*bool, bool)`

GetExistOrdOk returns a tuple with the ExistOrd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistOrd

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) SetExistOrd(v bool)`

SetExistOrd sets ExistOrd field to given value.

### HasExistOrd

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) HasExistOrd() bool`

HasExistOrd returns a boolean if a field has been set.

### GetMaxLever

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) GetMaxLever() string`

GetMaxLever returns the MaxLever field if non-nil, zero value otherwise.

### GetMaxLeverOk

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) GetMaxLeverOk() (*string, bool)`

GetMaxLeverOk returns a tuple with the MaxLever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLever

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) SetMaxLever(v string)`

SetMaxLever sets MaxLever field to given value.

### HasMaxLever

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) HasMaxLever() bool`

HasMaxLever returns a boolean if a field has been set.

### GetMinLever

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) GetMinLever() string`

GetMinLever returns the MinLever field if non-nil, zero value otherwise.

### GetMinLeverOk

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) GetMinLeverOk() (*string, bool)`

GetMinLeverOk returns a tuple with the MinLever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLever

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) SetMinLever(v string)`

SetMinLever sets MinLever field to given value.

### HasMinLever

`func (o *GetAccountAdjustLeverageInfoV5RespDataInner) HasMinLever() bool`

HasMinLever returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


