# GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllAcctRemainingQuota** | Pointer to **string** | Total remaining quota for master account and sub-accounts | [optional] [default to ""]
**CurAcctRemainingQuota** | Pointer to **string** | The remaining quota for the current account.  Only applicable to the case in which the sub-account is assigned the loan allocation | [optional] [default to ""]
**PlatRemainingQuota** | Pointer to **string** | Remaining quota for the platform.  The format like  \&quot;600\&quot; will be returned when it is more than &#x60;curAcctRemainingQuota&#x60; or &#x60;allAcctRemainingQuota&#x60; | [optional] [default to ""]

## Methods

### NewGetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails

`func NewGetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails() *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails`

NewGetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails instantiates a new GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetailsWithDefaults

`func NewGetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetailsWithDefaults() *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails`

NewGetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetailsWithDefaults instantiates a new GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllAcctRemainingQuota

`func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) GetAllAcctRemainingQuota() string`

GetAllAcctRemainingQuota returns the AllAcctRemainingQuota field if non-nil, zero value otherwise.

### GetAllAcctRemainingQuotaOk

`func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) GetAllAcctRemainingQuotaOk() (*string, bool)`

GetAllAcctRemainingQuotaOk returns a tuple with the AllAcctRemainingQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAcctRemainingQuota

`func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) SetAllAcctRemainingQuota(v string)`

SetAllAcctRemainingQuota sets AllAcctRemainingQuota field to given value.

### HasAllAcctRemainingQuota

`func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) HasAllAcctRemainingQuota() bool`

HasAllAcctRemainingQuota returns a boolean if a field has been set.

### GetCurAcctRemainingQuota

`func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) GetCurAcctRemainingQuota() string`

GetCurAcctRemainingQuota returns the CurAcctRemainingQuota field if non-nil, zero value otherwise.

### GetCurAcctRemainingQuotaOk

`func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) GetCurAcctRemainingQuotaOk() (*string, bool)`

GetCurAcctRemainingQuotaOk returns a tuple with the CurAcctRemainingQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurAcctRemainingQuota

`func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) SetCurAcctRemainingQuota(v string)`

SetCurAcctRemainingQuota sets CurAcctRemainingQuota field to given value.

### HasCurAcctRemainingQuota

`func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) HasCurAcctRemainingQuota() bool`

HasCurAcctRemainingQuota returns a boolean if a field has been set.

### GetPlatRemainingQuota

`func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) GetPlatRemainingQuota() string`

GetPlatRemainingQuota returns the PlatRemainingQuota field if non-nil, zero value otherwise.

### GetPlatRemainingQuotaOk

`func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) GetPlatRemainingQuotaOk() (*string, bool)`

GetPlatRemainingQuotaOk returns a tuple with the PlatRemainingQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatRemainingQuota

`func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) SetPlatRemainingQuota(v string)`

SetPlatRemainingQuota sets PlatRemainingQuota field to given value.

### HasPlatRemainingQuota

`func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) HasPlatRemainingQuota() bool`

HasPlatRemainingQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


