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

// MappersListAllOf struct for MappersListAllOf
type MappersListAllOf struct {
	Resources []Mapper `json:"Resources"`
}

// NewMappersListAllOf instantiates a new MappersListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMappersListAllOf(resources []Mapper, ) *MappersListAllOf {
	this := MappersListAllOf{}
	this.Resources = resources
	return &this
}

// NewMappersListAllOfWithDefaults instantiates a new MappersListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMappersListAllOfWithDefaults() *MappersListAllOf {
	this := MappersListAllOf{}
	return &this
}

// GetResources returns the Resources field value
func (o *MappersListAllOf) GetResources() []Mapper {
	if o == nil  {
		var ret []Mapper
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *MappersListAllOf) GetResourcesOk() (*[]Mapper, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *MappersListAllOf) SetResources(v []Mapper) {
	o.Resources = v
}

func (o MappersListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullableMappersListAllOf struct {
	value *MappersListAllOf
	isSet bool
}

func (v NullableMappersListAllOf) Get() *MappersListAllOf {
	return v.value
}

func (v *NullableMappersListAllOf) Set(val *MappersListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMappersListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMappersListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMappersListAllOf(val *MappersListAllOf) *NullableMappersListAllOf {
	return &NullableMappersListAllOf{value: val, isSet: true}
}

func (v NullableMappersListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMappersListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

