# GetMarketBooksFullV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asks** | Pointer to **string** | Order book on sell side | [optional] [default to ""]
**Bids** | Pointer to **string** | Order book on buy side | [optional] [default to ""]
**Ts** | Pointer to **string** | Order book generation time | [optional] [default to ""]

## Methods

### NewGetMarketBooksFullV5RespData

`func NewGetMarketBooksFullV5RespData() *GetMarketBooksFullV5RespData`

NewGetMarketBooksFullV5RespData instantiates a new GetMarketBooksFullV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketBooksFullV5RespDataWithDefaults

`func NewGetMarketBooksFullV5RespDataWithDefaults() *GetMarketBooksFullV5RespData`

NewGetMarketBooksFullV5RespDataWithDefaults instantiates a new GetMarketBooksFullV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsks

`func (o *GetMarketBooksFullV5RespData) GetAsks() string`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *GetMarketBooksFullV5RespData) GetAsksOk() (*string, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *GetMarketBooksFullV5RespData) SetAsks(v string)`

SetAsks sets Asks field to given value.

### HasAsks

`func (o *GetMarketBooksFullV5RespData) HasAsks() bool`

HasAsks returns a boolean if a field has been set.

### GetBids

`func (o *GetMarketBooksFullV5RespData) GetBids() string`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *GetMarketBooksFullV5RespData) GetBidsOk() (*string, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *GetMarketBooksFullV5RespData) SetBids(v string)`

SetBids sets Bids field to given value.

### HasBids

`func (o *GetMarketBooksFullV5RespData) HasBids() bool`

HasBids returns a boolean if a field has been set.

### GetTs

`func (o *GetMarketBooksFullV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetMarketBooksFullV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetMarketBooksFullV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetMarketBooksFullV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


