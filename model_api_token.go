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

// ApiToken struct for ApiToken
type ApiToken struct {
	Active *bool `json:"active,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	Id *string `json:"id,omitempty"`
	IdentityId *string `json:"identityId,omitempty"`
	IssuedAt *time.Time `json:"issuedAt,omitempty"`
	Meta *ApiTokenMeta `json:"meta,omitempty"`
	SignedToken *string `json:"signedToken,omitempty"`
}

// NewApiToken instantiates a new ApiToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiToken() *ApiToken {
	this := ApiToken{}
	return &this
}

// NewApiTokenWithDefaults instantiates a new ApiToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTokenWithDefaults() *ApiToken {
	this := ApiToken{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ApiToken) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ApiToken) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ApiToken) SetActive(v bool) {
	o.Active = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ApiToken) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ApiToken) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ApiToken) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *ApiToken) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ApiToken) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *ApiToken) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiToken) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiToken) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiToken) SetId(v string) {
	o.Id = &v
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *ApiToken) GetIdentityId() string {
	if o == nil || o.IdentityId == nil {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetIdentityIdOk() (*string, bool) {
	if o == nil || o.IdentityId == nil {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *ApiToken) HasIdentityId() bool {
	if o != nil && o.IdentityId != nil {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *ApiToken) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetIssuedAt returns the IssuedAt field value if set, zero value otherwise.
func (o *ApiToken) GetIssuedAt() time.Time {
	if o == nil || o.IssuedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.IssuedAt
}

// GetIssuedAtOk returns a tuple with the IssuedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetIssuedAtOk() (*time.Time, bool) {
	if o == nil || o.IssuedAt == nil {
		return nil, false
	}
	return o.IssuedAt, true
}

// HasIssuedAt returns a boolean if a field has been set.
func (o *ApiToken) HasIssuedAt() bool {
	if o != nil && o.IssuedAt != nil {
		return true
	}

	return false
}

// SetIssuedAt gets a reference to the given time.Time and assigns it to the IssuedAt field.
func (o *ApiToken) SetIssuedAt(v time.Time) {
	o.IssuedAt = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ApiToken) GetMeta() ApiTokenMeta {
	if o == nil || o.Meta == nil {
		var ret ApiTokenMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetMetaOk() (*ApiTokenMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ApiToken) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiTokenMeta and assigns it to the Meta field.
func (o *ApiToken) SetMeta(v ApiTokenMeta) {
	o.Meta = &v
}

// GetSignedToken returns the SignedToken field value if set, zero value otherwise.
func (o *ApiToken) GetSignedToken() string {
	if o == nil || o.SignedToken == nil {
		var ret string
		return ret
	}
	return *o.SignedToken
}

// GetSignedTokenOk returns a tuple with the SignedToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetSignedTokenOk() (*string, bool) {
	if o == nil || o.SignedToken == nil {
		return nil, false
	}
	return o.SignedToken, true
}

// HasSignedToken returns a boolean if a field has been set.
func (o *ApiToken) HasSignedToken() bool {
	if o != nil && o.SignedToken != nil {
		return true
	}

	return false
}

// SetSignedToken gets a reference to the given string and assigns it to the SignedToken field.
func (o *ApiToken) SetSignedToken(v string) {
	o.SignedToken = &v
}

func (o ApiToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.ExpiresAt != nil {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IdentityId != nil {
		toSerialize["identityId"] = o.IdentityId
	}
	if o.IssuedAt != nil {
		toSerialize["issuedAt"] = o.IssuedAt
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.SignedToken != nil {
		toSerialize["signedToken"] = o.SignedToken
	}
	return json.Marshal(toSerialize)
}

type NullableApiToken struct {
	value *ApiToken
	isSet bool
}

func (v NullableApiToken) Get() *ApiToken {
	return v.value
}

func (v *NullableApiToken) Set(val *ApiToken) {
	v.value = val
	v.isSet = true
}

func (v NullableApiToken) IsSet() bool {
	return v.isSet
}

func (v *NullableApiToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiToken(val *ApiToken) *NullableApiToken {
	return &NullableApiToken{value: val, isSet: true}
}

func (v NullableApiToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

