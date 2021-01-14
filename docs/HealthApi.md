# \HealthApi

All URIs are relative to *http://dev.local.crossid.io:8000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HealthAlive**](HealthApi.md#HealthAlive) | **Get** /health/alive | Check aliveness
[**HealthReady**](HealthApi.md#HealthReady) | **Get** /health/ready | Check readiness



## HealthAlive

> HealthResponse HealthAlive(ctx).Execute()

Check aliveness



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HealthApi.HealthAlive(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthApi.HealthAlive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HealthAlive`: HealthResponse
    fmt.Fprintf(os.Stdout, "Response from `HealthApi.HealthAlive`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthAliveRequest struct via the builder pattern


### Return type

[**HealthResponse**](HealthResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealthReady

> HealthResponse HealthReady(ctx).Execute()

Check readiness



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HealthApi.HealthReady(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthApi.HealthReady``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HealthReady`: HealthResponse
    fmt.Fprintf(os.Stdout, "Response from `HealthApi.HealthReady`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthReadyRequest struct via the builder pattern


### Return type

[**HealthResponse**](HealthResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

