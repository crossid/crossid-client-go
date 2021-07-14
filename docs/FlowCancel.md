# FlowCancel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanceledJobs** | Pointer to **[]string** | jobs created by the flow that were due to run and were canceled | [optional] 
**FlowId** | Pointer to **string** | flow instance id that was canceled | [optional] 
**FlowStatus** | Pointer to **string** | flow status after cancellation. always \&quot;canceled\&quot; | [optional] 
**NotCanceledJobs** | Pointer to **[]string** | jobs created by the flow that were due to run and failed cancel | [optional] 
**TaskIds** | Pointer to **[]string** | flow tasks that were canceled along with the flow | [optional] 

## Methods

### NewFlowCancel

`func NewFlowCancel() *FlowCancel`

NewFlowCancel instantiates a new FlowCancel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowCancelWithDefaults

`func NewFlowCancelWithDefaults() *FlowCancel`

NewFlowCancelWithDefaults instantiates a new FlowCancel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanceledJobs

`func (o *FlowCancel) GetCanceledJobs() []string`

GetCanceledJobs returns the CanceledJobs field if non-nil, zero value otherwise.

### GetCanceledJobsOk

`func (o *FlowCancel) GetCanceledJobsOk() (*[]string, bool)`

GetCanceledJobsOk returns a tuple with the CanceledJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledJobs

`func (o *FlowCancel) SetCanceledJobs(v []string)`

SetCanceledJobs sets CanceledJobs field to given value.

### HasCanceledJobs

`func (o *FlowCancel) HasCanceledJobs() bool`

HasCanceledJobs returns a boolean if a field has been set.

### GetFlowId

`func (o *FlowCancel) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *FlowCancel) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *FlowCancel) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *FlowCancel) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetFlowStatus

`func (o *FlowCancel) GetFlowStatus() string`

GetFlowStatus returns the FlowStatus field if non-nil, zero value otherwise.

### GetFlowStatusOk

`func (o *FlowCancel) GetFlowStatusOk() (*string, bool)`

GetFlowStatusOk returns a tuple with the FlowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowStatus

`func (o *FlowCancel) SetFlowStatus(v string)`

SetFlowStatus sets FlowStatus field to given value.

### HasFlowStatus

`func (o *FlowCancel) HasFlowStatus() bool`

HasFlowStatus returns a boolean if a field has been set.

### GetNotCanceledJobs

`func (o *FlowCancel) GetNotCanceledJobs() []string`

GetNotCanceledJobs returns the NotCanceledJobs field if non-nil, zero value otherwise.

### GetNotCanceledJobsOk

`func (o *FlowCancel) GetNotCanceledJobsOk() (*[]string, bool)`

GetNotCanceledJobsOk returns a tuple with the NotCanceledJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotCanceledJobs

`func (o *FlowCancel) SetNotCanceledJobs(v []string)`

SetNotCanceledJobs sets NotCanceledJobs field to given value.

### HasNotCanceledJobs

`func (o *FlowCancel) HasNotCanceledJobs() bool`

HasNotCanceledJobs returns a boolean if a field has been set.

### GetTaskIds

`func (o *FlowCancel) GetTaskIds() []string`

GetTaskIds returns the TaskIds field if non-nil, zero value otherwise.

### GetTaskIdsOk

`func (o *FlowCancel) GetTaskIdsOk() (*[]string, bool)`

GetTaskIdsOk returns a tuple with the TaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIds

`func (o *FlowCancel) SetTaskIds(v []string)`

SetTaskIds sets TaskIds field to given value.

### HasTaskIds

`func (o *FlowCancel) HasTaskIds() bool`

HasTaskIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


