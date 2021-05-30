# EventLogClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Geo** | Pointer to [**EventLogClientGeo**](EventLog_client_geo.md) |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**UserAgent** | Pointer to [**EventLogClientUserAgent**](EventLog_client_userAgent.md) |  | [optional] 

## Methods

### NewEventLogClient

`func NewEventLogClient() *EventLogClient`

NewEventLogClient instantiates a new EventLogClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventLogClientWithDefaults

`func NewEventLogClientWithDefaults() *EventLogClient`

NewEventLogClientWithDefaults instantiates a new EventLogClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeo

`func (o *EventLogClient) GetGeo() EventLogClientGeo`

GetGeo returns the Geo field if non-nil, zero value otherwise.

### GetGeoOk

`func (o *EventLogClient) GetGeoOk() (*EventLogClientGeo, bool)`

GetGeoOk returns a tuple with the Geo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeo

`func (o *EventLogClient) SetGeo(v EventLogClientGeo)`

SetGeo sets Geo field to given value.

### HasGeo

`func (o *EventLogClient) HasGeo() bool`

HasGeo returns a boolean if a field has been set.

### GetIp

`func (o *EventLogClient) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *EventLogClient) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *EventLogClient) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *EventLogClient) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetUserAgent

`func (o *EventLogClient) GetUserAgent() EventLogClientUserAgent`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *EventLogClient) GetUserAgentOk() (*EventLogClientUserAgent, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *EventLogClient) SetUserAgent(v EventLogClientUserAgent)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *EventLogClient) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


