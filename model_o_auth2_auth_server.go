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

// OAuth2AuthServer struct for OAuth2AuthServer
type OAuth2AuthServer struct {
	AccessTokenLifespan *int64 `json:"accessTokenLifespan,omitempty"`
	AccessTokenStrategy *string `json:"accessTokenStrategy,omitempty"`
	Active *bool `json:"active,omitempty"`
	ConsentRequestMaxAge *int64 `json:"consentRequestMaxAge,omitempty"`
	CookieSameSiteLegacyWorkaround *bool `json:"cookieSameSiteLegacyWorkaround,omitempty"`
	CookieSameSiteMode *string `json:"cookieSameSiteMode,omitempty"`
	Debug *bool `json:"debug,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ExternalId *string `json:"externalId,omitempty"`
	ForceHTTP *bool `json:"forceHTTP,omitempty"`
	Id *string `json:"id,omitempty"`
	IdTokenLifespan *int64 `json:"idTokenLifespan,omitempty"`
	Meta *ApiTokenMeta `json:"meta,omitempty"`
	RefreshTokenLifespan *int64 `json:"refreshTokenLifespan,omitempty"`
	ScopeStrategy *string `json:"scopeStrategy,omitempty"`
	SubjectIdentifierAlgorithmSalt *string `json:"subjectIdentifierAlgorithmSalt,omitempty"`
	SubjectTypesSupported *[]string `json:"subjectTypesSupported,omitempty"`
	WellKnownKeys *[]string `json:"wellKnownKeys,omitempty"`
}

// NewOAuth2AuthServer instantiates a new OAuth2AuthServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2AuthServer() *OAuth2AuthServer {
	this := OAuth2AuthServer{}
	return &this
}

// NewOAuth2AuthServerWithDefaults instantiates a new OAuth2AuthServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2AuthServerWithDefaults() *OAuth2AuthServer {
	this := OAuth2AuthServer{}
	return &this
}

// GetAccessTokenLifespan returns the AccessTokenLifespan field value if set, zero value otherwise.
func (o *OAuth2AuthServer) GetAccessTokenLifespan() int64 {
	if o == nil || o.AccessTokenLifespan == nil {
		var ret int64
		return ret
	}
	return *o.AccessTokenLifespan
}

// GetAccessTokenLifespanOk returns a tuple with the AccessTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AuthServer) GetAccessTokenLifespanOk() (*int64, bool) {
	if o == nil || o.AccessTokenLifespan == nil {
		return nil, false
	}
	return o.AccessTokenLifespan, true
}

// HasAccessTokenLifespan returns a boolean if a field has been set.
func (o *OAuth2AuthServer) HasAccessTokenLifespan() bool {
	if o != nil && o.AccessTokenLifespan != nil {
		return true
	}

	return false
}

// SetAccessTokenLifespan gets a reference to the given int64 and assigns it to the AccessTokenLifespan field.
func (o *OAuth2AuthServer) SetAccessTokenLifespan(v int64) {
	o.AccessTokenLifespan = &v
}

// GetAccessTokenStrategy returns the AccessTokenStrategy field value if set, zero value otherwise.
func (o *OAuth2AuthServer) GetAccessTokenStrategy() string {
	if o == nil || o.AccessTokenStrategy == nil {
		var ret string
		return ret
	}
	return *o.AccessTokenStrategy
}

// GetAccessTokenStrategyOk returns a tuple with the AccessTokenStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AuthServer) GetAccessTokenStrategyOk() (*string, bool) {
	if o == nil || o.AccessTokenStrategy == nil {
		return nil, false
	}
	return o.AccessTokenStrategy, true
}

// HasAccessTokenStrategy returns a boolean if a field has been set.
func (o *OAuth2AuthServer) HasAccessTokenStrategy() bool {
	if o != nil && o.AccessTokenStrategy != nil {
		return true
	}

	return false
}

// SetAccessTokenStrategy gets a reference to the given string and assigns it to the AccessTokenStrategy field.
func (o *OAuth2AuthServer) SetAccessTokenStrategy(v string) {
	o.AccessTokenStrategy = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *OAuth2AuthServer) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AuthServer) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *OAuth2AuthServer) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *OAuth2AuthServer) SetActive(v bool) {
	o.Active = &v
}

// GetConsentRequestMaxAge returns the ConsentRequestMaxAge field value if set, zero value otherwise.
func (o *OAuth2AuthServer) GetConsentRequestMaxAge() int64 {
	if o == nil || o.ConsentRequestMaxAge == nil {
		var ret int64
		return ret
	}
	return *o.ConsentRequestMaxAge
}

// GetConsentRequestMaxAgeOk returns a tuple with the ConsentRequestMaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AuthServer) GetConsentRequestMaxAgeOk() (*int64, bool) {
	if o == nil || o.ConsentRequestMaxAge == nil {
		return nil, false
	}
	return o.ConsentRequestMaxAge, true
}

// HasConsentRequestMaxAge returns a boolean if a field has been set.
func (o *OAuth2AuthServer) HasConsentRequestMaxAge() bool {
	if o != nil && o.ConsentRequestMaxAge != nil {
		return true
	}

	return false
}

// SetConsentRequestMaxAge gets a reference to the given int64 and assigns it to the ConsentRequestMaxAge field.
func (o *OAuth2AuthServer) SetConsentRequestMaxAge(v int64) {
	o.ConsentRequestMaxAge = &v
}

// GetCookieSameSiteLegacyWorkaround returns the CookieSameSiteLegacyWorkaround field value if set, zero value otherwise.
func (o *OAuth2AuthServer) GetCookieSameSiteLegacyWorkaround() bool {
	if o == nil || o.CookieSameSiteLegacyWorkaround == nil {
		var ret bool
		return ret
	}
	return *o.CookieSameSiteLegacyWorkaround
}

// GetCookieSameSiteLegacyWorkaroundOk returns a tuple with the CookieSameSiteLegacyWorkaround field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AuthServer) GetCookieSameSiteLegacyWorkaroundOk() (*bool, bool) {
	if o == nil || o.CookieSameSiteLegacyWorkaround == nil {
		return nil, false
	}
	return o.CookieSameSiteLegacyWorkaround, true
}

// HasCookieSameSiteLegacyWorkaround returns a boolean if a field has been set.
func (o *OAuth2AuthServer) HasCookieSameSiteLegacyWorkaround() bool {
	if o != nil && o.CookieSameSiteLegacyWorkaround != nil {
		return true
	}

	return false
}

// SetCookieSameSiteLegacyWorkaround gets a reference to the given bool and assigns it to the CookieSameSiteLegacyWorkaround field.
func (o *OAuth2AuthServer) SetCookieSameSiteLegacyWorkaround(v bool) {
	o.CookieSameSiteLegacyWorkaround = &v
}

// GetCookieSameSiteMode returns the CookieSameSiteMode field value if set, zero value otherwise.
func (o *OAuth2AuthServer) GetCookieSameSiteMode() string {
	if o == nil || o.CookieSameSiteMode == nil {
		var ret string
		return ret
	}
	return *o.CookieSameSiteMode
}

// GetCookieSameSiteModeOk returns a tuple with the CookieSameSiteMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AuthServer) GetCookieSameSiteModeOk() (*string, bool) {
	if o == nil || o.CookieSameSiteMode == nil {
		return nil, false
	}
	return o.CookieSameSiteMode, true
}

// HasCookieSameSiteMode returns a boolean if a field has been set.
func (o *OAuth2AuthServer) HasCookieSameSiteMode() bool {
	if o != nil && o.CookieSameSiteMode != nil {
		return true
	}

	return false
}

// SetCookieSameSiteMode gets a reference to the given string and assigns it to the CookieSameSiteMode field.
func (o *OAuth2AuthServer) SetCookieSameSiteMode(v string) {
	o.CookieSameSiteMode = &v
}

// GetDebug returns the Debug field value if set, zero value otherwise.
func (o *OAuth2AuthServer) GetDebug() bool {
	if o == nil || o.Debug == nil {
		var ret bool
		return ret
	}
	return *o.Debug
}

// GetDebugOk returns a tuple with the Debug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AuthServer) GetDebugOk() (*bool, bool) {
	if o == nil || o.Debug == nil {
		return nil, false
	}
	return o.Debug, true
}

// HasDebug returns a boolean if a field has been set.
func (o *OAuth2AuthServer) HasDebug() bool {
	if o != nil && o.Debug != nil {
		return true
	}

	return false
}

// SetDebug gets a reference to the given bool and assigns it to the Debug field.
func (o *OAuth2AuthServer) SetDebug(v bool) {
	o.Debug = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *OAuth2AuthServer) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AuthServer) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *OAuth2AuthServer) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *OAuth2AuthServer) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *OAuth2AuthServer) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AuthServer) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *OAuth2AuthServer) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *OAuth2AuthServer) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetForceHTTP returns the ForceHTTP field value if set, zero value otherwise.
func (o *OAuth2AuthServer) GetForceHTTP() bool {
	if o == nil || o.ForceHTTP == nil {
		var ret bool
		return ret
	}
	return *o.ForceHTTP
}

// GetForceHTTPOk returns a tuple with the ForceHTTP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AuthServer) GetForceHTTPOk() (*bool, bool) {
	if o == nil || o.ForceHTTP == nil {
		return nil, false
	}
	return o.ForceHTTP, true
}

// HasForceHTTP returns a boolean if a field has been set.
func (o *OAuth2AuthServer) HasForceHTTP() bool {
	if o != nil && o.ForceHTTP != nil {
		return true
	}

	return false
}

// SetForceHTTP gets a reference to the given bool and assigns it to the ForceHTTP field.
func (o *OAuth2AuthServer) SetForceHTTP(v bool) {
	o.ForceHTTP = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OAuth2AuthServer) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AuthServer) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OAuth2AuthServer) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OAuth2AuthServer) SetId(v string) {
	o.Id = &v
}

// GetIdTokenLifespan returns the IdTokenLifespan field value if set, zero value otherwise.
func (o *OAuth2AuthServer) GetIdTokenLifespan() int64 {
	if o == nil || o.IdTokenLifespan == nil {
		var ret int64
		return ret
	}
	return *o.IdTokenLifespan
}

// GetIdTokenLifespanOk returns a tuple with the IdTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AuthServer) GetIdTokenLifespanOk() (*int64, bool) {
	if o == nil || o.IdTokenLifespan == nil {
		return nil, false
	}
	return o.IdTokenLifespan, true
}

// HasIdTokenLifespan returns a boolean if a field has been set.
func (o *OAuth2AuthServer) HasIdTokenLifespan() bool {
	if o != nil && o.IdTokenLifespan != nil {
		return true
	}

	return false
}

// SetIdTokenLifespan gets a reference to the given int64 and assigns it to the IdTokenLifespan field.
func (o *OAuth2AuthServer) SetIdTokenLifespan(v int64) {
	o.IdTokenLifespan = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *OAuth2AuthServer) GetMeta() ApiTokenMeta {
	if o == nil || o.Meta == nil {
		var ret ApiTokenMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AuthServer) GetMetaOk() (*ApiTokenMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *OAuth2AuthServer) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiTokenMeta and assigns it to the Meta field.
func (o *OAuth2AuthServer) SetMeta(v ApiTokenMeta) {
	o.Meta = &v
}

// GetRefreshTokenLifespan returns the RefreshTokenLifespan field value if set, zero value otherwise.
func (o *OAuth2AuthServer) GetRefreshTokenLifespan() int64 {
	if o == nil || o.RefreshTokenLifespan == nil {
		var ret int64
		return ret
	}
	return *o.RefreshTokenLifespan
}

// GetRefreshTokenLifespanOk returns a tuple with the RefreshTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AuthServer) GetRefreshTokenLifespanOk() (*int64, bool) {
	if o == nil || o.RefreshTokenLifespan == nil {
		return nil, false
	}
	return o.RefreshTokenLifespan, true
}

// HasRefreshTokenLifespan returns a boolean if a field has been set.
func (o *OAuth2AuthServer) HasRefreshTokenLifespan() bool {
	if o != nil && o.RefreshTokenLifespan != nil {
		return true
	}

	return false
}

// SetRefreshTokenLifespan gets a reference to the given int64 and assigns it to the RefreshTokenLifespan field.
func (o *OAuth2AuthServer) SetRefreshTokenLifespan(v int64) {
	o.RefreshTokenLifespan = &v
}

// GetScopeStrategy returns the ScopeStrategy field value if set, zero value otherwise.
func (o *OAuth2AuthServer) GetScopeStrategy() string {
	if o == nil || o.ScopeStrategy == nil {
		var ret string
		return ret
	}
	return *o.ScopeStrategy
}

// GetScopeStrategyOk returns a tuple with the ScopeStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AuthServer) GetScopeStrategyOk() (*string, bool) {
	if o == nil || o.ScopeStrategy == nil {
		return nil, false
	}
	return o.ScopeStrategy, true
}

// HasScopeStrategy returns a boolean if a field has been set.
func (o *OAuth2AuthServer) HasScopeStrategy() bool {
	if o != nil && o.ScopeStrategy != nil {
		return true
	}

	return false
}

// SetScopeStrategy gets a reference to the given string and assigns it to the ScopeStrategy field.
func (o *OAuth2AuthServer) SetScopeStrategy(v string) {
	o.ScopeStrategy = &v
}

// GetSubjectIdentifierAlgorithmSalt returns the SubjectIdentifierAlgorithmSalt field value if set, zero value otherwise.
func (o *OAuth2AuthServer) GetSubjectIdentifierAlgorithmSalt() string {
	if o == nil || o.SubjectIdentifierAlgorithmSalt == nil {
		var ret string
		return ret
	}
	return *o.SubjectIdentifierAlgorithmSalt
}

// GetSubjectIdentifierAlgorithmSaltOk returns a tuple with the SubjectIdentifierAlgorithmSalt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AuthServer) GetSubjectIdentifierAlgorithmSaltOk() (*string, bool) {
	if o == nil || o.SubjectIdentifierAlgorithmSalt == nil {
		return nil, false
	}
	return o.SubjectIdentifierAlgorithmSalt, true
}

// HasSubjectIdentifierAlgorithmSalt returns a boolean if a field has been set.
func (o *OAuth2AuthServer) HasSubjectIdentifierAlgorithmSalt() bool {
	if o != nil && o.SubjectIdentifierAlgorithmSalt != nil {
		return true
	}

	return false
}

// SetSubjectIdentifierAlgorithmSalt gets a reference to the given string and assigns it to the SubjectIdentifierAlgorithmSalt field.
func (o *OAuth2AuthServer) SetSubjectIdentifierAlgorithmSalt(v string) {
	o.SubjectIdentifierAlgorithmSalt = &v
}

// GetSubjectTypesSupported returns the SubjectTypesSupported field value if set, zero value otherwise.
func (o *OAuth2AuthServer) GetSubjectTypesSupported() []string {
	if o == nil || o.SubjectTypesSupported == nil {
		var ret []string
		return ret
	}
	return *o.SubjectTypesSupported
}

// GetSubjectTypesSupportedOk returns a tuple with the SubjectTypesSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AuthServer) GetSubjectTypesSupportedOk() (*[]string, bool) {
	if o == nil || o.SubjectTypesSupported == nil {
		return nil, false
	}
	return o.SubjectTypesSupported, true
}

// HasSubjectTypesSupported returns a boolean if a field has been set.
func (o *OAuth2AuthServer) HasSubjectTypesSupported() bool {
	if o != nil && o.SubjectTypesSupported != nil {
		return true
	}

	return false
}

// SetSubjectTypesSupported gets a reference to the given []string and assigns it to the SubjectTypesSupported field.
func (o *OAuth2AuthServer) SetSubjectTypesSupported(v []string) {
	o.SubjectTypesSupported = &v
}

// GetWellKnownKeys returns the WellKnownKeys field value if set, zero value otherwise.
func (o *OAuth2AuthServer) GetWellKnownKeys() []string {
	if o == nil || o.WellKnownKeys == nil {
		var ret []string
		return ret
	}
	return *o.WellKnownKeys
}

// GetWellKnownKeysOk returns a tuple with the WellKnownKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AuthServer) GetWellKnownKeysOk() (*[]string, bool) {
	if o == nil || o.WellKnownKeys == nil {
		return nil, false
	}
	return o.WellKnownKeys, true
}

// HasWellKnownKeys returns a boolean if a field has been set.
func (o *OAuth2AuthServer) HasWellKnownKeys() bool {
	if o != nil && o.WellKnownKeys != nil {
		return true
	}

	return false
}

// SetWellKnownKeys gets a reference to the given []string and assigns it to the WellKnownKeys field.
func (o *OAuth2AuthServer) SetWellKnownKeys(v []string) {
	o.WellKnownKeys = &v
}

func (o OAuth2AuthServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessTokenLifespan != nil {
		toSerialize["accessTokenLifespan"] = o.AccessTokenLifespan
	}
	if o.AccessTokenStrategy != nil {
		toSerialize["accessTokenStrategy"] = o.AccessTokenStrategy
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.ConsentRequestMaxAge != nil {
		toSerialize["consentRequestMaxAge"] = o.ConsentRequestMaxAge
	}
	if o.CookieSameSiteLegacyWorkaround != nil {
		toSerialize["cookieSameSiteLegacyWorkaround"] = o.CookieSameSiteLegacyWorkaround
	}
	if o.CookieSameSiteMode != nil {
		toSerialize["cookieSameSiteMode"] = o.CookieSameSiteMode
	}
	if o.Debug != nil {
		toSerialize["debug"] = o.Debug
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.ForceHTTP != nil {
		toSerialize["forceHTTP"] = o.ForceHTTP
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IdTokenLifespan != nil {
		toSerialize["idTokenLifespan"] = o.IdTokenLifespan
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.RefreshTokenLifespan != nil {
		toSerialize["refreshTokenLifespan"] = o.RefreshTokenLifespan
	}
	if o.ScopeStrategy != nil {
		toSerialize["scopeStrategy"] = o.ScopeStrategy
	}
	if o.SubjectIdentifierAlgorithmSalt != nil {
		toSerialize["subjectIdentifierAlgorithmSalt"] = o.SubjectIdentifierAlgorithmSalt
	}
	if o.SubjectTypesSupported != nil {
		toSerialize["subjectTypesSupported"] = o.SubjectTypesSupported
	}
	if o.WellKnownKeys != nil {
		toSerialize["wellKnownKeys"] = o.WellKnownKeys
	}
	return json.Marshal(toSerialize)
}

type NullableOAuth2AuthServer struct {
	value *OAuth2AuthServer
	isSet bool
}

func (v NullableOAuth2AuthServer) Get() *OAuth2AuthServer {
	return v.value
}

func (v *NullableOAuth2AuthServer) Set(val *OAuth2AuthServer) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2AuthServer) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2AuthServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2AuthServer(val *OAuth2AuthServer) *NullableOAuth2AuthServer {
	return &NullableOAuth2AuthServer{value: val, isSet: true}
}

func (v NullableOAuth2AuthServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2AuthServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


