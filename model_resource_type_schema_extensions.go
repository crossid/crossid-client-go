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

// ResourceTypeSchemaExtensions struct for ResourceTypeSchemaExtensions
type ResourceTypeSchemaExtensions struct {
	Required *bool `json:"required,omitempty"`
	Schema *string `json:"schema,omitempty"`
}

// NewResourceTypeSchemaExtensions instantiates a new ResourceTypeSchemaExtensions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceTypeSchemaExtensions() *ResourceTypeSchemaExtensions {
	this := ResourceTypeSchemaExtensions{}
	return &this
}

// NewResourceTypeSchemaExtensionsWithDefaults instantiates a new ResourceTypeSchemaExtensions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceTypeSchemaExtensionsWithDefaults() *ResourceTypeSchemaExtensions {
	this := ResourceTypeSchemaExtensions{}
	return &this
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *ResourceTypeSchemaExtensions) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTypeSchemaExtensions) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *ResourceTypeSchemaExtensions) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *ResourceTypeSchemaExtensions) SetRequired(v bool) {
	o.Required = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *ResourceTypeSchemaExtensions) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTypeSchemaExtensions) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *ResourceTypeSchemaExtensions) HasSchema() bool {
	if o != nil && o.Schema != nil {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *ResourceTypeSchemaExtensions) SetSchema(v string) {
	o.Schema = &v
}

func (o ResourceTypeSchemaExtensions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	return json.Marshal(toSerialize)
}

type NullableResourceTypeSchemaExtensions struct {
	value *ResourceTypeSchemaExtensions
	isSet bool
}

func (v NullableResourceTypeSchemaExtensions) Get() *ResourceTypeSchemaExtensions {
	return v.value
}

func (v *NullableResourceTypeSchemaExtensions) Set(val *ResourceTypeSchemaExtensions) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceTypeSchemaExtensions) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceTypeSchemaExtensions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceTypeSchemaExtensions(val *ResourceTypeSchemaExtensions) *NullableResourceTypeSchemaExtensions {
	return &NullableResourceTypeSchemaExtensions{value: val, isSet: true}
}

func (v NullableResourceTypeSchemaExtensions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceTypeSchemaExtensions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

