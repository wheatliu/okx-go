# GetPublicOptionTradesV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FillVol** | Pointer to **string** | Implied volatility while trading (Correspond to trade price) | [optional] [default to ""]
**FwdPx** | Pointer to **string** | Forward price while trading | [optional] [default to ""]
**IdxPx** | Pointer to **string** | Index price while trading | [optional] [default to ""]
**InstFamily** | Pointer to **string** | Instrument family | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**MarkPx** | Pointer to **string** | Mark price while trading | [optional] [default to ""]
**OptType** | Pointer to **string** | Option type, C: Call P: Put | [optional] [default to ""]
**Px** | Pointer to **string** | Trade price | [optional] [default to ""]
**Side** | Pointer to **string** | Trade side   &#x60;buy&#x60;    &#x60;sell&#x60; | [optional] [default to ""]
**TradeId** | Pointer to **string** | Trade ID | [optional] [default to ""]
**Ts** | Pointer to **string** | Trade time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60;. | [optional] [default to ""]

## Methods

### NewGetPublicOptionTradesV5RespData

`func NewGetPublicOptionTradesV5RespData() *GetPublicOptionTradesV5RespData`

NewGetPublicOptionTradesV5RespData instantiates a new GetPublicOptionTradesV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicOptionTradesV5RespDataWithDefaults

`func NewGetPublicOptionTradesV5RespDataWithDefaults() *GetPublicOptionTradesV5RespData`

NewGetPublicOptionTradesV5RespDataWithDefaults instantiates a new GetPublicOptionTradesV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFillVol

`func (o *GetPublicOptionTradesV5RespData) GetFillVol() string`

GetFillVol returns the FillVol field if non-nil, zero value otherwise.

### GetFillVolOk

`func (o *GetPublicOptionTradesV5RespData) GetFillVolOk() (*string, bool)`

GetFillVolOk returns a tuple with the FillVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillVol

`func (o *GetPublicOptionTradesV5RespData) SetFillVol(v string)`

SetFillVol sets FillVol field to given value.

### HasFillVol

`func (o *GetPublicOptionTradesV5RespData) HasFillVol() bool`

HasFillVol returns a boolean if a field has been set.

### GetFwdPx

`func (o *GetPublicOptionTradesV5RespData) GetFwdPx() string`

GetFwdPx returns the FwdPx field if non-nil, zero value otherwise.

### GetFwdPxOk

`func (o *GetPublicOptionTradesV5RespData) GetFwdPxOk() (*string, bool)`

GetFwdPxOk returns a tuple with the FwdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwdPx

`func (o *GetPublicOptionTradesV5RespData) SetFwdPx(v string)`

SetFwdPx sets FwdPx field to given value.

### HasFwdPx

`func (o *GetPublicOptionTradesV5RespData) HasFwdPx() bool`

HasFwdPx returns a boolean if a field has been set.

### GetIdxPx

`func (o *GetPublicOptionTradesV5RespData) GetIdxPx() string`

GetIdxPx returns the IdxPx field if non-nil, zero value otherwise.

### GetIdxPxOk

`func (o *GetPublicOptionTradesV5RespData) GetIdxPxOk() (*string, bool)`

GetIdxPxOk returns a tuple with the IdxPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdxPx

`func (o *GetPublicOptionTradesV5RespData) SetIdxPx(v string)`

SetIdxPx sets IdxPx field to given value.

### HasIdxPx

`func (o *GetPublicOptionTradesV5RespData) HasIdxPx() bool`

HasIdxPx returns a boolean if a field has been set.

### GetInstFamily

`func (o *GetPublicOptionTradesV5RespData) GetInstFamily() string`

GetInstFamily returns the InstFamily field if non-nil, zero value otherwise.

### GetInstFamilyOk

`func (o *GetPublicOptionTradesV5RespData) GetInstFamilyOk() (*string, bool)`

GetInstFamilyOk returns a tuple with the InstFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstFamily

`func (o *GetPublicOptionTradesV5RespData) SetInstFamily(v string)`

SetInstFamily sets InstFamily field to given value.

### HasInstFamily

`func (o *GetPublicOptionTradesV5RespData) HasInstFamily() bool`

HasInstFamily returns a boolean if a field has been set.

### GetInstId

`func (o *GetPublicOptionTradesV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetPublicOptionTradesV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetPublicOptionTradesV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetPublicOptionTradesV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetMarkPx

`func (o *GetPublicOptionTradesV5RespData) GetMarkPx() string`

GetMarkPx returns the MarkPx field if non-nil, zero value otherwise.

### GetMarkPxOk

`func (o *GetPublicOptionTradesV5RespData) GetMarkPxOk() (*string, bool)`

GetMarkPxOk returns a tuple with the MarkPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPx

`func (o *GetPublicOptionTradesV5RespData) SetMarkPx(v string)`

SetMarkPx sets MarkPx field to given value.

### HasMarkPx

`func (o *GetPublicOptionTradesV5RespData) HasMarkPx() bool`

HasMarkPx returns a boolean if a field has been set.

### GetOptType

`func (o *GetPublicOptionTradesV5RespData) GetOptType() string`

GetOptType returns the OptType field if non-nil, zero value otherwise.

### GetOptTypeOk

`func (o *GetPublicOptionTradesV5RespData) GetOptTypeOk() (*string, bool)`

GetOptTypeOk returns a tuple with the OptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptType

`func (o *GetPublicOptionTradesV5RespData) SetOptType(v string)`

SetOptType sets OptType field to given value.

### HasOptType

`func (o *GetPublicOptionTradesV5RespData) HasOptType() bool`

HasOptType returns a boolean if a field has been set.

### GetPx

`func (o *GetPublicOptionTradesV5RespData) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetPublicOptionTradesV5RespData) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetPublicOptionTradesV5RespData) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetPublicOptionTradesV5RespData) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSide

`func (o *GetPublicOptionTradesV5RespData) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetPublicOptionTradesV5RespData) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetPublicOptionTradesV5RespData) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetPublicOptionTradesV5RespData) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetTradeId

`func (o *GetPublicOptionTradesV5RespData) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *GetPublicOptionTradesV5RespData) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *GetPublicOptionTradesV5RespData) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *GetPublicOptionTradesV5RespData) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.

### GetTs

`func (o *GetPublicOptionTradesV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetPublicOptionTradesV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetPublicOptionTradesV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetPublicOptionTradesV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


