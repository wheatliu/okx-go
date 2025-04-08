# GetSprdSpreadsV5RespDataInnerLegsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstId** | Pointer to **string** | Instrument ID, e.g. BTC-USD-SWAP | [optional] [default to ""]
**Side** | Pointer to **string** | The direction of the leg of the spread. Valid Values include &#x60;buy&#x60; and &#x60;sell&#x60;. | [optional] [default to ""]

## Methods

### NewGetSprdSpreadsV5RespDataInnerLegsInner

`func NewGetSprdSpreadsV5RespDataInnerLegsInner() *GetSprdSpreadsV5RespDataInnerLegsInner`

NewGetSprdSpreadsV5RespDataInnerLegsInner instantiates a new GetSprdSpreadsV5RespDataInnerLegsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSprdSpreadsV5RespDataInnerLegsInnerWithDefaults

`func NewGetSprdSpreadsV5RespDataInnerLegsInnerWithDefaults() *GetSprdSpreadsV5RespDataInnerLegsInner`

NewGetSprdSpreadsV5RespDataInnerLegsInnerWithDefaults instantiates a new GetSprdSpreadsV5RespDataInnerLegsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstId

`func (o *GetSprdSpreadsV5RespDataInnerLegsInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetSprdSpreadsV5RespDataInnerLegsInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetSprdSpreadsV5RespDataInnerLegsInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetSprdSpreadsV5RespDataInnerLegsInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetSide

`func (o *GetSprdSpreadsV5RespDataInnerLegsInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetSprdSpreadsV5RespDataInnerLegsInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetSprdSpreadsV5RespDataInnerLegsInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetSprdSpreadsV5RespDataInnerLegsInner) HasSide() bool`

HasSide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


