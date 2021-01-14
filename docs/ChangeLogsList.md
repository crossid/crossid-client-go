# ChangeLogsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemsPerPage** | Pointer to **int32** |  | [optional] 
**StartIndex** | Pointer to **int64** |  | [optional] 
**TotalResults** | Pointer to **int64** |  | [optional] 
**Resources** | [**[]ChangeLog**](ChangeLog.md) |  | 

## Methods

### NewChangeLogsList

`func NewChangeLogsList(resources []ChangeLog, ) *ChangeLogsList`

NewChangeLogsList instantiates a new ChangeLogsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeLogsListWithDefaults

`func NewChangeLogsListWithDefaults() *ChangeLogsList`

NewChangeLogsListWithDefaults instantiates a new ChangeLogsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemsPerPage

`func (o *ChangeLogsList) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *ChangeLogsList) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *ChangeLogsList) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *ChangeLogsList) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetStartIndex

`func (o *ChangeLogsList) GetStartIndex() int64`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *ChangeLogsList) GetStartIndexOk() (*int64, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *ChangeLogsList) SetStartIndex(v int64)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *ChangeLogsList) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### GetTotalResults

`func (o *ChangeLogsList) GetTotalResults() int64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *ChangeLogsList) GetTotalResultsOk() (*int64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *ChangeLogsList) SetTotalResults(v int64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *ChangeLogsList) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetResources

`func (o *ChangeLogsList) GetResources() []ChangeLog`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ChangeLogsList) GetResourcesOk() (*[]ChangeLog, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ChangeLogsList) SetResources(v []ChangeLog)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


