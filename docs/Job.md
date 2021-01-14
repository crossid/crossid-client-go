# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**By** | Pointer to **string** |  | [optional] 
**CommitId** | Pointer to **string** |  | [optional] 
**Correlation** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**FailureSize** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**LastExecutedAt** | Pointer to **time.Time** |  | [optional] 
**Meta** | Pointer to [**ApiTokenMeta**](ApiToken_meta.md) |  | [optional] 
**NotBefore** | Pointer to **time.Time** |  | [optional] 
**Payload** | Pointer to **interface{}** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int64** |  | [optional] 

## Methods

### NewJob

`func NewJob() *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBy

`func (o *Job) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *Job) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *Job) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *Job) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetCommitId

`func (o *Job) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *Job) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *Job) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.

### HasCommitId

`func (o *Job) HasCommitId() bool`

HasCommitId returns a boolean if a field has been set.

### GetCorrelation

`func (o *Job) GetCorrelation() string`

GetCorrelation returns the Correlation field if non-nil, zero value otherwise.

### GetCorrelationOk

`func (o *Job) GetCorrelationOk() (*string, bool)`

GetCorrelationOk returns a tuple with the Correlation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelation

`func (o *Job) SetCorrelation(v string)`

SetCorrelation sets Correlation field to given value.

### HasCorrelation

`func (o *Job) HasCorrelation() bool`

HasCorrelation returns a boolean if a field has been set.

### GetError

`func (o *Job) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Job) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Job) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Job) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFailureSize

`func (o *Job) GetFailureSize() int64`

GetFailureSize returns the FailureSize field if non-nil, zero value otherwise.

### GetFailureSizeOk

`func (o *Job) GetFailureSizeOk() (*int64, bool)`

GetFailureSizeOk returns a tuple with the FailureSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureSize

`func (o *Job) SetFailureSize(v int64)`

SetFailureSize sets FailureSize field to given value.

### HasFailureSize

`func (o *Job) HasFailureSize() bool`

HasFailureSize returns a boolean if a field has been set.

### GetId

`func (o *Job) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Job) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Job) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Job) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Job) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Job) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Job) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Job) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLastExecutedAt

`func (o *Job) GetLastExecutedAt() time.Time`

GetLastExecutedAt returns the LastExecutedAt field if non-nil, zero value otherwise.

### GetLastExecutedAtOk

`func (o *Job) GetLastExecutedAtOk() (*time.Time, bool)`

GetLastExecutedAtOk returns a tuple with the LastExecutedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutedAt

`func (o *Job) SetLastExecutedAt(v time.Time)`

SetLastExecutedAt sets LastExecutedAt field to given value.

### HasLastExecutedAt

`func (o *Job) HasLastExecutedAt() bool`

HasLastExecutedAt returns a boolean if a field has been set.

### GetMeta

`func (o *Job) GetMeta() ApiTokenMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Job) GetMetaOk() (*ApiTokenMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Job) SetMeta(v ApiTokenMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Job) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetNotBefore

`func (o *Job) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *Job) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *Job) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *Job) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetPayload

`func (o *Job) GetPayload() interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Job) GetPayloadOk() (*interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Job) SetPayload(v interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *Job) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *Job) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *Job) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetReason

`func (o *Job) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Job) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Job) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Job) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *Job) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Job) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Job) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Job) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *Job) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *Job) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *Job) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *Job) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


