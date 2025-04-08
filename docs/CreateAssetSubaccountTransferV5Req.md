# CreateAssetSubaccountTransferV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | **string** | Transfer amount | [default to ""]
**Ccy** | **string** | Currency | [default to ""]
**From** | **string** | Account type of transfer from sub-account  &#x60;6&#x60;: Funding Account  &#x60;18&#x60;: Trading account | [default to ""]
**FromSubAccount** | **string** | Sub-account name of the account that transfers funds out. | [default to ""]
**LoanTrans** | Pointer to **bool** | Whether or not borrowed coins can be transferred out under &#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60;  The default is &#x60;false&#x60; | [optional] 
**OmitPosRisk** | Pointer to **string** | Ignore position risk  Default is &#x60;false&#x60;  Applicable to &#x60;Portfolio margin&#x60; | [optional] [default to ""]
**To** | **string** | Account type of transfer to sub-account  &#x60;6&#x60;: Funding Account  &#x60;18&#x60;: Trading account | [default to ""]
**ToSubAccount** | **string** | Sub-account name of the account that transfers funds in. | [default to ""]

## Methods

### NewCreateAssetSubaccountTransferV5Req

`func NewCreateAssetSubaccountTransferV5Req(amt string, ccy string, from string, fromSubAccount string, to string, toSubAccount string, ) *CreateAssetSubaccountTransferV5Req`

NewCreateAssetSubaccountTransferV5Req instantiates a new CreateAssetSubaccountTransferV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAssetSubaccountTransferV5ReqWithDefaults

`func NewCreateAssetSubaccountTransferV5ReqWithDefaults() *CreateAssetSubaccountTransferV5Req`

NewCreateAssetSubaccountTransferV5ReqWithDefaults instantiates a new CreateAssetSubaccountTransferV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *CreateAssetSubaccountTransferV5Req) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *CreateAssetSubaccountTransferV5Req) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *CreateAssetSubaccountTransferV5Req) SetAmt(v string)`

SetAmt sets Amt field to given value.


### GetCcy

`func (o *CreateAssetSubaccountTransferV5Req) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateAssetSubaccountTransferV5Req) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateAssetSubaccountTransferV5Req) SetCcy(v string)`

SetCcy sets Ccy field to given value.


### GetFrom

`func (o *CreateAssetSubaccountTransferV5Req) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CreateAssetSubaccountTransferV5Req) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CreateAssetSubaccountTransferV5Req) SetFrom(v string)`

SetFrom sets From field to given value.


### GetFromSubAccount

`func (o *CreateAssetSubaccountTransferV5Req) GetFromSubAccount() string`

GetFromSubAccount returns the FromSubAccount field if non-nil, zero value otherwise.

### GetFromSubAccountOk

`func (o *CreateAssetSubaccountTransferV5Req) GetFromSubAccountOk() (*string, bool)`

GetFromSubAccountOk returns a tuple with the FromSubAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromSubAccount

`func (o *CreateAssetSubaccountTransferV5Req) SetFromSubAccount(v string)`

SetFromSubAccount sets FromSubAccount field to given value.


### GetLoanTrans

`func (o *CreateAssetSubaccountTransferV5Req) GetLoanTrans() bool`

GetLoanTrans returns the LoanTrans field if non-nil, zero value otherwise.

### GetLoanTransOk

`func (o *CreateAssetSubaccountTransferV5Req) GetLoanTransOk() (*bool, bool)`

GetLoanTransOk returns a tuple with the LoanTrans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanTrans

`func (o *CreateAssetSubaccountTransferV5Req) SetLoanTrans(v bool)`

SetLoanTrans sets LoanTrans field to given value.

### HasLoanTrans

`func (o *CreateAssetSubaccountTransferV5Req) HasLoanTrans() bool`

HasLoanTrans returns a boolean if a field has been set.

### GetOmitPosRisk

`func (o *CreateAssetSubaccountTransferV5Req) GetOmitPosRisk() string`

GetOmitPosRisk returns the OmitPosRisk field if non-nil, zero value otherwise.

### GetOmitPosRiskOk

`func (o *CreateAssetSubaccountTransferV5Req) GetOmitPosRiskOk() (*string, bool)`

GetOmitPosRiskOk returns a tuple with the OmitPosRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitPosRisk

`func (o *CreateAssetSubaccountTransferV5Req) SetOmitPosRisk(v string)`

SetOmitPosRisk sets OmitPosRisk field to given value.

### HasOmitPosRisk

`func (o *CreateAssetSubaccountTransferV5Req) HasOmitPosRisk() bool`

HasOmitPosRisk returns a boolean if a field has been set.

### GetTo

`func (o *CreateAssetSubaccountTransferV5Req) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CreateAssetSubaccountTransferV5Req) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CreateAssetSubaccountTransferV5Req) SetTo(v string)`

SetTo sets To field to given value.


### GetToSubAccount

`func (o *CreateAssetSubaccountTransferV5Req) GetToSubAccount() string`

GetToSubAccount returns the ToSubAccount field if non-nil, zero value otherwise.

### GetToSubAccountOk

`func (o *CreateAssetSubaccountTransferV5Req) GetToSubAccountOk() (*string, bool)`

GetToSubAccountOk returns a tuple with the ToSubAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToSubAccount

`func (o *CreateAssetSubaccountTransferV5Req) SetToSubAccount(v string)`

SetToSubAccount sets ToSubAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


