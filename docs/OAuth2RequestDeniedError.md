# OAuth2RequestDeniedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**ErrorDebug** | Pointer to **string** |  | [optional] 
**ErrorDescription** | Pointer to **string** |  | [optional] 
**ErrorHint** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewOAuth2RequestDeniedError

`func NewOAuth2RequestDeniedError() *OAuth2RequestDeniedError`

NewOAuth2RequestDeniedError instantiates a new OAuth2RequestDeniedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2RequestDeniedErrorWithDefaults

`func NewOAuth2RequestDeniedErrorWithDefaults() *OAuth2RequestDeniedError`

NewOAuth2RequestDeniedErrorWithDefaults instantiates a new OAuth2RequestDeniedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *OAuth2RequestDeniedError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OAuth2RequestDeniedError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OAuth2RequestDeniedError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *OAuth2RequestDeniedError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorDebug

`func (o *OAuth2RequestDeniedError) GetErrorDebug() string`

GetErrorDebug returns the ErrorDebug field if non-nil, zero value otherwise.

### GetErrorDebugOk

`func (o *OAuth2RequestDeniedError) GetErrorDebugOk() (*string, bool)`

GetErrorDebugOk returns a tuple with the ErrorDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDebug

`func (o *OAuth2RequestDeniedError) SetErrorDebug(v string)`

SetErrorDebug sets ErrorDebug field to given value.

### HasErrorDebug

`func (o *OAuth2RequestDeniedError) HasErrorDebug() bool`

HasErrorDebug returns a boolean if a field has been set.

### GetErrorDescription

`func (o *OAuth2RequestDeniedError) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *OAuth2RequestDeniedError) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *OAuth2RequestDeniedError) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *OAuth2RequestDeniedError) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorHint

`func (o *OAuth2RequestDeniedError) GetErrorHint() string`

GetErrorHint returns the ErrorHint field if non-nil, zero value otherwise.

### GetErrorHintOk

`func (o *OAuth2RequestDeniedError) GetErrorHintOk() (*string, bool)`

GetErrorHintOk returns a tuple with the ErrorHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorHint

`func (o *OAuth2RequestDeniedError) SetErrorHint(v string)`

SetErrorHint sets ErrorHint field to given value.

### HasErrorHint

`func (o *OAuth2RequestDeniedError) HasErrorHint() bool`

HasErrorHint returns a boolean if a field has been set.

### GetStatusCode

`func (o *OAuth2RequestDeniedError) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *OAuth2RequestDeniedError) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *OAuth2RequestDeniedError) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *OAuth2RequestDeniedError) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetType

`func (o *OAuth2RequestDeniedError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OAuth2RequestDeniedError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OAuth2RequestDeniedError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OAuth2RequestDeniedError) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


