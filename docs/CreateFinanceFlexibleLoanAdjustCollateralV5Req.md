# CreateFinanceFlexibleLoanAdjustCollateralV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollateralAmt** | **string** | Collateral amount | [default to ""]
**CollateralCcy** | **string** | Collateral currency, e.g. &#x60;BTC&#x60; | [default to ""]
**Type** | **string** | Operation type  &#x60;add&#x60;: Add collateral  &#x60;reduce&#x60;: Reduce collateral | [default to ""]

## Methods

### NewCreateFinanceFlexibleLoanAdjustCollateralV5Req

`func NewCreateFinanceFlexibleLoanAdjustCollateralV5Req(collateralAmt string, collateralCcy string, type_ string, ) *CreateFinanceFlexibleLoanAdjustCollateralV5Req`

NewCreateFinanceFlexibleLoanAdjustCollateralV5Req instantiates a new CreateFinanceFlexibleLoanAdjustCollateralV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFinanceFlexibleLoanAdjustCollateralV5ReqWithDefaults

`func NewCreateFinanceFlexibleLoanAdjustCollateralV5ReqWithDefaults() *CreateFinanceFlexibleLoanAdjustCollateralV5Req`

NewCreateFinanceFlexibleLoanAdjustCollateralV5ReqWithDefaults instantiates a new CreateFinanceFlexibleLoanAdjustCollateralV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollateralAmt

`func (o *CreateFinanceFlexibleLoanAdjustCollateralV5Req) GetCollateralAmt() string`

GetCollateralAmt returns the CollateralAmt field if non-nil, zero value otherwise.

### GetCollateralAmtOk

`func (o *CreateFinanceFlexibleLoanAdjustCollateralV5Req) GetCollateralAmtOk() (*string, bool)`

GetCollateralAmtOk returns a tuple with the CollateralAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralAmt

`func (o *CreateFinanceFlexibleLoanAdjustCollateralV5Req) SetCollateralAmt(v string)`

SetCollateralAmt sets CollateralAmt field to given value.


### GetCollateralCcy

`func (o *CreateFinanceFlexibleLoanAdjustCollateralV5Req) GetCollateralCcy() string`

GetCollateralCcy returns the CollateralCcy field if non-nil, zero value otherwise.

### GetCollateralCcyOk

`func (o *CreateFinanceFlexibleLoanAdjustCollateralV5Req) GetCollateralCcyOk() (*string, bool)`

GetCollateralCcyOk returns a tuple with the CollateralCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralCcy

`func (o *CreateFinanceFlexibleLoanAdjustCollateralV5Req) SetCollateralCcy(v string)`

SetCollateralCcy sets CollateralCcy field to given value.


### GetType

`func (o *CreateFinanceFlexibleLoanAdjustCollateralV5Req) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateFinanceFlexibleLoanAdjustCollateralV5Req) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateFinanceFlexibleLoanAdjustCollateralV5Req) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


