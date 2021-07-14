# FormElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**OneOfFormInputAttributesFormAnchorAttributesFormButtonAttributesFormTextAttributes**](oneOf&lt;FormInputAttributes,FormAnchorAttributes,FormButtonAttributes,FormTextAttributes&gt;.md) |  | 
**Group** | **string** |  | 
**Id** | **string** |  | 
**Tag** | **string** |  | 

## Methods

### NewFormElement

`func NewFormElement(attributes OneOfFormInputAttributesFormAnchorAttributesFormButtonAttributesFormTextAttributes, group string, id string, tag string, ) *FormElement`

NewFormElement instantiates a new FormElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormElementWithDefaults

`func NewFormElementWithDefaults() *FormElement`

NewFormElementWithDefaults instantiates a new FormElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *FormElement) GetAttributes() OneOfFormInputAttributesFormAnchorAttributesFormButtonAttributesFormTextAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FormElement) GetAttributesOk() (*OneOfFormInputAttributesFormAnchorAttributesFormButtonAttributesFormTextAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FormElement) SetAttributes(v OneOfFormInputAttributesFormAnchorAttributesFormButtonAttributesFormTextAttributes)`

SetAttributes sets Attributes field to given value.


### GetGroup

`func (o *FormElement) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *FormElement) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *FormElement) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetId

`func (o *FormElement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormElement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormElement) SetId(v string)`

SetId sets Id field to given value.


### GetTag

`func (o *FormElement) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *FormElement) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *FormElement) SetTag(v string)`

SetTag sets Tag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


