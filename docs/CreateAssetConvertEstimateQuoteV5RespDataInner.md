# CreateAssetConvertEstimateQuoteV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCcy** | Pointer to **string** | Base currency, e.g. &#x60;BTC&#x60; in &#x60;BTC-USDT&#x60; | [optional] [default to ""]
**BaseSz** | Pointer to **string** | Convert amount of base currency | [optional] [default to ""]
**ClQReqId** | Pointer to **string** | Client Order ID as assigned by the client | [optional] [default to ""]
**CnvtPx** | Pointer to **string** | Convert price based on quote currency | [optional] [default to ""]
**OrigRfqSz** | Pointer to **string** | Original RFQ amount | [optional] [default to ""]
**QuoteCcy** | Pointer to **string** | Quote currency, e.g. &#x60;USDT&#x60; in &#x60;BTC-USDT&#x60; | [optional] [default to ""]
**QuoteId** | Pointer to **string** | Quote ID | [optional] [default to ""]
**QuoteSz** | Pointer to **string** | Convert amount of quote currency | [optional] [default to ""]
**QuoteTime** | Pointer to **string** | Quotation generation time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**RfqSz** | Pointer to **string** | Real RFQ amount | [optional] [default to ""]
**RfqSzCcy** | Pointer to **string** | RFQ currency | [optional] [default to ""]
**Side** | Pointer to **string** | Trade side based on &#x60;baseCcy&#x60; | [optional] [default to ""]
**TtlMs** | Pointer to **string** | Validity period of quotation in milliseconds | [optional] [default to ""]

## Methods

### NewCreateAssetConvertEstimateQuoteV5RespDataInner

`func NewCreateAssetConvertEstimateQuoteV5RespDataInner() *CreateAssetConvertEstimateQuoteV5RespDataInner`

NewCreateAssetConvertEstimateQuoteV5RespDataInner instantiates a new CreateAssetConvertEstimateQuoteV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAssetConvertEstimateQuoteV5RespDataInnerWithDefaults

`func NewCreateAssetConvertEstimateQuoteV5RespDataInnerWithDefaults() *CreateAssetConvertEstimateQuoteV5RespDataInner`

NewCreateAssetConvertEstimateQuoteV5RespDataInnerWithDefaults instantiates a new CreateAssetConvertEstimateQuoteV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseCcy

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetBaseCcy() string`

GetBaseCcy returns the BaseCcy field if non-nil, zero value otherwise.

### GetBaseCcyOk

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetBaseCcyOk() (*string, bool)`

GetBaseCcyOk returns a tuple with the BaseCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCcy

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) SetBaseCcy(v string)`

SetBaseCcy sets BaseCcy field to given value.

### HasBaseCcy

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) HasBaseCcy() bool`

HasBaseCcy returns a boolean if a field has been set.

### GetBaseSz

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetBaseSz() string`

GetBaseSz returns the BaseSz field if non-nil, zero value otherwise.

### GetBaseSzOk

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetBaseSzOk() (*string, bool)`

GetBaseSzOk returns a tuple with the BaseSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSz

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) SetBaseSz(v string)`

SetBaseSz sets BaseSz field to given value.

### HasBaseSz

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) HasBaseSz() bool`

HasBaseSz returns a boolean if a field has been set.

### GetClQReqId

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetClQReqId() string`

GetClQReqId returns the ClQReqId field if non-nil, zero value otherwise.

### GetClQReqIdOk

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetClQReqIdOk() (*string, bool)`

GetClQReqIdOk returns a tuple with the ClQReqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClQReqId

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) SetClQReqId(v string)`

SetClQReqId sets ClQReqId field to given value.

### HasClQReqId

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) HasClQReqId() bool`

HasClQReqId returns a boolean if a field has been set.

### GetCnvtPx

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetCnvtPx() string`

GetCnvtPx returns the CnvtPx field if non-nil, zero value otherwise.

### GetCnvtPxOk

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetCnvtPxOk() (*string, bool)`

GetCnvtPxOk returns a tuple with the CnvtPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnvtPx

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) SetCnvtPx(v string)`

SetCnvtPx sets CnvtPx field to given value.

### HasCnvtPx

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) HasCnvtPx() bool`

HasCnvtPx returns a boolean if a field has been set.

### GetOrigRfqSz

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetOrigRfqSz() string`

GetOrigRfqSz returns the OrigRfqSz field if non-nil, zero value otherwise.

### GetOrigRfqSzOk

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetOrigRfqSzOk() (*string, bool)`

GetOrigRfqSzOk returns a tuple with the OrigRfqSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigRfqSz

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) SetOrigRfqSz(v string)`

SetOrigRfqSz sets OrigRfqSz field to given value.

### HasOrigRfqSz

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) HasOrigRfqSz() bool`

HasOrigRfqSz returns a boolean if a field has been set.

### GetQuoteCcy

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetQuoteCcy() string`

GetQuoteCcy returns the QuoteCcy field if non-nil, zero value otherwise.

### GetQuoteCcyOk

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetQuoteCcyOk() (*string, bool)`

GetQuoteCcyOk returns a tuple with the QuoteCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCcy

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) SetQuoteCcy(v string)`

SetQuoteCcy sets QuoteCcy field to given value.

### HasQuoteCcy

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) HasQuoteCcy() bool`

HasQuoteCcy returns a boolean if a field has been set.

### GetQuoteId

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.

### HasQuoteId

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) HasQuoteId() bool`

HasQuoteId returns a boolean if a field has been set.

### GetQuoteSz

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetQuoteSz() string`

GetQuoteSz returns the QuoteSz field if non-nil, zero value otherwise.

### GetQuoteSzOk

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetQuoteSzOk() (*string, bool)`

GetQuoteSzOk returns a tuple with the QuoteSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteSz

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) SetQuoteSz(v string)`

SetQuoteSz sets QuoteSz field to given value.

### HasQuoteSz

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) HasQuoteSz() bool`

HasQuoteSz returns a boolean if a field has been set.

### GetQuoteTime

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetQuoteTime() string`

GetQuoteTime returns the QuoteTime field if non-nil, zero value otherwise.

### GetQuoteTimeOk

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetQuoteTimeOk() (*string, bool)`

GetQuoteTimeOk returns a tuple with the QuoteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteTime

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) SetQuoteTime(v string)`

SetQuoteTime sets QuoteTime field to given value.

### HasQuoteTime

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) HasQuoteTime() bool`

HasQuoteTime returns a boolean if a field has been set.

### GetRfqSz

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetRfqSz() string`

GetRfqSz returns the RfqSz field if non-nil, zero value otherwise.

### GetRfqSzOk

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetRfqSzOk() (*string, bool)`

GetRfqSzOk returns a tuple with the RfqSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfqSz

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) SetRfqSz(v string)`

SetRfqSz sets RfqSz field to given value.

### HasRfqSz

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) HasRfqSz() bool`

HasRfqSz returns a boolean if a field has been set.

### GetRfqSzCcy

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetRfqSzCcy() string`

GetRfqSzCcy returns the RfqSzCcy field if non-nil, zero value otherwise.

### GetRfqSzCcyOk

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetRfqSzCcyOk() (*string, bool)`

GetRfqSzCcyOk returns a tuple with the RfqSzCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfqSzCcy

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) SetRfqSzCcy(v string)`

SetRfqSzCcy sets RfqSzCcy field to given value.

### HasRfqSzCcy

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) HasRfqSzCcy() bool`

HasRfqSzCcy returns a boolean if a field has been set.

### GetSide

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetTtlMs

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetTtlMs() string`

GetTtlMs returns the TtlMs field if non-nil, zero value otherwise.

### GetTtlMsOk

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) GetTtlMsOk() (*string, bool)`

GetTtlMsOk returns a tuple with the TtlMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlMs

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) SetTtlMs(v string)`

SetTtlMs sets TtlMs field to given value.

### HasTtlMs

`func (o *CreateAssetConvertEstimateQuoteV5RespDataInner) HasTtlMs() bool`

HasTtlMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


