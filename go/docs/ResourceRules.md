# ResourceRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Patch** | Pointer to [**[]PatchOP**](PatchOP.md) |  | [optional] 
**Filter** | Pointer to **string** | which rules to check against | [optional] 

## Methods

### NewResourceRules

`func NewResourceRules() *ResourceRules`

NewResourceRules instantiates a new ResourceRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRulesWithDefaults

`func NewResourceRulesWithDefaults() *ResourceRules`

NewResourceRulesWithDefaults instantiates a new ResourceRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatch

`func (o *ResourceRules) GetPatch() []PatchOP`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *ResourceRules) GetPatchOk() (*[]PatchOP, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *ResourceRules) SetPatch(v []PatchOP)`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *ResourceRules) HasPatch() bool`

HasPatch returns a boolean if a field has been set.

### GetFilter

`func (o *ResourceRules) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ResourceRules) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ResourceRules) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ResourceRules) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


