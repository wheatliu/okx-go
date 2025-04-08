# CreateCopytradingCloseSubpositionV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstType** | Pointer to **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;SWAP&#x60;, the default value | [optional] [default to ""]
**OrdType** | Pointer to **string** | Order type  &#x60;market&#x60;：Market order, the default value  &#x60;limit&#x60;：Limit order   | [optional] [default to ""]
**Px** | Pointer to **string** | Order price. Only applicable to &#x60;limit&#x60; order and &#x60;SPOT&#x60; lead trader   If the price is 0, the pending order will be canceled.   It is modifying order if you set &#x60;px&#x60; after placing limit order. | [optional] [default to ""]
**SubPosId** | **string** | Lead position ID | [default to ""]
**Tag** | Pointer to **string** | Order tag  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 16 characters. | [optional] [default to ""]

## Methods

### NewCreateCopytradingCloseSubpositionV5Req

`func NewCreateCopytradingCloseSubpositionV5Req(subPosId string, ) *CreateCopytradingCloseSubpositionV5Req`

NewCreateCopytradingCloseSubpositionV5Req instantiates a new CreateCopytradingCloseSubpositionV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCopytradingCloseSubpositionV5ReqWithDefaults

`func NewCreateCopytradingCloseSubpositionV5ReqWithDefaults() *CreateCopytradingCloseSubpositionV5Req`

NewCreateCopytradingCloseSubpositionV5ReqWithDefaults instantiates a new CreateCopytradingCloseSubpositionV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstType

`func (o *CreateCopytradingCloseSubpositionV5Req) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *CreateCopytradingCloseSubpositionV5Req) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *CreateCopytradingCloseSubpositionV5Req) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *CreateCopytradingCloseSubpositionV5Req) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetOrdType

`func (o *CreateCopytradingCloseSubpositionV5Req) GetOrdType() string`

GetOrdType returns the OrdType field if non-nil, zero value otherwise.

### GetOrdTypeOk

`func (o *CreateCopytradingCloseSubpositionV5Req) GetOrdTypeOk() (*string, bool)`

GetOrdTypeOk returns a tuple with the OrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdType

`func (o *CreateCopytradingCloseSubpositionV5Req) SetOrdType(v string)`

SetOrdType sets OrdType field to given value.

### HasOrdType

`func (o *CreateCopytradingCloseSubpositionV5Req) HasOrdType() bool`

HasOrdType returns a boolean if a field has been set.

### GetPx

`func (o *CreateCopytradingCloseSubpositionV5Req) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *CreateCopytradingCloseSubpositionV5Req) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *CreateCopytradingCloseSubpositionV5Req) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *CreateCopytradingCloseSubpositionV5Req) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSubPosId

`func (o *CreateCopytradingCloseSubpositionV5Req) GetSubPosId() string`

GetSubPosId returns the SubPosId field if non-nil, zero value otherwise.

### GetSubPosIdOk

`func (o *CreateCopytradingCloseSubpositionV5Req) GetSubPosIdOk() (*string, bool)`

GetSubPosIdOk returns a tuple with the SubPosId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPosId

`func (o *CreateCopytradingCloseSubpositionV5Req) SetSubPosId(v string)`

SetSubPosId sets SubPosId field to given value.


### GetTag

`func (o *CreateCopytradingCloseSubpositionV5Req) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateCopytradingCloseSubpositionV5Req) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateCopytradingCloseSubpositionV5Req) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateCopytradingCloseSubpositionV5Req) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


