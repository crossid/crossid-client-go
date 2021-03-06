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

// Commit struct for Commit
type Commit struct {
	By *string `json:"by,omitempty"`
	Correlation *string `json:"correlation,omitempty"`
	Id *string `json:"id,omitempty"`
	Meta *AppMeta `json:"meta,omitempty"`
	Reason *string `json:"reason,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
}

// NewCommit instantiates a new Commit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommit() *Commit {
	this := Commit{}
	return &this
}

// NewCommitWithDefaults instantiates a new Commit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitWithDefaults() *Commit {
	this := Commit{}
	return &this
}

// GetBy returns the By field value if set, zero value otherwise.
func (o *Commit) GetBy() string {
	if o == nil || o.By == nil {
		var ret string
		return ret
	}
	return *o.By
}

// GetByOk returns a tuple with the By field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commit) GetByOk() (*string, bool) {
	if o == nil || o.By == nil {
		return nil, false
	}
	return o.By, true
}

// HasBy returns a boolean if a field has been set.
func (o *Commit) HasBy() bool {
	if o != nil && o.By != nil {
		return true
	}

	return false
}

// SetBy gets a reference to the given string and assigns it to the By field.
func (o *Commit) SetBy(v string) {
	o.By = &v
}

// GetCorrelation returns the Correlation field value if set, zero value otherwise.
func (o *Commit) GetCorrelation() string {
	if o == nil || o.Correlation == nil {
		var ret string
		return ret
	}
	return *o.Correlation
}

// GetCorrelationOk returns a tuple with the Correlation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commit) GetCorrelationOk() (*string, bool) {
	if o == nil || o.Correlation == nil {
		return nil, false
	}
	return o.Correlation, true
}

// HasCorrelation returns a boolean if a field has been set.
func (o *Commit) HasCorrelation() bool {
	if o != nil && o.Correlation != nil {
		return true
	}

	return false
}

// SetCorrelation gets a reference to the given string and assigns it to the Correlation field.
func (o *Commit) SetCorrelation(v string) {
	o.Correlation = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Commit) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commit) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Commit) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Commit) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Commit) GetMeta() AppMeta {
	if o == nil || o.Meta == nil {
		var ret AppMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commit) GetMetaOk() (*AppMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Commit) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given AppMeta and assigns it to the Meta field.
func (o *Commit) SetMeta(v AppMeta) {
	o.Meta = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *Commit) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commit) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *Commit) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *Commit) SetReason(v string) {
	o.Reason = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *Commit) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commit) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *Commit) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *Commit) SetRequestId(v string) {
	o.RequestId = &v
}

func (o Commit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.By != nil {
		toSerialize["by"] = o.By
	}
	if o.Correlation != nil {
		toSerialize["correlation"] = o.Correlation
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.RequestId != nil {
		toSerialize["requestId"] = o.RequestId
	}
	return json.Marshal(toSerialize)
}

type NullableCommit struct {
	value *Commit
	isSet bool
}

func (v NullableCommit) Get() *Commit {
	return v.value
}

func (v *NullableCommit) Set(val *Commit) {
	v.value = val
	v.isSet = true
}

func (v NullableCommit) IsSet() bool {
	return v.isSet
}

func (v *NullableCommit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommit(val *Commit) *NullableCommit {
	return &NullableCommit{value: val, isSet: true}
}

func (v NullableCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


