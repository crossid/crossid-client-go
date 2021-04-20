# FlowInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusExplanations** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**EndedAt** | Pointer to **time.Time** |  | [optional] 
**Flow** | Pointer to [**FlowDefinitionFlow**](FlowDefinition_flow.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewFlowInstance

`func NewFlowInstance() *FlowInstance`

NewFlowInstance instantiates a new FlowInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowInstanceWithDefaults

`func NewFlowInstanceWithDefaults() *FlowInstance`

NewFlowInstanceWithDefaults instantiates a new FlowInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusExplanations

`func (o *FlowInstance) GetStatusExplanations() []map[string]interface{}`

GetStatusExplanations returns the StatusExplanations field if non-nil, zero value otherwise.

### GetStatusExplanationsOk

`func (o *FlowInstance) GetStatusExplanationsOk() (*[]map[string]interface{}, bool)`

GetStatusExplanationsOk returns a tuple with the StatusExplanations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusExplanations

`func (o *FlowInstance) SetStatusExplanations(v []map[string]interface{})`

SetStatusExplanations sets StatusExplanations field to given value.

### HasStatusExplanations

`func (o *FlowInstance) HasStatusExplanations() bool`

HasStatusExplanations returns a boolean if a field has been set.

### GetDescription

`func (o *FlowInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlowInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlowInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FlowInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *FlowInstance) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FlowInstance) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FlowInstance) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FlowInstance) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEndedAt

`func (o *FlowInstance) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *FlowInstance) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *FlowInstance) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *FlowInstance) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetFlow

`func (o *FlowInstance) GetFlow() FlowDefinitionFlow`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *FlowInstance) GetFlowOk() (*FlowDefinitionFlow, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *FlowInstance) SetFlow(v FlowDefinitionFlow)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *FlowInstance) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetId

`func (o *FlowInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FlowInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartedAt

`func (o *FlowInstance) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *FlowInstance) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *FlowInstance) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *FlowInstance) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *FlowInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlowInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlowInstance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlowInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVariables

`func (o *FlowInstance) GetVariables() interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *FlowInstance) GetVariablesOk() (*interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *FlowInstance) SetVariables(v interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *FlowInstance) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *FlowInstance) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *FlowInstance) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


