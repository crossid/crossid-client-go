# SCIMSchemaAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanonicalValues** | Pointer to **[]string** | A collection of suggested canonical values that MAY be used (e.g., \&quot;work\&quot;, \&quot;home\&quot;) | [optional] 
**CaseExact** | Pointer to **bool** | true if a string attribute is case sensitive. | [optional] 
**Description** | Pointer to **string** | a more detailed description. | [optional] 
**MultiValued** | Pointer to **bool** | true if this is a multi valued (array) attribute | [optional] 
**Mutability** | Pointer to **string** |  | [optional] 
**Name** | **string** | the name of the attribute. | 
**NoRevision** | Pointer to **bool** | true if a change of this attribute should not create a revision in history. | [optional] 
**ReferenceTypes** | Pointer to **[]string** | indicates the resource types thatare referenced. | [optional] 
**Required** | Pointer to **bool** | true if this attribute is required. | [optional] 
**Returned** | Pointer to **string** |  | [optional] 
**Search** | Pointer to [**SCIMSchemaAttributeSearch**](SCIMSchemaAttribute_search.md) |  | [optional] 
**SubAttributes** | Pointer to [**[]SCIMSchemaAttribute**](SCIMSchemaAttribute.md) |  | [optional] 
**Type** | **string** |  | 
**Uniqueness** | Pointer to **string** |  | [optional] 

## Methods

### NewSCIMSchemaAttribute

`func NewSCIMSchemaAttribute(name string, type_ string, ) *SCIMSchemaAttribute`

NewSCIMSchemaAttribute instantiates a new SCIMSchemaAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCIMSchemaAttributeWithDefaults

`func NewSCIMSchemaAttributeWithDefaults() *SCIMSchemaAttribute`

NewSCIMSchemaAttributeWithDefaults instantiates a new SCIMSchemaAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanonicalValues

`func (o *SCIMSchemaAttribute) GetCanonicalValues() []string`

GetCanonicalValues returns the CanonicalValues field if non-nil, zero value otherwise.

### GetCanonicalValuesOk

`func (o *SCIMSchemaAttribute) GetCanonicalValuesOk() (*[]string, bool)`

GetCanonicalValuesOk returns a tuple with the CanonicalValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalValues

`func (o *SCIMSchemaAttribute) SetCanonicalValues(v []string)`

SetCanonicalValues sets CanonicalValues field to given value.

### HasCanonicalValues

`func (o *SCIMSchemaAttribute) HasCanonicalValues() bool`

HasCanonicalValues returns a boolean if a field has been set.

### GetCaseExact

`func (o *SCIMSchemaAttribute) GetCaseExact() bool`

GetCaseExact returns the CaseExact field if non-nil, zero value otherwise.

### GetCaseExactOk

`func (o *SCIMSchemaAttribute) GetCaseExactOk() (*bool, bool)`

GetCaseExactOk returns a tuple with the CaseExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseExact

`func (o *SCIMSchemaAttribute) SetCaseExact(v bool)`

SetCaseExact sets CaseExact field to given value.

### HasCaseExact

`func (o *SCIMSchemaAttribute) HasCaseExact() bool`

HasCaseExact returns a boolean if a field has been set.

### GetDescription

`func (o *SCIMSchemaAttribute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SCIMSchemaAttribute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SCIMSchemaAttribute) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SCIMSchemaAttribute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMultiValued

`func (o *SCIMSchemaAttribute) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *SCIMSchemaAttribute) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *SCIMSchemaAttribute) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.

### HasMultiValued

`func (o *SCIMSchemaAttribute) HasMultiValued() bool`

HasMultiValued returns a boolean if a field has been set.

### GetMutability

`func (o *SCIMSchemaAttribute) GetMutability() string`

GetMutability returns the Mutability field if non-nil, zero value otherwise.

### GetMutabilityOk

`func (o *SCIMSchemaAttribute) GetMutabilityOk() (*string, bool)`

GetMutabilityOk returns a tuple with the Mutability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutability

`func (o *SCIMSchemaAttribute) SetMutability(v string)`

SetMutability sets Mutability field to given value.

### HasMutability

`func (o *SCIMSchemaAttribute) HasMutability() bool`

HasMutability returns a boolean if a field has been set.

### GetName

`func (o *SCIMSchemaAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SCIMSchemaAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SCIMSchemaAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetNoRevision

`func (o *SCIMSchemaAttribute) GetNoRevision() bool`

GetNoRevision returns the NoRevision field if non-nil, zero value otherwise.

### GetNoRevisionOk

`func (o *SCIMSchemaAttribute) GetNoRevisionOk() (*bool, bool)`

GetNoRevisionOk returns a tuple with the NoRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoRevision

`func (o *SCIMSchemaAttribute) SetNoRevision(v bool)`

SetNoRevision sets NoRevision field to given value.

### HasNoRevision

`func (o *SCIMSchemaAttribute) HasNoRevision() bool`

HasNoRevision returns a boolean if a field has been set.

### GetReferenceTypes

`func (o *SCIMSchemaAttribute) GetReferenceTypes() []string`

GetReferenceTypes returns the ReferenceTypes field if non-nil, zero value otherwise.

### GetReferenceTypesOk

`func (o *SCIMSchemaAttribute) GetReferenceTypesOk() (*[]string, bool)`

GetReferenceTypesOk returns a tuple with the ReferenceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceTypes

`func (o *SCIMSchemaAttribute) SetReferenceTypes(v []string)`

SetReferenceTypes sets ReferenceTypes field to given value.

### HasReferenceTypes

`func (o *SCIMSchemaAttribute) HasReferenceTypes() bool`

HasReferenceTypes returns a boolean if a field has been set.

### GetRequired

`func (o *SCIMSchemaAttribute) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *SCIMSchemaAttribute) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *SCIMSchemaAttribute) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *SCIMSchemaAttribute) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetReturned

`func (o *SCIMSchemaAttribute) GetReturned() string`

GetReturned returns the Returned field if non-nil, zero value otherwise.

### GetReturnedOk

`func (o *SCIMSchemaAttribute) GetReturnedOk() (*string, bool)`

GetReturnedOk returns a tuple with the Returned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturned

`func (o *SCIMSchemaAttribute) SetReturned(v string)`

SetReturned sets Returned field to given value.

### HasReturned

`func (o *SCIMSchemaAttribute) HasReturned() bool`

HasReturned returns a boolean if a field has been set.

### GetSearch

`func (o *SCIMSchemaAttribute) GetSearch() SCIMSchemaAttributeSearch`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *SCIMSchemaAttribute) GetSearchOk() (*SCIMSchemaAttributeSearch, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *SCIMSchemaAttribute) SetSearch(v SCIMSchemaAttributeSearch)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *SCIMSchemaAttribute) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetSubAttributes

`func (o *SCIMSchemaAttribute) GetSubAttributes() []SCIMSchemaAttribute`

GetSubAttributes returns the SubAttributes field if non-nil, zero value otherwise.

### GetSubAttributesOk

`func (o *SCIMSchemaAttribute) GetSubAttributesOk() (*[]SCIMSchemaAttribute, bool)`

GetSubAttributesOk returns a tuple with the SubAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAttributes

`func (o *SCIMSchemaAttribute) SetSubAttributes(v []SCIMSchemaAttribute)`

SetSubAttributes sets SubAttributes field to given value.

### HasSubAttributes

`func (o *SCIMSchemaAttribute) HasSubAttributes() bool`

HasSubAttributes returns a boolean if a field has been set.

### GetType

`func (o *SCIMSchemaAttribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SCIMSchemaAttribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SCIMSchemaAttribute) SetType(v string)`

SetType sets Type field to given value.


### GetUniqueness

`func (o *SCIMSchemaAttribute) GetUniqueness() string`

GetUniqueness returns the Uniqueness field if non-nil, zero value otherwise.

### GetUniquenessOk

`func (o *SCIMSchemaAttribute) GetUniquenessOk() (*string, bool)`

GetUniquenessOk returns a tuple with the Uniqueness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueness

`func (o *SCIMSchemaAttribute) SetUniqueness(v string)`

SetUniqueness sets Uniqueness field to given value.

### HasUniqueness

`func (o *SCIMSchemaAttribute) HasUniqueness() bool`

HasUniqueness returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


