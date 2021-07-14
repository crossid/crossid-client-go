# CrossFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**EndedAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Id** | **string** |  | 
**StartedAt** | **time.Time** |  | 
**Status** | **string** |  | 
**StepId** | **string** |  | 
**Steps** | [**CrossFlowSteps**](CrossFlowSteps.md) |  | 
**Vars** | **map[string]map[string]interface{}** |  | 

## Methods

### NewCrossFlow

`func NewCrossFlow(displayName string, id string, startedAt time.Time, status string, stepId string, steps CrossFlowSteps, vars map[string]map[string]interface{}, ) *CrossFlow`

NewCrossFlow instantiates a new CrossFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossFlowWithDefaults

`func NewCrossFlowWithDefaults() *CrossFlow`

NewCrossFlowWithDefaults instantiates a new CrossFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CrossFlow) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CrossFlow) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CrossFlow) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEndedAt

`func (o *CrossFlow) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *CrossFlow) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *CrossFlow) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *CrossFlow) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CrossFlow) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CrossFlow) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CrossFlow) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CrossFlow) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *CrossFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CrossFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CrossFlow) SetId(v string)`

SetId sets Id field to given value.


### GetStartedAt

`func (o *CrossFlow) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CrossFlow) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CrossFlow) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetStatus

`func (o *CrossFlow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CrossFlow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CrossFlow) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStepId

`func (o *CrossFlow) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *CrossFlow) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *CrossFlow) SetStepId(v string)`

SetStepId sets StepId field to given value.


### GetSteps

`func (o *CrossFlow) GetSteps() CrossFlowSteps`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *CrossFlow) GetStepsOk() (*CrossFlowSteps, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *CrossFlow) SetSteps(v CrossFlowSteps)`

SetSteps sets Steps field to given value.


### GetVars

`func (o *CrossFlow) GetVars() map[string]map[string]interface{}`

GetVars returns the Vars field if non-nil, zero value otherwise.

### GetVarsOk

`func (o *CrossFlow) GetVarsOk() (*map[string]map[string]interface{}, bool)`

GetVarsOk returns a tuple with the Vars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVars

`func (o *CrossFlow) SetVars(v map[string]map[string]interface{})`

SetVars sets Vars field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


