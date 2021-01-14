# \ErrorsApi

All URIs are relative to *http://dev.local.crossid.io:8000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetErrors**](ErrorsApi.md#GetErrors) | **Get** /auth/errors/ | Info for a specific errors.



## GetErrors

> AuthErrors GetErrors(ctx).Error_(error_).Execute()

Info for a specific errors.

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
    error_ := "error__example" // string | The id of the errors document to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ErrorsApi.GetErrors(context.Background()).Error_(error_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ErrorsApi.GetErrors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetErrors`: AuthErrors
    fmt.Fprintf(os.Stdout, "Response from `ErrorsApi.GetErrors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetErrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **error_** | **string** | The id of the errors document to retrieve | 

### Return type

[**AuthErrors**](AuthErrors.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

