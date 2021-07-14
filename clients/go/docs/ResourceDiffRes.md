# ResourceDiffRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | Pointer to [**[]PatchOP**](PatchOP.md) |  | [optional] 

## Methods

### NewResourceDiffRes

`func NewResourceDiffRes() *ResourceDiffRes`

NewResourceDiffRes instantiates a new ResourceDiffRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceDiffResWithDefaults

`func NewResourceDiffResWithDefaults() *ResourceDiffRes`

NewResourceDiffResWithDefaults instantiates a new ResourceDiffRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperations

`func (o *ResourceDiffRes) GetOperations() []PatchOP`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *ResourceDiffRes) GetOperationsOk() (*[]PatchOP, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *ResourceDiffRes) SetOperations(v []PatchOP)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *ResourceDiffRes) HasOperations() bool`

HasOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


