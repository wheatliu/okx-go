# GetUsersSubaccountListV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanTransOut** | Pointer to **bool** | Whether the sub-account has the right to transfer out.   &#x60;true&#x60;: can transfer out   &#x60;false&#x60;: cannot transfer out | [optional] 
**Enable** | Pointer to **bool** | Sub-account status  &#x60;true&#x60;: Normal  &#x60;false&#x60;: Frozen (global) | [optional] 
**FirstLvSubAcct** | Pointer to **string** | The first level sub-account.   For subAcctLv: 1, firstLvSubAcct is equal to subAcct  For subAcctLv: 2, subAcct belongs to firstLvSubAcct. | [optional] [default to ""]
**FrozenFunc** | Pointer to **[]string** | Frozen functions  &#x60;trading&#x60;  &#x60;convert&#x60;  &#x60;transfer&#x60;  &#x60;withdrawal&#x60;  &#x60;deposit&#x60;  &#x60;flexible_loan&#x60; | [optional] 
**GAuth** | Pointer to **bool** | If the sub-account switches on the Google Authenticator for login authentication.   &#x60;true&#x60;: On  &#x60;false&#x60;: Off | [optional] 
**IfDma** | Pointer to **bool** | Whether it is dma broker sub-account.   &#x60;true&#x60;: Dma broker sub-account   &#x60;false&#x60;: It is not dma broker sub-account. | [optional] 
**Label** | Pointer to **string** | Sub-account note | [optional] [default to ""]
**Mobile** | Pointer to **string** | Mobile number that linked with the sub-account. | [optional] [default to ""]
**SubAcct** | Pointer to **string** | Sub-account name | [optional] [default to ""]
**SubAcctLv** | Pointer to **string** | Sub-account level   &#x60;1&#x60;: First level sub-account  &#x60;2&#x60;: Second level sub-account. | [optional] [default to ""]
**Ts** | Pointer to **string** | Sub-account creation time, Unix timestamp in millisecond format. e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Type** | Pointer to **string** | Sub-account type   &#x60;1&#x60;: Standard sub-account   &#x60;2&#x60;: Managed trading sub-account   &#x60;5&#x60;: Custody trading sub-account - Copper  &#x60;9&#x60;: Managed trading sub-account - Copper   &#x60;12&#x60;: Custody trading sub-account - Komainu | [optional] [default to ""]
**Uid** | Pointer to **string** | Sub-account uid | [optional] [default to ""]

## Methods

### NewGetUsersSubaccountListV5RespDataInner

`func NewGetUsersSubaccountListV5RespDataInner() *GetUsersSubaccountListV5RespDataInner`

NewGetUsersSubaccountListV5RespDataInner instantiates a new GetUsersSubaccountListV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUsersSubaccountListV5RespDataInnerWithDefaults

`func NewGetUsersSubaccountListV5RespDataInnerWithDefaults() *GetUsersSubaccountListV5RespDataInner`

NewGetUsersSubaccountListV5RespDataInnerWithDefaults instantiates a new GetUsersSubaccountListV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanTransOut

`func (o *GetUsersSubaccountListV5RespDataInner) GetCanTransOut() bool`

GetCanTransOut returns the CanTransOut field if non-nil, zero value otherwise.

### GetCanTransOutOk

`func (o *GetUsersSubaccountListV5RespDataInner) GetCanTransOutOk() (*bool, bool)`

GetCanTransOutOk returns a tuple with the CanTransOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTransOut

`func (o *GetUsersSubaccountListV5RespDataInner) SetCanTransOut(v bool)`

SetCanTransOut sets CanTransOut field to given value.

### HasCanTransOut

`func (o *GetUsersSubaccountListV5RespDataInner) HasCanTransOut() bool`

HasCanTransOut returns a boolean if a field has been set.

### GetEnable

`func (o *GetUsersSubaccountListV5RespDataInner) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GetUsersSubaccountListV5RespDataInner) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GetUsersSubaccountListV5RespDataInner) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GetUsersSubaccountListV5RespDataInner) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetFirstLvSubAcct

`func (o *GetUsersSubaccountListV5RespDataInner) GetFirstLvSubAcct() string`

GetFirstLvSubAcct returns the FirstLvSubAcct field if non-nil, zero value otherwise.

### GetFirstLvSubAcctOk

`func (o *GetUsersSubaccountListV5RespDataInner) GetFirstLvSubAcctOk() (*string, bool)`

GetFirstLvSubAcctOk returns a tuple with the FirstLvSubAcct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstLvSubAcct

`func (o *GetUsersSubaccountListV5RespDataInner) SetFirstLvSubAcct(v string)`

SetFirstLvSubAcct sets FirstLvSubAcct field to given value.

### HasFirstLvSubAcct

`func (o *GetUsersSubaccountListV5RespDataInner) HasFirstLvSubAcct() bool`

HasFirstLvSubAcct returns a boolean if a field has been set.

### GetFrozenFunc

`func (o *GetUsersSubaccountListV5RespDataInner) GetFrozenFunc() []string`

GetFrozenFunc returns the FrozenFunc field if non-nil, zero value otherwise.

### GetFrozenFuncOk

`func (o *GetUsersSubaccountListV5RespDataInner) GetFrozenFuncOk() (*[]string, bool)`

GetFrozenFuncOk returns a tuple with the FrozenFunc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenFunc

`func (o *GetUsersSubaccountListV5RespDataInner) SetFrozenFunc(v []string)`

SetFrozenFunc sets FrozenFunc field to given value.

### HasFrozenFunc

`func (o *GetUsersSubaccountListV5RespDataInner) HasFrozenFunc() bool`

HasFrozenFunc returns a boolean if a field has been set.

### GetGAuth

`func (o *GetUsersSubaccountListV5RespDataInner) GetGAuth() bool`

GetGAuth returns the GAuth field if non-nil, zero value otherwise.

### GetGAuthOk

`func (o *GetUsersSubaccountListV5RespDataInner) GetGAuthOk() (*bool, bool)`

GetGAuthOk returns a tuple with the GAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGAuth

`func (o *GetUsersSubaccountListV5RespDataInner) SetGAuth(v bool)`

SetGAuth sets GAuth field to given value.

### HasGAuth

`func (o *GetUsersSubaccountListV5RespDataInner) HasGAuth() bool`

HasGAuth returns a boolean if a field has been set.

### GetIfDma

`func (o *GetUsersSubaccountListV5RespDataInner) GetIfDma() bool`

GetIfDma returns the IfDma field if non-nil, zero value otherwise.

### GetIfDmaOk

`func (o *GetUsersSubaccountListV5RespDataInner) GetIfDmaOk() (*bool, bool)`

GetIfDmaOk returns a tuple with the IfDma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfDma

`func (o *GetUsersSubaccountListV5RespDataInner) SetIfDma(v bool)`

SetIfDma sets IfDma field to given value.

### HasIfDma

`func (o *GetUsersSubaccountListV5RespDataInner) HasIfDma() bool`

HasIfDma returns a boolean if a field has been set.

### GetLabel

`func (o *GetUsersSubaccountListV5RespDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetUsersSubaccountListV5RespDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetUsersSubaccountListV5RespDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetUsersSubaccountListV5RespDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMobile

`func (o *GetUsersSubaccountListV5RespDataInner) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *GetUsersSubaccountListV5RespDataInner) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *GetUsersSubaccountListV5RespDataInner) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *GetUsersSubaccountListV5RespDataInner) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetSubAcct

`func (o *GetUsersSubaccountListV5RespDataInner) GetSubAcct() string`

GetSubAcct returns the SubAcct field if non-nil, zero value otherwise.

### GetSubAcctOk

`func (o *GetUsersSubaccountListV5RespDataInner) GetSubAcctOk() (*string, bool)`

GetSubAcctOk returns a tuple with the SubAcct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAcct

`func (o *GetUsersSubaccountListV5RespDataInner) SetSubAcct(v string)`

SetSubAcct sets SubAcct field to given value.

### HasSubAcct

`func (o *GetUsersSubaccountListV5RespDataInner) HasSubAcct() bool`

HasSubAcct returns a boolean if a field has been set.

### GetSubAcctLv

`func (o *GetUsersSubaccountListV5RespDataInner) GetSubAcctLv() string`

GetSubAcctLv returns the SubAcctLv field if non-nil, zero value otherwise.

### GetSubAcctLvOk

`func (o *GetUsersSubaccountListV5RespDataInner) GetSubAcctLvOk() (*string, bool)`

GetSubAcctLvOk returns a tuple with the SubAcctLv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAcctLv

`func (o *GetUsersSubaccountListV5RespDataInner) SetSubAcctLv(v string)`

SetSubAcctLv sets SubAcctLv field to given value.

### HasSubAcctLv

`func (o *GetUsersSubaccountListV5RespDataInner) HasSubAcctLv() bool`

HasSubAcctLv returns a boolean if a field has been set.

### GetTs

`func (o *GetUsersSubaccountListV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetUsersSubaccountListV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetUsersSubaccountListV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetUsersSubaccountListV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetType

`func (o *GetUsersSubaccountListV5RespDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetUsersSubaccountListV5RespDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetUsersSubaccountListV5RespDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetUsersSubaccountListV5RespDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUid

`func (o *GetUsersSubaccountListV5RespDataInner) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *GetUsersSubaccountListV5RespDataInner) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *GetUsersSubaccountListV5RespDataInner) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *GetUsersSubaccountListV5RespDataInner) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


