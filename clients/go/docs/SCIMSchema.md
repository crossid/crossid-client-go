# SCIMSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**[]SCIMSchemaAttribute**](SCIMSchemaAttribute.md) |  | [optional] 
**Description** | Pointer to **string** | a more detailed description. | [optional] 
**Id** | Pointer to **string** | unique identifier of the schema. | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 
**Name** | Pointer to **string** | the name of the schema. | [optional] 

## Methods

### NewSCIMSchema

`func NewSCIMSchema() *SCIMSchema`

NewSCIMSchema instantiates a new SCIMSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCIMSchemaWithDefaults

`func NewSCIMSchemaWithDefaults() *SCIMSchema`

NewSCIMSchemaWithDefaults instantiates a new SCIMSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *SCIMSchema) GetAttributes() []SCIMSchemaAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SCIMSchema) GetAttributesOk() (*[]SCIMSchemaAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SCIMSchema) SetAttributes(v []SCIMSchemaAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SCIMSchema) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDescription

`func (o *SCIMSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SCIMSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SCIMSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SCIMSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *SCIMSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SCIMSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SCIMSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SCIMSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMeta

`func (o *SCIMSchema) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SCIMSchema) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SCIMSchema) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SCIMSchema) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetName

`func (o *SCIMSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SCIMSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SCIMSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SCIMSchema) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


