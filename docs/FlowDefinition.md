# FlowDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusExplanations** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**EndedAt** | Pointer to **time.Time** |  | [optional] 
**Flow** | Pointer to [**FlowDefinitionFlow**](FlowDefinitionFlow.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewFlowDefinition

`func NewFlowDefinition() *FlowDefinition`

NewFlowDefinition instantiates a new FlowDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowDefinitionWithDefaults

`func NewFlowDefinitionWithDefaults() *FlowDefinition`

NewFlowDefinitionWithDefaults instantiates a new FlowDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusExplanations

`func (o *FlowDefinition) GetStatusExplanations() []map[string]interface{}`

GetStatusExplanations returns the StatusExplanations field if non-nil, zero value otherwise.

### GetStatusExplanationsOk

`func (o *FlowDefinition) GetStatusExplanationsOk() (*[]map[string]interface{}, bool)`

GetStatusExplanationsOk returns a tuple with the StatusExplanations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusExplanations

`func (o *FlowDefinition) SetStatusExplanations(v []map[string]interface{})`

SetStatusExplanations sets StatusExplanations field to given value.

### HasStatusExplanations

`func (o *FlowDefinition) HasStatusExplanations() bool`

HasStatusExplanations returns a boolean if a field has been set.

### GetDescription

`func (o *FlowDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlowDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlowDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FlowDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *FlowDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FlowDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FlowDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FlowDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEndedAt

`func (o *FlowDefinition) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *FlowDefinition) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *FlowDefinition) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *FlowDefinition) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetFlow

`func (o *FlowDefinition) GetFlow() FlowDefinitionFlow`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *FlowDefinition) GetFlowOk() (*FlowDefinitionFlow, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *FlowDefinition) SetFlow(v FlowDefinitionFlow)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *FlowDefinition) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetId

`func (o *FlowDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FlowDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartedAt

`func (o *FlowDefinition) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *FlowDefinition) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *FlowDefinition) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *FlowDefinition) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *FlowDefinition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlowDefinition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlowDefinition) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlowDefinition) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVariables

`func (o *FlowDefinition) GetVariables() interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *FlowDefinition) GetVariablesOk() (*interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *FlowDefinition) SetVariables(v interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *FlowDefinition) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *FlowDefinition) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *FlowDefinition) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


