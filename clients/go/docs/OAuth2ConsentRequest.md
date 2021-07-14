# OAuth2ConsentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acr** | Pointer to **string** |  | [optional] 
**Challenge** | Pointer to **string** |  | [optional] 
**Client** | Pointer to [**OAuth2Client**](OAuth2Client.md) |  | [optional] 
**Csrf** | Pointer to **string** |  | [optional] 
**LoginChallenge** | Pointer to **string** |  | [optional] 
**LoginSessionId** | Pointer to **string** |  | [optional] 
**RequestUrl** | Pointer to **string** |  | [optional] 
**RequestedAccessTokenAudience** | Pointer to **[]string** |  | [optional] 
**RequestedScope** | Pointer to **[]string** |  | [optional] 
**Skip** | Pointer to **bool** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 

## Methods

### NewOAuth2ConsentRequest

`func NewOAuth2ConsentRequest() *OAuth2ConsentRequest`

NewOAuth2ConsentRequest instantiates a new OAuth2ConsentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ConsentRequestWithDefaults

`func NewOAuth2ConsentRequestWithDefaults() *OAuth2ConsentRequest`

NewOAuth2ConsentRequestWithDefaults instantiates a new OAuth2ConsentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcr

`func (o *OAuth2ConsentRequest) GetAcr() string`

GetAcr returns the Acr field if non-nil, zero value otherwise.

### GetAcrOk

`func (o *OAuth2ConsentRequest) GetAcrOk() (*string, bool)`

GetAcrOk returns a tuple with the Acr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcr

`func (o *OAuth2ConsentRequest) SetAcr(v string)`

SetAcr sets Acr field to given value.

### HasAcr

`func (o *OAuth2ConsentRequest) HasAcr() bool`

HasAcr returns a boolean if a field has been set.

### GetChallenge

`func (o *OAuth2ConsentRequest) GetChallenge() string`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *OAuth2ConsentRequest) GetChallengeOk() (*string, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *OAuth2ConsentRequest) SetChallenge(v string)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *OAuth2ConsentRequest) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetClient

`func (o *OAuth2ConsentRequest) GetClient() OAuth2Client`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *OAuth2ConsentRequest) GetClientOk() (*OAuth2Client, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *OAuth2ConsentRequest) SetClient(v OAuth2Client)`

SetClient sets Client field to given value.

### HasClient

`func (o *OAuth2ConsentRequest) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetCsrf

`func (o *OAuth2ConsentRequest) GetCsrf() string`

GetCsrf returns the Csrf field if non-nil, zero value otherwise.

### GetCsrfOk

`func (o *OAuth2ConsentRequest) GetCsrfOk() (*string, bool)`

GetCsrfOk returns a tuple with the Csrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrf

`func (o *OAuth2ConsentRequest) SetCsrf(v string)`

SetCsrf sets Csrf field to given value.

### HasCsrf

`func (o *OAuth2ConsentRequest) HasCsrf() bool`

HasCsrf returns a boolean if a field has been set.

### GetLoginChallenge

`func (o *OAuth2ConsentRequest) GetLoginChallenge() string`

GetLoginChallenge returns the LoginChallenge field if non-nil, zero value otherwise.

### GetLoginChallengeOk

`func (o *OAuth2ConsentRequest) GetLoginChallengeOk() (*string, bool)`

GetLoginChallengeOk returns a tuple with the LoginChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginChallenge

`func (o *OAuth2ConsentRequest) SetLoginChallenge(v string)`

SetLoginChallenge sets LoginChallenge field to given value.

### HasLoginChallenge

`func (o *OAuth2ConsentRequest) HasLoginChallenge() bool`

HasLoginChallenge returns a boolean if a field has been set.

### GetLoginSessionId

`func (o *OAuth2ConsentRequest) GetLoginSessionId() string`

GetLoginSessionId returns the LoginSessionId field if non-nil, zero value otherwise.

### GetLoginSessionIdOk

`func (o *OAuth2ConsentRequest) GetLoginSessionIdOk() (*string, bool)`

GetLoginSessionIdOk returns a tuple with the LoginSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginSessionId

`func (o *OAuth2ConsentRequest) SetLoginSessionId(v string)`

SetLoginSessionId sets LoginSessionId field to given value.

### HasLoginSessionId

`func (o *OAuth2ConsentRequest) HasLoginSessionId() bool`

HasLoginSessionId returns a boolean if a field has been set.

### GetRequestUrl

`func (o *OAuth2ConsentRequest) GetRequestUrl() string`

GetRequestUrl returns the RequestUrl field if non-nil, zero value otherwise.

### GetRequestUrlOk

`func (o *OAuth2ConsentRequest) GetRequestUrlOk() (*string, bool)`

GetRequestUrlOk returns a tuple with the RequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUrl

`func (o *OAuth2ConsentRequest) SetRequestUrl(v string)`

SetRequestUrl sets RequestUrl field to given value.

### HasRequestUrl

`func (o *OAuth2ConsentRequest) HasRequestUrl() bool`

HasRequestUrl returns a boolean if a field has been set.

### GetRequestedAccessTokenAudience

`func (o *OAuth2ConsentRequest) GetRequestedAccessTokenAudience() []string`

GetRequestedAccessTokenAudience returns the RequestedAccessTokenAudience field if non-nil, zero value otherwise.

### GetRequestedAccessTokenAudienceOk

`func (o *OAuth2ConsentRequest) GetRequestedAccessTokenAudienceOk() (*[]string, bool)`

GetRequestedAccessTokenAudienceOk returns a tuple with the RequestedAccessTokenAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAccessTokenAudience

`func (o *OAuth2ConsentRequest) SetRequestedAccessTokenAudience(v []string)`

SetRequestedAccessTokenAudience sets RequestedAccessTokenAudience field to given value.

### HasRequestedAccessTokenAudience

`func (o *OAuth2ConsentRequest) HasRequestedAccessTokenAudience() bool`

HasRequestedAccessTokenAudience returns a boolean if a field has been set.

### GetRequestedScope

`func (o *OAuth2ConsentRequest) GetRequestedScope() []string`

GetRequestedScope returns the RequestedScope field if non-nil, zero value otherwise.

### GetRequestedScopeOk

`func (o *OAuth2ConsentRequest) GetRequestedScopeOk() (*[]string, bool)`

GetRequestedScopeOk returns a tuple with the RequestedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedScope

`func (o *OAuth2ConsentRequest) SetRequestedScope(v []string)`

SetRequestedScope sets RequestedScope field to given value.

### HasRequestedScope

`func (o *OAuth2ConsentRequest) HasRequestedScope() bool`

HasRequestedScope returns a boolean if a field has been set.

### GetSkip

`func (o *OAuth2ConsentRequest) GetSkip() bool`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *OAuth2ConsentRequest) GetSkipOk() (*bool, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *OAuth2ConsentRequest) SetSkip(v bool)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *OAuth2ConsentRequest) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetSubject

`func (o *OAuth2ConsentRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *OAuth2ConsentRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *OAuth2ConsentRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *OAuth2ConsentRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


