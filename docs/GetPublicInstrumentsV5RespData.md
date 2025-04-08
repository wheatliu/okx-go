# GetPublicInstrumentsV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | Alias  &#x60;this_week&#x60;  &#x60;next_week&#x60;  &#x60;this_month&#x60;  &#x60;next_month&#x60;  &#x60;quarter&#x60;  &#x60;next_quarter&#x60;  &#x60;third_quarter&#x60;  Only applicable to &#x60;FUTURES&#x60;    | [optional] [default to ""]
**AuctionEndTime** | Pointer to **string** | The end time of call auction, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60;   Only applicable to &#x60;SPOT&#x60; that are listed through call auctions, return \&quot;\&quot; in other cases | [optional] [default to ""]
**BaseCcy** | Pointer to **string** | Base currency, e.g. &#x60;BTC&#x60;  in&#x60;BTC-USDT&#x60;   Only applicable to &#x60;SPOT&#x60;/&#x60;MARGIN&#x60; | [optional] [default to ""]
**Category** | Pointer to **string** | Currency category. Note: this parameter is already deprecated | [optional] [default to ""]
**CtMult** | Pointer to **string** | Contract multiplier     Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [optional] [default to ""]
**CtType** | Pointer to **string** | Contract type  &#x60;linear&#x60;: linear contract  &#x60;inverse&#x60;: inverse contract   Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60; | [optional] [default to ""]
**CtVal** | Pointer to **string** | Contract value     Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [optional] [default to ""]
**CtValCcy** | Pointer to **string** | Contract value currency    Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [optional] [default to ""]
**ExpTime** | Pointer to **string** | Expiry time   Applicable to &#x60;SPOT&#x60;/&#x60;MARGIN&#x60;/&#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;. For &#x60;FUTURES&#x60;/&#x60;OPTION&#x60;, it is natural delivery/exercise time. It is the instrument offline time when there is &#x60;SPOT/MARGIN/FUTURES/SWAP/&#x60; manual offline. Update once change. | [optional] [default to ""]
**FutureSettlement** | Pointer to **bool** | Whether daily settlement for expiry feature is enabled  Applicable to &#x60;FUTURES&#x60; &#x60;cross&#x60; | [optional] 
**InstFamily** | Pointer to **string** | Instrument family, e.g. &#x60;BTC-USD&#x60;   Only applicable to &#x60;MARGIN/FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID,  e.g. &#x60;BTC-USD-SWAP&#x60; | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**Lever** | Pointer to **string** | Max Leverage,     Not applicable to &#x60;SPOT&#x60;, &#x60;OPTION&#x60; | [optional] [default to ""]
**ListTime** | Pointer to **string** | Listing time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**LotSz** | Pointer to **string** | Lot size  If it is a derivatives contract, the value is the number of contracts.  If it is &#x60;SPOT&#x60;/&#x60;MARGIN&#x60;, the value is the quantity in &#x60;base currency&#x60;. | [optional] [default to ""]
**MaxIcebergSz** | Pointer to **string** | The maximum order quantity of a single iceBerg order.  If it is a derivatives contract, the value is the number of contracts.  If it is &#x60;SPOT&#x60;/&#x60;MARGIN&#x60;, the value is the quantity in &#x60;base currency&#x60;. | [optional] [default to ""]
**MaxLmtAmt** | Pointer to **string** | Max USD amount for a single limit order | [optional] [default to ""]
**MaxLmtSz** | Pointer to **string** | The maximum order quantity of a single limit order.  If it is a derivatives contract, the value is the number of contracts.  If it is &#x60;SPOT&#x60;/&#x60;MARGIN&#x60;, the value is the quantity in &#x60;base currency&#x60;. | [optional] [default to ""]
**MaxMktAmt** | Pointer to **string** | Max USD amount for a single market order   Only applicable to &#x60;SPOT&#x60;/&#x60;MARGIN&#x60; | [optional] [default to ""]
**MaxMktSz** | Pointer to **string** | The maximum order quantity of a single market order.  If it is a derivatives contract, the value is the number of contracts.  If it is &#x60;SPOT&#x60;/&#x60;MARGIN&#x60;, the value is the quantity in &#x60;USDT&#x60;. | [optional] [default to ""]
**MaxStopSz** | Pointer to **string** | The maximum order quantity of a single stop market order.  If it is a derivatives contract, the value is the number of contracts.  If it is &#x60;SPOT&#x60;/&#x60;MARGIN&#x60;, the value is the quantity in &#x60;USDT&#x60;. | [optional] [default to ""]
**MaxTriggerSz** | Pointer to **string** | The maximum order quantity of a single trigger order.  If it is a derivatives contract, the value is the number of contracts.  If it is &#x60;SPOT&#x60;/&#x60;MARGIN&#x60;, the value is the quantity in &#x60;base currency&#x60;. | [optional] [default to ""]
**MaxTwapSz** | Pointer to **string** | The maximum order quantity of a single TWAP order.  If it is a derivatives contract, the value is the number of contracts.  If it is &#x60;SPOT&#x60;/&#x60;MARGIN&#x60;, the value is the quantity in &#x60;base currency&#x60;.    The minimum order quantity of a single TWAP order is minSz*2 | [optional] [default to ""]
**MinSz** | Pointer to **string** | Minimum order size  If it is a derivatives contract, the value is the number of contracts.  If it is &#x60;SPOT&#x60;/&#x60;MARGIN&#x60;, the value is the quantity in &#x60;base currency&#x60;. | [optional] [default to ""]
**OptType** | Pointer to **string** | Option type, &#x60;C&#x60;: Call  &#x60;P&#x60;: put   Only applicable to &#x60;OPTION&#x60; | [optional] [default to ""]
**QuoteCcy** | Pointer to **string** | Quote currency, e.g. &#x60;USDT&#x60; in &#x60;BTC-USDT&#x60;     Only applicable to &#x60;SPOT&#x60;/&#x60;MARGIN&#x60; | [optional] [default to ""]
**RuleType** | Pointer to **string** | Trading rule types   &#x60;normal&#x60;: normal trading   &#x60;pre_market&#x60;: pre-market trading | [optional] [default to ""]
**SettleCcy** | Pointer to **string** | Settlement and margin currency, e.g. &#x60;BTC&#x60;    Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [optional] [default to ""]
**State** | Pointer to **string** | Instrument status  &#x60;live&#x60;   &#x60;suspend&#x60;  &#x60;preopen&#x60;. e.g. There will be &#x60;preopen&#x60; before the Futures and Options new contracts state is live.  &#x60;test&#x60;: Test pairs, can&#39;t be traded | [optional] [default to ""]
**Stk** | Pointer to **string** | Strike price    Only applicable to &#x60;OPTION&#x60; | [optional] [default to ""]
**TickSz** | Pointer to **string** | Tick size,  e.g. &#x60;0.0001&#x60;  For Option, it is minimum tickSz among tick band, please use \&quot;Get option tick bands\&quot; if you want get option tickBands. | [optional] [default to ""]
**Uly** | Pointer to **string** | Underlying, e.g. &#x60;BTC-USD&#x60;   Only applicable to &#x60;MARGIN/FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [optional] [default to ""]

## Methods

### NewGetPublicInstrumentsV5RespData

`func NewGetPublicInstrumentsV5RespData() *GetPublicInstrumentsV5RespData`

NewGetPublicInstrumentsV5RespData instantiates a new GetPublicInstrumentsV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicInstrumentsV5RespDataWithDefaults

`func NewGetPublicInstrumentsV5RespDataWithDefaults() *GetPublicInstrumentsV5RespData`

NewGetPublicInstrumentsV5RespDataWithDefaults instantiates a new GetPublicInstrumentsV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *GetPublicInstrumentsV5RespData) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *GetPublicInstrumentsV5RespData) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *GetPublicInstrumentsV5RespData) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *GetPublicInstrumentsV5RespData) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetAuctionEndTime

`func (o *GetPublicInstrumentsV5RespData) GetAuctionEndTime() string`

GetAuctionEndTime returns the AuctionEndTime field if non-nil, zero value otherwise.

### GetAuctionEndTimeOk

`func (o *GetPublicInstrumentsV5RespData) GetAuctionEndTimeOk() (*string, bool)`

GetAuctionEndTimeOk returns a tuple with the AuctionEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionEndTime

`func (o *GetPublicInstrumentsV5RespData) SetAuctionEndTime(v string)`

SetAuctionEndTime sets AuctionEndTime field to given value.

### HasAuctionEndTime

`func (o *GetPublicInstrumentsV5RespData) HasAuctionEndTime() bool`

HasAuctionEndTime returns a boolean if a field has been set.

### GetBaseCcy

`func (o *GetPublicInstrumentsV5RespData) GetBaseCcy() string`

GetBaseCcy returns the BaseCcy field if non-nil, zero value otherwise.

### GetBaseCcyOk

`func (o *GetPublicInstrumentsV5RespData) GetBaseCcyOk() (*string, bool)`

GetBaseCcyOk returns a tuple with the BaseCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCcy

`func (o *GetPublicInstrumentsV5RespData) SetBaseCcy(v string)`

SetBaseCcy sets BaseCcy field to given value.

### HasBaseCcy

`func (o *GetPublicInstrumentsV5RespData) HasBaseCcy() bool`

HasBaseCcy returns a boolean if a field has been set.

### GetCategory

`func (o *GetPublicInstrumentsV5RespData) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetPublicInstrumentsV5RespData) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetPublicInstrumentsV5RespData) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetPublicInstrumentsV5RespData) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCtMult

`func (o *GetPublicInstrumentsV5RespData) GetCtMult() string`

GetCtMult returns the CtMult field if non-nil, zero value otherwise.

### GetCtMultOk

`func (o *GetPublicInstrumentsV5RespData) GetCtMultOk() (*string, bool)`

GetCtMultOk returns a tuple with the CtMult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtMult

`func (o *GetPublicInstrumentsV5RespData) SetCtMult(v string)`

SetCtMult sets CtMult field to given value.

### HasCtMult

`func (o *GetPublicInstrumentsV5RespData) HasCtMult() bool`

HasCtMult returns a boolean if a field has been set.

### GetCtType

`func (o *GetPublicInstrumentsV5RespData) GetCtType() string`

GetCtType returns the CtType field if non-nil, zero value otherwise.

### GetCtTypeOk

`func (o *GetPublicInstrumentsV5RespData) GetCtTypeOk() (*string, bool)`

GetCtTypeOk returns a tuple with the CtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtType

`func (o *GetPublicInstrumentsV5RespData) SetCtType(v string)`

SetCtType sets CtType field to given value.

### HasCtType

`func (o *GetPublicInstrumentsV5RespData) HasCtType() bool`

HasCtType returns a boolean if a field has been set.

### GetCtVal

`func (o *GetPublicInstrumentsV5RespData) GetCtVal() string`

GetCtVal returns the CtVal field if non-nil, zero value otherwise.

### GetCtValOk

`func (o *GetPublicInstrumentsV5RespData) GetCtValOk() (*string, bool)`

GetCtValOk returns a tuple with the CtVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtVal

`func (o *GetPublicInstrumentsV5RespData) SetCtVal(v string)`

SetCtVal sets CtVal field to given value.

### HasCtVal

`func (o *GetPublicInstrumentsV5RespData) HasCtVal() bool`

HasCtVal returns a boolean if a field has been set.

### GetCtValCcy

`func (o *GetPublicInstrumentsV5RespData) GetCtValCcy() string`

GetCtValCcy returns the CtValCcy field if non-nil, zero value otherwise.

### GetCtValCcyOk

`func (o *GetPublicInstrumentsV5RespData) GetCtValCcyOk() (*string, bool)`

GetCtValCcyOk returns a tuple with the CtValCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtValCcy

`func (o *GetPublicInstrumentsV5RespData) SetCtValCcy(v string)`

SetCtValCcy sets CtValCcy field to given value.

### HasCtValCcy

`func (o *GetPublicInstrumentsV5RespData) HasCtValCcy() bool`

HasCtValCcy returns a boolean if a field has been set.

### GetExpTime

`func (o *GetPublicInstrumentsV5RespData) GetExpTime() string`

GetExpTime returns the ExpTime field if non-nil, zero value otherwise.

### GetExpTimeOk

`func (o *GetPublicInstrumentsV5RespData) GetExpTimeOk() (*string, bool)`

GetExpTimeOk returns a tuple with the ExpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpTime

`func (o *GetPublicInstrumentsV5RespData) SetExpTime(v string)`

SetExpTime sets ExpTime field to given value.

### HasExpTime

`func (o *GetPublicInstrumentsV5RespData) HasExpTime() bool`

HasExpTime returns a boolean if a field has been set.

### GetFutureSettlement

`func (o *GetPublicInstrumentsV5RespData) GetFutureSettlement() bool`

GetFutureSettlement returns the FutureSettlement field if non-nil, zero value otherwise.

### GetFutureSettlementOk

`func (o *GetPublicInstrumentsV5RespData) GetFutureSettlementOk() (*bool, bool)`

GetFutureSettlementOk returns a tuple with the FutureSettlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureSettlement

`func (o *GetPublicInstrumentsV5RespData) SetFutureSettlement(v bool)`

SetFutureSettlement sets FutureSettlement field to given value.

### HasFutureSettlement

`func (o *GetPublicInstrumentsV5RespData) HasFutureSettlement() bool`

HasFutureSettlement returns a boolean if a field has been set.

### GetInstFamily

`func (o *GetPublicInstrumentsV5RespData) GetInstFamily() string`

GetInstFamily returns the InstFamily field if non-nil, zero value otherwise.

### GetInstFamilyOk

`func (o *GetPublicInstrumentsV5RespData) GetInstFamilyOk() (*string, bool)`

GetInstFamilyOk returns a tuple with the InstFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstFamily

`func (o *GetPublicInstrumentsV5RespData) SetInstFamily(v string)`

SetInstFamily sets InstFamily field to given value.

### HasInstFamily

`func (o *GetPublicInstrumentsV5RespData) HasInstFamily() bool`

HasInstFamily returns a boolean if a field has been set.

### GetInstId

`func (o *GetPublicInstrumentsV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetPublicInstrumentsV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetPublicInstrumentsV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetPublicInstrumentsV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetPublicInstrumentsV5RespData) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetPublicInstrumentsV5RespData) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetPublicInstrumentsV5RespData) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetPublicInstrumentsV5RespData) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetLever

`func (o *GetPublicInstrumentsV5RespData) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetPublicInstrumentsV5RespData) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetPublicInstrumentsV5RespData) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetPublicInstrumentsV5RespData) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetListTime

`func (o *GetPublicInstrumentsV5RespData) GetListTime() string`

GetListTime returns the ListTime field if non-nil, zero value otherwise.

### GetListTimeOk

`func (o *GetPublicInstrumentsV5RespData) GetListTimeOk() (*string, bool)`

GetListTimeOk returns a tuple with the ListTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListTime

`func (o *GetPublicInstrumentsV5RespData) SetListTime(v string)`

SetListTime sets ListTime field to given value.

### HasListTime

`func (o *GetPublicInstrumentsV5RespData) HasListTime() bool`

HasListTime returns a boolean if a field has been set.

### GetLotSz

`func (o *GetPublicInstrumentsV5RespData) GetLotSz() string`

GetLotSz returns the LotSz field if non-nil, zero value otherwise.

### GetLotSzOk

`func (o *GetPublicInstrumentsV5RespData) GetLotSzOk() (*string, bool)`

GetLotSzOk returns a tuple with the LotSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotSz

`func (o *GetPublicInstrumentsV5RespData) SetLotSz(v string)`

SetLotSz sets LotSz field to given value.

### HasLotSz

`func (o *GetPublicInstrumentsV5RespData) HasLotSz() bool`

HasLotSz returns a boolean if a field has been set.

### GetMaxIcebergSz

`func (o *GetPublicInstrumentsV5RespData) GetMaxIcebergSz() string`

GetMaxIcebergSz returns the MaxIcebergSz field if non-nil, zero value otherwise.

### GetMaxIcebergSzOk

`func (o *GetPublicInstrumentsV5RespData) GetMaxIcebergSzOk() (*string, bool)`

GetMaxIcebergSzOk returns a tuple with the MaxIcebergSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIcebergSz

`func (o *GetPublicInstrumentsV5RespData) SetMaxIcebergSz(v string)`

SetMaxIcebergSz sets MaxIcebergSz field to given value.

### HasMaxIcebergSz

`func (o *GetPublicInstrumentsV5RespData) HasMaxIcebergSz() bool`

HasMaxIcebergSz returns a boolean if a field has been set.

### GetMaxLmtAmt

`func (o *GetPublicInstrumentsV5RespData) GetMaxLmtAmt() string`

GetMaxLmtAmt returns the MaxLmtAmt field if non-nil, zero value otherwise.

### GetMaxLmtAmtOk

`func (o *GetPublicInstrumentsV5RespData) GetMaxLmtAmtOk() (*string, bool)`

GetMaxLmtAmtOk returns a tuple with the MaxLmtAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLmtAmt

`func (o *GetPublicInstrumentsV5RespData) SetMaxLmtAmt(v string)`

SetMaxLmtAmt sets MaxLmtAmt field to given value.

### HasMaxLmtAmt

`func (o *GetPublicInstrumentsV5RespData) HasMaxLmtAmt() bool`

HasMaxLmtAmt returns a boolean if a field has been set.

### GetMaxLmtSz

`func (o *GetPublicInstrumentsV5RespData) GetMaxLmtSz() string`

GetMaxLmtSz returns the MaxLmtSz field if non-nil, zero value otherwise.

### GetMaxLmtSzOk

`func (o *GetPublicInstrumentsV5RespData) GetMaxLmtSzOk() (*string, bool)`

GetMaxLmtSzOk returns a tuple with the MaxLmtSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLmtSz

`func (o *GetPublicInstrumentsV5RespData) SetMaxLmtSz(v string)`

SetMaxLmtSz sets MaxLmtSz field to given value.

### HasMaxLmtSz

`func (o *GetPublicInstrumentsV5RespData) HasMaxLmtSz() bool`

HasMaxLmtSz returns a boolean if a field has been set.

### GetMaxMktAmt

`func (o *GetPublicInstrumentsV5RespData) GetMaxMktAmt() string`

GetMaxMktAmt returns the MaxMktAmt field if non-nil, zero value otherwise.

### GetMaxMktAmtOk

`func (o *GetPublicInstrumentsV5RespData) GetMaxMktAmtOk() (*string, bool)`

GetMaxMktAmtOk returns a tuple with the MaxMktAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMktAmt

`func (o *GetPublicInstrumentsV5RespData) SetMaxMktAmt(v string)`

SetMaxMktAmt sets MaxMktAmt field to given value.

### HasMaxMktAmt

`func (o *GetPublicInstrumentsV5RespData) HasMaxMktAmt() bool`

HasMaxMktAmt returns a boolean if a field has been set.

### GetMaxMktSz

`func (o *GetPublicInstrumentsV5RespData) GetMaxMktSz() string`

GetMaxMktSz returns the MaxMktSz field if non-nil, zero value otherwise.

### GetMaxMktSzOk

`func (o *GetPublicInstrumentsV5RespData) GetMaxMktSzOk() (*string, bool)`

GetMaxMktSzOk returns a tuple with the MaxMktSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMktSz

`func (o *GetPublicInstrumentsV5RespData) SetMaxMktSz(v string)`

SetMaxMktSz sets MaxMktSz field to given value.

### HasMaxMktSz

`func (o *GetPublicInstrumentsV5RespData) HasMaxMktSz() bool`

HasMaxMktSz returns a boolean if a field has been set.

### GetMaxStopSz

`func (o *GetPublicInstrumentsV5RespData) GetMaxStopSz() string`

GetMaxStopSz returns the MaxStopSz field if non-nil, zero value otherwise.

### GetMaxStopSzOk

`func (o *GetPublicInstrumentsV5RespData) GetMaxStopSzOk() (*string, bool)`

GetMaxStopSzOk returns a tuple with the MaxStopSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStopSz

`func (o *GetPublicInstrumentsV5RespData) SetMaxStopSz(v string)`

SetMaxStopSz sets MaxStopSz field to given value.

### HasMaxStopSz

`func (o *GetPublicInstrumentsV5RespData) HasMaxStopSz() bool`

HasMaxStopSz returns a boolean if a field has been set.

### GetMaxTriggerSz

`func (o *GetPublicInstrumentsV5RespData) GetMaxTriggerSz() string`

GetMaxTriggerSz returns the MaxTriggerSz field if non-nil, zero value otherwise.

### GetMaxTriggerSzOk

`func (o *GetPublicInstrumentsV5RespData) GetMaxTriggerSzOk() (*string, bool)`

GetMaxTriggerSzOk returns a tuple with the MaxTriggerSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTriggerSz

`func (o *GetPublicInstrumentsV5RespData) SetMaxTriggerSz(v string)`

SetMaxTriggerSz sets MaxTriggerSz field to given value.

### HasMaxTriggerSz

`func (o *GetPublicInstrumentsV5RespData) HasMaxTriggerSz() bool`

HasMaxTriggerSz returns a boolean if a field has been set.

### GetMaxTwapSz

`func (o *GetPublicInstrumentsV5RespData) GetMaxTwapSz() string`

GetMaxTwapSz returns the MaxTwapSz field if non-nil, zero value otherwise.

### GetMaxTwapSzOk

`func (o *GetPublicInstrumentsV5RespData) GetMaxTwapSzOk() (*string, bool)`

GetMaxTwapSzOk returns a tuple with the MaxTwapSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTwapSz

`func (o *GetPublicInstrumentsV5RespData) SetMaxTwapSz(v string)`

SetMaxTwapSz sets MaxTwapSz field to given value.

### HasMaxTwapSz

`func (o *GetPublicInstrumentsV5RespData) HasMaxTwapSz() bool`

HasMaxTwapSz returns a boolean if a field has been set.

### GetMinSz

`func (o *GetPublicInstrumentsV5RespData) GetMinSz() string`

GetMinSz returns the MinSz field if non-nil, zero value otherwise.

### GetMinSzOk

`func (o *GetPublicInstrumentsV5RespData) GetMinSzOk() (*string, bool)`

GetMinSzOk returns a tuple with the MinSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSz

`func (o *GetPublicInstrumentsV5RespData) SetMinSz(v string)`

SetMinSz sets MinSz field to given value.

### HasMinSz

`func (o *GetPublicInstrumentsV5RespData) HasMinSz() bool`

HasMinSz returns a boolean if a field has been set.

### GetOptType

`func (o *GetPublicInstrumentsV5RespData) GetOptType() string`

GetOptType returns the OptType field if non-nil, zero value otherwise.

### GetOptTypeOk

`func (o *GetPublicInstrumentsV5RespData) GetOptTypeOk() (*string, bool)`

GetOptTypeOk returns a tuple with the OptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptType

`func (o *GetPublicInstrumentsV5RespData) SetOptType(v string)`

SetOptType sets OptType field to given value.

### HasOptType

`func (o *GetPublicInstrumentsV5RespData) HasOptType() bool`

HasOptType returns a boolean if a field has been set.

### GetQuoteCcy

`func (o *GetPublicInstrumentsV5RespData) GetQuoteCcy() string`

GetQuoteCcy returns the QuoteCcy field if non-nil, zero value otherwise.

### GetQuoteCcyOk

`func (o *GetPublicInstrumentsV5RespData) GetQuoteCcyOk() (*string, bool)`

GetQuoteCcyOk returns a tuple with the QuoteCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCcy

`func (o *GetPublicInstrumentsV5RespData) SetQuoteCcy(v string)`

SetQuoteCcy sets QuoteCcy field to given value.

### HasQuoteCcy

`func (o *GetPublicInstrumentsV5RespData) HasQuoteCcy() bool`

HasQuoteCcy returns a boolean if a field has been set.

### GetRuleType

`func (o *GetPublicInstrumentsV5RespData) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *GetPublicInstrumentsV5RespData) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *GetPublicInstrumentsV5RespData) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *GetPublicInstrumentsV5RespData) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetSettleCcy

`func (o *GetPublicInstrumentsV5RespData) GetSettleCcy() string`

GetSettleCcy returns the SettleCcy field if non-nil, zero value otherwise.

### GetSettleCcyOk

`func (o *GetPublicInstrumentsV5RespData) GetSettleCcyOk() (*string, bool)`

GetSettleCcyOk returns a tuple with the SettleCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettleCcy

`func (o *GetPublicInstrumentsV5RespData) SetSettleCcy(v string)`

SetSettleCcy sets SettleCcy field to given value.

### HasSettleCcy

`func (o *GetPublicInstrumentsV5RespData) HasSettleCcy() bool`

HasSettleCcy returns a boolean if a field has been set.

### GetState

`func (o *GetPublicInstrumentsV5RespData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetPublicInstrumentsV5RespData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetPublicInstrumentsV5RespData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetPublicInstrumentsV5RespData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStk

`func (o *GetPublicInstrumentsV5RespData) GetStk() string`

GetStk returns the Stk field if non-nil, zero value otherwise.

### GetStkOk

`func (o *GetPublicInstrumentsV5RespData) GetStkOk() (*string, bool)`

GetStkOk returns a tuple with the Stk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStk

`func (o *GetPublicInstrumentsV5RespData) SetStk(v string)`

SetStk sets Stk field to given value.

### HasStk

`func (o *GetPublicInstrumentsV5RespData) HasStk() bool`

HasStk returns a boolean if a field has been set.

### GetTickSz

`func (o *GetPublicInstrumentsV5RespData) GetTickSz() string`

GetTickSz returns the TickSz field if non-nil, zero value otherwise.

### GetTickSzOk

`func (o *GetPublicInstrumentsV5RespData) GetTickSzOk() (*string, bool)`

GetTickSzOk returns a tuple with the TickSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickSz

`func (o *GetPublicInstrumentsV5RespData) SetTickSz(v string)`

SetTickSz sets TickSz field to given value.

### HasTickSz

`func (o *GetPublicInstrumentsV5RespData) HasTickSz() bool`

HasTickSz returns a boolean if a field has been set.

### GetUly

`func (o *GetPublicInstrumentsV5RespData) GetUly() string`

GetUly returns the Uly field if non-nil, zero value otherwise.

### GetUlyOk

`func (o *GetPublicInstrumentsV5RespData) GetUlyOk() (*string, bool)`

GetUlyOk returns a tuple with the Uly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUly

`func (o *GetPublicInstrumentsV5RespData) SetUly(v string)`

SetUly sets Uly field to given value.

### HasUly

`func (o *GetPublicInstrumentsV5RespData) HasUly() bool`

HasUly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


