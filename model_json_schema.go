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

// JSONSchema struct for JSONSchema
type JSONSchema struct {
	// externally added unique identifier.
	ExternalId *string `json:"externalId,omitempty"`
	// unique identifier of the schema.
	Id *string `json:"id,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _JSONSchema JSONSchema

// NewJSONSchema instantiates a new JSONSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSONSchema() *JSONSchema {
	this := JSONSchema{}
	return &this
}

// NewJSONSchemaWithDefaults instantiates a new JSONSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSONSchemaWithDefaults() *JSONSchema {
	this := JSONSchema{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *JSONSchema) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONSchema) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *JSONSchema) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *JSONSchema) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *JSONSchema) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONSchema) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *JSONSchema) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *JSONSchema) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *JSONSchema) GetMeta() Meta {
	if o == nil || o.Meta == nil {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONSchema) GetMetaOk() (*Meta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *JSONSchema) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *JSONSchema) SetMeta(v Meta) {
	o.Meta = &v
}

func (o JSONSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *JSONSchema) UnmarshalJSON(bytes []byte) (err error) {
	varJSONSchema := _JSONSchema{}

	if err = json.Unmarshal(bytes, &varJSONSchema); err == nil {
		*o = JSONSchema(varJSONSchema)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJSONSchema struct {
	value *JSONSchema
	isSet bool
}

func (v NullableJSONSchema) Get() *JSONSchema {
	return v.value
}

func (v *NullableJSONSchema) Set(val *JSONSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableJSONSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableJSONSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSONSchema(val *JSONSchema) *NullableJSONSchema {
	return &NullableJSONSchema{value: val, isSet: true}
}

func (v NullableJSONSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSONSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


