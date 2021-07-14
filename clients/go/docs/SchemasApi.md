# \SchemasApi

All URIs are relative to *http://dev.local.crossid.io:8000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJSONSchema**](SchemasApi.md#CreateJSONSchema) | **Post** /json-schemas/ | Create JSON Schema
[**CreateSCIMSchema**](SchemasApi.md#CreateSCIMSchema) | **Post** /scim-schemas/ | Create a SCIM Schema
[**GetJSONSchema**](SchemasApi.md#GetJSONSchema) | **Get** /json-schemas/{id} | Get a JSON Schema
[**GetSCIMSchema**](SchemasApi.md#GetSCIMSchema) | **Get** /scim-schemas/{id} | Get a SCIM Schema
[**ReplaceJSONSchema**](SchemasApi.md#ReplaceJSONSchema) | **Put** /json-schemas/{id} | replace a JSON Schema
[**ReplaceSCIMSchema**](SchemasApi.md#ReplaceSCIMSchema) | **Put** /scim-schemas/{id} | Replace a SCIM Schema



## CreateJSONSchema

> JSONSchema CreateJSONSchema(ctx).Reason(reason).JSONSchema(jSONSchema).Execute()

Create JSON Schema

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
    jSONSchema := *openapiclient.NewJSONSchema() // JSONSchema |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemasApi.CreateJSONSchema(context.Background()).Reason(reason).JSONSchema(jSONSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.CreateJSONSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateJSONSchema`: JSONSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.CreateJSONSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJSONSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reason** | **string** | a descriptive reason of the change  | 
 **jSONSchema** | [**JSONSchema**](JSONSchema.md) |  | 

### Return type

[**JSONSchema**](JSONSchema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSCIMSchema

> SCIMSchema CreateSCIMSchema(ctx).Reason(reason).SCIMSchema(sCIMSchema).Execute()

Create a SCIM Schema

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
    sCIMSchema := *openapiclient.NewSCIMSchema() // SCIMSchema |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemasApi.CreateSCIMSchema(context.Background()).Reason(reason).SCIMSchema(sCIMSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.CreateSCIMSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSCIMSchema`: SCIMSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.CreateSCIMSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSCIMSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reason** | **string** | a descriptive reason of the change  | 
 **sCIMSchema** | [**SCIMSchema**](SCIMSchema.md) |  | 

### Return type

[**SCIMSchema**](SCIMSchema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJSONSchema

> JSONSchema GetJSONSchema(ctx, id).Execute()

Get a JSON Schema

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
    id := "id_example" // string | The id of the schema to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemasApi.GetJSONSchema(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.GetJSONSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJSONSchema`: JSONSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.GetJSONSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the schema to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJSONSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JSONSchema**](JSONSchema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSCIMSchema

> SCIMSchema GetSCIMSchema(ctx, id).Execute()

Get a SCIM Schema

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
    id := "id_example" // string | The id of the schema to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemasApi.GetSCIMSchema(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.GetSCIMSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSCIMSchema`: SCIMSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.GetSCIMSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the schema to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSCIMSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SCIMSchema**](SCIMSchema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceJSONSchema

> JSONSchema ReplaceJSONSchema(ctx, id).Reason(reason).JSONSchema(jSONSchema).Execute()

replace a JSON Schema

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
    id := "id_example" // string | The id of the schema to replace.
    reason := "required due to..." // string | a descriptive reason of the change  (optional)
    jSONSchema := *openapiclient.NewJSONSchema() // JSONSchema |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemasApi.ReplaceJSONSchema(context.Background(), id).Reason(reason).JSONSchema(jSONSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.ReplaceJSONSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceJSONSchema`: JSONSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.ReplaceJSONSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the schema to replace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceJSONSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **string** | a descriptive reason of the change  | 
 **jSONSchema** | [**JSONSchema**](JSONSchema.md) |  | 

### Return type

[**JSONSchema**](JSONSchema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceSCIMSchema

> SCIMSchema ReplaceSCIMSchema(ctx, id).Reason(reason).SCIMSchema(sCIMSchema).Execute()

Replace a SCIM Schema

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
    id := "id_example" // string | The id of the schema to replace.
    reason := "required due to..." // string | a descriptive reason of the change  (optional)
    sCIMSchema := *openapiclient.NewSCIMSchema() // SCIMSchema |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemasApi.ReplaceSCIMSchema(context.Background(), id).Reason(reason).SCIMSchema(sCIMSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.ReplaceSCIMSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceSCIMSchema`: SCIMSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.ReplaceSCIMSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the schema to replace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceSCIMSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **string** | a descriptive reason of the change  | 
 **sCIMSchema** | [**SCIMSchema**](SCIMSchema.md) |  | 

### Return type

[**SCIMSchema**](SCIMSchema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

