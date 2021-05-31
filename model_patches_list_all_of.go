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

// PatchesListAllOf struct for PatchesListAllOf
type PatchesListAllOf struct {
	Resources []DocPatch `json:"Resources"`
}

// NewPatchesListAllOf instantiates a new PatchesListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchesListAllOf(resources []DocPatch) *PatchesListAllOf {
	this := PatchesListAllOf{}
	this.Resources = resources
	return &this
}

// NewPatchesListAllOfWithDefaults instantiates a new PatchesListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchesListAllOfWithDefaults() *PatchesListAllOf {
	this := PatchesListAllOf{}
	return &this
}

// GetResources returns the Resources field value
func (o *PatchesListAllOf) GetResources() []DocPatch {
	if o == nil {
		var ret []DocPatch
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *PatchesListAllOf) GetResourcesOk() (*[]DocPatch, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *PatchesListAllOf) SetResources(v []DocPatch) {
	o.Resources = v
}

func (o PatchesListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullablePatchesListAllOf struct {
	value *PatchesListAllOf
	isSet bool
}

func (v NullablePatchesListAllOf) Get() *PatchesListAllOf {
	return v.value
}

func (v *NullablePatchesListAllOf) Set(val *PatchesListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchesListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchesListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchesListAllOf(val *PatchesListAllOf) *NullablePatchesListAllOf {
	return &NullablePatchesListAllOf{value: val, isSet: true}
}

func (v NullablePatchesListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchesListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


