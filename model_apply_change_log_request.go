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

// ApplyChangeLogRequest struct for ApplyChangeLogRequest
type ApplyChangeLogRequest struct {
	AdditionalProperties map[string]interface{}
}

type _ApplyChangeLogRequest ApplyChangeLogRequest

// NewApplyChangeLogRequest instantiates a new ApplyChangeLogRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyChangeLogRequest() *ApplyChangeLogRequest {
	this := ApplyChangeLogRequest{}
	return &this
}

// NewApplyChangeLogRequestWithDefaults instantiates a new ApplyChangeLogRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyChangeLogRequestWithDefaults() *ApplyChangeLogRequest {
	this := ApplyChangeLogRequest{}
	return &this
}

func (o ApplyChangeLogRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplyChangeLogRequest) UnmarshalJSON(bytes []byte) (err error) {
	varApplyChangeLogRequest := _ApplyChangeLogRequest{}

	if err = json.Unmarshal(bytes, &varApplyChangeLogRequest); err == nil {
		*o = ApplyChangeLogRequest(varApplyChangeLogRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplyChangeLogRequest struct {
	value *ApplyChangeLogRequest
	isSet bool
}

func (v NullableApplyChangeLogRequest) Get() *ApplyChangeLogRequest {
	return v.value
}

func (v *NullableApplyChangeLogRequest) Set(val *ApplyChangeLogRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyChangeLogRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyChangeLogRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyChangeLogRequest(val *ApplyChangeLogRequest) *NullableApplyChangeLogRequest {
	return &NullableApplyChangeLogRequest{value: val, isSet: true}
}

func (v NullableApplyChangeLogRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyChangeLogRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


