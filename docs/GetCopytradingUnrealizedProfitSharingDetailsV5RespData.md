# GetCopytradingUnrealizedProfitSharingDetailsV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | Pointer to **string** | The currency of profit sharing. e.g. USDT | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**NickName** | Pointer to **string** | Nickname of copy trader. | [optional] [default to ""]
**PortLink** | Pointer to **string** | Portrait link | [optional] [default to ""]
**Ts** | Pointer to **string** | Update time. | [optional] [default to ""]
**UnrealizedProfitSharingAmt** | Pointer to **string** | Unrealized profit sharing amount. | [optional] [default to ""]

## Methods

### NewGetCopytradingUnrealizedProfitSharingDetailsV5RespData

`func NewGetCopytradingUnrealizedProfitSharingDetailsV5RespData() *GetCopytradingUnrealizedProfitSharingDetailsV5RespData`

NewGetCopytradingUnrealizedProfitSharingDetailsV5RespData instantiates a new GetCopytradingUnrealizedProfitSharingDetailsV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCopytradingUnrealizedProfitSharingDetailsV5RespDataWithDefaults

`func NewGetCopytradingUnrealizedProfitSharingDetailsV5RespDataWithDefaults() *GetCopytradingUnrealizedProfitSharingDetailsV5RespData`

NewGetCopytradingUnrealizedProfitSharingDetailsV5RespDataWithDefaults instantiates a new GetCopytradingUnrealizedProfitSharingDetailsV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetInstType

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetNickName

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### GetPortLink

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) GetPortLink() string`

GetPortLink returns the PortLink field if non-nil, zero value otherwise.

### GetPortLinkOk

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) GetPortLinkOk() (*string, bool)`

GetPortLinkOk returns a tuple with the PortLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortLink

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) SetPortLink(v string)`

SetPortLink sets PortLink field to given value.

### HasPortLink

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) HasPortLink() bool`

HasPortLink returns a boolean if a field has been set.

### GetTs

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetUnrealizedProfitSharingAmt

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) GetUnrealizedProfitSharingAmt() string`

GetUnrealizedProfitSharingAmt returns the UnrealizedProfitSharingAmt field if non-nil, zero value otherwise.

### GetUnrealizedProfitSharingAmtOk

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) GetUnrealizedProfitSharingAmtOk() (*string, bool)`

GetUnrealizedProfitSharingAmtOk returns a tuple with the UnrealizedProfitSharingAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitSharingAmt

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) SetUnrealizedProfitSharingAmt(v string)`

SetUnrealizedProfitSharingAmt sets UnrealizedProfitSharingAmt field to given value.

### HasUnrealizedProfitSharingAmt

`func (o *GetCopytradingUnrealizedProfitSharingDetailsV5RespData) HasUnrealizedProfitSharingAmt() bool`

HasUnrealizedProfitSharingAmt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


