# Form

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**Elements** | [**FormElements**](FormElements.md) |  | 
**ExternalId** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 

## Methods

### NewForm

`func NewForm(action string, elements FormElements, id string, ) *Form`

NewForm instantiates a new Form object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormWithDefaults

`func NewFormWithDefaults() *Form`

NewFormWithDefaults instantiates a new Form object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Form) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Form) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Form) SetAction(v string)`

SetAction sets Action field to given value.


### GetDisplayName

`func (o *Form) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Form) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Form) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Form) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElements

`func (o *Form) GetElements() FormElements`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *Form) GetElementsOk() (*FormElements, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *Form) SetElements(v FormElements)`

SetElements sets Elements field to given value.


### GetExternalId

`func (o *Form) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Form) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Form) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Form) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *Form) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Form) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Form) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


