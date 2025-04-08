# CreateFinanceStakingDefiCancelV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrdId** | **string** | Order ID | [default to ""]
**ProtocolType** | **string** | Protocol type  &#x60;defi&#x60;: on-chain earn | [default to ""]

## Methods

### NewCreateFinanceStakingDefiCancelV5Req

`func NewCreateFinanceStakingDefiCancelV5Req(ordId string, protocolType string, ) *CreateFinanceStakingDefiCancelV5Req`

NewCreateFinanceStakingDefiCancelV5Req instantiates a new CreateFinanceStakingDefiCancelV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFinanceStakingDefiCancelV5ReqWithDefaults

`func NewCreateFinanceStakingDefiCancelV5ReqWithDefaults() *CreateFinanceStakingDefiCancelV5Req`

NewCreateFinanceStakingDefiCancelV5ReqWithDefaults instantiates a new CreateFinanceStakingDefiCancelV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrdId

`func (o *CreateFinanceStakingDefiCancelV5Req) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *CreateFinanceStakingDefiCancelV5Req) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *CreateFinanceStakingDefiCancelV5Req) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.


### GetProtocolType

`func (o *CreateFinanceStakingDefiCancelV5Req) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *CreateFinanceStakingDefiCancelV5Req) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *CreateFinanceStakingDefiCancelV5Req) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


