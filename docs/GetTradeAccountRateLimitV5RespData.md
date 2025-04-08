# GetTradeAccountRateLimitV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccRateLimit** | Pointer to **string** | Current sub-account rate limit per two seconds | [optional] [default to ""]
**FillRatio** | Pointer to **string** | Sub account fill ratio during the monitoring period   Applicable for users with trading fee level &gt;&#x3D; VIP 5 and return \&quot;\&quot; for others   For accounts with no trading volume during the monitoring period, return \&quot;0\&quot;. For accounts with trading volume but no order count due to our counting logic, return \&quot;9999\&quot;. | [optional] [default to ""]
**MainFillRatio** | Pointer to **string** | Master account aggregated fill ratio during the monitoring period   Applicable for users with trading fee level &gt;&#x3D; VIP 5 and return \&quot;\&quot; for others   For accounts with no trading volume during the monitoring period, return \&quot;0\&quot; | [optional] [default to ""]
**NextAccRateLimit** | Pointer to **string** | Expected sub-account rate limit (per two seconds) in the next period   Applicable for users with trading fee level &gt;&#x3D; VIP 5 and return \&quot;\&quot; for others | [optional] [default to ""]
**Ts** | Pointer to **string** | Data update time   For users with trading fee level &gt;&#x3D; VIP 5, the data will be generated at 08:00 am (UTC)   For users with trading fee level &lt; VIP 5, return the current timestamp | [optional] [default to ""]

## Methods

### NewGetTradeAccountRateLimitV5RespData

`func NewGetTradeAccountRateLimitV5RespData() *GetTradeAccountRateLimitV5RespData`

NewGetTradeAccountRateLimitV5RespData instantiates a new GetTradeAccountRateLimitV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradeAccountRateLimitV5RespDataWithDefaults

`func NewGetTradeAccountRateLimitV5RespDataWithDefaults() *GetTradeAccountRateLimitV5RespData`

NewGetTradeAccountRateLimitV5RespDataWithDefaults instantiates a new GetTradeAccountRateLimitV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccRateLimit

`func (o *GetTradeAccountRateLimitV5RespData) GetAccRateLimit() string`

GetAccRateLimit returns the AccRateLimit field if non-nil, zero value otherwise.

### GetAccRateLimitOk

`func (o *GetTradeAccountRateLimitV5RespData) GetAccRateLimitOk() (*string, bool)`

GetAccRateLimitOk returns a tuple with the AccRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccRateLimit

`func (o *GetTradeAccountRateLimitV5RespData) SetAccRateLimit(v string)`

SetAccRateLimit sets AccRateLimit field to given value.

### HasAccRateLimit

`func (o *GetTradeAccountRateLimitV5RespData) HasAccRateLimit() bool`

HasAccRateLimit returns a boolean if a field has been set.

### GetFillRatio

`func (o *GetTradeAccountRateLimitV5RespData) GetFillRatio() string`

GetFillRatio returns the FillRatio field if non-nil, zero value otherwise.

### GetFillRatioOk

`func (o *GetTradeAccountRateLimitV5RespData) GetFillRatioOk() (*string, bool)`

GetFillRatioOk returns a tuple with the FillRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillRatio

`func (o *GetTradeAccountRateLimitV5RespData) SetFillRatio(v string)`

SetFillRatio sets FillRatio field to given value.

### HasFillRatio

`func (o *GetTradeAccountRateLimitV5RespData) HasFillRatio() bool`

HasFillRatio returns a boolean if a field has been set.

### GetMainFillRatio

`func (o *GetTradeAccountRateLimitV5RespData) GetMainFillRatio() string`

GetMainFillRatio returns the MainFillRatio field if non-nil, zero value otherwise.

### GetMainFillRatioOk

`func (o *GetTradeAccountRateLimitV5RespData) GetMainFillRatioOk() (*string, bool)`

GetMainFillRatioOk returns a tuple with the MainFillRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainFillRatio

`func (o *GetTradeAccountRateLimitV5RespData) SetMainFillRatio(v string)`

SetMainFillRatio sets MainFillRatio field to given value.

### HasMainFillRatio

`func (o *GetTradeAccountRateLimitV5RespData) HasMainFillRatio() bool`

HasMainFillRatio returns a boolean if a field has been set.

### GetNextAccRateLimit

`func (o *GetTradeAccountRateLimitV5RespData) GetNextAccRateLimit() string`

GetNextAccRateLimit returns the NextAccRateLimit field if non-nil, zero value otherwise.

### GetNextAccRateLimitOk

`func (o *GetTradeAccountRateLimitV5RespData) GetNextAccRateLimitOk() (*string, bool)`

GetNextAccRateLimitOk returns a tuple with the NextAccRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAccRateLimit

`func (o *GetTradeAccountRateLimitV5RespData) SetNextAccRateLimit(v string)`

SetNextAccRateLimit sets NextAccRateLimit field to given value.

### HasNextAccRateLimit

`func (o *GetTradeAccountRateLimitV5RespData) HasNextAccRateLimit() bool`

HasNextAccRateLimit returns a boolean if a field has been set.

### GetTs

`func (o *GetTradeAccountRateLimitV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetTradeAccountRateLimitV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetTradeAccountRateLimitV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetTradeAccountRateLimitV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


