# GetAssetSubaccountManagedSubaccountBillsV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | Pointer to **string** | Transfer amount | [optional] [default to ""]
**BillId** | Pointer to **string** | Bill ID | [optional] [default to ""]
**Ccy** | Pointer to **string** | Transfer currency | [optional] [default to ""]
**SubAcct** | Pointer to **string** | Sub-account name | [optional] [default to ""]
**SubUid** | Pointer to **string** | Sub-account UID | [optional] [default to ""]
**Ts** | Pointer to **string** | Bill ID creation time, Unix timestamp in millisecond format, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Type** | Pointer to **string** | Bill type | [optional] [default to ""]

## Methods

### NewGetAssetSubaccountManagedSubaccountBillsV5RespData

`func NewGetAssetSubaccountManagedSubaccountBillsV5RespData() *GetAssetSubaccountManagedSubaccountBillsV5RespData`

NewGetAssetSubaccountManagedSubaccountBillsV5RespData instantiates a new GetAssetSubaccountManagedSubaccountBillsV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetSubaccountManagedSubaccountBillsV5RespDataWithDefaults

`func NewGetAssetSubaccountManagedSubaccountBillsV5RespDataWithDefaults() *GetAssetSubaccountManagedSubaccountBillsV5RespData`

NewGetAssetSubaccountManagedSubaccountBillsV5RespDataWithDefaults instantiates a new GetAssetSubaccountManagedSubaccountBillsV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetBillId

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) GetBillId() string`

GetBillId returns the BillId field if non-nil, zero value otherwise.

### GetBillIdOk

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) GetBillIdOk() (*string, bool)`

GetBillIdOk returns a tuple with the BillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillId

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) SetBillId(v string)`

SetBillId sets BillId field to given value.

### HasBillId

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) HasBillId() bool`

HasBillId returns a boolean if a field has been set.

### GetCcy

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetSubAcct

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) GetSubAcct() string`

GetSubAcct returns the SubAcct field if non-nil, zero value otherwise.

### GetSubAcctOk

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) GetSubAcctOk() (*string, bool)`

GetSubAcctOk returns a tuple with the SubAcct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAcct

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) SetSubAcct(v string)`

SetSubAcct sets SubAcct field to given value.

### HasSubAcct

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) HasSubAcct() bool`

HasSubAcct returns a boolean if a field has been set.

### GetSubUid

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) GetSubUid() string`

GetSubUid returns the SubUid field if non-nil, zero value otherwise.

### GetSubUidOk

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) GetSubUidOk() (*string, bool)`

GetSubUidOk returns a tuple with the SubUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubUid

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) SetSubUid(v string)`

SetSubUid sets SubUid field to given value.

### HasSubUid

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) HasSubUid() bool`

HasSubUid returns a boolean if a field has been set.

### GetTs

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetType

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAssetSubaccountManagedSubaccountBillsV5RespData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


