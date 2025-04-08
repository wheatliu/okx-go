# GetPublicInstrumentsV5RespDataInner

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

### NewGetPublicInstrumentsV5RespDataInner

`func NewGetPublicInstrumentsV5RespDataInner() *GetPublicInstrumentsV5RespDataInner`

NewGetPublicInstrumentsV5RespDataInner instantiates a new GetPublicInstrumentsV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicInstrumentsV5RespDataInnerWithDefaults

`func NewGetPublicInstrumentsV5RespDataInnerWithDefaults() *GetPublicInstrumentsV5RespDataInner`

NewGetPublicInstrumentsV5RespDataInnerWithDefaults instantiates a new GetPublicInstrumentsV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *GetPublicInstrumentsV5RespDataInner) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *GetPublicInstrumentsV5RespDataInner) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *GetPublicInstrumentsV5RespDataInner) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetAuctionEndTime

`func (o *GetPublicInstrumentsV5RespDataInner) GetAuctionEndTime() string`

GetAuctionEndTime returns the AuctionEndTime field if non-nil, zero value otherwise.

### GetAuctionEndTimeOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetAuctionEndTimeOk() (*string, bool)`

GetAuctionEndTimeOk returns a tuple with the AuctionEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionEndTime

`func (o *GetPublicInstrumentsV5RespDataInner) SetAuctionEndTime(v string)`

SetAuctionEndTime sets AuctionEndTime field to given value.

### HasAuctionEndTime

`func (o *GetPublicInstrumentsV5RespDataInner) HasAuctionEndTime() bool`

HasAuctionEndTime returns a boolean if a field has been set.

### GetBaseCcy

`func (o *GetPublicInstrumentsV5RespDataInner) GetBaseCcy() string`

GetBaseCcy returns the BaseCcy field if non-nil, zero value otherwise.

### GetBaseCcyOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetBaseCcyOk() (*string, bool)`

GetBaseCcyOk returns a tuple with the BaseCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCcy

`func (o *GetPublicInstrumentsV5RespDataInner) SetBaseCcy(v string)`

SetBaseCcy sets BaseCcy field to given value.

### HasBaseCcy

`func (o *GetPublicInstrumentsV5RespDataInner) HasBaseCcy() bool`

HasBaseCcy returns a boolean if a field has been set.

### GetCategory

`func (o *GetPublicInstrumentsV5RespDataInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetPublicInstrumentsV5RespDataInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetPublicInstrumentsV5RespDataInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCtMult

`func (o *GetPublicInstrumentsV5RespDataInner) GetCtMult() string`

GetCtMult returns the CtMult field if non-nil, zero value otherwise.

### GetCtMultOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetCtMultOk() (*string, bool)`

GetCtMultOk returns a tuple with the CtMult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtMult

`func (o *GetPublicInstrumentsV5RespDataInner) SetCtMult(v string)`

SetCtMult sets CtMult field to given value.

### HasCtMult

`func (o *GetPublicInstrumentsV5RespDataInner) HasCtMult() bool`

HasCtMult returns a boolean if a field has been set.

### GetCtType

`func (o *GetPublicInstrumentsV5RespDataInner) GetCtType() string`

GetCtType returns the CtType field if non-nil, zero value otherwise.

### GetCtTypeOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetCtTypeOk() (*string, bool)`

GetCtTypeOk returns a tuple with the CtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtType

`func (o *GetPublicInstrumentsV5RespDataInner) SetCtType(v string)`

SetCtType sets CtType field to given value.

### HasCtType

`func (o *GetPublicInstrumentsV5RespDataInner) HasCtType() bool`

HasCtType returns a boolean if a field has been set.

### GetCtVal

`func (o *GetPublicInstrumentsV5RespDataInner) GetCtVal() string`

GetCtVal returns the CtVal field if non-nil, zero value otherwise.

### GetCtValOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetCtValOk() (*string, bool)`

GetCtValOk returns a tuple with the CtVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtVal

`func (o *GetPublicInstrumentsV5RespDataInner) SetCtVal(v string)`

SetCtVal sets CtVal field to given value.

### HasCtVal

`func (o *GetPublicInstrumentsV5RespDataInner) HasCtVal() bool`

HasCtVal returns a boolean if a field has been set.

### GetCtValCcy

`func (o *GetPublicInstrumentsV5RespDataInner) GetCtValCcy() string`

GetCtValCcy returns the CtValCcy field if non-nil, zero value otherwise.

### GetCtValCcyOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetCtValCcyOk() (*string, bool)`

GetCtValCcyOk returns a tuple with the CtValCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtValCcy

`func (o *GetPublicInstrumentsV5RespDataInner) SetCtValCcy(v string)`

SetCtValCcy sets CtValCcy field to given value.

### HasCtValCcy

`func (o *GetPublicInstrumentsV5RespDataInner) HasCtValCcy() bool`

HasCtValCcy returns a boolean if a field has been set.

### GetExpTime

`func (o *GetPublicInstrumentsV5RespDataInner) GetExpTime() string`

GetExpTime returns the ExpTime field if non-nil, zero value otherwise.

### GetExpTimeOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetExpTimeOk() (*string, bool)`

GetExpTimeOk returns a tuple with the ExpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpTime

`func (o *GetPublicInstrumentsV5RespDataInner) SetExpTime(v string)`

SetExpTime sets ExpTime field to given value.

### HasExpTime

`func (o *GetPublicInstrumentsV5RespDataInner) HasExpTime() bool`

HasExpTime returns a boolean if a field has been set.

### GetFutureSettlement

`func (o *GetPublicInstrumentsV5RespDataInner) GetFutureSettlement() bool`

GetFutureSettlement returns the FutureSettlement field if non-nil, zero value otherwise.

### GetFutureSettlementOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetFutureSettlementOk() (*bool, bool)`

GetFutureSettlementOk returns a tuple with the FutureSettlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureSettlement

`func (o *GetPublicInstrumentsV5RespDataInner) SetFutureSettlement(v bool)`

SetFutureSettlement sets FutureSettlement field to given value.

### HasFutureSettlement

`func (o *GetPublicInstrumentsV5RespDataInner) HasFutureSettlement() bool`

HasFutureSettlement returns a boolean if a field has been set.

### GetInstFamily

`func (o *GetPublicInstrumentsV5RespDataInner) GetInstFamily() string`

GetInstFamily returns the InstFamily field if non-nil, zero value otherwise.

### GetInstFamilyOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetInstFamilyOk() (*string, bool)`

GetInstFamilyOk returns a tuple with the InstFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstFamily

`func (o *GetPublicInstrumentsV5RespDataInner) SetInstFamily(v string)`

SetInstFamily sets InstFamily field to given value.

### HasInstFamily

`func (o *GetPublicInstrumentsV5RespDataInner) HasInstFamily() bool`

HasInstFamily returns a boolean if a field has been set.

### GetInstId

`func (o *GetPublicInstrumentsV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetPublicInstrumentsV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetPublicInstrumentsV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetPublicInstrumentsV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetPublicInstrumentsV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetPublicInstrumentsV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetLever

`func (o *GetPublicInstrumentsV5RespDataInner) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetPublicInstrumentsV5RespDataInner) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetPublicInstrumentsV5RespDataInner) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetListTime

`func (o *GetPublicInstrumentsV5RespDataInner) GetListTime() string`

GetListTime returns the ListTime field if non-nil, zero value otherwise.

### GetListTimeOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetListTimeOk() (*string, bool)`

GetListTimeOk returns a tuple with the ListTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListTime

`func (o *GetPublicInstrumentsV5RespDataInner) SetListTime(v string)`

SetListTime sets ListTime field to given value.

### HasListTime

`func (o *GetPublicInstrumentsV5RespDataInner) HasListTime() bool`

HasListTime returns a boolean if a field has been set.

### GetLotSz

`func (o *GetPublicInstrumentsV5RespDataInner) GetLotSz() string`

GetLotSz returns the LotSz field if non-nil, zero value otherwise.

### GetLotSzOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetLotSzOk() (*string, bool)`

GetLotSzOk returns a tuple with the LotSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotSz

`func (o *GetPublicInstrumentsV5RespDataInner) SetLotSz(v string)`

SetLotSz sets LotSz field to given value.

### HasLotSz

`func (o *GetPublicInstrumentsV5RespDataInner) HasLotSz() bool`

HasLotSz returns a boolean if a field has been set.

### GetMaxIcebergSz

`func (o *GetPublicInstrumentsV5RespDataInner) GetMaxIcebergSz() string`

GetMaxIcebergSz returns the MaxIcebergSz field if non-nil, zero value otherwise.

### GetMaxIcebergSzOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetMaxIcebergSzOk() (*string, bool)`

GetMaxIcebergSzOk returns a tuple with the MaxIcebergSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIcebergSz

`func (o *GetPublicInstrumentsV5RespDataInner) SetMaxIcebergSz(v string)`

SetMaxIcebergSz sets MaxIcebergSz field to given value.

### HasMaxIcebergSz

`func (o *GetPublicInstrumentsV5RespDataInner) HasMaxIcebergSz() bool`

HasMaxIcebergSz returns a boolean if a field has been set.

### GetMaxLmtAmt

`func (o *GetPublicInstrumentsV5RespDataInner) GetMaxLmtAmt() string`

GetMaxLmtAmt returns the MaxLmtAmt field if non-nil, zero value otherwise.

### GetMaxLmtAmtOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetMaxLmtAmtOk() (*string, bool)`

GetMaxLmtAmtOk returns a tuple with the MaxLmtAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLmtAmt

`func (o *GetPublicInstrumentsV5RespDataInner) SetMaxLmtAmt(v string)`

SetMaxLmtAmt sets MaxLmtAmt field to given value.

### HasMaxLmtAmt

`func (o *GetPublicInstrumentsV5RespDataInner) HasMaxLmtAmt() bool`

HasMaxLmtAmt returns a boolean if a field has been set.

### GetMaxLmtSz

`func (o *GetPublicInstrumentsV5RespDataInner) GetMaxLmtSz() string`

GetMaxLmtSz returns the MaxLmtSz field if non-nil, zero value otherwise.

### GetMaxLmtSzOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetMaxLmtSzOk() (*string, bool)`

GetMaxLmtSzOk returns a tuple with the MaxLmtSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLmtSz

`func (o *GetPublicInstrumentsV5RespDataInner) SetMaxLmtSz(v string)`

SetMaxLmtSz sets MaxLmtSz field to given value.

### HasMaxLmtSz

`func (o *GetPublicInstrumentsV5RespDataInner) HasMaxLmtSz() bool`

HasMaxLmtSz returns a boolean if a field has been set.

### GetMaxMktAmt

`func (o *GetPublicInstrumentsV5RespDataInner) GetMaxMktAmt() string`

GetMaxMktAmt returns the MaxMktAmt field if non-nil, zero value otherwise.

### GetMaxMktAmtOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetMaxMktAmtOk() (*string, bool)`

GetMaxMktAmtOk returns a tuple with the MaxMktAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMktAmt

`func (o *GetPublicInstrumentsV5RespDataInner) SetMaxMktAmt(v string)`

SetMaxMktAmt sets MaxMktAmt field to given value.

### HasMaxMktAmt

`func (o *GetPublicInstrumentsV5RespDataInner) HasMaxMktAmt() bool`

HasMaxMktAmt returns a boolean if a field has been set.

### GetMaxMktSz

`func (o *GetPublicInstrumentsV5RespDataInner) GetMaxMktSz() string`

GetMaxMktSz returns the MaxMktSz field if non-nil, zero value otherwise.

### GetMaxMktSzOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetMaxMktSzOk() (*string, bool)`

GetMaxMktSzOk returns a tuple with the MaxMktSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMktSz

`func (o *GetPublicInstrumentsV5RespDataInner) SetMaxMktSz(v string)`

SetMaxMktSz sets MaxMktSz field to given value.

### HasMaxMktSz

`func (o *GetPublicInstrumentsV5RespDataInner) HasMaxMktSz() bool`

HasMaxMktSz returns a boolean if a field has been set.

### GetMaxStopSz

`func (o *GetPublicInstrumentsV5RespDataInner) GetMaxStopSz() string`

GetMaxStopSz returns the MaxStopSz field if non-nil, zero value otherwise.

### GetMaxStopSzOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetMaxStopSzOk() (*string, bool)`

GetMaxStopSzOk returns a tuple with the MaxStopSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStopSz

`func (o *GetPublicInstrumentsV5RespDataInner) SetMaxStopSz(v string)`

SetMaxStopSz sets MaxStopSz field to given value.

### HasMaxStopSz

`func (o *GetPublicInstrumentsV5RespDataInner) HasMaxStopSz() bool`

HasMaxStopSz returns a boolean if a field has been set.

### GetMaxTriggerSz

`func (o *GetPublicInstrumentsV5RespDataInner) GetMaxTriggerSz() string`

GetMaxTriggerSz returns the MaxTriggerSz field if non-nil, zero value otherwise.

### GetMaxTriggerSzOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetMaxTriggerSzOk() (*string, bool)`

GetMaxTriggerSzOk returns a tuple with the MaxTriggerSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTriggerSz

`func (o *GetPublicInstrumentsV5RespDataInner) SetMaxTriggerSz(v string)`

SetMaxTriggerSz sets MaxTriggerSz field to given value.

### HasMaxTriggerSz

`func (o *GetPublicInstrumentsV5RespDataInner) HasMaxTriggerSz() bool`

HasMaxTriggerSz returns a boolean if a field has been set.

### GetMaxTwapSz

`func (o *GetPublicInstrumentsV5RespDataInner) GetMaxTwapSz() string`

GetMaxTwapSz returns the MaxTwapSz field if non-nil, zero value otherwise.

### GetMaxTwapSzOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetMaxTwapSzOk() (*string, bool)`

GetMaxTwapSzOk returns a tuple with the MaxTwapSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTwapSz

`func (o *GetPublicInstrumentsV5RespDataInner) SetMaxTwapSz(v string)`

SetMaxTwapSz sets MaxTwapSz field to given value.

### HasMaxTwapSz

`func (o *GetPublicInstrumentsV5RespDataInner) HasMaxTwapSz() bool`

HasMaxTwapSz returns a boolean if a field has been set.

### GetMinSz

`func (o *GetPublicInstrumentsV5RespDataInner) GetMinSz() string`

GetMinSz returns the MinSz field if non-nil, zero value otherwise.

### GetMinSzOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetMinSzOk() (*string, bool)`

GetMinSzOk returns a tuple with the MinSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSz

`func (o *GetPublicInstrumentsV5RespDataInner) SetMinSz(v string)`

SetMinSz sets MinSz field to given value.

### HasMinSz

`func (o *GetPublicInstrumentsV5RespDataInner) HasMinSz() bool`

HasMinSz returns a boolean if a field has been set.

### GetOptType

`func (o *GetPublicInstrumentsV5RespDataInner) GetOptType() string`

GetOptType returns the OptType field if non-nil, zero value otherwise.

### GetOptTypeOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetOptTypeOk() (*string, bool)`

GetOptTypeOk returns a tuple with the OptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptType

`func (o *GetPublicInstrumentsV5RespDataInner) SetOptType(v string)`

SetOptType sets OptType field to given value.

### HasOptType

`func (o *GetPublicInstrumentsV5RespDataInner) HasOptType() bool`

HasOptType returns a boolean if a field has been set.

### GetQuoteCcy

`func (o *GetPublicInstrumentsV5RespDataInner) GetQuoteCcy() string`

GetQuoteCcy returns the QuoteCcy field if non-nil, zero value otherwise.

### GetQuoteCcyOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetQuoteCcyOk() (*string, bool)`

GetQuoteCcyOk returns a tuple with the QuoteCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCcy

`func (o *GetPublicInstrumentsV5RespDataInner) SetQuoteCcy(v string)`

SetQuoteCcy sets QuoteCcy field to given value.

### HasQuoteCcy

`func (o *GetPublicInstrumentsV5RespDataInner) HasQuoteCcy() bool`

HasQuoteCcy returns a boolean if a field has been set.

### GetRuleType

`func (o *GetPublicInstrumentsV5RespDataInner) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *GetPublicInstrumentsV5RespDataInner) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *GetPublicInstrumentsV5RespDataInner) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetSettleCcy

`func (o *GetPublicInstrumentsV5RespDataInner) GetSettleCcy() string`

GetSettleCcy returns the SettleCcy field if non-nil, zero value otherwise.

### GetSettleCcyOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetSettleCcyOk() (*string, bool)`

GetSettleCcyOk returns a tuple with the SettleCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettleCcy

`func (o *GetPublicInstrumentsV5RespDataInner) SetSettleCcy(v string)`

SetSettleCcy sets SettleCcy field to given value.

### HasSettleCcy

`func (o *GetPublicInstrumentsV5RespDataInner) HasSettleCcy() bool`

HasSettleCcy returns a boolean if a field has been set.

### GetState

`func (o *GetPublicInstrumentsV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetPublicInstrumentsV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetPublicInstrumentsV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStk

`func (o *GetPublicInstrumentsV5RespDataInner) GetStk() string`

GetStk returns the Stk field if non-nil, zero value otherwise.

### GetStkOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetStkOk() (*string, bool)`

GetStkOk returns a tuple with the Stk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStk

`func (o *GetPublicInstrumentsV5RespDataInner) SetStk(v string)`

SetStk sets Stk field to given value.

### HasStk

`func (o *GetPublicInstrumentsV5RespDataInner) HasStk() bool`

HasStk returns a boolean if a field has been set.

### GetTickSz

`func (o *GetPublicInstrumentsV5RespDataInner) GetTickSz() string`

GetTickSz returns the TickSz field if non-nil, zero value otherwise.

### GetTickSzOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetTickSzOk() (*string, bool)`

GetTickSzOk returns a tuple with the TickSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickSz

`func (o *GetPublicInstrumentsV5RespDataInner) SetTickSz(v string)`

SetTickSz sets TickSz field to given value.

### HasTickSz

`func (o *GetPublicInstrumentsV5RespDataInner) HasTickSz() bool`

HasTickSz returns a boolean if a field has been set.

### GetUly

`func (o *GetPublicInstrumentsV5RespDataInner) GetUly() string`

GetUly returns the Uly field if non-nil, zero value otherwise.

### GetUlyOk

`func (o *GetPublicInstrumentsV5RespDataInner) GetUlyOk() (*string, bool)`

GetUlyOk returns a tuple with the Uly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUly

`func (o *GetPublicInstrumentsV5RespDataInner) SetUly(v string)`

SetUly sets Uly field to given value.

### HasUly

`func (o *GetPublicInstrumentsV5RespDataInner) HasUly() bool`

HasUly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


