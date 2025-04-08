# GetPublicInterestRateLoanQuotaV5RespDataRegularInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IrDiscount** | Pointer to **string** | Interest rate discount(Deprecated) | [optional] [default to ""]
**Level** | Pointer to **string** | Regular user Level, e.g. &#x60;Lv1&#x60; | [optional] [default to ""]
**LoanQuotaCoef** | Pointer to **string** | Loan quota coefficient. Loan quota &#x3D; &#x60;quota&#x60; * &#x60;level&#x60; | [optional] [default to ""]

## Methods

### NewGetPublicInterestRateLoanQuotaV5RespDataRegularInner

`func NewGetPublicInterestRateLoanQuotaV5RespDataRegularInner() *GetPublicInterestRateLoanQuotaV5RespDataRegularInner`

NewGetPublicInterestRateLoanQuotaV5RespDataRegularInner instantiates a new GetPublicInterestRateLoanQuotaV5RespDataRegularInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicInterestRateLoanQuotaV5RespDataRegularInnerWithDefaults

`func NewGetPublicInterestRateLoanQuotaV5RespDataRegularInnerWithDefaults() *GetPublicInterestRateLoanQuotaV5RespDataRegularInner`

NewGetPublicInterestRateLoanQuotaV5RespDataRegularInnerWithDefaults instantiates a new GetPublicInterestRateLoanQuotaV5RespDataRegularInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIrDiscount

`func (o *GetPublicInterestRateLoanQuotaV5RespDataRegularInner) GetIrDiscount() string`

GetIrDiscount returns the IrDiscount field if non-nil, zero value otherwise.

### GetIrDiscountOk

`func (o *GetPublicInterestRateLoanQuotaV5RespDataRegularInner) GetIrDiscountOk() (*string, bool)`

GetIrDiscountOk returns a tuple with the IrDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIrDiscount

`func (o *GetPublicInterestRateLoanQuotaV5RespDataRegularInner) SetIrDiscount(v string)`

SetIrDiscount sets IrDiscount field to given value.

### HasIrDiscount

`func (o *GetPublicInterestRateLoanQuotaV5RespDataRegularInner) HasIrDiscount() bool`

HasIrDiscount returns a boolean if a field has been set.

### GetLevel

`func (o *GetPublicInterestRateLoanQuotaV5RespDataRegularInner) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *GetPublicInterestRateLoanQuotaV5RespDataRegularInner) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *GetPublicInterestRateLoanQuotaV5RespDataRegularInner) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *GetPublicInterestRateLoanQuotaV5RespDataRegularInner) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetLoanQuotaCoef

`func (o *GetPublicInterestRateLoanQuotaV5RespDataRegularInner) GetLoanQuotaCoef() string`

GetLoanQuotaCoef returns the LoanQuotaCoef field if non-nil, zero value otherwise.

### GetLoanQuotaCoefOk

`func (o *GetPublicInterestRateLoanQuotaV5RespDataRegularInner) GetLoanQuotaCoefOk() (*string, bool)`

GetLoanQuotaCoefOk returns a tuple with the LoanQuotaCoef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanQuotaCoef

`func (o *GetPublicInterestRateLoanQuotaV5RespDataRegularInner) SetLoanQuotaCoef(v string)`

SetLoanQuotaCoef sets LoanQuotaCoef field to given value.

### HasLoanQuotaCoef

`func (o *GetPublicInterestRateLoanQuotaV5RespDataRegularInner) HasLoanQuotaCoef() bool`

HasLoanQuotaCoef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


