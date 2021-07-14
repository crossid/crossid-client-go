# OAuth2HandledConsentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** |  | [optional] 
**GrantAccessTokenAudience** | Pointer to **[]string** |  | [optional] 
**GrantScope** | Pointer to **[]string** |  | [optional] 
**HandledAt** | Pointer to **NullableTime** |  | [optional] 
**Remember** | Pointer to **bool** |  | [optional] 
**RememberFor** | Pointer to **int32** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 

## Methods

### NewOAuth2HandledConsentRequest

`func NewOAuth2HandledConsentRequest() *OAuth2HandledConsentRequest`

NewOAuth2HandledConsentRequest instantiates a new OAuth2HandledConsentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2HandledConsentRequestWithDefaults

`func NewOAuth2HandledConsentRequestWithDefaults() *OAuth2HandledConsentRequest`

NewOAuth2HandledConsentRequestWithDefaults instantiates a new OAuth2HandledConsentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OAuth2HandledConsentRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuth2HandledConsentRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuth2HandledConsentRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OAuth2HandledConsentRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetGrantAccessTokenAudience

`func (o *OAuth2HandledConsentRequest) GetGrantAccessTokenAudience() []string`

GetGrantAccessTokenAudience returns the GrantAccessTokenAudience field if non-nil, zero value otherwise.

### GetGrantAccessTokenAudienceOk

`func (o *OAuth2HandledConsentRequest) GetGrantAccessTokenAudienceOk() (*[]string, bool)`

GetGrantAccessTokenAudienceOk returns a tuple with the GrantAccessTokenAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantAccessTokenAudience

`func (o *OAuth2HandledConsentRequest) SetGrantAccessTokenAudience(v []string)`

SetGrantAccessTokenAudience sets GrantAccessTokenAudience field to given value.

### HasGrantAccessTokenAudience

`func (o *OAuth2HandledConsentRequest) HasGrantAccessTokenAudience() bool`

HasGrantAccessTokenAudience returns a boolean if a field has been set.

### GetGrantScope

`func (o *OAuth2HandledConsentRequest) GetGrantScope() []string`

GetGrantScope returns the GrantScope field if non-nil, zero value otherwise.

### GetGrantScopeOk

`func (o *OAuth2HandledConsentRequest) GetGrantScopeOk() (*[]string, bool)`

GetGrantScopeOk returns a tuple with the GrantScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantScope

`func (o *OAuth2HandledConsentRequest) SetGrantScope(v []string)`

SetGrantScope sets GrantScope field to given value.

### HasGrantScope

`func (o *OAuth2HandledConsentRequest) HasGrantScope() bool`

HasGrantScope returns a boolean if a field has been set.

### GetHandledAt

`func (o *OAuth2HandledConsentRequest) GetHandledAt() time.Time`

GetHandledAt returns the HandledAt field if non-nil, zero value otherwise.

### GetHandledAtOk

`func (o *OAuth2HandledConsentRequest) GetHandledAtOk() (*time.Time, bool)`

GetHandledAtOk returns a tuple with the HandledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandledAt

`func (o *OAuth2HandledConsentRequest) SetHandledAt(v time.Time)`

SetHandledAt sets HandledAt field to given value.

### HasHandledAt

`func (o *OAuth2HandledConsentRequest) HasHandledAt() bool`

HasHandledAt returns a boolean if a field has been set.

### SetHandledAtNil

`func (o *OAuth2HandledConsentRequest) SetHandledAtNil(b bool)`

 SetHandledAtNil sets the value for HandledAt to be an explicit nil

### UnsetHandledAt
`func (o *OAuth2HandledConsentRequest) UnsetHandledAt()`

UnsetHandledAt ensures that no value is present for HandledAt, not even an explicit nil
### GetRemember

`func (o *OAuth2HandledConsentRequest) GetRemember() bool`

GetRemember returns the Remember field if non-nil, zero value otherwise.

### GetRememberOk

`func (o *OAuth2HandledConsentRequest) GetRememberOk() (*bool, bool)`

GetRememberOk returns a tuple with the Remember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemember

`func (o *OAuth2HandledConsentRequest) SetRemember(v bool)`

SetRemember sets Remember field to given value.

### HasRemember

`func (o *OAuth2HandledConsentRequest) HasRemember() bool`

HasRemember returns a boolean if a field has been set.

### GetRememberFor

`func (o *OAuth2HandledConsentRequest) GetRememberFor() int32`

GetRememberFor returns the RememberFor field if non-nil, zero value otherwise.

### GetRememberForOk

`func (o *OAuth2HandledConsentRequest) GetRememberForOk() (*int32, bool)`

GetRememberForOk returns a tuple with the RememberFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberFor

`func (o *OAuth2HandledConsentRequest) SetRememberFor(v int32)`

SetRememberFor sets RememberFor field to given value.

### HasRememberFor

`func (o *OAuth2HandledConsentRequest) HasRememberFor() bool`

HasRememberFor returns a boolean if a field has been set.

### GetSubject

`func (o *OAuth2HandledConsentRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *OAuth2HandledConsentRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *OAuth2HandledConsentRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *OAuth2HandledConsentRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


