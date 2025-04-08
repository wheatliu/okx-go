# GetAssetSubaccountBillsV5RespDataInner

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

### NewGetAssetSubaccountBillsV5RespDataInner

`func NewGetAssetSubaccountBillsV5RespDataInner() *GetAssetSubaccountBillsV5RespDataInner`

NewGetAssetSubaccountBillsV5RespDataInner instantiates a new GetAssetSubaccountBillsV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetSubaccountBillsV5RespDataInnerWithDefaults

`func NewGetAssetSubaccountBillsV5RespDataInnerWithDefaults() *GetAssetSubaccountBillsV5RespDataInner`

NewGetAssetSubaccountBillsV5RespDataInnerWithDefaults instantiates a new GetAssetSubaccountBillsV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *GetAssetSubaccountBillsV5RespDataInner) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetAssetSubaccountBillsV5RespDataInner) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetAssetSubaccountBillsV5RespDataInner) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetAssetSubaccountBillsV5RespDataInner) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetBillId

`func (o *GetAssetSubaccountBillsV5RespDataInner) GetBillId() string`

GetBillId returns the BillId field if non-nil, zero value otherwise.

### GetBillIdOk

`func (o *GetAssetSubaccountBillsV5RespDataInner) GetBillIdOk() (*string, bool)`

GetBillIdOk returns a tuple with the BillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillId

`func (o *GetAssetSubaccountBillsV5RespDataInner) SetBillId(v string)`

SetBillId sets BillId field to given value.

### HasBillId

`func (o *GetAssetSubaccountBillsV5RespDataInner) HasBillId() bool`

HasBillId returns a boolean if a field has been set.

### GetCcy

`func (o *GetAssetSubaccountBillsV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAssetSubaccountBillsV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAssetSubaccountBillsV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAssetSubaccountBillsV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetSubAcct

`func (o *GetAssetSubaccountBillsV5RespDataInner) GetSubAcct() string`

GetSubAcct returns the SubAcct field if non-nil, zero value otherwise.

### GetSubAcctOk

`func (o *GetAssetSubaccountBillsV5RespDataInner) GetSubAcctOk() (*string, bool)`

GetSubAcctOk returns a tuple with the SubAcct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAcct

`func (o *GetAssetSubaccountBillsV5RespDataInner) SetSubAcct(v string)`

SetSubAcct sets SubAcct field to given value.

### HasSubAcct

`func (o *GetAssetSubaccountBillsV5RespDataInner) HasSubAcct() bool`

HasSubAcct returns a boolean if a field has been set.

### GetTs

`func (o *GetAssetSubaccountBillsV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetAssetSubaccountBillsV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetAssetSubaccountBillsV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetAssetSubaccountBillsV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetType

`func (o *GetAssetSubaccountBillsV5RespDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAssetSubaccountBillsV5RespDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAssetSubaccountBillsV5RespDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAssetSubaccountBillsV5RespDataInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


