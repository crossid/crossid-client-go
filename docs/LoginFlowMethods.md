# LoginFlowMethods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oidc** | Pointer to [**LoginFlowMethod**](LoginFlowMethod.md) |  | [optional] 
**Otp** | Pointer to [**LoginFlowMethod**](LoginFlowMethod.md) |  | [optional] 
**Password** | Pointer to [**LoginFlowMethod**](LoginFlowMethod.md) |  | [optional] 
**Spnego** | Pointer to [**LoginFlowMethod**](LoginFlowMethod.md) |  | [optional] 
**Webauthn** | Pointer to [**LoginFlowMethod**](LoginFlowMethod.md) |  | [optional] 

## Methods

### NewLoginFlowMethods

`func NewLoginFlowMethods() *LoginFlowMethods`

NewLoginFlowMethods instantiates a new LoginFlowMethods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginFlowMethodsWithDefaults

`func NewLoginFlowMethodsWithDefaults() *LoginFlowMethods`

NewLoginFlowMethodsWithDefaults instantiates a new LoginFlowMethods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOidc

`func (o *LoginFlowMethods) GetOidc() LoginFlowMethod`

GetOidc returns the Oidc field if non-nil, zero value otherwise.

### GetOidcOk

`func (o *LoginFlowMethods) GetOidcOk() (*LoginFlowMethod, bool)`

GetOidcOk returns a tuple with the Oidc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidc

`func (o *LoginFlowMethods) SetOidc(v LoginFlowMethod)`

SetOidc sets Oidc field to given value.

### HasOidc

`func (o *LoginFlowMethods) HasOidc() bool`

HasOidc returns a boolean if a field has been set.

### GetOtp

`func (o *LoginFlowMethods) GetOtp() LoginFlowMethod`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *LoginFlowMethods) GetOtpOk() (*LoginFlowMethod, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *LoginFlowMethods) SetOtp(v LoginFlowMethod)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *LoginFlowMethods) HasOtp() bool`

HasOtp returns a boolean if a field has been set.

### GetPassword

`func (o *LoginFlowMethods) GetPassword() LoginFlowMethod`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoginFlowMethods) GetPasswordOk() (*LoginFlowMethod, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoginFlowMethods) SetPassword(v LoginFlowMethod)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LoginFlowMethods) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSpnego

`func (o *LoginFlowMethods) GetSpnego() LoginFlowMethod`

GetSpnego returns the Spnego field if non-nil, zero value otherwise.

### GetSpnegoOk

`func (o *LoginFlowMethods) GetSpnegoOk() (*LoginFlowMethod, bool)`

GetSpnegoOk returns a tuple with the Spnego field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpnego

`func (o *LoginFlowMethods) SetSpnego(v LoginFlowMethod)`

SetSpnego sets Spnego field to given value.

### HasSpnego

`func (o *LoginFlowMethods) HasSpnego() bool`

HasSpnego returns a boolean if a field has been set.

### GetWebauthn

`func (o *LoginFlowMethods) GetWebauthn() LoginFlowMethod`

GetWebauthn returns the Webauthn field if non-nil, zero value otherwise.

### GetWebauthnOk

`func (o *LoginFlowMethods) GetWebauthnOk() (*LoginFlowMethod, bool)`

GetWebauthnOk returns a tuple with the Webauthn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebauthn

`func (o *LoginFlowMethods) SetWebauthn(v LoginFlowMethod)`

SetWebauthn sets Webauthn field to given value.

### HasWebauthn

`func (o *LoginFlowMethods) HasWebauthn() bool`

HasWebauthn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


