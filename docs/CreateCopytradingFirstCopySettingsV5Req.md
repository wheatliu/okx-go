# CreateCopytradingFirstCopySettingsV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopyAmt** | Pointer to **string** | Copy amount per order in USDT. | [optional] [default to ""]
**CopyInstIdType** | **string** | Copy contract type setted  &#x60;custom&#x60;: custom by &#x60;instId&#x60; which is required；  &#x60;copy&#x60;: Keep your contracts consistent with this trader by automatically adding or removing contracts when they do | [default to ""]
**CopyMgnMode** | **string** | Copy margin mode  &#x60;cross&#x60;: cross  &#x60;isolated&#x60;: isolated  &#x60;copy&#x60;: Use the same margin mode as lead trader when opening positions | [default to ""]
**CopyMode** | Pointer to **string** | Copy mode  &#x60;fixed_amount&#x60;: set the same fixed amount for each order, and &#x60;copyAmt&#x60; is required；  &#x60;ratio_copy&#x60;: set amount as a multiple of the lead trader’s order value, and &#x60;copyRatio&#x60; is required   The default is &#x60;fixed_amount&#x60; | [optional] [default to ""]
**CopyRatio** | Pointer to **string** | Copy ratio per order. | [optional] [default to ""]
**CopyTotalAmt** | **string** | Maximum total amount in USDT.   The maximum total amount you&#39;ll invest at any given time across all orders in this copy trade  You won’t copy new orders if you exceed this amount | [default to ""]
**InstId** | Pointer to **string** | Instrument ID.   If there are multiple instruments, separate them with commas. | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type  &#x60;SWAP&#x60;, the default value | [optional] [default to ""]
**SlRatio** | Pointer to **string** | Stop loss per order. 0.1 represents 10% | [optional] [default to ""]
**SlTotalAmt** | Pointer to **string** | Total stop loss in USDT for trader.   If your net loss (total profit - total loss) reaches this amount, you&#39;ll stop copying this trader | [optional] [default to ""]
**SubPosCloseType** | **string** | Action type for open positions  &#x60;market_close&#x60;: immediately close at market price  &#x60;copy_close&#x60;：close when trader closes  &#x60;manual_close&#x60;: close manually  The default is &#x60;copy_close&#x60; | [default to ""]
**Tag** | Pointer to **string** | Order tag  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 16 characters. | [optional] [default to ""]
**TpRatio** | Pointer to **string** | Take profit per order. 0.1 represents 10% | [optional] [default to ""]
**UniqueCode** | **string** | Lead trader unique code  A combination of case-sensitive alphanumerics, all numbers and the length is 16 characters, e.g. 213E8C92DC61EFAC | [default to ""]

## Methods

### NewCreateCopytradingFirstCopySettingsV5Req

`func NewCreateCopytradingFirstCopySettingsV5Req(copyInstIdType string, copyMgnMode string, copyTotalAmt string, subPosCloseType string, uniqueCode string, ) *CreateCopytradingFirstCopySettingsV5Req`

NewCreateCopytradingFirstCopySettingsV5Req instantiates a new CreateCopytradingFirstCopySettingsV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCopytradingFirstCopySettingsV5ReqWithDefaults

`func NewCreateCopytradingFirstCopySettingsV5ReqWithDefaults() *CreateCopytradingFirstCopySettingsV5Req`

NewCreateCopytradingFirstCopySettingsV5ReqWithDefaults instantiates a new CreateCopytradingFirstCopySettingsV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyAmt

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetCopyAmt() string`

GetCopyAmt returns the CopyAmt field if non-nil, zero value otherwise.

### GetCopyAmtOk

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetCopyAmtOk() (*string, bool)`

GetCopyAmtOk returns a tuple with the CopyAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyAmt

`func (o *CreateCopytradingFirstCopySettingsV5Req) SetCopyAmt(v string)`

SetCopyAmt sets CopyAmt field to given value.

### HasCopyAmt

`func (o *CreateCopytradingFirstCopySettingsV5Req) HasCopyAmt() bool`

HasCopyAmt returns a boolean if a field has been set.

### GetCopyInstIdType

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetCopyInstIdType() string`

GetCopyInstIdType returns the CopyInstIdType field if non-nil, zero value otherwise.

### GetCopyInstIdTypeOk

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetCopyInstIdTypeOk() (*string, bool)`

GetCopyInstIdTypeOk returns a tuple with the CopyInstIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyInstIdType

`func (o *CreateCopytradingFirstCopySettingsV5Req) SetCopyInstIdType(v string)`

SetCopyInstIdType sets CopyInstIdType field to given value.


### GetCopyMgnMode

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetCopyMgnMode() string`

GetCopyMgnMode returns the CopyMgnMode field if non-nil, zero value otherwise.

### GetCopyMgnModeOk

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetCopyMgnModeOk() (*string, bool)`

GetCopyMgnModeOk returns a tuple with the CopyMgnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyMgnMode

`func (o *CreateCopytradingFirstCopySettingsV5Req) SetCopyMgnMode(v string)`

SetCopyMgnMode sets CopyMgnMode field to given value.


### GetCopyMode

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetCopyMode() string`

GetCopyMode returns the CopyMode field if non-nil, zero value otherwise.

### GetCopyModeOk

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetCopyModeOk() (*string, bool)`

GetCopyModeOk returns a tuple with the CopyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyMode

`func (o *CreateCopytradingFirstCopySettingsV5Req) SetCopyMode(v string)`

SetCopyMode sets CopyMode field to given value.

### HasCopyMode

`func (o *CreateCopytradingFirstCopySettingsV5Req) HasCopyMode() bool`

HasCopyMode returns a boolean if a field has been set.

### GetCopyRatio

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetCopyRatio() string`

GetCopyRatio returns the CopyRatio field if non-nil, zero value otherwise.

### GetCopyRatioOk

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetCopyRatioOk() (*string, bool)`

GetCopyRatioOk returns a tuple with the CopyRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyRatio

`func (o *CreateCopytradingFirstCopySettingsV5Req) SetCopyRatio(v string)`

SetCopyRatio sets CopyRatio field to given value.

### HasCopyRatio

`func (o *CreateCopytradingFirstCopySettingsV5Req) HasCopyRatio() bool`

HasCopyRatio returns a boolean if a field has been set.

### GetCopyTotalAmt

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetCopyTotalAmt() string`

GetCopyTotalAmt returns the CopyTotalAmt field if non-nil, zero value otherwise.

### GetCopyTotalAmtOk

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetCopyTotalAmtOk() (*string, bool)`

GetCopyTotalAmtOk returns a tuple with the CopyTotalAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTotalAmt

`func (o *CreateCopytradingFirstCopySettingsV5Req) SetCopyTotalAmt(v string)`

SetCopyTotalAmt sets CopyTotalAmt field to given value.


### GetInstId

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateCopytradingFirstCopySettingsV5Req) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *CreateCopytradingFirstCopySettingsV5Req) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *CreateCopytradingFirstCopySettingsV5Req) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *CreateCopytradingFirstCopySettingsV5Req) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetSlRatio

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetSlRatio() string`

GetSlRatio returns the SlRatio field if non-nil, zero value otherwise.

### GetSlRatioOk

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetSlRatioOk() (*string, bool)`

GetSlRatioOk returns a tuple with the SlRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlRatio

`func (o *CreateCopytradingFirstCopySettingsV5Req) SetSlRatio(v string)`

SetSlRatio sets SlRatio field to given value.

### HasSlRatio

`func (o *CreateCopytradingFirstCopySettingsV5Req) HasSlRatio() bool`

HasSlRatio returns a boolean if a field has been set.

### GetSlTotalAmt

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetSlTotalAmt() string`

GetSlTotalAmt returns the SlTotalAmt field if non-nil, zero value otherwise.

### GetSlTotalAmtOk

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetSlTotalAmtOk() (*string, bool)`

GetSlTotalAmtOk returns a tuple with the SlTotalAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTotalAmt

`func (o *CreateCopytradingFirstCopySettingsV5Req) SetSlTotalAmt(v string)`

SetSlTotalAmt sets SlTotalAmt field to given value.

### HasSlTotalAmt

`func (o *CreateCopytradingFirstCopySettingsV5Req) HasSlTotalAmt() bool`

HasSlTotalAmt returns a boolean if a field has been set.

### GetSubPosCloseType

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetSubPosCloseType() string`

GetSubPosCloseType returns the SubPosCloseType field if non-nil, zero value otherwise.

### GetSubPosCloseTypeOk

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetSubPosCloseTypeOk() (*string, bool)`

GetSubPosCloseTypeOk returns a tuple with the SubPosCloseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPosCloseType

`func (o *CreateCopytradingFirstCopySettingsV5Req) SetSubPosCloseType(v string)`

SetSubPosCloseType sets SubPosCloseType field to given value.


### GetTag

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateCopytradingFirstCopySettingsV5Req) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateCopytradingFirstCopySettingsV5Req) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTpRatio

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetTpRatio() string`

GetTpRatio returns the TpRatio field if non-nil, zero value otherwise.

### GetTpRatioOk

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetTpRatioOk() (*string, bool)`

GetTpRatioOk returns a tuple with the TpRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpRatio

`func (o *CreateCopytradingFirstCopySettingsV5Req) SetTpRatio(v string)`

SetTpRatio sets TpRatio field to given value.

### HasTpRatio

`func (o *CreateCopytradingFirstCopySettingsV5Req) HasTpRatio() bool`

HasTpRatio returns a boolean if a field has been set.

### GetUniqueCode

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetUniqueCode() string`

GetUniqueCode returns the UniqueCode field if non-nil, zero value otherwise.

### GetUniqueCodeOk

`func (o *CreateCopytradingFirstCopySettingsV5Req) GetUniqueCodeOk() (*string, bool)`

GetUniqueCodeOk returns a tuple with the UniqueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueCode

`func (o *CreateCopytradingFirstCopySettingsV5Req) SetUniqueCode(v string)`

SetUniqueCode sets UniqueCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


