# \LoginApi

All URIs are relative to *http://dev.local.crossid.io:8000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLoginFlow**](LoginApi.md#GetLoginFlow) | **Get** /auth/self-service/login/flows | Info for a specific login flow.



## GetLoginFlow

> LoginFlow GetLoginFlow(ctx).Id(id).Execute()

Info for a specific login flow.

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
    id := "id_example" // string | The id of the flow to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoginApi.GetLoginFlow(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginApi.GetLoginFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoginFlow`: LoginFlow
    fmt.Fprintf(os.Stdout, "Response from `LoginApi.GetLoginFlow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoginFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of the flow to retrieve | 

### Return type

[**LoginFlow**](LoginFlow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

