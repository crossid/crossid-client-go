# HealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **map[string]string** |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewHealthResponse

`func NewHealthResponse(status string, ) *HealthResponse`

NewHealthResponse instantiates a new HealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthResponseWithDefaults

`func NewHealthResponseWithDefaults() *HealthResponse`

NewHealthResponseWithDefaults instantiates a new HealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *HealthResponse) GetErrors() map[string]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *HealthResponse) GetErrorsOk() (*map[string]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *HealthResponse) SetErrors(v map[string]string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *HealthResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetStatus

`func (o *HealthResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


