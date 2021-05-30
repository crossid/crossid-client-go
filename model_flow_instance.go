/*
 * CrossID Public API
 *
 * CrossID API Reference 
 *
 * API version: 1.0.0
 * Contact: developer@crossid.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cidclient

import (
	"encoding/json"
	"time"
)

// FlowInstance struct for FlowInstance
type FlowInstance struct {
	StatusExplanations *[]map[string]interface{} `json:"StatusExplanations,omitempty"`
	Description *string `json:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	EndedAt *time.Time `json:"endedAt,omitempty"`
	Flow *FlowDefinitionFlow `json:"flow,omitempty"`
	Id *string `json:"id,omitempty"`
	StartedAt *time.Time `json:"startedAt,omitempty"`
	Status *string `json:"status,omitempty"`
	Variables interface{} `json:"variables,omitempty"`
}

// NewFlowInstance instantiates a new FlowInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowInstance() *FlowInstance {
	this := FlowInstance{}
	return &this
}

// NewFlowInstanceWithDefaults instantiates a new FlowInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowInstanceWithDefaults() *FlowInstance {
	this := FlowInstance{}
	return &this
}

// GetStatusExplanations returns the StatusExplanations field value if set, zero value otherwise.
func (o *FlowInstance) GetStatusExplanations() []map[string]interface{} {
	if o == nil || o.StatusExplanations == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.StatusExplanations
}

// GetStatusExplanationsOk returns a tuple with the StatusExplanations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowInstance) GetStatusExplanationsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.StatusExplanations == nil {
		return nil, false
	}
	return o.StatusExplanations, true
}

// HasStatusExplanations returns a boolean if a field has been set.
func (o *FlowInstance) HasStatusExplanations() bool {
	if o != nil && o.StatusExplanations != nil {
		return true
	}

	return false
}

// SetStatusExplanations gets a reference to the given []map[string]interface{} and assigns it to the StatusExplanations field.
func (o *FlowInstance) SetStatusExplanations(v []map[string]interface{}) {
	o.StatusExplanations = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FlowInstance) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowInstance) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FlowInstance) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FlowInstance) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *FlowInstance) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowInstance) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *FlowInstance) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *FlowInstance) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEndedAt returns the EndedAt field value if set, zero value otherwise.
func (o *FlowInstance) GetEndedAt() time.Time {
	if o == nil || o.EndedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.EndedAt
}

// GetEndedAtOk returns a tuple with the EndedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowInstance) GetEndedAtOk() (*time.Time, bool) {
	if o == nil || o.EndedAt == nil {
		return nil, false
	}
	return o.EndedAt, true
}

// HasEndedAt returns a boolean if a field has been set.
func (o *FlowInstance) HasEndedAt() bool {
	if o != nil && o.EndedAt != nil {
		return true
	}

	return false
}

// SetEndedAt gets a reference to the given time.Time and assigns it to the EndedAt field.
func (o *FlowInstance) SetEndedAt(v time.Time) {
	o.EndedAt = &v
}

// GetFlow returns the Flow field value if set, zero value otherwise.
func (o *FlowInstance) GetFlow() FlowDefinitionFlow {
	if o == nil || o.Flow == nil {
		var ret FlowDefinitionFlow
		return ret
	}
	return *o.Flow
}

// GetFlowOk returns a tuple with the Flow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowInstance) GetFlowOk() (*FlowDefinitionFlow, bool) {
	if o == nil || o.Flow == nil {
		return nil, false
	}
	return o.Flow, true
}

// HasFlow returns a boolean if a field has been set.
func (o *FlowInstance) HasFlow() bool {
	if o != nil && o.Flow != nil {
		return true
	}

	return false
}

// SetFlow gets a reference to the given FlowDefinitionFlow and assigns it to the Flow field.
func (o *FlowInstance) SetFlow(v FlowDefinitionFlow) {
	o.Flow = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FlowInstance) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowInstance) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FlowInstance) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FlowInstance) SetId(v string) {
	o.Id = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *FlowInstance) GetStartedAt() time.Time {
	if o == nil || o.StartedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowInstance) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *FlowInstance) HasStartedAt() bool {
	if o != nil && o.StartedAt != nil {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *FlowInstance) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FlowInstance) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowInstance) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FlowInstance) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FlowInstance) SetStatus(v string) {
	o.Status = &v
}

// GetVariables returns the Variables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlowInstance) GetVariables() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlowInstance) GetVariablesOk() (*interface{}, bool) {
	if o == nil || o.Variables == nil {
		return nil, false
	}
	return &o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *FlowInstance) HasVariables() bool {
	if o != nil && o.Variables != nil {
		return true
	}

	return false
}

// SetVariables gets a reference to the given interface{} and assigns it to the Variables field.
func (o *FlowInstance) SetVariables(v interface{}) {
	o.Variables = v
}

func (o FlowInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StatusExplanations != nil {
		toSerialize["StatusExplanations"] = o.StatusExplanations
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.EndedAt != nil {
		toSerialize["endedAt"] = o.EndedAt
	}
	if o.Flow != nil {
		toSerialize["flow"] = o.Flow
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.StartedAt != nil {
		toSerialize["startedAt"] = o.StartedAt
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	return json.Marshal(toSerialize)
}

type NullableFlowInstance struct {
	value *FlowInstance
	isSet bool
}

func (v NullableFlowInstance) Get() *FlowInstance {
	return v.value
}

func (v *NullableFlowInstance) Set(val *FlowInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowInstance(val *FlowInstance) *NullableFlowInstance {
	return &NullableFlowInstance{value: val, isSet: true}
}

func (v NullableFlowInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

