# GetAccountMaxSizeV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | Pointer to **string** | Currency used for margin | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**MaxBuy** | Pointer to **string** | &#x60;SPOT&#x60;/&#x60;MARGIN&#x60;: The maximum quantity in base currency that you can buy  The cross-margin order under &#x60;Spot and futures mode&#x60; mode, quantity of coins is based on base currency.  &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTIONS&#x60;: The maximum quantity of contracts that you can buy | [optional] [default to ""]
**MaxSell** | Pointer to **string** | &#x60;SPOT&#x60;/&#x60;MARGIN&#x60;: The maximum quantity in quote currency that you can sell  The cross-margin order under &#x60;Spot and futures mode&#x60; mode, quantity of coins is based on base currency.  &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTIONS&#x60;: The maximum quantity of contracts that you can sell | [optional] [default to ""]

## Methods

### NewGetAccountMaxSizeV5RespDataInner

`func NewGetAccountMaxSizeV5RespDataInner() *GetAccountMaxSizeV5RespDataInner`

NewGetAccountMaxSizeV5RespDataInner instantiates a new GetAccountMaxSizeV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountMaxSizeV5RespDataInnerWithDefaults

`func NewGetAccountMaxSizeV5RespDataInnerWithDefaults() *GetAccountMaxSizeV5RespDataInner`

NewGetAccountMaxSizeV5RespDataInnerWithDefaults instantiates a new GetAccountMaxSizeV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *GetAccountMaxSizeV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAccountMaxSizeV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAccountMaxSizeV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAccountMaxSizeV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetInstId

`func (o *GetAccountMaxSizeV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetAccountMaxSizeV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetAccountMaxSizeV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetAccountMaxSizeV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetMaxBuy

`func (o *GetAccountMaxSizeV5RespDataInner) GetMaxBuy() string`

GetMaxBuy returns the MaxBuy field if non-nil, zero value otherwise.

### GetMaxBuyOk

`func (o *GetAccountMaxSizeV5RespDataInner) GetMaxBuyOk() (*string, bool)`

GetMaxBuyOk returns a tuple with the MaxBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBuy

`func (o *GetAccountMaxSizeV5RespDataInner) SetMaxBuy(v string)`

SetMaxBuy sets MaxBuy field to given value.

### HasMaxBuy

`func (o *GetAccountMaxSizeV5RespDataInner) HasMaxBuy() bool`

HasMaxBuy returns a boolean if a field has been set.

### GetMaxSell

`func (o *GetAccountMaxSizeV5RespDataInner) GetMaxSell() string`

GetMaxSell returns the MaxSell field if non-nil, zero value otherwise.

### GetMaxSellOk

`func (o *GetAccountMaxSizeV5RespDataInner) GetMaxSellOk() (*string, bool)`

GetMaxSellOk returns a tuple with the MaxSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSell

`func (o *GetAccountMaxSizeV5RespDataInner) SetMaxSell(v string)`

SetMaxSell sets MaxSell field to given value.

### HasMaxSell

`func (o *GetAccountMaxSizeV5RespDataInner) HasMaxSell() bool`

HasMaxSell returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


