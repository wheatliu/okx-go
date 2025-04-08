# GetAccountSubaccountMaxWithdrawalV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | Pointer to **string** | Currency | [optional] [default to ""]
**MaxWd** | Pointer to **string** | Max withdrawal (excluding borrowed assets under &#x60;Multi-currency margin&#x60;) | [optional] [default to ""]
**MaxWdEx** | Pointer to **string** | Max withdrawal (including borrowed assets under &#x60;Multi-currency margin&#x60;) | [optional] [default to ""]
**SpotOffsetMaxWd** | Pointer to **string** | Max withdrawal under Spot-Derivatives risk offset mode (excluding borrowed assets under &#x60;Portfolio margin&#x60;)   Applicable to &#x60;Portfolio margin&#x60; | [optional] [default to ""]
**SpotOffsetMaxWdEx** | Pointer to **string** | Max withdrawal under Spot-Derivatives risk offset mode (including borrowed assets under &#x60;Portfolio margin&#x60;)   Applicable to &#x60;Portfolio margin&#x60; | [optional] [default to ""]

## Methods

### NewGetAccountSubaccountMaxWithdrawalV5RespDataInner

`func NewGetAccountSubaccountMaxWithdrawalV5RespDataInner() *GetAccountSubaccountMaxWithdrawalV5RespDataInner`

NewGetAccountSubaccountMaxWithdrawalV5RespDataInner instantiates a new GetAccountSubaccountMaxWithdrawalV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountSubaccountMaxWithdrawalV5RespDataInnerWithDefaults

`func NewGetAccountSubaccountMaxWithdrawalV5RespDataInnerWithDefaults() *GetAccountSubaccountMaxWithdrawalV5RespDataInner`

NewGetAccountSubaccountMaxWithdrawalV5RespDataInnerWithDefaults instantiates a new GetAccountSubaccountMaxWithdrawalV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *GetAccountSubaccountMaxWithdrawalV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAccountSubaccountMaxWithdrawalV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAccountSubaccountMaxWithdrawalV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAccountSubaccountMaxWithdrawalV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetMaxWd

`func (o *GetAccountSubaccountMaxWithdrawalV5RespDataInner) GetMaxWd() string`

GetMaxWd returns the MaxWd field if non-nil, zero value otherwise.

### GetMaxWdOk

`func (o *GetAccountSubaccountMaxWithdrawalV5RespDataInner) GetMaxWdOk() (*string, bool)`

GetMaxWdOk returns a tuple with the MaxWd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWd

`func (o *GetAccountSubaccountMaxWithdrawalV5RespDataInner) SetMaxWd(v string)`

SetMaxWd sets MaxWd field to given value.

### HasMaxWd

`func (o *GetAccountSubaccountMaxWithdrawalV5RespDataInner) HasMaxWd() bool`

HasMaxWd returns a boolean if a field has been set.

### GetMaxWdEx

`func (o *GetAccountSubaccountMaxWithdrawalV5RespDataInner) GetMaxWdEx() string`

GetMaxWdEx returns the MaxWdEx field if non-nil, zero value otherwise.

### GetMaxWdExOk

`func (o *GetAccountSubaccountMaxWithdrawalV5RespDataInner) GetMaxWdExOk() (*string, bool)`

GetMaxWdExOk returns a tuple with the MaxWdEx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWdEx

`func (o *GetAccountSubaccountMaxWithdrawalV5RespDataInner) SetMaxWdEx(v string)`

SetMaxWdEx sets MaxWdEx field to given value.

### HasMaxWdEx

`func (o *GetAccountSubaccountMaxWithdrawalV5RespDataInner) HasMaxWdEx() bool`

HasMaxWdEx returns a boolean if a field has been set.

### GetSpotOffsetMaxWd

`func (o *GetAccountSubaccountMaxWithdrawalV5RespDataInner) GetSpotOffsetMaxWd() string`

GetSpotOffsetMaxWd returns the SpotOffsetMaxWd field if non-nil, zero value otherwise.

### GetSpotOffsetMaxWdOk

`func (o *GetAccountSubaccountMaxWithdrawalV5RespDataInner) GetSpotOffsetMaxWdOk() (*string, bool)`

GetSpotOffsetMaxWdOk returns a tuple with the SpotOffsetMaxWd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotOffsetMaxWd

`func (o *GetAccountSubaccountMaxWithdrawalV5RespDataInner) SetSpotOffsetMaxWd(v string)`

SetSpotOffsetMaxWd sets SpotOffsetMaxWd field to given value.

### HasSpotOffsetMaxWd

`func (o *GetAccountSubaccountMaxWithdrawalV5RespDataInner) HasSpotOffsetMaxWd() bool`

HasSpotOffsetMaxWd returns a boolean if a field has been set.

### GetSpotOffsetMaxWdEx

`func (o *GetAccountSubaccountMaxWithdrawalV5RespDataInner) GetSpotOffsetMaxWdEx() string`

GetSpotOffsetMaxWdEx returns the SpotOffsetMaxWdEx field if non-nil, zero value otherwise.

### GetSpotOffsetMaxWdExOk

`func (o *GetAccountSubaccountMaxWithdrawalV5RespDataInner) GetSpotOffsetMaxWdExOk() (*string, bool)`

GetSpotOffsetMaxWdExOk returns a tuple with the SpotOffsetMaxWdEx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotOffsetMaxWdEx

`func (o *GetAccountSubaccountMaxWithdrawalV5RespDataInner) SetSpotOffsetMaxWdEx(v string)`

SetSpotOffsetMaxWdEx sets SpotOffsetMaxWdEx field to given value.

### HasSpotOffsetMaxWdEx

`func (o *GetAccountSubaccountMaxWithdrawalV5RespDataInner) HasSpotOffsetMaxWdEx() bool`

HasSpotOffsetMaxWdEx returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


