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

// JobsListAllOf struct for JobsListAllOf
type JobsListAllOf struct {
	Resources []Job `json:"Resources"`
}

// NewJobsListAllOf instantiates a new JobsListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobsListAllOf(resources []Job, ) *JobsListAllOf {
	this := JobsListAllOf{}
	this.Resources = resources
	return &this
}

// NewJobsListAllOfWithDefaults instantiates a new JobsListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobsListAllOfWithDefaults() *JobsListAllOf {
	this := JobsListAllOf{}
	return &this
}

// GetResources returns the Resources field value
func (o *JobsListAllOf) GetResources() []Job {
	if o == nil  {
		var ret []Job
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *JobsListAllOf) GetResourcesOk() (*[]Job, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *JobsListAllOf) SetResources(v []Job) {
	o.Resources = v
}

func (o JobsListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullableJobsListAllOf struct {
	value *JobsListAllOf
	isSet bool
}

func (v NullableJobsListAllOf) Get() *JobsListAllOf {
	return v.value
}

func (v *NullableJobsListAllOf) Set(val *JobsListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableJobsListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableJobsListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobsListAllOf(val *JobsListAllOf) *NullableJobsListAllOf {
	return &NullableJobsListAllOf{value: val, isSet: true}
}

func (v NullableJobsListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobsListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


