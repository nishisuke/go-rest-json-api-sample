# \AccountsApi

All URIs are relative to *http://petstore.swagger.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountsGet**](AccountsApi.md#AccountsGet) | **Get** /accounts | List accounts



## AccountsGet

> []RoutesUser AccountsGet(ctx).Q(q).Execute()

List accounts



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
    q := "q_example" // string | name search by q (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.AccountsGet(context.Background()).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsGet`: []RoutesUser
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | name search by q | 

### Return type

[**[]RoutesUser**](RoutesUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

