# GetAccountTradeFeeV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Currency category. Note: this parameter is already deprecated | [optional] [default to ""]
**Delivery** | Pointer to **string** | Delivery fee rate | [optional] [default to ""]
**Exercise** | Pointer to **string** | Fee rate for exercising the option | [optional] [default to ""]
**Fiat** | Pointer to [**[]GetAccountTradeFeeV5RespDataInnerFiatInner**](GetAccountTradeFeeV5RespDataInnerFiatInner.md) | Details of fiat fee rate | [optional] 
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**Level** | Pointer to **string** | Fee rate Level | [optional] [default to ""]
**Maker** | Pointer to **string** | For &#x60;SPOT&#x60;/&#x60;MARGIN&#x60;, it is maker fee rate of the USDT trading pairs.   For &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;, it is the fee rate of crypto-margined contracts | [optional] [default to ""]
**MakerU** | Pointer to **string** | Maker fee rate of USDT-margined contracts, only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60; | [optional] [default to ""]
**MakerUSDC** | Pointer to **string** | For &#x60;SPOT&#x60;/&#x60;MARGIN&#x60;, it is maker fee rate of the USDⓈ&amp;Crypto trading pairs.  For &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;, it is the fee rate of USDC-margined contracts | [optional] [default to ""]
**RuleType** | Pointer to **string** | Trading rule types   &#x60;normal&#x60;: normal trading   &#x60;pre_market&#x60;: pre-market trading | [optional] [default to ""]
**Taker** | Pointer to **string** | For &#x60;SPOT&#x60;/&#x60;MARGIN&#x60;, it is taker fee rate of the USDT trading pairs.   For &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;, it is the fee rate of crypto-margined contracts | [optional] [default to ""]
**TakerU** | Pointer to **string** | Taker fee rate of USDT-margined contracts, only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60; | [optional] [default to ""]
**TakerUSDC** | Pointer to **string** | For &#x60;SPOT&#x60;/&#x60;MARGIN&#x60;, it is taker fee rate of the USDⓈ&amp;Crypto trading pairs.  For &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;, it is the fee rate of USDC-margined contracts | [optional] [default to ""]
**Ts** | Pointer to **string** | Data return time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetAccountTradeFeeV5RespDataInner

`func NewGetAccountTradeFeeV5RespDataInner() *GetAccountTradeFeeV5RespDataInner`

NewGetAccountTradeFeeV5RespDataInner instantiates a new GetAccountTradeFeeV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountTradeFeeV5RespDataInnerWithDefaults

`func NewGetAccountTradeFeeV5RespDataInnerWithDefaults() *GetAccountTradeFeeV5RespDataInner`

NewGetAccountTradeFeeV5RespDataInnerWithDefaults instantiates a new GetAccountTradeFeeV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *GetAccountTradeFeeV5RespDataInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetAccountTradeFeeV5RespDataInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetAccountTradeFeeV5RespDataInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetAccountTradeFeeV5RespDataInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDelivery

`func (o *GetAccountTradeFeeV5RespDataInner) GetDelivery() string`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *GetAccountTradeFeeV5RespDataInner) GetDeliveryOk() (*string, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *GetAccountTradeFeeV5RespDataInner) SetDelivery(v string)`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *GetAccountTradeFeeV5RespDataInner) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.

### GetExercise

`func (o *GetAccountTradeFeeV5RespDataInner) GetExercise() string`

GetExercise returns the Exercise field if non-nil, zero value otherwise.

### GetExerciseOk

`func (o *GetAccountTradeFeeV5RespDataInner) GetExerciseOk() (*string, bool)`

GetExerciseOk returns a tuple with the Exercise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExercise

`func (o *GetAccountTradeFeeV5RespDataInner) SetExercise(v string)`

SetExercise sets Exercise field to given value.

### HasExercise

`func (o *GetAccountTradeFeeV5RespDataInner) HasExercise() bool`

HasExercise returns a boolean if a field has been set.

### GetFiat

`func (o *GetAccountTradeFeeV5RespDataInner) GetFiat() []GetAccountTradeFeeV5RespDataInnerFiatInner`

GetFiat returns the Fiat field if non-nil, zero value otherwise.

### GetFiatOk

`func (o *GetAccountTradeFeeV5RespDataInner) GetFiatOk() (*[]GetAccountTradeFeeV5RespDataInnerFiatInner, bool)`

GetFiatOk returns a tuple with the Fiat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiat

`func (o *GetAccountTradeFeeV5RespDataInner) SetFiat(v []GetAccountTradeFeeV5RespDataInnerFiatInner)`

SetFiat sets Fiat field to given value.

### HasFiat

`func (o *GetAccountTradeFeeV5RespDataInner) HasFiat() bool`

HasFiat returns a boolean if a field has been set.

### GetInstType

`func (o *GetAccountTradeFeeV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetAccountTradeFeeV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetAccountTradeFeeV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetAccountTradeFeeV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetLevel

`func (o *GetAccountTradeFeeV5RespDataInner) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *GetAccountTradeFeeV5RespDataInner) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *GetAccountTradeFeeV5RespDataInner) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *GetAccountTradeFeeV5RespDataInner) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetMaker

`func (o *GetAccountTradeFeeV5RespDataInner) GetMaker() string`

GetMaker returns the Maker field if non-nil, zero value otherwise.

### GetMakerOk

`func (o *GetAccountTradeFeeV5RespDataInner) GetMakerOk() (*string, bool)`

GetMakerOk returns a tuple with the Maker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaker

`func (o *GetAccountTradeFeeV5RespDataInner) SetMaker(v string)`

SetMaker sets Maker field to given value.

### HasMaker

`func (o *GetAccountTradeFeeV5RespDataInner) HasMaker() bool`

HasMaker returns a boolean if a field has been set.

### GetMakerU

`func (o *GetAccountTradeFeeV5RespDataInner) GetMakerU() string`

GetMakerU returns the MakerU field if non-nil, zero value otherwise.

### GetMakerUOk

`func (o *GetAccountTradeFeeV5RespDataInner) GetMakerUOk() (*string, bool)`

GetMakerUOk returns a tuple with the MakerU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerU

`func (o *GetAccountTradeFeeV5RespDataInner) SetMakerU(v string)`

SetMakerU sets MakerU field to given value.

### HasMakerU

`func (o *GetAccountTradeFeeV5RespDataInner) HasMakerU() bool`

HasMakerU returns a boolean if a field has been set.

### GetMakerUSDC

`func (o *GetAccountTradeFeeV5RespDataInner) GetMakerUSDC() string`

GetMakerUSDC returns the MakerUSDC field if non-nil, zero value otherwise.

### GetMakerUSDCOk

`func (o *GetAccountTradeFeeV5RespDataInner) GetMakerUSDCOk() (*string, bool)`

GetMakerUSDCOk returns a tuple with the MakerUSDC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerUSDC

`func (o *GetAccountTradeFeeV5RespDataInner) SetMakerUSDC(v string)`

SetMakerUSDC sets MakerUSDC field to given value.

### HasMakerUSDC

`func (o *GetAccountTradeFeeV5RespDataInner) HasMakerUSDC() bool`

HasMakerUSDC returns a boolean if a field has been set.

### GetRuleType

`func (o *GetAccountTradeFeeV5RespDataInner) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *GetAccountTradeFeeV5RespDataInner) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *GetAccountTradeFeeV5RespDataInner) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *GetAccountTradeFeeV5RespDataInner) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetTaker

`func (o *GetAccountTradeFeeV5RespDataInner) GetTaker() string`

GetTaker returns the Taker field if non-nil, zero value otherwise.

### GetTakerOk

`func (o *GetAccountTradeFeeV5RespDataInner) GetTakerOk() (*string, bool)`

GetTakerOk returns a tuple with the Taker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaker

`func (o *GetAccountTradeFeeV5RespDataInner) SetTaker(v string)`

SetTaker sets Taker field to given value.

### HasTaker

`func (o *GetAccountTradeFeeV5RespDataInner) HasTaker() bool`

HasTaker returns a boolean if a field has been set.

### GetTakerU

`func (o *GetAccountTradeFeeV5RespDataInner) GetTakerU() string`

GetTakerU returns the TakerU field if non-nil, zero value otherwise.

### GetTakerUOk

`func (o *GetAccountTradeFeeV5RespDataInner) GetTakerUOk() (*string, bool)`

GetTakerUOk returns a tuple with the TakerU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerU

`func (o *GetAccountTradeFeeV5RespDataInner) SetTakerU(v string)`

SetTakerU sets TakerU field to given value.

### HasTakerU

`func (o *GetAccountTradeFeeV5RespDataInner) HasTakerU() bool`

HasTakerU returns a boolean if a field has been set.

### GetTakerUSDC

`func (o *GetAccountTradeFeeV5RespDataInner) GetTakerUSDC() string`

GetTakerUSDC returns the TakerUSDC field if non-nil, zero value otherwise.

### GetTakerUSDCOk

`func (o *GetAccountTradeFeeV5RespDataInner) GetTakerUSDCOk() (*string, bool)`

GetTakerUSDCOk returns a tuple with the TakerUSDC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerUSDC

`func (o *GetAccountTradeFeeV5RespDataInner) SetTakerUSDC(v string)`

SetTakerUSDC sets TakerUSDC field to given value.

### HasTakerUSDC

`func (o *GetAccountTradeFeeV5RespDataInner) HasTakerUSDC() bool`

HasTakerUSDC returns a boolean if a field has been set.

### GetTs

`func (o *GetAccountTradeFeeV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetAccountTradeFeeV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetAccountTradeFeeV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetAccountTradeFeeV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


