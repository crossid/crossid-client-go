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

// RegistrationFlow struct for RegistrationFlow
type RegistrationFlow struct {
	Id *string `json:"id,omitempty"`
	Method *string `json:"method,omitempty"`
	Methods *RegistrationFlowMethods `json:"methods,omitempty"`
}

// NewRegistrationFlow instantiates a new RegistrationFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationFlow() *RegistrationFlow {
	this := RegistrationFlow{}
	return &this
}

// NewRegistrationFlowWithDefaults instantiates a new RegistrationFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationFlowWithDefaults() *RegistrationFlow {
	this := RegistrationFlow{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RegistrationFlow) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationFlow) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RegistrationFlow) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RegistrationFlow) SetId(v string) {
	o.Id = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *RegistrationFlow) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationFlow) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *RegistrationFlow) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *RegistrationFlow) SetMethod(v string) {
	o.Method = &v
}

// GetMethods returns the Methods field value if set, zero value otherwise.
func (o *RegistrationFlow) GetMethods() RegistrationFlowMethods {
	if o == nil || o.Methods == nil {
		var ret RegistrationFlowMethods
		return ret
	}
	return *o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationFlow) GetMethodsOk() (*RegistrationFlowMethods, bool) {
	if o == nil || o.Methods == nil {
		return nil, false
	}
	return o.Methods, true
}

// HasMethods returns a boolean if a field has been set.
func (o *RegistrationFlow) HasMethods() bool {
	if o != nil && o.Methods != nil {
		return true
	}

	return false
}

// SetMethods gets a reference to the given RegistrationFlowMethods and assigns it to the Methods field.
func (o *RegistrationFlow) SetMethods(v RegistrationFlowMethods) {
	o.Methods = &v
}

func (o RegistrationFlow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Methods != nil {
		toSerialize["methods"] = o.Methods
	}
	return json.Marshal(toSerialize)
}

type NullableRegistrationFlow struct {
	value *RegistrationFlow
	isSet bool
}

func (v NullableRegistrationFlow) Get() *RegistrationFlow {
	return v.value
}

func (v *NullableRegistrationFlow) Set(val *RegistrationFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationFlow(val *RegistrationFlow) *NullableRegistrationFlow {
	return &NullableRegistrationFlow{value: val, isSet: true}
}

func (v NullableRegistrationFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

