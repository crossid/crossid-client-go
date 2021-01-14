# ResourceEval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effective** | Pointer to [**[]RuleResult**](RuleResult.md) |  | [optional] 
**Explanation** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewResourceEval

`func NewResourceEval() *ResourceEval`

NewResourceEval instantiates a new ResourceEval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceEvalWithDefaults

`func NewResourceEvalWithDefaults() *ResourceEval`

NewResourceEvalWithDefaults instantiates a new ResourceEval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffective

`func (o *ResourceEval) GetEffective() []RuleResult`

GetEffective returns the Effective field if non-nil, zero value otherwise.

### GetEffectiveOk

`func (o *ResourceEval) GetEffectiveOk() (*[]RuleResult, bool)`

GetEffectiveOk returns a tuple with the Effective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffective

`func (o *ResourceEval) SetEffective(v []RuleResult)`

SetEffective sets Effective field to given value.

### HasEffective

`func (o *ResourceEval) HasEffective() bool`

HasEffective returns a boolean if a field has been set.

### GetExplanation

`func (o *ResourceEval) GetExplanation() map[string]map[string]interface{}`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *ResourceEval) GetExplanationOk() (*map[string]map[string]interface{}, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *ResourceEval) SetExplanation(v map[string]map[string]interface{})`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *ResourceEval) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


