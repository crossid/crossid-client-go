# ApiTokensList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemsPerPage** | Pointer to **int32** |  | [optional] 
**StartIndex** | Pointer to **int64** |  | [optional] 
**TotalResults** | Pointer to **int64** |  | [optional] 
**Resources** | [**[]ApiToken**](ApiToken.md) |  | 

## Methods

### NewApiTokensList

`func NewApiTokensList(resources []ApiToken, ) *ApiTokensList`

NewApiTokensList instantiates a new ApiTokensList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTokensListWithDefaults

`func NewApiTokensListWithDefaults() *ApiTokensList`

NewApiTokensListWithDefaults instantiates a new ApiTokensList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemsPerPage

`func (o *ApiTokensList) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *ApiTokensList) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *ApiTokensList) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *ApiTokensList) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetStartIndex

`func (o *ApiTokensList) GetStartIndex() int64`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *ApiTokensList) GetStartIndexOk() (*int64, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *ApiTokensList) SetStartIndex(v int64)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *ApiTokensList) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### GetTotalResults

`func (o *ApiTokensList) GetTotalResults() int64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *ApiTokensList) GetTotalResultsOk() (*int64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *ApiTokensList) SetTotalResults(v int64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *ApiTokensList) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetResources

`func (o *ApiTokensList) GetResources() []ApiToken`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ApiTokensList) GetResourcesOk() (*[]ApiToken, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ApiTokensList) SetResources(v []ApiToken)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


