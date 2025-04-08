# CreateAssetConvertEstimateQuoteV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCcy** | **string** | Base currency, e.g. &#x60;BTC&#x60; in &#x60;BTC-USDT&#x60; | [default to ""]
**ClQReqId** | Pointer to **string** | Client Order ID as assigned by the client  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. | [optional] [default to ""]
**QuoteCcy** | **string** | Quote currency, e.g. &#x60;USDT&#x60; in &#x60;BTC-USDT&#x60; | [default to ""]
**RfqSz** | **string** | RFQ amount | [default to ""]
**RfqSzCcy** | **string** | RFQ currency | [default to ""]
**Side** | **string** | Trade side based on &#x60;baseCcy&#x60;  &#x60;buy&#x60; &#x60;sell&#x60; | [default to ""]
**Tag** | Pointer to **string** | Order tag  Applicable to broker user | [optional] [default to ""]

## Methods

### NewCreateAssetConvertEstimateQuoteV5Req

`func NewCreateAssetConvertEstimateQuoteV5Req(baseCcy string, quoteCcy string, rfqSz string, rfqSzCcy string, side string, ) *CreateAssetConvertEstimateQuoteV5Req`

NewCreateAssetConvertEstimateQuoteV5Req instantiates a new CreateAssetConvertEstimateQuoteV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAssetConvertEstimateQuoteV5ReqWithDefaults

`func NewCreateAssetConvertEstimateQuoteV5ReqWithDefaults() *CreateAssetConvertEstimateQuoteV5Req`

NewCreateAssetConvertEstimateQuoteV5ReqWithDefaults instantiates a new CreateAssetConvertEstimateQuoteV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseCcy

`func (o *CreateAssetConvertEstimateQuoteV5Req) GetBaseCcy() string`

GetBaseCcy returns the BaseCcy field if non-nil, zero value otherwise.

### GetBaseCcyOk

`func (o *CreateAssetConvertEstimateQuoteV5Req) GetBaseCcyOk() (*string, bool)`

GetBaseCcyOk returns a tuple with the BaseCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCcy

`func (o *CreateAssetConvertEstimateQuoteV5Req) SetBaseCcy(v string)`

SetBaseCcy sets BaseCcy field to given value.


### GetClQReqId

`func (o *CreateAssetConvertEstimateQuoteV5Req) GetClQReqId() string`

GetClQReqId returns the ClQReqId field if non-nil, zero value otherwise.

### GetClQReqIdOk

`func (o *CreateAssetConvertEstimateQuoteV5Req) GetClQReqIdOk() (*string, bool)`

GetClQReqIdOk returns a tuple with the ClQReqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClQReqId

`func (o *CreateAssetConvertEstimateQuoteV5Req) SetClQReqId(v string)`

SetClQReqId sets ClQReqId field to given value.

### HasClQReqId

`func (o *CreateAssetConvertEstimateQuoteV5Req) HasClQReqId() bool`

HasClQReqId returns a boolean if a field has been set.

### GetQuoteCcy

`func (o *CreateAssetConvertEstimateQuoteV5Req) GetQuoteCcy() string`

GetQuoteCcy returns the QuoteCcy field if non-nil, zero value otherwise.

### GetQuoteCcyOk

`func (o *CreateAssetConvertEstimateQuoteV5Req) GetQuoteCcyOk() (*string, bool)`

GetQuoteCcyOk returns a tuple with the QuoteCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCcy

`func (o *CreateAssetConvertEstimateQuoteV5Req) SetQuoteCcy(v string)`

SetQuoteCcy sets QuoteCcy field to given value.


### GetRfqSz

`func (o *CreateAssetConvertEstimateQuoteV5Req) GetRfqSz() string`

GetRfqSz returns the RfqSz field if non-nil, zero value otherwise.

### GetRfqSzOk

`func (o *CreateAssetConvertEstimateQuoteV5Req) GetRfqSzOk() (*string, bool)`

GetRfqSzOk returns a tuple with the RfqSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfqSz

`func (o *CreateAssetConvertEstimateQuoteV5Req) SetRfqSz(v string)`

SetRfqSz sets RfqSz field to given value.


### GetRfqSzCcy

`func (o *CreateAssetConvertEstimateQuoteV5Req) GetRfqSzCcy() string`

GetRfqSzCcy returns the RfqSzCcy field if non-nil, zero value otherwise.

### GetRfqSzCcyOk

`func (o *CreateAssetConvertEstimateQuoteV5Req) GetRfqSzCcyOk() (*string, bool)`

GetRfqSzCcyOk returns a tuple with the RfqSzCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfqSzCcy

`func (o *CreateAssetConvertEstimateQuoteV5Req) SetRfqSzCcy(v string)`

SetRfqSzCcy sets RfqSzCcy field to given value.


### GetSide

`func (o *CreateAssetConvertEstimateQuoteV5Req) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateAssetConvertEstimateQuoteV5Req) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateAssetConvertEstimateQuoteV5Req) SetSide(v string)`

SetSide sets Side field to given value.


### GetTag

`func (o *CreateAssetConvertEstimateQuoteV5Req) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateAssetConvertEstimateQuoteV5Req) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateAssetConvertEstimateQuoteV5Req) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateAssetConvertEstimateQuoteV5Req) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


