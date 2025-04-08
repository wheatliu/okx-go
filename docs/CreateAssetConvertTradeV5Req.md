# CreateAssetConvertTradeV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCcy** | **string** | Base currency, e.g. &#x60;BTC&#x60; in &#x60;BTC-USDT&#x60; | [default to ""]
**ClTReqId** | Pointer to **string** | Client Order ID as assigned by the client  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. | [optional] [default to ""]
**QuoteCcy** | **string** | Quote currency, e.g. &#x60;USDT&#x60; in &#x60;BTC-USDT&#x60; | [default to ""]
**QuoteId** | **string** | Quote ID | [default to ""]
**Side** | **string** | Trade side based on &#x60;baseCcy&#x60;  &#x60;buy&#x60; &#x60;sell&#x60; | [default to ""]
**Sz** | **string** | Quote amount  The quote amount should no more then RFQ amount | [default to ""]
**SzCcy** | **string** | Quote currency | [default to ""]
**Tag** | Pointer to **string** | Order tag  Applicable to broker user | [optional] [default to ""]

## Methods

### NewCreateAssetConvertTradeV5Req

`func NewCreateAssetConvertTradeV5Req(baseCcy string, quoteCcy string, quoteId string, side string, sz string, szCcy string, ) *CreateAssetConvertTradeV5Req`

NewCreateAssetConvertTradeV5Req instantiates a new CreateAssetConvertTradeV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAssetConvertTradeV5ReqWithDefaults

`func NewCreateAssetConvertTradeV5ReqWithDefaults() *CreateAssetConvertTradeV5Req`

NewCreateAssetConvertTradeV5ReqWithDefaults instantiates a new CreateAssetConvertTradeV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseCcy

`func (o *CreateAssetConvertTradeV5Req) GetBaseCcy() string`

GetBaseCcy returns the BaseCcy field if non-nil, zero value otherwise.

### GetBaseCcyOk

`func (o *CreateAssetConvertTradeV5Req) GetBaseCcyOk() (*string, bool)`

GetBaseCcyOk returns a tuple with the BaseCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCcy

`func (o *CreateAssetConvertTradeV5Req) SetBaseCcy(v string)`

SetBaseCcy sets BaseCcy field to given value.


### GetClTReqId

`func (o *CreateAssetConvertTradeV5Req) GetClTReqId() string`

GetClTReqId returns the ClTReqId field if non-nil, zero value otherwise.

### GetClTReqIdOk

`func (o *CreateAssetConvertTradeV5Req) GetClTReqIdOk() (*string, bool)`

GetClTReqIdOk returns a tuple with the ClTReqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClTReqId

`func (o *CreateAssetConvertTradeV5Req) SetClTReqId(v string)`

SetClTReqId sets ClTReqId field to given value.

### HasClTReqId

`func (o *CreateAssetConvertTradeV5Req) HasClTReqId() bool`

HasClTReqId returns a boolean if a field has been set.

### GetQuoteCcy

`func (o *CreateAssetConvertTradeV5Req) GetQuoteCcy() string`

GetQuoteCcy returns the QuoteCcy field if non-nil, zero value otherwise.

### GetQuoteCcyOk

`func (o *CreateAssetConvertTradeV5Req) GetQuoteCcyOk() (*string, bool)`

GetQuoteCcyOk returns a tuple with the QuoteCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCcy

`func (o *CreateAssetConvertTradeV5Req) SetQuoteCcy(v string)`

SetQuoteCcy sets QuoteCcy field to given value.


### GetQuoteId

`func (o *CreateAssetConvertTradeV5Req) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *CreateAssetConvertTradeV5Req) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *CreateAssetConvertTradeV5Req) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetSide

`func (o *CreateAssetConvertTradeV5Req) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateAssetConvertTradeV5Req) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateAssetConvertTradeV5Req) SetSide(v string)`

SetSide sets Side field to given value.


### GetSz

`func (o *CreateAssetConvertTradeV5Req) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *CreateAssetConvertTradeV5Req) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *CreateAssetConvertTradeV5Req) SetSz(v string)`

SetSz sets Sz field to given value.


### GetSzCcy

`func (o *CreateAssetConvertTradeV5Req) GetSzCcy() string`

GetSzCcy returns the SzCcy field if non-nil, zero value otherwise.

### GetSzCcyOk

`func (o *CreateAssetConvertTradeV5Req) GetSzCcyOk() (*string, bool)`

GetSzCcyOk returns a tuple with the SzCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSzCcy

`func (o *CreateAssetConvertTradeV5Req) SetSzCcy(v string)`

SetSzCcy sets SzCcy field to given value.


### GetTag

`func (o *CreateAssetConvertTradeV5Req) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateAssetConvertTradeV5Req) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateAssetConvertTradeV5Req) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateAssetConvertTradeV5Req) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


