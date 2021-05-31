# ChangeLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apply** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiTokenMeta**](ApiTokenMeta.md) |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewChangeLog

`func NewChangeLog() *ChangeLog`

NewChangeLog instantiates a new ChangeLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeLogWithDefaults

`func NewChangeLogWithDefaults() *ChangeLog`

NewChangeLogWithDefaults instantiates a new ChangeLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApply

`func (o *ChangeLog) GetApply() string`

GetApply returns the Apply field if non-nil, zero value otherwise.

### GetApplyOk

`func (o *ChangeLog) GetApplyOk() (*string, bool)`

GetApplyOk returns a tuple with the Apply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApply

`func (o *ChangeLog) SetApply(v string)`

SetApply sets Apply field to given value.

### HasApply

`func (o *ChangeLog) HasApply() bool`

HasApply returns a boolean if a field has been set.

### GetId

`func (o *ChangeLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChangeLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChangeLog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChangeLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMeta

`func (o *ChangeLog) GetMeta() ApiTokenMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ChangeLog) GetMetaOk() (*ApiTokenMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ChangeLog) SetMeta(v ApiTokenMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ChangeLog) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetNote

`func (o *ChangeLog) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ChangeLog) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ChangeLog) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ChangeLog) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetState

`func (o *ChangeLog) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ChangeLog) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ChangeLog) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ChangeLog) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


