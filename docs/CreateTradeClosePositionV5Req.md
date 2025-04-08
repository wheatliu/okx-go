# CreateTradeClosePositionV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoCxl** | Pointer to **bool** | Whether any pending orders for closing out needs to be automatically canceled when close position via a market order.  &#x60;false&#x60; or &#x60;true&#x60;, the default is &#x60;false&#x60;. | [optional] 
**Ccy** | Pointer to **string** | Margin currency, required in the case of closing &#x60;cross&#x60; &#x60;MARGIN&#x60; position for &#x60;Spot and futures mode&#x60;. | [optional] [default to ""]
**ClOrdId** | Pointer to **string** | Client-supplied ID  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. | [optional] [default to ""]
**InstId** | **string** | Instrument ID | [default to ""]
**MgnMode** | **string** | Margin mode  &#x60;cross&#x60; &#x60;isolated&#x60; | [default to ""]
**PosSide** | Pointer to **string** | Position side   This parameter can be omitted in &#x60;net&#x60; mode, and the default value is &#x60;net&#x60;. You can only fill with &#x60;net&#x60;.  This parameter must be filled in under the &#x60;long/short&#x60; mode. Fill in &#x60;long&#x60; for close-long and &#x60;short&#x60; for close-short. | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 16 characters. | [optional] [default to ""]

## Methods

### NewCreateTradeClosePositionV5Req

`func NewCreateTradeClosePositionV5Req(instId string, mgnMode string, ) *CreateTradeClosePositionV5Req`

NewCreateTradeClosePositionV5Req instantiates a new CreateTradeClosePositionV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeClosePositionV5ReqWithDefaults

`func NewCreateTradeClosePositionV5ReqWithDefaults() *CreateTradeClosePositionV5Req`

NewCreateTradeClosePositionV5ReqWithDefaults instantiates a new CreateTradeClosePositionV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoCxl

`func (o *CreateTradeClosePositionV5Req) GetAutoCxl() bool`

GetAutoCxl returns the AutoCxl field if non-nil, zero value otherwise.

### GetAutoCxlOk

`func (o *CreateTradeClosePositionV5Req) GetAutoCxlOk() (*bool, bool)`

GetAutoCxlOk returns a tuple with the AutoCxl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCxl

`func (o *CreateTradeClosePositionV5Req) SetAutoCxl(v bool)`

SetAutoCxl sets AutoCxl field to given value.

### HasAutoCxl

`func (o *CreateTradeClosePositionV5Req) HasAutoCxl() bool`

HasAutoCxl returns a boolean if a field has been set.

### GetCcy

`func (o *CreateTradeClosePositionV5Req) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateTradeClosePositionV5Req) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateTradeClosePositionV5Req) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *CreateTradeClosePositionV5Req) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetClOrdId

`func (o *CreateTradeClosePositionV5Req) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *CreateTradeClosePositionV5Req) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *CreateTradeClosePositionV5Req) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *CreateTradeClosePositionV5Req) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetInstId

`func (o *CreateTradeClosePositionV5Req) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateTradeClosePositionV5Req) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateTradeClosePositionV5Req) SetInstId(v string)`

SetInstId sets InstId field to given value.


### GetMgnMode

`func (o *CreateTradeClosePositionV5Req) GetMgnMode() string`

GetMgnMode returns the MgnMode field if non-nil, zero value otherwise.

### GetMgnModeOk

`func (o *CreateTradeClosePositionV5Req) GetMgnModeOk() (*string, bool)`

GetMgnModeOk returns a tuple with the MgnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnMode

`func (o *CreateTradeClosePositionV5Req) SetMgnMode(v string)`

SetMgnMode sets MgnMode field to given value.


### GetPosSide

`func (o *CreateTradeClosePositionV5Req) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *CreateTradeClosePositionV5Req) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *CreateTradeClosePositionV5Req) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *CreateTradeClosePositionV5Req) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetTag

`func (o *CreateTradeClosePositionV5Req) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateTradeClosePositionV5Req) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateTradeClosePositionV5Req) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateTradeClosePositionV5Req) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


