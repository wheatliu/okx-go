# CreateFinanceStakingDefiRedeemV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowEarlyRedeem** | Pointer to **bool** | Whether allows early redemption  Default is &#x60;false&#x60; | [optional] 
**OrdId** | **string** | Order ID | [default to ""]
**ProtocolType** | **string** | Protocol type  &#x60;defi&#x60;: on-chain earn | [default to ""]

## Methods

### NewCreateFinanceStakingDefiRedeemV5Req

`func NewCreateFinanceStakingDefiRedeemV5Req(ordId string, protocolType string, ) *CreateFinanceStakingDefiRedeemV5Req`

NewCreateFinanceStakingDefiRedeemV5Req instantiates a new CreateFinanceStakingDefiRedeemV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFinanceStakingDefiRedeemV5ReqWithDefaults

`func NewCreateFinanceStakingDefiRedeemV5ReqWithDefaults() *CreateFinanceStakingDefiRedeemV5Req`

NewCreateFinanceStakingDefiRedeemV5ReqWithDefaults instantiates a new CreateFinanceStakingDefiRedeemV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowEarlyRedeem

`func (o *CreateFinanceStakingDefiRedeemV5Req) GetAllowEarlyRedeem() bool`

GetAllowEarlyRedeem returns the AllowEarlyRedeem field if non-nil, zero value otherwise.

### GetAllowEarlyRedeemOk

`func (o *CreateFinanceStakingDefiRedeemV5Req) GetAllowEarlyRedeemOk() (*bool, bool)`

GetAllowEarlyRedeemOk returns a tuple with the AllowEarlyRedeem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEarlyRedeem

`func (o *CreateFinanceStakingDefiRedeemV5Req) SetAllowEarlyRedeem(v bool)`

SetAllowEarlyRedeem sets AllowEarlyRedeem field to given value.

### HasAllowEarlyRedeem

`func (o *CreateFinanceStakingDefiRedeemV5Req) HasAllowEarlyRedeem() bool`

HasAllowEarlyRedeem returns a boolean if a field has been set.

### GetOrdId

`func (o *CreateFinanceStakingDefiRedeemV5Req) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *CreateFinanceStakingDefiRedeemV5Req) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *CreateFinanceStakingDefiRedeemV5Req) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.


### GetProtocolType

`func (o *CreateFinanceStakingDefiRedeemV5Req) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *CreateFinanceStakingDefiRedeemV5Req) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *CreateFinanceStakingDefiRedeemV5Req) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


