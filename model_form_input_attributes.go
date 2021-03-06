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

// FormInputAttributes struct for FormInputAttributes
type FormInputAttributes struct {
	Autocomplete *string `json:"autocomplete,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	Label *string `json:"label,omitempty"`
	MaxLength *float32 `json:"maxLength,omitempty"`
	MinLength *float32 `json:"minLength,omitempty"`
	Name string `json:"name"`
	Pattern *string `json:"pattern,omitempty"`
	Placeholder *string `json:"placeholder,omitempty"`
	Required *bool `json:"required,omitempty"`
	Transient *bool `json:"transient,omitempty"`
	Type string `json:"type"`
}

// NewFormInputAttributes instantiates a new FormInputAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormInputAttributes(name string, type_ string) *FormInputAttributes {
	this := FormInputAttributes{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewFormInputAttributesWithDefaults instantiates a new FormInputAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormInputAttributesWithDefaults() *FormInputAttributes {
	this := FormInputAttributes{}
	return &this
}

// GetAutocomplete returns the Autocomplete field value if set, zero value otherwise.
func (o *FormInputAttributes) GetAutocomplete() string {
	if o == nil || o.Autocomplete == nil {
		var ret string
		return ret
	}
	return *o.Autocomplete
}

// GetAutocompleteOk returns a tuple with the Autocomplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInputAttributes) GetAutocompleteOk() (*string, bool) {
	if o == nil || o.Autocomplete == nil {
		return nil, false
	}
	return o.Autocomplete, true
}

// HasAutocomplete returns a boolean if a field has been set.
func (o *FormInputAttributes) HasAutocomplete() bool {
	if o != nil && o.Autocomplete != nil {
		return true
	}

	return false
}

// SetAutocomplete gets a reference to the given string and assigns it to the Autocomplete field.
func (o *FormInputAttributes) SetAutocomplete(v string) {
	o.Autocomplete = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *FormInputAttributes) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInputAttributes) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *FormInputAttributes) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *FormInputAttributes) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FormInputAttributes) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInputAttributes) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FormInputAttributes) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FormInputAttributes) SetLabel(v string) {
	o.Label = &v
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise.
func (o *FormInputAttributes) GetMaxLength() float32 {
	if o == nil || o.MaxLength == nil {
		var ret float32
		return ret
	}
	return *o.MaxLength
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInputAttributes) GetMaxLengthOk() (*float32, bool) {
	if o == nil || o.MaxLength == nil {
		return nil, false
	}
	return o.MaxLength, true
}

// HasMaxLength returns a boolean if a field has been set.
func (o *FormInputAttributes) HasMaxLength() bool {
	if o != nil && o.MaxLength != nil {
		return true
	}

	return false
}

// SetMaxLength gets a reference to the given float32 and assigns it to the MaxLength field.
func (o *FormInputAttributes) SetMaxLength(v float32) {
	o.MaxLength = &v
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *FormInputAttributes) GetMinLength() float32 {
	if o == nil || o.MinLength == nil {
		var ret float32
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInputAttributes) GetMinLengthOk() (*float32, bool) {
	if o == nil || o.MinLength == nil {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *FormInputAttributes) HasMinLength() bool {
	if o != nil && o.MinLength != nil {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given float32 and assigns it to the MinLength field.
func (o *FormInputAttributes) SetMinLength(v float32) {
	o.MinLength = &v
}

// GetName returns the Name field value
func (o *FormInputAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormInputAttributes) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FormInputAttributes) SetName(v string) {
	o.Name = v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *FormInputAttributes) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInputAttributes) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *FormInputAttributes) HasPattern() bool {
	if o != nil && o.Pattern != nil {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *FormInputAttributes) SetPattern(v string) {
	o.Pattern = &v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *FormInputAttributes) GetPlaceholder() string {
	if o == nil || o.Placeholder == nil {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInputAttributes) GetPlaceholderOk() (*string, bool) {
	if o == nil || o.Placeholder == nil {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *FormInputAttributes) HasPlaceholder() bool {
	if o != nil && o.Placeholder != nil {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *FormInputAttributes) SetPlaceholder(v string) {
	o.Placeholder = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *FormInputAttributes) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInputAttributes) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *FormInputAttributes) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *FormInputAttributes) SetRequired(v bool) {
	o.Required = &v
}

// GetTransient returns the Transient field value if set, zero value otherwise.
func (o *FormInputAttributes) GetTransient() bool {
	if o == nil || o.Transient == nil {
		var ret bool
		return ret
	}
	return *o.Transient
}

// GetTransientOk returns a tuple with the Transient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInputAttributes) GetTransientOk() (*bool, bool) {
	if o == nil || o.Transient == nil {
		return nil, false
	}
	return o.Transient, true
}

// HasTransient returns a boolean if a field has been set.
func (o *FormInputAttributes) HasTransient() bool {
	if o != nil && o.Transient != nil {
		return true
	}

	return false
}

// SetTransient gets a reference to the given bool and assigns it to the Transient field.
func (o *FormInputAttributes) SetTransient(v bool) {
	o.Transient = &v
}

// GetType returns the Type field value
func (o *FormInputAttributes) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FormInputAttributes) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FormInputAttributes) SetType(v string) {
	o.Type = v
}

func (o FormInputAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Autocomplete != nil {
		toSerialize["autocomplete"] = o.Autocomplete
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.MaxLength != nil {
		toSerialize["maxLength"] = o.MaxLength
	}
	if o.MinLength != nil {
		toSerialize["minLength"] = o.MinLength
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if o.Placeholder != nil {
		toSerialize["placeholder"] = o.Placeholder
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Transient != nil {
		toSerialize["transient"] = o.Transient
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableFormInputAttributes struct {
	value *FormInputAttributes
	isSet bool
}

func (v NullableFormInputAttributes) Get() *FormInputAttributes {
	return v.value
}

func (v *NullableFormInputAttributes) Set(val *FormInputAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableFormInputAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableFormInputAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormInputAttributes(val *FormInputAttributes) *NullableFormInputAttributes {
	return &NullableFormInputAttributes{value: val, isSet: true}
}

func (v NullableFormInputAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormInputAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


