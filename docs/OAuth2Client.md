# OAuth2Client

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** |  | [optional] 
**ClientName** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**GrantTypes** | Pointer to **string** |  | [optional] 
**RedirectUris** | Pointer to **[]string** |  | [optional] 
**ResponseTypes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOAuth2Client

`func NewOAuth2Client() *OAuth2Client`

NewOAuth2Client instantiates a new OAuth2Client object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ClientWithDefaults

`func NewOAuth2ClientWithDefaults() *OAuth2Client`

NewOAuth2ClientWithDefaults instantiates a new OAuth2Client object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OAuth2Client) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuth2Client) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuth2Client) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OAuth2Client) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientName

`func (o *OAuth2Client) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *OAuth2Client) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *OAuth2Client) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *OAuth2Client) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetClientSecret

`func (o *OAuth2Client) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuth2Client) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuth2Client) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuth2Client) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetGrantTypes

`func (o *OAuth2Client) GetGrantTypes() string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *OAuth2Client) GetGrantTypesOk() (*string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *OAuth2Client) SetGrantTypes(v string)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *OAuth2Client) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetRedirectUris

`func (o *OAuth2Client) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *OAuth2Client) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *OAuth2Client) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *OAuth2Client) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetResponseTypes

`func (o *OAuth2Client) GetResponseTypes() []string`

GetResponseTypes returns the ResponseTypes field if non-nil, zero value otherwise.

### GetResponseTypesOk

`func (o *OAuth2Client) GetResponseTypesOk() (*[]string, bool)`

GetResponseTypesOk returns a tuple with the ResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypes

`func (o *OAuth2Client) SetResponseTypes(v []string)`

SetResponseTypes sets ResponseTypes field to given value.

### HasResponseTypes

`func (o *OAuth2Client) HasResponseTypes() bool`

HasResponseTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


