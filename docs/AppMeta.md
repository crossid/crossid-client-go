# AppMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** |  | [optional] 
**CommitId** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**LastModified** | Pointer to **time.Time** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**RefersToId** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **int64** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 

## Methods

### NewAppMeta

`func NewAppMeta() *AppMeta`

NewAppMeta instantiates a new AppMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppMetaWithDefaults

`func NewAppMetaWithDefaults() *AppMeta`

NewAppMetaWithDefaults instantiates a new AppMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *AppMeta) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AppMeta) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AppMeta) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *AppMeta) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetCommitId

`func (o *AppMeta) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *AppMeta) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *AppMeta) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.

### HasCommitId

`func (o *AppMeta) HasCommitId() bool`

HasCommitId returns a boolean if a field has been set.

### GetCreated

`func (o *AppMeta) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AppMeta) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AppMeta) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AppMeta) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastModified

`func (o *AppMeta) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AppMeta) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AppMeta) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *AppMeta) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetLocation

`func (o *AppMeta) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AppMeta) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AppMeta) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AppMeta) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetRefersToId

`func (o *AppMeta) GetRefersToId() string`

GetRefersToId returns the RefersToId field if non-nil, zero value otherwise.

### GetRefersToIdOk

`func (o *AppMeta) GetRefersToIdOk() (*string, bool)`

GetRefersToIdOk returns a tuple with the RefersToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefersToId

`func (o *AppMeta) SetRefersToId(v string)`

SetRefersToId sets RefersToId field to given value.

### HasRefersToId

`func (o *AppMeta) HasRefersToId() bool`

HasRefersToId returns a boolean if a field has been set.

### GetResourceType

`func (o *AppMeta) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AppMeta) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AppMeta) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AppMeta) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetRevision

`func (o *AppMeta) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *AppMeta) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *AppMeta) SetRevision(v int64)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *AppMeta) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetTenantId

`func (o *AppMeta) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AppMeta) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AppMeta) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AppMeta) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


