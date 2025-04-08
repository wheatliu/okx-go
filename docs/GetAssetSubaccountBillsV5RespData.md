# GetAssetSubaccountBillsV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | Pointer to **string** | Transfer amount | [optional] [default to ""]
**BillId** | Pointer to **string** | Bill ID | [optional] [default to ""]
**Ccy** | Pointer to **string** | Transfer currency | [optional] [default to ""]
**SubAcct** | Pointer to **string** | Sub-account name | [optional] [default to ""]
**Ts** | Pointer to **string** | Bill ID creation time, Unix timestamp in millisecond format, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Type** | Pointer to **string** | Bill type | [optional] [default to ""]

## Methods

### NewGetAssetSubaccountBillsV5RespData

`func NewGetAssetSubaccountBillsV5RespData() *GetAssetSubaccountBillsV5RespData`

NewGetAssetSubaccountBillsV5RespData instantiates a new GetAssetSubaccountBillsV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetSubaccountBillsV5RespDataWithDefaults

`func NewGetAssetSubaccountBillsV5RespDataWithDefaults() *GetAssetSubaccountBillsV5RespData`

NewGetAssetSubaccountBillsV5RespDataWithDefaults instantiates a new GetAssetSubaccountBillsV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *GetAssetSubaccountBillsV5RespData) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetAssetSubaccountBillsV5RespData) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetAssetSubaccountBillsV5RespData) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetAssetSubaccountBillsV5RespData) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetBillId

`func (o *GetAssetSubaccountBillsV5RespData) GetBillId() string`

GetBillId returns the BillId field if non-nil, zero value otherwise.

### GetBillIdOk

`func (o *GetAssetSubaccountBillsV5RespData) GetBillIdOk() (*string, bool)`

GetBillIdOk returns a tuple with the BillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillId

`func (o *GetAssetSubaccountBillsV5RespData) SetBillId(v string)`

SetBillId sets BillId field to given value.

### HasBillId

`func (o *GetAssetSubaccountBillsV5RespData) HasBillId() bool`

HasBillId returns a boolean if a field has been set.

### GetCcy

`func (o *GetAssetSubaccountBillsV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAssetSubaccountBillsV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAssetSubaccountBillsV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAssetSubaccountBillsV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetSubAcct

`func (o *GetAssetSubaccountBillsV5RespData) GetSubAcct() string`

GetSubAcct returns the SubAcct field if non-nil, zero value otherwise.

### GetSubAcctOk

`func (o *GetAssetSubaccountBillsV5RespData) GetSubAcctOk() (*string, bool)`

GetSubAcctOk returns a tuple with the SubAcct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAcct

`func (o *GetAssetSubaccountBillsV5RespData) SetSubAcct(v string)`

SetSubAcct sets SubAcct field to given value.

### HasSubAcct

`func (o *GetAssetSubaccountBillsV5RespData) HasSubAcct() bool`

HasSubAcct returns a boolean if a field has been set.

### GetTs

`func (o *GetAssetSubaccountBillsV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetAssetSubaccountBillsV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetAssetSubaccountBillsV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetAssetSubaccountBillsV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetType

`func (o *GetAssetSubaccountBillsV5RespData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAssetSubaccountBillsV5RespData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAssetSubaccountBillsV5RespData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAssetSubaccountBillsV5RespData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


