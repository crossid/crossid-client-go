# DocPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 
**Patch** | Pointer to [**[]PatchOP**](PatchOP.md) |  | [optional] 

## Methods

### NewDocPatch

`func NewDocPatch() *DocPatch`

NewDocPatch instantiates a new DocPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocPatchWithDefaults

`func NewDocPatchWithDefaults() *DocPatch`

NewDocPatchWithDefaults instantiates a new DocPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *DocPatch) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DocPatch) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DocPatch) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DocPatch) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *DocPatch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocPatch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocPatch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocPatch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMeta

`func (o *DocPatch) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DocPatch) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DocPatch) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DocPatch) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPatch

`func (o *DocPatch) GetPatch() []PatchOP`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *DocPatch) GetPatchOk() (*[]PatchOP, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *DocPatch) SetPatch(v []PatchOP)`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *DocPatch) HasPatch() bool`

HasPatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


