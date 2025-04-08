# GetAssetBillsV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bal** | Pointer to **string** | Balance at the account level | [optional] [default to ""]
**BalChg** | Pointer to **string** | Change in balance at the account level | [optional] [default to ""]
**BillId** | Pointer to **string** | Bill ID | [optional] [default to ""]
**Ccy** | Pointer to **string** | Account balance currency | [optional] [default to ""]
**ClientId** | Pointer to **string** | Client-supplied ID for transfer or withdrawal | [optional] [default to ""]
**Ts** | Pointer to **string** | Creation time, Unix timestamp format in milliseconds, e.g.&#x60;1597026383085&#x60; | [optional] [default to ""]
**Type** | Pointer to **string** | Bill type | [optional] [default to ""]

## Methods

### NewGetAssetBillsV5RespData

`func NewGetAssetBillsV5RespData() *GetAssetBillsV5RespData`

NewGetAssetBillsV5RespData instantiates a new GetAssetBillsV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetBillsV5RespDataWithDefaults

`func NewGetAssetBillsV5RespDataWithDefaults() *GetAssetBillsV5RespData`

NewGetAssetBillsV5RespDataWithDefaults instantiates a new GetAssetBillsV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBal

`func (o *GetAssetBillsV5RespData) GetBal() string`

GetBal returns the Bal field if non-nil, zero value otherwise.

### GetBalOk

`func (o *GetAssetBillsV5RespData) GetBalOk() (*string, bool)`

GetBalOk returns a tuple with the Bal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBal

`func (o *GetAssetBillsV5RespData) SetBal(v string)`

SetBal sets Bal field to given value.

### HasBal

`func (o *GetAssetBillsV5RespData) HasBal() bool`

HasBal returns a boolean if a field has been set.

### GetBalChg

`func (o *GetAssetBillsV5RespData) GetBalChg() string`

GetBalChg returns the BalChg field if non-nil, zero value otherwise.

### GetBalChgOk

`func (o *GetAssetBillsV5RespData) GetBalChgOk() (*string, bool)`

GetBalChgOk returns a tuple with the BalChg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalChg

`func (o *GetAssetBillsV5RespData) SetBalChg(v string)`

SetBalChg sets BalChg field to given value.

### HasBalChg

`func (o *GetAssetBillsV5RespData) HasBalChg() bool`

HasBalChg returns a boolean if a field has been set.

### GetBillId

`func (o *GetAssetBillsV5RespData) GetBillId() string`

GetBillId returns the BillId field if non-nil, zero value otherwise.

### GetBillIdOk

`func (o *GetAssetBillsV5RespData) GetBillIdOk() (*string, bool)`

GetBillIdOk returns a tuple with the BillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillId

`func (o *GetAssetBillsV5RespData) SetBillId(v string)`

SetBillId sets BillId field to given value.

### HasBillId

`func (o *GetAssetBillsV5RespData) HasBillId() bool`

HasBillId returns a boolean if a field has been set.

### GetCcy

`func (o *GetAssetBillsV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAssetBillsV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAssetBillsV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAssetBillsV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetClientId

`func (o *GetAssetBillsV5RespData) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetAssetBillsV5RespData) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetAssetBillsV5RespData) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GetAssetBillsV5RespData) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetTs

`func (o *GetAssetBillsV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetAssetBillsV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetAssetBillsV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetAssetBillsV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetType

`func (o *GetAssetBillsV5RespData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAssetBillsV5RespData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAssetBillsV5RespData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAssetBillsV5RespData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


