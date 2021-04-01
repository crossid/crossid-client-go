# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssigneeId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Due** | Pointer to **time.Time** |  | [optional] 
**Ended** | Pointer to **time.Time** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Flow** | Pointer to [**TaskFlow**](TaskFlow.md) |  | [optional] 
**FormId** | Pointer to **string** |  | [optional] 
**GroupsIds** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiTokenMeta**](ApiTokenMeta.md) |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UsersIds** | Pointer to **[]string** |  | [optional] 
**Variables** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewTask

`func NewTask() *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssigneeId

`func (o *Task) GetAssigneeId() string`

GetAssigneeId returns the AssigneeId field if non-nil, zero value otherwise.

### GetAssigneeIdOk

`func (o *Task) GetAssigneeIdOk() (*string, bool)`

GetAssigneeIdOk returns a tuple with the AssigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeId

`func (o *Task) SetAssigneeId(v string)`

SetAssigneeId sets AssigneeId field to given value.

### HasAssigneeId

`func (o *Task) HasAssigneeId() bool`

HasAssigneeId returns a boolean if a field has been set.

### GetDescription

`func (o *Task) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Task) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Task) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Task) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *Task) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Task) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Task) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Task) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDue

`func (o *Task) GetDue() time.Time`

GetDue returns the Due field if non-nil, zero value otherwise.

### GetDueOk

`func (o *Task) GetDueOk() (*time.Time, bool)`

GetDueOk returns a tuple with the Due field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDue

`func (o *Task) SetDue(v time.Time)`

SetDue sets Due field to given value.

### HasDue

`func (o *Task) HasDue() bool`

HasDue returns a boolean if a field has been set.

### GetEnded

`func (o *Task) GetEnded() time.Time`

GetEnded returns the Ended field if non-nil, zero value otherwise.

### GetEndedOk

`func (o *Task) GetEndedOk() (*time.Time, bool)`

GetEndedOk returns a tuple with the Ended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnded

`func (o *Task) SetEnded(v time.Time)`

SetEnded sets Ended field to given value.

### HasEnded

`func (o *Task) HasEnded() bool`

HasEnded returns a boolean if a field has been set.

### GetExternalId

`func (o *Task) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Task) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Task) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Task) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetFlow

`func (o *Task) GetFlow() TaskFlow`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *Task) GetFlowOk() (*TaskFlow, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *Task) SetFlow(v TaskFlow)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *Task) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetFormId

`func (o *Task) GetFormId() string`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *Task) GetFormIdOk() (*string, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *Task) SetFormId(v string)`

SetFormId sets FormId field to given value.

### HasFormId

`func (o *Task) HasFormId() bool`

HasFormId returns a boolean if a field has been set.

### GetGroupsIds

`func (o *Task) GetGroupsIds() []string`

GetGroupsIds returns the GroupsIds field if non-nil, zero value otherwise.

### GetGroupsIdsOk

`func (o *Task) GetGroupsIdsOk() (*[]string, bool)`

GetGroupsIdsOk returns a tuple with the GroupsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsIds

`func (o *Task) SetGroupsIds(v []string)`

SetGroupsIds sets GroupsIds field to given value.

### HasGroupsIds

`func (o *Task) HasGroupsIds() bool`

HasGroupsIds returns a boolean if a field has been set.

### GetId

`func (o *Task) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Task) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Task) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Task) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMeta

`func (o *Task) GetMeta() ApiTokenMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Task) GetMetaOk() (*ApiTokenMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Task) SetMeta(v ApiTokenMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Task) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetOwnerId

`func (o *Task) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Task) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Task) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Task) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetStatus

`func (o *Task) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Task) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Task) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Task) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsersIds

`func (o *Task) GetUsersIds() []string`

GetUsersIds returns the UsersIds field if non-nil, zero value otherwise.

### GetUsersIdsOk

`func (o *Task) GetUsersIdsOk() (*[]string, bool)`

GetUsersIdsOk returns a tuple with the UsersIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersIds

`func (o *Task) SetUsersIds(v []string)`

SetUsersIds sets UsersIds field to given value.

### HasUsersIds

`func (o *Task) HasUsersIds() bool`

HasUsersIds returns a boolean if a field has been set.

### GetVariables

`func (o *Task) GetVariables() interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Task) GetVariablesOk() (*interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Task) SetVariables(v interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *Task) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *Task) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *Task) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


