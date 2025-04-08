# CreateAssetTransferV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | **string** | Amount to be transferred | [default to ""]
**Ccy** | **string** | Transfer currency, e.g. &#x60;USDT&#x60; | [default to ""]
**ClientId** | Pointer to **string** | Client-supplied ID  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. | [optional] [default to ""]
**From** | **string** | The remitting account  &#x60;6&#x60;: Funding account  &#x60;18&#x60;: Trading account | [default to ""]
**LoanTrans** | Pointer to **bool** | Whether or not borrowed coins can be transferred out under &#x60;Spot mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60;  &#x60;true&#x60;: borrowed coins can be transferred out  &#x60;false&#x60;: borrowed coins cannot be transferred out  the default is &#x60;false&#x60; | [optional] 
**OmitPosRisk** | Pointer to **string** | Ignore position risk  Default is &#x60;false&#x60;  Applicable to &#x60;Portfolio margin&#x60; | [optional] [default to ""]
**SubAcct** | Pointer to **string** | Name of the sub-account  When &#x60;type&#x60; is &#x60;1&#x60;/&#x60;2&#x60;/&#x60;4&#x60;, this parameter is required. | [optional] [default to ""]
**To** | **string** | The beneficiary account  &#x60;6&#x60;: Funding account  &#x60;18&#x60;: Trading account | [default to ""]
**Type** | Pointer to **string** | Transfer type  &#x60;0&#x60;: transfer within account  &#x60;1&#x60;: master account to sub-account (Only applicable to API Key from master account)  &#x60;2&#x60;: sub-account to master account (Only applicable to API Key from master account)  &#x60;3&#x60;: sub-account to master account (Only applicable to APIKey from sub-account)  &#x60;4&#x60;: sub-account to sub-account (Only applicable to APIKey from sub-account, and target account needs to be another sub-account which belongs to same master account. Sub-account directly transfer out permission is disabled by default, set permission please refer to )  The default is &#x60;0&#x60;.  If you want to make transfer between sub-accounts by master account API key, refer to  | [optional] [default to ""]

## Methods

### NewCreateAssetTransferV5Req

`func NewCreateAssetTransferV5Req(amt string, ccy string, from string, to string, ) *CreateAssetTransferV5Req`

NewCreateAssetTransferV5Req instantiates a new CreateAssetTransferV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAssetTransferV5ReqWithDefaults

`func NewCreateAssetTransferV5ReqWithDefaults() *CreateAssetTransferV5Req`

NewCreateAssetTransferV5ReqWithDefaults instantiates a new CreateAssetTransferV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *CreateAssetTransferV5Req) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *CreateAssetTransferV5Req) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *CreateAssetTransferV5Req) SetAmt(v string)`

SetAmt sets Amt field to given value.


### GetCcy

`func (o *CreateAssetTransferV5Req) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateAssetTransferV5Req) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateAssetTransferV5Req) SetCcy(v string)`

SetCcy sets Ccy field to given value.


### GetClientId

`func (o *CreateAssetTransferV5Req) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateAssetTransferV5Req) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateAssetTransferV5Req) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CreateAssetTransferV5Req) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetFrom

`func (o *CreateAssetTransferV5Req) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CreateAssetTransferV5Req) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CreateAssetTransferV5Req) SetFrom(v string)`

SetFrom sets From field to given value.


### GetLoanTrans

`func (o *CreateAssetTransferV5Req) GetLoanTrans() bool`

GetLoanTrans returns the LoanTrans field if non-nil, zero value otherwise.

### GetLoanTransOk

`func (o *CreateAssetTransferV5Req) GetLoanTransOk() (*bool, bool)`

GetLoanTransOk returns a tuple with the LoanTrans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanTrans

`func (o *CreateAssetTransferV5Req) SetLoanTrans(v bool)`

SetLoanTrans sets LoanTrans field to given value.

### HasLoanTrans

`func (o *CreateAssetTransferV5Req) HasLoanTrans() bool`

HasLoanTrans returns a boolean if a field has been set.

### GetOmitPosRisk

`func (o *CreateAssetTransferV5Req) GetOmitPosRisk() string`

GetOmitPosRisk returns the OmitPosRisk field if non-nil, zero value otherwise.

### GetOmitPosRiskOk

`func (o *CreateAssetTransferV5Req) GetOmitPosRiskOk() (*string, bool)`

GetOmitPosRiskOk returns a tuple with the OmitPosRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitPosRisk

`func (o *CreateAssetTransferV5Req) SetOmitPosRisk(v string)`

SetOmitPosRisk sets OmitPosRisk field to given value.

### HasOmitPosRisk

`func (o *CreateAssetTransferV5Req) HasOmitPosRisk() bool`

HasOmitPosRisk returns a boolean if a field has been set.

### GetSubAcct

`func (o *CreateAssetTransferV5Req) GetSubAcct() string`

GetSubAcct returns the SubAcct field if non-nil, zero value otherwise.

### GetSubAcctOk

`func (o *CreateAssetTransferV5Req) GetSubAcctOk() (*string, bool)`

GetSubAcctOk returns a tuple with the SubAcct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAcct

`func (o *CreateAssetTransferV5Req) SetSubAcct(v string)`

SetSubAcct sets SubAcct field to given value.

### HasSubAcct

`func (o *CreateAssetTransferV5Req) HasSubAcct() bool`

HasSubAcct returns a boolean if a field has been set.

### GetTo

`func (o *CreateAssetTransferV5Req) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CreateAssetTransferV5Req) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CreateAssetTransferV5Req) SetTo(v string)`

SetTo sets To field to given value.


### GetType

`func (o *CreateAssetTransferV5Req) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateAssetTransferV5Req) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateAssetTransferV5Req) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateAssetTransferV5Req) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


