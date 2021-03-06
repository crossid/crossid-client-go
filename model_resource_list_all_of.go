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

// ResourceListAllOf struct for ResourceListAllOf
type ResourceListAllOf struct {
	Resources []Resource `json:"Resources"`
}

// NewResourceListAllOf instantiates a new ResourceListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceListAllOf(resources []Resource) *ResourceListAllOf {
	this := ResourceListAllOf{}
	this.Resources = resources
	return &this
}

// NewResourceListAllOfWithDefaults instantiates a new ResourceListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceListAllOfWithDefaults() *ResourceListAllOf {
	this := ResourceListAllOf{}
	return &this
}

// GetResources returns the Resources field value
func (o *ResourceListAllOf) GetResources() []Resource {
	if o == nil {
		var ret []Resource
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *ResourceListAllOf) GetResourcesOk() (*[]Resource, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *ResourceListAllOf) SetResources(v []Resource) {
	o.Resources = v
}

func (o ResourceListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullableResourceListAllOf struct {
	value *ResourceListAllOf
	isSet bool
}

func (v NullableResourceListAllOf) Get() *ResourceListAllOf {
	return v.value
}

func (v *NullableResourceListAllOf) Set(val *ResourceListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceListAllOf(val *ResourceListAllOf) *NullableResourceListAllOf {
	return &NullableResourceListAllOf{value: val, isSet: true}
}

func (v NullableResourceListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


