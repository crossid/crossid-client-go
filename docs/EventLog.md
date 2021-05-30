# EventLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to [**EventLogActor**](EventLog_actor.md) |  | [optional] 
**Client** | Pointer to [**EventLogClient**](EventLog_client.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**EventLogResult**](EventLog_result.md) |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewEventLog

`func NewEventLog() *EventLog`

NewEventLog instantiates a new EventLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventLogWithDefaults

`func NewEventLogWithDefaults() *EventLog`

NewEventLogWithDefaults instantiates a new EventLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *EventLog) GetActor() EventLogActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *EventLog) GetActorOk() (*EventLogActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *EventLog) SetActor(v EventLogActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *EventLog) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetClient

`func (o *EventLog) GetClient() EventLogClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *EventLog) GetClientOk() (*EventLogClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *EventLog) SetClient(v EventLogClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *EventLog) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EventLog) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventLog) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventLog) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EventLog) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *EventLog) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EventLog) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EventLog) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EventLog) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *EventLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventLog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRequestId

`func (o *EventLog) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *EventLog) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *EventLog) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *EventLog) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetResult

`func (o *EventLog) GetResult() EventLogResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *EventLog) GetResultOk() (*EventLogResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *EventLog) SetResult(v EventLogResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *EventLog) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSeverity

`func (o *EventLog) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *EventLog) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *EventLog) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *EventLog) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetType

`func (o *EventLog) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventLog) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventLog) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventLog) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


