# JSONSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | externally added unique identifier. | [optional] 
**Id** | Pointer to **string** | unique identifier of the schema. | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewJSONSchema

`func NewJSONSchema() *JSONSchema`

NewJSONSchema instantiates a new JSONSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSONSchemaWithDefaults

`func NewJSONSchemaWithDefaults() *JSONSchema`

NewJSONSchemaWithDefaults instantiates a new JSONSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *JSONSchema) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *JSONSchema) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *JSONSchema) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *JSONSchema) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *JSONSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JSONSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JSONSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JSONSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMeta

`func (o *JSONSchema) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *JSONSchema) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *JSONSchema) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *JSONSchema) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


