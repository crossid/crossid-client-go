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
	"time"
)

// OAuth2HandledConsentRequest struct for OAuth2HandledConsentRequest
type OAuth2HandledConsentRequest struct {
	ClientId *string `json:"clientId,omitempty"`
	GrantAccessTokenAudience *[]string `json:"grant_access_token_audience,omitempty"`
	GrantScope *[]string `json:"grant_scope,omitempty"`
	HandledAt NullableTime `json:"handled_at,omitempty"`
	Remember *bool `json:"remember,omitempty"`
	RememberFor *int32 `json:"remember_for,omitempty"`
	Subject *string `json:"subject,omitempty"`
}

// NewOAuth2HandledConsentRequest instantiates a new OAuth2HandledConsentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2HandledConsentRequest() *OAuth2HandledConsentRequest {
	this := OAuth2HandledConsentRequest{}
	return &this
}

// NewOAuth2HandledConsentRequestWithDefaults instantiates a new OAuth2HandledConsentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2HandledConsentRequestWithDefaults() *OAuth2HandledConsentRequest {
	this := OAuth2HandledConsentRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OAuth2HandledConsentRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2HandledConsentRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OAuth2HandledConsentRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OAuth2HandledConsentRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetGrantAccessTokenAudience returns the GrantAccessTokenAudience field value if set, zero value otherwise.
func (o *OAuth2HandledConsentRequest) GetGrantAccessTokenAudience() []string {
	if o == nil || o.GrantAccessTokenAudience == nil {
		var ret []string
		return ret
	}
	return *o.GrantAccessTokenAudience
}

// GetGrantAccessTokenAudienceOk returns a tuple with the GrantAccessTokenAudience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2HandledConsentRequest) GetGrantAccessTokenAudienceOk() (*[]string, bool) {
	if o == nil || o.GrantAccessTokenAudience == nil {
		return nil, false
	}
	return o.GrantAccessTokenAudience, true
}

// HasGrantAccessTokenAudience returns a boolean if a field has been set.
func (o *OAuth2HandledConsentRequest) HasGrantAccessTokenAudience() bool {
	if o != nil && o.GrantAccessTokenAudience != nil {
		return true
	}

	return false
}

// SetGrantAccessTokenAudience gets a reference to the given []string and assigns it to the GrantAccessTokenAudience field.
func (o *OAuth2HandledConsentRequest) SetGrantAccessTokenAudience(v []string) {
	o.GrantAccessTokenAudience = &v
}

// GetGrantScope returns the GrantScope field value if set, zero value otherwise.
func (o *OAuth2HandledConsentRequest) GetGrantScope() []string {
	if o == nil || o.GrantScope == nil {
		var ret []string
		return ret
	}
	return *o.GrantScope
}

// GetGrantScopeOk returns a tuple with the GrantScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2HandledConsentRequest) GetGrantScopeOk() (*[]string, bool) {
	if o == nil || o.GrantScope == nil {
		return nil, false
	}
	return o.GrantScope, true
}

// HasGrantScope returns a boolean if a field has been set.
func (o *OAuth2HandledConsentRequest) HasGrantScope() bool {
	if o != nil && o.GrantScope != nil {
		return true
	}

	return false
}

// SetGrantScope gets a reference to the given []string and assigns it to the GrantScope field.
func (o *OAuth2HandledConsentRequest) SetGrantScope(v []string) {
	o.GrantScope = &v
}

// GetHandledAt returns the HandledAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2HandledConsentRequest) GetHandledAt() time.Time {
	if o == nil || o.HandledAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.HandledAt.Get()
}

// GetHandledAtOk returns a tuple with the HandledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2HandledConsentRequest) GetHandledAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HandledAt.Get(), o.HandledAt.IsSet()
}

// HasHandledAt returns a boolean if a field has been set.
func (o *OAuth2HandledConsentRequest) HasHandledAt() bool {
	if o != nil && o.HandledAt.IsSet() {
		return true
	}

	return false
}

// SetHandledAt gets a reference to the given NullableTime and assigns it to the HandledAt field.
func (o *OAuth2HandledConsentRequest) SetHandledAt(v time.Time) {
	o.HandledAt.Set(&v)
}
// SetHandledAtNil sets the value for HandledAt to be an explicit nil
func (o *OAuth2HandledConsentRequest) SetHandledAtNil() {
	o.HandledAt.Set(nil)
}

// UnsetHandledAt ensures that no value is present for HandledAt, not even an explicit nil
func (o *OAuth2HandledConsentRequest) UnsetHandledAt() {
	o.HandledAt.Unset()
}

// GetRemember returns the Remember field value if set, zero value otherwise.
func (o *OAuth2HandledConsentRequest) GetRemember() bool {
	if o == nil || o.Remember == nil {
		var ret bool
		return ret
	}
	return *o.Remember
}

// GetRememberOk returns a tuple with the Remember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2HandledConsentRequest) GetRememberOk() (*bool, bool) {
	if o == nil || o.Remember == nil {
		return nil, false
	}
	return o.Remember, true
}

// HasRemember returns a boolean if a field has been set.
func (o *OAuth2HandledConsentRequest) HasRemember() bool {
	if o != nil && o.Remember != nil {
		return true
	}

	return false
}

// SetRemember gets a reference to the given bool and assigns it to the Remember field.
func (o *OAuth2HandledConsentRequest) SetRemember(v bool) {
	o.Remember = &v
}

// GetRememberFor returns the RememberFor field value if set, zero value otherwise.
func (o *OAuth2HandledConsentRequest) GetRememberFor() int32 {
	if o == nil || o.RememberFor == nil {
		var ret int32
		return ret
	}
	return *o.RememberFor
}

// GetRememberForOk returns a tuple with the RememberFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2HandledConsentRequest) GetRememberForOk() (*int32, bool) {
	if o == nil || o.RememberFor == nil {
		return nil, false
	}
	return o.RememberFor, true
}

// HasRememberFor returns a boolean if a field has been set.
func (o *OAuth2HandledConsentRequest) HasRememberFor() bool {
	if o != nil && o.RememberFor != nil {
		return true
	}

	return false
}

// SetRememberFor gets a reference to the given int32 and assigns it to the RememberFor field.
func (o *OAuth2HandledConsentRequest) SetRememberFor(v int32) {
	o.RememberFor = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *OAuth2HandledConsentRequest) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2HandledConsentRequest) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *OAuth2HandledConsentRequest) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *OAuth2HandledConsentRequest) SetSubject(v string) {
	o.Subject = &v
}

func (o OAuth2HandledConsentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.GrantAccessTokenAudience != nil {
		toSerialize["grant_access_token_audience"] = o.GrantAccessTokenAudience
	}
	if o.GrantScope != nil {
		toSerialize["grant_scope"] = o.GrantScope
	}
	if o.HandledAt.IsSet() {
		toSerialize["handled_at"] = o.HandledAt.Get()
	}
	if o.Remember != nil {
		toSerialize["remember"] = o.Remember
	}
	if o.RememberFor != nil {
		toSerialize["remember_for"] = o.RememberFor
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	return json.Marshal(toSerialize)
}

type NullableOAuth2HandledConsentRequest struct {
	value *OAuth2HandledConsentRequest
	isSet bool
}

func (v NullableOAuth2HandledConsentRequest) Get() *OAuth2HandledConsentRequest {
	return v.value
}

func (v *NullableOAuth2HandledConsentRequest) Set(val *OAuth2HandledConsentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2HandledConsentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2HandledConsentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2HandledConsentRequest(val *OAuth2HandledConsentRequest) *NullableOAuth2HandledConsentRequest {
	return &NullableOAuth2HandledConsentRequest{value: val, isSet: true}
}

func (v NullableOAuth2HandledConsentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2HandledConsentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

