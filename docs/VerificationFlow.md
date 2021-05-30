# VerificationFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Methods** | Pointer to [**VerificationFlowMethods**](VerificationFlow_methods.md) |  | [optional] 

## Methods

### NewVerificationFlow

`func NewVerificationFlow() *VerificationFlow`

NewVerificationFlow instantiates a new VerificationFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationFlowWithDefaults

`func NewVerificationFlowWithDefaults() *VerificationFlow`

NewVerificationFlowWithDefaults instantiates a new VerificationFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VerificationFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VerificationFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VerificationFlow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VerificationFlow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMethod

`func (o *VerificationFlow) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *VerificationFlow) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *VerificationFlow) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *VerificationFlow) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethods

`func (o *VerificationFlow) GetMethods() VerificationFlowMethods`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *VerificationFlow) GetMethodsOk() (*VerificationFlowMethods, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *VerificationFlow) SetMethods(v VerificationFlowMethods)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *VerificationFlow) HasMethods() bool`

HasMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


