# RegistrationFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Methods** | Pointer to [**RegistrationFlowMethods**](RegistrationFlowMethods.md) |  | [optional] 

## Methods

### NewRegistrationFlow

`func NewRegistrationFlow() *RegistrationFlow`

NewRegistrationFlow instantiates a new RegistrationFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationFlowWithDefaults

`func NewRegistrationFlowWithDefaults() *RegistrationFlow`

NewRegistrationFlowWithDefaults instantiates a new RegistrationFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegistrationFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistrationFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistrationFlow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RegistrationFlow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMethod

`func (o *RegistrationFlow) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RegistrationFlow) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RegistrationFlow) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RegistrationFlow) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethods

`func (o *RegistrationFlow) GetMethods() RegistrationFlowMethods`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *RegistrationFlow) GetMethodsOk() (*RegistrationFlowMethods, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *RegistrationFlow) SetMethods(v RegistrationFlowMethods)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *RegistrationFlow) HasMethods() bool`

HasMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


