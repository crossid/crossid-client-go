# OAuth2LoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Challenge** | Pointer to **string** |  | [optional] 
**Client** | Pointer to [**OAuth2Client**](OAuth2Client.md) |  | [optional] 
**Csrf** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**RequestUrl** | Pointer to **string** |  | [optional] 
**RequestedAccessTokenAudience** | Pointer to **[]string** |  | [optional] 
**RequestedScope** | Pointer to **[]string** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**Skip** | Pointer to **bool** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 

## Methods

### NewOAuth2LoginRequest

`func NewOAuth2LoginRequest() *OAuth2LoginRequest`

NewOAuth2LoginRequest instantiates a new OAuth2LoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2LoginRequestWithDefaults

`func NewOAuth2LoginRequestWithDefaults() *OAuth2LoginRequest`

NewOAuth2LoginRequestWithDefaults instantiates a new OAuth2LoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChallenge

`func (o *OAuth2LoginRequest) GetChallenge() string`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *OAuth2LoginRequest) GetChallengeOk() (*string, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *OAuth2LoginRequest) SetChallenge(v string)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *OAuth2LoginRequest) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetClient

`func (o *OAuth2LoginRequest) GetClient() OAuth2Client`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *OAuth2LoginRequest) GetClientOk() (*OAuth2Client, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *OAuth2LoginRequest) SetClient(v OAuth2Client)`

SetClient sets Client field to given value.

### HasClient

`func (o *OAuth2LoginRequest) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetCsrf

`func (o *OAuth2LoginRequest) GetCsrf() string`

GetCsrf returns the Csrf field if non-nil, zero value otherwise.

### GetCsrfOk

`func (o *OAuth2LoginRequest) GetCsrfOk() (*string, bool)`

GetCsrfOk returns a tuple with the Csrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrf

`func (o *OAuth2LoginRequest) SetCsrf(v string)`

SetCsrf sets Csrf field to given value.

### HasCsrf

`func (o *OAuth2LoginRequest) HasCsrf() bool`

HasCsrf returns a boolean if a field has been set.

### GetId

`func (o *OAuth2LoginRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuth2LoginRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuth2LoginRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OAuth2LoginRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRequestUrl

`func (o *OAuth2LoginRequest) GetRequestUrl() string`

GetRequestUrl returns the RequestUrl field if non-nil, zero value otherwise.

### GetRequestUrlOk

`func (o *OAuth2LoginRequest) GetRequestUrlOk() (*string, bool)`

GetRequestUrlOk returns a tuple with the RequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUrl

`func (o *OAuth2LoginRequest) SetRequestUrl(v string)`

SetRequestUrl sets RequestUrl field to given value.

### HasRequestUrl

`func (o *OAuth2LoginRequest) HasRequestUrl() bool`

HasRequestUrl returns a boolean if a field has been set.

### GetRequestedAccessTokenAudience

`func (o *OAuth2LoginRequest) GetRequestedAccessTokenAudience() []string`

GetRequestedAccessTokenAudience returns the RequestedAccessTokenAudience field if non-nil, zero value otherwise.

### GetRequestedAccessTokenAudienceOk

`func (o *OAuth2LoginRequest) GetRequestedAccessTokenAudienceOk() (*[]string, bool)`

GetRequestedAccessTokenAudienceOk returns a tuple with the RequestedAccessTokenAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAccessTokenAudience

`func (o *OAuth2LoginRequest) SetRequestedAccessTokenAudience(v []string)`

SetRequestedAccessTokenAudience sets RequestedAccessTokenAudience field to given value.

### HasRequestedAccessTokenAudience

`func (o *OAuth2LoginRequest) HasRequestedAccessTokenAudience() bool`

HasRequestedAccessTokenAudience returns a boolean if a field has been set.

### GetRequestedScope

`func (o *OAuth2LoginRequest) GetRequestedScope() []string`

GetRequestedScope returns the RequestedScope field if non-nil, zero value otherwise.

### GetRequestedScopeOk

`func (o *OAuth2LoginRequest) GetRequestedScopeOk() (*[]string, bool)`

GetRequestedScopeOk returns a tuple with the RequestedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedScope

`func (o *OAuth2LoginRequest) SetRequestedScope(v []string)`

SetRequestedScope sets RequestedScope field to given value.

### HasRequestedScope

`func (o *OAuth2LoginRequest) HasRequestedScope() bool`

HasRequestedScope returns a boolean if a field has been set.

### GetSessionId

`func (o *OAuth2LoginRequest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *OAuth2LoginRequest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *OAuth2LoginRequest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *OAuth2LoginRequest) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSkip

`func (o *OAuth2LoginRequest) GetSkip() bool`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *OAuth2LoginRequest) GetSkipOk() (*bool, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *OAuth2LoginRequest) SetSkip(v bool)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *OAuth2LoginRequest) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetSubject

`func (o *OAuth2LoginRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *OAuth2LoginRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *OAuth2LoginRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *OAuth2LoginRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


