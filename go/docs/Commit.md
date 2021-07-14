# Commit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**By** | Pointer to **string** |  | [optional] 
**Correlation** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**AppMeta**](AppMeta.md) |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 

## Methods

### NewCommit

`func NewCommit() *Commit`

NewCommit instantiates a new Commit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitWithDefaults

`func NewCommitWithDefaults() *Commit`

NewCommitWithDefaults instantiates a new Commit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBy

`func (o *Commit) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *Commit) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *Commit) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *Commit) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetCorrelation

`func (o *Commit) GetCorrelation() string`

GetCorrelation returns the Correlation field if non-nil, zero value otherwise.

### GetCorrelationOk

`func (o *Commit) GetCorrelationOk() (*string, bool)`

GetCorrelationOk returns a tuple with the Correlation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelation

`func (o *Commit) SetCorrelation(v string)`

SetCorrelation sets Correlation field to given value.

### HasCorrelation

`func (o *Commit) HasCorrelation() bool`

HasCorrelation returns a boolean if a field has been set.

### GetId

`func (o *Commit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Commit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Commit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Commit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMeta

`func (o *Commit) GetMeta() AppMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Commit) GetMetaOk() (*AppMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Commit) SetMeta(v AppMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Commit) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetReason

`func (o *Commit) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Commit) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Commit) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Commit) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRequestId

`func (o *Commit) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Commit) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *Commit) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *Commit) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


