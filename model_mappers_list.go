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

// MappersList struct for MappersList
type MappersList struct {
	ItemsPerPage *int32 `json:"itemsPerPage,omitempty"`
	StartIndex *int64 `json:"startIndex,omitempty"`
	TotalResults *int64 `json:"totalResults,omitempty"`
	Resources []Mapper `json:"Resources"`
}

// NewMappersList instantiates a new MappersList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMappersList(resources []Mapper) *MappersList {
	this := MappersList{}
	this.Resources = resources
	return &this
}

// NewMappersListWithDefaults instantiates a new MappersList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMappersListWithDefaults() *MappersList {
	this := MappersList{}
	return &this
}

// GetItemsPerPage returns the ItemsPerPage field value if set, zero value otherwise.
func (o *MappersList) GetItemsPerPage() int32 {
	if o == nil || o.ItemsPerPage == nil {
		var ret int32
		return ret
	}
	return *o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MappersList) GetItemsPerPageOk() (*int32, bool) {
	if o == nil || o.ItemsPerPage == nil {
		return nil, false
	}
	return o.ItemsPerPage, true
}

// HasItemsPerPage returns a boolean if a field has been set.
func (o *MappersList) HasItemsPerPage() bool {
	if o != nil && o.ItemsPerPage != nil {
		return true
	}

	return false
}

// SetItemsPerPage gets a reference to the given int32 and assigns it to the ItemsPerPage field.
func (o *MappersList) SetItemsPerPage(v int32) {
	o.ItemsPerPage = &v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *MappersList) GetStartIndex() int64 {
	if o == nil || o.StartIndex == nil {
		var ret int64
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MappersList) GetStartIndexOk() (*int64, bool) {
	if o == nil || o.StartIndex == nil {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *MappersList) HasStartIndex() bool {
	if o != nil && o.StartIndex != nil {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given int64 and assigns it to the StartIndex field.
func (o *MappersList) SetStartIndex(v int64) {
	o.StartIndex = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *MappersList) GetTotalResults() int64 {
	if o == nil || o.TotalResults == nil {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MappersList) GetTotalResultsOk() (*int64, bool) {
	if o == nil || o.TotalResults == nil {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *MappersList) HasTotalResults() bool {
	if o != nil && o.TotalResults != nil {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *MappersList) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetResources returns the Resources field value
func (o *MappersList) GetResources() []Mapper {
	if o == nil {
		var ret []Mapper
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *MappersList) GetResourcesOk() (*[]Mapper, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *MappersList) SetResources(v []Mapper) {
	o.Resources = v
}

func (o MappersList) MarshalJSON() ([]byte, error) {
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

type NullableMappersList struct {
	value *MappersList
	isSet bool
}

func (v NullableMappersList) Get() *MappersList {
	return v.value
}

func (v *NullableMappersList) Set(val *MappersList) {
	v.value = val
	v.isSet = true
}

func (v NullableMappersList) IsSet() bool {
	return v.isSet
}

func (v *NullableMappersList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMappersList(val *MappersList) *NullableMappersList {
	return &NullableMappersList{value: val, isSet: true}
}

func (v NullableMappersList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMappersList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


