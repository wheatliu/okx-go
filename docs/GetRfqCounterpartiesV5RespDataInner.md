# GetRfqCounterpartiesV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TraderCode** | Pointer to **string** | A unique identifier of maker which will be publicly visible on the platform. All RFQ and Quote endpoints will use this as the unique counterparty identifier. | [optional] [default to ""]
**TraderName** | Pointer to **string** | The long formative username of trader or entity on the platform. | [optional] [default to ""]
**Type** | Pointer to **string** | The counterparty type. &#x60;LP&#x60; refers to API connected auto market makers. | [optional] [default to ""]

## Methods

### NewGetRfqCounterpartiesV5RespDataInner

`func NewGetRfqCounterpartiesV5RespDataInner() *GetRfqCounterpartiesV5RespDataInner`

NewGetRfqCounterpartiesV5RespDataInner instantiates a new GetRfqCounterpartiesV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRfqCounterpartiesV5RespDataInnerWithDefaults

`func NewGetRfqCounterpartiesV5RespDataInnerWithDefaults() *GetRfqCounterpartiesV5RespDataInner`

NewGetRfqCounterpartiesV5RespDataInnerWithDefaults instantiates a new GetRfqCounterpartiesV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraderCode

`func (o *GetRfqCounterpartiesV5RespDataInner) GetTraderCode() string`

GetTraderCode returns the TraderCode field if non-nil, zero value otherwise.

### GetTraderCodeOk

`func (o *GetRfqCounterpartiesV5RespDataInner) GetTraderCodeOk() (*string, bool)`

GetTraderCodeOk returns a tuple with the TraderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraderCode

`func (o *GetRfqCounterpartiesV5RespDataInner) SetTraderCode(v string)`

SetTraderCode sets TraderCode field to given value.

### HasTraderCode

`func (o *GetRfqCounterpartiesV5RespDataInner) HasTraderCode() bool`

HasTraderCode returns a boolean if a field has been set.

### GetTraderName

`func (o *GetRfqCounterpartiesV5RespDataInner) GetTraderName() string`

GetTraderName returns the TraderName field if non-nil, zero value otherwise.

### GetTraderNameOk

`func (o *GetRfqCounterpartiesV5RespDataInner) GetTraderNameOk() (*string, bool)`

GetTraderNameOk returns a tuple with the TraderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraderName

`func (o *GetRfqCounterpartiesV5RespDataInner) SetTraderName(v string)`

SetTraderName sets TraderName field to given value.

### HasTraderName

`func (o *GetRfqCounterpartiesV5RespDataInner) HasTraderName() bool`

HasTraderName returns a boolean if a field has been set.

### GetType

`func (o *GetRfqCounterpartiesV5RespDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetRfqCounterpartiesV5RespDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetRfqCounterpartiesV5RespDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetRfqCounterpartiesV5RespDataInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


