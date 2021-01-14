# Mapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Lang** | Pointer to **string** |  | [optional] 
**Mapper** | Pointer to **interface{}** |  | [optional] 
**Meta** | Pointer to [**ApiTokenMeta**](ApiToken_meta.md) |  | [optional] 
**Variables** | Pointer to [**[]MapperVariables**](MapperVariables.md) |  | [optional] 

## Methods

### NewMapper

`func NewMapper() *Mapper`

NewMapper instantiates a new Mapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapperWithDefaults

`func NewMapperWithDefaults() *Mapper`

NewMapperWithDefaults instantiates a new Mapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *Mapper) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Mapper) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Mapper) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Mapper) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *Mapper) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Mapper) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Mapper) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Mapper) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLang

`func (o *Mapper) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *Mapper) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *Mapper) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *Mapper) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetMapper

`func (o *Mapper) GetMapper() interface{}`

GetMapper returns the Mapper field if non-nil, zero value otherwise.

### GetMapperOk

`func (o *Mapper) GetMapperOk() (*interface{}, bool)`

GetMapperOk returns a tuple with the Mapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapper

`func (o *Mapper) SetMapper(v interface{})`

SetMapper sets Mapper field to given value.

### HasMapper

`func (o *Mapper) HasMapper() bool`

HasMapper returns a boolean if a field has been set.

### SetMapperNil

`func (o *Mapper) SetMapperNil(b bool)`

 SetMapperNil sets the value for Mapper to be an explicit nil

### UnsetMapper
`func (o *Mapper) UnsetMapper()`

UnsetMapper ensures that no value is present for Mapper, not even an explicit nil
### GetMeta

`func (o *Mapper) GetMeta() ApiTokenMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Mapper) GetMetaOk() (*ApiTokenMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Mapper) SetMeta(v ApiTokenMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Mapper) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetVariables

`func (o *Mapper) GetVariables() []MapperVariables`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Mapper) GetVariablesOk() (*[]MapperVariables, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Mapper) SetVariables(v []MapperVariables)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *Mapper) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


