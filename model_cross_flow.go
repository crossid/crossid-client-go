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

// CrossFlow struct for CrossFlow
type CrossFlow struct {
	DisplayName string `json:"displayName"`
	EndedAt *time.Time `json:"endedAt,omitempty"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	Id string `json:"id"`
	StartedAt time.Time `json:"startedAt"`
	Status string `json:"status"`
	StepId string `json:"stepId"`
	Steps CrossFlowSteps `json:"steps"`
	Vars map[string]map[string]interface{} `json:"vars"`
}

// NewCrossFlow instantiates a new CrossFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrossFlow(displayName string, id string, startedAt time.Time, status string, stepId string, steps CrossFlowSteps, vars map[string]map[string]interface{}) *CrossFlow {
	this := CrossFlow{}
	this.DisplayName = displayName
	this.Id = id
	this.StartedAt = startedAt
	this.Status = status
	this.StepId = stepId
	this.Steps = steps
	this.Vars = vars
	return &this
}

// NewCrossFlowWithDefaults instantiates a new CrossFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrossFlowWithDefaults() *CrossFlow {
	this := CrossFlow{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *CrossFlow) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *CrossFlow) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *CrossFlow) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetEndedAt returns the EndedAt field value if set, zero value otherwise.
func (o *CrossFlow) GetEndedAt() time.Time {
	if o == nil || o.EndedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.EndedAt
}

// GetEndedAtOk returns a tuple with the EndedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossFlow) GetEndedAtOk() (*time.Time, bool) {
	if o == nil || o.EndedAt == nil {
		return nil, false
	}
	return o.EndedAt, true
}

// HasEndedAt returns a boolean if a field has been set.
func (o *CrossFlow) HasEndedAt() bool {
	if o != nil && o.EndedAt != nil {
		return true
	}

	return false
}

// SetEndedAt gets a reference to the given time.Time and assigns it to the EndedAt field.
func (o *CrossFlow) SetEndedAt(v time.Time) {
	o.EndedAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *CrossFlow) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossFlow) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CrossFlow) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *CrossFlow) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value
func (o *CrossFlow) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CrossFlow) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CrossFlow) SetId(v string) {
	o.Id = v
}

// GetStartedAt returns the StartedAt field value
func (o *CrossFlow) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *CrossFlow) GetStartedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *CrossFlow) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetStatus returns the Status field value
func (o *CrossFlow) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CrossFlow) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CrossFlow) SetStatus(v string) {
	o.Status = v
}

// GetStepId returns the StepId field value
func (o *CrossFlow) GetStepId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StepId
}

// GetStepIdOk returns a tuple with the StepId field value
// and a boolean to check if the value has been set.
func (o *CrossFlow) GetStepIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StepId, true
}

// SetStepId sets field value
func (o *CrossFlow) SetStepId(v string) {
	o.StepId = v
}

// GetSteps returns the Steps field value
func (o *CrossFlow) GetSteps() CrossFlowSteps {
	if o == nil {
		var ret CrossFlowSteps
		return ret
	}

	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *CrossFlow) GetStepsOk() (*CrossFlowSteps, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Steps, true
}

// SetSteps sets field value
func (o *CrossFlow) SetSteps(v CrossFlowSteps) {
	o.Steps = v
}

// GetVars returns the Vars field value
func (o *CrossFlow) GetVars() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.Vars
}

// GetVarsOk returns a tuple with the Vars field value
// and a boolean to check if the value has been set.
func (o *CrossFlow) GetVarsOk() (*map[string]map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vars, true
}

// SetVars sets field value
func (o *CrossFlow) SetVars(v map[string]map[string]interface{}) {
	o.Vars = v
}

func (o CrossFlow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.EndedAt != nil {
		toSerialize["endedAt"] = o.EndedAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["startedAt"] = o.StartedAt
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["stepId"] = o.StepId
	}
	if true {
		toSerialize["steps"] = o.Steps
	}
	if true {
		toSerialize["vars"] = o.Vars
	}
	return json.Marshal(toSerialize)
}

type NullableCrossFlow struct {
	value *CrossFlow
	isSet bool
}

func (v NullableCrossFlow) Get() *CrossFlow {
	return v.value
}

func (v *NullableCrossFlow) Set(val *CrossFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableCrossFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableCrossFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrossFlow(val *CrossFlow) *NullableCrossFlow {
	return &NullableCrossFlow{value: val, isSet: true}
}

func (v NullableCrossFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrossFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


