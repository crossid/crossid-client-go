# ApiToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IdentityId** | Pointer to **string** |  | [optional] 
**IssuedAt** | Pointer to **time.Time** |  | [optional] 
**Meta** | Pointer to [**ApiTokenMeta**](ApiTokenMeta.md) |  | [optional] 
**SignedToken** | Pointer to **string** |  | [optional] 

## Methods

### NewApiToken

`func NewApiToken() *ApiToken`

NewApiToken instantiates a new ApiToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTokenWithDefaults

`func NewApiTokenWithDefaults() *ApiToken`

NewApiTokenWithDefaults instantiates a new ApiToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ApiToken) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApiToken) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApiToken) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApiToken) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDisplayName

`func (o *ApiToken) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiToken) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiToken) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApiToken) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ApiToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ApiToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ApiToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ApiToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *ApiToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityId

`func (o *ApiToken) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *ApiToken) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *ApiToken) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *ApiToken) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetIssuedAt

`func (o *ApiToken) GetIssuedAt() time.Time`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *ApiToken) GetIssuedAtOk() (*time.Time, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *ApiToken) SetIssuedAt(v time.Time)`

SetIssuedAt sets IssuedAt field to given value.

### HasIssuedAt

`func (o *ApiToken) HasIssuedAt() bool`

HasIssuedAt returns a boolean if a field has been set.

### GetMeta

`func (o *ApiToken) GetMeta() ApiTokenMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ApiToken) GetMetaOk() (*ApiTokenMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ApiToken) SetMeta(v ApiTokenMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ApiToken) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSignedToken

`func (o *ApiToken) GetSignedToken() string`

GetSignedToken returns the SignedToken field if non-nil, zero value otherwise.

### GetSignedTokenOk

`func (o *ApiToken) GetSignedTokenOk() (*string, bool)`

GetSignedTokenOk returns a tuple with the SignedToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedToken

`func (o *ApiToken) SetSignedToken(v string)`

SetSignedToken sets SignedToken field to given value.

### HasSignedToken

`func (o *ApiToken) HasSignedToken() bool`

HasSignedToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


