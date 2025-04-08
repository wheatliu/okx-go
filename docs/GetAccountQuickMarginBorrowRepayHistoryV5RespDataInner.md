# GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccBorrowed** | Pointer to **string** | Accumulate borrow amount | [optional] [default to ""]
**Amt** | Pointer to **string** | borrow/repay amount | [optional] [default to ""]
**Ccy** | Pointer to **string** | Loan currency, e.g. &#x60;BTC&#x60; | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. BTC-USDT | [optional] [default to ""]
**RefId** | Pointer to **string** | The ID of borrowing or repayment | [optional] [default to ""]
**Side** | Pointer to **string** | &#x60;borrow&#x60;  &#x60;repay&#x60; | [optional] [default to ""]
**Ts** | Pointer to **string** | Timestamp for Borrow/Repay | [optional] [default to ""]

## Methods

### NewGetAccountQuickMarginBorrowRepayHistoryV5RespDataInner

`func NewGetAccountQuickMarginBorrowRepayHistoryV5RespDataInner() *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner`

NewGetAccountQuickMarginBorrowRepayHistoryV5RespDataInner instantiates a new GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountQuickMarginBorrowRepayHistoryV5RespDataInnerWithDefaults

`func NewGetAccountQuickMarginBorrowRepayHistoryV5RespDataInnerWithDefaults() *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner`

NewGetAccountQuickMarginBorrowRepayHistoryV5RespDataInnerWithDefaults instantiates a new GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccBorrowed

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) GetAccBorrowed() string`

GetAccBorrowed returns the AccBorrowed field if non-nil, zero value otherwise.

### GetAccBorrowedOk

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) GetAccBorrowedOk() (*string, bool)`

GetAccBorrowedOk returns a tuple with the AccBorrowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccBorrowed

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) SetAccBorrowed(v string)`

SetAccBorrowed sets AccBorrowed field to given value.

### HasAccBorrowed

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) HasAccBorrowed() bool`

HasAccBorrowed returns a boolean if a field has been set.

### GetAmt

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetCcy

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetInstId

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetRefId

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetSide

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetTs

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetAccountQuickMarginBorrowRepayHistoryV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


