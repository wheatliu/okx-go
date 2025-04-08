# CreateCopytradingSetInstrumentsV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstId** | **string** | Instrument ID, e.g. BTC-USDT-SWAP. If there are multiple instruments, separate them with commas. | [default to ""]
**InstType** | Pointer to **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;SWAP&#x60;, the default value | [optional] [default to ""]

## Methods

### NewCreateCopytradingSetInstrumentsV5Req

`func NewCreateCopytradingSetInstrumentsV5Req(instId string, ) *CreateCopytradingSetInstrumentsV5Req`

NewCreateCopytradingSetInstrumentsV5Req instantiates a new CreateCopytradingSetInstrumentsV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCopytradingSetInstrumentsV5ReqWithDefaults

`func NewCreateCopytradingSetInstrumentsV5ReqWithDefaults() *CreateCopytradingSetInstrumentsV5Req`

NewCreateCopytradingSetInstrumentsV5ReqWithDefaults instantiates a new CreateCopytradingSetInstrumentsV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstId

`func (o *CreateCopytradingSetInstrumentsV5Req) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateCopytradingSetInstrumentsV5Req) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateCopytradingSetInstrumentsV5Req) SetInstId(v string)`

SetInstId sets InstId field to given value.


### GetInstType

`func (o *CreateCopytradingSetInstrumentsV5Req) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *CreateCopytradingSetInstrumentsV5Req) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *CreateCopytradingSetInstrumentsV5Req) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *CreateCopytradingSetInstrumentsV5Req) HasInstType() bool`

HasInstType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


