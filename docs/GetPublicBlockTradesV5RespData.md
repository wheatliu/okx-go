# GetPublicBlockTradesV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FillVol** | Pointer to **string** | Implied volatility   Only applicable to &#x60;OPTION&#x60; | [optional] [default to ""]
**FwdPx** | Pointer to **string** | Forward price   Only applicable to &#x60;OPTION&#x60; | [optional] [default to ""]
**IdxPx** | Pointer to **string** | Index price    Applicable to &#x60;FUTURES&#x60;, &#x60;SWAP&#x60;, &#x60;OPTION&#x60; | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**MarkPx** | Pointer to **string** | Mark price    Applicable to &#x60;FUTURES&#x60;, &#x60;SWAP&#x60;, &#x60;OPTION&#x60; | [optional] [default to ""]
**Px** | Pointer to **string** | Trade price | [optional] [default to ""]
**Side** | Pointer to **string** | Trade side   &#x60;buy&#x60;    &#x60;sell&#x60; | [optional] [default to ""]
**Sz** | Pointer to **string** | Trade quantity   For spot trading, the unit is base currency  For &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;, the unit is contract. | [optional] [default to ""]
**TradeId** | Pointer to **string** | Trade ID | [optional] [default to ""]
**Ts** | Pointer to **string** | Trade time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60;. | [optional] [default to ""]

## Methods

### NewGetPublicBlockTradesV5RespData

`func NewGetPublicBlockTradesV5RespData() *GetPublicBlockTradesV5RespData`

NewGetPublicBlockTradesV5RespData instantiates a new GetPublicBlockTradesV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicBlockTradesV5RespDataWithDefaults

`func NewGetPublicBlockTradesV5RespDataWithDefaults() *GetPublicBlockTradesV5RespData`

NewGetPublicBlockTradesV5RespDataWithDefaults instantiates a new GetPublicBlockTradesV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFillVol

`func (o *GetPublicBlockTradesV5RespData) GetFillVol() string`

GetFillVol returns the FillVol field if non-nil, zero value otherwise.

### GetFillVolOk

`func (o *GetPublicBlockTradesV5RespData) GetFillVolOk() (*string, bool)`

GetFillVolOk returns a tuple with the FillVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillVol

`func (o *GetPublicBlockTradesV5RespData) SetFillVol(v string)`

SetFillVol sets FillVol field to given value.

### HasFillVol

`func (o *GetPublicBlockTradesV5RespData) HasFillVol() bool`

HasFillVol returns a boolean if a field has been set.

### GetFwdPx

`func (o *GetPublicBlockTradesV5RespData) GetFwdPx() string`

GetFwdPx returns the FwdPx field if non-nil, zero value otherwise.

### GetFwdPxOk

`func (o *GetPublicBlockTradesV5RespData) GetFwdPxOk() (*string, bool)`

GetFwdPxOk returns a tuple with the FwdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwdPx

`func (o *GetPublicBlockTradesV5RespData) SetFwdPx(v string)`

SetFwdPx sets FwdPx field to given value.

### HasFwdPx

`func (o *GetPublicBlockTradesV5RespData) HasFwdPx() bool`

HasFwdPx returns a boolean if a field has been set.

### GetIdxPx

`func (o *GetPublicBlockTradesV5RespData) GetIdxPx() string`

GetIdxPx returns the IdxPx field if non-nil, zero value otherwise.

### GetIdxPxOk

`func (o *GetPublicBlockTradesV5RespData) GetIdxPxOk() (*string, bool)`

GetIdxPxOk returns a tuple with the IdxPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdxPx

`func (o *GetPublicBlockTradesV5RespData) SetIdxPx(v string)`

SetIdxPx sets IdxPx field to given value.

### HasIdxPx

`func (o *GetPublicBlockTradesV5RespData) HasIdxPx() bool`

HasIdxPx returns a boolean if a field has been set.

### GetInstId

`func (o *GetPublicBlockTradesV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetPublicBlockTradesV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetPublicBlockTradesV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetPublicBlockTradesV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetMarkPx

`func (o *GetPublicBlockTradesV5RespData) GetMarkPx() string`

GetMarkPx returns the MarkPx field if non-nil, zero value otherwise.

### GetMarkPxOk

`func (o *GetPublicBlockTradesV5RespData) GetMarkPxOk() (*string, bool)`

GetMarkPxOk returns a tuple with the MarkPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPx

`func (o *GetPublicBlockTradesV5RespData) SetMarkPx(v string)`

SetMarkPx sets MarkPx field to given value.

### HasMarkPx

`func (o *GetPublicBlockTradesV5RespData) HasMarkPx() bool`

HasMarkPx returns a boolean if a field has been set.

### GetPx

`func (o *GetPublicBlockTradesV5RespData) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetPublicBlockTradesV5RespData) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetPublicBlockTradesV5RespData) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetPublicBlockTradesV5RespData) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSide

`func (o *GetPublicBlockTradesV5RespData) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetPublicBlockTradesV5RespData) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetPublicBlockTradesV5RespData) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetPublicBlockTradesV5RespData) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSz

`func (o *GetPublicBlockTradesV5RespData) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetPublicBlockTradesV5RespData) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetPublicBlockTradesV5RespData) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetPublicBlockTradesV5RespData) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTradeId

`func (o *GetPublicBlockTradesV5RespData) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *GetPublicBlockTradesV5RespData) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *GetPublicBlockTradesV5RespData) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *GetPublicBlockTradesV5RespData) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.

### GetTs

`func (o *GetPublicBlockTradesV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetPublicBlockTradesV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetPublicBlockTradesV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetPublicBlockTradesV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


