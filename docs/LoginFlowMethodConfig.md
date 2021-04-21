# LoginFlowMethodConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Messages** | Pointer to [**[]UserFacingMessage**](UserFacingMessage.md) |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 

## Methods

### NewLoginFlowMethodConfig

`func NewLoginFlowMethodConfig() *LoginFlowMethodConfig`

NewLoginFlowMethodConfig instantiates a new LoginFlowMethodConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginFlowMethodConfigWithDefaults

`func NewLoginFlowMethodConfigWithDefaults() *LoginFlowMethodConfig`

NewLoginFlowMethodConfigWithDefaults instantiates a new LoginFlowMethodConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *LoginFlowMethodConfig) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *LoginFlowMethodConfig) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *LoginFlowMethodConfig) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *LoginFlowMethodConfig) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetMessages

`func (o *LoginFlowMethodConfig) GetMessages() []UserFacingMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *LoginFlowMethodConfig) GetMessagesOk() (*[]UserFacingMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *LoginFlowMethodConfig) SetMessages(v []UserFacingMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *LoginFlowMethodConfig) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetMethod

`func (o *LoginFlowMethodConfig) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LoginFlowMethodConfig) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LoginFlowMethodConfig) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LoginFlowMethodConfig) HasMethod() bool`

HasMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


