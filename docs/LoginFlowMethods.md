# LoginFlowMethods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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


