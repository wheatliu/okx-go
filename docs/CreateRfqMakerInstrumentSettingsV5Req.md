# CreateRfqMakerInstrumentSettingsV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CreateRfqMakerInstrumentSettingsV5ReqDataInner**](CreateRfqMakerInstrumentSettingsV5ReqDataInner.md) | Elements of the instType. | 
**IncludeAll** | Pointer to **bool** | Receive all instruments or not under specific instType setting.   Valid value can be boolean (&#x60;True&#x60;/&#x60;False&#x60;). By default, the value will be &#x60;false&#x60;. | [optional] 
**InstType** | **string** | Type of instrument. Valid value can be &#x60;FUTURES&#x60;, &#x60;OPTION&#x60;, &#x60;SWAP&#x60; or &#x60;SPOT&#x60;. | [default to ""]

## Methods

### NewCreateRfqMakerInstrumentSettingsV5Req

`func NewCreateRfqMakerInstrumentSettingsV5Req(data []CreateRfqMakerInstrumentSettingsV5ReqDataInner, instType string, ) *CreateRfqMakerInstrumentSettingsV5Req`

NewCreateRfqMakerInstrumentSettingsV5Req instantiates a new CreateRfqMakerInstrumentSettingsV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqMakerInstrumentSettingsV5ReqWithDefaults

`func NewCreateRfqMakerInstrumentSettingsV5ReqWithDefaults() *CreateRfqMakerInstrumentSettingsV5Req`

NewCreateRfqMakerInstrumentSettingsV5ReqWithDefaults instantiates a new CreateRfqMakerInstrumentSettingsV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateRfqMakerInstrumentSettingsV5Req) GetData() []CreateRfqMakerInstrumentSettingsV5ReqDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateRfqMakerInstrumentSettingsV5Req) GetDataOk() (*[]CreateRfqMakerInstrumentSettingsV5ReqDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateRfqMakerInstrumentSettingsV5Req) SetData(v []CreateRfqMakerInstrumentSettingsV5ReqDataInner)`

SetData sets Data field to given value.


### GetIncludeAll

`func (o *CreateRfqMakerInstrumentSettingsV5Req) GetIncludeAll() bool`

GetIncludeAll returns the IncludeAll field if non-nil, zero value otherwise.

### GetIncludeAllOk

`func (o *CreateRfqMakerInstrumentSettingsV5Req) GetIncludeAllOk() (*bool, bool)`

GetIncludeAllOk returns a tuple with the IncludeAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAll

`func (o *CreateRfqMakerInstrumentSettingsV5Req) SetIncludeAll(v bool)`

SetIncludeAll sets IncludeAll field to given value.

### HasIncludeAll

`func (o *CreateRfqMakerInstrumentSettingsV5Req) HasIncludeAll() bool`

HasIncludeAll returns a boolean if a field has been set.

### GetInstType

`func (o *CreateRfqMakerInstrumentSettingsV5Req) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *CreateRfqMakerInstrumentSettingsV5Req) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *CreateRfqMakerInstrumentSettingsV5Req) SetInstType(v string)`

SetInstType sets InstType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


