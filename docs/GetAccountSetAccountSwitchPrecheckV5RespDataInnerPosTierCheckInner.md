# GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstFamily** | Pointer to **string** | Instrument family | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type   &#x60;SWAP&#x60;   &#x60;FUTURES&#x60;   &#x60;OPTION&#x60; | [optional] [default to ""]
**Lever** | Pointer to **string** | Leverage | [optional] [default to ""]
**MaxSz** | Pointer to **string** | If acctLv is &#x60;2/3&#x60;, it refers to the maximum position size allowed at the current leverage. If acctLv is &#x60;4&#x60;, it refers to the maximum position limit for cross-margin positions under the PM mode. | [optional] [default to ""]
**Pos** | Pointer to **string** | Quantity of position | [optional] [default to ""]

## Methods

### NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner

`func NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner() *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner`

NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner instantiates a new GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInnerWithDefaults

`func NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInnerWithDefaults() *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner`

NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInnerWithDefaults instantiates a new GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstFamily

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner) GetInstFamily() string`

GetInstFamily returns the InstFamily field if non-nil, zero value otherwise.

### GetInstFamilyOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner) GetInstFamilyOk() (*string, bool)`

GetInstFamilyOk returns a tuple with the InstFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstFamily

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner) SetInstFamily(v string)`

SetInstFamily sets InstFamily field to given value.

### HasInstFamily

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner) HasInstFamily() bool`

HasInstFamily returns a boolean if a field has been set.

### GetInstType

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetLever

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetMaxSz

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner) GetMaxSz() string`

GetMaxSz returns the MaxSz field if non-nil, zero value otherwise.

### GetMaxSzOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner) GetMaxSzOk() (*string, bool)`

GetMaxSzOk returns a tuple with the MaxSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSz

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner) SetMaxSz(v string)`

SetMaxSz sets MaxSz field to given value.

### HasMaxSz

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner) HasMaxSz() bool`

HasMaxSz returns a boolean if a field has been set.

### GetPos

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner) GetPos() string`

GetPos returns the Pos field if non-nil, zero value otherwise.

### GetPosOk

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner) GetPosOk() (*string, bool)`

GetPosOk returns a tuple with the Pos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPos

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner) SetPos(v string)`

SetPos sets Pos field to given value.

### HasPos

`func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerPosTierCheckInner) HasPos() bool`

HasPos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


