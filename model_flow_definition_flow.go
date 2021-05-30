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
)

// FlowDefinitionFlow struct for FlowDefinitionFlow
type FlowDefinitionFlow struct {
	DisplayName *string `json:"displayName,omitempty"`
	ExternalId *string `json:"externalId,omitempty"`
	Id *string `json:"id,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewFlowDefinitionFlow instantiates a new FlowDefinitionFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowDefinitionFlow() *FlowDefinitionFlow {
	this := FlowDefinitionFlow{}
	return &this
}

// NewFlowDefinitionFlowWithDefaults instantiates a new FlowDefinitionFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowDefinitionFlowWithDefaults() *FlowDefinitionFlow {
	this := FlowDefinitionFlow{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *FlowDefinitionFlow) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowDefinitionFlow) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *FlowDefinitionFlow) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *FlowDefinitionFlow) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *FlowDefinitionFlow) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowDefinitionFlow) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *FlowDefinitionFlow) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *FlowDefinitionFlow) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FlowDefinitionFlow) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowDefinitionFlow) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FlowDefinitionFlow) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FlowDefinitionFlow) SetId(v string) {
	o.Id = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *FlowDefinitionFlow) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowDefinitionFlow) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *FlowDefinitionFlow) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *FlowDefinitionFlow) SetVersion(v int64) {
	o.Version = &v
}

func (o FlowDefinitionFlow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableFlowDefinitionFlow struct {
	value *FlowDefinitionFlow
	isSet bool
}

func (v NullableFlowDefinitionFlow) Get() *FlowDefinitionFlow {
	return v.value
}

func (v *NullableFlowDefinitionFlow) Set(val *FlowDefinitionFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowDefinitionFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowDefinitionFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowDefinitionFlow(val *FlowDefinitionFlow) *NullableFlowDefinitionFlow {
	return &NullableFlowDefinitionFlow{value: val, isSet: true}
}

func (v NullableFlowDefinitionFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowDefinitionFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

