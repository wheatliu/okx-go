# GetAssetTransferStateV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | Pointer to **string** | Amount to be transferred | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency, e.g. &#x60;USDT&#x60; | [optional] [default to ""]
**ClientId** | Pointer to **string** | Client-supplied ID | [optional] [default to ""]
**From** | Pointer to **string** | The remitting account  &#x60;6&#x60;: Funding account  &#x60;18&#x60;: Trading account | [optional] [default to ""]
**InstId** | Pointer to **string** | deprecated | [optional] [default to ""]
**State** | Pointer to **string** | Transfer state  &#x60;success&#x60;  &#x60;pending&#x60;  &#x60;failed&#x60; | [optional] [default to ""]
**SubAcct** | Pointer to **string** | Name of the sub-account | [optional] [default to ""]
**To** | Pointer to **string** | The beneficiary account  &#x60;6&#x60;: Funding account  &#x60;18&#x60;: Trading account | [optional] [default to ""]
**ToInstId** | Pointer to **string** | deprecated | [optional] [default to ""]
**TransId** | Pointer to **string** | Transfer ID | [optional] [default to ""]
**Type** | Pointer to **string** | Transfer type  &#x60;0&#x60;: transfer within account  &#x60;1&#x60;: master account to sub-account (Only applicable to API Key from master account)   &#x60;2&#x60;: sub-account to master account (Only applicable to APIKey from master account)  &#x60;3&#x60;: sub-account to master account (Only applicable to APIKey from sub-account)  &#x60;4&#x60;: sub-account to sub-account (Only applicable to APIKey from sub-account, and target account needs to be another sub-account which belongs to same master account) | [optional] [default to ""]

## Methods

### NewGetAssetTransferStateV5RespDataInner

`func NewGetAssetTransferStateV5RespDataInner() *GetAssetTransferStateV5RespDataInner`

NewGetAssetTransferStateV5RespDataInner instantiates a new GetAssetTransferStateV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetTransferStateV5RespDataInnerWithDefaults

`func NewGetAssetTransferStateV5RespDataInnerWithDefaults() *GetAssetTransferStateV5RespDataInner`

NewGetAssetTransferStateV5RespDataInnerWithDefaults instantiates a new GetAssetTransferStateV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *GetAssetTransferStateV5RespDataInner) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetAssetTransferStateV5RespDataInner) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetAssetTransferStateV5RespDataInner) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetAssetTransferStateV5RespDataInner) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetCcy

`func (o *GetAssetTransferStateV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAssetTransferStateV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAssetTransferStateV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAssetTransferStateV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetClientId

`func (o *GetAssetTransferStateV5RespDataInner) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetAssetTransferStateV5RespDataInner) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetAssetTransferStateV5RespDataInner) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GetAssetTransferStateV5RespDataInner) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetFrom

`func (o *GetAssetTransferStateV5RespDataInner) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetAssetTransferStateV5RespDataInner) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetAssetTransferStateV5RespDataInner) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GetAssetTransferStateV5RespDataInner) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetInstId

`func (o *GetAssetTransferStateV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetAssetTransferStateV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetAssetTransferStateV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetAssetTransferStateV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetState

`func (o *GetAssetTransferStateV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetAssetTransferStateV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetAssetTransferStateV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetAssetTransferStateV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubAcct

`func (o *GetAssetTransferStateV5RespDataInner) GetSubAcct() string`

GetSubAcct returns the SubAcct field if non-nil, zero value otherwise.

### GetSubAcctOk

`func (o *GetAssetTransferStateV5RespDataInner) GetSubAcctOk() (*string, bool)`

GetSubAcctOk returns a tuple with the SubAcct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAcct

`func (o *GetAssetTransferStateV5RespDataInner) SetSubAcct(v string)`

SetSubAcct sets SubAcct field to given value.

### HasSubAcct

`func (o *GetAssetTransferStateV5RespDataInner) HasSubAcct() bool`

HasSubAcct returns a boolean if a field has been set.

### GetTo

`func (o *GetAssetTransferStateV5RespDataInner) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GetAssetTransferStateV5RespDataInner) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GetAssetTransferStateV5RespDataInner) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *GetAssetTransferStateV5RespDataInner) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetToInstId

`func (o *GetAssetTransferStateV5RespDataInner) GetToInstId() string`

GetToInstId returns the ToInstId field if non-nil, zero value otherwise.

### GetToInstIdOk

`func (o *GetAssetTransferStateV5RespDataInner) GetToInstIdOk() (*string, bool)`

GetToInstIdOk returns a tuple with the ToInstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToInstId

`func (o *GetAssetTransferStateV5RespDataInner) SetToInstId(v string)`

SetToInstId sets ToInstId field to given value.

### HasToInstId

`func (o *GetAssetTransferStateV5RespDataInner) HasToInstId() bool`

HasToInstId returns a boolean if a field has been set.

### GetTransId

`func (o *GetAssetTransferStateV5RespDataInner) GetTransId() string`

GetTransId returns the TransId field if non-nil, zero value otherwise.

### GetTransIdOk

`func (o *GetAssetTransferStateV5RespDataInner) GetTransIdOk() (*string, bool)`

GetTransIdOk returns a tuple with the TransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransId

`func (o *GetAssetTransferStateV5RespDataInner) SetTransId(v string)`

SetTransId sets TransId field to given value.

### HasTransId

`func (o *GetAssetTransferStateV5RespDataInner) HasTransId() bool`

HasTransId returns a boolean if a field has been set.

### GetType

`func (o *GetAssetTransferStateV5RespDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAssetTransferStateV5RespDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAssetTransferStateV5RespDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAssetTransferStateV5RespDataInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


