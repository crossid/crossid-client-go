# \ApiTokensApi

All URIs are relative to *http://dev.local.crossid.io:8000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAPIToken**](ApiTokensApi.md#CreateAPIToken) | **Post** /auth/api-tokens/ | Insert an API Token
[**DeleteAPIToken**](ApiTokensApi.md#DeleteAPIToken) | **Delete** /auth/api-tokens/{id} | Delete an API token
[**ListAPITokens**](ApiTokensApi.md#ListAPITokens) | **Get** /auth/api-tokens/ | List API Tokens belonging to the authenticated user



## CreateAPIToken

> ApiToken CreateAPIToken(ctx).Reason(reason).ApiToken(apiToken).Execute()

Insert an API Token

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
    reason := "required due to..." // string | a descriptive reason of the change  (optional)
    apiToken := *openapiclient.NewApiToken() // ApiToken |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApiTokensApi.CreateAPIToken(context.Background()).Reason(reason).ApiToken(apiToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiTokensApi.CreateAPIToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAPIToken`: ApiToken
    fmt.Fprintf(os.Stdout, "Response from `ApiTokensApi.CreateAPIToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAPITokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reason** | **string** | a descriptive reason of the change  | 
 **apiToken** | [**ApiToken**](ApiToken.md) |  | 

### Return type

[**ApiToken**](ApiToken.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAPIToken

> DeleteAPIToken(ctx, id).Reason(reason).Execute()

Delete an API token

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
    id := "id_example" // string | The id of the API token to delete.
    reason := "required due to..." // string | a descriptive reason of the change  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApiTokensApi.DeleteAPIToken(context.Background(), id).Reason(reason).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiTokensApi.DeleteAPIToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the API token to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAPITokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **string** | a descriptive reason of the change  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAPITokens

> ApiTokensList ListAPITokens(ctx).Filter(filter).Count(count).StartIndex(startIndex).Execute()

List API Tokens belonging to the authenticated user

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
    filter := "displayName sw "smith"" // string | The filter string used to request a subset of models.  (optional)
    count := int64(789) // int64 | The desired maximum number of query results per page, e.g., 10. A negative value is interpreted as \"0\". A value of \"0\" indicates that no model results are to be returned except for \"totalResults\".  (optional) (default to 10)
    startIndex := int64(789) // int64 | The 1-based index of the first query result. A value less than 1 SHALL be interpreted as 1.  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApiTokensApi.ListAPITokens(context.Background()).Filter(filter).Count(count).StartIndex(startIndex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiTokensApi.ListAPITokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAPITokens`: ApiTokensList
    fmt.Fprintf(os.Stdout, "Response from `ApiTokensApi.ListAPITokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAPITokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | The filter string used to request a subset of models.  | 
 **count** | **int64** | The desired maximum number of query results per page, e.g., 10. A negative value is interpreted as \&quot;0\&quot;. A value of \&quot;0\&quot; indicates that no model results are to be returned except for \&quot;totalResults\&quot;.  | [default to 10]
 **startIndex** | **int64** | The 1-based index of the first query result. A value less than 1 SHALL be interpreted as 1.  | [default to 0]

### Return type

[**ApiTokensList**](apiTokensList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

