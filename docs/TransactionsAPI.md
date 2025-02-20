# \TransactionsAPI

All URIs are relative to *https://svc.blockdaemon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransactions**](TransactionsAPI.md#GetTransactions) | **Get** /transactions | Get transaction history for an account



## GetTransactions

> TransactionsResponse GetTransactions(ctx).ChainID(chainID).AccountAddress(accountAddress).Limit(limit).Page(page).Execute()

Get transaction history for an account



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
	chainID := "eip155:10" // string | The CAIP-2 identifier for a blockchain.
	accountAddress := "0xb8DC8a293fb4c6677dE0A29E505d33314eBB9161" // string | The address of an account.
	limit := int32(100) // int32 | The number of items per page. (default to 100)
	page := "1" // string | The page token/number for pagination. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransactions(context.Background()).ChainID(chainID).AccountAddress(accountAddress).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactions`: TransactionsResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chainID** | **string** | The CAIP-2 identifier for a blockchain. | 
 **accountAddress** | **string** | The address of an account. | 
 **limit** | **int32** | The number of items per page. | [default to 100]
 **page** | **string** | The page token/number for pagination. | 

### Return type

[**TransactionsResponse**](TransactionsResponse.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

