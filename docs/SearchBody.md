# SearchBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppIds** | Pointer to **[]string** |  | [optional] 
**Attributes** | Pointer to **[]string** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] [default to 10]
**SearchFor** | **string** |  | 
**StartIndex** | Pointer to **int32** |  | [optional] [default to 0]
**Term** | Pointer to **string** |  | [optional] 

## Methods

### NewSearchBody

`func NewSearchBody(searchFor string, ) *SearchBody`

NewSearchBody instantiates a new SearchBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchBodyWithDefaults

`func NewSearchBodyWithDefaults() *SearchBody`

NewSearchBodyWithDefaults instantiates a new SearchBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppIds

`func (o *SearchBody) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *SearchBody) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *SearchBody) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *SearchBody) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.

### GetAttributes

`func (o *SearchBody) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SearchBody) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SearchBody) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SearchBody) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetCount

`func (o *SearchBody) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SearchBody) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SearchBody) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SearchBody) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSearchFor

`func (o *SearchBody) GetSearchFor() string`

GetSearchFor returns the SearchFor field if non-nil, zero value otherwise.

### GetSearchForOk

`func (o *SearchBody) GetSearchForOk() (*string, bool)`

GetSearchForOk returns a tuple with the SearchFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFor

`func (o *SearchBody) SetSearchFor(v string)`

SetSearchFor sets SearchFor field to given value.


### GetStartIndex

`func (o *SearchBody) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *SearchBody) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *SearchBody) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *SearchBody) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### GetTerm

`func (o *SearchBody) GetTerm() string`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *SearchBody) GetTermOk() (*string, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *SearchBody) SetTerm(v string)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *SearchBody) HasTerm() bool`

HasTerm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


