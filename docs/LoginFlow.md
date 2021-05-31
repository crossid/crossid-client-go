# LoginFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Forced** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Methods** | Pointer to [**LoginFlowMethods**](LoginFlowMethods.md) |  | [optional] 
**SsoChallenge** | Pointer to **string** |  | [optional] 

## Methods

### NewLoginFlow

`func NewLoginFlow() *LoginFlow`

NewLoginFlow instantiates a new LoginFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginFlowWithDefaults

`func NewLoginFlowWithDefaults() *LoginFlow`

NewLoginFlowWithDefaults instantiates a new LoginFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForced

`func (o *LoginFlow) GetForced() bool`

GetForced returns the Forced field if non-nil, zero value otherwise.

### GetForcedOk

`func (o *LoginFlow) GetForcedOk() (*bool, bool)`

GetForcedOk returns a tuple with the Forced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForced

`func (o *LoginFlow) SetForced(v bool)`

SetForced sets Forced field to given value.

### HasForced

`func (o *LoginFlow) HasForced() bool`

HasForced returns a boolean if a field has been set.

### GetId

`func (o *LoginFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoginFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoginFlow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LoginFlow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMethod

`func (o *LoginFlow) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LoginFlow) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LoginFlow) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LoginFlow) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethods

`func (o *LoginFlow) GetMethods() LoginFlowMethods`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *LoginFlow) GetMethodsOk() (*LoginFlowMethods, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *LoginFlow) SetMethods(v LoginFlowMethods)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *LoginFlow) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### GetSsoChallenge

`func (o *LoginFlow) GetSsoChallenge() string`

GetSsoChallenge returns the SsoChallenge field if non-nil, zero value otherwise.

### GetSsoChallengeOk

`func (o *LoginFlow) GetSsoChallengeOk() (*string, bool)`

GetSsoChallengeOk returns a tuple with the SsoChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoChallenge

`func (o *LoginFlow) SetSsoChallenge(v string)`

SetSsoChallenge sets SsoChallenge field to given value.

### HasSsoChallenge

`func (o *LoginFlow) HasSsoChallenge() bool`

HasSsoChallenge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


