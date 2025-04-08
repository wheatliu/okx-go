# GetAccountMaxAvailSizeV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailBuy** | Pointer to **string** | Maximum available balance/equity to buy | [optional] [default to ""]
**AvailSell** | Pointer to **string** | Maximum available balance/equity to sell | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]

## Methods

### NewGetAccountMaxAvailSizeV5RespData

`func NewGetAccountMaxAvailSizeV5RespData() *GetAccountMaxAvailSizeV5RespData`

NewGetAccountMaxAvailSizeV5RespData instantiates a new GetAccountMaxAvailSizeV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountMaxAvailSizeV5RespDataWithDefaults

`func NewGetAccountMaxAvailSizeV5RespDataWithDefaults() *GetAccountMaxAvailSizeV5RespData`

NewGetAccountMaxAvailSizeV5RespDataWithDefaults instantiates a new GetAccountMaxAvailSizeV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailBuy

`func (o *GetAccountMaxAvailSizeV5RespData) GetAvailBuy() string`

GetAvailBuy returns the AvailBuy field if non-nil, zero value otherwise.

### GetAvailBuyOk

`func (o *GetAccountMaxAvailSizeV5RespData) GetAvailBuyOk() (*string, bool)`

GetAvailBuyOk returns a tuple with the AvailBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailBuy

`func (o *GetAccountMaxAvailSizeV5RespData) SetAvailBuy(v string)`

SetAvailBuy sets AvailBuy field to given value.

### HasAvailBuy

`func (o *GetAccountMaxAvailSizeV5RespData) HasAvailBuy() bool`

HasAvailBuy returns a boolean if a field has been set.

### GetAvailSell

`func (o *GetAccountMaxAvailSizeV5RespData) GetAvailSell() string`

GetAvailSell returns the AvailSell field if non-nil, zero value otherwise.

### GetAvailSellOk

`func (o *GetAccountMaxAvailSizeV5RespData) GetAvailSellOk() (*string, bool)`

GetAvailSellOk returns a tuple with the AvailSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailSell

`func (o *GetAccountMaxAvailSizeV5RespData) SetAvailSell(v string)`

SetAvailSell sets AvailSell field to given value.

### HasAvailSell

`func (o *GetAccountMaxAvailSizeV5RespData) HasAvailSell() bool`

HasAvailSell returns a boolean if a field has been set.

### GetInstId

`func (o *GetAccountMaxAvailSizeV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetAccountMaxAvailSizeV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetAccountMaxAvailSizeV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetAccountMaxAvailSizeV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


