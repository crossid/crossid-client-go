# Rule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**DisplayName** | **string** |  | 
**Filter** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Keywords** | Pointer to **[]string** |  | [optional] 
**Lifecycle** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 
**ModelType** | Pointer to **string** |  | [optional] 
**Result** | [**RuleResult**](RuleResult.md) |  | 

## Methods

### NewRule

`func NewRule(displayName string, id string, result RuleResult, ) *Rule`

NewRule instantiates a new Rule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleWithDefaults

`func NewRuleWithDefaults() *Rule`

NewRuleWithDefaults instantiates a new Rule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *Rule) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Rule) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Rule) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Rule) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDisplayName

`func (o *Rule) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Rule) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Rule) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetFilter

`func (o *Rule) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *Rule) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *Rule) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *Rule) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetId

`func (o *Rule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Rule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Rule) SetId(v string)`

SetId sets Id field to given value.


### GetKeywords

`func (o *Rule) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *Rule) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *Rule) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *Rule) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetLifecycle

`func (o *Rule) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *Rule) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *Rule) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *Rule) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetMeta

`func (o *Rule) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Rule) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Rule) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Rule) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetModelType

`func (o *Rule) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *Rule) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *Rule) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *Rule) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### GetResult

`func (o *Rule) GetResult() RuleResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Rule) GetResultOk() (*RuleResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Rule) SetResult(v RuleResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


