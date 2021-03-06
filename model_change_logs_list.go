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

// ChangeLogsList struct for ChangeLogsList
type ChangeLogsList struct {
	ItemsPerPage *int32 `json:"itemsPerPage,omitempty"`
	StartIndex *int64 `json:"startIndex,omitempty"`
	TotalResults *int64 `json:"totalResults,omitempty"`
	Resources []ChangeLog `json:"Resources"`
}

// NewChangeLogsList instantiates a new ChangeLogsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeLogsList(resources []ChangeLog) *ChangeLogsList {
	this := ChangeLogsList{}
	this.Resources = resources
	return &this
}

// NewChangeLogsListWithDefaults instantiates a new ChangeLogsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeLogsListWithDefaults() *ChangeLogsList {
	this := ChangeLogsList{}
	return &this
}

// GetItemsPerPage returns the ItemsPerPage field value if set, zero value otherwise.
func (o *ChangeLogsList) GetItemsPerPage() int32 {
	if o == nil || o.ItemsPerPage == nil {
		var ret int32
		return ret
	}
	return *o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeLogsList) GetItemsPerPageOk() (*int32, bool) {
	if o == nil || o.ItemsPerPage == nil {
		return nil, false
	}
	return o.ItemsPerPage, true
}

// HasItemsPerPage returns a boolean if a field has been set.
func (o *ChangeLogsList) HasItemsPerPage() bool {
	if o != nil && o.ItemsPerPage != nil {
		return true
	}

	return false
}

// SetItemsPerPage gets a reference to the given int32 and assigns it to the ItemsPerPage field.
func (o *ChangeLogsList) SetItemsPerPage(v int32) {
	o.ItemsPerPage = &v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *ChangeLogsList) GetStartIndex() int64 {
	if o == nil || o.StartIndex == nil {
		var ret int64
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeLogsList) GetStartIndexOk() (*int64, bool) {
	if o == nil || o.StartIndex == nil {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *ChangeLogsList) HasStartIndex() bool {
	if o != nil && o.StartIndex != nil {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given int64 and assigns it to the StartIndex field.
func (o *ChangeLogsList) SetStartIndex(v int64) {
	o.StartIndex = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *ChangeLogsList) GetTotalResults() int64 {
	if o == nil || o.TotalResults == nil {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeLogsList) GetTotalResultsOk() (*int64, bool) {
	if o == nil || o.TotalResults == nil {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *ChangeLogsList) HasTotalResults() bool {
	if o != nil && o.TotalResults != nil {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *ChangeLogsList) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetResources returns the Resources field value
func (o *ChangeLogsList) GetResources() []ChangeLog {
	if o == nil {
		var ret []ChangeLog
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *ChangeLogsList) GetResourcesOk() (*[]ChangeLog, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *ChangeLogsList) SetResources(v []ChangeLog) {
	o.Resources = v
}

func (o ChangeLogsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ItemsPerPage != nil {
		toSerialize["itemsPerPage"] = o.ItemsPerPage
	}
	if o.StartIndex != nil {
		toSerialize["startIndex"] = o.StartIndex
	}
	if o.TotalResults != nil {
		toSerialize["totalResults"] = o.TotalResults
	}
	if true {
		toSerialize["Resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullableChangeLogsList struct {
	value *ChangeLogsList
	isSet bool
}

func (v NullableChangeLogsList) Get() *ChangeLogsList {
	return v.value
}

func (v *NullableChangeLogsList) Set(val *ChangeLogsList) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeLogsList) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeLogsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeLogsList(val *ChangeLogsList) *NullableChangeLogsList {
	return &NullableChangeLogsList{value: val, isSet: true}
}

func (v NullableChangeLogsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeLogsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


