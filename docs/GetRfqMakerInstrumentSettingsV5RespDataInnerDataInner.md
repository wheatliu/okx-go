# GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CreateRfqMakerInstrumentSettingsV5ReqDataInner**](CreateRfqMakerInstrumentSettingsV5ReqDataInner.md) | Elements of the instType. | [optional] 
**IncludeAll** | Pointer to **bool** | Receive all instruments or not under specific instType setting.   Valid value can be boolean (&#x60;True&#x60;/&#x60;False&#x60;). By default, the value will be &#x60;false&#x60;. | [optional] 
**InstFamily** | Pointer to **string** | Instrument family. Required for &#x60;FUTURES&#x60;, &#x60;OPTION&#x60; and &#x60;SWAP&#x60; only. | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID. Required for &#x60;SPOT&#x60; only. | [optional] [default to ""]
**InstType** | Pointer to **string** | Type of instrument. Valid value can be &#x60;FUTURES&#x60;, &#x60;OPTION&#x60;, &#x60;SWAP&#x60; or &#x60;SPOT&#x60;. | [optional] [default to ""]
**MakerPxBand** | Pointer to **string** | Price bands in unit of ticks, measured against mark price.   Setting makerPxBand to 1 tick means:   If Bid price &gt; Mark + 1 tick, it will be stopped   If Ask price &lt; Mark - 1 tick, It will be stopped | [optional] [default to ""]
**MaxBlockSz** | Pointer to **string** | Max trade quantity for the product(s).   For &#x60;FUTURES&#x60;, &#x60;OPTION&#x60; and &#x60;SWAP&#x60;, the max quantity of the RFQ/Quote is in unit of contracts. For &#x60;SPOT&#x60;, this parameter is in base currency. | [optional] [default to ""]

## Methods

### NewGetRfqMakerInstrumentSettingsV5RespDataInnerDataInner

`func NewGetRfqMakerInstrumentSettingsV5RespDataInnerDataInner() *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner`

NewGetRfqMakerInstrumentSettingsV5RespDataInnerDataInner instantiates a new GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRfqMakerInstrumentSettingsV5RespDataInnerDataInnerWithDefaults

`func NewGetRfqMakerInstrumentSettingsV5RespDataInnerDataInnerWithDefaults() *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner`

NewGetRfqMakerInstrumentSettingsV5RespDataInnerDataInnerWithDefaults instantiates a new GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) GetData() []CreateRfqMakerInstrumentSettingsV5ReqDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) GetDataOk() (*[]CreateRfqMakerInstrumentSettingsV5ReqDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) SetData(v []CreateRfqMakerInstrumentSettingsV5ReqDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncludeAll

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) GetIncludeAll() bool`

GetIncludeAll returns the IncludeAll field if non-nil, zero value otherwise.

### GetIncludeAllOk

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) GetIncludeAllOk() (*bool, bool)`

GetIncludeAllOk returns a tuple with the IncludeAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAll

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) SetIncludeAll(v bool)`

SetIncludeAll sets IncludeAll field to given value.

### HasIncludeAll

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) HasIncludeAll() bool`

HasIncludeAll returns a boolean if a field has been set.

### GetInstFamily

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) GetInstFamily() string`

GetInstFamily returns the InstFamily field if non-nil, zero value otherwise.

### GetInstFamilyOk

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) GetInstFamilyOk() (*string, bool)`

GetInstFamilyOk returns a tuple with the InstFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstFamily

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) SetInstFamily(v string)`

SetInstFamily sets InstFamily field to given value.

### HasInstFamily

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) HasInstFamily() bool`

HasInstFamily returns a boolean if a field has been set.

### GetInstId

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetMakerPxBand

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) GetMakerPxBand() string`

GetMakerPxBand returns the MakerPxBand field if non-nil, zero value otherwise.

### GetMakerPxBandOk

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) GetMakerPxBandOk() (*string, bool)`

GetMakerPxBandOk returns a tuple with the MakerPxBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerPxBand

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) SetMakerPxBand(v string)`

SetMakerPxBand sets MakerPxBand field to given value.

### HasMakerPxBand

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) HasMakerPxBand() bool`

HasMakerPxBand returns a boolean if a field has been set.

### GetMaxBlockSz

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) GetMaxBlockSz() string`

GetMaxBlockSz returns the MaxBlockSz field if non-nil, zero value otherwise.

### GetMaxBlockSzOk

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) GetMaxBlockSzOk() (*string, bool)`

GetMaxBlockSzOk returns a tuple with the MaxBlockSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBlockSz

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) SetMaxBlockSz(v string)`

SetMaxBlockSz sets MaxBlockSz field to given value.

### HasMaxBlockSz

`func (o *GetRfqMakerInstrumentSettingsV5RespDataInnerDataInner) HasMaxBlockSz() bool`

HasMaxBlockSz returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


