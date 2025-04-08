# CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolSame** | Pointer to **map[string]interface{}** | When the volatility keep the same, the P&amp;L of stress tests under different price volatility ratios, format in {&#x60;change&#x60;: &#x60;value&#x60;,...}  &#x60;change&#x60;: price volatility ratio (in percentage), e.g. &#x60;0.01&#x60; representing &#x60;1%&#x60;  &#x60;value&#x60;: P&amp;L under stress tests, measured in &#x60;USD&#x60;  e.g. {\&quot;-0.15\&quot;:\&quot;-2333.23\&quot;, ...} | [optional] 
**VolShockDown** | Pointer to **map[string]interface{}** | When the volatility shock down, the P&amp;L of stress tests under different price volatility ratios, format in {&#x60;change&#x60;: &#x60;value&#x60;,...}  &#x60;change&#x60;: price volatility ratio (in percentage), e.g. &#x60;0.01&#x60; representing &#x60;1%&#x60;  &#x60;value&#x60;: P&amp;L under stress tests, measured in &#x60;USD&#x60;  e.g. {\&quot;-0.15\&quot;:\&quot;-2333.23\&quot;, ...} | [optional] 
**VolShockUp** | Pointer to **map[string]interface{}** | When the volatility shock up, the P&amp;L of stress tests under different price volatility ratios, format in {&#x60;change&#x60;: &#x60;value&#x60;,...}  &#x60;change&#x60;: price volatility ratio (in percentage), e.g. &#x60;0.01&#x60; representing &#x60;1%&#x60;  &#x60;value&#x60;: P&amp;L under stress tests, measured in &#x60;USD&#x60;  e.g. {\&quot;-0.15\&quot;:\&quot;-2333.23\&quot;, ...} | [optional] 

## Methods

### NewCreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios

`func NewCreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios() *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios`

NewCreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios instantiates a new CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1ScenariosWithDefaults

`func NewCreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1ScenariosWithDefaults() *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios`

NewCreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1ScenariosWithDefaults instantiates a new CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolSame

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios) GetVolSame() map[string]interface{}`

GetVolSame returns the VolSame field if non-nil, zero value otherwise.

### GetVolSameOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios) GetVolSameOk() (*map[string]interface{}, bool)`

GetVolSameOk returns a tuple with the VolSame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolSame

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios) SetVolSame(v map[string]interface{})`

SetVolSame sets VolSame field to given value.

### HasVolSame

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios) HasVolSame() bool`

HasVolSame returns a boolean if a field has been set.

### GetVolShockDown

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios) GetVolShockDown() map[string]interface{}`

GetVolShockDown returns the VolShockDown field if non-nil, zero value otherwise.

### GetVolShockDownOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios) GetVolShockDownOk() (*map[string]interface{}, bool)`

GetVolShockDownOk returns a tuple with the VolShockDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolShockDown

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios) SetVolShockDown(v map[string]interface{})`

SetVolShockDown sets VolShockDown field to given value.

### HasVolShockDown

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios) HasVolShockDown() bool`

HasVolShockDown returns a boolean if a field has been set.

### GetVolShockUp

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios) GetVolShockUp() map[string]interface{}`

GetVolShockUp returns the VolShockUp field if non-nil, zero value otherwise.

### GetVolShockUpOk

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios) GetVolShockUpOk() (*map[string]interface{}, bool)`

GetVolShockUpOk returns a tuple with the VolShockUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolShockUp

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios) SetVolShockUp(v map[string]interface{})`

SetVolShockUp sets VolShockUp field to given value.

### HasVolShockUp

`func (o *CreateAccountPositionBuilderV5RespDataInnerRiskUnitDataInnerMr1Scenarios) HasVolShockUp() bool`

HasVolShockUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


