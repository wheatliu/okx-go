# CreateCopytradingStopCopyTradingV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstType** | Pointer to **string** | Instrument type  &#x60;SWAP&#x60; | [optional] [default to ""]
**SubPosCloseType** | **string** | Action type for open positions, it is required if you have related copy position  &#x60;market_close&#x60;: immediately close at market price  &#x60;copy_close&#x60;：close when trader closes  &#x60;manual_close&#x60;: close manually | [default to ""]
**UniqueCode** | **string** | Lead trader unique code  A combination of case-sensitive alphanumerics, all numbers and the length is 16 characters, e.g. 213E8C92DC61EFAC | [default to ""]

## Methods

### NewCreateCopytradingStopCopyTradingV5Req

`func NewCreateCopytradingStopCopyTradingV5Req(subPosCloseType string, uniqueCode string, ) *CreateCopytradingStopCopyTradingV5Req`

NewCreateCopytradingStopCopyTradingV5Req instantiates a new CreateCopytradingStopCopyTradingV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCopytradingStopCopyTradingV5ReqWithDefaults

`func NewCreateCopytradingStopCopyTradingV5ReqWithDefaults() *CreateCopytradingStopCopyTradingV5Req`

NewCreateCopytradingStopCopyTradingV5ReqWithDefaults instantiates a new CreateCopytradingStopCopyTradingV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstType

`func (o *CreateCopytradingStopCopyTradingV5Req) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *CreateCopytradingStopCopyTradingV5Req) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *CreateCopytradingStopCopyTradingV5Req) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *CreateCopytradingStopCopyTradingV5Req) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetSubPosCloseType

`func (o *CreateCopytradingStopCopyTradingV5Req) GetSubPosCloseType() string`

GetSubPosCloseType returns the SubPosCloseType field if non-nil, zero value otherwise.

### GetSubPosCloseTypeOk

`func (o *CreateCopytradingStopCopyTradingV5Req) GetSubPosCloseTypeOk() (*string, bool)`

GetSubPosCloseTypeOk returns a tuple with the SubPosCloseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPosCloseType

`func (o *CreateCopytradingStopCopyTradingV5Req) SetSubPosCloseType(v string)`

SetSubPosCloseType sets SubPosCloseType field to given value.


### GetUniqueCode

`func (o *CreateCopytradingStopCopyTradingV5Req) GetUniqueCode() string`

GetUniqueCode returns the UniqueCode field if non-nil, zero value otherwise.

### GetUniqueCodeOk

`func (o *CreateCopytradingStopCopyTradingV5Req) GetUniqueCodeOk() (*string, bool)`

GetUniqueCodeOk returns a tuple with the UniqueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueCode

`func (o *CreateCopytradingStopCopyTradingV5Req) SetUniqueCode(v string)`

SetUniqueCode sets UniqueCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


