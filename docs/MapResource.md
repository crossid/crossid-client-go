# MapResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrId** | **string** | id of a filter template used to find the resources to map from | 
**Filter** | **string** | used to request a subset of resources and must be a valid FILTER expression. | 
**MapperId** | **string** | The mapper to be used on the matched resources to produce the attributes to be copied. If not set it tries to copy any possible attribute | 

## Methods

### NewMapResource

`func NewMapResource(corrId string, filter string, mapperId string, ) *MapResource`

NewMapResource instantiates a new MapResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapResourceWithDefaults

`func NewMapResourceWithDefaults() *MapResource`

NewMapResourceWithDefaults instantiates a new MapResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrId

`func (o *MapResource) GetCorrId() string`

GetCorrId returns the CorrId field if non-nil, zero value otherwise.

### GetCorrIdOk

`func (o *MapResource) GetCorrIdOk() (*string, bool)`

GetCorrIdOk returns a tuple with the CorrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrId

`func (o *MapResource) SetCorrId(v string)`

SetCorrId sets CorrId field to given value.


### GetFilter

`func (o *MapResource) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MapResource) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MapResource) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetMapperId

`func (o *MapResource) GetMapperId() string`

GetMapperId returns the MapperId field if non-nil, zero value otherwise.

### GetMapperIdOk

`func (o *MapResource) GetMapperIdOk() (*string, bool)`

GetMapperIdOk returns a tuple with the MapperId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapperId

`func (o *MapResource) SetMapperId(v string)`

SetMapperId sets MapperId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


