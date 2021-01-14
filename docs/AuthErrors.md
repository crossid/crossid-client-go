# AuthErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]AuthError**](AuthError.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthErrors

`func NewAuthErrors() *AuthErrors`

NewAuthErrors instantiates a new AuthErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthErrorsWithDefaults

`func NewAuthErrorsWithDefaults() *AuthErrors`

NewAuthErrorsWithDefaults instantiates a new AuthErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *AuthErrors) GetErrors() []AuthError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AuthErrors) GetErrorsOk() (*[]AuthError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AuthErrors) SetErrors(v []AuthError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AuthErrors) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetId

`func (o *AuthErrors) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthErrors) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthErrors) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthErrors) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


