# GetCopytradingInstrumentsV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether instrument is a lead instrument. &#x60;true&#x60; or &#x60;false&#x60; | [optional] 
**InstId** | Pointer to **string** | Instrument ID, e.g. BTC-USDT-SWAP | [optional] [default to ""]

## Methods

### NewGetCopytradingInstrumentsV5RespDataInner

`func NewGetCopytradingInstrumentsV5RespDataInner() *GetCopytradingInstrumentsV5RespDataInner`

NewGetCopytradingInstrumentsV5RespDataInner instantiates a new GetCopytradingInstrumentsV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCopytradingInstrumentsV5RespDataInnerWithDefaults

`func NewGetCopytradingInstrumentsV5RespDataInnerWithDefaults() *GetCopytradingInstrumentsV5RespDataInner`

NewGetCopytradingInstrumentsV5RespDataInnerWithDefaults instantiates a new GetCopytradingInstrumentsV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetCopytradingInstrumentsV5RespDataInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetCopytradingInstrumentsV5RespDataInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetCopytradingInstrumentsV5RespDataInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetCopytradingInstrumentsV5RespDataInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInstId

`func (o *GetCopytradingInstrumentsV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetCopytradingInstrumentsV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetCopytradingInstrumentsV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetCopytradingInstrumentsV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


