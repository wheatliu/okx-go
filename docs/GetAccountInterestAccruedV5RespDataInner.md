# GetAccountInterestAccruedV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | Pointer to **string** | Loan currency, e.g. &#x60;BTC&#x60; | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60;  Only applicable to &#x60;Market loans&#x60; | [optional] [default to ""]
**Interest** | Pointer to **string** | Interest | [optional] [default to ""]
**InterestRate** | Pointer to **string** | Interest rate (in hour) | [optional] [default to ""]
**Liab** | Pointer to **string** | Liability | [optional] [default to ""]
**MgnMode** | Pointer to **string** | Margin mode  &#x60;cross&#x60;    &#x60;isolated&#x60; | [optional] [default to ""]
**Ts** | Pointer to **string** | Timestamp for interest accured, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Type** | Pointer to **string** | Loan type  &#x60;2&#x60;: Market loans | [optional] [default to ""]

## Methods

### NewGetAccountInterestAccruedV5RespDataInner

`func NewGetAccountInterestAccruedV5RespDataInner() *GetAccountInterestAccruedV5RespDataInner`

NewGetAccountInterestAccruedV5RespDataInner instantiates a new GetAccountInterestAccruedV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountInterestAccruedV5RespDataInnerWithDefaults

`func NewGetAccountInterestAccruedV5RespDataInnerWithDefaults() *GetAccountInterestAccruedV5RespDataInner`

NewGetAccountInterestAccruedV5RespDataInnerWithDefaults instantiates a new GetAccountInterestAccruedV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *GetAccountInterestAccruedV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAccountInterestAccruedV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAccountInterestAccruedV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAccountInterestAccruedV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetInstId

`func (o *GetAccountInterestAccruedV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetAccountInterestAccruedV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetAccountInterestAccruedV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetAccountInterestAccruedV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInterest

`func (o *GetAccountInterestAccruedV5RespDataInner) GetInterest() string`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *GetAccountInterestAccruedV5RespDataInner) GetInterestOk() (*string, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *GetAccountInterestAccruedV5RespDataInner) SetInterest(v string)`

SetInterest sets Interest field to given value.

### HasInterest

`func (o *GetAccountInterestAccruedV5RespDataInner) HasInterest() bool`

HasInterest returns a boolean if a field has been set.

### GetInterestRate

`func (o *GetAccountInterestAccruedV5RespDataInner) GetInterestRate() string`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *GetAccountInterestAccruedV5RespDataInner) GetInterestRateOk() (*string, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *GetAccountInterestAccruedV5RespDataInner) SetInterestRate(v string)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *GetAccountInterestAccruedV5RespDataInner) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetLiab

`func (o *GetAccountInterestAccruedV5RespDataInner) GetLiab() string`

GetLiab returns the Liab field if non-nil, zero value otherwise.

### GetLiabOk

`func (o *GetAccountInterestAccruedV5RespDataInner) GetLiabOk() (*string, bool)`

GetLiabOk returns a tuple with the Liab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiab

`func (o *GetAccountInterestAccruedV5RespDataInner) SetLiab(v string)`

SetLiab sets Liab field to given value.

### HasLiab

`func (o *GetAccountInterestAccruedV5RespDataInner) HasLiab() bool`

HasLiab returns a boolean if a field has been set.

### GetMgnMode

`func (o *GetAccountInterestAccruedV5RespDataInner) GetMgnMode() string`

GetMgnMode returns the MgnMode field if non-nil, zero value otherwise.

### GetMgnModeOk

`func (o *GetAccountInterestAccruedV5RespDataInner) GetMgnModeOk() (*string, bool)`

GetMgnModeOk returns a tuple with the MgnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnMode

`func (o *GetAccountInterestAccruedV5RespDataInner) SetMgnMode(v string)`

SetMgnMode sets MgnMode field to given value.

### HasMgnMode

`func (o *GetAccountInterestAccruedV5RespDataInner) HasMgnMode() bool`

HasMgnMode returns a boolean if a field has been set.

### GetTs

`func (o *GetAccountInterestAccruedV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetAccountInterestAccruedV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetAccountInterestAccruedV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetAccountInterestAccruedV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetType

`func (o *GetAccountInterestAccruedV5RespDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAccountInterestAccruedV5RespDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAccountInterestAccruedV5RespDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAccountInterestAccruedV5RespDataInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


