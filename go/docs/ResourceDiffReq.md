# ResourceDiffReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | [**Resource**](Resource.md) |  | 
**Desired** | [**Resource**](Resource.md) |  | 
**IgnoreKeys** | Pointer to **[]string** |  | [optional] 
**Location** | **string** |  | 

## Methods

### NewResourceDiffReq

`func NewResourceDiffReq(current Resource, desired Resource, location string, ) *ResourceDiffReq`

NewResourceDiffReq instantiates a new ResourceDiffReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceDiffReqWithDefaults

`func NewResourceDiffReqWithDefaults() *ResourceDiffReq`

NewResourceDiffReqWithDefaults instantiates a new ResourceDiffReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *ResourceDiffReq) GetCurrent() Resource`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *ResourceDiffReq) GetCurrentOk() (*Resource, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *ResourceDiffReq) SetCurrent(v Resource)`

SetCurrent sets Current field to given value.


### GetDesired

`func (o *ResourceDiffReq) GetDesired() Resource`

GetDesired returns the Desired field if non-nil, zero value otherwise.

### GetDesiredOk

`func (o *ResourceDiffReq) GetDesiredOk() (*Resource, bool)`

GetDesiredOk returns a tuple with the Desired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesired

`func (o *ResourceDiffReq) SetDesired(v Resource)`

SetDesired sets Desired field to given value.


### GetIgnoreKeys

`func (o *ResourceDiffReq) GetIgnoreKeys() []string`

GetIgnoreKeys returns the IgnoreKeys field if non-nil, zero value otherwise.

### GetIgnoreKeysOk

`func (o *ResourceDiffReq) GetIgnoreKeysOk() (*[]string, bool)`

GetIgnoreKeysOk returns a tuple with the IgnoreKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreKeys

`func (o *ResourceDiffReq) SetIgnoreKeys(v []string)`

SetIgnoreKeys sets IgnoreKeys field to given value.

### HasIgnoreKeys

`func (o *ResourceDiffReq) HasIgnoreKeys() bool`

HasIgnoreKeys returns a boolean if a field has been set.

### GetLocation

`func (o *ResourceDiffReq) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ResourceDiffReq) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ResourceDiffReq) SetLocation(v string)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


