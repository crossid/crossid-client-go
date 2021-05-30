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

// LoginFlowMethodConfigFields struct for LoginFlowMethodConfigFields
type LoginFlowMethodConfigFields struct {
	Disabled *bool `json:"disabled,omitempty"`
	Messages *[]UserFacingMessage `json:"messages,omitempty"`
	Name *string `json:"name,omitempty"`
	Required *bool `json:"required,omitempty"`
	Type *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewLoginFlowMethodConfigFields instantiates a new LoginFlowMethodConfigFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginFlowMethodConfigFields() *LoginFlowMethodConfigFields {
	this := LoginFlowMethodConfigFields{}
	return &this
}

// NewLoginFlowMethodConfigFieldsWithDefaults instantiates a new LoginFlowMethodConfigFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginFlowMethodConfigFieldsWithDefaults() *LoginFlowMethodConfigFields {
	this := LoginFlowMethodConfigFields{}
	return &this
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *LoginFlowMethodConfigFields) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlowMethodConfigFields) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *LoginFlowMethodConfigFields) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *LoginFlowMethodConfigFields) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *LoginFlowMethodConfigFields) GetMessages() []UserFacingMessage {
	if o == nil || o.Messages == nil {
		var ret []UserFacingMessage
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlowMethodConfigFields) GetMessagesOk() (*[]UserFacingMessage, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *LoginFlowMethodConfigFields) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []UserFacingMessage and assigns it to the Messages field.
func (o *LoginFlowMethodConfigFields) SetMessages(v []UserFacingMessage) {
	o.Messages = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoginFlowMethodConfigFields) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlowMethodConfigFields) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoginFlowMethodConfigFields) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoginFlowMethodConfigFields) SetName(v string) {
	o.Name = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *LoginFlowMethodConfigFields) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlowMethodConfigFields) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *LoginFlowMethodConfigFields) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *LoginFlowMethodConfigFields) SetRequired(v bool) {
	o.Required = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LoginFlowMethodConfigFields) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlowMethodConfigFields) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LoginFlowMethodConfigFields) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LoginFlowMethodConfigFields) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *LoginFlowMethodConfigFields) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlowMethodConfigFields) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *LoginFlowMethodConfigFields) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *LoginFlowMethodConfigFields) SetValue(v string) {
	o.Value = &v
}

func (o LoginFlowMethodConfigFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableLoginFlowMethodConfigFields struct {
	value *LoginFlowMethodConfigFields
	isSet bool
}

func (v NullableLoginFlowMethodConfigFields) Get() *LoginFlowMethodConfigFields {
	return v.value
}

func (v *NullableLoginFlowMethodConfigFields) Set(val *LoginFlowMethodConfigFields) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginFlowMethodConfigFields) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginFlowMethodConfigFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginFlowMethodConfigFields(val *LoginFlowMethodConfigFields) *NullableLoginFlowMethodConfigFields {
	return &NullableLoginFlowMethodConfigFields{value: val, isSet: true}
}

func (v NullableLoginFlowMethodConfigFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginFlowMethodConfigFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


