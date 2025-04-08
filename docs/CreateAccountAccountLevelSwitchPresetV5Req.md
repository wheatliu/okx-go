# CreateAccountAccountLevelSwitchPresetV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctLv** | **string** | Account mode   &#x60;2&#x60;: Spot and futures mode   &#x60;3&#x60;: Multi-currency margin code   &#x60;4&#x60;: Portfolio margin mode | [default to ""]
**Lever** | Pointer to **string** | Leverage   Required when switching from Portfolio margin mode to Spot and futures mode or Multi-currency margin mode, and the user holds cross-margin positions. | [optional] [default to ""]
**RiskOffsetType** | Pointer to **string** | Risk offset type   &#x60;1&#x60;: Spot-derivatives (USDT) risk offset   &#x60;2&#x60;: Spot-derivatives (Crypto) risk offset   &#x60;3&#x60;: Derivatives only mode   &#x60;4&#x60;: Spot-derivatives (USDC) risk offset   Applicable when switching from Spot and futures mode or Multi-currency margin mode to Portfolio margin mode.(Deprecated) | [optional] [default to ""]

## Methods

### NewCreateAccountAccountLevelSwitchPresetV5Req

`func NewCreateAccountAccountLevelSwitchPresetV5Req(acctLv string, ) *CreateAccountAccountLevelSwitchPresetV5Req`

NewCreateAccountAccountLevelSwitchPresetV5Req instantiates a new CreateAccountAccountLevelSwitchPresetV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountAccountLevelSwitchPresetV5ReqWithDefaults

`func NewCreateAccountAccountLevelSwitchPresetV5ReqWithDefaults() *CreateAccountAccountLevelSwitchPresetV5Req`

NewCreateAccountAccountLevelSwitchPresetV5ReqWithDefaults instantiates a new CreateAccountAccountLevelSwitchPresetV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctLv

`func (o *CreateAccountAccountLevelSwitchPresetV5Req) GetAcctLv() string`

GetAcctLv returns the AcctLv field if non-nil, zero value otherwise.

### GetAcctLvOk

`func (o *CreateAccountAccountLevelSwitchPresetV5Req) GetAcctLvOk() (*string, bool)`

GetAcctLvOk returns a tuple with the AcctLv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctLv

`func (o *CreateAccountAccountLevelSwitchPresetV5Req) SetAcctLv(v string)`

SetAcctLv sets AcctLv field to given value.


### GetLever

`func (o *CreateAccountAccountLevelSwitchPresetV5Req) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *CreateAccountAccountLevelSwitchPresetV5Req) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *CreateAccountAccountLevelSwitchPresetV5Req) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *CreateAccountAccountLevelSwitchPresetV5Req) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetRiskOffsetType

`func (o *CreateAccountAccountLevelSwitchPresetV5Req) GetRiskOffsetType() string`

GetRiskOffsetType returns the RiskOffsetType field if non-nil, zero value otherwise.

### GetRiskOffsetTypeOk

`func (o *CreateAccountAccountLevelSwitchPresetV5Req) GetRiskOffsetTypeOk() (*string, bool)`

GetRiskOffsetTypeOk returns a tuple with the RiskOffsetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskOffsetType

`func (o *CreateAccountAccountLevelSwitchPresetV5Req) SetRiskOffsetType(v string)`

SetRiskOffsetType sets RiskOffsetType field to given value.

### HasRiskOffsetType

`func (o *CreateAccountAccountLevelSwitchPresetV5Req) HasRiskOffsetType() bool`

HasRiskOffsetType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


