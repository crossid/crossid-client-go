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

// FormTextAttributes struct for FormTextAttributes
type FormTextAttributes struct {
	Text string `json:"text"`
}

// NewFormTextAttributes instantiates a new FormTextAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormTextAttributes(text string) *FormTextAttributes {
	this := FormTextAttributes{}
	this.Text = text
	return &this
}

// NewFormTextAttributesWithDefaults instantiates a new FormTextAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormTextAttributesWithDefaults() *FormTextAttributes {
	this := FormTextAttributes{}
	return &this
}

// GetText returns the Text field value
func (o *FormTextAttributes) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *FormTextAttributes) GetTextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *FormTextAttributes) SetText(v string) {
	o.Text = v
}

func (o FormTextAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableFormTextAttributes struct {
	value *FormTextAttributes
	isSet bool
}

func (v NullableFormTextAttributes) Get() *FormTextAttributes {
	return v.value
}

func (v *NullableFormTextAttributes) Set(val *FormTextAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableFormTextAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableFormTextAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormTextAttributes(val *FormTextAttributes) *NullableFormTextAttributes {
	return &NullableFormTextAttributes{value: val, isSet: true}
}

func (v NullableFormTextAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormTextAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


