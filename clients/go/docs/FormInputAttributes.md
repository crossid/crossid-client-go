# FormInputAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autocomplete** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**MaxLength** | Pointer to **float32** |  | [optional] 
**MinLength** | Pointer to **float32** |  | [optional] 
**Name** | **string** |  | 
**Pattern** | Pointer to **string** |  | [optional] 
**Placeholder** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Transient** | Pointer to **bool** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewFormInputAttributes

`func NewFormInputAttributes(name string, type_ string, ) *FormInputAttributes`

NewFormInputAttributes instantiates a new FormInputAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormInputAttributesWithDefaults

`func NewFormInputAttributesWithDefaults() *FormInputAttributes`

NewFormInputAttributesWithDefaults instantiates a new FormInputAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutocomplete

`func (o *FormInputAttributes) GetAutocomplete() string`

GetAutocomplete returns the Autocomplete field if non-nil, zero value otherwise.

### GetAutocompleteOk

`func (o *FormInputAttributes) GetAutocompleteOk() (*string, bool)`

GetAutocompleteOk returns a tuple with the Autocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocomplete

`func (o *FormInputAttributes) SetAutocomplete(v string)`

SetAutocomplete sets Autocomplete field to given value.

### HasAutocomplete

`func (o *FormInputAttributes) HasAutocomplete() bool`

HasAutocomplete returns a boolean if a field has been set.

### GetDisabled

`func (o *FormInputAttributes) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *FormInputAttributes) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *FormInputAttributes) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *FormInputAttributes) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetLabel

`func (o *FormInputAttributes) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormInputAttributes) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormInputAttributes) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FormInputAttributes) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMaxLength

`func (o *FormInputAttributes) GetMaxLength() float32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *FormInputAttributes) GetMaxLengthOk() (*float32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *FormInputAttributes) SetMaxLength(v float32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *FormInputAttributes) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetMinLength

`func (o *FormInputAttributes) GetMinLength() float32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *FormInputAttributes) GetMinLengthOk() (*float32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *FormInputAttributes) SetMinLength(v float32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *FormInputAttributes) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetName

`func (o *FormInputAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormInputAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormInputAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetPattern

`func (o *FormInputAttributes) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *FormInputAttributes) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *FormInputAttributes) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *FormInputAttributes) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetPlaceholder

`func (o *FormInputAttributes) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *FormInputAttributes) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *FormInputAttributes) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *FormInputAttributes) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetRequired

`func (o *FormInputAttributes) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FormInputAttributes) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FormInputAttributes) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FormInputAttributes) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetTransient

`func (o *FormInputAttributes) GetTransient() bool`

GetTransient returns the Transient field if non-nil, zero value otherwise.

### GetTransientOk

`func (o *FormInputAttributes) GetTransientOk() (*bool, bool)`

GetTransientOk returns a tuple with the Transient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransient

`func (o *FormInputAttributes) SetTransient(v bool)`

SetTransient sets Transient field to given value.

### HasTransient

`func (o *FormInputAttributes) HasTransient() bool`

HasTransient returns a boolean if a field has been set.

### GetType

`func (o *FormInputAttributes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormInputAttributes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormInputAttributes) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


