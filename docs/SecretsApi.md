# \SecretsApi

All URIs are relative to *http://dev.local.crossid.io:8000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSecret**](SecretsApi.md#GetSecret) | **Get** /secrets/{path} | Get a Secret
[**PutSecret**](SecretsApi.md#PutSecret) | **Put** /secrets/{path} | Put a Secret



## GetSecret

> Secret GetSecret(ctx, path).Ciphertext(ciphertext).Execute()

Get a Secret

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
    path := "path_example" // string | The path to the secret
    ciphertext := "ciphertext_example" // string | the cipher text to decrypt, used by the ondemand vault engine. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecretsApi.GetSecret(context.Background(), path).Ciphertext(ciphertext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecret`: Secret
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.GetSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path to the secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ciphertext** | **string** | the cipher text to decrypt, used by the ondemand vault engine. | 

### Return type

[**Secret**](Secret.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSecret

> Secret PutSecret(ctx, path).Body(body).Execute()

Put a Secret

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
    path := "path_example" // string | The path to the secret
    body := Secret(987) // Secret | a JSON containing the information required to put a secret.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecretsApi.PutSecret(context.Background(), path).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.PutSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSecret`: Secret
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.PutSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path to the secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **Secret** | a JSON containing the information required to put a secret. | 

### Return type

[**Secret**](Secret.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

