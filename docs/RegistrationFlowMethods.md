# RegistrationFlowMethods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to [**RegistrationFlowMethod**](RegistrationFlowMethod.md) |  | [optional] 
**Spnego** | Pointer to [**RegistrationFlowMethod**](RegistrationFlowMethod.md) |  | [optional] 
**Webauthn** | Pointer to [**RegistrationFlowMethod**](RegistrationFlowMethod.md) |  | [optional] 

## Methods

### NewRegistrationFlowMethods

`func NewRegistrationFlowMethods() *RegistrationFlowMethods`

NewRegistrationFlowMethods instantiates a new RegistrationFlowMethods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationFlowMethodsWithDefaults

`func NewRegistrationFlowMethodsWithDefaults() *RegistrationFlowMethods`

NewRegistrationFlowMethodsWithDefaults instantiates a new RegistrationFlowMethods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *RegistrationFlowMethods) GetPassword() RegistrationFlowMethod`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegistrationFlowMethods) GetPasswordOk() (*RegistrationFlowMethod, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegistrationFlowMethods) SetPassword(v RegistrationFlowMethod)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RegistrationFlowMethods) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSpnego

`func (o *RegistrationFlowMethods) GetSpnego() RegistrationFlowMethod`

GetSpnego returns the Spnego field if non-nil, zero value otherwise.

### GetSpnegoOk

`func (o *RegistrationFlowMethods) GetSpnegoOk() (*RegistrationFlowMethod, bool)`

GetSpnegoOk returns a tuple with the Spnego field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpnego

`func (o *RegistrationFlowMethods) SetSpnego(v RegistrationFlowMethod)`

SetSpnego sets Spnego field to given value.

### HasSpnego

`func (o *RegistrationFlowMethods) HasSpnego() bool`

HasSpnego returns a boolean if a field has been set.

### GetWebauthn

`func (o *RegistrationFlowMethods) GetWebauthn() RegistrationFlowMethod`

GetWebauthn returns the Webauthn field if non-nil, zero value otherwise.

### GetWebauthnOk

`func (o *RegistrationFlowMethods) GetWebauthnOk() (*RegistrationFlowMethod, bool)`

GetWebauthnOk returns a tuple with the Webauthn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebauthn

`func (o *RegistrationFlowMethods) SetWebauthn(v RegistrationFlowMethod)`

SetWebauthn sets Webauthn field to given value.

### HasWebauthn

`func (o *RegistrationFlowMethods) HasWebauthn() bool`

HasWebauthn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

