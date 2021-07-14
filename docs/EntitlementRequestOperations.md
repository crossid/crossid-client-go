# EntitlementRequestOperations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedEntitlement** | Pointer to [**EntitlementRequestRequestedEntitlement**](EntitlementRequestRequestedEntitlement.md) |  | [optional] 
**RequestedFor** | Pointer to **string** |  | [optional] 

## Methods

### NewEntitlementRequestOperations

`func NewEntitlementRequestOperations() *EntitlementRequestOperations`

NewEntitlementRequestOperations instantiates a new EntitlementRequestOperations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementRequestOperationsWithDefaults

`func NewEntitlementRequestOperationsWithDefaults() *EntitlementRequestOperations`

NewEntitlementRequestOperationsWithDefaults instantiates a new EntitlementRequestOperations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedEntitlement

`func (o *EntitlementRequestOperations) GetRequestedEntitlement() EntitlementRequestRequestedEntitlement`

GetRequestedEntitlement returns the RequestedEntitlement field if non-nil, zero value otherwise.

### GetRequestedEntitlementOk

`func (o *EntitlementRequestOperations) GetRequestedEntitlementOk() (*EntitlementRequestRequestedEntitlement, bool)`

GetRequestedEntitlementOk returns a tuple with the RequestedEntitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedEntitlement

`func (o *EntitlementRequestOperations) SetRequestedEntitlement(v EntitlementRequestRequestedEntitlement)`

SetRequestedEntitlement sets RequestedEntitlement field to given value.

### HasRequestedEntitlement

`func (o *EntitlementRequestOperations) HasRequestedEntitlement() bool`

HasRequestedEntitlement returns a boolean if a field has been set.

### GetRequestedFor

`func (o *EntitlementRequestOperations) GetRequestedFor() string`

GetRequestedFor returns the RequestedFor field if non-nil, zero value otherwise.

### GetRequestedForOk

`func (o *EntitlementRequestOperations) GetRequestedForOk() (*string, bool)`

GetRequestedForOk returns a tuple with the RequestedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedFor

`func (o *EntitlementRequestOperations) SetRequestedFor(v string)`

SetRequestedFor sets RequestedFor field to given value.

### HasRequestedFor

`func (o *EntitlementRequestOperations) HasRequestedFor() bool`

HasRequestedFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


