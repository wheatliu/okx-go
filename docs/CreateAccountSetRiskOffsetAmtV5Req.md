# CreateAccountSetRiskOffsetAmtV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | **string** | Currency | [default to ""]
**ClSpotInUseAmt** | **string** | Spot risk offset amount defined by users | [default to ""]

## Methods

### NewCreateAccountSetRiskOffsetAmtV5Req

`func NewCreateAccountSetRiskOffsetAmtV5Req(ccy string, clSpotInUseAmt string, ) *CreateAccountSetRiskOffsetAmtV5Req`

NewCreateAccountSetRiskOffsetAmtV5Req instantiates a new CreateAccountSetRiskOffsetAmtV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountSetRiskOffsetAmtV5ReqWithDefaults

`func NewCreateAccountSetRiskOffsetAmtV5ReqWithDefaults() *CreateAccountSetRiskOffsetAmtV5Req`

NewCreateAccountSetRiskOffsetAmtV5ReqWithDefaults instantiates a new CreateAccountSetRiskOffsetAmtV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *CreateAccountSetRiskOffsetAmtV5Req) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateAccountSetRiskOffsetAmtV5Req) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateAccountSetRiskOffsetAmtV5Req) SetCcy(v string)`

SetCcy sets Ccy field to given value.


### GetClSpotInUseAmt

`func (o *CreateAccountSetRiskOffsetAmtV5Req) GetClSpotInUseAmt() string`

GetClSpotInUseAmt returns the ClSpotInUseAmt field if non-nil, zero value otherwise.

### GetClSpotInUseAmtOk

`func (o *CreateAccountSetRiskOffsetAmtV5Req) GetClSpotInUseAmtOk() (*string, bool)`

GetClSpotInUseAmtOk returns a tuple with the ClSpotInUseAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClSpotInUseAmt

`func (o *CreateAccountSetRiskOffsetAmtV5Req) SetClSpotInUseAmt(v string)`

SetClSpotInUseAmt sets ClSpotInUseAmt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


