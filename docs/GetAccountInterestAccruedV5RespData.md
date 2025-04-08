# GetAccountInterestAccruedV5RespData

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

### NewGetAccountInterestAccruedV5RespData

`func NewGetAccountInterestAccruedV5RespData() *GetAccountInterestAccruedV5RespData`

NewGetAccountInterestAccruedV5RespData instantiates a new GetAccountInterestAccruedV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountInterestAccruedV5RespDataWithDefaults

`func NewGetAccountInterestAccruedV5RespDataWithDefaults() *GetAccountInterestAccruedV5RespData`

NewGetAccountInterestAccruedV5RespDataWithDefaults instantiates a new GetAccountInterestAccruedV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *GetAccountInterestAccruedV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAccountInterestAccruedV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAccountInterestAccruedV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAccountInterestAccruedV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetInstId

`func (o *GetAccountInterestAccruedV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetAccountInterestAccruedV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetAccountInterestAccruedV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetAccountInterestAccruedV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInterest

`func (o *GetAccountInterestAccruedV5RespData) GetInterest() string`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *GetAccountInterestAccruedV5RespData) GetInterestOk() (*string, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *GetAccountInterestAccruedV5RespData) SetInterest(v string)`

SetInterest sets Interest field to given value.

### HasInterest

`func (o *GetAccountInterestAccruedV5RespData) HasInterest() bool`

HasInterest returns a boolean if a field has been set.

### GetInterestRate

`func (o *GetAccountInterestAccruedV5RespData) GetInterestRate() string`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *GetAccountInterestAccruedV5RespData) GetInterestRateOk() (*string, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *GetAccountInterestAccruedV5RespData) SetInterestRate(v string)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *GetAccountInterestAccruedV5RespData) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetLiab

`func (o *GetAccountInterestAccruedV5RespData) GetLiab() string`

GetLiab returns the Liab field if non-nil, zero value otherwise.

### GetLiabOk

`func (o *GetAccountInterestAccruedV5RespData) GetLiabOk() (*string, bool)`

GetLiabOk returns a tuple with the Liab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiab

`func (o *GetAccountInterestAccruedV5RespData) SetLiab(v string)`

SetLiab sets Liab field to given value.

### HasLiab

`func (o *GetAccountInterestAccruedV5RespData) HasLiab() bool`

HasLiab returns a boolean if a field has been set.

### GetMgnMode

`func (o *GetAccountInterestAccruedV5RespData) GetMgnMode() string`

GetMgnMode returns the MgnMode field if non-nil, zero value otherwise.

### GetMgnModeOk

`func (o *GetAccountInterestAccruedV5RespData) GetMgnModeOk() (*string, bool)`

GetMgnModeOk returns a tuple with the MgnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnMode

`func (o *GetAccountInterestAccruedV5RespData) SetMgnMode(v string)`

SetMgnMode sets MgnMode field to given value.

### HasMgnMode

`func (o *GetAccountInterestAccruedV5RespData) HasMgnMode() bool`

HasMgnMode returns a boolean if a field has been set.

### GetTs

`func (o *GetAccountInterestAccruedV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetAccountInterestAccruedV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetAccountInterestAccruedV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetAccountInterestAccruedV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetType

`func (o *GetAccountInterestAccruedV5RespData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAccountInterestAccruedV5RespData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAccountInterestAccruedV5RespData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAccountInterestAccruedV5RespData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


