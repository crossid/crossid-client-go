# \InternalApi

All URIs are relative to *http://dev.local.crossid.io:8000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EvaluateResource**](InternalApi.md#EvaluateResource) | **Get** /resources/{appID}/{resourceTypes}/{id}/.evaluate | test a resource against a rules filter and return an explanation of what needs to be patched
[**RemoveBrokenRefs**](InternalApi.md#RemoveBrokenRefs) | **Post** /internal/internal/remove-broken-refs | Returns a summary of all broken references in resources in the system
[**RemoveExpiredRefs**](InternalApi.md#RemoveExpiredRefs) | **Post** /internal/internal/remove-expired-refs | Returns a summary of all expired references in resources in the system.



## EvaluateResource

> ResourceEval EvaluateResource(ctx, appID, resourceTypes, id).Filter(filter).Execute()

test a resource against a rules filter and return an explanation of what needs to be patched

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appID := "azuread" // string | The id of the app the resource belongs to.
    resourceTypes := "Users" // string | the type of the resource.
    id := "111111" // string | the unique identifier of the resource.
    filter := "displayName sw "smith"" // string | The filter string used to request a subset of models.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternalApi.EvaluateResource(context.Background(), appID, resourceTypes, id).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.EvaluateResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EvaluateResource`: ResourceEval
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.EvaluateResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appID** | **string** | The id of the app the resource belongs to. | 
**resourceTypes** | **string** | the type of the resource. | 
**id** | **string** | the unique identifier of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **string** | The filter string used to request a subset of models.  | 

### Return type

[**ResourceEval**](ResourceEval.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveBrokenRefs

> ChangeLog RemoveBrokenRefs(ctx).RefsReqBody(refsReqBody).Reason(reason).Execute()

Returns a summary of all broken references in resources in the system

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    refsReqBody := *openapiclient.NewRefsReqBody([]string{"LimitAppIds_example"}) // RefsReqBody | Removing refs request body
    reason := "required due to..." // string | a descriptive reason of the change  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternalApi.RemoveBrokenRefs(context.Background()).RefsReqBody(refsReqBody).Reason(reason).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.RemoveBrokenRefs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveBrokenRefs`: ChangeLog
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.RemoveBrokenRefs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveBrokenRefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refsReqBody** | [**RefsReqBody**](RefsReqBody.md) | Removing refs request body | 
 **reason** | **string** | a descriptive reason of the change  | 

### Return type

[**ChangeLog**](ChangeLog.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveExpiredRefs

> ChangeLog RemoveExpiredRefs(ctx).RefsReqBody(refsReqBody).Reason(reason).Execute()

Returns a summary of all expired references in resources in the system.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    refsReqBody := *openapiclient.NewRefsReqBody([]string{"LimitAppIds_example"}) // RefsReqBody | Removing refs request body
    reason := "required due to..." // string | a descriptive reason of the change  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternalApi.RemoveExpiredRefs(context.Background()).RefsReqBody(refsReqBody).Reason(reason).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.RemoveExpiredRefs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveExpiredRefs`: ChangeLog
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.RemoveExpiredRefs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveExpiredRefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refsReqBody** | [**RefsReqBody**](RefsReqBody.md) | Removing refs request body | 
 **reason** | **string** | a descriptive reason of the change  | 

### Return type

[**ChangeLog**](ChangeLog.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

