# GetAccountLeverageInfoV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | Pointer to **string** | Currency，used for getting leverage of currency level.  Applicable to &#x60;cross&#x60; &#x60;MARGIN&#x60; of &#x60;Spot mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60;. | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**Lever** | Pointer to **string** | Leverage | [optional] [default to ""]
**MgnMode** | Pointer to **string** | Margin mode | [optional] [default to ""]
**PosSide** | Pointer to **string** | Position side  &#x60;long&#x60;   &#x60;short&#x60;   &#x60;net&#x60;  In &#x60;long/short&#x60; mode, the leverage in both directions &#x60;long&#x60;/&#x60;short&#x60; will be returned. | [optional] [default to ""]

## Methods

### NewGetAccountLeverageInfoV5RespData

`func NewGetAccountLeverageInfoV5RespData() *GetAccountLeverageInfoV5RespData`

NewGetAccountLeverageInfoV5RespData instantiates a new GetAccountLeverageInfoV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountLeverageInfoV5RespDataWithDefaults

`func NewGetAccountLeverageInfoV5RespDataWithDefaults() *GetAccountLeverageInfoV5RespData`

NewGetAccountLeverageInfoV5RespDataWithDefaults instantiates a new GetAccountLeverageInfoV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *GetAccountLeverageInfoV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAccountLeverageInfoV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAccountLeverageInfoV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAccountLeverageInfoV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetInstId

`func (o *GetAccountLeverageInfoV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetAccountLeverageInfoV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetAccountLeverageInfoV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetAccountLeverageInfoV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetLever

`func (o *GetAccountLeverageInfoV5RespData) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetAccountLeverageInfoV5RespData) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetAccountLeverageInfoV5RespData) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetAccountLeverageInfoV5RespData) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetMgnMode

`func (o *GetAccountLeverageInfoV5RespData) GetMgnMode() string`

GetMgnMode returns the MgnMode field if non-nil, zero value otherwise.

### GetMgnModeOk

`func (o *GetAccountLeverageInfoV5RespData) GetMgnModeOk() (*string, bool)`

GetMgnModeOk returns a tuple with the MgnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnMode

`func (o *GetAccountLeverageInfoV5RespData) SetMgnMode(v string)`

SetMgnMode sets MgnMode field to given value.

### HasMgnMode

`func (o *GetAccountLeverageInfoV5RespData) HasMgnMode() bool`

HasMgnMode returns a boolean if a field has been set.

### GetPosSide

`func (o *GetAccountLeverageInfoV5RespData) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *GetAccountLeverageInfoV5RespData) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *GetAccountLeverageInfoV5RespData) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *GetAccountLeverageInfoV5RespData) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


