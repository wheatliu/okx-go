# CreateCopytradingAlgoOrderV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstType** | Pointer to **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;SWAP&#x60;, the default value | [optional] [default to ""]
**SlOrdPx** | Pointer to **string** | Stop-loss order price  If the price is -1, stop-loss will be executed at the market price, the default is &#x60;-1&#x60;  Only applicable to &#x60;SPOT&#x60; lead trader | [optional] [default to ""]
**SlTriggerPx** | Pointer to **string** | Stop-loss trigger price. Stop-loss order price will be the market price after triggering. The stop loss order will be deleted if it is 0 | [optional] [default to ""]
**SlTriggerPxType** | Pointer to **string** | Stop-loss trigger price type  &#x60;last&#x60;: last price   &#x60;index&#x60;: index price   &#x60;mark&#x60;: mark price   Default is last | [optional] [default to ""]
**SubPosId** | **string** | Lead position ID | [default to ""]
**Tag** | Pointer to **string** | Order tag  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 16 characters. | [optional] [default to ""]
**TpOrdPx** | Pointer to **string** | Take-profit order price  If the price is -1, take-profit will be executed at the market price, the default is &#x60;-1&#x60;  Only applicable to &#x60;SPOT&#x60; lead trader | [optional] [default to ""]
**TpTriggerPx** | Pointer to **string** | Take-profit trigger price. Take-profit order price will be the market price after triggering. At least one of tpTriggerPx and slTriggerPx must be filled  The take profit order will be deleted if it is 0 | [optional] [default to ""]
**TpTriggerPxType** | Pointer to **string** | Take-profit trigger price type     &#x60;last&#x60;: last price  &#x60;index&#x60;: index price  &#x60;mark&#x60;: mark price   Default is &#x60;last&#x60; | [optional] [default to ""]

## Methods

### NewCreateCopytradingAlgoOrderV5Req

`func NewCreateCopytradingAlgoOrderV5Req(subPosId string, ) *CreateCopytradingAlgoOrderV5Req`

NewCreateCopytradingAlgoOrderV5Req instantiates a new CreateCopytradingAlgoOrderV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCopytradingAlgoOrderV5ReqWithDefaults

`func NewCreateCopytradingAlgoOrderV5ReqWithDefaults() *CreateCopytradingAlgoOrderV5Req`

NewCreateCopytradingAlgoOrderV5ReqWithDefaults instantiates a new CreateCopytradingAlgoOrderV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstType

`func (o *CreateCopytradingAlgoOrderV5Req) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *CreateCopytradingAlgoOrderV5Req) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *CreateCopytradingAlgoOrderV5Req) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *CreateCopytradingAlgoOrderV5Req) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetSlOrdPx

`func (o *CreateCopytradingAlgoOrderV5Req) GetSlOrdPx() string`

GetSlOrdPx returns the SlOrdPx field if non-nil, zero value otherwise.

### GetSlOrdPxOk

`func (o *CreateCopytradingAlgoOrderV5Req) GetSlOrdPxOk() (*string, bool)`

GetSlOrdPxOk returns a tuple with the SlOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlOrdPx

`func (o *CreateCopytradingAlgoOrderV5Req) SetSlOrdPx(v string)`

SetSlOrdPx sets SlOrdPx field to given value.

### HasSlOrdPx

`func (o *CreateCopytradingAlgoOrderV5Req) HasSlOrdPx() bool`

HasSlOrdPx returns a boolean if a field has been set.

### GetSlTriggerPx

`func (o *CreateCopytradingAlgoOrderV5Req) GetSlTriggerPx() string`

GetSlTriggerPx returns the SlTriggerPx field if non-nil, zero value otherwise.

### GetSlTriggerPxOk

`func (o *CreateCopytradingAlgoOrderV5Req) GetSlTriggerPxOk() (*string, bool)`

GetSlTriggerPxOk returns a tuple with the SlTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPx

`func (o *CreateCopytradingAlgoOrderV5Req) SetSlTriggerPx(v string)`

SetSlTriggerPx sets SlTriggerPx field to given value.

### HasSlTriggerPx

`func (o *CreateCopytradingAlgoOrderV5Req) HasSlTriggerPx() bool`

HasSlTriggerPx returns a boolean if a field has been set.

### GetSlTriggerPxType

`func (o *CreateCopytradingAlgoOrderV5Req) GetSlTriggerPxType() string`

GetSlTriggerPxType returns the SlTriggerPxType field if non-nil, zero value otherwise.

### GetSlTriggerPxTypeOk

`func (o *CreateCopytradingAlgoOrderV5Req) GetSlTriggerPxTypeOk() (*string, bool)`

GetSlTriggerPxTypeOk returns a tuple with the SlTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPxType

`func (o *CreateCopytradingAlgoOrderV5Req) SetSlTriggerPxType(v string)`

SetSlTriggerPxType sets SlTriggerPxType field to given value.

### HasSlTriggerPxType

`func (o *CreateCopytradingAlgoOrderV5Req) HasSlTriggerPxType() bool`

HasSlTriggerPxType returns a boolean if a field has been set.

### GetSubPosId

`func (o *CreateCopytradingAlgoOrderV5Req) GetSubPosId() string`

GetSubPosId returns the SubPosId field if non-nil, zero value otherwise.

### GetSubPosIdOk

`func (o *CreateCopytradingAlgoOrderV5Req) GetSubPosIdOk() (*string, bool)`

GetSubPosIdOk returns a tuple with the SubPosId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPosId

`func (o *CreateCopytradingAlgoOrderV5Req) SetSubPosId(v string)`

SetSubPosId sets SubPosId field to given value.


### GetTag

`func (o *CreateCopytradingAlgoOrderV5Req) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateCopytradingAlgoOrderV5Req) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateCopytradingAlgoOrderV5Req) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateCopytradingAlgoOrderV5Req) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTpOrdPx

`func (o *CreateCopytradingAlgoOrderV5Req) GetTpOrdPx() string`

GetTpOrdPx returns the TpOrdPx field if non-nil, zero value otherwise.

### GetTpOrdPxOk

`func (o *CreateCopytradingAlgoOrderV5Req) GetTpOrdPxOk() (*string, bool)`

GetTpOrdPxOk returns a tuple with the TpOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpOrdPx

`func (o *CreateCopytradingAlgoOrderV5Req) SetTpOrdPx(v string)`

SetTpOrdPx sets TpOrdPx field to given value.

### HasTpOrdPx

`func (o *CreateCopytradingAlgoOrderV5Req) HasTpOrdPx() bool`

HasTpOrdPx returns a boolean if a field has been set.

### GetTpTriggerPx

`func (o *CreateCopytradingAlgoOrderV5Req) GetTpTriggerPx() string`

GetTpTriggerPx returns the TpTriggerPx field if non-nil, zero value otherwise.

### GetTpTriggerPxOk

`func (o *CreateCopytradingAlgoOrderV5Req) GetTpTriggerPxOk() (*string, bool)`

GetTpTriggerPxOk returns a tuple with the TpTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPx

`func (o *CreateCopytradingAlgoOrderV5Req) SetTpTriggerPx(v string)`

SetTpTriggerPx sets TpTriggerPx field to given value.

### HasTpTriggerPx

`func (o *CreateCopytradingAlgoOrderV5Req) HasTpTriggerPx() bool`

HasTpTriggerPx returns a boolean if a field has been set.

### GetTpTriggerPxType

`func (o *CreateCopytradingAlgoOrderV5Req) GetTpTriggerPxType() string`

GetTpTriggerPxType returns the TpTriggerPxType field if non-nil, zero value otherwise.

### GetTpTriggerPxTypeOk

`func (o *CreateCopytradingAlgoOrderV5Req) GetTpTriggerPxTypeOk() (*string, bool)`

GetTpTriggerPxTypeOk returns a tuple with the TpTriggerPxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPxType

`func (o *CreateCopytradingAlgoOrderV5Req) SetTpTriggerPxType(v string)`

SetTpTriggerPxType sets TpTriggerPxType field to given value.

### HasTpTriggerPxType

`func (o *CreateCopytradingAlgoOrderV5Req) HasTpTriggerPxType() bool`

HasTpTriggerPxType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


