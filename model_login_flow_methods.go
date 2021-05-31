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

// LoginFlowMethods struct for LoginFlowMethods
type LoginFlowMethods struct {
	Otp *LoginFlowMethod `json:"otp,omitempty"`
	Password *LoginFlowMethod `json:"password,omitempty"`
	Spnego *LoginFlowMethod `json:"spnego,omitempty"`
	Webauthn *LoginFlowMethod `json:"webauthn,omitempty"`
}

// NewLoginFlowMethods instantiates a new LoginFlowMethods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginFlowMethods() *LoginFlowMethods {
	this := LoginFlowMethods{}
	return &this
}

// NewLoginFlowMethodsWithDefaults instantiates a new LoginFlowMethods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginFlowMethodsWithDefaults() *LoginFlowMethods {
	this := LoginFlowMethods{}
	return &this
}

// GetOtp returns the Otp field value if set, zero value otherwise.
func (o *LoginFlowMethods) GetOtp() LoginFlowMethod {
	if o == nil || o.Otp == nil {
		var ret LoginFlowMethod
		return ret
	}
	return *o.Otp
}

// GetOtpOk returns a tuple with the Otp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlowMethods) GetOtpOk() (*LoginFlowMethod, bool) {
	if o == nil || o.Otp == nil {
		return nil, false
	}
	return o.Otp, true
}

// HasOtp returns a boolean if a field has been set.
func (o *LoginFlowMethods) HasOtp() bool {
	if o != nil && o.Otp != nil {
		return true
	}

	return false
}

// SetOtp gets a reference to the given LoginFlowMethod and assigns it to the Otp field.
func (o *LoginFlowMethods) SetOtp(v LoginFlowMethod) {
	o.Otp = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *LoginFlowMethods) GetPassword() LoginFlowMethod {
	if o == nil || o.Password == nil {
		var ret LoginFlowMethod
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlowMethods) GetPasswordOk() (*LoginFlowMethod, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *LoginFlowMethods) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given LoginFlowMethod and assigns it to the Password field.
func (o *LoginFlowMethods) SetPassword(v LoginFlowMethod) {
	o.Password = &v
}

// GetSpnego returns the Spnego field value if set, zero value otherwise.
func (o *LoginFlowMethods) GetSpnego() LoginFlowMethod {
	if o == nil || o.Spnego == nil {
		var ret LoginFlowMethod
		return ret
	}
	return *o.Spnego
}

// GetSpnegoOk returns a tuple with the Spnego field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlowMethods) GetSpnegoOk() (*LoginFlowMethod, bool) {
	if o == nil || o.Spnego == nil {
		return nil, false
	}
	return o.Spnego, true
}

// HasSpnego returns a boolean if a field has been set.
func (o *LoginFlowMethods) HasSpnego() bool {
	if o != nil && o.Spnego != nil {
		return true
	}

	return false
}

// SetSpnego gets a reference to the given LoginFlowMethod and assigns it to the Spnego field.
func (o *LoginFlowMethods) SetSpnego(v LoginFlowMethod) {
	o.Spnego = &v
}

// GetWebauthn returns the Webauthn field value if set, zero value otherwise.
func (o *LoginFlowMethods) GetWebauthn() LoginFlowMethod {
	if o == nil || o.Webauthn == nil {
		var ret LoginFlowMethod
		return ret
	}
	return *o.Webauthn
}

// GetWebauthnOk returns a tuple with the Webauthn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlowMethods) GetWebauthnOk() (*LoginFlowMethod, bool) {
	if o == nil || o.Webauthn == nil {
		return nil, false
	}
	return o.Webauthn, true
}

// HasWebauthn returns a boolean if a field has been set.
func (o *LoginFlowMethods) HasWebauthn() bool {
	if o != nil && o.Webauthn != nil {
		return true
	}

	return false
}

// SetWebauthn gets a reference to the given LoginFlowMethod and assigns it to the Webauthn field.
func (o *LoginFlowMethods) SetWebauthn(v LoginFlowMethod) {
	o.Webauthn = &v
}

func (o LoginFlowMethods) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Otp != nil {
		toSerialize["otp"] = o.Otp
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Spnego != nil {
		toSerialize["spnego"] = o.Spnego
	}
	if o.Webauthn != nil {
		toSerialize["webauthn"] = o.Webauthn
	}
	return json.Marshal(toSerialize)
}

type NullableLoginFlowMethods struct {
	value *LoginFlowMethods
	isSet bool
}

func (v NullableLoginFlowMethods) Get() *LoginFlowMethods {
	return v.value
}

func (v *NullableLoginFlowMethods) Set(val *LoginFlowMethods) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginFlowMethods) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginFlowMethods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginFlowMethods(val *LoginFlowMethods) *NullableLoginFlowMethods {
	return &NullableLoginFlowMethods{value: val, isSet: true}
}

func (v NullableLoginFlowMethods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginFlowMethods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


