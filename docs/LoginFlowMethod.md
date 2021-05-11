# LoginFlowMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**LoginFlowMethodConfig**](LoginFlowMethod_config.md) |  | [optional] 
**Level** | Pointer to **int32** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**LoginFlowMethodState**](LoginFlowMethod_state.md) |  | [optional] 

## Methods

### NewLoginFlowMethod

`func NewLoginFlowMethod() *LoginFlowMethod`

NewLoginFlowMethod instantiates a new LoginFlowMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginFlowMethodWithDefaults

`func NewLoginFlowMethodWithDefaults() *LoginFlowMethod`

NewLoginFlowMethodWithDefaults instantiates a new LoginFlowMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *LoginFlowMethod) GetConfig() LoginFlowMethodConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *LoginFlowMethod) GetConfigOk() (*LoginFlowMethodConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *LoginFlowMethod) SetConfig(v LoginFlowMethodConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *LoginFlowMethod) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLevel

`func (o *LoginFlowMethod) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *LoginFlowMethod) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *LoginFlowMethod) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *LoginFlowMethod) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetMethod

`func (o *LoginFlowMethod) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LoginFlowMethod) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LoginFlowMethod) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LoginFlowMethod) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetState

`func (o *LoginFlowMethod) GetState() LoginFlowMethodState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LoginFlowMethod) GetStateOk() (*LoginFlowMethodState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LoginFlowMethod) SetState(v LoginFlowMethodState)`

SetState sets State field to given value.

### HasState

`func (o *LoginFlowMethod) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


