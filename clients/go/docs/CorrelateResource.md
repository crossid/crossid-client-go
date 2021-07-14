# CorrelateResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrId** | **string** | id of a filter template used to find the resources to correlate to. | 
**Filter** | **string** | used to request a subset of resources and must be a valid FILTER expression. | 
**RefPath** | **string** | The ref path create when resources are correlated. | 

## Methods

### NewCorrelateResource

`func NewCorrelateResource(corrId string, filter string, refPath string, ) *CorrelateResource`

NewCorrelateResource instantiates a new CorrelateResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorrelateResourceWithDefaults

`func NewCorrelateResourceWithDefaults() *CorrelateResource`

NewCorrelateResourceWithDefaults instantiates a new CorrelateResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrId

`func (o *CorrelateResource) GetCorrId() string`

GetCorrId returns the CorrId field if non-nil, zero value otherwise.

### GetCorrIdOk

`func (o *CorrelateResource) GetCorrIdOk() (*string, bool)`

GetCorrIdOk returns a tuple with the CorrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrId

`func (o *CorrelateResource) SetCorrId(v string)`

SetCorrId sets CorrId field to given value.


### GetFilter

`func (o *CorrelateResource) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CorrelateResource) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CorrelateResource) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetRefPath

`func (o *CorrelateResource) GetRefPath() string`

GetRefPath returns the RefPath field if non-nil, zero value otherwise.

### GetRefPathOk

`func (o *CorrelateResource) GetRefPathOk() (*string, bool)`

GetRefPathOk returns a tuple with the RefPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefPath

`func (o *CorrelateResource) SetRefPath(v string)`

SetRefPath sets RefPath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


