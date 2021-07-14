# PatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | Pointer to [**[]PatchOP**](PatchOP.md) |  | [optional] 
**Revision** | Pointer to **int64** |  | [optional] 

## Methods

### NewPatchRequest

`func NewPatchRequest() *PatchRequest`

NewPatchRequest instantiates a new PatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchRequestWithDefaults

`func NewPatchRequestWithDefaults() *PatchRequest`

NewPatchRequestWithDefaults instantiates a new PatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperations

`func (o *PatchRequest) GetOperations() []PatchOP`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *PatchRequest) GetOperationsOk() (*[]PatchOP, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *PatchRequest) SetOperations(v []PatchOP)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *PatchRequest) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetRevision

`func (o *PatchRequest) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *PatchRequest) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *PatchRequest) SetRevision(v int64)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *PatchRequest) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


