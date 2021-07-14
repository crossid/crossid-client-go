# App

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**AppLogic** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **interface{}** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Keywords** | Pointer to **[]string** |  | [optional] 
**LogoURL** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**AppMeta**](AppMeta.md) |  | [optional] 

## Methods

### NewApp

`func NewApp() *App`

NewApp instantiates a new App object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppWithDefaults

`func NewAppWithDefaults() *App`

NewAppWithDefaults instantiates a new App object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *App) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *App) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *App) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *App) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAppId

`func (o *App) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *App) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *App) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *App) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAppLogic

`func (o *App) GetAppLogic() string`

GetAppLogic returns the AppLogic field if non-nil, zero value otherwise.

### GetAppLogicOk

`func (o *App) GetAppLogicOk() (*string, bool)`

GetAppLogicOk returns a tuple with the AppLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLogic

`func (o *App) SetAppLogic(v string)`

SetAppLogic sets AppLogic field to given value.

### HasAppLogic

`func (o *App) HasAppLogic() bool`

HasAppLogic returns a boolean if a field has been set.

### GetConfig

`func (o *App) GetConfig() interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *App) GetConfigOk() (*interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *App) SetConfig(v interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *App) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *App) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *App) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetDisplayName

`func (o *App) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *App) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *App) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *App) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *App) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *App) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *App) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *App) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeywords

`func (o *App) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *App) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *App) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *App) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetLogoURL

`func (o *App) GetLogoURL() string`

GetLogoURL returns the LogoURL field if non-nil, zero value otherwise.

### GetLogoURLOk

`func (o *App) GetLogoURLOk() (*string, bool)`

GetLogoURLOk returns a tuple with the LogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoURL

`func (o *App) SetLogoURL(v string)`

SetLogoURL sets LogoURL field to given value.

### HasLogoURL

`func (o *App) HasLogoURL() bool`

HasLogoURL returns a boolean if a field has been set.

### GetMeta

`func (o *App) GetMeta() AppMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *App) GetMetaOk() (*AppMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *App) SetMeta(v AppMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *App) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


