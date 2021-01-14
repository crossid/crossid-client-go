# TasksList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemsPerPage** | Pointer to **int32** |  | [optional] 
**StartIndex** | Pointer to **int64** |  | [optional] 
**TotalResults** | Pointer to **int64** |  | [optional] 
**Resources** | [**[]Task**](Task.md) |  | 

## Methods

### NewTasksList

`func NewTasksList(resources []Task, ) *TasksList`

NewTasksList instantiates a new TasksList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTasksListWithDefaults

`func NewTasksListWithDefaults() *TasksList`

NewTasksListWithDefaults instantiates a new TasksList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemsPerPage

`func (o *TasksList) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *TasksList) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *TasksList) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *TasksList) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetStartIndex

`func (o *TasksList) GetStartIndex() int64`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *TasksList) GetStartIndexOk() (*int64, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *TasksList) SetStartIndex(v int64)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *TasksList) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### GetTotalResults

`func (o *TasksList) GetTotalResults() int64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *TasksList) GetTotalResultsOk() (*int64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *TasksList) SetTotalResults(v int64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *TasksList) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetResources

`func (o *TasksList) GetResources() []Task`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *TasksList) GetResourcesOk() (*[]Task, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *TasksList) SetResources(v []Task)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


