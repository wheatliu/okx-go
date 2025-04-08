# CreateRfqExecuteQuoteV5ReqLegsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstId** | Pointer to **string** | The Instrument ID, for example: \&quot;BTC-USDT-SWAP\&quot;. | [optional] [default to ""]
**Sz** | Pointer to **string** | The size of each leg | [optional] [default to ""]

## Methods

### NewCreateRfqExecuteQuoteV5ReqLegsInner

`func NewCreateRfqExecuteQuoteV5ReqLegsInner() *CreateRfqExecuteQuoteV5ReqLegsInner`

NewCreateRfqExecuteQuoteV5ReqLegsInner instantiates a new CreateRfqExecuteQuoteV5ReqLegsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqExecuteQuoteV5ReqLegsInnerWithDefaults

`func NewCreateRfqExecuteQuoteV5ReqLegsInnerWithDefaults() *CreateRfqExecuteQuoteV5ReqLegsInner`

NewCreateRfqExecuteQuoteV5ReqLegsInnerWithDefaults instantiates a new CreateRfqExecuteQuoteV5ReqLegsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstId

`func (o *CreateRfqExecuteQuoteV5ReqLegsInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateRfqExecuteQuoteV5ReqLegsInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateRfqExecuteQuoteV5ReqLegsInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *CreateRfqExecuteQuoteV5ReqLegsInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetSz

`func (o *CreateRfqExecuteQuoteV5ReqLegsInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *CreateRfqExecuteQuoteV5ReqLegsInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *CreateRfqExecuteQuoteV5ReqLegsInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *CreateRfqExecuteQuoteV5ReqLegsInner) HasSz() bool`

HasSz returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


