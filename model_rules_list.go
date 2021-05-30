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

// RulesList struct for RulesList
type RulesList struct {
	ItemsPerPage *int32 `json:"itemsPerPage,omitempty"`
	StartIndex *int64 `json:"startIndex,omitempty"`
	TotalResults *int64 `json:"totalResults,omitempty"`
	Resources []Rule `json:"Resources"`
}

// NewRulesList instantiates a new RulesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRulesList(resources []Rule, ) *RulesList {
	this := RulesList{}
	this.Resources = resources
	return &this
}

// NewRulesListWithDefaults instantiates a new RulesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRulesListWithDefaults() *RulesList {
	this := RulesList{}
	return &this
}

// GetItemsPerPage returns the ItemsPerPage field value if set, zero value otherwise.
func (o *RulesList) GetItemsPerPage() int32 {
	if o == nil || o.ItemsPerPage == nil {
		var ret int32
		return ret
	}
	return *o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesList) GetItemsPerPageOk() (*int32, bool) {
	if o == nil || o.ItemsPerPage == nil {
		return nil, false
	}
	return o.ItemsPerPage, true
}

// HasItemsPerPage returns a boolean if a field has been set.
func (o *RulesList) HasItemsPerPage() bool {
	if o != nil && o.ItemsPerPage != nil {
		return true
	}

	return false
}

// SetItemsPerPage gets a reference to the given int32 and assigns it to the ItemsPerPage field.
func (o *RulesList) SetItemsPerPage(v int32) {
	o.ItemsPerPage = &v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *RulesList) GetStartIndex() int64 {
	if o == nil || o.StartIndex == nil {
		var ret int64
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesList) GetStartIndexOk() (*int64, bool) {
	if o == nil || o.StartIndex == nil {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *RulesList) HasStartIndex() bool {
	if o != nil && o.StartIndex != nil {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given int64 and assigns it to the StartIndex field.
func (o *RulesList) SetStartIndex(v int64) {
	o.StartIndex = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *RulesList) GetTotalResults() int64 {
	if o == nil || o.TotalResults == nil {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesList) GetTotalResultsOk() (*int64, bool) {
	if o == nil || o.TotalResults == nil {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *RulesList) HasTotalResults() bool {
	if o != nil && o.TotalResults != nil {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *RulesList) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetResources returns the Resources field value
func (o *RulesList) GetResources() []Rule {
	if o == nil  {
		var ret []Rule
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *RulesList) GetResourcesOk() (*[]Rule, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *RulesList) SetResources(v []Rule) {
	o.Resources = v
}

func (o RulesList) MarshalJSON() ([]byte, error) {
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

type NullableRulesList struct {
	value *RulesList
	isSet bool
}

func (v NullableRulesList) Get() *RulesList {
	return v.value
}

func (v *NullableRulesList) Set(val *RulesList) {
	v.value = val
	v.isSet = true
}

func (v NullableRulesList) IsSet() bool {
	return v.isSet
}

func (v *NullableRulesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRulesList(val *RulesList) *NullableRulesList {
	return &NullableRulesList{value: val, isSet: true}
}

func (v NullableRulesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRulesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

