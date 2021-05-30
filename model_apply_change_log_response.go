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

// ApplyChangeLogResponse struct for ApplyChangeLogResponse
type ApplyChangeLogResponse struct {
	LogId *string `json:"logId,omitempty"`
	Report *ApplyChangesReport `json:"report,omitempty"`
}

// NewApplyChangeLogResponse instantiates a new ApplyChangeLogResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyChangeLogResponse() *ApplyChangeLogResponse {
	this := ApplyChangeLogResponse{}
	return &this
}

// NewApplyChangeLogResponseWithDefaults instantiates a new ApplyChangeLogResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyChangeLogResponseWithDefaults() *ApplyChangeLogResponse {
	this := ApplyChangeLogResponse{}
	return &this
}

// GetLogId returns the LogId field value if set, zero value otherwise.
func (o *ApplyChangeLogResponse) GetLogId() string {
	if o == nil || o.LogId == nil {
		var ret string
		return ret
	}
	return *o.LogId
}

// GetLogIdOk returns a tuple with the LogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyChangeLogResponse) GetLogIdOk() (*string, bool) {
	if o == nil || o.LogId == nil {
		return nil, false
	}
	return o.LogId, true
}

// HasLogId returns a boolean if a field has been set.
func (o *ApplyChangeLogResponse) HasLogId() bool {
	if o != nil && o.LogId != nil {
		return true
	}

	return false
}

// SetLogId gets a reference to the given string and assigns it to the LogId field.
func (o *ApplyChangeLogResponse) SetLogId(v string) {
	o.LogId = &v
}

// GetReport returns the Report field value if set, zero value otherwise.
func (o *ApplyChangeLogResponse) GetReport() ApplyChangesReport {
	if o == nil || o.Report == nil {
		var ret ApplyChangesReport
		return ret
	}
	return *o.Report
}

// GetReportOk returns a tuple with the Report field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyChangeLogResponse) GetReportOk() (*ApplyChangesReport, bool) {
	if o == nil || o.Report == nil {
		return nil, false
	}
	return o.Report, true
}

// HasReport returns a boolean if a field has been set.
func (o *ApplyChangeLogResponse) HasReport() bool {
	if o != nil && o.Report != nil {
		return true
	}

	return false
}

// SetReport gets a reference to the given ApplyChangesReport and assigns it to the Report field.
func (o *ApplyChangeLogResponse) SetReport(v ApplyChangesReport) {
	o.Report = &v
}

func (o ApplyChangeLogResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LogId != nil {
		toSerialize["logId"] = o.LogId
	}
	if o.Report != nil {
		toSerialize["report"] = o.Report
	}
	return json.Marshal(toSerialize)
}

type NullableApplyChangeLogResponse struct {
	value *ApplyChangeLogResponse
	isSet bool
}

func (v NullableApplyChangeLogResponse) Get() *ApplyChangeLogResponse {
	return v.value
}

func (v *NullableApplyChangeLogResponse) Set(val *ApplyChangeLogResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyChangeLogResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyChangeLogResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyChangeLogResponse(val *ApplyChangeLogResponse) *NullableApplyChangeLogResponse {
	return &NullableApplyChangeLogResponse{value: val, isSet: true}
}

func (v NullableApplyChangeLogResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyChangeLogResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


