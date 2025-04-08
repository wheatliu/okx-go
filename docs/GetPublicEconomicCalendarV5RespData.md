# GetPublicEconomicCalendarV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actual** | Pointer to **string** | The actual value of this event | [optional] [default to ""]
**CalendarId** | Pointer to **string** | Calendar ID | [optional] [default to ""]
**Category** | Pointer to **string** | Category name | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency of the data | [optional] [default to ""]
**Date** | Pointer to **string** | Estimated release time of the value of actual field, millisecond format of Unix timestamp, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**DateSpan** | Pointer to **string** | &#x60;0&#x60;: The time of the event is known  &#x60;1&#x60;: we only know the date of the event, the exact time of the event is unknown. | [optional] [default to ""]
**Event** | Pointer to **string** | Event name | [optional] [default to ""]
**Forecast** | Pointer to **string** | Average forecast among a representative group of economists | [optional] [default to ""]
**Importance** | Pointer to **string** | Level of importance   &#x60;1&#x60;: low   &#x60;2&#x60;: medium   &#x60;3&#x60;: high | [optional] [default to ""]
**PrevInitial** | Pointer to **string** | The initial value of the previous period   Only applicable when revision happens | [optional] [default to ""]
**Previous** | Pointer to **string** | Latest actual value of the previous period   The value will be revised if revision is applicable | [optional] [default to ""]
**RefDate** | Pointer to **string** | Date for which the datapoint refers to | [optional] [default to ""]
**Region** | Pointer to **string** | Country, region or entity | [optional] [default to ""]
**UTime** | Pointer to **string** | Update time of this record, millisecond format of Unix timestamp, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Unit** | Pointer to **string** | Unit of the data | [optional] [default to ""]

## Methods

### NewGetPublicEconomicCalendarV5RespData

`func NewGetPublicEconomicCalendarV5RespData() *GetPublicEconomicCalendarV5RespData`

NewGetPublicEconomicCalendarV5RespData instantiates a new GetPublicEconomicCalendarV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicEconomicCalendarV5RespDataWithDefaults

`func NewGetPublicEconomicCalendarV5RespDataWithDefaults() *GetPublicEconomicCalendarV5RespData`

NewGetPublicEconomicCalendarV5RespDataWithDefaults instantiates a new GetPublicEconomicCalendarV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActual

`func (o *GetPublicEconomicCalendarV5RespData) GetActual() string`

GetActual returns the Actual field if non-nil, zero value otherwise.

### GetActualOk

`func (o *GetPublicEconomicCalendarV5RespData) GetActualOk() (*string, bool)`

GetActualOk returns a tuple with the Actual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActual

`func (o *GetPublicEconomicCalendarV5RespData) SetActual(v string)`

SetActual sets Actual field to given value.

### HasActual

`func (o *GetPublicEconomicCalendarV5RespData) HasActual() bool`

HasActual returns a boolean if a field has been set.

### GetCalendarId

`func (o *GetPublicEconomicCalendarV5RespData) GetCalendarId() string`

GetCalendarId returns the CalendarId field if non-nil, zero value otherwise.

### GetCalendarIdOk

`func (o *GetPublicEconomicCalendarV5RespData) GetCalendarIdOk() (*string, bool)`

GetCalendarIdOk returns a tuple with the CalendarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarId

`func (o *GetPublicEconomicCalendarV5RespData) SetCalendarId(v string)`

SetCalendarId sets CalendarId field to given value.

### HasCalendarId

`func (o *GetPublicEconomicCalendarV5RespData) HasCalendarId() bool`

HasCalendarId returns a boolean if a field has been set.

### GetCategory

`func (o *GetPublicEconomicCalendarV5RespData) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetPublicEconomicCalendarV5RespData) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetPublicEconomicCalendarV5RespData) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetPublicEconomicCalendarV5RespData) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCcy

`func (o *GetPublicEconomicCalendarV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetPublicEconomicCalendarV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetPublicEconomicCalendarV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetPublicEconomicCalendarV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetDate

`func (o *GetPublicEconomicCalendarV5RespData) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetPublicEconomicCalendarV5RespData) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetPublicEconomicCalendarV5RespData) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *GetPublicEconomicCalendarV5RespData) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDateSpan

`func (o *GetPublicEconomicCalendarV5RespData) GetDateSpan() string`

GetDateSpan returns the DateSpan field if non-nil, zero value otherwise.

### GetDateSpanOk

`func (o *GetPublicEconomicCalendarV5RespData) GetDateSpanOk() (*string, bool)`

GetDateSpanOk returns a tuple with the DateSpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateSpan

`func (o *GetPublicEconomicCalendarV5RespData) SetDateSpan(v string)`

SetDateSpan sets DateSpan field to given value.

### HasDateSpan

`func (o *GetPublicEconomicCalendarV5RespData) HasDateSpan() bool`

HasDateSpan returns a boolean if a field has been set.

### GetEvent

`func (o *GetPublicEconomicCalendarV5RespData) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *GetPublicEconomicCalendarV5RespData) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *GetPublicEconomicCalendarV5RespData) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *GetPublicEconomicCalendarV5RespData) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetForecast

`func (o *GetPublicEconomicCalendarV5RespData) GetForecast() string`

GetForecast returns the Forecast field if non-nil, zero value otherwise.

### GetForecastOk

`func (o *GetPublicEconomicCalendarV5RespData) GetForecastOk() (*string, bool)`

GetForecastOk returns a tuple with the Forecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecast

`func (o *GetPublicEconomicCalendarV5RespData) SetForecast(v string)`

SetForecast sets Forecast field to given value.

### HasForecast

`func (o *GetPublicEconomicCalendarV5RespData) HasForecast() bool`

HasForecast returns a boolean if a field has been set.

### GetImportance

`func (o *GetPublicEconomicCalendarV5RespData) GetImportance() string`

GetImportance returns the Importance field if non-nil, zero value otherwise.

### GetImportanceOk

`func (o *GetPublicEconomicCalendarV5RespData) GetImportanceOk() (*string, bool)`

GetImportanceOk returns a tuple with the Importance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportance

`func (o *GetPublicEconomicCalendarV5RespData) SetImportance(v string)`

SetImportance sets Importance field to given value.

### HasImportance

`func (o *GetPublicEconomicCalendarV5RespData) HasImportance() bool`

HasImportance returns a boolean if a field has been set.

### GetPrevInitial

`func (o *GetPublicEconomicCalendarV5RespData) GetPrevInitial() string`

GetPrevInitial returns the PrevInitial field if non-nil, zero value otherwise.

### GetPrevInitialOk

`func (o *GetPublicEconomicCalendarV5RespData) GetPrevInitialOk() (*string, bool)`

GetPrevInitialOk returns a tuple with the PrevInitial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevInitial

`func (o *GetPublicEconomicCalendarV5RespData) SetPrevInitial(v string)`

SetPrevInitial sets PrevInitial field to given value.

### HasPrevInitial

`func (o *GetPublicEconomicCalendarV5RespData) HasPrevInitial() bool`

HasPrevInitial returns a boolean if a field has been set.

### GetPrevious

`func (o *GetPublicEconomicCalendarV5RespData) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *GetPublicEconomicCalendarV5RespData) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *GetPublicEconomicCalendarV5RespData) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *GetPublicEconomicCalendarV5RespData) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetRefDate

`func (o *GetPublicEconomicCalendarV5RespData) GetRefDate() string`

GetRefDate returns the RefDate field if non-nil, zero value otherwise.

### GetRefDateOk

`func (o *GetPublicEconomicCalendarV5RespData) GetRefDateOk() (*string, bool)`

GetRefDateOk returns a tuple with the RefDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefDate

`func (o *GetPublicEconomicCalendarV5RespData) SetRefDate(v string)`

SetRefDate sets RefDate field to given value.

### HasRefDate

`func (o *GetPublicEconomicCalendarV5RespData) HasRefDate() bool`

HasRefDate returns a boolean if a field has been set.

### GetRegion

`func (o *GetPublicEconomicCalendarV5RespData) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetPublicEconomicCalendarV5RespData) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetPublicEconomicCalendarV5RespData) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetPublicEconomicCalendarV5RespData) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetUTime

`func (o *GetPublicEconomicCalendarV5RespData) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetPublicEconomicCalendarV5RespData) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetPublicEconomicCalendarV5RespData) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetPublicEconomicCalendarV5RespData) HasUTime() bool`

HasUTime returns a boolean if a field has been set.

### GetUnit

`func (o *GetPublicEconomicCalendarV5RespData) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *GetPublicEconomicCalendarV5RespData) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *GetPublicEconomicCalendarV5RespData) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *GetPublicEconomicCalendarV5RespData) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


