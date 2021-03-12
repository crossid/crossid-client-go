# \InsightsApi

All URIs are relative to *http://dev.local.crossid.io:8000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InsightDiscovery**](InsightsApi.md#InsightDiscovery) | **Post** /insights/discover | Send classication data to ML server.



## InsightDiscovery

> InsightList InsightDiscovery(ctx).Execute()

Send classication data to ML server.

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
    resp, r, err := api_client.InsightsApi.InsightDiscovery(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsApi.InsightDiscovery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightDiscovery`: InsightList
    fmt.Fprintf(os.Stdout, "Response from `InsightsApi.InsightDiscovery`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInsightDiscoveryRequest struct via the builder pattern


### Return type

[**InsightList**](insightList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

