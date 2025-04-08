# GetAssetConvertHistoryV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCcy** | Pointer to **string** | Base currency, e.g. &#x60;BTC&#x60; in &#x60;BTC-USDT&#x60; | [optional] [default to ""]
**ClTReqId** | Pointer to **string** | Client Order ID as assigned by the client | [optional] [default to ""]
**FillBaseSz** | Pointer to **string** | Filled amount for base currency | [optional] [default to ""]
**FillPx** | Pointer to **string** | Filled price based on quote currency | [optional] [default to ""]
**FillQuoteSz** | Pointer to **string** | Filled amount for quote currency | [optional] [default to ""]
**InstId** | Pointer to **string** | Currency pair, e.g. &#x60;BTC-USDT&#x60; | [optional] [default to ""]
**QuoteCcy** | Pointer to **string** | Quote currency, e.g. &#x60;USDT&#x60; in &#x60;BTC-USDT&#x60; | [optional] [default to ""]
**Side** | Pointer to **string** | Trade side based on &#x60;baseCcy&#x60;  &#x60;buy&#x60; &#x60;sell&#x60; | [optional] [default to ""]
**State** | Pointer to **string** | Trade state  &#x60;fullyFilled&#x60; : success   &#x60;rejected&#x60; : failed | [optional] [default to ""]
**TradeId** | Pointer to **string** | Trade ID | [optional] [default to ""]
**Ts** | Pointer to **string** | Convert trade time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetAssetConvertHistoryV5RespDataInner

`func NewGetAssetConvertHistoryV5RespDataInner() *GetAssetConvertHistoryV5RespDataInner`

NewGetAssetConvertHistoryV5RespDataInner instantiates a new GetAssetConvertHistoryV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetConvertHistoryV5RespDataInnerWithDefaults

`func NewGetAssetConvertHistoryV5RespDataInnerWithDefaults() *GetAssetConvertHistoryV5RespDataInner`

NewGetAssetConvertHistoryV5RespDataInnerWithDefaults instantiates a new GetAssetConvertHistoryV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseCcy

`func (o *GetAssetConvertHistoryV5RespDataInner) GetBaseCcy() string`

GetBaseCcy returns the BaseCcy field if non-nil, zero value otherwise.

### GetBaseCcyOk

`func (o *GetAssetConvertHistoryV5RespDataInner) GetBaseCcyOk() (*string, bool)`

GetBaseCcyOk returns a tuple with the BaseCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCcy

`func (o *GetAssetConvertHistoryV5RespDataInner) SetBaseCcy(v string)`

SetBaseCcy sets BaseCcy field to given value.

### HasBaseCcy

`func (o *GetAssetConvertHistoryV5RespDataInner) HasBaseCcy() bool`

HasBaseCcy returns a boolean if a field has been set.

### GetClTReqId

`func (o *GetAssetConvertHistoryV5RespDataInner) GetClTReqId() string`

GetClTReqId returns the ClTReqId field if non-nil, zero value otherwise.

### GetClTReqIdOk

`func (o *GetAssetConvertHistoryV5RespDataInner) GetClTReqIdOk() (*string, bool)`

GetClTReqIdOk returns a tuple with the ClTReqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClTReqId

`func (o *GetAssetConvertHistoryV5RespDataInner) SetClTReqId(v string)`

SetClTReqId sets ClTReqId field to given value.

### HasClTReqId

`func (o *GetAssetConvertHistoryV5RespDataInner) HasClTReqId() bool`

HasClTReqId returns a boolean if a field has been set.

### GetFillBaseSz

`func (o *GetAssetConvertHistoryV5RespDataInner) GetFillBaseSz() string`

GetFillBaseSz returns the FillBaseSz field if non-nil, zero value otherwise.

### GetFillBaseSzOk

`func (o *GetAssetConvertHistoryV5RespDataInner) GetFillBaseSzOk() (*string, bool)`

GetFillBaseSzOk returns a tuple with the FillBaseSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillBaseSz

`func (o *GetAssetConvertHistoryV5RespDataInner) SetFillBaseSz(v string)`

SetFillBaseSz sets FillBaseSz field to given value.

### HasFillBaseSz

`func (o *GetAssetConvertHistoryV5RespDataInner) HasFillBaseSz() bool`

HasFillBaseSz returns a boolean if a field has been set.

### GetFillPx

`func (o *GetAssetConvertHistoryV5RespDataInner) GetFillPx() string`

GetFillPx returns the FillPx field if non-nil, zero value otherwise.

### GetFillPxOk

`func (o *GetAssetConvertHistoryV5RespDataInner) GetFillPxOk() (*string, bool)`

GetFillPxOk returns a tuple with the FillPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPx

`func (o *GetAssetConvertHistoryV5RespDataInner) SetFillPx(v string)`

SetFillPx sets FillPx field to given value.

### HasFillPx

`func (o *GetAssetConvertHistoryV5RespDataInner) HasFillPx() bool`

HasFillPx returns a boolean if a field has been set.

### GetFillQuoteSz

`func (o *GetAssetConvertHistoryV5RespDataInner) GetFillQuoteSz() string`

GetFillQuoteSz returns the FillQuoteSz field if non-nil, zero value otherwise.

### GetFillQuoteSzOk

`func (o *GetAssetConvertHistoryV5RespDataInner) GetFillQuoteSzOk() (*string, bool)`

GetFillQuoteSzOk returns a tuple with the FillQuoteSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillQuoteSz

`func (o *GetAssetConvertHistoryV5RespDataInner) SetFillQuoteSz(v string)`

SetFillQuoteSz sets FillQuoteSz field to given value.

### HasFillQuoteSz

`func (o *GetAssetConvertHistoryV5RespDataInner) HasFillQuoteSz() bool`

HasFillQuoteSz returns a boolean if a field has been set.

### GetInstId

`func (o *GetAssetConvertHistoryV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetAssetConvertHistoryV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetAssetConvertHistoryV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetAssetConvertHistoryV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetQuoteCcy

`func (o *GetAssetConvertHistoryV5RespDataInner) GetQuoteCcy() string`

GetQuoteCcy returns the QuoteCcy field if non-nil, zero value otherwise.

### GetQuoteCcyOk

`func (o *GetAssetConvertHistoryV5RespDataInner) GetQuoteCcyOk() (*string, bool)`

GetQuoteCcyOk returns a tuple with the QuoteCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCcy

`func (o *GetAssetConvertHistoryV5RespDataInner) SetQuoteCcy(v string)`

SetQuoteCcy sets QuoteCcy field to given value.

### HasQuoteCcy

`func (o *GetAssetConvertHistoryV5RespDataInner) HasQuoteCcy() bool`

HasQuoteCcy returns a boolean if a field has been set.

### GetSide

`func (o *GetAssetConvertHistoryV5RespDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetAssetConvertHistoryV5RespDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetAssetConvertHistoryV5RespDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetAssetConvertHistoryV5RespDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetState

`func (o *GetAssetConvertHistoryV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetAssetConvertHistoryV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetAssetConvertHistoryV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetAssetConvertHistoryV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTradeId

`func (o *GetAssetConvertHistoryV5RespDataInner) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *GetAssetConvertHistoryV5RespDataInner) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *GetAssetConvertHistoryV5RespDataInner) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *GetAssetConvertHistoryV5RespDataInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.

### GetTs

`func (o *GetAssetConvertHistoryV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetAssetConvertHistoryV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetAssetConvertHistoryV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetAssetConvertHistoryV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


