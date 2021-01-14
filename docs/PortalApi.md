# \PortalApi

All URIs are relative to *http://dev.local.crossid.io:8000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PortalFindRequest**](PortalApi.md#PortalFindRequest) | **Post** /portal/.gsearch | Make a portal specific resources search
[**PortalRequest**](PortalApi.md#PortalRequest) | **Post** /portal/request | Make a portal request



## PortalFindRequest

> ResourcesList PortalFindRequest(ctx).SearchBody(searchBody).Execute()

Make a portal specific resources search

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
    searchBody := *openapiclient.NewSearchBody("SearchFor_example") // SearchBody | Portal General Search parameters

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortalApi.PortalFindRequest(context.Background()).SearchBody(searchBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortalApi.PortalFindRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PortalFindRequest`: ResourcesList
    fmt.Fprintf(os.Stdout, "Response from `PortalApi.PortalFindRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPortalFindRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchBody** | [**SearchBody**](SearchBody.md) | Portal General Search parameters | 

### Return type

[**ResourcesList**](resourcesList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PortalRequest

> EntitlementReqResponse PortalRequest(ctx).EntitlementRequest(entitlementRequest).Reason(reason).Execute()

Make a portal request

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
    entitlementRequest := *openapiclient.NewEntitlementRequest() // EntitlementRequest | dsadsa
    reason := "required due to..." // string | a descriptive reason of the change  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortalApi.PortalRequest(context.Background()).EntitlementRequest(entitlementRequest).Reason(reason).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortalApi.PortalRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PortalRequest`: EntitlementReqResponse
    fmt.Fprintf(os.Stdout, "Response from `PortalApi.PortalRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPortalRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entitlementRequest** | [**EntitlementRequest**](EntitlementRequest.md) | dsadsa | 
 **reason** | **string** | a descriptive reason of the change  | 

### Return type

[**EntitlementReqResponse**](EntitlementReqResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

