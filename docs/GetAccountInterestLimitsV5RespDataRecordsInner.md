# GetAccountInterestLimitsV5RespDataRecordsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllAcctRemainingQuota** | Pointer to **string** | Total remaining quota for master account and sub-accounts | [optional] [default to ""]
**AvailLoan** | Pointer to **string** | Available amount for current account (Within the locked quota)  Only applicable to &#x60;VIP loans&#x60; | [optional] [default to ""]
**AvgRate** | Pointer to **string** | Average (hour) interest of already borrowed coin  only applicable to &#x60;VIP loans&#x60; | [optional] [default to ""]
**Ccy** | Pointer to **string** | Loan currency, e.g. &#x60;BTC&#x60; | [optional] [default to ""]
**CurAcctRemainingQuota** | Pointer to **string** | The remaining quota for the current account.  Only applicable to the case in which the sub-account is assigned the loan allocation | [optional] [default to ""]
**Interest** | Pointer to **string** | Interest to be deducted  Only applicable to &#x60;Market loans&#x60; | [optional] [default to ""]
**LoanQuota** | Pointer to **string** | Borrow limit of master account  If loan allocation has been assigned, then it is the borrow limit of the current trading account | [optional] [default to ""]
**PlatRemainingQuota** | Pointer to **string** | Remaining quota for the platform.  The format like  \&quot;600\&quot; will be returned when it is more than &#x60;curAcctRemainingQuota&#x60; or &#x60;allAcctRemainingQuota&#x60; | [optional] [default to ""]
**PosLoan** | Pointer to **string** | Frozen amount for current account (Within the locked quota)   Only applicable to &#x60;VIP loans&#x60; | [optional] [default to ""]
**Rate** | Pointer to **string** | Current daily rate | [optional] [default to ""]
**SurplusLmt** | Pointer to **string** | Available amount across all sub-accounts  If loan allocation has been assigned, then it is the available amount to borrow by the current trading account | [optional] [default to ""]
**SurplusLmtDetails** | Pointer to [**GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails**](GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails.md) |  | [optional] 
**UsedLmt** | Pointer to **string** | Borrowed amount across all sub-accounts  If loan allocation has been assigned, then it is the borrowed amount by the current trading account | [optional] [default to ""]
**UsedLoan** | Pointer to **string** | Borrowed amount for current account  Only applicable to &#x60;VIP loans&#x60; | [optional] [default to ""]

## Methods

### NewGetAccountInterestLimitsV5RespDataRecordsInner

`func NewGetAccountInterestLimitsV5RespDataRecordsInner() *GetAccountInterestLimitsV5RespDataRecordsInner`

NewGetAccountInterestLimitsV5RespDataRecordsInner instantiates a new GetAccountInterestLimitsV5RespDataRecordsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountInterestLimitsV5RespDataRecordsInnerWithDefaults

`func NewGetAccountInterestLimitsV5RespDataRecordsInnerWithDefaults() *GetAccountInterestLimitsV5RespDataRecordsInner`

NewGetAccountInterestLimitsV5RespDataRecordsInnerWithDefaults instantiates a new GetAccountInterestLimitsV5RespDataRecordsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllAcctRemainingQuota

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetAllAcctRemainingQuota() string`

GetAllAcctRemainingQuota returns the AllAcctRemainingQuota field if non-nil, zero value otherwise.

### GetAllAcctRemainingQuotaOk

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetAllAcctRemainingQuotaOk() (*string, bool)`

GetAllAcctRemainingQuotaOk returns a tuple with the AllAcctRemainingQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAcctRemainingQuota

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) SetAllAcctRemainingQuota(v string)`

SetAllAcctRemainingQuota sets AllAcctRemainingQuota field to given value.

### HasAllAcctRemainingQuota

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) HasAllAcctRemainingQuota() bool`

HasAllAcctRemainingQuota returns a boolean if a field has been set.

### GetAvailLoan

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetAvailLoan() string`

GetAvailLoan returns the AvailLoan field if non-nil, zero value otherwise.

### GetAvailLoanOk

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetAvailLoanOk() (*string, bool)`

GetAvailLoanOk returns a tuple with the AvailLoan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailLoan

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) SetAvailLoan(v string)`

SetAvailLoan sets AvailLoan field to given value.

### HasAvailLoan

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) HasAvailLoan() bool`

HasAvailLoan returns a boolean if a field has been set.

### GetAvgRate

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetAvgRate() string`

GetAvgRate returns the AvgRate field if non-nil, zero value otherwise.

### GetAvgRateOk

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetAvgRateOk() (*string, bool)`

GetAvgRateOk returns a tuple with the AvgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgRate

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) SetAvgRate(v string)`

SetAvgRate sets AvgRate field to given value.

### HasAvgRate

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) HasAvgRate() bool`

HasAvgRate returns a boolean if a field has been set.

### GetCcy

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetCurAcctRemainingQuota

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetCurAcctRemainingQuota() string`

GetCurAcctRemainingQuota returns the CurAcctRemainingQuota field if non-nil, zero value otherwise.

### GetCurAcctRemainingQuotaOk

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetCurAcctRemainingQuotaOk() (*string, bool)`

GetCurAcctRemainingQuotaOk returns a tuple with the CurAcctRemainingQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurAcctRemainingQuota

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) SetCurAcctRemainingQuota(v string)`

SetCurAcctRemainingQuota sets CurAcctRemainingQuota field to given value.

### HasCurAcctRemainingQuota

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) HasCurAcctRemainingQuota() bool`

HasCurAcctRemainingQuota returns a boolean if a field has been set.

### GetInterest

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetInterest() string`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetInterestOk() (*string, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) SetInterest(v string)`

SetInterest sets Interest field to given value.

### HasInterest

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) HasInterest() bool`

HasInterest returns a boolean if a field has been set.

### GetLoanQuota

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetLoanQuota() string`

GetLoanQuota returns the LoanQuota field if non-nil, zero value otherwise.

### GetLoanQuotaOk

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetLoanQuotaOk() (*string, bool)`

GetLoanQuotaOk returns a tuple with the LoanQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanQuota

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) SetLoanQuota(v string)`

SetLoanQuota sets LoanQuota field to given value.

### HasLoanQuota

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) HasLoanQuota() bool`

HasLoanQuota returns a boolean if a field has been set.

### GetPlatRemainingQuota

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetPlatRemainingQuota() string`

GetPlatRemainingQuota returns the PlatRemainingQuota field if non-nil, zero value otherwise.

### GetPlatRemainingQuotaOk

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetPlatRemainingQuotaOk() (*string, bool)`

GetPlatRemainingQuotaOk returns a tuple with the PlatRemainingQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatRemainingQuota

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) SetPlatRemainingQuota(v string)`

SetPlatRemainingQuota sets PlatRemainingQuota field to given value.

### HasPlatRemainingQuota

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) HasPlatRemainingQuota() bool`

HasPlatRemainingQuota returns a boolean if a field has been set.

### GetPosLoan

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetPosLoan() string`

GetPosLoan returns the PosLoan field if non-nil, zero value otherwise.

### GetPosLoanOk

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetPosLoanOk() (*string, bool)`

GetPosLoanOk returns a tuple with the PosLoan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosLoan

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) SetPosLoan(v string)`

SetPosLoan sets PosLoan field to given value.

### HasPosLoan

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) HasPosLoan() bool`

HasPosLoan returns a boolean if a field has been set.

### GetRate

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetSurplusLmt

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetSurplusLmt() string`

GetSurplusLmt returns the SurplusLmt field if non-nil, zero value otherwise.

### GetSurplusLmtOk

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetSurplusLmtOk() (*string, bool)`

GetSurplusLmtOk returns a tuple with the SurplusLmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurplusLmt

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) SetSurplusLmt(v string)`

SetSurplusLmt sets SurplusLmt field to given value.

### HasSurplusLmt

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) HasSurplusLmt() bool`

HasSurplusLmt returns a boolean if a field has been set.

### GetSurplusLmtDetails

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetSurplusLmtDetails() GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails`

GetSurplusLmtDetails returns the SurplusLmtDetails field if non-nil, zero value otherwise.

### GetSurplusLmtDetailsOk

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetSurplusLmtDetailsOk() (*GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails, bool)`

GetSurplusLmtDetailsOk returns a tuple with the SurplusLmtDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurplusLmtDetails

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) SetSurplusLmtDetails(v GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails)`

SetSurplusLmtDetails sets SurplusLmtDetails field to given value.

### HasSurplusLmtDetails

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) HasSurplusLmtDetails() bool`

HasSurplusLmtDetails returns a boolean if a field has been set.

### GetUsedLmt

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetUsedLmt() string`

GetUsedLmt returns the UsedLmt field if non-nil, zero value otherwise.

### GetUsedLmtOk

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetUsedLmtOk() (*string, bool)`

GetUsedLmtOk returns a tuple with the UsedLmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedLmt

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) SetUsedLmt(v string)`

SetUsedLmt sets UsedLmt field to given value.

### HasUsedLmt

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) HasUsedLmt() bool`

HasUsedLmt returns a boolean if a field has been set.

### GetUsedLoan

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetUsedLoan() string`

GetUsedLoan returns the UsedLoan field if non-nil, zero value otherwise.

### GetUsedLoanOk

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) GetUsedLoanOk() (*string, bool)`

GetUsedLoanOk returns a tuple with the UsedLoan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedLoan

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) SetUsedLoan(v string)`

SetUsedLoan sets UsedLoan field to given value.

### HasUsedLoan

`func (o *GetAccountInterestLimitsV5RespDataRecordsInner) HasUsedLoan() bool`

HasUsedLoan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


