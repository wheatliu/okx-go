/*
OKX v5 API

OpenAPI specification for Okx exchange - Rest API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rest

import (
	"encoding/json"
)

// checks if the GetTradingBotSignalEventHistoryV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTradingBotSignalEventHistoryV5RespDataInner{}

// GetTradingBotSignalEventHistoryV5RespDataInner struct for GetTradingBotSignalEventHistoryV5RespDataInner
type GetTradingBotSignalEventHistoryV5RespDataInner struct {
	// Alert message
	AlertMsg *string `json:"alertMsg,omitempty"`
	// Algo ID
	AlgoId *string `json:"algoId,omitempty"`
	// Event timestamp of creation. Unix timestamp format in milliseconds, e.g. `1597026383085`
	EventCtime *string `json:"eventCtime,omitempty"`
	// Event process message
	EventProcessMsg *string `json:"eventProcessMsg,omitempty"`
	// Event status  `success`  `failure`
	EventStatus *string `json:"eventStatus,omitempty"`
	// Event type  `system_action`  `user_action`  `signal_processing`
	EventType *string `json:"eventType,omitempty"`
	// Event timestamp of update. Unix timestamp format in milliseconds, e.g. `1597026383085`
	EventUtime *string `json:"eventUtime,omitempty"`
	// Triggered sub order data
	TriggeredOrdData []GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner `json:"triggeredOrdData,omitempty"`
}

// NewGetTradingBotSignalEventHistoryV5RespDataInner instantiates a new GetTradingBotSignalEventHistoryV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTradingBotSignalEventHistoryV5RespDataInner() *GetTradingBotSignalEventHistoryV5RespDataInner {
	this := GetTradingBotSignalEventHistoryV5RespDataInner{}
	var alertMsg string = ""
	this.AlertMsg = &alertMsg
	var algoId string = ""
	this.AlgoId = &algoId
	var eventCtime string = ""
	this.EventCtime = &eventCtime
	var eventProcessMsg string = ""
	this.EventProcessMsg = &eventProcessMsg
	var eventStatus string = ""
	this.EventStatus = &eventStatus
	var eventType string = ""
	this.EventType = &eventType
	var eventUtime string = ""
	this.EventUtime = &eventUtime
	return &this
}

// NewGetTradingBotSignalEventHistoryV5RespDataInnerWithDefaults instantiates a new GetTradingBotSignalEventHistoryV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTradingBotSignalEventHistoryV5RespDataInnerWithDefaults() *GetTradingBotSignalEventHistoryV5RespDataInner {
	this := GetTradingBotSignalEventHistoryV5RespDataInner{}
	var alertMsg string = ""
	this.AlertMsg = &alertMsg
	var algoId string = ""
	this.AlgoId = &algoId
	var eventCtime string = ""
	this.EventCtime = &eventCtime
	var eventProcessMsg string = ""
	this.EventProcessMsg = &eventProcessMsg
	var eventStatus string = ""
	this.EventStatus = &eventStatus
	var eventType string = ""
	this.EventType = &eventType
	var eventUtime string = ""
	this.EventUtime = &eventUtime
	return &this
}

// GetAlertMsg returns the AlertMsg field value if set, zero value otherwise.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetAlertMsg() string {
	if o == nil || IsNil(o.AlertMsg) {
		var ret string
		return ret
	}
	return *o.AlertMsg
}

// GetAlertMsgOk returns a tuple with the AlertMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetAlertMsgOk() (*string, bool) {
	if o == nil || IsNil(o.AlertMsg) {
		return nil, false
	}
	return o.AlertMsg, true
}

// HasAlertMsg returns a boolean if a field has been set.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) HasAlertMsg() bool {
	if o != nil && !IsNil(o.AlertMsg) {
		return true
	}

	return false
}

// SetAlertMsg gets a reference to the given string and assigns it to the AlertMsg field.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) SetAlertMsg(v string) {
	o.AlertMsg = &v
}

// GetAlgoId returns the AlgoId field value if set, zero value otherwise.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetAlgoId() string {
	if o == nil || IsNil(o.AlgoId) {
		var ret string
		return ret
	}
	return *o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetAlgoIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoId) {
		return nil, false
	}
	return o.AlgoId, true
}

// HasAlgoId returns a boolean if a field has been set.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) HasAlgoId() bool {
	if o != nil && !IsNil(o.AlgoId) {
		return true
	}

	return false
}

// SetAlgoId gets a reference to the given string and assigns it to the AlgoId field.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) SetAlgoId(v string) {
	o.AlgoId = &v
}

// GetEventCtime returns the EventCtime field value if set, zero value otherwise.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetEventCtime() string {
	if o == nil || IsNil(o.EventCtime) {
		var ret string
		return ret
	}
	return *o.EventCtime
}

// GetEventCtimeOk returns a tuple with the EventCtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetEventCtimeOk() (*string, bool) {
	if o == nil || IsNil(o.EventCtime) {
		return nil, false
	}
	return o.EventCtime, true
}

// HasEventCtime returns a boolean if a field has been set.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) HasEventCtime() bool {
	if o != nil && !IsNil(o.EventCtime) {
		return true
	}

	return false
}

// SetEventCtime gets a reference to the given string and assigns it to the EventCtime field.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) SetEventCtime(v string) {
	o.EventCtime = &v
}

// GetEventProcessMsg returns the EventProcessMsg field value if set, zero value otherwise.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetEventProcessMsg() string {
	if o == nil || IsNil(o.EventProcessMsg) {
		var ret string
		return ret
	}
	return *o.EventProcessMsg
}

// GetEventProcessMsgOk returns a tuple with the EventProcessMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetEventProcessMsgOk() (*string, bool) {
	if o == nil || IsNil(o.EventProcessMsg) {
		return nil, false
	}
	return o.EventProcessMsg, true
}

// HasEventProcessMsg returns a boolean if a field has been set.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) HasEventProcessMsg() bool {
	if o != nil && !IsNil(o.EventProcessMsg) {
		return true
	}

	return false
}

// SetEventProcessMsg gets a reference to the given string and assigns it to the EventProcessMsg field.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) SetEventProcessMsg(v string) {
	o.EventProcessMsg = &v
}

// GetEventStatus returns the EventStatus field value if set, zero value otherwise.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetEventStatus() string {
	if o == nil || IsNil(o.EventStatus) {
		var ret string
		return ret
	}
	return *o.EventStatus
}

// GetEventStatusOk returns a tuple with the EventStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetEventStatusOk() (*string, bool) {
	if o == nil || IsNil(o.EventStatus) {
		return nil, false
	}
	return o.EventStatus, true
}

// HasEventStatus returns a boolean if a field has been set.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) HasEventStatus() bool {
	if o != nil && !IsNil(o.EventStatus) {
		return true
	}

	return false
}

// SetEventStatus gets a reference to the given string and assigns it to the EventStatus field.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) SetEventStatus(v string) {
	o.EventStatus = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) SetEventType(v string) {
	o.EventType = &v
}

// GetEventUtime returns the EventUtime field value if set, zero value otherwise.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetEventUtime() string {
	if o == nil || IsNil(o.EventUtime) {
		var ret string
		return ret
	}
	return *o.EventUtime
}

// GetEventUtimeOk returns a tuple with the EventUtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetEventUtimeOk() (*string, bool) {
	if o == nil || IsNil(o.EventUtime) {
		return nil, false
	}
	return o.EventUtime, true
}

// HasEventUtime returns a boolean if a field has been set.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) HasEventUtime() bool {
	if o != nil && !IsNil(o.EventUtime) {
		return true
	}

	return false
}

// SetEventUtime gets a reference to the given string and assigns it to the EventUtime field.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) SetEventUtime(v string) {
	o.EventUtime = &v
}

// GetTriggeredOrdData returns the TriggeredOrdData field value if set, zero value otherwise.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetTriggeredOrdData() []GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner {
	if o == nil || IsNil(o.TriggeredOrdData) {
		var ret []GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner
		return ret
	}
	return o.TriggeredOrdData
}

// GetTriggeredOrdDataOk returns a tuple with the TriggeredOrdData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetTriggeredOrdDataOk() ([]GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner, bool) {
	if o == nil || IsNil(o.TriggeredOrdData) {
		return nil, false
	}
	return o.TriggeredOrdData, true
}

// HasTriggeredOrdData returns a boolean if a field has been set.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) HasTriggeredOrdData() bool {
	if o != nil && !IsNil(o.TriggeredOrdData) {
		return true
	}

	return false
}

// SetTriggeredOrdData gets a reference to the given []GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner and assigns it to the TriggeredOrdData field.
func (o *GetTradingBotSignalEventHistoryV5RespDataInner) SetTriggeredOrdData(v []GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner) {
	o.TriggeredOrdData = v
}

func (o GetTradingBotSignalEventHistoryV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTradingBotSignalEventHistoryV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlertMsg) {
		toSerialize["alertMsg"] = o.AlertMsg
	}
	if !IsNil(o.AlgoId) {
		toSerialize["algoId"] = o.AlgoId
	}
	if !IsNil(o.EventCtime) {
		toSerialize["eventCtime"] = o.EventCtime
	}
	if !IsNil(o.EventProcessMsg) {
		toSerialize["eventProcessMsg"] = o.EventProcessMsg
	}
	if !IsNil(o.EventStatus) {
		toSerialize["eventStatus"] = o.EventStatus
	}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.EventUtime) {
		toSerialize["eventUtime"] = o.EventUtime
	}
	if !IsNil(o.TriggeredOrdData) {
		toSerialize["triggeredOrdData"] = o.TriggeredOrdData
	}
	return toSerialize, nil
}

type NullableGetTradingBotSignalEventHistoryV5RespDataInner struct {
	value *GetTradingBotSignalEventHistoryV5RespDataInner
	isSet bool
}

func (v NullableGetTradingBotSignalEventHistoryV5RespDataInner) Get() *GetTradingBotSignalEventHistoryV5RespDataInner {
	return v.value
}

func (v *NullableGetTradingBotSignalEventHistoryV5RespDataInner) Set(val *GetTradingBotSignalEventHistoryV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTradingBotSignalEventHistoryV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTradingBotSignalEventHistoryV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTradingBotSignalEventHistoryV5RespDataInner(val *GetTradingBotSignalEventHistoryV5RespDataInner) *NullableGetTradingBotSignalEventHistoryV5RespDataInner {
	return &NullableGetTradingBotSignalEventHistoryV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetTradingBotSignalEventHistoryV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTradingBotSignalEventHistoryV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


