# \TokensAPI

All URIs are relative to *https://svc.blockdaemon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTokenTags**](TokensAPI.md#GetTokenTags) | **Get** /tokens/tags | Get list of token tags
[**GetTokens**](TokensAPI.md#GetTokens) | **Get** /tokens | Get list of supported tokens with metadata



## GetTokenTags

> Tags GetTokenTags(ctx).Execute()

Get list of token tags



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Blockdaemon/defi-api-go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.GetTokenTags(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.GetTokenTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenTags`: Tags
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.GetTokenTags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenTagsRequest struct via the builder pattern


### Return type

[**Tags**](Tags.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokens

> TokenList GetTokens(ctx).ChainID(chainID).TokenAddress(tokenAddress).TokenSymbol(tokenSymbol).TagLimit(tagLimit).Execute()

Get list of supported tokens with metadata



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Blockdaemon/defi-api-go-sdk"
)

func main() {
	chainID := "eip155:1" // string | The CAIP-2 identifier for a blockchain. (optional)
	tokenAddress := "0x7F5c764cBc14f9669B88837ca1490cCa17c31607" // string | The address of an token contract to query. If omitted, all supported tokens will be queried. (optional)
	tokenSymbol := "USDC" // string | The symbol of an token contract. (optional)
	tagLimit := []string{"defi"} // []string | A list of token tag identifiers to include in the query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.GetTokens(context.Background()).ChainID(chainID).TokenAddress(tokenAddress).TokenSymbol(tokenSymbol).TagLimit(tagLimit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.GetTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokens`: TokenList
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.GetTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chainID** | **string** | The CAIP-2 identifier for a blockchain. | 
 **tokenAddress** | **string** | The address of an token contract to query. If omitted, all supported tokens will be queried. | 
 **tokenSymbol** | **string** | The symbol of an token contract. | 
 **tagLimit** | **[]string** | A list of token tag identifiers to include in the query. | 

### Return type

[**TokenList**](TokenList.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

