# GetAccountSpotBorrowRepayHistoryV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccBorrowed** | Pointer to **string** | Accumulated borrow amount | [optional] [default to ""]
**Amt** | Pointer to **string** | Amount | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency, e.g. &#x60;BTC&#x60; | [optional] [default to ""]
**Ts** | Pointer to **string** | Timestamp for the event, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Type** | Pointer to **string** | Event type  &#x60;auto_borrow&#x60;  &#x60;auto_repay&#x60;  &#x60;manual_borrow&#x60;  &#x60;manual_repay&#x60; | [optional] [default to ""]

## Methods

### NewGetAccountSpotBorrowRepayHistoryV5RespData

`func NewGetAccountSpotBorrowRepayHistoryV5RespData() *GetAccountSpotBorrowRepayHistoryV5RespData`

NewGetAccountSpotBorrowRepayHistoryV5RespData instantiates a new GetAccountSpotBorrowRepayHistoryV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountSpotBorrowRepayHistoryV5RespDataWithDefaults

`func NewGetAccountSpotBorrowRepayHistoryV5RespDataWithDefaults() *GetAccountSpotBorrowRepayHistoryV5RespData`

NewGetAccountSpotBorrowRepayHistoryV5RespDataWithDefaults instantiates a new GetAccountSpotBorrowRepayHistoryV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccBorrowed

`func (o *GetAccountSpotBorrowRepayHistoryV5RespData) GetAccBorrowed() string`

GetAccBorrowed returns the AccBorrowed field if non-nil, zero value otherwise.

### GetAccBorrowedOk

`func (o *GetAccountSpotBorrowRepayHistoryV5RespData) GetAccBorrowedOk() (*string, bool)`

GetAccBorrowedOk returns a tuple with the AccBorrowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccBorrowed

`func (o *GetAccountSpotBorrowRepayHistoryV5RespData) SetAccBorrowed(v string)`

SetAccBorrowed sets AccBorrowed field to given value.

### HasAccBorrowed

`func (o *GetAccountSpotBorrowRepayHistoryV5RespData) HasAccBorrowed() bool`

HasAccBorrowed returns a boolean if a field has been set.

### GetAmt

`func (o *GetAccountSpotBorrowRepayHistoryV5RespData) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetAccountSpotBorrowRepayHistoryV5RespData) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetAccountSpotBorrowRepayHistoryV5RespData) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetAccountSpotBorrowRepayHistoryV5RespData) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetCcy

`func (o *GetAccountSpotBorrowRepayHistoryV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAccountSpotBorrowRepayHistoryV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAccountSpotBorrowRepayHistoryV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAccountSpotBorrowRepayHistoryV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetTs

`func (o *GetAccountSpotBorrowRepayHistoryV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetAccountSpotBorrowRepayHistoryV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetAccountSpotBorrowRepayHistoryV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetAccountSpotBorrowRepayHistoryV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetType

`func (o *GetAccountSpotBorrowRepayHistoryV5RespData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAccountSpotBorrowRepayHistoryV5RespData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAccountSpotBorrowRepayHistoryV5RespData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAccountSpotBorrowRepayHistoryV5RespData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


