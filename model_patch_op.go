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

// PatchOP struct for PatchOP
type PatchOP struct {
	Op string `json:"op"`
	Path string `json:"path"`
	Value interface{} `json:"value"`
}

// NewPatchOP instantiates a new PatchOP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchOP(op string, path string, value interface{}, ) *PatchOP {
	this := PatchOP{}
	this.Op = op
	this.Path = path
	this.Value = value
	return &this
}

// NewPatchOPWithDefaults instantiates a new PatchOP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchOPWithDefaults() *PatchOP {
	this := PatchOP{}
	return &this
}

// GetOp returns the Op field value
func (o *PatchOP) GetOp() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *PatchOP) GetOpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *PatchOP) SetOp(v string) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *PatchOP) GetPath() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *PatchOP) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *PatchOP) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PatchOP) GetValue() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchOP) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PatchOP) SetValue(v interface{}) {
	o.Value = v
}

func (o PatchOP) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["op"] = o.Op
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullablePatchOP struct {
	value *PatchOP
	isSet bool
}

func (v NullablePatchOP) Get() *PatchOP {
	return v.value
}

func (v *NullablePatchOP) Set(val *PatchOP) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchOP) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchOP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchOP(val *PatchOP) *NullablePatchOP {
	return &NullablePatchOP{value: val, isSet: true}
}

func (v NullablePatchOP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchOP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

