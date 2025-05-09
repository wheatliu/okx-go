# GetPublicInsuranceFundV5RespDataInnerDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdlType** | Pointer to **string** | ADL related events   &#x60;rate_adl_start&#x60;: ADL begins due to high insurance fund decline rate   &#x60;bal_adl_start&#x60;: ADL begins due to insurance fund balance falling   &#x60;pos_adl_start&#x60;：ADL begins due to the volume of liquidation orders falls to a certain level (only applicable to premarket symbols)   &#x60;adl_end&#x60;: ADL ends   Only applicable when type is &#x60;adl&#x60; | [optional] [default to ""]
**Amt** | Pointer to **string** | The change in the balance of insurance fund   Applicable when type is &#x60;liquidation_balance_deposit&#x60;, &#x60;bankruptcy_loss&#x60; or &#x60;platform_revenue&#x60; | [optional] [default to ""]
**Balance** | Pointer to **string** | The balance of insurance fund | [optional] [default to ""]
**Ccy** | Pointer to **string** | The currency of insurance fund | [optional] [default to ""]
**DecRate** | Pointer to **string** | Real-time insurance fund decline rate (compare balance and maxBal)   Only applicable when type is &#x60;adl&#x60;(Deprecated) | [optional] [default to ""]
**MaxBal** | Pointer to **string** | Maximum insurance fund balance in the past eight hours   Only applicable when type is &#x60;adl&#x60; | [optional] [default to ""]
**MaxBalTs** | Pointer to **string** | Timestamp when insurance fund balance reached maximum in the past eight hours, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60;   Only applicable when type is &#x60;adl&#x60; | [optional] [default to ""]
**Ts** | Pointer to **string** | The update timestamp of insurance fund. Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Type** | Pointer to **string** | The type of insurance fund | [optional] [default to ""]

## Methods

### NewGetPublicInsuranceFundV5RespDataInnerDetailsInner

`func NewGetPublicInsuranceFundV5RespDataInnerDetailsInner() *GetPublicInsuranceFundV5RespDataInnerDetailsInner`

NewGetPublicInsuranceFundV5RespDataInnerDetailsInner instantiates a new GetPublicInsuranceFundV5RespDataInnerDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicInsuranceFundV5RespDataInnerDetailsInnerWithDefaults

`func NewGetPublicInsuranceFundV5RespDataInnerDetailsInnerWithDefaults() *GetPublicInsuranceFundV5RespDataInnerDetailsInner`

NewGetPublicInsuranceFundV5RespDataInnerDetailsInnerWithDefaults instantiates a new GetPublicInsuranceFundV5RespDataInnerDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdlType

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) GetAdlType() string`

GetAdlType returns the AdlType field if non-nil, zero value otherwise.

### GetAdlTypeOk

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) GetAdlTypeOk() (*string, bool)`

GetAdlTypeOk returns a tuple with the AdlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdlType

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) SetAdlType(v string)`

SetAdlType sets AdlType field to given value.

### HasAdlType

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) HasAdlType() bool`

HasAdlType returns a boolean if a field has been set.

### GetAmt

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetBalance

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) SetBalance(v string)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCcy

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetDecRate

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) GetDecRate() string`

GetDecRate returns the DecRate field if non-nil, zero value otherwise.

### GetDecRateOk

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) GetDecRateOk() (*string, bool)`

GetDecRateOk returns a tuple with the DecRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecRate

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) SetDecRate(v string)`

SetDecRate sets DecRate field to given value.

### HasDecRate

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) HasDecRate() bool`

HasDecRate returns a boolean if a field has been set.

### GetMaxBal

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) GetMaxBal() string`

GetMaxBal returns the MaxBal field if non-nil, zero value otherwise.

### GetMaxBalOk

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) GetMaxBalOk() (*string, bool)`

GetMaxBalOk returns a tuple with the MaxBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBal

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) SetMaxBal(v string)`

SetMaxBal sets MaxBal field to given value.

### HasMaxBal

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) HasMaxBal() bool`

HasMaxBal returns a boolean if a field has been set.

### GetMaxBalTs

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) GetMaxBalTs() string`

GetMaxBalTs returns the MaxBalTs field if non-nil, zero value otherwise.

### GetMaxBalTsOk

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) GetMaxBalTsOk() (*string, bool)`

GetMaxBalTsOk returns a tuple with the MaxBalTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBalTs

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) SetMaxBalTs(v string)`

SetMaxBalTs sets MaxBalTs field to given value.

### HasMaxBalTs

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) HasMaxBalTs() bool`

HasMaxBalTs returns a boolean if a field has been set.

### GetTs

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetType

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetPublicInsuranceFundV5RespDataInnerDetailsInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


