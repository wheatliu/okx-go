# CreateRfqMakerInstrumentSettingsV5ReqDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstFamily** | Pointer to **string** | Instrument family. Required for &#x60;FUTURES&#x60;, &#x60;OPTION&#x60; and &#x60;SWAP&#x60; only. | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID. Required for &#x60;SPOT&#x60; only. | [optional] [default to ""]
**MakerPxBand** | Pointer to **string** | Price bands in unit of ticks, measured against mark price.   Setting makerPxBand to 1 tick means:   If Bid price &gt; Mark + 1 tick, it will be stopped   If Ask price &lt; Mark - 1 tick, It will be stopped | [optional] [default to ""]
**MaxBlockSz** | Pointer to **string** | Max trade quantity for the product(s).   For &#x60;FUTURES&#x60;, &#x60;OPTION&#x60; and &#x60;SWAP&#x60;, the max quantity of the RFQ/Quote is in unit of contracts. For &#x60;SPOT&#x60;, this parameter is in base currency. | [optional] [default to ""]

## Methods

### NewCreateRfqMakerInstrumentSettingsV5ReqDataInner

`func NewCreateRfqMakerInstrumentSettingsV5ReqDataInner() *CreateRfqMakerInstrumentSettingsV5ReqDataInner`

NewCreateRfqMakerInstrumentSettingsV5ReqDataInner instantiates a new CreateRfqMakerInstrumentSettingsV5ReqDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqMakerInstrumentSettingsV5ReqDataInnerWithDefaults

`func NewCreateRfqMakerInstrumentSettingsV5ReqDataInnerWithDefaults() *CreateRfqMakerInstrumentSettingsV5ReqDataInner`

NewCreateRfqMakerInstrumentSettingsV5ReqDataInnerWithDefaults instantiates a new CreateRfqMakerInstrumentSettingsV5ReqDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstFamily

`func (o *CreateRfqMakerInstrumentSettingsV5ReqDataInner) GetInstFamily() string`

GetInstFamily returns the InstFamily field if non-nil, zero value otherwise.

### GetInstFamilyOk

`func (o *CreateRfqMakerInstrumentSettingsV5ReqDataInner) GetInstFamilyOk() (*string, bool)`

GetInstFamilyOk returns a tuple with the InstFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstFamily

`func (o *CreateRfqMakerInstrumentSettingsV5ReqDataInner) SetInstFamily(v string)`

SetInstFamily sets InstFamily field to given value.

### HasInstFamily

`func (o *CreateRfqMakerInstrumentSettingsV5ReqDataInner) HasInstFamily() bool`

HasInstFamily returns a boolean if a field has been set.

### GetInstId

`func (o *CreateRfqMakerInstrumentSettingsV5ReqDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateRfqMakerInstrumentSettingsV5ReqDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateRfqMakerInstrumentSettingsV5ReqDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *CreateRfqMakerInstrumentSettingsV5ReqDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetMakerPxBand

`func (o *CreateRfqMakerInstrumentSettingsV5ReqDataInner) GetMakerPxBand() string`

GetMakerPxBand returns the MakerPxBand field if non-nil, zero value otherwise.

### GetMakerPxBandOk

`func (o *CreateRfqMakerInstrumentSettingsV5ReqDataInner) GetMakerPxBandOk() (*string, bool)`

GetMakerPxBandOk returns a tuple with the MakerPxBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerPxBand

`func (o *CreateRfqMakerInstrumentSettingsV5ReqDataInner) SetMakerPxBand(v string)`

SetMakerPxBand sets MakerPxBand field to given value.

### HasMakerPxBand

`func (o *CreateRfqMakerInstrumentSettingsV5ReqDataInner) HasMakerPxBand() bool`

HasMakerPxBand returns a boolean if a field has been set.

### GetMaxBlockSz

`func (o *CreateRfqMakerInstrumentSettingsV5ReqDataInner) GetMaxBlockSz() string`

GetMaxBlockSz returns the MaxBlockSz field if non-nil, zero value otherwise.

### GetMaxBlockSzOk

`func (o *CreateRfqMakerInstrumentSettingsV5ReqDataInner) GetMaxBlockSzOk() (*string, bool)`

GetMaxBlockSzOk returns a tuple with the MaxBlockSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBlockSz

`func (o *CreateRfqMakerInstrumentSettingsV5ReqDataInner) SetMaxBlockSz(v string)`

SetMaxBlockSz sets MaxBlockSz field to given value.

### HasMaxBlockSz

`func (o *CreateRfqMakerInstrumentSettingsV5ReqDataInner) HasMaxBlockSz() bool`

HasMaxBlockSz returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


