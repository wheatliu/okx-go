# CreateAccountBillsHistoryArchiveV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quarter** | **string** | Quarter, valid value is &#x60;Q1&#x60;, &#x60;Q2&#x60;, &#x60;Q3&#x60;, &#x60;Q4&#x60; | [default to ""]
**Year** | **string** | 4 digits year | [default to ""]

## Methods

### NewCreateAccountBillsHistoryArchiveV5Req

`func NewCreateAccountBillsHistoryArchiveV5Req(quarter string, year string, ) *CreateAccountBillsHistoryArchiveV5Req`

NewCreateAccountBillsHistoryArchiveV5Req instantiates a new CreateAccountBillsHistoryArchiveV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountBillsHistoryArchiveV5ReqWithDefaults

`func NewCreateAccountBillsHistoryArchiveV5ReqWithDefaults() *CreateAccountBillsHistoryArchiveV5Req`

NewCreateAccountBillsHistoryArchiveV5ReqWithDefaults instantiates a new CreateAccountBillsHistoryArchiveV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuarter

`func (o *CreateAccountBillsHistoryArchiveV5Req) GetQuarter() string`

GetQuarter returns the Quarter field if non-nil, zero value otherwise.

### GetQuarterOk

`func (o *CreateAccountBillsHistoryArchiveV5Req) GetQuarterOk() (*string, bool)`

GetQuarterOk returns a tuple with the Quarter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarter

`func (o *CreateAccountBillsHistoryArchiveV5Req) SetQuarter(v string)`

SetQuarter sets Quarter field to given value.


### GetYear

`func (o *CreateAccountBillsHistoryArchiveV5Req) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *CreateAccountBillsHistoryArchiveV5Req) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *CreateAccountBillsHistoryArchiveV5Req) SetYear(v string)`

SetYear sets Year field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


