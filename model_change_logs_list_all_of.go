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

// ChangeLogsListAllOf struct for ChangeLogsListAllOf
type ChangeLogsListAllOf struct {
	Resources []ChangeLog `json:"Resources"`
}

// NewChangeLogsListAllOf instantiates a new ChangeLogsListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeLogsListAllOf(resources []ChangeLog) *ChangeLogsListAllOf {
	this := ChangeLogsListAllOf{}
	this.Resources = resources
	return &this
}

// NewChangeLogsListAllOfWithDefaults instantiates a new ChangeLogsListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeLogsListAllOfWithDefaults() *ChangeLogsListAllOf {
	this := ChangeLogsListAllOf{}
	return &this
}

// GetResources returns the Resources field value
func (o *ChangeLogsListAllOf) GetResources() []ChangeLog {
	if o == nil {
		var ret []ChangeLog
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *ChangeLogsListAllOf) GetResourcesOk() (*[]ChangeLog, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *ChangeLogsListAllOf) SetResources(v []ChangeLog) {
	o.Resources = v
}

func (o ChangeLogsListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullableChangeLogsListAllOf struct {
	value *ChangeLogsListAllOf
	isSet bool
}

func (v NullableChangeLogsListAllOf) Get() *ChangeLogsListAllOf {
	return v.value
}

func (v *NullableChangeLogsListAllOf) Set(val *ChangeLogsListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeLogsListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeLogsListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeLogsListAllOf(val *ChangeLogsListAllOf) *NullableChangeLogsListAllOf {
	return &NullableChangeLogsListAllOf{value: val, isSet: true}
}

func (v NullableChangeLogsListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeLogsListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


