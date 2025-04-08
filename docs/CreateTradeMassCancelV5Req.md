# CreateTradeMassCancelV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstFamily** | **string** | Instrument family | [default to ""]
**InstType** | **string** | Instrument type  &#x60;OPTION&#x60; | [default to ""]
**LockInterval** | Pointer to **string** | Lock interval(ms)   The range should be [0, 10 000]   The default is 0. You can set it as \&quot;0\&quot; if you want to unlock it immediately.   Error 54008 will be returned when placing order during lock interval, it is different from 51034 which is thrown when MMP is triggered | [optional] [default to ""]

## Methods

### NewCreateTradeMassCancelV5Req

`func NewCreateTradeMassCancelV5Req(instFamily string, instType string, ) *CreateTradeMassCancelV5Req`

NewCreateTradeMassCancelV5Req instantiates a new CreateTradeMassCancelV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeMassCancelV5ReqWithDefaults

`func NewCreateTradeMassCancelV5ReqWithDefaults() *CreateTradeMassCancelV5Req`

NewCreateTradeMassCancelV5ReqWithDefaults instantiates a new CreateTradeMassCancelV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstFamily

`func (o *CreateTradeMassCancelV5Req) GetInstFamily() string`

GetInstFamily returns the InstFamily field if non-nil, zero value otherwise.

### GetInstFamilyOk

`func (o *CreateTradeMassCancelV5Req) GetInstFamilyOk() (*string, bool)`

GetInstFamilyOk returns a tuple with the InstFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstFamily

`func (o *CreateTradeMassCancelV5Req) SetInstFamily(v string)`

SetInstFamily sets InstFamily field to given value.


### GetInstType

`func (o *CreateTradeMassCancelV5Req) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *CreateTradeMassCancelV5Req) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *CreateTradeMassCancelV5Req) SetInstType(v string)`

SetInstType sets InstType field to given value.


### GetLockInterval

`func (o *CreateTradeMassCancelV5Req) GetLockInterval() string`

GetLockInterval returns the LockInterval field if non-nil, zero value otherwise.

### GetLockIntervalOk

`func (o *CreateTradeMassCancelV5Req) GetLockIntervalOk() (*string, bool)`

GetLockIntervalOk returns a tuple with the LockInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockInterval

`func (o *CreateTradeMassCancelV5Req) SetLockInterval(v string)`

SetLockInterval sets LockInterval field to given value.

### HasLockInterval

`func (o *CreateTradeMassCancelV5Req) HasLockInterval() bool`

HasLockInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


