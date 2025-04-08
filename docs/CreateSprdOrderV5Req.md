# CreateSprdOrderV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client   A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. | [optional] [default to ""]
**OrdType** | **string** | Order type  &#x60;market&#x60;: Market order   &#x60;limit&#x60;: Limit order   &#x60;post_only&#x60;: Post-only order  &#x60;ioc&#x60;: Immediate-or-cancel order | [default to ""]
**Px** | **string** | Order price. Only applicable to &#x60;limit&#x60;, &#x60;post_only&#x60;, &#x60;ioc&#x60; | [default to ""]
**Side** | **string** | Order side, &#x60;buy&#x60; &#x60;sell&#x60; | [default to ""]
**SprdId** | **string** | spread ID, e.g. BTC-USDT_BTC-USD-SWAP | [default to ""]
**Sz** | **string** | Quantity to buy or sell. The unit is USD for inverse spreads, and the corresponding baseCcy for linear and hybrid spreads. | [default to ""]
**Tag** | Pointer to **string** | Order tag   A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 16 characters. | [optional] [default to ""]

## Methods

### NewCreateSprdOrderV5Req

`func NewCreateSprdOrderV5Req(ordType string, px string, side string, sprdId string, sz string, ) *CreateSprdOrderV5Req`

NewCreateSprdOrderV5Req instantiates a new CreateSprdOrderV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSprdOrderV5ReqWithDefaults

`func NewCreateSprdOrderV5ReqWithDefaults() *CreateSprdOrderV5Req`

NewCreateSprdOrderV5ReqWithDefaults instantiates a new CreateSprdOrderV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClOrdId

`func (o *CreateSprdOrderV5Req) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *CreateSprdOrderV5Req) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *CreateSprdOrderV5Req) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *CreateSprdOrderV5Req) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetOrdType

`func (o *CreateSprdOrderV5Req) GetOrdType() string`

GetOrdType returns the OrdType field if non-nil, zero value otherwise.

### GetOrdTypeOk

`func (o *CreateSprdOrderV5Req) GetOrdTypeOk() (*string, bool)`

GetOrdTypeOk returns a tuple with the OrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdType

`func (o *CreateSprdOrderV5Req) SetOrdType(v string)`

SetOrdType sets OrdType field to given value.


### GetPx

`func (o *CreateSprdOrderV5Req) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *CreateSprdOrderV5Req) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *CreateSprdOrderV5Req) SetPx(v string)`

SetPx sets Px field to given value.


### GetSide

`func (o *CreateSprdOrderV5Req) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateSprdOrderV5Req) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateSprdOrderV5Req) SetSide(v string)`

SetSide sets Side field to given value.


### GetSprdId

`func (o *CreateSprdOrderV5Req) GetSprdId() string`

GetSprdId returns the SprdId field if non-nil, zero value otherwise.

### GetSprdIdOk

`func (o *CreateSprdOrderV5Req) GetSprdIdOk() (*string, bool)`

GetSprdIdOk returns a tuple with the SprdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSprdId

`func (o *CreateSprdOrderV5Req) SetSprdId(v string)`

SetSprdId sets SprdId field to given value.


### GetSz

`func (o *CreateSprdOrderV5Req) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *CreateSprdOrderV5Req) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *CreateSprdOrderV5Req) SetSz(v string)`

SetSz sets Sz field to given value.


### GetTag

`func (o *CreateSprdOrderV5Req) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateSprdOrderV5Req) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateSprdOrderV5Req) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateSprdOrderV5Req) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


