# CreateAccountPositionBuilderV5ReqSimPosInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgPx** | Pointer to **string** | Average open price | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60;  Applicable to &#x60;SWAP&#x60;/&#x60;FUTURES&#x60;/&#x60;OPTION&#x60; | [optional] [default to ""]
**Lever** | Pointer to **string** | leverage  Only applicable to &#x60;Multi-currency margin&#x60;  The default is &#x60;1&#x60;  If the allowed leverage is exceeded, set according to the maximum leverage. | [optional] [default to ""]
**Pos** | Pointer to **string** | Quantity of positions | [optional] [default to ""]

## Methods

### NewCreateAccountPositionBuilderV5ReqSimPosInner

`func NewCreateAccountPositionBuilderV5ReqSimPosInner() *CreateAccountPositionBuilderV5ReqSimPosInner`

NewCreateAccountPositionBuilderV5ReqSimPosInner instantiates a new CreateAccountPositionBuilderV5ReqSimPosInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountPositionBuilderV5ReqSimPosInnerWithDefaults

`func NewCreateAccountPositionBuilderV5ReqSimPosInnerWithDefaults() *CreateAccountPositionBuilderV5ReqSimPosInner`

NewCreateAccountPositionBuilderV5ReqSimPosInnerWithDefaults instantiates a new CreateAccountPositionBuilderV5ReqSimPosInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgPx

`func (o *CreateAccountPositionBuilderV5ReqSimPosInner) GetAvgPx() string`

GetAvgPx returns the AvgPx field if non-nil, zero value otherwise.

### GetAvgPxOk

`func (o *CreateAccountPositionBuilderV5ReqSimPosInner) GetAvgPxOk() (*string, bool)`

GetAvgPxOk returns a tuple with the AvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPx

`func (o *CreateAccountPositionBuilderV5ReqSimPosInner) SetAvgPx(v string)`

SetAvgPx sets AvgPx field to given value.

### HasAvgPx

`func (o *CreateAccountPositionBuilderV5ReqSimPosInner) HasAvgPx() bool`

HasAvgPx returns a boolean if a field has been set.

### GetInstId

`func (o *CreateAccountPositionBuilderV5ReqSimPosInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateAccountPositionBuilderV5ReqSimPosInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateAccountPositionBuilderV5ReqSimPosInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *CreateAccountPositionBuilderV5ReqSimPosInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetLever

`func (o *CreateAccountPositionBuilderV5ReqSimPosInner) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *CreateAccountPositionBuilderV5ReqSimPosInner) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *CreateAccountPositionBuilderV5ReqSimPosInner) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *CreateAccountPositionBuilderV5ReqSimPosInner) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetPos

`func (o *CreateAccountPositionBuilderV5ReqSimPosInner) GetPos() string`

GetPos returns the Pos field if non-nil, zero value otherwise.

### GetPosOk

`func (o *CreateAccountPositionBuilderV5ReqSimPosInner) GetPosOk() (*string, bool)`

GetPosOk returns a tuple with the Pos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPos

`func (o *CreateAccountPositionBuilderV5ReqSimPosInner) SetPos(v string)`

SetPos sets Pos field to given value.

### HasPos

`func (o *CreateAccountPositionBuilderV5ReqSimPosInner) HasPos() bool`

HasPos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


