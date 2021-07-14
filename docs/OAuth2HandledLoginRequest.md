# OAuth2HandledLoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acr** | Pointer to **string** |  | [optional] 
**Context** | Pointer to **map[string]string** |  | [optional] 
**ForceSubjectIdentifier** | Pointer to **string** |  | [optional] 
**Remember** | Pointer to **bool** |  | [optional] 
**RememberFor** | Pointer to **int32** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 

## Methods

### NewOAuth2HandledLoginRequest

`func NewOAuth2HandledLoginRequest() *OAuth2HandledLoginRequest`

NewOAuth2HandledLoginRequest instantiates a new OAuth2HandledLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2HandledLoginRequestWithDefaults

`func NewOAuth2HandledLoginRequestWithDefaults() *OAuth2HandledLoginRequest`

NewOAuth2HandledLoginRequestWithDefaults instantiates a new OAuth2HandledLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcr

`func (o *OAuth2HandledLoginRequest) GetAcr() string`

GetAcr returns the Acr field if non-nil, zero value otherwise.

### GetAcrOk

`func (o *OAuth2HandledLoginRequest) GetAcrOk() (*string, bool)`

GetAcrOk returns a tuple with the Acr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcr

`func (o *OAuth2HandledLoginRequest) SetAcr(v string)`

SetAcr sets Acr field to given value.

### HasAcr

`func (o *OAuth2HandledLoginRequest) HasAcr() bool`

HasAcr returns a boolean if a field has been set.

### GetContext

`func (o *OAuth2HandledLoginRequest) GetContext() map[string]string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *OAuth2HandledLoginRequest) GetContextOk() (*map[string]string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *OAuth2HandledLoginRequest) SetContext(v map[string]string)`

SetContext sets Context field to given value.

### HasContext

`func (o *OAuth2HandledLoginRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetForceSubjectIdentifier

`func (o *OAuth2HandledLoginRequest) GetForceSubjectIdentifier() string`

GetForceSubjectIdentifier returns the ForceSubjectIdentifier field if non-nil, zero value otherwise.

### GetForceSubjectIdentifierOk

`func (o *OAuth2HandledLoginRequest) GetForceSubjectIdentifierOk() (*string, bool)`

GetForceSubjectIdentifierOk returns a tuple with the ForceSubjectIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubjectIdentifier

`func (o *OAuth2HandledLoginRequest) SetForceSubjectIdentifier(v string)`

SetForceSubjectIdentifier sets ForceSubjectIdentifier field to given value.

### HasForceSubjectIdentifier

`func (o *OAuth2HandledLoginRequest) HasForceSubjectIdentifier() bool`

HasForceSubjectIdentifier returns a boolean if a field has been set.

### GetRemember

`func (o *OAuth2HandledLoginRequest) GetRemember() bool`

GetRemember returns the Remember field if non-nil, zero value otherwise.

### GetRememberOk

`func (o *OAuth2HandledLoginRequest) GetRememberOk() (*bool, bool)`

GetRememberOk returns a tuple with the Remember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemember

`func (o *OAuth2HandledLoginRequest) SetRemember(v bool)`

SetRemember sets Remember field to given value.

### HasRemember

`func (o *OAuth2HandledLoginRequest) HasRemember() bool`

HasRemember returns a boolean if a field has been set.

### GetRememberFor

`func (o *OAuth2HandledLoginRequest) GetRememberFor() int32`

GetRememberFor returns the RememberFor field if non-nil, zero value otherwise.

### GetRememberForOk

`func (o *OAuth2HandledLoginRequest) GetRememberForOk() (*int32, bool)`

GetRememberForOk returns a tuple with the RememberFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberFor

`func (o *OAuth2HandledLoginRequest) SetRememberFor(v int32)`

SetRememberFor sets RememberFor field to given value.

### HasRememberFor

`func (o *OAuth2HandledLoginRequest) HasRememberFor() bool`

HasRememberFor returns a boolean if a field has been set.

### GetSubject

`func (o *OAuth2HandledLoginRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *OAuth2HandledLoginRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *OAuth2HandledLoginRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *OAuth2HandledLoginRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


