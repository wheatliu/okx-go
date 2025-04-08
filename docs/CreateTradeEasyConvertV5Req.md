# CreateTradeEasyConvertV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromCcy** | **[]string** | Type of small payment currency convert from   Maximum 5 currencies can be selected in one order. If there are multiple currencies, separate them with commas. | 
**Source** | Pointer to **string** | Funding source  &#x60;1&#x60;: Trading account  &#x60;2&#x60;: Funding account  The default is &#x60;1&#x60;. | [optional] [default to ""]
**ToCcy** | **string** | Type of mainstream currency convert to   Only one receiving currency type can be selected in one order and cannot be the same as the small payment currencies. | [default to ""]

## Methods

### NewCreateTradeEasyConvertV5Req

`func NewCreateTradeEasyConvertV5Req(fromCcy []string, toCcy string, ) *CreateTradeEasyConvertV5Req`

NewCreateTradeEasyConvertV5Req instantiates a new CreateTradeEasyConvertV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeEasyConvertV5ReqWithDefaults

`func NewCreateTradeEasyConvertV5ReqWithDefaults() *CreateTradeEasyConvertV5Req`

NewCreateTradeEasyConvertV5ReqWithDefaults instantiates a new CreateTradeEasyConvertV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromCcy

`func (o *CreateTradeEasyConvertV5Req) GetFromCcy() []string`

GetFromCcy returns the FromCcy field if non-nil, zero value otherwise.

### GetFromCcyOk

`func (o *CreateTradeEasyConvertV5Req) GetFromCcyOk() (*[]string, bool)`

GetFromCcyOk returns a tuple with the FromCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCcy

`func (o *CreateTradeEasyConvertV5Req) SetFromCcy(v []string)`

SetFromCcy sets FromCcy field to given value.


### GetSource

`func (o *CreateTradeEasyConvertV5Req) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateTradeEasyConvertV5Req) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateTradeEasyConvertV5Req) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CreateTradeEasyConvertV5Req) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetToCcy

`func (o *CreateTradeEasyConvertV5Req) GetToCcy() string`

GetToCcy returns the ToCcy field if non-nil, zero value otherwise.

### GetToCcyOk

`func (o *CreateTradeEasyConvertV5Req) GetToCcyOk() (*string, bool)`

GetToCcyOk returns a tuple with the ToCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCcy

`func (o *CreateTradeEasyConvertV5Req) SetToCcy(v string)`

SetToCcy sets ToCcy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


