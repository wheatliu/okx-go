# GetAccountGreeksV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | Pointer to **string** | Currency | [optional] [default to ""]
**DeltaBS** | Pointer to **string** | delta: Black-Scholes Greeks in dollars | [optional] [default to ""]
**DeltaPA** | Pointer to **string** | delta: Greeks in coins | [optional] [default to ""]
**GammaBS** | Pointer to **string** | gamma: Black-Scholes Greeks in dollars, only applicable to OPTION | [optional] [default to ""]
**GammaPA** | Pointer to **string** | gamma: Greeks in coins, only applicable to OPTION | [optional] [default to ""]
**ThetaBS** | Pointer to **string** | theta: Black-Scholes Greeks in dollars, only applicable to &#x60;OPTION&#x60; | [optional] [default to ""]
**ThetaPA** | Pointer to **string** | theta: Greeks in coins, only applicable to &#x60;OPTION&#x60; | [optional] [default to ""]
**Ts** | Pointer to **string** | Time of getting Greeks, Unix timestamp format in milliseconds, e.g. 1597026383085 | [optional] [default to ""]
**VegaBS** | Pointer to **string** | vega: Black-Scholes Greeks in dollars, only applicable to &#x60;OPTION&#x60; | [optional] [default to ""]
**VegaPA** | Pointer to **string** | vega：Greeks in coins, only applicable to &#x60;OPTION&#x60; | [optional] [default to ""]

## Methods

### NewGetAccountGreeksV5RespDataInner

`func NewGetAccountGreeksV5RespDataInner() *GetAccountGreeksV5RespDataInner`

NewGetAccountGreeksV5RespDataInner instantiates a new GetAccountGreeksV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountGreeksV5RespDataInnerWithDefaults

`func NewGetAccountGreeksV5RespDataInnerWithDefaults() *GetAccountGreeksV5RespDataInner`

NewGetAccountGreeksV5RespDataInnerWithDefaults instantiates a new GetAccountGreeksV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *GetAccountGreeksV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAccountGreeksV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAccountGreeksV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAccountGreeksV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetDeltaBS

`func (o *GetAccountGreeksV5RespDataInner) GetDeltaBS() string`

GetDeltaBS returns the DeltaBS field if non-nil, zero value otherwise.

### GetDeltaBSOk

`func (o *GetAccountGreeksV5RespDataInner) GetDeltaBSOk() (*string, bool)`

GetDeltaBSOk returns a tuple with the DeltaBS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaBS

`func (o *GetAccountGreeksV5RespDataInner) SetDeltaBS(v string)`

SetDeltaBS sets DeltaBS field to given value.

### HasDeltaBS

`func (o *GetAccountGreeksV5RespDataInner) HasDeltaBS() bool`

HasDeltaBS returns a boolean if a field has been set.

### GetDeltaPA

`func (o *GetAccountGreeksV5RespDataInner) GetDeltaPA() string`

GetDeltaPA returns the DeltaPA field if non-nil, zero value otherwise.

### GetDeltaPAOk

`func (o *GetAccountGreeksV5RespDataInner) GetDeltaPAOk() (*string, bool)`

GetDeltaPAOk returns a tuple with the DeltaPA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaPA

`func (o *GetAccountGreeksV5RespDataInner) SetDeltaPA(v string)`

SetDeltaPA sets DeltaPA field to given value.

### HasDeltaPA

`func (o *GetAccountGreeksV5RespDataInner) HasDeltaPA() bool`

HasDeltaPA returns a boolean if a field has been set.

### GetGammaBS

`func (o *GetAccountGreeksV5RespDataInner) GetGammaBS() string`

GetGammaBS returns the GammaBS field if non-nil, zero value otherwise.

### GetGammaBSOk

`func (o *GetAccountGreeksV5RespDataInner) GetGammaBSOk() (*string, bool)`

GetGammaBSOk returns a tuple with the GammaBS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGammaBS

`func (o *GetAccountGreeksV5RespDataInner) SetGammaBS(v string)`

SetGammaBS sets GammaBS field to given value.

### HasGammaBS

`func (o *GetAccountGreeksV5RespDataInner) HasGammaBS() bool`

HasGammaBS returns a boolean if a field has been set.

### GetGammaPA

`func (o *GetAccountGreeksV5RespDataInner) GetGammaPA() string`

GetGammaPA returns the GammaPA field if non-nil, zero value otherwise.

### GetGammaPAOk

`func (o *GetAccountGreeksV5RespDataInner) GetGammaPAOk() (*string, bool)`

GetGammaPAOk returns a tuple with the GammaPA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGammaPA

`func (o *GetAccountGreeksV5RespDataInner) SetGammaPA(v string)`

SetGammaPA sets GammaPA field to given value.

### HasGammaPA

`func (o *GetAccountGreeksV5RespDataInner) HasGammaPA() bool`

HasGammaPA returns a boolean if a field has been set.

### GetThetaBS

`func (o *GetAccountGreeksV5RespDataInner) GetThetaBS() string`

GetThetaBS returns the ThetaBS field if non-nil, zero value otherwise.

### GetThetaBSOk

`func (o *GetAccountGreeksV5RespDataInner) GetThetaBSOk() (*string, bool)`

GetThetaBSOk returns a tuple with the ThetaBS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThetaBS

`func (o *GetAccountGreeksV5RespDataInner) SetThetaBS(v string)`

SetThetaBS sets ThetaBS field to given value.

### HasThetaBS

`func (o *GetAccountGreeksV5RespDataInner) HasThetaBS() bool`

HasThetaBS returns a boolean if a field has been set.

### GetThetaPA

`func (o *GetAccountGreeksV5RespDataInner) GetThetaPA() string`

GetThetaPA returns the ThetaPA field if non-nil, zero value otherwise.

### GetThetaPAOk

`func (o *GetAccountGreeksV5RespDataInner) GetThetaPAOk() (*string, bool)`

GetThetaPAOk returns a tuple with the ThetaPA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThetaPA

`func (o *GetAccountGreeksV5RespDataInner) SetThetaPA(v string)`

SetThetaPA sets ThetaPA field to given value.

### HasThetaPA

`func (o *GetAccountGreeksV5RespDataInner) HasThetaPA() bool`

HasThetaPA returns a boolean if a field has been set.

### GetTs

`func (o *GetAccountGreeksV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetAccountGreeksV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetAccountGreeksV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetAccountGreeksV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetVegaBS

`func (o *GetAccountGreeksV5RespDataInner) GetVegaBS() string`

GetVegaBS returns the VegaBS field if non-nil, zero value otherwise.

### GetVegaBSOk

`func (o *GetAccountGreeksV5RespDataInner) GetVegaBSOk() (*string, bool)`

GetVegaBSOk returns a tuple with the VegaBS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVegaBS

`func (o *GetAccountGreeksV5RespDataInner) SetVegaBS(v string)`

SetVegaBS sets VegaBS field to given value.

### HasVegaBS

`func (o *GetAccountGreeksV5RespDataInner) HasVegaBS() bool`

HasVegaBS returns a boolean if a field has been set.

### GetVegaPA

`func (o *GetAccountGreeksV5RespDataInner) GetVegaPA() string`

GetVegaPA returns the VegaPA field if non-nil, zero value otherwise.

### GetVegaPAOk

`func (o *GetAccountGreeksV5RespDataInner) GetVegaPAOk() (*string, bool)`

GetVegaPAOk returns a tuple with the VegaPA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVegaPA

`func (o *GetAccountGreeksV5RespDataInner) SetVegaPA(v string)`

SetVegaPA sets VegaPA field to given value.

### HasVegaPA

`func (o *GetAccountGreeksV5RespDataInner) HasVegaPA() bool`

HasVegaPA returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


