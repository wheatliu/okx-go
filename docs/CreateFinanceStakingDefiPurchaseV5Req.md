# CreateFinanceStakingDefiPurchaseV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvestData** | [**[]CreateFinanceStakingDefiPurchaseV5ReqInvestDataInner**](CreateFinanceStakingDefiPurchaseV5ReqInvestDataInner.md) | Investment data | 
**ProductId** | **string** | Product ID | [default to ""]
**Tag** | Pointer to **string** | Order tag  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 16 characters. | [optional] [default to ""]
**Term** | Pointer to **string** | Investment term  Investment term must be specified for fixed-term product | [optional] [default to ""]

## Methods

### NewCreateFinanceStakingDefiPurchaseV5Req

`func NewCreateFinanceStakingDefiPurchaseV5Req(investData []CreateFinanceStakingDefiPurchaseV5ReqInvestDataInner, productId string, ) *CreateFinanceStakingDefiPurchaseV5Req`

NewCreateFinanceStakingDefiPurchaseV5Req instantiates a new CreateFinanceStakingDefiPurchaseV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFinanceStakingDefiPurchaseV5ReqWithDefaults

`func NewCreateFinanceStakingDefiPurchaseV5ReqWithDefaults() *CreateFinanceStakingDefiPurchaseV5Req`

NewCreateFinanceStakingDefiPurchaseV5ReqWithDefaults instantiates a new CreateFinanceStakingDefiPurchaseV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestData

`func (o *CreateFinanceStakingDefiPurchaseV5Req) GetInvestData() []CreateFinanceStakingDefiPurchaseV5ReqInvestDataInner`

GetInvestData returns the InvestData field if non-nil, zero value otherwise.

### GetInvestDataOk

`func (o *CreateFinanceStakingDefiPurchaseV5Req) GetInvestDataOk() (*[]CreateFinanceStakingDefiPurchaseV5ReqInvestDataInner, bool)`

GetInvestDataOk returns a tuple with the InvestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestData

`func (o *CreateFinanceStakingDefiPurchaseV5Req) SetInvestData(v []CreateFinanceStakingDefiPurchaseV5ReqInvestDataInner)`

SetInvestData sets InvestData field to given value.


### GetProductId

`func (o *CreateFinanceStakingDefiPurchaseV5Req) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CreateFinanceStakingDefiPurchaseV5Req) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CreateFinanceStakingDefiPurchaseV5Req) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetTag

`func (o *CreateFinanceStakingDefiPurchaseV5Req) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateFinanceStakingDefiPurchaseV5Req) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateFinanceStakingDefiPurchaseV5Req) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateFinanceStakingDefiPurchaseV5Req) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTerm

`func (o *CreateFinanceStakingDefiPurchaseV5Req) GetTerm() string`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *CreateFinanceStakingDefiPurchaseV5Req) GetTermOk() (*string, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *CreateFinanceStakingDefiPurchaseV5Req) SetTerm(v string)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *CreateFinanceStakingDefiPurchaseV5Req) HasTerm() bool`

HasTerm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


