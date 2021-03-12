# Insight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]InsightActions**](InsightActions.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InsightId** | Pointer to **string** |  | [optional] 
**InsightModelType** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiTokenMeta**](ApiToken_meta.md) |  | [optional] 
**ModelIds** | Pointer to **[]string** |  | [optional] 
**Payload** | Pointer to **interface{}** |  | [optional] 
**ReadBy** | Pointer to **[]string** |  | [optional] 
**Severity** | Pointer to **int64** |  | [optional] 
**SnoozeUntil** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInsight

`func NewInsight() *Insight`

NewInsight instantiates a new Insight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightWithDefaults

`func NewInsightWithDefaults() *Insight`

NewInsightWithDefaults instantiates a new Insight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *Insight) GetActions() []InsightActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *Insight) GetActionsOk() (*[]InsightActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *Insight) SetActions(v []InsightActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *Insight) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetDisplayName

`func (o *Insight) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Insight) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Insight) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Insight) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *Insight) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Insight) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Insight) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Insight) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsightId

`func (o *Insight) GetInsightId() string`

GetInsightId returns the InsightId field if non-nil, zero value otherwise.

### GetInsightIdOk

`func (o *Insight) GetInsightIdOk() (*string, bool)`

GetInsightIdOk returns a tuple with the InsightId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightId

`func (o *Insight) SetInsightId(v string)`

SetInsightId sets InsightId field to given value.

### HasInsightId

`func (o *Insight) HasInsightId() bool`

HasInsightId returns a boolean if a field has been set.

### GetInsightModelType

`func (o *Insight) GetInsightModelType() string`

GetInsightModelType returns the InsightModelType field if non-nil, zero value otherwise.

### GetInsightModelTypeOk

`func (o *Insight) GetInsightModelTypeOk() (*string, bool)`

GetInsightModelTypeOk returns a tuple with the InsightModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightModelType

`func (o *Insight) SetInsightModelType(v string)`

SetInsightModelType sets InsightModelType field to given value.

### HasInsightModelType

`func (o *Insight) HasInsightModelType() bool`

HasInsightModelType returns a boolean if a field has been set.

### GetKind

`func (o *Insight) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Insight) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Insight) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Insight) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *Insight) GetMeta() ApiTokenMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Insight) GetMetaOk() (*ApiTokenMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Insight) SetMeta(v ApiTokenMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Insight) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetModelIds

`func (o *Insight) GetModelIds() []string`

GetModelIds returns the ModelIds field if non-nil, zero value otherwise.

### GetModelIdsOk

`func (o *Insight) GetModelIdsOk() (*[]string, bool)`

GetModelIdsOk returns a tuple with the ModelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelIds

`func (o *Insight) SetModelIds(v []string)`

SetModelIds sets ModelIds field to given value.

### HasModelIds

`func (o *Insight) HasModelIds() bool`

HasModelIds returns a boolean if a field has been set.

### GetPayload

`func (o *Insight) GetPayload() interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Insight) GetPayloadOk() (*interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Insight) SetPayload(v interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *Insight) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *Insight) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *Insight) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetReadBy

`func (o *Insight) GetReadBy() []string`

GetReadBy returns the ReadBy field if non-nil, zero value otherwise.

### GetReadByOk

`func (o *Insight) GetReadByOk() (*[]string, bool)`

GetReadByOk returns a tuple with the ReadBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadBy

`func (o *Insight) SetReadBy(v []string)`

SetReadBy sets ReadBy field to given value.

### HasReadBy

`func (o *Insight) HasReadBy() bool`

HasReadBy returns a boolean if a field has been set.

### GetSeverity

`func (o *Insight) GetSeverity() int64`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Insight) GetSeverityOk() (*int64, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Insight) SetSeverity(v int64)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Insight) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSnoozeUntil

`func (o *Insight) GetSnoozeUntil() time.Time`

GetSnoozeUntil returns the SnoozeUntil field if non-nil, zero value otherwise.

### GetSnoozeUntilOk

`func (o *Insight) GetSnoozeUntilOk() (*time.Time, bool)`

GetSnoozeUntilOk returns a tuple with the SnoozeUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnoozeUntil

`func (o *Insight) SetSnoozeUntil(v time.Time)`

SetSnoozeUntil sets SnoozeUntil field to given value.

### HasSnoozeUntil

`func (o *Insight) HasSnoozeUntil() bool`

HasSnoozeUntil returns a boolean if a field has been set.

### GetStatus

`func (o *Insight) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Insight) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Insight) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Insight) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *Insight) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Insight) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Insight) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Insight) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


